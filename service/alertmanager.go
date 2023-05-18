package service

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/gds/timeinterval"
	"github.com/woocoos/msgcenter/dispatch"
	"github.com/woocoos/msgcenter/inhibit"
	"github.com/woocoos/msgcenter/nflog"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/provider"
	"github.com/woocoos/msgcenter/provider/mem"
	"github.com/woocoos/msgcenter/silence"
	"time"
)

// NotifyMinTimeout is the minimum timeout that is set for the context of a call
// to a notification pipeline.
const NotifyMinTimeout = 10 * time.Second

type AlertManager struct {
	Cnf             *conf.Configuration
	NotificationLog notify.NotificationLog
	Silences        *silence.Silences
	Marker          alert.Marker
	Alerts          provider.Alerts
	Dispatcher      *dispatch.Dispatcher
	Inhibitor       *inhibit.Inhibitor
	Silencer        *silence.Silencer
}

func DefaultAlertManager(cnf *conf.Configuration) (am *AlertManager, err error) {
	am = &AlertManager{Cnf: cnf}
	am.NotificationLog, err = nflog.NewFromConfiguration(cnf)
	if err != nil {
		return nil, err
	}
	am.Silences, err = silence.NewFromConfiguration(cnf)
	if err != nil {
		return nil, err
	}
	am.Marker = alert.NewMarker(prometheus.DefaultRegisterer)
	am.Alerts, err = mem.NewAlerts(context.Background(), am.Marker,
		am.Cnf.Duration("alerts.gcInterval"), nil)
	if err != nil {
		return nil, err
	}
	return am, nil
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
	)
	am.Dispatcher = dispatch.NewDispatcher(am.Alerts, routes, pipeline, am.Marker, timeoutFunc)
	routes.Apply(am.Cnf)

	go am.Dispatcher.Run()
	go am.Inhibitor.Start(context.Background())

	return nil
}

func (am *AlertManager) Stop() {
	am.Dispatcher.Stop()
	am.Inhibitor.Stop(context.Background())
}
