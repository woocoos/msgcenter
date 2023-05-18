package silence

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/benbjohnson/clock"
	"github.com/google/uuid"
	"github.com/matttproud/golang_protobuf_extensions/pbutil"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"sync"
	"time"
	"unicode/utf8"

	pb "github.com/woocoos/msgcenter/silence/silencepb"
)

var logger = log.Component("silence")

var (
	// ErrNotFound is returned if a silence was not found.
	ErrNotFound = fmt.Errorf("silence not found")
	// ErrInvalidState is returned if the state isn't valid.
	ErrInvalidState = fmt.Errorf("invalid state")
)

type matcherCache map[*pb.Silence]label.Matchers

// Get retrieves the matchers for a given silence. If it is a missed cache
// access, it compiles and adds the matchers of the requested silence to the
// cache.
func (c matcherCache) Get(s *pb.Silence) (label.Matchers, error) {
	if m, ok := c[s]; ok {
		return m, nil
	}
	return c.add(s)
}

// add compiles a silences' matchers and adds them to the cache.
// It returns the compiled matchers.
func (c matcherCache) add(s *pb.Silence) (label.Matchers, error) {
	ms := make(label.Matchers, len(s.Matchers))

	for i, m := range s.Matchers {
		var mt label.MatchType
		switch m.Type {
		case pb.Matcher_EQUAL:
			mt = label.MatchEqual
		case pb.Matcher_NOT_EQUAL:
			mt = label.MatchNotEqual
		case pb.Matcher_REGEXP:
			mt = label.MatchRegexp
		case pb.Matcher_NOT_REGEXP:
			mt = label.MatchNotRegexp
		default:
			return nil, fmt.Errorf("unknown matcher type %q", m.Type)
		}
		matcher, err := label.NewMatcher(mt, m.Name, m.Pattern)
		if err != nil {
			return nil, err
		}

		ms[i] = matcher
	}

	c[s] = ms
	return ms, nil
}

// Silencer binds together a Marker and a Silences to implement the Muter
// interface.
type Silencer struct {
	silences *Silences
	marker   alert.Marker
}

// NewSilencer returns a new Silencer.
func NewSilencer(s *Silences, m alert.Marker) *Silencer {
	return &Silencer{
		silences: s,
		marker:   m,
	}
}

// Mutes implements the Muter interface.
func (s *Silencer) Mutes(lset label.LabelSet) bool {
	fp := lset.Fingerprint()
	activeIDs, pendingIDs, markerVersion, _ := s.marker.Silenced(fp)

	var (
		err        error
		allSils    []*pb.Silence
		newVersion = markerVersion
	)
	if markerVersion == s.silences.Version() {
		totalSilences := len(activeIDs) + len(pendingIDs)
		// No new silences added, just need to check which of the old
		// silences are still relevant and which of the pending ones
		// have become active.
		if totalSilences == 0 {
			// Super fast path: No silences ever applied to this
			// alert, none have been added. We are done.
			return false
		}
		// This is still a quite fast path: No silences have been added,
		// we only need to check which of the applicable silences are
		// currently active. Note that newVersion is left at
		// markerVersion because the Query call might already return a
		// newer version, which is not the version our old list of
		// applicable silences is based on.
		allIDs := append(append(make([]string, 0, totalSilences), activeIDs...), pendingIDs...)
		allSils, _, err = s.silences.Query(
			QIDs(allIDs...),
			QState(alert.SilenceStateActive, alert.SilenceStatePending),
		)
	} else {
		// New silences have been added, do a full query.
		allSils, newVersion, err = s.silences.Query(
			QState(alert.SilenceStateActive, alert.SilenceStatePending),
			QMatches(lset),
		)
	}
	if err != nil {
		logger.Error("Querying silences failed, alerts might not get silenced correctly", zap.Error(err))
	}
	if len(allSils) == 0 {
		// Easy case, neither active nor pending silences anymore.
		s.marker.SetActiveOrSilenced(fp, newVersion, nil, nil)
		return false
	}
	// It is still possible that nothing has changed, but finding out is not
	// much less effort than just recreating the IDs from the query
	// result. So let's do it in any case. Note that we cannot reuse the
	// current ID slices for concurrency reasons.
	activeIDs, pendingIDs = nil, nil
	now := s.silences.nowUTC()
	for _, sil := range allSils {
		switch getState(sil, now) {
		case alert.SilenceStatePending:
			pendingIDs = append(pendingIDs, sil.Id)
		case alert.SilenceStateActive:
			activeIDs = append(activeIDs, sil.Id)
		default:
			// Do nothing, silence has expired in the meantime.
		}
	}
	logger.Debug("determined current silences state",
		zap.Int("total", len(allSils)),
		zap.Time("now", now),
		zap.Int("active", len(activeIDs)),
		zap.Int("pending", len(pendingIDs)),
	)

	sort.Strings(activeIDs)
	sort.Strings(pendingIDs)

	s.marker.SetActiveOrSilenced(fp, newVersion, activeIDs, pendingIDs)

	return len(activeIDs) > 0
}

