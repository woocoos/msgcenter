// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
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

// NlogQuery is the builder for querying Nlog entities.
type NlogQuery struct {
	config
	ctx                *QueryContext
	order              []nlog.OrderOption
	inters             []Interceptor
	predicates         []predicate.Nlog
	withAlerts         *MsgAlertQuery
	withNlogAlert      *NlogAlertQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*Nlog) error
	withNamedAlerts    map[string]*MsgAlertQuery
	withNamedNlogAlert map[string]*NlogAlertQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NlogQuery builder.
func (nq *NlogQuery) Where(ps ...predicate.Nlog) *NlogQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NlogQuery) Limit(limit int) *NlogQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NlogQuery) Offset(offset int) *NlogQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NlogQuery) Unique(unique bool) *NlogQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NlogQuery) Order(o ...nlog.OrderOption) *NlogQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryAlerts chains the current query on the "alerts" edge.
func (nq *NlogQuery) QueryAlerts() *MsgAlertQuery {
	query := (&MsgAlertClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nlog.Table, nlog.FieldID, selector),
			sqlgraph.To(msgalert.Table, msgalert.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, nlog.AlertsTable, nlog.AlertsPrimaryKey...),
		)
		schemaConfig := nq.schemaConfig
		step.To.Schema = schemaConfig.MsgAlert
		step.Edge.Schema = schemaConfig.NlogAlert
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNlogAlert chains the current query on the "nlog_alert" edge.
func (nq *NlogQuery) QueryNlogAlert() *NlogAlertQuery {
	query := (&NlogAlertClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nlog.Table, nlog.FieldID, selector),
			sqlgraph.To(nlogalert.Table, nlogalert.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, nlog.NlogAlertTable, nlog.NlogAlertColumn),
		)
		schemaConfig := nq.schemaConfig
		step.To.Schema = schemaConfig.NlogAlert
		step.Edge.Schema = schemaConfig.NlogAlert
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Nlog entity from the query.
// Returns a *NotFoundError when no Nlog was found.
func (nq *NlogQuery) First(ctx context.Context) (*Nlog, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NlogQuery) FirstX(ctx context.Context) *Nlog {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Nlog ID from the query.
// Returns a *NotFoundError when no Nlog ID was found.
func (nq *NlogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NlogQuery) FirstIDX(ctx context.Context) int {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Nlog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Nlog entity is found.
// Returns a *NotFoundError when no Nlog entities are found.
func (nq *NlogQuery) Only(ctx context.Context) (*Nlog, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nlog.Label}
	default:
		return nil, &NotSingularError{nlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NlogQuery) OnlyX(ctx context.Context) *Nlog {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Nlog ID in the query.
// Returns a *NotSingularError when more than one Nlog ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NlogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nlog.Label}
	default:
		err = &NotSingularError{nlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NlogQuery) OnlyIDX(ctx context.Context) int {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Nlogs.
func (nq *NlogQuery) All(ctx context.Context) ([]*Nlog, error) {
	ctx = setContextOp(ctx, nq.ctx, "All")
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Nlog, *NlogQuery]()
	return withInterceptors[[]*Nlog](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NlogQuery) AllX(ctx context.Context) []*Nlog {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Nlog IDs.
func (nq *NlogQuery) IDs(ctx context.Context) (ids []int, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, "IDs")
	if err = nq.Select(nlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NlogQuery) IDsX(ctx context.Context) []int {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NlogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, "Count")
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NlogQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NlogQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NlogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, "Exist")
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NlogQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NlogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NlogQuery) Clone() *NlogQuery {
	if nq == nil {
		return nil
	}
	return &NlogQuery{
		config:        nq.config,
		ctx:           nq.ctx.Clone(),
		order:         append([]nlog.OrderOption{}, nq.order...),
		inters:        append([]Interceptor{}, nq.inters...),
		predicates:    append([]predicate.Nlog{}, nq.predicates...),
		withAlerts:    nq.withAlerts.Clone(),
		withNlogAlert: nq.withNlogAlert.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithAlerts tells the query-builder to eager-load the nodes that are connected to
// the "alerts" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NlogQuery) WithAlerts(opts ...func(*MsgAlertQuery)) *NlogQuery {
	query := (&MsgAlertClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withAlerts = query
	return nq
}

// WithNlogAlert tells the query-builder to eager-load the nodes that are connected to
// the "nlog_alert" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NlogQuery) WithNlogAlert(opts ...func(*NlogAlertQuery)) *NlogQuery {
	query := (&NlogAlertClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withNlogAlert = query
	return nq
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
//	client.Nlog.Query().
//		GroupBy(nlog.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NlogQuery) GroupBy(field string, fields ...string) *NlogGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NlogGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = nlog.Label
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
//	client.Nlog.Query().
//		Select(nlog.FieldTenantID).
//		Scan(ctx, &v)
func (nq *NlogQuery) Select(fields ...string) *NlogSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NlogSelect{NlogQuery: nq}
	sbuild.label = nlog.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NlogSelect configured with the given aggregations.
func (nq *NlogQuery) Aggregate(fns ...AggregateFunc) *NlogSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NlogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !nlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NlogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Nlog, error) {
	var (
		nodes       = []*Nlog{}
		_spec       = nq.querySpec()
		loadedTypes = [2]bool{
			nq.withAlerts != nil,
			nq.withNlogAlert != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Nlog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Nlog{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = nq.schemaConfig.Nlog
	ctx = internal.NewSchemaConfigContext(ctx, nq.schemaConfig)
	if len(nq.modifiers) > 0 {
		_spec.Modifiers = nq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withAlerts; query != nil {
		if err := nq.loadAlerts(ctx, query, nodes,
			func(n *Nlog) { n.Edges.Alerts = []*MsgAlert{} },
			func(n *Nlog, e *MsgAlert) { n.Edges.Alerts = append(n.Edges.Alerts, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withNlogAlert; query != nil {
		if err := nq.loadNlogAlert(ctx, query, nodes,
			func(n *Nlog) { n.Edges.NlogAlert = []*NlogAlert{} },
			func(n *Nlog, e *NlogAlert) { n.Edges.NlogAlert = append(n.Edges.NlogAlert, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range nq.withNamedAlerts {
		if err := nq.loadAlerts(ctx, query, nodes,
			func(n *Nlog) { n.appendNamedAlerts(name) },
			func(n *Nlog, e *MsgAlert) { n.appendNamedAlerts(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range nq.withNamedNlogAlert {
		if err := nq.loadNlogAlert(ctx, query, nodes,
			func(n *Nlog) { n.appendNamedNlogAlert(name) },
			func(n *Nlog, e *NlogAlert) { n.appendNamedNlogAlert(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range nq.loadTotal {
		if err := nq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NlogQuery) loadAlerts(ctx context.Context, query *MsgAlertQuery, nodes []*Nlog, init func(*Nlog), assign func(*Nlog, *MsgAlert)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Nlog)
	nids := make(map[int]map[*Nlog]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(nlog.AlertsTable)
		joinT.Schema(nq.schemaConfig.NlogAlert)
		s.Join(joinT).On(s.C(msgalert.FieldID), joinT.C(nlog.AlertsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(nlog.AlertsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(nlog.AlertsPrimaryKey[0]))
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
					nids[inValue] = map[*Nlog]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*MsgAlert](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "alerts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (nq *NlogQuery) loadNlogAlert(ctx context.Context, query *NlogAlertQuery, nodes []*Nlog, init func(*Nlog), assign func(*Nlog, *NlogAlert)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Nlog)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(nlogalert.FieldNlogID)
	}
	query.Where(predicate.NlogAlert(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(nlog.NlogAlertColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.NlogID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "nlog_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (nq *NlogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Schema = nq.schemaConfig.Nlog
	ctx = internal.NewSchemaConfigContext(ctx, nq.schemaConfig)
	if len(nq.modifiers) > 0 {
		_spec.Modifiers = nq.modifiers
	}
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NlogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(nlog.Table, nlog.Columns, sqlgraph.NewFieldSpec(nlog.FieldID, field.TypeInt))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nlog.FieldID)
		for i := range fields {
			if fields[i] != nlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NlogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(nlog.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = nlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(nq.schemaConfig.Nlog)
	ctx = internal.NewSchemaConfigContext(ctx, nq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAlerts tells the query-builder to eager-load the nodes that are connected to the "alerts"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (nq *NlogQuery) WithNamedAlerts(name string, opts ...func(*MsgAlertQuery)) *NlogQuery {
	query := (&MsgAlertClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if nq.withNamedAlerts == nil {
		nq.withNamedAlerts = make(map[string]*MsgAlertQuery)
	}
	nq.withNamedAlerts[name] = query
	return nq
}

// WithNamedNlogAlert tells the query-builder to eager-load the nodes that are connected to the "nlog_alert"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (nq *NlogQuery) WithNamedNlogAlert(name string, opts ...func(*NlogAlertQuery)) *NlogQuery {
	query := (&NlogAlertClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if nq.withNamedNlogAlert == nil {
		nq.withNamedNlogAlert = make(map[string]*NlogAlertQuery)
	}
	nq.withNamedNlogAlert[name] = query
	return nq
}

// NlogGroupBy is the group-by builder for Nlog entities.
type NlogGroupBy struct {
	selector
	build *NlogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NlogGroupBy) Aggregate(fns ...AggregateFunc) *NlogGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NlogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, "GroupBy")
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NlogQuery, *NlogGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NlogGroupBy) sqlScan(ctx context.Context, root *NlogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NlogSelect is the builder for selecting fields of Nlog entities.
type NlogSelect struct {
	*NlogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NlogSelect) Aggregate(fns ...AggregateFunc) *NlogSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NlogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, "Select")
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NlogQuery, *NlogSelect](ctx, ns.NlogQuery, ns, ns.inters, v)
}

func (ns *NlogSelect) sqlScan(ctx context.Context, root *NlogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
