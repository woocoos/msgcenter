package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/msgcenter/api/oas"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/metrics"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/provider/mem"
	"github.com/woocoos/msgcenter/service"
	"github.com/woocoos/msgcenter/test"
	"github.com/woocoos/msgcenter/test/maildev"
	"math/rand"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

func open(ctx context.Context) *ent.Client {
	client, err := ent.Open("sqlite3", "file:msgcenter?mode=memory&cache=shared&_fk=1", ent.Debug())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

// ServiceSuite is the service test suite
type serviceSuite struct {
	suite.Suite
	Service      *Service
	AlertManager *service.AlertManager
	server       *web.Server
	db           *ent.Client
	shutdowns    []func()
	maildev      maildev.MailDev
}

// TestServiceSuite runs the service test suite
func TestServiceSuite(t *testing.T) {
	suite.Run(t, &serviceSuite{
		maildev: maildev.MailDev{
			URL: &url.URL{
				Host:   "localhost:8025",
				Scheme: "http",
			},
		},
	})
}

func (s *serviceSuite) initData(ctx context.Context) {
	s.db.MsgType.Create().SetName("alert").SetID(1).SetStatus(typex.SimpleStatusActive).SetCreatedBy(1).
		SetAppID(1).SetCategory("账户安全").SetCanSubs(true).SetCanCustom(true).SaveX(ctx)
	s.db.MsgEvent.Create().SetID(1).SetMsgTypeID(1).SetName("AlterPassword").SetStatus(typex.SimpleStatusActive).
		SetCreatedBy(1).SetModes("email,internal").
		SetRoute(&profile.Route{
			Name:     "AlterPassword",
			Receiver: "email",
			Matchers: label.Matchers{
				{
					Name:  "app",
					Value: "1",
					Type:  label.MatchEqual,
				},
				{
					Name:  "alertname",
					Value: "AlterPassword",
					Type:  label.MatchEqual,
				},
			},
			Routes: []*profile.Route{
				{
					Matchers: label.Matchers{
						{
							Name:  "receiver",
							Value: "email",
							Type:  label.MatchRegexp,
						},
					},
					Receiver: "email",
					Continue: true,
				},
				{
					Matchers: label.Matchers{
						{
							Name:  "receiver",
							Value: "internal",
							Type:  label.MatchRegexp,
						},
					},
					Receiver: "internal",
				},
			},
		}).SaveX(ctx)
	s.db.MsgTemplate.Create().SetMsgTypeID(1).SetEventID(1).SetTenantID(1000).SetName("AlterPassword").SetCreatedBy(1).
		SetStatus(typex.SimpleStatusActive).SetFormat(msgtemplate.FormatTxt).SetReceiverType(profile.ReceiverEmail).SetTo(`{{ template "email.to" . }}`).
		SetSubject(`{{ with .CommonAnnotations }}{{.uid}}{{end}}密码到期提醒`).SetCc(`{{ template "email.cc" . }}`).
		SetBcc(`{{ template "email.bcc" . }}`).SetFrom(`custom <test@localhost>`).
		SetBody(`{{ template "1000.alterpwd.txt" . }}`).SetAttachments("1000/alterpwd.tmpl").
		SetTpl("1000/alterpwd.tmpl").SaveX(ctx)

	s.db.MsgChannel.Create().SetName("email").SetStatus(typex.SimpleStatusActive).SetCreatedBy(1).
		SetTenantID(1000).SetReceiverType(profile.ReceiverEmail).
		SetReceiver(&profile.Receiver{
			Name: "email",
			EmailConfigs: []*profile.EmailConfig{
				{
					SmartHost: profile.HostPort{Host: "localhost", Port: "1025"},
					To:        `{{ template "email.to" . }}`,
					From:      "serviceSuite@localhost",
				},
			},
		}).SaveX(ctx)
}

// SetupSuite sets up the test suite
func (s *serviceSuite) SetupSuite() {
	file := filepath.Join(test.BaseDir(), "testdata", "etc", "app.yaml")
	bs, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	app := woocoo.New(woocoo.WithAppConfiguration(
		conf.NewFromBytes(bs, conf.WithBaseDir(test.BaseDir())).AsGlobal()),
	)
	cnf := app.AppConfiguration()
	cnf.Parser().Set("alertManager.storage.path", filepath.Join(test.BaseDir(), "testdata", "tmp"))
	cnf.Parser().Set("alertManager.config.file", filepath.Join(test.BaseDir(), "testdata", "etc", "alertmanager.yaml"))

	s.db = open(context.Background())
	s.initData(context.Background())

	metrics.BuildGlobal(prometheus.DefaultRegisterer)

	alertManagerCnf := cnf.Sub("alertManager")
	s.AlertManager, err = service.DefaultAlertManager(alertManagerCnf)
	s.Require().NoError(err)

	configCoordinator := service.NewCoordinator(alertManagerCnf)

	s.server = web.New(web.WithConfiguration(cnf.Sub("web")))
	s.Service, err = NewService(
		&Options{
			Coordinator: configCoordinator,
			Alerts:      s.AlertManager.Alerts,
			Silences:    s.AlertManager.Silences,
			StatusFunc:  s.AlertManager.Marker.Status,
			Registry:    prometheus.DefaultRegisterer,
			GroupFunc:   s.AlertManager.Dispatcher.Groups,
		})
	s.Require().NoError(err)

	configCoordinator.SetDBClient(s.db)
	configCoordinator.Subscribe(func(c *profile.Config) error {
		configCoordinator.Template.ExternalURL, err = url.Parse("http://localhost:9093")
		s.Require().NoError(err)
		s.Require().NoError(s.AlertManager.Start(configCoordinator, c))

		s.Service.Update(c, func(labels label.LabelSet) {
			s.AlertManager.Inhibitor.Mutes(labels)
			s.AlertManager.Silencer.Mutes(labels)
		})

		return nil
	})

	err = configCoordinator.Reload()
	s.Require().NoError(err)
	alerts := s.AlertManager.Alerts.(*mem.Alerts)
	go alerts.Start(nil)
	s.shutdowns = append(s.shutdowns, func() {
		s.AlertManager.Stop()
		alerts.Stop(context.Background())
	})
}

// TearDownSuite tears down the test suite
func (s *serviceSuite) TearDownSuite() {
	for _, shutdown := range s.shutdowns {
		shutdown()
	}
}

// test postalerts
func (s *serviceSuite) TestPostAlerts() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := oas.PostableAlerts{
		{
			Alert: &oas.Alert{
				Labels: map[string]string{
					"alertname": "test",
				},
			},
			Annotations: map[string]string{
				"summary": "test",
			},
			EndsAt:   time.Now().Add(time.Hour),
			StartsAt: time.Now(),
		},
	}
	s.Require().NoError(s.Service.PostAlerts(ctx, &oas.PostAlertsRequest{Body: req}))
	time.Sleep(time.Second * 2)
	mail, err := s.maildev.GetLastEmail()
	s.Require().NoError(err)
	s.Require().NotNil(mail)
	s.Require().Equal("alerts@example.com", mail.To[0]["Address"])
}

// TestPostAlertsWithTenant tenant with custom template and attachment
func (s *serviceSuite) TestPostAlertsWithTenant() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	// rand a string
	to := fmt.Sprintf("%d@localhost", rand.Intn(10000000))
	req := oas.PostableAlerts{
		{
			Alert: &oas.Alert{
				Labels: map[string]string{
					"receiver":  "email|webhook",
					"alertname": "AlterPassword",
					"tenant":    "1000",
				},
			},
			Annotations: map[string]string{
				"to":       to,
				"summary":  "test",
				"nickname": "woocoos",
			},
			EndsAt:   time.Now().Add(time.Second * 2),
			StartsAt: time.Now(),
		},
	}
	s.Require().NoError(s.Service.PostAlerts(ctx, &oas.PostAlertsRequest{Body: req}))
	time.Sleep(time.Second * 3)
	mail, err := s.maildev.GetLastEmail()
	s.Require().NoError(err)
	s.Require().NotNil(mail)
	s.Require().Equal(to, mail.To[0]["Address"])
}
