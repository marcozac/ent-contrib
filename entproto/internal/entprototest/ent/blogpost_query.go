// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/blogpost"
	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/category"
	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/predicate"
	"github.com/marcozac/ent-contrib/entproto/internal/entprototest/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogPostQuery is the builder for querying BlogPost entities.
type BlogPostQuery struct {
	config
	ctx            *QueryContext
	order          []blogpost.OrderOption
	inters         []Interceptor
	predicates     []predicate.BlogPost
	withAuthor     *UserQuery
	withCategories *CategoryQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlogPostQuery builder.
func (bpq *BlogPostQuery) Where(ps ...predicate.BlogPost) *BlogPostQuery {
	bpq.predicates = append(bpq.predicates, ps...)
	return bpq
}

// Limit the number of records to be returned by this query.
func (bpq *BlogPostQuery) Limit(limit int) *BlogPostQuery {
	bpq.ctx.Limit = &limit
	return bpq
}

// Offset to start from.
func (bpq *BlogPostQuery) Offset(offset int) *BlogPostQuery {
	bpq.ctx.Offset = &offset
	return bpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bpq *BlogPostQuery) Unique(unique bool) *BlogPostQuery {
	bpq.ctx.Unique = &unique
	return bpq
}

// Order specifies how the records should be ordered.
func (bpq *BlogPostQuery) Order(o ...blogpost.OrderOption) *BlogPostQuery {
	bpq.order = append(bpq.order, o...)
	return bpq
}

// QueryAuthor chains the current query on the "author" edge.
func (bpq *BlogPostQuery) QueryAuthor() *UserQuery {
	query := (&UserClient{config: bpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blogpost.Table, blogpost.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, blogpost.AuthorTable, blogpost.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(bpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCategories chains the current query on the "categories" edge.
func (bpq *BlogPostQuery) QueryCategories() *CategoryQuery {
	query := (&CategoryClient{config: bpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blogpost.Table, blogpost.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, blogpost.CategoriesTable, blogpost.CategoriesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BlogPost entity from the query.
// Returns a *NotFoundError when no BlogPost was found.
func (bpq *BlogPostQuery) First(ctx context.Context) (*BlogPost, error) {
	nodes, err := bpq.Limit(1).All(setContextOp(ctx, bpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{blogpost.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bpq *BlogPostQuery) FirstX(ctx context.Context) *BlogPost {
	node, err := bpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BlogPost ID from the query.
// Returns a *NotFoundError when no BlogPost ID was found.
func (bpq *BlogPostQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bpq.Limit(1).IDs(setContextOp(ctx, bpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blogpost.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bpq *BlogPostQuery) FirstIDX(ctx context.Context) int {
	id, err := bpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BlogPost entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BlogPost entity is found.
// Returns a *NotFoundError when no BlogPost entities are found.
func (bpq *BlogPostQuery) Only(ctx context.Context) (*BlogPost, error) {
	nodes, err := bpq.Limit(2).All(setContextOp(ctx, bpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{blogpost.Label}
	default:
		return nil, &NotSingularError{blogpost.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bpq *BlogPostQuery) OnlyX(ctx context.Context) *BlogPost {
	node, err := bpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BlogPost ID in the query.
// Returns a *NotSingularError when more than one BlogPost ID is found.
// Returns a *NotFoundError when no entities are found.
func (bpq *BlogPostQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bpq.Limit(2).IDs(setContextOp(ctx, bpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = &NotSingularError{blogpost.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bpq *BlogPostQuery) OnlyIDX(ctx context.Context) int {
	id, err := bpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BlogPosts.
func (bpq *BlogPostQuery) All(ctx context.Context) ([]*BlogPost, error) {
	ctx = setContextOp(ctx, bpq.ctx, "All")
	if err := bpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BlogPost, *BlogPostQuery]()
	return withInterceptors[[]*BlogPost](ctx, bpq, qr, bpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bpq *BlogPostQuery) AllX(ctx context.Context) []*BlogPost {
	nodes, err := bpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BlogPost IDs.
func (bpq *BlogPostQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bpq.ctx.Unique == nil && bpq.path != nil {
		bpq.Unique(true)
	}
	ctx = setContextOp(ctx, bpq.ctx, "IDs")
	if err = bpq.Select(blogpost.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bpq *BlogPostQuery) IDsX(ctx context.Context) []int {
	ids, err := bpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bpq *BlogPostQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bpq.ctx, "Count")
	if err := bpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bpq, querierCount[*BlogPostQuery](), bpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bpq *BlogPostQuery) CountX(ctx context.Context) int {
	count, err := bpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bpq *BlogPostQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bpq.ctx, "Exist")
	switch _, err := bpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bpq *BlogPostQuery) ExistX(ctx context.Context) bool {
	exist, err := bpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlogPostQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bpq *BlogPostQuery) Clone() *BlogPostQuery {
	if bpq == nil {
		return nil
	}
	return &BlogPostQuery{
		config:         bpq.config,
		ctx:            bpq.ctx.Clone(),
		order:          append([]blogpost.OrderOption{}, bpq.order...),
		inters:         append([]Interceptor{}, bpq.inters...),
		predicates:     append([]predicate.BlogPost{}, bpq.predicates...),
		withAuthor:     bpq.withAuthor.Clone(),
		withCategories: bpq.withCategories.Clone(),
		// clone intermediate query.
		sql:  bpq.sql.Clone(),
		path: bpq.path,
	}
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (bpq *BlogPostQuery) WithAuthor(opts ...func(*UserQuery)) *BlogPostQuery {
	query := (&UserClient{config: bpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bpq.withAuthor = query
	return bpq
}

// WithCategories tells the query-builder to eager-load the nodes that are connected to
// the "categories" edge. The optional arguments are used to configure the query builder of the edge.
func (bpq *BlogPostQuery) WithCategories(opts ...func(*CategoryQuery)) *BlogPostQuery {
	query := (&CategoryClient{config: bpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bpq.withCategories = query
	return bpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BlogPost.Query().
//		GroupBy(blogpost.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bpq *BlogPostQuery) GroupBy(field string, fields ...string) *BlogPostGroupBy {
	bpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BlogPostGroupBy{build: bpq}
	grbuild.flds = &bpq.ctx.Fields
	grbuild.label = blogpost.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.BlogPost.Query().
//		Select(blogpost.FieldTitle).
//		Scan(ctx, &v)
func (bpq *BlogPostQuery) Select(fields ...string) *BlogPostSelect {
	bpq.ctx.Fields = append(bpq.ctx.Fields, fields...)
	sbuild := &BlogPostSelect{BlogPostQuery: bpq}
	sbuild.label = blogpost.Label
	sbuild.flds, sbuild.scan = &bpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BlogPostSelect configured with the given aggregations.
func (bpq *BlogPostQuery) Aggregate(fns ...AggregateFunc) *BlogPostSelect {
	return bpq.Select().Aggregate(fns...)
}

func (bpq *BlogPostQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bpq); err != nil {
				return err
			}
		}
	}
	for _, f := range bpq.ctx.Fields {
		if !blogpost.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bpq.path != nil {
		prev, err := bpq.path(ctx)
		if err != nil {
			return err
		}
		bpq.sql = prev
	}
	return nil
}

func (bpq *BlogPostQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BlogPost, error) {
	var (
		nodes       = []*BlogPost{}
		withFKs     = bpq.withFKs
		_spec       = bpq.querySpec()
		loadedTypes = [2]bool{
			bpq.withAuthor != nil,
			bpq.withCategories != nil,
		}
	)
	if bpq.withAuthor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, blogpost.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BlogPost).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BlogPost{config: bpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bpq.withAuthor; query != nil {
		if err := bpq.loadAuthor(ctx, query, nodes, nil,
			func(n *BlogPost, e *User) { n.Edges.Author = e }); err != nil {
			return nil, err
		}
	}
	if query := bpq.withCategories; query != nil {
		if err := bpq.loadCategories(ctx, query, nodes,
			func(n *BlogPost) { n.Edges.Categories = []*Category{} },
			func(n *BlogPost, e *Category) { n.Edges.Categories = append(n.Edges.Categories, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bpq *BlogPostQuery) loadAuthor(ctx context.Context, query *UserQuery, nodes []*BlogPost, init func(*BlogPost), assign func(*BlogPost, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*BlogPost)
	for i := range nodes {
		if nodes[i].blog_post_author == nil {
			continue
		}
		fk := *nodes[i].blog_post_author
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "blog_post_author" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bpq *BlogPostQuery) loadCategories(ctx context.Context, query *CategoryQuery, nodes []*BlogPost, init func(*BlogPost), assign func(*BlogPost, *Category)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*BlogPost)
	nids := make(map[int]map[*BlogPost]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(blogpost.CategoriesTable)
		s.Join(joinT).On(s.C(category.FieldID), joinT.C(blogpost.CategoriesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(blogpost.CategoriesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(blogpost.CategoriesPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*BlogPost]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Category](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "categories" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (bpq *BlogPostQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bpq.querySpec()
	_spec.Node.Columns = bpq.ctx.Fields
	if len(bpq.ctx.Fields) > 0 {
		_spec.Unique = bpq.ctx.Unique != nil && *bpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bpq.driver, _spec)
}

func (bpq *BlogPostQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(blogpost.Table, blogpost.Columns, sqlgraph.NewFieldSpec(blogpost.FieldID, field.TypeInt))
	_spec.From = bpq.sql
	if unique := bpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bpq.path != nil {
		_spec.Unique = true
	}
	if fields := bpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blogpost.FieldID)
		for i := range fields {
			if fields[i] != blogpost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bpq *BlogPostQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bpq.driver.Dialect())
	t1 := builder.Table(blogpost.Table)
	columns := bpq.ctx.Fields
	if len(columns) == 0 {
		columns = blogpost.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bpq.sql != nil {
		selector = bpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bpq.ctx.Unique != nil && *bpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bpq.predicates {
		p(selector)
	}
	for _, p := range bpq.order {
		p(selector)
	}
	if offset := bpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BlogPostGroupBy is the group-by builder for BlogPost entities.
type BlogPostGroupBy struct {
	selector
	build *BlogPostQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bpgb *BlogPostGroupBy) Aggregate(fns ...AggregateFunc) *BlogPostGroupBy {
	bpgb.fns = append(bpgb.fns, fns...)
	return bpgb
}

// Scan applies the selector query and scans the result into the given value.
func (bpgb *BlogPostGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bpgb.build.ctx, "GroupBy")
	if err := bpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlogPostQuery, *BlogPostGroupBy](ctx, bpgb.build, bpgb, bpgb.build.inters, v)
}

func (bpgb *BlogPostGroupBy) sqlScan(ctx context.Context, root *BlogPostQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bpgb.fns))
	for _, fn := range bpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bpgb.flds)+len(bpgb.fns))
		for _, f := range *bpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BlogPostSelect is the builder for selecting fields of BlogPost entities.
type BlogPostSelect struct {
	*BlogPostQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bps *BlogPostSelect) Aggregate(fns ...AggregateFunc) *BlogPostSelect {
	bps.fns = append(bps.fns, fns...)
	return bps
}

// Scan applies the selector query and scans the result into the given value.
func (bps *BlogPostSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bps.ctx, "Select")
	if err := bps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlogPostQuery, *BlogPostSelect](ctx, bps.BlogPostQuery, bps, bps.inters, v)
}

func (bps *BlogPostSelect) sqlScan(ctx context.Context, root *BlogPostQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bps.fns))
	for _, fn := range bps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
