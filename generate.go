// Package tools hosts the repo-root go:generate entrypoints. Three artifacts
// are produced when you run `go generate ./...` from the repo root:
//
//  1. assets/openapi.yaml                       — OpenAPI 3.1 spec
//  2. assets/openapi-generator/*.cfg            — client-generator name mappings
//  3. assets/oapi-codegen-exp/{unifi,settings}.yaml — oapi-codegen-exp configs
//
// Steps 1-3 are all written by the first directive (cmd/fields). The two
// oapi-codegen directives that follow then consume the freshly-written configs
// to emit the typed Go clients into clients/go/.
//
// All paths are relative to the repo root, which is where go generate sets
// the working directory when run from this file.
package tools

//go:generate go run ./cmd/fields -latest
//go:generate go tool oapi-codegen -config assets/oapi-codegen-exp/unifi.yaml assets/openapi.yaml
//go:generate go tool oapi-codegen -config assets/oapi-codegen-exp/settings.yaml assets/openapi.yaml
