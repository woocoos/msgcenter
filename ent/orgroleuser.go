// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/msgcenter/ent/orgroleuser"
)

// OrgRoleUser is the model entity for the OrgRoleUser schema.
type OrgRoleUser struct {
	config `json:"-"`
	// ID of the ent.
	// ID
	ID int `json:"id,omitempty"`
	// 组织角色ID
	OrgRoleID int `json:"org_role_id,omitempty"`
	// 组织用户ID
	OrgUserID int `json:"org_user_id,omitempty"`
	// 组织ID
	OrgID int `json:"org_id,omitempty"`
	// 用户ID
	UserID       int `json:"user_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgRoleUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orgroleuser.FieldID, orgroleuser.FieldOrgRoleID, orgroleuser.FieldOrgUserID, orgroleuser.FieldOrgID, orgroleuser.FieldUserID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgRoleUser fields.
func (oru *OrgRoleUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orgroleuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oru.ID = int(value.Int64)
		case orgroleuser.FieldOrgRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_role_id", values[i])
			} else if value.Valid {
				oru.OrgRoleID = int(value.Int64)
			}
		case orgroleuser.FieldOrgUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_user_id", values[i])
			} else if value.Valid {
				oru.OrgUserID = int(value.Int64)
			}
		case orgroleuser.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				oru.OrgID = int(value.Int64)
			}
		case orgroleuser.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				oru.UserID = int(value.Int64)
			}
		default:
			oru.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrgRoleUser.
// This includes values selected through modifiers, order, etc.
func (oru *OrgRoleUser) Value(name string) (ent.Value, error) {
	return oru.selectValues.Get(name)
}

// Update returns a builder for updating this OrgRoleUser.
// Note that you need to call OrgRoleUser.Unwrap() before calling this method if this OrgRoleUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (oru *OrgRoleUser) Update() *OrgRoleUserUpdateOne {
	return NewOrgRoleUserClient(oru.config).UpdateOne(oru)
}

// Unwrap unwraps the OrgRoleUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oru *OrgRoleUser) Unwrap() *OrgRoleUser {
	_tx, ok := oru.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgRoleUser is not a transactional entity")
	}
	oru.config.driver = _tx.drv
	return oru
}

// String implements the fmt.Stringer.
func (oru *OrgRoleUser) String() string {
	var builder strings.Builder
	builder.WriteString("OrgRoleUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oru.ID))
	builder.WriteString("org_role_id=")
	builder.WriteString(fmt.Sprintf("%v", oru.OrgRoleID))
	builder.WriteString(", ")
	builder.WriteString("org_user_id=")
	builder.WriteString(fmt.Sprintf("%v", oru.OrgUserID))
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", oru.OrgID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", oru.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// OrgRoleUsers is a parsable slice of OrgRoleUser.
type OrgRoleUsers []*OrgRoleUser
