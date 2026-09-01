package profile

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"net/textproto"

	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/httpx"
)

var (
	DefaultMessageConfig = MessageConfig{
		SendResolved: false,
		Subject:      `{{ template "message.default.subject" . }}`,
		HTML:         `{{ template "message.default.html" . }}`,
		Extras:       make(map[string]string),
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
	// DefaultUmengConfig defines default values for Umeng Push configurations.
	DefaultUmengConfig = UmengConfig{
		SendResolved: false,
		APIURL:       "https://msgapi.umeng.com/api/send",
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
	To           string   `yaml:"to,omitempty" json:"to,omitempty"`
	From         string   `yaml:"from,omitempty" json:"from,omitempty"`
	Subject      string   `yaml:"subject,omitempty" json:"subject,omitempty"`
	SmartHost    HostPort `yaml:"smartHost,omitempty" json:"smartHost,omitempty"`
	AuthType     string   `yaml:"authType,omitempty" json:"authType,omitempty"`
	AuthUsername string   `yaml:"authUsername,omitempty" json:"authUsername,omitempty"`
	AuthPassword string   `yaml:"authPassword,omitempty" json:"authPassword,omitempty"`
	// AuthPasswordFile is a file containing the auth password.
	AuthPasswordFile string            `yaml:"authPasswordFile,omitempty" json:"authPasswordFile,omitempty"`
	AuthSecret       string            `yaml:"authSecret,omitempty" json:"authSecret,omitempty"`
	AuthIdentity     string            `yaml:"authIdentity,omitempty" json:"authIdentity,omitempty"`
	Headers          map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
	HTML             string            `yaml:"html,omitempty" json:"html,omitempty"`
	Text             string            `yaml:"text,omitempty" json:"text,omitempty"`
	RequireTLS       bool              `yaml:"requireTls,omitempty" json:"requireTls,omitempty"`
	TLSConfig        *conf.TLS         `yaml:"tls,omitempty" json:"tls,omitempty"`
	// ForceImplicitTLS forces implicit TLS (direct TLS connection, typically port 465).
	// When nil, port 465 defaults to implicit TLS.
	ForceImplicitTLS *bool `yaml:"forceImplicitTls,omitempty" json:"forceImplicitTls,omitempty"`
	// Threading configures email threading via References/In-Reply-To headers.
	Threading EmailThreading `yaml:"threading,omitempty" json:"threading,omitempty"`
}

// EmailThreading configures email threading.
type EmailThreading struct {
	// Enabled enables email threading.
	Enabled bool `yaml:"enabled,omitempty" json:"enabled,omitempty"`
	// ThreadByDate groups emails by date. "none" or "daily".
	ThreadByDate string `yaml:"threadByDate,omitempty" json:"threadByDate,omitempty"`
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

// WebhookReceiveType defines the type of webhook receiver.
type WebhookReceiveType string

const (
	// WebhookReceiveTypeGeneric is a generic webhook receiver.
	WebhookReceiveTypeGeneric WebhookReceiveType = ""
	// WebhookReceiveTypeDingtalk is a DingTalk webhook receiver.
	WebhookReceiveTypeDingtalk WebhookReceiveType = "dingtalk"
)

func (r WebhookReceiveType) String() string {
	return string(r)
}

func (r WebhookReceiveType) Values() []string {
	return []string{
		WebhookReceiveTypeGeneric.String(),
		WebhookReceiveTypeDingtalk.String(),
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (r WebhookReceiveType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(r.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (r *WebhookReceiveType) UnmarshalGQL(val any) error {
	str, ok := val.(string)
	if !ok {
		return nil
	}
	*r = WebhookReceiveType(str)
	if err := WebhookReceiveTypeValidator(*r); err != nil {
		return fmt.Errorf("%s is not a valid WebhookReceiveType", str)
	}
	return nil
}

func WebhookReceiveTypeValidator(input WebhookReceiveType) error {
	switch input {
	case WebhookReceiveTypeGeneric, WebhookReceiveTypeDingtalk:
		return nil
	default:
		return fmt.Errorf("invalid enum value for webhook receive type field: %q", input)
	}
}

// WebhookConfig configures notifications via a generic webhook.
//
// Because the configuration of httpconfig is dynamic and requires initialization,
// the original configuration needs to be retained as `HttpConfigOri`.
// When OAuth2 is used, Webhook needs a token storage to store token, such as memory, redis, etc.
// The kind of storage depends on Run Mod: cluster or not.
type WebhookConfig struct {
	SendResolved bool `yaml:"sendResolved" json:"sendResolved"`
	// ReceiveType is the type of webhook receiver (e.g., "dingtalk", "" for generic).
	ReceiveType WebhookReceiveType `yaml:"receiveType,omitempty" json:"receiveType,omitempty"`
	// Secret is the signing secret for webhook receivers that require HMAC-SHA256 signing
	// (e.g., DingTalk custom robots). When set, timestamp and sign query parameters are
	// appended to the webhook URL.
	Secret string `yaml:"secret,omitempty" json:"secret,omitempty"`
	// HTTPConfig configures the HTTP client used to send the webhook. Unmarshalled by custom logic.
	HTTPConfig    *httpx.ClientConfig `yaml:"httpConfig" json:"httpConfig"`
	HttpConfigOri *conf.Configuration `yaml:"-" json:"-"`
	// URL to send POST request to.
	URL *URL `yaml:"url" json:"url"`
	// URLFile is a file containing the URL to send POST request to.
	URLFile string `yaml:"urlFile" json:"urlFile"`
	// MaxAlerts is the maximum number of alerts to be sent per webhook message.
	// Alerts exceeding this threshold will be truncated. Setting this to 0
	// allows an unlimited number of alerts.
	MaxAlerts uint64 `yaml:"maxAlerts" json:"maxAlerts"`
	// Timeout for the webhook request.
	Timeout time.Duration `yaml:"timeout" json:"timeout"`
	// HTTP Headers.
	Headers map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
	// Template for POST message body.
	Subject string `yaml:"subject,omitempty" json:"subject,omitempty"`
	// Body is a template with JSON-string. WebHook uses application/json content type.
	Body string `yaml:"body,omitempty" json:"body,omitempty"`
}

func (c *WebhookConfig) Validate() error {
	if c.URL == nil && c.URLFile == "" {
		return fmt.Errorf("one of url or urlFile must be configured")
	}
	if c.URL != nil && c.URLFile != "" {
		return fmt.Errorf("at most one of url and urlFile must be configured")
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
	if p.IsSet("httpConfig") {
		hp, err := p.Sub("httpConfig")
		if err != nil {
			return err
		}
		c.HttpConfigOri = conf.NewFromParse(hp)
		c.HTTPConfig, err = httpx.NewClientConfig(c.HttpConfigOri)
	}
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

// UmengAppConfig holds credentials for a specific Umeng application.
type UmengAppConfig struct {
	AppKey          string `yaml:"appKey" json:"appKey"`
	AppMasterSecret string `yaml:"appMasterSecret" json:"appMasterSecret"`
	// Platform is the target platform: android, ios, or harmonyos.
	Platform string `yaml:"platform" json:"platform"`
	// AppSet is the business application set name (e.g., "xiaohongshu").
	// Multiple Umeng apps (iOS/Android) can belong to the same AppSet.
	AppSet string `yaml:"appSet" json:"appSet"`
	// AliasType is the alias type for customizedcast push.
	// This should match the alias_type set in the client SDK.
	AliasType string `yaml:"aliasType,omitempty" json:"aliasType,omitempty"`
	// Android-specific: after_open behavior (go_app, go_url, go_activity, go_custom).
	AfterOpen string `yaml:"afterOpen,omitempty" json:"afterOpen,omitempty"`
	// Android-specific: activity to open when after_open=go_activity.
	Activity string `yaml:"activity,omitempty" json:"activity,omitempty"`
}

// UmengConfig configures notifications via Umeng Push API.
type UmengConfig struct {
	SendResolved bool `yaml:"sendResolved" json:"sendResolved"`
	// HTTPConfig configures the HTTP client used to send the request.
	HTTPConfig    *httpx.ClientConfig `yaml:"httpConfig" json:"httpConfig"`
	HttpConfigOri *conf.Configuration `yaml:"-" json:"-"`
	// APIURL is the Umeng Push API endpoint. Defaults to https://msgapi.umeng.com/api/send.
	APIURL string `yaml:"apiURL" json:"apiURL"`
	// Apps maps application name to application-specific credentials.
	// When set, the notifier selects credentials based on the alert's "app" label.
	Apps map[string]*UmengAppConfig `yaml:"apps,omitempty" json:"apps,omitempty"`
	// ProductionMode indicates whether to use production environment.
	ProductionMode *bool `yaml:"productionMode,omitempty" json:"productionMode,omitempty"`
}

func (c *UmengConfig) Validate() error {
	if len(c.Apps) == 0 {
		return fmt.Errorf("missing apps in umeng config")
	}
	// Validate each app config.
	for name, ac := range c.Apps {
		if ac.AppKey == "" {
			return fmt.Errorf("missing appKey for app %q in umeng config", name)
		}
		if ac.AppMasterSecret == "" {
			return fmt.Errorf("missing appMasterSecret for app %q in umeng config", name)
		}
		if ac.Platform == "" {
			return fmt.Errorf("missing platform for app %q in umeng config", name)
		}
		switch ac.Platform {
		case "android", "ios", "harmonyos":
			// valid
		default:
			return fmt.Errorf("invalid platform %q for app %q in umeng config", ac.Platform, name)
		}
	}
	if c.HTTPConfig != nil {
		if err := c.HTTPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (c *UmengConfig) UnmarshalJSON(data []byte) error {
	*c = DefaultUmengConfig
	p, err := NewJsonParse(data)
	if err != nil {
		return err
	}
	if err := p.Unmarshal("", c); err != nil {
		return err
	}
	if p.IsSet("httpConfig") {
		hp, err := p.Sub("httpConfig")
		if err != nil {
			return err
		}
		c.HttpConfigOri = conf.NewFromParse(hp)
		c.HTTPConfig, err = httpx.NewClientConfig(c.HttpConfigOri)
		if err != nil {
			return err
		}
	}
	if err := c.Validate(); err != nil {
		return err
	}
	return nil
}

// Clone returns a deep clone of the UmengConfig.
func (c *UmengConfig) Clone() *UmengConfig {
	clone := *c
	if c.Apps != nil {
		clone.Apps = make(map[string]*UmengAppConfig, len(c.Apps))
		for k, v := range c.Apps {
			pv := *v
			clone.Apps[k] = &pv
		}
	}
	return &clone
}
