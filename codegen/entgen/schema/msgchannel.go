package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	gen "github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/hook"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/pkg/profile"
	"strings"
)

// MsgChannel 消息通道配置.消息通道是特定类型的receiver,每一消息通道对应着一个receiver配置.
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
		schemax.NotifyMixin{},
	}
}

// Fields of the MsgChannel.
func (MsgChannel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(45).Comment("消息通道名称"),
		field.Int("tenant_id").Comment("组织ID").Annotations(entgql.Type("ID")),
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

func (MsgChannel) Hooks() []ent.Hook {
	return []ent.Hook{
		receiverHook(),
	}
}

// receiverHook 检查receiver name是否符合receiverType
func receiverHook() ent.Hook {
	return hook.If(
		func(next ent.Mutator) ent.Mutator {
			return hook.MsgChannelFunc(func(ctx context.Context, m *gen.MsgChannelMutation) (ent.Value, error) {
				rts := ""
				if m.Op() == gen.OpCreate {
					rt, ok := m.ReceiverType()
					if !ok {
						return next.Mutate(ctx, m)
					}
					rts = rt.String()
				} else if m.Op() == gen.OpUpdate || m.Op() == gen.OpUpdateOne {
					rt, _ := m.ReceiverType()
					if rt.String() != "" {
						rts = rt.String()
					} else {
						mcid, ok := m.ID()
						if !ok {
							return next.Mutate(ctx, m)
						}
						receiverType, err := m.Client().MsgChannel.Query().Where(msgchannel.ID(mcid)).Select(msgchannel.FieldReceiverType).String(ctx)
						if err != nil {
							return nil, err
						}
						rts = receiverType
					}
				}
				if rts == "" {
					return next.Mutate(ctx, m)
				}
				r, ok := m.Receiver()
				if !ok {
					return next.Mutate(ctx, m)
				}
				// 验证receiver name
				if !strings.HasPrefix(r.Name, rts) {
					return nil, fmt.Errorf("the receiver name must be receiverType or prefixed with receiverType")
				}
				// 检查receiver name是否存在
				return next.Mutate(ctx, m)
			})
		}, hook.HasFields("receiver"))
}
