package profile

import (
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
			name: "Webhook URL is required",
			cfgStr: `
url: 'http://example.com'
urlFile: 'http://example.com'
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "at most one of url & url_file must be configured")
			},
		},
		{
			name: "Webhook http config is valid",
			cfgStr: `
url: 'http://example.com'
httpConfig:
  bearer_token: foo
  bearer_token_file: /tmp/bar
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "at most one of bearer_token & bearer_token_file must be configured")

			},
		},
		{
			name: "Webhook http config optional",
			cfgStr: `
url: 'http://example.com'
`,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var config = DefaultWebhookConfig
			cfg := conf.NewFromBytes([]byte(tt.cfgStr))
			err := cfg.Unmarshal(&config)
			require.NoError(t, err)
			err = config.Validate()
			tt.wantErr(t, err)
		})
	}
}
