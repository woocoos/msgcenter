package msg

import (
	"context"
	"time"

	"entgo.io/contrib/entgql"
	gqlgen "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/google/uuid"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/store/redisx"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler/authz"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout-go/pkg/middleware"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/ent"
)

// Server alert server, includes: API提醒服务,包括API及消息分发功能,可选服务包括: UI
type Server struct {
	appCnf                *conf.AppConfiguration
	dbClient              *ent.Client
	msgClient             *redisx.Client
	webSrv                *web.Server
	subs                  *PubSub
	serverID              string
	keepAlivePingInterval time.Duration
}

func NewServer(cnf *conf.AppConfiguration) *Server {
	s := &Server{
		appCnf:                cnf,
		keepAlivePingInterval: 30 * time.Second,
	}
	// 初始化错误处理
	if err := fmterr.InitErrorHandler(cnf.Sub("errors")); err != nil {
		panic(err)
	}
	s.buildEntClient()
	s.buildPubSub()
	s.buildWebServer(cnf)
	return s
}

func (s *Server) buildEntClient() {
	ents := koapp.BuildEntComponents(s.appCnf)
	drv := ents["msgcenter"]

	scfg := ent.AlternateSchema(ent.SchemaConfig{
		User:        "portal",
		Org:         "portal",
		OrgRoleUser: "portal",
		UserAddr:    "portal",
		UserDevice:  "portal",
	})
	if s.appCnf.Development {
		s.dbClient = ent.NewClient(ent.Driver(drv), ent.Debug(), scfg)
	} else {
		s.dbClient = ent.NewClient(ent.Driver(drv), scfg)
	}
}

func (s *Server) buildPubSub() {
	cli, err := redisx.NewClient(s.appCnf.Sub("store.redis"))
	if err != nil {
		panic(err)
	}
	s.msgClient = cli
	s.serverID = uuid.New().String()
	s.subs = NewPubSub(cli, s.serverID)
}

func (s *Server) buildWebServer(cnf *conf.AppConfiguration) {
	s.webSrv = web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		gql.RegisterMiddleware(),
		middleware.RegisterTenantID(),
		middleware.RegisterTokenSigner(),
		otelweb.RegisterMiddleware(),
		web.WithMiddlewareNewFunc("authz", authz.Middleware),
	)

	// 手动创建 gqlgen server，确保 SSE transport 在注册前已添加，
	// 使 buildGraphqlServer 中 SupportStream() 能正确检测并设置 isSupportStream=true
	gqlsrv := gqlgen.New(NewSchema(
		graphql.WithClient(s.dbClient),
		graphql.WithMsgClient(s.msgClient.UniversalClient),
		graphql.WithPubSub(s.subs),
	))
	// SSE 必须在 POST 之前，否则 POST transport 会匹配所有 POST 请求（它不检查 Accept 头），
	// SSE transport 永远不会被选中
	gqlsrv.AddTransport(transport.SSE{
		KeepAlivePingInterval: s.keepAlivePingInterval,
	})
	gqlsrv.AddTransport(transport.Options{})
	gqlsrv.AddTransport(transport.GET{})
	gqlsrv.AddTransport(transport.POST{})
	gqlsrv.AddTransport(transport.MultipartForm{})
	gqlsrv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	// RegisterGraphqlServer 内部调用 SupportStream 发送测试请求，
	// AroundResponses/Use 必须在其之后注册，避免测试请求触发中间件 panic
	if err := gql.RegisterGraphqlServer(s.webSrv, gqlsrv); err != nil {
		panic(err)
	}

	gqlsrv.AroundResponses(middleware.SimplePagination())
	// mutation事务
	gqlsrv.Use(entgql.Transactioner{TxOpener: s.dbClient})
}

func (s *Server) Start(ctx context.Context) error {
	err := s.subs.Start(ctx)
	if err != nil {
		return err
	}
	return s.webSrv.Start(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	s.webSrv.Stop(ctx)
	s.dbClient.Close()
	s.msgClient.Close()
	s.subs.Stop(ctx)
	return nil
}
