package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/httpx"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/gqlx"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/api/oas/server"
	"github.com/woocoos/msgcenter/dispatch"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/metrics"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/provider/mem"
	"github.com/woocoos/msgcenter/service"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

var (
	dbClient *ent.Client
	//casbinClient *casbinent.Client
)

func main() {
	app := woocoo.New()
	cnf := app.AppConfiguration()
	dbClient = buildEntClient(cnf)
	defer dbClient.Close()

	alertManagerCnf := cnf.Sub("alertManager")
	dataDir := alertManagerCnf.String("storage.path")
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	am, err := service.DefaultAlertManager(alertManagerCnf, service.WithClient(dbClient))
	if err != nil {
		log.Fatal(err)
	}
	defer am.Stop()

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

	configCoordinator.SetDBClient(dbClient)
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

	apiSrv := buildApiServer(cnf, api, cfgopts)
	app.RegisterServer(am.Alerts.(*mem.Alerts), am.NotificationLog.(*notify.Log), am.Peer, am.Silences, apiSrv,
		newPromHttp(cnf.Sub("prometheus")),
	)

	if cnf.IsSet("ui.enabled") {
		app.RegisterServer(buildUiServer(cnf))
	}

	// join before service run
	if err = am.Peer.Join(context.Background()); err != nil {
		log.Fatal(err)
	}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func buildEntClient(cnf *conf.AppConfiguration) *ent.Client {
	pd := oteldriver.BuildOTELDriver(cnf, "store.msgcenter")
	pd = ecx.BuildEntCacheDriver(cnf, pd)
	scfg := ent.AlternateSchema(ent.SchemaConfig{
		User:        "portal",
		OrgRoleUser: "portal",
	})
	if cnf.Development {
		dbClient = ent.NewClient(ent.Driver(pd), ent.Debug(), scfg)
		//casbinClient = casbinent.NewClient(casbinent.Driver(pd), casbinent.Debug())
	} else {
		dbClient = ent.NewClient(ent.Driver(pd), scfg)
		//casbinClient = casbinent.NewClient(casbinent.Driver(pd))
	}
	return dbClient
}

func buildApiServer(cnf *conf.AppConfiguration, ws *server.Service, amopt *server.Options) *web.Server {
	webSrv := web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(gql.New()),
		identity.RegistryTenantIDMiddleware(),
		//web.RegisterMiddleware(otelweb.NewMiddleware()),
	)
	server.RegisterHandlers(webSrv.Router().FindGroup("/api/v2").Group, ws)
	//gql
	gqlsrv := handler.NewDefaultServer(
		graphql.NewSchema(
			graphql.WithClient(dbClient),
			graphql.WithCoordinator(amopt.Coordinator),
			graphql.WithSilences(amopt.Silences),
		),
	)
	gqlsrv.AroundResponses(gqlx.ContextCache())
	gqlsrv.AroundResponses(gqlx.SimplePagination())
	// mutation事务
	gqlsrv.Use(entgql.Transactioner{TxOpener: dbClient})
	if err := gql.RegisterGraphqlServer(webSrv, gqlsrv); err != nil {
		log.Fatal(err)
	}
	return webSrv
}

func buildUiServer(cnf *conf.AppConfiguration) *web.Server {
	uiSrv := web.New(web.WithConfiguration(cnf.Sub("ui")),
		web.WithGracefulStop(),
	)
	uiSrv.Router().StaticFS("/", http.Dir("../../web/build"))
	return uiSrv
}

func newPromHttp(cnf *conf.Configuration) woocoo.Server {
	hd := web.New(web.WithConfiguration(cnf))
	hd.Router().GET("/metrics", gin.WrapH(promhttp.Handler()))
	return hd
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
