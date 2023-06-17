// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/marcozac/ent-contrib/entproto/internal/todo/ent/pony"
	"github.com/marcozac/ent-contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PonyUpdate is the builder for updating Pony entities.
type PonyUpdate struct {
	config
	hooks    []Hook
	mutation *PonyMutation
}

// Where appends a list predicates to the PonyUpdate builder.
func (pu *PonyUpdate) Where(ps ...predicate.Pony) *PonyUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PonyUpdate) SetName(s string) *PonyUpdate {
	pu.mutation.SetName(s)
	return pu
}

// Mutation returns the PonyMutation object of the builder.
func (pu *PonyUpdate) Mutation() *PonyMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PonyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PonyMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PonyUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PonyUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PonyUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PonyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(pony.Table, pony.Columns, sqlgraph.NewFieldSpec(pony.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(pony.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pony.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PonyUpdateOne is the builder for updating a single Pony entity.
type PonyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PonyMutation
}

// SetName sets the "name" field.
func (puo *PonyUpdateOne) SetName(s string) *PonyUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// Mutation returns the PonyMutation object of the builder.
func (puo *PonyUpdateOne) Mutation() *PonyMutation {
	return puo.mutation
}

// Where appends a list predicates to the PonyUpdate builder.
func (puo *PonyUpdateOne) Where(ps ...predicate.Pony) *PonyUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PonyUpdateOne) Select(field string, fields ...string) *PonyUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pony entity.
func (puo *PonyUpdateOne) Save(ctx context.Context) (*Pony, error) {
	return withHooks[*Pony, PonyMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PonyUpdateOne) SaveX(ctx context.Context) *Pony {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PonyUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PonyUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PonyUpdateOne) sqlSave(ctx context.Context) (_node *Pony, err error) {
	_spec := sqlgraph.NewUpdateSpec(pony.Table, pony.Columns, sqlgraph.NewFieldSpec(pony.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pony.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pony.FieldID)
		for _, f := range fields {
			if !pony.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pony.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(pony.FieldName, field.TypeString, value)
	}
	_node = &Pony{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pony.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
