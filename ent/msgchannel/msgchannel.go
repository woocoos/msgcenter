// Code generated by ent, DO NOT EDIT.

package msgchannel

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/msgcenter/pkg/profile"
)

const (
	// Label holds the string label denoting the msgchannel type in the database.
	Label = "msg_channel"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldReceiverType holds the string denoting the receiver_type field in the database.
	FieldReceiverType = "receiver_type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldReceiver holds the string denoting the receiver field in the database.
	FieldReceiver = "receiver"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// Table holds the table name of the msgchannel in the database.
	Table = "msg_channel"
)

// Columns holds all SQL columns for msgchannel fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldName,
	FieldTenantID,
	FieldReceiverType,
	FieldStatus,
	FieldReceiver,
	FieldComments,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/woocoos/msgcenter/ent/runtime"
var (
	Hooks [3]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// ReceiverTypeValidator is a validator for the "receiver_type" field enum values. It is called by the builders before save.
func ReceiverTypeValidator(rt profile.ReceiverType) error {
	switch rt.String() {
	case "email", "message", "webhook":
		return nil
	default:
		return fmt.Errorf("msgchannel: invalid enum value for receiver_type field: %q", rt)
	}
}

const DefaultStatus typex.SimpleStatus = "inactive"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s typex.SimpleStatus) error {
	switch s.String() {
	case "active", "inactive", "processing", "disabled":
		return nil
	default:
		return fmt.Errorf("msgchannel: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the MsgChannel queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByReceiverType orders the results by the receiver_type field.
func ByReceiverType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReceiverType, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByComments orders the results by the comments field.
func ByComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComments, opts...).ToFunc()
}

var (
	// profile.ReceiverType must implement graphql.Marshaler.
	_ graphql.Marshaler = (*profile.ReceiverType)(nil)
	// profile.ReceiverType must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*profile.ReceiverType)(nil)
)

var (
	// typex.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*typex.SimpleStatus)(nil)
	// typex.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*typex.SimpleStatus)(nil)
)
