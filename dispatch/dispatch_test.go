// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package dispatch

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/marker"
	"github.com/woocoos/msgcenter/pkg/metrics"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service/provider"
)

const testMaintenanceInterval = 30 * time.Second

func TestAggrGroup(t *testing.T) {
	lset := label.LabelSet{"a": "v1", "b": "v2"}
	opts := &RouteOpts{
		Receiver: "n1",
		GroupBy: map[label.LabelName]struct{}{
			"a": {},
			"b": {},
		},
		GroupWait:      1 * time.Second,
		GroupInterval:  300 * time.Millisecond,
		RepeatInterval: 1 * time.Hour,
	}
	route := &Route{RouteOpts: *opts}

	a1 := &alert.Alert{
		Labels:   label.LabelSet{"a": "v1", "b": "v2", "c": "v3"},
		StartsAt: time.Now().Add(time.Minute),
		EndsAt:   time.Now().Add(time.Hour),
	}
	a2 := &alert.Alert{
		Labels:   label.LabelSet{"a": "v1", "b": "v2", "c": "v4"},
		StartsAt: time.Now().Add(-time.Hour),
		EndsAt:   time.Now().Add(2 * time.Hour),
	}
	a3 := &alert.Alert{
		Labels:   label.LabelSet{"a": "v1", "b": "v2", "c": "v5"},
		StartsAt: time.Now().Add(time.Minute),
		EndsAt:   time.Now().Add(5 * time.Minute),
	}

	type batch struct {
		alerts alert.Alerts
		now    time.Time
	}
	batchCh := make(chan batch)

	ntfy := func(ctx context.Context, alerts ...*alert.Alert) bool {
		now, ok := notify.Now(ctx)
		if !ok {
			t.Error("now missing from context")
		}
		if _, ok := notify.GroupKey(ctx); !ok {
			t.Error("group key missing from context")
		}
		if lbls, ok := notify.GroupLabels(ctx); !ok || !reflect.DeepEqual(lbls, lset) {
			t.Errorf("wrong group labels: %q", lbls)
		}
		if rcv, ok := notify.ReceiverName(ctx); !ok || rcv != opts.Receiver {
			t.Errorf("wrong receiver: %q", rcv)
		}
		if ri, ok := notify.RepeatInterval(ctx); !ok || ri != opts.RepeatInterval {
			t.Errorf("wrong repeat interval: %q", ri)
		}
		batchCh <- batch{alerts: alerts, now: now}
		return true
	}

	removeEndsAt := func(as alert.Alerts) alert.Alerts {
		for i, a := range as {
			ac := *a
			ac.EndsAt = time.Time{}
			as[i] = &ac
		}
		return as
	}

	receiveBatch := func(t *testing.T, since time.Time, minWait time.Duration, want alert.Alerts) time.Time {
		t.Helper()
		select {
		case <-time.After(2 * minWait):
			t.Fatalf("expected new batch after %v but received none", minWait)
		case b := <-batchCh:
			if got := b.now.Sub(since); got < minWait {
				t.Fatalf("received batch too early after %v (want >= %v)", got, minWait)
			}
			// Sort both slices by fingerprint for deterministic comparison.
			sort.Slice(b.alerts, func(i, j int) bool {
				return b.alerts[i].Fingerprint() < b.alerts[j].Fingerprint()
			})
			sort.Slice(want, func(i, j int) bool {
				return want[i].Fingerprint() < want[j].Fingerprint()
			})
			if !reflect.DeepEqual(b.alerts, want) {
				t.Fatalf("expected alerts %v but got %v", want, b.alerts)
			}
			return b.now
		}
		return time.Time{}
	}

	// Test regular situation: wait for group_wait to send out alerts.
	createdAt := time.Now()
	ag := newAggrGroup(context.Background(), lset, route, nil)
	go ag.run(ntfy)

	ctx := context.Background()
	ag.insert(ctx, a1)

	last := receiveBatch(t, createdAt, opts.GroupWait, removeEndsAt(alert.Alerts{a1}))

	for range 3 {
		ag.insert(ctx, a3)
		last = receiveBatch(t, last, opts.GroupInterval, removeEndsAt(alert.Alerts{a1, a3}))
	}

	ag.stop()

	// Test resolved alerts cleanup.
	createdAt = time.Now()
	ag = newAggrGroup(context.Background(), lset, route, nil)
	go ag.run(ntfy)

	ag.insert(ctx, a1)
	ag.insert(ctx, a2)

	last = receiveBatch(t, createdAt, opts.GroupWait, removeEndsAt(alert.Alerts{a1, a2}))

	for range 3 {
		ag.insert(ctx, a3)
		last = receiveBatch(t, last, opts.GroupInterval, removeEndsAt(alert.Alerts{a1, a2, a3}))
	}

	// Resolve an alert, it should be removed after the next batch.
	a1r := *a1
	a1r.EndsAt = time.Now()
	ag.insert(ctx, &a1r)
	last = receiveBatch(t, last, opts.GroupInterval, append(alert.Alerts{&a1r}, removeEndsAt(alert.Alerts{a2, a3})...))

	// Resolve all remaining alerts.
	a2r, a3r := *a2, *a3
	resolved := alert.Alerts{&a2r, &a3r}
	for _, a := range resolved {
		a.EndsAt = time.Now()
		ag.insert(ctx, a)
	}
	receiveBatch(t, last, opts.GroupInterval, resolved)

	require.Eventually(t, ag.empty, time.Second, 10*time.Millisecond,
		"Expected aggregation group to be empty after resolving alerts")

	ag.stop()
}

