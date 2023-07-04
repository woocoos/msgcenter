// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/msgcenter/ent/predicate"
	"github.com/woocoos/msgcenter/ent/silence"
	"github.com/woocoos/msgcenter/ent/user"

	"github.com/woocoos/msgcenter/ent/internal"
)

// SilenceQuery is the builder for querying Silence entities.
type SilenceQuery struct {
	config
	ctx        *QueryContext
	order      []silence.OrderOption
	inters     []Interceptor
	predicates []predicate.Silence
	withUser   *UserQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*Silence) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SilenceQuery builder.
func (sq *SilenceQuery) Where(ps ...predicate.Silence) *SilenceQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SilenceQuery) Limit(limit int) *SilenceQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SilenceQuery) Offset(offset int) *SilenceQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SilenceQuery) Unique(unique bool) *SilenceQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SilenceQuery) Order(o ...silence.OrderOption) *SilenceQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryUser chains the current query on the "user" edge.
func (sq *SilenceQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(silence.Table, silence.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, silence.UserTable, silence.UserColumn),
		)
		schemaConfig := sq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Silence
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Silence entity from the query.
// Returns a *NotFoundError when no Silence was found.
func (sq *SilenceQuery) First(ctx context.Context) (*Silence, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{silence.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SilenceQuery) FirstX(ctx context.Context) *Silence {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Silence ID from the query.
// Returns a *NotFoundError when no Silence ID was found.
func (sq *SilenceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{silence.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SilenceQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Silence entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Silence entity is found.
// Returns a *NotFoundError when no Silence entities are found.
func (sq *SilenceQuery) Only(ctx context.Context) (*Silence, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{silence.Label}
	default:
		return nil, &NotSingularError{silence.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SilenceQuery) OnlyX(ctx context.Context) *Silence {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Silence ID in the query.
// Returns a *NotSingularError when more than one Silence ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SilenceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{silence.Label}
	default:
		err = &NotSingularError{silence.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SilenceQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Silences.
func (sq *SilenceQuery) All(ctx context.Context) ([]*Silence, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Silence, *SilenceQuery]()
	return withInterceptors[[]*Silence](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SilenceQuery) AllX(ctx context.Context) []*Silence {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Silence IDs.
func (sq *SilenceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(silence.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SilenceQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SilenceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SilenceQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SilenceQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SilenceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SilenceQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SilenceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SilenceQuery) Clone() *SilenceQuery {
	if sq == nil {
		return nil
	}
	return &SilenceQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]silence.OrderOption{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Silence{}, sq.predicates...),
		withUser:   sq.withUser.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SilenceQuery) WithUser(opts ...func(*UserQuery)) *SilenceQuery {
	query := (&UserClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withUser = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int `json:"created_by,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Silence.Query().
//		GroupBy(silence.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SilenceQuery) GroupBy(field string, fields ...string) *SilenceGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SilenceGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = silence.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int `json:"created_by,omitempty"`
//	}
//
//	client.Silence.Query().
//		Select(silence.FieldCreatedBy).
//		Scan(ctx, &v)
func (sq *SilenceQuery) Select(fields ...string) *SilenceSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SilenceSelect{SilenceQuery: sq}
	sbuild.label = silence.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SilenceSelect configured with the given aggregations.
func (sq *SilenceQuery) Aggregate(fns ...AggregateFunc) *SilenceSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SilenceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !silence.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SilenceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Silence, error) {
	var (
		nodes       = []*Silence{}
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Silence).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Silence{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = sq.schemaConfig.Silence
	ctx = internal.NewSchemaConfigContext(ctx, sq.schemaConfig)
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withUser; query != nil {
		if err := sq.loadUser(ctx, query, nodes, nil,
			func(n *Silence, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	for i := range sq.loadTotal {
		if err := sq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SilenceQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Silence, init func(*Silence), assign func(*Silence, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Silence)
	for i := range nodes {
		fk := nodes[i].CreatedBy
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
			return fmt.Errorf(`unexpected foreign-key "created_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SilenceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Schema = sq.schemaConfig.Silence
	ctx = internal.NewSchemaConfigContext(ctx, sq.schemaConfig)
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SilenceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(silence.Table, silence.Columns, sqlgraph.NewFieldSpec(silence.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, silence.FieldID)
		for i := range fields {
			if fields[i] != silence.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if sq.withUser != nil {
			_spec.Node.AddColumnOnce(silence.FieldCreatedBy)
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SilenceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(silence.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = silence.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(sq.schemaConfig.Silence)
	ctx = internal.NewSchemaConfigContext(ctx, sq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SilenceGroupBy is the group-by builder for Silence entities.
type SilenceGroupBy struct {
	selector
	build *SilenceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SilenceGroupBy) Aggregate(fns ...AggregateFunc) *SilenceGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SilenceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SilenceQuery, *SilenceGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SilenceGroupBy) sqlScan(ctx context.Context, root *SilenceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SilenceSelect is the builder for selecting fields of Silence entities.
type SilenceSelect struct {
	*SilenceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SilenceSelect) Aggregate(fns ...AggregateFunc) *SilenceSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SilenceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SilenceQuery, *SilenceSelect](ctx, ss.SilenceQuery, ss, ss.inters, v)
}

func (ss *SilenceSelect) sqlScan(ctx context.Context, root *SilenceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}