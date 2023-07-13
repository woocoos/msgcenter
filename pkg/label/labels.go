package label

import (
	"crypto/sha1"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/gds"
	"sort"
	"strconv"
	"unicode/utf8"
)

const (
	AlertNameLabel = "alertname"
	TenantLabel    = "tenant"
	ToUserIDLabel  = "user"
)

type LabelName string

// IsValid is true iff the label name matches the pattern of LabelNameRE. This
// method, however, does not use LabelNameRE for the check but a much faster
// hardcoded implementation.
func (ln LabelName) IsValid() bool {
	if len(ln) == 0 {
		return false
	}
	for i, b := range ln {
		if !((b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_' || (b >= '0' && b <= '9' && i > 0)) {
			return false
		}
	}
	return true
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ln *LabelName) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if !LabelName(s).IsValid() {
		return fmt.Errorf("%q is not a valid label name", s)
	}
	*ln = LabelName(s)
	return nil
}

// A LabelSet is a collection of LabelName and LabelValue pairs.  The LabelSet
// may be fully-qualified down to the point where it may resolve to a single
// Metric in the data store or not.  All operations that occur within the realm
// of a LabelSet can emit a vector of Metric entities to which the LabelSet may
// match.
type LabelSet map[LabelName]string

// Validate checks whether all names and values in the label set
// are valid.
func (ls LabelSet) Validate() error {
	for ln, lv := range ls {
		if !ln.IsValid() {
			return fmt.Errorf("invalid name %q", ln)
		}
		if !utf8.ValidString(lv) {
			return fmt.Errorf("invalid value %q", lv)
		}
	}
	return nil
}

// Fingerprint returns the LabelSet's fingerprint.
func (ls LabelSet) Fingerprint() Fingerprint {
	return labelSetToFingerprint(ls)
}

// Clone returns a copy of the label set.
func (ls LabelSet) Clone() LabelSet {
	clone := make(LabelSet, len(ls))
	for k, v := range ls {
		clone[k] = v
	}
	return clone
}

// Equal returns true iff both label sets have exactly the same key/value pairs.
func (ls LabelSet) Equal(o LabelSet) bool {
	if len(ls) != len(o) {
		return false
	}
	for ln, lv := range ls {
		olv, ok := o[ln]
		if !ok {
			return false
		}
		if olv != lv {
			return false
		}
	}
	return true
}

func (ls LabelSet) SortedKV() []gds.KeyValue {
	kvs := make([]gds.KeyValue, 0, len(ls))
	for k, v := range ls {
		kvs = append(kvs, gds.KeyValue{Key: string(k), Value: v})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Key < kvs[j].Key
	})
	return kvs
}

// Before compares the metrics, using the following criteria:
//
// If m has fewer labels than o, it is before o. If it has more, it is not.
//
// If the number of labels is the same, the superset of all label names is
// sorted alphanumerically. The first differing label pair found in that order
// determines the outcome: If the label does not exist at all in m, then m is
// before o, and vice versa. Otherwise the label value is compared
// alphanumerically.
//
// If m and o are equal, the method returns false.
func (ls LabelSet) Before(o LabelSet) bool {
	if len(ls) < len(o) {
		return true
	}
	if len(ls) > len(o) {
		return false
	}

	lns := make([]string, 0, len(ls)+len(o))
	for ln := range ls {
		lns = append(lns, string(ln))
	}
	for ln := range o {
		lns = append(lns, string(ln))
	}
	// It's probably not worth it to de-dup lns.
	sort.Strings(lns)
	for _, ln := range lns {
		mlv, ok := ls[LabelName(ln)]
		if !ok {
			return true
		}
		olv, ok := o[LabelName(ln)]
		if !ok {
			return false
		}
		if mlv < olv {
			return true
		}
		if mlv > olv {
			return false
		}
	}
	return false
}

type Fingerprint uint64

func (f Fingerprint) String() string {
	return fmt.Sprintf("%016x", uint64(f))
}

func (f *Fingerprint) Parse(s string) error {
	v, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return err
	}
	*f = Fingerprint(v)
	return nil
}

// StringToFingerprint parses a string representation of a Fingerprint.
func StringToFingerprint(s string) (Fingerprint, error) {
	var f Fingerprint
	err := f.Parse(s)
	return f, err
}

// labelSetToFingerprint works exactly as LabelsToSignature but takes a LabelSet as
// parameter (rather than a label map) and returns a Fingerprint.
func labelSetToFingerprint(ls LabelSet) Fingerprint {
	if len(ls) == 0 {
		return 0
	}
	keys := make([]string, 0, len(ls))
	for labelName := range ls {
		keys = append(keys, string(labelName))
	}
	sort.Strings(keys)

	hash := sha1.New()
	for _, k := range keys {
		hash.Write([]byte(k))
		hash.Write([]byte(ls[LabelName(k)]))
	}
	sum := hash.Sum(nil)
	return Fingerprint(binary.BigEndian.Uint64(sum[:8]))
}
