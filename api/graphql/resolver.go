package graphql

import (
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Coordinator *service.Coordinator
	Client      *ent.Client
}
