// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package dispatch

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync/atomic"
	"time"

	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/marker"
	"github.com/woocoos/msgcenter/service/store"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// AlertGroup represents how alerts exist within an aggrGroup.
type AlertGroup struct {
	Alerts   alert.Alerts
	Labels   label.LabelSet
	Receiver string
}

type AlertGroups []*AlertGroup

func (ag AlertGroups) Swap(i, j int) { ag[i], ag[j] = ag[j], ag[i] }
func (ag AlertGroups) Less(i, j int) bool {
	if ag[i].Labels.Equal(ag[j].Labels) {
		return ag[i].Receiver < ag[j].Receiver
	}
	return ag[i].Labels.Before(ag[j].Labels)
}
func (ag AlertGroups) Len() int { return len(ag) }

// aggrGroup aggregates alert fingerprints into groups to which a
// common set of routing options applies.
// It emits notifications in the specified intervals.
type aggrGroup struct {
	labels   label.LabelSet
	opts     *RouteOpts
	routeKey string

	alerts  *store.Alerts
	marker  marker.AlertMarker
	ctx     context.Context
	cancel  func()
	done    chan struct{}
	next    *time.Timer
	timeout func(time.Duration) time.Duration

	logField zap.Field

	running atomic.Bool
}

// newAggrGroup returns a new aggregation group.
func newAggrGroup(ctx context.Context, labels label.LabelSet, r *Route, to func(time.Duration) time.Duration) *aggrGroup {
	if to == nil {
		to = func(d time.Duration) time.Duration { return d }
	}
	ag := &aggrGroup{
		labels:   labels,
		routeKey: r.Key(),
		opts:     &r.RouteOpts,
		timeout:  to,
		alerts:   store.NewAlerts(),
		marker:   marker.NewAlertMarker(),
		done:     make(chan struct{}),
	}
	ag.ctx, ag.cancel = context.WithCancel(ctx)

	// Set an initial one-time wait before flushing
	// the first batch of notifications.
	ag.next = time.NewTimer(ag.opts.GroupWait)

	ag.logField = zap.Stringer("aggrGroup", ag)
	return ag
}

func (ag *aggrGroup) fingerprint() label.Fingerprint {
	return ag.labels.Fingerprint()
}

func (ag *aggrGroup) GroupKey() string {
	return fmt.Sprintf("%s:%s", ag.routeKey, ag.labels)
}

func (ag *aggrGroup) String() string {
	return ag.GroupKey()
}

func (ag *aggrGroup) run(nf notifyFunc) {
	defer close(ag.done)
	defer ag.next.Stop()

	for {
		select {
		case now := <-ag.next.C:
			// Give the notifications time until the next flush to
			// finish before terminating them.
			ctx, cancel := context.WithTimeout(ag.ctx, ag.timeout(ag.opts.GroupInterval))

			// The now time we retrieve from the ticker is the only reliable
			// point of time reference for the subsequent notification pipeline.
			ctx = notify.WithNow(ctx, now)

			// Populate context with information needed along the pipeline.
			ctx = notify.WithGroupKey(ctx, ag.GroupKey())
			ctx = notify.WithGroupLabels(ctx, ag.labels)
			ctx = notify.WithReceiverName(ctx, ag.opts.Receiver)
			ctx = notify.WithRepeatInterval(ctx, ag.opts.RepeatInterval)
			ctx = notify.WithMuteTimeIntervals(ctx, ag.opts.MuteTimeIntervals)
			ctx = notify.WithActiveTimeIntervals(ctx, ag.opts.ActiveTimeIntervals)
			ctx = marker.WithContext(ctx, ag.marker)

			// Wait the configured interval before calling flush again.
			ag.next.Reset(ag.opts.GroupInterval)

			ag.flush(func(alerts ...*alert.Alert) bool {
				ctx, span := tracer.Start(ctx, "dispatch.AggregationGroup.flush",
					trace.WithAttributes(
						attribute.String("alerting.aggregation_group.key", ag.GroupKey()),
						attribute.Int("alerting.alerts.count", len(alerts)),
					),
					trace.WithSpanKind(trace.SpanKindInternal),
				)
				defer span.End()

				success := nf(ctx, alerts...)
				if !success {
					span.SetStatus(codes.Error, "notification failed")
				}
				return success
			})

			cancel()

			// If destroyed, exit: this alert group won't be used anymore.
			if ag.destroyed() {
				return
			}

		case <-ag.ctx.Done():
			return
		}
	}
}

