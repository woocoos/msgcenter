//go:build ignore

package main

import (
	atlas "ariga.io/atlas/sql/schema"
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
		// Hook into Atlas Diff process.
		schema.WithDiffHook(func(next schema.Differ) schema.Differ {
			return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
				var dt []*atlas.Table
				for i, table := range desired.Tables {
					if !(table.Name == "user" || table.Name == "org_role_user") {
						dt = append(dt, desired.Tables[i])
					}
				}
				desired.Tables = dt
				// Before calculating changes.
				changes, err := next.Diff(current, desired)
				if err != nil {
					return nil, err
				}
				// After diff, you can filter
				// changes or return new ones.
				return changes, nil
			})
		}),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
