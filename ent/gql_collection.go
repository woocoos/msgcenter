// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/entco/pkg/pagination"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/user"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (mc *MsgChannelQuery) CollectFields(ctx context.Context, satisfies ...string) (*MsgChannelQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return mc, nil
	}
	if err := mc.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return mc, nil
}

func (mc *MsgChannelQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(msgchannel.Columns))
		selectedFields = []string{msgchannel.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "createdBy":
			if _, ok := fieldSeen[msgchannel.FieldCreatedBy]; !ok {
				selectedFields = append(selectedFields, msgchannel.FieldCreatedBy)
				fieldSeen[msgchannel.FieldCreatedBy] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[msgchannel.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, msgchannel.FieldCreatedAt)
				fieldSeen[msgchannel.FieldCreatedAt] = struct{}{}
			}
		case "updatedBy":
			if _, ok := fieldSeen[msgchannel.FieldUpdatedBy]; !ok {
				selectedFields = append(selectedFields, msgchannel.FieldUpdatedBy)
				fieldSeen[msgchannel.FieldUpdatedBy] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[msgchannel.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, msgchannel.FieldUpdatedAt)
				fieldSeen[msgchannel.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[msgchannel.FieldName]; !ok {
				selectedFields = append(selectedFields, msgchannel.FieldName)
				fieldSeen[msgchannel.FieldName] = struct{}{}
			}
		case "tenantID":
			if _, ok := fieldSeen[msgchannel.FieldTenantID]; !ok {
				selectedFields = append(selectedFields, msgchannel.FieldTenantID)
				fieldSeen[msgchannel.FieldTenantID] = struct{}{}
			}
		case "receiverType":
			if _, ok := fieldSeen[msgchannel.FieldReceiverType]; !ok {
				selectedFields = append(selectedFields, msgchannel.FieldReceiverType)
				fieldSeen[msgchannel.FieldReceiverType] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[msgchannel.FieldStatus]; !ok {
				selectedFields = append(selectedFields, msgchannel.FieldStatus)
				fieldSeen[msgchannel.FieldStatus] = struct{}{}
			}
		case "receiver":
			if _, ok := fieldSeen[msgchannel.FieldReceiver]; !ok {
				selectedFields = append(selectedFields, msgchannel.FieldReceiver)
				fieldSeen[msgchannel.FieldReceiver] = struct{}{}
			}
		case "comments":
			if _, ok := fieldSeen[msgchannel.FieldComments]; !ok {
				selectedFields = append(selectedFields, msgchannel.FieldComments)
				fieldSeen[msgchannel.FieldComments] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		mc.Select(selectedFields...)
	}
	return nil
}

type msgchannelPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MsgChannelPaginateOption
}

func newMsgChannelPaginateArgs(rv map[string]any) *msgchannelPaginateArgs {
	args := &msgchannelPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &MsgChannelOrder{Field: &MsgChannelOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithMsgChannelOrder(order))
			}
		case *MsgChannelOrder:
			if v != nil {
				args.opts = append(args.opts, WithMsgChannelOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*MsgChannelWhereInput); ok {
		args.opts = append(args.opts, WithMsgChannelFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (me *MsgEventQuery) CollectFields(ctx context.Context, satisfies ...string) (*MsgEventQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return me, nil
	}
	if err := me.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return me, nil
}

func (me *MsgEventQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(msgevent.Columns))
		selectedFields = []string{msgevent.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "msgType":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MsgTypeClient{config: me.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, "MsgType")...); err != nil {
				return err
			}
			me.withMsgType = query
			if _, ok := fieldSeen[msgevent.FieldMsgTypeID]; !ok {
				selectedFields = append(selectedFields, msgevent.FieldMsgTypeID)
				fieldSeen[msgevent.FieldMsgTypeID] = struct{}{}
			}
		case "customerTemplate":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MsgTemplateClient{config: me.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, "MsgTemplate")...); err != nil {
				return err
			}
			me.WithNamedCustomerTemplate(alias, func(wq *MsgTemplateQuery) {
				*wq = *query
			})
		case "createdBy":
			if _, ok := fieldSeen[msgevent.FieldCreatedBy]; !ok {
				selectedFields = append(selectedFields, msgevent.FieldCreatedBy)
				fieldSeen[msgevent.FieldCreatedBy] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[msgevent.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, msgevent.FieldCreatedAt)
				fieldSeen[msgevent.FieldCreatedAt] = struct{}{}
			}
		case "updatedBy":
			if _, ok := fieldSeen[msgevent.FieldUpdatedBy]; !ok {
				selectedFields = append(selectedFields, msgevent.FieldUpdatedBy)
				fieldSeen[msgevent.FieldUpdatedBy] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[msgevent.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, msgevent.FieldUpdatedAt)
				fieldSeen[msgevent.FieldUpdatedAt] = struct{}{}
			}
		case "msgTypeID":
			if _, ok := fieldSeen[msgevent.FieldMsgTypeID]; !ok {
				selectedFields = append(selectedFields, msgevent.FieldMsgTypeID)
				fieldSeen[msgevent.FieldMsgTypeID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[msgevent.FieldName]; !ok {
				selectedFields = append(selectedFields, msgevent.FieldName)
				fieldSeen[msgevent.FieldName] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[msgevent.FieldStatus]; !ok {
				selectedFields = append(selectedFields, msgevent.FieldStatus)
				fieldSeen[msgevent.FieldStatus] = struct{}{}
			}
		case "comments":
			if _, ok := fieldSeen[msgevent.FieldComments]; !ok {
				selectedFields = append(selectedFields, msgevent.FieldComments)
				fieldSeen[msgevent.FieldComments] = struct{}{}
			}
		case "route":
			if _, ok := fieldSeen[msgevent.FieldRoute]; !ok {
				selectedFields = append(selectedFields, msgevent.FieldRoute)
				fieldSeen[msgevent.FieldRoute] = struct{}{}
			}
		case "modes":
			if _, ok := fieldSeen[msgevent.FieldModes]; !ok {
				selectedFields = append(selectedFields, msgevent.FieldModes)
				fieldSeen[msgevent.FieldModes] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		me.Select(selectedFields...)
	}
	return nil
}

type msgeventPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MsgEventPaginateOption
}

func newMsgEventPaginateArgs(rv map[string]any) *msgeventPaginateArgs {
	args := &msgeventPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &MsgEventOrder{Field: &MsgEventOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithMsgEventOrder(order))
			}
		case *MsgEventOrder:
			if v != nil {
				args.opts = append(args.opts, WithMsgEventOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*MsgEventWhereInput); ok {
		args.opts = append(args.opts, WithMsgEventFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ms *MsgSubscriberQuery) CollectFields(ctx context.Context, satisfies ...string) (*MsgSubscriberQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ms, nil
	}
	if err := ms.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ms, nil
}

func (ms *MsgSubscriberQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(msgsubscriber.Columns))
		selectedFields = []string{msgsubscriber.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "msgType":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MsgTypeClient{config: ms.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, "MsgType")...); err != nil {
				return err
			}
			ms.withMsgType = query
			if _, ok := fieldSeen[msgsubscriber.FieldMsgTypeID]; !ok {
				selectedFields = append(selectedFields, msgsubscriber.FieldMsgTypeID)
				fieldSeen[msgsubscriber.FieldMsgTypeID] = struct{}{}
			}
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: ms.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, "User")...); err != nil {
				return err
			}
			ms.withUser = query
			if _, ok := fieldSeen[msgsubscriber.FieldUserID]; !ok {
				selectedFields = append(selectedFields, msgsubscriber.FieldUserID)
				fieldSeen[msgsubscriber.FieldUserID] = struct{}{}
			}
		case "createdBy":
			if _, ok := fieldSeen[msgsubscriber.FieldCreatedBy]; !ok {
				selectedFields = append(selectedFields, msgsubscriber.FieldCreatedBy)
				fieldSeen[msgsubscriber.FieldCreatedBy] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[msgsubscriber.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, msgsubscriber.FieldCreatedAt)
				fieldSeen[msgsubscriber.FieldCreatedAt] = struct{}{}
			}
		case "updatedBy":
			if _, ok := fieldSeen[msgsubscriber.FieldUpdatedBy]; !ok {
				selectedFields = append(selectedFields, msgsubscriber.FieldUpdatedBy)
				fieldSeen[msgsubscriber.FieldUpdatedBy] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[msgsubscriber.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, msgsubscriber.FieldUpdatedAt)
				fieldSeen[msgsubscriber.FieldUpdatedAt] = struct{}{}
			}
		case "msgTypeID":
			if _, ok := fieldSeen[msgsubscriber.FieldMsgTypeID]; !ok {
				selectedFields = append(selectedFields, msgsubscriber.FieldMsgTypeID)
				fieldSeen[msgsubscriber.FieldMsgTypeID] = struct{}{}
			}
		case "tenantID":
			if _, ok := fieldSeen[msgsubscriber.FieldTenantID]; !ok {
				selectedFields = append(selectedFields, msgsubscriber.FieldTenantID)
				fieldSeen[msgsubscriber.FieldTenantID] = struct{}{}
			}
		case "userID":
			if _, ok := fieldSeen[msgsubscriber.FieldUserID]; !ok {
				selectedFields = append(selectedFields, msgsubscriber.FieldUserID)
				fieldSeen[msgsubscriber.FieldUserID] = struct{}{}
			}
		case "orgRoleID":
			if _, ok := fieldSeen[msgsubscriber.FieldOrgRoleID]; !ok {
				selectedFields = append(selectedFields, msgsubscriber.FieldOrgRoleID)
				fieldSeen[msgsubscriber.FieldOrgRoleID] = struct{}{}
			}
		case "exclude":
			if _, ok := fieldSeen[msgsubscriber.FieldExclude]; !ok {
				selectedFields = append(selectedFields, msgsubscriber.FieldExclude)
				fieldSeen[msgsubscriber.FieldExclude] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ms.Select(selectedFields...)
	}
	return nil
}

type msgsubscriberPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MsgSubscriberPaginateOption
}

func newMsgSubscriberPaginateArgs(rv map[string]any) *msgsubscriberPaginateArgs {
	args := &msgsubscriberPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &MsgSubscriberOrder{Field: &MsgSubscriberOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithMsgSubscriberOrder(order))
			}
		case *MsgSubscriberOrder:
			if v != nil {
				args.opts = append(args.opts, WithMsgSubscriberOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*MsgSubscriberWhereInput); ok {
		args.opts = append(args.opts, WithMsgSubscriberFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (mt *MsgTemplateQuery) CollectFields(ctx context.Context, satisfies ...string) (*MsgTemplateQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return mt, nil
	}
	if err := mt.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return mt, nil
}

func (mt *MsgTemplateQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(msgtemplate.Columns))
		selectedFields = []string{msgtemplate.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "event":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MsgEventClient{config: mt.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, "MsgEvent")...); err != nil {
				return err
			}
			mt.withEvent = query
			if _, ok := fieldSeen[msgtemplate.FieldMsgEventID]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldMsgEventID)
				fieldSeen[msgtemplate.FieldMsgEventID] = struct{}{}
			}
		case "createdBy":
			if _, ok := fieldSeen[msgtemplate.FieldCreatedBy]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldCreatedBy)
				fieldSeen[msgtemplate.FieldCreatedBy] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[msgtemplate.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldCreatedAt)
				fieldSeen[msgtemplate.FieldCreatedAt] = struct{}{}
			}
		case "updatedBy":
			if _, ok := fieldSeen[msgtemplate.FieldUpdatedBy]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldUpdatedBy)
				fieldSeen[msgtemplate.FieldUpdatedBy] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[msgtemplate.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldUpdatedAt)
				fieldSeen[msgtemplate.FieldUpdatedAt] = struct{}{}
			}
		case "msgTypeID":
			if _, ok := fieldSeen[msgtemplate.FieldMsgTypeID]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldMsgTypeID)
				fieldSeen[msgtemplate.FieldMsgTypeID] = struct{}{}
			}
		case "msgEventID":
			if _, ok := fieldSeen[msgtemplate.FieldMsgEventID]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldMsgEventID)
				fieldSeen[msgtemplate.FieldMsgEventID] = struct{}{}
			}
		case "tenantID":
			if _, ok := fieldSeen[msgtemplate.FieldTenantID]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldTenantID)
				fieldSeen[msgtemplate.FieldTenantID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[msgtemplate.FieldName]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldName)
				fieldSeen[msgtemplate.FieldName] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[msgtemplate.FieldStatus]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldStatus)
				fieldSeen[msgtemplate.FieldStatus] = struct{}{}
			}
		case "receiverType":
			if _, ok := fieldSeen[msgtemplate.FieldReceiverType]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldReceiverType)
				fieldSeen[msgtemplate.FieldReceiverType] = struct{}{}
			}
		case "format":
			if _, ok := fieldSeen[msgtemplate.FieldFormat]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldFormat)
				fieldSeen[msgtemplate.FieldFormat] = struct{}{}
			}
		case "subject":
			if _, ok := fieldSeen[msgtemplate.FieldSubject]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldSubject)
				fieldSeen[msgtemplate.FieldSubject] = struct{}{}
			}
		case "from":
			if _, ok := fieldSeen[msgtemplate.FieldFrom]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldFrom)
				fieldSeen[msgtemplate.FieldFrom] = struct{}{}
			}
		case "to":
			if _, ok := fieldSeen[msgtemplate.FieldTo]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldTo)
				fieldSeen[msgtemplate.FieldTo] = struct{}{}
			}
		case "cc":
			if _, ok := fieldSeen[msgtemplate.FieldCc]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldCc)
				fieldSeen[msgtemplate.FieldCc] = struct{}{}
			}
		case "bcc":
			if _, ok := fieldSeen[msgtemplate.FieldBcc]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldBcc)
				fieldSeen[msgtemplate.FieldBcc] = struct{}{}
			}
		case "body":
			if _, ok := fieldSeen[msgtemplate.FieldBody]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldBody)
				fieldSeen[msgtemplate.FieldBody] = struct{}{}
			}
		case "tpl":
			if _, ok := fieldSeen[msgtemplate.FieldTpl]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldTpl)
				fieldSeen[msgtemplate.FieldTpl] = struct{}{}
			}
		case "attachments":
			if _, ok := fieldSeen[msgtemplate.FieldAttachments]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldAttachments)
				fieldSeen[msgtemplate.FieldAttachments] = struct{}{}
			}
		case "comments":
			if _, ok := fieldSeen[msgtemplate.FieldComments]; !ok {
				selectedFields = append(selectedFields, msgtemplate.FieldComments)
				fieldSeen[msgtemplate.FieldComments] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		mt.Select(selectedFields...)
	}
	return nil
}

type msgtemplatePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MsgTemplatePaginateOption
}

func newMsgTemplatePaginateArgs(rv map[string]any) *msgtemplatePaginateArgs {
	args := &msgtemplatePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &MsgTemplateOrder{Field: &MsgTemplateOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithMsgTemplateOrder(order))
			}
		case *MsgTemplateOrder:
			if v != nil {
				args.opts = append(args.opts, WithMsgTemplateOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*MsgTemplateWhereInput); ok {
		args.opts = append(args.opts, WithMsgTemplateFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (mt *MsgTypeQuery) CollectFields(ctx context.Context, satisfies ...string) (*MsgTypeQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return mt, nil
	}
	if err := mt.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return mt, nil
}

func (mt *MsgTypeQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(msgtype.Columns))
		selectedFields = []string{msgtype.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "events":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MsgEventClient{config: mt.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, "MsgEvent")...); err != nil {
				return err
			}
			mt.WithNamedEvents(alias, func(wq *MsgEventQuery) {
				*wq = *query
			})
		case "subscribers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MsgSubscriberClient{config: mt.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, "MsgSubscriber")...); err != nil {
				return err
			}
			mt.WithNamedSubscribers(alias, func(wq *MsgSubscriberQuery) {
				*wq = *query
			})
		case "createdBy":
			if _, ok := fieldSeen[msgtype.FieldCreatedBy]; !ok {
				selectedFields = append(selectedFields, msgtype.FieldCreatedBy)
				fieldSeen[msgtype.FieldCreatedBy] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[msgtype.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, msgtype.FieldCreatedAt)
				fieldSeen[msgtype.FieldCreatedAt] = struct{}{}
			}
		case "updatedBy":
			if _, ok := fieldSeen[msgtype.FieldUpdatedBy]; !ok {
				selectedFields = append(selectedFields, msgtype.FieldUpdatedBy)
				fieldSeen[msgtype.FieldUpdatedBy] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[msgtype.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, msgtype.FieldUpdatedAt)
				fieldSeen[msgtype.FieldUpdatedAt] = struct{}{}
			}
		case "appID":
			if _, ok := fieldSeen[msgtype.FieldAppID]; !ok {
				selectedFields = append(selectedFields, msgtype.FieldAppID)
				fieldSeen[msgtype.FieldAppID] = struct{}{}
			}
		case "category":
			if _, ok := fieldSeen[msgtype.FieldCategory]; !ok {
				selectedFields = append(selectedFields, msgtype.FieldCategory)
				fieldSeen[msgtype.FieldCategory] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[msgtype.FieldName]; !ok {
				selectedFields = append(selectedFields, msgtype.FieldName)
				fieldSeen[msgtype.FieldName] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[msgtype.FieldStatus]; !ok {
				selectedFields = append(selectedFields, msgtype.FieldStatus)
				fieldSeen[msgtype.FieldStatus] = struct{}{}
			}
		case "comments":
			if _, ok := fieldSeen[msgtype.FieldComments]; !ok {
				selectedFields = append(selectedFields, msgtype.FieldComments)
				fieldSeen[msgtype.FieldComments] = struct{}{}
			}
		case "canSubs":
			if _, ok := fieldSeen[msgtype.FieldCanSubs]; !ok {
				selectedFields = append(selectedFields, msgtype.FieldCanSubs)
				fieldSeen[msgtype.FieldCanSubs] = struct{}{}
			}
		case "canCustom":
			if _, ok := fieldSeen[msgtype.FieldCanCustom]; !ok {
				selectedFields = append(selectedFields, msgtype.FieldCanCustom)
				fieldSeen[msgtype.FieldCanCustom] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		mt.Select(selectedFields...)
	}
	return nil
}

type msgtypePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MsgTypePaginateOption
}

func newMsgTypePaginateArgs(rv map[string]any) *msgtypePaginateArgs {
	args := &msgtypePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &MsgTypeOrder{Field: &MsgTypeOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithMsgTypeOrder(order))
			}
		case *MsgTypeOrder:
			if v != nil {
				args.opts = append(args.opts, WithMsgTypeOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*MsgTypeWhereInput); ok {
		args.opts = append(args.opts, WithMsgTypeFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "principalName":
			if _, ok := fieldSeen[user.FieldPrincipalName]; !ok {
				selectedFields = append(selectedFields, user.FieldPrincipalName)
				fieldSeen[user.FieldPrincipalName] = struct{}{}
			}
		case "displayName":
			if _, ok := fieldSeen[user.FieldDisplayName]; !ok {
				selectedFields = append(selectedFields, user.FieldDisplayName)
				fieldSeen[user.FieldDisplayName] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[user.FieldEmail]; !ok {
				selectedFields = append(selectedFields, user.FieldEmail)
				fieldSeen[user.FieldEmail] = struct{}{}
			}
		case "mobile":
			if _, ok := fieldSeen[user.FieldMobile]; !ok {
				selectedFields = append(selectedFields, user.FieldMobile)
				fieldSeen[user.FieldMobile] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(ctx context.Context, partitionBy string, limit int, first, last *int, orderBy ...sql.Querier) func(s *sql.Selector) {
	offset := 0
	if sp, ok := pagination.SimplePaginationFromContext(ctx); ok {
		if first != nil {
			offset = (sp.PageIndex - sp.CurrentIndex - 1) * *first
		}
		if last != nil {
			offset = (sp.CurrentIndex - sp.PageIndex - 1) * *last
		}
	}
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		if offset != 0 {
			*s = *d.Select(s.UnqualifiedColumns()...).
				From(t).
				Where(sql.GT(t.C("row_number"), offset)).Limit(limit).
				Prefix(with)
		} else {
			*s = *d.Select(s.UnqualifiedColumns()...).
				From(t).
				Where(sql.LTE(t.C("row_number"), limit)).
				Prefix(with)
		}
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
