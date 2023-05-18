// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/ent/predicate"
)

// MsgChannelDelete is the builder for deleting a MsgChannel entity.
type MsgChannelDelete struct {
	config
	hooks    []Hook
	mutation *MsgChannelMutation
}

// Where appends a list predicates to the MsgChannelDelete builder.
func (mcd *MsgChannelDelete) Where(ps ...predicate.MsgChannel) *MsgChannelDelete {
	mcd.mutation.Where(ps...)
	return mcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mcd *MsgChannelDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mcd.sqlExec, mcd.mutation, mcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mcd *MsgChannelDelete) ExecX(ctx context.Context) int {
	n, err := mcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mcd *MsgChannelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(msgchannel.Table, sqlgraph.NewFieldSpec(msgchannel.FieldID, field.TypeInt))
	if ps := mcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mcd.mutation.done = true
	return affected, err
}

// MsgChannelDeleteOne is the builder for deleting a single MsgChannel entity.
type MsgChannelDeleteOne struct {
	mcd *MsgChannelDelete
}

// Where appends a list predicates to the MsgChannelDelete builder.
func (mcdo *MsgChannelDeleteOne) Where(ps ...predicate.MsgChannel) *MsgChannelDeleteOne {
	mcdo.mcd.mutation.Where(ps...)
	return mcdo
}

// Exec executes the deletion query.
func (mcdo *MsgChannelDeleteOne) Exec(ctx context.Context) error {
	n, err := mcdo.mcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{msgchannel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mcdo *MsgChannelDeleteOne) ExecX(ctx context.Context) {
	if err := mcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
