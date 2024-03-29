// Code generated by ent, DO NOT EDIT.

package orgroleuser

import (
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/msgcenter/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLTE(FieldID, id))
}

// OrgRoleID applies equality check predicate on the "org_role_id" field. It's identical to OrgRoleIDEQ.
func OrgRoleID(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldOrgRoleID, v))
}

// OrgUserID applies equality check predicate on the "org_user_id" field. It's identical to OrgUserIDEQ.
func OrgUserID(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldOrgUserID, v))
}

// OrgID applies equality check predicate on the "org_id" field. It's identical to OrgIDEQ.
func OrgID(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldOrgID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldUserID, v))
}

// OrgRoleIDEQ applies the EQ predicate on the "org_role_id" field.
func OrgRoleIDEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldOrgRoleID, v))
}

// OrgRoleIDNEQ applies the NEQ predicate on the "org_role_id" field.
func OrgRoleIDNEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldOrgRoleID, v))
}

// OrgRoleIDIn applies the In predicate on the "org_role_id" field.
func OrgRoleIDIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldOrgRoleID, vs...))
}

// OrgRoleIDNotIn applies the NotIn predicate on the "org_role_id" field.
func OrgRoleIDNotIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldOrgRoleID, vs...))
}

// OrgRoleIDGT applies the GT predicate on the "org_role_id" field.
func OrgRoleIDGT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGT(FieldOrgRoleID, v))
}

// OrgRoleIDGTE applies the GTE predicate on the "org_role_id" field.
func OrgRoleIDGTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGTE(FieldOrgRoleID, v))
}

// OrgRoleIDLT applies the LT predicate on the "org_role_id" field.
func OrgRoleIDLT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLT(FieldOrgRoleID, v))
}

// OrgRoleIDLTE applies the LTE predicate on the "org_role_id" field.
func OrgRoleIDLTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLTE(FieldOrgRoleID, v))
}

// OrgUserIDEQ applies the EQ predicate on the "org_user_id" field.
func OrgUserIDEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldOrgUserID, v))
}

// OrgUserIDNEQ applies the NEQ predicate on the "org_user_id" field.
func OrgUserIDNEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldOrgUserID, v))
}

// OrgUserIDIn applies the In predicate on the "org_user_id" field.
func OrgUserIDIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldOrgUserID, vs...))
}

// OrgUserIDNotIn applies the NotIn predicate on the "org_user_id" field.
func OrgUserIDNotIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldOrgUserID, vs...))
}

// OrgUserIDGT applies the GT predicate on the "org_user_id" field.
func OrgUserIDGT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGT(FieldOrgUserID, v))
}

// OrgUserIDGTE applies the GTE predicate on the "org_user_id" field.
func OrgUserIDGTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGTE(FieldOrgUserID, v))
}

// OrgUserIDLT applies the LT predicate on the "org_user_id" field.
func OrgUserIDLT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLT(FieldOrgUserID, v))
}

// OrgUserIDLTE applies the LTE predicate on the "org_user_id" field.
func OrgUserIDLTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLTE(FieldOrgUserID, v))
}

// OrgIDEQ applies the EQ predicate on the "org_id" field.
func OrgIDEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldOrgID, v))
}

// OrgIDNEQ applies the NEQ predicate on the "org_id" field.
func OrgIDNEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldOrgID, v))
}

// OrgIDIn applies the In predicate on the "org_id" field.
func OrgIDIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldOrgID, vs...))
}

// OrgIDNotIn applies the NotIn predicate on the "org_id" field.
func OrgIDNotIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldOrgID, vs...))
}

// OrgIDGT applies the GT predicate on the "org_id" field.
func OrgIDGT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGT(FieldOrgID, v))
}

// OrgIDGTE applies the GTE predicate on the "org_id" field.
func OrgIDGTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGTE(FieldOrgID, v))
}

// OrgIDLT applies the LT predicate on the "org_id" field.
func OrgIDLT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLT(FieldOrgID, v))
}

// OrgIDLTE applies the LTE predicate on the "org_id" field.
func OrgIDLTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLTE(FieldOrgID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLTE(FieldUserID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrgRoleUser) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrgRoleUser) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrgRoleUser) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.NotPredicates(p))
}