func TestGroupLabels(t *testing.T) {
	t.Parallel()
	a := &alert.Alert{
		Labels: label.LabelSet{"a": "v1", "b": "v2", "c": "v3"},
	}
	route := &Route{
		RouteOpts: RouteOpts{
			GroupBy: map[label.LabelName]struct{}{
				"a": {},
				"b": {},
			},
		},
	}
	expLs := label.LabelSet{"a": "v1", "b": "v2"}
	ls := getGroupLabels(a, route)
	assert.Equal(t, expLs, ls)
}

func TestGroupByAllLabels(t *testing.T) {
	t.Parallel()
	a := &alert.Alert{
		Labels: label.LabelSet{"a": "v1", "b": "v2", "c": "v3"},
	}
	route := &Route{
		RouteOpts: RouteOpts{GroupByAll: true},
	}
	ls := getGroupLabels(a, route)
	assert.Equal(t, a.Labels, ls)
}

func TestGroups(t *testing.T) {
	confData := `
route:
  receiver: 'prod'
  groupBy: ['alertname']
  groupWait: 10ms
  groupInterval: 10ms
  routes:
    - matchers: ['env="testing"']
      receiver: 'testing'
      groupBy: ['alertname', 'service']
    - matchers: ['env="prod"']
      receiver: 'prod'
      groupBy: ['alertname', 'service', 'cluster']
      continue: true
    - matchers: ['kafka="yes"']
      receiver: 'kafka'
      groupBy: ['alertname', 'service', 'cluster']
receivers:
  - name: 'prod'
  - name: 'testing'
  - name: 'kafka'
`
	cfg, err := profile.Load([]byte(confData))
	require.NoError(t, err)

	route := NewRoute(cfg.Route, nil)
	mk := marker.NewGroupMarker()
	alerts := newMockAlerts()

	timeout := func(d time.Duration) time.Duration { return 0 }
	recorder := newRecordStage()
	dispatcher := NewDispatcher(alerts, route, recorder, mk, timeout, testMaintenanceInterval, nil, nil)
	go dispatcher.Run(time.Now())
	defer dispatcher.Stop()

	inputAlerts := []*alert.Alert{
		// Matches the parent route.
		newTestAlert(label.LabelSet{"alertname": "OtherAlert", "cluster": "cc", "service": "dd"}),
		// Matches the first sub-route.
		newTestAlert(label.LabelSet{"env": "testing", "alertname": "TestingAlert", "service": "api", "instance": "inst1"}),
		// Matches the second sub-route.
		newTestAlert(label.LabelSet{"env": "prod", "alertname": "HighErrorRate", "cluster": "aa", "service": "api", "instance": "inst1"}),
		newTestAlert(label.LabelSet{"env": "prod", "alertname": "HighErrorRate", "cluster": "aa", "service": "api", "instance": "inst2"}),
		// Matches the second sub-route.
		newTestAlert(label.LabelSet{"env": "prod", "alertname": "HighErrorRate", "cluster": "bb", "service": "api", "instance": "inst1"}),
		// Matches the second and third sub-route.
		newTestAlert(label.LabelSet{"env": "prod", "alertname": "HighLatency", "cluster": "bb", "service": "db", "kafka": "yes", "instance": "inst3"}),
	}
	for _, a := range inputAlerts {
		require.NoError(t, alerts.Put(context.Background(), a))
	}

	// Let alerts get processed.
	require.Eventually(t, func() bool {
		return len(recorder.Alerts()) == 7
	}, 5*time.Second, 200*time.Millisecond, "not all alerts were notified")

	alertGroups, receivers := dispatcher.Groups(func(*Route) bool { return true }, func(*alert.Alert, time.Time) bool { return true })

	require.Len(t, alertGroups, 6)
	require.Equal(t, map[label.Fingerprint][]string{
		inputAlerts[0].Fingerprint(): {"prod"},
		inputAlerts[1].Fingerprint(): {"testing"},
		inputAlerts[2].Fingerprint(): {"prod"},
		inputAlerts[3].Fingerprint(): {"prod"},
		inputAlerts[4].Fingerprint(): {"prod"},
		inputAlerts[5].Fingerprint(): {"kafka", "prod"},
	}, receivers)

	// Verify group labels and receivers.
	for _, ag := range alertGroups {
		switch {
		case ag.Labels.Equal(label.LabelSet{"alertname": "OtherAlert"}):
			require.Equal(t, "prod", ag.Receiver)
			require.Len(t, ag.Alerts, 1)
		case ag.Labels.Equal(label.LabelSet{"alertname": "TestingAlert", "service": "api"}):
			require.Equal(t, "testing", ag.Receiver)
			require.Len(t, ag.Alerts, 1)
		case ag.Labels.Equal(label.LabelSet{"alertname": "HighErrorRate", "service": "api", "cluster": "aa"}):
			require.Equal(t, "prod", ag.Receiver)
			require.Len(t, ag.Alerts, 2)
		case ag.Labels.Equal(label.LabelSet{"alertname": "HighErrorRate", "service": "api", "cluster": "bb"}):
			require.Equal(t, "prod", ag.Receiver)
			require.Len(t, ag.Alerts, 1)
		case ag.Labels.Equal(label.LabelSet{"alertname": "HighLatency", "service": "db", "cluster": "bb"}):
			// Two groups with same labels: one for kafka, one for prod.
			require.Contains(t, []string{"kafka", "prod"}, ag.Receiver)
			require.Len(t, ag.Alerts, 1)
		default:
			t.Fatalf("unexpected group: labels=%v receiver=%s", ag.Labels, ag.Receiver)
		}
	}
}

