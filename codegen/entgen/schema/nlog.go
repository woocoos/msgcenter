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
	"github.com/woocoos/msgcenter/pkg/profile"
	"time"
)

// Nlog 提醒的发送记录,记录着由哪个接收者,发送通道类型,发送执行通道.
//
// 由于一个Receivers可能会有多个通道,而具体的通道并不记录名称,因此记录的是该发送通道的索引位置.
type Nlog struct {
	ent.Schema
}

func (Nlog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_nlog"},
		entgql.RelayConnection(),
		schemax.TenantField("tenant_id"),
	}
}

func (Nlog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.NewTenantMixin[intercept.Query, *gen.Client](intercept.NewQuery),
	}
}

// Fields of the Nlog.
func (Nlog) Fields() []ent.Field {
	return []ent.Field{
		field.String("group_key").Comment("分组键"),
		field.String("receiver").Comment("接收组名称"),
		field.Enum("receiver_type").Comment("支持的消息模式:站内信,app推送,邮件,短信,微信等").
			GoType(profile.ReceiverType("")),
		field.Int("idx").Comment("通道的索引位置"),
		field.Time("send_at").Comment("发送时间"),
		field.Time("created_at").Immutable().Default(time.Now).Immutable().
			Annotations(entgql.OrderField("createdAt"), entgql.Skip(entgql.SkipMutationCreateInput)),
		field.Time("updated_at").Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("expires_at").Comment("驱逐时间"),
	}
}

// Edges of the Nlog.
func (Nlog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("alerts", MsgAlert.Type).Through("nlog_alert", NlogAlert.Type),
	}
}
