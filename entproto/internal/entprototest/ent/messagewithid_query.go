// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/messagewithid"
	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithIDQuery is the builder for querying MessageWithID entities.
type MessageWithIDQuery struct {
	config
	ctx        *QueryContext
	order      []messagewithid.OrderOption
	inters     []Interceptor
	predicates []predicate.MessageWithID
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageWithIDQuery builder.
func (mwiq *MessageWithIDQuery) Where(ps ...predicate.MessageWithID) *MessageWithIDQuery {
	mwiq.predicates = append(mwiq.predicates, ps...)
	return mwiq
}

// Limit the number of records to be returned by this query.
func (mwiq *MessageWithIDQuery) Limit(limit int) *MessageWithIDQuery {
	mwiq.ctx.Limit = &limit
	return mwiq
}

// Offset to start from.
func (mwiq *MessageWithIDQuery) Offset(offset int) *MessageWithIDQuery {
	mwiq.ctx.Offset = &offset
	return mwiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mwiq *MessageWithIDQuery) Unique(unique bool) *MessageWithIDQuery {
	mwiq.ctx.Unique = &unique
	return mwiq
}

// Order specifies how the records should be ordered.
func (mwiq *MessageWithIDQuery) Order(o ...messagewithid.OrderOption) *MessageWithIDQuery {
	mwiq.order = append(mwiq.order, o...)
	return mwiq
}

// First returns the first MessageWithID entity from the query.
// Returns a *NotFoundError when no MessageWithID was found.
func (mwiq *MessageWithIDQuery) First(ctx context.Context) (*MessageWithID, error) {
	nodes, err := mwiq.Limit(1).All(setContextOp(ctx, mwiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{messagewithid.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mwiq *MessageWithIDQuery) FirstX(ctx context.Context) *MessageWithID {
	node, err := mwiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MessageWithID ID from the query.
// Returns a *NotFoundError when no MessageWithID ID was found.
func (mwiq *MessageWithIDQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = mwiq.Limit(1).IDs(setContextOp(ctx, mwiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{messagewithid.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mwiq *MessageWithIDQuery) FirstIDX(ctx context.Context) int32 {
	id, err := mwiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MessageWithID entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MessageWithID entity is found.
// Returns a *NotFoundError when no MessageWithID entities are found.
func (mwiq *MessageWithIDQuery) Only(ctx context.Context) (*MessageWithID, error) {
	nodes, err := mwiq.Limit(2).All(setContextOp(ctx, mwiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{messagewithid.Label}
	default:
		return nil, &NotSingularError{messagewithid.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mwiq *MessageWithIDQuery) OnlyX(ctx context.Context) *MessageWithID {
	node, err := mwiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MessageWithID ID in the query.
// Returns a *NotSingularError when more than one MessageWithID ID is found.
// Returns a *NotFoundError when no entities are found.
func (mwiq *MessageWithIDQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = mwiq.Limit(2).IDs(setContextOp(ctx, mwiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{messagewithid.Label}
	default:
		err = &NotSingularError{messagewithid.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mwiq *MessageWithIDQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := mwiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MessageWithIDs.
func (mwiq *MessageWithIDQuery) All(ctx context.Context) ([]*MessageWithID, error) {
	ctx = setContextOp(ctx, mwiq.ctx, "All")
	if err := mwiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MessageWithID, *MessageWithIDQuery]()
	return withInterceptors[[]*MessageWithID](ctx, mwiq, qr, mwiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mwiq *MessageWithIDQuery) AllX(ctx context.Context) []*MessageWithID {
	nodes, err := mwiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MessageWithID IDs.
func (mwiq *MessageWithIDQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if mwiq.ctx.Unique == nil && mwiq.path != nil {
		mwiq.Unique(true)
	}
	ctx = setContextOp(ctx, mwiq.ctx, "IDs")
	if err = mwiq.Select(messagewithid.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mwiq *MessageWithIDQuery) IDsX(ctx context.Context) []int32 {
	ids, err := mwiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mwiq *MessageWithIDQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mwiq.ctx, "Count")
	if err := mwiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mwiq, querierCount[*MessageWithIDQuery](), mwiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mwiq *MessageWithIDQuery) CountX(ctx context.Context) int {
	count, err := mwiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mwiq *MessageWithIDQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mwiq.ctx, "Exist")
	switch _, err := mwiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mwiq *MessageWithIDQuery) ExistX(ctx context.Context) bool {
	exist, err := mwiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageWithIDQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mwiq *MessageWithIDQuery) Clone() *MessageWithIDQuery {
	if mwiq == nil {
		return nil
	}
	return &MessageWithIDQuery{
		config:     mwiq.config,
		ctx:        mwiq.ctx.Clone(),
		order:      append([]messagewithid.OrderOption{}, mwiq.order...),
		inters:     append([]Interceptor{}, mwiq.inters...),
		predicates: append([]predicate.MessageWithID{}, mwiq.predicates...),
		// clone intermediate query.
		sql:  mwiq.sql.Clone(),
		path: mwiq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (mwiq *MessageWithIDQuery) GroupBy(field string, fields ...string) *MessageWithIDGroupBy {
	mwiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MessageWithIDGroupBy{build: mwiq}
	grbuild.flds = &mwiq.ctx.Fields
	grbuild.label = messagewithid.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (mwiq *MessageWithIDQuery) Select(fields ...string) *MessageWithIDSelect {
	mwiq.ctx.Fields = append(mwiq.ctx.Fields, fields...)
	sbuild := &MessageWithIDSelect{MessageWithIDQuery: mwiq}
	sbuild.label = messagewithid.Label
	sbuild.flds, sbuild.scan = &mwiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MessageWithIDSelect configured with the given aggregations.
func (mwiq *MessageWithIDQuery) Aggregate(fns ...AggregateFunc) *MessageWithIDSelect {
	return mwiq.Select().Aggregate(fns...)
}

func (mwiq *MessageWithIDQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mwiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mwiq); err != nil {
				return err
			}
		}
	}
	for _, f := range mwiq.ctx.Fields {
		if !messagewithid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mwiq.path != nil {
		prev, err := mwiq.path(ctx)
		if err != nil {
			return err
		}
		mwiq.sql = prev
	}
	return nil
}

func (mwiq *MessageWithIDQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MessageWithID, error) {
	var (
		nodes = []*MessageWithID{}
		_spec = mwiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MessageWithID).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MessageWithID{config: mwiq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mwiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mwiq *MessageWithIDQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mwiq.querySpec()
	_spec.Node.Columns = mwiq.ctx.Fields
	if len(mwiq.ctx.Fields) > 0 {
		_spec.Unique = mwiq.ctx.Unique != nil && *mwiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mwiq.driver, _spec)
}

func (mwiq *MessageWithIDQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(messagewithid.Table, messagewithid.Columns, sqlgraph.NewFieldSpec(messagewithid.FieldID, field.TypeInt32))
	_spec.From = mwiq.sql
	if unique := mwiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mwiq.path != nil {
		_spec.Unique = true
	}
	if fields := mwiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithid.FieldID)
		for i := range fields {
			if fields[i] != messagewithid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mwiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mwiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mwiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mwiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mwiq *MessageWithIDQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mwiq.driver.Dialect())
	t1 := builder.Table(messagewithid.Table)
	columns := mwiq.ctx.Fields
	if len(columns) == 0 {
		columns = messagewithid.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mwiq.sql != nil {
		selector = mwiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mwiq.ctx.Unique != nil && *mwiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mwiq.predicates {
		p(selector)
	}
	for _, p := range mwiq.order {
		p(selector)
	}
	if offset := mwiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mwiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MessageWithIDGroupBy is the group-by builder for MessageWithID entities.
type MessageWithIDGroupBy struct {
	selector
	build *MessageWithIDQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mwigb *MessageWithIDGroupBy) Aggregate(fns ...AggregateFunc) *MessageWithIDGroupBy {
	mwigb.fns = append(mwigb.fns, fns...)
	return mwigb
}

// Scan applies the selector query and scans the result into the given value.
func (mwigb *MessageWithIDGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mwigb.build.ctx, "GroupBy")
	if err := mwigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MessageWithIDQuery, *MessageWithIDGroupBy](ctx, mwigb.build, mwigb, mwigb.build.inters, v)
}

func (mwigb *MessageWithIDGroupBy) sqlScan(ctx context.Context, root *MessageWithIDQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mwigb.fns))
	for _, fn := range mwigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mwigb.flds)+len(mwigb.fns))
		for _, f := range *mwigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mwigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mwigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MessageWithIDSelect is the builder for selecting fields of MessageWithID entities.
type MessageWithIDSelect struct {
	*MessageWithIDQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mwis *MessageWithIDSelect) Aggregate(fns ...AggregateFunc) *MessageWithIDSelect {
	mwis.fns = append(mwis.fns, fns...)
	return mwis
}

// Scan applies the selector query and scans the result into the given value.
func (mwis *MessageWithIDSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mwis.ctx, "Select")
	if err := mwis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MessageWithIDQuery, *MessageWithIDSelect](ctx, mwis.MessageWithIDQuery, mwis, mwis.inters, v)
}

func (mwis *MessageWithIDSelect) sqlScan(ctx context.Context, root *MessageWithIDQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mwis.fns))
	for _, fn := range mwis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mwis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mwis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