func TestGroupsWithLimits(t *testing.T) {
	confData := `
route:
  receiver: 'prod'
  groupBy: ['alertname']
  groupWait: 10ms
  groupInterval: 10ms
  routes:
    - matchers: ['env="testing"']
      receiver: 'testing'
      groupBy: ['alertname', 'service']
    - matchers: ['env="prod"']
      receiver: 'prod'
      groupBy: ['alertname', 'service', 'cluster']
      continue: true
    - matchers: ['kafka="yes"']
      receiver: 'kafka'
      groupBy: ['alertname', 'service', 'cluster']
receivers:
  - name: 'prod'
  - name: 'testing'
  - name: 'kafka'
`
	cfg, err := profile.Load([]byte(confData))
	require.NoError(t, err)

	route := NewRoute(cfg.Route, nil)
	reg := prometheus.NewRegistry()
	mk := marker.NewGroupMarker()
	alerts := newMockAlerts()

	timeout := func(d time.Duration) time.Duration { return 0 }
	recorder := newRecordStage()
	lim := testLimits{groups: 6}
	m := metrics.NewDispatcherMetrics(true, reg)
	dispatcher := NewDispatcher(alerts, route, recorder, mk, timeout, testMaintenanceInterval, lim, m)
	go dispatcher.Run(time.Now())
	defer dispatcher.Stop()

	inputAlerts := []*alert.Alert{
		newTestAlert(label.LabelSet{"alertname": "OtherAlert", "cluster": "cc", "service": "dd"}),
		newTestAlert(label.LabelSet{"env": "testing", "alertname": "TestingAlert", "service": "api", "instance": "inst1"}),
		newTestAlert(label.LabelSet{"env": "prod", "alertname": "HighErrorRate", "cluster": "aa", "service": "api", "instance": "inst1"}),
		newTestAlert(label.LabelSet{"env": "prod", "alertname": "HighErrorRate", "cluster": "aa", "service": "api", "instance": "inst2"}),
		newTestAlert(label.LabelSet{"env": "prod", "alertname": "HighErrorRate", "cluster": "bb", "service": "api", "instance": "inst1"}),
		newTestAlert(label.LabelSet{"env": "prod", "alertname": "HighLatency", "cluster": "bb", "service": "db", "kafka": "yes", "instance": "inst3"}),
	}
	for _, a := range inputAlerts {
		require.NoError(t, alerts.Put(context.Background(), a))
	}

	require.Eventually(t, func() bool {
		return len(recorder.Alerts()) == 7
	}, 5*time.Second, 200*time.Millisecond)

	routeFilter := func(*Route) bool { return true }
	alertFilter := func(*alert.Alert, time.Time) bool { return true }

	alertGroups, _ := dispatcher.Groups(routeFilter, alertFilter)
	require.Len(t, alertGroups, 6)

	require.Equal(t, 0.0, testutil.ToFloat64(m.AggrGroupLimitReached))

	// Try to store a new alert that would create a 7th group, hitting the limit.
	require.NoError(t, alerts.Put(context.Background(),
		newTestAlert(label.LabelSet{"env": "prod", "alertname": "NewAlert", "cluster": "new-cluster", "service": "db"})))

	require.Eventually(t, func() bool {
		return testutil.ToFloat64(m.AggrGroupLimitReached) > 0
	}, 5*time.Second, 200*time.Millisecond)

	// Verify there are still only 6 groups.
	alertGroups, _ = dispatcher.Groups(routeFilter, alertFilter)
	require.Len(t, alertGroups, 6)
}

