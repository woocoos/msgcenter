// Code generated by ent, DO NOT EDIT.

package msgchannel

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/msgcenter/ent/predicate"
	"github.com/woocoos/msgcenter/pkg/profile"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldName, v))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldTenantID, v))
}

// Comments applies equality check predicate on the "comments" field. It's identical to CommentsEQ.
func Comments(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldComments, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotNull(FieldUpdatedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldContainsFold(FieldName, v))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v int) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLTE(FieldTenantID, v))
}

// ReceiverTypeEQ applies the EQ predicate on the "receiver_type" field.
func ReceiverTypeEQ(v profile.ReceiverType) predicate.MsgChannel {
	vc := v
	return predicate.MsgChannel(sql.FieldEQ(FieldReceiverType, vc))
}

// ReceiverTypeNEQ applies the NEQ predicate on the "receiver_type" field.
func ReceiverTypeNEQ(v profile.ReceiverType) predicate.MsgChannel {
	vc := v
	return predicate.MsgChannel(sql.FieldNEQ(FieldReceiverType, vc))
}

// ReceiverTypeIn applies the In predicate on the "receiver_type" field.
func ReceiverTypeIn(vs ...profile.ReceiverType) predicate.MsgChannel {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MsgChannel(sql.FieldIn(FieldReceiverType, v...))
}

// ReceiverTypeNotIn applies the NotIn predicate on the "receiver_type" field.
func ReceiverTypeNotIn(vs ...profile.ReceiverType) predicate.MsgChannel {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MsgChannel(sql.FieldNotIn(FieldReceiverType, v...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v typex.SimpleStatus) predicate.MsgChannel {
	vc := v
	return predicate.MsgChannel(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v typex.SimpleStatus) predicate.MsgChannel {
	vc := v
	return predicate.MsgChannel(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...typex.SimpleStatus) predicate.MsgChannel {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MsgChannel(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...typex.SimpleStatus) predicate.MsgChannel {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MsgChannel(sql.FieldNotIn(FieldStatus, v...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotNull(FieldStatus))
}

// ReceiverIsNil applies the IsNil predicate on the "receiver" field.
func ReceiverIsNil() predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIsNull(FieldReceiver))
}

// ReceiverNotNil applies the NotNil predicate on the "receiver" field.
func ReceiverNotNil() predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotNull(FieldReceiver))
}

// CommentsEQ applies the EQ predicate on the "comments" field.
func CommentsEQ(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEQ(FieldComments, v))
}

// CommentsNEQ applies the NEQ predicate on the "comments" field.
func CommentsNEQ(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNEQ(FieldComments, v))
}

// CommentsIn applies the In predicate on the "comments" field.
func CommentsIn(vs ...string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIn(FieldComments, vs...))
}

// CommentsNotIn applies the NotIn predicate on the "comments" field.
func CommentsNotIn(vs ...string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotIn(FieldComments, vs...))
}

// CommentsGT applies the GT predicate on the "comments" field.
func CommentsGT(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGT(FieldComments, v))
}

// CommentsGTE applies the GTE predicate on the "comments" field.
func CommentsGTE(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldGTE(FieldComments, v))
}

// CommentsLT applies the LT predicate on the "comments" field.
func CommentsLT(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLT(FieldComments, v))
}

// CommentsLTE applies the LTE predicate on the "comments" field.
func CommentsLTE(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldLTE(FieldComments, v))
}

// CommentsContains applies the Contains predicate on the "comments" field.
func CommentsContains(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldContains(FieldComments, v))
}

// CommentsHasPrefix applies the HasPrefix predicate on the "comments" field.
func CommentsHasPrefix(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldHasPrefix(FieldComments, v))
}

// CommentsHasSuffix applies the HasSuffix predicate on the "comments" field.
func CommentsHasSuffix(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldHasSuffix(FieldComments, v))
}

// CommentsIsNil applies the IsNil predicate on the "comments" field.
func CommentsIsNil() predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldIsNull(FieldComments))
}

// CommentsNotNil applies the NotNil predicate on the "comments" field.
func CommentsNotNil() predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldNotNull(FieldComments))
}

// CommentsEqualFold applies the EqualFold predicate on the "comments" field.
func CommentsEqualFold(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldEqualFold(FieldComments, v))
}

// CommentsContainsFold applies the ContainsFold predicate on the "comments" field.
func CommentsContainsFold(v string) predicate.MsgChannel {
	return predicate.MsgChannel(sql.FieldContainsFold(FieldComments, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MsgChannel) predicate.MsgChannel {
	return predicate.MsgChannel(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MsgChannel) predicate.MsgChannel {
	return predicate.MsgChannel(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MsgChannel) predicate.MsgChannel {
	return predicate.MsgChannel(sql.NotPredicates(p))
}
