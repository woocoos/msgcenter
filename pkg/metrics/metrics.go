package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Registerer  = prometheus.DefaultRegisterer
	Dispatcher  *DispatcherMetrics
	Notify      *NotifyMetrics
	Coordinator *CoordinatorMetrics
	Silences    *SilencesMetrics
	Marker      *MarkerMetrics
	Nflog       *NflogMetrics
)

// BuildGlobal build default metrics
func BuildGlobal() {
	Dispatcher = NewDispatcherMetrics(false, Registerer)
	Notify = NewNotifyMetrics([]string{"other", "clientError", "serverError"}, prometheus.DefaultRegisterer)
	Coordinator = NewCoordinatorMetrics(prometheus.DefaultRegisterer)
}

// Alerts stores metrics for alerts which are common across all API versions.
type Alerts struct {
	firing   prometheus.Counter
	resolved prometheus.Counter
	invalid  prometheus.Counter
}

// NewAlerts returns an *Alerts struct for the given API version.
func NewAlerts(version string) *Alerts {
	numReceivedAlerts := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:        "msgcenter_alerts_received_total",
		Help:        "The total number of received alerts.",
		ConstLabels: prometheus.Labels{"version": version},
	}, []string{"status"})
	numInvalidAlerts := prometheus.NewCounter(prometheus.CounterOpts{
		Name:        "msgcenter_alerts_invalid_total",
		Help:        "The total number of received alerts that were invalid.",
		ConstLabels: prometheus.Labels{"version": version},
	})
	Registerer.MustRegister(numReceivedAlerts, numInvalidAlerts)
	return &Alerts{
		firing:   numReceivedAlerts.WithLabelValues("firing"),
		resolved: numReceivedAlerts.WithLabelValues("resolved"),
		invalid:  numInvalidAlerts,
	}
}

// Firing returns a counter of firing alerts.
func (a *Alerts) Firing() prometheus.Counter { return a.firing }

// Resolved returns a counter of resolved alerts.
func (a *Alerts) Resolved() prometheus.Counter { return a.resolved }

// Invalid returns a counter of invalid alerts.
func (a *Alerts) Invalid() prometheus.Counter { return a.invalid }

// CoordinatorMetrics is a metrics for config.coordinator
type CoordinatorMetrics struct {
	ConfigSuccess          prometheus.Gauge
	ConfigSuccessTime      prometheus.Gauge
	ConfiguredReceivers    prometheus.Gauge
	ConfiguredIntegrations prometheus.Gauge
}

func NewCoordinatorMetrics(r prometheus.Registerer) *CoordinatorMetrics {
	configSuccess := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "msgcenter_config_last_reload_successful",
		Help: "Whether the last configuration reload attempt was successful.",
	})
	configSuccessTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "msgcenter_config_last_reload_success_timestamp_seconds",
		Help: "Timestamp of the last successful configuration reload.",
	})
	configuredReceivers := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "msgcenter_receivers",
			Help: "Number of configured receivers.",
		},
	)
	configuredIntegrations := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "msgcenter_integrations",
			Help: "Number of configured integrations.",
		},
	)
	if r != nil {
		r.MustRegister(configSuccess, configSuccessTime, configuredReceivers, configuredIntegrations)
	}
	return &CoordinatorMetrics{
		ConfigSuccess:          configSuccess,
		ConfigSuccessTime:      configSuccessTime,
		ConfiguredReceivers:    configuredReceivers,
		ConfiguredIntegrations: configuredIntegrations,
	}
}

type NotifyMetrics struct {
	NumNotifications                   *prometheus.CounterVec
	NumTotalFailedNotifications        *prometheus.CounterVec
	NumNotificationRequestsTotal       *prometheus.CounterVec
	NumNotificationRequestsFailedTotal *prometheus.CounterVec
	NotificationLatencySeconds         *prometheus.HistogramVec
}

