// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/predicate"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The MsgChannelFunc type is an adapter to allow the use of ordinary function as a Querier.
type MsgChannelFunc func(context.Context, *ent.MsgChannelQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MsgChannelFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MsgChannelQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MsgChannelQuery", q)
}

// The TraverseMsgChannel type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMsgChannel func(context.Context, *ent.MsgChannelQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMsgChannel) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMsgChannel) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MsgChannelQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MsgChannelQuery", q)
}

// The MsgEventFunc type is an adapter to allow the use of ordinary function as a Querier.
type MsgEventFunc func(context.Context, *ent.MsgEventQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MsgEventFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MsgEventQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MsgEventQuery", q)
}

// The TraverseMsgEvent type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMsgEvent func(context.Context, *ent.MsgEventQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMsgEvent) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMsgEvent) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MsgEventQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MsgEventQuery", q)
}

// The MsgSubscriberFunc type is an adapter to allow the use of ordinary function as a Querier.
type MsgSubscriberFunc func(context.Context, *ent.MsgSubscriberQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MsgSubscriberFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MsgSubscriberQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MsgSubscriberQuery", q)
}

// The TraverseMsgSubscriber type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMsgSubscriber func(context.Context, *ent.MsgSubscriberQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMsgSubscriber) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMsgSubscriber) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MsgSubscriberQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MsgSubscriberQuery", q)
}

// The MsgTemplateFunc type is an adapter to allow the use of ordinary function as a Querier.
type MsgTemplateFunc func(context.Context, *ent.MsgTemplateQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MsgTemplateFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MsgTemplateQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MsgTemplateQuery", q)
}

// The TraverseMsgTemplate type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMsgTemplate func(context.Context, *ent.MsgTemplateQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMsgTemplate) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMsgTemplate) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MsgTemplateQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MsgTemplateQuery", q)
}

// The MsgTypeFunc type is an adapter to allow the use of ordinary function as a Querier.
type MsgTypeFunc func(context.Context, *ent.MsgTypeQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MsgTypeFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MsgTypeQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MsgTypeQuery", q)
}

// The TraverseMsgType type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMsgType func(context.Context, *ent.MsgTypeQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMsgType) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMsgType) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MsgTypeQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MsgTypeQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.MsgChannelQuery:
		return &query[*ent.MsgChannelQuery, predicate.MsgChannel, msgchannel.OrderOption]{typ: ent.TypeMsgChannel, tq: q}, nil
	case *ent.MsgEventQuery:
		return &query[*ent.MsgEventQuery, predicate.MsgEvent, msgevent.OrderOption]{typ: ent.TypeMsgEvent, tq: q}, nil
	case *ent.MsgSubscriberQuery:
		return &query[*ent.MsgSubscriberQuery, predicate.MsgSubscriber, msgsubscriber.OrderOption]{typ: ent.TypeMsgSubscriber, tq: q}, nil
	case *ent.MsgTemplateQuery:
		return &query[*ent.MsgTemplateQuery, predicate.MsgTemplate, msgtemplate.OrderOption]{typ: ent.TypeMsgTemplate, tq: q}, nil
	case *ent.MsgTypeQuery:
		return &query[*ent.MsgTypeQuery, predicate.MsgType, msgtype.OrderOption]{typ: ent.TypeMsgType, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
