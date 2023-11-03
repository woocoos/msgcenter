package silence

import (
	"bytes"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"io"
	"time"
)

type (
	Entry struct {
		ID        int                `json:"id,omitempty"`
		UpdatedAt time.Time          `json:"created_at,omitempty"`
		Matchers  []*label.Matcher   `json:"matchers,omitempty"`
		StartsAt  time.Time          `json:"starts_at,omitempty"`
		EndsAt    time.Time          `json:"ends_at,omitempty"`
		State     alert.SilenceState `json:"state,omitempty"`
	}
	state map[int]*Entry

	EntryQuery func(*Entry) (bool, error)
)

// getState returns a silence's SilenceState at the given timestamp.
func getState(sil *Entry, ts time.Time) alert.SilenceState {
	if ts.Before(sil.StartsAt) {
		return alert.SilenceStatePending
	}
	if ts.After(sil.EndsAt) {
		return alert.SilenceStateExpired
	}
	return alert.SilenceStateActive
}

func (s state) merge(e *Entry, now time.Time) bool {
	id := e.ID
	if e.EndsAt.Before(now) {
		return false
	}

	prev, ok := s[id]
	if !ok || prev.UpdatedAt.Before(e.UpdatedAt) {
		s[id] = e
		return true
	}
	return false
}

func (s state) Merge(bs []byte) error {
	if len(bs) == 0 {
		return nil
	}
	dec := msgpack.NewDecoder(bytes.NewReader(bs))
	for {
		var e *Entry
		if err := dec.Decode(e); err != nil {
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

func (s state) marshalBinary(e *Entry) ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	if err := enc.Encode(e); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s state) query(qs ...EntryQuery) (res []*Entry, err error) {
	for _, e := range s {
		var ok bool
		for _, q := range qs {
			ok, err = q(e)
			if !ok {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, cloneSilence(e))
		}
	}
	return res, nil
}

// cloneSilence returns a shallow copy of a silence.
func cloneSilence(sil *Entry) *Entry {
	s := *sil
	return &s
}

func QState(now time.Time, states ...alert.SilenceState) EntryQuery {
	return func(e *Entry) (bool, error) {
		s := getState(e, now)
		for _, st := range states {
			if s == st {
				return true, nil
			}
		}
		return false, nil
	}
}

func QIDs(ids []int) EntryQuery {
	return func(e *Entry) (bool, error) {
		for _, id := range ids {
			if e.ID == id {
				return true, nil
			}
		}
		return false, nil
	}
}

// QMatchers returns a EntryQuery that matches silences with the given.if not found matcher return true
func QMatchers(set label.LabelSet, mc matcherCache) EntryQuery {
	return func(e *Entry) (bool, error) {
		m, err := mc.Get(e)
		if err != nil {
			return true, err
		}
		return m.Matches(set), nil
	}
}
