package graphql

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/knockout-go/pkg/middleware"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/service"
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

	//gql without websocket
	gqlsrv := handler.New(NewSchema(
		WithClient(s.DB),
		WithCoordinator(am.Coordinator),
		WithSilences(am.Silences),
		WithKOClient(am.Coordinator.KOSdk),
	))

	gqlsrv.AddTransport(transport.Options{})
	gqlsrv.AddTransport(transport.GET{})
	gqlsrv.AddTransport(transport.POST{})
	gqlsrv.AddTransport(transport.MultipartForm{})

	gqlsrv.SetQueryCache(lru.New(1000))

	gqlsrv.Use(extension.Introspection{})
	gqlsrv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	gqlsrv.AroundResponses(middleware.SimplePagination())
	// mutation事务
	gqlsrv.Use(entgql.Transactioner{TxOpener: s.DB})
	if err := gql.RegisterGraphqlServer(s.WebServer, gqlsrv); err != nil {
		panic(err)
	}
}
