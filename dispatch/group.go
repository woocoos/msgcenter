package dispatch

import (
	"context"
	"fmt"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/store"
	"go.uber.org/zap"
	"sort"
	"sync"
	"time"
)

// AlertGroup represents how alerts exist within an aggrGroup.
type AlertGroup struct {
	Alerts   alert.Alerts
	Labels   label.LabelSet
	Receiver string
}

type AlertGroups []*AlertGroup

func (ag AlertGroups) Swap(i, j int) { ag[i], ag[j] = ag[j], ag[i] }
func (ag AlertGroups) Less(i, j int) bool {
	if ag[i].Labels.Equal(ag[j].Labels) {
		return len(ag[i].Receiver) < len(ag[j].Receiver)
	}
	return ag[i].Labels.Before(ag[j].Labels)
}
func (ag AlertGroups) Len() int { return len(ag) }

// aggrGroup aggregates alert fingerprints into groups to which a
// common set of routing options applies.
// It emits notifications in the specified intervals.
type aggrGroup struct {
	labels   label.LabelSet
	opts     *RouteOpts
	routeKey string

	alerts  *store.Alerts
	ctx     context.Context
	cancel  func()
	done    chan struct{}
	next    *time.Timer
	timeout func(time.Duration) time.Duration

	logField   zap.Field
	mtx        sync.RWMutex
	hasFlushed bool
}

// newAggrGroup returns a new aggregation group.
func newAggrGroup(ctx context.Context, labels label.LabelSet, r *Route, to func(time.Duration) time.Duration) *aggrGroup {
	if to == nil {
		to = func(d time.Duration) time.Duration { return d }
	}
	ag := &aggrGroup{
		labels:   labels,
		routeKey: r.Key(),
		opts:     &r.RouteOpts,
		timeout:  to,
		alerts:   store.NewAlerts(),
		done:     make(chan struct{}),
	}
	ag.ctx, ag.cancel = context.WithCancel(ctx)

	// Set an initial one-time wait before flushing
	// the first batch of notifications.
	ag.next = time.NewTimer(ag.opts.GroupWait)

	ag.logField = zap.Stringer("aggrGroup", ag)
	return ag
}

func (ag *aggrGroup) fingerprint() label.Fingerprint {
	return ag.labels.Fingerprint()
}

func (ag *aggrGroup) GroupKey() string {
	return fmt.Sprintf("%s:%s", ag.routeKey, ag.labels)
}

func (ag *aggrGroup) String() string {
	return ag.GroupKey()
}

func (ag *aggrGroup) run(nf notifyFunc) {
	defer close(ag.done)
	defer ag.next.Stop()

	for {
		select {
		case now := <-ag.next.C:
			// Give the notifications time until the next flush to
			// finish before terminating them.
			ctx, cancel := context.WithTimeout(ag.ctx, ag.timeout(ag.opts.GroupInterval))

			// The now time we retrieve from the ticker is the only reliable
			// point of time reference for the subsequent notification pipeline.
			// Calculating the current time directly is prone to flaky behavior,
			// which usually only becomes apparent in tests.
			ctx = notify.WithNow(ctx, now)

			// Populate context with information needed along the pipeline.
			ctx = notify.WithGroupKey(ctx, ag.GroupKey())
			ctx = notify.WithGroupLabels(ctx, ag.labels)
			ctx = notify.WithReceiverName(ctx, ag.opts.Receiver)
			ctx = notify.WithRepeatInterval(ctx, ag.opts.RepeatInterval)
			ctx = notify.WithMuteTimeIntervals(ctx, ag.opts.MuteTimeIntervals)
			ctx = notify.WithActiveTimeIntervals(ctx, ag.opts.ActiveTimeIntervals)

			// Wait the configured interval before calling flush again.
			ag.mtx.Lock()
			ag.next.Reset(ag.opts.GroupInterval)
			ag.hasFlushed = true
			ag.mtx.Unlock()

			ag.flush(func(alerts ...*alert.Alert) bool {
				return nf(ctx, alerts...)
			})

			cancel()

		case <-ag.ctx.Done():
			return
		}
	}
}

func (ag *aggrGroup) stop() {
	// Calling cancel will terminate all in-process notifications
	// and the run() loop.
	ag.cancel()
	<-ag.done
}

// insert inserts the alert into the aggregation group.
func (ag *aggrGroup) insert(alert *alert.Alert) {
	if err := ag.alerts.Set(alert); err != nil {
		logger.Error("error on set alert", ag.logField, zap.Error(err))
	}

	// Immediately trigger a flush if the wait duration for this
	// alert is already over.
	ag.mtx.Lock()
	defer ag.mtx.Unlock()
	if !ag.hasFlushed && alert.StartsAt.Add(ag.opts.GroupWait).Before(time.Now()) {
		ag.next.Reset(0)
	}
}

func (ag *aggrGroup) empty() bool {
	return ag.alerts.Empty()
}

// flush sends notifications for all new alerts.
func (ag *aggrGroup) flush(notify func(...*alert.Alert) bool) {
	if ag.empty() {
		return
	}

	var (
		alerts      = ag.alerts.List()
		alertsSlice = make(alert.Alerts, 0, len(alerts))
		now         = time.Now()
	)
	for _, item := range alerts {
		a := *item
		// Ensure that alerts don't resolve as time move forwards.
		if !a.ResolvedAt(now) {
			a.EndsAt = time.Time{}
		}
		alertsSlice = append(alertsSlice, &a)
	}
	sort.Sort(alertsSlice)

	logger.Debug("flushing", ag.logField, zap.String("alerts", fmt.Sprintf("%v", alertsSlice)))

	if notify(alertsSlice...) {
		for _, a := range alertsSlice {
			// Only delete if the fingerprint has not been inserted
			// again since we notified about it.
			fp := a.Fingerprint()
			got, err := ag.alerts.Get(fp)
			if err != nil {
				// This should never happen.
				logger.Error("failed to get alert", ag.logField, zap.Error(err), zap.String("alert", a.String()))
				continue
			}
			if a.Resolved() && got.UpdatedAt == a.UpdatedAt {
				if err := ag.alerts.Delete(fp); err != nil {
					logger.Error("error on delete alert", ag.logField, zap.Error(err), zap.String("alert", a.String()))
				}
			}
		}
	}
}

func getGroupLabels(alert *alert.Alert, route *Route) label.LabelSet {
	groupLabels := label.LabelSet{}
	for ln, lv := range alert.Labels {
		if _, ok := route.RouteOpts.GroupBy[ln]; ok || route.RouteOpts.GroupByAll {
			groupLabels[ln] = lv
		}
	}

	return groupLabels
}
