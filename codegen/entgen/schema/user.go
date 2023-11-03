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
		field.String("email").MaxLen(45).Optional().Comment("邮箱"),
		field.String("mobile").MaxLen(45).Optional().Comment("手机"),
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
		edge.To("silences", Silence.Type).Comment("静默"),
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
