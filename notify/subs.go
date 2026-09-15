package notify

import (
	"context"
	"strconv"

	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/service/provider"
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
// Need to not handle subscribing.
func (u EventSubscribeStage) Exec(ctx context.Context, alerts ...*alert.Alert) (context.Context, []*alert.Alert, error) {
	for _, a := range alerts {
		if _, ok := a.Labels[label.SkipSubscribeLabel]; ok {
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

	// Collect user IDs from subscribers.
	seen := make(map[string]struct{})
	var userIDs []string
	for _, ui := range uis {
		if ui.UserID == "" {
			continue
		}
		if _, ok := seen[ui.UserID]; !ok {
			seen[ui.UserID] = struct{}{}
			userIDs = append(userIDs, ui.UserID)
		}
	}

	// 没有订阅用户，走正常流程
	if len(userIDs) == 0 {
		return ctx, alerts, nil
	}

	// Also extract user IDs from alert labels and merge (deduplicated).
	labelUIDs, _ := label.UserIDsFromLabels(ga.Labels)
	for _, id := range labelUIDs {
		uid := strconv.Itoa(id)
		if _, ok := seen[uid]; !ok {
			seen[uid] = struct{}{}
			userIDs = append(userIDs, uid)
		}
	}

	for _, uid := range userIDs {
		uls := make([]*alert.Alert, len(alerts))
		for i, a := range alerts {
			ac := a.Clone()
			ac.Labels[label.ToUserIDLabel] = uid
			ac.Labels[label.SkipSubscribeLabel] = "Y"
			uls[i] = ac
		}
		if err := u.alerts.Put(ctx, uls...); err != nil {
			return ctx, nil, err
		}
	}
	return ctx, nil, nil
}
