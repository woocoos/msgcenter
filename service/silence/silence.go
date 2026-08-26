// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package silence

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tsingsun/members"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/knockout-go/pkg/snowflake"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/marker"
	"github.com/woocoos/msgcenter/pkg/metrics"
	"github.com/woocoos/msgcenter/pkg/tracing"
	"github.com/woocoos/msgcenter/service/provider/mem"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var (
	logger = log.Component("silence")
	tracer = tracing.NewTracer("github.com/woocoos/msgcenter/silence")
)

var (
	// ErrNotFound is returned if a silence was not found.
	ErrNotFound = fmt.Errorf("silence not found")
)

// matcherIndex stores pre-compiled matcher sets keyed by silence ID.
// Matchers are compiled once when the silence is added, avoiding repeated
// regexp compilation on every query.
type matcherIndex map[int]label.MatcherSet

func (c matcherIndex) get(id int) (label.MatcherSet, error) {
	if m, ok := c[id]; ok {
		return m, nil
	}
	return nil, ErrNotFound
}

// add compiles a silence's matchers and adds them to the cache.
// It returns the compiled matcher set.
// Entry.MatcherSets must be populated before calling add (done by validateSilence).
func (c matcherIndex) add(e *Entry) (label.MatcherSet, error) {
	matcherSet := make(label.MatcherSet, 0, len(e.MatcherSets))

	for _, ms := range e.MatcherSets {
		matchers := make(label.Matchers, len(*ms))
		for i, m := range *ms {
			var mt label.MatchType
			switch m.Type {
			case label.MatchEqual, label.MatchNotEqual, label.MatchRegexp, label.MatchNotRegexp:
				mt = m.Type
			default:
				return nil, fmt.Errorf("unknown matcher type %q", m.Type)
			}
			matcher, err := label.NewMatcher(mt, m.Name, m.Value)
			if err != nil {
				return nil, err
			}
			matchers[i] = matcher
		}

		matcherSet = append(matcherSet, &matchers)
	}

	c[e.ID] = matcherSet
	return matcherSet, nil
}

// versionIndex is an append-only slice sorted by version. It enables O(log n)
// binary search for incremental queries via QSince.
type versionIndex []silenceVersion

type silenceVersion struct {
	id      int
	version int
}

func (vi *versionIndex) add(version int, id int) {
	*vi = append(*vi, silenceVersion{version: version, id: id})
}

// findVersionGreaterThan returns the index of the first entry with version > the given version.
func (vi versionIndex) findVersionGreaterThan(version int) (int, bool) {
	idx := sort.Search(len(vi), func(i int) bool {
		return vi[i].version > version
	})
	return idx, idx < len(vi)
}

var _ members.Shard = (*Silences)(nil)

// Silencer binds together a Silences and an internal cache to implement the
// Muter interface. It no longer depends on a global Marker.
type Silencer struct {
	silences *Silences
	cache    *cache
	// underlying handler for store in db
	callback mem.AlertStoreCallback
}

// NewSilencer returns a new Silencer.
func NewSilencer(s *Silences, callback mem.AlertStoreCallback) *Silencer {
	return &Silencer{
		silences: s,
		cache:    &cache{entries: map[label.Fingerprint]*cacheEntry{}},
		callback: callback,
	}
}

