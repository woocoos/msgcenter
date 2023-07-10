package silence

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tsingsun/members"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/vmihailenco/msgpack/v4"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/metrics"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"go.uber.org/zap"
	"reflect"
	"regexp"
	"sort"
	"sync"
	"time"
	"unicode/utf8"
)

var (
	logger               = log.Component("silence")
	_      members.Shard = (*Silences)(nil)
)

var (
	// ErrNotFound is returned if a silence was not found.
	ErrNotFound = fmt.Errorf("silence not found")
	// ErrInvalidState is returned if the state isn't valid.
	ErrInvalidState = fmt.Errorf("invalid state")
)

type matcherCache map[*ent.Silence]label.Matchers

// Get retrieves the matchers for a given silence. If it is a missed cache
// access, it compiles and adds the matchers of the requested silence to the
// cache.
func (c matcherCache) Get(s *ent.Silence) (label.Matchers, error) {
	if m, ok := c[s]; ok {
		return m, nil
	}
	return c.add(s)
}

// add compiles a silences' matchers and adds them to the cache.
// It returns the compiled matchers.
func (c matcherCache) add(s *ent.Silence) (label.Matchers, error) {
	ms := make(label.Matchers, len(s.Matchers))

	for i, m := range s.Matchers {
		matcher, err := label.NewMatcher(m.Type, m.Name, m.Value)
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
		allSils    []*ent.Silence
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
		allIDs := append(append(make([]int, 0, totalSilences), activeIDs...), pendingIDs...)
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
			pendingIDs = append(pendingIDs, sil.ID)
		case alert.SilenceStateActive:
			activeIDs = append(activeIDs, sil.ID)
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

	sort.Ints(activeIDs)
	sort.Ints(pendingIDs)

	s.marker.SetActiveOrSilenced(fp, newVersion, activeIDs, pendingIDs)

	return len(activeIDs) > 0
}

// Silences holds a silence state that can be modified, queried, and snapshot.
type Silences struct {
	Options

	mtx     sync.RWMutex
	st      state
	version int // Increments whenever silences are added.
	mc      matcherCache
}

// MaintenanceFunc represents the function to run as part of the periodic maintenance for silences.
type MaintenanceFunc func() error

type Option func(*Options)

// WithDataLoader sets the data loader function for the silences.
// if not set, the silences will use the memory to store silences.
// in distributed mode, this should be set to use the sync data.
func WithDataLoader(fn func(ids ...int) ([]*ent.Silence, error)) Option {
	return func(o *Options) {
		o.DataLoader = fn
	}
}

// Options exposes configuration options for creating a new Silences object.
// Its zero value is a safe default.
type Options struct {
	// Retention time for newly created Silences. Silences may be
	// garbage collected after the given duration after they ended.
	Retention time.Duration

	MaintenanceInterval time.Duration
	MaintenanceFunc     MaintenanceFunc

	DataLoader func(ids ...int) ([]*ent.Silence, error)
	Spreader   members.Spreader
}

func (o *Options) validate() error {
	if o.MaintenanceInterval == 0 {
		return errors.New("interval or stop signal are missing - not running maintenance")
	}

	return nil
}

// NewFromConfiguration returns a new Silences object with the given configuration.
func NewFromConfiguration(cfg *conf.Configuration, opts ...Option) (*Silences, error) {
	silenceOpts := Options{
		Retention:           cfg.Duration("data.retention"),
		MaintenanceInterval: cfg.Duration("data.maintenanceInterval"),
	}
	for _, opt := range opts {
		opt(&silenceOpts)
	}
	return New(silenceOpts)
}

// New returns a new Silences object with the given configuration.
func New(o Options) (*Silences, error) {
	if err := o.validate(); err != nil {
		return nil, err
	}
	s := &Silences{
		Options: o,
		mc:      matcherCache{},
		st:      state{},
	}

	if metrics.Silences == nil {
		metrics.Silences = metrics.NewSilencesMetrics(func(state string) float64 {
			count, err := s.CountState(alert.SilenceState(state))
			if err != nil {
				logger.Error("Counting silences failed", zap.Error(err))
			}
			return float64(count)
		})
	}

	if s.MaintenanceFunc == nil {
		s.MaintenanceFunc = func() error {
			_, err := s.GC()
			if err != nil {
				return err
			}
			return nil
		}
	}
	if s.DataLoader != nil {
		if err := s.loadData(); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s *Silences) Name() string {
	return "silences"
}

func (s *Silences) MarshalBinary() ([]byte, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.st.MarshalBinary()
}

func (s *Silences) Merge(b []byte) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.st.Merge(b)
}

func (s *Silences) nowUTC() time.Time {
	return time.Now().UTC()
}

func (s *Silences) Start(ctx context.Context) error {
	t := time.NewTicker(s.MaintenanceInterval)
	defer t.Stop()

	runMaintenance := func(do MaintenanceFunc) error {
		metrics.Silences.MaintenanceTotal.Inc()
		logger.Debug("Running maintenance")
		start := s.nowUTC()
		err := do()
		if err != nil {
			metrics.Silences.MaintenanceErrorsTotal.Inc()
			return err
		}
		logger.Debug("Maintenance done", zap.Duration("duration", time.Since(start)))
		return nil
	}

Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		case <-t.C:
			if err := runMaintenance(s.MaintenanceFunc); err != nil {
				logger.Info("Running maintenance failed", zap.Error(err))
			}
		}
	}

	if err := runMaintenance(s.MaintenanceFunc); err != nil {
		logger.Info("Running final maintenance failed", zap.Error(err))
	}
	return nil
}