func NewNotifyMetrics(reasons []string, r prometheus.Registerer) *NotifyMetrics {
	m := &NotifyMetrics{
		NumNotifications: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "msgcenter",
			Name:      "notifications_total",
			Help:      "The total number of attempted notifications.",
		}, []string{"integration"}),
		NumTotalFailedNotifications: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "msgcenter",
			Name:      "notifications_failed_total",
			Help:      "The total number of failed notifications.",
		}, []string{"integration", "reason"}),
		NumNotificationRequestsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "msgcenter",
			Name:      "notification_requests_total",
			Help:      "The total number of attempted notification requests.",
		}, []string{"integration"}),
		NumNotificationRequestsFailedTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "msgcenter",
			Name:      "notification_requests_failed_total",
			Help:      "The total number of failed notification requests.",
		}, []string{"integration"}),
		NotificationLatencySeconds: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "msgcenter",
			Name:      "notification_latency_seconds",
			Help:      "The latency of notifications in seconds.",
			Buckets:   []float64{1, 5, 10, 15, 20},
		}, []string{"integration"}),
	}
	for _, integration := range []string{
		"email",
		"pagerduty",
		"wechat",
		"pushover",
		"slack",
		"opsgenie",
		"webhook",
		"victorops",
		"sns",
		"telegram",
	} {
		m.NumNotifications.WithLabelValues(integration)
		m.NumNotificationRequestsTotal.WithLabelValues(integration)
		m.NumNotificationRequestsFailedTotal.WithLabelValues(integration)
		m.NotificationLatencySeconds.WithLabelValues(integration)

		for _, reason := range reasons {
			m.NumTotalFailedNotifications.WithLabelValues(integration, reason)
		}
	}
	r.MustRegister(
		m.NumNotifications, m.NumTotalFailedNotifications,
		m.NumNotificationRequestsTotal, m.NumNotificationRequestsFailedTotal,
		m.NotificationLatencySeconds,
	)
	return m
}

// DispatcherMetrics represents metrics associated to a dispatcher.
type DispatcherMetrics struct {
	AggrGroups            prometheus.Gauge
	ProcessingDuration    prometheus.Summary
	AggrGroupLimitReached prometheus.Counter
}

// NewDispatcherMetrics returns a new registered DispatchMetrics.
func NewDispatcherMetrics(registerLimitMetrics bool, r prometheus.Registerer) *DispatcherMetrics {
	m := &DispatcherMetrics{
		AggrGroups: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "msgcenter_dispatcher_aggregation_groups",
				Help: "Number of active aggregation groups",
			},
		),
		ProcessingDuration: prometheus.NewSummary(
			prometheus.SummaryOpts{
				Name: "msgcenter_dispatcher_alert_processing_duration_seconds",
				Help: "Summary of latencies for the processing of alerts.",
			},
		),
		AggrGroupLimitReached: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "msgcenter_dispatcher_aggregation_group_limit_reached_total",
				Help: "Number of times when dispatcher failed to create new aggregation group due to limit.",
			},
		),
	}

	if r != nil {
		r.MustRegister(m.AggrGroups, m.ProcessingDuration)
		if registerLimitMetrics {
			r.MustRegister(m.AggrGroupLimitReached)
		}
	}

	return m
}

type NflogMetrics struct {
	GCDuration              prometheus.Summary
	SnapshotDuration        prometheus.Summary
	QueriesTotal            prometheus.Counter
	QueryErrorsTotal        prometheus.Counter
	QueryDuration           prometheus.Histogram
	PropagatedMessagesTotal prometheus.Counter
	MaintenanceTotal        prometheus.Counter
	MaintenanceErrorsTotal  prometheus.Counter
}

func NewNflogMetrics() *NflogMetrics {
	m := &NflogMetrics{}

	m.GCDuration = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "msgcenter_nflog_gc_duration_seconds",
		Help:       "Duration of the last notification log garbage collection cycle.",
		Objectives: map[float64]float64{},
	})
	m.MaintenanceTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "msgcenter_nflog_maintenance_total",
		Help: "How many maintenances were executed for the notification log.",
	})
	m.MaintenanceErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "msgcenter_nflog_maintenance_errors_total",
		Help: "How many maintenances were executed for the notification log that failed.",
	})
	m.QueriesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "msgcenter_nflog_queries_total",
		Help: "Number of notification log queries were received.",
	})
	m.QueryErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "msgcenter_nflog_query_errors_total",
		Help: "Number notification log received queries that failed.",
	})
	m.QueryDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "msgcenter_nflog_query_duration_seconds",
		Help: "Duration of notification log query evaluation.",
	})
	m.PropagatedMessagesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "msgcenter_nflog_gossip_messages_propagated_total",
		Help: "Number of received gossip messages that have been further gossiped.",
	})

	Registerer.MustRegister(
		m.GCDuration,
		m.QueriesTotal,
		m.QueryErrorsTotal,
		m.QueryDuration,
		m.PropagatedMessagesTotal,
		m.MaintenanceTotal,
		m.MaintenanceErrorsTotal,
	)
	return m
}