// Mutes implements the Muter interface.
func (s *Silencer) Mutes(ctx context.Context, lset label.LabelSet) bool {
	fp := lset.Fingerprint()
	ctx, span := tracer.Start(ctx, "silence.Silencer.Mutes",
		trace.WithAttributes(
			attribute.String("alerting.alert.fingerprint", fp.String()),
		),
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// Track the silences to set on marker.
	var markedSilences []string
	defer func() {
		m, ok := marker.FromContext(ctx)
		if ok {
			m.SetSilenced(fp, markedSilences)
		}
	}()

	// Get the cached entry for this fingerprint.
	cachedEntry := s.cache.get(fp)

	var (
		err        error
		oldSils    []*Entry
		newSils    []*Entry
		newVersion = cachedEntry.version
	)
	cacheIsUpToDate := cachedEntry.version == s.silences.Version()

	if cacheIsUpToDate && cachedEntry.count() == 0 {
		// Very fast path: no new silences have been added and this lset
		// was not silenced last time we checked.
		span.AddEvent("No new silences to match since last check",
			trace.WithAttributes(
				attribute.Int("alerting.silences.cache.count", cachedEntry.count()),
			),
		)
		return false
	}
	// Either there are new silences and we need to check if those match lset or there were
	// silences last time we queried so we need to see if those are still active/have become
	// active. It's possible for there to be both old and new silences.

	if cachedEntry.count() > 0 {
		// There were old silences for this lset, we need to find them
		// to check if they are still active/pending, or have ended.
		oldSils, _, err = s.silences.Query(
			QIDs(cachedEntry.silenceIDs),
			QState(time.Now(), SilenceStateActive, SilenceStatePending),
		)
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)
			logger.Error("Querying old silences failed, alerts might not get silenced correctly", zap.Error(err))
		}
	}

	if !cacheIsUpToDate {
		// New silences have been added since the last check. Do a full
		// query for any silences newer than the cached version that
		// match the lset.
		newSils, newVersion, err = s.silences.Query(
			QSince(cachedEntry.version),
			QState(time.Now(), SilenceStateActive, SilenceStatePending),
			QMatchers(lset, s.silences.mi),
		)
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)
			logger.Error("Querying silences failed, alerts might not get silenced correctly", zap.Error(err))
		}
	}

	totalSilences := len(oldSils) + len(newSils)
	if totalSilences == 0 {
		// Easy case, neither active nor pending silences anymore.
		s.cache.set(fp, newCacheEntry(newVersion))
		span.AddEvent("No silences to match", trace.WithAttributes(
			attribute.Int("alerting.silences.count", totalSilences),
		))
		return false
	}

	// Categorize old and new silences by their current state.
	// oldSils and newSils may overlap if a cached silence was updated,
	// so we deduplicate by ID.
	activeIDs := make([]string, 0, totalSilences)
	allIDs := make([]int, 0, totalSilences)
	seen := make(map[string]struct{}, totalSilences)
	now := time.Now()

	for _, sils := range [][]*Entry{oldSils, newSils} {
		for _, sil := range sils {
			sid := strconv.Itoa(sil.ID)
			if _, ok := seen[sid]; ok {
				continue
			}
			seen[sid] = struct{}{}
			switch getState(sil, now) {
			case SilenceStatePending:
				allIDs = append(allIDs, sil.ID)
			case SilenceStateActive:
				activeIDs = append(activeIDs, sid)
				allIDs = append(allIDs, sil.ID)
			default:
				// Do nothing, silence has expired in the meantime.
			}
		}
	}
	logger.Debug("determined current silences state",
		zap.Int("total", len(allIDs)),
		zap.Time("now", now),
		zap.Int("active", len(activeIDs)),
		zap.Int("pending", len(allIDs)-len(activeIDs)),
	)

	s.cache.set(fp, newCacheEntry(newVersion, allIDs...))

	t := trace.WithAttributes(
		attribute.Int("alerting.silences.active.count", len(activeIDs)),
		attribute.Int("alerting.silences.pending.count", len(allIDs)-len(activeIDs)),
		attribute.Int("alerting.silences.total.count", len(allIDs)),
	)

	mutes := len(activeIDs) > 0
	if mutes {
		markedSilences = activeIDs
		span.AddEvent("Silencer mutes alert", t)
	} else {
		span.AddEvent("Silencer does not mute alert", t)
	}
	return mutes
}

// The following methods implement mem.AlertStoreCallback.
func (s *Silencer) PreStore(alert *alert.Alert, existing bool) error {
	if s.callback != nil {
		return s.callback.PreStore(alert, existing)
	}
	return nil
}

func (s *Silencer) PostStore(alert *alert.Alert, existing bool) {
	if s.callback != nil {
		s.callback.PostStore(alert, existing)
	}
}

