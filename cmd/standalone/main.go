package main

import (
	"flag"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	ecx "github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/msgcenter/cmd/internal/ams"
	"github.com/woocoos/msgcenter/cmd/internal/msg"
)

var (
	amsConfig = flag.String("r", "../ams", "rms etc dir")
	msgConfig = flag.String("m", "../msg", "msg etc dir")
)

func main() {
	flag.Parse()
	app := woocoo.New()

	rmsc := conf.New(conf.WithBaseDir(*amsConfig), conf.WithGlobal(false)).Load()
	rmscnf := &conf.AppConfiguration{Configuration: rmsc}

	rmsSvr := ams.NewServer(rmscnf)
	app.RegisterServer(rmsSvr.Servers()...)

	msgc := conf.New(conf.WithBaseDir(*msgConfig), conf.WithGlobal(false)).Load()
	msgcnf := &conf.AppConfiguration{Configuration: msgc}
	msgSvr := msg.NewServer(msgcnf)
	app.RegisterServer(msgSvr)

	app.RegisterServer(ecx.ChangeSet)
	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
