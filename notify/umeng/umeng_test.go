package umeng

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
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

func newTestNotifier(t *testing.T, cfg *profile.UmengConfig) *Notifier {
	t.Helper()
	tmpl, err := template.New()
	require.NoError(t, err)
	n, err := New(cfg, tmpl, nil, nil)
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

func testContext() context.Context {
	ctx := context.Background()
	ctx = notify.WithReceiverName(ctx, "test-receiver")
	ctx = notify.WithGroupKey(ctx, "test-group-key")
	ctx = notify.WithGroupLabels(ctx, label.LabelSet{"alertname": "test"})
	return ctx
}

func TestMd5Hex(t *testing.T) {
	t.Parallel()
	input := "app_key1234567890master_secret"
	h := md5.New()
	h.Write([]byte(input))
	expected := hex.EncodeToString(h.Sum(nil))
	assert.Equal(t, expected, md5Hex(input))
}

func TestMd5Hex_Empty(t *testing.T) {
	t.Parallel()
	h := md5.New()
	h.Write([]byte(""))
	expected := hex.EncodeToString(h.Sum(nil))
	assert.Equal(t, expected, md5Hex(""))
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

func newTestUmengConfig(appKey, appSecret, platform string) *profile.UmengConfig {
	return &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test": {
				AppKey:          appKey,
				AppMasterSecret: appSecret,
				Platform:        platform,
			},
		},
	}
}

func TestNotifier_SendResolved(t *testing.T) {
	t.Parallel()
	n := newTestNotifier(t, &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test": {
				AppKey:          "test_key",
				AppMasterSecret: "test_secret",
				Platform:        "android",
			},
		},
		SendResolved: true,
	})
	assert.True(t, n.SendResolved())

	n2 := newTestNotifier(t, &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test": {
				AppKey:          "test_key",
				AppMasterSecret: "test_secret",
				Platform:        "android",
			},
		},
		SendResolved: false,
	})
	assert.False(t, n2.SendResolved())
}

func TestNew_DefaultClient(t *testing.T) {
	t.Parallel()
	cfg := newTestUmengConfig("test_key", "test_secret", "android")
	n, err := New(cfg, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, http.DefaultClient, n.client)
}

func TestBuildRequest_Broadcast(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New()
	require.NoError(t, err)
	cfg := &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test": {
				AppKey:          "test_key",
				AppMasterSecret: "test_secret",
				Platform:        "android",
			},
		},
		ProductionMode: boolPtr(true),
	}
	n, err := New(cfg, tmpl, nil, nil)
	require.NoError(t, err)

	msg := &Message{
		Data: notify.GetTemplateData(testContext(), tmpl, []*alert.Alert{testAlert()}),
	}
	req, err := n.buildRequestForApp(cfg, msg, "")
	require.NoError(t, err)

	assert.Equal(t, "test_key", req.AppKey)
	assert.Equal(t, "broadcast", req.Type)
	// Title and text now use default templates
	assert.NotNil(t, req.Payload["body"])
	body := req.Payload["body"].(map[string]any)
	assert.NotNil(t, body["title"])
	assert.NotNil(t, body["text"])
	assert.Equal(t, "notification", req.Payload["display_type"])
	assert.NotNil(t, req.Production)
	assert.True(t, *req.Production)
	assert.Empty(t, req.DeviceTokens)
}

func TestBuildRequest_Unicast(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New()
	require.NoError(t, err)
	cfg := newTestUmengConfig("test_key", "test_secret", "android")
	n, err := New(cfg, tmpl, nil, nil)
	require.NoError(t, err)

	msg := &Message{
		Data:         notify.GetTemplateData(testContext(), tmpl, []*alert.Alert{testAlert()}),
		DeviceTokens: []string{"token123"},
	}
	req, err := n.buildRequestForApp(cfg, msg, "")
	require.NoError(t, err)

	assert.Equal(t, "unicast", req.Type)
	assert.Equal(t, "token123", req.DeviceTokens)
}

