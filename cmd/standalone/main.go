package main

import (
	"context"
	"flag"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/api/oas"
	"github.com/woocoos/msgcenter/cmd/internal/ams"
	"github.com/woocoos/msgcenter/cmd/internal/msg"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/metrics"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

var (
	amsConfig = flag.String("r", "../ams", "rms etc dir")
	msgConfig = flag.String("m", "../msg", "msg etc dir")
)

func main() {
	flag.Parse()
	app := woocoo.New()
	metrics.BuildGlobal()

	amscnf := conf.New(conf.WithBaseDir(*amsConfig), conf.WithGlobal(false)).Load()
	app.AppConfiguration().Configuration = amscnf
	koapp.BuildCacheComponents(app.AppConfiguration())
	am, err := service.NewAlertManager(app)
	if err != nil {
		panic(err)
	}
	defer func() {
		am.Stop()
	}()

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

	if amscnf.IsSet("ui.enabled") {
		app.RegisterServer(ams.BuildUiServer(amscnf.Sub("ui")))
	}

	app.RegisterServer(ams.NewPromHttp(amscnf.Sub("prometheus")))

	// join before service run
	if am.Peer != nil {
		if err = am.Peer.Join(context.Background()); err != nil {
			panic(err)
		}
	}

	msgc := conf.New(conf.WithBaseDir(*msgConfig), conf.WithGlobal(false)).Load()
	msgcnf := &conf.AppConfiguration{Configuration: msgc}
	koapp.BuildCacheComponents(msgcnf)
	msgSvr := msg.NewServer(msgcnf)
	app.RegisterServer(msgSvr)
	if clientx.ChangeSet != nil {
		app.RegisterServer(clientx.ChangeSet)
	}
	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
