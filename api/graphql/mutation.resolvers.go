package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"fmt"
	"strconv"

	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/msgcenter/api/graphql/generated"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
)

// CreateMsgType is the resolver for the createMsgType field.
func (r *mutationResolver) CreateMsgType(ctx context.Context, input ent.CreateMsgTypeInput) (*ent.MsgType, error) {
	return ent.FromContext(ctx).MsgType.Create().SetInput(input).Save(ctx)
}

// UpdateMsgType is the resolver for the updateMsgType field.
func (r *mutationResolver) UpdateMsgType(ctx context.Context, id int, input ent.UpdateMsgTypeInput) (*ent.MsgType, error) {
	return ent.FromContext(ctx).MsgType.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteMsgType is the resolver for the deleteMsgType field.
func (r *mutationResolver) DeleteMsgType(ctx context.Context, id int) (bool, error) {
	client := ent.FromContext(ctx)
	if has, err := client.MsgType.Query().Where(msgtype.HasEvents()).Exist(ctx); err != nil {
		return false, err
	} else if has {
		return false, fmt.Errorf("cannot be deleted. msgtype is associated with msgevent")
	}
	if has, err := client.MsgType.Query().Where(msgtype.HasSubscribers()).Exist(ctx); err != nil {
		return false, err
	} else if has {
		return false, fmt.Errorf("cannot be deleted, msgtype is subscribed")
	}
	err := client.MsgType.DeleteOneID(id).Exec(ctx)
	return err == nil, err
}

// CreateMsgEvent is the resolver for the createMsgEvent field.
func (r *mutationResolver) CreateMsgEvent(ctx context.Context, input ent.CreateMsgEventInput) (*ent.MsgEvent, error) {
	return ent.FromContext(ctx).MsgEvent.Create().SetInput(input).Save(ctx)
}

// UpdateMsgEvent is the resolver for the updateMsgEvent field.
func (r *mutationResolver) UpdateMsgEvent(ctx context.Context, id int, input ent.UpdateMsgEventInput) (*ent.MsgEvent, error) {
	if input.Route != nil {
		// route不为空，验证是否符合modes
		//modes := strings.Split(*input.Modes, ",")

	}
	return ent.FromContext(ctx).MsgEvent.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteMsgEvent is the resolver for the deleteMsgEvent field.
func (r *mutationResolver) DeleteMsgEvent(ctx context.Context, id int) (bool, error) {
	client := ent.FromContext(ctx)
	if has, err := client.MsgEvent.Query().Where(msgevent.ID(id), msgevent.StatusEQ(typex.SimpleStatusActive)).Exist(ctx); err != nil {
		return false, err
	} else if has {
		return false, fmt.Errorf("the active status cannot be deleted")
	}
	if has, err := client.MsgEvent.Query().Where(msgevent.HasCustomerTemplate()).Exist(ctx); err != nil {
		return false, err
	} else if has {
		return false, fmt.Errorf("cannot be deleted. msgevent is associated with msgtemplate")
	}
	err := ent.FromContext(ctx).MsgEvent.DeleteOneID(id).Exec(ctx)
	return err == nil, err
}

// EnableMsgEvent is the resolver for the enableMsgEvent field.
func (r *mutationResolver) EnableMsgEvent(ctx context.Context, id int) (*ent.MsgEvent, error) {
	event, err := ent.FromContext(ctx).MsgEvent.Query().Where(msgevent.ID(id)).WithMsgType().Only(ctx)
	if err != nil {
		return nil, err
	}
	if event.Route == nil {
		return nil, fmt.Errorf("route cannot nil")
	}
	event.Route.Name = profile.AppRouteName(strconv.Itoa(event.Edges.MsgType.AppID), event.Name)
	if err = r.Coordinator.AddNamedRoute([]*profile.Route{event.Route}); err != nil {
		return nil, err
	}
	return event.Update().SetStatus(typex.SimpleStatusActive).Save(ctx)
}

// DisableMsgEvent is the resolver for the disableMsgEvent field.
func (r *mutationResolver) DisableMsgEvent(ctx context.Context, id int) (*ent.MsgEvent, error) {
	event, err := ent.FromContext(ctx).MsgEvent.Query().Where(msgevent.ID(id)).WithMsgType().Only(ctx)
	if err != nil {
		return nil, err
	}
	if err = r.Coordinator.RemoveNamedRoute([]string{profile.AppRouteName(strconv.Itoa(event.Edges.MsgType.AppID), event.Name)}); err != nil {
		return nil, err
	}
	return event.Update().SetStatus(typex.SimpleStatusInactive).Save(ctx)
}

// CreateMsgChannel is the resolver for the createMsgChannel field.
func (r *mutationResolver) CreateMsgChannel(ctx context.Context, input ent.CreateMsgChannelInput) (*ent.MsgChannel, error) {
	return ent.FromContext(ctx).MsgChannel.Create().SetInput(input).Save(ctx)
}

// UpdateMsgChannel is the resolver for the updateMsgChannel field.
func (r *mutationResolver) UpdateMsgChannel(ctx context.Context, id int, input ent.UpdateMsgChannelInput) (*ent.MsgChannel, error) {
	return ent.FromContext(ctx).MsgChannel.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteMsgChannel is the resolver for the deleteMsgChannel field.
func (r *mutationResolver) DeleteMsgChannel(ctx context.Context, id int) (bool, error) {
	client := ent.FromContext(ctx)
	if has, err := client.MsgChannel.Query().Where(msgchannel.ID(id), msgchannel.StatusEQ(typex.SimpleStatusActive)).Exist(ctx); err != nil {
		return false, err
	} else if has {
		return false, fmt.Errorf("the active status cannot be deleted")
	}
	err := client.MsgChannel.DeleteOneID(id).Exec(ctx)
	return err == nil, err
}

// EnableMsgChannel is the resolver for the enableMsgChannel field.
func (r *mutationResolver) EnableMsgChannel(ctx context.Context, id int) (*ent.MsgChannel, error) {
	channel, err := ent.FromContext(ctx).MsgChannel.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if channel.Receiver == nil {
		return nil, fmt.Errorf("receiver cannot nil")
	}
	channel.Receiver.Name = profile.TenantReceiverName(strconv.Itoa(channel.TenantID), channel.Name)
	if err = r.Coordinator.AddTenantReceiver([]*profile.Receiver{channel.Receiver}); err != nil {
		return nil, err
	}
	return channel.Update().SetStatus(typex.SimpleStatusActive).Save(ctx)
}

// DisableMsgChannel is the resolver for the disableMsgChannel field.
func (r *mutationResolver) DisableMsgChannel(ctx context.Context, id int) (*ent.MsgChannel, error) {
	channel, err := ent.FromContext(ctx).MsgChannel.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if err = r.Coordinator.RemoveTenantReceiver([]string{profile.TenantReceiverName(strconv.Itoa(channel.TenantID), channel.Name)}); err != nil {
		return nil, err
	}
	return channel.Update().SetStatus(typex.SimpleStatusInactive).Save(ctx)
}

// CreateMsgTemplate is the resolver for the createMsgTemplate field.
func (r *mutationResolver) CreateMsgTemplate(ctx context.Context, input ent.CreateMsgTemplateInput) (*ent.MsgTemplate, error) {
	return ent.FromContext(ctx).MsgTemplate.Create().SetInput(input).Save(ctx)
}

// UpdateMsgTemplate is the resolver for the updateMsgTemplate field.
func (r *mutationResolver) UpdateMsgTemplate(ctx context.Context, id int, input ent.UpdateMsgTemplateInput) (*ent.MsgTemplate, error) {
	return ent.FromContext(ctx).MsgTemplate.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteMsgTemplate is the resolver for the deleteMsgTemplate field.
func (r *mutationResolver) DeleteMsgTemplate(ctx context.Context, id int) (bool, error) {
	client := ent.FromContext(ctx)
	if has, err := client.MsgTemplate.Query().Where(msgtemplate.ID(id), msgtemplate.StatusEQ(typex.SimpleStatusActive)).Exist(ctx); err != nil {
		return false, err
	} else if has {
		return false, fmt.Errorf("the active status cannot be deleted")
	}
	err := client.MsgTemplate.DeleteOneID(id).Exec(ctx)
	return err == nil, err
}

// EnableMsgTemplate is the resolver for the enableMsgTemplate field.
func (r *mutationResolver) EnableMsgTemplate(ctx context.Context, id int) (*ent.MsgTemplate, error) {
	temp, err := ent.FromContext(ctx).MsgTemplate.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// TODO 启用模板时需加载模板
	return temp.Update().SetStatus(typex.SimpleStatusActive).Save(ctx)
}

// DisableMsgTemplate is the resolver for the disableMsgTemplate field.
func (r *mutationResolver) DisableMsgTemplate(ctx context.Context, id int) (*ent.MsgTemplate, error) {
	temp, err := ent.FromContext(ctx).MsgTemplate.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// TODO 禁用模板时需移除模板
	return temp.Update().SetStatus(typex.SimpleStatusInactive).Save(ctx)
}

// CreateMsgSubscriber is the resolver for the createMsgSubscriber field.
func (r *mutationResolver) CreateMsgSubscriber(ctx context.Context, inputs []*ent.CreateMsgSubscriberInput) ([]*ent.MsgSubscriber, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	mss := make([]*ent.MsgSubscriberCreate, 0)
	for _, v := range inputs {
		if v.TenantID != tid {
			return nil, fmt.Errorf("invalid tenantID")
		}
		if v.Exclude != nil && *v.Exclude && v.OrgRoleID != nil {
			return nil, fmt.Errorf("orgRole cannot exclude")
		}
		if (v.UserID == nil && v.OrgRoleID == nil) || (v.UserID != nil && v.OrgRoleID != nil) {
			return nil, fmt.Errorf("only one of userID and orgRoleID")
		}
		// 检查是否已订阅
		exclude := false
		if v.Exclude != nil {
			exclude = *v.Exclude
		}
		uid := 0
		if v.UserID != nil {
			uid = *v.UserID
		}
		orid := 0
		if v.OrgRoleID != nil {
			orid = *v.OrgRoleID
		}
		has, err := client.MsgSubscriber.Query().Where(
			msgsubscriber.TenantID(v.TenantID),
			msgsubscriber.MsgTypeID(v.MsgTypeID),
			msgsubscriber.ExcludeEQ(exclude),
			msgsubscriber.Or(msgsubscriber.UserIDEQ(uid), msgsubscriber.OrgRoleIDEQ(orid))).Exist(ctx)
		if err != nil {
			return nil, err
		}
		if has {
			return nil, fmt.Errorf("subscription has existed")
		}
		mss = append(mss, client.MsgSubscriber.Create().SetInput(*v))
	}
	return client.MsgSubscriber.CreateBulk(mss...).Save(ctx)
}

// DeleteMsgSubscriber is the resolver for the deleteMsgSubscriber field.
func (r *mutationResolver) DeleteMsgSubscriber(ctx context.Context, ids []int) (bool, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	_, err = ent.FromContext(ctx).MsgSubscriber.Delete().Where(msgsubscriber.IDIn(ids...), msgsubscriber.TenantID(tid)).Exec(ctx)
	return err == nil, err
}

// Matchers is the resolver for the matchers field.
func (r *routeInputResolver) Matchers(ctx context.Context, obj *profile.Route, data []*label.Matcher) error {
	obj.Matchers = data
	return nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// RouteInput returns generated.RouteInputResolver implementation.
func (r *Resolver) RouteInput() generated.RouteInputResolver { return &routeInputResolver{r} }

type mutationResolver struct{ *Resolver }
type routeInputResolver struct{ *Resolver }
