package server

import (
	"errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/woocoos/msgcenter/dispatch"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/provider"
	"github.com/woocoos/msgcenter/service"
	"github.com/woocoos/msgcenter/silence"
	"time"
)

type Options struct {
	Coordinator *service.Coordinator
	// Alerts to be used by the API. Mandatory.
	Alerts provider.Alerts
	// Silences to be used by the API. Mandatory.
	Silences *silence.Silences
	// StatusFunc is used be the API to retrieve the AlertStatus of an
	// alert. Mandatory.
	StatusFunc func(label.Fingerprint) alert.MarkerStatus
	// Registry is used to register Prometheus metrics. If nil, no metrics
	// registration will happen.
	Registry prometheus.Registerer
	// GroupFunc returns a list of alert groups. The alerts are grouped
	// according to the current active configuration. Alerts returned are
	// filtered by the arguments provided to the function.
	GroupFunc func(func(*dispatch.Route) bool, func(*alert.Alert, time.Time) bool) (dispatch.AlertGroups, map[label.Fingerprint][]string)
}

func (o Options) Validate() error {
	if o.Alerts == nil {
		return errors.New("mandatory field Alerts not set")
	}
	if o.Silences == nil {
		return errors.New("mandatory field Silences not set")
	}
	if o.StatusFunc == nil {
		return errors.New("mandatory field StatusFunc not set")
	}
	if o.GroupFunc == nil {
		return errors.New("mandatory field GroupFunc not set")
	}
	return nil
}
