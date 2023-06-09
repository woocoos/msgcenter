package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	Dispatcher  *DispatcherMetrics
	Notify      *NotifyMetrics
	Coordinator *CoordinatorMetrics
)

// BuildGlobal build default metrics
func BuildGlobal(registerer prometheus.Registerer) {
	if registerer == nil {
		registerer = prometheus.DefaultRegisterer
	}
	Dispatcher = NewDispatcherMetrics(false, registerer)
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
func NewAlerts(version string, r prometheus.Registerer) *Alerts {
	numReceivedAlerts := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:        "alertmanager_alerts_received_total",
		Help:        "The total number of received alerts.",
		ConstLabels: prometheus.Labels{"version": version},
	}, []string{"status"})
	numInvalidAlerts := prometheus.NewCounter(prometheus.CounterOpts{
		Name:        "alertmanager_alerts_invalid_total",
		Help:        "The total number of received alerts that were invalid.",
		ConstLabels: prometheus.Labels{"version": version},
	})
	if r != nil {
		r.MustRegister(numReceivedAlerts, numInvalidAlerts)
	}
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

// CoordinatorMetrics is a metrics for config.Coordinator
type CoordinatorMetrics struct {
	ConfigSuccess          prometheus.Gauge
	ConfigSuccessTime      prometheus.Gauge
	ConfiguredReceivers    prometheus.Gauge
	ConfiguredIntegrations prometheus.Gauge
}

func NewCoordinatorMetrics(r prometheus.Registerer) *CoordinatorMetrics {
	configSuccess := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "alertmanager_config_last_reload_successful",
		Help: "Whether the last configuration reload attempt was successful.",
	})
	configSuccessTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "alertmanager_config_last_reload_success_timestamp_seconds",
		Help: "Timestamp of the last successful configuration reload.",
	})
	configuredReceivers := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "alertmanager_receivers",
			Help: "Number of configured receivers.",
		},
	)
	configuredIntegrations := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "alertmanager_integrations",
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
			Namespace: "alertmanager",
			Name:      "notifications_total",
			Help:      "The total number of attempted notifications.",
		}, []string{"integration"}),
		NumTotalFailedNotifications: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "alertmanager",
			Name:      "notifications_failed_total",
			Help:      "The total number of failed notifications.",
		}, []string{"integration", "reason"}),
		NumNotificationRequestsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "alertmanager",
			Name:      "notification_requests_total",
			Help:      "The total number of attempted notification requests.",
		}, []string{"integration"}),
		NumNotificationRequestsFailedTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "alertmanager",
			Name:      "notification_requests_failed_total",
			Help:      "The total number of failed notification requests.",
		}, []string{"integration"}),
		NotificationLatencySeconds: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "alertmanager",
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
				Name: "alertmanager_dispatcher_aggregation_groups",
				Help: "Number of active aggregation groups",
			},
		),
		ProcessingDuration: prometheus.NewSummary(
			prometheus.SummaryOpts{
				Name: "alertmanager_dispatcher_alert_processing_duration_seconds",
				Help: "Summary of latencies for the processing of alerts.",
			},
		),
		AggrGroupLimitReached: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "alertmanager_dispatcher_aggregation_group_limit_reached_total",
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
	SnapshotSize            prometheus.Gauge
	QueriesTotal            prometheus.Counter
	QueryErrorsTotal        prometheus.Counter
	QueryDuration           prometheus.Histogram
	PropagatedMessagesTotal prometheus.Counter
	MaintenanceTotal        prometheus.Counter
	MaintenanceErrorsTotal  prometheus.Counter
}

func NewNflogMetrics(r prometheus.Registerer) *NflogMetrics {
	m := &NflogMetrics{}

	m.GCDuration = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "alertmanager_nflog_gc_duration_seconds",
		Help:       "Duration of the last notification log garbage collection cycle.",
		Objectives: map[float64]float64{},
	})
	m.SnapshotDuration = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "alertmanager_nflog_snapshot_duration_seconds",
		Help:       "Duration of the last notification log snapshot.",
		Objectives: map[float64]float64{},
	})
	m.SnapshotSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "alertmanager_nflog_snapshot_size_bytes",
		Help: "Size of the last notification log snapshot in bytes.",
	})
	m.MaintenanceTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "alertmanager_nflog_maintenance_total",
		Help: "How many maintenances were executed for the notification log.",
	})
	m.MaintenanceErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "alertmanager_nflog_maintenance_errors_total",
		Help: "How many maintenances were executed for the notification log that failed.",
	})
	m.QueriesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "alertmanager_nflog_queries_total",
		Help: "Number of notification log queries were received.",
	})
	m.QueryErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "alertmanager_nflog_query_errors_total",
		Help: "Number notification log received queries that failed.",
	})
	m.QueryDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "alertmanager_nflog_query_duration_seconds",
		Help: "Duration of notification log query evaluation.",
	})
	m.PropagatedMessagesTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "alertmanager_nflog_gossip_messages_propagated_total",
		Help: "Number of received gossip messages that have been further gossiped.",
	})

	if r != nil {
		r.MustRegister(
			m.GCDuration,
			m.SnapshotDuration,
			m.SnapshotSize,
			m.QueriesTotal,
			m.QueryErrorsTotal,
			m.QueryDuration,
			m.PropagatedMessagesTotal,
			m.MaintenanceTotal,
			m.MaintenanceErrorsTotal,
		)
	}
	return m
}
