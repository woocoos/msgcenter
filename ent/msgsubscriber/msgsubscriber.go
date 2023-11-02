// Code generated by ent, DO NOT EDIT.

package msgsubscriber

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the msgsubscriber type in the database.
	Label = "msg_subscriber"
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
	// FieldMsgTypeID holds the string denoting the msg_type_id field in the database.
	FieldMsgTypeID = "msg_type_id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldOrgRoleID holds the string denoting the org_role_id field in the database.
	FieldOrgRoleID = "org_role_id"
	// FieldExclude holds the string denoting the exclude field in the database.
	FieldExclude = "exclude"
	// EdgeMsgType holds the string denoting the msg_type edge name in mutations.
	EdgeMsgType = "msg_type"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the msgsubscriber in the database.
	Table = "msg_subscriber"
	// MsgTypeTable is the table that holds the msg_type relation/edge.
	MsgTypeTable = "msg_subscriber"
	// MsgTypeInverseTable is the table name for the MsgType entity.
	// It exists in this package in order to avoid circular dependency with the "msgtype" package.
	MsgTypeInverseTable = "msg_type"
	// MsgTypeColumn is the table column denoting the msg_type relation/edge.
	MsgTypeColumn = "msg_type_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "msg_subscriber"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "user"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for msgsubscriber fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldMsgTypeID,
	FieldTenantID,
	FieldUserID,
	FieldOrgRoleID,
	FieldExclude,
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
	// DefaultExclude holds the default value on creation for the "exclude" field.
	DefaultExclude bool
)

// OrderOption defines the ordering options for the MsgSubscriber queries.
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

// ByMsgTypeID orders the results by the msg_type_id field.
func ByMsgTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMsgTypeID, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByOrgRoleID orders the results by the org_role_id field.
func ByOrgRoleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrgRoleID, opts...).ToFunc()
}

// ByExclude orders the results by the exclude field.
func ByExclude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExclude, opts...).ToFunc()
}

// ByMsgTypeField orders the results by msg_type field.
func ByMsgTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMsgTypeStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newMsgTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MsgTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MsgTypeTable, MsgTypeColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
