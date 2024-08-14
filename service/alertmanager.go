package service

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tsingsun/members"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/gds/timeinterval"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/msgcenter/dispatch"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service/fsclient"
	"github.com/woocoos/msgcenter/service/inhibit"
	"github.com/woocoos/msgcenter/service/provider"
	"github.com/woocoos/msgcenter/service/provider/mem"
	"github.com/woocoos/msgcenter/service/silence"
	"os"
	"time"
)

// NotifyMinTimeout is the minimum timeout that is set for the context of a call
// to a notification pipeline.
const NotifyMinTimeout = 10 * time.Second

type AmOption func(*AlertManager)

func WithClient(client *ent.Client) AmOption {
	return func(am *AlertManager) {
		am.DB = client
	}
}

func WithPeer(p *members.Peer) AmOption {
	return func(am *AlertManager) {
		am.Peer = p
	}
}

type AlertManager struct {
	cnf             *conf.Configuration
	Coordinator     *Coordinator
	NotificationLog notify.NotificationLog
	Silences        *silence.Silences
	Marker          alert.Marker
	Alerts          provider.Alerts
	Dispatcher      *dispatch.Dispatcher
	Inhibitor       *inhibit.Inhibitor
	Silencer        *silence.Silencer
	Subscribe       *UserSubscribe
	DB              *ent.Client
	Peer            *members.Peer
}

func NewAlertManager(app *woocoo.App, opts ...AmOption) (*AlertManager, error) {
	am := &AlertManager{
		cnf:       app.AppConfiguration().Sub("alertManager"),
		Subscribe: &UserSubscribe{},
	}
	for _, opt := range opts {
		opt(am)
	}
	if am.DB == nil {
		am.buildDBClient(app.AppConfiguration())
	}
	if err := am.Members(); err != nil {
		return nil, err
	} else {
		app.RegisterServer(am.Peer)
	}

	err := am.Apply(am.cnf)
	if err != nil {
		return nil, err
	}

	kosdk, err := api.NewSDK(app.AppConfiguration().Sub("kosdk"))
	if err != nil {
		return nil, err
	}
	fs, err := fsclient.NewClient(am.DB, kosdk)
	if err != nil {
		return nil, err
	}
	am.Coordinator = NewCoordinator(am.cnf)
	am.Coordinator.db = am.DB
	am.Subscribe.DB = am.DB
	am.Coordinator.KOSdk = kosdk
	am.Coordinator.FSClient = fs

	app.RegisterServer(am.Alerts, am.NotificationLog, am.Silences)

	return am, nil
}

func (am *AlertManager) Apply(cnf *conf.Configuration) error {
	dataDir := cnf.String("storage.path")
	err := os.MkdirAll(dataDir, os.ModePerm)
	if err != nil {
		return err
	}

	if nflog, err := notify.NewLog(cnf); err != nil {
		return err
	} else {
		if am.Peer != nil {
			if nflog.Spreader, err = am.Peer.AddShard(nflog); err != nil {
				return err
			}
		}
		nflog.NLogCallback = NlogCallback{db: am.DB}
		am.NotificationLog = nflog
	}

	am.Silences, err = silence.NewFromConfiguration(cnf, silence.WithDataLoader(SilencesDataLoad(am.DB)))
	if err != nil {
		return err
	}
	if am.Peer != nil {
		if am.Silences.Spreader, err = am.Peer.AddShard(am.Silences); err != nil {
			return err
		}
	}

	am.Marker = alert.NewMarker(prometheus.DefaultRegisterer)
	am.Alerts, err = mem.NewAlerts(context.Background(), am.Marker,
		am.cnf.Duration("alerts.gcInterval"), &AlertCallback{db: am.DB})
	if err != nil {
		return err
	}

	return nil
}

func (am *AlertManager) buildDBClient(cnf *conf.AppConfiguration) {
	ents := koapp.BuildEntComponents(cnf)
	drv := ents["msgcenter"]
	scfg := ent.AlternateSchema(ent.SchemaConfig{
		User:         "portal",
		OrgRoleUser:  "portal",
		FileIdentity: "portal",
		FileSource:   "portal",
	})
	if cnf.Development {
		am.DB = ent.NewClient(ent.Driver(drv), ent.Debug(), scfg)
	} else {
		am.DB = ent.NewClient(ent.Driver(drv), scfg)
	}
}

func (am *AlertManager) Start(co *Coordinator, config *profile.Config) error {
	waitFunc := func() time.Duration { return 0 }
	timeoutFunc := func(d time.Duration) time.Duration {
		if d < NotifyMinTimeout {
			d = NotifyMinTimeout
		}
		return d + waitFunc()
	}
	pipelineBuilder := notify.NewPipelineBuilder()
	// Build the routing tree and record which receivers are used.
	routes := dispatch.NewRoute(config.Route, nil)
	routes.Walk(func(r *dispatch.Route) {
		co.ActiveReceivers[r.RouteOpts.Receiver] = 0
	})
	// Build the map of receiver to integrations.
	receivers := make(map[string][]notify.Integration, len(co.ActiveReceivers))
	err := co.WalkReceivers(func(rcv profile.Receiver) error {
		integrations, err := co.buildReceiverIntegrations(rcv, co.Template)
		if err != nil {
			return err
		}
		receivers[rcv.Name] = integrations
		co.ActiveReceivers[rcv.Name] = len(integrations)
		return nil
	})
	if err != nil {
		return err
	}

	// Build the map of time interval names to time interval definitions.
	timeIntervals := make(map[string][]timeinterval.TimeInterval, len(config.TimeIntervals))

	for _, ti := range config.TimeIntervals {
		timeIntervals[ti.Name] = ti.TimeIntervals
	}

	am.Inhibitor.Stop(context.Background())
	am.Dispatcher.Stop()

	am.Inhibitor = inhibit.NewInhibitor(am.Alerts, config.InhibitRules, am.Marker)
	am.Silencer = silence.NewSilencer(am.Silences, am.Marker)

	pipeline := pipelineBuilder.New(
		receivers,
		waitFunc,
		am.Inhibitor,
		am.Silencer,
		timeIntervals,
		am.NotificationLog,
		am.Alerts,
		am.Subscribe,
	)
	am.Dispatcher = dispatch.NewDispatcher(am.Alerts, routes, pipeline, am.Marker, timeoutFunc)
	routes.Apply(am.cnf)

	go am.Dispatcher.Run()
	go am.Inhibitor.Start(context.Background())

	return nil
}

func (am *AlertManager) Stop() {
	am.Dispatcher.Stop()
	am.Inhibitor.Stop(context.Background())
	am.DB.Close()
}

func (am *AlertManager) Members() error {
	if !am.cnf.IsSet("cluster") {
		return nil
	}
	cnf := am.cnf.Sub("cluster")
	peer, err := members.NewPeer(members.WithConfiguration(cnf))
	if err != nil {
		return err
	}

	am.Peer = peer

	return nil
}
