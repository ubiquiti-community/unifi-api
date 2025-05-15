package main

//go:generate go run ./cmd/fields -latest
//go:generate go run ./cmd/openapi
//go:generate go tool tfplugingen-openapi generate --config ./assets/generator_config.yml --output ./assets/provider_code_spec.json ./assets/openapi.yaml
