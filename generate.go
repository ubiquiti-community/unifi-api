// Package tools hosts the repo-root go:generate entrypoint. unifi-api generates
// two artifacts from the UniFi controller field definitions, both produced by
// cmd/fields in a single pass:
//
//  1. assets/openapi.yaml            — the OpenAPI 3.1 spec
//  2. assets/openapi-generator/*.cfg — client-generator name mappings
//
// The directive lives at the module root so its relative paths (./cmd/fields,
// assets/) resolve against the repo root when `go generate ./...` runs it.
package tools

//go:generate go run ./cmd/fields -latest
