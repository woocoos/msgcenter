// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/msgcenter/ent/orgroleuser"
)

// OrgRoleUserCreate is the builder for creating a OrgRoleUser entity.
type OrgRoleUserCreate struct {
	config
	mutation *OrgRoleUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOrgRoleID sets the "org_role_id" field.
func (oruc *OrgRoleUserCreate) SetOrgRoleID(i int) *OrgRoleUserCreate {
	oruc.mutation.SetOrgRoleID(i)
	return oruc
}

// SetOrgUserID sets the "org_user_id" field.
func (oruc *OrgRoleUserCreate) SetOrgUserID(i int) *OrgRoleUserCreate {
	oruc.mutation.SetOrgUserID(i)
	return oruc
}

// SetOrgID sets the "org_id" field.
func (oruc *OrgRoleUserCreate) SetOrgID(i int) *OrgRoleUserCreate {
	oruc.mutation.SetOrgID(i)
	return oruc
}

// SetUserID sets the "user_id" field.
func (oruc *OrgRoleUserCreate) SetUserID(i int) *OrgRoleUserCreate {
	oruc.mutation.SetUserID(i)
	return oruc
}

// SetID sets the "id" field.
func (oruc *OrgRoleUserCreate) SetID(i int) *OrgRoleUserCreate {
	oruc.mutation.SetID(i)
	return oruc
}

// Mutation returns the OrgRoleUserMutation object of the builder.
func (oruc *OrgRoleUserCreate) Mutation() *OrgRoleUserMutation {
	return oruc.mutation
}

// Save creates the OrgRoleUser in the database.
func (oruc *OrgRoleUserCreate) Save(ctx context.Context) (*OrgRoleUser, error) {
	return withHooks(ctx, oruc.sqlSave, oruc.mutation, oruc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oruc *OrgRoleUserCreate) SaveX(ctx context.Context) *OrgRoleUser {
	v, err := oruc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oruc *OrgRoleUserCreate) Exec(ctx context.Context) error {
	_, err := oruc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oruc *OrgRoleUserCreate) ExecX(ctx context.Context) {
	if err := oruc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oruc *OrgRoleUserCreate) check() error {
	if _, ok := oruc.mutation.OrgRoleID(); !ok {
		return &ValidationError{Name: "org_role_id", err: errors.New(`ent: missing required field "OrgRoleUser.org_role_id"`)}
	}
	if _, ok := oruc.mutation.OrgUserID(); !ok {
		return &ValidationError{Name: "org_user_id", err: errors.New(`ent: missing required field "OrgRoleUser.org_user_id"`)}
	}
	if _, ok := oruc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "OrgRoleUser.org_id"`)}
	}
	if _, ok := oruc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "OrgRoleUser.user_id"`)}
	}
	return nil
}

func (oruc *OrgRoleUserCreate) sqlSave(ctx context.Context) (*OrgRoleUser, error) {
	if err := oruc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oruc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oruc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	oruc.mutation.id = &_node.ID
	oruc.mutation.done = true
	return _node, nil
}