func TestBuildRequest_CustomPayload(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New()
	require.NoError(t, err)
	cfg := newTestUmengConfig("test_key", "test_secret", "android")
	n, err := New(cfg, tmpl, nil, nil)
	require.NoError(t, err)

	msg := &Message{
		Data: notify.GetTemplateData(testContext(), tmpl, []*alert.Alert{testAlert()}),
	}
	req, err := n.buildRequestForApp(cfg, msg, "")
	require.NoError(t, err)

	// Default payload is used since custom payload is no longer supported
	assert.Equal(t, "notification", req.Payload["display_type"])
}

func TestBuildRequest_Sign(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New()
	require.NoError(t, err)
	cfg := newTestUmengConfig("my_app_key", "my_secret", "android")
	n, err := New(cfg, tmpl, nil, nil)
	require.NoError(t, err)

	msg := &Message{
		Data: notify.GetTemplateData(testContext(), tmpl, []*alert.Alert{testAlert()}),
	}
	req, err := n.buildRequestForApp(cfg, msg, "")
	require.NoError(t, err)

	// Verify sign = md5(appKey + timestamp + appMasterSecret)
	expectedSign := md5Hex("my_app_key" + req.Timestamp + "my_secret")
	assert.Equal(t, expectedSign, req.Sign)
}

func TestBuildRequest_WithPolicy(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New()
	require.NoError(t, err)
	cfg := newTestUmengConfig("test_key", "test_secret", "android")
	n, err := New(cfg, tmpl, nil, nil)
	require.NoError(t, err)

	msg := &Message{
		Data: notify.GetTemplateData(testContext(), tmpl, []*alert.Alert{testAlert()}),
	}
	req, err := n.buildRequestForApp(cfg, msg, "")
	require.NoError(t, err)

	// Policy is no longer supported in simplified config
	assert.Nil(t, req.Policy)
}

func TestBuildRequest_WithDescription(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New()
	require.NoError(t, err)
	cfg := newTestUmengConfig("test_key", "test_secret", "android")
	n, err := New(cfg, tmpl, nil, nil)
	require.NoError(t, err)

	msg := &Message{
		Data: notify.GetTemplateData(testContext(), tmpl, []*alert.Alert{testAlert()}),
	}
	req, err := n.buildRequestForApp(cfg, msg, "")
	require.NoError(t, err)

	// Description is no longer supported in simplified config
	assert.Empty(t, req.Description)
}

func TestBuildRequest_Alias(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New()
	require.NoError(t, err)
	cfg := newTestUmengConfig("test_key", "test_secret", "android")
	n, err := New(cfg, tmpl, nil, nil)
	require.NoError(t, err)

	// Create alert with user label to trigger alias-based push
	alertWithUser := &alert.Alert{
		Labels:    label.LabelSet{"alertname": "test", "user": "user_123"},
		StartsAt:  time.Now(),
		EndsAt:    time.Now().Add(time.Hour),
		UpdatedAt: time.Now(),
	}
	msg := &Message{
		Data: notify.GetTemplateData(testContext(), tmpl, []*alert.Alert{alertWithUser}),
	}
	req, err := n.buildRequestForApp(cfg, msg, "")
	require.NoError(t, err)

	assert.Equal(t, "user_123", req.Alias)
	assert.Equal(t, "uid", req.AliasType)
}

func TestNotify_Success(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		body, _ := io.ReadAll(r.Body)
		var req Request
		require.NoError(t, json.Unmarshal(body, &req))
		assert.Equal(t, "test_key", req.AppKey)
		assert.NotEmpty(t, req.Timestamp)
		assert.NotEmpty(t, req.Sign)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Ret: "SUCCESS", Data: struct {
			MsgID string `json:"msg_id"`
		}{MsgID: "msg123"}})
	}))
	defer srv.Close()

	cfg := &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test": {
				AppKey:          "test_key",
				AppMasterSecret: "test_secret",
				Platform:        "android",
			},
		},
		APIURL: srv.URL,
	}
	n := newTestNotifier(t, cfg)

	retry, err := n.Notify(testContext(), testAlert())
	require.NoError(t, err)
	assert.False(t, retry)
}

