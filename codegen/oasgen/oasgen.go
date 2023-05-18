package main

import (
	"github.com/tsingsun/woocoo/cmd/woco/oasgen"
	"github.com/tsingsun/woocoo/cmd/woco/oasgen/codegen"
	"log"
)

func main() {
	opts := []oasgen.Option{}
	cfg := &codegen.Config{
		OpenAPISchema: "./api/oas/openapi.yaml",
	}
	err := oasgen.LoadConfig(cfg, "./codegen/oasgen/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = oasgen.Generate(cfg.OpenAPISchema, cfg, opts...)
	if err != nil {
		log.Fatal(err)
	}
}
