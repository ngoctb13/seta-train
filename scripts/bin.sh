#!/usr/bin/env sh

SCRIPTPATH="$(
  cd "$(dirname "$0")"
  pwd -P
)"

CURRENT_DIR=$SCRIPTPATH
ROOT_DIR="$(dirname $CURRENT_DIR)"
PORT="8090"

INFRA_LOCAL_COMPOSE_FILE=$ROOT_DIR/build/docker-compose.dev.yaml

function local_infra() {
  docker-compose -f $INFRA_LOCAL_COMPOSE_FILE $@
}

function init() {
    cd $CURRENT_DIR/..
    goimports -w ./..
    go fmt ./...
}

function infra() {
  case $1 in
  up)
    local_infra up ${@:2}
    ;;
  down)
    local_infra down ${@:2}
    ;;
  build)
    local_infra build ${@:2}
    ;;
  *)
    echo "up|down|build [docker-compose command arguments]"
    ;;
  esac
}

function api_start() {
  echo "Starting infrastructure..."
  infra up -d
  setup_env_variables
  echo "Start api app config file: $CONFIG_FILE"
  ENTRY_FILE="$ROOT_DIR/cmd/service/main.go"
  go run $ENTRY_FILE --config-file=$CONFIG_FILE --port=$PORT
}

function setup_env_variables() {
    set -a
    export $(grep -v '^#' "$ROOT_DIR/build/.base.env" | xargs -0) >/dev/null 2>&1
    . $ROOT_DIR/build/.base.env
    set +a
    export CONFIG_FILE=$ROOT_DIR/build/app.yaml
    export PORT=$PORT
}

function api() {
    case $1 in
    start)
        api_start
        ;;
    worker_start)
        worker_start
        ;;
    migrate)
        migrate_db ${@:2}
        ;;
    *)
        echo "[test|start|worker_start|docs_gen|migrate|gqlgen|benchmark]"
        ;;
    esac
}

function migrate_db() {
    source_dir="file://$ROOT_DIR/migrations/sql"
    conn_str="postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"

    case $1 in
    up)
        migrate -source $source_dir -database "mysql://$conn_str" up
        ;;
    down)
        migrate -source $source_dir -database "mysql://$conn_str" down
        ;;
    *)
        echo "[up|down]"
        ;;
    esac
}

case $1 in
init)
    init
    ;;
infra)
    infra ${@:2}
    ;;
api)
    api ${@:2}
    ;;
*)
    echo "./scripts/bin.sh [infra|api|lint|add_version|test]"
    ;;
esac