func (s *Silencer) PostDelete(alert *alert.Alert) {
	if s.callback != nil {
		s.callback.PostDelete(alert)
	}
}

// PostGC cleans up cache entries for alerts that have been garbage collected.
// This prevents the cache from growing indefinitely with stale entries.
func (s *Silencer) PostGC(fps []label.Fingerprint) {
	for _, fp := range fps {
		s.cache.delete(fp)
	}
}

// Silences holds a silence state that can be modified, queried, and snapshot.
type Silences struct {
	Options

	mtx     sync.RWMutex
	st      state
	version int // Increments whenever silences are added.
	mi      matcherIndex
	vi      versionIndex
}

// MaintenanceFunc represents the function to run as part of the periodic maintenance for silences.
type MaintenanceFunc func() error

type Option func(*Options)

// WithDataLoader sets the data loader function for the silences.
// if not set, the silences will use the memory to store silences.
// in distributed mode, this should be set to use the sync data.
func WithDataLoader(fn func(ids ...int) ([]*Entry, error)) Option {
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

	DataLoader func(ids ...int) ([]*Entry, error)
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
	options := Options{
		Retention:           time.Hour * 120,
		MaintenanceInterval: time.Minute * 15,
	}
	if err := cfg.Unmarshal(&options); err != nil {
		return nil, err
	}
	for _, opt := range opts {
		opt(&options)
	}
	return New(options)
}

