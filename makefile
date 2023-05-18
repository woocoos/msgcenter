.PHONY: maildev
maildev:
	# https://github.com/axllent/mailpit
	/usr/local/opt/mailpit/bin/mailpit
maildev-withauth:
	docker run --rm -p 1081:1080 -p 1026:1025 maildev/maildev:2.0.5 --incoming-user user --incoming-pass pass -v -w 1080

ent-new:
	GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent --target codegen/entgen/schema new $(NAME)

test-db:
	go run test/initdb.go

.PHONY: gen genent gengql genoas
genent:
	go run codegen/entgen/entc.go
genoas:
	go run codegen/oasgen/oasgen.go
gengql:
	go run codegen/gqlgen/gqlgen.go
gengqlall:
	go run github.com/woocoos/entco/cmd/gqltools