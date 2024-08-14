package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FileSource holds the schema definition for the FileSource entity.
type FileSource struct {
	ent.Schema
}

// Annotations of the FileSource.
func (FileSource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "file_source"},
		entgql.Skip(entgql.SkipAll),
	}
}

// Fields of the FileSource.
func (FileSource) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Enum("kind").Values(
			"local", "minio", "aliOSS", "awsS3",
		).Comment("文件来源"),
		field.String("comments").Optional().Comment("备注").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.String("endpoint").MaxLen(255).Comment("对外服务的访问域名"),
		field.Bool("endpoint_immutable").Default(false).Comment("是否禁止修改endpoint，如果是自定义域名设为true"),
		field.String("sts_endpoint").MaxLen(255).Comment("sts服务的访问域名"),
		field.String("region").MaxLen(100).Comment("地域，数据存储的物理位置"),
		field.String("bucket").MaxLen(255).Comment("文件存储空间"),
		field.String("bucket_url").MaxLen(255).Comment("文件存储空间地址，用于匹配url"),
	}
}

// Hooks of the FileSource.
func (FileSource) Hooks() []ent.Hook {
	return []ent.Hook{
		readonlyHook(),
	}
}

func (FileSource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("identities", FileIdentity.Type).Comment("来源凭证").Annotations(entgql.Skip(entgql.SkipType)),
	}
}

type FileIdentity struct {
	ent.Schema
}

func (FileIdentity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "file_identity"},
		entgql.Skip(entgql.SkipAll),
	}
}

func (FileIdentity) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Comment("ID"),
		field.Int("tenant_id").Comment("组织ID").Annotations(entgql.Type("ID")),
		field.String("access_key_id").MaxLen(255).Comment("accesskey id"),
		field.String("access_key_secret").MaxLen(255).Comment("accesskey secret"),
		field.Int("file_source_id").Comment("文件来源ID"),
		field.String("role_arn").MaxLen(255).Comment("角色的资源名称(ARN)，用于STS"),
		field.Text("policy").Optional().Comment("指定返回的STS令牌的权限的策略"),
		field.Int("duration_seconds").Optional().Default(3600).Comment("STS令牌的有效期，默认3600s"),
		field.Bool("is_default").Default(false).Comment("租户默认的凭证").Annotations(entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipMutationCreateInput)),
		field.String("comments").Optional().Comment("备注").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}
func (FileIdentity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("source", FileSource.Type).Ref("identities").Required().Unique().Field("file_source_id"),
	}
}

func (FileIdentity) Hooks() []ent.Hook {
	return []ent.Hook{
		readonlyHook(),
	}
}
