package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
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
	}
}

// Fields of the MsgSubscriber.
func (MsgSubscriber) Fields() []ent.Field {
	return []ent.Field{
		field.Int("msg_type_id").Comment("应用消息类型ID"),
		field.Int("tenant_id").Comment("组织ID"),
		field.Int("user_id").Comment("用户ID"),
		field.Bool("exclude").Optional().Default(false).Comment("是否排除"),
	}
}

// Edges of the MsgSubscriber.
func (MsgSubscriber) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("msg_type", MsgType.Type).Ref("subscribers").Required().Unique().Field("msg_type_id"),
	}
}
