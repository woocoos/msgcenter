// Code generated by ent, DO NOT EDIT.

package msgtemplate

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/msgcenter/pkg/profile"
)

const (
	// Label holds the string label denoting the msgtemplate type in the database.
	Label = "msg_template"
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
	// FieldMsgEventID holds the string denoting the msg_event_id field in the database.
	FieldMsgEventID = "msg_event_id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldReceiverType holds the string denoting the receiver_type field in the database.
	FieldReceiverType = "receiver_type"
	// FieldFormat holds the string denoting the format field in the database.
	FieldFormat = "format"
	// FieldSubject holds the string denoting the subject field in the database.
	FieldSubject = "subject"
	// FieldFrom holds the string denoting the from field in the database.
	FieldFrom = "from"
	// FieldTo holds the string denoting the to field in the database.
	FieldTo = "to"
	// FieldCc holds the string denoting the cc field in the database.
	FieldCc = "cc"
	// FieldBcc holds the string denoting the bcc field in the database.
	FieldBcc = "bcc"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldTpl holds the string denoting the tpl field in the database.
	FieldTpl = "tpl"
	// FieldAttachments holds the string denoting the attachments field in the database.
	FieldAttachments = "attachments"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// Table holds the table name of the msgtemplate in the database.
	Table = "msg_template"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "msg_template"
	// EventInverseTable is the table name for the MsgEvent entity.
	// It exists in this package in order to avoid circular dependency with the "msgevent" package.
	EventInverseTable = "msg_event"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "msg_event_id"
)

// Columns holds all SQL columns for msgtemplate fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldMsgTypeID,
	FieldMsgEventID,
	FieldTenantID,
	FieldName,
	FieldStatus,
	FieldReceiverType,
	FieldFormat,
	FieldSubject,
	FieldFrom,
	FieldTo,
	FieldCc,
	FieldBcc,
	FieldBody,
	FieldTpl,
	FieldAttachments,
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
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

const DefaultStatus typex.SimpleStatus = "inactive"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s typex.SimpleStatus) error {
	switch s.String() {
	case "active", "inactive", "processing":
		return nil
	default:
		return fmt.Errorf("msgtemplate: invalid enum value for status field: %q", s)
	}
}

// ReceiverTypeValidator is a validator for the "receiver_type" field enum values. It is called by the builders before save.
func ReceiverTypeValidator(rt profile.ReceiverType) error {
	switch rt.String() {
	case "email", "internal", "webhook":
		return nil
	default:
		return fmt.Errorf("msgtemplate: invalid enum value for receiver_type field: %q", rt)
	}
}

// Format defines the type for the "format" enum field.
type Format string

// Format values.
const (
	FormatTxt  Format = "txt"
	FormatHTML Format = "html"
)

func (f Format) String() string {
	return string(f)
}

// FormatValidator is a validator for the "format" field enum values. It is called by the builders before save.
func FormatValidator(f Format) error {
	switch f {
	case FormatTxt, FormatHTML:
		return nil
	default:
		return fmt.Errorf("msgtemplate: invalid enum value for format field: %q", f)
	}
}

// OrderOption defines the ordering options for the MsgTemplate queries.
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

// ByMsgEventID orders the results by the msg_event_id field.
func ByMsgEventID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMsgEventID, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByReceiverType orders the results by the receiver_type field.
func ByReceiverType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReceiverType, opts...).ToFunc()
}

// ByFormat orders the results by the format field.
func ByFormat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFormat, opts...).ToFunc()
}

// BySubject orders the results by the subject field.
func BySubject(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubject, opts...).ToFunc()
}

// ByFrom orders the results by the from field.
func ByFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrom, opts...).ToFunc()
}

// ByTo orders the results by the to field.
func ByTo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTo, opts...).ToFunc()
}

// ByCc orders the results by the cc field.
func ByCc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCc, opts...).ToFunc()
}

// ByBcc orders the results by the bcc field.
func ByBcc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBcc, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByTpl orders the results by the tpl field.
func ByTpl(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTpl, opts...).ToFunc()
}

// ByAttachments orders the results by the attachments field.
func ByAttachments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttachments, opts...).ToFunc()
}

// ByComments orders the results by the comments field.
func ByComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComments, opts...).ToFunc()
}

// ByEventField orders the results by event field.
func ByEventField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventStep(), sql.OrderByField(field, opts...))
	}
}
func newEventStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, EventTable, EventColumn),
	)
}

var (
	// typex.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*typex.SimpleStatus)(nil)
	// typex.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*typex.SimpleStatus)(nil)
)

var (
	// profile.ReceiverType must implement graphql.Marshaler.
	_ graphql.Marshaler = (*profile.ReceiverType)(nil)
	// profile.ReceiverType must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*profile.ReceiverType)(nil)
)

// MarshalGQL implements graphql.Marshaler interface.
func (e Format) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Format) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Format(str)
	if err := FormatValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Format", str)
	}
	return nil
}
