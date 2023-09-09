package msg

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/gqlx"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/ent"
)

// Server alert server, includes: API提醒服务,包括API及消息分发功能,可选服务包括: UI
type Server struct {
	appCnf   *conf.AppConfiguration
	dbClient *ent.Client
	webSrv   *web.Server
}

func NewServer(cnf *conf.AppConfiguration) *Server {
	s := &Server{
		appCnf: cnf,
	}
	s.buildEntClient()
	s.buildWebServer(cnf)
	return s
}

func (s *Server) buildEntClient() {
	pd := oteldriver.BuildOTELDriver(s.appCnf, "store.msgcenter")
	pd = ecx.BuildEntCacheDriver(s.appCnf, pd)
	scfg := ent.AlternateSchema(ent.SchemaConfig{
		User:        "portal",
		OrgRoleUser: "portal",
	})
	if s.appCnf.Development {
		s.dbClient = ent.NewClient(ent.Driver(pd), ent.Debug(), scfg)
	} else {
		s.dbClient = ent.NewClient(ent.Driver(pd), scfg)
	}
}

func (s *Server) buildWebServer(cnf *conf.AppConfiguration) {
	s.webSrv = web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(gql.New()),
		identity.RegistryTenantIDMiddleware(),
		//web.RegisterMiddleware(otelweb.NewMiddleware()),
	)
	//gql
	gqlsrv := handler.NewDefaultServer(
		NewSchema(
			graphql.WithClient(s.dbClient),
		),
	)
	gqlsrv.AroundResponses(gqlx.ContextCache())
	gqlsrv.AroundResponses(gqlx.SimplePagination())
	// mutation事务
	gqlsrv.Use(entgql.Transactioner{TxOpener: s.dbClient})
	if err := gql.RegisterGraphqlServer(s.webSrv, gqlsrv); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Start(ctx context.Context) error {
	return s.webSrv.Start(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	s.webSrv.Stop(ctx)
	s.dbClient.Close()
	return nil
}
