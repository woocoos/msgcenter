package dispatch

import (
	"context"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/metrics"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/provider"
	"go.uber.org/zap"
	"sort"
	"sync"
	"time"
)

var logger = log.Component("dispatch")

// Dispatcher sorts incoming alerts into aggregation groups and
// assigns the correct notifiers to each.
type Dispatcher struct {
	route   *Route
	alerts  provider.Alerts
	stage   notify.Stage
	metrics *metrics.DispatcherMetrics
	limits  Limits

	timeout func(time.Duration) time.Duration

	mtx                sync.RWMutex
	aggrGroupsPerRoute map[*Route]map[label.Fingerprint]*aggrGroup
	aggrGroupsNum      int

	done   chan struct{}
	ctx    context.Context
	cancel func()
}

// Limits describes limits used by Dispatcher.
type Limits interface {
	// MaxNumberOfAggregationGroups returns max number of aggregation groups that dispatcher can have.
	// 0 or negative value = unlimited.
	// If dispatcher hits this limit, it will not create additional groups, but will log an error instead.
	MaxNumberOfAggregationGroups() int
}

// NewDispatcher returns a new Dispatcher.
func NewDispatcher(
	ap provider.Alerts,
	r *Route,
	s notify.Stage,
	mk alert.Marker,
	to func(time.Duration) time.Duration,
) *Dispatcher {

	disp := &Dispatcher{
		alerts:  ap,
		stage:   s,
		route:   r,
		timeout: to,
		metrics: metrics.Dispatcher,
		limits:  nilLimits{},
	}
	return disp
}

// SetMetrics sets the metrics for the dispatcher.
func (d *Dispatcher) SetMetrics(m *metrics.DispatcherMetrics) {
	d.metrics = m
}

// Run starts dispatching alerts incoming via the updates channel.
func (d *Dispatcher) Run() {
	d.done = make(chan struct{})

	d.mtx.Lock()
	d.aggrGroupsPerRoute = map[*Route]map[label.Fingerprint]*aggrGroup{}
	d.aggrGroupsNum = 0
	d.metrics.AggrGroups.Set(0)
	d.ctx, d.cancel = context.WithCancel(context.Background())
	d.mtx.Unlock()

	d.run(d.alerts.Subscribe())
	close(d.done)
}

func (d *Dispatcher) run(it provider.AlertIterator) {
	cleanup := time.NewTicker(30 * time.Second)
	defer cleanup.Stop()

	defer it.Close()

	for {
		select {
		case alt, ok := <-it.Next():
			if !ok {
				// Iterator exhausted for some reason.
				if err := it.Err(); err != nil {
					logger.Error("error on alert update", zap.Error(err))
				}
				return
			}

			logger.Debug("received alert", zap.Any("alert", alt))

			// Log errors but keep trying.
			if err := it.Err(); err != nil {
				logger.Error("error on alert update", zap.Error(err))
				continue
			}

			now := time.Now()
			for _, r := range d.route.Match(alt.Labels) {
				d.processAlert(alt, r)
			}
			d.metrics.ProcessingDuration.Observe(time.Since(now).Seconds())

		case <-cleanup.C:
			d.mtx.Lock()

			for _, groups := range d.aggrGroupsPerRoute {
				for _, ag := range groups {
					if ag.empty() {
						ag.stop()
						delete(groups, ag.fingerprint())
						d.aggrGroupsNum--
						d.metrics.AggrGroups.Dec()
					}
				}
			}

			d.mtx.Unlock()

		case <-d.ctx.Done():
			return
		}
	}
}

