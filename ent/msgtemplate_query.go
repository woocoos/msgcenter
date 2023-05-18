// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/ent/predicate"
)

// MsgTemplateQuery is the builder for querying MsgTemplate entities.
type MsgTemplateQuery struct {
	config
	ctx        *QueryContext
	order      []msgtemplate.OrderOption
	inters     []Interceptor
	predicates []predicate.MsgTemplate
	withEvent  *MsgEventQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*MsgTemplate) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MsgTemplateQuery builder.
func (mtq *MsgTemplateQuery) Where(ps ...predicate.MsgTemplate) *MsgTemplateQuery {
	mtq.predicates = append(mtq.predicates, ps...)
	return mtq
}

// Limit the number of records to be returned by this query.
func (mtq *MsgTemplateQuery) Limit(limit int) *MsgTemplateQuery {
	mtq.ctx.Limit = &limit
	return mtq
}

// Offset to start from.
func (mtq *MsgTemplateQuery) Offset(offset int) *MsgTemplateQuery {
	mtq.ctx.Offset = &offset
	return mtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mtq *MsgTemplateQuery) Unique(unique bool) *MsgTemplateQuery {
	mtq.ctx.Unique = &unique
	return mtq
}

// Order specifies how the records should be ordered.
func (mtq *MsgTemplateQuery) Order(o ...msgtemplate.OrderOption) *MsgTemplateQuery {
	mtq.order = append(mtq.order, o...)
	return mtq
}

