// Code generated by ent, DO NOT EDIT.

package msgalert

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/msgcenter/ent/predicate"
	"github.com/woocoos/msgcenter/pkg/alert"

	"github.com/woocoos/msgcenter/ent/internal"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLTE(FieldID, id))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldTenantID, v))
}

// StartsAt applies equality check predicate on the "starts_at" field. It's identical to StartsAtEQ.
func StartsAt(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldStartsAt, v))
}

// EndsAt applies equality check predicate on the "ends_at" field. It's identical to EndsAtEQ.
func EndsAt(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldEndsAt, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldURL, v))
}

// Timeout applies equality check predicate on the "timeout" field. It's identical to TimeoutEQ.
func Timeout(v bool) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldTimeout, v))
}

// Fingerprint applies equality check predicate on the "fingerprint" field. It's identical to FingerprintEQ.
func Fingerprint(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldFingerprint, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldUpdatedAt, v))
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v bool) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldDeleted, v))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v int) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLTE(FieldTenantID, v))
}

// LabelsIsNil applies the IsNil predicate on the "labels" field.
func LabelsIsNil() predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIsNull(FieldLabels))
}

// LabelsNotNil applies the NotNil predicate on the "labels" field.
func LabelsNotNil() predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotNull(FieldLabels))
}

// AnnotationsIsNil applies the IsNil predicate on the "annotations" field.
func AnnotationsIsNil() predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIsNull(FieldAnnotations))
}

// AnnotationsNotNil applies the NotNil predicate on the "annotations" field.
func AnnotationsNotNil() predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotNull(FieldAnnotations))
}

// StartsAtEQ applies the EQ predicate on the "starts_at" field.
func StartsAtEQ(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldStartsAt, v))
}

// StartsAtNEQ applies the NEQ predicate on the "starts_at" field.
func StartsAtNEQ(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNEQ(FieldStartsAt, v))
}

// StartsAtIn applies the In predicate on the "starts_at" field.
func StartsAtIn(vs ...time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIn(FieldStartsAt, vs...))
}

// StartsAtNotIn applies the NotIn predicate on the "starts_at" field.
func StartsAtNotIn(vs ...time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotIn(FieldStartsAt, vs...))
}

// StartsAtGT applies the GT predicate on the "starts_at" field.
func StartsAtGT(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGT(FieldStartsAt, v))
}

// StartsAtGTE applies the GTE predicate on the "starts_at" field.
func StartsAtGTE(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGTE(FieldStartsAt, v))
}

// StartsAtLT applies the LT predicate on the "starts_at" field.
func StartsAtLT(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLT(FieldStartsAt, v))
}

// StartsAtLTE applies the LTE predicate on the "starts_at" field.
func StartsAtLTE(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLTE(FieldStartsAt, v))
}

// EndsAtEQ applies the EQ predicate on the "ends_at" field.
func EndsAtEQ(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldEndsAt, v))
}

// EndsAtNEQ applies the NEQ predicate on the "ends_at" field.
func EndsAtNEQ(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNEQ(FieldEndsAt, v))
}

// EndsAtIn applies the In predicate on the "ends_at" field.
func EndsAtIn(vs ...time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIn(FieldEndsAt, vs...))
}

// EndsAtNotIn applies the NotIn predicate on the "ends_at" field.
func EndsAtNotIn(vs ...time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotIn(FieldEndsAt, vs...))
}

// EndsAtGT applies the GT predicate on the "ends_at" field.
func EndsAtGT(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGT(FieldEndsAt, v))
}

// EndsAtGTE applies the GTE predicate on the "ends_at" field.
func EndsAtGTE(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGTE(FieldEndsAt, v))
}

// EndsAtLT applies the LT predicate on the "ends_at" field.
func EndsAtLT(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLT(FieldEndsAt, v))
}

// EndsAtLTE applies the LTE predicate on the "ends_at" field.
func EndsAtLTE(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLTE(FieldEndsAt, v))
}

// EndsAtIsNil applies the IsNil predicate on the "ends_at" field.
func EndsAtIsNil() predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIsNull(FieldEndsAt))
}

// EndsAtNotNil applies the NotNil predicate on the "ends_at" field.
func EndsAtNotNil() predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotNull(FieldEndsAt))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldHasSuffix(FieldURL, v))
}

// URLIsNil applies the IsNil predicate on the "url" field.
func URLIsNil() predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIsNull(FieldURL))
}

// URLNotNil applies the NotNil predicate on the "url" field.
func URLNotNil() predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotNull(FieldURL))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldContainsFold(FieldURL, v))
}

