// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/msgcenter/ent/msginternal"
	"github.com/woocoos/msgcenter/ent/msginternalto"
	"github.com/woocoos/msgcenter/ent/predicate"

	"github.com/woocoos/msgcenter/ent/internal"
)

// MsgInternalQuery is the builder for querying MsgInternal entities.
type MsgInternalQuery struct {
	config
	ctx                    *QueryContext
	order                  []msginternal.OrderOption
	inters                 []Interceptor
	predicates             []predicate.MsgInternal
	withMsgInternalTo      *MsgInternalToQuery
	modifiers              []func(*sql.Selector)
	loadTotal              []func(context.Context, []*MsgInternal) error
	withNamedMsgInternalTo map[string]*MsgInternalToQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MsgInternalQuery builder.
func (miq *MsgInternalQuery) Where(ps ...predicate.MsgInternal) *MsgInternalQuery {
	miq.predicates = append(miq.predicates, ps...)
	return miq
}

// Limit the number of records to be returned by this query.
func (miq *MsgInternalQuery) Limit(limit int) *MsgInternalQuery {
	miq.ctx.Limit = &limit
	return miq
}

// Offset to start from.
func (miq *MsgInternalQuery) Offset(offset int) *MsgInternalQuery {
	miq.ctx.Offset = &offset
	return miq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (miq *MsgInternalQuery) Unique(unique bool) *MsgInternalQuery {
	miq.ctx.Unique = &unique
	return miq
}

// Order specifies how the records should be ordered.
func (miq *MsgInternalQuery) Order(o ...msginternal.OrderOption) *MsgInternalQuery {
	miq.order = append(miq.order, o...)
	return miq
}

// QueryMsgInternalTo chains the current query on the "msg_internal_to" edge.
func (miq *MsgInternalQuery) QueryMsgInternalTo() *MsgInternalToQuery {
	query := (&MsgInternalToClient{config: miq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := miq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(msginternal.Table, msginternal.FieldID, selector),
			sqlgraph.To(msginternalto.Table, msginternalto.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, msginternal.MsgInternalToTable, msginternal.MsgInternalToColumn),
		)
		schemaConfig := miq.schemaConfig
		step.To.Schema = schemaConfig.MsgInternalTo
		step.Edge.Schema = schemaConfig.MsgInternalTo
		fromU = sqlgraph.SetNeighbors(miq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MsgInternal entity from the query.
// Returns a *NotFoundError when no MsgInternal was found.
func (miq *MsgInternalQuery) First(ctx context.Context) (*MsgInternal, error) {
	nodes, err := miq.Limit(1).All(setContextOp(ctx, miq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{msginternal.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (miq *MsgInternalQuery) FirstX(ctx context.Context) *MsgInternal {
	node, err := miq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MsgInternal ID from the query.
// Returns a *NotFoundError when no MsgInternal ID was found.
func (miq *MsgInternalQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = miq.Limit(1).IDs(setContextOp(ctx, miq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{msginternal.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (miq *MsgInternalQuery) FirstIDX(ctx context.Context) int {
	id, err := miq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MsgInternal entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MsgInternal entity is found.
// Returns a *NotFoundError when no MsgInternal entities are found.
func (miq *MsgInternalQuery) Only(ctx context.Context) (*MsgInternal, error) {
	nodes, err := miq.Limit(2).All(setContextOp(ctx, miq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{msginternal.Label}
	default:
		return nil, &NotSingularError{msginternal.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (miq *MsgInternalQuery) OnlyX(ctx context.Context) *MsgInternal {
	node, err := miq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MsgInternal ID in the query.
// Returns a *NotSingularError when more than one MsgInternal ID is found.
// Returns a *NotFoundError when no entities are found.
func (miq *MsgInternalQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = miq.Limit(2).IDs(setContextOp(ctx, miq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{msginternal.Label}
	default:
		err = &NotSingularError{msginternal.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (miq *MsgInternalQuery) OnlyIDX(ctx context.Context) int {
	id, err := miq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MsgInternals.
func (miq *MsgInternalQuery) All(ctx context.Context) ([]*MsgInternal, error) {
	ctx = setContextOp(ctx, miq.ctx, ent.OpQueryAll)
	if err := miq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MsgInternal, *MsgInternalQuery]()
	return withInterceptors[[]*MsgInternal](ctx, miq, qr, miq.inters)
}

// AllX is like All, but panics if an error occurs.
func (miq *MsgInternalQuery) AllX(ctx context.Context) []*MsgInternal {
	nodes, err := miq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MsgInternal IDs.
func (miq *MsgInternalQuery) IDs(ctx context.Context) (ids []int, err error) {
	if miq.ctx.Unique == nil && miq.path != nil {
		miq.Unique(true)
	}
	ctx = setContextOp(ctx, miq.ctx, ent.OpQueryIDs)
	if err = miq.Select(msginternal.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (miq *MsgInternalQuery) IDsX(ctx context.Context) []int {
	ids, err := miq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (miq *MsgInternalQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, miq.ctx, ent.OpQueryCount)
	if err := miq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, miq, querierCount[*MsgInternalQuery](), miq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (miq *MsgInternalQuery) CountX(ctx context.Context) int {
	count, err := miq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (miq *MsgInternalQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, miq.ctx, ent.OpQueryExist)
	switch _, err := miq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (miq *MsgInternalQuery) ExistX(ctx context.Context) bool {
	exist, err := miq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MsgInternalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (miq *MsgInternalQuery) Clone() *MsgInternalQuery {
	if miq == nil {
		return nil
	}
	return &MsgInternalQuery{
		config:            miq.config,
		ctx:               miq.ctx.Clone(),
		order:             append([]msginternal.OrderOption{}, miq.order...),
		inters:            append([]Interceptor{}, miq.inters...),
		predicates:        append([]predicate.MsgInternal{}, miq.predicates...),
		withMsgInternalTo: miq.withMsgInternalTo.Clone(),
		// clone intermediate query.
		sql:  miq.sql.Clone(),
		path: miq.path,
	}
}

// WithMsgInternalTo tells the query-builder to eager-load the nodes that are connected to
// the "msg_internal_to" edge. The optional arguments are used to configure the query builder of the edge.
func (miq *MsgInternalQuery) WithMsgInternalTo(opts ...func(*MsgInternalToQuery)) *MsgInternalQuery {
	query := (&MsgInternalToClient{config: miq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	miq.withMsgInternalTo = query
	return miq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TenantID int `json:"tenant_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MsgInternal.Query().
//		GroupBy(msginternal.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (miq *MsgInternalQuery) GroupBy(field string, fields ...string) *MsgInternalGroupBy {
	miq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MsgInternalGroupBy{build: miq}
	grbuild.flds = &miq.ctx.Fields
	grbuild.label = msginternal.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TenantID int `json:"tenant_id,omitempty"`
//	}
//
//	client.MsgInternal.Query().
//		Select(msginternal.FieldTenantID).
//		Scan(ctx, &v)
func (miq *MsgInternalQuery) Select(fields ...string) *MsgInternalSelect {
	miq.ctx.Fields = append(miq.ctx.Fields, fields...)
	sbuild := &MsgInternalSelect{MsgInternalQuery: miq}
	sbuild.label = msginternal.Label
	sbuild.flds, sbuild.scan = &miq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MsgInternalSelect configured with the given aggregations.
func (miq *MsgInternalQuery) Aggregate(fns ...AggregateFunc) *MsgInternalSelect {
	return miq.Select().Aggregate(fns...)
}

func (miq *MsgInternalQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range miq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, miq); err != nil {
				return err
			}
		}
	}
	for _, f := range miq.ctx.Fields {
		if !msginternal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if miq.path != nil {
		prev, err := miq.path(ctx)
		if err != nil {
			return err
		}
		miq.sql = prev
	}
	return nil
}

func (miq *MsgInternalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MsgInternal, error) {
	var (
		nodes       = []*MsgInternal{}
		_spec       = miq.querySpec()
		loadedTypes = [1]bool{
			miq.withMsgInternalTo != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MsgInternal).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MsgInternal{config: miq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = miq.schemaConfig.MsgInternal
	ctx = internal.NewSchemaConfigContext(ctx, miq.schemaConfig)
	if len(miq.modifiers) > 0 {
		_spec.Modifiers = miq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, miq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := miq.withMsgInternalTo; query != nil {
		if err := miq.loadMsgInternalTo(ctx, query, nodes,
			func(n *MsgInternal) { n.Edges.MsgInternalTo = []*MsgInternalTo{} },
			func(n *MsgInternal, e *MsgInternalTo) { n.Edges.MsgInternalTo = append(n.Edges.MsgInternalTo, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range miq.withNamedMsgInternalTo {
		if err := miq.loadMsgInternalTo(ctx, query, nodes,
			func(n *MsgInternal) { n.appendNamedMsgInternalTo(name) },
			func(n *MsgInternal, e *MsgInternalTo) { n.appendNamedMsgInternalTo(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range miq.loadTotal {
		if err := miq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (miq *MsgInternalQuery) loadMsgInternalTo(ctx context.Context, query *MsgInternalToQuery, nodes []*MsgInternal, init func(*MsgInternal), assign func(*MsgInternal, *MsgInternalTo)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*MsgInternal)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(msginternalto.FieldMsgInternalID)
	}
	query.Where(predicate.MsgInternalTo(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(msginternal.MsgInternalToColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MsgInternalID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "msg_internal_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (miq *MsgInternalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := miq.querySpec()
	_spec.Node.Schema = miq.schemaConfig.MsgInternal
	ctx = internal.NewSchemaConfigContext(ctx, miq.schemaConfig)
	if len(miq.modifiers) > 0 {
		_spec.Modifiers = miq.modifiers
	}
	_spec.Node.Columns = miq.ctx.Fields
	if len(miq.ctx.Fields) > 0 {
		_spec.Unique = miq.ctx.Unique != nil && *miq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, miq.driver, _spec)
}

func (miq *MsgInternalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(msginternal.Table, msginternal.Columns, sqlgraph.NewFieldSpec(msginternal.FieldID, field.TypeInt))
	_spec.From = miq.sql
	if unique := miq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if miq.path != nil {
		_spec.Unique = true
	}
	if fields := miq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, msginternal.FieldID)
		for i := range fields {
			if fields[i] != msginternal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := miq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := miq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := miq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := miq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (miq *MsgInternalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(miq.driver.Dialect())
	t1 := builder.Table(msginternal.Table)
	columns := miq.ctx.Fields
	if len(columns) == 0 {
		columns = msginternal.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if miq.sql != nil {
		selector = miq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if miq.ctx.Unique != nil && *miq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(miq.schemaConfig.MsgInternal)
	ctx = internal.NewSchemaConfigContext(ctx, miq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range miq.predicates {
		p(selector)
	}
	for _, p := range miq.order {
		p(selector)
	}
	if offset := miq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := miq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedMsgInternalTo tells the query-builder to eager-load the nodes that are connected to the "msg_internal_to"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (miq *MsgInternalQuery) WithNamedMsgInternalTo(name string, opts ...func(*MsgInternalToQuery)) *MsgInternalQuery {
	query := (&MsgInternalToClient{config: miq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if miq.withNamedMsgInternalTo == nil {
		miq.withNamedMsgInternalTo = make(map[string]*MsgInternalToQuery)
	}
	miq.withNamedMsgInternalTo[name] = query
	return miq
}

// MsgInternalGroupBy is the group-by builder for MsgInternal entities.
type MsgInternalGroupBy struct {
	selector
	build *MsgInternalQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (migb *MsgInternalGroupBy) Aggregate(fns ...AggregateFunc) *MsgInternalGroupBy {
	migb.fns = append(migb.fns, fns...)
	return migb
}

// Scan applies the selector query and scans the result into the given value.
func (migb *MsgInternalGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, migb.build.ctx, ent.OpQueryGroupBy)
	if err := migb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MsgInternalQuery, *MsgInternalGroupBy](ctx, migb.build, migb, migb.build.inters, v)
}

func (migb *MsgInternalGroupBy) sqlScan(ctx context.Context, root *MsgInternalQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(migb.fns))
	for _, fn := range migb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*migb.flds)+len(migb.fns))
		for _, f := range *migb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*migb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := migb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MsgInternalSelect is the builder for selecting fields of MsgInternal entities.
type MsgInternalSelect struct {
	*MsgInternalQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mis *MsgInternalSelect) Aggregate(fns ...AggregateFunc) *MsgInternalSelect {
	mis.fns = append(mis.fns, fns...)
	return mis
}

// Scan applies the selector query and scans the result into the given value.
func (mis *MsgInternalSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mis.ctx, ent.OpQuerySelect)
	if err := mis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MsgInternalQuery, *MsgInternalSelect](ctx, mis.MsgInternalQuery, mis, mis.inters, v)
}

func (mis *MsgInternalSelect) sqlScan(ctx context.Context, root *MsgInternalQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mis.fns))
	for _, fn := range mis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
