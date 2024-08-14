package graphql

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/redis/go-redis/v9"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/msgcenter/api/graphql/generated"
	"github.com/woocoos/msgcenter/api/graphql/model"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/service"
	"github.com/woocoos/msgcenter/service/silence"
)

var (
	wsFilterCtxKey = "ko_msg_filter"
)

type Option func(*Resolver)

func WithCoordinator(coordinator *service.Coordinator) Option {
	return func(r *Resolver) {
		r.coordinator = coordinator
	}
}

func WithClient(client *ent.Client) Option {
	return func(r *Resolver) {
		r.client = client
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

func WithKOClient(client *api.SDK) Option {
	return func(r *Resolver) {
		r.kosdk = client
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
	coordinator *service.Coordinator
	client      *ent.Client
	Silences    *silence.Silences
	MsgClient   redis.UniversalClient
	PubSub      PubSub
	// knockout sdk
	kosdk *api.SDK
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

// RemoveDuplicateElement 去重
func RemoveDuplicateElement[T int | int64 | string | float32 | float64](arr []T) []T {
	if arr == nil {
		return nil
	}
	temp := make(map[string]bool)
	result := make([]T, 0, len(arr))
	for _, v := range arr {
		key := fmt.Sprint(v)
		if _, ok := temp[key]; !ok {
			temp[key] = true
			result = append(result, v)
		}
	}
	return result
}
