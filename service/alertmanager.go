package service

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/tsingsun/members"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/gds/timeinterval"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/msgcenter/dispatch"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/marker"
	"github.com/woocoos/msgcenter/pkg/metrics"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service/inhibit"
	"github.com/woocoos/msgcenter/service/kosdk"
	"github.com/woocoos/msgcenter/service/provider"
	"github.com/woocoos/msgcenter/service/provider/mem"
	"github.com/woocoos/msgcenter/service/silence"
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
	Alerts          provider.Alerts
	Dispatcher      atomic.Pointer[dispatch.Dispatcher]
	Inhibitor       atomic.Pointer[inhibit.Inhibitor]
	Silencer        *silence.Silencer
	Subscribe       *UserSubscribe
	DB              *ent.Client
	Peer            *members.Peer
	Route           *dispatch.Route
	groupMarker     marker.GroupMarker
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

	am.Coordinator = NewCoordinator(am.cnf)
	am.Coordinator.db = am.DB
	am.Subscribe.DB = am.DB

	// Load bucket mount paths mapping.
	if am.cnf.IsSet("mountPaths") {
		am.Coordinator.MountPaths = am.cnf.StringMap("mountPaths")
	}

	if app.AppConfiguration().IsSet("kosdk") {
		koSdk, err := kosdk.NewSDK(app.AppConfiguration().Sub("kosdk"), am.DB)
		if err != nil {
			return nil, err
		}
		am.Coordinator.KOSdk = koSdk
	}

	app.RegisterServer(am.Alerts, am.NotificationLog, am.Silences)

	return am, nil
}

func (am *AlertManager) Apply(cnf *conf.Configuration) error {
	var err error
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
	am.groupMarker = marker.NewGroupMarker()

	am.Silences, err = silence.NewFromConfiguration(cnf, silence.WithDataLoader(SilencesDataLoad(am.DB)))
	if err != nil {
		return err
	}
	if am.Peer != nil {
		if am.Silences.Spreader, err = am.Peer.AddShard(am.Silences); err != nil {
			return err
		}
	}

	am.Silencer = silence.NewSilencer(am.Silences, &AlertCallback{db: am.DB})

	am.Alerts, err = mem.NewAlerts(
		am.cnf.Duration("alerts.gcInterval"),
		am.cnf.Int("alerts.perAlertnameLimit"),
		am.Silencer)
	if err != nil {
		return err
	}

	return nil
}

func (am *AlertManager) buildDBClient(cnf *conf.AppConfiguration) {
	ents := koapp.BuildEntComponents(cnf)
	drv := ents["msgcenter"]
	scfg := ent.AlternateSchema(ent.SchemaConfig{
		User:        "portal",
		Org:         "portal",
		OrgRoleUser: "portal",
		UserAddr:    "portal",
	})
	if cnf.Development {
		am.DB = ent.NewClient(ent.Driver(drv), ent.Debug(), scfg)
	} else {
		am.DB = ent.NewClient(ent.Driver(drv), scfg)
	}
	am.DB.User.Use(ReadOnlyHook)
	am.DB.UserAddr.Use(ReadOnlyHook)
	am.DB.UserDevice.Use(ReadOnlyHook)
	am.DB.Org.Use(ReadOnlyHook)
	am.DB.OrgRoleUser.Use(ReadOnlyHook)

}

// ReadOnlyHook keep schema data readonly.
func ReadOnlyHook(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
		return nil, errors.New("not implemented")
	})
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

	if old := am.Inhibitor.Load(); old != nil {
		old.Stop()
	}
	if old := am.Dispatcher.Load(); old != nil {
		old.Stop()
	}

	newInhibitor := inhibit.NewInhibitor(am.Alerts, config.InhibitRules)

	pipeline := pipelineBuilder.New(
		receivers,
		waitFunc,
		newInhibitor,
		am.Silencer,
		timeIntervals,
		am.NotificationLog,
		am.Alerts,
		am.Subscribe,
	)
	newDispatcher := dispatch.NewDispatcher(am.Alerts, routes, pipeline, am.groupMarker, timeoutFunc,
		am.cnf.Duration("dispatch.maintenanceInterval"),
		nil,
		metrics.Dispatcher,
	)
	routes.Apply(am.cnf)
	am.Route = routes

	// First, start the inhibitor so the inhibition cache can populate.
	// Wait for it to load alerts before starting the dispatcher so we
	// don't accidentally notify for an alert that will be inhibited.
	// Publish it only after loading completes: the API mute callback
	// reads r.inhibitor.Load(), so swapping earlier would expose an
	// empty inhibition cache to concurrent requests during a reload (the
	// pipeline already holds newInhibitor directly, and no dispatcher is
	// running to drive it yet, so the old inhibitor stays authoritative
	// for the API until the new one is ready).
	go newInhibitor.Run()
	newInhibitor.WaitForLoading()
	am.Inhibitor.Store(newInhibitor)

	// Next, start the dispatcher and wait for it to load before swapping
	// the dispatcher pointer. This ensures that the API doesn't see the new
	// dispatcher before it finishes populating the aggrGroups.
	go newDispatcher.Run(time.Now().Add(am.cnf.Duration("dispatch.startDelay")))
	newDispatcher.WaitForLoading()
	am.Dispatcher.Store(newDispatcher)

	return nil
}

func (am *AlertManager) Stop() {
	if i := am.Inhibitor.Load(); i != nil {
		i.Stop()
	}
	if d := am.Dispatcher.Load(); d != nil {
		d.Stop()
	}
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
