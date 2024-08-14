// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FileIdentityColumns holds the columns for the "file_identity" table.
	FileIdentityColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tenant_id", Type: field.TypeInt},
		{Name: "access_key_id", Type: field.TypeString, Size: 255},
		{Name: "access_key_secret", Type: field.TypeString, Size: 255},
		{Name: "role_arn", Type: field.TypeString, Size: 255},
		{Name: "policy", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "duration_seconds", Type: field.TypeInt, Nullable: true, Default: 3600},
		{Name: "is_default", Type: field.TypeBool, Default: false},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "file_source_id", Type: field.TypeInt},
	}
	// FileIdentityTable holds the schema information for the "file_identity" table.
	FileIdentityTable = &schema.Table{
		Name:       "file_identity",
		Columns:    FileIdentityColumns,
		PrimaryKey: []*schema.Column{FileIdentityColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "file_identity_file_source_identities",
				Columns:    []*schema.Column{FileIdentityColumns[9]},
				RefColumns: []*schema.Column{FileSourceColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FileSourceColumns holds the columns for the "file_source" table.
	FileSourceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"local", "minio", "aliOSS", "awsS3"}},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "endpoint", Type: field.TypeString, Size: 255},
		{Name: "endpoint_immutable", Type: field.TypeBool, Default: false},
		{Name: "sts_endpoint", Type: field.TypeString, Size: 255},
		{Name: "region", Type: field.TypeString, Size: 100},
		{Name: "bucket", Type: field.TypeString, Size: 255},
		{Name: "bucket_url", Type: field.TypeString, Size: 255},
	}
	// FileSourceTable holds the schema information for the "file_source" table.
	FileSourceTable = &schema.Table{
		Name:       "file_source",
		Columns:    FileSourceColumns,
		PrimaryKey: []*schema.Column{FileSourceColumns[0]},
	}
	// MsgAlertColumns holds the columns for the "msg_alert" table.
	MsgAlertColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "tenant_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "labels", Type: field.TypeJSON, Nullable: true},
		{Name: "annotations", Type: field.TypeJSON, Nullable: true},
		{Name: "starts_at", Type: field.TypeTime},
		{Name: "ends_at", Type: field.TypeTime, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "timeout", Type: field.TypeBool, Default: false},
		{Name: "fingerprint", Type: field.TypeString},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"none", "firing", "resolved"}, Default: "none"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted", Type: field.TypeBool, Default: false},
	}
	// MsgAlertTable holds the schema information for the "msg_alert" table.
	MsgAlertTable = &schema.Table{
		Name:       "msg_alert",
		Columns:    MsgAlertColumns,
		PrimaryKey: []*schema.Column{MsgAlertColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "msgalert_fingerprint",
				Unique:  false,
				Columns: []*schema.Column{MsgAlertColumns[8]},
			},
		},
	}
	// MsgChannelColumns holds the columns for the "msg_channel" table.
	MsgChannelColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 45},
		{Name: "tenant_id", Type: field.TypeInt},
		{Name: "receiver_type", Type: field.TypeEnum, Enums: []string{"email", "message", "webhook"}},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing", "disabled"}, Default: "inactive"},
		{Name: "receiver", Type: field.TypeJSON, Nullable: true},
		{Name: "comments", Type: field.TypeString, Nullable: true},
	}
	// MsgChannelTable holds the schema information for the "msg_channel" table.
	MsgChannelTable = &schema.Table{
		Name:       "msg_channel",
		Columns:    MsgChannelColumns,
		PrimaryKey: []*schema.Column{MsgChannelColumns[0]},
	}
	// MsgEventColumns holds the columns for the "msg_event" table.
	MsgEventColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 45},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing", "disabled"}, Default: "inactive"},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "route", Type: field.TypeJSON, Nullable: true},
		{Name: "modes", Type: field.TypeString},
		{Name: "msg_type_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "int"}},
	}
	// MsgEventTable holds the schema information for the "msg_event" table.
	MsgEventTable = &schema.Table{
		Name:       "msg_event",
		Columns:    MsgEventColumns,
		PrimaryKey: []*schema.Column{MsgEventColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "msg_event_msg_type_events",
				Columns:    []*schema.Column{MsgEventColumns[10]},
				RefColumns: []*schema.Column{MsgTypeColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MsgInternalColumns holds the columns for the "msg_internal" table.
	MsgInternalColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "tenant_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "category", Type: field.TypeString, Size: 20},
		{Name: "subject", Type: field.TypeString},
		{Name: "body", Type: field.TypeString, Nullable: true},
		{Name: "format", Type: field.TypeString},
		{Name: "redirect", Type: field.TypeString, Nullable: true},
	}
	// MsgInternalTable holds the schema information for the "msg_internal" table.
	MsgInternalTable = &schema.Table{
		Name:       "msg_internal",
		Columns:    MsgInternalColumns,
		PrimaryKey: []*schema.Column{MsgInternalColumns[0]},
	}
	// MsgInternalToColumns holds the columns for the "msg_internal_to" table.
	MsgInternalToColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "tenant_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "read_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "msg_internal_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "user_id", Type: field.TypeInt},
	}
	// MsgInternalToTable holds the schema information for the "msg_internal_to" table.
	MsgInternalToTable = &schema.Table{
		Name:       "msg_internal_to",
		Columns:    MsgInternalToColumns,
		PrimaryKey: []*schema.Column{MsgInternalToColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "msg_internal_to_msg_internal_msg_internal_to",
				Columns:    []*schema.Column{MsgInternalToColumns[5]},
				RefColumns: []*schema.Column{MsgInternalColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "msg_internal_to_user_user",
				Columns:    []*schema.Column{MsgInternalToColumns[6]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MsgSubscriberColumns holds the columns for the "msg_subscriber" table.
	MsgSubscriberColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "tenant_id", Type: field.TypeInt},
		{Name: "org_role_id", Type: field.TypeInt, Nullable: true},
		{Name: "exclude", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
		{Name: "msg_type_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "int"}},
	}
	// MsgSubscriberTable holds the schema information for the "msg_subscriber" table.
	MsgSubscriberTable = &schema.Table{
		Name:       "msg_subscriber",
		Columns:    MsgSubscriberColumns,
		PrimaryKey: []*schema.Column{MsgSubscriberColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "msg_subscriber_user_user",
				Columns:    []*schema.Column{MsgSubscriberColumns[8]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "msg_subscriber_msg_type_subscribers",
				Columns:    []*schema.Column{MsgSubscriberColumns[9]},
				RefColumns: []*schema.Column{MsgTypeColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MsgTemplateColumns holds the columns for the "msg_template" table.
	MsgTemplateColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "msg_type_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "tenant_id", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Size: 45},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing", "disabled"}, Default: "inactive"},
		{Name: "receiver_type", Type: field.TypeEnum, Enums: []string{"email", "message", "webhook"}},
		{Name: "format", Type: field.TypeEnum, Enums: []string{"txt", "html"}},
		{Name: "subject", Type: field.TypeString, Nullable: true},
		{Name: "from", Type: field.TypeString, Nullable: true},
		{Name: "to", Type: field.TypeString, Nullable: true},
		{Name: "cc", Type: field.TypeString, Nullable: true},
		{Name: "bcc", Type: field.TypeString, Nullable: true},
		{Name: "body", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "tpl", Type: field.TypeString, Nullable: true},
		{Name: "attachments", Type: field.TypeJSON, Nullable: true},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "msg_event_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "int"}},
	}
	// MsgTemplateTable holds the schema information for the "msg_template" table.
	MsgTemplateTable = &schema.Table{
		Name:       "msg_template",
		Columns:    MsgTemplateColumns,
		PrimaryKey: []*schema.Column{MsgTemplateColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "msg_template_msg_event_customer_template",
				Columns:    []*schema.Column{MsgTemplateColumns[20]},
				RefColumns: []*schema.Column{MsgEventColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MsgTypeColumns holds the columns for the "msg_type" table.
	MsgTypeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "app_id", Type: field.TypeInt, Nullable: true},
		{Name: "category", Type: field.TypeString, Size: 20},
		{Name: "name", Type: field.TypeString, Size: 45},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing", "disabled"}, Default: "active"},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "can_subs", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "can_custom", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// MsgTypeTable holds the schema information for the "msg_type" table.
	MsgTypeTable = &schema.Table{
		Name:       "msg_type",
		Columns:    MsgTypeColumns,
		PrimaryKey: []*schema.Column{MsgTypeColumns[0]},
	}
	// MsgNlogColumns holds the columns for the "msg_nlog" table.
	MsgNlogColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "tenant_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "group_key", Type: field.TypeString},
		{Name: "receiver", Type: field.TypeString},
		{Name: "receiver_type", Type: field.TypeEnum, Enums: []string{"email", "message", "webhook"}},
		{Name: "idx", Type: field.TypeInt},
		{Name: "send_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "expires_at", Type: field.TypeTime},
	}
	// MsgNlogTable holds the schema information for the "msg_nlog" table.
	MsgNlogTable = &schema.Table{
		Name:       "msg_nlog",
		Columns:    MsgNlogColumns,
		PrimaryKey: []*schema.Column{MsgNlogColumns[0]},
	}
	// MsgNlogAlertColumns holds the columns for the "msg_nlog_alert" table.
	MsgNlogAlertColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "nlog_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "alert_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "int"}},
	}
	// MsgNlogAlertTable holds the schema information for the "msg_nlog_alert" table.
	MsgNlogAlertTable = &schema.Table{
		Name:       "msg_nlog_alert",
		Columns:    MsgNlogAlertColumns,
		PrimaryKey: []*schema.Column{MsgNlogAlertColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "msg_nlog_alert_msg_nlog_nlog",
				Columns:    []*schema.Column{MsgNlogAlertColumns[2]},
				RefColumns: []*schema.Column{MsgNlogColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "msg_nlog_alert_msg_alert_alert",
				Columns:    []*schema.Column{MsgNlogAlertColumns[3]},
				RefColumns: []*schema.Column{MsgAlertColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "nlogalert_nlog_id_alert_id",
				Unique:  true,
				Columns: []*schema.Column{MsgNlogAlertColumns[2], MsgNlogAlertColumns[3]},
			},
		},
	}
	// OrgRoleUserColumns holds the columns for the "org_role_user" table.
	OrgRoleUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "org_role_id", Type: field.TypeInt},
		{Name: "org_user_id", Type: field.TypeInt},
		{Name: "org_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// OrgRoleUserTable holds the schema information for the "org_role_user" table.
	OrgRoleUserTable = &schema.Table{
		Name:       "org_role_user",
		Columns:    OrgRoleUserColumns,
		PrimaryKey: []*schema.Column{OrgRoleUserColumns[0]},
	}
	// MsgSilenceColumns holds the columns for the "msg_silence" table.
	MsgSilenceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "tenant_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "matchers", Type: field.TypeJSON, Nullable: true},
		{Name: "starts_at", Type: field.TypeTime},
		{Name: "ends_at", Type: field.TypeTime},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"expired", "active", "pending"}, Default: "active"},
		{Name: "created_by", Type: field.TypeInt},
	}
	// MsgSilenceTable holds the schema information for the "msg_silence" table.
	MsgSilenceTable = &schema.Table{
		Name:       "msg_silence",
		Columns:    MsgSilenceColumns,
		PrimaryKey: []*schema.Column{MsgSilenceColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "msg_silence_user_silences",
				Columns:    []*schema.Column{MsgSilenceColumns[10]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "principal_name", Type: field.TypeString, Unique: true},
		{Name: "display_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "mobile", Type: field.TypeString, Nullable: true, Size: 45},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FileIdentityTable,
		FileSourceTable,
		MsgAlertTable,
		MsgChannelTable,
		MsgEventTable,
		MsgInternalTable,
		MsgInternalToTable,
		MsgSubscriberTable,
		MsgTemplateTable,
		MsgTypeTable,
		MsgNlogTable,
		MsgNlogAlertTable,
		OrgRoleUserTable,
		MsgSilenceTable,
		UserTable,
	}
)

func init() {
	FileIdentityTable.ForeignKeys[0].RefTable = FileSourceTable
	FileIdentityTable.Annotation = &entsql.Annotation{
		Table: "file_identity",
	}
	FileSourceTable.Annotation = &entsql.Annotation{
		Table: "file_source",
	}
	MsgAlertTable.Annotation = &entsql.Annotation{
		Table: "msg_alert",
	}
	MsgChannelTable.Annotation = &entsql.Annotation{
		Table: "msg_channel",
	}
	MsgEventTable.ForeignKeys[0].RefTable = MsgTypeTable
	MsgEventTable.Annotation = &entsql.Annotation{
		Table: "msg_event",
	}
	MsgInternalTable.Annotation = &entsql.Annotation{
		Table: "msg_internal",
	}
	MsgInternalToTable.ForeignKeys[0].RefTable = MsgInternalTable
	MsgInternalToTable.ForeignKeys[1].RefTable = UserTable
	MsgInternalToTable.Annotation = &entsql.Annotation{
		Table: "msg_internal_to",
	}
	MsgSubscriberTable.ForeignKeys[0].RefTable = UserTable
	MsgSubscriberTable.ForeignKeys[1].RefTable = MsgTypeTable
	MsgSubscriberTable.Annotation = &entsql.Annotation{
		Table: "msg_subscriber",
	}
	MsgTemplateTable.ForeignKeys[0].RefTable = MsgEventTable
	MsgTemplateTable.Annotation = &entsql.Annotation{
		Table: "msg_template",
	}
	MsgTypeTable.Annotation = &entsql.Annotation{
		Table: "msg_type",
	}
	MsgNlogTable.Annotation = &entsql.Annotation{
		Table: "msg_nlog",
	}
	MsgNlogAlertTable.ForeignKeys[0].RefTable = MsgNlogTable
	MsgNlogAlertTable.ForeignKeys[1].RefTable = MsgAlertTable
	MsgNlogAlertTable.Annotation = &entsql.Annotation{
		Table: "msg_nlog_alert",
	}
	OrgRoleUserTable.Annotation = &entsql.Annotation{
		Table: "org_role_user",
	}
	MsgSilenceTable.ForeignKeys[0].RefTable = UserTable
	MsgSilenceTable.Annotation = &entsql.Annotation{
		Table: "msg_silence",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}
