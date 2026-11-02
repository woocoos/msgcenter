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

	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
)

// fingerprintSet is a set of alert fingerprints.
type fingerprintSet map[label.Fingerprint]struct{}

// cache contains the runtime state of the inhibit rule.
// It merges alert storage and index under a single lock to ensure
// atomic operations and avoid race conditions.
type cache struct {
	equal map[label.LabelName]struct{}

	mtx sync.RWMutex
	// alerts is the map of alerts that match the source matchers of the inhibit rule.
	alerts map[label.Fingerprint]*alert.Alert
	// index is a map of equal label fingerprint to the set of source alert fingerprints
	// stored in the cache.
	index map[label.Fingerprint]fingerprintSet
}

func newCache(equal map[label.LabelName]struct{}) *cache {
	return &cache{
		equal:  equal,
		alerts: make(map[label.Fingerprint]*alert.Alert),
		index:  make(map[label.Fingerprint]fingerprintSet),
	}
}

// fingerprintEquals returns the fingerprint of the equal labels of the given label set.
func (c *cache) fingerprintEquals(lset label.LabelSet) label.Fingerprint {
	equalSet := make(label.LabelSet, len(c.equal))
	for n := range c.equal {
		equalSet[n] = lset[n]
	}
	return equalSet.Fingerprint()
}

// set adds or replaces the given source alert.
func (c *cache) set(a *alert.Alert) {
	fp := a.Fingerprint()
	eq := c.fingerprintEquals(a.Labels)

	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.alerts[fp] = a
	set, ok := c.index[eq]
	if !ok {
		set = make(fingerprintSet)
		c.index[eq] = set
	}
	set[fp] = struct{}{}
}

// find returns the fingerprint of a cached source alert that shares the equal
// labels of lset, is active at now, and satisfies match.
func (c *cache) find(lset label.LabelSet, now time.Time, match func(*alert.Alert) bool) (label.Fingerprint, bool) {
	eq := c.fingerprintEquals(lset)

	c.mtx.RLock()
	defer c.mtx.RUnlock()

	for fp := range c.index[eq] {
		a := c.alerts[fp]
		if a.ResolvedAt(now) {
			continue
		}
		if !match(a) {
			continue
		}
		return fp, true
	}

	return 0, false
}

func (c *cache) run(ctx context.Context, interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			c.gc()
		}
	}
}

func (c *cache) gc() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	for fp, a := range c.alerts {
		if !a.Resolved() {
			continue
		}
		delete(c.alerts, fp)

		eq := c.fingerprintEquals(a.Labels)
		set := c.index[eq]
		delete(set, fp)
		if len(set) == 0 {
			delete(c.index, eq)
		}
	}
}
