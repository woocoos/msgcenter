.PHONY: maildev
maildev:
	# https://github.com/axllent/mailpit UI: http://0.0.0.0:8025
	/usr/local/opt/mailpit/bin/mailpit

ent-new:
	GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent --target codegen/entgen/schema new $(NAME)

test-db:
	go run script/initdb.go

.PHONY: gen genent gengql genoas
genent:
	go run codegen/entgen/entc.go
genoas:
	# go install github.com/tsingsun/woocoo/cmd/woco@main
	woco oasgen -c ./codegen/oasgen/config.yaml
gengql:
	go run codegen/gqlgen/gqlgen.go