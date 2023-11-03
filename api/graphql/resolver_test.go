package graphql

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo/pkg/gds"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msginternalto"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service/provider/mem"
	"github.com/woocoos/msgcenter/test/testsuite"
	"net/url"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

type resolverSuite struct {
	testsuite.BaseSuite
	resolver  *Resolver
	mr        *mutationResolver
	qr        *queryResolver
	server    *web.Server
	shutdowns []func()
}

func TestRolverSuite(t *testing.T) {
	s := &resolverSuite{
		BaseSuite: testsuite.BaseSuite{
			DSN:        "file:msgcenter?mode=memory&cache=shared&_fk=1",
			DriverName: "sqlite3",
		},
	}
	suite.Run(t, s)
}

// SetupSuite sets up the test suite
func (s *resolverSuite) SetupSuite() {
	err := s.BaseSuite.Setup()
	s.Require().NoError(err)
	s.server = web.New(web.WithConfiguration(s.Cnf.Sub("web")))
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
	go alerts.Start(nil)
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

func (s *resolverSuite) TestCreateSilence() {
	ctx := s.NewTestCtx()
	silence, err := s.mr.CreateSilence(ctx, ent.CreateSilenceInput{
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
