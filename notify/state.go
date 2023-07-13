package notify

import (
	"bytes"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/woocoos/msgcenter/pkg/profile"
	"io"
	"time"
)

type (
	LogEntry struct {
		ID             int                  `json:"id,omitempty"`
		ExpiresAt      time.Time            `json:"expires_at,omitempty"`
		UpdatedAt      time.Time            `json:"updated_at,omitempty"`
		GroupKey       string               `json:"group_key,omitempty"`
		Receiver       string               `json:"receiver,omitempty"`
		ReceiverType   profile.ReceiverType `json:"receiver_type,omitempty"`
		Idx            int                  `json:"idx,omitempty"`
		FiringAlerts   []uint64             `json:"firing_alerts,omitempty"`
		ResolvedAlerts []uint64             `json:"resolved_alerts,omitempty"`
	}

	state map[int]*LogEntry

	EntryQuery func(*LogEntry) bool
)

func QReceiver(r *profile.ReceiverKey) EntryQuery {
	return func(e *LogEntry) bool {
		return e.Receiver == r.Name
	}
}

func QGroupKey(gk string) EntryQuery {
	return func(e *LogEntry) bool {
		return e.GroupKey == gk
	}
}

func (s state) query(qs ...EntryQuery) ([]*LogEntry, error) {
	var res []*LogEntry
	for _, e := range s {
		var ok bool
		for _, q := range qs {
			if !q(e) {
				ok = false
				break
			}
			ok = true
		}
		if ok {
			res = append(res, e)
		}
	}
	return res, nil
}

func (s state) clone() state {
	c := make(state, len(s))
	for k, v := range s {
		c[k] = v
	}
	return c
}

// merge returns true or false whether the MeshEntry was merged or
// not. This information is used to decide to gossip the message further.
func (s state) merge(e *LogEntry, now time.Time) bool {
	if e.ExpiresAt.Before(now) {
		return false
	}

	prev, ok := s[e.ID]
	if !ok || prev.UpdatedAt.Before(e.UpdatedAt) {
		s[e.ID] = e
		return true
	}
	return false
}

func (s state) Merge(bs []byte) error {
	dec := msgpack.NewDecoder(bytes.NewReader(bs))
	for {
		var e *LogEntry
		if err := dec.Decode(&e); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		s.merge(e, time.Now())
	}
	return nil
}

func (s state) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)

	for _, e := range s {
		if err := enc.Encode(e); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (s state) marshalBinary(e *LogEntry) ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	if err := enc.Encode(e); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// IsFiringSubset returns whether the given subset is a subset of the alerts
// that were firing at the time of the last notification.
func (m *LogEntry) IsFiringSubset(subset map[uint64]struct{}) bool {
	set := map[uint64]struct{}{}
	for i := range m.FiringAlerts {
		set[m.FiringAlerts[i]] = struct{}{}
	}

	return isSubset(set, subset)
}

// IsResolvedSubset returns whether the given subset is a subset of the alerts
// that were resolved at the time of the last notification.
func (m *LogEntry) IsResolvedSubset(subset map[uint64]struct{}) bool {
	set := map[uint64]struct{}{}
	for i := range m.ResolvedAlerts {
		set[m.ResolvedAlerts[i]] = struct{}{}
	}

	return isSubset(set, subset)
}

func isSubset(set, subset map[uint64]struct{}) bool {
	for k := range subset {
		_, exists := set[k]
		if !exists {
			return false
		}
	}

	return true
}
