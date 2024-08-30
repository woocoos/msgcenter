// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// SubscriptionAction is a generic type for all subscription actions
type Message struct {
	Topic   string            `json:"topic"`
	Title   string            `json:"title"`
	Content string            `json:"content"`
	Format  string            `json:"format"`
	URL     string            `json:"url"`
	SendAt  time.Time         `json:"sendAt"`
	Extras  map[string]string `json:"extras"`
}

// MessageFilter is a generic type for all subscription filters
type MessageFilter struct {
	TenantID int    `json:"tenantId"`
	AppCode  string `json:"appCode"`
	UserID   int    `json:"userId"`
	DeviceID string `json:"deviceId"`
}

type Subscription struct {
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
