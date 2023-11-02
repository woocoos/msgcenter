package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	gen "github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/hook"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/predicate"
	"github.com/woocoos/msgcenter/pkg/profile"
	"regexp"
	"strings"
)

// MsgEvent 消息事件,消息事件相当于消息路由的入口配置.
type MsgEvent struct {
	ent.Schema
}

func (MsgEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_event"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (MsgEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the MsgEvent.
func (MsgEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Int("msg_type_id").Comment("消息类型ID").SchemaType(schemax.IntID{}.SchemaType()),
		field.String("name").NotEmpty().MaxLen(45).Comment("消息事件名称,应用内唯一").
			Match(regexp.MustCompile("[a-zA-Z0-9_]+$")),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusInactive.String()).
			Optional().Comment("状态").Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("comments").Optional().Comment("备注").Annotations(
			entgql.Skip(entgql.SkipWhereInput)),
		field.JSON("route", &profile.Route{}).Optional().Comment("消息路由配置"),
		field.String("modes").Comment("根据route配置对应的以,分隔的mode列表"),
	}
}

// Edges of the MsgEvent.
func (MsgEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("msg_type", MsgType.Type).Ref("events").Required().Unique().Field("msg_type_id").
			Comment("消息类型"),
		edge.To("customer_template", MsgTemplate.Type).Comment("自定义的消息模板").Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
	}
}

func (MsgEvent) Hooks() []ent.Hook {
	return []ent.Hook{
		nameHook(),
		modesHook(),
		routeHook(),
	}
}

func nameHook() ent.Hook {
	return hook.If(func(next ent.Mutator) ent.Mutator {
		return hook.MsgEventFunc(func(ctx context.Context, m *gen.MsgEventMutation) (gen.Value, error) {
			name, ok := m.Name()
			if !ok {
				return next.Mutate(ctx, m)
			}
			appID := 0
			if m.Op() == gen.OpCreate {
				mtid, _ := m.MsgTypeID()
				aid, err := m.Client().MsgType.Query().Where(msgtype.ID(mtid)).Select(msgtype.FieldAppID).Int(ctx)
				if err != nil {
					return nil, err
				}
				appID = aid
			} else if m.Op() == gen.OpUpdate || m.Op() == gen.OpUpdateOne {
				mtid, _ := m.MsgTypeID()
				if mtid == 0 {
					meid, _ := m.ID()
					me, err := m.Client().MsgEvent.Query().Where(msgevent.ID(meid)).WithMsgType().Only(ctx)
					if err != nil {
						return nil, err
					}
					appID = me.Edges.MsgType.AppID
				} else {
					aid, err := m.Client().MsgType.Query().Where(msgtype.ID(mtid)).Select(msgtype.FieldAppID).Int(ctx)
					if err != nil {
						return nil, err
					}
					appID = aid
				}
			}
			var where []predicate.MsgEvent
			where = append(where, msgevent.Name(name))
			if appID != 0 {
				where = append(where, msgevent.HasMsgTypeWith(msgtype.AppID(appID)))
			}
			has, err := m.Client().MsgEvent.Query().Where(where...).Exist(ctx)
			if err != nil {
				return nil, err
			}
			if has {
				return nil, fmt.Errorf("the event name must be unique in the app")
			}
			return next.Mutate(ctx, m)
		})
	}, hook.HasFields("name"))
}

// 站内信是一种特殊的webhoook.
func modesHook() ent.Hook {
	return hook.If(func(next ent.Mutator) ent.Mutator {
		return hook.MsgEventFunc(func(ctx context.Context, m *gen.MsgEventMutation) (gen.Value, error) {
			modes, ok := m.Modes()
			if !ok {
				return next.Mutate(ctx, m)
			}
			rts := strings.Split(modes, ",")
			rtvs := profile.ReceiverType("").Values()
		check:
			for _, rt := range rts {
				for _, rtv := range rtvs {
					if rt == rtv {
						break check
					}
				}
				return nil, fmt.Errorf("invalid modes %s", rt)
			}
			return next.Mutate(ctx, m)
		})
	}, hook.HasFields("modes"))
}

// routeHook 检查route添加的receiver是否符合modes
func routeHook() ent.Hook {
	return hook.If(
		func(next ent.Mutator) ent.Mutator {
			return hook.MsgEventFunc(func(ctx context.Context, m *gen.MsgEventMutation) (ent.Value, error) {
				modes := ""
				if m.Op() == gen.OpCreate {
					ms, ok := m.Modes()
					if !ok {
						return next.Mutate(ctx, m)
					}
					modes = ms
				} else if m.Op() == gen.OpUpdateOne || m.Op() == gen.OpUpdate {
					ms, _ := m.Modes()
					if ms != "" {
						modes = ms
					} else {
						meid, _ := m.ID()
						me, err := m.Client().MsgEvent.Query().Where(msgevent.ID(meid)).Only(ctx)
						if err != nil {
							return nil, err
						}
						modes = me.Modes
					}
				}
				if modes == "" {
					return next.Mutate(ctx, m)
				}
				route, ok := m.Route()
				if !ok {
					return next.Mutate(ctx, m)
				}
				rts := strings.Split(modes, ",")
				err := checkReceiverName([]*profile.Route{route}, rts)
				if err != nil {
					return nil, err
				}
				return next.Mutate(ctx, m)
			})
		}, hook.HasFields("route"))
}

func checkReceiverName(routes []*profile.Route, receiverTypes []string) error {
	for _, r := range routes {
		found := false
		for _, rt := range receiverTypes {
			if strings.HasPrefix(r.Receiver, rt) {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("invalid receiver %s", r.Receiver)
		}
		if r.Routes != nil {
			if err := checkReceiverName(r.Routes, receiverTypes); err != nil {
				return err
			}
		}
	}
	return nil
}
