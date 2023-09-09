package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	gen "github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/intercept"
)

// MsgInternal 站内信.
type MsgInternal struct {
	ent.Schema
}

func (MsgInternal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_internal"},
		entgql.QueryField().Description("站内信查询"),
		entgql.RelayConnection(),
		schemax.TenantField("tenant_id"),
	}
}

func (MsgInternal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.NewTenantMixin[intercept.Query, *gen.Client](intercept.NewQuery),
		schemax.AuditMixin{},
	}
}

// Fields of the MsgInternal.
func (MsgInternal) Fields() []ent.Field {
	return []ent.Field{
		field.String("subject").Comment("标题"),
		field.String("body").Optional().Comment("消息体").Annotations(
			entgql.Skip(entgql.SkipWhereInput)),
		field.String("format").Comment("内容类型: html,txt"),
		field.String("redirect").Optional().Comment("消息跳转"),
	}
}

// Edges of the MsgInternal.
func (MsgInternal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("msg_internal_to", MsgInternalTo.Type),
	}
}
