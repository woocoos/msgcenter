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
	"github.com/woocoos/msgcenter/ent/silence"
	"github.com/woocoos/msgcenter/ent/user"
	"github.com/woocoos/msgcenter/pkg/label"
)

// SilenceCreate is the builder for creating a Silence entity.
type SilenceCreate struct {
	config
	mutation *SilenceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (sc *SilenceCreate) SetCreatedBy(i int) *SilenceCreate {
	sc.mutation.SetCreatedBy(i)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SilenceCreate) SetCreatedAt(t time.Time) *SilenceCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SilenceCreate) SetNillableCreatedAt(t *time.Time) *SilenceCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedBy sets the "updated_by" field.
func (sc *SilenceCreate) SetUpdatedBy(i int) *SilenceCreate {
	sc.mutation.SetUpdatedBy(i)
	return sc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sc *SilenceCreate) SetNillableUpdatedBy(i *int) *SilenceCreate {
	if i != nil {
		sc.SetUpdatedBy(*i)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SilenceCreate) SetUpdatedAt(t time.Time) *SilenceCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SilenceCreate) SetNillableUpdatedAt(t *time.Time) *SilenceCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *SilenceCreate) SetDeletedAt(t time.Time) *SilenceCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *SilenceCreate) SetNillableDeletedAt(t *time.Time) *SilenceCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// SetMatchers sets the "matchers" field.
func (sc *SilenceCreate) SetMatchers(l []label.Matcher) *SilenceCreate {
	sc.mutation.SetMatchers(l)
	return sc
}

// SetStartsAt sets the "starts_at" field.
func (sc *SilenceCreate) SetStartsAt(t time.Time) *SilenceCreate {
	sc.mutation.SetStartsAt(t)
	return sc
}

// SetEndsAt sets the "ends_at" field.
func (sc *SilenceCreate) SetEndsAt(t time.Time) *SilenceCreate {
	sc.mutation.SetEndsAt(t)
	return sc
}

// SetComments sets the "comments" field.
func (sc *SilenceCreate) SetComments(s string) *SilenceCreate {
	sc.mutation.SetComments(s)
	return sc
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (sc *SilenceCreate) SetNillableComments(s *string) *SilenceCreate {
	if s != nil {
		sc.SetComments(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SilenceCreate) SetID(i int) *SilenceCreate {
	sc.mutation.SetID(i)
	return sc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sc *SilenceCreate) SetUserID(id int) *SilenceCreate {
	sc.mutation.SetUserID(id)
	return sc
}

// SetUser sets the "user" edge to the User entity.
func (sc *SilenceCreate) SetUser(u *User) *SilenceCreate {
	return sc.SetUserID(u.ID)
}

// Mutation returns the SilenceMutation object of the builder.
func (sc *SilenceCreate) Mutation() *SilenceMutation {
	return sc.mutation
}

// Save creates the Silence in the database.
func (sc *SilenceCreate) Save(ctx context.Context) (*Silence, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SilenceCreate) SaveX(ctx context.Context) *Silence {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SilenceCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SilenceCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SilenceCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if silence.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized silence.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := silence.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SilenceCreate) check() error {
	if _, ok := sc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Silence.created_by"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Silence.created_at"`)}
	}
	if _, ok := sc.mutation.StartsAt(); !ok {
		return &ValidationError{Name: "starts_at", err: errors.New(`ent: missing required field "Silence.starts_at"`)}
	}
	if _, ok := sc.mutation.EndsAt(); !ok {
		return &ValidationError{Name: "ends_at", err: errors.New(`ent: missing required field "Silence.ends_at"`)}
	}
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Silence.user"`)}
	}
	return nil
}

