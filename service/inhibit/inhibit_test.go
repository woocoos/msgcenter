package inhibit

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
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

func resolvedAlert(labels label.LabelSet) *alert.Alert {
	return &alert.Alert{
		Labels:    labels,
		StartsAt:  time.Now().Add(-2 * time.Hour),
		EndsAt:    time.Now().Add(-time.Hour),
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

func TestInhibitRule_hasEqual(t *testing.T) {
	tests := []struct {
		name                 string
		rule                 *InhibitRule
		setupAlerts          []*alert.Alert
		target               label.LabelSet
		excludeTwoSidedMatch bool
		expectFound          bool
	}{
		{
			name:        "no source alerts",
			rule:        newTestRule("severity"),
			target:      label.LabelSet{"alertname": "target", "severity": "critical"},
			expectFound: false,
		},
		{
			name: "matching source alert",
			rule: newTestRule("severity"),
			setupAlerts: []*alert.Alert{
				firingAlert(label.LabelSet{"alertname": "source", "severity": "critical"}),
			},
			target:      label.LabelSet{"alertname": "target", "severity": "critical"},
			expectFound: true,
		},
		{
			name: "resolved source alert not found",
			rule: newTestRule("severity"),
			setupAlerts: []*alert.Alert{
				resolvedAlert(label.LabelSet{"alertname": "source", "severity": "critical"}),
			},
			target:      label.LabelSet{"alertname": "target", "severity": "critical"},
			expectFound: false,
		},
		{
			name: "non-matching equal labels",
			rule: newTestRule("severity"),
			setupAlerts: []*alert.Alert{
				firingAlert(label.LabelSet{"alertname": "source", "severity": "warning"}),
			},
			target:      label.LabelSet{"alertname": "target", "severity": "critical"},
			expectFound: false,
		},
		{
			name: "exclude two-sided match",
			rule: func() *InhibitRule {
				// Both source and target matchers match severity=critical,
				// so the source alert also matches target matchers.
				r := &InhibitRule{
					SourceMatchers: label.Matchers{
						&label.Matcher{Type: label.MatchEqual, Name: "severity", Value: "critical"},
					},
					TargetMatchers: label.Matchers{
						&label.Matcher{Type: label.MatchEqual, Name: "severity", Value: "critical"},
					},
					Equal:  map[label.LabelName]struct{}{"severity": {}},
					scache: newTestStore(),
					sindex: newIndex(),
				}
				r.scache.SetGCCallback(r.gcCallback)
				return r
			}(),
			setupAlerts: []*alert.Alert{
				firingAlert(label.LabelSet{"alertname": "source", "severity": "critical"}),
			},
			target:               label.LabelSet{"alertname": "source", "severity": "critical"},
			excludeTwoSidedMatch: true,
			expectFound:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, a := range tt.setupAlerts {
				require.NoError(t, tt.rule.scache.Set(a))
				tt.rule.updateIndex(a)
			}

			fp, found := tt.rule.hasEqual(tt.target, tt.excludeTwoSidedMatch)
			assert.Equal(t, tt.expectFound, found)
			if found {
				assert.NotZero(t, fp)
			}
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
