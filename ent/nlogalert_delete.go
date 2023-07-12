// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/msgcenter/ent/predicate"

	"github.com/woocoos/msgcenter/ent/internal"
	"github.com/woocoos/msgcenter/ent/nlogalert"
)

// NlogAlertDelete is the builder for deleting a NlogAlert entity.
type NlogAlertDelete struct {
	config
	hooks    []Hook
	mutation *NlogAlertMutation
}

// Where appends a list predicates to the NlogAlertDelete builder.
func (nad *NlogAlertDelete) Where(ps ...predicate.NlogAlert) *NlogAlertDelete {
	nad.mutation.Where(ps...)
	return nad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nad *NlogAlertDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nad.sqlExec, nad.mutation, nad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nad *NlogAlertDelete) ExecX(ctx context.Context) int {
	n, err := nad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nad *NlogAlertDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(nlogalert.Table, sqlgraph.NewFieldSpec(nlogalert.FieldID, field.TypeInt))
	_spec.Node.Schema = nad.schemaConfig.NlogAlert
	ctx = internal.NewSchemaConfigContext(ctx, nad.schemaConfig)
	if ps := nad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nad.mutation.done = true
	return affected, err
}

// NlogAlertDeleteOne is the builder for deleting a single NlogAlert entity.
type NlogAlertDeleteOne struct {
	nad *NlogAlertDelete
}

// Where appends a list predicates to the NlogAlertDelete builder.
func (nado *NlogAlertDeleteOne) Where(ps ...predicate.NlogAlert) *NlogAlertDeleteOne {
	nado.nad.mutation.Where(ps...)
	return nado
}

// Exec executes the deletion query.
func (nado *NlogAlertDeleteOne) Exec(ctx context.Context) error {
	n, err := nado.nad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{nlogalert.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nado *NlogAlertDeleteOne) ExecX(ctx context.Context) {
	if err := nado.Exec(ctx); err != nil {
		panic(err)
	}
}
