#!/usr/bin/env bash
#
# End-to-end: bring up the UniFi controller (simulation mode), wait for the API,
# then run the recursive UI crawler and emit a HAR dump for OpenAPI generation.
#
#   ./run.sh                # up + wait + crawl, leave controller running
#   KEEP_UP=0 ./run.sh      # tear the controller down afterwards
#   HEADLESS=0 ./run.sh     # watch the browser
#   MUTATE=0 ./run.sh       # read-only crawl (no create/update/delete)
#
set -euo pipefail

HERE="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$HERE/../.." && pwd)"
COMPOSE_FILE="$PROJECT_ROOT/docker-compose.yaml"
CONTAINER=unifi
IMAGE=jacobalberty/unifi:latest

export UNIFI_API="${UNIFI_API:-https://localhost:8443}"
export UNIFI_USERNAME="${UNIFI_USERNAME:-admin}"
export UNIFI_PASSWORD="${UNIFI_PASSWORD:-admin}"
export NODE_TLS_REJECT_UNAUTHORIZED=0   # accept the controller's self-signed cert
KEEP_UP="${KEEP_UP:-1}"

log() { echo "[run] $*"; }

# --- 1. bring up the controller --------------------------------------------
up() {
  if docker compose version >/dev/null 2>&1; then
    log "starting controller via 'docker compose'"
    docker compose -f "$COMPOSE_FILE" up -d
  else
    log "'docker compose' unavailable; using 'docker run' (compose translation)"
    if docker ps -a --format '{{.Names}}' | grep -qx "$CONTAINER"; then
      docker start "$CONTAINER" >/dev/null
    else
      docker run -d \
        --name "$CONTAINER" --init --user unifi \
        -e PKGURL="" -e UNIFI_STDOUT=true -e TZ=Etc/UTC \
        -v "$PROJECT_ROOT/scripts/init.d/demo-mode:/unifi/init.d/demo-mode:Z" \
        -p 8443:8443 -p 3478:3478/udp -p 10001:10001/udp -p 8080:8080 \
        --restart unless-stopped "$IMAGE" >/dev/null
    fi
  fi
}

down() {
  if docker compose version >/dev/null 2>&1; then
    docker compose -f "$COMPOSE_FILE" down
  else
    docker rm -f "$CONTAINER" >/dev/null 2>&1 || true
  fi
}

# --- 2. node deps ----------------------------------------------------------
deps() {
  cd "$HERE"
  if [ ! -d node_modules/playwright ]; then
    log "installing node dependencies"
    npm install --no-audit --no-fund
  fi
  # Ensure the chromium browser binary is present.
  npx playwright install chromium >/dev/null 2>&1 || npx playwright install-deps chromium || true
}

# --- main ------------------------------------------------------------------
up
deps

log "waiting for UniFi API readiness"
node "$HERE/setup.mjs"

log "starting recursive crawl"
node "$HERE/crawl.mjs"

if [ "$KEEP_UP" != "1" ]; then
  log "tearing down controller (KEEP_UP=0)"
  down
fi

log "done. artifacts in $HERE/out/"
