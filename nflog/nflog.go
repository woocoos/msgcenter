// Package nflog implements a garbage-collected and snapshottable append-only log of
// active/resolved notifications. Each log entry stores the active/resolved state,
// the notified receiver, and a hash digest of the notification's identifying contents.
// The log can be queried along different parameters.
package nflog

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/benbjohnson/clock"
	"github.com/matttproud/golang_protobuf_extensions/pbutil"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/metrics"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	pb "github.com/woocoos/msgcenter/nflog/nflogpb"
)

var logger = log.Component("nflog")

var (
	// ErrNotFound is returned for empty query results.
	ErrNotFound = errors.New("not found")
	// ErrInvalidState is returned if the state isn't valid.
	ErrInvalidState = errors.New("invalid state")
)

// query currently allows filtering by and/or receiver group key.
// It is configured via QueryParameter functions.
//
// TODO(fabxc): Future versions could allow querying a certain receiver
// group or a given time interval.
type query struct {
	recv     *pb.Receiver
	groupKey string
}

// QueryParam is a function that modifies a query to incorporate
// a set of parameters. Returns an error for invalid or conflicting
// parameters.
type QueryParam func(*query) error

// QReceiver adds a receiver parameter to a query.
func QReceiver(r *pb.Receiver) QueryParam {
	return func(q *query) error {
		q.recv = r
		return nil
	}
}

// QGroupKey adds a group key as querying argument.
func QGroupKey(gk string) QueryParam {
	return func(q *query) error {
		q.groupKey = gk
		return nil
	}
}

// Log holds the notification log state for alerts that have been notified.
type Log struct {
	Options
	clock clock.Clock

	metrics   *metrics.NflogMetrics
	retention time.Duration

	// For now we only store the most recently added log entry.
	// The key is a serialized concatenation of group key and receiver.
	mtx       sync.RWMutex
	st        state
	broadcast func([]byte)

	stopc chan struct{}
}

// MaintenanceFunc represents the function to run as part of the periodic maintenance for the nflog.
// It returns the size of the snapshot taken or an error if it failed.
type MaintenanceFunc func() (int64, error)

type state map[string]*pb.MeshEntry

func (s state) clone() state {
	c := make(state, len(s))
	for k, v := range s {
		c[k] = v
	}
	return c
}

// merge returns true or false whether the MeshEntry was merged or
// not. This information is used to decide to gossip the message further.
func (s state) merge(e *pb.MeshEntry, now time.Time) bool {
	if e.ExpiresAt.AsTime().Before(now) {
		return false
	}
	k := stateKey(string(e.Entry.GroupKey), e.Entry.Receiver)

	prev, ok := s[k]
	if !ok || prev.Entry.Timestamp.AsTime().Before(e.Entry.Timestamp.AsTime()) {
		s[k] = e
		return true
	}
	return false
}