func (sc *SilenceCreate) sqlSave(ctx context.Context) (*Silence, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SilenceCreate) createSpec() (*Silence, *sqlgraph.CreateSpec) {
	var (
		_node = &Silence{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(silence.Table, sqlgraph.NewFieldSpec(silence.FieldID, field.TypeInt))
	)
	_spec.Schema = sc.schemaConfig.Silence
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(silence.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedBy(); ok {
		_spec.SetField(silence.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(silence.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.SetField(silence.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := sc.mutation.Matchers(); ok {
		_spec.SetField(silence.FieldMatchers, field.TypeJSON, value)
		_node.Matchers = value
	}
	if value, ok := sc.mutation.StartsAt(); ok {
		_spec.SetField(silence.FieldStartsAt, field.TypeTime, value)
		_node.StartsAt = value
	}
	if value, ok := sc.mutation.EndsAt(); ok {
		_spec.SetField(silence.FieldEndsAt, field.TypeTime, value)
		_node.EndsAt = value
	}
	if value, ok := sc.mutation.Comments(); ok {
		_spec.SetField(silence.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   silence.UserTable,
			Columns: []string{silence.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = sc.schemaConfig.Silence
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreatedBy = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Silence.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SilenceUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (sc *SilenceCreate) OnConflict(opts ...sql.ConflictOption) *SilenceUpsertOne {
	sc.conflict = opts
	return &SilenceUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Silence.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SilenceCreate) OnConflictColumns(columns ...string) *SilenceUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SilenceUpsertOne{
		create: sc,
	}
}

type (
	// SilenceUpsertOne is the builder for "upsert"-ing
	//  one Silence node.
	SilenceUpsertOne struct {
		create *SilenceCreate
	}

	// SilenceUpsert is the "OnConflict" setter.
	SilenceUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *SilenceUpsert) SetUpdatedBy(v int) *SilenceUpsert {
	u.Set(silence.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *SilenceUpsert) UpdateUpdatedBy() *SilenceUpsert {
	u.SetExcluded(silence.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *SilenceUpsert) AddUpdatedBy(v int) *SilenceUpsert {
	u.Add(silence.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *SilenceUpsert) ClearUpdatedBy() *SilenceUpsert {
	u.SetNull(silence.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SilenceUpsert) SetUpdatedAt(v time.Time) *SilenceUpsert {
	u.Set(silence.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SilenceUpsert) UpdateUpdatedAt() *SilenceUpsert {
	u.SetExcluded(silence.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SilenceUpsert) ClearUpdatedAt() *SilenceUpsert {
	u.SetNull(silence.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SilenceUpsert) SetDeletedAt(v time.Time) *SilenceUpsert {
	u.Set(silence.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SilenceUpsert) UpdateDeletedAt() *SilenceUpsert {
	u.SetExcluded(silence.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SilenceUpsert) ClearDeletedAt() *SilenceUpsert {
	u.SetNull(silence.FieldDeletedAt)
	return u
}

// SetMatchers sets the "matchers" field.
func (u *SilenceUpsert) SetMatchers(v []label.Matcher) *SilenceUpsert {
	u.Set(silence.FieldMatchers, v)
	return u
}

// UpdateMatchers sets the "matchers" field to the value that was provided on create.
func (u *SilenceUpsert) UpdateMatchers() *SilenceUpsert {
	u.SetExcluded(silence.FieldMatchers)
	return u
}

// ClearMatchers clears the value of the "matchers" field.
func (u *SilenceUpsert) ClearMatchers() *SilenceUpsert {
	u.SetNull(silence.FieldMatchers)
	return u
}

// SetStartsAt sets the "starts_at" field.
func (u *SilenceUpsert) SetStartsAt(v time.Time) *SilenceUpsert {
	u.Set(silence.FieldStartsAt, v)
	return u
}

// UpdateStartsAt sets the "starts_at" field to the value that was provided on create.
func (u *SilenceUpsert) UpdateStartsAt() *SilenceUpsert {
	u.SetExcluded(silence.FieldStartsAt)
	return u
}

// SetEndsAt sets the "ends_at" field.
func (u *SilenceUpsert) SetEndsAt(v time.Time) *SilenceUpsert {
	u.Set(silence.FieldEndsAt, v)
	return u
}

// UpdateEndsAt sets the "ends_at" field to the value that was provided on create.
func (u *SilenceUpsert) UpdateEndsAt() *SilenceUpsert {
	u.SetExcluded(silence.FieldEndsAt)
	return u
}

// SetComments sets the "comments" field.
func (u *SilenceUpsert) SetComments(v string) *SilenceUpsert {
	u.Set(silence.FieldComments, v)
	return u
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *SilenceUpsert) UpdateComments() *SilenceUpsert {
	u.SetExcluded(silence.FieldComments)
	return u
}

// ClearComments clears the value of the "comments" field.
func (u *SilenceUpsert) ClearComments() *SilenceUpsert {
	u.SetNull(silence.FieldComments)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Silence.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(silence.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SilenceUpsertOne) UpdateNewValues() *SilenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(silence.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(silence.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(silence.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Silence.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SilenceUpsertOne) Ignore() *SilenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SilenceUpsertOne) DoNothing() *SilenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SilenceCreate.OnConflict
// documentation for more info.
func (u *SilenceUpsertOne) Update(set func(*SilenceUpsert)) *SilenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SilenceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *SilenceUpsertOne) SetUpdatedBy(v int) *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *SilenceUpsertOne) AddUpdatedBy(v int) *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *SilenceUpsertOne) UpdateUpdatedBy() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *SilenceUpsertOne) ClearUpdatedBy() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SilenceUpsertOne) SetUpdatedAt(v time.Time) *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SilenceUpsertOne) UpdateUpdatedAt() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SilenceUpsertOne) ClearUpdatedAt() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SilenceUpsertOne) SetDeletedAt(v time.Time) *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SilenceUpsertOne) UpdateDeletedAt() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SilenceUpsertOne) ClearDeletedAt() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.ClearDeletedAt()
	})
}

// SetMatchers sets the "matchers" field.
func (u *SilenceUpsertOne) SetMatchers(v []label.Matcher) *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.SetMatchers(v)
	})
}

// UpdateMatchers sets the "matchers" field to the value that was provided on create.
func (u *SilenceUpsertOne) UpdateMatchers() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateMatchers()
	})
}

// ClearMatchers clears the value of the "matchers" field.
func (u *SilenceUpsertOne) ClearMatchers() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.ClearMatchers()
	})
}

