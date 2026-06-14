# unifi-api — OpenAPI spec + multi-language client generation.
#
#   make openapi        regenerate assets/openapi.yaml + generator mappings
#   make client-go      generate Go API + clients/go/models package
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
GO_CLIENT_DIR     := clients/go
GO_MODEL_DIR      := $(GO_CLIENT_DIR)/models
GO_MODULE         := github.com/ubiquiti-community/unifi-api/clients/go

.PHONY: openapi clients client-go $(addprefix client-,$(LANGS))

openapi:
	go generate ./...

clients: $(addprefix client-,$(LANGS))

# Generate models and APIs separately. The postprocessor bundles nested structs
# with their owning top-level model and adds aliases that preserve the client's
# existing unifi.Type API while the structs live in the models subpackage.
client-go:
	rm -rf $(GO_CLIENT_DIR)
	$(OPENAPI_GENERATOR) -i $(SPEC) -o $(GO_MODEL_DIR) \
		-c $(GENDIR)/languages/go-model.yaml \
		--global-property 'models,modelDocs=false,modelTests=false,supportingFiles=utils.go' \
		--openapi-normalizer $(NORMALIZER) \
		--model-name-mappings '$(MODEL_MAPPINGS)' \
		--name-mappings '$(GO_NAME_MAPPINGS)'
	go run ./cmd/openapi-generator-postprocess \
		-models $(GO_MODEL_DIR) \
		-mappings $(GENDIR)/model-file-mappings.cfg \
		-aliases $(GO_CLIENT_DIR)/model_aliases.go \
		-alias-package unifi \
		-alias-import $(GO_MODULE)/models
	$(OPENAPI_GENERATOR) -i $(SPEC) -o $(GO_CLIENT_DIR) \
		-c $(GENDIR)/languages/go.yaml \
		--global-property 'apis,apiDocs=false,apiTests=false,supportingFiles' \
		--git-user-id ubiquiti-community \
		--git-repo-id unifi-api/clients/go \
		--openapi-normalizer $(NORMALIZER) \
		--model-name-mappings '$(MODEL_MAPPINGS)' \
		--name-mappings '$(GO_NAME_MAPPINGS)'

# Every other language: shared model-name mappings + idiomatic property casing.
client-%:
	$(OPENAPI_GENERATOR) -i $(SPEC) -o clients/$* \
		-c $(GENDIR)/languages/$*.yaml \
		--openapi-normalizer $(NORMALIZER) \
		--model-name-mappings '$(MODEL_MAPPINGS)'
