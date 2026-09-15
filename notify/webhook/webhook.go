package webhook

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
)

// Message is the data structure of the webhook message. It is passed to the webhook server.
type Message struct {
	*template.Data

	// The protocol version.
	Version         string `json:"version"`
	GroupKey        string `json:"groupKey"`
	TruncatedAlerts uint64 `json:"truncatedAlerts"`
}

// Notifier sends notifications via a generic webhook.
type Notifier struct {
	config        *profile.WebhookConfig
	tmpl          *template.Template
	customTplFunc notify.CustomerConfigFunc[profile.WebhookConfig]
	client        *http.Client
	retrier       *notify.Retrier
}

// New returns a new Webhook.
func New(cfg *profile.WebhookConfig, tmpl *template.Template,
	fn notify.CustomerConfigFunc[profile.WebhookConfig],
) (*Notifier, error) {
	nf := &Notifier{
		config:        cfg,
		tmpl:          tmpl,
		customTplFunc: fn,
		retrier: &notify.Retrier{
			CustomDetailsFunc: func(_ int, body io.Reader) string {
				return errDetails(body)
			},
		},
	}
	if cfg.HTTPConfig == nil {
		nf.client = http.DefaultClient
	} else {
		httpClient, err := cfg.HTTPConfig.Client(context.Background(), nil)
		if err != nil {
			return nil, err
		}
		nf.client = httpClient
	}
	return nf, nil
}

func (n *Notifier) SendResolved() bool {
	return n.config.SendResolved
}

// CustomConfig returns a custom config for the notifier.
func (n *Notifier) CustomConfig(ctx context.Context) (*profile.WebhookConfig, error) {
	if n.customTplFunc == nil {
		return n.config, nil
	}
	labels, ok := notify.GroupLabels(ctx)
	if !ok {
		return n.config, nil
	}
	cfg := n.config.Clone()
	err := n.customTplFunc(ctx, cfg, labels)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Notify implements the Notifier interface.
func (n *Notifier) Notify(ctx context.Context, alerts ...*alert.Alert) (bool, error) {
	alerts, numTruncated := truncateAlerts(n.config.MaxAlerts, alerts)
	data := notify.GetTemplateData(ctx, n.tmpl, alerts)
	var err error
	tmpl := notify.TmplText(n.tmpl, data, &err)
	if err != nil {
		return false, err
	}

	config, err := n.CustomConfig(ctx)
	if err != nil {
		return false, err
	}
	groupKey, err := notify.ExtractGroupKey(ctx)
	if err != nil {
		return false, err
	}
	url, err := n.resolveURL(ctx, config)
	if err != nil {
		return false, err
	}
	sub := tmpl(config.Subject)
	if err != nil {
		return false, fmt.Errorf("execute 'subject' template: %w", err)
	}

	msg := &Message{
		Version:         "4",
		Data:            data,
		GroupKey:        groupKey.String(),
		TruncatedAlerts: numTruncated,
	}

	var buf bytes.Buffer
	if config.Body == "" {
		if err := json.NewEncoder(&buf).Encode(msg); err != nil {
			return false, err
		}
	} else {
		body, err := n.tmpl.ExecuteTextString(config.Body, msg)
		if err != nil {
			return false, err
		}
		if config.ReceiveType == profile.WebhookReceiveTypeDingtalk {
			mdMsg := map[string]any{
				"msgtype": "markdown",
				"markdown": map[string]string{
					"title": sub,
					"text":  body,
				},
			}
			if err := json.NewEncoder(&buf).Encode(mdMsg); err != nil {
				return false, err
			}

			// Append timestamp and sign query parameters when secret is configured.
			if config.Secret != "" {
				url = signDingTalkURL(url, config.Secret)
			}
		} else {
			buf.WriteString(body)
		}
	}

	// Apply timeout if configured.
	if config.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, config.Timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return true, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Set custom headers.
	for k, v := range config.Headers {
		req.Header.Set(k, v)
	}

	resp, err := n.client.Do(req)
	if err != nil {
		return true, redactURLError(err)
	}
	defer drain(resp)

	shouldRetry, err := n.retrier.Check(resp.StatusCode, resp.Body)
	return shouldRetry, err
}

// resolveURL determines the webhook URL from config.URL or config.URLFile,
// with optional template rendering.
func (n *Notifier) resolveURL(ctx context.Context, config *profile.WebhookConfig) (string, error) {
	var rawURL string
	if config.URLFile != "" {
		content, err := os.ReadFile(config.URLFile)
		if err != nil {
			return "", fmt.Errorf("reading url_file %q: %w", config.URLFile, err)
		}
		rawURL = strings.TrimSpace(string(content))
	} else if config.URL != nil {
		rawURL = config.URL.String()
	}

	// Template rendering on URL.
	if strings.Contains(rawURL, "{{") {
		rendered, err := n.tmpl.ExecuteTextString(rawURL, notify.GetTemplateData(ctx, n.tmpl, nil))
		if err != nil {
			return "", fmt.Errorf("rendering webhook url template: %w", err)
		}
		rawURL = rendered
	}

	if rawURL == "" {
		return "", fmt.Errorf("webhook url is empty")
	}
	return rawURL, nil
}

func truncateAlerts(maxAlerts uint64, alerts []*alert.Alert) ([]*alert.Alert, uint64) {
	if maxAlerts != 0 && uint64(len(alerts)) > maxAlerts {
		return alerts[:maxAlerts], uint64(len(alerts)) - maxAlerts
	}
	return alerts, 0
}

func errDetails(body io.Reader) string {
	if body == nil {
		return ""
	}
	bs, err := io.ReadAll(body)
	if err != nil {
		return ""
	}
	return string(bs)
}

// drain consumes and closes the response body to enable connection reuse.
func drain(resp *http.Response) {
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
}

// redactURLError replaces the URL in a *url.Error with "<redacted>" to prevent
// leaking sensitive information in error messages.
func redactURLError(err error) error {
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		urlErr.URL = "<redacted>"
	}
	return err
}

// signDingTalkURL appends timestamp and sign query parameters to the URL using
// HMAC-SHA256 signing (compatible with DingTalk custom robot security settings).
func signDingTalkURL(rawURL, secret string) string {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	stringToSign := timestamp + "\n" + secret
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(stringToSign))
	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	sep := "&"
	if strings.Contains(rawURL, "?") {
		if !strings.HasSuffix(rawURL, "&") {
			sep = "&"
		}
	} else {
		sep = "?"
	}
	return rawURL + sep + "timestamp=" + timestamp + "&sign=" + sign
}