// QueryEvent chains the current query on the "event" edge.
func (mtq *MsgTemplateQuery) QueryEvent() *MsgEventQuery {
	query := (&MsgEventClient{config: mtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(msgtemplate.Table, msgtemplate.FieldID, selector),
			sqlgraph.To(msgevent.Table, msgevent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, msgtemplate.EventTable, msgtemplate.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(mtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MsgTemplate entity from the query.
// Returns a *NotFoundError when no MsgTemplate was found.
func (mtq *MsgTemplateQuery) First(ctx context.Context) (*MsgTemplate, error) {
	nodes, err := mtq.Limit(1).All(setContextOp(ctx, mtq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{msgtemplate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mtq *MsgTemplateQuery) FirstX(ctx context.Context) *MsgTemplate {
	node, err := mtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MsgTemplate ID from the query.
// Returns a *NotFoundError when no MsgTemplate ID was found.
func (mtq *MsgTemplateQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mtq.Limit(1).IDs(setContextOp(ctx, mtq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{msgtemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mtq *MsgTemplateQuery) FirstIDX(ctx context.Context) int {
	id, err := mtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MsgTemplate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MsgTemplate entity is found.
// Returns a *NotFoundError when no MsgTemplate entities are found.
func (mtq *MsgTemplateQuery) Only(ctx context.Context) (*MsgTemplate, error) {
	nodes, err := mtq.Limit(2).All(setContextOp(ctx, mtq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{msgtemplate.Label}
	default:
		return nil, &NotSingularError{msgtemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mtq *MsgTemplateQuery) OnlyX(ctx context.Context) *MsgTemplate {
	node, err := mtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MsgTemplate ID in the query.
// Returns a *NotSingularError when more than one MsgTemplate ID is found.
// Returns a *NotFoundError when no entities are found.
func (mtq *MsgTemplateQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mtq.Limit(2).IDs(setContextOp(ctx, mtq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{msgtemplate.Label}
	default:
		err = &NotSingularError{msgtemplate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mtq *MsgTemplateQuery) OnlyIDX(ctx context.Context) int {
	id, err := mtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MsgTemplates.
func (mtq *MsgTemplateQuery) All(ctx context.Context) ([]*MsgTemplate, error) {
	ctx = setContextOp(ctx, mtq.ctx, "All")
	if err := mtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MsgTemplate, *MsgTemplateQuery]()
	return withInterceptors[[]*MsgTemplate](ctx, mtq, qr, mtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mtq *MsgTemplateQuery) AllX(ctx context.Context) []*MsgTemplate {
	nodes, err := mtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MsgTemplate IDs.
func (mtq *MsgTemplateQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mtq.ctx.Unique == nil && mtq.path != nil {
		mtq.Unique(true)
	}
	ctx = setContextOp(ctx, mtq.ctx, "IDs")
	if err = mtq.Select(msgtemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mtq *MsgTemplateQuery) IDsX(ctx context.Context) []int {
	ids, err := mtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mtq *MsgTemplateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mtq.ctx, "Count")
	if err := mtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mtq, querierCount[*MsgTemplateQuery](), mtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mtq *MsgTemplateQuery) CountX(ctx context.Context) int {
	count, err := mtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mtq *MsgTemplateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mtq.ctx, "Exist")
	switch _, err := mtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mtq *MsgTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := mtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MsgTemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mtq *MsgTemplateQuery) Clone() *MsgTemplateQuery {
	if mtq == nil {
		return nil
	}
	return &MsgTemplateQuery{
		config:     mtq.config,
		ctx:        mtq.ctx.Clone(),
		order:      append([]msgtemplate.OrderOption{}, mtq.order...),
		inters:     append([]Interceptor{}, mtq.inters...),
		predicates: append([]predicate.MsgTemplate{}, mtq.predicates...),
		withEvent:  mtq.withEvent.Clone(),
		// clone intermediate query.
		sql:  mtq.sql.Clone(),
		path: mtq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (mtq *MsgTemplateQuery) WithEvent(opts ...func(*MsgEventQuery)) *MsgTemplateQuery {
	query := (&MsgEventClient{config: mtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mtq.withEvent = query
	return mtq
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
//	client.MsgTemplate.Query().
//		GroupBy(msgtemplate.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mtq *MsgTemplateQuery) GroupBy(field string, fields ...string) *MsgTemplateGroupBy {
	mtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MsgTemplateGroupBy{build: mtq}
	grbuild.flds = &mtq.ctx.Fields
	grbuild.label = msgtemplate.Label
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
//	client.MsgTemplate.Query().
//		Select(msgtemplate.FieldCreatedBy).
//		Scan(ctx, &v)
func (mtq *MsgTemplateQuery) Select(fields ...string) *MsgTemplateSelect {
	mtq.ctx.Fields = append(mtq.ctx.Fields, fields...)
	sbuild := &MsgTemplateSelect{MsgTemplateQuery: mtq}
	sbuild.label = msgtemplate.Label
	sbuild.flds, sbuild.scan = &mtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MsgTemplateSelect configured with the given aggregations.
func (mtq *MsgTemplateQuery) Aggregate(fns ...AggregateFunc) *MsgTemplateSelect {
	return mtq.Select().Aggregate(fns...)
}

func (mtq *MsgTemplateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mtq); err != nil {
				return err
			}
		}
	}
	for _, f := range mtq.ctx.Fields {
		if !msgtemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mtq.path != nil {
		prev, err := mtq.path(ctx)
		if err != nil {
			return err
		}
		mtq.sql = prev
	}
	return nil
}

func (mtq *MsgTemplateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MsgTemplate, error) {
	var (
		nodes       = []*MsgTemplate{}
		_spec       = mtq.querySpec()
		loadedTypes = [1]bool{
			mtq.withEvent != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MsgTemplate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MsgTemplate{config: mtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mtq.modifiers) > 0 {
		_spec.Modifiers = mtq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mtq.withEvent; query != nil {
		if err := mtq.loadEvent(ctx, query, nodes, nil,
			func(n *MsgTemplate, e *MsgEvent) { n.Edges.Event = e }); err != nil {
			return nil, err
		}
	}
	for i := range mtq.loadTotal {
		if err := mtq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mtq *MsgTemplateQuery) loadEvent(ctx context.Context, query *MsgEventQuery, nodes []*MsgTemplate, init func(*MsgTemplate), assign func(*MsgTemplate, *MsgEvent)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*MsgTemplate)
	for i := range nodes {
		fk := nodes[i].MsgEventID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(msgevent.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "msg_event_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mtq *MsgTemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mtq.querySpec()
	if len(mtq.modifiers) > 0 {
		_spec.Modifiers = mtq.modifiers
	}
	_spec.Node.Columns = mtq.ctx.Fields
	if len(mtq.ctx.Fields) > 0 {
		_spec.Unique = mtq.ctx.Unique != nil && *mtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mtq.driver, _spec)
}

func (mtq *MsgTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(msgtemplate.Table, msgtemplate.Columns, sqlgraph.NewFieldSpec(msgtemplate.FieldID, field.TypeInt))
	_spec.From = mtq.sql
	if unique := mtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mtq.path != nil {
		_spec.Unique = true
	}
	if fields := mtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, msgtemplate.FieldID)
		for i := range fields {
			if fields[i] != msgtemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mtq.withEvent != nil {
			_spec.Node.AddColumnOnce(msgtemplate.FieldMsgEventID)
		}
	}
	if ps := mtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mtq *MsgTemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mtq.driver.Dialect())
	t1 := builder.Table(msgtemplate.Table)
	columns := mtq.ctx.Fields
	if len(columns) == 0 {
		columns = msgtemplate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mtq.sql != nil {
		selector = mtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mtq.ctx.Unique != nil && *mtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mtq.predicates {
		p(selector)
	}
	for _, p := range mtq.order {
		p(selector)
	}
	if offset := mtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MsgTemplateGroupBy is the group-by builder for MsgTemplate entities.
type MsgTemplateGroupBy struct {
	selector
	build *MsgTemplateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mtgb *MsgTemplateGroupBy) Aggregate(fns ...AggregateFunc) *MsgTemplateGroupBy {
	mtgb.fns = append(mtgb.fns, fns...)
	return mtgb
}

// Scan applies the selector query and scans the result into the given value.
func (mtgb *MsgTemplateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mtgb.build.ctx, "GroupBy")
	if err := mtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MsgTemplateQuery, *MsgTemplateGroupBy](ctx, mtgb.build, mtgb, mtgb.build.inters, v)
}

func (mtgb *MsgTemplateGroupBy) sqlScan(ctx context.Context, root *MsgTemplateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mtgb.fns))
	for _, fn := range mtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mtgb.flds)+len(mtgb.fns))
		for _, f := range *mtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MsgTemplateSelect is the builder for selecting fields of MsgTemplate entities.
type MsgTemplateSelect struct {
	*MsgTemplateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mts *MsgTemplateSelect) Aggregate(fns ...AggregateFunc) *MsgTemplateSelect {
	mts.fns = append(mts.fns, fns...)
	return mts
}

// Scan applies the selector query and scans the result into the given value.
func (mts *MsgTemplateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mts.ctx, "Select")
	if err := mts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MsgTemplateQuery, *MsgTemplateSelect](ctx, mts.MsgTemplateQuery, mts, mts.inters, v)
}

func (mts *MsgTemplateSelect) sqlScan(ctx context.Context, root *MsgTemplateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mts.fns))
	for _, fn := range mts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
