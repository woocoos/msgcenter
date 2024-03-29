package alert

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/metrics"
	"io"
	"strconv"
	"sync"
	"time"
)

// Marker helps to mark alerts as silenced and/or inhibited.
// All methods are goroutine-safe.
type Marker interface {
	// SetActiveOrSilenced replaces the previous SilencedBy by the provided IDs of
	// active and pending silences, including the version number of the
	// silences state. The set of provided IDs is supposed to represent the
	// complete set of relevant silences. If no active silence IDs are provided and
	// InhibitedBy is already empty, it sets the provided alert to AlertStateActive.
	// Otherwise, it sets the provided alert to AlertStateSuppressed.
	SetActiveOrSilenced(alert label.Fingerprint, version int, activeSilenceIDs, pendingSilenceIDs []int)
	// SetInhibited replaces the previous InhibitedBy by the provided IDs of
	// alerts. In contrast to SetActiveOrSilenced, the set of provided IDs is not
	// expected to represent the complete set of inhibiting alerts. (In
	// practice, this method is only called with one or zero IDs. However,
	// this expectation might change in the future. If no IDs are provided
	// and InhibitedBy is already empty, it sets the provided alert to
	// AlertStateActive. Otherwise, it sets the provided alert to
	// AlertStateSuppressed.
	SetInhibited(alert label.Fingerprint, alertIDs ...string)

	// Count alerts of the given state(s). With no state provided, count all
	// alerts.
	Count(...AlertState) int

	// Status of the given alert.
	Status(label.Fingerprint) MarkerStatus
	// Delete the given alert.
	Delete(label.Fingerprint)

	// Various methods to inquire if the given alert is in a certain
	// AlertState. Silenced also returns all the active and pending
	// silences, while Inhibited may return only a subset of inhibiting
	// alerts. Silenced also returns the version of the silences state the
	// result is based on.
	Unprocessed(label.Fingerprint) bool
	Active(label.Fingerprint) bool
	Silenced(label.Fingerprint) (activeIDs, pendingIDs []int, version int, silenced bool)
	Inhibited(label.Fingerprint) ([]string, bool)
}

// AlertState is used as part of AlertStatus.
type AlertState string

// Possible values for AlertState.
const (
	AlertStateUnprocessed AlertState = "unprocessed"
	AlertStateActive      AlertState = "active"
	AlertStateSuppressed  AlertState = "suppressed"
)

// MarkerStatus stores the state of an alert and, as applicable, the IDs of
// silences silencing the alert and of other alerts inhibiting the alert. Note
// that currently, SilencedBy is supposed to be the complete set of the relevant
// silences while InhibitedBy may contain only a subset of the inhibiting alerts
// – in practice exactly one ID. (This somewhat confusing semantics might change
// in the future.)
type MarkerStatus struct {
	State       AlertState `json:"state"`
	SilencedBy  []int      `json:"silencedBy"`
	InhibitedBy []string   `json:"inhibitedBy"`

	// For internal tracking, not exposed in the API.
	pendingSilences []int
	silencesVersion int
}

// NewMarker returns an instance of a Marker implementation.
func NewMarker(r prometheus.Registerer) Marker {
	m := &memMarker{
		m: map[label.Fingerprint]*MarkerStatus{},
	}
	if metrics.Marker == nil {
		metrics.Marker = metrics.NewMarkerMetrics(r, func(state string) float64 {
			return float64(m.Count(AlertState(state)))
		})
	}

	return m
}

type memMarker struct {
	m map[label.Fingerprint]*MarkerStatus

	mtx sync.RWMutex
}

// Count implements Marker.
func (m *memMarker) Count(states ...AlertState) int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	if len(states) == 0 {
		return len(m.m)
	}

	var count int
	for _, status := range m.m {
		for _, state := range states {
			if status.State == state {
				count++
			}
		}
	}
	return count
}

// SetActiveOrSilenced implements Marker.
func (m *memMarker) SetActiveOrSilenced(alert label.Fingerprint, version int, activeIDs, pendingIDs []int) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	s, found := m.m[alert]
	if !found {
		s = &MarkerStatus{}
		m.m[alert] = s
	}
	s.SilencedBy = activeIDs
	s.pendingSilences = pendingIDs
	s.silencesVersion = version

	// If there are any silence or alert IDs associated with the
	// fingerprint, it is suppressed. Otherwise, set it to
	// AlertStateActive.
	if len(activeIDs) == 0 && len(s.InhibitedBy) == 0 {
		s.State = AlertStateActive
		return
	}

	s.State = AlertStateSuppressed
}

