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
)

// MsgType 应用的消息类型
type MsgType struct {
	ent.Schema
}

func (MsgType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_type"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (MsgType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the MsgType.
func (MsgType) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Optional().Comment("应用ID"),
		field.String("category").MaxLen(20).Comment("消息类型分类"),
		field.String("name").MaxLen(45).Comment("消息类型名称,应用内唯一"),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusActive.String()).
			Optional().Comment("状态"),
		field.String("comments").Optional().Comment("备注").Annotations(
			entgql.Skip(entgql.SkipWhereInput)),
		field.Bool("can_subs").Optional().Default(false).Comment("是否可订阅"),
		field.Bool("can_custom").Optional().Default(false).Comment("是否可定制"),
	}
}

// Edges of the MsgType.
func (MsgType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", MsgEvent.Type).Comment("消息事件").Annotations(
			entsql.Annotation{OnDelete: entsql.Cascade},
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		edge.To("subscribers", MsgSubscriber.Type).Comment("订阅者").Annotations(
			entsql.Annotation{OnDelete: entsql.Cascade},
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
	}
}
