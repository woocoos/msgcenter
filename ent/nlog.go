// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/msgcenter/ent/nlog"
	"github.com/woocoos/msgcenter/pkg/profile"
)

// Nlog is the model entity for the Nlog schema.
type Nlog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID int `json:"tenant_id,omitempty"`
	// 分组键
	GroupKey string `json:"group_key,omitempty"`
	// 接收组名称
	Receiver string `json:"receiver,omitempty"`
	// 支持的消息模式:站内信,app推送,邮件,短信,微信等
	ReceiverType profile.ReceiverType `json:"receiver_type,omitempty"`
	// 通道的索引位置
	Idx int `json:"idx,omitempty"`
	// 发送时间
	SendAt time.Time `json:"send_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 过期时间
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NlogQuery when eager-loading is set.
	Edges        NlogEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NlogEdges holds the relations/edges for other nodes in the graph.
type NlogEdges struct {
	// Alerts holds the value of the alerts edge.
	Alerts []*MsgAlert `json:"alerts,omitempty"`
	// NlogAlert holds the value of the nlog_alert edge.
	NlogAlert []*NlogAlert `json:"nlog_alert,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedAlerts    map[string][]*MsgAlert
	namedNlogAlert map[string][]*NlogAlert
}

// AlertsOrErr returns the Alerts value or an error if the edge
// was not loaded in eager-loading.
func (e NlogEdges) AlertsOrErr() ([]*MsgAlert, error) {
	if e.loadedTypes[0] {
		return e.Alerts, nil
	}
	return nil, &NotLoadedError{edge: "alerts"}
}

// NlogAlertOrErr returns the NlogAlert value or an error if the edge
// was not loaded in eager-loading.
func (e NlogEdges) NlogAlertOrErr() ([]*NlogAlert, error) {
	if e.loadedTypes[1] {
		return e.NlogAlert, nil
	}
	return nil, &NotLoadedError{edge: "nlog_alert"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Nlog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case nlog.FieldID, nlog.FieldTenantID, nlog.FieldIdx:
			values[i] = new(sql.NullInt64)
		case nlog.FieldGroupKey, nlog.FieldReceiver, nlog.FieldReceiverType:
			values[i] = new(sql.NullString)
		case nlog.FieldSendAt, nlog.FieldCreatedAt, nlog.FieldUpdatedAt, nlog.FieldExpiresAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Nlog fields.
func (n *Nlog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case nlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case nlog.FieldTenantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				n.TenantID = int(value.Int64)
			}
		case nlog.FieldGroupKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_key", values[i])
			} else if value.Valid {
				n.GroupKey = value.String
			}
		case nlog.FieldReceiver:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver", values[i])
			} else if value.Valid {
				n.Receiver = value.String
			}
		case nlog.FieldReceiverType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_type", values[i])
			} else if value.Valid {
				n.ReceiverType = profile.ReceiverType(value.String)
			}
		case nlog.FieldIdx:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field idx", values[i])
			} else if value.Valid {
				n.Idx = int(value.Int64)
			}
		case nlog.FieldSendAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field send_at", values[i])
			} else if value.Valid {
				n.SendAt = value.Time
			}
		case nlog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case nlog.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		case nlog.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				n.ExpiresAt = value.Time
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Nlog.
// This includes values selected through modifiers, order, etc.
func (n *Nlog) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryAlerts queries the "alerts" edge of the Nlog entity.
func (n *Nlog) QueryAlerts() *MsgAlertQuery {
	return NewNlogClient(n.config).QueryAlerts(n)
}

// QueryNlogAlert queries the "nlog_alert" edge of the Nlog entity.
func (n *Nlog) QueryNlogAlert() *NlogAlertQuery {
	return NewNlogClient(n.config).QueryNlogAlert(n)
}

// Update returns a builder for updating this Nlog.
// Note that you need to call Nlog.Unwrap() before calling this method if this Nlog
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Nlog) Update() *NlogUpdateOne {
	return NewNlogClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Nlog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Nlog) Unwrap() *Nlog {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Nlog is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Nlog) String() string {
	var builder strings.Builder
	builder.WriteString("Nlog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(fmt.Sprintf("%v", n.TenantID))
	builder.WriteString(", ")
	builder.WriteString("group_key=")
	builder.WriteString(n.GroupKey)
	builder.WriteString(", ")
	builder.WriteString("receiver=")
	builder.WriteString(n.Receiver)
	builder.WriteString(", ")
	builder.WriteString("receiver_type=")
	builder.WriteString(fmt.Sprintf("%v", n.ReceiverType))
	builder.WriteString(", ")
	builder.WriteString("idx=")
	builder.WriteString(fmt.Sprintf("%v", n.Idx))
	builder.WriteString(", ")
	builder.WriteString("send_at=")
	builder.WriteString(n.SendAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(n.ExpiresAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedAlerts returns the Alerts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (n *Nlog) NamedAlerts(name string) ([]*MsgAlert, error) {
	if n.Edges.namedAlerts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := n.Edges.namedAlerts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (n *Nlog) appendNamedAlerts(name string, edges ...*MsgAlert) {
	if n.Edges.namedAlerts == nil {
		n.Edges.namedAlerts = make(map[string][]*MsgAlert)
	}
	if len(edges) == 0 {
		n.Edges.namedAlerts[name] = []*MsgAlert{}
	} else {
		n.Edges.namedAlerts[name] = append(n.Edges.namedAlerts[name], edges...)
	}
}

// NamedNlogAlert returns the NlogAlert named value or an error if the edge was not
// loaded in eager-loading with this name.
func (n *Nlog) NamedNlogAlert(name string) ([]*NlogAlert, error) {
	if n.Edges.namedNlogAlert == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := n.Edges.namedNlogAlert[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (n *Nlog) appendNamedNlogAlert(name string, edges ...*NlogAlert) {
	if n.Edges.namedNlogAlert == nil {
		n.Edges.namedNlogAlert = make(map[string][]*NlogAlert)
	}
	if len(edges) == 0 {
		n.Edges.namedNlogAlert[name] = []*NlogAlert{}
	} else {
		n.Edges.namedNlogAlert[name] = append(n.Edges.namedNlogAlert[name], edges...)
	}
}

// Nlogs is a parsable slice of Nlog.
type Nlogs []*Nlog
