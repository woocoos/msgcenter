package ams

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/web"
	"path/filepath"
)

func BuildUiServer(cnf *conf.Configuration) *web.Server {
	staticDir := cnf.String("staticDir")
	uiSrv := web.New(web.WithConfiguration(cnf),
		web.WithGracefulStop(),
	)
	uiSrv.Router().NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(staticDir, c.Request.URL.Path+".html"))
	})
	uiSrv.Router().Use(static.Serve("/", static.LocalFile(staticDir, false)))
	return uiSrv
}

func NewPromHttp(cnf *conf.Configuration) woocoo.Server {
	hd := web.New(web.WithConfiguration(cnf))
	hd.Router().GET("/metrics", gin.WrapH(promhttp.Handler()))
	return hd
}
