package schema

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
	gen "github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/hook"
)

// MsgSubscriber holds the schema definition for the MsgSubscriber entity.
type MsgSubscriber struct {
	ent.Schema
}

func (MsgSubscriber) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_subscriber"},
		schemax.TenantField("tenant_id"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (MsgSubscriber) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the MsgSubscriber.
func (MsgSubscriber) Fields() []ent.Field {
	return []ent.Field{
		field.Int("msg_type_id").Optional().Comment("应用消息类型ID"),
		field.Int("msg_event_id").Optional().Comment("应用消息事件ID"),
		field.Int("tenant_id").Comment("组织ID").Annotations(entgql.Type("ID")),
		field.Int("user_id").Optional().Comment("用户ID"),
		field.Int("org_role_id").Optional().Comment("用户组ID").Annotations(entgql.Type("ID")),
		field.Bool("exclude").Optional().Default(false).Comment("是否排除"),
	}
}

// Edges of the MsgSubscriber.
func (MsgSubscriber) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("msg_type", MsgType.Type).Ref("subscribers").Unique().Field("msg_type_id"),
		edge.From("msg_event", MsgEvent.Type).Ref("subscribers").Unique().Field("msg_event_id"),
		edge.To("user", User.Type).Unique().Field("user_id"),
	}
}

func (MsgSubscriber) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.MsgSubscriberFunc(func(ctx context.Context, m *gen.MsgSubscriberMutation) (ent.Value, error) {
					// 限制用户和用户组只能存在一个
					if uid, ok := m.OrgRoleID(); ok && uid != 0 {
						m.ClearUserID()
					}

					// 验证 msg_type_id 和 msg_event_id 互斥且必选其一
					typeID, hasTypeID := m.MsgTypeID()
					eventID, hasEventID := m.MsgEventID()

					typeIDValid := hasTypeID && typeID != 0
					eventIDValid := hasEventID && eventID != 0

					if typeIDValid && eventIDValid {
						return nil, fmt.Errorf("msg_type_id and msg_event_id are mutually exclusive, only one can be set")
					}
					if !typeIDValid && !eventIDValid {
						return nil, fmt.Errorf("either msg_type_id or msg_event_id must be set")
					}

					return next.Mutate(ctx, m)
				})
			}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
