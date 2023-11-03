package msg

import (
	"github.com/99designs/gqlgen/graphql"
	apigraphql "github.com/woocoos/msgcenter/api/graphql"
	"github.com/woocoos/msgcenter/api/graphql/generated"
)

type Resolver struct {
	*apigraphql.Resolver
}

func NewResolver(opt ...apigraphql.Option) *Resolver {
	r := &Resolver{
		Resolver: apigraphql.NewResolver(opt...),
	}
	return r
}

// NewSchema creates a graphql executable schema.
func NewSchema(opts ...apigraphql.Option) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: NewResolver(opts...),
	})
}
