// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package dispatch

import (
	"context"
	"errors"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/tsingsun/woocoo/pkg/gds"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/marker"
	"github.com/woocoos/msgcenter/pkg/metrics"
	"github.com/woocoos/msgcenter/pkg/tracing"
	"github.com/woocoos/msgcenter/service/provider"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var logger = log.Component("dispatch")

var tracer = tracing.NewTracer("github.com/woocoos/msgcenter/dispatch")

const (
	dispatcherStateUnknown = iota
	dispatcherStateWaitingToStart
	dispatcherStateRunning
	dispatcherStateStopped
)

// Dispatcher sorts incoming alerts into aggregation groups and
// assigns the correct notifiers to each.
type Dispatcher struct {
	route      *Route
	alerts     provider.Alerts
	stage      notify.Stage
	marker     marker.GroupMarker
	metrics    *metrics.DispatcherMetrics
	limits     Limits
	propagator propagation.TextMapPropagator

	timeout func(time.Duration) time.Duration

	loaded   chan struct{}
	finished sync.WaitGroup
	ctx      context.Context
	cancel   func()

	routeGroupsSlice []routeAggrGroups
	aggrGroupsNum    atomic.Int32

	maintenanceInterval time.Duration
	concurrency         int // Number of goroutines for alert ingestion

	startTimer *time.Timer
	state      atomic.Int32
}

// Limits describes limits used by Dispatcher.
type Limits interface {
	MaxNumberOfAggregationGroups() int
}

// NewDispatcher returns a new Dispatcher.
func NewDispatcher(
	alerts provider.Alerts,
	route *Route,
	stage notify.Stage,
	marker marker.GroupMarker,
	timeout func(time.Duration) time.Duration,
	maintenanceInterval time.Duration,
	limits Limits,
	metric *metrics.DispatcherMetrics,
) *Dispatcher {
	if limits == nil {
		limits = nilLimits{}
	}
	if metric == nil {
		metric = metrics.NewNoopDispatcherMetrics()
	}

	// Calculate concurrency for ingestion.
	concurrency := min(max(runtime.GOMAXPROCS(0)/2, 2), 8)

	disp := &Dispatcher{
		alerts:              alerts,
		stage:               stage,
		route:               route,
		timeout:             timeout,
		metrics:             metric,
		limits:              limits,
		concurrency:         concurrency,
		maintenanceInterval: gds.IIF(maintenanceInterval == 0, 30*time.Second, maintenanceInterval),
		propagator:          otel.GetTextMapPropagator(),
		marker:              marker,
	}

	disp.state.Store(dispatcherStateUnknown)
	disp.loaded = make(chan struct{})
	disp.ctx, disp.cancel = context.WithCancel(context.Background())
	return disp
}

// Run starts dispatching alerts incoming via the updates channel.
// dispatchStartTime controls when aggregation groups begin flushing.
// Alerts received before this time are collected into groups but not flushed
// until the timer fires. Pass time.Time{} to start immediately.
func (d *Dispatcher) Run(dispatchStartTime time.Time) {
	if !d.state.CompareAndSwap(dispatcherStateUnknown, dispatcherStateWaitingToStart) {
		return
	}
	d.finished.Add(1)
	defer d.finished.Done()

	logger.Debug("preparing to start", zap.Time("startTime", dispatchStartTime))
	d.startTimer = time.NewTimer(time.Until(dispatchStartTime))
	logger.Debug("setting state", zap.String("state", "waiting_to_start"))
	d.routeGroupsSlice = make([]routeAggrGroups, d.route.Idx+1)
	d.route.Walk(func(r *Route) {
		d.routeGroupsSlice[r.Idx] = routeAggrGroups{route: r}
	})

	d.aggrGroupsNum.Store(0)
	d.metrics.AggrGroups.Set(0)

	// Use SlurpAndSubscribe: process existing alerts synchronously before
	// entering the main loop, then switch to incremental mode.
	initialAlerts, it := d.alerts.SlurpAndSubscribe("dispatcher")
	for _, a := range initialAlerts {
		d.routeAlert(d.ctx, a)
	}
	close(d.loaded)

	d.run(it)
}

func (d *Dispatcher) run(it provider.AlertIterator) {
	defer it.Close()

	// Start maintenance goroutine.
	d.finished.Go(func() {
		ticker := time.NewTicker(d.maintenanceInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				d.doMaintenance()
			case <-d.ctx.Done():
				return
			}
		}
	})

	// Start timer goroutine: waits for startTimer, then transitions to Running
	// and starts all existing aggrGroups created during WaitingToStart.
	d.finished.Go(func() {
		<-d.startTimer.C
		if d.state.CompareAndSwap(dispatcherStateWaitingToStart, dispatcherStateRunning) {
			logger.Debug("started", zap.String("state", "running"))
			logger.Debug("Starting all existing aggregation groups")
			for i := range d.routeGroupsSlice {
				d.routeGroupsSlice[i].groups.Range(func(_, ag any) bool {
					d.runAG(ag.(*aggrGroup))
					return true
				})
			}
		}
	})

	// Start concurrent alert ingestion workers.
	alertCh := it.Next()
	for i := 0; i < d.concurrency; i++ {
		d.finished.Add(1)
		go func(workerID int) {
			defer d.finished.Done()
			for {
				select {
				case pAlert, ok := <-alertCh:
					if !ok {
						if err := it.Err(); err != nil {
							logger.Error("error on alert update", zap.Error(err), zap.Int("workerID", workerID))
						}
						return
					}
					if err := it.Err(); err != nil {
						logger.Error("error on alert update", zap.Error(err), zap.Int("workerID", workerID))
						continue
					}
					// Extract tracing context from header if present.
					ctx := d.ctx
					if pAlert.Header != nil {
						ctx = d.propagator.Extract(ctx, propagation.MapCarrier(pAlert.Header))
					}
					d.routeAlert(ctx, pAlert.Data)
				case <-d.ctx.Done():
					return
				}
			}
		}(i)
	}

	<-d.ctx.Done()
}

