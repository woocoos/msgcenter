$version = Get-Date -Format "yyyy-MM-dd HH:mm"
$BUILD_NAME = "standalone"

function build {
    go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$version'" -o "./cmd/$BUILD_NAME/$BUILD_NAME" "./cmd/$BUILD_NAME/main.go"
}

function mac {
    $env:GOOS = "darwin"
    go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$version'" -o "./cmd/$BUILD_NAME-darwin" "./cmd/$BUILD_NAME/main.go"
    if (Get-Command upx -ErrorAction SilentlyContinue) {
        upx "$BUILD_NAME-darwin"
    }
}

function win {
    $env:GOOS = "windows"
    go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$version'" -o "./cmd/$BUILD_NAME/$BUILD_NAME.exe" "./cmd/$BUILD_NAME/main.go"
    if (Get-Command upx -ErrorAction SilentlyContinue) {
        upx "$BUILD_NAME.exe"
    }
}

function linux {
    $env:GOOS = "linux"
    go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$version'" -o "./cmd/$BUILD_NAME/$BUILD_NAME-linux" "./cmd/$BUILD_NAME/main.go"
    if (Get-Command upx -ErrorAction SilentlyContinue) {
        upx "$BUILD_NAME-linux"
    }
}


function new-ent {
    $env:GOWORK = "off"
    go run -mod=mod entgo.io/ent/cmd/ent --target codegen/entgen/schema new $env:NAME
}

function migrate-init {
    $env:GOWORK = "off"
    go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration ./ent/schema
}

function migrate {
    $env:GOWORK = "off"
    go run -mod=mod ent/migrate/main.go -dsn="$env:DSN" -name=$env:NAME
}

function migrate-lint {
    atlas migrate lint --dev-url="$env:DSN" --dir="file://ent/migrate/migrations" --latest=$env:LATEST
}

function migrate-apply {
    atlas migrate apply --dev-url="$env:DSN" --dir="file://ent/migrate/migrations" --latest=$env:LATEST
}

function db {
    initialize-schema
    initialize-basedata
    initialize-qql-action
    initialize-qql-menu
    initialize-app-policy
}

function initialize-schema {
    $env:GOWORK = "off"
    go run -mod=mod script/initschema.go
}

function initialize-basedata {
    $env:GOWORK = "off"
    go run -mod=mod script/initdata.go
}

function initialize-qql-action {
    kocli res gql-action -a resource -g codegen/gqlgen/gqlgen.yaml -f codegen/knockout.yaml
}

function initialize-qql-menu {
    # todo move adminx-ui to project
    Write-Output "kocli res menu -a resource -d {adminui}/src/components/layout/menu.json -f codegen/knockout.yaml"
}

function initialize-app-policy {
    $env:GOWORK = "off"
    go run -mod=mod script/initapppolicy.go
}

function gen {
    genent
    gengql
}

function genent {
    go run codegen/entgen/entc.go
}

function gengql {
    go run codegen/gqlgen/gqlgen.go
}

function genoas {
    # go run codegen/oasgen/oasgen.go
    # go install github.com/tsingsun/woocoo/cmd/woco@main
    woco oasgen -c ./codegen/oasgen/config.yaml
}