// Silences holds a silence state that can be modified, queried, and snapshot.
type Silences struct {
	Options

	clock clock.Clock

	metrics *metrics

	mtx       sync.RWMutex
	st        state
	version   int // Increments whenever silences are added.
	broadcast func([]byte)
	mc        matcherCache

	stopc chan struct{}
}

// MaintenanceFunc represents the function to run as part of the periodic maintenance for silences.
// It returns the size of the snapshot taken or an error if it failed.
type MaintenanceFunc func() (int64, error)

type metrics struct {
	gcDuration              prometheus.Summary
	snapshotDuration        prometheus.Summary
	snapshotSize            prometheus.Gauge
	queriesTotal            prometheus.Counter
	queryErrorsTotal        prometheus.Counter
	queryDuration           prometheus.Histogram
	silencesActive          prometheus.GaugeFunc
	silencesPending         prometheus.GaugeFunc
	silencesExpired         prometheus.GaugeFunc
	propagatedMessagesTotal prometheus.Counter
	maintenanceTotal        prometheus.Counter
	maintenanceErrorsTotal  prometheus.Counter
}

func newSilenceMetricByState(s *Silences, st alert.SilenceState) prometheus.GaugeFunc {
	return prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name:        "alertmanager_silences",
			Help:        "How many silences by state.",
			ConstLabels: prometheus.Labels{"state": string(st)},
		},
		func() float64 {
			count, err := s.CountState(st)
			if err != nil {
				logger.Error("Counting silences failed", zap.Error(err))
			}
			return float64(count)
		},
	)
}

func newMetrics(r prometheus.Registerer, s *Silences) *metrics {
	m := &metrics{}

	m.gcDuration = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "alertmanager_silences_gc_duration_seconds",
		Help:       "Duration of the last silence garbage collection cycle.",
		Objectives: map[float64]float64{},
	})
	m.snapshotDuration = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "alertmanager_silences_snapshot_duration_seconds",
		Help:       "Duration of the last silence snapshot.",
		Objectives: map[float64]float64{},
	})
	m.snapshotSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "alertmanager_silences_snapshot_size_bytes",
		Help: "Size of the last silence snapshot in bytes.",
	})
	m.maintenanceTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "alertmanager_silences_maintenance_total",
		Help: "How many maintenances were executed for silences.",
	})
	m.maintenanceErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "alertmanager_silences_maintenance_errors_total",
		Help: "How many maintenances were executed for silences that failed.",
	})
	m.queriesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "alertmanager_silences_queries_total",
		Help: "How many silence queries were received.",
	})
	m.queryErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "alertmanager_silences_query_errors_total",
		Help: "How many silence received queries did not succeed.",
	})
	m.queryDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "alertmanager_silences_query_duration_seconds",
		Help: "Duration of silence query evaluation.",
	})
	m.propagatedMessagesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "alertmanager_silences_gossip_messages_propagated_total",
		Help: "Number of received gossip messages that have been further gossiped.",
	})
	if s != nil {
		m.silencesActive = newSilenceMetricByState(s, alert.SilenceStateActive)
		m.silencesPending = newSilenceMetricByState(s, alert.SilenceStatePending)
		m.silencesExpired = newSilenceMetricByState(s, alert.SilenceStateExpired)
	}

	if r != nil {
		r.MustRegister(
			m.gcDuration,
			m.snapshotDuration,
			m.snapshotSize,
			m.queriesTotal,
			m.queryErrorsTotal,
			m.queryDuration,
			m.silencesActive,
			m.silencesPending,
			m.silencesExpired,
			m.propagatedMessagesTotal,
			m.maintenanceTotal,
			m.maintenanceErrorsTotal,
		)
	}
	return m
}

