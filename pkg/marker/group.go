// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package marker

import "sync"

// NewGroupMarker returns an instance of a GroupMarker implementation.
func NewGroupMarker() GroupMarker {
	return &groupMarker{
		groups: map[groupMarkerKey]*groupStatus{},
	}
}

type groupMarkerKey struct {
	routeID  string
	groupKey string
}

type groupStatus struct {
	mutedBy []string
}

type groupMarker struct {
	groups map[groupMarkerKey]*groupStatus
	mtx    sync.RWMutex
}

func (m *groupMarker) Muted(routeID, groupKey string) ([]string, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	status, ok := m.groups[groupMarkerKey{routeID: routeID, groupKey: groupKey}]
	if !ok {
		return nil, false
	}
	return status.mutedBy, len(status.mutedBy) > 0
}

func (m *groupMarker) SetMuted(routeID, groupKey string, timeIntervalNames []string) {
	key := groupMarkerKey{routeID: routeID, groupKey: groupKey}
	m.mtx.Lock()
	defer m.mtx.Unlock()
	status, ok := m.groups[key]
	if !ok {
		status = &groupStatus{}
		m.groups[key] = status
	}
	status.mutedBy = timeIntervalNames
}

func (m *groupMarker) DeleteByGroupKey(routeID, groupKey string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	delete(m.groups, groupMarkerKey{routeID: routeID, groupKey: groupKey})
}
