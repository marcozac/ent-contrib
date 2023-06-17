// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/marcozac/ent-contrib/entproto/internal/todo/ent/multiwordschema"
	"github.com/marcozac/ent-contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MultiWordSchemaUpdate is the builder for updating MultiWordSchema entities.
type MultiWordSchemaUpdate struct {
	config
	hooks    []Hook
	mutation *MultiWordSchemaMutation
}

// Where appends a list predicates to the MultiWordSchemaUpdate builder.
func (mwsu *MultiWordSchemaUpdate) Where(ps ...predicate.MultiWordSchema) *MultiWordSchemaUpdate {
	mwsu.mutation.Where(ps...)
	return mwsu
}

// SetUnit sets the "unit" field.
func (mwsu *MultiWordSchemaUpdate) SetUnit(m multiwordschema.Unit) *MultiWordSchemaUpdate {
	mwsu.mutation.SetUnit(m)
	return mwsu
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (mwsu *MultiWordSchemaUpdate) SetNillableUnit(m *multiwordschema.Unit) *MultiWordSchemaUpdate {
	if m != nil {
		mwsu.SetUnit(*m)
	}
	return mwsu
}

// Mutation returns the MultiWordSchemaMutation object of the builder.
func (mwsu *MultiWordSchemaUpdate) Mutation() *MultiWordSchemaMutation {
	return mwsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mwsu *MultiWordSchemaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, MultiWordSchemaMutation](ctx, mwsu.sqlSave, mwsu.mutation, mwsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mwsu *MultiWordSchemaUpdate) SaveX(ctx context.Context) int {
	affected, err := mwsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mwsu *MultiWordSchemaUpdate) Exec(ctx context.Context) error {
	_, err := mwsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwsu *MultiWordSchemaUpdate) ExecX(ctx context.Context) {
	if err := mwsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mwsu *MultiWordSchemaUpdate) check() error {
	if v, ok := mwsu.mutation.Unit(); ok {
		if err := multiwordschema.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf(`ent: validator failed for field "MultiWordSchema.unit": %w`, err)}
		}
	}
	return nil
}

func (mwsu *MultiWordSchemaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mwsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(multiwordschema.Table, multiwordschema.Columns, sqlgraph.NewFieldSpec(multiwordschema.FieldID, field.TypeInt))
	if ps := mwsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mwsu.mutation.Unit(); ok {
		_spec.SetField(multiwordschema.FieldUnit, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mwsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{multiwordschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mwsu.mutation.done = true
	return n, nil
}

// MultiWordSchemaUpdateOne is the builder for updating a single MultiWordSchema entity.
type MultiWordSchemaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MultiWordSchemaMutation
}

// SetUnit sets the "unit" field.
func (mwsuo *MultiWordSchemaUpdateOne) SetUnit(m multiwordschema.Unit) *MultiWordSchemaUpdateOne {
	mwsuo.mutation.SetUnit(m)
	return mwsuo
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (mwsuo *MultiWordSchemaUpdateOne) SetNillableUnit(m *multiwordschema.Unit) *MultiWordSchemaUpdateOne {
	if m != nil {
		mwsuo.SetUnit(*m)
	}
	return mwsuo
}

// Mutation returns the MultiWordSchemaMutation object of the builder.
func (mwsuo *MultiWordSchemaUpdateOne) Mutation() *MultiWordSchemaMutation {
	return mwsuo.mutation
}

// Where appends a list predicates to the MultiWordSchemaUpdate builder.
func (mwsuo *MultiWordSchemaUpdateOne) Where(ps ...predicate.MultiWordSchema) *MultiWordSchemaUpdateOne {
	mwsuo.mutation.Where(ps...)
	return mwsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mwsuo *MultiWordSchemaUpdateOne) Select(field string, fields ...string) *MultiWordSchemaUpdateOne {
	mwsuo.fields = append([]string{field}, fields...)
	return mwsuo
}

// Save executes the query and returns the updated MultiWordSchema entity.
func (mwsuo *MultiWordSchemaUpdateOne) Save(ctx context.Context) (*MultiWordSchema, error) {
	return withHooks[*MultiWordSchema, MultiWordSchemaMutation](ctx, mwsuo.sqlSave, mwsuo.mutation, mwsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mwsuo *MultiWordSchemaUpdateOne) SaveX(ctx context.Context) *MultiWordSchema {
	node, err := mwsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mwsuo *MultiWordSchemaUpdateOne) Exec(ctx context.Context) error {
	_, err := mwsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwsuo *MultiWordSchemaUpdateOne) ExecX(ctx context.Context) {
	if err := mwsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mwsuo *MultiWordSchemaUpdateOne) check() error {
	if v, ok := mwsuo.mutation.Unit(); ok {
		if err := multiwordschema.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf(`ent: validator failed for field "MultiWordSchema.unit": %w`, err)}
		}
	}
	return nil
}

func (mwsuo *MultiWordSchemaUpdateOne) sqlSave(ctx context.Context) (_node *MultiWordSchema, err error) {
	if err := mwsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(multiwordschema.Table, multiwordschema.Columns, sqlgraph.NewFieldSpec(multiwordschema.FieldID, field.TypeInt))
	id, ok := mwsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MultiWordSchema.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mwsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, multiwordschema.FieldID)
		for _, f := range fields {
			if !multiwordschema.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != multiwordschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mwsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mwsuo.mutation.Unit(); ok {
		_spec.SetField(multiwordschema.FieldUnit, field.TypeEnum, value)
	}
	_node = &MultiWordSchema{config: mwsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mwsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{multiwordschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mwsuo.mutation.done = true
	return _node, nil
}
