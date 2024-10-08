// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/msgcenter/ent/predicate"
	"github.com/woocoos/msgcenter/ent/silence"
	"github.com/woocoos/msgcenter/ent/user"

	"github.com/woocoos/msgcenter/ent/internal"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetPrincipalName sets the "principal_name" field.
func (uu *UserUpdate) SetPrincipalName(s string) *UserUpdate {
	uu.mutation.SetPrincipalName(s)
	return uu
}

// SetNillablePrincipalName sets the "principal_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePrincipalName(s *string) *UserUpdate {
	if s != nil {
		uu.SetPrincipalName(*s)
	}
	return uu
}

// SetDisplayName sets the "display_name" field.
func (uu *UserUpdate) SetDisplayName(s string) *UserUpdate {
	uu.mutation.SetDisplayName(s)
	return uu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDisplayName(s *string) *UserUpdate {
	if s != nil {
		uu.SetDisplayName(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of the "email" field.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetMobile sets the "mobile" field.
func (uu *UserUpdate) SetMobile(s string) *UserUpdate {
	uu.mutation.SetMobile(s)
	return uu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMobile(s *string) *UserUpdate {
	if s != nil {
		uu.SetMobile(*s)
	}
	return uu
}

// ClearMobile clears the value of the "mobile" field.
func (uu *UserUpdate) ClearMobile() *UserUpdate {
	uu.mutation.ClearMobile()
	return uu
}

// AddSilenceIDs adds the "silences" edge to the Silence entity by IDs.
func (uu *UserUpdate) AddSilenceIDs(ids ...int) *UserUpdate {
	uu.mutation.AddSilenceIDs(ids...)
	return uu
}

// AddSilences adds the "silences" edges to the Silence entity.
func (uu *UserUpdate) AddSilences(s ...*Silence) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddSilenceIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearSilences clears all "silences" edges to the Silence entity.
func (uu *UserUpdate) ClearSilences() *UserUpdate {
	uu.mutation.ClearSilences()
	return uu
}

// RemoveSilenceIDs removes the "silences" edge to Silence entities by IDs.
func (uu *UserUpdate) RemoveSilenceIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveSilenceIDs(ids...)
	return uu
}

// RemoveSilences removes "silences" edges to Silence entities.
func (uu *UserUpdate) RemoveSilences(s ...*Silence) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveSilenceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Mobile(); ok {
		if err := user.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "User.mobile": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.PrincipalName(); ok {
		_spec.SetField(user.FieldPrincipalName, field.TypeString, value)
	}
	if value, ok := uu.mutation.DisplayName(); ok {
		_spec.SetField(user.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uu.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uu.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if uu.mutation.MobileCleared() {
		_spec.ClearField(user.FieldMobile, field.TypeString)
	}
	if uu.mutation.SilencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SilencesTable,
			Columns: []string{user.SilencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(silence.FieldID, field.TypeInt),
			},
		}
		edge.Schema = uu.schemaConfig.Silence
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedSilencesIDs(); len(nodes) > 0 && !uu.mutation.SilencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SilencesTable,
			Columns: []string{user.SilencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(silence.FieldID, field.TypeInt),
			},
		}
		edge.Schema = uu.schemaConfig.Silence
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SilencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SilencesTable,
			Columns: []string{user.SilencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(silence.FieldID, field.TypeInt),
			},
		}
		edge.Schema = uu.schemaConfig.Silence
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = uu.schemaConfig.User
	ctx = internal.NewSchemaConfigContext(ctx, uu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetPrincipalName sets the "principal_name" field.
func (uuo *UserUpdateOne) SetPrincipalName(s string) *UserUpdateOne {
	uuo.mutation.SetPrincipalName(s)
	return uuo
}

// SetNillablePrincipalName sets the "principal_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePrincipalName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPrincipalName(*s)
	}
	return uuo
}

// SetDisplayName sets the "display_name" field.
func (uuo *UserUpdateOne) SetDisplayName(s string) *UserUpdateOne {
	uuo.mutation.SetDisplayName(s)
	return uuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDisplayName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDisplayName(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of the "email" field.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetMobile sets the "mobile" field.
func (uuo *UserUpdateOne) SetMobile(s string) *UserUpdateOne {
	uuo.mutation.SetMobile(s)
	return uuo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMobile(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetMobile(*s)
	}
	return uuo
}

// ClearMobile clears the value of the "mobile" field.
func (uuo *UserUpdateOne) ClearMobile() *UserUpdateOne {
	uuo.mutation.ClearMobile()
	return uuo
}

// AddSilenceIDs adds the "silences" edge to the Silence entity by IDs.
func (uuo *UserUpdateOne) AddSilenceIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddSilenceIDs(ids...)
	return uuo
}

// AddSilences adds the "silences" edges to the Silence entity.
func (uuo *UserUpdateOne) AddSilences(s ...*Silence) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddSilenceIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearSilences clears all "silences" edges to the Silence entity.
func (uuo *UserUpdateOne) ClearSilences() *UserUpdateOne {
	uuo.mutation.ClearSilences()
	return uuo
}

// RemoveSilenceIDs removes the "silences" edge to Silence entities by IDs.
func (uuo *UserUpdateOne) RemoveSilenceIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveSilenceIDs(ids...)
	return uuo
}

// RemoveSilences removes "silences" edges to Silence entities.
func (uuo *UserUpdateOne) RemoveSilences(s ...*Silence) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveSilenceIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Mobile(); ok {
		if err := user.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "User.mobile": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.PrincipalName(); ok {
		_spec.SetField(user.FieldPrincipalName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.DisplayName(); ok {
		_spec.SetField(user.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uuo.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uuo.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if uuo.mutation.MobileCleared() {
		_spec.ClearField(user.FieldMobile, field.TypeString)
	}
	if uuo.mutation.SilencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SilencesTable,
			Columns: []string{user.SilencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(silence.FieldID, field.TypeInt),
			},
		}
		edge.Schema = uuo.schemaConfig.Silence
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedSilencesIDs(); len(nodes) > 0 && !uuo.mutation.SilencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SilencesTable,
			Columns: []string{user.SilencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(silence.FieldID, field.TypeInt),
			},
		}
		edge.Schema = uuo.schemaConfig.Silence
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SilencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SilencesTable,
			Columns: []string{user.SilencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(silence.FieldID, field.TypeInt),
			},
		}
		edge.Schema = uuo.schemaConfig.Silence
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = uuo.schemaConfig.User
	ctx = internal.NewSchemaConfigContext(ctx, uuo.schemaConfig)
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
