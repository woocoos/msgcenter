package silence

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/marker"
)

func newTestSilences(t *testing.T) *Silences {
	t.Helper()
	s, err := New(Options{
		Retention:           time.Hour,
		MaintenanceInterval: time.Minute,
	})
	require.NoError(t, err)
	return s
}

func TestGC_ZeroEndsAt_ContinuesProcessing(t *testing.T) {
	s := newTestSilences(t)
	now := time.Now()

	// Silence with zero EndsAt — previously caused GC to abort immediately.
	bad := &Entry{
		ID:        1,
		Matchers:  []*label.Matcher{{Type: label.MatchEqual, Name: "a", Value: "b"}},
		StartsAt:  now.Add(-time.Hour),
		EndsAt:    time.Time{}, // zero value
		UpdatedAt: now,
	}
	// Expired silence that should also be collected.
	expired := &Entry{
		ID:        2,
		Matchers:  []*label.Matcher{{Type: label.MatchEqual, Name: "c", Value: "d"}},
		StartsAt:  now.Add(-2 * time.Hour),
		EndsAt:    now.Add(-time.Hour), // expired
		UpdatedAt: now.Add(-2 * time.Hour),
	}
	// Active silence that should NOT be collected.
	active := &Entry{
		ID:        3,
		Matchers:  []*label.Matcher{{Type: label.MatchEqual, Name: "e", Value: "f"}},
		StartsAt:  now.Add(-time.Minute),
		EndsAt:    now.Add(time.Hour), // still active
		UpdatedAt: now.Add(-time.Minute),
	}

	s.st[bad.ID] = bad
	s.st[expired.ID] = expired
	s.st[active.ID] = active

	n, err := s.GC()
	require.NoError(t, err)
	// Both bad (zero EndsAt) and expired should be removed.
	assert.Equal(t, 2, n)

	_, exists := s.st[bad.ID]
	assert.False(t, exists, "silence with zero EndsAt should be removed")

	_, exists = s.st[expired.ID]
	assert.False(t, exists, "expired silence should be removed")

	_, exists = s.st[active.ID]
	assert.True(t, exists, "active silence should be kept")
}

func TestGC_AllZeroEndsAt(t *testing.T) {
	s := newTestSilences(t)
	now := time.Now()

	for i := 1; i <= 3; i++ {
		s.st[i] = &Entry{
			ID:        i,
			Matchers:  []*label.Matcher{{Type: label.MatchEqual, Name: "a", Value: "b"}},
			StartsAt:  now.Add(-time.Hour),
			EndsAt:    time.Time{},
			UpdatedAt: now,
		}
	}

	n, err := s.GC()
	require.NoError(t, err)
	assert.Equal(t, 3, n)
	assert.Empty(t, s.st)
}

func TestGC_MatcherIndexCleaned(t *testing.T) {
	s := newTestSilences(t)
	now := time.Now()

	sil := &Entry{
		ID:        1,
		Matchers:  []*label.Matcher{{Type: label.MatchEqual, Name: "a", Value: "b"}},
		StartsAt:  now.Add(-time.Hour),
		EndsAt:    time.Time{},
		UpdatedAt: now,
	}
	s.st[sil.ID] = sil

	// Populate matcher index.
	_, err := s.mi.add(sil)
	require.NoError(t, err)
	assert.Len(t, s.mi, 1)

	_, err = s.GC()
	require.NoError(t, err)
	assert.Empty(t, s.mi, "matcher index should be cleaned for removed silences")
}

func TestNew_Validation(t *testing.T) {
	_, err := New(Options{})
	require.Error(t, err, "zero MaintenanceInterval should fail validation")

	s, err := New(Options{
		Retention:           time.Hour,
		MaintenanceInterval: time.Minute,
	})
	require.NoError(t, err)
	require.NotNil(t, s)
}

func TestSet_And_Query(t *testing.T) {
	s := newTestSilences(t)

	sil := &Entry{
		Matchers: []*label.Matcher{{Type: label.MatchEqual, Name: "alertname", Value: "test"}},
		StartsAt: time.Now(),
		EndsAt:   time.Now().Add(time.Hour),
	}

	id, err := s.Set(context.Background(), sil)
	require.NoError(t, err)
	require.NotZero(t, id)

	results, version, err := s.Query(QState(time.Now(), SilenceStateActive))
	require.NoError(t, err)
	assert.Greater(t, version, 0)
	require.Len(t, results, 1)
	assert.Equal(t, id, results[0].ID)
}

