package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/msgcenter/api/graphql/generated"
	"github.com/woocoos/msgcenter/api/graphql/model"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgalert"
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtype"
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

// SubscriberUsers is the resolver for the subscriberUsers field.
func (r *msgTypeResolver) SubscriberUsers(ctx context.Context, obj *ent.MsgType) ([]*ent.MsgSubscriber, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Client.MsgSubscriber.Query().Where(
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
	return r.Client.MsgSubscriber.Query().Where(
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
	return r.Client.MsgSubscriber.Query().Where(
		msgsubscriber.MsgTypeID(obj.ID),
		msgsubscriber.TenantID(tid),
		msgsubscriber.Exclude(true),
		msgsubscriber.UserIDNotNil(),
		msgsubscriber.OrgRoleIDIsNil(),
	).All(ctx)
}

// MsgChannels is the resolver for the msgChannels field.
func (r *queryResolver) MsgChannels(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MsgChannelOrder, where *ent.MsgChannelWhereInput) (*ent.MsgChannelConnection, error) {
	return r.Client.MsgChannel.Query().Paginate(ctx, after, first, before, last, ent.WithMsgChannelOrder(orderBy), ent.WithMsgChannelFilter(where.Filter))
}

// MsgTypes is the resolver for the msgTypes field.
func (r *queryResolver) MsgTypes(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MsgTypeOrder, where *ent.MsgTypeWhereInput) (*ent.MsgTypeConnection, error) {
	return r.Client.MsgType.Query().Paginate(ctx, after, first, before, last, ent.WithMsgTypeOrder(orderBy), ent.WithMsgTypeFilter(where.Filter))
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
	return r.Client.MsgType.Query().Where(where...).Select(msgtype.FieldCategory).GroupBy(msgtype.FieldCategory).Strings(ctx)
}

// MsgEvents is the resolver for the msgEvents field.
func (r *queryResolver) MsgEvents(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MsgEventOrder, where *ent.MsgEventWhereInput) (*ent.MsgEventConnection, error) {
	return r.Client.MsgEvent.Query().Paginate(ctx, after, first, before, last, ent.WithMsgEventOrder(orderBy), ent.WithMsgEventFilter(where.Filter))
}

// MsgTemplates is the resolver for the msgTemplates field.
func (r *queryResolver) MsgTemplates(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MsgTemplateOrder, where *ent.MsgTemplateWhereInput) (*ent.MsgTemplateConnection, error) {
	return r.Client.MsgTemplate.Query().Paginate(ctx, after, first, before, last, ent.WithMsgTemplateOrder(orderBy), ent.WithMsgTemplateFilter(where.Filter))
}

// Silences is the resolver for the silences field.
func (r *queryResolver) Silences(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.SilenceOrder, where *ent.SilenceWhereInput) (*ent.SilenceConnection, error) {
	return r.Client.Silence.Query().Paginate(ctx, after, first, before, last,
		ent.WithSilenceOrder(orderBy),
		ent.WithSilenceFilter(where.Filter))
}

// MsgAlerts is the resolver for the msgAlerts field.
func (r *queryResolver) MsgAlerts(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.MsgAlertOrder, where *ent.MsgAlertWhereInput) (*ent.MsgAlertConnection, error) {
	return r.Client.MsgAlert.Query().Where(msgalert.Deleted(false)).Paginate(ctx, after, first, before, last,
		ent.WithMsgAlertOrder(orderBy),
		ent.WithMsgAlertFilter(where.Filter))
}

// Matchers is the resolver for the matchers field.
func (r *routeResolver) Matchers(ctx context.Context, obj *profile.Route) ([]*label.Matcher, error) {
	return obj.Matchers, nil
}

// Route returns generated.RouteResolver implementation.
func (r *Resolver) Route() generated.RouteResolver { return &routeResolver{r} }

type routeResolver struct{ *Resolver }
