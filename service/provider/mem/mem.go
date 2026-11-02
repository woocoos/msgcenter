// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package mem

import (
	"context"
	"sync"
	"time"

	"github.com/woocoos/msgcenter/pkg/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/service/provider"
	"github.com/woocoos/msgcenter/service/store"
)

const alertChannelLength = 200

var logger = log.Component("provider")
var tracer = tracing.NewTracer("github.com/woocoos/msgcenter/provider/mem")

// Alerts gives access to a set of alerts. All methods are goroutine-safe.
type Alerts struct {
	intervalGC time.Duration
	alerts     *store.Alerts
	propagator propagation.TextMapPropagator

	mtx       sync.Mutex
	listeners map[int]listeningAlerts
	next      int

	callback AlertStoreCallback

	stop chan struct{}
}

// AlertStoreCallback defines callbacks for alert store operations.
type AlertStoreCallback interface {
	PreStore(alert *alert.Alert, existing bool) error
	PostStore(alert *alert.Alert, existing bool)
	PostDelete(alert *alert.Alert)
	// PostGC is called after alerts are garbage collected. The fingerprints
	// of deleted alerts are passed so that listeners can clean up caches.
	PostGC(fps []label.Fingerprint)
}

type listeningAlerts struct {
	name   string
	alerts chan *provider.Alert
	done   chan struct{}
}

// NewAlerts returns a new alert provider.
func NewAlerts(intervalGC time.Duration, perAlertNameLimit int,
	alertCallback AlertStoreCallback) (*Alerts, error) {
	if alertCallback == nil {
		alertCallback = noopCallback{}
	}

	if perAlertNameLimit > 0 {
		logger.Info("per alert name limit enabled", zap.Int("limit", perAlertNameLimit))
	}

	a := &Alerts{
		alerts:     store.NewAlerts().WithPerAlertLimit(perAlertNameLimit),
		intervalGC: intervalGC,
		listeners:  map[int]listeningAlerts{},
		next:       0,
		callback:   alertCallback,
		propagator: otel.GetTextMapPropagator(),
		stop:       make(chan struct{}),
	}
	return a, nil
}

func (a *Alerts) gcLoop(ctx context.Context, interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			a.gc()
		case <-a.stop:
			return
		}
	}
}

func (a *Alerts) gc() {
	a.gcListeners()

	// As we don't persist alerts, we no longer consider them after
	// they are resolved. Alerts waiting for resolved notifications are
	// held in memory in aggregation groups redundantly.
	deleted := a.gcAlerts()

	// If there are no deleted alerts, there is nothing to do.
	if len(deleted) == 0 {
		return
	}

	ff := make([]label.Fingerprint, len(deleted))
	for i, alert := range deleted {
		ff[i] = alert.Fingerprint()
		a.callback.PostDelete(alert)
	}
	a.callback.PostGC(ff)
}

func (a *Alerts) gcAlerts() []*alert.Alert {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	return a.alerts.GC()
}

func (a *Alerts) gcListeners() {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	for i, l := range a.listeners {
		select {
		case <-l.done:
			delete(a.listeners, i)
			close(l.alerts)
		default:
			// listener is not closed yet, hence proceed.
		}
	}
}

// Start the alert provider.
func (a *Alerts) Start(ctx context.Context) error {
	a.gcLoop(ctx, a.intervalGC)
	return nil
}

// Stop the alert provider. It works with Start
func (a *Alerts) Stop(_ context.Context) error {
	close(a.stop)
	return nil
}

// Subscribe returns an iterator over new alerts.
// Existing alerts are pre-filled into the channel.
func (a *Alerts) Subscribe(name string) provider.AlertIterator {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	var (
		done = make(chan struct{})
		al   = a.alerts.List()
		ch   = make(chan *provider.Alert, max(len(al), alertChannelLength))
	)

	for _, alt := range al {
		ch <- &provider.Alert{
			Header: map[string]string{},
			Data:   alt,
		}
	}

	a.listeners[a.next] = listeningAlerts{name: name, alerts: ch, done: done}
	a.next++

	return provider.NewAlertIterator(ch, done, nil)
}

// SlurpAndSubscribe returns all active alerts available before the call as a
// slice, and an iterator for alerts arriving after the call. Unlike Subscribe,
// existing alerts are NOT pre-filled into the channel.
func (a *Alerts) SlurpAndSubscribe(name string) ([]*alert.Alert, provider.AlertIterator) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	var (
		done = make(chan struct{})
		al   = a.alerts.List()
		ch   = make(chan *provider.Alert, alertChannelLength)
	)

	a.listeners[a.next] = listeningAlerts{name: name, alerts: ch, done: done}
	a.next++

	return al, provider.NewAlertIterator(ch, done, nil)
}

// GetPending returns an iterator over all the alerts that have pending notifications.
func (a *Alerts) GetPending() provider.AlertIterator {
	var (
		ch   = make(chan *provider.Alert, alertChannelLength)
		done = make(chan struct{})
	)
	a.mtx.Lock()
	defer a.mtx.Unlock()
	alerts := a.alerts.List()

	go func() {
		defer close(ch)
		for _, alt := range alerts {
			select {
			case ch <- &provider.Alert{Header: map[string]string{}, Data: alt}:
			case <-done:
				return
			}
		}
	}()

	return provider.NewAlertIterator(ch, done, nil)
}

// Get returns the alert for a given fingerprint.
func (a *Alerts) Get(fp label.Fingerprint) (*alert.Alert, error) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	return a.alerts.Get(fp)
}

// Put adds the given alerts to the set.
func (a *Alerts) Put(ctx context.Context, alerts ...*alert.Alert) error {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	ctx, span := tracer.Start(ctx, "provider.mem.Put",
		trace.WithAttributes(
			attribute.Int("alerting.alerts.count", len(alerts)),
		),
		trace.WithSpanKind(trace.SpanKindProducer),
	)
	defer span.End()

	for _, alt := range alerts {
		fp := alt.Fingerprint()

		existing := false

		// Check that there's an alert existing within the store before
		// trying to merge.
		if old, err := a.alerts.Get(fp); err == nil {
			existing = true
			// Merge alerts if there is an overlap in activity range.
			if (alt.EndsAt.After(old.StartsAt) && alt.EndsAt.Before(old.EndsAt)) ||
				(alt.StartsAt.After(old.StartsAt) && alt.StartsAt.Before(old.EndsAt)) {
				alt = old.Merge(alt)
			}
		}

		if err := a.callback.PreStore(alt, existing); err != nil {
			logger.Error("pre-store callback returned error on set alert", zap.Error(err))
			continue
		}

		if err := a.alerts.Set(alt); err != nil {
			logger.Warn("error on set alert", zap.Error(err))
			continue
		}

		a.callback.PostStore(alt, existing)

		// Inject tracing context into header for downstream consumers.
		metadata := map[string]string{}
		a.propagator.Inject(ctx, propagation.MapCarrier(metadata))
		msg := &provider.Alert{
			Data:   alt,
			Header: metadata,
		}

		for _, l := range a.listeners {
			select {
			case l.alerts <- msg:
			case <-l.done:
			}
		}
	}

	return nil
}

type noopCallback struct{}

func (n noopCallback) PreStore(_ *alert.Alert, _ bool) error { return nil }
func (n noopCallback) PostStore(_ *alert.Alert, _ bool)      {}
func (n noopCallback) PostDelete(_ *alert.Alert)             {}
func (n noopCallback) PostGC(_ []label.Fingerprint)          {}
