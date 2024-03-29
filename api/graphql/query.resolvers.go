package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql/generated"
	"github.com/woocoos/msgcenter/api/graphql/model"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msginternal"
	"github.com/woocoos/msgcenter/ent/msginternalto"
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/orgroleuser"
	"github.com/woocoos/msgcenter/ent/predicate"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	yaml "gopkg.in/yaml.v3"
)

// RouteStr is the resolver for the routeStr field.
func (r *msgEventResolver) RouteStr(ctx context.Context, obj *ent.MsgEvent, typeArg model.RouteStrType) (string, error) {
	rs, err := json.Marshal(obj.Route)
	if err != nil {
		return "", err
	}
	route := profile.Route{}
	err = json.Unmarshal(rs, &route)
	if err != nil {
		return "", err
	}
	route.Name = ""
	if typeArg == model.RouteStrTypeJSON {
		rs, err = json.Marshal(route)
	} else if typeArg == model.RouteStrTypeYaml {
		rs, err = yaml.Marshal(route)
	} else {
		return "", fmt.Errorf("invalid type")
	}
	if err != nil {
		return "", err
	}
	return string(rs), nil
}

// ToSendCounts is the resolver for the toSendCounts field.
func (r *msgInternalResolver) ToSendCounts(ctx context.Context, obj *ent.MsgInternal) (int, error) {
	return r.client.MsgInternalTo.Query().Where(msginternalto.MsgInternalID(obj.ID)).Count(ctx)
}

// HasReadCounts is the resolver for the hasReadCounts field.
func (r *msgInternalResolver) HasReadCounts(ctx context.Context, obj *ent.MsgInternal) (int, error) {
	return r.client.MsgInternalTo.Query().Where(msginternalto.MsgInternalID(obj.ID), msginternalto.ReadAtNotNil()).Count(ctx)
}

// SubscriberUsers is the resolver for the subscriberUsers field.
func (r *msgTypeResolver) SubscriberUsers(ctx context.Context, obj *ent.MsgType) ([]*ent.MsgSubscriber, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.MsgSubscriber.Query().Where(
		msgsubscriber.MsgTypeID(obj.ID),
		msgsubscriber.TenantID(tid),
		msgsubscriber.Exclude(false),
		msgsubscriber.UserIDNotNil(),
		msgsubscriber.OrgRoleIDIsNil(),
	).All(ctx)
}

// SubscriberRoles is the resolver for the subscriberRoles field.
func (r *msgTypeResolver) SubscriberRoles(ctx context.Context, obj *ent.MsgType) ([]*ent.MsgSubscriber, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.MsgSubscriber.Query().Where(
		msgsubscriber.MsgTypeID(obj.ID),
		msgsubscriber.TenantID(tid),
		msgsubscriber.Exclude(false),
		msgsubscriber.UserIDIsNil(),
		msgsubscriber.OrgRoleIDNotNil(),
	).All(ctx)
}

// ExcludeSubscriberUsers is the resolver for the excludeSubscriberUsers field.
func (r *msgTypeResolver) ExcludeSubscriberUsers(ctx context.Context, obj *ent.MsgType) ([]*ent.MsgSubscriber, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.MsgSubscriber.Query().Where(
		msgsubscriber.MsgTypeID(obj.ID),
		msgsubscriber.TenantID(tid),
		msgsubscriber.Exclude(true),
		msgsubscriber.UserIDNotNil(),
		msgsubscriber.OrgRoleIDIsNil(),
	).All(ctx)
}

// MsgChannels is the resolver for the msgChannels field.
func (r *queryResolver) MsgChannels(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MsgChannelOrder, where *ent.MsgChannelWhereInput) (*ent.MsgChannelConnection, error) {
	return r.client.MsgChannel.Query().Paginate(ctx, after, first, before, last, ent.WithMsgChannelOrder(orderBy), ent.WithMsgChannelFilter(where.Filter))
}

// MsgTypes is the resolver for the msgTypes field.
func (r *queryResolver) MsgTypes(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MsgTypeOrder, where *ent.MsgTypeWhereInput) (*ent.MsgTypeConnection, error) {
	return r.client.MsgType.Query().Paginate(ctx, after, first, before, last, ent.WithMsgTypeOrder(orderBy), ent.WithMsgTypeFilter(where.Filter))
}

// MsgTypeCategories is the resolver for the msgTypeCategories field.
func (r *queryResolver) MsgTypeCategories(ctx context.Context, keyword *string, appID *int) ([]string, error) {
	where := make([]predicate.MsgType, 0)
	if keyword != nil {
		where = append(where, msgtype.CategoryContains(*keyword))
	}
	if appID != nil {
		where = append(where, msgtype.AppID(*appID))
	}
	return r.client.MsgType.Query().Where(where...).Select(msgtype.FieldCategory).GroupBy(msgtype.FieldCategory).Strings(ctx)
}

