// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/msgcenter/ent/msgalert"
	"github.com/woocoos/msgcenter/ent/nlog"
	"github.com/woocoos/msgcenter/ent/nlogalert"
	"github.com/woocoos/msgcenter/ent/predicate"

	"github.com/woocoos/msgcenter/ent/internal"
)

// NlogAlertQuery is the builder for querying NlogAlert entities.
type NlogAlertQuery struct {
	config
	ctx        *QueryContext
	order      []nlogalert.OrderOption
	inters     []Interceptor
	predicates []predicate.NlogAlert
	withNlog   *NlogQuery
	withAlert  *MsgAlertQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*NlogAlert) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NlogAlertQuery builder.
func (naq *NlogAlertQuery) Where(ps ...predicate.NlogAlert) *NlogAlertQuery {
	naq.predicates = append(naq.predicates, ps...)
	return naq
}

// Limit the number of records to be returned by this query.
func (naq *NlogAlertQuery) Limit(limit int) *NlogAlertQuery {
	naq.ctx.Limit = &limit
	return naq
}

// Offset to start from.
func (naq *NlogAlertQuery) Offset(offset int) *NlogAlertQuery {
	naq.ctx.Offset = &offset
	return naq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (naq *NlogAlertQuery) Unique(unique bool) *NlogAlertQuery {
	naq.ctx.Unique = &unique
	return naq
}

// Order specifies how the records should be ordered.
func (naq *NlogAlertQuery) Order(o ...nlogalert.OrderOption) *NlogAlertQuery {
	naq.order = append(naq.order, o...)
	return naq
}

// QueryNlog chains the current query on the "nlog" edge.
func (naq *NlogAlertQuery) QueryNlog() *NlogQuery {
	query := (&NlogClient{config: naq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := naq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := naq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nlogalert.Table, nlogalert.FieldID, selector),
			sqlgraph.To(nlog.Table, nlog.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, nlogalert.NlogTable, nlogalert.NlogColumn),
		)
		schemaConfig := naq.schemaConfig
		step.To.Schema = schemaConfig.Nlog
		step.Edge.Schema = schemaConfig.NlogAlert
		fromU = sqlgraph.SetNeighbors(naq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAlert chains the current query on the "alert" edge.
func (naq *NlogAlertQuery) QueryAlert() *MsgAlertQuery {
	query := (&MsgAlertClient{config: naq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := naq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := naq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nlogalert.Table, nlogalert.FieldID, selector),
			sqlgraph.To(msgalert.Table, msgalert.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, nlogalert.AlertTable, nlogalert.AlertColumn),
		)
		schemaConfig := naq.schemaConfig
		step.To.Schema = schemaConfig.MsgAlert
		step.Edge.Schema = schemaConfig.NlogAlert
		fromU = sqlgraph.SetNeighbors(naq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NlogAlert entity from the query.
// Returns a *NotFoundError when no NlogAlert was found.
func (naq *NlogAlertQuery) First(ctx context.Context) (*NlogAlert, error) {
	nodes, err := naq.Limit(1).All(setContextOp(ctx, naq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nlogalert.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (naq *NlogAlertQuery) FirstX(ctx context.Context) *NlogAlert {
	node, err := naq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NlogAlert ID from the query.
// Returns a *NotFoundError when no NlogAlert ID was found.
func (naq *NlogAlertQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = naq.Limit(1).IDs(setContextOp(ctx, naq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nlogalert.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (naq *NlogAlertQuery) FirstIDX(ctx context.Context) int {
	id, err := naq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NlogAlert entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NlogAlert entity is found.
// Returns a *NotFoundError when no NlogAlert entities are found.
func (naq *NlogAlertQuery) Only(ctx context.Context) (*NlogAlert, error) {
	nodes, err := naq.Limit(2).All(setContextOp(ctx, naq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nlogalert.Label}
	default:
		return nil, &NotSingularError{nlogalert.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (naq *NlogAlertQuery) OnlyX(ctx context.Context) *NlogAlert {
	node, err := naq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NlogAlert ID in the query.
// Returns a *NotSingularError when more than one NlogAlert ID is found.
// Returns a *NotFoundError when no entities are found.
func (naq *NlogAlertQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = naq.Limit(2).IDs(setContextOp(ctx, naq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nlogalert.Label}
	default:
		err = &NotSingularError{nlogalert.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (naq *NlogAlertQuery) OnlyIDX(ctx context.Context) int {
	id, err := naq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NlogAlerts.
func (naq *NlogAlertQuery) All(ctx context.Context) ([]*NlogAlert, error) {
	ctx = setContextOp(ctx, naq.ctx, "All")
	if err := naq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NlogAlert, *NlogAlertQuery]()
	return withInterceptors[[]*NlogAlert](ctx, naq, qr, naq.inters)
}

// AllX is like All, but panics if an error occurs.
func (naq *NlogAlertQuery) AllX(ctx context.Context) []*NlogAlert {
	nodes, err := naq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NlogAlert IDs.
func (naq *NlogAlertQuery) IDs(ctx context.Context) (ids []int, err error) {
	if naq.ctx.Unique == nil && naq.path != nil {
		naq.Unique(true)
	}
	ctx = setContextOp(ctx, naq.ctx, "IDs")
	if err = naq.Select(nlogalert.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (naq *NlogAlertQuery) IDsX(ctx context.Context) []int {
	ids, err := naq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (naq *NlogAlertQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, naq.ctx, "Count")
	if err := naq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, naq, querierCount[*NlogAlertQuery](), naq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (naq *NlogAlertQuery) CountX(ctx context.Context) int {
	count, err := naq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (naq *NlogAlertQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, naq.ctx, "Exist")
	switch _, err := naq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (naq *NlogAlertQuery) ExistX(ctx context.Context) bool {
	exist, err := naq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NlogAlertQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (naq *NlogAlertQuery) Clone() *NlogAlertQuery {
	if naq == nil {
		return nil
	}
	return &NlogAlertQuery{
		config:     naq.config,
		ctx:        naq.ctx.Clone(),
		order:      append([]nlogalert.OrderOption{}, naq.order...),
		inters:     append([]Interceptor{}, naq.inters...),
		predicates: append([]predicate.NlogAlert{}, naq.predicates...),
		withNlog:   naq.withNlog.Clone(),
		withAlert:  naq.withAlert.Clone(),
		// clone intermediate query.
		sql:  naq.sql.Clone(),
		path: naq.path,
	}
}

// WithNlog tells the query-builder to eager-load the nodes that are connected to
// the "nlog" edge. The optional arguments are used to configure the query builder of the edge.
func (naq *NlogAlertQuery) WithNlog(opts ...func(*NlogQuery)) *NlogAlertQuery {
	query := (&NlogClient{config: naq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	naq.withNlog = query
	return naq
}

// WithAlert tells the query-builder to eager-load the nodes that are connected to
// the "alert" edge. The optional arguments are used to configure the query builder of the edge.
func (naq *NlogAlertQuery) WithAlert(opts ...func(*MsgAlertQuery)) *NlogAlertQuery {
	query := (&MsgAlertClient{config: naq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	naq.withAlert = query
	return naq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NlogID int `json:"nlog_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NlogAlert.Query().
//		GroupBy(nlogalert.FieldNlogID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (naq *NlogAlertQuery) GroupBy(field string, fields ...string) *NlogAlertGroupBy {
	naq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NlogAlertGroupBy{build: naq}
	grbuild.flds = &naq.ctx.Fields
	grbuild.label = nlogalert.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		NlogID int `json:"nlog_id,omitempty"`
//	}
//
//	client.NlogAlert.Query().
//		Select(nlogalert.FieldNlogID).
//		Scan(ctx, &v)
func (naq *NlogAlertQuery) Select(fields ...string) *NlogAlertSelect {
	naq.ctx.Fields = append(naq.ctx.Fields, fields...)
	sbuild := &NlogAlertSelect{NlogAlertQuery: naq}
	sbuild.label = nlogalert.Label
	sbuild.flds, sbuild.scan = &naq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NlogAlertSelect configured with the given aggregations.
func (naq *NlogAlertQuery) Aggregate(fns ...AggregateFunc) *NlogAlertSelect {
	return naq.Select().Aggregate(fns...)
}

func (naq *NlogAlertQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range naq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, naq); err != nil {
				return err
			}
		}
	}
	for _, f := range naq.ctx.Fields {
		if !nlogalert.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if naq.path != nil {
		prev, err := naq.path(ctx)
		if err != nil {
			return err
		}
		naq.sql = prev
	}
	return nil
}

func (naq *NlogAlertQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NlogAlert, error) {
	var (
		nodes       = []*NlogAlert{}
		_spec       = naq.querySpec()
		loadedTypes = [2]bool{
			naq.withNlog != nil,
			naq.withAlert != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NlogAlert).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NlogAlert{config: naq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = naq.schemaConfig.NlogAlert
	ctx = internal.NewSchemaConfigContext(ctx, naq.schemaConfig)
	if len(naq.modifiers) > 0 {
		_spec.Modifiers = naq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, naq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := naq.withNlog; query != nil {
		if err := naq.loadNlog(ctx, query, nodes, nil,
			func(n *NlogAlert, e *Nlog) { n.Edges.Nlog = e }); err != nil {
			return nil, err
		}
	}
	if query := naq.withAlert; query != nil {
		if err := naq.loadAlert(ctx, query, nodes, nil,
			func(n *NlogAlert, e *MsgAlert) { n.Edges.Alert = e }); err != nil {
			return nil, err
		}
	}
	for i := range naq.loadTotal {
		if err := naq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (naq *NlogAlertQuery) loadNlog(ctx context.Context, query *NlogQuery, nodes []*NlogAlert, init func(*NlogAlert), assign func(*NlogAlert, *Nlog)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*NlogAlert)
	for i := range nodes {
		fk := nodes[i].NlogID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(nlog.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "nlog_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (naq *NlogAlertQuery) loadAlert(ctx context.Context, query *MsgAlertQuery, nodes []*NlogAlert, init func(*NlogAlert), assign func(*NlogAlert, *MsgAlert)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*NlogAlert)
	for i := range nodes {
		fk := nodes[i].AlertID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(msgalert.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "alert_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (naq *NlogAlertQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := naq.querySpec()
	_spec.Node.Schema = naq.schemaConfig.NlogAlert
	ctx = internal.NewSchemaConfigContext(ctx, naq.schemaConfig)
	if len(naq.modifiers) > 0 {
		_spec.Modifiers = naq.modifiers
	}
	_spec.Node.Columns = naq.ctx.Fields
	if len(naq.ctx.Fields) > 0 {
		_spec.Unique = naq.ctx.Unique != nil && *naq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, naq.driver, _spec)
}

func (naq *NlogAlertQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(nlogalert.Table, nlogalert.Columns, sqlgraph.NewFieldSpec(nlogalert.FieldID, field.TypeInt))
	_spec.From = naq.sql
	if unique := naq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if naq.path != nil {
		_spec.Unique = true
	}
	if fields := naq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nlogalert.FieldID)
		for i := range fields {
			if fields[i] != nlogalert.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if naq.withNlog != nil {
			_spec.Node.AddColumnOnce(nlogalert.FieldNlogID)
		}
		if naq.withAlert != nil {
			_spec.Node.AddColumnOnce(nlogalert.FieldAlertID)
		}
	}
	if ps := naq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := naq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := naq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := naq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (naq *NlogAlertQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(naq.driver.Dialect())
	t1 := builder.Table(nlogalert.Table)
	columns := naq.ctx.Fields
	if len(columns) == 0 {
		columns = nlogalert.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if naq.sql != nil {
		selector = naq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if naq.ctx.Unique != nil && *naq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(naq.schemaConfig.NlogAlert)
	ctx = internal.NewSchemaConfigContext(ctx, naq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range naq.predicates {
		p(selector)
	}
	for _, p := range naq.order {
		p(selector)
	}
	if offset := naq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := naq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NlogAlertGroupBy is the group-by builder for NlogAlert entities.
type NlogAlertGroupBy struct {
	selector
	build *NlogAlertQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nagb *NlogAlertGroupBy) Aggregate(fns ...AggregateFunc) *NlogAlertGroupBy {
	nagb.fns = append(nagb.fns, fns...)
	return nagb
}

// Scan applies the selector query and scans the result into the given value.
func (nagb *NlogAlertGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nagb.build.ctx, "GroupBy")
	if err := nagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NlogAlertQuery, *NlogAlertGroupBy](ctx, nagb.build, nagb, nagb.build.inters, v)
}

func (nagb *NlogAlertGroupBy) sqlScan(ctx context.Context, root *NlogAlertQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(nagb.fns))
	for _, fn := range nagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*nagb.flds)+len(nagb.fns))
		for _, f := range *nagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*nagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NlogAlertSelect is the builder for selecting fields of NlogAlert entities.
type NlogAlertSelect struct {
	*NlogAlertQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nas *NlogAlertSelect) Aggregate(fns ...AggregateFunc) *NlogAlertSelect {
	nas.fns = append(nas.fns, fns...)
	return nas
}

// Scan applies the selector query and scans the result into the given value.
func (nas *NlogAlertSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nas.ctx, "Select")
	if err := nas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NlogAlertQuery, *NlogAlertSelect](ctx, nas.NlogAlertQuery, nas, nas.inters, v)
}

func (nas *NlogAlertSelect) sqlScan(ctx context.Context, root *NlogAlertQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nas.fns))
	for _, fn := range nas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
