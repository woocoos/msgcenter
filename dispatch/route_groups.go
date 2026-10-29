// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package dispatch

import (
	"sync"
	"sync/atomic"
)

// routeAggrGroups holds the aggregation groups for a single route.
// It uses sync.Map for lock-free concurrent access.
type routeAggrGroups struct {
	route     *Route
	groups    sync.Map // map[label.Fingerprint]*aggrGroup
	groupsLen atomic.Int64
}
