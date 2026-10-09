package webhook

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
)

func newTestNotifier(t *testing.T, cfg *profile.WebhookConfig) *Notifier {
	t.Helper()
	tmpl, err := template.New()
	require.NoError(t, err)
	n, err := New(cfg, tmpl, nil)
	require.NoError(t, err)
	return n
}

func testAlert() *alert.Alert {
	return &alert.Alert{
		Labels:    label.LabelSet{"alertname": "test", "severity": "critical"},
		StartsAt:  time.Now(),
		EndsAt:    time.Now().Add(time.Hour),
		UpdatedAt: time.Now(),
	}
}

func mustParseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

func testContext() context.Context {
	ctx := context.Background()
	ctx = notify.WithReceiverName(ctx, "test-receiver")
	ctx = notify.WithGroupKey(ctx, "test-group-key")
	ctx = notify.WithGroupLabels(ctx, label.LabelSet{"alertname": "test"})
	return ctx
}

func TestNotify_BasicPost(t *testing.T) {
	t.Parallel()
	var receivedBody []byte
	var receivedContentType string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedContentType = r.Header.Get("Content-Type")
		receivedBody, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	u, _ := url.Parse(srv.URL)
	cfg := &profile.WebhookConfig{URL: (*profile.URL)(u)}
	n := newTestNotifier(t, cfg)

	retry, err := n.Notify(testContext(), testAlert())
	require.NoError(t, err)
	assert.False(t, retry)
	assert.Equal(t, "application/json", receivedContentType)

	var msg Message
	require.NoError(t, json.Unmarshal(receivedBody, &msg))
	assert.Equal(t, "4", msg.Version)
}

func TestNotify_CustomHeaders(t *testing.T) {
	t.Parallel()
	var receivedHeaders http.Header
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedHeaders = r.Header
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	u, _ := url.Parse(srv.URL)
	cfg := &profile.WebhookConfig{
		URL:     (*profile.URL)(u),
		Headers: map[string]string{"X-Custom": "value1", "Authorization": "Bearer token"},
	}
	n := newTestNotifier(t, cfg)

	_, err := n.Notify(testContext(), testAlert())
	require.NoError(t, err)
	assert.Equal(t, "value1", receivedHeaders.Get("X-Custom"))
	assert.Equal(t, "Bearer token", receivedHeaders.Get("Authorization"))
}

func TestNotify_CustomBody(t *testing.T) {
	t.Parallel()
	var receivedBody string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		receivedBody = string(b)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	u, _ := url.Parse(srv.URL)
	cfg := &profile.WebhookConfig{
		URL:  (*profile.URL)(u),
		Body: `{"custom": true}`,
	}
	n := newTestNotifier(t, cfg)

	_, err := n.Notify(testContext(), testAlert())
	require.NoError(t, err)
	assert.Contains(t, receivedBody, `"custom": true`)
}

func TestNotify_Timeout(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	u, _ := url.Parse(srv.URL)
	cfg := &profile.WebhookConfig{
		URL:     (*profile.URL)(u),
		Timeout: 100 * time.Millisecond,
	}
	n := newTestNotifier(t, cfg)

	_, err := n.Notify(testContext(), testAlert())
	assert.Error(t, err)
}

func TestResolveURL_URLFile(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	urlFile := filepath.Join(dir, "webhook.url")
	require.NoError(t, os.WriteFile(urlFile, []byte("  https://example.com/webhook  \n"), 0644))

	cfg := &profile.WebhookConfig{URLFile: urlFile}
	n := newTestNotifier(t, cfg)

	got, err := n.resolveURL(context.Background(), cfg)
	require.NoError(t, err)
	assert.Equal(t, "https://example.com/webhook", got)
}

func TestResolveURL_URLFileMissing(t *testing.T) {
	t.Parallel()
	cfg := &profile.WebhookConfig{URLFile: "/nonexistent/file"}
	n := newTestNotifier(t, cfg)

	_, err := n.resolveURL(context.Background(), cfg)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "reading url_file")
}

func TestResolveURL_EmptyURL(t *testing.T) {
	t.Parallel()
	cfg := &profile.WebhookConfig{}
	n := newTestNotifier(t, cfg)

	_, err := n.resolveURL(context.Background(), cfg)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "empty")
}

func TestRedactURLError(t *testing.T) {
	t.Parallel()
	err := &url.Error{
		Op:  "Post",
		URL: "https://secret:password@example.com/webhook",
		Err: io.EOF,
	}
	redacted := redactURLError(err)
	assert.Contains(t, redacted.Error(), "<redacted>")
	assert.NotContains(t, redacted.Error(), "secret:password")
}

func TestRedactURLError_NonURLError(t *testing.T) {
	t.Parallel()
	err := io.EOF
	redacted := redactURLError(err)
	assert.Equal(t, io.EOF, redacted)
}

func TestDrain(t *testing.T) {
	t.Parallel()
	body := io.NopCloser(strings.NewReader("response body content"))
	resp := &http.Response{Body: body}
	drain(resp)
	b, _ := io.ReadAll(body)
	assert.Empty(t, b)
}

