package profile

import (
	"fmt"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/httpx"
	"net/textproto"
)

var (
	// DefaultWebhookConfig defines default values for Webhook configurations.
	DefaultWebhookConfig = WebhookConfig{
		SendResolved: true,
	}
	// DefaultEmailConfig defines default values for Email configurations.
	DefaultEmailConfig = EmailConfig{
		SendResolved: false,
		HTML:         `{{ template "email.default.html" . }}`,
		Text:         ``,
		Subject:      `{{ template "email.default.subject" . }}`,
	}
)

// EmailConfig configures notifications via mail.
type EmailConfig struct {
	SendResolved bool `yaml:"sendResolved" json:"sendResolved"`
	// Email address to notify.
	// To 一般采用模板的方式接收动态参数
	To           string            `yaml:"to,omitempty" json:"to,omitempty"`
	From         string            `yaml:"from,omitempty" json:"from,omitempty"`
	Subject      string            `yaml:"subject,omitempty" json:"subject,omitempty"`
	SmartHost    HostPort          `yaml:"smartHost,omitempty" json:"smartHost,omitempty"`
	AuthType     string            `yaml:"authType,omitempty" json:"authType,omitempty"`
	AuthUsername string            `yaml:"authUsername,omitempty" json:"authUsername,omitempty"`
	AuthPassword string            `yaml:"authPassword,omitempty" json:"authPassword,omitempty"`
	AuthSecret   string            `yaml:"authSecret,omitempty" json:"authSecret,omitempty"`
	AuthIdentity string            `yaml:"authIdentity,omitempty" json:"authIdentity,omitempty"`
	Headers      map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
	HTML         string            `yaml:"html,omitempty" json:"html,omitempty"`
	Text         string            `yaml:"text,omitempty" json:"text,omitempty"`
	RequireTLS   bool              `yaml:"requireTls,omitempty" json:"requireTls,omitempty"`
	TLSConfig    *conf.TLS         `yaml:"tls,omitempty" json:"tls,omitempty"`
}

func (ec *EmailConfig) Clone() EmailConfig {
	c := *ec
	c.Headers = CopyMap(ec.Headers)
	return c
}

func (ec *EmailConfig) UnmarshalJSON(bytes []byte) error {
	*ec = DefaultEmailConfig
	p, err := NewJsonParse(bytes)
	if err != nil {
		return err
	}
	if err := p.Unmarshal("", ec); err != nil {
		return err
	}
	if err := ec.Validate(); err != nil {
		return err
	}
	return nil
}

func (ec *EmailConfig) Validate() error {
	if ec.To == "" {
		return fmt.Errorf("missing to address in email config")
	}
	// Header names are case-insensitive, check for collisions.
	normalizedHeaders := map[string]string{}
	for h, v := range ec.Headers {
		normalized := textproto.CanonicalMIMEHeaderKey(h)
		if _, ok := normalizedHeaders[normalized]; ok {
			return fmt.Errorf("duplicate header %q in email config", normalized)
		}
		normalizedHeaders[normalized] = v
	}
	ec.Headers = normalizedHeaders

	return nil
}

// WebhookConfig configures notifications via a generic webhook.
type WebhookConfig struct {
	SendResolved bool                `yaml:"sendResolved" json:"sendResolved"`
	HTTPConfig   *httpx.ClientConfig `yaml:"httpConfig,omitempty" json:"httpConfig,omitempty"`

	// URL to send POST request to.
	URL *URL `yaml:"url" json:"url"`

	// MaxAlerts is the maximum number of alerts to be sent per webhook message.
	// Alerts exceeding this threshold will be truncated. Setting this to 0
	// allows an unlimited number of alerts.
	MaxAlerts uint64 `yaml:"maxAlerts" json:"maxAlerts"`
}

func (c *WebhookConfig) Validate() error {
	if c.URL == nil {
		return fmt.Errorf("one of url or url_file must be configured")
	}
	if c.URL != nil {
		if c.URL.Scheme != "https" && c.URL.Scheme != "http" {
			return fmt.Errorf("scheme required for webhook url")
		}
	}
	if c.HTTPConfig != nil {
		if err := c.HTTPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (c *WebhookConfig) UnmarshalJSON(data []byte) error {
	*c = WebhookConfig{
		SendResolved: true,
	}
	p, err := NewJsonParse(data)
	if err != nil {
		return err
	}
	if err := p.Unmarshal("", c); err != nil {
		return err
	}
	if err := c.Validate(); err != nil {
		return err
	}
	return nil
}