func TestVersionIndex_findAfter(t *testing.T) {
	t.Parallel()

	vi := versionIndex{
		{version: 1, id: 10},
		{version: 3, id: 20},
		{version: 5, id: 30},
		{version: 7, id: 40},
	}

	tests := []struct {
		name      string
		version   int
		wantIdx   int
		wantFound bool
	}{
		{"before all", 0, 0, true},
		{"at first", 1, 1, true},
		{"between", 4, 2, true},
		{"at last", 7, 4, false},
		{"after all", 100, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, found := vi.findVersionGreaterThan(tt.version)
			assert.Equal(t, tt.wantIdx, idx)
			assert.Equal(t, tt.wantFound, found)
		})
	}
}

func TestMatcherIndex_addAndGet(t *testing.T) {
	t.Parallel()

	mi := make(matcherIndex)
	matchers := label.Matchers{{Type: label.MatchEqual, Name: "a", Value: "b"}}
	e := &Entry{
		ID:          1,
		MatcherSets: label.MatcherSet{&matchers},
	}

	ms, err := mi.add(e)
	require.NoError(t, err)
	assert.Len(t, ms, 1)

	got, err := mi.get(1)
	require.NoError(t, err)
	assert.Equal(t, ms, got)

	_, err = mi.get(999)
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestQSince_IncrementalQuery(t *testing.T) {
	s := newTestSilences(t)
	now := time.Now()

	// Add first silence.
	s1 := &Entry{
		Matchers:  []*label.Matcher{{Type: label.MatchEqual, Name: "a", Value: "b"}},
		StartsAt:  now.Add(-time.Minute),
		EndsAt:    now.Add(time.Hour),
		UpdatedAt: now,
	}
	id1, err := s.Set(context.Background(), s1)
	require.NoError(t, err)
	v1 := s.Version()

	// Add second silence.
	s2 := &Entry{
		Matchers:  []*label.Matcher{{Type: label.MatchEqual, Name: "c", Value: "d"}},
		StartsAt:  now.Add(-time.Minute),
		EndsAt:    now.Add(time.Hour),
		UpdatedAt: now,
	}
	id2, err := s.Set(context.Background(), s2)
	require.NoError(t, err)
	_ = id1

	// QSince(v1) should only return silences created after v1 (i.e. s2).
	results, _, err := s.Query(QSince(v1), QState(time.Now(), SilenceStateActive))
	require.NoError(t, err)
	require.Len(t, results, 1)
	assert.Equal(t, id2, results[0].ID)

	// QSince(0) should return all silences.
	results, _, err = s.Query(QSince(0), QState(time.Now(), SilenceStateActive))
	require.NoError(t, err)
	assert.Len(t, results, 2)

	// QSince with a very high version should return nothing.
	results, _, err = s.Query(QSince(1000))
	require.NoError(t, err)
	assert.Empty(t, results)
}

func TestStateMerge_ReturnsAdded(t *testing.T) {
	t.Parallel()

	st := make(state)
	now := time.Now()

	e := &Entry{
		ID:        1,
		Matchers:  []*label.Matcher{{Type: label.MatchEqual, Name: "a", Value: "b"}},
		StartsAt:  now,
		EndsAt:    now.Add(time.Hour),
		UpdatedAt: now,
	}

	// First merge: should be merged and added.
	merged, added := st.merge(e, now)
	assert.True(t, merged)
	assert.True(t, added)

	// Second merge with same data: merged but not added.
	merged, added = st.merge(e, now)
	assert.False(t, merged) // same UpdatedAt, not before
	assert.False(t, added)

	// Merge with newer UpdatedAt: merged but not added (update).
	e2 := *e
	e2.UpdatedAt = now.Add(time.Second)
	merged, added = st.merge(&e2, now)
	assert.True(t, merged)
	assert.False(t, added)
}

func TestGC_CleansVersionIndex(t *testing.T) {
	s := newTestSilences(t)
	now := time.Now()

	// Add two silences and track their IDs.
	var ids []int
	for i := 0; i < 2; i++ {
		e := &Entry{
			Matchers:  []*label.Matcher{{Type: label.MatchEqual, Name: "a", Value: "b"}},
			StartsAt:  now.Add(-time.Hour),
			EndsAt:    now.Add(time.Hour),
			UpdatedAt: now,
		}
		id, err := s.Set(context.Background(), e)
		require.NoError(t, err)
		ids = append(ids, id)
	}
	assert.Equal(t, 2, len(s.vi))

	// Expire the first silence.
	s.st[ids[0]].EndsAt = now.Add(-time.Second)

	n, err := s.GC()
	require.NoError(t, err)
	assert.Equal(t, 1, n)
	// Version index should only contain the active silence.
	assert.Equal(t, 1, len(s.vi))
	assert.Equal(t, ids[1], s.vi[0].id)
}

func TestSetSilence_PrecompilesMatchers(t *testing.T) {
	s := newTestSilences(t)

	sil := &Entry{
		Matchers: []*label.Matcher{{Type: label.MatchEqual, Name: "alertname", Value: "test"}},
		StartsAt: time.Now(),
		EndsAt:   time.Now().Add(time.Hour),
	}

	id, err := s.Set(context.Background(), sil)
	require.NoError(t, err)

	// Matcher should be pre-compiled in the index.
	ms, err := s.mi.get(id)
	require.NoError(t, err)
	assert.Len(t, ms, 1)

	// Version index should have an entry.
	assert.Equal(t, 1, len(s.vi))
	assert.Equal(t, id, s.vi[0].id)
}

// BenchmarkQSince compares incremental vs full scan query performance.
func BenchmarkQSince(b *testing.B) {
	s, _ := New(Options{
		Retention:           time.Hour,
		MaintenanceInterval: time.Minute,
	})
	now := time.Now()

	// Populate 1000 silences.
	for i := 1; i <= 1000; i++ {
		e := &Entry{
			ID:        i,
			Matchers:  []*label.Matcher{{Type: label.MatchEqual, Name: "alertname", Value: fmt.Sprintf("alert%d", i)}},
			StartsAt:  now.Add(-time.Minute),
			EndsAt:    now.Add(time.Hour),
			UpdatedAt: now,
		}
		s.st[e.ID] = e
		s.indexSilence(e)
	}

	versionAt900 := s.vi[899].version

	b.Run("full_scan", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s.Query(QState(time.Now(), SilenceStateActive))
		}
	})

	b.Run("incremental_QSince", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s.Query(QSince(versionAt900), QState(time.Now(), SilenceStateActive))
		}
	})
}

