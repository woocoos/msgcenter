package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	"time"
)

// NlogAlert holds the schema definition for the NlogAlert entity.
type NlogAlert struct {
	ent.Schema
}

func (NlogAlert) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_nlog_alert"},
	}
}

// Fields of the NlogAlert.
func (NlogAlert) Fields() []ent.Field {
	return []ent.Field{
		field.Int("nlog_id").Comment("nlog id").SchemaType(schemax.IntID{}.SchemaType()),
		field.Int("alert_id").Comment("alert id").SchemaType(schemax.IntID{}.SchemaType()),
		field.Time("created_at").Immutable().Default(time.Now).Immutable().
			Annotations(entgql.OrderField("createdAt"), entgql.Skip(entgql.SkipMutationCreateInput)),
	}
}

// Edges of the NlogAlert.
func (NlogAlert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nlog", Nlog.Type).Unique().Required().Field("nlog_id"),
		edge.To("alert", MsgAlert.Type).Unique().Required().Field("alert_id"),
	}
}
