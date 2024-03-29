// Code generated by ent, DO NOT EDIT.

package msginternal

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the msginternal type in the database.
	Label = "msg_internal"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldSubject holds the string denoting the subject field in the database.
	FieldSubject = "subject"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldFormat holds the string denoting the format field in the database.
	FieldFormat = "format"
	// FieldRedirect holds the string denoting the redirect field in the database.
	FieldRedirect = "redirect"
	// EdgeMsgInternalTo holds the string denoting the msg_internal_to edge name in mutations.
	EdgeMsgInternalTo = "msg_internal_to"
	// Table holds the table name of the msginternal in the database.
	Table = "msg_internal"
	// MsgInternalToTable is the table that holds the msg_internal_to relation/edge.
	MsgInternalToTable = "msg_internal_to"
	// MsgInternalToInverseTable is the table name for the MsgInternalTo entity.
	// It exists in this package in order to avoid circular dependency with the "msginternalto" package.
	MsgInternalToInverseTable = "msg_internal_to"
	// MsgInternalToColumn is the table column denoting the msg_internal_to relation/edge.
	MsgInternalToColumn = "msg_internal_id"
)

// Columns holds all SQL columns for msginternal fields.
var Columns = []string{
	FieldID,
	FieldTenantID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldCategory,
	FieldSubject,
	FieldBody,
	FieldFormat,
	FieldRedirect,
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
	Hooks        [3]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	CategoryValidator func(string) error
)

// OrderOption defines the ordering options for the MsgInternal queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
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

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// BySubject orders the results by the subject field.
func BySubject(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubject, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByFormat orders the results by the format field.
func ByFormat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFormat, opts...).ToFunc()
}

// ByRedirect orders the results by the redirect field.
func ByRedirect(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRedirect, opts...).ToFunc()
}

// ByMsgInternalToCount orders the results by msg_internal_to count.
func ByMsgInternalToCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMsgInternalToStep(), opts...)
	}
}

// ByMsgInternalTo orders the results by msg_internal_to terms.
func ByMsgInternalTo(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMsgInternalToStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMsgInternalToStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MsgInternalToInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MsgInternalToTable, MsgInternalToColumn),
	)
}