// recordStage implements notify.Stage for testing. It records all alerts
// dispatched by the Dispatcher, keyed by group key.
type recordStage struct {
	mtx    sync.RWMutex
	alerts map[string]map[label.Fingerprint]*alert.Alert
}

func newRecordStage() *recordStage {
	return &recordStage{alerts: make(map[string]map[label.Fingerprint]*alert.Alert)}
}

func (r *recordStage) Alerts() []*alert.Alert {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	var result []*alert.Alert
	for _, m := range r.alerts {
		for _, a := range m {
			result = append(result, a)
		}
	}
	return result
}

func (r *recordStage) Exec(ctx context.Context, alerts ...*alert.Alert) (context.Context, []*alert.Alert, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	gk, ok := notify.GroupKey(ctx)
	if !ok {
		panic("GroupKey not present in context")
	}
	if _, ok := r.alerts[gk]; !ok {
		r.alerts[gk] = make(map[label.Fingerprint]*alert.Alert)
	}
	for _, a := range alerts {
		r.alerts[gk][a.Fingerprint()] = a
	}
	return ctx, alerts, nil
}

type testLimits struct {
	groups int
}

func (l testLimits) MaxNumberOfAggregationGroups() int { return l.groups }

// mockAlerts implements provider.Alerts for testing.
type mockAlerts struct {
	ch   chan *provider.Alert
	done chan struct{}
}

func newMockAlerts() *mockAlerts {
	return &mockAlerts{
		ch:   make(chan *provider.Alert, 1000),
		done: make(chan struct{}),
	}
}

func (m *mockAlerts) Subscribe(name string) provider.AlertIterator {
	return provider.NewAlertIterator(m.ch, m.done, nil)
}

func (m *mockAlerts) SlurpAndSubscribe(name string) ([]*alert.Alert, provider.AlertIterator) {
	return nil, provider.NewAlertIterator(m.ch, m.done, nil)
}

func (m *mockAlerts) GetPending() provider.AlertIterator {
	return provider.NewAlertIterator(make(chan *provider.Alert), make(chan struct{}), nil)
}
func (m *mockAlerts) Get(fp label.Fingerprint) (*alert.Alert, error) { return nil, nil }

func (m *mockAlerts) Put(ctx context.Context, alerts ...*alert.Alert) error {
	for _, a := range alerts {
		m.ch <- &provider.Alert{Data: a, Header: map[string]string{}}
	}
	return nil
}
func (m *mockAlerts) Start(ctx context.Context) error { return nil }
func (m *mockAlerts) Stop(ctx context.Context) error  { return nil }
func (m *mockAlerts) Name() string                    { return "mock" }

func (m *mockAlerts) Merge(b []byte) error { return nil }

var (
	// Set the start time in the past to trigger a flush immediately.
	t0 = time.Now().Add(-time.Minute)
	// Set the end time in the future to avoid deleting the alert.
	t1 = t0.Add(2 * time.Minute)
)

func newTestAlert(labels label.LabelSet) *alert.Alert {
	return &alert.Alert{
		Labels:       labels,
		Annotations:  label.LabelSet{"foo": "bar"},
		StartsAt:     t0,
		EndsAt:       t1,
		GeneratorURL: "http://example.com/prometheus",
		UpdatedAt:    t0,
	}
}

func TestDispatcherRace(t *testing.T) {
	t.Parallel()
	alerts := newMockAlerts()
	route := &Route{}
	d := NewDispatcher(alerts, route, nil, nil,
		func(d time.Duration) time.Duration { return d },
		testMaintenanceInterval, nil, nil)

	go d.Run(time.Now())
	d.Stop()
}

