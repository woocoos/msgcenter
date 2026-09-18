// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package marker

import "context"

type markerContextKey int

const (
	keyAlertMarker markerContextKey = iota
)

// WithContext returns a copy of ctx carrying the given AlertMarker.
// Inhibitor and Silencer extract it from the context to write per-group
// alert status.
func WithContext(ctx context.Context, m AlertMarker) context.Context {
	return context.WithValue(ctx, keyAlertMarker, m)
}

// FromContext extracts the AlertMarker from ctx.
// If no marker is present, it returns (nil, false).
func FromContext(ctx context.Context) (AlertMarker, bool) {
	m, ok := ctx.Value(keyAlertMarker).(AlertMarker)
	return m, ok
}