func (ag *aggrGroup) stop() {
	// Calling cancel will terminate all in-process notifications
	// and the run() loop.
	ag.cancel()
	<-ag.done
}

// insert inserts the alert into the aggregation group.
// Returns false if the group has been destroyed.
func (ag *aggrGroup) insert(ctx context.Context, alert *alert.Alert) bool {
	_, span := tracer.Start(ctx, "dispatch.AggregationGroup.insert",
		trace.WithAttributes(
			attribute.String("alerting.alert.name", alert.Name()),
			attribute.String("alerting.alert.fingerprint", alert.Fingerprint().String()),
			attribute.String("alerting.aggregation_group.key", ag.GroupKey()),
		),
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	if ag.destroyed() {
		return false
	}
	if err := ag.alerts.Set(alert); err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return false
		}
		span.SetStatus(codes.Error, "error on set alert")
		span.RecordError(err)
		logger.Error("error on set alert", ag.logField, zap.Error(err))
		return false
	}
	return true
}

// resetTimer resets the group's timer to the given duration.
func (ag *aggrGroup) resetTimer(d time.Duration) {
	ag.next.Reset(d)
}

func (ag *aggrGroup) empty() bool {
	return ag.alerts.Empty()
}

func (ag *aggrGroup) destroyed() bool {
	return ag.alerts.Destroyed()
}

// flush sends notifications for all new alerts.
func (ag *aggrGroup) flush(notify func(...*alert.Alert) bool) {
	if ag.empty() {
		return
	}

	var (
		alerts        = ag.alerts.List()
		alertsSlice   = make(alert.Alerts, 0, len(alerts))
		resolvedSlice = make(alert.Alerts, 0, len(alerts))
		now           = time.Now()
	)
	for _, item := range alerts {
		a := *item
		// Ensure that alerts don't resolve as time move forwards.
		if a.ResolvedAt(now) {
			resolvedSlice = append(resolvedSlice, &a)
		} else {
			a.EndsAt = time.Time{}
		}
		alertsSlice = append(alertsSlice, &a)
	}
	sort.Stable(alertsSlice)

	logger.Debug("flushing", ag.logField, zap.String("alerts", fmt.Sprintf("%v", alertsSlice)))

	if notify(alertsSlice...) {
		// Delete only resolved alerts as we just sent a notification for them,
		// and we don't want to send another one. However, we need to make sure
		// that each resolved alert has not fired again during the flush as then
		// we would delete an active alert thinking it was resolved.
		if err := ag.alerts.DeleteIfNotModified(resolvedSlice, true); err != nil {
			logger.Error("error on delete alerts", zap.Error(err))
		} else {
			// Delete markers for resolved alerts that are no longer in the store.
			for _, a := range resolvedSlice {
				_, err := ag.alerts.Get(a.Fingerprint())
				if errors.Is(err, store.ErrNotFound) {
					ag.marker.Delete(a.Fingerprint())
				}
			}
		}
	}
}

func getGroupLabels(a *alert.Alert, route *Route) label.LabelSet {
	capacity := len(route.RouteOpts.GroupBy)
	if route.RouteOpts.GroupByAll {
		capacity = len(a.Labels)
	}
	groupLabels := make(label.LabelSet, capacity)
	for ln, lv := range a.Labels {
		if _, ok := route.RouteOpts.GroupBy[ln]; ok || route.RouteOpts.GroupByAll {
			groupLabels[ln] = lv
		}
	}
	return groupLabels
}