// --- TestDispatcherRaceOnFirstAlertNotDeliveredWhenGroupWaitIsZero ---
// Regression test: with GroupWait=0, all alerts should be notified immediately.

func TestDispatcherRaceOnFirstAlertNotDeliveredWhenGroupWaitIsZero(t *testing.T) {
	const numAlerts = 5000

	alerts := newMockAlerts()
	gw := time.Duration(0)
	gi := time.Hour
	ri := time.Hour
	route := NewRoute(&profile.Route{
		Receiver:       "default",
		GroupBy:        []label.LabelName{"alertname"},
		GroupWait:      &gw,
		GroupInterval:  &gi,
		RepeatInterval: &ri,
	}, nil)

	recorder := newRecordStage()
	d := NewDispatcher(alerts, route, recorder, nil,
		func(d time.Duration) time.Duration { return d },
		testMaintenanceInterval, nil, nil)

	go d.Run(time.Now())
	defer d.Stop()

	// Push all alerts.
	for i := range numAlerts {
		a := newTestAlert(label.LabelSet{"alertname": fmt.Sprintf("Alert_%d", i)})
		require.NoError(t, alerts.Put(context.Background(), a))
	}

	// Wait until all alerts have been notified.
	require.Eventually(t, func() bool {
		return len(recorder.Alerts()) >= numAlerts
	}, 5*time.Second, 10*time.Millisecond, "not all alerts were notified")

	require.Len(t, recorder.Alerts(), numAlerts)
}

func TestDispatcher_DoMaintenance(t *testing.T) {
	t.Parallel()
	alerts := newMockAlerts()
	route := &Route{
		RouteOpts: RouteOpts{
			GroupBy:       map[label.LabelName]struct{}{"alertname": {}},
			GroupWait:     0,
			GroupInterval: 5 * time.Minute,
		},
		Idx: 0,
	}
	timeout := func(d time.Duration) time.Duration { return d }
	recorder := newRecordStage()
	mk := marker.NewGroupMarker()

	d := NewDispatcher(alerts, route, recorder, mk, timeout, testMaintenanceInterval, nil, nil)
	// Manually create routeGroupsSlice since we are not calling Run().
	d.routeGroupsSlice = make([]routeAggrGroups, route.Idx+1)
	d.routeGroupsSlice[route.Idx] = routeAggrGroups{route: route}

	// Insert an aggregation group with one resolved alert.
	labels := label.LabelSet{"alertname": "1"}
	ag := newAggrGroup(context.Background(), labels, route, timeout)
	d.routeGroupsSlice[route.Idx].groups.Store(ag.fingerprint(), ag)

	// Add a resolved alert.
	resolvedAlert := &alert.Alert{
		Labels:    labels,
		StartsAt:  time.Now().Add(-2 * time.Hour),
		EndsAt:    time.Now().Add(-1 * time.Hour),
		UpdatedAt: time.Now().Add(-1 * time.Hour),
	}
	ag.alerts.Set(resolvedAlert)

	// Flush will detect the resolved alert and delete it.
	notified := false
	ag.flush(func(alerts ...*alert.Alert) bool {
		require.Len(t, alerts, 1)
		require.Equal(t, labels, alerts[0].Labels)
		notified = true
		return true
	})
	require.True(t, notified, "flush should have called notify function")

	// Insert a marker for the aggregation group's group key.
	mk.SetMuted(route.ID(), ag.GroupKey(), []string{"weekends"})
	mutedBy, isMuted := mk.Muted(route.ID(), ag.GroupKey())
	require.True(t, isMuted)
	require.Equal(t, []string{"weekends"}, mutedBy)

	// Must run otherwise doMaintenance blocks on ag.stop().
	go ag.run(func(context.Context, ...*alert.Alert) bool { return true })

	// Run maintenance — the destroyed group should be cleaned up.
	d.doMaintenance()

	// The group should have been removed from the map.
	_, ok := d.routeGroupsSlice[route.Idx].groups.Load(ag.fingerprint())
	assert.False(t, ok, "destroyed group should have been removed by maintenance")

	// The marker should have been cleaned up.
	mutedBy, isMuted = mk.Muted(route.ID(), ag.GroupKey())
	assert.False(t, isMuted, "marker should have been cleaned up by maintenance")
	assert.Empty(t, mutedBy)
}

