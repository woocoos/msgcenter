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
)

func newTestRule(equal ...label.LabelName) *InhibitRule {
	rule := &InhibitRule{
		SourceMatchers: label.Matchers{
			&label.Matcher{Type: label.MatchEqual, Name: "alertname", Value: "source"},
		},
		TargetMatchers: label.Matchers{
			&label.Matcher{Type: label.MatchEqual, Name: "alertname", Value: "target"},
		},
		Equal: make(map[label.LabelName]struct{}),
	}
	for _, ln := range equal {
		rule.Equal[ln] = struct{}{}
	}
	rule.cache = newCache(rule.Equal)
	return rule
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
	c := newCache(map[label.LabelName]struct{}{
		"severity": {},
		"instance": {},
	})

	ls1 := label.LabelSet{"alertname": "src", "severity": "critical", "instance": "host1", "extra": "x"}
	ls2 := label.LabelSet{"alertname": "other", "severity": "critical", "instance": "host1", "extra": "y"}
	ls3 := label.LabelSet{"alertname": "src", "severity": "warning", "instance": "host1"}

	// 相同的 Equal 标签值 → 相同的 fingerprint, 与其余标签无关.
	assert.Equal(t, c.fingerprintEquals(ls1), c.fingerprintEquals(ls2))
	// 不同的 Equal 标签值 → 不同的 fingerprint.
	assert.NotEqual(t, c.fingerprintEquals(ls1), c.fingerprintEquals(ls3))
}

func TestCache_setAndFind(t *testing.T) {
	t.Parallel()
	rule := newTestRule("severity")

	src := firingAlert(label.LabelSet{"alertname": "source", "severity": "critical"})
	rule.cache.set(src)

	now := time.Now()
	// 通过 index 查找 source alert.
	target := label.LabelSet{"alertname": "target", "severity": "critical"}
	fp, ok := rule.cache.find(target, now, func(_ *alert.Alert) bool { return true })
	require.True(t, ok)
	assert.Equal(t, src.Fingerprint(), fp)

	// 不匹配的 Equal 标签不应找到.
	other := label.LabelSet{"alertname": "target", "severity": "warning"}
	_, ok = rule.cache.find(other, now, func(_ *alert.Alert) bool { return true })
	assert.False(t, ok)
}

func TestCache_multipleSourceAlertsSameEqualLabels(t *testing.T) {
	t.Parallel()
	rule := newTestRule("severity")

	// 两个 source alert 共享相同的 Equal 标签.
	src1 := firingAlert(label.LabelSet{"alertname": "source", "severity": "critical", "instance": "host1"})
	src2 := firingAlert(label.LabelSet{"alertname": "source", "severity": "critical", "instance": "host2"})
	rule.cache.set(src1)
	rule.cache.set(src2)

	now := time.Now()
	// 两个都应被索引.
	target := label.LabelSet{"alertname": "target", "severity": "critical"}
	fp, ok := rule.cache.find(target, now, func(_ *alert.Alert) bool { return true })
	require.True(t, ok)
	// 应找到其中一个.
	assert.True(t, fp == src1.Fingerprint() || fp == src2.Fingerprint())

	// 将第一个标记为 resolved.
	src1.EndsAt = now.Add(-time.Minute)
	rule.cache.set(src1)

	// GC 后应仍能找到第二个.
	rule.cache.gc()
	fp, ok = rule.cache.find(target, now, func(_ *alert.Alert) bool { return true })
	require.True(t, ok)
	assert.Equal(t, src2.Fingerprint(), fp)
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
			equal := map[label.LabelName]struct{}{}
			for _, ln := range c.equal {
				equal[ln] = struct{}{}
			}
			r := &InhibitRule{
				Equal: equal,
				cache: newCache(equal),
			}
			for _, a := range c.initial {
				r.cache.set(a)
			}

			_, have := r.hasEqual(c.input, false, now)
			require.Equal(t, c.result, have)
		})
	}
}

func TestCache_gcCleansIndex(t *testing.T) {
	t.Parallel()
	rule := newTestRule("severity")

	src := firingAlert(label.LabelSet{"alertname": "source", "severity": "critical"})
	rule.cache.set(src)
	assert.Len(t, rule.cache.alerts, 1)
	assert.Len(t, rule.cache.index, 1)

	// 标记为 resolved 后 GC.
	src.EndsAt = time.Now().Add(-time.Minute)
	rule.cache.set(src)
	rule.cache.gc()

	assert.Empty(t, rule.cache.alerts)
	assert.Empty(t, rule.cache.index)
}

