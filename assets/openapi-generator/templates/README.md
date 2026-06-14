# Custom templates (override hook)

OpenAPI Generator renders clients from per-generator [mustache] templates. To
override any of them for a language, drop the modified `.mustache` files in a
subdirectory here and point the generator at it.

1. Export a generator's default templates to copy from:

   ```bash
   docker run --rm -v "$PWD:/local" openapitools/openapi-generator-cli \
     author template -g go -o /local/assets/openapi-generator/templates/go
   ```

2. Edit only the templates you need to change; delete the rest (the generator
   falls back to its built-ins for anything not present).

3. Wire it into the language's generation by adding to `languages/<lang>.yaml`:

   ```yaml
   templateDir: assets/openapi-generator/templates/<lang>
   ```

No template overrides are committed by default — the generated `*-name-mappings`
files already reproduce go-unifi's naming, which is the only customization most
clients need.

[mustache]: https://mustache.github.io/