func (d *Dispatcher) routeAlert(ctx context.Context, a *alert.Alert) {
	ctx, span := tracer.Start(ctx, "dispatch.Dispatcher.routeAlert",
		trace.WithAttributes(
			attribute.String("alerting.alert.name", a.Name()),
			attribute.String("alerting.alert.fingerprint", a.Fingerprint().String()),
		),
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	now := time.Now()
	for _, r := range d.route.Match(a.Labels) {
		span.AddEvent("dispatching alert to route",
			trace.WithAttributes(
				attribute.String("alerting.route.receiver.name", r.RouteOpts.Receiver),
			),
		)
		d.groupAlert(ctx, a, r)
	}
	d.metrics.ProcessingDuration.Observe(time.Since(now).Seconds())
}

// groupAlert finds or creates the aggregation group for the alert and route,
// using lock-free CAS operations for concurrent safety.
func (d *Dispatcher) groupAlert(ctx context.Context, a *alert.Alert, route *Route) {
	_, span := tracer.Start(ctx, "dispatch.Dispatcher.groupAlert",
		trace.WithAttributes(
			attribute.String("alerting.alert.name", a.Name()),
			attribute.String("alerting.alert.fingerprint", a.Fingerprint().String()),
			attribute.String("alerting.route.receiver.name", route.RouteOpts.Receiver),
		),
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	now := time.Now()
	groupLabels := getGroupLabels(a, route)
	fp := groupLabels.Fingerprint()
	rg := &d.routeGroupsSlice[route.Idx]

	// Fast path: try to load existing group.
	el, loaded := rg.groups.Load(fp)
	if loaded {
		ag := el.(*aggrGroup)
		if ag.insert(ctx, a) {
			return
		}
	}

	// Check aggregation group limit.
	limit := d.limits.MaxNumberOfAggregationGroups()
	current := int(d.aggrGroupsNum.Load())
	if limit > 0 && current >= limit {
		d.metrics.AggrGroupLimitReached.Inc()
		err := errors.New("too many aggregation groups, cannot create new group for alert")
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err,
			trace.WithAttributes(
				attribute.Int("alerting.aggregation_group.count", current),
				attribute.Int("alerting.aggregation_group.limit", limit),
			),
		)
		logger.Error("too many aggregation groups",
			zap.Int("groups", current), zap.Int("limit", limit), zap.String("alert", a.Name()))
		return
	}

	// Create new aggrGroup.
	ag := newAggrGroup(d.ctx, groupLabels, route, d.timeout)
	ag.insert(ctx, a)

	// CAS retry loop for group creation/replacement.
	retries := 0
	for {
		if loaded {
			swapped := rg.groups.CompareAndSwap(fp, el, ag)
			if swapped {
				el.(*aggrGroup).cancel()
				break
			}
			loaded = false
		} else {
			el, loaded = rg.groups.LoadOrStore(fp, ag)
			if !loaded {
				rg.groupsLen.Add(1)
				d.aggrGroupsNum.Add(1)
				d.metrics.AggrGroups.Set(float64(d.aggrGroupsNum.Load()))
				break
			}
			if el == nil {
				continue
			}
			agExisting := el.(*aggrGroup)
			if agExisting.insert(ctx, a) {
				return
			}
		}

		retries++
		d.metrics.AggrGroupCreationRetries.Inc()
		if retries > 100 {
			d.metrics.AggrGroupCreationGivenUp.Inc()
			logger.Error("excessive retries creating aggregation group",
				zap.String("alert", a.Name()), zap.Int("retries", retries))
			return
		}
	}

	span.AddEvent("new AggregationGroup created",
		trace.WithAttributes(
			attribute.String("alerting.aggregation_group.key", ag.GroupKey()),
			attribute.Int("alerting.aggregation_group.count", int(d.aggrGroupsNum.Load())),
		),
	)

	if a.StartsAt.Add(ag.opts.GroupWait).Before(now) {
		ag.resetTimer(0)
	}

	switch d.state.Load() {
	case dispatcherStateWaitingToStart:
	case dispatcherStateRunning:
		d.runAG(ag)
	default:
		logger.Warn("unknown dispatcher state", zap.Int("state", int(d.state.Load())))
	}
}

// runAG starts the aggrGroup's run goroutine if not already running.
func (d *Dispatcher) runAG(ag *aggrGroup) {
	if !ag.running.CompareAndSwap(false, true) {
		return
	}
	go ag.run(func(ctx context.Context, alerts ...*alert.Alert) bool {
		_, _, err := d.stage.Exec(ctx, alerts...)
		if err != nil {
			fs := []zap.Field{zap.Int("num_alerts", len(alerts)), zap.Error(err)}
			if errors.Is(ctx.Err(), context.Canceled) {
				logger.Debug("notify for alerts failed", fs...)
			} else {
				logger.Error("notify for alerts failed", fs...)
			}
		}
		return err == nil
	})
}

// doMaintenance cleans up destroyed aggregation groups.
func (d *Dispatcher) doMaintenance() {
	for i := range d.routeGroupsSlice {
		d.routeGroupsSlice[i].groups.Range(func(_, el any) bool {
			ag := el.(*aggrGroup)
			if ag.destroyed() {
				ag.stop()
				if deleted := d.routeGroupsSlice[i].groups.CompareAndDelete(ag.fingerprint(), ag); deleted {
					d.marker.DeleteByGroupKey(ag.routeKey, ag.GroupKey())
					d.routeGroupsSlice[i].groupsLen.Add(-1)
					d.aggrGroupsNum.Add(-1)
					d.metrics.AggrGroups.Dec()
				}
			}
			return true
		})
	}
}

func (d *Dispatcher) WaitForLoading() {
	<-d.loaded
}

func (d *Dispatcher) LoadingDone() <-chan struct{} {
	return d.loaded
}

// Groups returns a slice of AlertGroups from the dispatcher's internal state.
func (d *Dispatcher) Groups(routeFilter func(*Route) bool, alertFilter func(*alert.Alert, time.Time) bool) (AlertGroups, map[label.Fingerprint][]string) {
	<-d.loaded

	groups := AlertGroups{}
	receivers := map[label.Fingerprint][]string{}
	now := time.Now()

	for i := range d.routeGroupsSlice {
		if !routeFilter(d.routeGroupsSlice[i].route) {
			continue
		}
		receiver := d.routeGroupsSlice[i].route.RouteOpts.Receiver

		snapshot := make([]*aggrGroup, 0, d.routeGroupsSlice[i].groupsLen.Load()+32)
		d.routeGroupsSlice[i].groups.Range(func(_, el any) bool {
			snapshot = append(snapshot, el.(*aggrGroup))
			return true
		})

		for _, ag := range snapshot {
			alerts := ag.alerts.List()
			filteredAlerts := make([]*alert.Alert, 0, len(alerts))
			for _, a := range alerts {
				if !alertFilter(a, now) {
					continue
				}
				fp := a.Fingerprint()
				if r, ok := receivers[fp]; ok {
					receivers[fp] = append(r, receiver)
				} else {
					receivers[fp] = []string{receiver}
				}
				filteredAlerts = append(filteredAlerts, a)
			}
			if len(filteredAlerts) == 0 {
				continue
			}
			alertGroup := &AlertGroup{
				Labels:   ag.labels,
				Receiver: receiver,
				Alerts:   filteredAlerts,
			}
			groups = append(groups, alertGroup)
		}
	}

	sort.Sort(groups)
	for i := range groups {
		sort.Sort(groups[i].Alerts)
	}
	for i := range receivers {
		sort.Strings(receivers[i])
	}

	return groups, receivers
}

// Stop the dispatcher.
func (d *Dispatcher) Stop() {
	if d == nil {
		return
	}
	d.state.Store(dispatcherStateStopped)
	if d.cancel == nil {
		return
	}
	d.cancel()
	d.cancel = nil
	d.finished.Wait()
}

type notifyFunc func(context.Context, ...*alert.Alert) bool

type nilLimits struct{}

func (n nilLimits) MaxNumberOfAggregationGroups() int { return 0 }