func (oruc *OrgRoleUserCreate) createSpec() (*OrgRoleUser, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgRoleUser{config: oruc.config}
		_spec = sqlgraph.NewCreateSpec(orgroleuser.Table, sqlgraph.NewFieldSpec(orgroleuser.FieldID, field.TypeInt))
	)
	_spec.Schema = oruc.schemaConfig.OrgRoleUser
	_spec.OnConflict = oruc.conflict
	if id, ok := oruc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oruc.mutation.OrgRoleID(); ok {
		_spec.SetField(orgroleuser.FieldOrgRoleID, field.TypeInt, value)
		_node.OrgRoleID = value
	}
	if value, ok := oruc.mutation.OrgUserID(); ok {
		_spec.SetField(orgroleuser.FieldOrgUserID, field.TypeInt, value)
		_node.OrgUserID = value
	}
	if value, ok := oruc.mutation.OrgID(); ok {
		_spec.SetField(orgroleuser.FieldOrgID, field.TypeInt, value)
		_node.OrgID = value
	}
	if value, ok := oruc.mutation.UserID(); ok {
		_spec.SetField(orgroleuser.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrgRoleUser.Create().
//		SetOrgRoleID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgRoleUserUpsert) {
//			SetOrgRoleID(v+v).
//		}).
//		Exec(ctx)
func (oruc *OrgRoleUserCreate) OnConflict(opts ...sql.ConflictOption) *OrgRoleUserUpsertOne {
	oruc.conflict = opts
	return &OrgRoleUserUpsertOne{
		create: oruc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgRoleUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oruc *OrgRoleUserCreate) OnConflictColumns(columns ...string) *OrgRoleUserUpsertOne {
	oruc.conflict = append(oruc.conflict, sql.ConflictColumns(columns...))
	return &OrgRoleUserUpsertOne{
		create: oruc,
	}
}

type (
	// OrgRoleUserUpsertOne is the builder for "upsert"-ing
	//  one OrgRoleUser node.
	OrgRoleUserUpsertOne struct {
		create *OrgRoleUserCreate
	}

	// OrgRoleUserUpsert is the "OnConflict" setter.
	OrgRoleUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetOrgRoleID sets the "org_role_id" field.
func (u *OrgRoleUserUpsert) SetOrgRoleID(v int) *OrgRoleUserUpsert {
	u.Set(orgroleuser.FieldOrgRoleID, v)
	return u
}

// UpdateOrgRoleID sets the "org_role_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsert) UpdateOrgRoleID() *OrgRoleUserUpsert {
	u.SetExcluded(orgroleuser.FieldOrgRoleID)
	return u
}

// AddOrgRoleID adds v to the "org_role_id" field.
func (u *OrgRoleUserUpsert) AddOrgRoleID(v int) *OrgRoleUserUpsert {
	u.Add(orgroleuser.FieldOrgRoleID, v)
	return u
}

// SetOrgUserID sets the "org_user_id" field.
func (u *OrgRoleUserUpsert) SetOrgUserID(v int) *OrgRoleUserUpsert {
	u.Set(orgroleuser.FieldOrgUserID, v)
	return u
}

// UpdateOrgUserID sets the "org_user_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsert) UpdateOrgUserID() *OrgRoleUserUpsert {
	u.SetExcluded(orgroleuser.FieldOrgUserID)
	return u
}

// AddOrgUserID adds v to the "org_user_id" field.
func (u *OrgRoleUserUpsert) AddOrgUserID(v int) *OrgRoleUserUpsert {
	u.Add(orgroleuser.FieldOrgUserID, v)
	return u
}

// SetOrgID sets the "org_id" field.
func (u *OrgRoleUserUpsert) SetOrgID(v int) *OrgRoleUserUpsert {
	u.Set(orgroleuser.FieldOrgID, v)
	return u
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsert) UpdateOrgID() *OrgRoleUserUpsert {
	u.SetExcluded(orgroleuser.FieldOrgID)
	return u
}

// AddOrgID adds v to the "org_id" field.
func (u *OrgRoleUserUpsert) AddOrgID(v int) *OrgRoleUserUpsert {
	u.Add(orgroleuser.FieldOrgID, v)
	return u
}

// SetUserID sets the "user_id" field.
func (u *OrgRoleUserUpsert) SetUserID(v int) *OrgRoleUserUpsert {
	u.Set(orgroleuser.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsert) UpdateUserID() *OrgRoleUserUpsert {
	u.SetExcluded(orgroleuser.FieldUserID)
	return u
}

// AddUserID adds v to the "user_id" field.
func (u *OrgRoleUserUpsert) AddUserID(v int) *OrgRoleUserUpsert {
	u.Add(orgroleuser.FieldUserID, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrgRoleUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orgroleuser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrgRoleUserUpsertOne) UpdateNewValues() *OrgRoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(orgroleuser.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrgRoleUser.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrgRoleUserUpsertOne) Ignore() *OrgRoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgRoleUserUpsertOne) DoNothing() *OrgRoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgRoleUserCreate.OnConflict
// documentation for more info.
func (u *OrgRoleUserUpsertOne) Update(set func(*OrgRoleUserUpsert)) *OrgRoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgRoleUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrgRoleID sets the "org_role_id" field.
func (u *OrgRoleUserUpsertOne) SetOrgRoleID(v int) *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.SetOrgRoleID(v)
	})
}

// AddOrgRoleID adds v to the "org_role_id" field.
func (u *OrgRoleUserUpsertOne) AddOrgRoleID(v int) *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.AddOrgRoleID(v)
	})
}

// UpdateOrgRoleID sets the "org_role_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsertOne) UpdateOrgRoleID() *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.UpdateOrgRoleID()
	})
}

// SetOrgUserID sets the "org_user_id" field.
func (u *OrgRoleUserUpsertOne) SetOrgUserID(v int) *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.SetOrgUserID(v)
	})
}

// AddOrgUserID adds v to the "org_user_id" field.
func (u *OrgRoleUserUpsertOne) AddOrgUserID(v int) *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.AddOrgUserID(v)
	})
}

// UpdateOrgUserID sets the "org_user_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsertOne) UpdateOrgUserID() *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.UpdateOrgUserID()
	})
}

// SetOrgID sets the "org_id" field.
func (u *OrgRoleUserUpsertOne) SetOrgID(v int) *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.SetOrgID(v)
	})
}

