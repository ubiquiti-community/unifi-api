# `oapi-codegen` V3

This is an experimental prototype of a V3 version of oapi-codegen. The generated code and command line options are not yet stable. Use at your
own risk.

## What is new in Version 3

This directory contains an experimental version of oapi-codegen's future V3 version, which is based on [libopenapi](https://github.com/pb33f/libopenapi),
instead of the prior [kin-openapi](https://github.com/getkin/kin-openapi). This change necessitated a nearly complete rewrite, but we strive to be as
compatible as possible.

What is working:

- All model, client and server generation as in earlier versions.
- We have added Webhook and Callback support, please see `./examples`, which contains the ubiquitous OpenAPI pet shop implemented in all supported servers
    and examples of webhooks and callbacks implemented on top of the `http.ServeMux` server, with no additional imports.
- Echo V5 support has been added (Go 1.25 required)
- The `runtime` has changed a lot. By default, we generate all the needed runtime
    functions into your generated code. You can, optionally, generate your own runtime
    package locally, to avoid duplication between multiple openapi specifications. This
    was done because the version pinning between runtime and the codegen was exceedingly
    annoying, so now, the runtime is embedded into the generator itself, and there is
    no versioning issue.
  
What is missing:

- Middleware, this is for someone else to solve.
- Good documentation. You'll have to read over the config file code to see how
    to configure.

## Differences in V3

V3 is a brand new implementation, and may (will) contain new bugs, but also strives to fix many current, existing bugs. We've run quite a few
conformance tests against specifications in old Issues, and we're looking pretty good! Please try this out, and if it failes in some way, please
file Issues.

### Normalized extension names

V3 normalizes all extension names under the `x-oapi-codegen-` prefix. The old names are still accepted for backwards compatibility.

| V2 | This version                           | Scope | Purpose |
|---|----------------------------------------|---|---|
| `x-go-type` + `x-go-type-import` | `x-oapi-codegen-type-override`         | Schema, Property | Use an external Go type instead of generating one. V3 combines type and import into a single value: `"TypeName;import/path"`. |
| `x-go-name` | `x-oapi-codegen-name-override`         | Property | Override the generated Go field name. |
| `x-go-type-name` | `x-oapi-codegen-type-name-override`    | Schema | Override the generated Go type name. |
| `x-go-type-skip-optional-pointer` | `x-oapi-codegen-skip-optional-pointer` | Property | Don't wrap optional fields in a pointer. |
| `x-go-json-ignore` | `x-oapi-codegen-json-ignore`           | Property | Exclude the field from JSON (`json:"-"`). |
| `x-omitempty` | `x-oapi-codegen-omitempty`             | Property | Explicitly control the `omitempty` JSON tag. |
| `x-omitzero` | `x-oapi-codegen-omitzero`              | Property | Add `omitzero` to the JSON tag (Go 1.24+ `encoding/json/v2`). |
| `x-enum-varnames` / `x-enumNames` | `x-oapi-codegen-enum-varnames`         | Schema (enum) | Override generated enum constant names. |
| `x-deprecated-reason` | `x-oapi-codegen-deprecated-reason`     | Schema, Operation | Provide a deprecation reason for documentation. |
| `x-order` | `x-oapi-codegen-order`                 | Property | Control field ordering in generated structs. |

### OpenAPI V3.1 Feature Support

Thanks to [libopenapi](https://github.com/pb33f/libopenapi), we are able to parse OpenAPI 3.1 and 3.2 specifications. They are functionally similar, you can
read the differences between `nullable` fields yourself, but they add some new functionality, namely `webhooks` and `callbacks`. We support all of them in
this prototype. `callbacks` and `webhooks` are basically the inverse of `paths`. Webhooks contain no URL element in their definition, so we can't register handlers
for you in your http router of choice, you have to do that yourself. Callbacks support complex request URL's which may reference the original request. This is
something you need to pull out of the request body, and doing it generically is difficult, so we punt this problem, for now, to our users.

Please see the [webhook example](examples/webhook/). It creates a little server that pretends to be a door badge reader, and it generates an event stream
about people coming and going. Any number of clients may subscribe to this event. See the [doc.go](examples/webhook/doc.go) for usage examples.

The [callback example](examples/callback), creates a little server that pretends to plant trees. Each tree planting request contains a callback to be notified
when tree planting is complete. We invoke those in a random order via delays, and the client prints out callbacks as they happen. Please see [doc.go](examples/callback/doc.go) for usage.

#### Enum via `oneOf` + `const`

OpenAPI 3.1 lets you express a named enum with per-value documentation by putting each variant in a `oneOf` branch with `const` and `title`:

```yaml
Severity:
  type: integer
  oneOf:
    - title: HIGH
      const: 2
      description: An urgent problem
    - title: MEDIUM
      const: 1
    - title: LOW
      const: 0
      description: Can wait forever
```

V3 detects this idiom and emits a regular Go enum (`type Severity int` with `HIGH`, `MEDIUM`, `LOW` constants) — with the `description` rendered as a per-value doc comment — instead of a `oneOf` union. All branches must carry both `const` and `title`, and the outer schema must declare a scalar `type` (`string` or `integer`); otherwise the schema falls through to the standard union generator. Set `generation.skip-enum-via-oneof: true` to disable detection.

### Flexible Configuration

oapi-codegen V3 tries to make no assumptions about which initialisms, struct tags, or name mangling that is correct for you. A very [flexible configuration file](Configuration.md) allows you to override anything.

### No runtime dependency

V2 generated code relied on `github.com/oapi-codegen/runtime` for parameter binding and styling. This was a complaint from lots of people due to various
audit requirements. V3 embeds all necessary helper functions and helper types into the spec. There are no longer generic, parameterized functions that
handle arbitrary parameters, but rather very specific functions for each kind of parameter, and we call the correct little helper versus a generic
runtime helper.

We still use the code generator to produce a pre-generated `runtime` package, which you are
welcome to use. It will always be consistent with the code generated with the corresponding
oapi-codegen. If you have lots of OpenAPI specs locally, you can also generate the runtime
package, as we do, in your own code to avoid bloat.

### Models now support default values configured in the spec

Every model which we generate supports an `ApplyDefaults()` function. It recursively applies defaults on
any unset optional fields. There's a little caveat here, in that some types are external references, so
we call `ApplyDefaults()` on them via reflection. This might call an `ApplyDefaults()` which is completely
unrelated to what we're doing. Please let me know if this feature is causing trouble.

## Installation

Go 1.25 is required, install like so:

    go get -tool github.com/oapi-codegen/oapi-codegen-exp/cmd/oapi-codegen@latest

You can then run the code generator

    //go:generate go run github.com/oapi-codegen/oapi-codegen-exp/cmd/oapi-codegen