func TestTruncateAlerts(t *testing.T) {
	t.Parallel()
	alerts := []*alert.Alert{testAlert(), testAlert(), testAlert()}

	truncated, num := truncateAlerts(2, alerts)
	assert.Len(t, truncated, 2)
	assert.Equal(t, uint64(1), num)

	truncated, num = truncateAlerts(0, alerts)
	assert.Len(t, truncated, 3)
	assert.Equal(t, uint64(0), num)

	truncated, num = truncateAlerts(10, alerts)
	assert.Len(t, truncated, 3)
	assert.Equal(t, uint64(0), num)
}

func TestNotify_Dingtalk(t *testing.T) {
	t.Parallel()
	var receivedBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedBody, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	u, _ := url.Parse(srv.URL)
	cfg := &profile.WebhookConfig{
		URL:         (*profile.URL)(u),
		ReceiveType: profile.WebhookReceiveTypeDingtalk,
		Body:        `## Alert\n{{ .CommonLabels.alertname }}`,
	}
	n := newTestNotifier(t, cfg)

	retry, err := n.Notify(testContext(), testAlert())
	require.NoError(t, err)
	assert.False(t, retry)

	// Body is rendered and sent directly; dingtalk JSON formatting
	// is handled entirely by the template.
	assert.Contains(t, string(receivedBody), "## Alert")
	assert.Contains(t, string(receivedBody), "test")
}

func TestNotify_Dingtalk_NoBody(t *testing.T) {
	t.Parallel()
	var receivedBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedBody, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	u, _ := url.Parse(srv.URL)
	cfg := &profile.WebhookConfig{
		URL:         (*profile.URL)(u),
		ReceiveType: profile.WebhookReceiveTypeDingtalk,
	}
	n := newTestNotifier(t, cfg)

	retry, err := n.Notify(testContext(), testAlert())
	require.NoError(t, err)
	assert.False(t, retry)

	// Body为空时走默认逻辑，序列化Message而非钉钉格式。
	var msg Message
	require.NoError(t, json.Unmarshal(receivedBody, &msg))
	assert.Equal(t, "4", msg.Version)
}

func TestSignDingTalkURL(t *testing.T) {
	t.Parallel()

	secret := "SECtest123"
	rawURL := "https://oapi.dingtalk.com/robot/send?access_token=abc"

	signed := signDingTalkURL(rawURL, secret)

	// Should contain timestamp and sign parameters.
	assert.Contains(t, signed, "timestamp=")
	assert.Contains(t, signed, "sign=")
	// URL already has ?, should use & separator.
	assert.True(t, strings.HasPrefix(signed, rawURL+"&"))

	// Verify sign is valid HMAC-SHA256 by parsing the raw query string.
	idx := strings.Index(signed, "timestamp=")
	queryPart := signed[idx:]
	params := strings.Split(queryPart, "&")
	var ts, signRaw string
	for _, p := range params {
		kv := strings.SplitN(p, "=", 2)
		if kv[0] == "timestamp" {
			ts = kv[1]
		} else if kv[0] == "sign" {
			signRaw = kv[1]
		}
	}
	assert.NotEmpty(t, ts)
	assert.NotEmpty(t, signRaw)

	// Verify timestamp is a valid millisecond timestamp.
	tsInt, err := strconv.ParseInt(ts, 10, 64)
	require.NoError(t, err)
	assert.Greater(t, tsInt, int64(0))

	// Verify sign matches HMAC-SHA256 computation (signRaw is URL-encoded).
	stringToSign := ts + "\n" + secret
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(stringToSign))
	expectedSign := url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	assert.Equal(t, expectedSign, signRaw)
}

func TestNotify_WithSecret(t *testing.T) {
	t.Parallel()
	var receivedURL string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedURL = r.URL.RequestURI()
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	u, _ := url.Parse(srv.URL)
	cfg := &profile.WebhookConfig{
		URL:         (*profile.URL)(u),
		Secret:      "SECtest123",
		Body:        `{"text":"hello"}`,
		ReceiveType: profile.WebhookReceiveTypeDingtalk,
	}
	n := newTestNotifier(t, cfg)

	retry, err := n.Notify(testContext(), testAlert())
	require.NoError(t, err)
	assert.False(t, retry)

	// Verify timestamp and sign are in the request URL.
	assert.Contains(t, receivedURL, "timestamp=")
	assert.Contains(t, receivedURL, "sign=")
}

func TestWebhookConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     profile.WebhookConfig
		wantErr string
	}{
		{
			name:    "no url or urlFile",
			cfg:     profile.WebhookConfig{},
			wantErr: "one of url or urlFile",
		},
		{
			name: "both url and urlFile",
			cfg: profile.WebhookConfig{
				URL:     (*profile.URL)(mustParseURL("http://example.com")),
				URLFile: "/some/file",
			},
			wantErr: "at most one of url and urlFile",
		},
		{
			name: "valid url",
			cfg: profile.WebhookConfig{
				URL: (*profile.URL)(mustParseURL("http://example.com")),
			},
		},
		{
			name: "valid urlFile only",
			cfg: profile.WebhookConfig{
				URLFile: "/some/file",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