// Groups returns a slice of AlertGroups from the dispatcher's internal state.
func (d *Dispatcher) Groups(routeFilter func(*Route) bool, alertFilter func(*alert.Alert, time.Time) bool) (AlertGroups, map[label.Fingerprint][]string) {
	groups := AlertGroups{}

	d.mtx.RLock()
	defer d.mtx.RUnlock()

	// Keep a list of receivers for an alert to prevent checking each alert
	// again against all routes. The alert has already matched against this
	// route on ingestion.
	receivers := map[label.Fingerprint][]string{}

	now := time.Now()
	for route, ags := range d.aggrGroupsPerRoute {
		if !routeFilter(route) {
			continue
		}

		for _, ag := range ags {
			receiver := route.RouteOpts.Receiver
			alertGroup := &AlertGroup{
				Labels:   ag.labels,
				Receiver: receiver,
			}

			alerts := ag.alerts.List()
			filteredAlerts := make([]*alert.Alert, 0, len(alerts))
			for _, a := range alerts {
				if !alertFilter(a, now) {
					continue
				}

				fp := a.Fingerprint()
				if r, ok := receivers[fp]; ok {
					// Receivers slice already exists. Add
					// the current receiver to the slice.
					receivers[fp] = append(r, receiver)
				} else {
					// First time we've seen this alert fingerprint.
					// Initialize a new receivers slice.
					receivers[fp] = []string{receiver}
				}

				filteredAlerts = append(filteredAlerts, a)
			}
			if len(filteredAlerts) == 0 {
				continue
			}
			alertGroup.Alerts = filteredAlerts

			groups = append(groups, alertGroup)
		}
	}
	sort.Sort(groups)
	for i := range groups {
		sort.Sort(groups[i].Alerts)
	}
	for i := range receivers {
		sort.Strings(receivers[i])
	}

	return groups, receivers
}

// Stop the dispatcher.
func (d *Dispatcher) Stop() {
	if d == nil {
		return
	}
	d.mtx.Lock()
	if d.cancel == nil {
		d.mtx.Unlock()
		return
	}
	d.cancel()
	d.cancel = nil
	d.mtx.Unlock()

	<-d.done
}

// notifyFunc is a function that performs notification for the alert
// with the given fingerprint. It aborts on context cancelation.
// Returns false iff notifying failed.
type notifyFunc func(context.Context, ...*alert.Alert) bool

// processAlert determines in which aggregation group the alert falls
// and inserts it.
func (d *Dispatcher) processAlert(alt *alert.Alert, route *Route) {
	groupLabels := getGroupLabels(alt, route)

	fp := groupLabels.Fingerprint()

	d.mtx.Lock()
	defer d.mtx.Unlock()

	routeGroups, ok := d.aggrGroupsPerRoute[route]
	if !ok {
		routeGroups = map[label.Fingerprint]*aggrGroup{}
		d.aggrGroupsPerRoute[route] = routeGroups
	}

	ag, ok := routeGroups[fp]
	if ok {
		ag.insert(alt)
		return
	}

	// If the group does not exist, create it. But check the limit first.
	if limit := d.limits.MaxNumberOfAggregationGroups(); limit > 0 && d.aggrGroupsNum >= limit {
		d.metrics.AggrGroupLimitReached.Inc()
		logger.Error("too many aggregation groups, cannot create new group for alert",
			zap.Int("groups", d.aggrGroupsNum), zap.Int("limit", limit), zap.String("alert", alt.Name()))
		return
	}

	ag = newAggrGroup(d.ctx, groupLabels, route, d.timeout)
	routeGroups[fp] = ag
	d.aggrGroupsNum++
	d.metrics.AggrGroups.Inc()

	// Insert the 1st alert in the group before starting the group's run()
	// function, to make sure that when the run() will be executed the 1st
	// alert is already there.
	ag.insert(alt)

	go ag.run(func(ctx context.Context, alerts ...*alert.Alert) bool {
		_, _, err := d.stage.Exec(ctx, alerts...)
		if err != nil {
			msg := "notify for alerts failed"
			fs := []zap.Field{zap.Int("num_alerts", len(alerts)), zap.Error(err)}
			if ctx.Err() == context.Canceled {
				// It is expected for the context to be canceled on
				// configuration reload or shutdown. In this case, the
				// message should only be logged at the debug level.
				logger.Debug(msg, fs...)
			}
			logger.Error(msg, fs...)
		}
		return err == nil
	})
}

type nilLimits struct{}

func (n nilLimits) MaxNumberOfAggregationGroups() int { return 0 }