// TimeoutEQ applies the EQ predicate on the "timeout" field.
func TimeoutEQ(v bool) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldTimeout, v))
}

// TimeoutNEQ applies the NEQ predicate on the "timeout" field.
func TimeoutNEQ(v bool) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNEQ(FieldTimeout, v))
}

// FingerprintEQ applies the EQ predicate on the "fingerprint" field.
func FingerprintEQ(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldFingerprint, v))
}

// FingerprintNEQ applies the NEQ predicate on the "fingerprint" field.
func FingerprintNEQ(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNEQ(FieldFingerprint, v))
}

// FingerprintIn applies the In predicate on the "fingerprint" field.
func FingerprintIn(vs ...string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIn(FieldFingerprint, vs...))
}

// FingerprintNotIn applies the NotIn predicate on the "fingerprint" field.
func FingerprintNotIn(vs ...string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotIn(FieldFingerprint, vs...))
}

// FingerprintGT applies the GT predicate on the "fingerprint" field.
func FingerprintGT(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGT(FieldFingerprint, v))
}

// FingerprintGTE applies the GTE predicate on the "fingerprint" field.
func FingerprintGTE(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGTE(FieldFingerprint, v))
}

// FingerprintLT applies the LT predicate on the "fingerprint" field.
func FingerprintLT(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLT(FieldFingerprint, v))
}

// FingerprintLTE applies the LTE predicate on the "fingerprint" field.
func FingerprintLTE(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLTE(FieldFingerprint, v))
}

// FingerprintContains applies the Contains predicate on the "fingerprint" field.
func FingerprintContains(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldContains(FieldFingerprint, v))
}

// FingerprintHasPrefix applies the HasPrefix predicate on the "fingerprint" field.
func FingerprintHasPrefix(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldHasPrefix(FieldFingerprint, v))
}

// FingerprintHasSuffix applies the HasSuffix predicate on the "fingerprint" field.
func FingerprintHasSuffix(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldHasSuffix(FieldFingerprint, v))
}

// FingerprintEqualFold applies the EqualFold predicate on the "fingerprint" field.
func FingerprintEqualFold(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEqualFold(FieldFingerprint, v))
}

// FingerprintContainsFold applies the ContainsFold predicate on the "fingerprint" field.
func FingerprintContainsFold(v string) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldContainsFold(FieldFingerprint, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v alert.AlertStatus) predicate.MsgAlert {
	vc := v
	return predicate.MsgAlert(sql.FieldEQ(FieldState, vc))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v alert.AlertStatus) predicate.MsgAlert {
	vc := v
	return predicate.MsgAlert(sql.FieldNEQ(FieldState, vc))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...alert.AlertStatus) predicate.MsgAlert {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MsgAlert(sql.FieldIn(FieldState, v...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...alert.AlertStatus) predicate.MsgAlert {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MsgAlert(sql.FieldNotIn(FieldState, v...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v bool) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldEQ(FieldDeleted, v))
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v bool) predicate.MsgAlert {
	return predicate.MsgAlert(sql.FieldNEQ(FieldDeleted, v))
}

// HasNlog applies the HasEdge predicate on the "nlog" edge.
func HasNlog() predicate.MsgAlert {
	return predicate.MsgAlert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, NlogTable, NlogPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Nlog
		step.Edge.Schema = schemaConfig.NlogAlert
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNlogWith applies the HasEdge predicate on the "nlog" edge with a given conditions (other predicates).
func HasNlogWith(preds ...predicate.Nlog) predicate.MsgAlert {
	return predicate.MsgAlert(func(s *sql.Selector) {
		step := newNlogStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Nlog
		step.Edge.Schema = schemaConfig.NlogAlert
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNlogAlerts applies the HasEdge predicate on the "nlog_alerts" edge.
func HasNlogAlerts() predicate.MsgAlert {
	return predicate.MsgAlert(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, NlogAlertsTable, NlogAlertsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.NlogAlert
		step.Edge.Schema = schemaConfig.NlogAlert
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNlogAlertsWith applies the HasEdge predicate on the "nlog_alerts" edge with a given conditions (other predicates).
func HasNlogAlertsWith(preds ...predicate.NlogAlert) predicate.MsgAlert {
	return predicate.MsgAlert(func(s *sql.Selector) {
		step := newNlogAlertsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.NlogAlert
		step.Edge.Schema = schemaConfig.NlogAlert
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MsgAlert) predicate.MsgAlert {
	return predicate.MsgAlert(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MsgAlert) predicate.MsgAlert {
	return predicate.MsgAlert(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MsgAlert) predicate.MsgAlert {
	return predicate.MsgAlert(sql.NotPredicates(p))
}
