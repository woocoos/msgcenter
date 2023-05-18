// Code generated by woco, DO NOT EDIT.

package oas

import "time"

type Alert struct {
	GeneratorURL string   `binding:"uri,omitempty" json:"generatorURL,omitempty"`
	Labels       LabelSet `binding:"required" json:"labels"`
}

type AlertGroup struct {
	Alerts   []*GettableAlert `binding:"required"`
	Labels   LabelSet         `binding:"required"`
	Receiver Receiver         `binding:"required"`
}

type AlertGroups []*AlertGroup

type AlertStatus struct {
	InhibitedBy []string `binding:"required" json:"inhibitedBy"`
	SilencedBy  []string `binding:"required" json:"silencedBy"`
	State       string   `binding:"required" json:"state"`
}

type AlertmanagerConfig struct {
	Original string `binding:"required" json:"original"`
}

type AlertmanagerStatus struct {
	Cluster     ClusterStatus      `binding:"required" json:"cluster"`
	Config      AlertmanagerConfig `json:"config"`
	Uptime      time.Time          `time_format:"2006-01-02T15:04:05Z07:00" binding:"required" json:"uptime"`
	VersionInfo VersionInfo        `binding:"required" json:"versionInfo"`
}

type ClusterStatus struct {
	Name   string        `json:"name,omitempty"`
	Peers  []*PeerStatus `json:"peers,omitempty"`
	Status string        `binding:"required" json:"status"`
}

type GettableAlert struct {
	*Alert      `json:",inline"`
	Annotations LabelSet    `binding:"required" json:"annotations"`
	EndsAt      time.Time   `time_format:"2006-01-02T15:04:05Z07:00" binding:"required" json:"endsAt"`
	Fingerprint string      `binding:"required" json:"fingerprint"`
	Receivers   []*Receiver `binding:"required" json:"receivers"`
	StartsAt    time.Time   `time_format:"2006-01-02T15:04:05Z07:00" binding:"required" json:"startsAt"`
	Status      AlertStatus `json:"status"`
	UpdatedAt   time.Time   `time_format:"2006-01-02T15:04:05Z07:00" binding:"required" json:"updatedAt"`
}

type GettableAlerts []*GettableAlert

type GettableSilence struct {
	*Silence  `json:",inline"`
	ID        string        `binding:"required" json:"id"`
	Status    SilenceStatus `binding:"required" json:"status"`
	UpdatedAt time.Time     `time_format:"2006-01-02T15:04:05Z07:00" binding:"required" json:"updatedAt"`
}

type GettableSilences []*GettableSilence

type LabelSet map[string]string

type Matcher struct {
	IsEqual bool
	IsRegex bool   `binding:"required"`
	Name    string `binding:"required"`
	Value   string `binding:"required"`
}

type Matchers []*Matcher

type PeerStatus struct {
	Address string `binding:"required"`
	Name    string `binding:"required"`
}

type PostableAlert struct {
	*Alert      `json:",inline"`
	Annotations LabelSet
	EndsAt      time.Time `time_format:"2006-01-02T15:04:05Z07:00"`
	StartsAt    time.Time `time_format:"2006-01-02T15:04:05Z07:00"`
}

type PostableAlerts []*PostableAlert

type PostableSilence struct {
	*Silence `json:",inline"`
	ID       string `json:"id,omitempty"`
}

type Receiver struct {
	Name string `binding:"required"`
}

type Silence struct {
	Comment   string    `binding:"required" json:"comment"`
	CreatedBy string    `binding:"required" json:"createdBy"`
	EndsAt    time.Time `time_format:"2006-01-02T15:04:05Z07:00" binding:"gt,required" json:"endsAt"`
	Matchers  Matchers  `binding:"min=1,omitempty" json:"matchers"`
	StartsAt  time.Time `time_format:"2006-01-02T15:04:05Z07:00" binding:"ltfield=EndsAt,required" json:"startsAt"`
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
