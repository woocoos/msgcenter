package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/msgcenter/pkg/profile"
)

// MsgEvent 消息事件,消息事件相当于消息路由的入口配置.
type MsgEvent struct {
	ent.Schema
}

func (MsgEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_event"},
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (MsgEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the MsgEvent.
func (MsgEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Int("msg_type_id").Comment("消息类型ID"),
		field.String("name").NotEmpty().MaxLen(45).Comment("消息事件名称,应用内唯一"),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusInactive.String()).
			Optional().Comment("状态").Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
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
