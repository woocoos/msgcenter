package store

import (
	"context"
	"errors"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"sync"
	"time"
)

// ErrNotFound is returned if a Store cannot find the Alert.
var ErrNotFound = errors.New("alert not found")

type Option func(alerts *Alerts)

// WithGCInterval sets the GC interval. The interval must be greater than 1 minute;
func WithGCInterval(interval time.Duration) Option {
	return func(a *Alerts) {
		if a.gcInterval < time.Minute {
			a.gcInterval = time.Minute
		}
		a.gcInterval = interval
	}
}

// Alerts provides lock-coordinated to an in-memory map of alerts, keyed by
// their fingerprint. Resolved alerts are removed from the map based on
// gcInterval. An optional callback can be set which receives a slice of all
// resolved alerts that have been removed.
type Alerts struct {
	sync.Mutex
	c          map[label.Fingerprint]*alert.Alert
	cb         func([]*alert.Alert)
	gcInterval time.Duration
}

// NewAlerts returns a new Alerts struct.
func NewAlerts(opts ...Option) *Alerts {
	a := &Alerts{
		c:          make(map[label.Fingerprint]*alert.Alert),
		cb:         func(_ []*alert.Alert) {},
		gcInterval: time.Minute * 30,
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// SetGCCallback sets a GC callback to be executed after each GC.
func (a *Alerts) SetGCCallback(cb func([]*alert.Alert)) {
	a.Lock()
	defer a.Unlock()

	a.cb = cb
}

func (a *Alerts) Start(ctx context.Context) error {
	t := time.NewTicker(a.gcInterval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			a.gc()
		}
	}
}

func (a *Alerts) Stop(ctx context.Context) error {
	return nil
}

func (a *Alerts) gc() {
	a.Lock()
	defer a.Unlock()

	var resolved []*alert.Alert
	for fp, c := range a.c {
		if c.Resolved() {
			delete(a.c, fp)
			resolved = append(resolved, c)
		}
	}
	a.cb(resolved)
}

// Get returns the Alert with the matching fingerprint, or an error if it is
// not found.
func (a *Alerts) Get(fp label.Fingerprint) (*alert.Alert, error) {
	a.Lock()
	defer a.Unlock()

	alert, prs := a.c[fp]
	if !prs {
		return nil, ErrNotFound
	}
	return alert, nil
}

// Set unconditionally sets the alert in memory.
func (a *Alerts) Set(alert *alert.Alert) error {
	a.Lock()
	defer a.Unlock()

	a.c[alert.Fingerprint()] = alert
	return nil
}

// Delete removes the Alert with the matching fingerprint from the store.
func (a *Alerts) Delete(fp label.Fingerprint) error {
	a.Lock()
	defer a.Unlock()

	delete(a.c, fp)
	return nil
}

// List returns a slice of Alerts currently held in memory.
func (a *Alerts) List() []*alert.Alert {
	a.Lock()
	defer a.Unlock()

	als := make([]*alert.Alert, 0, len(a.c))
	for _, alert := range a.c {
		als = append(als, alert)
	}

	return als
}

// Empty returns true if the store is empty.
func (a *Alerts) Empty() bool {
	a.Lock()
	defer a.Unlock()

	return len(a.c) == 0
}