func TestNotify_FailResponse(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Ret:       "FAIL",
			ErrorCode: "1001",
			ErrorMsg:  "Invalid app_key",
		})
	}))
	defer srv.Close()

	cfg := &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test": {
				AppKey:          "bad_key",
				AppMasterSecret: "bad_secret",
				Platform:        "android",
			},
		},
		APIURL: srv.URL,
	}
	n := newTestNotifier(t, cfg)

	retry, err := n.Notify(testContext(), testAlert())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "1001")
	assert.Contains(t, err.Error(), "Invalid app_key")
	assert.False(t, retry)
}

func TestNotify_ServerError(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	cfg := &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test": {
				AppKey:          "test_key",
				AppMasterSecret: "test_secret",
				Platform:        "android",
			},
		},
		APIURL: srv.URL,
	}
	n := newTestNotifier(t, cfg)

	retry, err := n.Notify(testContext(), testAlert())
	assert.Error(t, err)
	assert.True(t, retry)
}

func TestNotify_DefaultAPIURL(t *testing.T) {
	t.Parallel()
	var receivedPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Ret: "SUCCESS"})
	}))
	defer srv.Close()

	cfg := &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test": {
				AppKey:          "test_key",
				AppMasterSecret: "test_secret",
				Platform:        "android",
			},
		},
		APIURL: srv.URL + "/api/send",
	}
	n := newTestNotifier(t, cfg)

	_, err := n.Notify(testContext(), testAlert())
	require.NoError(t, err)
	assert.Equal(t, "/api/send", receivedPath)
}

func TestNotify_MultipleAlerts(t *testing.T) {
	t.Parallel()
	var receivedReq Request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &receivedReq)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Ret: "SUCCESS"})
	}))
	defer srv.Close()

	cfg := &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test": {
				AppKey:          "test_key",
				AppMasterSecret: "test_secret",
				Platform:        "android",
			},
		},
		APIURL: srv.URL,
	}
	n := newTestNotifier(t, cfg)

	alerts := []*alert.Alert{testAlert(), testAlert(), testAlert()}
	_, err := n.Notify(testContext(), alerts...)
	require.NoError(t, err)
}

func TestRedactURLError(t *testing.T) {
	t.Parallel()
	err := redactURLError(io.EOF)
	assert.Equal(t, io.EOF, err)
}

func TestDrain(t *testing.T) {
	t.Parallel()
	body := io.NopCloser(strings.NewReader("response body"))
	resp := &http.Response{Body: body}
	drain(resp)
	b, _ := io.ReadAll(body)
	assert.Empty(t, b)
}

func TestErrDetails(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "", errDetails(nil))
	assert.Equal(t, "error msg", errDetails(strings.NewReader("error msg")))
}

func TestUmengConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     profile.UmengConfig
		wantErr string
	}{
		{
			name:    "missing apps",
			cfg:     profile.UmengConfig{},
			wantErr: "missing apps",
		},
		{
			name: "missing appKey",
			cfg: profile.UmengConfig{
				Apps: map[string]*profile.UmengAppConfig{
					"test": {AppMasterSecret: "secret", Platform: "android"},
				},
			},
			wantErr: "missing appKey",
		},
		{
			name: "missing appMasterSecret",
			cfg: profile.UmengConfig{
				Apps: map[string]*profile.UmengAppConfig{
					"test": {AppKey: "key", Platform: "android"},
				},
			},
			wantErr: "missing appMasterSecret",
		},
		{
			name: "missing platform",
			cfg: profile.UmengConfig{
				Apps: map[string]*profile.UmengAppConfig{
					"test": {AppKey: "key", AppMasterSecret: "secret"},
				},
			},
			wantErr: "missing platform",
		},
		{
			name: "invalid platform",
			cfg: profile.UmengConfig{
				Apps: map[string]*profile.UmengAppConfig{
					"test": {AppKey: "key", AppMasterSecret: "secret", Platform: "invalid"},
				},
			},
			wantErr: "invalid platform",
		},
		{
			name: "valid android",
			cfg: profile.UmengConfig{
				Apps: map[string]*profile.UmengAppConfig{
					"test": {AppKey: "key", AppMasterSecret: "secret", Platform: "android"},
				},
			},
		},
		{
			name: "valid ios",
			cfg: profile.UmengConfig{
				Apps: map[string]*profile.UmengAppConfig{
					"test": {AppKey: "key", AppMasterSecret: "secret", Platform: "ios"},
				},
			},
		},
		{
			name: "valid harmonyos",
			cfg: profile.UmengConfig{
				Apps: map[string]*profile.UmengAppConfig{
					"test": {AppKey: "key", AppMasterSecret: "secret", Platform: "harmonyos"},
				},
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

func TestUmengConfig_Clone(t *testing.T) {
	t.Parallel()
	cfg := &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test": {
				AppKey:          "key",
				AppMasterSecret: "secret",
				Platform:        "android",
			},
		},
	}
	clone := cfg.Clone()
	assert.Equal(t, cfg.Apps["test"].AppKey, clone.Apps["test"].AppKey)
	assert.Equal(t, cfg.Apps["test"].AppMasterSecret, clone.Apps["test"].AppMasterSecret)
	assert.Equal(t, cfg.Apps["test"].Platform, clone.Apps["test"].Platform)

	clone.Apps["test"].AppKey = "modified"
	assert.NotEqual(t, cfg.Apps["test"].AppKey, clone.Apps["test"].AppKey)
}

func boolPtr(b bool) *bool {
	return &b
}

