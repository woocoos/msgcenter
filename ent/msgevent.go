// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/pkg/profile"
)

// MsgEvent is the model entity for the MsgEvent schema.
type MsgEvent struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int `json:"created_by,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int `json:"updated_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 消息类型ID
	MsgTypeID int `json:"msg_type_id,omitempty"`
	// 消息事件名称,应用内唯一
	Name string `json:"name,omitempty"`
	// 状态
	Status typex.SimpleStatus `json:"status,omitempty"`
	// 备注
	Comments string `json:"comments,omitempty"`
	// 消息路由配置
	Route *profile.Route `json:"route,omitempty"`
	// 根据route配置对应的以,分隔的mode列表
	Modes string `json:"modes,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MsgEventQuery when eager-loading is set.
	Edges        MsgEventEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MsgEventEdges holds the relations/edges for other nodes in the graph.
type MsgEventEdges struct {
	// 消息类型
	MsgType *MsgType `json:"msg_type,omitempty"`
	// 自定义的消息模板
	CustomerTemplate []*MsgTemplate `json:"customer_template,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedCustomerTemplate map[string][]*MsgTemplate
}

// MsgTypeOrErr returns the MsgType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MsgEventEdges) MsgTypeOrErr() (*MsgType, error) {
	if e.loadedTypes[0] {
		if e.MsgType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: msgtype.Label}
		}
		return e.MsgType, nil
	}
	return nil, &NotLoadedError{edge: "msg_type"}
}

// CustomerTemplateOrErr returns the CustomerTemplate value or an error if the edge
// was not loaded in eager-loading.
func (e MsgEventEdges) CustomerTemplateOrErr() ([]*MsgTemplate, error) {
	if e.loadedTypes[1] {
		return e.CustomerTemplate, nil
	}
	return nil, &NotLoadedError{edge: "customer_template"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MsgEvent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case msgevent.FieldRoute:
			values[i] = new([]byte)
		case msgevent.FieldID, msgevent.FieldCreatedBy, msgevent.FieldUpdatedBy, msgevent.FieldMsgTypeID:
			values[i] = new(sql.NullInt64)
		case msgevent.FieldName, msgevent.FieldStatus, msgevent.FieldComments, msgevent.FieldModes:
			values[i] = new(sql.NullString)
		case msgevent.FieldCreatedAt, msgevent.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MsgEvent fields.
func (me *MsgEvent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case msgevent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			me.ID = int(value.Int64)
		case msgevent.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				me.CreatedBy = int(value.Int64)
			}
		case msgevent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				me.CreatedAt = value.Time
			}
		case msgevent.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				me.UpdatedBy = int(value.Int64)
			}
		case msgevent.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				me.UpdatedAt = value.Time
			}
		case msgevent.FieldMsgTypeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field msg_type_id", values[i])
			} else if value.Valid {
				me.MsgTypeID = int(value.Int64)
			}
		case msgevent.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				me.Name = value.String
			}
		case msgevent.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				me.Status = typex.SimpleStatus(value.String)
			}
		case msgevent.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				me.Comments = value.String
			}
		case msgevent.FieldRoute:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field route", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &me.Route); err != nil {
					return fmt.Errorf("unmarshal field route: %w", err)
				}
			}
		case msgevent.FieldModes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field modes", values[i])
			} else if value.Valid {
				me.Modes = value.String
			}
		default:
			me.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MsgEvent.
// This includes values selected through modifiers, order, etc.
func (me *MsgEvent) Value(name string) (ent.Value, error) {
	return me.selectValues.Get(name)
}

// QueryMsgType queries the "msg_type" edge of the MsgEvent entity.
func (me *MsgEvent) QueryMsgType() *MsgTypeQuery {
	return NewMsgEventClient(me.config).QueryMsgType(me)
}

// QueryCustomerTemplate queries the "customer_template" edge of the MsgEvent entity.
func (me *MsgEvent) QueryCustomerTemplate() *MsgTemplateQuery {
	return NewMsgEventClient(me.config).QueryCustomerTemplate(me)
}

// Update returns a builder for updating this MsgEvent.
// Note that you need to call MsgEvent.Unwrap() before calling this method if this MsgEvent
// was returned from a transaction, and the transaction was committed or rolled back.
func (me *MsgEvent) Update() *MsgEventUpdateOne {
	return NewMsgEventClient(me.config).UpdateOne(me)
}

// Unwrap unwraps the MsgEvent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (me *MsgEvent) Unwrap() *MsgEvent {
	_tx, ok := me.config.driver.(*txDriver)
	if !ok {
		panic("ent: MsgEvent is not a transactional entity")
	}
	me.config.driver = _tx.drv
	return me
}

// String implements the fmt.Stringer.
func (me *MsgEvent) String() string {
	var builder strings.Builder
	builder.WriteString("MsgEvent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", me.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", me.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(me.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", me.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(me.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("msg_type_id=")
	builder.WriteString(fmt.Sprintf("%v", me.MsgTypeID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(me.Name)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", me.Status))
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(me.Comments)
	builder.WriteString(", ")
	builder.WriteString("route=")
	builder.WriteString(fmt.Sprintf("%v", me.Route))
	builder.WriteString(", ")
	builder.WriteString("modes=")
	builder.WriteString(me.Modes)
	builder.WriteByte(')')
	return builder.String()
}

// NamedCustomerTemplate returns the CustomerTemplate named value or an error if the edge was not
// loaded in eager-loading with this name.
func (me *MsgEvent) NamedCustomerTemplate(name string) ([]*MsgTemplate, error) {
	if me.Edges.namedCustomerTemplate == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := me.Edges.namedCustomerTemplate[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (me *MsgEvent) appendNamedCustomerTemplate(name string, edges ...*MsgTemplate) {
	if me.Edges.namedCustomerTemplate == nil {
		me.Edges.namedCustomerTemplate = make(map[string][]*MsgTemplate)
	}
	if len(edges) == 0 {
		me.Edges.namedCustomerTemplate[name] = []*MsgTemplate{}
	} else {
		me.Edges.namedCustomerTemplate[name] = append(me.Edges.namedCustomerTemplate[name], edges...)
	}
}

// MsgEvents is a parsable slice of MsgEvent.
type MsgEvents []*MsgEvent
