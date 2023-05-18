package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	kjson "github.com/knadh/koanf/parsers/json"
	"github.com/tsingsun/woocoo/pkg/conf"
	"os"
	"time"
)

// GlobalConfig Global configuration defines basic configuration items.
type GlobalConfig struct {
	// ResolveTimeout是声明警报的时间，如果警报没有更新，则在此之后警报被解决。
	ResolveTimeout time.Duration `yaml:"resolveTimeout" json:"resolveTimeout"`
	// SMTP 全局配置
	SMTPFrom             string    `yaml:"smtpFrom,omitempty" json:"smtpFrom,omitempty"`
	SMTPSmartHost        *HostPort `yaml:"smtpSmartHost,omitempty" json:"smtpSmartHost,omitempty"`
	SMTPAuthUsername     string    `yaml:"smtpAuthUsername,omitempty" json:"smtpAuthUsername,omitempty"`
	SMTPAuthPassword     string    `yaml:"smtpAuthPassword,omitempty" json:"smtpAuthPassword,omitempty"`
	SMTPAuthPasswordFile string    `yaml:"smtpAuthPasswordFile,omitempty" json:"smtpAuthPasswordFile,omitempty"`
	SMTPAuthSecret       string    `yaml:"smtpAuthSecret,omitempty" json:"smtpAuthSecret,omitempty"`
	SMTPAuthIdentity     string    `yaml:"smtpAuthIdentity,omitempty" json:"smtpAuthIdentity,omitempty"`
	SMTPRequireTLS       bool      `yaml:"smtpRequireTls" json:"smtpRequireTls,omitempty"`
}

// DefaultGlobalConfig returns GlobalConfig with default values.
func DefaultGlobalConfig() GlobalConfig {
	return GlobalConfig{
		ResolveTimeout: 5 * time.Minute,
		SMTPRequireTLS: true,
	}
}

func (g *GlobalConfig) UnmarshalJSON(data []byte) error {
	p, err := NewJsonParse(data)
	if err != nil {
		return err
	}
	if err := p.Unmarshal("", g); err != nil {
		return err
	}
	return nil
}

// Config 为根配置文件.所有配置的变化都会使体现在Config对象中.
type Config struct {
	Global       *GlobalConfig `yaml:"global,omitempty" json:"global,omitempty"`
	Route        *Route        `yaml:"route,omitempty" json:"route,omitempty"`
	InhibitRules []InhibitRule `yaml:"inhibitRules,omitempty" json:"inhibitRules,omitempty"`
	Receivers    []Receiver    `yaml:"receivers,omitempty" json:"receivers,omitempty"`
	Templates    []string      `yaml:"templates" json:"templates"`
	// time_interval指定了一个命名的时间间隔，该时间间隔可以在路由树中引用，以便在一天中的特定时间静音/激活特定路由。
	TimeIntervals []TimeInterval `yaml:"timeIntervals,omitempty" json:"timeIntervals,omitempty"`
}

func NewConfig(cfg *conf.Configuration) (c *Config, err error) {
	filename := cfg.String("config.file")
	if filename == "" {
		return nil, fmt.Errorf("no config file specified")
	}
	// read config file by filename
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Load(bytes)
}

// Load 加载主配置yaml文件
func Load(bytes []byte) (c *Config, err error) {
	sub := conf.NewFromBytes(bytes)
	jsStr, err := sub.ParserOperator().Marshal(kjson.Parser())
	if err != nil {
		return nil, err
	}
	c = &Config{}
	if err = json.Unmarshal(jsStr, c); err != nil {
		return nil, err
	}
	if c.Route == nil {
		return nil, ErrNoRouteProvided
	}
	if c.Route.Continue {
		return nil, errors.New("cannot have continue in root route")
	}
	return c, nil
}

