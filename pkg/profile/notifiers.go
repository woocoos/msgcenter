package profile

import (
	"encoding/json"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/httpx"
	"net/textproto"
)

var (
	DefaultMessageConfig = MessageConfig{
		SendResolved: false,
		Subject:      `{{ template "message.default.subject" . }}`,
		HTML:         `{{ template "message.default.html" . }}`,
	}
	// DefaultWebhookConfig defines default values for Webhook configurations.
	DefaultWebhookConfig = WebhookConfig{
		SendResolved: true,
		Subject:      `{{ template "webhook.default.subject" . }}`,
		Body:         `{{ template "webhook.default.body" . }}`,
	}
	// DefaultEmailConfig defines default values for Email configurations.
	DefaultEmailConfig = EmailConfig{
		SendResolved: false,
		HTML:         `{{ template "email.default.html" . }}`,
		Text:         ``,
		Subject:      `{{ template "email.default.subject" . }}`,
	}
)

// MessageConfig configures notifications via internal message.
type MessageConfig struct {
	SendResolved bool `yaml:"sendResolved" json:"sendResolved"`
	// To is user-ids.
	To      string `yaml:"to,omitempty" json:"to,omitempty"`
	Subject string `yaml:"subject,omitempty" json:"subject,omitempty"`
	HTML    string `yaml:"html,omitempty" json:"html,omitempty"`
	Text    string `yaml:"text,omitempty" json:"text,omitempty"`
	// URL is the url of message redirect.
	Redirect string `yaml:"url,omitempty" json:"url,omitempty"`
	// key-values
	Extras map[string]string `yaml:"extras,omitempty" json:"extras,omitempty"`
}

func (c *MessageConfig) UnmarshalJSON(bytes []byte) error {
	*c = DefaultMessageConfig
	type mc MessageConfig
	if err := json.Unmarshal(bytes, (*mc)(c)); err != nil {
		return err
	}
	return nil
}

func (c *MessageConfig) Clone() *MessageConfig {
	cc := *c
	return &cc
}

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

func (c *EmailConfig) Clone() *EmailConfig {
	cc := *c
	cc.Headers = CopyMap(c.Headers)
	return &cc
}

func (c *EmailConfig) UnmarshalJSON(bytes []byte) error {
	*c = DefaultEmailConfig
	p, err := NewJsonParse(bytes)
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

func (c *EmailConfig) Validate() error {
	if c.To == "" {
		return fmt.Errorf("missing to address in email config")
	}
	// Header names are case-insensitive, check for collisions.
	normalizedHeaders := map[string]string{}
	for h, v := range c.Headers {
		normalized := textproto.CanonicalMIMEHeaderKey(h)
		if _, ok := normalizedHeaders[normalized]; ok {
			return fmt.Errorf("duplicate header %q in email config", normalized)
		}
		normalizedHeaders[normalized] = v
	}
	c.Headers = normalizedHeaders

	return nil
}

// WebhookConfig configures notifications via a generic webhook.
//
// Because the configuration of httpconfig is dynamic and requires initialization,
// the original configuration needs to be retained as `HttpConfigOri`.
// When OAuth2 is used, Webhook needs a token storage to store token, such as memory, redis, etc.
// The kind of storage depends on Run Mod: cluster or not.
type WebhookConfig struct {
	SendResolved bool `yaml:"sendResolved" json:"sendResolved"`
	// HTTPConfig configures the HTTP client used to send the webhook. Unmarshalled by custom logic.
	HTTPConfig    *httpx.ClientConfig `yaml:"-" json:"-"`
	HttpConfigOri *conf.Configuration `yaml:"-" json:"-"`
	// URL to send POST request to.
	URL *URL `yaml:"url" json:"url"`
	// MaxAlerts is the maximum number of alerts to be sent per webhook message.
	// Alerts exceeding this threshold will be truncated. Setting this to 0
	// allows an unlimited number of alerts.
	MaxAlerts uint64 `yaml:"maxAlerts" json:"maxAlerts"`
	// HTTP Headers.
	Headers map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
	// Template for POST message body.
	Subject string `yaml:"subject,omitempty" json:"subject,omitempty"`
	// Body is a template with JSON-string. WebHook uses application/json content type.
	Body string `yaml:"body,omitempty" json:"body,omitempty"`
}

func (c *WebhookConfig) Validate() error {
	if c.URL == nil {
		return fmt.Errorf("url must be configured")
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
	*c = DefaultWebhookConfig
	p, err := NewJsonParse(data)
	if err != nil {
		return err
	}
	if err := p.Unmarshal("", c); err != nil {
		return err
	}

	c.HttpConfigOri = conf.NewFromParse(p)
	c.HTTPConfig, err = httpx.NewClientConfig(c.HttpConfigOri)
	if err := c.Validate(); err != nil {
		return err
	}
	return nil
}

// Clone returns a deep clone of the WebhookConfig.
func (c *WebhookConfig) Clone() *WebhookConfig {
	clone := *c
	clone.Headers = CopyMap(c.Headers)
	return &clone
}
