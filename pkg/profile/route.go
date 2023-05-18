package profile

import (
	"fmt"
	"github.com/woocoos/msgcenter/pkg/label"
	"time"
)

// A Route is a node that contains definitions of how to handle alerts.
type Route struct {
	Name                string            `yaml:"-" json:"-"`
	Receiver            string            `yaml:"receiver,omitempty" json:"receiver,omitempty"`
	GroupBy             []label.LabelName `yaml:"groupBy,omitempty" json:"groupBy,omitempty"`
	GroupByAll          bool              `yaml:"-" json:"-"`
	Matchers            label.Matchers    `yaml:"matchers,omitempty" json:"matchers,omitempty"`
	MuteTimeIntervals   []string          `yaml:"muteTimeIntervals,omitempty" json:"muteTimeIntervals,omitempty"`
	ActiveTimeIntervals []string          `yaml:"activeTimeIntervals,omitempty" json:"activeTimeIntervals,omitempty"`
	Continue            bool              `yaml:"continue" json:"continue,omitempty"`
	Routes              []*Route          `yaml:"routes,omitempty" json:"routes,omitempty"`

	GroupWait      *time.Duration `yaml:"groupWait,omitempty" json:"groupWait,omitempty"`
	GroupInterval  *time.Duration `yaml:"groupInterval,omitempty" json:"groupInterval,omitempty"`
	RepeatInterval *time.Duration `yaml:"repeatInterval,omitempty" json:"repeatInterval,omitempty"`
}

func (r *Route) Validate() error {
	groupBy := map[label.LabelName]struct{}{}
	var groupBytmp []label.LabelName
	for _, l := range r.GroupBy {
		if l == "..." {
			r.GroupByAll = true
			break
		}
		if !l.IsValid() {
			return fmt.Errorf("invalid label name %q in group_by list", l)
		}
		groupBytmp = append(groupBytmp, l)
		if _, ok := groupBy[l]; ok {
			return fmt.Errorf("duplicated label %q in group_by", l)
		}
		groupBy[l] = struct{}{}
	}
	r.GroupBy = groupBytmp
	if r.GroupByAll && len(r.GroupBy) > 1 {
		return fmt.Errorf("cannot have wildcard group_by (`...`) and other other labels at the same time")
	}

	if r.GroupInterval != nil && *r.GroupInterval == time.Duration(0) {
		return fmt.Errorf("groupInterval cannot be zero")
	}
	if r.RepeatInterval != nil && *r.RepeatInterval == time.Duration(0) {
		return fmt.Errorf("repeatInterval cannot be zero")
	}
	for _, route := range r.Routes {
		if err := route.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Route) UnmarshalJSON(data []byte) error {
	p, err := NewJsonParse(data)
	if err != nil {
		return err
	}
	if err := p.Unmarshal("", r); err != nil {
		return err
	}
	if err := r.Validate(); err != nil {
		return err
	}
	return nil
}