// checkMutes verifies the marker recorded the expected silenced state.
func checkMutes(t *testing.T, m marker.AlertMarker, target label.LabelSet, wantSilenced bool, msgAndArgs ...any) {
	t.Helper()
	fp := target.Fingerprint()
	status := m.Status(fp)
	if wantSilenced {
		require.Equal(t, alert.AlertStateSuppressed, status.State, msgAndArgs...)
	} else {
		require.NotEqual(t, alert.AlertStateSuppressed, status.State, msgAndArgs...)
	}
}

func TestSilencer(t *testing.T) {
	ss := newTestSilences(t)
	s := NewSilencer(ss, nil)

	lset := label.LabelSet{"foo": "bar"}

	// No silences — should not mute.
	m := marker.NewAlertMarker()
	ctx := marker.WithContext(context.Background(), m)
	require.False(t, s.Mutes(ctx, lset), "expected alert not silenced without any silences")
	checkMutes(t, m, lset, false, "expected marker not silenced without any silences")

	// Add a non-matching silence.
	now := time.Now()
	_, err := ss.Set(context.Background(), &Entry{
		Matchers: []*label.Matcher{{Type: label.MatchEqual, Name: "foo", Value: "baz"}},
		StartsAt: now.Add(-time.Hour),
		EndsAt:   now.Add(5 * time.Minute),
	})
	require.NoError(t, err)

	m = marker.NewAlertMarker()
	ctx = marker.WithContext(context.Background(), m)
	require.False(t, s.Mutes(ctx, lset), "expected alert not silenced by non-matching silence")
	checkMutes(t, m, lset, false, "expected marker not silenced by non-matching silence")

	// Add a matching silence (expires in 2 seconds from now).
	expireAt := time.Now().Add(2 * time.Second)
	_, err = ss.Set(context.Background(), &Entry{
		Matchers: []*label.Matcher{{Type: label.MatchEqual, Name: "foo", Value: "bar"}},
		StartsAt: now.Add(-time.Hour),
		EndsAt:   expireAt,
	})
	require.NoError(t, err)

	m = marker.NewAlertMarker()
	ctx = marker.WithContext(context.Background(), m)
	require.True(t, s.Mutes(ctx, lset), "expected alert silenced by matching silence")
	checkMutes(t, m, lset, true, "expected marker silenced by matching silence")

	// Wait for the silence to expire.
	time.Sleep(time.Until(expireAt) + time.Second)

	m = marker.NewAlertMarker()
	ctx = marker.WithContext(context.Background(), m)
	require.False(t, s.Mutes(ctx, lset), "expected alert not silenced by expired silence")
	checkMutes(t, m, lset, false, "expected marker not silenced by expired silence")
}

