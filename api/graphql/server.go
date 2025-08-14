package graphql

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/knockout-go/pkg/middleware"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/service"
	"github.com/woocoos/msgcenter/service/ams"
)

type Server struct {
	DB        *ent.Client
	WebServer *web.Server
}

func NewServer(app *woocoo.App, am *service.AlertManager) (*Server, error) {
	s := &Server{
		DB: am.DB,
	}
	s.buildWebEngine(app, am)

	app.RegisterServer(s.WebServer)
	return s, nil
}

func (s *Server) Start(ctx context.Context) error {
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.DB.Close()
}

func (s *Server) buildWebEngine(app *woocoo.App, am *service.AlertManager) {
	s.WebServer = web.New(web.WithConfiguration(app.AppConfiguration().Sub("web")),
		web.WithGracefulStop(),
		gql.RegisterMiddleware(),
		middleware.RegisterTenantID(),
		otelweb.RegisterMiddleware(),
	)

	amsSvc := ams.NewService(ams.WithClient(s.DB), ams.WithAlertManager(am))
	//gql without websocket
	gqlsrv := handler.NewDefaultServer(NewSchema(
		WithClient(s.DB),
		WithCoordinator(am.Coordinator),
		WithAmsService(amsSvc),
		WithSilences(am.Silences),
		WithKOClient(am.Coordinator.KOSdk),
	))
	gqlsrv.AroundResponses(middleware.SimplePagination())
	// mutation transaction
	gqlsrv.Use(entgql.Transactioner{TxOpener: s.DB})
	if err := gql.RegisterGraphqlServer(s.WebServer, gqlsrv); err != nil {
		panic(err)
	}
}
