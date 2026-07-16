package msg

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/pkg/store/redisx"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler/authz"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout-go/pkg/middleware"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/ent"
	"go.uber.org/zap"
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
	ss, err := gql.RegisterSchema(s.webSrv, NewSchema(
		graphql.WithClient(s.dbClient),
		graphql.WithMsgClient(s.msgClient.UniversalClient),
		graphql.WithPubSub(s.subs),
	))
	if err != nil {
		panic(err)
	}
	//gql use msg resolver
	gqlsrv := ss[0]
	gqlsrv.AroundResponses(middleware.SimplePagination())
	gqlsrv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: s.keepAlivePingInterval,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc:  s.wsInit,
		CloseFunc: s.wsClose,
		ErrorFunc: s.wsError,
	})

	// mutation事务
	gqlsrv.Use(entgql.Transactioner{TxOpener: s.dbClient})
	if err := gql.RegisterGraphqlServer(s.webSrv, gqlsrv); err != nil {
		log.Fatal(err)
	}
}

// websocket 初始化连接,只做了简单的验证,根据需求.看是否需要提前验证.
func (s *Server) wsInit(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
	bearer := initPayload.Authorization()
	if bearer == "" {
		return nil, nil, errors.New("authorization required")
	}
	gctx, err := gql.FromIncomingContext(ctx)
	if err != nil {
		return nil, nil, err
	}
	gctx.Request.Header.Set("Authorization", bearer)
	tidstr := initPayload.GetString(identity.TenantHeaderKey)
	tid, _ := strconv.Atoi(tidstr)
	if tid == 0 {
		return nil, nil, identity.ErrMisTenantID
	}
	gctx.Request.Header.Set(identity.TenantHeaderKey, tidstr)
	ctx = context.WithValue(ctx, connectionIDKey, uuid.New())
	return ctx, &initPayload, nil
}

func (s *Server) wsClose(ctx context.Context, code int) {
	if err := s.subs.RemoveConn(ctx); err != nil {
		logger.Error("remove ws connect error", zap.Error(err))
	}
}

func (s *Server) wsError(ctx context.Context, err error) {
	var websocketError transport.WebsocketError
	if !errors.As(err, &websocketError) {
		logger.Error("remove ws connect error", zap.Error(err))
	}
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
