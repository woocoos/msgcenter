// Copyright 2023 woocoos
//
// Derived from Prometheus Alertmanager (https://github.com/prometheus/alertmanager).
// Original Copyright 2016-2026 The Prometheus Authors.
// Licensed under the Apache License 2.0.

package silence

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/vmihailenco/msgpack/v5"
	"github.com/woocoos/msgcenter/pkg/label"
)

// SilenceState is used as part of SilenceStatus.
type SilenceState string

// Possible values for SilenceState.
const (
	SilenceStateExpired SilenceState = "expired"
	SilenceStateActive  SilenceState = "active"
	SilenceStatePending SilenceState = "pending"
)

func (s *SilenceState) UnmarshalGQL(v any) error {
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

// SilenceStatus stores the state of a silence.
type SilenceStatus struct {
	State SilenceState `json:"state"`
}

type (
	Entry struct {
		ID          int              `json:"id,omitempty"`
		UpdatedAt   time.Time        `json:"created_at,omitempty"`
		Matchers    []*label.Matcher `json:"matchers,omitempty"`
		MatcherSets label.MatcherSet `json:"matcher_sets,omitempty"`
		StartsAt    time.Time        `json:"starts_at,omitempty"`
		EndsAt      time.Time        `json:"ends_at,omitempty"`
		State       SilenceState     `json:"state,omitempty"`
	}
	state map[int]*Entry

	EntryQuery func(*Entry) (bool, error)
)

// getState returns a silence's SilenceState at the given timestamp.
func getState(sil *Entry, ts time.Time) SilenceState {
	if ts.Before(sil.StartsAt) {
		return SilenceStatePending
	}
	if ts.After(sil.EndsAt) {
		return SilenceStateExpired
	}
	return SilenceStateActive
}

func (s state) merge(e *Entry, now time.Time) (merged bool, added bool) {
	id := e.ID
	if e.EndsAt.Before(now) {
		return false, false
	}

	prev, ok := s[id]
	if !ok || prev.UpdatedAt.Before(e.UpdatedAt) {
		s[id] = e
		return true, !ok
	}
	return false, false
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

// mergeStream decodes and merges entries from binary data, calling onAdded for
// each newly added entry. Returns the decoded entries in order.
func (s state) mergeStream(bs []byte, now time.Time, onAdded func(*Entry)) error {
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
		if _, added := s.merge(e, now); added && onAdded != nil {
			onAdded(e)
		}
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

func QState(now time.Time, states ...SilenceState) EntryQuery {
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

// QMatchers returns a EntryQuery that matches silences with the given.
// It uses the pre-compiled matcherIndex for O(1) lookup.
func QMatchers(set label.LabelSet, mi matcherIndex) EntryQuery {
	return func(e *Entry) (bool, error) {
		m, err := mi.get(e.ID)
		if err != nil {
			return true, err
		}
		return m.Matches(set), nil
	}
}

// sinceQuery is a marker type for QSince. It is handled specially by Silences.query()
// and is not applied as a per-entry filter.
type sinceQuery struct {
	version int
}

// QSince returns silences created after the provided version.
// This enables incremental queries that skip scanning old silences.
func QSince(version int) sinceQuery {
	return sinceQuery{version: version}
}