func TestSilencerPendingThenActive(t *testing.T) {
	ss := newTestSilences(t)
	s := NewSilencer(ss, nil)

	lset := label.LabelSet{"foo": "bar"}
	now := time.Now()

	// Add a pending silence (starts in the future).
	_, err := ss.Set(context.Background(), &Entry{
		Matchers: []*label.Matcher{{Type: label.MatchEqual, Name: "foo", Value: "bar"}},
		StartsAt: now.Add(time.Hour),
		EndsAt:   now.Add(3 * time.Hour),
	})
	require.NoError(t, err)

	m := marker.NewAlertMarker()
	ctx := marker.WithContext(context.Background(), m)
	require.False(t, s.Mutes(ctx, lset), "expected alert not silenced by future silence")
	checkMutes(t, m, lset, false, "expected marker not silenced by future silence")
}

func TestSilencerOverlappingSilences(t *testing.T) {
	ss := newTestSilences(t)
	s := NewSilencer(ss, nil)

	lset := label.LabelSet{"foo": "bar"}
	now := time.Now()

	// Add first matching silence (short-lived).
	_, err := ss.Set(context.Background(), &Entry{
		Matchers: []*label.Matcher{{Type: label.MatchEqual, Name: "foo", Value: "bar"}},
		StartsAt: now.Add(-time.Hour),
		EndsAt:   now.Add(2 * time.Second),
	})
	require.NoError(t, err)

	m := marker.NewAlertMarker()
	ctx := marker.WithContext(context.Background(), m)
	require.True(t, s.Mutes(ctx, lset), "expected alert silenced by first silence")
	checkMutes(t, m, lset, true)

	// Add a second matching silence with regexp (longer-lived).
	_, err = ss.Set(context.Background(), &Entry{
		Matchers: []*label.Matcher{{Type: label.MatchRegexp, Name: "foo", Value: "b.."}},
		StartsAt: now.Add(-time.Hour),
		EndsAt:   now.Add(time.Hour),
	})
	require.NoError(t, err)

	// Still silenced (both match).
	m = marker.NewAlertMarker()
	ctx = marker.WithContext(context.Background(), m)
	require.True(t, s.Mutes(ctx, lset), "expected alert still silenced")
	checkMutes(t, m, lset, true)

	// Wait for first silence to expire; second should still be active.
	time.Sleep(3 * time.Second)

	m = marker.NewAlertMarker()
	ctx = marker.WithContext(context.Background(), m)
	require.True(t, s.Mutes(ctx, lset), "expected alert silenced by overlapping second silence")
	checkMutes(t, m, lset, true)
}

func TestSilencerPostDeleteEvictsCache(t *testing.T) {
	ss := newTestSilences(t)
	s := NewSilencer(ss, nil)

	lset := label.LabelSet{"foo": "bar"}
	fp := lset.Fingerprint()
	now := time.Now()

	// Create a matching silence.
	_, err := ss.Set(context.Background(), &Entry{
		Matchers: []*label.Matcher{{Type: label.MatchEqual, Name: "foo", Value: "bar"}},
		StartsAt: now.Add(-time.Hour),
		EndsAt:   now.Add(time.Hour),
	})
	require.NoError(t, err)

	// Mutes populates the cache.
	m := marker.NewAlertMarker()
	ctx := marker.WithContext(context.Background(), m)
	require.True(t, s.Mutes(ctx, lset))
	checkMutes(t, m, lset, true, "expected marker silenced after initial Mutes")
	entry := s.cache.get(fp)
	require.Positive(t, entry.count(), "cache should have entries after Mutes()")

	// PostGC evicts the cache entry for this fingerprint.
	s.PostGC([]label.Fingerprint{fp})
	entry = s.cache.get(fp)
	require.Equal(t, 0, entry.count(), "cache should be empty after PostGC()")
	require.Equal(t, 0, entry.version, "version should be zero for evicted entry")

	// Mutes re-evaluates from scratch (cache miss) and still finds the silence.
	m = marker.NewAlertMarker()
	ctx = marker.WithContext(context.Background(), m)
	require.True(t, s.Mutes(ctx, lset), "expected alert still silenced after cache eviction")
	checkMutes(t, m, lset, true, "expected marker silenced after cache eviction")
	entry = s.cache.get(fp)
	require.Positive(t, entry.count(), "cache should be repopulated after Mutes()")

	// PostGC for a different fingerprint should not affect this entry.
	otherLset := label.LabelSet{"other": "alert"}
	s.PostGC([]label.Fingerprint{otherLset.Fingerprint()})
	entry = s.cache.get(fp)
	require.Positive(t, entry.count(), "unrelated PostGC should not evict other entries")
}