func (s state) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer

	for _, e := range s {
		if _, err := pbutil.WriteDelimited(&buf, e); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func decodeState(r io.Reader) (state, error) {
	st := state{}
	for {
		var e pb.MeshEntry
		_, err := pbutil.ReadDelimited(r, &e)
		if err == nil {
			if e.Entry == nil || e.Entry.Receiver == nil {
				return nil, ErrInvalidState
			}
			st[stateKey(string(e.Entry.GroupKey), e.Entry.Receiver)] = &e
			continue
		}
		if err == io.EOF {
			break
		}
		return nil, err
	}
	return st, nil
}

func marshalMeshEntry(e *pb.MeshEntry) ([]byte, error) {
	var buf bytes.Buffer
	if _, err := pbutil.WriteDelimited(&buf, e); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Options configures a new Log implementation.
type Options struct {
	SnapshotReader io.Reader
	SnapshotFile   string

	Retention time.Duration

	// MaintenanceInterval alert manager garbage collects the notification log state at the given interval
	MaintenanceInterval time.Duration
	MaintenanceFunc     MaintenanceFunc
}

func (o *Options) validate() error {
	if o.MaintenanceInterval == 0 {
		return errors.New("interval or stop signal are missing - not running maintenance")
	}
	if o.SnapshotFile != "" && o.SnapshotReader != nil {
		return errors.New("only one of SnapshotFile and SnapshotReader must be set")
	}

	return nil
}

func NewFromConfiguration(cfg *conf.Configuration) (*Log, error) {
	silenceOpts := Options{
		SnapshotFile:        filepath.Join(cfg.String("storage.path"), "snapshot"),
		Retention:           cfg.Duration("data.retention"),
		MaintenanceInterval: cfg.Duration("data.maintenanceInterval"),
	}
	return New(silenceOpts)
}

// New creates a new notification log based on the provided options.
// The snapshot is loaded into the Log if it is set.
func New(o Options) (*Log, error) {
	if err := o.validate(); err != nil {
		return nil, err
	}

	l := &Log{
		Options:   o,
		clock:     clock.New(),
		st:        state{},
		broadcast: func([]byte) {},
		metrics:   metrics.NewNflogMetrics(prometheus.DefaultRegisterer),
		stopc:     make(chan struct{}),
	}

	if l.MaintenanceFunc == nil {
		l.MaintenanceFunc = func() (int64, error) {
			var size int64
			if _, err := l.GC(); err != nil {
				return size, err
			}
			if l.SnapshotFile == "" {
				return size, nil
			}
			f, err := openReplace(l.SnapshotFile)
			if err != nil {
				return size, err
			}
			if size, err = l.Snapshot(f); err != nil {
				f.Close()
				return size, err
			}
			return size, f.Close()
		}
	}

	if o.SnapshotFile != "" {
		if r, err := os.Open(o.SnapshotFile); err != nil {
			if !os.IsNotExist(err) {
				return nil, err
			}
			logger.Debug("notification log snapshot file doesn't exist", zap.Error(err))
		} else {
			o.SnapshotReader = r
			defer r.Close()
		}
	}

	if o.SnapshotReader != nil {
		if err := l.loadSnapshot(o.SnapshotReader); err != nil {
			return l, err
		}
	}

	return l, nil
}

func (l *Log) now() time.Time {
	return l.clock.Now()
}

func (l *Log) Start(ctx context.Context) error {
	t := l.clock.Ticker(l.MaintenanceInterval)

	runMaintenance := func(do func() (int64, error)) error {
		l.metrics.MaintenanceTotal.Inc()
		start := l.now().UTC()
		logger.Debug("Running maintenance")
		size, err := do()
		l.metrics.SnapshotSize.Set(float64(size))
		if err != nil {
			l.metrics.MaintenanceErrorsTotal.Inc()
			return err
		}
		logger.Debug("Maintenance done", zap.Duration("duration", l.now().Sub(start)), zap.Int64("size", size))
		return nil
	}

Loop:
	for {
		select {
		case <-l.stopc:
			break Loop
		case <-t.C:
			if err := runMaintenance(l.MaintenanceFunc); err != nil {
				logger.Error("Running maintenance failed", zap.Error(err))
			}
		}
	}

	// No need to run final maintenance if we don't want to snapshot.
	if l.SnapshotFile == "" {
		return nil
	}
	if err := runMaintenance(l.MaintenanceFunc); err != nil {
		logger.Error("Running final maintenance failed", zap.Error(err))
	}
	return nil
}

func (l *Log) Stop(ctx context.Context) error {
	close(l.stopc)
	return nil
}

func receiverKey(r *pb.Receiver) string {
	return fmt.Sprintf("%s/%s/%d", r.GroupName, r.Integration, r.Idx)
}

// stateKey returns a string key for a log entry consisting of the group key
// and receiver.
func stateKey(k string, r *pb.Receiver) string {
	return fmt.Sprintf("%s:%s", k, receiverKey(r))
}

func (l *Log) Log(r *pb.Receiver, gkey string, firingAlerts, resolvedAlerts []uint64, expiry time.Duration) error {
	// Write all st with the same timestamp.
	now := l.now()
	key := stateKey(gkey, r)

	l.mtx.Lock()
	defer l.mtx.Unlock()

	if prevle, ok := l.st[key]; ok {
		// Entry already exists, only overwrite if timestamp is newer.
		// This may happen with raciness or clock-drift across AM nodes.
		if prevle.Entry.Timestamp.AsTime().After(now) {
			return nil
		}
	}

	expiresAt := now.Add(l.Retention)
	if expiry > 0 && l.Retention > expiry {
		expiresAt = now.Add(expiry)
	}

	e := &pb.MeshEntry{
		Entry: &pb.Entry{
			Receiver:       r,
			GroupKey:       []byte(gkey),
			Timestamp:      timestamppb.New(now),
			FiringAlerts:   firingAlerts,
			ResolvedAlerts: resolvedAlerts,
		},
		ExpiresAt: timestamppb.New(expiresAt),
	}

	b, err := marshalMeshEntry(e)
	if err != nil {
		return err
	}
	l.st.merge(e, l.now())
	l.broadcast(b)

	return nil
}

// GC implements the Log interface.
func (l *Log) GC() (int, error) {
	start := time.Now()
	defer func() { l.metrics.GCDuration.Observe(time.Since(start).Seconds()) }()

	now := l.now()
	var n int

	l.mtx.Lock()
	defer l.mtx.Unlock()

	for k, le := range l.st {
		if le.ExpiresAt.AsTime().IsZero() {
			return n, errors.New("unexpected zero expiration timestamp")
		}
		if !le.ExpiresAt.AsTime().After(now) {
			delete(l.st, k)
			n++
		}
	}

	return n, nil
}

// Query implements the Log interface.
func (l *Log) Query(params ...QueryParam) ([]*pb.Entry, error) {
	start := time.Now()
	l.metrics.QueriesTotal.Inc()

	entries, err := func() ([]*pb.Entry, error) {
		q := &query{}
		for _, p := range params {
			if err := p(q); err != nil {
				return nil, err
			}
		}
		// TODO(fabxc): For now our only query mode is the most recent entry for a
		// receiver/group_key combination.
		if q.recv == nil || q.groupKey == "" {
			// TODO(fabxc): allow more complex queries in the future.
			// How to enable pagination?
			return nil, errors.New("no query parameters specified")
		}

		l.mtx.RLock()
		defer l.mtx.RUnlock()

		if le, ok := l.st[stateKey(q.groupKey, q.recv)]; ok {
			return []*pb.Entry{le.Entry}, nil
		}
		return nil, ErrNotFound
	}()
	if err != nil {
		l.metrics.QueryErrorsTotal.Inc()
	}
	l.metrics.QueryDuration.Observe(time.Since(start).Seconds())
	return entries, err
}

// loadSnapshot loads a snapshot generated by Snapshot() into the state.
func (l *Log) loadSnapshot(r io.Reader) error {
	st, err := decodeState(r)
	if err != nil {
		return err
	}

	l.mtx.Lock()
	l.st = st
	l.mtx.Unlock()

	return nil
}

// Snapshot implements the Log interface.
func (l *Log) Snapshot(w io.Writer) (int64, error) {
	start := time.Now()
	defer func() { l.metrics.SnapshotDuration.Observe(time.Since(start).Seconds()) }()

	l.mtx.RLock()
	defer l.mtx.RUnlock()

	b, err := l.st.MarshalBinary()
	if err != nil {
		return 0, err
	}

	return io.Copy(w, bytes.NewReader(b))
}

// MarshalBinary serializes all contents of the notification log.
func (l *Log) MarshalBinary() ([]byte, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	return l.st.MarshalBinary()
}

// Merge merges notification log state received from the cluster with the local state.
func (l *Log) Merge(b []byte) error {
	st, err := decodeState(bytes.NewReader(b))
	if err != nil {
		return err
	}
	l.mtx.Lock()
	defer l.mtx.Unlock()
	now := l.now()

	for _, e := range st {
		if merged := l.st.merge(e, now); merged {
			// If this is the first we've seen the message and it's
			// not oversized, gossip it to other nodes. We don't
			// propagate oversized messages because they're sent to
			// all nodes already.
			l.broadcast(b)
			l.metrics.PropagatedMessagesTotal.Inc()
			logger.Debug("gossiping new entry", zap.String("entry", e.String()))
		}
	}
	return nil
}

// SetBroadcast sets a broadcast callback that will be invoked with serialized state
// on updates.
func (l *Log) SetBroadcast(f func([]byte)) {
	l.mtx.Lock()
	l.broadcast = f
	l.mtx.Unlock()
}

// replaceFile wraps a file that is moved to another filename on closing.
type replaceFile struct {
	*os.File
	filename string
}

func (f *replaceFile) Close() error {
	if err := f.File.Sync(); err != nil {
		return err
	}
	if err := f.File.Close(); err != nil {
		return err
	}
	return os.Rename(f.File.Name(), f.filename)
}

// openReplace opens a new temporary file that is moved to filename on closing.
func openReplace(filename string) (*replaceFile, error) {
	tmpFilename := fmt.Sprintf("%s.%x", filename, uint64(rand.Int63()))

	f, err := os.Create(tmpFilename)
	if err != nil {
		return nil, err
	}

	rf := &replaceFile{
		File:     f,
		filename: filename,
	}
	return rf, nil
}
