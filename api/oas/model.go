// Code generated by woco, DO NOT EDIT.

package oas

import "time"

type Alert struct {
	GeneratorURL string `binding:"omitempty,uri" json:"generatorURL,omitempty"`
	// Labels A set of labels. Labels are key/value pairs that are attached to
	// alerts. Labels are used to specify identifying attributes of alerts,
	// such as their tenant, user , instance, and job.
	// tenant: specific tenant id.
	// user: specific user id. the user is the notify target. Some notification need info from user, such as email address.
	// alertname: the name of alert.it is also the event name.
	Labels LabelSet `binding:"required" json:"labels"`
}

type AlertGroup struct {
	Alerts []*GettableAlert `binding:"required" json:"alerts"`
	// Labels A set of labels. Labels are key/value pairs that are attached to
	// alerts. Labels are used to specify identifying attributes of alerts,
	// such as their tenant, user , instance, and job.
	// tenant: specific tenant id.
	// user: specific user id. the user is the notify target. Some notification need info from user, such as email address.
	// alertname: the name of alert.it is also the event name.
	Labels   LabelSet `binding:"required" json:"labels"`
	Receiver Receiver `binding:"required" json:"receiver"`
}

type AlertGroups []*AlertGroup

type AlertStatus struct {
	InhibitedBy []string `binding:"required" json:"inhibitedBy"`
	SilencedBy  []int    `binding:"required" json:"silencedBy"`
	State       string   `binding:"required" json:"state"`
}

type AlertmanagerConfig struct {
	Original string `binding:"required" json:"original"`
}

type AlertmanagerStatus struct {
	Cluster     ClusterStatus      `binding:"required" json:"cluster"`
	Config      AlertmanagerConfig `json:"config"`
	Uptime      time.Time          `binding:"required" json:"uptime" time_format:"2006-01-02T15:04:05Z07:00"`
	VersionInfo VersionInfo        `binding:"required" json:"versionInfo"`
}

type ClusterStatus struct {
	Name   string        `json:"name,omitempty"`
	Peers  []*PeerStatus `json:"peers,omitempty"`
	Status string        `binding:"required" json:"status"`
}

type GettableAlert struct {
	*Alert `json:",inline"`
	// Annotations A set of labels. Labels are key/value pairs that are attached to
	// alerts. Labels are used to specify identifying attributes of alerts,
	// such as their tenant, user , instance, and job.
	// tenant: specific tenant id.
	// user: specific user id. the user is the notify target. Some notification need info from user, such as email address.
	// alertname: the name of alert.it is also the event name.
	Annotations LabelSet    `binding:"required" json:"annotations"`
	EndsAt      time.Time   `binding:"required" json:"endsAt" time_format:"2006-01-02T15:04:05Z07:00"`
	Fingerprint string      `binding:"required" json:"fingerprint"`
	Receivers   []*Receiver `binding:"required" json:"receivers"`
	StartsAt    time.Time   `binding:"required" json:"startsAt" time_format:"2006-01-02T15:04:05Z07:00"`
	Status      AlertStatus `json:"status"`
	UpdatedAt   time.Time   `binding:"required" json:"updatedAt" time_format:"2006-01-02T15:04:05Z07:00"`
}

type GettableAlerts []*GettableAlert

type GettableSilence struct {
	*Silence  `json:",inline"`
	ID        int           `binding:"required" json:"id"`
	Status    SilenceStatus `binding:"required" json:"status"`
	UpdatedAt time.Time     `binding:"required" json:"updatedAt" time_format:"2006-01-02T15:04:05Z07:00"`
}

type GettableSilences []*GettableSilence

// LabelSet A set of labels. Labels are key/value pairs that are attached to
// alerts. Labels are used to specify identifying attributes of alerts,
// such as their tenant, user , instance, and job.
// tenant: specific tenant id.
// user: specific user id. the user is the notify target. Some notification need info from user, such as email address.
// alertname: the name of alert.it is also the event name.
type LabelSet map[string]string

type Matcher struct {
	IsEqual bool   `json:"isEqual,omitempty"`
	IsRegex bool   `binding:"required" json:"isRegex"`
	Name    string `binding:"required" json:"name"`
	Value   string `binding:"required" json:"value"`
}

type Matchers []*Matcher

type PeerStatus struct {
	Address string `binding:"required" json:"address"`
	Name    string `binding:"required" json:"name"`
}

type PostableAlert struct {
	*Alert `json:",inline"`
	// Annotations A set of labels. Labels are key/value pairs that are attached to
	// alerts. Labels are used to specify identifying attributes of alerts,
	// such as their tenant, user , instance, and job.
	// tenant: specific tenant id.
	// user: specific user id. the user is the notify target. Some notification need info from user, such as email address.
	// alertname: the name of alert.it is also the event name.
	Annotations LabelSet  `json:"annotations,omitempty"`
	EndsAt      time.Time `json:"endsAt,omitempty" time_format:"2006-01-02T15:04:05Z07:00"`
	StartsAt    time.Time `json:"startsAt,omitempty" time_format:"2006-01-02T15:04:05Z07:00"`
}

type PostableAlerts []*PostableAlert

type PostableSilence struct {
	*Silence `json:",inline"`
	ID       int `json:"id,omitempty"`
}

type Receiver struct {
	Name string `binding:"required" json:"name"`
}

type Silence struct {
	Comment   string    `binding:"required" json:"comment"`
	CreatedBy int       `binding:"required" json:"createdBy"`
	EndsAt    time.Time `binding:"gt,required" json:"endsAt" time_format:"2006-01-02T15:04:05Z07:00"`
	Matchers  Matchers  `binding:"omitempty,min=1" json:"matchers"`
	StartsAt  time.Time `binding:"ltfield=EndsAt,required" json:"startsAt" time_format:"2006-01-02T15:04:05Z07:00"`
	TenantID  int       `binding:"required" json:"tenantID"`
}

type SilenceStatus struct {
	State string `binding:"required" json:"state"`
}

type VersionInfo struct {
	Branch    string `binding:"required" json:"branch"`
	BuildDate string `binding:"required" json:"buildDate"`
	BuildUser string `binding:"required" json:"buildUser"`
	GoVersion string `binding:"required" json:"goVersion"`
	Revision  string `binding:"required" json:"revision"`
	Version   string `binding:"required" json:"version"`
}
