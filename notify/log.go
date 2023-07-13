package notify

import (
	"context"
	"errors"
	"github.com/tsingsun/members"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/msgcenter/metrics"
	"github.com/woocoos/msgcenter/pkg/profile"
	"go.uber.org/zap"
	"sync"
	"time"
)

var (
	defaultLogOptions = LogOptions{
		Retention:           time.Hour * 120,
		MaintenanceInterval: time.Minute * 15,
	}
)

type (
	// NLogCallback is the interface that must be implemented by the notification log callback.
	NLogCallback interface {
		LoadData() ([]*LogEntry, error)
		CreateLog(ctx context.Context, r *profile.ReceiverKey, gkey string, firingAlerts, resolvedAlerts []uint64, expiresAt time.Time) (int, error)
		EvictLog(ctx context.Context, ids []int)
	}
	// Log holds the notification log state for alerts that have been notified.
	Log struct {
		LogOptions

		st state
		mu sync.RWMutex
	}
	// LogOptions configures a new Log implementation.
	LogOptions struct {
		// Retention is the duration for which notifications log are kept in the state. default 120h
		Retention    time.Duration
		NLogCallback NLogCallback
		Spreader     members.Spreader
		// MaintenanceInterval alert manager garbage collects the notification log state at the given interval.
		// default 15m
		MaintenanceInterval time.Duration
		// MaintenanceFunc represents the function to run as part of the periodic maintenance for the nflog.
		MaintenanceFunc MaintenanceFunc
	}
	// MaintenanceFunc represents the function to run as part of the periodic maintenance for the nflog.
	MaintenanceFunc func() error
)

func (o *LogOptions) validate() error {
	if o.MaintenanceInterval == 0 {
		return errors.New("interval or stop signal are missing - not running maintenance")
	}

	return nil
}

// NewLog creates a new notification log based on the provided options.
func NewLog(cfg *conf.Configuration) (*Log, error) {
	o := defaultLogOptions
	if err := cfg.Unmarshal(&o); err != nil {
		return nil, err
	}
	if err := o.validate(); err != nil {
		return nil, err
	}

	l := &Log{
		LogOptions: o,
		st:         state{},
	}

	if metrics.Marker == nil {
		metrics.Nflog = metrics.NewNflogMetrics()
	}

	if l.MaintenanceFunc == nil {
		l.MaintenanceFunc = func() error {
			if _, err := l.GC(); err != nil {
				return err
			}

			return nil
		}
	}

	if l.NLogCallback != nil {
		if err := l.loadData(); err != nil {
			return nil, err
		}
	}

	return l, nil
}

// loadData loads all silences data from DataLoader
func (l *Log) loadData() error {
	datas, err := l.NLogCallback.LoadData()
	if err != nil {
		return err
	}
	st := make(state)
	for _, e := range datas {
		st[e.ID] = e
	}
	l.mu.Lock()
	l.st = st
	l.mu.Unlock()

	return nil
}

func (l *Log) Start(ctx context.Context) error {
	t := time.NewTicker(l.MaintenanceInterval)

	runMaintenance := func(do MaintenanceFunc) error {
		metrics.Nflog.MaintenanceTotal.Inc()
		if err := do(); err != nil {
			metrics.Nflog.MaintenanceErrorsTotal.Inc()
			return err
		}
		return nil
	}
	logger.Debug("notification log maintenance started")
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		case <-t.C:
			if err := runMaintenance(l.MaintenanceFunc); err != nil {
				logger.Error("Running maintenance failed", zap.Error(err))
			}
		}
	}

	if err := runMaintenance(l.MaintenanceFunc); err != nil {
		logger.Error("Running final maintenance failed", zap.Error(err))
	}
	return nil
}

func (l *Log) Stop(ctx context.Context) error {
	return nil
}

func (l *Log) Log(ctx context.Context, r *profile.ReceiverKey, gkey string, firingAlerts, resolvedAlerts []uint64, expiry time.Duration) error {
	// Write all st with the same timestamp.
	now := time.Now()
	expiresAt := now.Add(l.Retention)
	if expiry > 0 && l.Retention > expiry {
		expiresAt = now.Add(expiry)
	}
	//key := stateKey(gkey, r)
	key, err := l.NLogCallback.CreateLog(ctx, r, gkey, firingAlerts, resolvedAlerts, expiresAt)
	if err != nil {
		return err
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if prevle, ok := l.st[key]; ok {
		// LogEntry already exists, only overwrite if timestamp is newer.
		// This may happen with raciness or clock-drift across AM nodes.
		if prevle.UpdatedAt.After(now) {
			return nil
		}
	}

	if l.Spreader != nil {
		b, err := l.st.marshalBinary(&LogEntry{
			ID:             key,
			GroupKey:       gkey,
			Receiver:       r.Integration,
			ExpiresAt:      expiresAt,
			UpdatedAt:      now,
			FiringAlerts:   firingAlerts,
			ResolvedAlerts: resolvedAlerts,
		})
		if err != nil {
			return err
		}
		return l.Spreader.Broadcast(b)
	}
	return nil
}

// GC implements the Log interface.
func (l *Log) GC() (int, error) {
	start := time.Now()
	defer func() { metrics.Nflog.GCDuration.Observe(time.Since(start).Seconds()) }()

	now := time.Now()
	var n int

	l.mu.Lock()
	defer l.mu.Unlock()

	dl := make([]int, 0, len(l.st))
	for k, le := range l.st {
		if le.ExpiresAt.IsZero() {
			return n, errors.New("unexpected zero expiration timestamp")
		}
		if !le.ExpiresAt.After(now) {
			delete(l.st, k)
			dl = append(dl, k)
			n++
		}
	}
	if len(dl) > 0 {
		l.NLogCallback.EvictLog(context.Background(), dl)
	}

	return n, nil
}

// Query implements the Log interface.
func (l *Log) Query(params ...EntryQuery) ([]*LogEntry, error) {
	start := time.Now()
	metrics.Nflog.QueriesTotal.Inc()

	entries, err := func() ([]*LogEntry, error) {
		l.mu.RLock()
		defer l.mu.RUnlock()
		return l.st.query(params...)
	}()
	if err != nil {
		metrics.Nflog.QueryErrorsTotal.Inc()
	}
	metrics.Nflog.QueryDuration.Observe(time.Since(start).Seconds())
	return entries, err
}

func (l *Log) Name() string {
	return "nflog"
}

// MarshalBinary serializes all contents of the notification log.
func (l *Log) MarshalBinary() ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.st.MarshalBinary()
}

// Merge merges notification log state received from the cluster with the local state.
func (l *Log) Merge(b []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.st.Merge(b)
}
