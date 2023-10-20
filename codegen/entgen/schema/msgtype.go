package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	gen "github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/hook"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/predicate"
)

// MsgType 应用的消息类型
type MsgType struct {
	ent.Schema
}

func (MsgType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "msg_type"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (MsgType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the MsgType.
func (MsgType) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Optional().Comment("应用ID").Annotations(entgql.Type("ID")),
		field.String("category").MaxLen(20).Comment("消息类型分类"),
		field.String("name").MaxLen(45).Comment("消息类型名称,应用内唯一"),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusActive.String()).
			Optional().Comment("状态"),
		field.String("comments").Optional().Comment("备注").Annotations(
			entgql.Skip(entgql.SkipWhereInput)),
		field.Bool("can_subs").Optional().Default(false).Comment("是否可订阅"),
		field.Bool("can_custom").Optional().Default(false).Comment("是否可定制"),
	}
}

// Edges of the MsgType.
func (MsgType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", MsgEvent.Type).Comment("消息事件").Annotations(
			entsql.Annotation{OnDelete: entsql.Cascade},
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		edge.To("subscribers", MsgSubscriber.Type).Comment("订阅者").Annotations(
			entsql.Annotation{OnDelete: entsql.Cascade},
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
	}
}

func (MsgType) Hooks() []ent.Hook {
	return []ent.Hook{
		typeNameHook(),
	}
}

func typeNameHook() ent.Hook {
	return hook.If(func(next ent.Mutator) ent.Mutator {
		return hook.MsgTypeFunc(func(ctx context.Context, m *gen.MsgTypeMutation) (gen.Value, error) {
			name, ok := m.Name()
			if !ok {
				return next.Mutate(ctx, m)
			}
			appID := 0
			if m.Op() == gen.OpCreate {
				appID, _ = m.AppID()
			} else if m.Op() == gen.OpUpdate || m.Op() == gen.OpUpdateOne {
				appID, _ = m.AppID()
				if appID == 0 {
					mtID, _ := m.ID()
					mt, err := m.Client().MsgType.Query().Where(msgtype.ID(mtID)).Select(msgtype.FieldAppID).Only(ctx)
					if err != nil {
						return nil, err
					}
					appID = mt.AppID
				}
			}
			var where []predicate.MsgType
			where = append(where, msgtype.Name(name))
			if appID != 0 {
				where = append(where, msgtype.AppID(appID))
			}
			has, err := m.Client().MsgType.Query().Where(where...).Exist(ctx)
			if err != nil {
				return nil, err
			}
			if has {
				return nil, fmt.Errorf("the msgtype name must be unique in the app")
			}
			return next.Mutate(ctx, m)
		})
	}, hook.HasFields("name"))
}
