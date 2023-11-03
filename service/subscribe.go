package service

import (
	"context"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/orgroleuser"
	"github.com/woocoos/msgcenter/ent/user"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"strconv"
)

var _ notify.Subscriber = (*UserSubscribe)(nil)

type UserSubscribe struct {
	DB *ent.Client
}

func (u *UserSubscribe) SubUsers(ctx context.Context, al *alert.Alert) ([]notify.UserInfo, error) {
	// alert must have alter name map to an event name.
	eventname, ok := al.Labels[label.AlertNameLabel]
	if !ok {
		return nil, nil
	}
	query := u.DB.MsgSubscriber.Query().Where(
		msgsubscriber.HasMsgTypeWith(msgtype.HasEventsWith(msgevent.Name(eventname))))
	if tid, ok := al.Labels[label.TenantLabel]; ok {
		id, err := strconv.Atoi(tid)
		if err != nil {
			return nil, err
		}
		query.Where(msgsubscriber.TenantID(id))
	}

	subs, err := query.Select(msgsubscriber.FieldOrgRoleID, msgsubscriber.FieldUserID, msgsubscriber.FieldExclude).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var users, groups, excludes []int
	for _, sub := range subs {
		if sub.Exclude {
			excludes = append(excludes, sub.UserID)
		}
		if sub.UserID != 0 {
			users = append(users, sub.UserID)
		}
		if sub.OrgRoleID != 0 {
			groups = append(groups, sub.OrgRoleID)
		}
	}
	ids, err := u.DB.OrgRoleUser.Query().Where(
		orgroleuser.Or(
			orgroleuser.UserIDIn(users...),
			orgroleuser.OrgRoleIDIn(groups...),
		),
		orgroleuser.UserIDNotIn(excludes...),
	).IDs(ctx)
	if err != nil {
		return nil, err
	}
	ul, err := u.DB.User.Query().Where(user.IDIn(ids...)).
		Select(user.FieldID, user.FieldDisplayName, user.FieldEmail, user.FieldMobile).All(ctx)
	if err != nil {
		return nil, err
	}
	var uis []notify.UserInfo
	for _, eu := range ul {
		uis = append(uis, notify.UserInfo{
			UserID: strconv.Itoa(eu.ID),
			Name:   eu.DisplayName,
			Email:  eu.Email,
			Mobile: eu.Mobile,
		})
	}
	return uis, nil
}
