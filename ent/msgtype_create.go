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
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtype"
)

// MsgTypeCreate is the builder for creating a MsgType entity.
type MsgTypeCreate struct {
	config
	mutation *MsgTypeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (mtc *MsgTypeCreate) SetCreatedBy(i int) *MsgTypeCreate {
	mtc.mutation.SetCreatedBy(i)
	return mtc
}

// SetCreatedAt sets the "created_at" field.
func (mtc *MsgTypeCreate) SetCreatedAt(t time.Time) *MsgTypeCreate {
	mtc.mutation.SetCreatedAt(t)
	return mtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mtc *MsgTypeCreate) SetNillableCreatedAt(t *time.Time) *MsgTypeCreate {
	if t != nil {
		mtc.SetCreatedAt(*t)
	}
	return mtc
}

// SetUpdatedBy sets the "updated_by" field.
func (mtc *MsgTypeCreate) SetUpdatedBy(i int) *MsgTypeCreate {
	mtc.mutation.SetUpdatedBy(i)
	return mtc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mtc *MsgTypeCreate) SetNillableUpdatedBy(i *int) *MsgTypeCreate {
	if i != nil {
		mtc.SetUpdatedBy(*i)
	}
	return mtc
}

// SetUpdatedAt sets the "updated_at" field.
func (mtc *MsgTypeCreate) SetUpdatedAt(t time.Time) *MsgTypeCreate {
	mtc.mutation.SetUpdatedAt(t)
	return mtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mtc *MsgTypeCreate) SetNillableUpdatedAt(t *time.Time) *MsgTypeCreate {
	if t != nil {
		mtc.SetUpdatedAt(*t)
	}
	return mtc
}

// SetAppID sets the "app_id" field.
func (mtc *MsgTypeCreate) SetAppID(i int) *MsgTypeCreate {
	mtc.mutation.SetAppID(i)
	return mtc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (mtc *MsgTypeCreate) SetNillableAppID(i *int) *MsgTypeCreate {
	if i != nil {
		mtc.SetAppID(*i)
	}
	return mtc
}

// SetCategory sets the "category" field.
func (mtc *MsgTypeCreate) SetCategory(s string) *MsgTypeCreate {
	mtc.mutation.SetCategory(s)
	return mtc
}

// SetName sets the "name" field.
func (mtc *MsgTypeCreate) SetName(s string) *MsgTypeCreate {
	mtc.mutation.SetName(s)
	return mtc
}

// SetStatus sets the "status" field.
func (mtc *MsgTypeCreate) SetStatus(ts typex.SimpleStatus) *MsgTypeCreate {
	mtc.mutation.SetStatus(ts)
	return mtc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mtc *MsgTypeCreate) SetNillableStatus(ts *typex.SimpleStatus) *MsgTypeCreate {
	if ts != nil {
		mtc.SetStatus(*ts)
	}
	return mtc
}

// SetComments sets the "comments" field.
func (mtc *MsgTypeCreate) SetComments(s string) *MsgTypeCreate {
	mtc.mutation.SetComments(s)
	return mtc
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (mtc *MsgTypeCreate) SetNillableComments(s *string) *MsgTypeCreate {
	if s != nil {
		mtc.SetComments(*s)
	}
	return mtc
}

// SetCanSubs sets the "can_subs" field.
func (mtc *MsgTypeCreate) SetCanSubs(b bool) *MsgTypeCreate {
	mtc.mutation.SetCanSubs(b)
	return mtc
}

// SetNillableCanSubs sets the "can_subs" field if the given value is not nil.
func (mtc *MsgTypeCreate) SetNillableCanSubs(b *bool) *MsgTypeCreate {
	if b != nil {
		mtc.SetCanSubs(*b)
	}
	return mtc
}

// SetCanCustom sets the "can_custom" field.
func (mtc *MsgTypeCreate) SetCanCustom(b bool) *MsgTypeCreate {
	mtc.mutation.SetCanCustom(b)
	return mtc
}

// SetNillableCanCustom sets the "can_custom" field if the given value is not nil.
func (mtc *MsgTypeCreate) SetNillableCanCustom(b *bool) *MsgTypeCreate {
	if b != nil {
		mtc.SetCanCustom(*b)
	}
	return mtc
}

// SetID sets the "id" field.
func (mtc *MsgTypeCreate) SetID(i int) *MsgTypeCreate {
	mtc.mutation.SetID(i)
	return mtc
}

// AddEventIDs adds the "events" edge to the MsgEvent entity by IDs.
func (mtc *MsgTypeCreate) AddEventIDs(ids ...int) *MsgTypeCreate {
	mtc.mutation.AddEventIDs(ids...)
	return mtc
}

// AddEvents adds the "events" edges to the MsgEvent entity.
func (mtc *MsgTypeCreate) AddEvents(m ...*MsgEvent) *MsgTypeCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mtc.AddEventIDs(ids...)
}

// AddSubscriberIDs adds the "subscribers" edge to the MsgSubscriber entity by IDs.
func (mtc *MsgTypeCreate) AddSubscriberIDs(ids ...int) *MsgTypeCreate {
	mtc.mutation.AddSubscriberIDs(ids...)
	return mtc
}

// AddSubscribers adds the "subscribers" edges to the MsgSubscriber entity.
func (mtc *MsgTypeCreate) AddSubscribers(m ...*MsgSubscriber) *MsgTypeCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mtc.AddSubscriberIDs(ids...)
}

// Mutation returns the MsgTypeMutation object of the builder.
func (mtc *MsgTypeCreate) Mutation() *MsgTypeMutation {
	return mtc.mutation
}

// Save creates the MsgType in the database.
func (mtc *MsgTypeCreate) Save(ctx context.Context) (*MsgType, error) {
	if err := mtc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, mtc.sqlSave, mtc.mutation, mtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mtc *MsgTypeCreate) SaveX(ctx context.Context) *MsgType {
	v, err := mtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mtc *MsgTypeCreate) Exec(ctx context.Context) error {
	_, err := mtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtc *MsgTypeCreate) ExecX(ctx context.Context) {
	if err := mtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mtc *MsgTypeCreate) defaults() error {
	if _, ok := mtc.mutation.CreatedAt(); !ok {
		if msgtype.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized msgtype.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := msgtype.DefaultCreatedAt()
		mtc.mutation.SetCreatedAt(v)
	}
	if _, ok := mtc.mutation.Status(); !ok {
		v := msgtype.DefaultStatus
		mtc.mutation.SetStatus(v)
	}
	if _, ok := mtc.mutation.CanSubs(); !ok {
		v := msgtype.DefaultCanSubs
		mtc.mutation.SetCanSubs(v)
	}
	if _, ok := mtc.mutation.CanCustom(); !ok {
		v := msgtype.DefaultCanCustom
		mtc.mutation.SetCanCustom(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mtc *MsgTypeCreate) check() error {
	if _, ok := mtc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "MsgType.created_by"`)}
	}
	if _, ok := mtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MsgType.created_at"`)}
	}
	if _, ok := mtc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "MsgType.category"`)}
	}
	if v, ok := mtc.mutation.Category(); ok {
		if err := msgtype.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "MsgType.category": %w`, err)}
		}
	}
	if _, ok := mtc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MsgType.name"`)}
	}
	if v, ok := mtc.mutation.Name(); ok {
		if err := msgtype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MsgType.name": %w`, err)}
		}
	}
	if v, ok := mtc.mutation.Status(); ok {
		if err := msgtype.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "MsgType.status": %w`, err)}
		}
	}
	return nil
}

func (mtc *MsgTypeCreate) sqlSave(ctx context.Context) (*MsgType, error) {
	if err := mtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	mtc.mutation.id = &_node.ID
	mtc.mutation.done = true
	return _node, nil
}

func (mtc *MsgTypeCreate) createSpec() (*MsgType, *sqlgraph.CreateSpec) {
	var (
		_node = &MsgType{config: mtc.config}
		_spec = sqlgraph.NewCreateSpec(msgtype.Table, sqlgraph.NewFieldSpec(msgtype.FieldID, field.TypeInt))
	)
	_spec.Schema = mtc.schemaConfig.MsgType
	_spec.OnConflict = mtc.conflict
	if id, ok := mtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mtc.mutation.CreatedBy(); ok {
		_spec.SetField(msgtype.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := mtc.mutation.CreatedAt(); ok {
		_spec.SetField(msgtype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mtc.mutation.UpdatedBy(); ok {
		_spec.SetField(msgtype.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := mtc.mutation.UpdatedAt(); ok {
		_spec.SetField(msgtype.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mtc.mutation.AppID(); ok {
		_spec.SetField(msgtype.FieldAppID, field.TypeInt, value)
		_node.AppID = value
	}
	if value, ok := mtc.mutation.Category(); ok {
		_spec.SetField(msgtype.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := mtc.mutation.Name(); ok {
		_spec.SetField(msgtype.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mtc.mutation.Status(); ok {
		_spec.SetField(msgtype.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := mtc.mutation.Comments(); ok {
		_spec.SetField(msgtype.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if value, ok := mtc.mutation.CanSubs(); ok {
		_spec.SetField(msgtype.FieldCanSubs, field.TypeBool, value)
		_node.CanSubs = value
	}
	if value, ok := mtc.mutation.CanCustom(); ok {
		_spec.SetField(msgtype.FieldCanCustom, field.TypeBool, value)
		_node.CanCustom = value
	}
	if nodes := mtc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   msgtype.EventsTable,
			Columns: []string{msgtype.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgevent.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mtc.schemaConfig.MsgEvent
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mtc.mutation.SubscribersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   msgtype.SubscribersTable,
			Columns: []string{msgtype.SubscribersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgsubscriber.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mtc.schemaConfig.MsgSubscriber
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
//	client.MsgType.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MsgTypeUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (mtc *MsgTypeCreate) OnConflict(opts ...sql.ConflictOption) *MsgTypeUpsertOne {
	mtc.conflict = opts
	return &MsgTypeUpsertOne{
		create: mtc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MsgType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mtc *MsgTypeCreate) OnConflictColumns(columns ...string) *MsgTypeUpsertOne {
	mtc.conflict = append(mtc.conflict, sql.ConflictColumns(columns...))
	return &MsgTypeUpsertOne{
		create: mtc,
	}
}

type (
	// MsgTypeUpsertOne is the builder for "upsert"-ing
	//  one MsgType node.
	MsgTypeUpsertOne struct {
		create *MsgTypeCreate
	}

	// MsgTypeUpsert is the "OnConflict" setter.
	MsgTypeUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *MsgTypeUpsert) SetUpdatedBy(v int) *MsgTypeUpsert {
	u.Set(msgtype.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MsgTypeUpsert) UpdateUpdatedBy() *MsgTypeUpsert {
	u.SetExcluded(msgtype.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MsgTypeUpsert) AddUpdatedBy(v int) *MsgTypeUpsert {
	u.Add(msgtype.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *MsgTypeUpsert) ClearUpdatedBy() *MsgTypeUpsert {
	u.SetNull(msgtype.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MsgTypeUpsert) SetUpdatedAt(v time.Time) *MsgTypeUpsert {
	u.Set(msgtype.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MsgTypeUpsert) UpdateUpdatedAt() *MsgTypeUpsert {
	u.SetExcluded(msgtype.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MsgTypeUpsert) ClearUpdatedAt() *MsgTypeUpsert {
	u.SetNull(msgtype.FieldUpdatedAt)
	return u
}

// SetAppID sets the "app_id" field.
func (u *MsgTypeUpsert) SetAppID(v int) *MsgTypeUpsert {
	u.Set(msgtype.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MsgTypeUpsert) UpdateAppID() *MsgTypeUpsert {
	u.SetExcluded(msgtype.FieldAppID)
	return u
}

// AddAppID adds v to the "app_id" field.
func (u *MsgTypeUpsert) AddAppID(v int) *MsgTypeUpsert {
	u.Add(msgtype.FieldAppID, v)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *MsgTypeUpsert) ClearAppID() *MsgTypeUpsert {
	u.SetNull(msgtype.FieldAppID)
	return u
}

// SetCategory sets the "category" field.
func (u *MsgTypeUpsert) SetCategory(v string) *MsgTypeUpsert {
	u.Set(msgtype.FieldCategory, v)
	return u
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *MsgTypeUpsert) UpdateCategory() *MsgTypeUpsert {
	u.SetExcluded(msgtype.FieldCategory)
	return u
}

// SetName sets the "name" field.
func (u *MsgTypeUpsert) SetName(v string) *MsgTypeUpsert {
	u.Set(msgtype.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MsgTypeUpsert) UpdateName() *MsgTypeUpsert {
	u.SetExcluded(msgtype.FieldName)
	return u
}

// SetStatus sets the "status" field.
func (u *MsgTypeUpsert) SetStatus(v typex.SimpleStatus) *MsgTypeUpsert {
	u.Set(msgtype.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MsgTypeUpsert) UpdateStatus() *MsgTypeUpsert {
	u.SetExcluded(msgtype.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *MsgTypeUpsert) ClearStatus() *MsgTypeUpsert {
	u.SetNull(msgtype.FieldStatus)
	return u
}

// SetComments sets the "comments" field.
func (u *MsgTypeUpsert) SetComments(v string) *MsgTypeUpsert {
	u.Set(msgtype.FieldComments, v)
	return u
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *MsgTypeUpsert) UpdateComments() *MsgTypeUpsert {
	u.SetExcluded(msgtype.FieldComments)
	return u
}

// ClearComments clears the value of the "comments" field.
func (u *MsgTypeUpsert) ClearComments() *MsgTypeUpsert {
	u.SetNull(msgtype.FieldComments)
	return u
}

// SetCanSubs sets the "can_subs" field.
func (u *MsgTypeUpsert) SetCanSubs(v bool) *MsgTypeUpsert {
	u.Set(msgtype.FieldCanSubs, v)
	return u
}

// UpdateCanSubs sets the "can_subs" field to the value that was provided on create.
func (u *MsgTypeUpsert) UpdateCanSubs() *MsgTypeUpsert {
	u.SetExcluded(msgtype.FieldCanSubs)
	return u
}

// ClearCanSubs clears the value of the "can_subs" field.
func (u *MsgTypeUpsert) ClearCanSubs() *MsgTypeUpsert {
	u.SetNull(msgtype.FieldCanSubs)
	return u
}

// SetCanCustom sets the "can_custom" field.
func (u *MsgTypeUpsert) SetCanCustom(v bool) *MsgTypeUpsert {
	u.Set(msgtype.FieldCanCustom, v)
	return u
}

// UpdateCanCustom sets the "can_custom" field to the value that was provided on create.
func (u *MsgTypeUpsert) UpdateCanCustom() *MsgTypeUpsert {
	u.SetExcluded(msgtype.FieldCanCustom)
	return u
}

// ClearCanCustom clears the value of the "can_custom" field.
func (u *MsgTypeUpsert) ClearCanCustom() *MsgTypeUpsert {
	u.SetNull(msgtype.FieldCanCustom)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MsgType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(msgtype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MsgTypeUpsertOne) UpdateNewValues() *MsgTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(msgtype.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(msgtype.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(msgtype.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MsgType.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MsgTypeUpsertOne) Ignore() *MsgTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MsgTypeUpsertOne) DoNothing() *MsgTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MsgTypeCreate.OnConflict
// documentation for more info.
func (u *MsgTypeUpsertOne) Update(set func(*MsgTypeUpsert)) *MsgTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MsgTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MsgTypeUpsertOne) SetUpdatedBy(v int) *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MsgTypeUpsertOne) AddUpdatedBy(v int) *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MsgTypeUpsertOne) UpdateUpdatedBy() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *MsgTypeUpsertOne) ClearUpdatedBy() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MsgTypeUpsertOne) SetUpdatedAt(v time.Time) *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MsgTypeUpsertOne) UpdateUpdatedAt() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MsgTypeUpsertOne) ClearUpdatedAt() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *MsgTypeUpsertOne) SetAppID(v int) *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetAppID(v)
	})
}

// AddAppID adds v to the "app_id" field.
func (u *MsgTypeUpsertOne) AddAppID(v int) *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.AddAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MsgTypeUpsertOne) UpdateAppID() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *MsgTypeUpsertOne) ClearAppID() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearAppID()
	})
}

// SetCategory sets the "category" field.
func (u *MsgTypeUpsertOne) SetCategory(v string) *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetCategory(v)
	})
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *MsgTypeUpsertOne) UpdateCategory() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateCategory()
	})
}

// SetName sets the "name" field.
func (u *MsgTypeUpsertOne) SetName(v string) *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MsgTypeUpsertOne) UpdateName() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateName()
	})
}

// SetStatus sets the "status" field.
func (u *MsgTypeUpsertOne) SetStatus(v typex.SimpleStatus) *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MsgTypeUpsertOne) UpdateStatus() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *MsgTypeUpsertOne) ClearStatus() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearStatus()
	})
}

// SetComments sets the "comments" field.
func (u *MsgTypeUpsertOne) SetComments(v string) *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *MsgTypeUpsertOne) UpdateComments() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *MsgTypeUpsertOne) ClearComments() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearComments()
	})
}

// SetCanSubs sets the "can_subs" field.
func (u *MsgTypeUpsertOne) SetCanSubs(v bool) *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetCanSubs(v)
	})
}

// UpdateCanSubs sets the "can_subs" field to the value that was provided on create.
func (u *MsgTypeUpsertOne) UpdateCanSubs() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateCanSubs()
	})
}

// ClearCanSubs clears the value of the "can_subs" field.
func (u *MsgTypeUpsertOne) ClearCanSubs() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearCanSubs()
	})
}

// SetCanCustom sets the "can_custom" field.
func (u *MsgTypeUpsertOne) SetCanCustom(v bool) *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetCanCustom(v)
	})
}

// UpdateCanCustom sets the "can_custom" field to the value that was provided on create.
func (u *MsgTypeUpsertOne) UpdateCanCustom() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateCanCustom()
	})
}

// ClearCanCustom clears the value of the "can_custom" field.
func (u *MsgTypeUpsertOne) ClearCanCustom() *MsgTypeUpsertOne {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearCanCustom()
	})
}

// Exec executes the query.
func (u *MsgTypeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MsgTypeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MsgTypeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MsgTypeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MsgTypeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MsgTypeCreateBulk is the builder for creating many MsgType entities in bulk.
type MsgTypeCreateBulk struct {
	config
	err      error
	builders []*MsgTypeCreate
	conflict []sql.ConflictOption
}

// Save creates the MsgType entities in the database.
func (mtcb *MsgTypeCreateBulk) Save(ctx context.Context) ([]*MsgType, error) {
	if mtcb.err != nil {
		return nil, mtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mtcb.builders))
	nodes := make([]*MsgType, len(mtcb.builders))
	mutators := make([]Mutator, len(mtcb.builders))
	for i := range mtcb.builders {
		func(i int, root context.Context) {
			builder := mtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MsgTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, mtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mtcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mtcb *MsgTypeCreateBulk) SaveX(ctx context.Context) []*MsgType {
	v, err := mtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mtcb *MsgTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := mtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mtcb *MsgTypeCreateBulk) ExecX(ctx context.Context) {
	if err := mtcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MsgType.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MsgTypeUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (mtcb *MsgTypeCreateBulk) OnConflict(opts ...sql.ConflictOption) *MsgTypeUpsertBulk {
	mtcb.conflict = opts
	return &MsgTypeUpsertBulk{
		create: mtcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MsgType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mtcb *MsgTypeCreateBulk) OnConflictColumns(columns ...string) *MsgTypeUpsertBulk {
	mtcb.conflict = append(mtcb.conflict, sql.ConflictColumns(columns...))
	return &MsgTypeUpsertBulk{
		create: mtcb,
	}
}

// MsgTypeUpsertBulk is the builder for "upsert"-ing
// a bulk of MsgType nodes.
type MsgTypeUpsertBulk struct {
	create *MsgTypeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MsgType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(msgtype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MsgTypeUpsertBulk) UpdateNewValues() *MsgTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(msgtype.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(msgtype.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(msgtype.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MsgType.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MsgTypeUpsertBulk) Ignore() *MsgTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MsgTypeUpsertBulk) DoNothing() *MsgTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MsgTypeCreateBulk.OnConflict
// documentation for more info.
func (u *MsgTypeUpsertBulk) Update(set func(*MsgTypeUpsert)) *MsgTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MsgTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MsgTypeUpsertBulk) SetUpdatedBy(v int) *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MsgTypeUpsertBulk) AddUpdatedBy(v int) *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MsgTypeUpsertBulk) UpdateUpdatedBy() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *MsgTypeUpsertBulk) ClearUpdatedBy() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MsgTypeUpsertBulk) SetUpdatedAt(v time.Time) *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MsgTypeUpsertBulk) UpdateUpdatedAt() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MsgTypeUpsertBulk) ClearUpdatedAt() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *MsgTypeUpsertBulk) SetAppID(v int) *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetAppID(v)
	})
}

// AddAppID adds v to the "app_id" field.
func (u *MsgTypeUpsertBulk) AddAppID(v int) *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.AddAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MsgTypeUpsertBulk) UpdateAppID() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *MsgTypeUpsertBulk) ClearAppID() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearAppID()
	})
}

// SetCategory sets the "category" field.
func (u *MsgTypeUpsertBulk) SetCategory(v string) *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetCategory(v)
	})
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *MsgTypeUpsertBulk) UpdateCategory() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateCategory()
	})
}

// SetName sets the "name" field.
func (u *MsgTypeUpsertBulk) SetName(v string) *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MsgTypeUpsertBulk) UpdateName() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateName()
	})
}

// SetStatus sets the "status" field.
func (u *MsgTypeUpsertBulk) SetStatus(v typex.SimpleStatus) *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MsgTypeUpsertBulk) UpdateStatus() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *MsgTypeUpsertBulk) ClearStatus() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearStatus()
	})
}

// SetComments sets the "comments" field.
func (u *MsgTypeUpsertBulk) SetComments(v string) *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *MsgTypeUpsertBulk) UpdateComments() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *MsgTypeUpsertBulk) ClearComments() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearComments()
	})
}

// SetCanSubs sets the "can_subs" field.
func (u *MsgTypeUpsertBulk) SetCanSubs(v bool) *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetCanSubs(v)
	})
}

// UpdateCanSubs sets the "can_subs" field to the value that was provided on create.
func (u *MsgTypeUpsertBulk) UpdateCanSubs() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateCanSubs()
	})
}

// ClearCanSubs clears the value of the "can_subs" field.
func (u *MsgTypeUpsertBulk) ClearCanSubs() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearCanSubs()
	})
}

// SetCanCustom sets the "can_custom" field.
func (u *MsgTypeUpsertBulk) SetCanCustom(v bool) *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.SetCanCustom(v)
	})
}

// UpdateCanCustom sets the "can_custom" field to the value that was provided on create.
func (u *MsgTypeUpsertBulk) UpdateCanCustom() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.UpdateCanCustom()
	})
}

// ClearCanCustom clears the value of the "can_custom" field.
func (u *MsgTypeUpsertBulk) ClearCanCustom() *MsgTypeUpsertBulk {
	return u.Update(func(s *MsgTypeUpsert) {
		s.ClearCanCustom()
	})
}

// Exec executes the query.
func (u *MsgTypeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MsgTypeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MsgTypeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MsgTypeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
