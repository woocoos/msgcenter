package profile

import (
	"fmt"
	"io"
	"strconv"
)

// ReceiverConfigs is a union type for all receiver configs.
type ReceiverConfigs interface {
	EmailConfig | WebhookConfig | MessageConfig
}

type ReceiverType string

const (
	ReceiverEmail   ReceiverType = "email"
	ReceiverMessage ReceiverType = "message"
	ReceiverWebhook ReceiverType = "webhook"
)

func (r ReceiverType) String() string {
	return string(r)
}

func (r ReceiverType) Values() []string {
	return []string{
		ReceiverEmail.String(),
		ReceiverMessage.String(),
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
	case ReceiverEmail, ReceiverMessage, ReceiverWebhook:
		return nil
	default:
		return fmt.Errorf("invalid enum value for receiver field: %q", input)
	}
}

// Receiver configuration provides configuration on how to contact a receiver.
type Receiver struct {
	// A unique identifier for this receiver.
	Name string `yaml:"name" json:"name"`

	// MessageConfigs indicates whether using the internal message system.
	MessageConfig  *MessageConfig   `yaml:"messageConfig,omitempty" json:"messageConfig,omitempty"`
	EmailConfigs   []*EmailConfig   `yaml:"emailConfigs,omitempty" json:"emailConfigs,omitempty"`
	WebhookConfigs []*WebhookConfig `yaml:"webhookConfigs,omitempty" json:"webhookConfigs,omitempty"`
}

// TenantReceiverName extends the receiver name with the tenant ID.
func TenantReceiverName(tid string, ori string) string {
	return ori + "_" + tid
}

func AppRouteName(aid string, rname string) string {
	return rname + "_" + aid
}

// ReceiverKey identifies a receiver with the position of a receiver group.
type ReceiverKey struct {
	// Configured name of the receiver.
	Name string
	// Name of the integration of the receiver.
	Integration string
	// Index of the receiver with respect to the integration.
	// Every integration in a group may have 0..N configurations.
	Index uint32
}
