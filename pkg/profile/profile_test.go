package profile

import (
	stdjson "encoding/json"
	"github.com/knadh/koanf/parsers/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tsingsun/woocoo/pkg/conf"
	"testing"
)

func TestReceiver(t *testing.T) {
	tests := []struct {
		name    string
		cfgStr  string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Default receiver exists",
			cfgStr: `
route:
  groupWait: 30s
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrRootMissReceiver)
			},
		},
		{
			name: "Receiver name is unique",
			cfgStr: `
route:
  receiver: team-X
receivers:
- name: team-X
- name: team-X
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "notification config name \"team-X\" is not unique")
			},
		},
		{
			name: "Receiver exists",
			cfgStr: `
route:
  receiver: team-X
receivers:
- name: team-Y
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "undefined receiver \"team-X\" used in route")
			},
		},
		{
			name: "Receiver exists for deep sub route",
			cfgStr: `
route:
  receiver: team-X
  routes:
  - match:
      foo: bar
    routes:
    - match:
      foo: bar
      receiver: nonexistent
receivers:
- name: team-X
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "undefined receiver \"nonexistent\" used in route")
			},
		},
		{
			name: "Receiver has name",
			cfgStr: `
route:
  receiver: ""
receivers:
- name: ""
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrRootMissReceiver)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Load([]byte(tt.cfgStr))
			tt.wantErr(t, err)
		})
	}
}

func TestMuteAndActiveTimeExists(t *testing.T) {
	tests := []struct {
		name    string
		cfgStr  string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "MuteTime exists",
			cfgStr: `
route:
    receiver: team-Y
    routes:
    -  match:
        severity: critical
       muteTimeIntervals:
       - businessHours

receivers:
- name: 'team-Y'
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "undefined time interval \"businessHours\" used in route")
			},
		},
		{
			name: "ActiveTime exists",
			cfgStr: `
route:
    receiver: team-Y
    routes:
    -  match:
        severity: critical
       ActiveTimeIntervals:
       - businessHours

receivers:
- name: 'team-Y'
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "undefined time interval \"businessHours\" used in route")
			},
		},
		{
			name: "Time interval has name",
			cfgStr: `
timeIntervals:
- name: 
  timeIntervals:
  - times: '09:00~17:00'

receivers:
- name: 'team-X-mails'

route:
  receiver: 'team-X-mails'
  routes:
  -  match:
      severity: critical
     mute_time_intervals:
     - business_hours
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrNeedTimeIntervalName)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var config Config
			cfg := conf.NewFromBytes([]byte(tt.cfgStr))
			bs, err := cfg.ParserOperator().Marshal(json.Parser())
			require.NoError(t, err)
			tt.wantErr(t, stdjson.Unmarshal(bs, &config))
		})
	}
}

func TestGroupBy(t *testing.T) {
	tests := []struct {
		name    string
		cfgStr  string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GroupBy has no duplicated labels",
			cfgStr: `
route:
  receiver: 'team-X'
  groupBy: ['alertname', 'cluster', 'service', 'cluster']

receivers:
- name: 'team-X'
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "duplicated label \"cluster\" in group_by")
			},
		},
		{
			name: "Wildcard group by others",
			cfgStr: `
route:
  groupBy: ['alertname', 'cluster', '...']
  receiver: team-X
receivers:
- name: 'team-X'
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "cannot have wildcard group_by (`...`) and other other labels at the same time")
			},
		},
		{
			name: "GroupBy invalid label",
			cfgStr: `
route:
  groupBy: ['-invalid-']
  receiver: team-X
receivers:
- name: 'team-X'
`,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "invalid label name \"-invalid-\" in group_by list")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				config Config
			)
			cfg := conf.NewFromBytes([]byte(tt.cfgStr))
			bs, err := cfg.ParserOperator().Marshal(json.Parser())
			require.NoError(t, err)
			tt.wantErr(t, stdjson.Unmarshal(bs, &config))
		})
	}
}
