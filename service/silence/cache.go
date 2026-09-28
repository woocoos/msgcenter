// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package silence

import (
	"sync"

	"github.com/woocoos/msgcenter/pkg/label"
)

// cacheEntry stores the IDs of silences that match an alert and the version
// of the silences state the result is based on.
type cacheEntry struct {
	silenceIDs []int
	version    int
}

func newCacheEntry(version int, silenceIDs ...int) *cacheEntry {
	return &cacheEntry{
		silenceIDs: silenceIDs,
		version:    version,
	}
}

func (e *cacheEntry) count() int {
	return len(e.silenceIDs)
}

// cache stores the IDs of silences that match an alert and the version of the
// silences state the result is based on.
type cache struct {
	entries map[label.Fingerprint]*cacheEntry
	mtx     sync.RWMutex
}

func (c *cache) delete(fp label.Fingerprint) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	delete(c.entries, fp)
}

func (c *cache) get(fp label.Fingerprint) *cacheEntry {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	if e, found := c.entries[fp]; found {
		return e
	}
	return &cacheEntry{}
}

func (c *cache) set(fp label.Fingerprint, entry *cacheEntry) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.entries[fp] = entry
}