func TestExtractUserIDs(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New()
	require.NoError(t, err)

	tests := []struct {
		name   string
		alerts []*alert.Alert
		want   []string
	}{
		{
			name:   "no alerts",
			alerts: nil,
			want:   nil,
		},
		{
			name: "single user",
			alerts: []*alert.Alert{
				{Labels: label.LabelSet{"user": "100"}},
			},
			want: []string{"100"},
		},
		{
			name: "multiple users in one alert",
			alerts: []*alert.Alert{
				{Labels: label.LabelSet{"user": "100,200"}},
			},
			want: []string{"100", "200"},
		},
		{
			name: "users across multiple alerts deduped",
			alerts: []*alert.Alert{
				{Labels: label.LabelSet{"user": "100"}},
				{Labels: label.LabelSet{"user": "100,200"}},
				{Labels: label.LabelSet{"user": "300"}},
			},
			want: []string{"100", "200", "300"},
		},
		{
			name: "no user label",
			alerts: []*alert.Alert{
				{Labels: label.LabelSet{"alertname": "test"}},
			},
			want: nil,
		},
		{
			name: "empty user label",
			alerts: []*alert.Alert{
				{Labels: label.LabelSet{"user": ""}},
			},
			want: nil,
		},
		{
			name: "whitespace trimmed",
			alerts: []*alert.Alert{
				{Labels: label.LabelSet{"user": " 100 , 200 "}},
			},
			want: []string{"100", "200"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			data := notify.GetTemplateData(testContext(), tmpl, tt.alerts)
			msg := &Message{Data: data}
			got := extractUserIDs(msg)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBuildRequest_PushTypeFromUserIDs(t *testing.T) {
	t.Parallel()
	tmpl, err := template.New()
	require.NoError(t, err)

	t.Run("single user becomes customizedcast", func(t *testing.T) {
		t.Parallel()
		cfg := newTestUmengConfig("key", "secret", "android")
		n, err := New(cfg, tmpl, nil, nil)
		require.NoError(t, err)

		alerts := []*alert.Alert{{Labels: label.LabelSet{"user": "100"}}}
		data := notify.GetTemplateData(testContext(), tmpl, alerts)
		msg := &Message{Data: data}

		req, err := n.buildRequestForApp(cfg, msg, "")
		require.NoError(t, err)
		assert.Equal(t, "customizedcast", req.Type)
		assert.Equal(t, "100", req.Alias)
		assert.Equal(t, "uid", req.AliasType)
	})

	t.Run("multiple users become customizedcast", func(t *testing.T) {
		t.Parallel()
		cfg := newTestUmengConfig("key", "secret", "android")
		n, err := New(cfg, tmpl, nil, nil)
		require.NoError(t, err)

		alerts := []*alert.Alert{{Labels: label.LabelSet{"user": "100,200,300"}}}
		data := notify.GetTemplateData(testContext(), tmpl, alerts)
		msg := &Message{Data: data}

		req, err := n.buildRequestForApp(cfg, msg, "")
		require.NoError(t, err)
		assert.Equal(t, "customizedcast", req.Type)
		assert.Equal(t, "100,200,300", req.Alias)
		assert.Equal(t, "uid", req.AliasType)
	})

	t.Run("no users stays broadcast", func(t *testing.T) {
		t.Parallel()
		cfg := newTestUmengConfig("key", "secret", "android")
		n, err := New(cfg, tmpl, nil, nil)
		require.NoError(t, err)

		alerts := []*alert.Alert{{Labels: label.LabelSet{"alertname": "test"}}}
		data := notify.GetTemplateData(testContext(), tmpl, alerts)
		msg := &Message{Data: data}

		req, err := n.buildRequestForApp(cfg, msg, "")
		require.NoError(t, err)
		assert.Equal(t, TypeBroadcast, req.Type)
		assert.Empty(t, req.Alias)
	})

	t.Run("custom alias type", func(t *testing.T) {
		t.Parallel()
		cfg := newTestUmengConfig("key", "secret", "android")
		n, err := New(cfg, tmpl, nil, nil)
		require.NoError(t, err)

		alerts := []*alert.Alert{{Labels: label.LabelSet{"user": "100"}}}
		data := notify.GetTemplateData(testContext(), tmpl, alerts)
		msg := &Message{Data: data}

		req, err := n.buildRequestForApp(cfg, msg, "")
		require.NoError(t, err)
		assert.Equal(t, "customizedcast", req.Type)
		assert.Equal(t, "100", req.Alias)
		// AliasType is now hardcoded to "uid" in simplified config
		assert.Equal(t, "uid", req.AliasType)
	})

	t.Run("users across alerts deduped", func(t *testing.T) {
		t.Parallel()
		cfg := newTestUmengConfig("key", "secret", "android")
		n, err := New(cfg, tmpl, nil, nil)
		require.NoError(t, err)

		alerts := []*alert.Alert{
			{Labels: label.LabelSet{"user": "100"}},
			{Labels: label.LabelSet{"user": "100,200"}},
		}
		data := notify.GetTemplateData(testContext(), tmpl, alerts)
		msg := &Message{Data: data}

		req, err := n.buildRequestForApp(cfg, msg, "")
		require.NoError(t, err)
		assert.Equal(t, "customizedcast", req.Type)
		assert.Equal(t, "100,200", req.Alias)
	})
}

// loadTestEnv reads .env.local file for integration testing.
// Returns nil if the file does not exist.
func loadTestEnv() map[string]string {
	envFile := filepath.Join("testdata", ".env.local")
	f, err := os.Open(envFile)
	if err != nil {
		return nil
	}
	defer f.Close()

	env := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			if val != "" {
				env[key] = val
			}
		}
	}
	if len(env) == 0 {
		return nil
	}
	return env
}

// expandEnvVars replaces ${VAR_NAME} patterns in a string with environment variable values.
func expandEnvVars(s string, env map[string]string) string {
	result := s
	for key, val := range env {
		result = strings.ReplaceAll(result, "${"+key+"}", val)
	}
	return result
}

// TestIntegration_RealAPI_Android sends a real push to Umeng API for Android.
// Run with: go test -v -run TestIntegration_RealAPI_Android ./notify/umeng/...
// Requires: notify/umeng/testdata/.env.local with valid Android credentials.
func TestIntegration_RealAPI_Android(t *testing.T) {
	env := loadTestEnv()
	if env == nil {
		t.Skip("testdata/.env.local not found or empty, skipping integration test")
	}

	appKey := env["UMENG_ANDROID_APP_KEY"]
	appSecret := env["UMENG_ANDROID_APP_MASTER_SECRET"]
	if appKey == "" || appSecret == "" {
		t.Skip("UMENG_ANDROID_APP_KEY or UMENG_ANDROID_APP_MASTER_SECRET not set, skipping")
	}

	cfg := &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test-android": {
				AppKey:          appKey,
				AppMasterSecret: appSecret,
				Platform:        "android",
				AppSet:          env["UMENG_ANDROID_APP_SET"],
				AliasType:       env["UMENG_ANDROID_ALIAS_TYPE"],
				AfterOpen:       env["UMENG_ANDROID_AFTER_OPEN"],
				Activity:        env["UMENG_ANDROID_ACTIVITY"],
			},
		},
		ProductionMode: boolPtr(false),
	}

	if cfg.Apps["test-android"].AliasType == "" {
		cfg.Apps["test-android"].AliasType = "uid"
	}

	n := newTestNotifier(t, cfg)

	ctx := testContext()
	alert := &alert.Alert{
		Labels: label.LabelSet{
			"alertname": "IntegrationTestAndroid",
			"severity":  "info",
			"user":      "test-user-123", // 测试用户ID
		},
		StartsAt:  time.Now(),
		EndsAt:    time.Now().Add(time.Hour),
		UpdatedAt: time.Now(),
		Annotations: label.LabelSet{
			"summary": "msgcenter Android integration test",
		},
	}

	retry, err := n.Notify(ctx, alert)
	require.NoError(t, err, "Umeng API call should succeed")
	assert.False(t, retry, "should not retry on success")
	t.Logf("Android push sent successfully, retry=%v", retry)
}

