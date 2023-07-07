// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/pkg/profile"
)

// MsgTemplate is the model entity for the MsgTemplate schema.
type MsgTemplate struct {
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
	// 应用消息类型ID
	MsgTypeID int `json:"msg_type_id,omitempty"`
	// 消息事件ID
	MsgEventID int `json:"msg_event_id,omitempty"`
	// 组织ID
	TenantID int `json:"tenant_id,omitempty"`
	// 消息模板名称
	Name string `json:"name,omitempty"`
	// 状态
	Status typex.SimpleStatus `json:"status,omitempty"`
	// 消息模式:站内信,app推送,邮件,短信,微信等
	ReceiverType profile.ReceiverType `json:"receiver_type,omitempty"`
	// 消息类型:文本,网页,需要结合mod确定支持的格式
	Format msgtemplate.Format `json:"format,omitempty"`
	// 标题
	Subject string `json:"subject,omitempty"`
	// 发件人
	From string `json:"from,omitempty"`
	// 收件人
	To string `json:"to,omitempty"`
	// 抄送
	Cc string `json:"cc,omitempty"`
	// 密送
	Bcc string `json:"bcc,omitempty"`
	// 消息体
	Body string `json:"body,omitempty"`
	// 模板地址
	Tpl string `json:"tpl,omitempty"`
	// 模板地址
	TplFileID int `json:"tpl_file_id,omitempty"`
	// 附件地址,多个附件用逗号分隔
	Attachments string `json:"attachments,omitempty"`
	// 附件地址,多个附件用逗号分隔
	AttachmentsFileIds string `json:"attachments_file_ids,omitempty"`
	// 备注
	Comments string `json:"comments,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MsgTemplateQuery when eager-loading is set.
	Edges        MsgTemplateEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MsgTemplateEdges holds the relations/edges for other nodes in the graph.
type MsgTemplateEdges struct {
	// Event holds the value of the event edge.
	Event *MsgEvent `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MsgTemplateEdges) EventOrErr() (*MsgEvent, error) {
	if e.loadedTypes[0] {
		if e.Event == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: msgevent.Label}
		}
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MsgTemplate) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case msgtemplate.FieldID, msgtemplate.FieldCreatedBy, msgtemplate.FieldUpdatedBy, msgtemplate.FieldMsgTypeID, msgtemplate.FieldMsgEventID, msgtemplate.FieldTenantID, msgtemplate.FieldTplFileID:
			values[i] = new(sql.NullInt64)
		case msgtemplate.FieldName, msgtemplate.FieldStatus, msgtemplate.FieldReceiverType, msgtemplate.FieldFormat, msgtemplate.FieldSubject, msgtemplate.FieldFrom, msgtemplate.FieldTo, msgtemplate.FieldCc, msgtemplate.FieldBcc, msgtemplate.FieldBody, msgtemplate.FieldTpl, msgtemplate.FieldAttachments, msgtemplate.FieldAttachmentsFileIds, msgtemplate.FieldComments:
			values[i] = new(sql.NullString)
		case msgtemplate.FieldCreatedAt, msgtemplate.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MsgTemplate fields.
