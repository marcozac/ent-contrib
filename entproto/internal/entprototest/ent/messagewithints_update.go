// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/messagewithints"
	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// MessageWithIntsUpdate is the builder for updating MessageWithInts entities.
type MessageWithIntsUpdate struct {
	config
	hooks    []Hook
	mutation *MessageWithIntsMutation
}

// Where appends a list predicates to the MessageWithIntsUpdate builder.
func (mwiu *MessageWithIntsUpdate) Where(ps ...predicate.MessageWithInts) *MessageWithIntsUpdate {
	mwiu.mutation.Where(ps...)
	return mwiu
}

// SetInt32s sets the "int32s" field.
func (mwiu *MessageWithIntsUpdate) SetInt32s(i []int32) *MessageWithIntsUpdate {
	mwiu.mutation.SetInt32s(i)
	return mwiu
}

// AppendInt32s appends i to the "int32s" field.
func (mwiu *MessageWithIntsUpdate) AppendInt32s(i []int32) *MessageWithIntsUpdate {
	mwiu.mutation.AppendInt32s(i)
	return mwiu
}

// SetInt64s sets the "int64s" field.
func (mwiu *MessageWithIntsUpdate) SetInt64s(i []int64) *MessageWithIntsUpdate {
	mwiu.mutation.SetInt64s(i)
	return mwiu
}

// AppendInt64s appends i to the "int64s" field.
func (mwiu *MessageWithIntsUpdate) AppendInt64s(i []int64) *MessageWithIntsUpdate {
	mwiu.mutation.AppendInt64s(i)
	return mwiu
}

// SetUint32s sets the "uint32s" field.
func (mwiu *MessageWithIntsUpdate) SetUint32s(u []uint32) *MessageWithIntsUpdate {
	mwiu.mutation.SetUint32s(u)
	return mwiu
}

// AppendUint32s appends u to the "uint32s" field.
func (mwiu *MessageWithIntsUpdate) AppendUint32s(u []uint32) *MessageWithIntsUpdate {
	mwiu.mutation.AppendUint32s(u)
	return mwiu
}

// SetUint64s sets the "uint64s" field.
func (mwiu *MessageWithIntsUpdate) SetUint64s(u []uint64) *MessageWithIntsUpdate {
	mwiu.mutation.SetUint64s(u)
	return mwiu
}

// AppendUint64s appends u to the "uint64s" field.
func (mwiu *MessageWithIntsUpdate) AppendUint64s(u []uint64) *MessageWithIntsUpdate {
	mwiu.mutation.AppendUint64s(u)
	return mwiu
}

// Mutation returns the MessageWithIntsMutation object of the builder.
func (mwiu *MessageWithIntsUpdate) Mutation() *MessageWithIntsMutation {
	return mwiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mwiu *MessageWithIntsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, MessageWithIntsMutation](ctx, mwiu.sqlSave, mwiu.mutation, mwiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mwiu *MessageWithIntsUpdate) SaveX(ctx context.Context) int {
	affected, err := mwiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mwiu *MessageWithIntsUpdate) Exec(ctx context.Context) error {
	_, err := mwiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwiu *MessageWithIntsUpdate) ExecX(ctx context.Context) {
	if err := mwiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwiu *MessageWithIntsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(messagewithints.Table, messagewithints.Columns, sqlgraph.NewFieldSpec(messagewithints.FieldID, field.TypeInt))
	if ps := mwiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mwiu.mutation.Int32s(); ok {
		_spec.SetField(messagewithints.FieldInt32s, field.TypeJSON, value)
	}
	if value, ok := mwiu.mutation.AppendedInt32s(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, messagewithints.FieldInt32s, value)
		})
	}
	if value, ok := mwiu.mutation.Int64s(); ok {
		_spec.SetField(messagewithints.FieldInt64s, field.TypeJSON, value)
	}
	if value, ok := mwiu.mutation.AppendedInt64s(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, messagewithints.FieldInt64s, value)
		})
	}
	if value, ok := mwiu.mutation.Uint32s(); ok {
		_spec.SetField(messagewithints.FieldUint32s, field.TypeJSON, value)
	}
	if value, ok := mwiu.mutation.AppendedUint32s(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, messagewithints.FieldUint32s, value)
		})
	}
	if value, ok := mwiu.mutation.Uint64s(); ok {
		_spec.SetField(messagewithints.FieldUint64s, field.TypeJSON, value)
	}
	if value, ok := mwiu.mutation.AppendedUint64s(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, messagewithints.FieldUint64s, value)
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mwiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithints.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mwiu.mutation.done = true
	return n, nil
}

// MessageWithIntsUpdateOne is the builder for updating a single MessageWithInts entity.
type MessageWithIntsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageWithIntsMutation
}

// SetInt32s sets the "int32s" field.
func (mwiuo *MessageWithIntsUpdateOne) SetInt32s(i []int32) *MessageWithIntsUpdateOne {
	mwiuo.mutation.SetInt32s(i)
	return mwiuo
}

