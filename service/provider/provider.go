// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package provider

import (
	"context"

	"github.com/tsingsun/woocoo"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
)

// Alert wraps an alert with metadata such as propagated tracing information.
type Alert struct {
	// Header contains metadata, for example propagated tracing information.
	Header map[string]string
	Data   *alert.Alert
}

// Iterator provides the functions common to all iterators.
type Iterator interface {
	Err() error
	Close()
}

// AlertIterator is an Iterator for Alerts.
type AlertIterator interface {
	Iterator
	Next() <-chan *Alert
}

// NewAlertIterator returns a new AlertIterator.
func NewAlertIterator(ch <-chan *Alert, done chan struct{}, err error) AlertIterator {
	return &alertIterator{
		ch:   ch,
		done: done,
		err:  err,
	}
}

type alertIterator struct {
	ch   <-chan *Alert
	done chan struct{}
	err  error
}

func (ai alertIterator) Next() <-chan *Alert { return ai.ch }
func (ai alertIterator) Err() error          { return ai.err }
func (ai alertIterator) Close()              { close(ai.done) }

// Alerts gives access to a set of alerts. All methods are goroutine-safe.
type Alerts interface {
	woocoo.Server
	// Subscribe returns an iterator over new alerts.
	Subscribe(name string) AlertIterator
	// SlurpAndSubscribe returns a list of all active alerts available before
	// the call and an iterator for all alerts available after the call.
	SlurpAndSubscribe(name string) ([]*alert.Alert, AlertIterator)
	// GetPending returns an iterator over all alerts that have pending notifications.
	GetPending() AlertIterator
	// Get returns the alert for a given fingerprint.
	Get(label.Fingerprint) (*alert.Alert, error)
	// Put adds the given set of alerts to the set.
	Put(ctx context.Context, alerts ...*alert.Alert) error
}
