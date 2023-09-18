package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/msgcenter/api/graphql/generated"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msginternalto"
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service"
	"github.com/woocoos/msgcenter/silence"
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
	client := ent.FromContext(ctx)
	ome, err := client.MsgEvent.Query().Where(msgevent.ID(id)).WithMsgType().Only(ctx)
	if err != nil {
		return nil, err
	}
	me, err := client.MsgEvent.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	mt, err := me.QueryMsgType().Only(ctx)
	if err != nil {
		return nil, err
	}
	// 修改route，同步更新NamedRoute
	if input.Route != nil && me.Status == typex.SimpleStatusActive {
		// 移除
		if err = r.Coordinator.RemoveNamedRoute([]string{profile.AppRouteName(strconv.Itoa(ome.Edges.MsgType.AppID), ome.Name)}); err != nil {
			return nil, err
		}
		// 添加
		route := *me.Route
		route.Name = profile.AppRouteName(strconv.Itoa(mt.AppID), me.Name)
		if err = r.Coordinator.AddNamedRoute([]*profile.Route{&route}); err != nil {
			return nil, err
		}
	}
	return me, nil
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
	client := ent.FromContext(ctx)
	omc, err := client.MsgChannel.Query().Where(msgchannel.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	mc, err := client.MsgChannel.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	// 修改receiver，同步修改TenantReceiver
	if input.Receiver != nil && mc.Status == typex.SimpleStatusActive {
		// 移除
		if err = r.Coordinator.RemoveTenantReceiver([]string{profile.TenantReceiverName(strconv.Itoa(omc.TenantID), omc.Name)}); err != nil {
			return nil, err
		}
		// 添加
		receiver := *mc.Receiver
		receiver.Name = profile.TenantReceiverName(strconv.Itoa(mc.TenantID), mc.Name)
		if err = r.Coordinator.AddTenantReceiver([]*profile.Receiver{&receiver}); err != nil {
			return nil, err
		}
	}
	return mc, nil
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
	if (input.Tpl != nil && input.TplFileID == nil) || (input.Tpl == nil && input.TplFileID != nil) {
		return nil, fmt.Errorf("tpl and tplFileID is required")
	}
	if ((input.Attachments != nil && len(input.Attachments) != 0) && input.AttachmentsFileIds == nil) ||
		(input.Attachments == nil && (input.AttachmentsFileIds != nil && len(input.AttachmentsFileIds) != 0)) ||
		(len(input.Attachments) != len(input.AttachmentsFileIds)) {
		return nil, fmt.Errorf("attachments and attachmentsFileIds is required and lengths must be same")
	}
	newFileIDs := make([]int, 0)
	// 验证模板路径
	if input.Tpl != nil {
		err := r.Coordinator.ValidateFilePath(ctx, *input.Tpl, service.TempRelativePathTplTemp)
		if err != nil {
			return nil, err
		}
		newFileIDs = append(newFileIDs, *input.TplFileID)
	}
	// 验证附件路径
	if input.Attachments != nil {
		for _, att := range input.Attachments {
			err := r.Coordinator.ValidateFilePath(ctx, att, service.TempRelativePathAttachment)
			if err != nil {
				return nil, err
			}
		}
		newFileIDs = append(newFileIDs, input.AttachmentsFileIds...)
	}
	temp, err := ent.FromContext(ctx).MsgTemplate.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	// 上报文件引用次数
	err = r.Coordinator.ReportFileRefCount(ctx, newFileIDs, nil)
	if err != nil {
		return nil, err
	}
	return temp, nil
}

// UpdateMsgTemplate is the resolver for the updateMsgTemplate field.
func (r *mutationResolver) UpdateMsgTemplate(ctx context.Context, id int, input ent.UpdateMsgTemplateInput) (*ent.MsgTemplate, error) {
	if (input.Tpl != nil && input.TplFileID == nil) || (input.Tpl == nil && input.TplFileID != nil) {
		return nil, fmt.Errorf("tpl and tplFileID must exist at the same time")
	}
	if ((input.Attachments != nil && len(input.Attachments) != 0) && input.AttachmentsFileIds == nil) ||
		(input.Attachments == nil && (input.AttachmentsFileIds != nil && len(input.AttachmentsFileIds) != 0)) {
		return nil, fmt.Errorf("attachments and attachmentsFileIds must exist at the same time")
	}
	otemp, err := ent.FromContext(ctx).MsgTemplate.Query().Where(msgtemplate.ID(id)).Select(msgtemplate.FieldAttachmentsFileIds, msgtemplate.FieldTplFileID).Only(ctx)
	if err != nil {
		return nil, err
	}
	newFileIDs := make([]int, 0)
	oldFileIDs := make([]int, 0)
	// 验证模板路径
	if input.Tpl != nil {
		err := r.Coordinator.ValidateFilePath(ctx, *input.Tpl, service.TempRelativePathTplTemp)
		if err != nil {
			return nil, err
		}
		newFileIDs = append(newFileIDs, *input.TplFileID)

		if otemp.TplFileID != nil {
			oldFileIDs = append(oldFileIDs, *otemp.TplFileID)
		}
	}
	// 验证附件路径
	if input.Attachments != nil && len(input.Attachments) > 0 {
		for _, att := range input.Attachments {
			err := r.Coordinator.ValidateFilePath(ctx, att, service.TempRelativePathAttachment)
			if err != nil {
				return nil, err
			}
		}
		newFileIDs = append(newFileIDs, input.AttachmentsFileIds...)

		if otemp.AttachmentsFileIds != nil {
			oldFileIDs = append(oldFileIDs, otemp.AttachmentsFileIds...)
		}
	}

	// 更新模板
	temp, err := ent.FromContext(ctx).MsgTemplate.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	// 模板文件更新，则删除旧模板文件
	if input.Tpl != nil && input.TplFileID != nil && temp.Status == typex.SimpleStatusActive {
		// 移除data下的旧模板
		err = r.Coordinator.RemoveTplDataFile(otemp.Tpl)
		if err != nil {
			return nil, err
		}
		// 启用新模板
		err = r.Coordinator.EnableTplDataFile(temp.Tpl)
		if err != nil {
			return nil, err
		}
	}
	// 上报文件引用次数
	err = r.Coordinator.ReportFileRefCount(ctx, newFileIDs, oldFileIDs)
	if err != nil {
		return nil, err
	}
	return temp, nil
}

