package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
	gen "github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/intercept"
	"github.com/woocoos/msgcenter/version"
	"time"
)

// MsgInternalTo 站内信.
type MsgInternalTo struct {
	ent.Schema
}

func (MsgInternalTo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_internal_to"},
		entgql.QueryField().Description("站内信明细查询"),
		entgql.RelayConnection(),
		schemax.TenantField("tenant_id"),
	}
}

func (MsgInternalTo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.NewTenantMixin[intercept.Query, *gen.Client](version.AppCode, intercept.NewQuery),
		schemax.NotifyMixin{},
	}
}

// Fields of the MsgInternalTo.
func (MsgInternalTo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("msg_internal_id").Comment("站内信ID").Immutable().SchemaType(schemax.IntID{}.SchemaType()),
		field.Int("user_id").Comment("用户ID").Immutable(),
		field.Time("read_at").Optional().Comment("阅读时间"),
		field.Time("delete_at").Optional().Comment("删除时间"),
		field.Time("created_at").Immutable().Default(time.Now).Annotations(entgql.OrderField("createdAt")),
	}
}

// Edges of the MsgInternalTo.
func (MsgInternalTo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("msg_internal", MsgInternal.Type).Ref("msg_internal_to").Required().Unique().
			Immutable().Field("msg_internal_id"),
		edge.To("user", User.Type).Required().Immutable().Unique().Field("user_id"),
	}
}
