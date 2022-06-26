// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent/groupmember"
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupMemberDelete is the builder for deleting a GroupMember entity.
type GroupMemberDelete struct {
	config
	hooks    []Hook
	mutation *GroupMemberMutation
}

// Where appends a list predicates to the GroupMemberDelete builder.
func (gmd *GroupMemberDelete) Where(ps ...predicate.GroupMember) *GroupMemberDelete {
	gmd.mutation.Where(ps...)
	return gmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gmd *GroupMemberDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gmd.hooks) == 0 {
		affected, err = gmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gmd.mutation = mutation
			affected, err = gmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gmd.hooks) - 1; i >= 0; i-- {
			if gmd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmd *GroupMemberDelete) ExecX(ctx context.Context) int {
	n, err := gmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gmd *GroupMemberDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: groupmember.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupmember.FieldID,
			},
		},
	}
	if ps := gmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gmd.driver, _spec)
}

// GroupMemberDeleteOne is the builder for deleting a single GroupMember entity.
type GroupMemberDeleteOne struct {
	gmd *GroupMemberDelete
}

// Exec executes the deletion query.
func (gmdo *GroupMemberDeleteOne) Exec(ctx context.Context) error {
	n, err := gmdo.gmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{groupmember.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gmdo *GroupMemberDeleteOne) ExecX(ctx context.Context) {
	gmdo.gmd.ExecX(ctx)
}
