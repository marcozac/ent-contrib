// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/marcozac/ent-contrib/entgql/internal/todouuid/ent/friendship"
	"github.com/marcozac/ent-contrib/entgql/internal/todouuid/ent/user"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FriendshipCreate is the builder for creating a Friendship entity.
type FriendshipCreate struct {
	config
	mutation *FriendshipMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (fc *FriendshipCreate) SetCreatedAt(t time.Time) *FriendshipCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FriendshipCreate) SetNillableCreatedAt(t *time.Time) *FriendshipCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUserID sets the "user_id" field.
func (fc *FriendshipCreate) SetUserID(u uuid.UUID) *FriendshipCreate {
	fc.mutation.SetUserID(u)
	return fc
}

// SetFriendID sets the "friend_id" field.
func (fc *FriendshipCreate) SetFriendID(u uuid.UUID) *FriendshipCreate {
	fc.mutation.SetFriendID(u)
	return fc
}

// SetID sets the "id" field.
func (fc *FriendshipCreate) SetID(u uuid.UUID) *FriendshipCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FriendshipCreate) SetNillableID(u *uuid.UUID) *FriendshipCreate {
	if u != nil {
		fc.SetID(*u)
	}
	return fc
}

// SetUser sets the "user" edge to the User entity.
func (fc *FriendshipCreate) SetUser(u *User) *FriendshipCreate {
	return fc.SetUserID(u.ID)
}

// SetFriend sets the "friend" edge to the User entity.
func (fc *FriendshipCreate) SetFriend(u *User) *FriendshipCreate {
	return fc.SetFriendID(u.ID)
}

// Mutation returns the FriendshipMutation object of the builder.
func (fc *FriendshipCreate) Mutation() *FriendshipMutation {
	return fc.mutation
}

// Save creates the Friendship in the database.
func (fc *FriendshipCreate) Save(ctx context.Context) (*Friendship, error) {
	fc.defaults()
	return withHooks[*Friendship, FriendshipMutation](ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FriendshipCreate) SaveX(ctx context.Context) *Friendship {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FriendshipCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FriendshipCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FriendshipCreate) defaults() {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := friendship.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := friendship.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FriendshipCreate) check() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Friendship.created_at"`)}
	}
	if _, ok := fc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Friendship.user_id"`)}
	}
	if _, ok := fc.mutation.FriendID(); !ok {
		return &ValidationError{Name: "friend_id", err: errors.New(`ent: missing required field "Friendship.friend_id"`)}
	}
	if _, ok := fc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Friendship.user"`)}
	}
	if _, ok := fc.mutation.FriendID(); !ok {
		return &ValidationError{Name: "friend", err: errors.New(`ent: missing required edge "Friendship.friend"`)}
	}
	return nil
}

func (fc *FriendshipCreate) sqlSave(ctx context.Context) (*Friendship, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FriendshipCreate) createSpec() (*Friendship, *sqlgraph.CreateSpec) {
	var (
		_node = &Friendship{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(friendship.Table, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeUUID))
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(friendship.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := fc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.UserTable,
			Columns: []string{friendship.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FriendID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FriendshipCreateBulk is the builder for creating many Friendship entities in bulk.
type FriendshipCreateBulk struct {
	config
	builders []*FriendshipCreate
}

// Save creates the Friendship entities in the database.
func (fcb *FriendshipCreateBulk) Save(ctx context.Context) ([]*Friendship, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Friendship, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FriendshipMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FriendshipCreateBulk) SaveX(ctx context.Context) []*Friendship {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FriendshipCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FriendshipCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
