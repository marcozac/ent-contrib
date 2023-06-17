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

	"github.com/marcozac/ent-contrib/entgql/internal/todo/ent/onetomany"
	"github.com/marcozac/ent-contrib/entgql/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OneToManyUpdate is the builder for updating OneToMany entities.
type OneToManyUpdate struct {
	config
	hooks    []Hook
	mutation *OneToManyMutation
}

// Where appends a list predicates to the OneToManyUpdate builder.
func (otmu *OneToManyUpdate) Where(ps ...predicate.OneToMany) *OneToManyUpdate {
	otmu.mutation.Where(ps...)
	return otmu
}

// SetName sets the "name" field.
func (otmu *OneToManyUpdate) SetName(s string) *OneToManyUpdate {
	otmu.mutation.SetName(s)
	return otmu
}

// SetField2 sets the "field2" field.
func (otmu *OneToManyUpdate) SetField2(s string) *OneToManyUpdate {
	otmu.mutation.SetField2(s)
	return otmu
}

// SetNillableField2 sets the "field2" field if the given value is not nil.
func (otmu *OneToManyUpdate) SetNillableField2(s *string) *OneToManyUpdate {
	if s != nil {
		otmu.SetField2(*s)
	}
	return otmu
}

// ClearField2 clears the value of the "field2" field.
func (otmu *OneToManyUpdate) ClearField2() *OneToManyUpdate {
	otmu.mutation.ClearField2()
	return otmu
}

