package profile

import (
	"encoding/json"
	"fmt"
	kfjson "github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"github.com/tsingsun/woocoo/pkg/gds/timeinterval"
	"io"
	"net"
	"net/url"
)

// NewJsonParse returns a new Koanf instance with the JSON parser loaded.
// Use koanf for unmarshalling JSON data to map[string]interface{}.
func NewJsonParse(jsonData []byte) (*koanf.Koanf, error) {
	k := koanf.New(".")
	if err := k.Load(rawbytes.Provider(jsonData), kfjson.Parser()); err != nil {
		return nil, err
	}
	return k, nil
}

// HostPort represents a "host:port" network address.
type HostPort struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}

func (h *HostPort) UnmarshalText(in []byte) (err error) {
	h.Host, h.Port, err = net.SplitHostPort(string(in))
	return err
}

func (h *HostPort) UnmarshalGQL(v interface{}) (err error) {
	hp, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid type %T, expect string", v)
	}
	h.Host, h.Port, err = net.SplitHostPort(hp)
	return
}

func (h HostPort) MarshalGQL(w io.Writer) {
	w.Write([]byte(h.String()))
}

func (h *HostPort) String() string {
	if h == nil {
		return ""
	}
	if h.Host == "" && h.Port == "" {
		return ""
	}
	return fmt.Sprintf(`"%s:%s"`, h.Host, h.Port)
}

type URL url.URL

// MarshalJSON implements the json.Marshaler interface for URL.
func (u URL) MarshalJSON() ([]byte, error) {
	u2 := url.URL(u)
	return json.Marshal(u2.String())
}

// UnmarshalJSON implements the json.Marshaler interface for URL.
func (u *URL) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return nil
	}
	return u.UnmarshalText(data[1 : len(data)-1])
}

func (u *URL) UnmarshalText(data []byte) error {
	t, err := url.Parse(string(data))
	if err != nil {
		return err
	}
	*u = URL(*t)
	return nil
}

func (u URL) String() string {
	u2 := url.URL(u)
	return u2.String()
}

// TimeInterval represents a named set of time intervals for which a route should be muted.
type TimeInterval struct {
	Name          string                      `yaml:"name" json:"name"`
	TimeIntervals []timeinterval.TimeInterval `yaml:"timeIntervals" json:"timeIntervals"`
}

// UnmarshalJSON implements the json interface.
func (t *TimeInterval) UnmarshalJSON(b []byte) error {
	p, err := NewJsonParse(b)
	if err != nil {
		return err
	}
	if err := p.Unmarshal("", t); err != nil {
		return err
	}
	return nil
}

func CopyMap[K, V comparable](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}
