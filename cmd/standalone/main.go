package main

import (
	"flag"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/cmd/internal/ams"
)

var (
	amsConfig = flag.String("r", "../ams", "rms etc dir")
)

func main() {
	flag.Parse()
	app := woocoo.New()

	rmsc := conf.New(conf.WithBaseDir(*amsConfig), conf.WithGlobal(false)).Load()
	rmscnf := &conf.AppConfiguration{Configuration: rmsc}

	rmsSvr := ams.NewServer(rmscnf)

	app.RegisterServer(rmsSvr.Servers()...)

	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
