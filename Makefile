# unifi-api — OpenAPI spec + multi-language client generation.
#
#   make openapi        regenerate assets/openapi.yaml + generator mappings
#   make client-go      generate one client into clients/go/
#   make clients        generate every curated client
#
# Client generation runs OpenAPI Generator via Docker (no local Java needed).

SPEC          := assets/openapi.yaml
GENDIR        := assets/openapi-generator
LANGS         := go python typescript java rust
OPENAPI_GENERATOR_VERSION ?= v7.10.0
OPENAPI_GENERATOR ?= docker run --rm -u $$(id -u):$$(id -g) \
	-v "$$PWD:/local" -w /local \
	openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_VERSION) generate

# Shared "modification overrides" applied to every language.
NORMALIZER := SIMPLIFY_ONEOF_ANYOF=true,REFACTOR_ALLOF_WITH_PROPERTIES_ONLY=true

# Mapping CLI values, joined from the generated .cfg files (comments stripped).
HASH := \#
MODEL_MAPPINGS = $(shell grep -v '^$(HASH)' $(GENDIR)/model-name-mappings.cfg | paste -sd,)
GO_NAME_MAPPINGS = $(shell grep -v '^$(HASH)' $(GENDIR)/go-name-mappings.cfg | paste -sd,)

.PHONY: openapi clients $(addprefix client-,$(LANGS))

openapi:
	go generate ./...

clients: $(addprefix client-,$(LANGS))

# Go gets the property-name mappings too (go-unifi field-name parity).
client-go:
	$(OPENAPI_GENERATOR) -i $(SPEC) -o clients/go \
		-c $(GENDIR)/languages/go.yaml \
		--openapi-normalizer $(NORMALIZER) \
		--model-name-mappings '$(MODEL_MAPPINGS)' \
		--name-mappings '$(GO_NAME_MAPPINGS)'

# Every other language: shared model-name mappings + idiomatic property casing.
client-%:
	$(OPENAPI_GENERATOR) -i $(SPEC) -o clients/$* \
		-c $(GENDIR)/languages/$*.yaml \
		--openapi-normalizer $(NORMALIZER) \
		--model-name-mappings '$(MODEL_MAPPINGS)'