type MarkerMetrics struct {
	alertsActive          prometheus.GaugeFunc
	alertsSuppressed      prometheus.GaugeFunc
	alertStateUnprocessed prometheus.GaugeFunc
}

func NewMarkerMetrics(r prometheus.Registerer, gfunc func(state string) float64) *MarkerMetrics {
	m := &MarkerMetrics{}
	newMarkedAlertMetricByState := func(state string) prometheus.GaugeFunc {
		return prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name:        "msgcenter_marked_alerts",
				Help:        "How many alerts by state are currently marked in the Alertmanager regardless of their expiry.",
				ConstLabels: prometheus.Labels{"state": state},
			},
			func() float64 {
				return gfunc(state)
			},
		)
	}
	m.alertsActive = newMarkedAlertMetricByState("active")
	m.alertsSuppressed = newMarkedAlertMetricByState("suppressed")
	m.alertStateUnprocessed = newMarkedAlertMetricByState("unprocessed")
	r.MustRegister(m.alertsActive)
	r.MustRegister(m.alertsSuppressed)
	r.MustRegister(m.alertStateUnprocessed)
	return m
}

type SilencesMetrics struct {
	GcDuration              prometheus.Summary
	SnapshotDuration        prometheus.Summary
	QueriesTotal            prometheus.Counter
	QueryErrorsTotal        prometheus.Counter
	QueryDuration           prometheus.Histogram
	silencesActive          prometheus.GaugeFunc
	silencesPending         prometheus.GaugeFunc
	silencesExpired         prometheus.GaugeFunc
	propagatedMessagesTotal prometheus.Counter
	MaintenanceTotal        prometheus.Counter
	MaintenanceErrorsTotal  prometheus.Counter
}

func NewSilencesMetrics(sfunc func(state string) float64) *SilencesMetrics {
	m := &SilencesMetrics{}
	m.GcDuration = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "msgcenter_silences_gc_duration_seconds",
		Help:       "Duration of the last silence garbage collection cycle.",
		Objectives: map[float64]float64{},
	})
	m.SnapshotDuration = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "msgcenter_silences_snapshot_duration_seconds",
		Help:       "Duration of the last silence snapshot.",
		Objectives: map[float64]float64{},
	})
	m.MaintenanceTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "msgcenter_silences_maintenance_total",
		Help: "How many maintenances were executed for silences.",
	})
	m.MaintenanceErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "msgcenter_silences_maintenance_errors_total",
		Help: "How many maintenances were executed for silences that failed.",
	})
	m.QueriesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "msgcenter_silences_queries_total",
		Help: "How many silence queries were received.",
	})
	m.QueryErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "msgcenter_silences_query_errors_total",
		Help: "How many silence received queries did not succeed.",
	})
	m.QueryDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "msgcenter_silences_query_duration_seconds",
		Help: "Duration of silence query evaluation.",
	})
	m.propagatedMessagesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "msgcenter_silences_gossip_messages_propagated_total",
		Help: "Number of received gossip messages that have been further gossiped.",
	})
	newSilenceMetricByState := func(st string, sfunc func(state string) float64) prometheus.GaugeFunc {
		return prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name:        "msgcenter_silences",
				Help:        "How many silences by state.",
				ConstLabels: prometheus.Labels{"state": st},
			},
			func() float64 {
				return sfunc(st)
			},
		)
	}
	m.silencesActive = newSilenceMetricByState("active", sfunc)
	m.silencesPending = newSilenceMetricByState("pending", sfunc)
	m.silencesExpired = newSilenceMetricByState("expired", sfunc)

	Registerer.MustRegister(
		m.GcDuration,
		m.SnapshotDuration,
		m.QueriesTotal,
		m.QueryErrorsTotal,
		m.QueryDuration,
		m.silencesActive,
		m.silencesPending,
		m.silencesExpired,
		m.propagatedMessagesTotal,
		m.MaintenanceTotal,
		m.MaintenanceErrorsTotal,
	)
	return m
}
