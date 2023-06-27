package notify

import (
	"context"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/provider"
)

// UserInfo include base user notify info for event type alert which not specify the recipient.
type UserInfo struct {
	UserID string
	Name   string
	Email  string
	Mobile string
}

// Subscriber is an interface for subscription
type Subscriber interface {
	// SubUsers returns a list of subscribers by alert event name.if you don't have an event name,will return empty.
	SubUsers(context.Context, *alert.Alert) ([]UserInfo, error)
}

// EventSubscribeStage is a stage for if the alert is event type and not specify the recipient,then subscribe the user.
type EventSubscribeStage struct {
	alerts provider.Alerts
	Subs   Subscriber
}

func NewEventSubscribeStage(alerts provider.Alerts, subs Subscriber) *EventSubscribeStage {
	return &EventSubscribeStage{
		alerts: alerts,
		Subs:   subs,
	}
}

// Exec implements the Stage interface.
// If the alert has a label "to", it will be used as the recipient.means the alert will be sent to the user.
// need not handle subscribe.
func (u EventSubscribeStage) Exec(ctx context.Context, alerts ...*alert.Alert) (context.Context, []*alert.Alert, error) {
	for _, a := range alerts {
		if _, ok := a.Labels[label.ToUserIDLabel]; ok {
			return ctx, alerts, nil
		}
	}
	return u.exec(ctx, alerts...)
}

func (u EventSubscribeStage) exec(ctx context.Context, alerts ...*alert.Alert) (context.Context, []*alert.Alert, error) {
	ga := alerts[0]
	uis, err := u.Subs.SubUsers(ctx, ga)
	if err != nil {
		return ctx, alerts, err
	}
	if len(uis) == 0 {
		return ctx, alerts, nil
	}
	for _, ui := range uis {
		// copy alerts
		uls := make([]*alert.Alert, len(alerts))
		for i, a := range alerts {
			ac := a.Clone()
			ac.Labels[label.ToUserIDLabel] = ui.UserID
			uls[i] = ac
		}
		if err := u.alerts.Put(uls...); err != nil {
			return ctx, nil, err
		}
	}
	return ctx, nil, nil
}
