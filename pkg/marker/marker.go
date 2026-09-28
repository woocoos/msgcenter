// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package marker

import (
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
)

// AlertMarker tracks per-alert silenced/inhibited status within a single
// aggregation group. Each aggregation group owns its own instance.
// All methods are goroutine-safe.
type AlertMarker interface {
	// SetInhibited sets the inhibitedBy for the given fingerprint.
	// If inhibitedBy is empty, it clears the inhibitedBy.
	SetInhibited(fp label.Fingerprint, inhibitedBy []string)

	// SetSilenced sets the silencedBy for the given fingerprint.
	// If silencedBy is empty, it clears the silencedBy.
	SetSilenced(fp label.Fingerprint, silencedBy []string)

	// Status returns the MarkerStatus for the given fingerprint.
	// If the fingerprint is not found, it returns an unprocessed status.
	Status(fp label.Fingerprint) alert.MarkerStatus

	// Delete removes markers for the given fingerprints.
	Delete(fps ...label.Fingerprint)
}

// GroupMarker helps to mark groups as muted.
// All methods are goroutine-safe.
type GroupMarker interface {
	// Muted returns true if the group is muted, otherwise false. If the group
	// is muted then it also returns the names of the time intervals that muted it.
	Muted(routeID, groupKey string) ([]string, bool)

	// SetMuted marks the group as muted, and sets the names of the time
	// intervals that mute it. If the list of names is nil or the empty slice
	// then the muted marker is removed.
	SetMuted(routeID, groupKey string, timeIntervalNames []string)

	// DeleteByGroupKey removes all markers for the GroupKey.
	DeleteByGroupKey(routeID, groupKey string)
}