// AddOrgID adds v to the "org_id" field.
func (u *OrgRoleUserUpsertOne) AddOrgID(v int) *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.AddOrgID(v)
	})
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsertOne) UpdateOrgID() *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.UpdateOrgID()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrgRoleUserUpsertOne) SetUserID(v int) *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "user_id" field.
func (u *OrgRoleUserUpsertOne) AddUserID(v int) *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsertOne) UpdateUserID() *OrgRoleUserUpsertOne {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *OrgRoleUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrgRoleUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgRoleUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrgRoleUserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrgRoleUserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrgRoleUserCreateBulk is the builder for creating many OrgRoleUser entities in bulk.
type OrgRoleUserCreateBulk struct {
	config
	err      error
	builders []*OrgRoleUserCreate
	conflict []sql.ConflictOption
}

// Save creates the OrgRoleUser entities in the database.
func (orucb *OrgRoleUserCreateBulk) Save(ctx context.Context) ([]*OrgRoleUser, error) {
	if orucb.err != nil {
		return nil, orucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(orucb.builders))
	nodes := make([]*OrgRoleUser, len(orucb.builders))
	mutators := make([]Mutator, len(orucb.builders))
	for i := range orucb.builders {
		func(i int, root context.Context) {
			builder := orucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgRoleUserMutation)
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
					_, err = mutators[i+1].Mutate(root, orucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = orucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, orucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, orucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (orucb *OrgRoleUserCreateBulk) SaveX(ctx context.Context) []*OrgRoleUser {
	v, err := orucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (orucb *OrgRoleUserCreateBulk) Exec(ctx context.Context) error {
	_, err := orucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (orucb *OrgRoleUserCreateBulk) ExecX(ctx context.Context) {
	if err := orucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrgRoleUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgRoleUserUpsert) {
//			SetOrgRoleID(v+v).
//		}).
//		Exec(ctx)
func (orucb *OrgRoleUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrgRoleUserUpsertBulk {
	orucb.conflict = opts
	return &OrgRoleUserUpsertBulk{
		create: orucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgRoleUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (orucb *OrgRoleUserCreateBulk) OnConflictColumns(columns ...string) *OrgRoleUserUpsertBulk {
	orucb.conflict = append(orucb.conflict, sql.ConflictColumns(columns...))
	return &OrgRoleUserUpsertBulk{
		create: orucb,
	}
}

// OrgRoleUserUpsertBulk is the builder for "upsert"-ing
// a bulk of OrgRoleUser nodes.
type OrgRoleUserUpsertBulk struct {
	create *OrgRoleUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrgRoleUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orgroleuser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrgRoleUserUpsertBulk) UpdateNewValues() *OrgRoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(orgroleuser.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrgRoleUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrgRoleUserUpsertBulk) Ignore() *OrgRoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgRoleUserUpsertBulk) DoNothing() *OrgRoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgRoleUserCreateBulk.OnConflict
// documentation for more info.
func (u *OrgRoleUserUpsertBulk) Update(set func(*OrgRoleUserUpsert)) *OrgRoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgRoleUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrgRoleID sets the "org_role_id" field.
func (u *OrgRoleUserUpsertBulk) SetOrgRoleID(v int) *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.SetOrgRoleID(v)
	})
}

// AddOrgRoleID adds v to the "org_role_id" field.
func (u *OrgRoleUserUpsertBulk) AddOrgRoleID(v int) *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.AddOrgRoleID(v)
	})
}

// UpdateOrgRoleID sets the "org_role_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsertBulk) UpdateOrgRoleID() *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.UpdateOrgRoleID()
	})
}

// SetOrgUserID sets the "org_user_id" field.
func (u *OrgRoleUserUpsertBulk) SetOrgUserID(v int) *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.SetOrgUserID(v)
	})
}

// AddOrgUserID adds v to the "org_user_id" field.
func (u *OrgRoleUserUpsertBulk) AddOrgUserID(v int) *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.AddOrgUserID(v)
	})
}

// UpdateOrgUserID sets the "org_user_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsertBulk) UpdateOrgUserID() *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.UpdateOrgUserID()
	})
}

// SetOrgID sets the "org_id" field.
func (u *OrgRoleUserUpsertBulk) SetOrgID(v int) *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.SetOrgID(v)
	})
}

// AddOrgID adds v to the "org_id" field.
func (u *OrgRoleUserUpsertBulk) AddOrgID(v int) *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.AddOrgID(v)
	})
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsertBulk) UpdateOrgID() *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.UpdateOrgID()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrgRoleUserUpsertBulk) SetUserID(v int) *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "user_id" field.
func (u *OrgRoleUserUpsertBulk) AddUserID(v int) *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrgRoleUserUpsertBulk) UpdateUserID() *OrgRoleUserUpsertBulk {
	return u.Update(func(s *OrgRoleUserUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *OrgRoleUserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrgRoleUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrgRoleUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgRoleUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
