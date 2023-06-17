// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/invalidfieldmessage"
	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/schema"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InvalidFieldMessageCreate is the builder for creating a InvalidFieldMessage entity.
type InvalidFieldMessageCreate struct {
	config
	mutation *InvalidFieldMessageMutation
	hooks    []Hook
}

// SetJSON sets the "json" field.
func (ifmc *InvalidFieldMessageCreate) SetJSON(sj *schema.SomeJSON) *InvalidFieldMessageCreate {
	ifmc.mutation.SetJSON(sj)
	return ifmc
}

// Mutation returns the InvalidFieldMessageMutation object of the builder.
func (ifmc *InvalidFieldMessageCreate) Mutation() *InvalidFieldMessageMutation {
	return ifmc.mutation
}

// Save creates the InvalidFieldMessage in the database.
func (ifmc *InvalidFieldMessageCreate) Save(ctx context.Context) (*InvalidFieldMessage, error) {
	return withHooks[*InvalidFieldMessage, InvalidFieldMessageMutation](ctx, ifmc.sqlSave, ifmc.mutation, ifmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ifmc *InvalidFieldMessageCreate) SaveX(ctx context.Context) *InvalidFieldMessage {
	v, err := ifmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ifmc *InvalidFieldMessageCreate) Exec(ctx context.Context) error {
	_, err := ifmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifmc *InvalidFieldMessageCreate) ExecX(ctx context.Context) {
	if err := ifmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ifmc *InvalidFieldMessageCreate) check() error {
	if _, ok := ifmc.mutation.JSON(); !ok {
		return &ValidationError{Name: "json", err: errors.New(`ent: missing required field "InvalidFieldMessage.json"`)}
	}
	return nil
}

func (ifmc *InvalidFieldMessageCreate) sqlSave(ctx context.Context) (*InvalidFieldMessage, error) {
	if err := ifmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ifmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ifmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ifmc.mutation.id = &_node.ID
	ifmc.mutation.done = true
	return _node, nil
}

func (ifmc *InvalidFieldMessageCreate) createSpec() (*InvalidFieldMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &InvalidFieldMessage{config: ifmc.config}
		_spec = sqlgraph.NewCreateSpec(invalidfieldmessage.Table, sqlgraph.NewFieldSpec(invalidfieldmessage.FieldID, field.TypeInt))
	)
	if value, ok := ifmc.mutation.JSON(); ok {
		_spec.SetField(invalidfieldmessage.FieldJSON, field.TypeJSON, value)
		_node.JSON = value
	}
	return _node, _spec
}

// InvalidFieldMessageCreateBulk is the builder for creating many InvalidFieldMessage entities in bulk.
type InvalidFieldMessageCreateBulk struct {
	config
	builders []*InvalidFieldMessageCreate
}

// Save creates the InvalidFieldMessage entities in the database.
func (ifmcb *InvalidFieldMessageCreateBulk) Save(ctx context.Context) ([]*InvalidFieldMessage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ifmcb.builders))
	nodes := make([]*InvalidFieldMessage, len(ifmcb.builders))
	mutators := make([]Mutator, len(ifmcb.builders))
	for i := range ifmcb.builders {
		func(i int, root context.Context) {
			builder := ifmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvalidFieldMessageMutation)
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
					_, err = mutators[i+1].Mutate(root, ifmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ifmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ifmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ifmcb *InvalidFieldMessageCreateBulk) SaveX(ctx context.Context) []*InvalidFieldMessage {
	v, err := ifmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ifmcb *InvalidFieldMessageCreateBulk) Exec(ctx context.Context) error {
	_, err := ifmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ifmcb *InvalidFieldMessageCreateBulk) ExecX(ctx context.Context) {
	if err := ifmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