// AppendInt32s appends i to the "int32s" field.
func (mwiuo *MessageWithIntsUpdateOne) AppendInt32s(i []int32) *MessageWithIntsUpdateOne {
	mwiuo.mutation.AppendInt32s(i)
	return mwiuo
}

// SetInt64s sets the "int64s" field.
func (mwiuo *MessageWithIntsUpdateOne) SetInt64s(i []int64) *MessageWithIntsUpdateOne {
	mwiuo.mutation.SetInt64s(i)
	return mwiuo
}

// AppendInt64s appends i to the "int64s" field.
func (mwiuo *MessageWithIntsUpdateOne) AppendInt64s(i []int64) *MessageWithIntsUpdateOne {
	mwiuo.mutation.AppendInt64s(i)
	return mwiuo
}

// SetUint32s sets the "uint32s" field.
func (mwiuo *MessageWithIntsUpdateOne) SetUint32s(u []uint32) *MessageWithIntsUpdateOne {
	mwiuo.mutation.SetUint32s(u)
	return mwiuo
}

// AppendUint32s appends u to the "uint32s" field.
func (mwiuo *MessageWithIntsUpdateOne) AppendUint32s(u []uint32) *MessageWithIntsUpdateOne {
	mwiuo.mutation.AppendUint32s(u)
	return mwiuo
}

// SetUint64s sets the "uint64s" field.
func (mwiuo *MessageWithIntsUpdateOne) SetUint64s(u []uint64) *MessageWithIntsUpdateOne {
	mwiuo.mutation.SetUint64s(u)
	return mwiuo
}

// AppendUint64s appends u to the "uint64s" field.
func (mwiuo *MessageWithIntsUpdateOne) AppendUint64s(u []uint64) *MessageWithIntsUpdateOne {
	mwiuo.mutation.AppendUint64s(u)
	return mwiuo
}

// Mutation returns the MessageWithIntsMutation object of the builder.
func (mwiuo *MessageWithIntsUpdateOne) Mutation() *MessageWithIntsMutation {
	return mwiuo.mutation
}

// Where appends a list predicates to the MessageWithIntsUpdate builder.
func (mwiuo *MessageWithIntsUpdateOne) Where(ps ...predicate.MessageWithInts) *MessageWithIntsUpdateOne {
	mwiuo.mutation.Where(ps...)
	return mwiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mwiuo *MessageWithIntsUpdateOne) Select(field string, fields ...string) *MessageWithIntsUpdateOne {
	mwiuo.fields = append([]string{field}, fields...)
	return mwiuo
}

// Save executes the query and returns the updated MessageWithInts entity.
func (mwiuo *MessageWithIntsUpdateOne) Save(ctx context.Context) (*MessageWithInts, error) {
	return withHooks[*MessageWithInts, MessageWithIntsMutation](ctx, mwiuo.sqlSave, mwiuo.mutation, mwiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mwiuo *MessageWithIntsUpdateOne) SaveX(ctx context.Context) *MessageWithInts {
	node, err := mwiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mwiuo *MessageWithIntsUpdateOne) Exec(ctx context.Context) error {
	_, err := mwiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwiuo *MessageWithIntsUpdateOne) ExecX(ctx context.Context) {
	if err := mwiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwiuo *MessageWithIntsUpdateOne) sqlSave(ctx context.Context) (_node *MessageWithInts, err error) {
	_spec := sqlgraph.NewUpdateSpec(messagewithints.Table, messagewithints.Columns, sqlgraph.NewFieldSpec(messagewithints.FieldID, field.TypeInt))
	id, ok := mwiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MessageWithInts.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mwiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithints.FieldID)
		for _, f := range fields {
			if !messagewithints.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != messagewithints.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mwiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mwiuo.mutation.Int32s(); ok {
		_spec.SetField(messagewithints.FieldInt32s, field.TypeJSON, value)
	}
	if value, ok := mwiuo.mutation.AppendedInt32s(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, messagewithints.FieldInt32s, value)
		})
	}
	if value, ok := mwiuo.mutation.Int64s(); ok {
		_spec.SetField(messagewithints.FieldInt64s, field.TypeJSON, value)
	}
	if value, ok := mwiuo.mutation.AppendedInt64s(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, messagewithints.FieldInt64s, value)
		})
	}
	if value, ok := mwiuo.mutation.Uint32s(); ok {
		_spec.SetField(messagewithints.FieldUint32s, field.TypeJSON, value)
	}
	if value, ok := mwiuo.mutation.AppendedUint32s(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, messagewithints.FieldUint32s, value)
		})
	}
	if value, ok := mwiuo.mutation.Uint64s(); ok {
		_spec.SetField(messagewithints.FieldUint64s, field.TypeJSON, value)
	}
	if value, ok := mwiuo.mutation.AppendedUint64s(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, messagewithints.FieldUint64s, value)
		})
	}
	_node = &MessageWithInts{config: mwiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mwiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithints.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mwiuo.mutation.done = true
	return _node, nil
}
