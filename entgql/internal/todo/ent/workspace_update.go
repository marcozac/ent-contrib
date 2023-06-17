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

	"github.com/marcozac/ent-contrib/entgql/internal/todo/ent/predicate"
	"github.com/marcozac/ent-contrib/entgql/internal/todo/ent/workspace"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceUpdate is the builder for updating Workspace entities.
type WorkspaceUpdate struct {
	config
	hooks    []Hook
	mutation *WorkspaceMutation
}

// Where appends a list predicates to the WorkspaceUpdate builder.
func (wu *WorkspaceUpdate) Where(ps ...predicate.Workspace) *WorkspaceUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetName sets the "name" field.
func (wu *WorkspaceUpdate) SetName(s string) *WorkspaceUpdate {
	wu.mutation.SetName(s)
	return wu
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wu *WorkspaceUpdate) Mutation() *WorkspaceMutation {
	return wu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkspaceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, WorkspaceMutation](ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkspaceUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkspaceUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkspaceUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wu *WorkspaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(workspace.Table, workspace.Columns, sqlgraph.NewFieldSpec(workspace.FieldID, field.TypeInt))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Name(); ok {
		_spec.SetField(workspace.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WorkspaceUpdateOne is the builder for updating a single Workspace entity.
type WorkspaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkspaceMutation
}

// SetName sets the "name" field.
func (wuo *WorkspaceUpdateOne) SetName(s string) *WorkspaceUpdateOne {
	wuo.mutation.SetName(s)
	return wuo
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wuo *WorkspaceUpdateOne) Mutation() *WorkspaceMutation {
	return wuo.mutation
}

// Where appends a list predicates to the WorkspaceUpdate builder.
func (wuo *WorkspaceUpdateOne) Where(ps ...predicate.Workspace) *WorkspaceUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkspaceUpdateOne) Select(field string, fields ...string) *WorkspaceUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Workspace entity.
func (wuo *WorkspaceUpdateOne) Save(ctx context.Context) (*Workspace, error) {
	return withHooks[*Workspace, WorkspaceMutation](ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkspaceUpdateOne) SaveX(ctx context.Context) *Workspace {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkspaceUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkspaceUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wuo *WorkspaceUpdateOne) sqlSave(ctx context.Context) (_node *Workspace, err error) {
	_spec := sqlgraph.NewUpdateSpec(workspace.Table, workspace.Columns, sqlgraph.NewFieldSpec(workspace.FieldID, field.TypeInt))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Workspace.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workspace.FieldID)
		for _, f := range fields {
			if !workspace.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workspace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Name(); ok {
		_spec.SetField(workspace.FieldName, field.TypeString, value)
	}
	_node = &Workspace{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
