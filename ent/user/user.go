// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPrincipalName holds the string denoting the principal_name field in the database.
	FieldPrincipalName = "principal_name"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// Table holds the table name of the user in the database.
	Table = "user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldPrincipalName,
	FieldDisplayName,
	FieldEmail,
	FieldMobile,
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
	Hooks [1]ent.Hook
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	MobileValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPrincipalName orders the results by the principal_name field.
func ByPrincipalName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrincipalName, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}
