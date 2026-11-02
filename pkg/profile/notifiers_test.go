package profile

import (
	"encoding/json"
	jsonparse "github.com/knadh/koanf/parsers/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsingsun/woocoo/pkg/conf"
	"testing"
)

func TestEmail(t *testing.T) {
	tests := []struct {
		name    string
		cfgStr  string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Email is required",
			cfgStr: `
to: ''
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "missing to address in email config")
			},
		},
		{
			name: "Email headers duplicate",
			cfgStr: `
to: 'to@email.com'
headers:
  Subject: 'Alert'
  subject: 'New Alert'
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "duplicate header \"Subject\" in email config")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var config = DefaultEmailConfig
			cfg := conf.NewFromBytes([]byte(tt.cfgStr))
			err := cfg.Unmarshal(&config)
			require.NoError(t, err)
			err = config.Validate()
			tt.wantErr(t, err)
		})
	}
}

func TestWebHook(t *testing.T) {
	tests := []struct {
		name    string
		cfgStr  string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Webhook URL",
			cfgStr: `
url: 'http://example.com'
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				wc := i[0].(*WebhookConfig)
				return assert.Equal(t, wc.URL.String(), "http://example.com")
			},
		},
		{
			name:   "Webhook URL is required",
			cfgStr: ``,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "one of url or urlFile must be configured")
			},
		},
		{
			name: "oauth2",
			cfgStr: `
url: "http://127.0.0.1:5001/"
httpConfig:
  timeout: 1s
  oauth2:
    clientID: 206734260394752
    clientSecret: T2UlqISVFq4DR9InXamj3l74iWdu3Tyr
    endpoint:
      tokenURL: http://127.0.0.1:5001/token
    scopes:
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				wc := i[0].(*WebhookConfig)
				var cc = DefaultWebhookConfig
				js, _ := json.Marshal(wc)
				if assert.NoError(t, json.Unmarshal(js, &cc)) {
					assert.Equal(t, cc.URL.String(), "http://127.0.0.1:5001/")
					assert.Equal(t, cc.HTTPConfig.OAuth2.ClientID, "206734260394752")
					assert.Equal(t, cc.HTTPConfig.OAuth2.ClientSecret, "T2UlqISVFq4DR9InXamj3l74iWdu3Tyr")
					assert.Equal(t, cc.HTTPConfig.OAuth2.Endpoint.TokenURL, "http://127.0.0.1:5001/token")
					assert.Equal(t, cc.HTTPConfig.OAuth2.Scopes, []string(nil))
					assert.NotNil(t, cc.HttpConfigOri)
				}
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var config = DefaultWebhookConfig
			cfg := conf.NewFromBytes([]byte(tt.cfgStr))
			bs, err := cfg.Parser().ToBytes(jsonparse.Parser())
			require.NoError(t, err)
			if err = json.Unmarshal(bs, &config); err != nil {
				tt.wantErr(t, err)
			} else {
				err = config.Validate()
				tt.wantErr(t, err, &config)
			}
		})
	}
}

func TestWebhookConfig_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "oauth2",
			args: args{
				data: []byte(`
{
  "sendResolved": true,
  "httpConfig": {
    "timeout": 1000000000,
    "oauth2": {
      "clientID": "206734260394752",
      "clientSecret": "T2UlqISVFq4DR9InXamj3l74iWdu3Tyr",
      "endpoint": {
        "tokenURL": "http://127.0.0.1:5001/token"
      }
    }
  },
  "url": "http://127.0.0.1:5001/",
  "maxAlerts": 0
}
`)},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if !assert.NoError(t, err) {
					return false
				}
				got := i[0].(*WebhookConfig)
				assert.Equal(t, got.URL.String(), "http://127.0.0.1:5001/")
				assert.Equal(t, got.HTTPConfig.OAuth2.ClientID, "206734260394752")
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &WebhookConfig{}
			err := got.UnmarshalJSON(tt.args.data)
			tt.wantErr(t, err, got)
		})
	}
}

// TestMessageConfigClone_DeepCopyExtras 测试 Clone 深拷贝 Extras。
// 浅拷贝会与原配置共享底层 map——并发 Notify 时 overrideMessageConfig 对 Extras 的
// 写入相互竞争（fatal: concurrent map writes），且基础配置被写入污染、并发串数据。
func TestMessageConfigClone_DeepCopyExtras(t *testing.T) {
	cfg := &MessageConfig{Extras: map[string]string{"category": "cat"}}

	clone := cfg.Clone()

	// 对 clone 的 Extras 写入不得影响原配置（map 必须独立）
	clone.Extras["msgTypeName"] = "type1"
	assert.Empty(t, cfg.Extras["msgTypeName"], "clone must not share the Extras map with the original")
	assert.Equal(t, "cat", clone.Extras["category"], "existing Extras entries must be copied to the clone")

	// 对原配置的 Extras 写入也不得影响 clone（双向隔离）
	cfg.Extras["category"] = "cat2"
	assert.Equal(t, "cat", clone.Extras["category"], "clone Extras must be independent of the original")
}
