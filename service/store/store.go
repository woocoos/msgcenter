// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package store

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/limit"
)

// ErrLimited is returned if a Store has reached the per-alert limit.
var ErrLimited = errors.New("alert limited")

// ErrNotFound is returned if a Store cannot find the Alert.
var ErrNotFound = errors.New("alert not found")

// ErrDestroyed is returned if a Store has been destroyed.
var ErrDestroyed = errors.New("alert store destroyed")

// Alerts provides lock-coordinated to an in-memory map of alerts, keyed by
// their fingerprint. Resolved alerts are removed from the map based on
// gcInterval. An optional callback can be set which receives a slice of all
// resolved alerts that have been removed.
type Alerts struct {
	sync.Mutex
	alerts        map[label.Fingerprint]*alert.Alert
	gcCallback    func([]*alert.Alert)
	limits        map[string]*limit.Bucket[label.Fingerprint]
	perAlertLimit int
	destroyed     bool
}

// NewAlerts returns a new Alerts struct.
func NewAlerts() *Alerts {
	a := &Alerts{
		alerts:     make(map[label.Fingerprint]*alert.Alert),
		gcCallback: func(_ []*alert.Alert) {},
	}

	return a
}

// WithPerAlertLimit sets the per-alert limit for the Alerts struct.
func (a *Alerts) WithPerAlertLimit(lim int) *Alerts {
	a.Lock()
	defer a.Unlock()

	a.limits = make(map[string]*limit.Bucket[label.Fingerprint])
	a.perAlertLimit = lim

	return a
}

// SetGCCallback sets a GC callback to be executed after each GC.
func (a *Alerts) SetGCCallback(cb func([]*alert.Alert)) {
	a.Lock()
	defer a.Unlock()

	a.gcCallback = cb
}

func (a *Alerts) Run(ctx context.Context, interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			a.GC()
		}
	}
}

func (a *Alerts) GC() (deleted []*alert.Alert) {
	// Remove stale alert limit buckets.
	a.gcLimitBuckets()

	// Delete resolved alerts.
	deleted = a.gcAlerts()

	// Execute GC callback if needed.
	if len(deleted) > 0 {
		a.gcCallback(deleted)
	}

	return deleted
}

// gcAlerts deletes resolved alerts and returns a copy of them.
func (a *Alerts) gcAlerts() (deleted []*alert.Alert) {
	a.Lock()
	defer a.Unlock()
	for fp, alert := range a.alerts {
		if alert.Resolved() {
			deleted = append(deleted, alert)
			delete(a.alerts, fp)
		}
	}
	return deleted
}

// gcLimitBuckets removes stale alert limit buckets.
func (a *Alerts) gcLimitBuckets() {
	a.Lock()
	defer a.Unlock()

	for alertName, bucket := range a.limits {
		if bucket.IsStale() {
			delete(a.limits, alertName)
		}
	}
}

// Get returns the Alert with the matching fingerprint, or an error if it is
// not found.
func (a *Alerts) Get(fp label.Fingerprint) (*alert.Alert, error) {
	a.Lock()
	defer a.Unlock()

	alert, prs := a.alerts[fp]
	if !prs {
		return nil, ErrNotFound
	}
	return alert, nil
}

// Set unconditionally sets the alert in memory.
func (a *Alerts) Set(alert *alert.Alert) error {
	a.Lock()
	defer a.Unlock()

	if a.destroyed {
		return ErrDestroyed
	}
	fp := alert.Fingerprint()
	name := alert.Name()

	// Apply per alert limits if necessary
	if a.perAlertLimit > 0 {
		bucket, ok := a.limits[name]
		if !ok {
			bucket = limit.NewBucket[label.Fingerprint](a.perAlertLimit)
			a.limits[name] = bucket
		}
		if !bucket.Upsert(fp, alert.EndsAt) {
			return ErrLimited
		}
	}

	a.alerts[fp] = alert
	return nil
}

// DeleteIfNotModified deletes the slice of Alerts from the store if not
// modified.
func (a *Alerts) DeleteIfNotModified(alerts alert.Alerts, destroyIfEmpty bool) error {
	a.Lock()
	defer a.Unlock()
	for _, alert := range alerts {
		fp := alert.Fingerprint()
		if other, ok := a.alerts[fp]; ok && alert.UpdatedAt.Equal(other.UpdatedAt) {
			delete(a.alerts, fp)
		}
	}

	// If the store is now empty, mark it as destroyed
	if len(a.alerts) == 0 && destroyIfEmpty {
		a.destroyed = true
	}

	return nil
}

// List returns a slice of Alerts currently held in memory.
func (a *Alerts) List() []*alert.Alert {
	a.Lock()
	defer a.Unlock()

	als := make([]*alert.Alert, 0, len(a.alerts))
	for _, alert := range a.alerts {
		als = append(als, alert)
	}

	return als
}

// Empty returns true if the store is empty.
func (a *Alerts) Empty() bool {
	a.Lock()
	defer a.Unlock()

	return len(a.alerts) == 0
}

func (a *Alerts) Destroyed() bool {
	a.Lock()
	defer a.Unlock()

	return a.destroyed
}

// Len returns the number of alerts in the store.
func (a *Alerts) Len() int {
	a.Lock()
	defer a.Unlock()

	return len(a.alerts)
}
