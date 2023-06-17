// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/messagewithfieldone"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithFieldOneCreate is the builder for creating a MessageWithFieldOne entity.
type MessageWithFieldOneCreate struct {
	config
	mutation *MessageWithFieldOneMutation
	hooks    []Hook
}

// SetFieldOne sets the "field_one" field.
func (mwfoc *MessageWithFieldOneCreate) SetFieldOne(i int32) *MessageWithFieldOneCreate {
	mwfoc.mutation.SetFieldOne(i)
	return mwfoc
}

// Mutation returns the MessageWithFieldOneMutation object of the builder.
func (mwfoc *MessageWithFieldOneCreate) Mutation() *MessageWithFieldOneMutation {
	return mwfoc.mutation
}

// Save creates the MessageWithFieldOne in the database.
func (mwfoc *MessageWithFieldOneCreate) Save(ctx context.Context) (*MessageWithFieldOne, error) {
	return withHooks[*MessageWithFieldOne, MessageWithFieldOneMutation](ctx, mwfoc.sqlSave, mwfoc.mutation, mwfoc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mwfoc *MessageWithFieldOneCreate) SaveX(ctx context.Context) *MessageWithFieldOne {
	v, err := mwfoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mwfoc *MessageWithFieldOneCreate) Exec(ctx context.Context) error {
	_, err := mwfoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwfoc *MessageWithFieldOneCreate) ExecX(ctx context.Context) {
	if err := mwfoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mwfoc *MessageWithFieldOneCreate) check() error {
	if _, ok := mwfoc.mutation.FieldOne(); !ok {
		return &ValidationError{Name: "field_one", err: errors.New(`ent: missing required field "MessageWithFieldOne.field_one"`)}
	}
	return nil
}

func (mwfoc *MessageWithFieldOneCreate) sqlSave(ctx context.Context) (*MessageWithFieldOne, error) {
	if err := mwfoc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mwfoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mwfoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mwfoc.mutation.id = &_node.ID
	mwfoc.mutation.done = true
	return _node, nil
}

func (mwfoc *MessageWithFieldOneCreate) createSpec() (*MessageWithFieldOne, *sqlgraph.CreateSpec) {
	var (
		_node = &MessageWithFieldOne{config: mwfoc.config}
		_spec = sqlgraph.NewCreateSpec(messagewithfieldone.Table, sqlgraph.NewFieldSpec(messagewithfieldone.FieldID, field.TypeInt))
	)
	if value, ok := mwfoc.mutation.FieldOne(); ok {
		_spec.SetField(messagewithfieldone.FieldFieldOne, field.TypeInt32, value)
		_node.FieldOne = value
	}
	return _node, _spec
}

// MessageWithFieldOneCreateBulk is the builder for creating many MessageWithFieldOne entities in bulk.
type MessageWithFieldOneCreateBulk struct {
	config
	builders []*MessageWithFieldOneCreate
}

// Save creates the MessageWithFieldOne entities in the database.
func (mwfocb *MessageWithFieldOneCreateBulk) Save(ctx context.Context) ([]*MessageWithFieldOne, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mwfocb.builders))
	nodes := make([]*MessageWithFieldOne, len(mwfocb.builders))
	mutators := make([]Mutator, len(mwfocb.builders))
	for i := range mwfocb.builders {
		func(i int, root context.Context) {
			builder := mwfocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageWithFieldOneMutation)
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
					_, err = mutators[i+1].Mutate(root, mwfocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mwfocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mwfocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mwfocb *MessageWithFieldOneCreateBulk) SaveX(ctx context.Context) []*MessageWithFieldOne {
	v, err := mwfocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mwfocb *MessageWithFieldOneCreateBulk) Exec(ctx context.Context) error {
	_, err := mwfocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwfocb *MessageWithFieldOneCreateBulk) ExecX(ctx context.Context) {
	if err := mwfocb.Exec(ctx); err != nil {
		panic(err)
	}
}
