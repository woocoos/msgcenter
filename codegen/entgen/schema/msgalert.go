package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/woocoos/entco/schemax"
	gen "github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/intercept"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"time"
)

// MsgAlert holds the schema definition for the MsgAlert entity.
// if MsgAlert has state value that means it has been notified.
type MsgAlert struct {
	ent.Schema
}

func (MsgAlert) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_alert"},
		entgql.RelayConnection(),
		schemax.TenantField("tenant_id"),
	}
}

func (MsgAlert) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.NewTenantMixin[intercept.Query, *gen.Client](intercept.NewQuery),
	}
}

// Fields of the MsgAlert.
func (MsgAlert) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("labels", &label.LabelSet{}).Optional().Comment("标签").
			Annotations(entgql.Skip(entgql.SkipWhereInput), entgql.Type("MapString")),
		field.JSON("annotations", &label.LabelSet{}).Optional().Comment("注解").
			Annotations(entgql.Skip(entgql.SkipWhereInput), entgql.Type("MapString")),
		field.Time("starts_at").Comment("开始时间"),
		field.Time("ends_at").Comment("结束时间"),
		field.String("url").Optional().Comment("generatorURL"),
		field.Bool("timeout").Default(false).Comment("状态"),
		field.String("fingerprint").Comment("指纹hash值"),
		field.Enum("state").GoType(alert.AlertStatus("")).Default(string(alert.AlertNone)).
			Comment("通知状态,firing: 触发通知,resolved: 已处理过"),
		field.Time("created_at").Immutable().Default(time.Now).Immutable().
			Annotations(entgql.OrderField("createdAt"), entgql.Skip(entgql.SkipMutationCreateInput)),
		field.Time("updated_at").Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Bool("deleted").Default(false).Comment("是否移除"),
	}
}

// Edges of the MsgAlert.
func (MsgAlert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nlog", Nlog.Type).Ref("alerts").Through("nlog_alerts", NlogAlert.Type).
			Annotations(entgql.RelayConnection()).Comment("消息日志"),
	}
}

// Indexes of the MsgAlert.
func (MsgAlert) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("fingerprint"),
	}
}
