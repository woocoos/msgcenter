// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// SubscriptionAction is a generic type for all subscription actions
type Message struct {
	Action  string `json:"action"`
	Payload string `json:"payload"`
	Key     string `json:"key"`
	Topic   string `json:"topic"`
	SendAt  string `json:"sendAt"`
}

// MessageFilter is a generic type for all subscription filters
type MessageFilter struct {
	AppCode  string `json:"appCode"`
	UserID   int    `json:"userId"`
	DeviceID string `json:"deviceId"`
}

type RouteStrType string

const (
	RouteStrTypeJSON RouteStrType = "Json"
	RouteStrTypeYaml RouteStrType = "Yaml"
)

var AllRouteStrType = []RouteStrType{
	RouteStrTypeJSON,
	RouteStrTypeYaml,
}

func (e RouteStrType) IsValid() bool {
	switch e {
	case RouteStrTypeJSON, RouteStrTypeYaml:
		return true
	}
	return false
}

func (e RouteStrType) String() string {
	return string(e)
}

func (e *RouteStrType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RouteStrType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RouteStrType", str)
	}
	return nil
}

func (e RouteStrType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
