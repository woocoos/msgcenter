package inhibit

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/marker"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service/provider"
	"github.com/woocoos/msgcenter/service/store"
)

func newTestRule(equal ...label.LabelName) *InhibitRule {
	rule := &InhibitRule{
		SourceMatchers: label.Matchers{
			&label.Matcher{Type: label.MatchEqual, Name: "alertname", Value: "source"},
		},
		TargetMatchers: label.Matchers{
			&label.Matcher{Type: label.MatchEqual, Name: "alertname", Value: "target"},
		},
		Equal:  make(map[label.LabelName]struct{}),
		scache: newTestStore(),
		sindex: newIndex(),
	}
	for _, ln := range equal {
		rule.Equal[ln] = struct{}{}
	}
	rule.scache.SetGCCallback(rule.gcCallback)
	return rule
}

func newTestStore() *store.Alerts {
	return store.NewAlerts()
}

func firingAlert(labels label.LabelSet) *alert.Alert {
	return &alert.Alert{
		Labels:    labels,
		StartsAt:  time.Now(),
		EndsAt:    time.Now().Add(time.Hour),
		UpdatedAt: time.Now(),
	}
}

func TestInhibitRule_fingerprintEquals(t *testing.T) {
	t.Parallel()
	rule := newTestRule("severity", "instance")

	ls1 := label.LabelSet{"alertname": "src", "severity": "critical", "instance": "host1", "extra": "x"}
	ls2 := label.LabelSet{"alertname": "other", "severity": "critical", "instance": "host1", "extra": "y"}
	ls3 := label.LabelSet{"alertname": "src", "severity": "warning", "instance": "host1"}

	// Same Equal-label values → same fingerprint, regardless of other labels.
	assert.Equal(t, rule.fingerprintEquals(ls1), rule.fingerprintEquals(ls2))
	// Different Equal-label values → different fingerprint.
	assert.NotEqual(t, rule.fingerprintEquals(ls1), rule.fingerprintEquals(ls3))
}

func TestInhibitRule_updateIndex_And_findEqualSourceAlert(t *testing.T) {
	t.Parallel()
	rule := newTestRule("severity")

	src := firingAlert(label.LabelSet{"alertname": "source", "severity": "critical"})
	require.NoError(t, rule.scache.Set(src))
	rule.updateIndex(src)

	// Should find the source alert via index.
	target := label.LabelSet{"alertname": "target", "severity": "critical"}
	found, ok := rule.findEqualSourceAlert(target)
	require.True(t, ok)
	assert.Equal(t, src.Fingerprint(), found.Fingerprint())

	// Non-matching Equal labels should not find anything.
	other := label.LabelSet{"alertname": "target", "severity": "warning"}
	_, ok = rule.findEqualSourceAlert(other)
	assert.False(t, ok)
}

func TestInhibitRule_updateIndex_ResolvedOverride(t *testing.T) {
	t.Parallel()
	rule := newTestRule("severity")

	// First source alert.
	src1 := firingAlert(label.LabelSet{"alertname": "source", "severity": "critical"})
	require.NoError(t, rule.scache.Set(src1))
	rule.updateIndex(src1)

	// Second source alert with same Equal labels but later EndsAt.
	src2 := firingAlert(label.LabelSet{"alertname": "source", "severity": "critical", "instance": "host2"})
	src2.EndsAt = time.Now().Add(2 * time.Hour)
	require.NoError(t, rule.scache.Set(src2))
	rule.updateIndex(src2)

	// Index should point to the alert with later EndsAt.
	target := label.LabelSet{"alertname": "target", "severity": "critical"}
	found, ok := rule.findEqualSourceAlert(target)
	require.True(t, ok)
	assert.Equal(t, src2.Fingerprint(), found.Fingerprint())
}

