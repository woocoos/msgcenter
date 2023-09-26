package ams

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/httpx"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/pkg/store/redisx"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/gqlx"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/api/oas/server"
	"github.com/woocoos/msgcenter/dispatch"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/metrics"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/provider/mem"
	"github.com/woocoos/msgcenter/service"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

// Server alert server, includes: API提醒服务,包括API及消息分发功能,可选服务包括: UI
type Server struct {
	appCnf   *conf.AppConfiguration
	apiSrv   *web.Server
	dbClient *ent.Client
	apps     []woocoo.Server
	am       *service.AlertManager
}

func NewServer(cnf *conf.AppConfiguration) *Server {
	s := &Server{
		appCnf: cnf,
	}
	s.buildEntClient()
	alertManagerCnf := cnf.Sub("alertManager")
	dataDir := alertManagerCnf.String("storage.path")
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	am, err := service.DefaultAlertManager(alertManagerCnf, service.WithClient(s.dbClient))
	if err != nil {
		log.Fatal(err)
	}
	s.am = am

	metrics.BuildGlobal()
	configCoordinator := service.NewCoordinator(alertManagerCnf)

	cfgopts := &server.Options{
		Coordinator: configCoordinator,
		Alerts:      am.Alerts,
		Silences:    am.Silences,
		StatusFunc:  am.Marker.Status,
		GroupFunc: func(routeFilter func(*dispatch.Route) bool, alertFilter func(*alert.Alert, time.Time) bool) (
			dispatch.AlertGroups, map[label.Fingerprint][]string) {
			return am.Dispatcher.Groups(routeFilter, alertFilter)
		},
	}
	api, err := server.NewService(cfgopts)
	if err != nil {
		log.Fatal(err)
	}

	configCoordinator.SetHttpClient(buildHttpClient(cnf))
	configCoordinator.SetAlertServer(api)

	configCoordinator.SetDBClient(s.dbClient)
	configCoordinator.ReloadHooks(func(cfg *profile.Config) error {
		// TODO tmpl.ExternalURL = am
		err := am.Start(configCoordinator, cfg)
		if err != nil {
			return err
		}
		api.Update(cfg, func(labels label.LabelSet) {
			am.Inhibitor.Mutes(labels)
			am.Silencer.Mutes(labels)
		})

		return nil
	})
	if err := configCoordinator.Reload(); err != nil {
		log.Fatal(err)
	}
	s.buildApiService(cnf, api, cfgopts)
	s.buildPushService()

	s.apps = []woocoo.Server{
		am.Alerts.(*mem.Alerts), am.NotificationLog.(*notify.Log), am.Peer, am.Silences,
		s.apiSrv,
		newPromHttp(cnf.Sub("prometheus")),
		s, // self
	}
	if cnf.IsSet("ui.enabled") {
		s.apps = append(s.apps, buildUiServer(cnf))
	}

	// join before service run
	if err = am.Peer.Join(context.Background()); err != nil {
		log.Fatal(err)
	}

	return s
}

// Servers 需要被woocoo.App 注册的服务
func (s *Server) Servers() []woocoo.Server {
	return s.apps
}

func (s *Server) Start(ctx context.Context) error {
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.am.Stop()
	s.dbClient.Close()
	return nil
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

func (s *Server) buildApiService(cnf *conf.AppConfiguration, ws *server.Service, amopt *server.Options) {
	webSrv := web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		gql.RegistryMiddleware(),
		identity.RegistryTenantIDMiddleware(),
		//web.RegisterMiddleware(otelweb.NewMiddleware()),
	)
	server.RegisterHandlers(webSrv.Router().FindGroup("/api/v2").Group, ws)
	//gql without websocket
	gqlsrv := handler.New(graphql.NewSchema(
		graphql.WithClient(s.dbClient),
		graphql.WithCoordinator(amopt.Coordinator),
		graphql.WithSilences(amopt.Silences),
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

	gqlsrv.AroundResponses(gqlx.ContextCache())
	gqlsrv.AroundResponses(gqlx.SimplePagination())
	// mutation事务
	gqlsrv.Use(entgql.Transactioner{TxOpener: s.dbClient})
	if err := gql.RegisterGraphqlServer(webSrv, gqlsrv); err != nil {
		log.Fatal(err)
	}
	s.apiSrv = webSrv
}

func (s *Server) buildPushService() {
	cli, err := redisx.NewClient(s.appCnf.Sub("store.redis"))
	if err != nil {
		panic(err)
	}
	server.RegisterPushHandlers(s.apiSrv.Router().FindGroup("/api/v2").Group,
		server.NewPushService(cli.UniversalClient),
	)
}

func buildUiServer(cnf *conf.AppConfiguration) *web.Server {
	uiSrv := web.New(web.WithConfiguration(cnf.Sub("ui")),
		web.WithGracefulStop(),
	)
	uiSrv.Router().StaticFS("/", http.Dir(cnf.String("ui.staticDir")))
	return uiSrv
}

func buildHttpClient(cnf *conf.AppConfiguration) *http.Client {
	cfg, err := httpx.NewClientConfig(cnf.Sub("oauth-with-cache"))
	if err != nil {
		log.Fatal(err)
	}
	httpClient, err := cfg.Client(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return httpClient
}

func newPromHttp(cnf *conf.Configuration) woocoo.Server {
	hd := web.New(web.WithConfiguration(cnf))
	hd.Router().GET("/metrics", gin.WrapH(promhttp.Handler()))
	return hd
}