func TestInhibitRuleIndexSurvivesGC(t *testing.T) {
	t.Parallel()
	now := time.Now()
	r := NewInhibitRule(profile.InhibitRule{Equal: []label.LabelName{"cluster"}})

	active := &alert.Alert{
		Labels:   label.LabelSet{"alertname": "S1", "cluster": "c1"},
		StartsAt: now.Add(-time.Hour),
		EndsAt:   now.Add(2 * time.Hour),
	}
	resolved := &alert.Alert{
		Labels:   label.LabelSet{"alertname": "S2", "cluster": "c1"},
		StartsAt: now.Add(-time.Hour),
		EndsAt:   now.Add(-time.Minute),
	}
	r.cache.set(active)
	r.cache.set(resolved)

	target := label.LabelSet{"alertname": "T", "cluster": "c1"}
	fp, ok := r.hasEqual(target, false, now)
	require.True(t, ok)
	require.Equal(t, active.Fingerprint(), fp)

	r.cache.gc()
	assert.Len(t, r.cache.alerts, 1)
	assert.Contains(t, r.cache.alerts, active.Fingerprint())

	fp, ok = r.hasEqual(target, false, now)
	require.True(t, ok, "active source alert must still inhibit after GC of a sibling")
	require.Equal(t, active.Fingerprint(), fp)
	assert.Len(t, r.cache.index, 1)

	active.EndsAt = now.Add(-time.Second)
	r.cache.set(active)
	r.cache.gc()
	_, ok = r.hasEqual(target, false, now)
	assert.False(t, ok)
	assert.Empty(t, r.cache.alerts)
	assert.Empty(t, r.cache.index, "empty index keys must be removed")
}

func TestInhibitRuleTwoSidedDoesNotShadow(t *testing.T) {
	t.Parallel()
	now := time.Now()
	r := NewInhibitRule(profile.InhibitRule{
		TargetMatchers: label.Matchers{&label.Matcher{Type: label.MatchEqual, Name: "severity", Value: "warning"}},
		Equal:          []label.LabelName{"cluster"},
	})

	sourceOnly := &alert.Alert{
		Labels:   label.LabelSet{"alertname": "S1", "cluster": "c1", "severity": "critical"},
		StartsAt: now.Add(-time.Hour),
		EndsAt:   now.Add(time.Hour),
	}
	twoSided := &alert.Alert{
		Labels:   label.LabelSet{"alertname": "S2", "cluster": "c1", "severity": "warning"},
		StartsAt: now.Add(-time.Hour),
		EndsAt:   now.Add(2 * time.Hour),
	}
	r.cache.set(sourceOnly)
	r.cache.set(twoSided)

	target := label.LabelSet{"alertname": "T", "cluster": "c1", "severity": "warning"}
	fp, ok := r.hasEqual(target, true, now)
	require.True(t, ok)
	require.Equal(t, sourceOnly.Fingerprint(), fp)
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

	ih.rules[0].cache.set(sourceAlert1)
	ih.rules[1].cache.set(sourceAlert2)

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

	ih.rules[0].cache.set(sourceAlert1)
	ih.rules[1].cache.set(sourceAlert2)

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

func (f *fakeAlerts) Start(context.Context) error                        { return nil }
func (f *fakeAlerts) Stop(context.Context) error                         { return nil }
func (f *fakeAlerts) GetPending() provider.AlertIterator                 { return nil }
func (f *fakeAlerts) Get(label.Fingerprint) (*alert.Alert, error)        { return nil, nil }
func (f *fakeAlerts) Put(context.Context, ...*alert.Alert) error         { return nil }
func (f *fakeAlerts) Subscribe(name string) provider.AlertIterator       { return nil }
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
				rule.cache.set(a)
			}

			target := label.LabelSet{
				"alertname": "target",
				"severity":  "critical",
				"instance":  fmt.Sprintf("host%d", numSources-1),
			}

			b.ResetTimer()
			b.ReportAllocs()
			now := time.Now()
			for i := 0; i < b.N; i++ {
				rule.hasEqual(target, false, now)
			}
		})
	}
}
