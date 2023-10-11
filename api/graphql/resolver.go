package graphql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/redis/go-redis/v9"
	"github.com/woocoos/msgcenter/api/graphql/generated"
	"github.com/woocoos/msgcenter/api/graphql/model"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/service"
	"github.com/woocoos/msgcenter/silence"
)

var (
	wsFilterCtxKey = "ko_msg_filter"
)

type Option func(*Resolver)

func WithCoordinator(coordinator *service.Coordinator) Option {
	return func(r *Resolver) {
		r.Coordinator = coordinator
	}
}

func WithClient(client *ent.Client) Option {
	return func(r *Resolver) {
		r.Client = client
	}
}

func WithSilences(silences *silence.Silences) Option {
	return func(r *Resolver) {
		r.Silences = silences
	}
}

func WithMsgClient(client redis.UniversalClient) Option {
	return func(r *Resolver) {
		r.MsgClient = client
	}
}

func WithPubSub(pubSub PubSub) Option {
	return func(r *Resolver) {
		r.PubSub = pubSub
	}
}

type PubSub interface {
	Subscribe(ctx context.Context, topic string) (chan *model.Message, error)
}

// Resolver is the root resolver.
type Resolver struct {
	Coordinator *service.Coordinator
	Client      *ent.Client
	Silences    *silence.Silences
	MsgClient   redis.UniversalClient
	PubSub      PubSub
}

func NewResolver(opt ...Option) *Resolver {
	r := &Resolver{}
	for _, option := range opt {
		option(r)
	}
	return r
}

// NewSchema creates a graphql executable schema.
func NewSchema(opts ...Option) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: NewResolver(opts...),
	})
}
