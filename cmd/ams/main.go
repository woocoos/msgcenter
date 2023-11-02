package main

import (
	"context"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/api/oas"
	"github.com/woocoos/msgcenter/cmd/internal/ams"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/metrics"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

func main() {
	app := koapp.New()
	cnf := app.AppConfiguration()

	metrics.BuildGlobal()

	am, err := service.NewAlertManager(app)
	if err != nil {
		panic(err)
	}

	adm, err := graphql.NewServer(app, am)
	if err != nil {
		panic(err)
	}
	api, err := oas.NewServer(app, am, adm.WebServer)
	if err != nil {
		panic(err)
	}

	am.Coordinator.ReloadHooks(func(cfg *profile.Config) error {
		// TODO tmpl.ExternalURL = am
		err := am.Start(am.Coordinator, cfg)
		if err != nil {
			return err
		}
		api.Update(cfg, func(labels label.LabelSet) {
			am.Inhibitor.Mutes(labels)
			am.Silencer.Mutes(labels)
		})
		return nil
	})
	if err := am.Coordinator.Reload(); err != nil {
		panic(err)
	}

	if cnf.IsSet("ui.enabled") {
		app.RegisterServer(ams.BuildUiServer(cnf.Sub("ui")))
	}

	app.RegisterServer(ams.NewPromHttp(cnf.Sub("prometheus")))

	// join before service run
	if am.Peer != nil {
		if err = am.Peer.Join(context.Background()); err != nil {
			panic(err)
		}
	}
	defer func() {
		am.Stop()
	}()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
