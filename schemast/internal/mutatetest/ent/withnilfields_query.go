// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"github.com/marcozac/ent-contrib/schemast/internal/mutatetest/ent/predicate"
	"github.com/marcozac/ent-contrib/schemast/internal/mutatetest/ent/withnilfields"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithNilFieldsQuery is the builder for querying WithNilFields entities.
type WithNilFieldsQuery struct {
	config
	ctx        *QueryContext
	order      []withnilfields.OrderOption
	inters     []Interceptor
	predicates []predicate.WithNilFields
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WithNilFieldsQuery builder.
func (wnfq *WithNilFieldsQuery) Where(ps ...predicate.WithNilFields) *WithNilFieldsQuery {
	wnfq.predicates = append(wnfq.predicates, ps...)
	return wnfq
}

// Limit the number of records to be returned by this query.
func (wnfq *WithNilFieldsQuery) Limit(limit int) *WithNilFieldsQuery {
	wnfq.ctx.Limit = &limit
	return wnfq
}

// Offset to start from.
func (wnfq *WithNilFieldsQuery) Offset(offset int) *WithNilFieldsQuery {
	wnfq.ctx.Offset = &offset
	return wnfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wnfq *WithNilFieldsQuery) Unique(unique bool) *WithNilFieldsQuery {
	wnfq.ctx.Unique = &unique
	return wnfq
}

// Order specifies how the records should be ordered.
func (wnfq *WithNilFieldsQuery) Order(o ...withnilfields.OrderOption) *WithNilFieldsQuery {
	wnfq.order = append(wnfq.order, o...)
	return wnfq
}

// First returns the first WithNilFields entity from the query.
// Returns a *NotFoundError when no WithNilFields was found.
func (wnfq *WithNilFieldsQuery) First(ctx context.Context) (*WithNilFields, error) {
	nodes, err := wnfq.Limit(1).All(setContextOp(ctx, wnfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{withnilfields.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) FirstX(ctx context.Context) *WithNilFields {
	node, err := wnfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WithNilFields ID from the query.
// Returns a *NotFoundError when no WithNilFields ID was found.
func (wnfq *WithNilFieldsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wnfq.Limit(1).IDs(setContextOp(ctx, wnfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{withnilfields.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) FirstIDX(ctx context.Context) int {
	id, err := wnfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WithNilFields entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WithNilFields entity is found.
// Returns a *NotFoundError when no WithNilFields entities are found.
func (wnfq *WithNilFieldsQuery) Only(ctx context.Context) (*WithNilFields, error) {
	nodes, err := wnfq.Limit(2).All(setContextOp(ctx, wnfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{withnilfields.Label}
	default:
		return nil, &NotSingularError{withnilfields.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) OnlyX(ctx context.Context) *WithNilFields {
	node, err := wnfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WithNilFields ID in the query.
// Returns a *NotSingularError when more than one WithNilFields ID is found.
// Returns a *NotFoundError when no entities are found.
func (wnfq *WithNilFieldsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wnfq.Limit(2).IDs(setContextOp(ctx, wnfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{withnilfields.Label}
	default:
		err = &NotSingularError{withnilfields.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) OnlyIDX(ctx context.Context) int {
	id, err := wnfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WithNilFieldsSlice.
func (wnfq *WithNilFieldsQuery) All(ctx context.Context) ([]*WithNilFields, error) {
	ctx = setContextOp(ctx, wnfq.ctx, "All")
	if err := wnfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WithNilFields, *WithNilFieldsQuery]()
	return withInterceptors[[]*WithNilFields](ctx, wnfq, qr, wnfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) AllX(ctx context.Context) []*WithNilFields {
	nodes, err := wnfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WithNilFields IDs.
func (wnfq *WithNilFieldsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if wnfq.ctx.Unique == nil && wnfq.path != nil {
		wnfq.Unique(true)
	}
	ctx = setContextOp(ctx, wnfq.ctx, "IDs")
	if err = wnfq.Select(withnilfields.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) IDsX(ctx context.Context) []int {
	ids, err := wnfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wnfq *WithNilFieldsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wnfq.ctx, "Count")
	if err := wnfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wnfq, querierCount[*WithNilFieldsQuery](), wnfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) CountX(ctx context.Context) int {
	count, err := wnfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wnfq *WithNilFieldsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wnfq.ctx, "Exist")
	switch _, err := wnfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wnfq *WithNilFieldsQuery) ExistX(ctx context.Context) bool {
	exist, err := wnfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WithNilFieldsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wnfq *WithNilFieldsQuery) Clone() *WithNilFieldsQuery {
	if wnfq == nil {
		return nil
	}
	return &WithNilFieldsQuery{
		config:     wnfq.config,
		ctx:        wnfq.ctx.Clone(),
		order:      append([]withnilfields.OrderOption{}, wnfq.order...),
		inters:     append([]Interceptor{}, wnfq.inters...),
		predicates: append([]predicate.WithNilFields{}, wnfq.predicates...),
		// clone intermediate query.
		sql:  wnfq.sql.Clone(),
		path: wnfq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (wnfq *WithNilFieldsQuery) GroupBy(field string, fields ...string) *WithNilFieldsGroupBy {
	wnfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WithNilFieldsGroupBy{build: wnfq}
	grbuild.flds = &wnfq.ctx.Fields
	grbuild.label = withnilfields.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (wnfq *WithNilFieldsQuery) Select(fields ...string) *WithNilFieldsSelect {
	wnfq.ctx.Fields = append(wnfq.ctx.Fields, fields...)
	sbuild := &WithNilFieldsSelect{WithNilFieldsQuery: wnfq}
	sbuild.label = withnilfields.Label
	sbuild.flds, sbuild.scan = &wnfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WithNilFieldsSelect configured with the given aggregations.
func (wnfq *WithNilFieldsQuery) Aggregate(fns ...AggregateFunc) *WithNilFieldsSelect {
	return wnfq.Select().Aggregate(fns...)
}

func (wnfq *WithNilFieldsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wnfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wnfq); err != nil {
				return err
			}
		}
	}
	for _, f := range wnfq.ctx.Fields {
		if !withnilfields.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wnfq.path != nil {
		prev, err := wnfq.path(ctx)
		if err != nil {
			return err
		}
		wnfq.sql = prev
	}
	return nil
}

func (wnfq *WithNilFieldsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WithNilFields, error) {
	var (
		nodes = []*WithNilFields{}
		_spec = wnfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WithNilFields).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WithNilFields{config: wnfq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wnfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wnfq *WithNilFieldsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wnfq.querySpec()
	_spec.Node.Columns = wnfq.ctx.Fields
	if len(wnfq.ctx.Fields) > 0 {
		_spec.Unique = wnfq.ctx.Unique != nil && *wnfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wnfq.driver, _spec)
}

func (wnfq *WithNilFieldsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(withnilfields.Table, withnilfields.Columns, sqlgraph.NewFieldSpec(withnilfields.FieldID, field.TypeInt))
	_spec.From = wnfq.sql
	if unique := wnfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wnfq.path != nil {
		_spec.Unique = true
	}
	if fields := wnfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withnilfields.FieldID)
		for i := range fields {
			if fields[i] != withnilfields.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wnfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wnfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wnfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wnfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wnfq *WithNilFieldsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wnfq.driver.Dialect())
	t1 := builder.Table(withnilfields.Table)
	columns := wnfq.ctx.Fields
	if len(columns) == 0 {
		columns = withnilfields.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wnfq.sql != nil {
		selector = wnfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wnfq.ctx.Unique != nil && *wnfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wnfq.predicates {
		p(selector)
	}
	for _, p := range wnfq.order {
		p(selector)
	}
	if offset := wnfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wnfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNilFieldsGroupBy is the group-by builder for WithNilFields entities.
type WithNilFieldsGroupBy struct {
	selector
	build *WithNilFieldsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wnfgb *WithNilFieldsGroupBy) Aggregate(fns ...AggregateFunc) *WithNilFieldsGroupBy {
	wnfgb.fns = append(wnfgb.fns, fns...)
	return wnfgb
}

// Scan applies the selector query and scans the result into the given value.
func (wnfgb *WithNilFieldsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wnfgb.build.ctx, "GroupBy")
	if err := wnfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WithNilFieldsQuery, *WithNilFieldsGroupBy](ctx, wnfgb.build, wnfgb, wnfgb.build.inters, v)
}

func (wnfgb *WithNilFieldsGroupBy) sqlScan(ctx context.Context, root *WithNilFieldsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wnfgb.fns))
	for _, fn := range wnfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wnfgb.flds)+len(wnfgb.fns))
		for _, f := range *wnfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wnfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wnfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WithNilFieldsSelect is the builder for selecting fields of WithNilFields entities.
type WithNilFieldsSelect struct {
	*WithNilFieldsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wnfs *WithNilFieldsSelect) Aggregate(fns ...AggregateFunc) *WithNilFieldsSelect {
	wnfs.fns = append(wnfs.fns, fns...)
	return wnfs
}

// Scan applies the selector query and scans the result into the given value.
func (wnfs *WithNilFieldsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wnfs.ctx, "Select")
	if err := wnfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WithNilFieldsQuery, *WithNilFieldsSelect](ctx, wnfs.WithNilFieldsQuery, wnfs, wnfs.inters, v)
}

func (wnfs *WithNilFieldsSelect) sqlScan(ctx context.Context, root *WithNilFieldsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wnfs.fns))
	for _, fn := range wnfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wnfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wnfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
