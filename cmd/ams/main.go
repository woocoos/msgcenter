package main

import (
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/cmd/internal/ams"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

func main() {
	app := woocoo.New()
	rmsSvr := ams.NewServer(app.AppConfiguration())

	app.RegisterServer(rmsSvr.Servers()...)

	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