func TestInhibitRuleHasEqual(t *testing.T) {
	t.Parallel()

	now := time.Now()
	cases := []struct {
		name    string
		initial []*alert.Alert
		equal   []label.LabelName
		input   label.LabelSet
		result  bool
	}{
		{
			name:    "no source alerts",
			initial: nil,
			input:   label.LabelSet{"a": "b"},
			result:  false,
		},
		{
			name: "no equal labels, any source alert satisfies the requirement",
			initial: []*alert.Alert{
				{Labels: label.LabelSet{"a": "b"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(time.Hour), UpdatedAt: now},
			},
			input:  label.LabelSet{"a": "b"},
			result: true,
		},
		{
			name: "matching but already resolved",
			initial: []*alert.Alert{
				{Labels: label.LabelSet{"a": "b", "b": "f"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(-time.Second), UpdatedAt: now},
				{Labels: label.LabelSet{"a": "b", "b": "c"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(-time.Second), UpdatedAt: now},
			},
			equal:  []label.LabelName{"a", "b"},
			input:  label.LabelSet{"a": "b", "b": "c"},
			result: false,
		},
		{
			name: "matching and unresolved",
			initial: []*alert.Alert{
				{Labels: label.LabelSet{"a": "b", "c": "d"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(-time.Second), UpdatedAt: now},
				{Labels: label.LabelSet{"a": "b", "c": "f"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(time.Hour), UpdatedAt: now},
			},
			equal:  []label.LabelName{"a"},
			input:  label.LabelSet{"a": "b"},
			result: true,
		},
		{
			name: "equal label does not match",
			initial: []*alert.Alert{
				{Labels: label.LabelSet{"a": "c", "c": "d"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(-time.Second), UpdatedAt: now},
				{Labels: label.LabelSet{"a": "c", "c": "f"}, StartsAt: now.Add(-time.Minute), EndsAt: now.Add(-time.Second), UpdatedAt: now},
			},
			equal:  []label.LabelName{"a"},
			input:  label.LabelSet{"a": "b"},
			result: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := &InhibitRule{
				Equal:  map[label.LabelName]struct{}{},
				scache: newTestStore(),
				sindex: newIndex(),
			}
			for _, ln := range c.equal {
				r.Equal[ln] = struct{}{}
			}
			for _, a := range c.initial {
				require.NoError(t, r.scache.Set(a))
				r.updateIndex(a)
			}

			_, have := r.hasEqual(c.input, false)
			require.Equal(t, c.result, have)
		})
	}
}

func TestInhibitRule_gcCallback(t *testing.T) {
	t.Parallel()
	rule := newTestRule("severity")

	src := firingAlert(label.LabelSet{"alertname": "source", "severity": "critical"})
	require.NoError(t, rule.scache.Set(src))
	rule.updateIndex(src)
	assert.Equal(t, 1, rule.sindex.Len())

	// Simulate GC callback with the resolved alert.
	rule.gcCallback([]*alert.Alert{src})
	assert.Equal(t, 0, rule.sindex.Len())
}

func TestIndex_ConcurrentAccess(t *testing.T) {
	t.Parallel()
	idx := newIndex()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 1000; i++ {
			idx.Set(label.Fingerprint(i), label.Fingerprint(i*2))
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			idx.Get(label.Fingerprint(i))
		}
	}()

	<-done
	assert.Equal(t, 1000, idx.Len())
}

// checkMutes calls ih.Mutes with a fresh AlertMarker in the context
// and asserts the mute result matches wantMuted.
func checkMutes(t *testing.T, ih *Inhibitor, target label.LabelSet, wantMuted bool, msgAndArgs ...any) {
	t.Helper()
	m := marker.NewAlertMarker()
	ctx := marker.WithContext(context.Background(), m)
	got := ih.Mutes(ctx, target)
	require.Equal(t, wantMuted, got, msgAndArgs...)
}

// TestInhibitRuleMatches verifies inhibition through the full Mutes path,
// adapted from upstream TestInhibitRuleMatches.
func TestInhibitRuleMatches(t *testing.T) {
	t.Parallel()

	rule1 := profile.InhibitRule{
		SourceMatchers: label.Matchers{&label.Matcher{Type: label.MatchEqual, Name: "s1", Value: "1"}},
		TargetMatchers: label.Matchers{&label.Matcher{Type: label.MatchEqual, Name: "t1", Value: "1"}},
		Equal:          []label.LabelName{"e"},
	}
	rule2 := profile.InhibitRule{
		SourceMatchers: label.Matchers{&label.Matcher{Type: label.MatchEqual, Name: "s2", Value: "1"}},
		TargetMatchers: label.Matchers{&label.Matcher{Type: label.MatchEqual, Name: "t2", Value: "1"}},
		Equal:          []label.LabelName{"e"},
	}

	ih := NewInhibitor(nil, []profile.InhibitRule{rule1, rule2})
	now := time.Now()
	sourceAlert1 := &alert.Alert{
		Labels:   label.LabelSet{"s1": "1", "t1": "2", "e": "1"},
		StartsAt: now.Add(-time.Minute),
		EndsAt:   now.Add(time.Hour),
	}
	sourceAlert2 := &alert.Alert{
		Labels:   label.LabelSet{"s2": "1", "t2": "1", "e": "1"},
		StartsAt: now.Add(-time.Minute),
		EndsAt:   now.Add(time.Hour),
	}

	ih.rules[0].scache = store.NewAlerts()
	ih.rules[0].scache.Set(sourceAlert1)
	ih.rules[0].sindex = newIndex()
	ih.rules[0].updateIndex(sourceAlert1)

	ih.rules[1].scache = store.NewAlerts()
	ih.rules[1].scache.Set(sourceAlert2)
	ih.rules[1].sindex = newIndex()
	ih.rules[1].updateIndex(sourceAlert2)

	cases := []struct {
		target   label.LabelSet
		expected bool
	}{
		{target: label.LabelSet{"t1": "1", "e": "1"}, expected: true},
		{target: label.LabelSet{"t2": "1", "e": "1"}, expected: true},
		{target: label.LabelSet{"t1": "1", "t3": "1", "e": "1"}, expected: true},
		{target: label.LabelSet{"t1": "1", "t2": "1", "e": "1"}, expected: true},
		{target: label.LabelSet{"t1": "0", "e": "1"}, expected: false},
		{target: label.LabelSet{"s1": "1", "t1": "1", "e": "1"}, expected: true},
		{target: label.LabelSet{"s2": "1", "t2": "1", "e": "1"}, expected: false},
		{target: label.LabelSet{"t1": "1", "e": "0"}, expected: false},
	}

	for _, c := range cases {
		checkMutes(t, ih, c.target, c.expected, "target %v", c.target)
	}
}

// TestInhibitRuleMatchers verifies inhibition with mixed matcher types,
// adapted from upstream TestInhibitRuleMatchers.
func TestInhibitRuleMatchers(t *testing.T) {
	t.Parallel()

	rule1 := profile.InhibitRule{
		SourceMatchers: label.Matchers{&label.Matcher{Type: label.MatchEqual, Name: "s1", Value: "1"}},
		TargetMatchers: label.Matchers{&label.Matcher{Type: label.MatchNotEqual, Name: "t1", Value: "1"}},
		Equal:          []label.LabelName{"e"},
	}
	rule2 := profile.InhibitRule{
		SourceMatchers: label.Matchers{&label.Matcher{Type: label.MatchEqual, Name: "s2", Value: "1"}},
		TargetMatchers: label.Matchers{&label.Matcher{Type: label.MatchEqual, Name: "t2", Value: "1"}},
		Equal:          []label.LabelName{"e"},
	}

	ih := NewInhibitor(nil, []profile.InhibitRule{rule1, rule2})
	now := time.Now()
	sourceAlert1 := &alert.Alert{
		Labels:   label.LabelSet{"s1": "1", "t1": "2", "e": "1"},
		StartsAt: now.Add(-time.Minute),
		EndsAt:   now.Add(time.Hour),
	}
	sourceAlert2 := &alert.Alert{
		Labels:   label.LabelSet{"s2": "1", "t2": "1", "e": "1"},
		StartsAt: now.Add(-time.Minute),
		EndsAt:   now.Add(time.Hour),
	}

	ih.rules[0].scache = store.NewAlerts()
	ih.rules[0].scache.Set(sourceAlert1)
	ih.rules[0].sindex = newIndex()
	ih.rules[0].updateIndex(sourceAlert1)

	ih.rules[1].scache = store.NewAlerts()
	ih.rules[1].scache.Set(sourceAlert2)
	ih.rules[1].sindex = newIndex()
	ih.rules[1].updateIndex(sourceAlert2)

	cases := []struct {
		target   label.LabelSet
		expected bool
	}{
		{target: label.LabelSet{"t1": "1", "e": "1"}, expected: false},
		{target: label.LabelSet{"t2": "1", "e": "1"}, expected: true},
		{target: label.LabelSet{"t1": "1", "t3": "1", "e": "1"}, expected: false},
		{target: label.LabelSet{"t1": "1", "t2": "1", "e": "1"}, expected: true},
		{target: label.LabelSet{"t1": "0", "e": "1"}, expected: true},
		{target: label.LabelSet{"s1": "1", "t1": "1", "e": "1"}, expected: false},
		{target: label.LabelSet{"s2": "1", "t2": "1", "e": "1"}, expected: true},
		{target: label.LabelSet{"t1": "1", "e": "0"}, expected: false},
	}

	for _, c := range cases {
		checkMutes(t, ih, c.target, c.expected, "target %v", c.target)
	}
}

// TestInhibitRuleName verifies that named and unnamed rules are handled correctly,
// adapted from upstream TestInhibitRuleName.
func TestInhibitRuleName(t *testing.T) {
	t.Parallel()

	config1 := profile.InhibitRule{
		Name: "test-rule",
		SourceMatchers: label.Matchers{
			&label.Matcher{Type: label.MatchEqual, Name: "severity", Value: "critical"},
		},
		TargetMatchers: label.Matchers{
			&label.Matcher{Type: label.MatchEqual, Name: "severity", Value: "warning"},
		},
		Equal: []label.LabelName{"instance"},
	}
	config2 := profile.InhibitRule{
		SourceMatchers: label.Matchers{
			&label.Matcher{Type: label.MatchEqual, Name: "severity", Value: "critical"},
		},
		TargetMatchers: label.Matchers{
			&label.Matcher{Type: label.MatchEqual, Name: "severity", Value: "warning"},
		},
		Equal: []label.LabelName{"instance"},
	}

	rule1 := NewInhibitRule(config1)
	rule2 := NewInhibitRule(config2)

	require.Equal(t, "test-rule", rule1.Name, "Expected named rule to adopt name from config")
	require.Empty(t, rule2.Name, "Expected unnamed rule to have empty name")
}

// fakeAlerts implements provider.Alerts for integration testing.
type fakeAlerts struct {
	alerts   []*alert.Alert
	finished chan struct{}
}

func newFakeAlerts(alerts []*alert.Alert) *fakeAlerts {
	return &fakeAlerts{
		alerts:   alerts,
		finished: make(chan struct{}),
	}
}

func (f *fakeAlerts) Start(context.Context) error                          { return nil }
func (f *fakeAlerts) Stop(context.Context) error                           { return nil }
func (f *fakeAlerts) GetPending() provider.AlertIterator                   { return nil }
func (f *fakeAlerts) Get(label.Fingerprint) (*alert.Alert, error)          { return nil, nil }
func (f *fakeAlerts) Put(context.Context, ...*alert.Alert) error           { return nil }
func (f *fakeAlerts) Subscribe(name string) provider.AlertIterator         { return nil }
func (f *fakeAlerts) SlurpAndSubscribe(name string) ([]*alert.Alert, provider.AlertIterator) {
	ch := make(chan *provider.Alert)
	done := make(chan struct{})
	go func() {
		for _, a := range f.alerts {
			ch <- &provider.Alert{Data: a, Header: map[string]string{}}
		}
		ch <- &provider.Alert{
			Data:   &alert.Alert{Labels: label.LabelSet{}, StartsAt: time.Now()},
			Header: map[string]string{},
		}
		close(f.finished)
		<-done
	}()
	return nil, provider.NewAlertIterator(ch, done, nil)
}

// TestInhibit is a full integration test that exercises the Inhibitor lifecycle,
// adapted from upstream TestInhibit.
func TestInhibit(t *testing.T) {
	t.Parallel()

	now := time.Now()
	inhibitRule := func() profile.InhibitRule {
		return profile.InhibitRule{
			SourceMatchers: label.Matchers{&label.Matcher{Type: label.MatchEqual, Name: "s", Value: "1"}},
			TargetMatchers: label.Matchers{&label.Matcher{Type: label.MatchEqual, Name: "t", Value: "1"}},
			Equal:          []label.LabelName{"e"},
		}
	}
	alertOne := func() *alert.Alert {
		return &alert.Alert{
			Labels:   label.LabelSet{"t": "1", "e": "f"},
			StartsAt: now.Add(-time.Minute),
			EndsAt:   now.Add(time.Hour),
		}
	}
	alertTwo := func(resolved bool) *alert.Alert {
		var end time.Time
		if resolved {
			end = now.Add(-time.Second)
		} else {
			end = now.Add(time.Hour)
		}
		return &alert.Alert{
			Labels:   label.LabelSet{"s": "1", "e": "f"},
			StartsAt: now.Add(-time.Minute),
			EndsAt:   end,
		}
	}

	type exp struct {
		lbls  label.LabelSet
		muted bool
	}
	for i, tc := range []struct {
		alerts   []*alert.Alert
		expected []exp
	}{
		{
			alerts: []*alert.Alert{alertOne()},
			expected: []exp{
				{lbls: label.LabelSet{"t": "1", "e": "f"}, muted: false},
			},
		},
		{
			alerts: []*alert.Alert{alertOne(), alertTwo(false)},
			expected: []exp{
				{lbls: label.LabelSet{"t": "1", "e": "f"}, muted: true},
				{lbls: label.LabelSet{"s": "1", "e": "f"}, muted: false},
			},
		},
		{
			alerts: []*alert.Alert{alertOne(), alertTwo(false), alertTwo(true)},
			expected: []exp{
				{lbls: label.LabelSet{"t": "1", "e": "f"}, muted: false},
				{lbls: label.LabelSet{"s": "1", "e": "f"}, muted: false},
			},
		},
	} {
		ap := newFakeAlerts(tc.alerts)
		inhibitor := NewInhibitor(ap, []profile.InhibitRule{inhibitRule()})

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ap.finished != nil {
				select {
				case <-ap.finished:
					ap.finished = nil
				default:
				}
			}
			inhibitor.Stop()
		}()
		inhibitor.Run()
		wg.Wait()

		for _, expected := range tc.expected {
			checkMutes(t, inhibitor, expected.lbls, expected.muted, "tc: %d, labels %q", i, expected.lbls)
		}
	}
}

// BenchmarkMutes compares hasEqual performance with many source alerts cached.
func BenchmarkMutes(b *testing.B) {
	for _, numSources := range []int{10, 100, 1000} {
		b.Run(fmt.Sprintf("sources=%d", numSources), func(b *testing.B) {
			rule := newTestRule("severity", "instance")

			// Populate source alerts with varying "instance" labels.
			for i := 0; i < numSources; i++ {
				a := firingAlert(label.LabelSet{
					"alertname": "source",
					"severity":  "critical",
					"instance":  fmt.Sprintf("host%d", i),
				})
				_ = rule.scache.Set(a)
				rule.updateIndex(a)
			}

			target := label.LabelSet{
				"alertname": "target",
				"severity":  "critical",
				"instance":  fmt.Sprintf("host%d", numSources-1),
			}

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				rule.hasEqual(target, false)
			}
		})
	}
}