// SetParentID sets the "parent_id" field.
func (otmu *OneToManyUpdate) SetParentID(i int) *OneToManyUpdate {
	otmu.mutation.SetParentID(i)
	return otmu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (otmu *OneToManyUpdate) SetNillableParentID(i *int) *OneToManyUpdate {
	if i != nil {
		otmu.SetParentID(*i)
	}
	return otmu
}

// ClearParentID clears the value of the "parent_id" field.
func (otmu *OneToManyUpdate) ClearParentID() *OneToManyUpdate {
	otmu.mutation.ClearParentID()
	return otmu
}

// SetParent sets the "parent" edge to the OneToMany entity.
func (otmu *OneToManyUpdate) SetParent(o *OneToMany) *OneToManyUpdate {
	return otmu.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the OneToMany entity by IDs.
func (otmu *OneToManyUpdate) AddChildIDs(ids ...int) *OneToManyUpdate {
	otmu.mutation.AddChildIDs(ids...)
	return otmu
}

// AddChildren adds the "children" edges to the OneToMany entity.
func (otmu *OneToManyUpdate) AddChildren(o ...*OneToMany) *OneToManyUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return otmu.AddChildIDs(ids...)
}

// Mutation returns the OneToManyMutation object of the builder.
func (otmu *OneToManyUpdate) Mutation() *OneToManyMutation {
	return otmu.mutation
}

// ClearParent clears the "parent" edge to the OneToMany entity.
func (otmu *OneToManyUpdate) ClearParent() *OneToManyUpdate {
	otmu.mutation.ClearParent()
	return otmu
}

// ClearChildren clears all "children" edges to the OneToMany entity.
func (otmu *OneToManyUpdate) ClearChildren() *OneToManyUpdate {
	otmu.mutation.ClearChildren()
	return otmu
}

// RemoveChildIDs removes the "children" edge to OneToMany entities by IDs.
func (otmu *OneToManyUpdate) RemoveChildIDs(ids ...int) *OneToManyUpdate {
	otmu.mutation.RemoveChildIDs(ids...)
	return otmu
}

// RemoveChildren removes "children" edges to OneToMany entities.
func (otmu *OneToManyUpdate) RemoveChildren(o ...*OneToMany) *OneToManyUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return otmu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (otmu *OneToManyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OneToManyMutation](ctx, otmu.sqlSave, otmu.mutation, otmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (otmu *OneToManyUpdate) SaveX(ctx context.Context) int {
	affected, err := otmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (otmu *OneToManyUpdate) Exec(ctx context.Context) error {
	_, err := otmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otmu *OneToManyUpdate) ExecX(ctx context.Context) {
	if err := otmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otmu *OneToManyUpdate) check() error {
	if v, ok := otmu.mutation.Name(); ok {
		if err := onetomany.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OneToMany.name": %w`, err)}
		}
	}
	return nil
}

func (otmu *OneToManyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := otmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(onetomany.Table, onetomany.Columns, sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt))
	if ps := otmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := otmu.mutation.Name(); ok {
		_spec.SetField(onetomany.FieldName, field.TypeString, value)
	}
	if value, ok := otmu.mutation.Field2(); ok {
		_spec.SetField(onetomany.FieldField2, field.TypeString, value)
	}
	if otmu.mutation.Field2Cleared() {
		_spec.ClearField(onetomany.FieldField2, field.TypeString)
	}
	if otmu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   onetomany.ParentTable,
			Columns: []string{onetomany.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otmu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   onetomany.ParentTable,
			Columns: []string{onetomany.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if otmu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   onetomany.ChildrenTable,
			Columns: []string{onetomany.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otmu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !otmu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   onetomany.ChildrenTable,
			Columns: []string{onetomany.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otmu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   onetomany.ChildrenTable,
			Columns: []string{onetomany.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, otmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{onetomany.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	otmu.mutation.done = true
	return n, nil
}

// OneToManyUpdateOne is the builder for updating a single OneToMany entity.
type OneToManyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OneToManyMutation
}

// SetName sets the "name" field.
func (otmuo *OneToManyUpdateOne) SetName(s string) *OneToManyUpdateOne {
	otmuo.mutation.SetName(s)
	return otmuo
}

// SetField2 sets the "field2" field.
func (otmuo *OneToManyUpdateOne) SetField2(s string) *OneToManyUpdateOne {
	otmuo.mutation.SetField2(s)
	return otmuo
}

// SetNillableField2 sets the "field2" field if the given value is not nil.
func (otmuo *OneToManyUpdateOne) SetNillableField2(s *string) *OneToManyUpdateOne {
	if s != nil {
		otmuo.SetField2(*s)
	}
	return otmuo
}

// ClearField2 clears the value of the "field2" field.
func (otmuo *OneToManyUpdateOne) ClearField2() *OneToManyUpdateOne {
	otmuo.mutation.ClearField2()
	return otmuo
}

// SetParentID sets the "parent_id" field.
func (otmuo *OneToManyUpdateOne) SetParentID(i int) *OneToManyUpdateOne {
	otmuo.mutation.SetParentID(i)
	return otmuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (otmuo *OneToManyUpdateOne) SetNillableParentID(i *int) *OneToManyUpdateOne {
	if i != nil {
		otmuo.SetParentID(*i)
	}
	return otmuo
}

// ClearParentID clears the value of the "parent_id" field.
func (otmuo *OneToManyUpdateOne) ClearParentID() *OneToManyUpdateOne {
	otmuo.mutation.ClearParentID()
	return otmuo
}

// SetParent sets the "parent" edge to the OneToMany entity.
func (otmuo *OneToManyUpdateOne) SetParent(o *OneToMany) *OneToManyUpdateOne {
	return otmuo.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the OneToMany entity by IDs.
func (otmuo *OneToManyUpdateOne) AddChildIDs(ids ...int) *OneToManyUpdateOne {
	otmuo.mutation.AddChildIDs(ids...)
	return otmuo
}

// AddChildren adds the "children" edges to the OneToMany entity.
func (otmuo *OneToManyUpdateOne) AddChildren(o ...*OneToMany) *OneToManyUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return otmuo.AddChildIDs(ids...)
}

// Mutation returns the OneToManyMutation object of the builder.
func (otmuo *OneToManyUpdateOne) Mutation() *OneToManyMutation {
	return otmuo.mutation
}

// ClearParent clears the "parent" edge to the OneToMany entity.
func (otmuo *OneToManyUpdateOne) ClearParent() *OneToManyUpdateOne {
	otmuo.mutation.ClearParent()
	return otmuo
}

// ClearChildren clears all "children" edges to the OneToMany entity.
func (otmuo *OneToManyUpdateOne) ClearChildren() *OneToManyUpdateOne {
	otmuo.mutation.ClearChildren()
	return otmuo
}

// RemoveChildIDs removes the "children" edge to OneToMany entities by IDs.
func (otmuo *OneToManyUpdateOne) RemoveChildIDs(ids ...int) *OneToManyUpdateOne {
	otmuo.mutation.RemoveChildIDs(ids...)
	return otmuo
}

// RemoveChildren removes "children" edges to OneToMany entities.
func (otmuo *OneToManyUpdateOne) RemoveChildren(o ...*OneToMany) *OneToManyUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return otmuo.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the OneToManyUpdate builder.
func (otmuo *OneToManyUpdateOne) Where(ps ...predicate.OneToMany) *OneToManyUpdateOne {
	otmuo.mutation.Where(ps...)
	return otmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (otmuo *OneToManyUpdateOne) Select(field string, fields ...string) *OneToManyUpdateOne {
	otmuo.fields = append([]string{field}, fields...)
	return otmuo
}

// Save executes the query and returns the updated OneToMany entity.
func (otmuo *OneToManyUpdateOne) Save(ctx context.Context) (*OneToMany, error) {
	return withHooks[*OneToMany, OneToManyMutation](ctx, otmuo.sqlSave, otmuo.mutation, otmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (otmuo *OneToManyUpdateOne) SaveX(ctx context.Context) *OneToMany {
	node, err := otmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (otmuo *OneToManyUpdateOne) Exec(ctx context.Context) error {
	_, err := otmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otmuo *OneToManyUpdateOne) ExecX(ctx context.Context) {
	if err := otmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otmuo *OneToManyUpdateOne) check() error {
	if v, ok := otmuo.mutation.Name(); ok {
		if err := onetomany.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OneToMany.name": %w`, err)}
		}
	}
	return nil
}

func (otmuo *OneToManyUpdateOne) sqlSave(ctx context.Context) (_node *OneToMany, err error) {
	if err := otmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(onetomany.Table, onetomany.Columns, sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt))
	id, ok := otmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OneToMany.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := otmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, onetomany.FieldID)
		for _, f := range fields {
			if !onetomany.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != onetomany.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := otmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := otmuo.mutation.Name(); ok {
		_spec.SetField(onetomany.FieldName, field.TypeString, value)
	}
	if value, ok := otmuo.mutation.Field2(); ok {
		_spec.SetField(onetomany.FieldField2, field.TypeString, value)
	}
	if otmuo.mutation.Field2Cleared() {
		_spec.ClearField(onetomany.FieldField2, field.TypeString)
	}
	if otmuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   onetomany.ParentTable,
			Columns: []string{onetomany.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otmuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   onetomany.ParentTable,
			Columns: []string{onetomany.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if otmuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   onetomany.ChildrenTable,
			Columns: []string{onetomany.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otmuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !otmuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   onetomany.ChildrenTable,
			Columns: []string{onetomany.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := otmuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   onetomany.ChildrenTable,
			Columns: []string{onetomany.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OneToMany{config: otmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, otmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{onetomany.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	otmuo.mutation.done = true
	return _node, nil
}
