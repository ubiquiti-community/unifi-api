# Unifi OpenAPI Definition

Unifi API spec <https://10.0.0.1/proxy/network/api-docs/integration.json>

```bash
go generate ./...
go tool tfplugingen-framework generate all --input ./assets/provider_code_spec.json --output internal/provider
```
