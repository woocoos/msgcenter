package types

import (
	"fmt"
	"io"
	"strconv"
)

type Mode string

const (
	ChannelModeNone     Mode = "none"
	ChannelModeInternal Mode = "internal"
	ChannelModeEmail    Mode = "email"
	ChannelModeSMS      Mode = "sms"
	ChannelModeWebhook  Mode = "webhook"
	ChannelModeWechat   Mode = "wechat"
	ChannelModeDingtalk Mode = "dingtalk"
)

func (m Mode) Values() (kinds []string) {
	for _, s := range []Mode{ChannelModeNone, ChannelModeInternal, ChannelModeEmail,
		ChannelModeSMS, ChannelModeWebhook, ChannelModeWechat, ChannelModeDingtalk} {
		kinds = append(kinds, string(s))
	}
	return
}

// ModeValidator is a validator for the "mode" field enum values. It is called by the builders before save.
func ModeValidator(st Mode) error {
	switch st {
	case ChannelModeNone, ChannelModeInternal, ChannelModeEmail,
		ChannelModeSMS, ChannelModeWebhook, ChannelModeWechat, ChannelModeDingtalk:
		return nil
	default:
		return fmt.Errorf("mode: invalid enum value for mode field: %q", st)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (st Mode) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(string(st)))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (st *Mode) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*st = Mode(str)
	if err := ModeValidator(*st); err != nil {
		return fmt.Errorf("%s is not a valid SimpleStatus", str)
	}
	return nil
}
