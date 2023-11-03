package mem

import (
	"context"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/service/provider"
	"github.com/woocoos/msgcenter/service/store"
	"go.uber.org/zap"
	"sync"
	"time"
)

const alertChannelLength = 200

var logger = log.Component("provider")

// Alerts gives access to a set of alerts. All methods are goroutine-safe.
type Alerts struct {
	ctx        context.Context
	cancel     context.CancelFunc
	intervalGC time.Duration
	alerts     *store.Alerts
	marker     alert.Marker

	mtx       sync.Mutex
	listeners map[int]listeningAlerts
	next      int

	callback AlertStoreCallback
}

type AlertStoreCallback interface {
	// PreStore is called before alert is stored into the store. If this method returns error,
	// alert is not stored.
	// Existing flag indicates whether alert has existed before (and is only updated) or not.
	// If alert has existed before, then alert passed to PreStore is result of merging existing alert with new alert.
	PreStore(alert *alert.Alert, existing bool) error

	// PostStore is called after alert has been put into store.
	PostStore(alert *alert.Alert, existing bool)

	// PostDelete is called after alert has been removed from the store due to alert garbage collection.
	PostDelete(alert *alert.Alert)
}

type listeningAlerts struct {
	alerts chan *alert.Alert
	done   chan struct{}
}

// NewAlerts returns a new alert provider.
func NewAlerts(ctx context.Context, m alert.Marker, intervalGC time.Duration, alertCallback AlertStoreCallback) (*Alerts, error) {
	if alertCallback == nil {
		alertCallback = noopCallback{}
	}
	if intervalGC == 0 {
		intervalGC = 30 * time.Minute
	}
	ctx, cancel := context.WithCancel(ctx)
	a := &Alerts{
		marker:     m,
		alerts:     store.NewAlerts(store.WithGCInterval(intervalGC)),
		intervalGC: intervalGC,
		ctx:        ctx,
		cancel:     cancel,
		listeners:  map[int]listeningAlerts{},
		next:       0,
		callback:   alertCallback,
	}
	return a, nil
}

func (a *Alerts) Start(_ context.Context) error {
	a.alerts.SetGCCallback(func(alerts []*alert.Alert) {
		for _, alt := range alerts {
			// As we don't persist alerts, we no longer consider them after
			// they are resolved. Alerts waiting for resolved notifications are
			// held in memory in aggregation groups redundantly.
			a.marker.Delete(alt.Fingerprint())
			a.callback.PostDelete(alt)
		}

		a.mtx.Lock()
		for i, l := range a.listeners {
			select {
			case <-l.done:
				delete(a.listeners, i)
				close(l.alerts)
			default:
				// listener is not closed yet, hence proceed.
			}
		}
		a.mtx.Unlock()
	})
	a.alerts.Start(a.ctx) //nolint:errcheck
	return nil
}

// Stop the alert provider.
func (a *Alerts) Stop(ctx context.Context) error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

// Subscribe returns an iterator over active alerts that have not been
// resolved and successfully notified about.
// They are not guaranteed to be in chronological order.
func (a *Alerts) Subscribe() provider.AlertIterator {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	var (
		done = make(chan struct{})
		al   = a.alerts.List()
		ch   = make(chan *alert.Alert, max(len(al), alertChannelLength))
	)

	for _, a := range al {
		ch <- a
	}

	a.listeners[a.next] = listeningAlerts{alerts: ch, done: done}
	a.next++

	return provider.NewAlertIterator(ch, done, nil)
}

// GetPending returns an iterator over all the alerts that have
// pending notifications.
func (a *Alerts) GetPending() provider.AlertIterator {
	var (
		ch   = make(chan *alert.Alert, alertChannelLength)
		done = make(chan struct{})
	)

	go func() {
		defer close(ch)

		for _, alt := range a.alerts.List() {
			select {
			case ch <- alt:
			case <-done:
				return
			}
		}
	}()

	return provider.NewAlertIterator(ch, done, nil)
}

// Get returns the alert for a given fingerprint.
func (a *Alerts) Get(fp label.Fingerprint) (*alert.Alert, error) {
	return a.alerts.Get(fp)
}

// Put adds the given alert to the set.
func (a *Alerts) Put(alerts ...*alert.Alert) error {
	for _, alt := range alerts {
		fp := alt.Fingerprint()

		existing := false

		// Check that there's an alert existing within the store before
		// trying to merge.
		if old, err := a.alerts.Get(fp); err == nil {
			existing = true

			// Merge alerts if there is an overlap in activity range.
			if (alt.EndsAt.After(old.StartsAt) && alt.EndsAt.Before(old.EndsAt)) ||
				(alt.StartsAt.After(old.StartsAt) && alt.StartsAt.Before(old.EndsAt)) {
				alt = old.Merge(alt)
			}
		}

		if err := a.callback.PreStore(alt, existing); err != nil {
			logger.Error("pre-store callback returned error on set alert", zap.Error(err))
			continue
		}

		if err := a.alerts.Set(alt); err != nil {
			logger.Error("error on set alert", zap.Error(err))
			continue
		}

		a.callback.PostStore(alt, existing)

		a.mtx.Lock()
		for _, l := range a.listeners {
			select {
			case l.alerts <- alt:
			case <-l.done:
			}
		}
		a.mtx.Unlock()
	}

	return nil
}

// count returns the number of non-resolved alerts we currently have stored filtered by the provided state.
func (a *Alerts) count(state alert.AlertState) int {
	var count int
	for _, alt := range a.alerts.List() {
		if alt.Resolved() {
			continue
		}

		status := a.marker.Status(alt.Fingerprint())
		if status.State != state {
			continue
		}

		count++
	}

	return count
}

type noopCallback struct{}

func (n noopCallback) PreStore(_ *alert.Alert, _ bool) error { return nil }
func (n noopCallback) PostStore(_ *alert.Alert, _ bool)      {}
func (n noopCallback) PostDelete(_ *alert.Alert)             {}
