package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/msgcenter/pkg/profile"
)

// MsgChannel 消息通道配置.
type MsgChannel struct {
	ent.Schema
}

func (MsgChannel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_channel"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		schemax.TenantField("tenant_id"),
	}
}

func (MsgChannel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the MsgChannel.
func (MsgChannel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(45).Comment("消息通道名称"),
		field.Int("tenant_id").Comment("组织ID"),
		field.Enum("receiver_type").Comment("支持的消息模式:站内信,app推送,邮件,短信,微信等").
			GoType(profile.ReceiverType("")),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusInactive.String()).
			Optional().Comment("状态").Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.JSON("receiver", &profile.Receiver{}).Optional().Comment("通道配置Json格式"),
		field.String("comments").Optional().Comment("备注"),
	}
}

// Edges of the MsgChannel.
func (MsgChannel) Edges() []ent.Edge {
	return nil
}
