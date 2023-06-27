//go:build ignore

package main

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"flag"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/migrate"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// receive two arguments: the migration name and the database dsn.
var (
	dsn  = flag.String("dsn", "root:@tcp(localhost:3306)/msgcenter", "")
	name = flag.String("name", "mysql", "driver name")
)

func main() {
	client, err := ent.Open(*name, *dsn)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
		schema.WithHooks(func(creator schema.Creator) schema.Creator {
			return schema.CreateFunc(func(ctx context.Context, table ...*schema.Table) error {
				var rt []*schema.Table
				for _, t := range table {
					if t.Name == "user" || t.Name == "org_role_user" {
					} else {
						rt = append(rt, t)
					}
				}
				return creator.Create(ctx, rt...)
			})
		}))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
