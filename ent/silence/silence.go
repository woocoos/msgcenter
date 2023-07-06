// Code generated by ent, DO NOT EDIT.

package silence

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/msgcenter/pkg/alert"
)

const (
	// Label holds the string label denoting the silence type in the database.
	Label = "silence"
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
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldMatchers holds the string denoting the matchers field in the database.
	FieldMatchers = "matchers"
	// FieldStartsAt holds the string denoting the starts_at field in the database.
	FieldStartsAt = "starts_at"
	// FieldEndsAt holds the string denoting the ends_at field in the database.
	FieldEndsAt = "ends_at"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the silence in the database.
	Table = "msg_silence"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "msg_silence"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "user"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "created_by"
)

// Columns holds all SQL columns for silence fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldTenantID,
	FieldMatchers,
	FieldStartsAt,
	FieldEndsAt,
	FieldComments,
	FieldState,
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
	Hooks        [2]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int
)

const DefaultState alert.SilenceState = "active"

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s alert.SilenceState) error {
	switch s.String() {
	case "expired", "active", "pending":
		return nil
	default:
		return fmt.Errorf("silence: invalid enum value for state field: %q", s)
	}
}

// OrderOption defines the ordering options for the Silence queries.
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

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByStartsAt orders the results by the starts_at field.
func ByStartsAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartsAt, opts...).ToFunc()
}

// ByEndsAt orders the results by the ends_at field.
func ByEndsAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndsAt, opts...).ToFunc()
}

// ByComments orders the results by the comments field.
func ByComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComments, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}

var (
	// alert.SilenceState must implement graphql.Marshaler.
	_ graphql.Marshaler = (*alert.SilenceState)(nil)
	// alert.SilenceState must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*alert.SilenceState)(nil)
)
