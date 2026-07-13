package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"errors"
	"github.com/woocoos/msgcenter/ent/hook"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
		entgql.Skip(entgql.SkipEnumField, entgql.SkipOrderField, entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.String("principal_name").Unique().Comment("登陆名称").
			Annotations(entproto.Field(7)),
		field.String("display_name").Comment("显示名"),
	}
}

func readonlyHook() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			if mutation.Op().Is(ent.OpCreate) {
				// notice: for test,if run in production,should return error
				return next.Mutate(ctx, mutation)
			}
			return nil, errors.New("not implemented")
		})
	}, ent.OpCreate|ent.OpUpdate|ent.OpDelete)
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		readonlyHook(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("silences", MsgSilence.Type).Comment("静默"),
		edge.To("addresses", UserAddr.Type).Comment("用户联系信息"),
	}
}

// Org holds the schema definition for the Org entity.
type Org struct {
	ent.Schema
}

// Annotations of the Org.
func (Org) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org"},
		entgql.Skip(entgql.SkipEnumField, entgql.SkipOrderField, entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
	}
}

// Fields of the Org.
func (Org) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("组织ID"),
		field.Int("owner_id").Optional(),
		field.String("kind").Optional(),
		field.Int("parent_id").Optional(),
		field.Text("path").Optional(),
	}
}

// Hooks of the Org.
func (Org) Hooks() []ent.Hook {
	return []ent.Hook{
		readonlyHook(),
	}
}

func (Org) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("msg_alerts", MsgAlert.Type).Comment("消息列表"),
	}
}

type OrgRoleUser struct {
	ent.Schema
}

func (OrgRoleUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_role_user"},
		entgql.Skip(entgql.SkipAll),
	}
}

func (OrgRoleUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("org_role_id").Comment("组织角色ID"),
		field.Int("org_user_id").Comment("组织用户ID"),
		field.Int("org_id").Comment("组织ID"),
		field.Int("user_id").Comment("用户ID"),
	}
}

func (OrgRoleUser) Hooks() []ent.Hook {
	return []ent.Hook{
		readonlyHook(),
	}
}

// UserAddr holds the schema definition for the UserAddr entity.
type UserAddr struct {
	ent.Schema
}

func (UserAddr) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_addr"},
		entgql.Skip(entgql.SkipAll),
	}
}

// Fields of the UserAddr.
func (UserAddr) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("user_id").Optional().Immutable(),
		field.Enum("addr_type").NamedValues(
			"contact", "contact",
			"delivery", "delivery",
		).Comment("地址类型，contact：基本信息，delivery：收货地址").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Int("region_id").Optional().Nillable().Comment("地址地区：市"),
		field.String("addr").Optional().Comment("详细地址"),
		field.String("email").MaxLen(45).Optional().Comment("邮箱"),
		field.String("fax").MaxLen(45).Optional().Comment("传真"),
		field.String("zip_code").MaxLen(45).Optional().Comment("邮编"),
		field.String("tel").MaxLen(45).Optional().Comment("电话"),
		field.String("mobile").MaxLen(45).Optional().Comment("手机"),
		field.String("name").MaxLen(45).Optional().Comment("联系人名称"),
		field.Bool("is_default").Default(false).Comment("是否默认地址，类型为delivery时使用"),
	}
}

// Edges of the UserAddr.
func (UserAddr) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("addresses").Field("user_id").Unique().Immutable(),
	}
}

func (UserAddr) Hooks() []ent.Hook {
	return []ent.Hook{
		readonlyHook(),
	}
}