// New returns a new Silences object with the given configuration.
func New(o Options) (*Silences, error) {
	if err := o.validate(); err != nil {
		return nil, err
	}
	s := &Silences{
		Options: o,
		mi:      make(matcherIndex),
		vi:      make(versionIndex, 0),
		st:      state{},
	}

	if metrics.Silences == nil {
		metrics.Silences = metrics.NewSilencesMetrics(func(state string) float64 {
			count, err := s.CountState(SilenceState(state))
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
	return s.st.mergeStream(b, time.Now(), func(e *Entry) {
		// Normalize: convert Matchers to MatcherSets if needed.
		if len(e.MatcherSets) == 0 && len(e.Matchers) > 0 {
			ms := label.Matchers(e.Matchers)
			e.MatcherSets = label.MatcherSet{&ms}
		}
		s.indexSilence(e)
	})
}

func (s *Silences) Start(ctx context.Context) error {
	t := time.NewTicker(s.MaintenanceInterval)
	defer t.Stop()

	runMaintenance := func(do MaintenanceFunc) error {
		metrics.Silences.MaintenanceTotal.Inc()
		if err := do(); err != nil {
			metrics.Silences.MaintenanceErrorsTotal.Inc()
			return err
		}
		return nil
	}
	logger.Debug("silences maintenance started")
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
	now := time.Now()
	defer func() { metrics.Silences.GcDuration.Observe(time.Since(now).Seconds()) }()

	var n int

	s.mtx.Lock()
	defer s.mtx.Unlock()

	for id, sil := range s.st {
		if sil.EndsAt.IsZero() {
			logger.Error("unexpected zero expiration timestamp", zap.Int("id", id))
			delete(s.st, id)
			delete(s.mi, id)
			n++
			continue
		}
		if !sil.EndsAt.After(now) {
			delete(s.st, id)
			delete(s.mi, id)
			n++
		}
	}

	// Rebuild version index, keeping only entries that still exist in state.
	targetVi := s.vi[:0]
	for _, sv := range s.vi {
		if _, ok := s.st[sv.id]; ok {
			targetVi = append(targetVi, sv)
		}
	}
	clear(s.vi[len(targetVi):])
	s.vi = targetVi

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

func validateSilence(s *Entry) error {
	if s.ID == 0 {
		return errors.New("ID missing")
	}

	// Normalize: if MatcherSets is empty but Matchers is set, convert.
	if len(s.MatcherSets) == 0 && len(s.Matchers) > 0 {
		ms := label.Matchers(s.Matchers)
		s.MatcherSets = label.MatcherSet{&ms}
	}

	if len(s.MatcherSets) == 0 {
		return errors.New("at least one matcher required")
	}

	allMatchEmpty := true
	for setIdx, matchers := range s.MatcherSets {
		if len(*matchers) == 0 {
			return fmt.Errorf("matcher set %d is empty", setIdx)
		}
		for i, m := range *matchers {
			if err := ValidateMatcher(m); err != nil {
				return fmt.Errorf("invalid label matcher %d in set %d: %s", i, setIdx, err)
			}
			allMatchEmpty = allMatchEmpty && matchesEmpty(m)
		}
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

func (s *Silences) getSilence(id int) (*Entry, bool) {
	sil, ok := s.st[id]
	if !ok {
		return nil, false
	}
	return sil, true
}

// indexSilence registers a newly added silence in the version index and
// pre-compiles its matchers into the matcher index.
func (s *Silences) indexSilence(e *Entry) {
	s.version++
	s.vi.add(s.version, e.ID)
	if _, err := s.mi.add(e); err != nil {
		logger.Error("failed to compile silence matchers", zap.Int("id", e.ID), zap.Error(err))
	}
}

func (s *Silences) setSilence(sil *Entry, now time.Time) (changed, added bool, err error) {
	sil.UpdatedAt = now

	if err := validateSilence(sil); err != nil {
		return false, false, fmt.Errorf("silence invalid %w", err)
	}
	changed, added = s.st.merge(sil, now)
	if added {
		s.indexSilence(sil)
	}
	if changed {
		b, err := s.st.marshalBinary(sil)
		if err != nil {
			return false, false, err
		}
		if s.Spreader != nil {
			if err := s.Spreader.Broadcast(b); err != nil {
				return false, false, err
			}
		}
	}
	return changed, added, nil
}

// Set the specified silence. If a silence with the ID already exists and the modification
// modifies history, the old silence gets expired and a new one is created.
func (s *Silences) Set(ctx context.Context, sil *Entry) (int, error) {
	_, span := tracer.Start(ctx, "silences.Set")
	defer span.End()

	s.mtx.Lock()
	defer s.mtx.Unlock()

	now := time.Now()
	prev, ok := s.getSilence(sil.ID)

	if ok && canUpdate(prev, sil, now) {
		_, _, err := s.setSilence(sil, now)
		if err != nil {
			return 0, err
		}
		return sil.ID, nil
	}

	// Generate new ID only for silences without one (ID==0).
	// If sil.ID is already set (e.g. from DB sync), keep it.
	if sil.ID == 0 {
		sil.ID = int(snowflake.New().Int64())
	}

	// Set default StartsAt if before now.
	if sil.StartsAt.Before(now) {
		sil.StartsAt = now
	}

	// Set default EndsAt only for new silences.
	if sil.EndsAt.IsZero() {
		sil.EndsAt = now.Add(s.Retention)
	}

	if ok && getState(prev, now) != SilenceStateExpired {
		// We cannot update the silence, expire the old one.
		if err := s.expire(prev.ID); err != nil {
			return 0, fmt.Errorf("expire previous silence %w", err)
		}
	}

	_, _, err := s.setSilence(sil, now)
	if err != nil {
		return 0, err
	}
	return sil.ID, nil
}

// effectiveMatcherSet returns the MatcherSets for an entry, converting from
// Matchers if MatcherSets is empty (backward compatibility).
func effectiveMatcherSet(e *Entry) label.MatcherSet {
	if len(e.MatcherSets) > 0 {
		return e.MatcherSets
	}
	if len(e.Matchers) > 0 {
		ms := label.Matchers(e.Matchers)
		return label.MatcherSet{&ms}
	}
	return nil
}

// canUpdate returns true if silence a can be updated to b without
// affecting the historic view of silencing.
func canUpdate(a, b *Entry, now time.Time) bool {
	// Compare effective matcher (MatcherSets preferred, fallback to Matchers).
	aMatchers := effectiveMatcherSet(a)
	bMatchers := effectiveMatcherSet(b)
	if !reflect.DeepEqual(aMatchers, bMatchers) {
		return false
	}
	// Allowed timestamp modifications depend on the current time.
	switch st := getState(a, now); st {
	case SilenceStateActive:
		if b.StartsAt.Unix() != a.StartsAt.Unix() {
			return false
		}
		if b.EndsAt.Before(now) {
			return false
		}
	case SilenceStatePending:
		if b.StartsAt.Before(now) {
			return false
		}
	case SilenceStateExpired:
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
	now := time.Now()

	switch getState(sil, now) {
	case SilenceStateExpired:
		return nil
	case SilenceStateActive:
		sil.EndsAt = now
	case SilenceStatePending:
		// Set both to now to make Silence move to "expired" state
		sil.StartsAt = now
		sil.EndsAt = now
	}

	_, _, err := s.setSilence(sil, now)
	return err
}

// QueryOne queries with the given parameters and returns the first result.
// Returns ErrNotFound if the query result is empty.
func (s *Silences) QueryOne(qs ...any) (*Entry, error) {
	res, _, err := s.Query(qs...)
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
// Parameters can be EntryQuery (per-entry filters) or QSince (incremental scan).
func (s *Silences) Query(qs ...any) ([]*Entry, int, error) {
	metrics.Silences.QueriesTotal.Inc()
	defer prometheus.NewTimer(metrics.Silences.QueryDuration).ObserveDuration()

	sils, version, err := s.query(qs...)
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
func (s *Silences) CountState(states ...SilenceState) (int, error) {
	// This could probably be optimized.
	sils, _, err := s.Query(QState(time.Now(), states...))
	if err != nil {
		return -1, err
	}
	return len(sils), nil
}

func (s *Silences) query(qs ...any) ([]*Entry, int, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	// Separate QSince from per-entry filters.
	var since *int
	var filters []EntryQuery
	for _, q := range qs {
		switch v := q.(type) {
		case sinceQuery:
			since = &v.version
		case EntryQuery:
			filters = append(filters, v)
		}
	}

	// Determine which entries to scan.
	var entries []*Entry
	if since != nil {
		// Incremental query: only scan silences created after the given version.
		start, found := s.vi.findVersionGreaterThan(*since)
		if !found {
			return nil, s.version, nil
		}
		entries = make([]*Entry, 0, len(s.vi)-start)
		for _, sv := range s.vi[start:] {
			if sil, ok := s.st[sv.id]; ok {
				entries = append(entries, sil)
			}
		}
	} else {
		// Full scan.
		entries = make([]*Entry, 0, len(s.st))
		for _, e := range s.st {
			entries = append(entries, e)
		}
	}

	// Apply per-entry filters.
	var res []*Entry
	for _, e := range entries {
		ok := true
		for _, f := range filters {
			match, err := f(e)
			if err != nil {
				return nil, s.version, err
			}
			if !match {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, cloneSilence(e))
		}
	}
	return res, s.version, nil
}

// loadData loads all silences data from DataLoader
func (s *Silences) loadData() error {
	datas, err := s.DataLoader()
	if err != nil {
		return err
	}
	st := make(state)
	mi := make(matcherIndex, len(datas))
	vi := make(versionIndex, 0, len(datas))
	for _, e := range datas {
		// Normalize: convert Matchers to MatcherSets if needed.
		if len(e.MatcherSets) == 0 && len(e.Matchers) > 0 {
			ms := label.Matchers(e.Matchers)
			e.MatcherSets = label.MatcherSet{&ms}
		}
		st[e.ID] = e
		if _, err := mi.add(e); err != nil {
			logger.Error("failed to compile silence matchers on load", zap.Int("id", e.ID), zap.Error(err))
		}
	}
	s.mtx.Lock()
	s.st = st
	s.mi = mi
	s.version++
	for id := range st {
		vi.add(s.version, id)
	}
	s.vi = vi
	s.mtx.Unlock()

	return nil
}
