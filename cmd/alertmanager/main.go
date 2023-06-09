package main

import (
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/gqlx"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/api/graphql/generated"
	"github.com/woocoos/msgcenter/api/oas/server"
	"github.com/woocoos/msgcenter/dispatch"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/metrics"
	"github.com/woocoos/msgcenter/nflog"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/provider/mem"
	"github.com/woocoos/msgcenter/service"
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

	am, err := service.DefaultAlertManager(alertManagerCnf)
	if err != nil {
		log.Fatal(err)
	}
	defer am.Stop()

	metrics.BuildGlobal(prometheus.DefaultRegisterer)
	configCoordinator := service.NewCoordinator(alertManagerCnf)

	cfgopts := &server.Options{
		Coordinator: configCoordinator,
		Alerts:      am.Alerts,
		Silences:    am.Silences,
		StatusFunc:  am.Marker.Status,
		Registry:    prometheus.DefaultRegisterer,
		GroupFunc: func(routeFilter func(*dispatch.Route) bool, alertFilter func(*alert.Alert, time.Time) bool) (
			dispatch.AlertGroups, map[label.Fingerprint][]string) {
			return am.Dispatcher.Groups(routeFilter, alertFilter)
		},
	}
	api, err := server.NewService(cfgopts)
	if err != nil {
		log.Fatal(err)
	}

	configCoordinator.SetDBClient(dbClient)
	configCoordinator.Subscribe(func(cfg *profile.Config) error {
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

	apiSrv := buildWebServer(cnf, api, configCoordinator)
	app.RegisterServer(am.Alerts.(*mem.Alerts), am.NotificationLog.(*nflog.Log), am.Silences, apiSrv)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func buildEntClient(cnf *conf.AppConfiguration) *ent.Client {
	pd := oteldriver.BuildOTELDriver(cnf, "store.msgcenter")
	pd = ecx.BuildEntCacheDriver(cnf, pd)
	if cnf.Development {
		dbClient = ent.NewClient(ent.Driver(pd), ent.Debug())
		//casbinClient = casbinent.NewClient(casbinent.Driver(pd), casbinent.Debug())
	} else {
		dbClient = ent.NewClient(ent.Driver(pd))
		//casbinClient = casbinent.NewClient(casbinent.Driver(pd))
	}
	return dbClient
}

func buildWebServer(cnf *conf.AppConfiguration, ws *server.Service, co *service.Coordinator) *web.Server {
	webSrv := web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(gql.New()),
		//web.RegisterMiddleware(otelweb.NewMiddleware()),
	)
	server.RegisterHandlers(webSrv.Router().FindGroup("/api/v2").Group, ws)
	//gql
	gqlsrv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graphql.Resolver{
				Client:      dbClient,
				Coordinator: co,
			},
		}))
	gqlsrv.AroundResponses(gqlx.ContextCache())
	gqlsrv.AroundResponses(gqlx.SimplePagination())
	// mutation事务
	gqlsrv.Use(entgql.Transactioner{TxOpener: dbClient})
	if err := gql.RegisterGraphqlServer(webSrv, gqlsrv); err != nil {
		log.Fatal(err)
	}
	return webSrv
}