// MsgEvents is the resolver for the msgEvents field.
func (r *queryResolver) MsgEvents(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MsgEventOrder, where *ent.MsgEventWhereInput) (*ent.MsgEventConnection, error) {
	return r.client.MsgEvent.Query().Paginate(ctx, after, first, before, last, ent.WithMsgEventOrder(orderBy), ent.WithMsgEventFilter(where.Filter))
}

// MsgTemplates is the resolver for the msgTemplates field.
func (r *queryResolver) MsgTemplates(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MsgTemplateOrder, where *ent.MsgTemplateWhereInput) (*ent.MsgTemplateConnection, error) {
	return r.client.MsgTemplate.Query().Paginate(ctx, after, first, before, last, ent.WithMsgTemplateOrder(orderBy), ent.WithMsgTemplateFilter(where.Filter))
}

// Silences is the resolver for the silences field.
func (r *queryResolver) Silences(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.SilenceOrder, where *ent.SilenceWhereInput) (*ent.SilenceConnection, error) {
	return r.client.Silence.Query().Paginate(ctx, after, first, before, last,
		ent.WithSilenceOrder(orderBy),
		ent.WithSilenceFilter(where.Filter))
}

// MsgAlerts is the resolver for the msgAlerts field.
func (r *queryResolver) MsgAlerts(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MsgAlertOrder, where *ent.MsgAlertWhereInput) (*ent.MsgAlertConnection, error) {
	return r.client.MsgAlert.Query().Paginate(ctx, after, first, before, last,
		ent.WithMsgAlertOrder(orderBy),
		ent.WithMsgAlertFilter(where.Filter))
}

// UserMsgInternalTos is the resolver for the userMsgInternalTos field.
func (r *queryResolver) UserMsgInternalTos(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MsgInternalToOrder, where *ent.MsgInternalToWhereInput) (*ent.MsgInternalToConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.MsgInternalTo.Query().Where(msginternalto.UserID(uid), msginternalto.TenantID(tid)).
		Paginate(ctx, after, first, before, last, ent.WithMsgInternalToOrder(orderBy), ent.WithMsgInternalToFilter(where.Filter))
}

// UserSubMsgCategory is the resolver for the userSubMsgCategory field.
func (r *queryResolver) UserSubMsgCategory(ctx context.Context) ([]string, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// 用户加入的role
	orIDs, err := r.client.OrgRoleUser.Query().Where(
		orgroleuser.OrgID(tid), orgroleuser.UserID(uid),
	).Select(orgroleuser.FieldOrgRoleID).Ints(ctx)
	if err != nil {
		return nil, err
	}
	// 查询用户定义的消息类型分类
	categories, err := r.client.MsgType.Query().Where(
		msgtype.HasSubscribersWith(msgsubscriber.Or(msgsubscriber.UserID(uid), msgsubscriber.OrgRoleIDIn(orIDs...)),
			msgsubscriber.TenantID(tid), msgsubscriber.Exclude(false)),
	).Select(msgtype.FieldCategory).Strings(ctx)
	if err != nil {
		return nil, err
	}
	return RemoveDuplicateElement(categories), nil
}

// UserUnreadMsgInternalsFromMsgCategory is the resolver for the userUnreadMsgInternalsFromMsgCategory field.
func (r *queryResolver) UserUnreadMsgInternalsFromMsgCategory(ctx context.Context, categories []string) ([]int, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return []int{}, err
	}
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return []int{}, err
	}
	counts := make([]int, 0)
	for _, v := range categories {
		count, err := r.client.MsgInternalTo.Query().Where(
			msginternalto.UserID(uid), msginternalto.TenantID(tid), msginternalto.ReadAtIsNil(), msginternalto.HasMsgInternalWith(msginternal.Category(v)),
		).Count(ctx)
		if err != nil {
			return []int{}, err
		}
		counts = append(counts, count)
	}
	return counts, nil
}

// UserUnreadMsgInternals is the resolver for the userUnreadMsgInternals field.
func (r *queryResolver) UserUnreadMsgInternals(ctx context.Context) (int, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return 0, err
	}
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return 0, err
	}
	return r.client.MsgInternalTo.Query().Where(
		msginternalto.UserID(uid), msginternalto.TenantID(tid), msginternalto.ReadAtIsNil(),
	).Count(ctx)
}

// Matchers is the resolver for the matchers field.
func (r *routeResolver) Matchers(ctx context.Context, obj *profile.Route) ([]*label.Matcher, error) {
	return obj.Matchers, nil
}

// Route returns generated.RouteResolver implementation.
func (r *Resolver) Route() generated.RouteResolver { return &routeResolver{r} }

type routeResolver struct{ *Resolver }
