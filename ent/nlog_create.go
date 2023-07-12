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
	"github.com/woocoos/msgcenter/pkg/profile"
)

// NlogCreate is the builder for creating a Nlog entity.
type NlogCreate struct {
	config
	mutation *NlogMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTenantID sets the "tenant_id" field.
func (nc *NlogCreate) SetTenantID(i int) *NlogCreate {
	nc.mutation.SetTenantID(i)
	return nc
}

// SetGroupKey sets the "group_key" field.
func (nc *NlogCreate) SetGroupKey(s string) *NlogCreate {
	nc.mutation.SetGroupKey(s)
	return nc
}

// SetReceiver sets the "receiver" field.
func (nc *NlogCreate) SetReceiver(s string) *NlogCreate {
	nc.mutation.SetReceiver(s)
	return nc
}

// SetReceiverType sets the "receiver_type" field.
func (nc *NlogCreate) SetReceiverType(pt profile.ReceiverType) *NlogCreate {
	nc.mutation.SetReceiverType(pt)
	return nc
}

// SetIdx sets the "idx" field.
func (nc *NlogCreate) SetIdx(i int) *NlogCreate {
	nc.mutation.SetIdx(i)
	return nc
}

// SetSendAt sets the "send_at" field.
func (nc *NlogCreate) SetSendAt(t time.Time) *NlogCreate {
	nc.mutation.SetSendAt(t)
	return nc
}

// SetCreatedAt sets the "created_at" field.
func (nc *NlogCreate) SetCreatedAt(t time.Time) *NlogCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NlogCreate) SetNillableCreatedAt(t *time.Time) *NlogCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NlogCreate) SetUpdatedAt(t time.Time) *NlogCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NlogCreate) SetNillableUpdatedAt(t *time.Time) *NlogCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetExpiresAt sets the "expires_at" field.
func (nc *NlogCreate) SetExpiresAt(t time.Time) *NlogCreate {
	nc.mutation.SetExpiresAt(t)
	return nc
}

// SetID sets the "id" field.
func (nc *NlogCreate) SetID(i int) *NlogCreate {
	nc.mutation.SetID(i)
	return nc
}

// AddAlertIDs adds the "alerts" edge to the MsgAlert entity by IDs.
func (nc *NlogCreate) AddAlertIDs(ids ...int) *NlogCreate {
	nc.mutation.AddAlertIDs(ids...)
	return nc
}

// AddAlerts adds the "alerts" edges to the MsgAlert entity.
func (nc *NlogCreate) AddAlerts(m ...*MsgAlert) *NlogCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return nc.AddAlertIDs(ids...)
}

// AddNlogAlertIDs adds the "nlog_alert" edge to the NlogAlert entity by IDs.
func (nc *NlogCreate) AddNlogAlertIDs(ids ...int) *NlogCreate {
	nc.mutation.AddNlogAlertIDs(ids...)
	return nc
}

// AddNlogAlert adds the "nlog_alert" edges to the NlogAlert entity.
func (nc *NlogCreate) AddNlogAlert(n ...*NlogAlert) *NlogCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nc.AddNlogAlertIDs(ids...)
}

// Mutation returns the NlogMutation object of the builder.
func (nc *NlogCreate) Mutation() *NlogMutation {
	return nc.mutation
}

