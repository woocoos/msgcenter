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
	"github.com/woocoos/msgcenter/ent/msginternal"
	"github.com/woocoos/msgcenter/ent/msginternalto"
)

// MsgInternalCreate is the builder for creating a MsgInternal entity.
type MsgInternalCreate struct {
	config
	mutation *MsgInternalMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTenantID sets the "tenant_id" field.
func (mic *MsgInternalCreate) SetTenantID(i int) *MsgInternalCreate {
	mic.mutation.SetTenantID(i)
	return mic
}

// SetCreatedBy sets the "created_by" field.
func (mic *MsgInternalCreate) SetCreatedBy(i int) *MsgInternalCreate {
	mic.mutation.SetCreatedBy(i)
	return mic
}

// SetCreatedAt sets the "created_at" field.
func (mic *MsgInternalCreate) SetCreatedAt(t time.Time) *MsgInternalCreate {
	mic.mutation.SetCreatedAt(t)
	return mic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mic *MsgInternalCreate) SetNillableCreatedAt(t *time.Time) *MsgInternalCreate {
	if t != nil {
		mic.SetCreatedAt(*t)
	}
	return mic
}

// SetUpdatedBy sets the "updated_by" field.
func (mic *MsgInternalCreate) SetUpdatedBy(i int) *MsgInternalCreate {
	mic.mutation.SetUpdatedBy(i)
	return mic
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mic *MsgInternalCreate) SetNillableUpdatedBy(i *int) *MsgInternalCreate {
	if i != nil {
		mic.SetUpdatedBy(*i)
	}
	return mic
}

// SetUpdatedAt sets the "updated_at" field.
func (mic *MsgInternalCreate) SetUpdatedAt(t time.Time) *MsgInternalCreate {
	mic.mutation.SetUpdatedAt(t)
	return mic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mic *MsgInternalCreate) SetNillableUpdatedAt(t *time.Time) *MsgInternalCreate {
	if t != nil {
		mic.SetUpdatedAt(*t)
	}
	return mic
}

// SetSubject sets the "subject" field.
func (mic *MsgInternalCreate) SetSubject(s string) *MsgInternalCreate {
	mic.mutation.SetSubject(s)
	return mic
}

// SetBody sets the "body" field.
func (mic *MsgInternalCreate) SetBody(s string) *MsgInternalCreate {
	mic.mutation.SetBody(s)
	return mic
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (mic *MsgInternalCreate) SetNillableBody(s *string) *MsgInternalCreate {
	if s != nil {
		mic.SetBody(*s)
	}
	return mic
}

// SetFormat sets the "format" field.
func (mic *MsgInternalCreate) SetFormat(s string) *MsgInternalCreate {
	mic.mutation.SetFormat(s)
	return mic
}

// SetRedirect sets the "redirect" field.
func (mic *MsgInternalCreate) SetRedirect(s string) *MsgInternalCreate {
	mic.mutation.SetRedirect(s)
	return mic
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (mic *MsgInternalCreate) SetNillableRedirect(s *string) *MsgInternalCreate {
	if s != nil {
		mic.SetRedirect(*s)
	}
	return mic
}

// SetID sets the "id" field.
func (mic *MsgInternalCreate) SetID(i int) *MsgInternalCreate {
	mic.mutation.SetID(i)
	return mic
}

// AddMsgInternalToIDs adds the "msg_internal_to" edge to the MsgInternalTo entity by IDs.
func (mic *MsgInternalCreate) AddMsgInternalToIDs(ids ...int) *MsgInternalCreate {
	mic.mutation.AddMsgInternalToIDs(ids...)
	return mic
}

// AddMsgInternalTo adds the "msg_internal_to" edges to the MsgInternalTo entity.
func (mic *MsgInternalCreate) AddMsgInternalTo(m ...*MsgInternalTo) *MsgInternalCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mic.AddMsgInternalToIDs(ids...)
}

// Mutation returns the MsgInternalMutation object of the builder.
func (mic *MsgInternalCreate) Mutation() *MsgInternalMutation {
	return mic.mutation
}

// Save creates the MsgInternal in the database.
func (mic *MsgInternalCreate) Save(ctx context.Context) (*MsgInternal, error) {
	if err := mic.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, mic.sqlSave, mic.mutation, mic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mic *MsgInternalCreate) SaveX(ctx context.Context) *MsgInternal {
	v, err := mic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mic *MsgInternalCreate) Exec(ctx context.Context) error {
	_, err := mic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mic *MsgInternalCreate) ExecX(ctx context.Context) {
	if err := mic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mic *MsgInternalCreate) defaults() error {
	if _, ok := mic.mutation.CreatedAt(); !ok {
		if msginternal.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized msginternal.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := msginternal.DefaultCreatedAt()
		mic.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mic *MsgInternalCreate) check() error {
	if _, ok := mic.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "MsgInternal.tenant_id"`)}
	}
	if _, ok := mic.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "MsgInternal.created_by"`)}
	}
	if _, ok := mic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MsgInternal.created_at"`)}
	}
	if _, ok := mic.mutation.Subject(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required field "MsgInternal.subject"`)}
	}
	if _, ok := mic.mutation.Format(); !ok {
		return &ValidationError{Name: "format", err: errors.New(`ent: missing required field "MsgInternal.format"`)}
	}
	return nil
}

func (mic *MsgInternalCreate) sqlSave(ctx context.Context) (*MsgInternal, error) {
	if err := mic.check(); err != nil {
		return nil, err
	}
	_node, _spec := mic.createSpec()
	if err := sqlgraph.CreateNode(ctx, mic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	mic.mutation.id = &_node.ID
	mic.mutation.done = true
	return _node, nil
}

func (mic *MsgInternalCreate) createSpec() (*MsgInternal, *sqlgraph.CreateSpec) {
	var (
		_node = &MsgInternal{config: mic.config}
		_spec = sqlgraph.NewCreateSpec(msginternal.Table, sqlgraph.NewFieldSpec(msginternal.FieldID, field.TypeInt))
	)
	_spec.Schema = mic.schemaConfig.MsgInternal
	_spec.OnConflict = mic.conflict
	if id, ok := mic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mic.mutation.TenantID(); ok {
		_spec.SetField(msginternal.FieldTenantID, field.TypeInt, value)
		_node.TenantID = value
	}
	if value, ok := mic.mutation.CreatedBy(); ok {
		_spec.SetField(msginternal.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := mic.mutation.CreatedAt(); ok {
		_spec.SetField(msginternal.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mic.mutation.UpdatedBy(); ok {
		_spec.SetField(msginternal.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := mic.mutation.UpdatedAt(); ok {
		_spec.SetField(msginternal.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mic.mutation.Subject(); ok {
		_spec.SetField(msginternal.FieldSubject, field.TypeString, value)
		_node.Subject = value
	}
	if value, ok := mic.mutation.Body(); ok {
		_spec.SetField(msginternal.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := mic.mutation.Format(); ok {
		_spec.SetField(msginternal.FieldFormat, field.TypeString, value)
		_node.Format = value
	}
	if value, ok := mic.mutation.Redirect(); ok {
		_spec.SetField(msginternal.FieldRedirect, field.TypeString, value)
		_node.Redirect = value
	}
	if nodes := mic.mutation.MsgInternalToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   msginternal.MsgInternalToTable,
			Columns: []string{msginternal.MsgInternalToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msginternalto.FieldID, field.TypeInt),
			},
		}
		edge.Schema = mic.schemaConfig.MsgInternalTo
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
//	client.MsgInternal.Create().
//		SetTenantID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MsgInternalUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
func (mic *MsgInternalCreate) OnConflict(opts ...sql.ConflictOption) *MsgInternalUpsertOne {
	mic.conflict = opts
	return &MsgInternalUpsertOne{
		create: mic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MsgInternal.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mic *MsgInternalCreate) OnConflictColumns(columns ...string) *MsgInternalUpsertOne {
	mic.conflict = append(mic.conflict, sql.ConflictColumns(columns...))
	return &MsgInternalUpsertOne{
		create: mic,
	}
}

type (
	// MsgInternalUpsertOne is the builder for "upsert"-ing
	//  one MsgInternal node.
	MsgInternalUpsertOne struct {
		create *MsgInternalCreate
	}

	// MsgInternalUpsert is the "OnConflict" setter.
	MsgInternalUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *MsgInternalUpsert) SetUpdatedBy(v int) *MsgInternalUpsert {
	u.Set(msginternal.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MsgInternalUpsert) UpdateUpdatedBy() *MsgInternalUpsert {
	u.SetExcluded(msginternal.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MsgInternalUpsert) AddUpdatedBy(v int) *MsgInternalUpsert {
	u.Add(msginternal.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *MsgInternalUpsert) ClearUpdatedBy() *MsgInternalUpsert {
	u.SetNull(msginternal.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MsgInternalUpsert) SetUpdatedAt(v time.Time) *MsgInternalUpsert {
	u.Set(msginternal.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MsgInternalUpsert) UpdateUpdatedAt() *MsgInternalUpsert {
	u.SetExcluded(msginternal.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MsgInternalUpsert) ClearUpdatedAt() *MsgInternalUpsert {
	u.SetNull(msginternal.FieldUpdatedAt)
	return u
}

// SetSubject sets the "subject" field.
func (u *MsgInternalUpsert) SetSubject(v string) *MsgInternalUpsert {
	u.Set(msginternal.FieldSubject, v)
	return u
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *MsgInternalUpsert) UpdateSubject() *MsgInternalUpsert {
	u.SetExcluded(msginternal.FieldSubject)
	return u
}

// SetBody sets the "body" field.
func (u *MsgInternalUpsert) SetBody(v string) *MsgInternalUpsert {
	u.Set(msginternal.FieldBody, v)
	return u
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *MsgInternalUpsert) UpdateBody() *MsgInternalUpsert {
	u.SetExcluded(msginternal.FieldBody)
	return u
}

// ClearBody clears the value of the "body" field.
func (u *MsgInternalUpsert) ClearBody() *MsgInternalUpsert {
	u.SetNull(msginternal.FieldBody)
	return u
}

// SetFormat sets the "format" field.
func (u *MsgInternalUpsert) SetFormat(v string) *MsgInternalUpsert {
	u.Set(msginternal.FieldFormat, v)
	return u
}

// UpdateFormat sets the "format" field to the value that was provided on create.
func (u *MsgInternalUpsert) UpdateFormat() *MsgInternalUpsert {
	u.SetExcluded(msginternal.FieldFormat)
	return u
}

// SetRedirect sets the "redirect" field.
func (u *MsgInternalUpsert) SetRedirect(v string) *MsgInternalUpsert {
	u.Set(msginternal.FieldRedirect, v)
	return u
}

// UpdateRedirect sets the "redirect" field to the value that was provided on create.
func (u *MsgInternalUpsert) UpdateRedirect() *MsgInternalUpsert {
	u.SetExcluded(msginternal.FieldRedirect)
	return u
}

// ClearRedirect clears the value of the "redirect" field.
func (u *MsgInternalUpsert) ClearRedirect() *MsgInternalUpsert {
	u.SetNull(msginternal.FieldRedirect)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MsgInternal.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(msginternal.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MsgInternalUpsertOne) UpdateNewValues() *MsgInternalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(msginternal.FieldID)
		}
		if _, exists := u.create.mutation.TenantID(); exists {
			s.SetIgnore(msginternal.FieldTenantID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(msginternal.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(msginternal.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MsgInternal.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MsgInternalUpsertOne) Ignore() *MsgInternalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MsgInternalUpsertOne) DoNothing() *MsgInternalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MsgInternalCreate.OnConflict
// documentation for more info.
func (u *MsgInternalUpsertOne) Update(set func(*MsgInternalUpsert)) *MsgInternalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MsgInternalUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MsgInternalUpsertOne) SetUpdatedBy(v int) *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MsgInternalUpsertOne) AddUpdatedBy(v int) *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MsgInternalUpsertOne) UpdateUpdatedBy() *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *MsgInternalUpsertOne) ClearUpdatedBy() *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MsgInternalUpsertOne) SetUpdatedAt(v time.Time) *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MsgInternalUpsertOne) UpdateUpdatedAt() *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MsgInternalUpsertOne) ClearUpdatedAt() *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetSubject sets the "subject" field.
func (u *MsgInternalUpsertOne) SetSubject(v string) *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetSubject(v)
	})
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *MsgInternalUpsertOne) UpdateSubject() *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateSubject()
	})
}

// SetBody sets the "body" field.
func (u *MsgInternalUpsertOne) SetBody(v string) *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetBody(v)
	})
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *MsgInternalUpsertOne) UpdateBody() *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateBody()
	})
}

// ClearBody clears the value of the "body" field.
func (u *MsgInternalUpsertOne) ClearBody() *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.ClearBody()
	})
}

// SetFormat sets the "format" field.
func (u *MsgInternalUpsertOne) SetFormat(v string) *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetFormat(v)
	})
}

// UpdateFormat sets the "format" field to the value that was provided on create.
func (u *MsgInternalUpsertOne) UpdateFormat() *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateFormat()
	})
}

// SetRedirect sets the "redirect" field.
func (u *MsgInternalUpsertOne) SetRedirect(v string) *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetRedirect(v)
	})
}

// UpdateRedirect sets the "redirect" field to the value that was provided on create.
func (u *MsgInternalUpsertOne) UpdateRedirect() *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateRedirect()
	})
}

// ClearRedirect clears the value of the "redirect" field.
func (u *MsgInternalUpsertOne) ClearRedirect() *MsgInternalUpsertOne {
	return u.Update(func(s *MsgInternalUpsert) {
		s.ClearRedirect()
	})
}

// Exec executes the query.
func (u *MsgInternalUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MsgInternalCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MsgInternalUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MsgInternalUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MsgInternalUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MsgInternalCreateBulk is the builder for creating many MsgInternal entities in bulk.
type MsgInternalCreateBulk struct {
	config
	builders []*MsgInternalCreate
	conflict []sql.ConflictOption
}

// Save creates the MsgInternal entities in the database.
func (micb *MsgInternalCreateBulk) Save(ctx context.Context) ([]*MsgInternal, error) {
	specs := make([]*sqlgraph.CreateSpec, len(micb.builders))
	nodes := make([]*MsgInternal, len(micb.builders))
	mutators := make([]Mutator, len(micb.builders))
	for i := range micb.builders {
		func(i int, root context.Context) {
			builder := micb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MsgInternalMutation)
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
					_, err = mutators[i+1].Mutate(root, micb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = micb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, micb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, micb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (micb *MsgInternalCreateBulk) SaveX(ctx context.Context) []*MsgInternal {
	v, err := micb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (micb *MsgInternalCreateBulk) Exec(ctx context.Context) error {
	_, err := micb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (micb *MsgInternalCreateBulk) ExecX(ctx context.Context) {
	if err := micb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MsgInternal.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MsgInternalUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
func (micb *MsgInternalCreateBulk) OnConflict(opts ...sql.ConflictOption) *MsgInternalUpsertBulk {
	micb.conflict = opts
	return &MsgInternalUpsertBulk{
		create: micb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MsgInternal.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (micb *MsgInternalCreateBulk) OnConflictColumns(columns ...string) *MsgInternalUpsertBulk {
	micb.conflict = append(micb.conflict, sql.ConflictColumns(columns...))
	return &MsgInternalUpsertBulk{
		create: micb,
	}
}

// MsgInternalUpsertBulk is the builder for "upsert"-ing
// a bulk of MsgInternal nodes.
type MsgInternalUpsertBulk struct {
	create *MsgInternalCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MsgInternal.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(msginternal.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MsgInternalUpsertBulk) UpdateNewValues() *MsgInternalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(msginternal.FieldID)
			}
			if _, exists := b.mutation.TenantID(); exists {
				s.SetIgnore(msginternal.FieldTenantID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(msginternal.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(msginternal.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MsgInternal.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MsgInternalUpsertBulk) Ignore() *MsgInternalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MsgInternalUpsertBulk) DoNothing() *MsgInternalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MsgInternalCreateBulk.OnConflict
// documentation for more info.
func (u *MsgInternalUpsertBulk) Update(set func(*MsgInternalUpsert)) *MsgInternalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MsgInternalUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *MsgInternalUpsertBulk) SetUpdatedBy(v int) *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *MsgInternalUpsertBulk) AddUpdatedBy(v int) *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *MsgInternalUpsertBulk) UpdateUpdatedBy() *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *MsgInternalUpsertBulk) ClearUpdatedBy() *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MsgInternalUpsertBulk) SetUpdatedAt(v time.Time) *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MsgInternalUpsertBulk) UpdateUpdatedAt() *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *MsgInternalUpsertBulk) ClearUpdatedAt() *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetSubject sets the "subject" field.
func (u *MsgInternalUpsertBulk) SetSubject(v string) *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetSubject(v)
	})
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *MsgInternalUpsertBulk) UpdateSubject() *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateSubject()
	})
}

// SetBody sets the "body" field.
func (u *MsgInternalUpsertBulk) SetBody(v string) *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetBody(v)
	})
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *MsgInternalUpsertBulk) UpdateBody() *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateBody()
	})
}

// ClearBody clears the value of the "body" field.
func (u *MsgInternalUpsertBulk) ClearBody() *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.ClearBody()
	})
}

// SetFormat sets the "format" field.
func (u *MsgInternalUpsertBulk) SetFormat(v string) *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetFormat(v)
	})
}

// UpdateFormat sets the "format" field to the value that was provided on create.
func (u *MsgInternalUpsertBulk) UpdateFormat() *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateFormat()
	})
}

// SetRedirect sets the "redirect" field.
func (u *MsgInternalUpsertBulk) SetRedirect(v string) *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.SetRedirect(v)
	})
}

// UpdateRedirect sets the "redirect" field to the value that was provided on create.
func (u *MsgInternalUpsertBulk) UpdateRedirect() *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.UpdateRedirect()
	})
}

// ClearRedirect clears the value of the "redirect" field.
func (u *MsgInternalUpsertBulk) ClearRedirect() *MsgInternalUpsertBulk {
	return u.Update(func(s *MsgInternalUpsert) {
		s.ClearRedirect()
	})
}

// Exec executes the query.
func (u *MsgInternalUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MsgInternalCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MsgInternalCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MsgInternalUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
