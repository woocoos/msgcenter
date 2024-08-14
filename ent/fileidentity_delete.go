// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/msgcenter/ent/predicate"

	"github.com/woocoos/msgcenter/ent/fileidentity"
	"github.com/woocoos/msgcenter/ent/internal"
)

// FileIdentityDelete is the builder for deleting a FileIdentity entity.
type FileIdentityDelete struct {
	config
	hooks    []Hook
	mutation *FileIdentityMutation
}

// Where appends a list predicates to the FileIdentityDelete builder.
func (fid *FileIdentityDelete) Where(ps ...predicate.FileIdentity) *FileIdentityDelete {
	fid.mutation.Where(ps...)
	return fid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fid *FileIdentityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fid.sqlExec, fid.mutation, fid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fid *FileIdentityDelete) ExecX(ctx context.Context) int {
	n, err := fid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fid *FileIdentityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(fileidentity.Table, sqlgraph.NewFieldSpec(fileidentity.FieldID, field.TypeInt))
	_spec.Node.Schema = fid.schemaConfig.FileIdentity
	ctx = internal.NewSchemaConfigContext(ctx, fid.schemaConfig)
	if ps := fid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fid.mutation.done = true
	return affected, err
}

// FileIdentityDeleteOne is the builder for deleting a single FileIdentity entity.
type FileIdentityDeleteOne struct {
	fid *FileIdentityDelete
}

// Where appends a list predicates to the FileIdentityDelete builder.
func (fido *FileIdentityDeleteOne) Where(ps ...predicate.FileIdentity) *FileIdentityDeleteOne {
	fido.fid.mutation.Where(ps...)
	return fido
}

// Exec executes the deletion query.
func (fido *FileIdentityDeleteOne) Exec(ctx context.Context) error {
	n, err := fido.fid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fileidentity.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fido *FileIdentityDeleteOne) ExecX(ctx context.Context) {
	if err := fido.Exec(ctx); err != nil {
		panic(err)
	}
}
