// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package inhibit

import (
	"context"
	"sync"
	"time"

	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/marker"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/pkg/tracing"
	"github.com/woocoos/msgcenter/service/provider"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var logger = log.Component("inhibit")
var tracer = tracing.NewTracer("github.com/woocoos/msgcenter/inhibit")

// An Inhibitor determines whether a given label set is muted based on the
// currently active alerts and a set of inhibition rules. It implements the
// Muter interface.
type Inhibitor struct {
	alerts     provider.Alerts
	rules      []*InhibitRule
	propagator propagation.TextMapPropagator

	mtx             sync.RWMutex
	loadingFinished sync.WaitGroup
	cancel          func()
}

// NewInhibitor returns a new Inhibitor.
func NewInhibitor(ap provider.Alerts, rs []profile.InhibitRule) *Inhibitor {
	ih := &Inhibitor{
		alerts:     ap,
		propagator: otel.GetTextMapPropagator(),
	}

	ih.loadingFinished.Add(1)
	ruleNames := make(map[string]struct{})
	for i, cr := range rs {
		if _, ok := ruleNames[cr.Name]; ok {
			logger.Debug("duplicate inhibition rule name", zap.Int("index", i), zap.String("name", cr.Name))
		}
		r := NewInhibitRule(cr)
		ih.rules = append(ih.rules, r)

		if cr.Name != "" {
			ruleNames[cr.Name] = struct{}{}
		}
	}
	return ih
}

func (ih *Inhibitor) run(ctx context.Context) {
	initalAlerts, it := ih.alerts.SlurpAndSubscribe("inhibitor")
	defer it.Close()

	for _, a := range initalAlerts {
		ih.processAlert(ctx, a)
	}

	ih.loadingFinished.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case pAlert := <-it.Next():
			if err := it.Err(); err != nil {
				logger.Error("Error iterating alerts", zap.Error(err))
				continue
			}
			traceCtx := context.Background()
			if pAlert.Header != nil {
				traceCtx = ih.propagator.Extract(traceCtx, propagation.MapCarrier(pAlert.Header))
			}

			ih.processAlert(traceCtx, pAlert.Data)
		}
	}
}

func (ih *Inhibitor) processAlert(ctx context.Context, a *alert.Alert) {
	_, span := tracer.Start(ctx, "inhibit.Inhibitor.processAlert",
		trace.WithAttributes(
			attribute.String("alerting.alert.name", a.Name()),
			attribute.String("alerting.alert.fingerprint", a.Fingerprint().String()),
		),
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// Update the inhibition rules' cache.
	for _, r := range ih.rules {
		if r.SourceMatchers.Matches(a.Labels) {
			attr := attribute.String("alerting.inhibit_rule.name", r.Name)
			span.AddEvent("alert matched rule source", trace.WithAttributes(attr))
			span.SetAttributes(attr)
			r.cache.set(a)
		}
	}
}

func (ih *Inhibitor) WaitForLoading() {
	ih.loadingFinished.Wait()
}

// Run the Inhibitor's background processing.
func (ih *Inhibitor) Run() {
	var ctx context.Context
	ih.mtx.Lock()
	ctx, ih.cancel = context.WithCancel(context.Background())
	ih.mtx.Unlock()

	for _, rule := range ih.rules {
		go rule.cache.run(ctx, 15*time.Minute)
	}
	ih.run(ctx)   // 直接阻塞
	ih.cancel()   // 退出时清理 cache
}

// Stop the Inhibitor's background processing.
func (ih *Inhibitor) Stop() {
	if ih == nil {
		return
	}

	ih.mtx.RLock()
	defer ih.mtx.RUnlock()
	if ih.cancel != nil {
		ih.cancel()
	}
}

// Mutes returns true if the given label set is muted. It implements the Muter
// interface.
func (ih *Inhibitor) Mutes(ctx context.Context, lset label.LabelSet) bool {
	fp := lset.Fingerprint()

	_, span := tracer.Start(ctx, "inhibit.Inhibitor.Mutes",
		trace.WithAttributes(attribute.String("alerting.alert.fingerprint", fp.String())),
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	var inhibitedBy []string
	defer func() {
		// 从 context 获取 marker, 设置抑制状态.
		m, ok := marker.FromContext(ctx)
		if ok {
			m.SetInhibited(fp, inhibitedBy)
		}
	}()

	now := time.Now()
	for _, r := range ih.rules {
		if !r.TargetMatchers.Matches(lset) {
			// If target side of rule doesn't match, we don't need to look any further.
			continue
		}
		span.AddEvent("alert matched rule target",
			trace.WithAttributes(attribute.String("alerting.inhibit_rule.name", r.Name)),
		)
		// If we are here, the target side matches. If the source side matches, too, we
		// need to exclude inhibiting alerts for which the same is true.
		if inhibitedByFP, eq := r.hasEqual(lset, r.SourceMatchers.Matches(lset), now); eq {
			inhibitedBy = append(inhibitedBy, inhibitedByFP.String())
			span.AddEvent("alert inhibited",
				trace.WithAttributes(attribute.String("alerting.inhibit_rule.source.fingerprint", inhibitedByFP.String())),
			)
			return true
		}
	}
	span.AddEvent("alert not inhibited")

	return false
}

// An InhibitRule specifies that a class of (source) alerts should inhibit
// notifications for another class of (target) alerts if all specified matching
// labels are equal between the two alerts. This may be used to inhibit alerts
// from sending notifications if their meaning is logically a subset of a
// higher-level alert.
type InhibitRule struct {
	// Name is an optional name for the inhibition rule.
	Name string
	// The set of Filters which define the group of source alerts (which inhibit
	// the target alerts).
	SourceMatchers label.Matchers
	// The set of Filters which define the group of target alerts (which are
	// inhibited by the source alerts).
	TargetMatchers label.Matchers
	// A set of label names whose label values need to be identical in source and
	// target alerts in order for the inhibition to take effect.
	Equal map[label.LabelName]struct{}

	// Cache of alerts matching source labels.
	cache *cache
}

// NewInhibitRule returns a new InhibitRule based on a configuration definition.
func NewInhibitRule(cr profile.InhibitRule) *InhibitRule {
	var (
		sourcem label.Matchers
		targetm label.Matchers
	)
	// We append the new-style matchers. This can be simplified once the deprecated matcher syntax is removed.
	sourcem = append(sourcem, cr.SourceMatchers...)

	// We append the new-style matchers. This can be simplified once the deprecated matcher syntax is removed.
	targetm = append(targetm, cr.TargetMatchers...)

	equal := map[label.LabelName]struct{}{}
	for _, ln := range cr.Equal {
		equal[ln] = struct{}{}
	}

	return &InhibitRule{
		Name:           cr.Name,
		SourceMatchers: sourcem,
		TargetMatchers: targetm,
		Equal:          equal,
		cache:          newCache(equal),
	}
}

// hasEqual checks whether the source cache contains alerts matching the equal
// labels for the given label set. If so, the fingerprint of one of those alerts
// is returned. If excludeTwoSidedMatch is true, alerts that match both the
// source and the target side of the rule are disregarded.
func (r *InhibitRule) hasEqual(lset label.LabelSet, excludeTwoSidedMatch bool, now time.Time) (label.Fingerprint, bool) {
	return r.cache.find(lset, now, func(a *alert.Alert) bool {
		return !excludeTwoSidedMatch || !r.TargetMatchers.Matches(a.Labels)
	})
}
