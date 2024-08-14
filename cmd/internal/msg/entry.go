package msg

import (
	"context"
	"entgo.io/contrib/entgql"
	"errors"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/pkg/store/redisx"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout-go/pkg/middleware"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/ent"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// Server alert server, includes: API提醒服务,包括API及消息分发功能,可选服务包括: UI
type Server struct {
	appCnf    *conf.AppConfiguration
	dbClient  *ent.Client
	msgClient *redisx.Client
	webSrv    *web.Server
	subs      *PubSub
}

func NewServer(cnf *conf.AppConfiguration) *Server {
	s := &Server{
		appCnf: cnf,
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
		User:         "portal",
		OrgRoleUser:  "portal",
		FileIdentity: "portal",
		FileSource:   "portal",
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
	s.subs = NewPubSub(cli)
}

func (s *Server) buildWebServer(cnf *conf.AppConfiguration) {
	s.webSrv = web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		gql.RegisterMiddleware(),
		middleware.RegisterTenantID(),
		middleware.RegisterTokenSigner(),
		//web.RegisterMiddleware(otelweb.NewMiddleware()),
	)
	//gql use msg resolver
	gqlsrv := handler.New(NewSchema(
		graphql.WithClient(s.dbClient),
		graphql.WithMsgClient(s.msgClient.UniversalClient),
		graphql.WithPubSub(s.subs),
	))

	gqlsrv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc:  s.wsInit,
		CloseFunc: s.wsClose,
		ErrorFunc: s.wsError,
	})

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
