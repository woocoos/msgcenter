// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package marker

import (
	"sync"

	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
)

// NewAlertMarker returns a new AlertMarker backed by an in-memory map.
func NewAlertMarker() AlertMarker {
	return &alertMarker{
		status: map[label.Fingerprint]*alertStatus{},
	}
}

type alertMarker struct {
	status map[label.Fingerprint]*alertStatus
	mtx    sync.RWMutex
}

func (m *alertMarker) SetSilenced(fp label.Fingerprint, silencedBy []string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	s, found := m.status[fp]
	if !found {
		s = &alertStatus{}
		m.status[fp] = s
	}
	s.SilencedBy = cloneStrings(silencedBy)
}

func (m *alertMarker) SetInhibited(fp label.Fingerprint, inhibitedBy []string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	s, found := m.status[fp]
	if !found {
		s = &alertStatus{}
		m.status[fp] = s
	}
	s.InhibitedBy = cloneStrings(inhibitedBy)
}

func (m *alertMarker) Status(fp label.Fingerprint) alert.MarkerStatus {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	status := alert.MarkerStatus{
		State:       alert.AlertStateUnprocessed,
		SilencedBy:  []int{},
		InhibitedBy: []string{},
	}

	s, found := m.status[fp]
	if !found {
		return status
	}

	status.State = s.state()
	if s.InhibitedBy != nil {
		status.InhibitedBy = cloneStrings(s.InhibitedBy)
	}
	return status
}

func (m *alertMarker) Delete(fps ...label.Fingerprint) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	for _, fp := range fps {
		delete(m.status, fp)
	}
}

func cloneStrings(s []string) []string {
	if s == nil {
		return nil
	}
	c := make([]string, len(s))
	copy(c, s)
	return c
}