func (c *Config) Validate() error {
	names := map[string]struct{}{}
	for _, rcv := range c.Receivers {
		if _, ok := names[rcv.Name]; ok {
			return fmt.Errorf("notification config name %q is not unique", rcv.Name)
		}
		names[rcv.Name] = struct{}{}
	}
	// The root route must not have any matchers as it is the fallback node
	// for all alerts.
	if c.Route == nil {
		return ErrNoRouteProvided
	}

	if c.Route.Receiver == "" {
		return ErrRootMissReceiver
	}
	if len(c.Route.Matchers) > 0 {
		return ErrRootMustNoMatcher
	}
	if len(c.Route.MuteTimeIntervals) > 0 {
		return ErrRootMustNoMute
	}
	if len(c.Route.ActiveTimeIntervals) > 0 {
		return ErrRootMustNoActive
	}
	// Validate that all receivers used in the routing tree are defined.
	if err := checkReceiver(c.Route, names); err != nil {
		return err
	}

	tiNames := make(map[string]struct{})

	for _, mt := range c.TimeIntervals {
		if len(mt.Name) == 0 {
			return ErrNeedTimeIntervalName
		}
		if _, ok := tiNames[mt.Name]; ok {
			return fmt.Errorf("time interval %q is not unique", mt.Name)
		}
		tiNames[mt.Name] = struct{}{}
	}

	return checkTimeInterval(c.Route, tiNames)
}

func (c *Config) String() string {
	js, _ := json.Marshal(c)
	return string(js)
}

func (c *Config) applyGlobal(names map[string]struct{}) (err error) {
	//names = make(map[string]struct{})
	for _, rcv := range c.Receivers {
		if _, ok := names[rcv.Name]; ok {
			return fmt.Errorf("notification config name %q is not unique", rcv.Name)
		}
		for _, ec := range rcv.EmailConfigs {
			if ec.Headers == nil {
				ec.Headers = make(map[string]string)
			}
			if ec.SmartHost.String() == "" {
				if c.Global.SMTPSmartHost.String() == "" {
					return fmt.Errorf("no global SMTP smarthost set")
				}
				ec.SmartHost = *c.Global.SMTPSmartHost
			}
			if ec.From == "" {
				if c.Global.SMTPFrom == "" {
					return fmt.Errorf("no global SMTP from set")
				}
				ec.From = c.Global.SMTPFrom
			}
			if ec.AuthUsername == "" {
				ec.AuthUsername = c.Global.SMTPAuthUsername
			}
			if ec.AuthPassword == "" {
				ec.AuthPassword = c.Global.SMTPAuthPassword
			}
			if ec.AuthSecret == "" {
				ec.AuthSecret = c.Global.SMTPAuthSecret
			}
			if ec.AuthIdentity == "" {
				ec.AuthIdentity = c.Global.SMTPAuthIdentity
			}
			if ec.RequireTLS {
				ec.RequireTLS = c.Global.SMTPRequireTLS
			}
		}
		names[rcv.Name] = struct{}{}
	}
	return nil
}

func (c *Config) UnmarshalJSON(bytes []byte) error {
	type config Config
	var cc config
	if err := json.Unmarshal(bytes, &cc); err != nil {
		return err
	}
	*c = Config(cc)
	// If a global block was open but empty the default global config is overwritten.
	// We have to restore it here.
	if c.Global == nil {
		c.Global = &GlobalConfig{}
		*c.Global = DefaultGlobalConfig()
	}
	names := map[string]struct{}{}
	if err := c.applyGlobal(names); err != nil {
		return err
	}
	return c.Validate()
}

// DeepClone returns a deep clone of the config.
func (c *Config) DeepClone() (*Config, error) {
	var cc Config
	js, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(js, &cc); err != nil {
		return nil, err
	}
	return &cc, nil
}

// checkReceiver returns an error if a node in the routing tree
// references a receiver not in the given map.
func checkReceiver(r *Route, receivers map[string]struct{}) error {
	for _, sr := range r.Routes {
		if err := checkReceiver(sr, receivers); err != nil {
			return err
		}
	}
	if r.Receiver == "" {
		return nil
	}
	if _, ok := receivers[r.Receiver]; !ok {
		return fmt.Errorf("undefined receiver %q used in route", r.Receiver)
	}
	return nil
}

func checkTimeInterval(r *Route, timeIntervals map[string]struct{}) error {
	for _, sr := range r.Routes {
		if err := checkTimeInterval(sr, timeIntervals); err != nil {
			return err
		}
	}

	for _, ti := range r.ActiveTimeIntervals {
		if _, ok := timeIntervals[ti]; !ok {
			return fmt.Errorf("undefined time interval %q used in route", ti)
		}
	}

	for _, tm := range r.MuteTimeIntervals {
		if _, ok := timeIntervals[tm]; !ok {
			return fmt.Errorf("undefined time interval %q used in route", tm)
		}
	}
	return nil
}
