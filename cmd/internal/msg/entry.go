package msg

import (
	"context"
	"entgo.io/contrib/entgql"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/pkg/store/redisx"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/gqlx"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/ent"
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
		gql.RegistryMiddleware(),
		identity.RegistryTenantIDMiddleware(),
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
			CheckOrigin: webSocketCheckOrigin(cnf.Sub("webSocket")),
		},
		InitFunc:  s.wsInit,
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

	gqlsrv.AroundResponses(gqlx.ContextCache())
	gqlsrv.AroundResponses(gqlx.SimplePagination())
	// mutation事务
	gqlsrv.Use(entgql.Transactioner{TxOpener: s.dbClient})
	if err := gql.RegisterGraphqlServer(s.webSrv, gqlsrv); err != nil {
		log.Fatal(err)
	}
}

// websocket 初始化连接,只做了简单的验证,根据需求.看是否需要提前验证.
func (s *Server) wsInit(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
	bearer := initPayload.Authorization()
	if bearer == "" {
		return nil, errors.New("authorization required")
	}
	gctx, err := gql.FromIncomingContext(ctx)
	if err != nil {
		return nil, err
	}
	gctx.Request.Header.Set("Authorization", bearer)
	tidstr := initPayload.GetString(identity.TenantHeaderKey)
	tid, err := strconv.Atoi(tidstr)
	if err != nil {
		return nil, fmt.Errorf("invalid tenant id %s:%v", tidstr, err)
	}
	gctx.Request.Header.Set(identity.TenantHeaderKey, tidstr)
	ctx = identity.WithTenantID(ctx, tid)
	return ctx, nil
}

func (s *Server) wsError(ctx context.Context, err error) {
	_, ok := err.(transport.WebsocketError)
	if ok {
		s.subs.RemoveConn(ctx)
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

type wsConf struct {
	Origins []string `json:"origins"`
}

func webSocketCheckOrigin(cnf *conf.Configuration) func(r *http.Request) bool {
	wsc := &wsConf{}
	cnf.Unmarshal(wsc)
	return func(r *http.Request) bool {
		if len(wsc.Origins) > 0 {
			ro := r.Header.Get("Origin")
			for _, o := range wsc.Origins {
				if o == ro {
					return true
				}
			}
			return false
		}
		return true
	}
}
