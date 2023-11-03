package label

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// MatchType is an enum for label matching types.
type MatchType int

// Possible MatchTypes.
const (
	MatchEqual MatchType = iota
	MatchNotEqual
	MatchRegexp
	MatchNotRegexp
)

func (m MatchType) String() string {
	typeToStr := map[MatchType]string{
		MatchEqual:     "=",
		MatchNotEqual:  "!=",
		MatchRegexp:    "=~",
		MatchNotRegexp: "!~",
	}
	if str, ok := typeToStr[m]; ok {
		return str
	}
	panic("unknown match type")
}

func (m *MatchType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", v)
	}
	switch str {
	case "MatchEqual":
		*m = MatchEqual
	case "MatchNotEqual":
		*m = MatchNotEqual
	case "MatchRegexp":
		*m = MatchRegexp
	case "MatchNotRegexp":
		*m = MatchNotRegexp
	default:
		return fmt.Errorf("unknown match type %q", str)
	}
	return nil
}

func (m MatchType) MarshalGQL(w io.Writer) {
	gqlM := ""
	switch m {
	case MatchEqual:
		gqlM = "MatchEqual"
	case MatchNotEqual:
		gqlM = "MatchNotEqual"
	case MatchRegexp:
		gqlM = "MatchRegexp"
	case MatchNotRegexp:
		gqlM = "MatchNotRegexp"
	}
	w.Write([]byte(strconv.Quote(gqlM)))
}

func (m MatchType) MarshalYAML() (interface{}, error) {
	return m.String(), nil
}

type Matcher struct {
	Type  MatchType `json:"type" yaml:"type"`
	Name  string    `json:"name" yaml:"name"`
	Value string    `json:"value" yaml:"value"`
	re    *regexp.Regexp
}

// NewMatcher returns a matcher object.
func NewMatcher(t MatchType, n, v string) (*Matcher, error) {
	m := &Matcher{
		Type:  t,
		Name:  n,
		Value: v,
	}
	if t == MatchRegexp || t == MatchNotRegexp {
		re, err := regexp.Compile("^(?:" + v + ")$")
		if err != nil {
			return nil, err
		}
		m.re = re
	}
	return m, nil
}

func (m *Matcher) String() string {
	return fmt.Sprintf(`%s%s"%s"`, m.Name, m.Type, openMetricsEscape(m.Value))
}

// Matches returns whether the matcher matches the given string value.
func (m *Matcher) Matches(s string) bool {
	switch m.Type {
	case MatchEqual:
		return s == m.Value
	case MatchNotEqual:
		return s != m.Value
	case MatchRegexp:
		return m.re.MatchString(s)
	case MatchNotRegexp:
		return !m.re.MatchString(s)
	}
	panic("labels.Matcher.Matches: invalid match type")
}

func (m *Matcher) UnmarshalText(in []byte) error {
	tmp, err := ParseMatcher(string(in))
	if err != nil {
		return err
	}
	*m = *tmp
	return nil
}

func (m *Matcher) UnmarshalJSON(data []byte) error {
	type cm Matcher
	return json.Unmarshal(data, (*cm)(m))
}

// openMetricsEscape is similar to the usual string escaping, but more
// restricted. It merely replaces a new-line character with '\n', a double-quote
// character with '\"', and a backslash with '\\', which is the escaping used by
// OpenMetrics.
func openMetricsEscape(s string) string {
	r := strings.NewReplacer(
		`\`, `\\`,
		"\n", `\n`,
		`"`, `\"`,
	)
	return r.Replace(s)
}

// Matchers is a slice of Matchers that is sortable, implements Stringer, and
// provides a Matches method to match a LabelSet against all Matchers in the
// slice. Note that some users of Matchers might require it to be sorted.
type Matchers []*Matcher

func (ms Matchers) Len() int      { return len(ms) }
func (ms Matchers) Swap(i, j int) { ms[i], ms[j] = ms[j], ms[i] }

func (ms Matchers) Less(i, j int) bool {
	if ms[i].Name > ms[j].Name {
		return false
	}
	if ms[i].Name < ms[j].Name {
		return true
	}
	if ms[i].Value > ms[j].Value {
		return false
	}
	if ms[i].Value < ms[j].Value {
		return true
	}
	return ms[i].Type < ms[j].Type
}

// Matches checks whether all matchers are fulfilled against the given label set.
func (ms Matchers) Matches(lset LabelSet) bool {
	for _, m := range ms {
		if !m.Matches(lset[LabelName(m.Name)]) {
			return false
		}
	}
	return true
}

func (ms Matchers) String() string {
	var buf bytes.Buffer

	buf.WriteByte('{')
	for i, m := range ms {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(m.String())
	}
	buf.WriteByte('}')

	return buf.String()
}

// UnmarshalJSON implements the json.Unmarshaler interface for Matchers.
func (m *Matchers) UnmarshalJSON(data []byte) error {
	var lines []string
	if err := json.Unmarshal(data, &lines); err != nil {
		return err
	}
	for _, line := range lines {
		pm, err := ParseMatchers(line)
		if err != nil {
			return err
		}
		*m = append(*m, pm...)
	}
	sort.Sort(Matchers(*m))
	return nil
}

// MarshalJSON implements the json.Marshaler interface for Matchers.
func (m Matchers) MarshalJSON() ([]byte, error) {
	if len(m) == 0 {
		return []byte("[]"), nil
	}
	result := make([]string, len(m))
	for i, matcher := range m {
		result[i] = matcher.String()
	}
	return json.Marshal(result)
}