// Save creates the Nlog in the database.
func (nc *NlogCreate) Save(ctx context.Context) (*Nlog, error) {
	if err := nc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NlogCreate) SaveX(ctx context.Context) *Nlog {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NlogCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NlogCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NlogCreate) defaults() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		if nlog.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized nlog.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := nlog.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (nc *NlogCreate) check() error {
	if _, ok := nc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "Nlog.tenant_id"`)}
	}
	if _, ok := nc.mutation.GroupKey(); !ok {
		return &ValidationError{Name: "group_key", err: errors.New(`ent: missing required field "Nlog.group_key"`)}
	}
	if _, ok := nc.mutation.Receiver(); !ok {
		return &ValidationError{Name: "receiver", err: errors.New(`ent: missing required field "Nlog.receiver"`)}
	}
	if _, ok := nc.mutation.ReceiverType(); !ok {
		return &ValidationError{Name: "receiver_type", err: errors.New(`ent: missing required field "Nlog.receiver_type"`)}
	}
	if v, ok := nc.mutation.ReceiverType(); ok {
		if err := nlog.ReceiverTypeValidator(v); err != nil {
			return &ValidationError{Name: "receiver_type", err: fmt.Errorf(`ent: validator failed for field "Nlog.receiver_type": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Idx(); !ok {
		return &ValidationError{Name: "idx", err: errors.New(`ent: missing required field "Nlog.idx"`)}
	}
	if _, ok := nc.mutation.SendAt(); !ok {
		return &ValidationError{Name: "send_at", err: errors.New(`ent: missing required field "Nlog.send_at"`)}
	}
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Nlog.created_at"`)}
	}
	if _, ok := nc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "Nlog.expires_at"`)}
	}
	return nil
}

func (nc *NlogCreate) sqlSave(ctx context.Context) (*Nlog, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NlogCreate) createSpec() (*Nlog, *sqlgraph.CreateSpec) {
	var (
		_node = &Nlog{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(nlog.Table, sqlgraph.NewFieldSpec(nlog.FieldID, field.TypeInt))
	)
	_spec.Schema = nc.schemaConfig.Nlog
	_spec.OnConflict = nc.conflict
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nc.mutation.TenantID(); ok {
		_spec.SetField(nlog.FieldTenantID, field.TypeInt, value)
		_node.TenantID = value
	}
	if value, ok := nc.mutation.GroupKey(); ok {
		_spec.SetField(nlog.FieldGroupKey, field.TypeString, value)
		_node.GroupKey = value
	}
	if value, ok := nc.mutation.Receiver(); ok {
		_spec.SetField(nlog.FieldReceiver, field.TypeString, value)
		_node.Receiver = value
	}
	if value, ok := nc.mutation.ReceiverType(); ok {
		_spec.SetField(nlog.FieldReceiverType, field.TypeEnum, value)
		_node.ReceiverType = value
	}
	if value, ok := nc.mutation.Idx(); ok {
		_spec.SetField(nlog.FieldIdx, field.TypeInt, value)
		_node.Idx = value
	}
	if value, ok := nc.mutation.SendAt(); ok {
		_spec.SetField(nlog.FieldSendAt, field.TypeTime, value)
		_node.SendAt = value
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(nlog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.SetField(nlog.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.ExpiresAt(); ok {
		_spec.SetField(nlog.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if nodes := nc.mutation.AlertsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   nlog.AlertsTable,
			Columns: nlog.AlertsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msgalert.FieldID, field.TypeInt),
			},
		}
		edge.Schema = nc.schemaConfig.NlogAlert
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NlogAlertCreate{config: nc.config, mutation: newNlogAlertMutation(nc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.NlogAlertIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   nlog.NlogAlertTable,
			Columns: []string{nlog.NlogAlertColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nlogalert.FieldID, field.TypeInt),
			},
		}
		edge.Schema = nc.schemaConfig.NlogAlert
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
//	client.Nlog.Create().
//		SetTenantID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NlogUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
func (nc *NlogCreate) OnConflict(opts ...sql.ConflictOption) *NlogUpsertOne {
	nc.conflict = opts
	return &NlogUpsertOne{
		create: nc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Nlog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nc *NlogCreate) OnConflictColumns(columns ...string) *NlogUpsertOne {
	nc.conflict = append(nc.conflict, sql.ConflictColumns(columns...))
	return &NlogUpsertOne{
		create: nc,
	}
}

type (
	// NlogUpsertOne is the builder for "upsert"-ing
	//  one Nlog node.
	NlogUpsertOne struct {
		create *NlogCreate
	}

	// NlogUpsert is the "OnConflict" setter.
	NlogUpsert struct {
		*sql.UpdateSet
	}
)

// SetGroupKey sets the "group_key" field.
func (u *NlogUpsert) SetGroupKey(v string) *NlogUpsert {
	u.Set(nlog.FieldGroupKey, v)
	return u
}

// UpdateGroupKey sets the "group_key" field to the value that was provided on create.
func (u *NlogUpsert) UpdateGroupKey() *NlogUpsert {
	u.SetExcluded(nlog.FieldGroupKey)
	return u
}

// SetReceiver sets the "receiver" field.
func (u *NlogUpsert) SetReceiver(v string) *NlogUpsert {
	u.Set(nlog.FieldReceiver, v)
	return u
}

// UpdateReceiver sets the "receiver" field to the value that was provided on create.
func (u *NlogUpsert) UpdateReceiver() *NlogUpsert {
	u.SetExcluded(nlog.FieldReceiver)
	return u
}

// SetReceiverType sets the "receiver_type" field.
func (u *NlogUpsert) SetReceiverType(v profile.ReceiverType) *NlogUpsert {
	u.Set(nlog.FieldReceiverType, v)
	return u
}

// UpdateReceiverType sets the "receiver_type" field to the value that was provided on create.
func (u *NlogUpsert) UpdateReceiverType() *NlogUpsert {
	u.SetExcluded(nlog.FieldReceiverType)
	return u
}

// SetIdx sets the "idx" field.
func (u *NlogUpsert) SetIdx(v int) *NlogUpsert {
	u.Set(nlog.FieldIdx, v)
	return u
}

// UpdateIdx sets the "idx" field to the value that was provided on create.
func (u *NlogUpsert) UpdateIdx() *NlogUpsert {
	u.SetExcluded(nlog.FieldIdx)
	return u
}

// AddIdx adds v to the "idx" field.
func (u *NlogUpsert) AddIdx(v int) *NlogUpsert {
	u.Add(nlog.FieldIdx, v)
	return u
}

// SetSendAt sets the "send_at" field.
func (u *NlogUpsert) SetSendAt(v time.Time) *NlogUpsert {
	u.Set(nlog.FieldSendAt, v)
	return u
}

// UpdateSendAt sets the "send_at" field to the value that was provided on create.
func (u *NlogUpsert) UpdateSendAt() *NlogUpsert {
	u.SetExcluded(nlog.FieldSendAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NlogUpsert) SetUpdatedAt(v time.Time) *NlogUpsert {
	u.Set(nlog.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NlogUpsert) UpdateUpdatedAt() *NlogUpsert {
	u.SetExcluded(nlog.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *NlogUpsert) ClearUpdatedAt() *NlogUpsert {
	u.SetNull(nlog.FieldUpdatedAt)
	return u
}

// SetExpiresAt sets the "expires_at" field.
func (u *NlogUpsert) SetExpiresAt(v time.Time) *NlogUpsert {
	u.Set(nlog.FieldExpiresAt, v)
	return u
}

// UpdateExpiresAt sets the "expires_at" field to the value that was provided on create.
func (u *NlogUpsert) UpdateExpiresAt() *NlogUpsert {
	u.SetExcluded(nlog.FieldExpiresAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Nlog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(nlog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NlogUpsertOne) UpdateNewValues() *NlogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(nlog.FieldID)
		}
		if _, exists := u.create.mutation.TenantID(); exists {
			s.SetIgnore(nlog.FieldTenantID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(nlog.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Nlog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NlogUpsertOne) Ignore() *NlogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NlogUpsertOne) DoNothing() *NlogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NlogCreate.OnConflict
// documentation for more info.
func (u *NlogUpsertOne) Update(set func(*NlogUpsert)) *NlogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NlogUpsert{UpdateSet: update})
	}))
	return u
}

// SetGroupKey sets the "group_key" field.
func (u *NlogUpsertOne) SetGroupKey(v string) *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.SetGroupKey(v)
	})
}

// UpdateGroupKey sets the "group_key" field to the value that was provided on create.
func (u *NlogUpsertOne) UpdateGroupKey() *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateGroupKey()
	})
}

// SetReceiver sets the "receiver" field.
func (u *NlogUpsertOne) SetReceiver(v string) *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.SetReceiver(v)
	})
}

// UpdateReceiver sets the "receiver" field to the value that was provided on create.
func (u *NlogUpsertOne) UpdateReceiver() *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateReceiver()
	})
}

// SetReceiverType sets the "receiver_type" field.
func (u *NlogUpsertOne) SetReceiverType(v profile.ReceiverType) *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.SetReceiverType(v)
	})
}

// UpdateReceiverType sets the "receiver_type" field to the value that was provided on create.
func (u *NlogUpsertOne) UpdateReceiverType() *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateReceiverType()
	})
}

// SetIdx sets the "idx" field.
func (u *NlogUpsertOne) SetIdx(v int) *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.SetIdx(v)
	})
}

// AddIdx adds v to the "idx" field.
func (u *NlogUpsertOne) AddIdx(v int) *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.AddIdx(v)
	})
}

// UpdateIdx sets the "idx" field to the value that was provided on create.
func (u *NlogUpsertOne) UpdateIdx() *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateIdx()
	})
}

// SetSendAt sets the "send_at" field.
func (u *NlogUpsertOne) SetSendAt(v time.Time) *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.SetSendAt(v)
	})
}

// UpdateSendAt sets the "send_at" field to the value that was provided on create.
func (u *NlogUpsertOne) UpdateSendAt() *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateSendAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NlogUpsertOne) SetUpdatedAt(v time.Time) *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NlogUpsertOne) UpdateUpdatedAt() *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *NlogUpsertOne) ClearUpdatedAt() *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetExpiresAt sets the "expires_at" field.
func (u *NlogUpsertOne) SetExpiresAt(v time.Time) *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.SetExpiresAt(v)
	})
}

// UpdateExpiresAt sets the "expires_at" field to the value that was provided on create.
func (u *NlogUpsertOne) UpdateExpiresAt() *NlogUpsertOne {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateExpiresAt()
	})
}

// Exec executes the query.
func (u *NlogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NlogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NlogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NlogUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NlogUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NlogCreateBulk is the builder for creating many Nlog entities in bulk.
type NlogCreateBulk struct {
	config
	builders []*NlogCreate
	conflict []sql.ConflictOption
}

// Save creates the Nlog entities in the database.
func (ncb *NlogCreateBulk) Save(ctx context.Context) ([]*Nlog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Nlog, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NlogMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ncb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NlogCreateBulk) SaveX(ctx context.Context) []*Nlog {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NlogCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NlogCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Nlog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NlogUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
func (ncb *NlogCreateBulk) OnConflict(opts ...sql.ConflictOption) *NlogUpsertBulk {
	ncb.conflict = opts
	return &NlogUpsertBulk{
		create: ncb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Nlog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ncb *NlogCreateBulk) OnConflictColumns(columns ...string) *NlogUpsertBulk {
	ncb.conflict = append(ncb.conflict, sql.ConflictColumns(columns...))
	return &NlogUpsertBulk{
		create: ncb,
	}
}

// NlogUpsertBulk is the builder for "upsert"-ing
// a bulk of Nlog nodes.
type NlogUpsertBulk struct {
	create *NlogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Nlog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(nlog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NlogUpsertBulk) UpdateNewValues() *NlogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(nlog.FieldID)
			}
			if _, exists := b.mutation.TenantID(); exists {
				s.SetIgnore(nlog.FieldTenantID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(nlog.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Nlog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NlogUpsertBulk) Ignore() *NlogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NlogUpsertBulk) DoNothing() *NlogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NlogCreateBulk.OnConflict
// documentation for more info.
func (u *NlogUpsertBulk) Update(set func(*NlogUpsert)) *NlogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NlogUpsert{UpdateSet: update})
	}))
	return u
}

// SetGroupKey sets the "group_key" field.
func (u *NlogUpsertBulk) SetGroupKey(v string) *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.SetGroupKey(v)
	})
}

// UpdateGroupKey sets the "group_key" field to the value that was provided on create.
func (u *NlogUpsertBulk) UpdateGroupKey() *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateGroupKey()
	})
}

// SetReceiver sets the "receiver" field.
func (u *NlogUpsertBulk) SetReceiver(v string) *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.SetReceiver(v)
	})
}

// UpdateReceiver sets the "receiver" field to the value that was provided on create.
func (u *NlogUpsertBulk) UpdateReceiver() *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateReceiver()
	})
}

// SetReceiverType sets the "receiver_type" field.
func (u *NlogUpsertBulk) SetReceiverType(v profile.ReceiverType) *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.SetReceiverType(v)
	})
}

// UpdateReceiverType sets the "receiver_type" field to the value that was provided on create.
func (u *NlogUpsertBulk) UpdateReceiverType() *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateReceiverType()
	})
}

// SetIdx sets the "idx" field.
func (u *NlogUpsertBulk) SetIdx(v int) *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.SetIdx(v)
	})
}

// AddIdx adds v to the "idx" field.
func (u *NlogUpsertBulk) AddIdx(v int) *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.AddIdx(v)
	})
}

// UpdateIdx sets the "idx" field to the value that was provided on create.
func (u *NlogUpsertBulk) UpdateIdx() *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateIdx()
	})
}

// SetSendAt sets the "send_at" field.
func (u *NlogUpsertBulk) SetSendAt(v time.Time) *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.SetSendAt(v)
	})
}

// UpdateSendAt sets the "send_at" field to the value that was provided on create.
func (u *NlogUpsertBulk) UpdateSendAt() *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateSendAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NlogUpsertBulk) SetUpdatedAt(v time.Time) *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NlogUpsertBulk) UpdateUpdatedAt() *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *NlogUpsertBulk) ClearUpdatedAt() *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetExpiresAt sets the "expires_at" field.
func (u *NlogUpsertBulk) SetExpiresAt(v time.Time) *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.SetExpiresAt(v)
	})
}

// UpdateExpiresAt sets the "expires_at" field to the value that was provided on create.
func (u *NlogUpsertBulk) UpdateExpiresAt() *NlogUpsertBulk {
	return u.Update(func(s *NlogUpsert) {
		s.UpdateExpiresAt()
	})
}

// Exec executes the query.
func (u *NlogUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NlogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NlogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NlogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