func (mt *MsgTemplate) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case msgtemplate.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mt.ID = int(value.Int64)
		case msgtemplate.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				mt.CreatedBy = int(value.Int64)
			}
		case msgtemplate.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mt.CreatedAt = value.Time
			}
		case msgtemplate.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				mt.UpdatedBy = int(value.Int64)
			}
		case msgtemplate.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mt.UpdatedAt = value.Time
			}
		case msgtemplate.FieldMsgTypeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field msg_type_id", values[i])
			} else if value.Valid {
				mt.MsgTypeID = int(value.Int64)
			}
		case msgtemplate.FieldMsgEventID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field msg_event_id", values[i])
			} else if value.Valid {
				mt.MsgEventID = int(value.Int64)
			}
		case msgtemplate.FieldTenantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				mt.TenantID = int(value.Int64)
			}
		case msgtemplate.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mt.Name = value.String
			}
		case msgtemplate.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				mt.Status = typex.SimpleStatus(value.String)
			}
		case msgtemplate.FieldReceiverType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver_type", values[i])
			} else if value.Valid {
				mt.ReceiverType = profile.ReceiverType(value.String)
			}
		case msgtemplate.FieldFormat:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field format", values[i])
			} else if value.Valid {
				mt.Format = msgtemplate.Format(value.String)
			}
		case msgtemplate.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				mt.Subject = value.String
			}
		case msgtemplate.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				mt.From = value.String
			}
		case msgtemplate.FieldTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				mt.To = value.String
			}
		case msgtemplate.FieldCc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cc", values[i])
			} else if value.Valid {
				mt.Cc = value.String
			}
		case msgtemplate.FieldBcc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bcc", values[i])
			} else if value.Valid {
				mt.Bcc = value.String
			}
		case msgtemplate.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				mt.Body = value.String
			}
		case msgtemplate.FieldTpl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tpl", values[i])
			} else if value.Valid {
				mt.Tpl = value.String
			}
		case msgtemplate.FieldTplFileID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tpl_file_id", values[i])
			} else if value.Valid {
				mt.TplFileID = int(value.Int64)
			}
		case msgtemplate.FieldAttachments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attachments", values[i])
			} else if value.Valid {
				mt.Attachments = value.String
			}
		case msgtemplate.FieldAttachmentsFileIds:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attachments_file_ids", values[i])
			} else if value.Valid {
				mt.AttachmentsFileIds = value.String
			}
		case msgtemplate.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				mt.Comments = value.String
			}
		default:
			mt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MsgTemplate.
// This includes values selected through modifiers, order, etc.
func (mt *MsgTemplate) Value(name string) (ent.Value, error) {
	return mt.selectValues.Get(name)
}

// QueryEvent queries the "event" edge of the MsgTemplate entity.
func (mt *MsgTemplate) QueryEvent() *MsgEventQuery {
	return NewMsgTemplateClient(mt.config).QueryEvent(mt)
}

// Update returns a builder for updating this MsgTemplate.
// Note that you need to call MsgTemplate.Unwrap() before calling this method if this MsgTemplate
// was returned from a transaction, and the transaction was committed or rolled back.
func (mt *MsgTemplate) Update() *MsgTemplateUpdateOne {
	return NewMsgTemplateClient(mt.config).UpdateOne(mt)
}

// Unwrap unwraps the MsgTemplate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mt *MsgTemplate) Unwrap() *MsgTemplate {
	_tx, ok := mt.config.driver.(*txDriver)
	if !ok {
		panic("ent: MsgTemplate is not a transactional entity")
	}
	mt.config.driver = _tx.drv
	return mt
}

// String implements the fmt.Stringer.
func (mt *MsgTemplate) String() string {
	var builder strings.Builder
	builder.WriteString("MsgTemplate(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mt.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", mt.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", mt.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("msg_type_id=")
	builder.WriteString(fmt.Sprintf("%v", mt.MsgTypeID))
	builder.WriteString(", ")
	builder.WriteString("msg_event_id=")
	builder.WriteString(fmt.Sprintf("%v", mt.MsgEventID))
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(fmt.Sprintf("%v", mt.TenantID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(mt.Name)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", mt.Status))
	builder.WriteString(", ")
	builder.WriteString("receiver_type=")
	builder.WriteString(fmt.Sprintf("%v", mt.ReceiverType))
	builder.WriteString(", ")
	builder.WriteString("format=")
	builder.WriteString(fmt.Sprintf("%v", mt.Format))
	builder.WriteString(", ")
	builder.WriteString("subject=")
	builder.WriteString(mt.Subject)
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(mt.From)
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(mt.To)
	builder.WriteString(", ")
	builder.WriteString("cc=")
	builder.WriteString(mt.Cc)
	builder.WriteString(", ")
	builder.WriteString("bcc=")
	builder.WriteString(mt.Bcc)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(mt.Body)
	builder.WriteString(", ")
	builder.WriteString("tpl=")
	builder.WriteString(mt.Tpl)
	builder.WriteString(", ")
	builder.WriteString("tpl_file_id=")
	builder.WriteString(fmt.Sprintf("%v", mt.TplFileID))
	builder.WriteString(", ")
	builder.WriteString("attachments=")
	builder.WriteString(mt.Attachments)
	builder.WriteString(", ")
	builder.WriteString("attachments_file_ids=")
	builder.WriteString(mt.AttachmentsFileIds)
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(mt.Comments)
	builder.WriteByte(')')
	return builder.String()
}

// MsgTemplates is a parsable slice of MsgTemplate.
type MsgTemplates []*MsgTemplate
