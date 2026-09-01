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
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/service/silence"
	"github.com/woocoos/msgcenter/version"
)

// MsgSilence holds the schema definition for the Silence entity.
type MsgSilence struct {
	ent.Schema
}

func (MsgSilence) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_silence"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		schemax.TenantField("tenant_id"),
	}
}

func (MsgSilence) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
		schemax.NewTenantMixin[intercept.Query, *gen.Client](version.AppCode, intercept.NewQuery),
		schemax.NotifyMixin{},
	}
}

// Fields of the MsgSilence.
func (MsgSilence) Fields() []ent.Field {
	return []ent.Field{
		field.Int(schemax.FieldTenantID).Immutable().Comment("租户ID"),
		field.JSON("matchers", []*label.Matcher{}).Optional().Comment("应用ID").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Time("starts_at").Comment("开始时间"),
		field.Time("ends_at").Comment("结束时间"),
		field.String("comments").Optional().Comment("备注").Annotations(
			entgql.Skip(entgql.SkipWhereInput)),
		field.Enum("state").GoType(silence.SilenceState("")).
			Default(silence.SilenceStateActive.String()).Comment("状态"),
	}
}

// Edges of the MsgSilence.
func (MsgSilence) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("silences").Comment("创建人").Unique().
			Required().Immutable().Field("created_by").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
	}
}
