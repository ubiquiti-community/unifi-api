This project generates OpenAPI specifications using the `go generate ./...` command.
You will be updating templates and adding additional configuration files to match the pre-generated json files.
This project first runs `cmd/fields`, which populates the `cmd/openapi` files with the necessary info to generate OpenAPI specifications using `swaggest/openapi-go`.
It also generates some additional metadata for use with the `tfplugingen-openapi` generator.
Finally, it runs `go tool tfplugingen-openapi generate --config ./assets/generator_config.yml --output ./assets/provider_code_spec.json ./assets/openapi.yaml` to generate a terraform provider.
Complete the following tasks:
1. Update the go templates to add full CRUD operations for each resource. This will exclude the Settings resources as they only have update and read operations.
2. Complete the generation of the assets/generator_config.yml and the provider_code_spec.json files.
3. Ensure that the generated OpenAPI specifications are valid and conform to the expected structure.
4. Implement the terraform provider code to match the generated OpenAPI specifications.