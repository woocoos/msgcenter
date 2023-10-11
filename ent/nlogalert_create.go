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
	"github.com/woocoos/msgcenter/ent/msgalert"
	"github.com/woocoos/msgcenter/ent/nlog"
	"github.com/woocoos/msgcenter/ent/nlogalert"
)

// NlogAlertCreate is the builder for creating a NlogAlert entity.
type NlogAlertCreate struct {
	config
	mutation *NlogAlertMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNlogID sets the "nlog_id" field.
func (nac *NlogAlertCreate) SetNlogID(i int) *NlogAlertCreate {
	nac.mutation.SetNlogID(i)
	return nac
}

// SetAlertID sets the "alert_id" field.
func (nac *NlogAlertCreate) SetAlertID(i int) *NlogAlertCreate {
	nac.mutation.SetAlertID(i)
	return nac
}

// SetCreatedAt sets the "created_at" field.
func (nac *NlogAlertCreate) SetCreatedAt(t time.Time) *NlogAlertCreate {
	nac.mutation.SetCreatedAt(t)
	return nac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nac *NlogAlertCreate) SetNillableCreatedAt(t *time.Time) *NlogAlertCreate {
	if t != nil {
		nac.SetCreatedAt(*t)
	}
	return nac
}

// SetNlog sets the "nlog" edge to the Nlog entity.
func (nac *NlogAlertCreate) SetNlog(n *Nlog) *NlogAlertCreate {
	return nac.SetNlogID(n.ID)
}

// SetAlert sets the "alert" edge to the MsgAlert entity.
func (nac *NlogAlertCreate) SetAlert(m *MsgAlert) *NlogAlertCreate {
	return nac.SetAlertID(m.ID)
}

// Mutation returns the NlogAlertMutation object of the builder.
func (nac *NlogAlertCreate) Mutation() *NlogAlertMutation {
	return nac.mutation
}

// Save creates the NlogAlert in the database.
func (nac *NlogAlertCreate) Save(ctx context.Context) (*NlogAlert, error) {
	nac.defaults()
	return withHooks(ctx, nac.sqlSave, nac.mutation, nac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nac *NlogAlertCreate) SaveX(ctx context.Context) *NlogAlert {
	v, err := nac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nac *NlogAlertCreate) Exec(ctx context.Context) error {
	_, err := nac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nac *NlogAlertCreate) ExecX(ctx context.Context) {
	if err := nac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nac *NlogAlertCreate) defaults() {
	if _, ok := nac.mutation.CreatedAt(); !ok {
		v := nlogalert.DefaultCreatedAt()
		nac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nac *NlogAlertCreate) check() error {
	if _, ok := nac.mutation.NlogID(); !ok {
		return &ValidationError{Name: "nlog_id", err: errors.New(`ent: missing required field "NlogAlert.nlog_id"`)}
	}
	if _, ok := nac.mutation.AlertID(); !ok {
		return &ValidationError{Name: "alert_id", err: errors.New(`ent: missing required field "NlogAlert.alert_id"`)}
	}
	if _, ok := nac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "NlogAlert.created_at"`)}
	}
	if _, ok := nac.mutation.NlogID(); !ok {
		return &ValidationError{Name: "nlog", err: errors.New(`ent: missing required edge "NlogAlert.nlog"`)}
	}
	if _, ok := nac.mutation.AlertID(); !ok {
		return &ValidationError{Name: "alert", err: errors.New(`ent: missing required edge "NlogAlert.alert"`)}
	}
	return nil
}

func (nac *NlogAlertCreate) sqlSave(ctx context.Context) (*NlogAlert, error) {
	if err := nac.check(); err != nil {
		return nil, err
	}
	_node, _spec := nac.createSpec()
	if err := sqlgraph.CreateNode(ctx, nac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	nac.mutation.id = &_node.ID
	nac.mutation.done = true
	return _node, nil
}

func (nac *NlogAlertCreate) createSpec() (*NlogAlert, *sqlgraph.CreateSpec) {
	var (
		_node = &NlogAlert{config: nac.config}
		_spec = sqlgraph.NewCreateSpec(nlogalert.Table, sqlgraph.NewFieldSpec(nlogalert.FieldID, field.TypeInt))
	)
	_spec.Schema = nac.schemaConfig.NlogAlert
	_spec.OnConflict = nac.conflict
	if value, ok := nac.mutation.CreatedAt(); ok {
		_spec.SetField(nlogalert.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := nac.mutation.NlogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   nlogalert.NlogTable,
			Columns: []string{nlogalert.NlogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nlog.FieldID, field.TypeInt),
			},
		}
		edge.Schema = nac.schemaConfig.NlogAlert
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NlogID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nac.mutation.AlertIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   nlogalert.AlertTable,
			Columns: []string{nlogalert.AlertColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgalert.FieldID, field.TypeInt),
			},
		}
		edge.Schema = nac.schemaConfig.NlogAlert
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AlertID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NlogAlert.Create().
//		SetNlogID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NlogAlertUpsert) {
//			SetNlogID(v+v).
//		}).
//		Exec(ctx)
func (nac *NlogAlertCreate) OnConflict(opts ...sql.ConflictOption) *NlogAlertUpsertOne {
	nac.conflict = opts
	return &NlogAlertUpsertOne{
		create: nac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NlogAlert.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nac *NlogAlertCreate) OnConflictColumns(columns ...string) *NlogAlertUpsertOne {
	nac.conflict = append(nac.conflict, sql.ConflictColumns(columns...))
	return &NlogAlertUpsertOne{
		create: nac,
	}
}

type (
	// NlogAlertUpsertOne is the builder for "upsert"-ing
	//  one NlogAlert node.
	NlogAlertUpsertOne struct {
		create *NlogAlertCreate
	}

	// NlogAlertUpsert is the "OnConflict" setter.
	NlogAlertUpsert struct {
		*sql.UpdateSet
	}
)

// SetNlogID sets the "nlog_id" field.
func (u *NlogAlertUpsert) SetNlogID(v int) *NlogAlertUpsert {
	u.Set(nlogalert.FieldNlogID, v)
	return u
}

// UpdateNlogID sets the "nlog_id" field to the value that was provided on create.
func (u *NlogAlertUpsert) UpdateNlogID() *NlogAlertUpsert {
	u.SetExcluded(nlogalert.FieldNlogID)
	return u
}

// SetAlertID sets the "alert_id" field.
func (u *NlogAlertUpsert) SetAlertID(v int) *NlogAlertUpsert {
	u.Set(nlogalert.FieldAlertID, v)
	return u
}

// UpdateAlertID sets the "alert_id" field to the value that was provided on create.
func (u *NlogAlertUpsert) UpdateAlertID() *NlogAlertUpsert {
	u.SetExcluded(nlogalert.FieldAlertID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.NlogAlert.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *NlogAlertUpsertOne) UpdateNewValues() *NlogAlertUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(nlogalert.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NlogAlert.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NlogAlertUpsertOne) Ignore() *NlogAlertUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NlogAlertUpsertOne) DoNothing() *NlogAlertUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NlogAlertCreate.OnConflict
// documentation for more info.
func (u *NlogAlertUpsertOne) Update(set func(*NlogAlertUpsert)) *NlogAlertUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NlogAlertUpsert{UpdateSet: update})
	}))
	return u
}

// SetNlogID sets the "nlog_id" field.
func (u *NlogAlertUpsertOne) SetNlogID(v int) *NlogAlertUpsertOne {
	return u.Update(func(s *NlogAlertUpsert) {
		s.SetNlogID(v)
	})
}

// UpdateNlogID sets the "nlog_id" field to the value that was provided on create.
func (u *NlogAlertUpsertOne) UpdateNlogID() *NlogAlertUpsertOne {
	return u.Update(func(s *NlogAlertUpsert) {
		s.UpdateNlogID()
	})
}

// SetAlertID sets the "alert_id" field.
func (u *NlogAlertUpsertOne) SetAlertID(v int) *NlogAlertUpsertOne {
	return u.Update(func(s *NlogAlertUpsert) {
		s.SetAlertID(v)
	})
}

// UpdateAlertID sets the "alert_id" field to the value that was provided on create.
func (u *NlogAlertUpsertOne) UpdateAlertID() *NlogAlertUpsertOne {
	return u.Update(func(s *NlogAlertUpsert) {
		s.UpdateAlertID()
	})
}

// Exec executes the query.
func (u *NlogAlertUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NlogAlertCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NlogAlertUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NlogAlertUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NlogAlertUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NlogAlertCreateBulk is the builder for creating many NlogAlert entities in bulk.
type NlogAlertCreateBulk struct {
	config
	err      error
	builders []*NlogAlertCreate
	conflict []sql.ConflictOption
}

// Save creates the NlogAlert entities in the database.
func (nacb *NlogAlertCreateBulk) Save(ctx context.Context) ([]*NlogAlert, error) {
	if nacb.err != nil {
		return nil, nacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(nacb.builders))
	nodes := make([]*NlogAlert, len(nacb.builders))
	mutators := make([]Mutator, len(nacb.builders))
	for i := range nacb.builders {
		func(i int, root context.Context) {
			builder := nacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NlogAlertMutation)
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
					_, err = mutators[i+1].Mutate(root, nacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = nacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, nacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nacb *NlogAlertCreateBulk) SaveX(ctx context.Context) []*NlogAlert {
	v, err := nacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nacb *NlogAlertCreateBulk) Exec(ctx context.Context) error {
	_, err := nacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nacb *NlogAlertCreateBulk) ExecX(ctx context.Context) {
	if err := nacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NlogAlert.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NlogAlertUpsert) {
//			SetNlogID(v+v).
//		}).
//		Exec(ctx)
func (nacb *NlogAlertCreateBulk) OnConflict(opts ...sql.ConflictOption) *NlogAlertUpsertBulk {
	nacb.conflict = opts
	return &NlogAlertUpsertBulk{
		create: nacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NlogAlert.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nacb *NlogAlertCreateBulk) OnConflictColumns(columns ...string) *NlogAlertUpsertBulk {
	nacb.conflict = append(nacb.conflict, sql.ConflictColumns(columns...))
	return &NlogAlertUpsertBulk{
		create: nacb,
	}
}

// NlogAlertUpsertBulk is the builder for "upsert"-ing
// a bulk of NlogAlert nodes.
type NlogAlertUpsertBulk struct {
	create *NlogAlertCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.NlogAlert.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *NlogAlertUpsertBulk) UpdateNewValues() *NlogAlertUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(nlogalert.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NlogAlert.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NlogAlertUpsertBulk) Ignore() *NlogAlertUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NlogAlertUpsertBulk) DoNothing() *NlogAlertUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NlogAlertCreateBulk.OnConflict
// documentation for more info.
func (u *NlogAlertUpsertBulk) Update(set func(*NlogAlertUpsert)) *NlogAlertUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NlogAlertUpsert{UpdateSet: update})
	}))
	return u
}

// SetNlogID sets the "nlog_id" field.
func (u *NlogAlertUpsertBulk) SetNlogID(v int) *NlogAlertUpsertBulk {
	return u.Update(func(s *NlogAlertUpsert) {
		s.SetNlogID(v)
	})
}

// UpdateNlogID sets the "nlog_id" field to the value that was provided on create.
func (u *NlogAlertUpsertBulk) UpdateNlogID() *NlogAlertUpsertBulk {
	return u.Update(func(s *NlogAlertUpsert) {
		s.UpdateNlogID()
	})
}

// SetAlertID sets the "alert_id" field.
func (u *NlogAlertUpsertBulk) SetAlertID(v int) *NlogAlertUpsertBulk {
	return u.Update(func(s *NlogAlertUpsert) {
		s.SetAlertID(v)
	})
}

// UpdateAlertID sets the "alert_id" field to the value that was provided on create.
func (u *NlogAlertUpsertBulk) UpdateAlertID() *NlogAlertUpsertBulk {
	return u.Update(func(s *NlogAlertUpsert) {
		s.UpdateAlertID()
	})
}

// Exec executes the query.
func (u *NlogAlertUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NlogAlertCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NlogAlertCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NlogAlertUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