// SetInhibited implements Marker.
func (m *memMarker) SetInhibited(alert label.Fingerprint, ids ...string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	s, found := m.m[alert]
	if !found {
		s = &MarkerStatus{}
		m.m[alert] = s
	}
	s.InhibitedBy = ids

	// If there are any silence or alert IDs associated with the
	// fingerprint, it is suppressed. Otherwise, set it to
	// AlertStateActive.
	if len(ids) == 0 && len(s.SilencedBy) == 0 {
		s.State = AlertStateActive
		return
	}

	s.State = AlertStateSuppressed
}

// Status implements Marker.
func (m *memMarker) Status(alert label.Fingerprint) MarkerStatus {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	if s, found := m.m[alert]; found {
		return *s
	}
	return MarkerStatus{
		State:       AlertStateUnprocessed,
		SilencedBy:  []int{},
		InhibitedBy: []string{},
	}
}

// Delete implements Marker.
func (m *memMarker) Delete(alert label.Fingerprint) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	delete(m.m, alert)
}

// Unprocessed implements Marker.
func (m *memMarker) Unprocessed(alert label.Fingerprint) bool {
	return m.Status(alert).State == AlertStateUnprocessed
}

// Active implements Marker.
func (m *memMarker) Active(alert label.Fingerprint) bool {
	return m.Status(alert).State == AlertStateActive
}

// Inhibited implements Marker.
func (m *memMarker) Inhibited(alert label.Fingerprint) ([]string, bool) {
	s := m.Status(alert)
	return s.InhibitedBy,
		s.State == AlertStateSuppressed && len(s.InhibitedBy) > 0
}

// Silenced returns whether the alert for the given Fingerprint is in the
// Silenced state, any associated silence IDs, and the silences state version
// the result is based on.
func (m *memMarker) Silenced(alert label.Fingerprint) (activeIDs, pendingIDs []int, version int, silenced bool) {
	s := m.Status(alert)
	return s.SilencedBy, s.pendingSilences, s.silencesVersion,
		s.State == AlertStateSuppressed && len(s.SilencedBy) > 0
}

// A Muter determines whether a given label set is muted. Implementers that
// maintain an underlying Marker are expected to update it during a call of
// Mutes.
type Muter interface {
	Mutes(label.LabelSet) bool
}

// A MuteFunc is a function that implements the Muter interface.
type MuteFunc func(label.LabelSet) bool

// Mutes implements the Muter interface.
func (f MuteFunc) Mutes(lset label.LabelSet) bool { return f(lset) }

// SilenceStatus stores the state of a silence.
type SilenceStatus struct {
	State SilenceState `json:"state"`
}

// SilenceState is used as part of SilenceStatus.
type SilenceState string

// Possible values for SilenceState.
const (
	SilenceStateExpired SilenceState = "expired"
	SilenceStateActive  SilenceState = "active"
	SilenceStatePending SilenceState = "pending"
)

func (s *SilenceState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", v)
	}
	*s = SilenceState(str)
	if !s.IsValid() {
		return fmt.Errorf("%s is not a valid SilenceState", str)
	}
	return nil
}

func (s SilenceState) IsValid() bool {
	switch s {
	case SilenceStateExpired, SilenceStateActive, SilenceStatePending:
		return true
	}
	return false
}

func (s SilenceState) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(s.String()))
}

func (s SilenceState) Values() []string {
	return []string{
		string(SilenceStateExpired),
		string(SilenceStateActive),
		string(SilenceStatePending),
	}
}

// String implements fmt.Stringer.
func (s SilenceState) String() string {
	return string(s)
}

// CalcSilenceState returns the SilenceState that a silence with the given start
// and end time would have right now.
func CalcSilenceState(start, end time.Time) SilenceState {
	current := time.Now()
	if current.Before(start) {
		return SilenceStatePending
	}
	if current.Before(end) {
		return SilenceStateActive
	}
	return SilenceStateExpired
}
