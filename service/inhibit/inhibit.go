package inhibit

import (
	"context"
	"fmt"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/provider"
	"github.com/woocoos/msgcenter/service/store"
	"go.uber.org/zap"
	"sync"
)

var logger = log.Component("inhibit")

// An Inhibitor determines whether a given label set is muted based on the
// currently active alerts and a set of inhibition rules. It implements the
// Muter interface.
type Inhibitor struct {
	alerts provider.Alerts
	rules  []*InhibitRule
	marker alert.Marker

	mtx    sync.RWMutex
	cancel func() error
}

// NewInhibitor returns a new Inhibitor.
func NewInhibitor(ap provider.Alerts, rs []profile.InhibitRule, mk alert.Marker) *Inhibitor {
	ih := &Inhibitor{
		alerts: ap,
		marker: mk,
	}
	for _, cr := range rs {
		r := NewInhibitRule(cr)
		ih.rules = append(ih.rules, r)
	}
	return ih
}

func (ih *Inhibitor) run(ctx context.Context) {
	it := ih.alerts.Subscribe()
	defer it.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case a := <-it.Next():
			if err := it.Err(); err != nil {
				logger.Error("Error iterating alerts", zap.Error(err))
				continue
			}
			// Update the inhibition rules' cache.
			for _, r := range ih.rules {
				if r.SourceMatchers.Matches(a.Labels) {
					if err := r.scache.Set(a); err != nil {
						logger.Error("error on set alert", zap.Error(err))
					}
				}
			}
		}
	}
}

// Start the Inhibitor's background processing.
func (ih *Inhibitor) Start(ctx context.Context) error {
	var srcs []woocoo.Server
	for _, rule := range ih.rules {
		srcs = append(srcs, rule.scache)
	}
	srcs = append(srcs, ih)
	run, stop := woocoo.MiniApp(ctx, 0, srcs...)
	ih.mtx.Lock()
	ih.cancel = stop
	ih.mtx.Unlock()
	if err := run(); err != nil {
		return fmt.Errorf("error running inhibitor: %w", err)
	}
	return nil
}

// Stop the Inhibitor's background processing.
func (ih *Inhibitor) Stop(ctx context.Context) error {
	if ih == nil {
		return nil
	}

	ih.mtx.RLock()
	defer ih.mtx.RUnlock()
	if ih.cancel != nil {
		return ih.cancel()
	}
	return nil
}

// Mutes returns true if the given label set is muted. It implements the Muter
// interface.
func (ih *Inhibitor) Mutes(lset label.LabelSet) bool {
	fp := lset.Fingerprint()

	for _, r := range ih.rules {
		if !r.TargetMatchers.Matches(lset) {
			// If target side of rule doesn't match, we don't need to look any further.
			continue
		}
		// If we are here, the target side matches. If the source side matches, too, we
		// need to exclude inhibiting alerts for which the same is true.
		if inhibitedByFP, eq := r.hasEqual(lset, r.SourceMatchers.Matches(lset)); eq {
			ih.marker.SetInhibited(fp, inhibitedByFP.String())
			return true
		}
	}
	ih.marker.SetInhibited(fp)

	return false
}

// An InhibitRule specifies that a class of (source) alerts should inhibit
// notifications for another class of (target) alerts if all specified matching
// labels are equal between the two alerts. This may be used to inhibit alerts
// from sending notifications if their meaning is logically a subset of a
// higher-level alert.
type InhibitRule struct {
	// The set of Filters which define the group of source alerts (which inhibit
	// the target alerts).
	SourceMatchers label.Matchers
	// The set of Filters which define the group of target alerts (which are
	// inhibited by the source alerts).
	TargetMatchers label.Matchers
	// A set of label names whose label values need to be identical in source and
	// target alerts in order for the inhibition to take effect.
	Equal map[label.LabelName]struct{}

	// Cache of alerts matching source labels.
	scache *store.Alerts
}

// NewInhibitRule returns a new InhibitRule based on a configuration definition.
func NewInhibitRule(cr profile.InhibitRule) *InhibitRule {
	var (
		sourcem label.Matchers
		targetm label.Matchers
	)
	// We append the new-style matchers. This can be simplified once the deprecated matcher syntax is removed.
	sourcem = append(sourcem, cr.SourceMatchers...)

	// We append the new-style matchers. This can be simplified once the deprecated matcher syntax is removed.
	targetm = append(targetm, cr.TargetMatchers...)

	equal := map[label.LabelName]struct{}{}
	for _, ln := range cr.Equal {
		equal[ln] = struct{}{}
	}

	return &InhibitRule{
		SourceMatchers: sourcem,
		TargetMatchers: targetm,
		Equal:          equal,
		scache:         store.NewAlerts(),
	}
}

// hasEqual checks whether the source cache contains alerts matching the equal
// labels for the given label set. If so, the fingerprint of one of those alerts
// is returned. If excludeTwoSidedMatch is true, alerts that match both the
// source and the target side of the rule are disregarded.
func (r *InhibitRule) hasEqual(lset label.LabelSet, excludeTwoSidedMatch bool) (label.Fingerprint, bool) {
Outer:
	for _, a := range r.scache.List() {
		// The cache might be stale and contain resolved alerts.
		if a.Resolved() {
			continue
		}
		for n := range r.Equal {
			if a.Labels[n] != lset[n] {
				continue Outer
			}
		}
		if excludeTwoSidedMatch && r.TargetMatchers.Matches(a.Labels) {
			continue Outer
		}
		return a.Fingerprint(), true
	}
	return 0, false
}