func (s *Silences) Stop(ctx context.Context) error {
	return nil
}

// GC runs a garbage collection that removes silences that have ended longer
// than the configured retention time ago.
func (s *Silences) GC() (int, error) {
	start := time.Now()
	defer func() { metrics.Silences.GcDuration.Observe(time.Since(start).Seconds()) }()

	now := s.nowUTC()
	var n int

	s.mtx.Lock()
	defer s.mtx.Unlock()

	for id, sil := range s.st {
		if sil.EndsAt.IsZero() {
			return n, errors.New("unexpected zero expiration timestamp")
		}
		if !sil.EndsAt.After(now) {
			delete(s.st, id)
			delete(s.mc, sil)
			n++
		}
	}

	return n, nil
}

// ValidateMatcher runs validation on the matcher name, type, and pattern.
var ValidateMatcher = func(m *label.Matcher) error {
	if !label.LabelName(m.Name).IsValid() {
		return fmt.Errorf("invalid label name %q", m.Name)
	}
	switch m.Type {
	case label.MatchEqual, label.MatchNotEqual:
		if !utf8.ValidString(m.Value) {
			return fmt.Errorf("invalid label value %q", m.Value)
		}
	case label.MatchRegexp, label.MatchNotRegexp:
		if _, err := regexp.Compile(m.Value); err != nil {
			return fmt.Errorf("invalid regular expression %q: %s", m.Value, err)
		}
	default:
		return fmt.Errorf("unknown matcher type %q", m.Type)
	}
	return nil
}

func matchesEmpty(m *label.Matcher) bool {
	switch m.Type {
	case label.MatchEqual:
		return m.Value == ""
	case label.MatchRegexp:
		matched, _ := regexp.MatchString(m.Value, "")
		return matched
	default:
		return false
	}
}