func TestGroupAlert_RecoversWhenCASFails(t *testing.T) {
	const (
		alertsPerRound = 200
		maxRounds      = 50
	)

	alerts := newMockAlerts()
	route := &Route{
		RouteOpts: RouteOpts{
			Receiver:       "test",
			GroupBy:        map[label.LabelName]struct{}{"alertname": {}},
			GroupWait:      time.Hour,
			GroupInterval:  time.Hour,
			RepeatInterval: time.Hour,
		},
		Idx: 0,
	}
	timeout := func(d time.Duration) time.Duration { return d }
	recorder := newRecordStage()
	reg := prometheus.NewRegistry()
	m := metrics.NewDispatcherMetrics(false, reg)

	d := NewDispatcher(alerts, route, recorder, nil, timeout, testMaintenanceInterval, nil, m)
	d.routeGroupsSlice = []routeAggrGroups{{route: route}}
	d.state.Store(dispatcherStateWaitingToStart)

	rounds := 0
	for rounds < maxRounds {
		groupLabels := label.LabelSet{"alertname": fmt.Sprintf("shared-%d", rounds)}
		destroyedAg := newAggrGroup(context.Background(), groupLabels, route, timeout)
		// Mark destroyed via DeleteIfNotModified with empty slice + destroyIfEmpty=true.
		require.NoError(t, destroyedAg.alerts.DeleteIfNotModified(alert.Alerts{}, true))
		require.True(t, destroyedAg.destroyed())

		fp := destroyedAg.fingerprint()
		d.routeGroupsSlice[0].groups.Store(fp, destroyedAg)

		var ready, done sync.WaitGroup
		ready.Add(alertsPerRound)
		done.Add(alertsPerRound)
		start := make(chan struct{})
		for i := range alertsPerRound {
			go func(idx int) {
				defer done.Done()
				a := newTestAlert(label.LabelSet{
					"alertname": string(groupLabels["alertname"]),
					"instance":  fmt.Sprintf("inst-%d", idx),
				})
				ready.Done()
				<-start
				d.groupAlert(context.Background(), a, route)
			}(i)
		}
		ready.Wait()
		close(start)
		done.Wait()

		el, ok := d.routeGroupsSlice[0].groups.Load(fp)
		require.True(t, ok, "round %d: a live group must occupy the fp", rounds)
		finalAg := el.(*aggrGroup)
		require.False(t, finalAg.destroyed(), "round %d: final group must not be destroyed", rounds)
		require.NotSame(t, destroyedAg, finalAg, "round %d: destroyed group must have been replaced", rounds)
		require.Len(t, finalAg.alerts.List(), alertsPerRound, "round %d: all alerts must land in the final group", rounds)
		rounds++

		// If we've observed retries, we've exercised the contended path.
		if m.AggrGroupCreationRetries != nil {
			break
		}
	}

	// With GOMAXPROCS=1 the contention can't be fully exercised.
	if runtime.GOMAXPROCS(0) == 1 {
		return
	}
}

func TestGroupAlert_DisplacedAggrGroupGoroutineExits(t *testing.T) {
	t.Parallel()
	alerts := newMockAlerts()
	route := &Route{
		RouteOpts: RouteOpts{
			Receiver:       "test",
			GroupBy:        map[label.LabelName]struct{}{"alertname": {}},
			GroupWait:      time.Hour,
			GroupInterval:  time.Hour,
			RepeatInterval: time.Hour,
		},
		Idx: 0,
	}
	timeout := func(d time.Duration) time.Duration { return d }
	recorder := newRecordStage()

	d := NewDispatcher(alerts, route, recorder, nil, timeout, testMaintenanceInterval, nil, nil)
	d.routeGroupsSlice = []routeAggrGroups{{route: route}}
	d.state.Store(dispatcherStateWaitingToStart)

	groupLabels := label.LabelSet{"alertname": "displaced"}
	displaced := newAggrGroup(context.Background(), groupLabels, route, timeout)
	require.NoError(t, displaced.alerts.DeleteIfNotModified(alert.Alerts{}, true))
	require.True(t, displaced.destroyed())
	d.routeGroupsSlice[0].groups.Store(displaced.fingerprint(), displaced)

	// Start the run goroutine — this is the orphan candidate.
	go displaced.run(func(context.Context, ...*alert.Alert) bool { return true })

	// Trigger the CAS replacement.
	d.groupAlert(context.Background(), newTestAlert(groupLabels), route)

	// The displaced group should have been swapped out.
	el, ok := d.routeGroupsSlice[0].groups.Load(displaced.fingerprint())
	require.True(t, ok)
	require.NotSame(t, displaced, el.(*aggrGroup), "destroyed group must have been replaced")

	// Its run goroutine must have exited.
	select {
	case <-displaced.done:
	case <-time.After(2 * time.Second):
		t.Fatal("displaced aggrGroup.run goroutine did not exit after CAS replacement")
	}
}