// TestIntegration_RealAPI_IOS sends a real push to Umeng API for iOS.
// Run with: go test -v -run TestIntegration_RealAPI_IOS ./notify/umeng/...
// Requires: notify/umeng/testdata/.env.local with valid iOS credentials.
func TestIntegration_RealAPI_IOS(t *testing.T) {
	env := loadTestEnv()
	if env == nil {
		t.Skip("testdata/.env.local not found or empty, skipping integration test")
	}

	appKey := env["UMENG_IOS_APP_KEY"]
	appSecret := env["UMENG_IOS_APP_MASTER_SECRET"]
	if appKey == "" || appSecret == "" {
		t.Skip("UMENG_IOS_APP_KEY or UMENG_IOS_APP_MASTER_SECRET not set, skipping")
	}

	cfg := &profile.UmengConfig{
		Apps: map[string]*profile.UmengAppConfig{
			"test-ios": {
				AppKey:          appKey,
				AppMasterSecret: appSecret,
				Platform:        "ios",
				AppSet:          env["UMENG_IOS_APP_SET"],
				AliasType:       env["UMENG_IOS_ALIAS_TYPE"],
			},
		},
		ProductionMode: boolPtr(false),
	}

	if cfg.Apps["test-ios"].AliasType == "" {
		cfg.Apps["test-ios"].AliasType = "uid"
	}

	n := newTestNotifier(t, cfg)

	ctx := testContext()
	alert := &alert.Alert{
		Labels: label.LabelSet{
			"alertname": "IntegrationTestIOS",
			"severity":  "info",
			"user":      "test-user-456", // 测试用户ID
		},
		StartsAt:  time.Now(),
		EndsAt:    time.Now().Add(time.Hour),
		UpdatedAt: time.Now(),
		Annotations: label.LabelSet{
			"summary": "msgcenter iOS integration test",
		},
	}

	retry, err := n.Notify(ctx, alert)
	require.NoError(t, err, "Umeng API call should succeed")
	assert.False(t, retry, "should not retry on success")
	t.Logf("iOS push sent successfully, retry=%v", retry)
}
