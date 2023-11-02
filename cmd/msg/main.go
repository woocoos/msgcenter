package main

import (
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/msgcenter/cmd/internal/msg"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

func main() {
	app := woocoo.New()
	msgSvr := msg.NewServer(app.AppConfiguration())

	app.RegisterServer(msgSvr, clientx.ChangeSet)

	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