// Options exposes configuration options for creating a new Silences object.
// Its zero value is a safe default.
type Options struct {
	// A snapshot file or reader from which the initial state is loaded.
	// None or only one of them must be set.
	SnapshotFile   string
	SnapshotReader io.Reader

	// Retention time for newly created Silences. Silences may be
	// garbage collected after the given duration after they ended.
	Retention time.Duration

	// A logger used by background processing.
	Logger  log.ComponentLogger
	Metrics prometheus.Registerer

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

func NewFromConfiguration(cfg *conf.Configuration) (*Silences, error) {
	silenceOpts := Options{
		SnapshotFile:        filepath.Join(cfg.String("storage.path"), "silences"),
		Retention:           cfg.Duration("data.retention"),
		Logger:              logger,
		Metrics:             prometheus.DefaultRegisterer,
		MaintenanceInterval: cfg.Duration("data.maintenanceInterval"),
	}
	return New(silenceOpts)
}

// New returns a new Silences object with the given configuration.
func New(o Options) (*Silences, error) {
	if err := o.validate(); err != nil {
		return nil, err
	}
	s := &Silences{
		Options:   o,
		clock:     clock.New(),
		mc:        matcherCache{},
		broadcast: func([]byte) {},
		st:        state{},
		stopc:     make(chan struct{}),
	}
	s.metrics = newMetrics(o.Metrics, s)

	if o.Logger == nil {
		s.Logger = log.Global()
	}

	if s.MaintenanceFunc == nil {
		s.MaintenanceFunc = func() (int64, error) {
			var size int64

			if _, err := s.GC(); err != nil {
				return size, err
			}
			if s.SnapshotFile == "" {
				return size, nil
			}
			f, err := openReplace(s.SnapshotFile)
			if err != nil {
				return size, err
			}
			if size, err = s.Snapshot(f); err != nil {
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
			logger.Debug("silences snapshot file doesn't exist", zap.Error(err))
		} else {
			o.SnapshotReader = r
			defer r.Close()
		}
	}

	if o.SnapshotReader != nil {
		if err := s.loadSnapshot(o.SnapshotReader); err != nil {
			return s, err
		}
	}
	return s, nil
}

func (s *Silences) nowUTC() time.Time {
	return s.clock.Now().UTC()
}

func (s *Silences) Start(ctx context.Context) error {
	t := s.clock.Ticker(s.MaintenanceInterval)
	defer t.Stop()

	runMaintenance := func(do MaintenanceFunc) error {
		s.metrics.maintenanceTotal.Inc()
		logger.Debug("Running maintenance")
		start := s.nowUTC()
		size, err := do()
		s.metrics.snapshotSize.Set(float64(size))
		if err != nil {
			s.metrics.maintenanceErrorsTotal.Inc()
			return err
		}
		logger.Debug("Maintenance done", zap.Duration("duration", s.clock.Since(start)), zap.Int64("size", size))
		return nil
	}

Loop:
	for {
		select {
		case <-s.stopc:
			break Loop
		case <-t.C:
			if err := runMaintenance(s.MaintenanceFunc); err != nil {
				logger.Info("Running maintenance failed", zap.Error(err))
			}
		}
	}

	// No need for final maintenance if we don't want to snapshot.
	if s.SnapshotFile == "" {
		return nil
	}
	if err := runMaintenance(s.MaintenanceFunc); err != nil {
		logger.Info("Running final maintenance failed", zap.Error(err))
	}
	return nil
}

func (s *Silences) Stop(ctx context.Context) error {
	close(s.stopc)
	return nil
}

// GC runs a garbage collection that removes silences that have ended longer
// than the configured retention time ago.
func (s *Silences) GC() (int, error) {
	start := time.Now()
	defer func() { s.metrics.gcDuration.Observe(time.Since(start).Seconds()) }()

	now := s.nowUTC()
	var n int

	s.mtx.Lock()
	defer s.mtx.Unlock()

	for id, sil := range s.st {
		if sil.ExpiresAt.AsTime().IsZero() {
			return n, errors.New("unexpected zero expiration timestamp")
		}
		if !sil.ExpiresAt.AsTime().After(now) {
			delete(s.st, id)
			delete(s.mc, sil.Silence)
			n++
		}
	}

	return n, nil
}

// ValidateMatcher runs validation on the matcher name, type, and pattern.
var ValidateMatcher = func(m *pb.Matcher) error {
	if !label.LabelName(m.Name).IsValid() {
		return fmt.Errorf("invalid label name %q", m.Name)
	}
	switch m.Type {
	case pb.Matcher_EQUAL, pb.Matcher_NOT_EQUAL:
		if !utf8.ValidString(m.Pattern) {
			return fmt.Errorf("invalid label value %q", m.Pattern)
		}
	case pb.Matcher_REGEXP, pb.Matcher_NOT_REGEXP:
		if _, err := regexp.Compile(m.Pattern); err != nil {
			return fmt.Errorf("invalid regular expression %q: %s", m.Pattern, err)
		}
	default:
		return fmt.Errorf("unknown matcher type %q", m.Type)
	}
	return nil
}

func matchesEmpty(m *pb.Matcher) bool {
	switch m.Type {
	case pb.Matcher_EQUAL:
		return m.Pattern == ""
	case pb.Matcher_REGEXP:
		matched, _ := regexp.MatchString(m.Pattern, "")
		return matched
	default:
		return false
	}
}

func validateSilence(s *pb.Silence) error {
	if s.Id == "" {
		return errors.New("ID missing")
	}
	if len(s.Matchers) == 0 {
		return errors.New("at least one matcher required")
	}
	allMatchEmpty := true
	for i, m := range s.Matchers {
		if err := ValidateMatcher(m); err != nil {
			return fmt.Errorf("invalid label matcher %d: %s", i, err)
		}
		allMatchEmpty = allMatchEmpty && matchesEmpty(m)
	}
	if allMatchEmpty {
		return errors.New("at least one matcher must not match the empty string")
	}
	if s.StartsAt.AsTime().IsZero() {
		return errors.New("invalid zero start timestamp")
	}
	if s.EndsAt.AsTime().IsZero() {
		return errors.New("invalid zero end timestamp")
	}
	if s.EndsAt.AsTime().Before(s.StartsAt.AsTime()) {
		return errors.New("end time must not be before start time")
	}
	if s.UpdatedAt.AsTime().IsZero() {
		return errors.New("invalid zero update timestamp")
	}
	return nil
}

// cloneSilence returns a shallow copy of a silence.
func cloneSilence(sil *pb.Silence) *pb.Silence {
	s := *sil
	return &s
}

func (s *Silences) getSilence(id string) (*pb.Silence, bool) {
	msil, ok := s.st[id]
	if !ok {
		return nil, false
	}
	return msil.Silence, true
}

func (s *Silences) setSilence(sil *pb.Silence, now time.Time) error {
	sil.UpdatedAt = timestamppb.New(now)

	if err := validateSilence(sil); err != nil {
		return fmt.Errorf("silence invalid %w", err)
	}

	msil := &pb.MeshSilence{
		Silence:   sil,
		ExpiresAt: timestamppb.New(sil.EndsAt.AsTime().Add(s.Retention)),
	}
	b, err := marshalMeshSilence(msil)
	if err != nil {
		return err
	}

	if s.st.merge(msil, now) {
		s.version++
	}
	s.broadcast(b)

	return nil
}

// Set the specified silence. If a silence with the ID already exists and the modification
// modifies history, the old silence gets expired and a new one is created.
func (s *Silences) Set(sil *pb.Silence) (string, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	now := s.nowUTC()
	prev, ok := s.getSilence(sil.Id)

	if sil.Id != "" && !ok {
		return "", ErrNotFound
	}
	if ok {
		if canUpdate(prev, sil, now) {
			return sil.Id, s.setSilence(sil, now)
		}
		if getState(prev, s.nowUTC()) != alert.SilenceStateExpired {
			// We cannot update the silence, expire the old one.
			if err := s.expire(prev.Id); err != nil {
				return "", fmt.Errorf("expire previous silence %w", err)
			}
		}
	}
	// If we got here it's either a new silence or a replacing one.
	uid := uuid.New()

	sil.Id = uid.String()

	if sil.StartsAt.AsTime().Before(now) {
		sil.StartsAt = timestamppb.New(now)
	}

	return sil.Id, s.setSilence(sil, now)
}

// canUpdate returns true if silence a can be updated to b without
// affecting the historic view of silencing.
func canUpdate(a, b *pb.Silence, now time.Time) bool {
	if !reflect.DeepEqual(a.Matchers, b.Matchers) {
		return false
	}
	// Allowed timestamp modifications depend on the current time.
	switch st := getState(a, now); st {
	case alert.SilenceStateActive:
		if b.StartsAt.AsTime().Unix() != a.StartsAt.AsTime().Unix() {
			return false
		}
		if b.EndsAt.AsTime().Before(now) {
			return false
		}
	case alert.SilenceStatePending:
		if b.StartsAt.AsTime().Before(now) {
			return false
		}
	case alert.SilenceStateExpired:
		return false
	default:
		panic("unknown silence state")
	}
	return true
}

// Expire the silence with the given ID immediately.
func (s *Silences) Expire(id string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.expire(id)
}

// Expire the silence with the given ID immediately.
// It is idempotent, nil is returned if the silence already expired before it is GC'd.
// If the silence is not found an error is returned.
func (s *Silences) expire(id string) error {
	sil, ok := s.getSilence(id)
	if !ok {
		return ErrNotFound
	}
	sil = cloneSilence(sil)
	now := s.nowUTC()

	switch getState(sil, now) {
	case alert.SilenceStateExpired:
		return nil
	case alert.SilenceStateActive:
		sil.EndsAt = timestamppb.New(now)
	case alert.SilenceStatePending:
		// Set both to now to make Silence move to "expired" state
		sil.StartsAt = timestamppb.New(now)
		sil.EndsAt = timestamppb.New(now)
	}

	return s.setSilence(sil, now)
}

// QueryParam expresses parameters along which silences are queried.
type QueryParam func(*query) error

type query struct {
	ids     []string
	filters []silenceFilter
}

// silenceFilter is a function that returns true if a silence
// should be dropped from a result set for a given time.
type silenceFilter func(*pb.Silence, *Silences, time.Time) (bool, error)

// QIDs configures a query to select the given silence IDs.
func QIDs(ids ...string) QueryParam {
	return func(q *query) error {
		q.ids = append(q.ids, ids...)
		return nil
	}
}

// QMatches returns silences that match the given label set.
func QMatches(set label.LabelSet) QueryParam {
	return func(q *query) error {
		f := func(sil *pb.Silence, s *Silences, _ time.Time) (bool, error) {
			m, err := s.mc.Get(sil)
			if err != nil {
				return true, err
			}
			return m.Matches(set), nil
		}
		q.filters = append(q.filters, f)
		return nil
	}
}

// getState returns a silence's SilenceState at the given timestamp.
func getState(sil *pb.Silence, ts time.Time) alert.SilenceState {
	if ts.Before(sil.StartsAt.AsTime()) {
		return alert.SilenceStatePending
	}
	if ts.After(sil.EndsAt.AsTime()) {
		return alert.SilenceStateExpired
	}
	return alert.SilenceStateActive
}

// QState filters queried silences by the given states.
func QState(states ...alert.SilenceState) QueryParam {
	return func(q *query) error {
		f := func(sil *pb.Silence, _ *Silences, now time.Time) (bool, error) {
			s := getState(sil, now)

			for _, ps := range states {
				if s == ps {
					return true, nil
				}
			}
			return false, nil
		}
		q.filters = append(q.filters, f)
		return nil
	}
}

// QueryOne queries with the given parameters and returns the first result.
// Returns ErrNotFound if the query result is empty.
func (s *Silences) QueryOne(params ...QueryParam) (*pb.Silence, error) {
	res, _, err := s.Query(params...)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, ErrNotFound
	}
	return res[0], nil
}

// Query for silences based on the given query parameters. It returns the
// resulting silences and the state version the result is based on.
func (s *Silences) Query(params ...QueryParam) ([]*pb.Silence, int, error) {
	s.metrics.queriesTotal.Inc()
	defer prometheus.NewTimer(s.metrics.queryDuration).ObserveDuration()

	q := &query{}
	for _, p := range params {
		if err := p(q); err != nil {
			s.metrics.queryErrorsTotal.Inc()
			return nil, s.Version(), err
		}
	}
	sils, version, err := s.query(q, s.nowUTC())
	if err != nil {
		s.metrics.queryErrorsTotal.Inc()
	}
	return sils, version, err
}

// Version of the silence state.
func (s *Silences) Version() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.version
}

