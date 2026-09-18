package service

import (
	"context"
	"strconv"

	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/orgroleuser"
	"github.com/woocoos/msgcenter/ent/predicate"
	"github.com/woocoos/msgcenter/ent/user"
	"github.com/woocoos/msgcenter/ent/useraddr"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
)

var _ notify.Subscriber = (*UserSubscribe)(nil)

type UserSubscribe struct {
	DB *ent.Client
}

func (u *UserSubscribe) SubUsers(ctx context.Context, al *alert.Alert) ([]notify.UserInfo, error) {
	// alert must have alert name map to an event name.
	eventname, ok := al.Labels[label.AlertNameLabel]
	if !ok {
		return nil, nil
	}

	// 获取租户 ID
	tenantID := 0
	if tid, ok := al.Labels[label.TenantLabel]; ok {
		id, err := strconv.Atoi(tid)
		if err != nil {
			return nil, err
		}
		tenantID = id
	}

	// 如果没有租户 ID，无法确定订阅范围
	if tenantID == 0 {
		return nil, nil
	}

	// 查找消息事件，获取 event ID 和 type ID
	// 注意：事件名称在应用内唯一，租户过滤在订阅查询中进行
	event, err := u.DB.MsgEvent.Query().
		Where(msgevent.Name(eventname)).
		WithMsgType().
		First(ctx)
	if err != nil {
		return nil, nil // 事件不存在，返回空
	}

	// 检查事件是否允许订阅
	if !event.CanSubs {
		return nil, nil // 事件不允许订阅，返回空
	}

	// 构建查询条件：支持事件级订阅和类型级订阅
	var predicates []predicate.MsgSubscriber
	predicates = append(predicates, msgsubscriber.TenantID(tenantID))

	// 匹配条件：msg_event_id = eventID OR msg_type_id = event.MsgTypeID
	predicates = append(predicates, msgsubscriber.Or(
		msgsubscriber.MsgEventID(event.ID),
		msgsubscriber.MsgTypeID(event.Edges.MsgType.ID),
	))

	subs, err := u.DB.MsgSubscriber.Query().Where(predicates...).
		Select(msgsubscriber.FieldOrgRoleID, msgsubscriber.FieldUserID, msgsubscriber.FieldExclude).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var users, groups, excludes []int
	for _, sub := range subs {
		if sub.Exclude {
			if sub.UserID != 0 {
				excludes = append(excludes, sub.UserID)
			}
			continue
		}
		if sub.UserID != 0 {
			users = append(users, sub.UserID)
		}
		if sub.OrgRoleID != 0 {
			groups = append(groups, sub.OrgRoleID)
		}
	}

	// 如果没有订阅者，直接返回
	if len(users) == 0 && len(groups) == 0 {
		return nil, nil
	}

	// 查询用户（使用 map 去重）
	userMap := make(map[int]struct{})

	// 添加直接订阅的用户
	for _, uid := range users {
		userMap[uid] = struct{}{}
	}

	// 添加用户组中的用户
	if len(groups) > 0 {
		userIDs, err := u.DB.OrgRoleUser.Query().Where(
			orgroleuser.OrgRoleIDIn(groups...),
		).Select(orgroleuser.FieldUserID).Ints(ctx)
		if err != nil {
			return nil, err
		}
		for _, uid := range userIDs {
			userMap[uid] = struct{}{}
		}
	}

	// 如果没有用户，返回空
	if len(userMap) == 0 {
		return nil, nil
	}

	// 转换为切片
	allUserIDs := make([]int, 0, len(userMap))
	for uid := range userMap {
		allUserIDs = append(allUserIDs, uid)
	}

	// 构建用户查询
	userQuery := user.IDIn(allUserIDs...)

	// 排除排除列表中的用户
	if len(excludes) > 0 {
		userQuery = user.And(userQuery, user.IDNotIn(excludes...))
	}

	ul, err := u.DB.User.Query().Where(userQuery).
		Select(user.FieldID, user.FieldDisplayName).All(ctx)
	if err != nil {
		return nil, err
	}

	var uis []notify.UserInfo
	for _, eu := range ul {
		// 查询用户联系方式
		addr, err := eu.QueryAddresses().Where(useraddr.AddrTypeEQ(useraddr.AddrTypeContact)).First(ctx)
		if err != nil {
			// 没有联系方式的用户跳过
			continue
		}
		uis = append(uis, notify.UserInfo{
			UserID: strconv.Itoa(eu.ID),
			Name:   eu.DisplayName,
			Email:  addr.Email,
			Mobile: addr.Mobile,
		})
	}
	return uis, nil
}
