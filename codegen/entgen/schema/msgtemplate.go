package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/fieldx"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/msgcenter/pkg/profile"
)

// MsgTemplate 消息模板用于提供给用户自行定义模板.
type MsgTemplate struct {
	ent.Schema
}

func (MsgTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_template"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		schemax.TenantField("tenant_id"),
	}
}

func (MsgTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the MsgTemplate.
func (MsgTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Int("msg_type_id").Comment("应用消息类型ID").SchemaType(schemax.IntID{}.SchemaType()),
		field.Int("msg_event_id").Comment("消息事件ID").SchemaType(schemax.IntID{}.SchemaType()),
		field.Int("tenant_id").Comment("组织ID").Annotations(entgql.Type("ID")),
		field.String("name").MaxLen(45).Comment("消息模板名称"),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusInactive.String()).
			Optional().Comment("状态").Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Enum("receiver_type").Comment("消息模式:站内信,app推送,邮件,短信,微信等").
			GoType(profile.ReceiverType("")),
		field.Enum("format").Values("txt", "html").Comment("消息类型:文本,网页,需要结合mod确定支持的格式"),
		field.String("subject").Optional().Comment("标题"),
		field.String("from").Optional().Comment("发件人"),
		field.String("to").Optional().Comment("收件人"),
		field.String("cc").Optional().Comment("抄送"),
		field.String("bcc").Optional().Comment("密送"),
		field.Text("body").Optional().Comment("消息体").Annotations(
			entgql.Skip(entgql.SkipWhereInput)),
		fieldx.File("tpl").Optional().Comment("模板地址").Annotations(
			entgql.Skip(entgql.SkipWhereInput)),
		field.Strings("attachments").Optional().Comment("附件地址").Annotations(
			entgql.Skip(entgql.SkipWhereInput)),
		field.String("comments").Optional().Comment("备注").Annotations(
			entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the MsgTemplate.
func (MsgTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", MsgEvent.Type).Ref("customer_template").Required().Unique().
			Field("msg_event_id"),
	}
}
