#!/usr/bin/env sh

SCRIPTPATH="$(
  cd "$(dirname "$0")"
  pwd -P
)"

CURRENT_DIR="$SCRIPTPATH"
ROOT_DIR="$(dirname "$CURRENT_DIR")"
AUTH_PORT="8080"
REST_PORT="8090"

INFRA_LOCAL_COMPOSE_FILE="$ROOT_DIR/seta-train/build/docker-compose.dev.yaml"
MONITORING_COMPOSE_FILE="$ROOT_DIR/seta-train/build/docker-compose.monitoring.yaml"

function local_infra() {
  docker-compose -f "$INFRA_LOCAL_COMPOSE_FILE" "$@"
}

function monitoring() {
  docker-compose -f "$MONITORING_COMPOSE_FILE" "$@"
}

function init() {
    cd "$CURRENT_DIR/.."
    goimports -w ./..
    go fmt ./...
}

function infra() {
  case "$1" in
  up)
    local_infra up "${@:2}"
    ;;
  down)
    local_infra down "${@:2}"
    ;;
  build)
    local_infra build "${@:2}"
    ;;
  *)
    echo "up|down|build [docker-compose command arguments]"
    ;;
  esac
}

function api_start() {
  echo "Starting infrastructure..."
  infra up -d
  migrate_db up
  auth_api_start &
  rest_api_start &
  wait
}

function monitoring_start() {
  echo "Starting monitoring stack..."
  monitoring up -d
  echo "Monitoring stack started!"
  echo "Grafana: http://localhost:3000 (admin/admin)"
  echo "Loki: http://localhost:3100"
}

function auth_api_start() {
    setup_auth_env_variables
    echo "Start auth-service config file: $AUTH_CONFIG_FILE"
    ENTRY_FILE="$ROOT_DIR/seta-train/auth-service/cmd/main.go"
    go run "$ENTRY_FILE" --config-file="$AUTH_CONFIG_FILE" --port="$AUTH_PORT"
}

function rest_api_start() {
    setup_rest_env_variables
    echo "Start rest-service config file: $REST_CONFIG_FILE"
    ENTRY_FILE="$ROOT_DIR/seta-train/rest-service/cmd/service/main.go"
    go run "$ENTRY_FILE" --config-file="$REST_CONFIG_FILE" --port="$REST_PORT"
}

function worker_start() {
  setup_rest_env_variables
  echo "Start worker config file: $REST_CONFIG_FILE"
  ENTRY_FILE="$ROOT_DIR/seta-train/rest-service/cmd/worker/main.go"
  go run "$ENTRY_FILE" --config-file="$REST_CONFIG_FILE"
}

function setup_auth_env_variables() {
    set -a
    export $(grep -v '^#' "$ROOT_DIR/seta-train/auth-service/build/.base.env" | xargs -0) >/dev/null 2>&1
    . "$ROOT_DIR/seta-train/auth-service/build/.base.env"
    set +a
    export AUTH_CONFIG_FILE="$ROOT_DIR/seta-train/auth-service/build/app.yaml"
    export AUTH_PORT="$AUTH_PORT"
}

function setup_rest_env_variables() {
    set -a
    export $(grep -v '^#' "$ROOT_DIR/seta-train/rest-service/build/.base.env" | xargs -0) >/dev/null 2>&1
    . "$ROOT_DIR/seta-train/rest-service/build/.base.env"
    set +a
    export REST_CONFIG_FILE="$ROOT_DIR/seta-train/rest-service/build/app.yaml"
    export REST_PORT="$REST_PORT"
}

function api() {
    case "$1" in
    start)
        api_start
        ;;
    worker_start)
        worker_start
        ;;
    migrate)
        migrate_db "${@:2}"
        ;;
    *)
        echo "[test|start|worker_start|docs_gen|migrate|gqlgen|benchmark]"
        ;;
    esac
}

function migrate_db() {
    
    # source_dir="file://$ROOT_DIR/seta-train/shared-modules/migrations/sql"
    source_dir="file://$(cygpath -m "$ROOT_DIR/seta-train/shared-modules/migrations/sql")"
    conn_str="myuser:mypassword@localhost:5432/mydb?sslmode=disable"

    case "$1" in
    up)
        docker run --rm -v "C:/Users/Ngoc Tran/Desktop/seta-train/shared-modules/migrations/sql:/migrations" migrate/migrate -source="file:///migrations" -database "postgres://myuser:mypassword@host.docker.internal:5432/mydb?sslmode=disable" up
        ;;
    down)
        docker run --rm -v "C:/Users/Ngoc Tran/Desktop/seta-train/shared-modules/migrations/sql:/migrations" migrate/migrate -source="file:///migrations" -database "postgres://myuser:mypassword@host.docker.internal:5432/mydb?sslmode=disable" down
        ;;
    *)
        echo "[up|down]"
        ;;
    esac
}

case "$1" in
init)
    init
    ;;
infra)
    infra "${@:2}"
    ;;
api)
    api "${@:2}"
    ;;
migrate)
    migrate_db "${@:2}"
    ;;
monitoring)
    case "$2" in
    start)
        monitoring_start
        ;;
    stop)
        monitoring down
        ;;
    restart)
        monitoring restart
        ;;
    logs)
        monitoring logs "${@:3}"
        ;;
    *)
        echo "monitoring [start|stop|restart|logs]"
        ;;
    esac
    ;;
*)
    echo "./bin.sh [infra|api|migrate|monitoring|lint|add_version|test]"
    ;;
esac