// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package inhibit

import (
	"sync"

	"github.com/woocoos/msgcenter/pkg/label"
)

// index is a thread-safe fingerprint-to-fingerprint mapping used to accelerate
// inhibition lookups. The key is the fingerprint of the Equal-label subset of
// a source alert; the value is the source alert's own fingerprint.
type index struct {
	mtx   sync.RWMutex
	items map[label.Fingerprint]label.Fingerprint
}

func newIndex() *index {
	return &index{
		items: make(map[label.Fingerprint]label.Fingerprint),
	}
}

func (c *index) Get(key label.Fingerprint) (label.Fingerprint, bool) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	v, ok := c.items[key]
	return v, ok
}

func (c *index) Set(key, value label.Fingerprint) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.items[key] = value
}

func (c *index) Delete(key label.Fingerprint) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	delete(c.items, key)
}

func (c *index) Len() int {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return len(c.items)
}
