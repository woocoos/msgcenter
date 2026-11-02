// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package marker

import "github.com/woocoos/msgcenter/pkg/alert"

type alertStatus struct {
	InhibitedBy []string
	SilencedBy  []string
}

// state calculates the alert state based on inhibition and silence status.
func (s *alertStatus) state() alert.AlertState {
	if len(s.InhibitedBy) > 0 || len(s.SilencedBy) > 0 {
		return alert.AlertStateSuppressed
	}
	return alert.AlertStateActive
}
