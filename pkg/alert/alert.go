package alert

import (
	"fmt"
	"github.com/woocoos/msgcenter/pkg/label"
	"io"
	"strconv"
	"time"
)

type AlertStatus string

const (
	AlertNone     AlertStatus = "none"
	AlertFiring   AlertStatus = "firing"
	AlertResolved AlertStatus = "resolved"
)

func (a *AlertStatus) UnmarshalGQL(v interface{}) error {
	if v, ok := v.(string); ok {
		switch v {
		case string(AlertNone):
			*a = AlertNone
		case string(AlertFiring):
			*a = AlertFiring
		case string(AlertResolved):
			*a = AlertResolved
		default:
			return fmt.Errorf("invalid AlertStatus %s", v)
		}
		return nil
	}
	return fmt.Errorf("%T is not a string", v)
}

func (a AlertStatus) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(string(a)))
}

func (a AlertStatus) Values() []string {
	return []string{string(AlertNone), string(AlertFiring), string(AlertResolved)}
}

// Alert is a single alert. it is a copy of the Alert struct from the Alertmanager
type Alert struct {
	// Label value pairs for purpose of aggregation, matching, and disposition
	// dispatching. This must minimally include an "alertname" label.
	Labels label.LabelSet `json:"labels"`

	// Extra key/value information which does not define alert identity.
	Annotations label.LabelSet `json:"annotations"`

	// The known time range for this alert. Both ends are optional.
	StartsAt     time.Time `json:"startsAt,omitempty"`
	EndsAt       time.Time `json:"endsAt,omitempty"`
	GeneratorURL string    `json:"generatorURL"`

	// The authoritative timestamp.
	UpdatedAt time.Time
	Timeout   bool
}

// Name returns the name of the alert. It is equivalent to the "alertname" label.
func (a *Alert) Name() string {
	return a.Labels[label.AlertNameLabel]
}

// Resolved returns true if the activity interval ended in the past.
func (a *Alert) Resolved() bool {
	return a.ResolvedAt(time.Now())
}

// ResolvedAt returns true if the activity interval ended before
// the given timestamp.
func (a *Alert) ResolvedAt(ts time.Time) bool {
	if a.EndsAt.IsZero() {
		return false
	}
	return !a.EndsAt.After(ts)
}

// Status returns the status of the alert.
func (a *Alert) Status() AlertStatus {
	if a.Resolved() {
		return AlertResolved
	}
	return AlertFiring
}

func (a *Alert) Fingerprint() label.Fingerprint {
	return a.Labels.Fingerprint()
}

// Merge merges the timespan of two alerts based and overwrites annotations
// based on the authoritative timestamp.  A new alert is returned, the labels
// are assumed to be equal.
func (a *Alert) Merge(o *Alert) *Alert {
	// Let o always be the younger alert.
	if o.UpdatedAt.Before(a.UpdatedAt) {
		return o.Merge(a)
	}

	res := *o

	// Always pick the earliest starting time.
	if a.StartsAt.Before(o.StartsAt) {
		res.StartsAt = a.StartsAt
	}

	if o.Resolved() {
		// The latest explicit resolved timestamp wins if both alerts are effectively resolved.
		if a.Resolved() && a.EndsAt.After(o.EndsAt) {
			res.EndsAt = a.EndsAt
		}
	} else {
		// A non-timeout timestamp always rules if it is the latest.
		if a.EndsAt.After(o.EndsAt) && !a.Timeout {
			res.EndsAt = a.EndsAt
		}
	}

	return &res
}

func (a *Alert) String() string {
	s := fmt.Sprintf("%s[%s]", a.Name(), a.Fingerprint().String()[:7])
	if a.Resolved() {
		return s + "[resolved]"
	}
	return s + "[active]"
}

// Validate checks whether the alert data is inconsistent.
func (a *Alert) Validate() error {
	if a.StartsAt.IsZero() {
		return fmt.Errorf("start time missing")
	}
	if !a.EndsAt.IsZero() && a.EndsAt.Before(a.StartsAt) {
		return fmt.Errorf("start time must be before end time")
	}
	if err := a.Labels.Validate(); err != nil {
		return fmt.Errorf("invalid label set: %s", err)
	}
	if len(a.Labels) == 0 {
		return fmt.Errorf("at least one label pair required")
	}
	if err := a.Annotations.Validate(); err != nil {
		return fmt.Errorf("invalid annotations: %s", err)
	}
	return nil
}

// Clone returns a deep copy of the alert.
func (a *Alert) Clone() *Alert {
	return &Alert{
		Labels:      a.Labels.Clone(),
		Annotations: a.Annotations.Clone(),
		StartsAt:    a.StartsAt,
		EndsAt:      a.EndsAt,
		UpdatedAt:   a.UpdatedAt,
		Timeout:     a.Timeout,
	}
}

// Alerts is a slice of alerts that implements sort.Interface.
type Alerts []*Alert

func (as Alerts) Len() int      { return len(as) }
func (as Alerts) Swap(i, j int) { as[i], as[j] = as[j], as[i] }
func (as Alerts) Less(i, j int) bool {
	if as[i].StartsAt.Before(as[j].StartsAt) {
		return true
	}
	if as[i].EndsAt.Before(as[j].EndsAt) {
		return true
	}
	return as[i].Fingerprint() < as[j].Fingerprint()
}

// HasFiring returns true iff one of the alerts is not resolved.
func (as Alerts) HasFiring() bool {
	for _, a := range as {
		if !a.Resolved() {
			return true
		}
	}
	return false
}

// Status returns StatusFiring iff at least one of the alerts is firing.
func (as Alerts) Status() AlertStatus {
	if as.HasFiring() {
		return AlertFiring
	}
	return AlertResolved
}
