// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/category"
	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/portal"
	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PortalQuery is the builder for querying Portal entities.
type PortalQuery struct {
	config
	ctx          *QueryContext
	order        []portal.OrderOption
	inters       []Interceptor
	predicates   []predicate.Portal
	withCategory *CategoryQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PortalQuery builder.
func (pq *PortalQuery) Where(ps ...predicate.Portal) *PortalQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PortalQuery) Limit(limit int) *PortalQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PortalQuery) Offset(offset int) *PortalQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PortalQuery) Unique(unique bool) *PortalQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PortalQuery) Order(o ...portal.OrderOption) *PortalQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryCategory chains the current query on the "category" edge.
func (pq *PortalQuery) QueryCategory() *CategoryQuery {
	query := (&CategoryClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(portal.Table, portal.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, portal.CategoryTable, portal.CategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Portal entity from the query.
// Returns a *NotFoundError when no Portal was found.
func (pq *PortalQuery) First(ctx context.Context) (*Portal, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{portal.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PortalQuery) FirstX(ctx context.Context) *Portal {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Portal ID from the query.
// Returns a *NotFoundError when no Portal ID was found.
func (pq *PortalQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{portal.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PortalQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Portal entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Portal entity is found.
// Returns a *NotFoundError when no Portal entities are found.
func (pq *PortalQuery) Only(ctx context.Context) (*Portal, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{portal.Label}
	default:
		return nil, &NotSingularError{portal.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PortalQuery) OnlyX(ctx context.Context) *Portal {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Portal ID in the query.
// Returns a *NotSingularError when more than one Portal ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PortalQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{portal.Label}
	default:
		err = &NotSingularError{portal.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PortalQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Portals.
func (pq *PortalQuery) All(ctx context.Context) ([]*Portal, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Portal, *PortalQuery]()
	return withInterceptors[[]*Portal](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PortalQuery) AllX(ctx context.Context) []*Portal {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Portal IDs.
func (pq *PortalQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(portal.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PortalQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PortalQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PortalQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PortalQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PortalQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PortalQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PortalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PortalQuery) Clone() *PortalQuery {
	if pq == nil {
		return nil
	}
	return &PortalQuery{
		config:       pq.config,
		ctx:          pq.ctx.Clone(),
		order:        append([]portal.OrderOption{}, pq.order...),
		inters:       append([]Interceptor{}, pq.inters...),
		predicates:   append([]predicate.Portal{}, pq.predicates...),
		withCategory: pq.withCategory.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithCategory tells the query-builder to eager-load the nodes that are connected to
// the "category" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PortalQuery) WithCategory(opts ...func(*CategoryQuery)) *PortalQuery {
	query := (&CategoryClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withCategory = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Portal.Query().
//		GroupBy(portal.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PortalQuery) GroupBy(field string, fields ...string) *PortalGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PortalGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = portal.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Portal.Query().
//		Select(portal.FieldName).
//		Scan(ctx, &v)
func (pq *PortalQuery) Select(fields ...string) *PortalSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PortalSelect{PortalQuery: pq}
	sbuild.label = portal.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PortalSelect configured with the given aggregations.
func (pq *PortalQuery) Aggregate(fns ...AggregateFunc) *PortalSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PortalQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !portal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PortalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Portal, error) {
	var (
		nodes       = []*Portal{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withCategory != nil,
		}
	)
	if pq.withCategory != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, portal.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Portal).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Portal{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withCategory; query != nil {
		if err := pq.loadCategory(ctx, query, nodes, nil,
			func(n *Portal, e *Category) { n.Edges.Category = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PortalQuery) loadCategory(ctx context.Context, query *CategoryQuery, nodes []*Portal, init func(*Portal), assign func(*Portal, *Category)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Portal)
	for i := range nodes {
		if nodes[i].portal_category == nil {
			continue
		}
		fk := *nodes[i].portal_category
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(category.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "portal_category" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PortalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PortalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(portal.Table, portal.Columns, sqlgraph.NewFieldSpec(portal.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, portal.FieldID)
		for i := range fields {
			if fields[i] != portal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PortalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(portal.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = portal.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PortalGroupBy is the group-by builder for Portal entities.
type PortalGroupBy struct {
	selector
	build *PortalQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PortalGroupBy) Aggregate(fns ...AggregateFunc) *PortalGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PortalGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PortalQuery, *PortalGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PortalGroupBy) sqlScan(ctx context.Context, root *PortalQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PortalSelect is the builder for selecting fields of Portal entities.
type PortalSelect struct {
	*PortalQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PortalSelect) Aggregate(fns ...AggregateFunc) *PortalSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PortalSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PortalQuery, *PortalSelect](ctx, ps.PortalQuery, ps, ps.inters, v)
}

func (ps *PortalSelect) sqlScan(ctx context.Context, root *PortalQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
