package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/msgcenter/api/oas"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/provider/mem"
	"github.com/woocoos/msgcenter/test/maildev"
	"github.com/woocoos/msgcenter/test/testsuite"
	"math/rand"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

var (
	subEventName = "SubEvent"
)

func open(ctx context.Context) *ent.Client {
	client, err := ent.Open("sqlite3", "file:msgcenter?mode=memory&cache=shared&_fk=1",
		ent.Debug(), ent.AlternateSchema(ent.SchemaConfig{
			User:        "portal",
			OrgRoleUser: "portal",
		}),
	)
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
	testsuite.BaseSuite
	Service   *Service
	server    *web.Server
	shutdowns []func()
	maildev   maildev.MailDev
}

// TestServiceSuite runs the service test suite
func TestServiceSuite(t *testing.T) {
	s := &serviceSuite{
		maildev: maildev.MailDev{
			URL: &url.URL{
				Host:   "localhost:8025",
				Scheme: "http",
			},
		},
	}
	s.DSN = "file:msgcenter?mode=memory&cache=shared&_fk=1"
	s.DriverName = "sqlite3"
	suite.Run(t, s)
}

// SetupSuite sets up the test suite
func (s *serviceSuite) SetupSuite() {
	err := s.BaseSuite.Setup()
	s.Require().NoError(err)

	s.server = web.New(web.WithConfiguration(s.Cnf.Sub("web")))
	s.Service, err = NewService(
		&Options{
			Coordinator: s.ConfigCoordinator,
			Alerts:      s.AlertManager.Alerts,
			Silences:    s.AlertManager.Silences,
			StatusFunc:  s.AlertManager.Marker.Status,
			Registry:    prometheus.DefaultRegisterer,
			GroupFunc:   s.AlertManager.Dispatcher.Groups,
		})
	s.Require().NoError(err)

	s.ConfigCoordinator.ReloadHooks(func(c *profile.Config) error {
		s.ConfigCoordinator.Template.ExternalURL, err = url.Parse("http://localhost:9093")
		s.Require().NoError(err)
		s.Require().NoError(s.AlertManager.Start(s.ConfigCoordinator, c))

		s.Service.Update(c, func(labels label.LabelSet) {
			s.AlertManager.Inhibitor.Mutes(labels)
			s.AlertManager.Silencer.Mutes(labels)
		})

		return nil
	})

	err = s.ConfigCoordinator.Reload()
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
	s.Require().NoError(s.Service.PostAlerts(ctx, &oas.PostAlertsRequest{PostableAlerts: req}))
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
					"alertname": "noSubscribe",
					"tenant":    "1",
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
	s.Require().NoError(s.Service.PostAlerts(ctx, &oas.PostAlertsRequest{PostableAlerts: req}))
	time.Sleep(time.Second * 3)
	mail, err := s.maildev.GetLastEmail()
	s.Require().NoError(err)
	s.Require().NotNil(mail)
	s.Require().Equal(to, mail.To[0]["Address"])
}

// TestPostAlertsWithTenant tenant with custom template and attachment
func (s *serviceSuite) TestUserSubscribe() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	req := oas.PostableAlerts{
		{
			Alert: &oas.Alert{
				Labels: map[string]string{
					"receiver":  "email|webhook",
					"alertname": subEventName,
					"tenant":    "1",
				},
			},
			Annotations: map[string]string{
				"summary":  "test",
				"nickname": "woocoos",
			},
			EndsAt:   time.Now().Add(time.Second * 5),
			StartsAt: time.Now(),
		},
	}
	s.Require().NoError(s.Service.PostAlerts(ctx, &oas.PostAlertsRequest{PostableAlerts: req}))
	time.Sleep(time.Second * 3)
	mail, err := s.maildev.GetLastEmail()
	s.Require().NoError(err)
	s.Require().NotNil(mail)
	// user 1 or user2,but notify not keep order
	s.Require().Equal("订阅事件提醒", mail.Subject)
}

func (s *serviceSuite) TestPostSilence() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	req := oas.PostableSilence{
		Silence: &oas.Silence{
			Comment:   "test",
			CreatedBy: 1,
			StartsAt:  time.Now(),
			EndsAt:    time.Now().Add(time.Hour),
			Matchers: []*oas.Matcher{
				{
					Name:  "alertname",
					Value: "test",
				},
			},
		},
	}
	res, err := s.Service.PostSilences(ctx, &oas.PostSilencesRequest{PostableSilence: req})
	s.Require().NoError(err)
	s.NotZero(res.SilenceID)
}
