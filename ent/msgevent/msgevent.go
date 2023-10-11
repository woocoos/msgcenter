// Code generated by ent, DO NOT EDIT.

package msgevent

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/entco/schemax/typex"
)

const (
	// Label holds the string label denoting the msgevent type in the database.
	Label = "msg_event"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldRoute holds the string denoting the route field in the database.
	FieldRoute = "route"
	// FieldModes holds the string denoting the modes field in the database.
	FieldModes = "modes"
	// EdgeMsgType holds the string denoting the msg_type edge name in mutations.
	EdgeMsgType = "msg_type"
	// EdgeCustomerTemplate holds the string denoting the customer_template edge name in mutations.
	EdgeCustomerTemplate = "customer_template"
	// Table holds the table name of the msgevent in the database.
	Table = "msg_event"
	// MsgTypeTable is the table that holds the msg_type relation/edge.
	MsgTypeTable = "msg_event"
	// MsgTypeInverseTable is the table name for the MsgType entity.
	// It exists in this package in order to avoid circular dependency with the "msgtype" package.
	MsgTypeInverseTable = "msg_type"
	// MsgTypeColumn is the table column denoting the msg_type relation/edge.
	MsgTypeColumn = "msg_type_id"
	// CustomerTemplateTable is the table that holds the customer_template relation/edge.
	CustomerTemplateTable = "msg_template"
	// CustomerTemplateInverseTable is the table name for the MsgTemplate entity.
	// It exists in this package in order to avoid circular dependency with the "msgtemplate" package.
	CustomerTemplateInverseTable = "msg_template"
	// CustomerTemplateColumn is the table column denoting the customer_template relation/edge.
	CustomerTemplateColumn = "msg_event_id"
)

// Columns holds all SQL columns for msgevent fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldMsgTypeID,
	FieldName,
	FieldStatus,
	FieldComments,
	FieldRoute,
	FieldModes,
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
	Hooks [4]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

const DefaultStatus typex.SimpleStatus = "inactive"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s typex.SimpleStatus) error {
	switch s.String() {
	case "active", "inactive", "processing", "disabled":
		return nil
	default:
		return fmt.Errorf("msgevent: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the MsgEvent queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByComments orders the results by the comments field.
func ByComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComments, opts...).ToFunc()
}

// ByModes orders the results by the modes field.
func ByModes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModes, opts...).ToFunc()
}

// ByMsgTypeField orders the results by msg_type field.
func ByMsgTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMsgTypeStep(), sql.OrderByField(field, opts...))
	}
}

// ByCustomerTemplateCount orders the results by customer_template count.
func ByCustomerTemplateCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCustomerTemplateStep(), opts...)
	}
}

// ByCustomerTemplate orders the results by customer_template terms.
func ByCustomerTemplate(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerTemplateStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMsgTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MsgTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MsgTypeTable, MsgTypeColumn),
	)
}
func newCustomerTemplateStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerTemplateInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CustomerTemplateTable, CustomerTemplateColumn),
	)
}

var (
	// typex.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*typex.SimpleStatus)(nil)
	// typex.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*typex.SimpleStatus)(nil)
)
