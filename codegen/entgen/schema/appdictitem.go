package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
)

// AppDictItem holds the schema definition for the AppDictItem entity.
type AppDictItem struct {
	ent.Schema
}

// Annotations of the AppDictItem.
func (AppDictItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_dict_item"},
		entgql.Skip(entgql.SkipEnumField, entgql.SkipOrderField, entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
	}
}

// Fields of the User.
func (AppDictItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Optional().Immutable().Comment("租户ID,空为全局字典"),
		field.String("ref_code").Comment("关联代码,由app_code和dict_code组成"),
		field.String("code").MinLen(3).MaxLen(45).Immutable().Comment("字典值唯一编码,生效后不可修改."),
		field.String("name").MaxLen(255).Comment("名称"),
		field.Enum("status").GoType(typex.SimpleStatus("")).Optional().Comment("状态"),
	}
}

// Hooks of the AppDictItem.
func (AppDictItem) Hooks() []ent.Hook {
	return []ent.Hook{
		readonlyHook(),
	}
}