func TestDispatcher_DeleteResolvedAlertsFromMarker(t *testing.T) {
	t.Run("successful flush deletes markers for resolved alerts", func(t *testing.T) {
		labels := label.LabelSet{"alertname": "TestAlert"}
		route := &Route{
			RouteOpts: RouteOpts{
				Receiver:       "test",
				GroupBy:        map[label.LabelName]struct{}{"alertname": {}},
				GroupWait:      0,
				GroupInterval:  time.Minute,
				RepeatInterval: time.Hour,
			},
		}
		timeout := func(d time.Duration) time.Duration { return d }

		ag := newAggrGroup(context.Background(), labels, route, timeout)

		now := time.Now()
		activeAlert := &alert.Alert{
			Labels:   label.LabelSet{"alertname": "TestAlert", "instance": "1"},
			StartsAt: now.Add(-time.Hour),
			EndsAt:   now.Add(time.Hour),
		}
		resolvedAlert := &alert.Alert{
			Labels:   label.LabelSet{"alertname": "TestAlert", "instance": "2"},
			StartsAt: now.Add(-time.Hour),
			EndsAt:   now.Add(-time.Minute),
		}

		ag.insert(context.Background(), activeAlert)
		ag.insert(context.Background(), resolvedAlert)

		ag.marker.SetSilenced(activeAlert.Fingerprint(), nil)
		ag.marker.SetSilenced(resolvedAlert.Fingerprint(), nil)

		require.Equal(t, alert.AlertStateActive, ag.marker.Status(activeAlert.Fingerprint()).State)
		require.Equal(t, alert.AlertStateActive, ag.marker.Status(resolvedAlert.Fingerprint()).State)

		ag.flush(func(alerts ...*alert.Alert) bool { return true })

		require.Equal(t, alert.AlertStateActive, ag.marker.Status(activeAlert.Fingerprint()).State,
			"active alert marker should still exist")
		require.Equal(t, alert.AlertStateUnprocessed, ag.marker.Status(resolvedAlert.Fingerprint()).State,
			"resolved alert marker should be deleted")
	})

	t.Run("failed flush does not delete markers", func(t *testing.T) {
		labels := label.LabelSet{"alertname": "TestAlert"}
		route := &Route{
			RouteOpts: RouteOpts{
				Receiver:       "test",
				GroupBy:        map[label.LabelName]struct{}{"alertname": {}},
				GroupWait:      0,
				GroupInterval:  time.Minute,
				RepeatInterval: time.Hour,
			},
		}
		timeout := func(d time.Duration) time.Duration { return d }

		ag := newAggrGroup(context.Background(), labels, route, timeout)

		now := time.Now()
		resolvedAlert := &alert.Alert{
			Labels:   label.LabelSet{"alertname": "TestAlert", "instance": "1"},
			StartsAt: now.Add(-time.Hour),
			EndsAt:   now.Add(-time.Minute),
		}

		ag.insert(context.Background(), resolvedAlert)
		ag.marker.SetSilenced(resolvedAlert.Fingerprint(), nil)

		require.Equal(t, alert.AlertStateActive, ag.marker.Status(resolvedAlert.Fingerprint()).State)

		ag.flush(func(alerts ...*alert.Alert) bool { return false })

		require.Equal(t, alert.AlertStateActive, ag.marker.Status(resolvedAlert.Fingerprint()).State,
			"marker should not be deleted when notify fails")
	})

	t.Run("markers not deleted when alert is modified during flush", func(t *testing.T) {
		labels := label.LabelSet{"alertname": "TestAlert"}
		route := &Route{
			RouteOpts: RouteOpts{
				Receiver:       "test",
				GroupBy:        map[label.LabelName]struct{}{"alertname": {}},
				GroupWait:      0,
				GroupInterval:  time.Minute,
				RepeatInterval: time.Hour,
			},
		}
		timeout := func(d time.Duration) time.Duration { return d }

		ag := newAggrGroup(context.Background(), labels, route, timeout)

		now := time.Now()
		resolvedAlert := &alert.Alert{
			Labels:    label.LabelSet{"alertname": "TestAlert", "instance": "1"},
			StartsAt:  now.Add(-time.Hour),
			EndsAt:    now.Add(-time.Minute),
			UpdatedAt: now,
		}

		ag.insert(context.Background(), resolvedAlert)
		ag.marker.SetSilenced(resolvedAlert.Fingerprint(), nil)

		require.Equal(t, alert.AlertStateActive, ag.marker.Status(resolvedAlert.Fingerprint()).State)

		ag.flush(func(alerts ...*alert.Alert) bool {
			modifiedAlert := &alert.Alert{
				Labels:    label.LabelSet{"alertname": "TestAlert", "instance": "1"},
				StartsAt:  now.Add(-time.Hour),
				EndsAt:    now.Add(time.Hour),
				UpdatedAt: now.Add(time.Second),
			}
			ag.alerts.Set(modifiedAlert)
			return true
		})

		require.Equal(t, alert.AlertStateActive, ag.marker.Status(resolvedAlert.Fingerprint()).State,
			"marker should not be deleted when alert is modified during flush")
	})
}

