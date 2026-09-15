package graphql

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo/pkg/gds"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msginternalto"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service"
	"github.com/woocoos/msgcenter/service/provider/mem"
	"github.com/woocoos/msgcenter/test/testsuite"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

type resolverSuite struct {
	testsuite.BaseSuite
	resolver *Resolver
	mr       *mutationResolver
	qr       *queryResolver
	webhook  *httptest.Server
	//server    *web.Server
	shutdowns []func()
}

func TestRolverSuite(t *testing.T) {
	s := &resolverSuite{
		BaseSuite: testsuite.BaseSuite{
			DSN:        "file:msgcenter?mode=memory&cache=shared&_fk=1",
			DriverName: "sqlite3",
		},
	}
	s.webhook = httptest.NewUnstartedServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			if r.URL.Path == "/token" {
				d, _ := json.Marshal(map[string]string{
					"access_token": "90d64460d14870c08c81352a05dedd3465940a7c",
					"expires_in":   "7200",
					"scope":        "user",
					"token_type":   "bearer",
				})
				w.Write(d)
				return
			} else if r.URL.Path == "/graphql/query" {
				w.Write([]byte(`{"data":{}}`))
				return
			} else if r.URL.Path == "/org/domain" {
				d, _ := json.Marshal(map[string]any{
					"id":        1,
					"name":      "test",
					"parent_id": 0,
				})
				w.Write(d)
				return
			}
			w.WriteHeader(http.StatusOK)
		}))
	var err error
	s.webhook.Listener, err = net.Listen("tcp", "127.0.0.1:5001")
	require.NoError(t, err)
	s.webhook.Start()
	defer s.webhook.Close()
	suite.Run(t, s)
}

// SetupSuite sets up the test suite
func (s *resolverSuite) SetupSuite() {
	err := s.BaseSuite.Setup()
	s.Require().NoError(err)
	s.AlertManager, err = service.NewAlertManager(s.App, service.WithClient(s.Client))
	s.Require().NoError(err)
	//s.server = web.New(web.WithConfiguration(s.Cnf.Sub("web")))
	s.resolver = &Resolver{
		coordinator: s.AlertManager.Coordinator,
		client:      s.Client,
		Silences:    s.AlertManager.Silences,
	}
	s.mr = &mutationResolver{
		Resolver: s.resolver,
	}
	s.qr = &queryResolver{
		Resolver: s.resolver,
	}

	s.AlertManager.Coordinator.ReloadHooks(func(c *profile.Config) error {
		s.AlertManager.Coordinator.Template.ExternalURL, err = url.Parse("http://localhost:9093")
		s.Require().NoError(err)
		s.Require().NoError(s.AlertManager.Start(s.AlertManager.Coordinator, c))
		return nil
	})

	err = s.AlertManager.Coordinator.Reload()
	s.Require().NoError(err)
	alerts := s.AlertManager.Alerts.(*mem.Alerts)
	go alerts.Start(context.Background())
	s.shutdowns = append(s.shutdowns, func() {
		s.AlertManager.Stop()
		alerts.Stop(context.Background())
	})
}

// TearDownSuite tears down the test suite
func (s *resolverSuite) TearDownSuite() {
	for _, shutdown := range s.shutdowns {
		shutdown()
	}
}

func (s *resolverSuite) TestCreateMsgSilence() {
	ctx := s.NewTestCtx()
	silence, err := s.mr.CreateMsgSilence(ctx, ent.CreateMsgSilenceInput{
		Comments: gds.Ptr("test"),
		EndsAt:   time.Now().Add(time.Second * 10),
		StartsAt: time.Now().Add(time.Second * -5),
		Matchers: []*label.Matcher{
			{
				Name:  "alertname",
				Value: "test",
			},
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(silence)
}

func (s *resolverSuite) TestUserSubMsgCategory() {
	ctx := s.NewTestCtx()
	category, err := s.qr.UserSubMsgCategory(ctx)
	s.Require().NoError(err)
	s.Require().NotEmpty(category)
	s.Require().Equal(category[0], "订阅类型")
}

func (s *resolverSuite) TestMessageHandler() {
	s.Run("UserUnreadMessages", func() {
		ctx := s.NewTestCtx()
		num, err := s.qr.UserUnreadMsgInternals(ctx)
		s.Require().NoError(err)
		s.Require().Equal(2, num)
	})
	s.Run("UserUnreadMessagesFromMsgCategory", func() {
		ctx := s.NewTestCtx()
		nums, err := s.qr.UserUnreadMsgInternalsFromMsgCategory(ctx, []string{"订阅类型"})
		s.Require().NoError(err)
		s.Require().NotEmpty(nums)
		s.Require().Equal(2, nums[0])
	})
	s.Run("MarkMessageReaOrUnRead", func() {
		ctx := s.NewTestCtx()
		suc, err := s.mr.MarkMsgInternalToReadOrUnRead(ctx, []int{1}, true)
		s.Require().NoError(err)
		s.Require().True(suc)
		has, err := s.Client.MsgInternalTo.Query().Where(msginternalto.IDIn(1), msginternalto.ReadAtNotNil()).Exist(ctx)
		s.Require().NoError(err)
		s.Require().True(has)
	})
	s.Run("MarkMessageDeleted", func() {
		ctx := s.NewTestCtx()
		suc, err := s.mr.MarkMsgInternalToDeleted(ctx, []int{2})
		s.Require().NoError(err)
		s.Require().True(suc)
		has, err := s.Client.MsgInternalTo.Query().Where(msginternalto.IDIn(2), msginternalto.DeleteAtNotNil()).Exist(ctx)
		s.Require().NoError(err)
		s.Require().True(has)
	})
}

func (s *resolverSuite) TestMsgTemplateDefineByName() {
	ctx := s.NewTestCtx()
	txt, err := s.qr.MsgTemplateDefineByName(ctx, "html", `{{ template "1.dingtalk.txt" . }}`)
	s.Require().NoError(err)
	s.Require().NotEmpty(txt)
}