// CountState counts silences by state.
func (s *Silences) CountState(states ...alert.SilenceState) (int, error) {
	// This could probably be optimized.
	sils, _, err := s.Query(QState(states...))
	if err != nil {
		return -1, err
	}
	return len(sils), nil
}

func (s *Silences) query(q *query, now time.Time) ([]*pb.Silence, int, error) {
	// If we have no ID constraint, all silences are our base set.  This and
	// the use of post-filter functions is the trivial solution for now.
	var res []*pb.Silence

	s.mtx.Lock()
	defer s.mtx.Unlock()

	if q.ids != nil {
		for _, id := range q.ids {
			if s, ok := s.st[id]; ok {
				res = append(res, s.Silence)
			}
		}
	} else {
		for _, sil := range s.st {
			res = append(res, sil.Silence)
		}
	}

	var resf []*pb.Silence
	for _, sil := range res {
		remove := false
		for _, f := range q.filters {
			ok, err := f(sil, s, now)
			if err != nil {
				return nil, s.version, err
			}
			if !ok {
				remove = true
				break
			}
		}
		if !remove {
			resf = append(resf, cloneSilence(sil))
		}
	}

	return resf, s.version, nil
}

// loadSnapshot loads a snapshot generated by Snapshot() into the state.
// Any previous state is wiped.
func (s *Silences) loadSnapshot(r io.Reader) error {
	st, err := decodeState(r)
	if err != nil {
		return err
	}
	for _, e := range st {
		st[e.Silence.Id] = e
	}
	s.mtx.Lock()
	s.st = st
	s.version++
	s.mtx.Unlock()

	return nil
}

