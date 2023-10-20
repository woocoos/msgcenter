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
	"github.com/woocoos/msgcenter/pkg/profile"
)

// MsgEventCreate is the builder for creating a MsgEvent entity.
type MsgEventCreate struct {
	config
	mutation *MsgEventMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (mec *MsgEventCreate) SetCreatedBy(i int) *MsgEventCreate {
	mec.mutation.SetCreatedBy(i)
	return mec
}

// SetCreatedAt sets the "created_at" field.
func (mec *MsgEventCreate) SetCreatedAt(t time.Time) *MsgEventCreate {
	mec.mutation.SetCreatedAt(t)
	return mec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mec *MsgEventCreate) SetNillableCreatedAt(t *time.Time) *MsgEventCreate {
	if t != nil {
		mec.SetCreatedAt(*t)
	}
	return mec
}

// SetUpdatedBy sets the "updated_by" field.
func (mec *MsgEventCreate) SetUpdatedBy(i int) *MsgEventCreate {
	mec.mutation.SetUpdatedBy(i)
	return mec
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mec *MsgEventCreate) SetNillableUpdatedBy(i *int) *MsgEventCreate {
	if i != nil {
		mec.SetUpdatedBy(*i)
	}
	return mec
}

// SetUpdatedAt sets the "updated_at" field.
func (mec *MsgEventCreate) SetUpdatedAt(t time.Time) *MsgEventCreate {
	mec.mutation.SetUpdatedAt(t)
	return mec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mec *MsgEventCreate) SetNillableUpdatedAt(t *time.Time) *MsgEventCreate {
	if t != nil {
		mec.SetUpdatedAt(*t)
	}
	return mec
}

// SetMsgTypeID sets the "msg_type_id" field.
func (mec *MsgEventCreate) SetMsgTypeID(i int) *MsgEventCreate {
	mec.mutation.SetMsgTypeID(i)
	return mec
}

// SetName sets the "name" field.
func (mec *MsgEventCreate) SetName(s string) *MsgEventCreate {
	mec.mutation.SetName(s)
	return mec
}

// SetStatus sets the "status" field.
func (mec *MsgEventCreate) SetStatus(ts typex.SimpleStatus) *MsgEventCreate {
	mec.mutation.SetStatus(ts)
	return mec
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mec *MsgEventCreate) SetNillableStatus(ts *typex.SimpleStatus) *MsgEventCreate {
	if ts != nil {
		mec.SetStatus(*ts)
	}
	return mec
}

// SetComments sets the "comments" field.
func (mec *MsgEventCreate) SetComments(s string) *MsgEventCreate {
	mec.mutation.SetComments(s)
	return mec
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (mec *MsgEventCreate) SetNillableComments(s *string) *MsgEventCreate {
	if s != nil {
		mec.SetComments(*s)
	}
	return mec
}

// SetRoute sets the "route" field.
func (mec *MsgEventCreate) SetRoute(pr *profile.Route) *MsgEventCreate {
	mec.mutation.SetRoute(pr)
	return mec
}

// SetModes sets the "modes" field.
func (mec *MsgEventCreate) SetModes(s string) *MsgEventCreate {
	mec.mutation.SetModes(s)
	return mec
}

// SetID sets the "id" field.
func (mec *MsgEventCreate) SetID(i int) *MsgEventCreate {
	mec.mutation.SetID(i)
	return mec
}

// SetMsgType sets the "msg_type" edge to the MsgType entity.
func (mec *MsgEventCreate) SetMsgType(m *MsgType) *MsgEventCreate {
	return mec.SetMsgTypeID(m.ID)
}

// AddCustomerTemplateIDs adds the "customer_template" edge to the MsgTemplate entity by IDs.
func (mec *MsgEventCreate) AddCustomerTemplateIDs(ids ...int) *MsgEventCreate {
	mec.mutation.AddCustomerTemplateIDs(ids...)
	return mec
}

// AddCustomerTemplate adds the "customer_template" edges to the MsgTemplate entity.
func (mec *MsgEventCreate) AddCustomerTemplate(m ...*MsgTemplate) *MsgEventCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mec.AddCustomerTemplateIDs(ids...)
}

// Mutation returns the MsgEventMutation object of the builder.
func (mec *MsgEventCreate) Mutation() *MsgEventMutation {
	return mec.mutation
}

// Save creates the MsgEvent in the database.
func (mec *MsgEventCreate) Save(ctx context.Context) (*MsgEvent, error) {
	if err := mec.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, mec.sqlSave, mec.mutation, mec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mec *MsgEventCreate) SaveX(ctx context.Context) *MsgEvent {
	v, err := mec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mec *MsgEventCreate) Exec(ctx context.Context) error {
	_, err := mec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mec *MsgEventCreate) ExecX(ctx context.Context) {
	if err := mec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mec *MsgEventCreate) defaults() error {
	if _, ok := mec.mutation.CreatedAt(); !ok {
		if msgevent.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized msgevent.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := msgevent.DefaultCreatedAt()
		mec.mutation.SetCreatedAt(v)
	}
	if _, ok := mec.mutation.Status(); !ok {
		v := msgevent.DefaultStatus
		mec.mutation.SetStatus(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mec *MsgEventCreate) check() error {
	if _, ok := mec.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "MsgEvent.created_by"`)}
	}
	if _, ok := mec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MsgEvent.created_at"`)}
	}
	if _, ok := mec.mutation.MsgTypeID(); !ok {
		return &ValidationError{Name: "msg_type_id", err: errors.New(`ent: missing required field "MsgEvent.msg_type_id"`)}
	}
	if _, ok := mec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MsgEvent.name"`)}
	}
	if v, ok := mec.mutation.Name(); ok {
		if err := msgevent.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MsgEvent.name": %w`, err)}
		}
	}
	if v, ok := mec.mutation.Status(); ok {
		if err := msgevent.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "MsgEvent.status": %w`, err)}
		}
	}
	if v, ok := mec.mutation.Route(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "route", err: fmt.Errorf(`ent: validator failed for field "MsgEvent.route": %w`, err)}
		}
	}
	if _, ok := mec.mutation.Modes(); !ok {
		return &ValidationError{Name: "modes", err: errors.New(`ent: missing required field "MsgEvent.modes"`)}
	}
	if _, ok := mec.mutation.MsgTypeID(); !ok {
		return &ValidationError{Name: "msg_type", err: errors.New(`ent: missing required edge "MsgEvent.msg_type"`)}
	}
	return nil
}

func (mec *MsgEventCreate) sqlSave(ctx context.Context) (*MsgEvent, error) {
	if err := mec.check(); err != nil {
		return nil, err
	}
	_node, _spec := mec.createSpec()
	if err := sqlgraph.CreateNode(ctx, mec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	mec.mutation.id = &_node.ID
	mec.mutation.done = true
	return _node, nil
}

func (mec *MsgEventCreate) createSpec() (*MsgEvent, *sqlgraph.CreateSpec) {
	var (
		_node = &MsgEvent{config: mec.config}
		_spec = sqlgraph.NewCreateSpec(msgevent.Table, sqlgraph.NewFieldSpec(msgevent.FieldID, field.TypeInt))
	)
	_spec.Schema = mec.schemaConfig.MsgEvent
	_spec.OnConflict = mec.conflict
	if id, ok := mec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mec.mutation.CreatedBy(); ok {
		_spec.SetField(msgevent.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := mec.mutation.CreatedAt(); ok {
		_spec.SetField(msgevent.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mec.mutation.UpdatedBy(); ok {
		_spec.SetField(msgevent.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := mec.mutation.UpdatedAt(); ok {
		_spec.SetField(msgevent.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mec.mutation.Name(); ok {
		_spec.SetField(msgevent.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mec.mutation.Status(); ok {
		_spec.SetField(msgevent.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := mec.mutation.Comments(); ok {
		_spec.SetField(msgevent.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if value, ok := mec.mutation.Route(); ok {
		_spec.SetField(msgevent.FieldRoute, field.TypeJSON, value)
		_node.Route = value
	}
	if value, ok := mec.mutation.Modes(); ok {
		_spec.SetField(msgevent.FieldModes, field.TypeString, value)
		_node.Modes = value
	}
	if nodes := mec.mutation.MsgTypeIDs(); len(nodes) > 0 {
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
		edge.Schema = mec.schemaConfig.MsgEvent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MsgTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mec.mutation.CustomerTemplateIDs(); len(nodes) > 0 {
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
		edge.Schema = mec.schemaConfig.MsgTemplate
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MsgEvent.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MsgEventUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (mec *MsgEventCreate) OnConflict(opts ...sql.ConflictOption) *MsgEventUpsertOne {
	mec.conflict = opts
	return &MsgEventUpsertOne{
		create: mec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MsgEvent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mec *MsgEventCreate) OnConflictColumns(columns ...string) *MsgEventUpsertOne {
	mec.conflict = append(mec.conflict, sql.ConflictColumns(columns...))
	return &MsgEventUpsertOne{
		create: mec,
	}
}

type (
	// MsgEventUpsertOne is the builder for "upsert"-ing
	//  one MsgEvent node.
	MsgEventUpsertOne struct {
		create *MsgEventCreate
	}

	// MsgEventUpsert is the "OnConflict" setter.
	MsgEventUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *MsgEventUpsert) SetUpdatedBy(v int) *MsgEventUpsert {
	u.Set(msgevent.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MsgEventUpsert) UpdateUpdatedBy() *MsgEventUpsert {
	u.SetExcluded(msgevent.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MsgEventUpsert) AddUpdatedBy(v int) *MsgEventUpsert {
	u.Add(msgevent.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *MsgEventUpsert) ClearUpdatedBy() *MsgEventUpsert {
	u.SetNull(msgevent.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MsgEventUpsert) SetUpdatedAt(v time.Time) *MsgEventUpsert {
	u.Set(msgevent.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MsgEventUpsert) UpdateUpdatedAt() *MsgEventUpsert {
	u.SetExcluded(msgevent.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MsgEventUpsert) ClearUpdatedAt() *MsgEventUpsert {
	u.SetNull(msgevent.FieldUpdatedAt)
	return u
}

// SetMsgTypeID sets the "msg_type_id" field.
func (u *MsgEventUpsert) SetMsgTypeID(v int) *MsgEventUpsert {
	u.Set(msgevent.FieldMsgTypeID, v)
	return u
}

// UpdateMsgTypeID sets the "msg_type_id" field to the value that was provided on create.
func (u *MsgEventUpsert) UpdateMsgTypeID() *MsgEventUpsert {
	u.SetExcluded(msgevent.FieldMsgTypeID)
	return u
}

// SetName sets the "name" field.
func (u *MsgEventUpsert) SetName(v string) *MsgEventUpsert {
	u.Set(msgevent.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MsgEventUpsert) UpdateName() *MsgEventUpsert {
	u.SetExcluded(msgevent.FieldName)
	return u
}

// SetStatus sets the "status" field.
func (u *MsgEventUpsert) SetStatus(v typex.SimpleStatus) *MsgEventUpsert {
	u.Set(msgevent.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MsgEventUpsert) UpdateStatus() *MsgEventUpsert {
	u.SetExcluded(msgevent.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *MsgEventUpsert) ClearStatus() *MsgEventUpsert {
	u.SetNull(msgevent.FieldStatus)
	return u
}

// SetComments sets the "comments" field.
func (u *MsgEventUpsert) SetComments(v string) *MsgEventUpsert {
	u.Set(msgevent.FieldComments, v)
	return u
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *MsgEventUpsert) UpdateComments() *MsgEventUpsert {
	u.SetExcluded(msgevent.FieldComments)
	return u
}

// ClearComments clears the value of the "comments" field.
func (u *MsgEventUpsert) ClearComments() *MsgEventUpsert {
	u.SetNull(msgevent.FieldComments)
	return u
}

// SetRoute sets the "route" field.
func (u *MsgEventUpsert) SetRoute(v *profile.Route) *MsgEventUpsert {
	u.Set(msgevent.FieldRoute, v)
	return u
}

// UpdateRoute sets the "route" field to the value that was provided on create.
func (u *MsgEventUpsert) UpdateRoute() *MsgEventUpsert {
	u.SetExcluded(msgevent.FieldRoute)
	return u
}

// ClearRoute clears the value of the "route" field.
func (u *MsgEventUpsert) ClearRoute() *MsgEventUpsert {
	u.SetNull(msgevent.FieldRoute)
	return u
}

// SetModes sets the "modes" field.
func (u *MsgEventUpsert) SetModes(v string) *MsgEventUpsert {
	u.Set(msgevent.FieldModes, v)
	return u
}

// UpdateModes sets the "modes" field to the value that was provided on create.
func (u *MsgEventUpsert) UpdateModes() *MsgEventUpsert {
	u.SetExcluded(msgevent.FieldModes)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MsgEvent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(msgevent.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MsgEventUpsertOne) UpdateNewValues() *MsgEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(msgevent.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(msgevent.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(msgevent.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MsgEvent.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MsgEventUpsertOne) Ignore() *MsgEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MsgEventUpsertOne) DoNothing() *MsgEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MsgEventCreate.OnConflict
// documentation for more info.
func (u *MsgEventUpsertOne) Update(set func(*MsgEventUpsert)) *MsgEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MsgEventUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MsgEventUpsertOne) SetUpdatedBy(v int) *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MsgEventUpsertOne) AddUpdatedBy(v int) *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MsgEventUpsertOne) UpdateUpdatedBy() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *MsgEventUpsertOne) ClearUpdatedBy() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MsgEventUpsertOne) SetUpdatedAt(v time.Time) *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MsgEventUpsertOne) UpdateUpdatedAt() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MsgEventUpsertOne) ClearUpdatedAt() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetMsgTypeID sets the "msg_type_id" field.
func (u *MsgEventUpsertOne) SetMsgTypeID(v int) *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetMsgTypeID(v)
	})
}

// UpdateMsgTypeID sets the "msg_type_id" field to the value that was provided on create.
func (u *MsgEventUpsertOne) UpdateMsgTypeID() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateMsgTypeID()
	})
}

// SetName sets the "name" field.
func (u *MsgEventUpsertOne) SetName(v string) *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MsgEventUpsertOne) UpdateName() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateName()
	})
}

// SetStatus sets the "status" field.
func (u *MsgEventUpsertOne) SetStatus(v typex.SimpleStatus) *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MsgEventUpsertOne) UpdateStatus() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *MsgEventUpsertOne) ClearStatus() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.ClearStatus()
	})
}

// SetComments sets the "comments" field.
func (u *MsgEventUpsertOne) SetComments(v string) *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *MsgEventUpsertOne) UpdateComments() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *MsgEventUpsertOne) ClearComments() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.ClearComments()
	})
}

// SetRoute sets the "route" field.
func (u *MsgEventUpsertOne) SetRoute(v *profile.Route) *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetRoute(v)
	})
}

// UpdateRoute sets the "route" field to the value that was provided on create.
func (u *MsgEventUpsertOne) UpdateRoute() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateRoute()
	})
}

// ClearRoute clears the value of the "route" field.
func (u *MsgEventUpsertOne) ClearRoute() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.ClearRoute()
	})
}

// SetModes sets the "modes" field.
func (u *MsgEventUpsertOne) SetModes(v string) *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetModes(v)
	})
}

// UpdateModes sets the "modes" field to the value that was provided on create.
func (u *MsgEventUpsertOne) UpdateModes() *MsgEventUpsertOne {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateModes()
	})
}

// Exec executes the query.
func (u *MsgEventUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MsgEventCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MsgEventUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MsgEventUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MsgEventUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MsgEventCreateBulk is the builder for creating many MsgEvent entities in bulk.
type MsgEventCreateBulk struct {
	config
	err      error
	builders []*MsgEventCreate
	conflict []sql.ConflictOption
}

// Save creates the MsgEvent entities in the database.
func (mecb *MsgEventCreateBulk) Save(ctx context.Context) ([]*MsgEvent, error) {
	if mecb.err != nil {
		return nil, mecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mecb.builders))
	nodes := make([]*MsgEvent, len(mecb.builders))
	mutators := make([]Mutator, len(mecb.builders))
	for i := range mecb.builders {
		func(i int, root context.Context) {
			builder := mecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MsgEventMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mecb *MsgEventCreateBulk) SaveX(ctx context.Context) []*MsgEvent {
	v, err := mecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mecb *MsgEventCreateBulk) Exec(ctx context.Context) error {
	_, err := mecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mecb *MsgEventCreateBulk) ExecX(ctx context.Context) {
	if err := mecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MsgEvent.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MsgEventUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (mecb *MsgEventCreateBulk) OnConflict(opts ...sql.ConflictOption) *MsgEventUpsertBulk {
	mecb.conflict = opts
	return &MsgEventUpsertBulk{
		create: mecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MsgEvent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mecb *MsgEventCreateBulk) OnConflictColumns(columns ...string) *MsgEventUpsertBulk {
	mecb.conflict = append(mecb.conflict, sql.ConflictColumns(columns...))
	return &MsgEventUpsertBulk{
		create: mecb,
	}
}

// MsgEventUpsertBulk is the builder for "upsert"-ing
// a bulk of MsgEvent nodes.
type MsgEventUpsertBulk struct {
	create *MsgEventCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MsgEvent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(msgevent.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MsgEventUpsertBulk) UpdateNewValues() *MsgEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(msgevent.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(msgevent.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(msgevent.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MsgEvent.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MsgEventUpsertBulk) Ignore() *MsgEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MsgEventUpsertBulk) DoNothing() *MsgEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MsgEventCreateBulk.OnConflict
// documentation for more info.
func (u *MsgEventUpsertBulk) Update(set func(*MsgEventUpsert)) *MsgEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MsgEventUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MsgEventUpsertBulk) SetUpdatedBy(v int) *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MsgEventUpsertBulk) AddUpdatedBy(v int) *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MsgEventUpsertBulk) UpdateUpdatedBy() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *MsgEventUpsertBulk) ClearUpdatedBy() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MsgEventUpsertBulk) SetUpdatedAt(v time.Time) *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MsgEventUpsertBulk) UpdateUpdatedAt() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MsgEventUpsertBulk) ClearUpdatedAt() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetMsgTypeID sets the "msg_type_id" field.
func (u *MsgEventUpsertBulk) SetMsgTypeID(v int) *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetMsgTypeID(v)
	})
}

// UpdateMsgTypeID sets the "msg_type_id" field to the value that was provided on create.
func (u *MsgEventUpsertBulk) UpdateMsgTypeID() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateMsgTypeID()
	})
}

// SetName sets the "name" field.
func (u *MsgEventUpsertBulk) SetName(v string) *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MsgEventUpsertBulk) UpdateName() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateName()
	})
}

// SetStatus sets the "status" field.
func (u *MsgEventUpsertBulk) SetStatus(v typex.SimpleStatus) *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MsgEventUpsertBulk) UpdateStatus() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *MsgEventUpsertBulk) ClearStatus() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.ClearStatus()
	})
}

// SetComments sets the "comments" field.
func (u *MsgEventUpsertBulk) SetComments(v string) *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *MsgEventUpsertBulk) UpdateComments() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *MsgEventUpsertBulk) ClearComments() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.ClearComments()
	})
}

// SetRoute sets the "route" field.
func (u *MsgEventUpsertBulk) SetRoute(v *profile.Route) *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetRoute(v)
	})
}

// UpdateRoute sets the "route" field to the value that was provided on create.
func (u *MsgEventUpsertBulk) UpdateRoute() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateRoute()
	})
}

// ClearRoute clears the value of the "route" field.
func (u *MsgEventUpsertBulk) ClearRoute() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.ClearRoute()
	})
}

// SetModes sets the "modes" field.
func (u *MsgEventUpsertBulk) SetModes(v string) *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.SetModes(v)
	})
}

// UpdateModes sets the "modes" field to the value that was provided on create.
func (u *MsgEventUpsertBulk) UpdateModes() *MsgEventUpsertBulk {
	return u.Update(func(s *MsgEventUpsert) {
		s.UpdateModes()
	})
}

// Exec executes the query.
func (u *MsgEventUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MsgEventCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MsgEventCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MsgEventUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