// DeleteMsgTemplate is the resolver for the deleteMsgTemplate field.
func (r *mutationResolver) DeleteMsgTemplate(ctx context.Context, id int) (bool, error) {
	client := ent.FromContext(ctx)
	// 激活状态不可删
	if has, err := client.MsgTemplate.Query().Where(msgtemplate.ID(id), msgtemplate.StatusEQ(typex.SimpleStatusActive)).Exist(ctx); err != nil {
		return false, err
	} else if has {
		return false, fmt.Errorf("the active status cannot be deleted")
	}
	// 获取模板数据
	otemp, err := ent.FromContext(ctx).MsgTemplate.Query().Where(msgtemplate.ID(id)).Select(msgtemplate.FieldAttachmentsFileIds, msgtemplate.FieldTplFileID).Only(ctx)
	if err != nil {
		return false, err
	}
	// 删除模板
	err = client.MsgTemplate.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return false, err
	}
	// 更新文件引用次数
	oldFileIDs := make([]int, 0)
	if otemp.TplFileID != nil {
		oldFileIDs = append(oldFileIDs, *otemp.TplFileID)
	}
	if otemp.AttachmentsFileIds != nil {
		oldFileIDs = append(oldFileIDs, otemp.AttachmentsFileIds...)
	}
	err = r.Coordinator.ReportFileRefCount(ctx, nil, oldFileIDs)
	if err != nil {
		return false, err
	}
	return true, nil
}

// EnableMsgTemplate is the resolver for the enableMsgTemplate field.
func (r *mutationResolver) EnableMsgTemplate(ctx context.Context, id int) (*ent.MsgTemplate, error) {
	temp, err := ent.FromContext(ctx).MsgTemplate.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	temp, err = temp.Update().SetStatus(typex.SimpleStatusActive).Save(ctx)
	if err != nil {
		return nil, err
	}
	// 启用模板
	err = r.Coordinator.EnableTplDataFile(temp.Tpl)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return temp, nil
}

// DisableMsgTemplate is the resolver for the disableMsgTemplate field.
func (r *mutationResolver) DisableMsgTemplate(ctx context.Context, id int) (*ent.MsgTemplate, error) {
	temp, err := ent.FromContext(ctx).MsgTemplate.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// 移除data目录模板
	err = r.Coordinator.RemoveTplDataFile(temp.Tpl)
	if err != nil {
		return nil, err
	}
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

// CreateSilence is the resolver for the createSilence field.
func (r *mutationResolver) CreateSilence(ctx context.Context, input ent.CreateSilenceInput) (*ent.Silence, error) {
	sil, err := ent.FromContext(ctx).Silence.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.Silences.Set(&silence.Entry{
		ID:        sil.ID,
		UpdatedAt: sil.UpdatedAt,
		Matchers:  sil.Matchers,
		StartsAt:  sil.StartsAt,
		EndsAt:    sil.EndsAt,
		State:     sil.State,
	})
	return sil, err
}

// UpdateSilence is the resolver for the updateSilence field.
func (r *mutationResolver) UpdateSilence(ctx context.Context, id int, input ent.UpdateSilenceInput) (*ent.Silence, error) {
	client := ent.FromContext(ctx)
	sil, err := client.Silence.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	id, err = r.Silences.Set(&silence.Entry{
		ID:        sil.ID,
		UpdatedAt: sil.UpdatedAt,
		Matchers:  sil.Matchers,
		StartsAt:  sil.StartsAt,
		EndsAt:    sil.EndsAt,
		State:     sil.State,
	})
	if err != nil {
		return nil, err
	}
	if sil.ID == id {
		return sil, nil
	}
	mu := client.Silence.UpdateOne(sil).Mutation()
	mu.SetOp(ent.OpCreate)
	v, err := client.Mutate(ctx, mu)
	return v.(*ent.Silence), err
}

// DeleteSilence is the resolver for the deleteSilence field.
func (r *mutationResolver) DeleteSilence(ctx context.Context, id int) (bool, error) {
	err := ent.FromContext(ctx).Silence.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return false, err
	}
	err = r.Silences.Expire(id)
	return err == nil, err
}

// MarkMessageReadOrUnRead is the resolver for the markMessageReadOrUnRead field.
func (r *mutationResolver) MarkMessageReadOrUnRead(ctx context.Context, ids []int, read bool) (bool, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	update := ent.FromContext(ctx).MsgInternalTo.Update().Where(
		msginternalto.IDIn(ids...), msginternalto.UserID(uid), msginternalto.TenantID(tid),
	)
	if read {
		update.SetReadAt(time.Now())
	} else {
		update.ClearReadAt()
	}
	err = update.Exec(ctx)
	return err == nil, err
}

// MarkMessageDeleted is the resolver for the markMessageDeleted field.
func (r *mutationResolver) MarkMessageDeleted(ctx context.Context, ids []int) (bool, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	err = ent.FromContext(ctx).MsgInternalTo.Update().Where(
		msginternalto.IDIn(ids...), msginternalto.UserID(uid), msginternalto.TenantID(tid),
	).SetDeleteAt(time.Now()).Exec(ctx)
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
