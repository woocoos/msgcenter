package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
	"io"
	"net/http"
)

// Message is the data structure of the webhook message. It is passed to the webhook server.
type Message struct {
	*template.Data

	// The protocol version.
	Version         string `json:"version"`
	GroupKey        string `json:"groupKey"`
	TruncatedAlerts uint64 `json:"truncatedAlerts"`
}

// Notifier email notifier
//
// tmpl include all of receiver's template.
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
		// Webhooks are assumed to respond with 2xx response codes on a successful
		// request and 5xx response codes are assumed to be recoverable.
		retrier: &notify.Retrier{
			CustomDetailsFunc: func(_ int, body io.Reader) string {
				return errDetails(body, cfg.URL.String())
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

	config, err := n.CustomConfig(ctx)
	if err != nil {
		return false, err
	}
	groupKey, err := notify.ExtractGroupKey(ctx)
	if err != nil {
		return false, err
	}

	msg := &Message{
		Version:         "4",
		Data:            data,
		GroupKey:        groupKey.String(),
		TruncatedAlerts: numTruncated,
	}
	var buf bytes.Buffer
	if config.Body == "" { // if body is empty, just send the data
		if err := json.NewEncoder(&buf).Encode(msg); err != nil {
			return false, err
		}
	} else {
		body, err := n.tmpl.ExecuteTextString(config.Body, msg)
		if err != nil {
			return false, err
		}
		buf.WriteString(body)
	}

	url := n.config.URL.String()
	resp, err := n.client.Post(url, "application/json", &buf)
	if err != nil {
		return true, err
	}
	defer resp.Body.Close()

	shouldRetry, err := n.retrier.Check(resp.StatusCode, resp.Body)
	return shouldRetry, err
}

func truncateAlerts(maxAlerts uint64, alerts []*alert.Alert) ([]*alert.Alert, uint64) {
	if maxAlerts != 0 && uint64(len(alerts)) > maxAlerts {
		return alerts[:maxAlerts], uint64(len(alerts)) - maxAlerts
	}

	return alerts, 0
}

func errDetails(body io.Reader, url string) string {
	if body == nil {
		return url
	}
	bs, err := io.ReadAll(body)
	if err != nil {
		return url
	}
	return fmt.Sprintf("%s: %s", url, string(bs))
}
