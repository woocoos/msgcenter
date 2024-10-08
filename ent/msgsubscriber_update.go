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
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/predicate"
	"github.com/woocoos/msgcenter/ent/user"

	"github.com/woocoos/msgcenter/ent/internal"
)

// MsgSubscriberUpdate is the builder for updating MsgSubscriber entities.
type MsgSubscriberUpdate struct {
	config
	hooks    []Hook
	mutation *MsgSubscriberMutation
}

// Where appends a list predicates to the MsgSubscriberUpdate builder.
func (msu *MsgSubscriberUpdate) Where(ps ...predicate.MsgSubscriber) *MsgSubscriberUpdate {
	msu.mutation.Where(ps...)
	return msu
}

// SetUpdatedBy sets the "updated_by" field.
func (msu *MsgSubscriberUpdate) SetUpdatedBy(i int) *MsgSubscriberUpdate {
	msu.mutation.ResetUpdatedBy()
	msu.mutation.SetUpdatedBy(i)
	return msu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (msu *MsgSubscriberUpdate) SetNillableUpdatedBy(i *int) *MsgSubscriberUpdate {
	if i != nil {
		msu.SetUpdatedBy(*i)
	}
	return msu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (msu *MsgSubscriberUpdate) AddUpdatedBy(i int) *MsgSubscriberUpdate {
	msu.mutation.AddUpdatedBy(i)
	return msu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (msu *MsgSubscriberUpdate) ClearUpdatedBy() *MsgSubscriberUpdate {
	msu.mutation.ClearUpdatedBy()
	return msu
}

// SetUpdatedAt sets the "updated_at" field.
func (msu *MsgSubscriberUpdate) SetUpdatedAt(t time.Time) *MsgSubscriberUpdate {
	msu.mutation.SetUpdatedAt(t)
	return msu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (msu *MsgSubscriberUpdate) SetNillableUpdatedAt(t *time.Time) *MsgSubscriberUpdate {
	if t != nil {
		msu.SetUpdatedAt(*t)
	}
	return msu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (msu *MsgSubscriberUpdate) ClearUpdatedAt() *MsgSubscriberUpdate {
	msu.mutation.ClearUpdatedAt()
	return msu
}

// SetMsgTypeID sets the "msg_type_id" field.
func (msu *MsgSubscriberUpdate) SetMsgTypeID(i int) *MsgSubscriberUpdate {
	msu.mutation.SetMsgTypeID(i)
	return msu
}

// SetNillableMsgTypeID sets the "msg_type_id" field if the given value is not nil.
func (msu *MsgSubscriberUpdate) SetNillableMsgTypeID(i *int) *MsgSubscriberUpdate {
	if i != nil {
		msu.SetMsgTypeID(*i)
	}
	return msu
}

// SetTenantID sets the "tenant_id" field.
func (msu *MsgSubscriberUpdate) SetTenantID(i int) *MsgSubscriberUpdate {
	msu.mutation.ResetTenantID()
	msu.mutation.SetTenantID(i)
	return msu
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (msu *MsgSubscriberUpdate) SetNillableTenantID(i *int) *MsgSubscriberUpdate {
	if i != nil {
		msu.SetTenantID(*i)
	}
	return msu
}

// AddTenantID adds i to the "tenant_id" field.
func (msu *MsgSubscriberUpdate) AddTenantID(i int) *MsgSubscriberUpdate {
	msu.mutation.AddTenantID(i)
	return msu
}

// SetUserID sets the "user_id" field.
func (msu *MsgSubscriberUpdate) SetUserID(i int) *MsgSubscriberUpdate {
	msu.mutation.SetUserID(i)
	return msu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (msu *MsgSubscriberUpdate) SetNillableUserID(i *int) *MsgSubscriberUpdate {
	if i != nil {
		msu.SetUserID(*i)
	}
	return msu
}

// ClearUserID clears the value of the "user_id" field.
func (msu *MsgSubscriberUpdate) ClearUserID() *MsgSubscriberUpdate {
	msu.mutation.ClearUserID()
	return msu
}

// SetOrgRoleID sets the "org_role_id" field.
func (msu *MsgSubscriberUpdate) SetOrgRoleID(i int) *MsgSubscriberUpdate {
	msu.mutation.ResetOrgRoleID()
	msu.mutation.SetOrgRoleID(i)
	return msu
}

// SetNillableOrgRoleID sets the "org_role_id" field if the given value is not nil.
func (msu *MsgSubscriberUpdate) SetNillableOrgRoleID(i *int) *MsgSubscriberUpdate {
	if i != nil {
		msu.SetOrgRoleID(*i)
	}
	return msu
}

// AddOrgRoleID adds i to the "org_role_id" field.
func (msu *MsgSubscriberUpdate) AddOrgRoleID(i int) *MsgSubscriberUpdate {
	msu.mutation.AddOrgRoleID(i)
	return msu
}

// ClearOrgRoleID clears the value of the "org_role_id" field.
func (msu *MsgSubscriberUpdate) ClearOrgRoleID() *MsgSubscriberUpdate {
	msu.mutation.ClearOrgRoleID()
	return msu
}

// SetExclude sets the "exclude" field.
func (msu *MsgSubscriberUpdate) SetExclude(b bool) *MsgSubscriberUpdate {
	msu.mutation.SetExclude(b)
	return msu
}

// SetNillableExclude sets the "exclude" field if the given value is not nil.
func (msu *MsgSubscriberUpdate) SetNillableExclude(b *bool) *MsgSubscriberUpdate {
	if b != nil {
		msu.SetExclude(*b)
	}
	return msu
}

// ClearExclude clears the value of the "exclude" field.
func (msu *MsgSubscriberUpdate) ClearExclude() *MsgSubscriberUpdate {
	msu.mutation.ClearExclude()
	return msu
}

// SetMsgType sets the "msg_type" edge to the MsgType entity.
func (msu *MsgSubscriberUpdate) SetMsgType(m *MsgType) *MsgSubscriberUpdate {
	return msu.SetMsgTypeID(m.ID)
}

// SetUser sets the "user" edge to the User entity.
func (msu *MsgSubscriberUpdate) SetUser(u *User) *MsgSubscriberUpdate {
	return msu.SetUserID(u.ID)
}

// Mutation returns the MsgSubscriberMutation object of the builder.
func (msu *MsgSubscriberUpdate) Mutation() *MsgSubscriberMutation {
	return msu.mutation
}

// ClearMsgType clears the "msg_type" edge to the MsgType entity.
func (msu *MsgSubscriberUpdate) ClearMsgType() *MsgSubscriberUpdate {
	msu.mutation.ClearMsgType()
	return msu
}

// ClearUser clears the "user" edge to the User entity.
func (msu *MsgSubscriberUpdate) ClearUser() *MsgSubscriberUpdate {
	msu.mutation.ClearUser()
	return msu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (msu *MsgSubscriberUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, msu.sqlSave, msu.mutation, msu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (msu *MsgSubscriberUpdate) SaveX(ctx context.Context) int {
	affected, err := msu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (msu *MsgSubscriberUpdate) Exec(ctx context.Context) error {
	_, err := msu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msu *MsgSubscriberUpdate) ExecX(ctx context.Context) {
	if err := msu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msu *MsgSubscriberUpdate) check() error {
	if msu.mutation.MsgTypeCleared() && len(msu.mutation.MsgTypeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "MsgSubscriber.msg_type"`)
	}
	return nil
}

func (msu *MsgSubscriberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := msu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(msgsubscriber.Table, msgsubscriber.Columns, sqlgraph.NewFieldSpec(msgsubscriber.FieldID, field.TypeInt))
	if ps := msu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := msu.mutation.UpdatedBy(); ok {
		_spec.SetField(msgsubscriber.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := msu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(msgsubscriber.FieldUpdatedBy, field.TypeInt, value)
	}
	if msu.mutation.UpdatedByCleared() {
		_spec.ClearField(msgsubscriber.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := msu.mutation.UpdatedAt(); ok {
		_spec.SetField(msgsubscriber.FieldUpdatedAt, field.TypeTime, value)
	}
	if msu.mutation.UpdatedAtCleared() {
		_spec.ClearField(msgsubscriber.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := msu.mutation.TenantID(); ok {
		_spec.SetField(msgsubscriber.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := msu.mutation.AddedTenantID(); ok {
		_spec.AddField(msgsubscriber.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := msu.mutation.OrgRoleID(); ok {
		_spec.SetField(msgsubscriber.FieldOrgRoleID, field.TypeInt, value)
	}
	if value, ok := msu.mutation.AddedOrgRoleID(); ok {
		_spec.AddField(msgsubscriber.FieldOrgRoleID, field.TypeInt, value)
	}
	if msu.mutation.OrgRoleIDCleared() {
		_spec.ClearField(msgsubscriber.FieldOrgRoleID, field.TypeInt)
	}
	if value, ok := msu.mutation.Exclude(); ok {
		_spec.SetField(msgsubscriber.FieldExclude, field.TypeBool, value)
	}
	if msu.mutation.ExcludeCleared() {
		_spec.ClearField(msgsubscriber.FieldExclude, field.TypeBool)
	}
	if msu.mutation.MsgTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   msgsubscriber.MsgTypeTable,
			Columns: []string{msgsubscriber.MsgTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtype.FieldID, field.TypeInt),
			},
		}
		edge.Schema = msu.schemaConfig.MsgSubscriber
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msu.mutation.MsgTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   msgsubscriber.MsgTypeTable,
			Columns: []string{msgsubscriber.MsgTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtype.FieldID, field.TypeInt),
			},
		}
		edge.Schema = msu.schemaConfig.MsgSubscriber
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if msu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   msgsubscriber.UserTable,
			Columns: []string{msgsubscriber.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = msu.schemaConfig.MsgSubscriber
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   msgsubscriber.UserTable,
			Columns: []string{msgsubscriber.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = msu.schemaConfig.MsgSubscriber
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = msu.schemaConfig.MsgSubscriber
	ctx = internal.NewSchemaConfigContext(ctx, msu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, msu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{msgsubscriber.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	msu.mutation.done = true
	return n, nil
}

// MsgSubscriberUpdateOne is the builder for updating a single MsgSubscriber entity.
type MsgSubscriberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MsgSubscriberMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (msuo *MsgSubscriberUpdateOne) SetUpdatedBy(i int) *MsgSubscriberUpdateOne {
	msuo.mutation.ResetUpdatedBy()
	msuo.mutation.SetUpdatedBy(i)
	return msuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (msuo *MsgSubscriberUpdateOne) SetNillableUpdatedBy(i *int) *MsgSubscriberUpdateOne {
	if i != nil {
		msuo.SetUpdatedBy(*i)
	}
	return msuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (msuo *MsgSubscriberUpdateOne) AddUpdatedBy(i int) *MsgSubscriberUpdateOne {
	msuo.mutation.AddUpdatedBy(i)
	return msuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (msuo *MsgSubscriberUpdateOne) ClearUpdatedBy() *MsgSubscriberUpdateOne {
	msuo.mutation.ClearUpdatedBy()
	return msuo
}

// SetUpdatedAt sets the "updated_at" field.
func (msuo *MsgSubscriberUpdateOne) SetUpdatedAt(t time.Time) *MsgSubscriberUpdateOne {
	msuo.mutation.SetUpdatedAt(t)
	return msuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (msuo *MsgSubscriberUpdateOne) SetNillableUpdatedAt(t *time.Time) *MsgSubscriberUpdateOne {
	if t != nil {
		msuo.SetUpdatedAt(*t)
	}
	return msuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (msuo *MsgSubscriberUpdateOne) ClearUpdatedAt() *MsgSubscriberUpdateOne {
	msuo.mutation.ClearUpdatedAt()
	return msuo
}

// SetMsgTypeID sets the "msg_type_id" field.
func (msuo *MsgSubscriberUpdateOne) SetMsgTypeID(i int) *MsgSubscriberUpdateOne {
	msuo.mutation.SetMsgTypeID(i)
	return msuo
}

// SetNillableMsgTypeID sets the "msg_type_id" field if the given value is not nil.
func (msuo *MsgSubscriberUpdateOne) SetNillableMsgTypeID(i *int) *MsgSubscriberUpdateOne {
	if i != nil {
		msuo.SetMsgTypeID(*i)
	}
	return msuo
}

// SetTenantID sets the "tenant_id" field.
func (msuo *MsgSubscriberUpdateOne) SetTenantID(i int) *MsgSubscriberUpdateOne {
	msuo.mutation.ResetTenantID()
	msuo.mutation.SetTenantID(i)
	return msuo
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (msuo *MsgSubscriberUpdateOne) SetNillableTenantID(i *int) *MsgSubscriberUpdateOne {
	if i != nil {
		msuo.SetTenantID(*i)
	}
	return msuo
}

// AddTenantID adds i to the "tenant_id" field.
func (msuo *MsgSubscriberUpdateOne) AddTenantID(i int) *MsgSubscriberUpdateOne {
	msuo.mutation.AddTenantID(i)
	return msuo
}

// SetUserID sets the "user_id" field.
func (msuo *MsgSubscriberUpdateOne) SetUserID(i int) *MsgSubscriberUpdateOne {
	msuo.mutation.SetUserID(i)
	return msuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (msuo *MsgSubscriberUpdateOne) SetNillableUserID(i *int) *MsgSubscriberUpdateOne {
	if i != nil {
		msuo.SetUserID(*i)
	}
	return msuo
}

// ClearUserID clears the value of the "user_id" field.
func (msuo *MsgSubscriberUpdateOne) ClearUserID() *MsgSubscriberUpdateOne {
	msuo.mutation.ClearUserID()
	return msuo
}

// SetOrgRoleID sets the "org_role_id" field.
func (msuo *MsgSubscriberUpdateOne) SetOrgRoleID(i int) *MsgSubscriberUpdateOne {
	msuo.mutation.ResetOrgRoleID()
	msuo.mutation.SetOrgRoleID(i)
	return msuo
}

// SetNillableOrgRoleID sets the "org_role_id" field if the given value is not nil.
func (msuo *MsgSubscriberUpdateOne) SetNillableOrgRoleID(i *int) *MsgSubscriberUpdateOne {
	if i != nil {
		msuo.SetOrgRoleID(*i)
	}
	return msuo
}

// AddOrgRoleID adds i to the "org_role_id" field.
func (msuo *MsgSubscriberUpdateOne) AddOrgRoleID(i int) *MsgSubscriberUpdateOne {
	msuo.mutation.AddOrgRoleID(i)
	return msuo
}

// ClearOrgRoleID clears the value of the "org_role_id" field.
func (msuo *MsgSubscriberUpdateOne) ClearOrgRoleID() *MsgSubscriberUpdateOne {
	msuo.mutation.ClearOrgRoleID()
	return msuo
}

// SetExclude sets the "exclude" field.
func (msuo *MsgSubscriberUpdateOne) SetExclude(b bool) *MsgSubscriberUpdateOne {
	msuo.mutation.SetExclude(b)
	return msuo
}

// SetNillableExclude sets the "exclude" field if the given value is not nil.
func (msuo *MsgSubscriberUpdateOne) SetNillableExclude(b *bool) *MsgSubscriberUpdateOne {
	if b != nil {
		msuo.SetExclude(*b)
	}
	return msuo
}

// ClearExclude clears the value of the "exclude" field.
func (msuo *MsgSubscriberUpdateOne) ClearExclude() *MsgSubscriberUpdateOne {
	msuo.mutation.ClearExclude()
	return msuo
}

// SetMsgType sets the "msg_type" edge to the MsgType entity.
func (msuo *MsgSubscriberUpdateOne) SetMsgType(m *MsgType) *MsgSubscriberUpdateOne {
	return msuo.SetMsgTypeID(m.ID)
}

// SetUser sets the "user" edge to the User entity.
func (msuo *MsgSubscriberUpdateOne) SetUser(u *User) *MsgSubscriberUpdateOne {
	return msuo.SetUserID(u.ID)
}

// Mutation returns the MsgSubscriberMutation object of the builder.
func (msuo *MsgSubscriberUpdateOne) Mutation() *MsgSubscriberMutation {
	return msuo.mutation
}

// ClearMsgType clears the "msg_type" edge to the MsgType entity.
func (msuo *MsgSubscriberUpdateOne) ClearMsgType() *MsgSubscriberUpdateOne {
	msuo.mutation.ClearMsgType()
	return msuo
}

// ClearUser clears the "user" edge to the User entity.
func (msuo *MsgSubscriberUpdateOne) ClearUser() *MsgSubscriberUpdateOne {
	msuo.mutation.ClearUser()
	return msuo
}

// Where appends a list predicates to the MsgSubscriberUpdate builder.
func (msuo *MsgSubscriberUpdateOne) Where(ps ...predicate.MsgSubscriber) *MsgSubscriberUpdateOne {
	msuo.mutation.Where(ps...)
	return msuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (msuo *MsgSubscriberUpdateOne) Select(field string, fields ...string) *MsgSubscriberUpdateOne {
	msuo.fields = append([]string{field}, fields...)
	return msuo
}

// Save executes the query and returns the updated MsgSubscriber entity.
func (msuo *MsgSubscriberUpdateOne) Save(ctx context.Context) (*MsgSubscriber, error) {
	return withHooks(ctx, msuo.sqlSave, msuo.mutation, msuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (msuo *MsgSubscriberUpdateOne) SaveX(ctx context.Context) *MsgSubscriber {
	node, err := msuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (msuo *MsgSubscriberUpdateOne) Exec(ctx context.Context) error {
	_, err := msuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msuo *MsgSubscriberUpdateOne) ExecX(ctx context.Context) {
	if err := msuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msuo *MsgSubscriberUpdateOne) check() error {
	if msuo.mutation.MsgTypeCleared() && len(msuo.mutation.MsgTypeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "MsgSubscriber.msg_type"`)
	}
	return nil
}

func (msuo *MsgSubscriberUpdateOne) sqlSave(ctx context.Context) (_node *MsgSubscriber, err error) {
	if err := msuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(msgsubscriber.Table, msgsubscriber.Columns, sqlgraph.NewFieldSpec(msgsubscriber.FieldID, field.TypeInt))
	id, ok := msuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MsgSubscriber.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := msuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, msgsubscriber.FieldID)
		for _, f := range fields {
			if !msgsubscriber.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != msgsubscriber.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := msuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := msuo.mutation.UpdatedBy(); ok {
		_spec.SetField(msgsubscriber.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := msuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(msgsubscriber.FieldUpdatedBy, field.TypeInt, value)
	}
	if msuo.mutation.UpdatedByCleared() {
		_spec.ClearField(msgsubscriber.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := msuo.mutation.UpdatedAt(); ok {
		_spec.SetField(msgsubscriber.FieldUpdatedAt, field.TypeTime, value)
	}
	if msuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(msgsubscriber.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := msuo.mutation.TenantID(); ok {
		_spec.SetField(msgsubscriber.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := msuo.mutation.AddedTenantID(); ok {
		_spec.AddField(msgsubscriber.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := msuo.mutation.OrgRoleID(); ok {
		_spec.SetField(msgsubscriber.FieldOrgRoleID, field.TypeInt, value)
	}
	if value, ok := msuo.mutation.AddedOrgRoleID(); ok {
		_spec.AddField(msgsubscriber.FieldOrgRoleID, field.TypeInt, value)
	}
	if msuo.mutation.OrgRoleIDCleared() {
		_spec.ClearField(msgsubscriber.FieldOrgRoleID, field.TypeInt)
	}
	if value, ok := msuo.mutation.Exclude(); ok {
		_spec.SetField(msgsubscriber.FieldExclude, field.TypeBool, value)
	}
	if msuo.mutation.ExcludeCleared() {
		_spec.ClearField(msgsubscriber.FieldExclude, field.TypeBool)
	}
	if msuo.mutation.MsgTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   msgsubscriber.MsgTypeTable,
			Columns: []string{msgsubscriber.MsgTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtype.FieldID, field.TypeInt),
			},
		}
		edge.Schema = msuo.schemaConfig.MsgSubscriber
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msuo.mutation.MsgTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   msgsubscriber.MsgTypeTable,
			Columns: []string{msgsubscriber.MsgTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgtype.FieldID, field.TypeInt),
			},
		}
		edge.Schema = msuo.schemaConfig.MsgSubscriber
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if msuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   msgsubscriber.UserTable,
			Columns: []string{msgsubscriber.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = msuo.schemaConfig.MsgSubscriber
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := msuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   msgsubscriber.UserTable,
			Columns: []string{msgsubscriber.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = msuo.schemaConfig.MsgSubscriber
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = msuo.schemaConfig.MsgSubscriber
	ctx = internal.NewSchemaConfigContext(ctx, msuo.schemaConfig)
	_node = &MsgSubscriber{config: msuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, msuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{msgsubscriber.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	msuo.mutation.done = true
	return _node, nil
}
