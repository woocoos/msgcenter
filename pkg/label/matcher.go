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
	"unicode"
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

func (m *MatchType) UnmarshalGQL(v any) error {
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

func (m MatchType) MarshalYAML() (any, error) {
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
	if strings.ContainsFunc(m.Name, isReserved) {
		return fmt.Sprintf(`%s%s%s`, strconv.Quote(m.Name), m.Type, strconv.Quote(m.Value))
	}
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
	// JSON object format: {"type": ..., "name": ..., "value": ...}
	if len(data) > 0 && data[0] == '{' {
		type plain Matcher
		return json.Unmarshal(data, (*plain)(m))
	}
	// Text format: alertname="value"
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		return m.UnmarshalText([]byte(s))
	}
	return m.UnmarshalText(data)
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

// MatcherSet is a slice of Matchers pointers that implements OR logic across
// multiple matcher sets. At least one matcher set must match for the MatcherSet
// to match.
type MatcherSet []*Matchers

// UnmarshalJSON implements the json.Unmarshaler interface for Matchers.
func (ms *Matchers) UnmarshalJSON(data []byte) error {
	var lines []string
	if err := json.Unmarshal(data, &lines); err != nil {
		return err
	}
	for _, line := range lines {
		pm, err := ParseMatchers(line)
		if err != nil {
			return err
		}
		*ms = append(*ms, pm...)
	}
	sort.Sort(Matchers(*ms))
	return nil
}

// MarshalJSON implements the json.Marshaler interface for Matchers.
func (ms Matchers) MarshalJSON() ([]byte, error) {
	if len(ms) == 0 {
		return []byte("[]"), nil
	}
	result := make([]string, len(ms))
	for i, matcher := range ms {
		result[i] = matcher.String()
	}
	return json.Marshal(result)
}

// Matches checks whether at least one matcher set is fulfilled against the given
// label set (OR logic across matcher sets, AND logic within each set).
func (ms MatcherSet) Matches(lset LabelSet) bool {
	for _, matchers := range ms {
		if (*matchers).Matches(lset) {
			return true
		}
	}
	return false
}

// This is copied from matcher/parse/lexer.go. It will be removed when
// the transition window from classic matchers to UTF-8 matchers is complete,
// as then we can use double quotes when printing the label name for all
// matchers. Until then, the classic parser does not understand double quotes
// around the label name, so we use this function as a heuristic to tell if
// the matcher was parsed with the UTF-8 parser or the classic parser.
func isReserved(r rune) bool {
	return unicode.IsSpace(r) || strings.ContainsRune("{}!=~,\\\"'`", r)
}
