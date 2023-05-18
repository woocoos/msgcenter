package profile

import "github.com/woocoos/msgcenter/pkg/label"

// InhibitRule 抑制规则.
type InhibitRule struct {
	// SourceMatchers defines a set of label matchers that have to be fulfilled for source alerts.
	SourceMatchers label.Matchers `yaml:"sourceMatchers,omitempty" json:"sourceMatchers,omitempty"`
	// TargetMatchers defines a set of label matchers that have to be fulfilled for target alerts.
	TargetMatchers label.Matchers `yaml:"targetMatchers,omitempty" json:"targetMatchers,omitempty"`
	// A set of labels that must be equal between the source and target alert
	// for them to be a match.
	Equal []label.LabelName `yaml:"equal,omitempty" json:"equal,omitempty"`
}
