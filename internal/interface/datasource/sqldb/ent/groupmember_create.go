// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent/groupmember"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GroupMemberCreate is the builder for creating a GroupMember entity.
type GroupMemberCreate struct {
	config
	mutation *GroupMemberMutation
	hooks    []Hook
}

// SetGroupID sets the "group_id" field.
func (gmc *GroupMemberCreate) SetGroupID(i int) *GroupMemberCreate {
	gmc.mutation.SetGroupID(i)
	return gmc
}

// SetMemberID sets the "member_id" field.
func (gmc *GroupMemberCreate) SetMemberID(i int) *GroupMemberCreate {
	gmc.mutation.SetMemberID(i)
	return gmc
}

// Mutation returns the GroupMemberMutation object of the builder.
func (gmc *GroupMemberCreate) Mutation() *GroupMemberMutation {
	return gmc.mutation
}

// Save creates the GroupMember in the database.
func (gmc *GroupMemberCreate) Save(ctx context.Context) (*GroupMember, error) {
	var (
		err  error
		node *GroupMember
	)
	if len(gmc.hooks) == 0 {
		if err = gmc.check(); err != nil {
			return nil, err
		}
		node, err = gmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gmc.check(); err != nil {
				return nil, err
			}
			gmc.mutation = mutation
			if node, err = gmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gmc.hooks) - 1; i >= 0; i-- {
			if gmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gmc *GroupMemberCreate) SaveX(ctx context.Context) *GroupMember {
	v, err := gmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmc *GroupMemberCreate) Exec(ctx context.Context) error {
	_, err := gmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmc *GroupMemberCreate) ExecX(ctx context.Context) {
	if err := gmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmc *GroupMemberCreate) check() error {
	if _, ok := gmc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "GroupMember.group_id"`)}
	}
	if _, ok := gmc.mutation.MemberID(); !ok {
		return &ValidationError{Name: "member_id", err: errors.New(`ent: missing required field "GroupMember.member_id"`)}
	}
	return nil
}

func (gmc *GroupMemberCreate) sqlSave(ctx context.Context) (*GroupMember, error) {
	_node, _spec := gmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gmc *GroupMemberCreate) createSpec() (*GroupMember, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupMember{config: gmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: groupmember.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupmember.FieldID,
			},
		}
	)
	if value, ok := gmc.mutation.GroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupmember.FieldGroupID,
		})
		_node.GroupID = value
	}
	if value, ok := gmc.mutation.MemberID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupmember.FieldMemberID,
		})
		_node.MemberID = value
	}
	return _node, _spec
}

// GroupMemberCreateBulk is the builder for creating many GroupMember entities in bulk.
type GroupMemberCreateBulk struct {
	config
	builders []*GroupMemberCreate
}

// Save creates the GroupMember entities in the database.
func (gmcb *GroupMemberCreateBulk) Save(ctx context.Context) ([]*GroupMember, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gmcb.builders))
	nodes := make([]*GroupMember, len(gmcb.builders))
	mutators := make([]Mutator, len(gmcb.builders))
	for i := range gmcb.builders {
		func(i int, root context.Context) {
			builder := gmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMemberMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gmcb *GroupMemberCreateBulk) SaveX(ctx context.Context) []*GroupMember {
	v, err := gmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmcb *GroupMemberCreateBulk) Exec(ctx context.Context) error {
	_, err := gmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmcb *GroupMemberCreateBulk) ExecX(ctx context.Context) {
	if err := gmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