func validateSilence(s *ent.Silence) error {
	if s.ID == 0 {
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
	if s.StartsAt.IsZero() {
		return errors.New("invalid zero start timestamp")
	}
	if s.EndsAt.IsZero() {
		return errors.New("invalid zero end timestamp")
	}
	if s.EndsAt.Before(s.StartsAt) {
		return errors.New("end time must not be before start time")
	}
	if s.UpdatedAt.IsZero() {
		return errors.New("invalid zero update timestamp")
	}
	return nil
}

// cloneSilence returns a shallow copy of a silence.
func cloneSilence(sil *ent.Silence) *ent.Silence {
	s := *sil
	return &s
}

func (s *Silences) getSilence(id int) (*ent.Silence, bool) {
	sil, ok := s.st[id]
	if !ok {
		return nil, false
	}
	return sil, true
}

func (s *Silences) setSilence(sil *ent.Silence, now time.Time) error {
	sil.UpdatedAt = now

	if err := validateSilence(sil); err != nil {
		return fmt.Errorf("silence invalid %w", err)
	}
	sil.EndsAt = time.Now().Add(s.Retention)
	if s.st.merge(sil, now) {
		s.version++
	}
	b, err := s.st.marshalBinary(sil)
	if err != nil {
		return err
	}
	if s.Spreader != nil {
		return s.Spreader.Broadcast(b)
	}
	return nil
}

// Set the specified silence. If a silence with the ID already exists and the modification
// modifies history, the old silence gets expired and a new one is created.
func (s *Silences) Set(sil *ent.Silence) (int, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	now := s.nowUTC()
	prev, ok := s.getSilence(sil.ID)
	if ok {
		if canUpdate(prev, sil, now) {
			return sil.ID, s.setSilence(sil, now)
		}
		if getState(prev, s.nowUTC()) != alert.SilenceStateExpired {
			// We cannot update the silence, expire the old one.
			if err := s.expire(prev.ID); err != nil {
				return 0, fmt.Errorf("expire previous silence %w", err)
			}
		}
	}
	sil.ID = int(snowflake.New().Int64())
	// If we got here it's either a new silence or a replacing one.
	if sil.StartsAt.Before(now) {
		sil.StartsAt = now
	}

	return sil.ID, s.setSilence(sil, now)
}

// canUpdate returns true if silence a can be updated to b without
// affecting the historic view of silencing.
func canUpdate(a, b *ent.Silence, now time.Time) bool {
	if !reflect.DeepEqual(a.Matchers, b.Matchers) {
		return false
	}
	// Allowed timestamp modifications depend on the current time.
	switch st := getState(a, now); st {
	case alert.SilenceStateActive:
		if b.StartsAt.Unix() != a.StartsAt.Unix() {
			return false
		}
		if b.EndsAt.Before(now) {
			return false
		}
	case alert.SilenceStatePending:
		if b.StartsAt.Before(now) {
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
func (s *Silences) Expire(id int) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.expire(id)
}

// Expire the silence with the given ID immediately.
// It is idempotent, nil is returned if the silence already expired before it is GC'd.
// If the silence is not found an error is returned.
func (s *Silences) expire(id int) error {
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
		sil.EndsAt = now
	case alert.SilenceStatePending:
		// Set both to now to make Silence move to "expired" state
		sil.StartsAt = now
		sil.EndsAt = now
	}

	return s.setSilence(sil, now)
}

// QueryParam expresses parameters along which silences are queried.
type QueryParam func(*query) error

type query struct {
	ids     []int
	filters []silenceFilter
}

// silenceFilter is a function that returns true if a silence
// should be dropped from a result set for a given time.
type silenceFilter func(*ent.Silence, *Silences, time.Time) (bool, error)

// QIDs configures a query to select the given silence IDs.
func QIDs(ids ...int) QueryParam {
	return func(q *query) error {
		q.ids = append(q.ids, ids...)
		return nil
	}
}

// QMatches returns silences that match the given label set.
func QMatches(set label.LabelSet) QueryParam {
	return func(q *query) error {
		f := func(sil *ent.Silence, s *Silences, _ time.Time) (bool, error) {
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
func getState(sil *ent.Silence, ts time.Time) alert.SilenceState {
	if ts.Before(sil.StartsAt) {
		return alert.SilenceStatePending
	}
	if ts.After(sil.EndsAt) {
		return alert.SilenceStateExpired
	}
	return alert.SilenceStateActive
}

// QState filters queried silences by the given states.
func QState(states ...alert.SilenceState) QueryParam {
	return func(q *query) error {
		f := func(sil *ent.Silence, _ *Silences, now time.Time) (bool, error) {
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
func (s *Silences) QueryOne(params ...QueryParam) (*ent.Silence, error) {
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
func (s *Silences) Query(params ...QueryParam) ([]*ent.Silence, int, error) {
	metrics.Silences.QueriesTotal.Inc()
	defer prometheus.NewTimer(metrics.Silences.QueryDuration).ObserveDuration()

	q := &query{}
	for _, p := range params {
		if err := p(q); err != nil {
			metrics.Silences.QueryErrorsTotal.Inc()
			return nil, s.Version(), err
		}
	}
	sils, version, err := s.query(q, s.nowUTC())
	if err != nil {
		metrics.Silences.QueryErrorsTotal.Inc()
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

func (s *Silences) query(q *query, now time.Time) ([]*ent.Silence, int, error) {
	// If we have no ID constraint, all silences are our base set.  This and
	// the use of post-filter functions is the trivial solution for now.
	var res []*ent.Silence

	s.mtx.Lock()
	defer s.mtx.Unlock()

	if q.ids != nil {
		for _, id := range q.ids {
			if s, ok := s.st[id]; ok {
				res = append(res, s)
			}
		}
	} else {
		for _, sil := range s.st {
			res = append(res, sil)
		}
	}

	var resf []*ent.Silence
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

// loadSnapshot loads all silences data from DataLoader
func (s *Silences) loadData() error {
	datas, err := s.DataLoader()
	if err != nil {
		return err
	}
	st := make(state)
	for _, e := range datas {
		st[e.ID] = e
	}
	s.mtx.Lock()
	s.st = st
	s.version++
	s.mtx.Unlock()

	return nil
}

type state map[int]*ent.Silence

func (s state) merge(e *ent.Silence, now time.Time) bool {
	id := e.ID
	if e.EndsAt.Before(now) {
		return false
	}

	prev, ok := s[id]
	if !ok || prev.UpdatedAt.Before(e.UpdatedAt) {
		s[id] = e
		return true
	}
	return false
}

func (s state) Merge(bs []byte) error {
	var msg []*ent.Silence
	if err := msgpack.Unmarshal(bs, &msg); err != nil {
		return err
	}
	for _, e := range msg {
		s.merge(e, time.Now())
	}
	return nil
}

func (s state) MarshalBinary() ([]byte, error) {
	var msg []*ent.Silence

	for _, e := range s {
		msg = append(msg, e)
	}
	return msgpack.Marshal(msg)
}

func (s state) marshalBinary(e *ent.Silence) ([]byte, error) {
	return msgpack.Marshal([]*ent.Silence{e})
}
