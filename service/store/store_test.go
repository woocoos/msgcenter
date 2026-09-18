// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2018 Prometheus Team.
// Licensed under the Apache License 2.0.

package store

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
)

func TestSetGet(t *testing.T) {
	a := NewAlerts()
	a1 := &alert.Alert{
		Labels:    label.LabelSet{"foo": "bar"},
		UpdatedAt: time.Now(),
	}
	require.NoError(t, a.Set(a1))
	want := a1.Fingerprint()
	got, err := a.Get(want)

	require.NoError(t, err)
	require.Equal(t, want, got.Fingerprint())
}

func TestDeleteIfNotModified(t *testing.T) {
	t.Run("unmodified alert should be deleted", func(t *testing.T) {
		a := NewAlerts()
		a1 := &alert.Alert{
			Labels:    label.LabelSet{"foo": "bar"},
			UpdatedAt: time.Now().Add(-time.Second),
		}
		require.NoError(t, a.Set(a1))

		// a1 should be deleted as it has not been modified.
		require.NoError(t, a.DeleteIfNotModified(alert.Alerts{a1}, false))
		got, err := a.Get(a1.Fingerprint())
		require.Equal(t, ErrNotFound, err)
		require.Nil(t, got)
	})

	t.Run("modified alert should not be deleted", func(t *testing.T) {
		a := NewAlerts()
		a1 := &alert.Alert{
			Labels:    label.LabelSet{"foo": "bar"},
			UpdatedAt: time.Now(),
		}
		require.NoError(t, a.Set(a1))

		// Make a copy of a1 that is older, but do not put it.
		// We want to make sure a1 is not deleted.
		a2 := &alert.Alert{
			Labels:    label.LabelSet{"foo": "bar"},
			UpdatedAt: time.Now().Add(-time.Second),
		}
		require.True(t, a2.UpdatedAt.Before(a1.UpdatedAt))
		require.NoError(t, a.DeleteIfNotModified(alert.Alerts{a2}, false))
		// a1 should not be deleted.
		got, err := a.Get(a1.Fingerprint())
		require.NoError(t, err)
		require.Equal(t, a1, got)

		// Make another copy of a1 that is newer, but do not put it.
		// We want to make sure a1 is not deleted here either.
		a3 := &alert.Alert{
			Labels:    label.LabelSet{"foo": "bar"},
			UpdatedAt: time.Now().Add(time.Second),
		}
		require.True(t, a3.UpdatedAt.After(a1.UpdatedAt))
		require.NoError(t, a.DeleteIfNotModified(alert.Alerts{a3}, false))
		// a1 should not be deleted.
		got, err = a.Get(a1.Fingerprint())
		require.NoError(t, err)
		require.Equal(t, a1, got)
	})

	t.Run("should not delete other alerts", func(t *testing.T) {
		a := NewAlerts()
		a1 := &alert.Alert{
			Labels:    label.LabelSet{"foo": "bar"},
			UpdatedAt: time.Now(),
		}
		a2 := &alert.Alert{
			Labels:    label.LabelSet{"bar": "baz"},
			UpdatedAt: time.Now(),
		}
		require.NoError(t, a.Set(a1))
		require.NoError(t, a.Set(a2))

		// Deleting a1 should not delete a2.
		require.NoError(t, a.DeleteIfNotModified(alert.Alerts{a1}, true))
		// a1 should be deleted.
		got, err := a.Get(a1.Fingerprint())
		require.Equal(t, ErrNotFound, err)
		require.False(t, a.Destroyed())
		require.Nil(t, got)
		// a2 should not be deleted.
		got, err = a.Get(a2.Fingerprint())
		require.NoError(t, err)
		require.Equal(t, a2, got)
	})
}

func TestGC(t *testing.T) {
	now := time.Now()
	newAlert := func(key string, start, end time.Duration) *alert.Alert {
		return &alert.Alert{
			Labels:   label.LabelSet{label.LabelName(key): "b"},
			StartsAt: now.Add(start * time.Minute),
			EndsAt:   now.Add(end * time.Minute),
		}
	}
	active := []*alert.Alert{
		newAlert("b", 10, 20),
		newAlert("c", -10, 10),
	}
	resolved := []*alert.Alert{
		newAlert("a", -10, -5),
		newAlert("d", -10, -1),
	}
	s := NewAlerts()
	var (
		n           int
		done        = make(chan struct{})
		ctx, cancel = context.WithCancel(context.Background())
	)
	s.SetGCCallback(func(a []*alert.Alert) {
		n += len(a)
		if n >= len(resolved) {
			cancel()
		}
	})
	for _, a := range append(active, resolved...) {
		require.NoError(t, s.Set(a))
	}
	go func() {
		s.Run(ctx, 10*time.Millisecond)
		close(done)
	}()
	select {
	case <-done:
		break
	case <-time.After(1 * time.Second):
		t.Fatal("garbage collection didn't complete in time")
	}

	for _, a := range active {
		if _, err := s.Get(a.Fingerprint()); err != nil {
			t.Errorf("alert %v should not have been gc'd", a)
		}
	}
	for _, a := range resolved {
		if _, err := s.Get(a.Fingerprint()); err == nil {
			t.Errorf("alert %v should have been gc'd", a)
		}
	}
	require.Len(t, resolved, n)
}
