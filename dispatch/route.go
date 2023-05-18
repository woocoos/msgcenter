package dispatch

import (
	"encoding/json"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"go.uber.org/zap"
	"sort"
	"strings"
	"time"
)

// DefaultRouteOpts are the defaulting routing options which apply
// to the root route of a routing tree.
var DefaultRouteOpts = RouteOpts{
	GroupWait:         30 * time.Second,
	GroupInterval:     5 * time.Minute,
	RepeatInterval:    4 * time.Hour,
	GroupBy:           map[label.LabelName]struct{}{},
	GroupByAll:        false,
	MuteTimeIntervals: []string{},
}

// A Route is a node that contains definitions of how to handle alerts.
type Route struct {
	parent *Route

	// The configuration parameters for matches of this route.
	RouteOpts RouteOpts

	// Matchers an alert has to fulfill to match
	// this route.
	Matchers label.Matchers

	// If true, an alert matches further routes on the same level.
	Continue bool

	// Children routes of this route.
	Routes []*Route
}

// NewRoute returns a new route.
func NewRoute(cr *profile.Route, parent *Route) *Route {
	// Create default and overwrite with configured settings.
	opts := DefaultRouteOpts
	if parent != nil {
		opts = parent.RouteOpts
	}

	if len(cr.Receiver) > 0 {
		opts.Receiver = cr.Receiver
	}

	if cr.GroupBy != nil {
		opts.GroupBy = map[label.LabelName]struct{}{}
		for _, ln := range cr.GroupBy {
			opts.GroupBy[ln] = struct{}{}
		}
		opts.GroupByAll = false
	} else {
		if cr.GroupByAll {
			opts.GroupByAll = cr.GroupByAll
		}
	}

	if cr.GroupWait != nil {
		opts.GroupWait = time.Duration(*cr.GroupWait)
	}
	if cr.GroupInterval != nil {
		opts.GroupInterval = time.Duration(*cr.GroupInterval)
	}
	if cr.RepeatInterval != nil {
		opts.RepeatInterval = time.Duration(*cr.RepeatInterval)
	}

	// Build matchers.
	var matchers label.Matchers

	// We append the new-style matchers. This can be simplified once the deprecated matcher syntax is removed.
	matchers = append(matchers, cr.Matchers...)

	sort.Sort(matchers)

	opts.MuteTimeIntervals = cr.MuteTimeIntervals
	opts.ActiveTimeIntervals = cr.ActiveTimeIntervals

	route := &Route{
		parent:    parent,
		RouteOpts: opts,
		Matchers:  matchers,
		Continue:  cr.Continue,
	}

	route.Routes = NewRoutes(cr.Routes, route)

	return route
}

func (r *Route) Apply(cfg *conf.Configuration) {
	retention := cfg.Duration("data.retention")
	r.Walk(func(r *Route) {
		if r.RouteOpts.RepeatInterval > retention {
			log.Component("config").Warn("repeatInterval is greater than the data retention period. It can lead to notifications being repeated more often than expected.",
				zap.Duration("repeatInterval", r.RouteOpts.RepeatInterval),
				zap.Duration("retention", retention),
				zap.String("route", r.Key()),
			)
		}
	})
}

// NewRoutes returns a slice of routes.
func NewRoutes(croutes []*profile.Route, parent *Route) []*Route {
	res := []*Route{}
	for _, cr := range croutes {
		res = append(res, NewRoute(cr, parent))
	}
	return res
}

// Match does a depth-first left-to-right search through the route tree
// and returns the matching routing nodes.
func (r *Route) Match(lset label.LabelSet) []*Route {
	if !r.Matchers.Matches(lset) {
		return nil
	}

	var all []*Route

	for _, cr := range r.Routes {
		matches := cr.Match(lset)

		all = append(all, matches...)

		if matches != nil && !cr.Continue {
			break
		}
	}

	// If no child nodes were matches, the current node itself is a match.
	if len(all) == 0 {
		all = append(all, r)
	}

	return all
}

// Key returns a key for the route. It does not uniquely identify the route in general.
func (r *Route) Key() string {
	b := strings.Builder{}

	if r.parent != nil {
		b.WriteString(r.parent.Key())
		b.WriteRune('/')
	}
	b.WriteString(r.Matchers.String())
	return b.String()
}

// Walk traverses the route tree in depth-first order.
func (r *Route) Walk(visit func(*Route)) {
	visit(r)
	for i := range r.Routes {
		r.Routes[i].Walk(visit)
	}
}

// RouteOpts holds various routing options necessary for processing alerts
// that match a given route.
type RouteOpts struct {
	// The identifiers of the associated notification configuration.
	Receiver string

	// What labels to group alerts by for notifications.
	GroupBy map[label.LabelName]struct{}

	// Use all alert labels to group.
	GroupByAll bool

	// How long to wait to group matching alerts before sending
	// a notification.
	GroupWait      time.Duration
	GroupInterval  time.Duration
	RepeatInterval time.Duration

	// A list of time intervals for which the route is muted.
	MuteTimeIntervals []string

	// A list of time intervals for which the route is active.
	ActiveTimeIntervals []string
}

func (ro *RouteOpts) String() string {
	var labels []label.LabelName
	for ln := range ro.GroupBy {
		labels = append(labels, ln)
	}
	return fmt.Sprintf("<RouteOpts send_to:%q group_by:%q group_by_all:%t timers:%q|%q>",
		ro.Receiver, labels, ro.GroupByAll, ro.GroupWait, ro.GroupInterval)
}

// MarshalJSON returns a JSON representation of the routing options.
func (ro *RouteOpts) MarshalJSON() ([]byte, error) {
	v := struct {
		Receiver       string            `json:"receivers"`
		GroupBy        []label.LabelName `json:"groupBy"`
		GroupByAll     bool              `json:"groupByAll"`
		GroupWait      time.Duration     `json:"groupWait"`
		GroupInterval  time.Duration     `json:"groupInterval"`
		RepeatInterval time.Duration     `json:"repeatInterval"`
	}{
		Receiver:       ro.Receiver,
		GroupByAll:     ro.GroupByAll,
		GroupWait:      ro.GroupWait,
		GroupInterval:  ro.GroupInterval,
		RepeatInterval: ro.RepeatInterval,
	}
	for ln := range ro.GroupBy {
		v.GroupBy = append(v.GroupBy, ln)
	}

	return json.Marshal(&v)
}
