package profile

import (
	"fmt"
	"io"
	"strconv"
)

// ReceiverConfigs is a union type for all receiver configs.
type ReceiverConfigs interface {
	EmailConfig | WebhookConfig
}

type ReceiverType string

const (
	ReceiverEmail   ReceiverType = "email"
	ReceiverWebhook ReceiverType = "webhook"
)

func (r ReceiverType) String() string {
	return string(r)
}

func (r ReceiverType) Values() []string {
	return []string{
		ReceiverEmail.String(),
		ReceiverWebhook.String(),
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (r ReceiverType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(r.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (r *ReceiverType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return nil
	}
	*r = ReceiverType(str)
	if err := ReceiverTypeValidator(*r); err != nil {
		return fmt.Errorf("%s is not a valid ReceiverType", str)
	}
	return nil
}

func ReceiverTypeValidator(input ReceiverType) error {
	switch input {
	case ReceiverEmail, ReceiverWebhook:
		return nil
	default:
		return fmt.Errorf("invalid enum value for receiver field: %q", input)
	}
}

// Receiver configuration provides configuration on how to contact a receiver.
type Receiver struct {
	// A unique identifier for this receiver.
	Name string `yaml:"name" json:"name"`

	EmailConfigs   []*EmailConfig   `yaml:"emailConfigs,omitempty" json:"emailConfigs,omitempty"`
	WebhookConfigs []*WebhookConfig `yaml:"webhookConfigs,omitempty" json:"webhookConfigs,omitempty"`
}

func TenantReceiverName(tid string, ori string) string {
	return ori + "_" + tid
}
