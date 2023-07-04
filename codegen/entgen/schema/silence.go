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
	"github.com/woocoos/msgcenter/pkg/label"
)

// Silence holds the schema definition for the Silence entity.
type Silence struct {
	ent.Schema
}

func (Silence) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_silence"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		schemax.TenantField("tenant_id"),
	}
}

func (Silence) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NewSoftDeleteMixin[intercept.Query, *gen.Client](intercept.NewQuery),
	}
}

// Fields of the Silence.
func (Silence) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("matchers", []label.Matcher{}).Optional().Comment("应用ID").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Time("starts_at").Comment("开始时间"),
		field.Time("ends_at").Comment("结束时间"),
		field.String("comments").Optional().Comment("备注").Annotations(
			entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the Silence.
func (Silence) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("silences").Comment("创建人").Unique().
			Required().Immutable().Field("created_by"),
	}
}