// SetStartsAt sets the "starts_at" field.
func (u *SilenceUpsertOne) SetStartsAt(v time.Time) *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.SetStartsAt(v)
	})
}

// UpdateStartsAt sets the "starts_at" field to the value that was provided on create.
func (u *SilenceUpsertOne) UpdateStartsAt() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateStartsAt()
	})
}

// SetEndsAt sets the "ends_at" field.
func (u *SilenceUpsertOne) SetEndsAt(v time.Time) *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.SetEndsAt(v)
	})
}

// UpdateEndsAt sets the "ends_at" field to the value that was provided on create.
func (u *SilenceUpsertOne) UpdateEndsAt() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateEndsAt()
	})
}

// SetComments sets the "comments" field.
func (u *SilenceUpsertOne) SetComments(v string) *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *SilenceUpsertOne) UpdateComments() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *SilenceUpsertOne) ClearComments() *SilenceUpsertOne {
	return u.Update(func(s *SilenceUpsert) {
		s.ClearComments()
	})
}

// Exec executes the query.
func (u *SilenceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SilenceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SilenceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SilenceUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SilenceUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SilenceCreateBulk is the builder for creating many Silence entities in bulk.
type SilenceCreateBulk struct {
	config
	builders []*SilenceCreate
	conflict []sql.ConflictOption
}

// Save creates the Silence entities in the database.
func (scb *SilenceCreateBulk) Save(ctx context.Context) ([]*Silence, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Silence, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SilenceMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SilenceCreateBulk) SaveX(ctx context.Context) []*Silence {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SilenceCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SilenceCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Silence.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SilenceUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (scb *SilenceCreateBulk) OnConflict(opts ...sql.ConflictOption) *SilenceUpsertBulk {
	scb.conflict = opts
	return &SilenceUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Silence.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SilenceCreateBulk) OnConflictColumns(columns ...string) *SilenceUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SilenceUpsertBulk{
		create: scb,
	}
}

// SilenceUpsertBulk is the builder for "upsert"-ing
// a bulk of Silence nodes.
type SilenceUpsertBulk struct {
	create *SilenceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Silence.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(silence.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SilenceUpsertBulk) UpdateNewValues() *SilenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(silence.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(silence.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(silence.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Silence.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SilenceUpsertBulk) Ignore() *SilenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SilenceUpsertBulk) DoNothing() *SilenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SilenceCreateBulk.OnConflict
// documentation for more info.
func (u *SilenceUpsertBulk) Update(set func(*SilenceUpsert)) *SilenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SilenceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *SilenceUpsertBulk) SetUpdatedBy(v int) *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *SilenceUpsertBulk) AddUpdatedBy(v int) *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *SilenceUpsertBulk) UpdateUpdatedBy() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *SilenceUpsertBulk) ClearUpdatedBy() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SilenceUpsertBulk) SetUpdatedAt(v time.Time) *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SilenceUpsertBulk) UpdateUpdatedAt() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SilenceUpsertBulk) ClearUpdatedAt() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SilenceUpsertBulk) SetDeletedAt(v time.Time) *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SilenceUpsertBulk) UpdateDeletedAt() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SilenceUpsertBulk) ClearDeletedAt() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.ClearDeletedAt()
	})
}

// SetMatchers sets the "matchers" field.
func (u *SilenceUpsertBulk) SetMatchers(v []label.Matcher) *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.SetMatchers(v)
	})
}

// UpdateMatchers sets the "matchers" field to the value that was provided on create.
func (u *SilenceUpsertBulk) UpdateMatchers() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateMatchers()
	})
}

// ClearMatchers clears the value of the "matchers" field.
func (u *SilenceUpsertBulk) ClearMatchers() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.ClearMatchers()
	})
}

// SetStartsAt sets the "starts_at" field.
func (u *SilenceUpsertBulk) SetStartsAt(v time.Time) *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.SetStartsAt(v)
	})
}

// UpdateStartsAt sets the "starts_at" field to the value that was provided on create.
func (u *SilenceUpsertBulk) UpdateStartsAt() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateStartsAt()
	})
}

// SetEndsAt sets the "ends_at" field.
func (u *SilenceUpsertBulk) SetEndsAt(v time.Time) *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.SetEndsAt(v)
	})
}

// UpdateEndsAt sets the "ends_at" field to the value that was provided on create.
func (u *SilenceUpsertBulk) UpdateEndsAt() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateEndsAt()
	})
}

// SetComments sets the "comments" field.
func (u *SilenceUpsertBulk) SetComments(v string) *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *SilenceUpsertBulk) UpdateComments() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *SilenceUpsertBulk) ClearComments() *SilenceUpsertBulk {
	return u.Update(func(s *SilenceUpsert) {
		s.ClearComments()
	})
}

// Exec executes the query.
func (u *SilenceUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SilenceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SilenceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SilenceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}