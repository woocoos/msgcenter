package profile

import (
	"encoding/json"
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
				return assert.ErrorContains(t, err, "url must be configured")
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
    cache:
      type: redis
      addrs:
        - 127.0.0.1:6379
      db: 1
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
			err := cfg.Unmarshal(&config)
			require.NoError(t, err)
			err = config.Validate()
			tt.wantErr(t, err, &config)
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
