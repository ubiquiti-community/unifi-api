# UniFi Network Application UI Crawler

Recursively drives the **UniFi Network Application** SPA and exports a **HAR dump
of every XHR** so the controller's HTTP API can be reverse-engineered into an
OpenAPI spec.

It spins up the controller in **simulation mode** (the `scripts/init.d/demo-mode`
init seeds an `admin`/`admin` account plus demo sites/devices/networks — the same
fixture `terraform-provider-unifi`'s acceptance tests use), waits for the API to
be ready (mirroring `provider_test.go :: waitForUniFiAPI`), then clicks through
the UI exhaustively.

## Quick start

```bash
cd scripts/crawler
./run.sh                 # up + wait + crawl, controller left running
KEEP_UP=0 ./run.sh       # tear the controller down afterwards
HEADLESS=0 ./run.sh      # watch the browser drive itself
MUTATE=0  ./run.sh       # read-only crawl (no create/update/delete)
```

Artifacts land in `scripts/crawler/out/`:

| File | What it is |
|------|------------|
| `unifi-ui.har`        | HAR of every XHR/fetch the UI made (filtered to `/api`, `/proxy`, `/v2`, `/wss`, `/status`, `/guest`). Feed this to an OpenAPI generator. |
| `api-endpoints.json`  | De-duplicated inventory: each distinct path with the set of HTTP **methods** and response **statuses** observed, plus example calls. The OpenAPI skeleton. |
| `routes-visited.json` | Every SPA route the crawler reached. |

## How it works

1. **Bring-up** (`run.sh`) — starts the controller. Uses `docker compose` when the
   plugin is present, otherwise falls back to a `docker run` translation of
   `docker-compose.yaml` (the dev box here has no compose CLI).
2. **Readiness** (`setup.mjs`) — polls `POST /api/login`, `/api/self/sites` and
   `/api/s/<site>/rest/networkconf` until the controller answers, and creates a
   default `Internet 1` WAN network if missing — a faithful port of
   `waitForUniFiAPI`. The interactive first-run wizard itself is dismissed in the
   browser by the crawler (per the brief: *click through the UI setup*).
3. **Crawl** (`crawl.mjs`) — a Playwright/Chromium recursion:
   - Log in (`admin`/`admin`) and dismiss any first-run / cookie / "what's new"
     dialogs.
   - **BFS over every in-app route** reachable from the side nav `<a href>`s,
     seeded with a known-route list so coverage doesn't depend on one UI version's
     DOM. This alone enumerates the entire **GET** read surface.
   - On each page, **click every tab/segment** to reveal sub-panels (more XHR).
   - When `MUTATE=1` (default), **exercise the forms** to capture writes:
     1. submit the **empty** form → `400`/validation XHR (documents required
        fields),
     2. fill with plausible dummy data → **POST** (create),
     3. re-open the created row, tweak it → **PUT/PATCH** (update),
     4. delete it → **DELETE**.
   - Everything is wrapped in timeouts/try-catch; a failed step never aborts the
     crawl, and **failed requests are still recorded** — a validation `400` is as
     useful for the spec as a `200`.
4. **Flush** — closing the browser context writes the HAR; the script also emits
   the `api-endpoints.json` / `routes-visited.json` summaries.

The controller runs in disposable simulation mode, so the destructive
create/update/delete passes are safe.

## Configuration (env vars)

| Var | Default | Meaning |
|-----|---------|---------|
| `UNIFI_API` | `https://localhost:8443` | controller base URL |
| `UNIFI_USERNAME` / `UNIFI_PASSWORD` | `admin` / `admin` | simulation creds |
| `UNIFI_SITE` | `default` | site to crawl |
| `MUTATE` | `1` | `0` = read-only |
| `HEADLESS` | `1` | `0` = show the browser |
| `BUDGET_MS` | `1200000` | hard wall-clock cap; the HAR is always flushed |
| `OUT_DIR` | `./out` | artifact directory |
| `NAV_TIMEOUT_MS` / `IDLE_TIMEOUT_MS` | `30000` / `8000` | per-nav / network-idle waits |

## Turning the HAR into OpenAPI

`api-endpoints.json` is already a method/path/status matrix. For full
request/response schemas, run the HAR through any HAR→OpenAPI tool, e.g.:

```bash
npx @har-sdk/openapi-generator out/unifi-ui.har -o openapi.json   # example
```

then refine against the captured request/response bodies embedded in the HAR.