// Snapshot writes the full internal state into the writer and returns the number of bytes
// written.
func (s *Silences) Snapshot(w io.Writer) (int64, error) {
	start := time.Now()
	defer func() { s.metrics.snapshotDuration.Observe(time.Since(start).Seconds()) }()

	s.mtx.RLock()
	defer s.mtx.RUnlock()

	b, err := s.st.MarshalBinary()
	if err != nil {
		return 0, err
	}

	return io.Copy(w, bytes.NewReader(b))
}

// MarshalBinary serializes all silences.
func (s *Silences) MarshalBinary() ([]byte, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.st.MarshalBinary()
}

// Merge merges silence state received from the cluster with the local state.
func (s *Silences) Merge(b []byte) error {
	st, err := decodeState(bytes.NewReader(b))
	if err != nil {
		return err
	}
	s.mtx.Lock()
	defer s.mtx.Unlock()

	now := s.nowUTC()

	for _, e := range st {
		if merged := s.st.merge(e, now); merged {
			s.version++
		}
	}
	return nil
}

// SetBroadcast sets the provided function as the one creating data to be
// broadcast.
func (s *Silences) SetBroadcast(f func([]byte)) {
	s.mtx.Lock()
	s.broadcast = f
	s.mtx.Unlock()
}

type state map[string]*pb.MeshSilence

func (s state) merge(e *pb.MeshSilence, now time.Time) bool {
	id := e.Silence.Id
	if e.ExpiresAt.AsTime().Before(now) {
		return false
	}

	prev, ok := s[id]
	if !ok || prev.Silence.UpdatedAt.AsTime().Before(e.Silence.UpdatedAt.AsTime()) {
		s[id] = e
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
		var s pb.MeshSilence
		_, err := pbutil.ReadDelimited(r, &s)
		if err == nil {
			if s.Silence == nil {
				return nil, ErrInvalidState
			}
			st[s.Silence.Id] = &s
			continue
		}
		if err == io.EOF {
			break
		}
		return nil, err
	}
	return st, nil
}

func marshalMeshSilence(e *pb.MeshSilence) ([]byte, error) {
	var buf bytes.Buffer
	if _, err := pbutil.WriteDelimited(&buf, e); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
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
