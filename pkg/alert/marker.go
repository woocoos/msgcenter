package alert

import (
	"context"

	"github.com/woocoos/msgcenter/pkg/label"
)

// AlertState is used as part of MarkerStatus.
type AlertState string

// Possible values for AlertState.
const (
	AlertStateUnprocessed AlertState = "unprocessed"
	AlertStateActive      AlertState = "active"
	AlertStateSuppressed  AlertState = "suppressed"
)

// MarkerStatus stores the state of an alert and, as applicable, the IDs of
// silences silencing the alert and of other alerts inhibiting the alert.
type MarkerStatus struct {
	State       AlertState `json:"state"`
	SilencedBy  []int      `json:"silencedBy"`
	InhibitedBy []string   `json:"inhibitedBy"`
}

// A Muter determines whether a given label set is muted. Implementers that
// maintain an underlying Marker are expected to update it during a call of
// Mutes.
type Muter interface {
	Mutes(ctx context.Context, lset label.LabelSet) bool
}

// A MuteFunc is a function that implements the Muter interface.
type MuteFunc func(ctx context.Context, lset label.LabelSet) bool

// Mutes implements the Muter interface.
func (f MuteFunc) Mutes(ctx context.Context, lset label.LabelSet) bool { return f(ctx, lset) }