func TestDispatchOnStartup(t *testing.T) {
	alerts := newMockAlerts()
	route := &Route{
		RouteOpts: RouteOpts{
			Receiver:       "default",
			GroupBy:        map[label.LabelName]struct{}{"instance": {}},
			GroupWait:      1 * time.Second,
			GroupInterval:  3 * time.Minute,
			RepeatInterval: 1 * time.Hour,
		},
	}

	recorder := newRecordStage()
	timeout := func(d time.Duration) time.Duration { return d }

	startDelay := 2 * time.Second
	startTime := time.Now().Add(startDelay)

	d := NewDispatcher(alerts, route, recorder, nil, timeout, testMaintenanceInterval, nil, nil)
	go d.Run(startTime)
	defer d.Stop()

	now := time.Now()
	alert1 := &alert.Alert{
		Labels:       label.LabelSet{"alertname": "TestAlert1", "instance": "1"},
		Annotations:  label.LabelSet{"foo": "bar"},
		StartsAt:     now.Add(-time.Hour),
		EndsAt:       now.Add(time.Hour),
		GeneratorURL: "http://example.com/prometheus",
		UpdatedAt:    now,
	}

	require.NoError(t, alerts.Put(context.Background(), alert1))

	// Expect alert1 to be dispatched after startTime + GroupWait.
	require.Eventually(t, func() bool {
		return len(recorder.Alerts()) == 1
	}, startDelay+route.RouteOpts.GroupWait+time.Second, 500*time.Millisecond)

	require.Equal(t, alert1.Fingerprint(), recorder.Alerts()[0].Fingerprint())
}

func TestGetGroupLabels(t *testing.T) {
	t.Parallel()
	a := &alert.Alert{
		Labels: label.LabelSet{
			"alertname": "TestAlert",
			"job":       "prometheus",
			"instance":  "localhost:9090",
			"severity":  "critical",
		},
	}

	t.Run("specific labels", func(t *testing.T) {
		route := &Route{
			RouteOpts: RouteOpts{
				GroupBy: map[label.LabelName]struct{}{
					"alertname": {},
					"job":       {},
				},
			},
		}
		labels := getGroupLabels(a, route)
		require.Len(t, labels, 2)
		assert.Equal(t, "TestAlert", labels["alertname"])
		assert.Equal(t, "prometheus", labels["job"])
	})

	t.Run("group by all", func(t *testing.T) {
		route := &Route{
			RouteOpts: RouteOpts{GroupByAll: true},
		}
		labels := getGroupLabels(a, route)
		require.Len(t, labels, 4)
		assert.Equal(t, a.Labels, labels)
	})
}

func BenchmarkGetGroupLabels(b *testing.B) {
	now := time.Now()
	a := &alert.Alert{
		Labels: label.LabelSet{
			"alertname":  "TestAlert",
			"severity":   "critical",
			"job":        "prometheus",
			"instance":   "localhost:9090",
			"namespace":  "monitoring",
			"cluster":    "prod-us-east-1",
			"datacenter": "dc1",
			"env":        "production",
			"team":       "platform",
			"service":    "alertmanager",
		},
		StartsAt: now.Add(-time.Hour),
		EndsAt:   now.Add(time.Hour),
	}

	b.Run("specific_labels", func(b *testing.B) {
		route := &Route{
			RouteOpts: RouteOpts{
				GroupBy: map[label.LabelName]struct{}{
					"alertname": {},
					"job":       {},
					"severity":  {},
				},
			},
		}
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = getGroupLabels(a, route)
		}
	})

	b.Run("group_by_all", func(b *testing.B) {
		route := &Route{
			RouteOpts: RouteOpts{GroupByAll: true},
		}
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = getGroupLabels(a, route)
		}
	})
}
