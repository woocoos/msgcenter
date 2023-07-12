// Package nflog implements a garbage-collected and snapshottable append-only log of
// active/resolved notifications. Each log entry stores the active/resolved state,
// the notified receiver, and a hash digest of the notification's identifying contents.
// The log can be queried along different parameters.
package nflog

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsingsun/members"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/metrics"
	"go.uber.org/zap"
	"sync"
	"time"
)

var logger = log.Component("nflog")

var (
	// ErrNotFound is returned for empty query results.
	ErrNotFound = errors.New("not found")
	// ErrInvalidState is returned if the state isn't valid.
	ErrInvalidState = errors.New("invalid state")
)

type NLogCallback interface {
	LoadData() ([]*Entry, error)
	CreateLog(ctx context.Context, r *Receiver, gkey string, firingAlerts, resolvedAlerts []uint64, expiresAt time.Time) (int, error)
}

type Receiver struct {
	// Configured name of the receiver.
	Name string
	// Name of the integration of the receiver.
	Integration string
	// Index of the receiver with respect to the integration.
	// Every integration in a group may have 0..N configurations.
	Index uint32
}

// Log holds the notification log state for alerts that have been notified.
type Log struct {
	Options
	NLogCallback NLogCallback
	st           state

	Spreader members.Spreader

	// For now we only store the most recently added log entry.
	// The key is a serialized concatenation of group key and receiver.
	mu sync.RWMutex
}

// MaintenanceFunc represents the function to run as part of the periodic maintenance for the nflog.
// It returns the size of the snapshot taken or an error if it failed.
type MaintenanceFunc func() error

// Options configures a new Log implementation.
type Options struct {
	Retention time.Duration

	// MaintenanceInterval alert manager garbage collects the notification log state at the given interval
	MaintenanceInterval time.Duration
	MaintenanceFunc     MaintenanceFunc
}

func (o *Options) validate() error {
	if o.MaintenanceInterval == 0 {
		return errors.New("interval or stop signal are missing - not running maintenance")
	}

	return nil
}

func NewFromConfiguration(cfg *conf.Configuration) (*Log, error) {
	silenceOpts := Options{
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
		Options: o,
		st:      state{},
	}

	if metrics.Marker == nil {
		metrics.Nflog = metrics.NewNflogMetrics()
	}

	if l.MaintenanceFunc == nil {
		l.MaintenanceFunc = func() error {
			if _, err := l.GC(); err != nil {
				return err
			}

			return nil
		}
	}

	if l.NLogCallback != nil {
		if err := l.loadData(); err != nil {
			return nil, err
		}
	}

	return l, nil
}

// loadData loads all silences data from DataLoader
func (l *Log) loadData() error {
	datas, err := l.NLogCallback.LoadData()
	if err != nil {
		return err
	}
	st := make(state)
	for _, e := range datas {
		st[e.ID] = e
	}
	l.mu.Lock()
	l.st = st
	l.mu.Unlock()

	return nil
}

func (l *Log) now() time.Time {
	return time.Now()
}

func (l *Log) Start(ctx context.Context) error {
	t := time.NewTicker(l.MaintenanceInterval)

	runMaintenance := func(do func() error) error {
		metrics.Nflog.MaintenanceTotal.Inc()
		start := l.now().UTC()
		logger.Debug("Running maintenance")
		if err := do(); err != nil {
			metrics.Nflog.MaintenanceErrorsTotal.Inc()
			return err
		}
		logger.Debug("Maintenance done", zap.Duration("duration", l.now().Sub(start)))
		return nil
	}

Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		case <-t.C:
			if err := runMaintenance(l.MaintenanceFunc); err != nil {
				logger.Error("Running maintenance failed", zap.Error(err))
			}
		}
	}

	if err := runMaintenance(l.MaintenanceFunc); err != nil {
		logger.Error("Running final maintenance failed", zap.Error(err))
	}
	return nil
}

func (l *Log) Stop(ctx context.Context) error {
	return nil
}

func receiverKey(r *Receiver) string {
	return fmt.Sprintf("%s/%s/%d", r.Name, r.Integration, r.Index)
}

// stateKey returns a string key for a log entry consisting of the group key
// and receiver.
func stateKey(k string, r *Receiver) string {
	return fmt.Sprintf("%s:%s", k, receiverKey(r))
}

func (l *Log) Log(ctx context.Context, r *Receiver, gkey string, firingAlerts, resolvedAlerts []uint64, expiry time.Duration) error {
	// Write all st with the same timestamp.
	now := l.now()
	expiresAt := now.Add(l.Retention)
	if expiry > 0 && l.Retention > expiry {
		expiresAt = now.Add(expiry)
	}
	//key := stateKey(gkey, r)
	key, err := l.NLogCallback.CreateLog(ctx, r, gkey, firingAlerts, resolvedAlerts, expiresAt)
	if err != nil {
		return err
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if prevle, ok := l.st[key]; ok {
		// Entry already exists, only overwrite if timestamp is newer.
		// This may happen with raciness or clock-drift across AM nodes.
		if prevle.UpdatedAt.After(now) {
			return nil
		}
	}

	if l.Spreader != nil {
		b, err := l.st.marshalBinary(&Entry{
			ID:             key,
			GroupKey:       gkey,
			Receiver:       r.Integration,
			ExpiresAt:      expiresAt,
			UpdatedAt:      now,
			FiringAlerts:   firingAlerts,
			ResolvedAlerts: resolvedAlerts,
		})
		if err != nil {
			return err
		}
		return l.Spreader.Broadcast(b)
	}
	return nil
}

// GC implements the Log interface.
func (l *Log) GC() (int, error) {
	start := time.Now()
	defer func() { metrics.Nflog.GCDuration.Observe(time.Since(start).Seconds()) }()

	now := l.now()
	var n int

	l.mu.Lock()
	defer l.mu.Unlock()

	for k, le := range l.st {
		if le.ExpiresAt.IsZero() {
			return n, errors.New("unexpected zero expiration timestamp")
		}
		if !le.ExpiresAt.After(now) {
			delete(l.st, k)
			n++
		}
	}

	return n, nil
}

// Query implements the Log interface.
func (l *Log) Query(params ...EntryQuery) ([]*Entry, error) {
	start := time.Now()
	metrics.Nflog.QueriesTotal.Inc()

	entries, err := func() ([]*Entry, error) {
		l.mu.RLock()
		defer l.mu.RUnlock()
		return l.st.query(params...)
	}()
	if err != nil {
		metrics.Nflog.QueryErrorsTotal.Inc()
	}
	metrics.Nflog.QueryDuration.Observe(time.Since(start).Seconds())
	return entries, err
}

func (l *Log) Name() string {
	return "nflog"
}

// MarshalBinary serializes all contents of the notification log.
func (l *Log) MarshalBinary() ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.st.MarshalBinary()
}

// Merge merges notification log state received from the cluster with the local state.
func (l *Log) Merge(b []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.st.Merge(b)
}
