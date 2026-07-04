# DutchCustomerData Golang SDK



The Golang SDK for the DutchCustomerData API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/dutch-customer-data-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/dutch-customer-data-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/dutch-customer-data-sdk/go=../dutch-customer-data-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    sdk "github.com/voxgig-sdk/dutch-customer-data-sdk/go"
)

func main() {
    client := sdk.New()

    // List euapi records — the value is the array of records itself.
    euapis, err := client.EuApI(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range euapis.([]any) {
        fmt.Println(item)
    }

    // Load a single euapi — the value is the loaded record.
    euapi, err := client.EuApI(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(euapi)
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

euapi, err := client.EuApI(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(euapi) // the loaded mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewDutchCustomerDataSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
DUTCH_CUSTOMER_DATA_TEST_LIVE=TRUE
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewDutchCustomerDataSDK

```go
func NewDutchCustomerDataSDK(options map[string]any) *DutchCustomerDataSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *DutchCustomerDataSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### DutchCustomerDataSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `EuApI` | `(data map[string]any) DutchCustomerDataEntity` | Create an EuApI entity instance. |
| `GlobalApI` | `(data map[string]any) DutchCustomerDataEntity` | Create a GlobalApI entity instance. |
| `NetherlandsApI` | `(data map[string]any) DutchCustomerDataEntity` | Create a NetherlandsApI entity instance. |

### Entity interface (DutchCustomerDataEntity)

All entities implement the `DutchCustomerDataEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    euapi, err := client.EuApI(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil { /* handle */ }
    // euapi is the loaded record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### EuApI

| Field | Description |
| --- | --- |
| `"buyer"` |  |
| `"buyer_country"` |  |
| `"contract_nature"` |  |
| `"html"` |  |
| `"id"` |  |
| `"link"` |  |
| `"notice_type"` |  |
| `"official_language"` |  |
| `"pdf"` |  |
| `"place_of_performance"` |  |
| `"procedure_type"` |  |
| `"publication_date"` |  |
| `"status"` |  |
| `"title"` |  |
| `"vat"` |  |

Operations: List, Load.

API path: `/tender`

#### GlobalApI

| Field | Description |
| --- | --- |
| `"addition"` |  |
| `"admin1"` |  |
| `"admin2"` |  |
| `"admin3"` |  |
| `"bic"` |  |
| `"city"` |  |
| `"currency"` |  |
| `"date"` |  |
| `"dns"` |  |
| `"email"` |  |
| `"found"` |  |
| `"freeformaddress"` |  |
| `"from_currency"` |  |
| `"iban"` |  |
| `"ip"` |  |
| `"lat"` |  |
| `"lei"` |  |
| `"letter"` |  |
| `"lon"` |  |
| `"municipality"` |  |
| `"number"` |  |
| `"password"` |  |
| `"phone"` |  |
| `"population"` |  |
| `"postcode"` |  |
| `"province"` |  |
| `"province_code"` |  |
| `"score"` |  |
| `"status"` |  |
| `"street"` |  |
| `"type"` |  |
| `"url"` |  |
| `"webrank"` |  |

Operations: Create, List, Load.

API path: `/password`

#### NetherlandsApI

| Field | Description |
| --- | --- |
| `"active"` |  |
| `"addition"` |  |
| `"city"` |  |
| `"coc"` |  |
| `"construction_year"` |  |
| `"floor_area"` |  |
| `"freeformaddress"` |  |
| `"id"` |  |
| `"lat"` |  |
| `"letter"` |  |
| `"lon"` |  |
| `"municipality"` |  |
| `"name"` |  |
| `"number"` |  |
| `"postcode"` |  |
| `"province"` |  |
| `"province_code"` |  |
| `"purpose"` |  |
| `"street"` |  |
| `"type"` |  |
| `"vestiging"` |  |

Operations: List.

API path: `/bag`



## Entities


### EuApI

Create an instance: `eu_ap_i := client.EuApI(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `buyer` | ``$STRING`` |  |
| `buyer_country` | ``$STRING`` |  |
| `contract_nature` | ``$STRING`` |  |
| `html` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `link` | ``$STRING`` |  |
| `notice_type` | ``$STRING`` |  |
| `official_language` | ``$STRING`` |  |
| `pdf` | ``$STRING`` |  |
| `place_of_performance` | ``$STRING`` |  |
| `procedure_type` | ``$STRING`` |  |
| `publication_date` | ``$STRING`` |  |
| `status` | ``$STRING`` |  |
| `title` | ``$STRING`` |  |
| `vat` | ``$OBJECT`` |  |

#### Example: Load

```go
eu_ap_i, err := client.EuApI(nil).Load(map[string]any{"id": "eu_ap_i_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(eu_ap_i) // the loaded record
```

#### Example: List

```go
eu_ap_is, err := client.EuApI(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(eu_ap_is) // the array of records
```


### GlobalApI

Create an instance: `global_ap_i := client.GlobalApI(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addition` | ``$STRING`` |  |
| `admin1` | ``$STRING`` |  |
| `admin2` | ``$STRING`` |  |
| `admin3` | ``$STRING`` |  |
| `bic` | ``$OBJECT`` |  |
| `city` | ``$STRING`` |  |
| `currency` | ``$OBJECT`` |  |
| `date` | ``$STRING`` |  |
| `dns` | ``$OBJECT`` |  |
| `email` | ``$OBJECT`` |  |
| `found` | ``$INTEGER`` |  |
| `freeformaddress` | ``$STRING`` |  |
| `from_currency` | ``$STRING`` |  |
| `iban` | ``$OBJECT`` |  |
| `ip` | ``$OBJECT`` |  |
| `lat` | ``$NUMBER`` |  |
| `lei` | ``$OBJECT`` |  |
| `letter` | ``$STRING`` |  |
| `lon` | ``$NUMBER`` |  |
| `municipality` | ``$STRING`` |  |
| `number` | ``$INTEGER`` |  |
| `password` | ``$OBJECT`` |  |
| `phone` | ``$OBJECT`` |  |
| `population` | ``$INTEGER`` |  |
| `postcode` | ``$STRING`` |  |
| `province` | ``$STRING`` |  |
| `province_code` | ``$STRING`` |  |
| `score` | ``$NUMBER`` |  |
| `status` | ``$STRING`` |  |
| `street` | ``$STRING`` |  |
| `type` | ``$STRING`` |  |
| `url` | ``$OBJECT`` |  |
| `webrank` | ``$OBJECT`` |  |

#### Example: Load

```go
global_ap_i, err := client.GlobalApI(nil).Load(map[string]any{"id": "global_ap_i_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(global_ap_i) // the loaded record
```

#### Example: List

```go
global_ap_is, err := client.GlobalApI(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(global_ap_is) // the array of records
```

#### Example: Create

```go
result, err := client.GlobalApI(nil).Create(map[string]any{
}, nil)
```


### NetherlandsApI

Create an instance: `netherlands_ap_i := client.NetherlandsApI(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `active` | ``$INTEGER`` |  |
| `addition` | ``$STRING`` |  |
| `city` | ``$STRING`` |  |
| `coc` | ``$STRING`` |  |
| `construction_year` | ``$INTEGER`` |  |
| `floor_area` | ``$INTEGER`` |  |
| `freeformaddress` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `lat` | ``$NUMBER`` |  |
| `letter` | ``$STRING`` |  |
| `lon` | ``$NUMBER`` |  |
| `municipality` | ``$STRING`` |  |
| `name` | ``$STRING`` |  |
| `number` | ``$STRING`` |  |
| `postcode` | ``$STRING`` |  |
| `province` | ``$STRING`` |  |
| `province_code` | ``$STRING`` |  |
| `purpose` | ``$STRING`` |  |
| `street` | ``$STRING`` |  |
| `type` | ``$STRING`` |  |
| `vestiging` | ``$STRING`` |  |

#### Example: List

```go
netherlands_ap_is, err := client.NetherlandsApI(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(netherlands_ap_is) // the array of records
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller. An unexpected panic triggers the
`PreUnexpected` hook.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/dutch-customer-data-sdk/go/
├── dutch-customer-data.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/dutch-customer-data-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
euapi := client.EuApI(nil)
euapi.Load(map[string]any{"id": "example_id"}, nil)

// euapi.Data() now returns the loaded euapi data
// euapi.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
