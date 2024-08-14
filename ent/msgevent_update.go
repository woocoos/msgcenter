// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/predicate"
	"github.com/woocoos/msgcenter/pkg/profile"

	"github.com/woocoos/msgcenter/ent/internal"
)

// MsgEventUpdate is the builder for updating MsgEvent entities.
type MsgEventUpdate struct {
	config
	hooks    []Hook
	mutation *MsgEventMutation
}

// Where appends a list predicates to the MsgEventUpdate builder.
func (meu *MsgEventUpdate) Where(ps ...predicate.MsgEvent) *MsgEventUpdate {
	meu.mutation.Where(ps...)
	return meu
}

// SetUpdatedBy sets the "updated_by" field.
func (meu *MsgEventUpdate) SetUpdatedBy(i int) *MsgEventUpdate {
	meu.mutation.ResetUpdatedBy()
	meu.mutation.SetUpdatedBy(i)
	return meu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (meu *MsgEventUpdate) SetNillableUpdatedBy(i *int) *MsgEventUpdate {
	if i != nil {
		meu.SetUpdatedBy(*i)
	}
	return meu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (meu *MsgEventUpdate) AddUpdatedBy(i int) *MsgEventUpdate {
	meu.mutation.AddUpdatedBy(i)
	return meu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (meu *MsgEventUpdate) ClearUpdatedBy() *MsgEventUpdate {
	meu.mutation.ClearUpdatedBy()
	return meu
}

// SetUpdatedAt sets the "updated_at" field.
func (meu *MsgEventUpdate) SetUpdatedAt(t time.Time) *MsgEventUpdate {
	meu.mutation.SetUpdatedAt(t)
	return meu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (meu *MsgEventUpdate) SetNillableUpdatedAt(t *time.Time) *MsgEventUpdate {
	if t != nil {
		meu.SetUpdatedAt(*t)
	}
	return meu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (meu *MsgEventUpdate) ClearUpdatedAt() *MsgEventUpdate {
	meu.mutation.ClearUpdatedAt()
	return meu
}

// SetMsgTypeID sets the "msg_type_id" field.
func (meu *MsgEventUpdate) SetMsgTypeID(i int) *MsgEventUpdate {
	meu.mutation.SetMsgTypeID(i)
	return meu
}

// SetNillableMsgTypeID sets the "msg_type_id" field if the given value is not nil.
func (meu *MsgEventUpdate) SetNillableMsgTypeID(i *int) *MsgEventUpdate {
	if i != nil {
		meu.SetMsgTypeID(*i)
	}
	return meu
}

// SetName sets the "name" field.
func (meu *MsgEventUpdate) SetName(s string) *MsgEventUpdate {
	meu.mutation.SetName(s)
	return meu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (meu *MsgEventUpdate) SetNillableName(s *string) *MsgEventUpdate {
	if s != nil {
		meu.SetName(*s)
	}
	return meu
}

// SetStatus sets the "status" field.
func (meu *MsgEventUpdate) SetStatus(ts typex.SimpleStatus) *MsgEventUpdate {
	meu.mutation.SetStatus(ts)
	return meu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (meu *MsgEventUpdate) SetNillableStatus(ts *typex.SimpleStatus) *MsgEventUpdate {
	if ts != nil {
		meu.SetStatus(*ts)
	}
	return meu
}

// ClearStatus clears the value of the "status" field.
func (meu *MsgEventUpdate) ClearStatus() *MsgEventUpdate {
	meu.mutation.ClearStatus()
	return meu
}

// SetComments sets the "comments" field.
func (meu *MsgEventUpdate) SetComments(s string) *MsgEventUpdate {
	meu.mutation.SetComments(s)
	return meu
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (meu *MsgEventUpdate) SetNillableComments(s *string) *MsgEventUpdate {
	if s != nil {
		meu.SetComments(*s)
	}
	return meu
}

// ClearComments clears the value of the "comments" field.
func (meu *MsgEventUpdate) ClearComments() *MsgEventUpdate {
	meu.mutation.ClearComments()
	return meu
}

// SetRoute sets the "route" field.
func (meu *MsgEventUpdate) SetRoute(pr *profile.Route) *MsgEventUpdate {
	meu.mutation.SetRoute(pr)
	return meu
}

// ClearRoute clears the value of the "route" field.
func (meu *MsgEventUpdate) ClearRoute() *MsgEventUpdate {
	meu.mutation.ClearRoute()
	return meu
}

// SetModes sets the "modes" field.
func (meu *MsgEventUpdate) SetModes(s string) *MsgEventUpdate {
	meu.mutation.SetModes(s)
	return meu
}

// SetNillableModes sets the "modes" field if the given value is not nil.
func (meu *MsgEventUpdate) SetNillableModes(s *string) *MsgEventUpdate {
	if s != nil {
		meu.SetModes(*s)
	}
	return meu
}

// SetMsgType sets the "msg_type" edge to the MsgType entity.
func (meu *MsgEventUpdate) SetMsgType(m *MsgType) *MsgEventUpdate {
	return meu.SetMsgTypeID(m.ID)
}

// AddCustomerTemplateIDs adds the "customer_template" edge to the MsgTemplate entity by IDs.
func (meu *MsgEventUpdate) AddCustomerTemplateIDs(ids ...int) *MsgEventUpdate {
	meu.mutation.AddCustomerTemplateIDs(ids...)
	return meu
}

// AddCustomerTemplate adds the "customer_template" edges to the MsgTemplate entity.
func (meu *MsgEventUpdate) AddCustomerTemplate(m ...*MsgTemplate) *MsgEventUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return meu.AddCustomerTemplateIDs(ids...)
}

// Mutation returns the MsgEventMutation object of the builder.
func (meu *MsgEventUpdate) Mutation() *MsgEventMutation {
	return meu.mutation
}

// ClearMsgType clears the "msg_type" edge to the MsgType entity.
func (meu *MsgEventUpdate) ClearMsgType() *MsgEventUpdate {
	meu.mutation.ClearMsgType()
	return meu
}

// ClearCustomerTemplate clears all "customer_template" edges to the MsgTemplate entity.
func (meu *MsgEventUpdate) ClearCustomerTemplate() *MsgEventUpdate {
	meu.mutation.ClearCustomerTemplate()
	return meu
}

// RemoveCustomerTemplateIDs removes the "customer_template" edge to MsgTemplate entities by IDs.
func (meu *MsgEventUpdate) RemoveCustomerTemplateIDs(ids ...int) *MsgEventUpdate {
	meu.mutation.RemoveCustomerTemplateIDs(ids...)
	return meu
}

// RemoveCustomerTemplate removes "customer_template" edges to MsgTemplate entities.
func (meu *MsgEventUpdate) RemoveCustomerTemplate(m ...*MsgTemplate) *MsgEventUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return meu.RemoveCustomerTemplateIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (meu *MsgEventUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, meu.sqlSave, meu.mutation, meu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (meu *MsgEventUpdate) SaveX(ctx context.Context) int {
	affected, err := meu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (meu *MsgEventUpdate) Exec(ctx context.Context) error {
	_, err := meu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (meu *MsgEventUpdate) ExecX(ctx context.Context) {
	if err := meu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (meu *MsgEventUpdate) check() error {
	if v, ok := meu.mutation.Name(); ok {
		if err := msgevent.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MsgEvent.name": %w`, err)}
		}
	}
	if v, ok := meu.mutation.Status(); ok {
		if err := msgevent.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "MsgEvent.status": %w`, err)}
		}
	}
	if v, ok := meu.mutation.Route(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "route", err: fmt.Errorf(`ent: validator failed for field "MsgEvent.route": %w`, err)}
		}
	}
	if meu.mutation.MsgTypeCleared() && len(meu.mutation.MsgTypeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "MsgEvent.msg_type"`)
	}
	return nil
}

func (meu *MsgEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := meu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(msgevent.Table, msgevent.Columns, sqlgraph.NewFieldSpec(msgevent.FieldID, field.TypeInt))
	if ps := meu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := meu.mutation.UpdatedBy(); ok {
		_spec.SetField(msgevent.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := meu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(msgevent.FieldUpdatedBy, field.TypeInt, value)
	}
	if meu.mutation.UpdatedByCleared() {
		_spec.ClearField(msgevent.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := meu.mutation.UpdatedAt(); ok {
		_spec.SetField(msgevent.FieldUpdatedAt, field.TypeTime, value)
	}
	if meu.mutation.UpdatedAtCleared() {
		_spec.ClearField(msgevent.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := meu.mutation.Name(); ok {
		_spec.SetField(msgevent.FieldName, field.TypeString, value)
	}
	if value, ok := meu.mutation.Status(); ok {
		_spec.SetField(msgevent.FieldStatus, field.TypeEnum, value)
	}
	if meu.mutation.StatusCleared() {
		_spec.ClearField(msgevent.FieldStatus, field.TypeEnum)
	}
	if value, ok := meu.mutation.Comments(); ok {
		_spec.SetField(msgevent.FieldComments, field.TypeString, value)
	}
	if meu.mutation.CommentsCleared() {
		_spec.ClearField(msgevent.FieldComments, field.TypeString)
	}
	if value, ok := meu.mutation.Route(); ok {
		_spec.SetField(msgevent.FieldRoute, field.TypeJSON, value)
	}
	if meu.mutation.RouteCleared() {
		_spec.ClearField(msgevent.FieldRoute, field.TypeJSON)
	}
	if value, ok := meu.mutation.Modes(); ok {
		_spec.SetField(msgevent.FieldModes, field.TypeString, value)
	}
	if meu.mutation.MsgTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   msgevent.MsgTypeTable,
			Columns: []string{msgevent.MsgTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtype.FieldID, field.TypeInt),
			},
		}
		edge.Schema = meu.schemaConfig.MsgEvent
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := meu.mutation.MsgTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   msgevent.MsgTypeTable,
			Columns: []string{msgevent.MsgTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtype.FieldID, field.TypeInt),
			},
		}
		edge.Schema = meu.schemaConfig.MsgEvent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if meu.mutation.CustomerTemplateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   msgevent.CustomerTemplateTable,
			Columns: []string{msgevent.CustomerTemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtemplate.FieldID, field.TypeInt),
			},
		}
		edge.Schema = meu.schemaConfig.MsgTemplate
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := meu.mutation.RemovedCustomerTemplateIDs(); len(nodes) > 0 && !meu.mutation.CustomerTemplateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   msgevent.CustomerTemplateTable,
			Columns: []string{msgevent.CustomerTemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtemplate.FieldID, field.TypeInt),
			},
		}
		edge.Schema = meu.schemaConfig.MsgTemplate
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := meu.mutation.CustomerTemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   msgevent.CustomerTemplateTable,
			Columns: []string{msgevent.CustomerTemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtemplate.FieldID, field.TypeInt),
			},
		}
		edge.Schema = meu.schemaConfig.MsgTemplate
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = meu.schemaConfig.MsgEvent
	ctx = internal.NewSchemaConfigContext(ctx, meu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, meu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{msgevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	meu.mutation.done = true
	return n, nil
}

// MsgEventUpdateOne is the builder for updating a single MsgEvent entity.
type MsgEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MsgEventMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (meuo *MsgEventUpdateOne) SetUpdatedBy(i int) *MsgEventUpdateOne {
	meuo.mutation.ResetUpdatedBy()
	meuo.mutation.SetUpdatedBy(i)
	return meuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (meuo *MsgEventUpdateOne) SetNillableUpdatedBy(i *int) *MsgEventUpdateOne {
	if i != nil {
		meuo.SetUpdatedBy(*i)
	}
	return meuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (meuo *MsgEventUpdateOne) AddUpdatedBy(i int) *MsgEventUpdateOne {
	meuo.mutation.AddUpdatedBy(i)
	return meuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (meuo *MsgEventUpdateOne) ClearUpdatedBy() *MsgEventUpdateOne {
	meuo.mutation.ClearUpdatedBy()
	return meuo
}

// SetUpdatedAt sets the "updated_at" field.
func (meuo *MsgEventUpdateOne) SetUpdatedAt(t time.Time) *MsgEventUpdateOne {
	meuo.mutation.SetUpdatedAt(t)
	return meuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (meuo *MsgEventUpdateOne) SetNillableUpdatedAt(t *time.Time) *MsgEventUpdateOne {
	if t != nil {
		meuo.SetUpdatedAt(*t)
	}
	return meuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (meuo *MsgEventUpdateOne) ClearUpdatedAt() *MsgEventUpdateOne {
	meuo.mutation.ClearUpdatedAt()
	return meuo
}

// SetMsgTypeID sets the "msg_type_id" field.
func (meuo *MsgEventUpdateOne) SetMsgTypeID(i int) *MsgEventUpdateOne {
	meuo.mutation.SetMsgTypeID(i)
	return meuo
}

// SetNillableMsgTypeID sets the "msg_type_id" field if the given value is not nil.
func (meuo *MsgEventUpdateOne) SetNillableMsgTypeID(i *int) *MsgEventUpdateOne {
	if i != nil {
		meuo.SetMsgTypeID(*i)
	}
	return meuo
}

// SetName sets the "name" field.
func (meuo *MsgEventUpdateOne) SetName(s string) *MsgEventUpdateOne {
	meuo.mutation.SetName(s)
	return meuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (meuo *MsgEventUpdateOne) SetNillableName(s *string) *MsgEventUpdateOne {
	if s != nil {
		meuo.SetName(*s)
	}
	return meuo
}

// SetStatus sets the "status" field.
func (meuo *MsgEventUpdateOne) SetStatus(ts typex.SimpleStatus) *MsgEventUpdateOne {
	meuo.mutation.SetStatus(ts)
	return meuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (meuo *MsgEventUpdateOne) SetNillableStatus(ts *typex.SimpleStatus) *MsgEventUpdateOne {
	if ts != nil {
		meuo.SetStatus(*ts)
	}
	return meuo
}

// ClearStatus clears the value of the "status" field.
func (meuo *MsgEventUpdateOne) ClearStatus() *MsgEventUpdateOne {
	meuo.mutation.ClearStatus()
	return meuo
}

// SetComments sets the "comments" field.
func (meuo *MsgEventUpdateOne) SetComments(s string) *MsgEventUpdateOne {
	meuo.mutation.SetComments(s)
	return meuo
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (meuo *MsgEventUpdateOne) SetNillableComments(s *string) *MsgEventUpdateOne {
	if s != nil {
		meuo.SetComments(*s)
	}
	return meuo
}

// ClearComments clears the value of the "comments" field.
func (meuo *MsgEventUpdateOne) ClearComments() *MsgEventUpdateOne {
	meuo.mutation.ClearComments()
	return meuo
}

// SetRoute sets the "route" field.
func (meuo *MsgEventUpdateOne) SetRoute(pr *profile.Route) *MsgEventUpdateOne {
	meuo.mutation.SetRoute(pr)
	return meuo
}

// ClearRoute clears the value of the "route" field.
func (meuo *MsgEventUpdateOne) ClearRoute() *MsgEventUpdateOne {
	meuo.mutation.ClearRoute()
	return meuo
}

// SetModes sets the "modes" field.
func (meuo *MsgEventUpdateOne) SetModes(s string) *MsgEventUpdateOne {
	meuo.mutation.SetModes(s)
	return meuo
}

// SetNillableModes sets the "modes" field if the given value is not nil.
func (meuo *MsgEventUpdateOne) SetNillableModes(s *string) *MsgEventUpdateOne {
	if s != nil {
		meuo.SetModes(*s)
	}
	return meuo
}

// SetMsgType sets the "msg_type" edge to the MsgType entity.
func (meuo *MsgEventUpdateOne) SetMsgType(m *MsgType) *MsgEventUpdateOne {
	return meuo.SetMsgTypeID(m.ID)
}

// AddCustomerTemplateIDs adds the "customer_template" edge to the MsgTemplate entity by IDs.
func (meuo *MsgEventUpdateOne) AddCustomerTemplateIDs(ids ...int) *MsgEventUpdateOne {
	meuo.mutation.AddCustomerTemplateIDs(ids...)
	return meuo
}

// AddCustomerTemplate adds the "customer_template" edges to the MsgTemplate entity.
func (meuo *MsgEventUpdateOne) AddCustomerTemplate(m ...*MsgTemplate) *MsgEventUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return meuo.AddCustomerTemplateIDs(ids...)
}

// Mutation returns the MsgEventMutation object of the builder.
func (meuo *MsgEventUpdateOne) Mutation() *MsgEventMutation {
	return meuo.mutation
}

// ClearMsgType clears the "msg_type" edge to the MsgType entity.
func (meuo *MsgEventUpdateOne) ClearMsgType() *MsgEventUpdateOne {
	meuo.mutation.ClearMsgType()
	return meuo
}

// ClearCustomerTemplate clears all "customer_template" edges to the MsgTemplate entity.
func (meuo *MsgEventUpdateOne) ClearCustomerTemplate() *MsgEventUpdateOne {
	meuo.mutation.ClearCustomerTemplate()
	return meuo
}

// RemoveCustomerTemplateIDs removes the "customer_template" edge to MsgTemplate entities by IDs.
func (meuo *MsgEventUpdateOne) RemoveCustomerTemplateIDs(ids ...int) *MsgEventUpdateOne {
	meuo.mutation.RemoveCustomerTemplateIDs(ids...)
	return meuo
}

// RemoveCustomerTemplate removes "customer_template" edges to MsgTemplate entities.
func (meuo *MsgEventUpdateOne) RemoveCustomerTemplate(m ...*MsgTemplate) *MsgEventUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return meuo.RemoveCustomerTemplateIDs(ids...)
}

// Where appends a list predicates to the MsgEventUpdate builder.
func (meuo *MsgEventUpdateOne) Where(ps ...predicate.MsgEvent) *MsgEventUpdateOne {
	meuo.mutation.Where(ps...)
	return meuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (meuo *MsgEventUpdateOne) Select(field string, fields ...string) *MsgEventUpdateOne {
	meuo.fields = append([]string{field}, fields...)
	return meuo
}

// Save executes the query and returns the updated MsgEvent entity.
func (meuo *MsgEventUpdateOne) Save(ctx context.Context) (*MsgEvent, error) {
	return withHooks(ctx, meuo.sqlSave, meuo.mutation, meuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (meuo *MsgEventUpdateOne) SaveX(ctx context.Context) *MsgEvent {
	node, err := meuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (meuo *MsgEventUpdateOne) Exec(ctx context.Context) error {
	_, err := meuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (meuo *MsgEventUpdateOne) ExecX(ctx context.Context) {
	if err := meuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (meuo *MsgEventUpdateOne) check() error {
	if v, ok := meuo.mutation.Name(); ok {
		if err := msgevent.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MsgEvent.name": %w`, err)}
		}
	}
	if v, ok := meuo.mutation.Status(); ok {
		if err := msgevent.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "MsgEvent.status": %w`, err)}
		}
	}
	if v, ok := meuo.mutation.Route(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "route", err: fmt.Errorf(`ent: validator failed for field "MsgEvent.route": %w`, err)}
		}
	}
	if meuo.mutation.MsgTypeCleared() && len(meuo.mutation.MsgTypeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "MsgEvent.msg_type"`)
	}
	return nil
}

func (meuo *MsgEventUpdateOne) sqlSave(ctx context.Context) (_node *MsgEvent, err error) {
	if err := meuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(msgevent.Table, msgevent.Columns, sqlgraph.NewFieldSpec(msgevent.FieldID, field.TypeInt))
	id, ok := meuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MsgEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := meuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, msgevent.FieldID)
		for _, f := range fields {
			if !msgevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != msgevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := meuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := meuo.mutation.UpdatedBy(); ok {
		_spec.SetField(msgevent.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := meuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(msgevent.FieldUpdatedBy, field.TypeInt, value)
	}
	if meuo.mutation.UpdatedByCleared() {
		_spec.ClearField(msgevent.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := meuo.mutation.UpdatedAt(); ok {
		_spec.SetField(msgevent.FieldUpdatedAt, field.TypeTime, value)
	}
	if meuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(msgevent.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := meuo.mutation.Name(); ok {
		_spec.SetField(msgevent.FieldName, field.TypeString, value)
	}
	if value, ok := meuo.mutation.Status(); ok {
		_spec.SetField(msgevent.FieldStatus, field.TypeEnum, value)
	}
	if meuo.mutation.StatusCleared() {
		_spec.ClearField(msgevent.FieldStatus, field.TypeEnum)
	}
	if value, ok := meuo.mutation.Comments(); ok {
		_spec.SetField(msgevent.FieldComments, field.TypeString, value)
	}
	if meuo.mutation.CommentsCleared() {
		_spec.ClearField(msgevent.FieldComments, field.TypeString)
	}
	if value, ok := meuo.mutation.Route(); ok {
		_spec.SetField(msgevent.FieldRoute, field.TypeJSON, value)
	}
	if meuo.mutation.RouteCleared() {
		_spec.ClearField(msgevent.FieldRoute, field.TypeJSON)
	}
	if value, ok := meuo.mutation.Modes(); ok {
		_spec.SetField(msgevent.FieldModes, field.TypeString, value)
	}
	if meuo.mutation.MsgTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   msgevent.MsgTypeTable,
			Columns: []string{msgevent.MsgTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtype.FieldID, field.TypeInt),
			},
		}
		edge.Schema = meuo.schemaConfig.MsgEvent
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := meuo.mutation.MsgTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   msgevent.MsgTypeTable,
			Columns: []string{msgevent.MsgTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtype.FieldID, field.TypeInt),
			},
		}
		edge.Schema = meuo.schemaConfig.MsgEvent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if meuo.mutation.CustomerTemplateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   msgevent.CustomerTemplateTable,
			Columns: []string{msgevent.CustomerTemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtemplate.FieldID, field.TypeInt),
			},
		}
		edge.Schema = meuo.schemaConfig.MsgTemplate
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := meuo.mutation.RemovedCustomerTemplateIDs(); len(nodes) > 0 && !meuo.mutation.CustomerTemplateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   msgevent.CustomerTemplateTable,
			Columns: []string{msgevent.CustomerTemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtemplate.FieldID, field.TypeInt),
			},
		}
		edge.Schema = meuo.schemaConfig.MsgTemplate
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := meuo.mutation.CustomerTemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   msgevent.CustomerTemplateTable,
			Columns: []string{msgevent.CustomerTemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtemplate.FieldID, field.TypeInt),
			},
		}
		edge.Schema = meuo.schemaConfig.MsgTemplate
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = meuo.schemaConfig.MsgEvent
	ctx = internal.NewSchemaConfigContext(ctx, meuo.schemaConfig)
	_node = &MsgEvent{config: meuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, meuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{msgevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	meuo.mutation.done = true
	return _node, nil
}
