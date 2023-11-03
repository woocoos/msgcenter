// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"github.com/woocoos/entcache"
	"github.com/woocoos/msgcenter/ent/msgalert"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msginternal"
	"github.com/woocoos/msgcenter/ent/msginternalto"
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/nlog"
	"github.com/woocoos/msgcenter/ent/nlogalert"
	"github.com/woocoos/msgcenter/ent/silence"
	"github.com/woocoos/msgcenter/ent/user"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

var msgalertImplementors = []string{"MsgAlert", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*MsgAlert) IsNode() {}

var msgchannelImplementors = []string{"MsgChannel", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*MsgChannel) IsNode() {}

var msgeventImplementors = []string{"MsgEvent", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*MsgEvent) IsNode() {}

var msginternalImplementors = []string{"MsgInternal", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*MsgInternal) IsNode() {}

var msginternaltoImplementors = []string{"MsgInternalTo", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*MsgInternalTo) IsNode() {}

var msgsubscriberImplementors = []string{"MsgSubscriber", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*MsgSubscriber) IsNode() {}

var msgtemplateImplementors = []string{"MsgTemplate", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*MsgTemplate) IsNode() {}

var msgtypeImplementors = []string{"MsgType", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*MsgType) IsNode() {}

var nlogImplementors = []string{"Nlog", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Nlog) IsNode() {}

var nlogalertImplementors = []string{"NlogAlert", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*NlogAlert) IsNode() {}

var silenceImplementors = []string{"Silence", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Silence) IsNode() {}

var userImplementors = []string{"User", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*User) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case msgalert.Table:
		query := c.MsgAlert.Query().
			Where(msgalert.ID(id))
		query, err := query.CollectFields(ctx, msgalertImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "MsgAlert", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case msgchannel.Table:
		query := c.MsgChannel.Query().
			Where(msgchannel.ID(id))
		query, err := query.CollectFields(ctx, msgchannelImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "MsgChannel", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case msgevent.Table:
		query := c.MsgEvent.Query().
			Where(msgevent.ID(id))
		query, err := query.CollectFields(ctx, msgeventImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "MsgEvent", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case msginternal.Table:
		query := c.MsgInternal.Query().
			Where(msginternal.ID(id))
		query, err := query.CollectFields(ctx, msginternalImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "MsgInternal", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case msginternalto.Table:
		query := c.MsgInternalTo.Query().
			Where(msginternalto.ID(id))
		query, err := query.CollectFields(ctx, msginternaltoImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "MsgInternalTo", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case msgsubscriber.Table:
		query := c.MsgSubscriber.Query().
			Where(msgsubscriber.ID(id))
		query, err := query.CollectFields(ctx, msgsubscriberImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "MsgSubscriber", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case msgtemplate.Table:
		query := c.MsgTemplate.Query().
			Where(msgtemplate.ID(id))
		query, err := query.CollectFields(ctx, msgtemplateImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "MsgTemplate", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case msgtype.Table:
		query := c.MsgType.Query().
			Where(msgtype.ID(id))
		query, err := query.CollectFields(ctx, msgtypeImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "MsgType", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case nlog.Table:
		query := c.Nlog.Query().
			Where(nlog.ID(id))
		query, err := query.CollectFields(ctx, nlogImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "Nlog", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case nlogalert.Table:
		query := c.NlogAlert.Query().
			Where(nlogalert.ID(id))
		query, err := query.CollectFields(ctx, nlogalertImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "NlogAlert", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case silence.Table:
		query := c.Silence.Query().
			Where(silence.ID(id))
		query, err := query.CollectFields(ctx, silenceImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "Silence", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		query, err := query.CollectFields(ctx, userImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "User", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case msgalert.Table:
		query := c.MsgAlert.Query().
			Where(msgalert.IDIn(ids...))
		query, err := query.CollectFields(ctx, msgalertImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case msgchannel.Table:
		query := c.MsgChannel.Query().
			Where(msgchannel.IDIn(ids...))
		query, err := query.CollectFields(ctx, msgchannelImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case msgevent.Table:
		query := c.MsgEvent.Query().
			Where(msgevent.IDIn(ids...))
		query, err := query.CollectFields(ctx, msgeventImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case msginternal.Table:
		query := c.MsgInternal.Query().
			Where(msginternal.IDIn(ids...))
		query, err := query.CollectFields(ctx, msginternalImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case msginternalto.Table:
		query := c.MsgInternalTo.Query().
			Where(msginternalto.IDIn(ids...))
		query, err := query.CollectFields(ctx, msginternaltoImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case msgsubscriber.Table:
		query := c.MsgSubscriber.Query().
			Where(msgsubscriber.IDIn(ids...))
		query, err := query.CollectFields(ctx, msgsubscriberImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case msgtemplate.Table:
		query := c.MsgTemplate.Query().
			Where(msgtemplate.IDIn(ids...))
		query, err := query.CollectFields(ctx, msgtemplateImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case msgtype.Table:
		query := c.MsgType.Query().
			Where(msgtype.IDIn(ids...))
		query, err := query.CollectFields(ctx, msgtypeImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case nlog.Table:
		query := c.Nlog.Query().
			Where(nlog.IDIn(ids...))
		query, err := query.CollectFields(ctx, nlogImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case nlogalert.Table:
		query := c.NlogAlert.Query().
			Where(nlogalert.IDIn(ids...))
		query, err := query.CollectFields(ctx, nlogalertImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case silence.Table:
		query := c.Silence.Query().
			Where(silence.IDIn(ids...))
		query, err := query.CollectFields(ctx, silenceImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, userImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
