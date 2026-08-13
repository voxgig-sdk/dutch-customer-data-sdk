# DutchCustomerData Ruby SDK



The Ruby SDK for the DutchCustomerData API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.EuApI` — with named operations (`list`/`load`/`create`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/dutch-customer-data-sdk/releases](https://github.com/voxgig-sdk/dutch-customer-data-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "DutchCustomerData_sdk"

client = DutchCustomerDataSDK.new
```

### 2. List euapi records

```ruby
begin
  # list returns an Array of EuApI records — iterate directly.
  euapis = client.EuApI.list
  euapis.each do |item|
    puts "#{item["id"]} #{item["active"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load an euapi

```ruby
begin
  # load returns the ENTITY — call data_get for the EuApI record (raises on error).
  euapi = client.EuApI.load({ "id" => "example_id" })
  puts euapi
rescue => err
  warn "load failed: #{err}"
end
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  euapis = client.EuApI.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```ruby
client = DutchCustomerDataSDK.test({
  "entity" => { "euapi" => { "test01" => { "id" => "test01" } } },
})

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
euapi = client.EuApI.list()
puts euapi
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = DutchCustomerDataSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
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
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### DutchCustomerDataSDK

```ruby
require_relative "DutchCustomerData_sdk"
client = DutchCustomerDataSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = DutchCustomerDataSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### DutchCustomerDataSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `EuApI` | `(data) -> EuApIEntity` | Create an EuApI entity instance. |
| `GlobalApI` | `(data) -> GlobalApIEntity` | Create a GlobalApI entity instance. |
| `NetherlandsApI` | `(data) -> NetherlandsApIEntity` | Create a NetherlandsApI entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `DutchCustomerDataError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### EuApI

| Field | Description |
| --- | --- |
| `active` |  |
| `address` |  |
| `buyer` |  |
| `buyer_country` |  |
| `city` |  |
| `contract_nature` |  |
| `country` |  |
| `html` |  |
| `id` |  |
| `link` |  |
| `name` |  |
| `notice_type` |  |
| `official_language` |  |
| `pdf` |  |
| `place_of_performance` |  |
| `postcode` |  |
| `procedure_type` |  |
| `publication_date` |  |
| `response_date` |  |
| `title` |  |
| `vat` |  |

Operations: List, Load.

API path: `/tender`

#### GlobalApI

| Field | Description |
| --- | --- |
| `addition` |  |
| `address` |  |
| `admin1` |  |
| `admin2` |  |
| `admin3` |  |
| `bank` |  |
| `bic` |  |
| `browser` |  |
| `builtwith` |  |
| `carrier` |  |
| `city` |  |
| `cloudflare` |  |
| `commoncrawl` |  |
| `content_length` |  |
| `content_type` |  |
| `country` |  |
| `country_code` |  |
| `crux` |  |
| `device_family` |  |
| `device_name` |  |
| `device_type` |  |
| `disposable` |  |
| `dns_a` |  |
| `dns_mx` |  |
| `dns_ns` |  |
| `dns_soa` |  |
| `dns_txt` |  |
| `dns_www_a` |  |
| `dnsserver` |  |
| `domain` |  |
| `domcop` |  |
| `email` |  |
| `found` |  |
| `free` |  |
| `freeformaddress` |  |
| `host` |  |
| `host_type` |  |
| `hostio` |  |
| `http_code` |  |
| `iban` |  |
| `iban_human` |  |
| `int` |  |
| `international` |  |
| `ip` |  |
| `ipint` |  |
| `ismobile` |  |
| `lat` |  |
| `lei` |  |
| `letter` |  |
| `local_id` |  |
| `lon` |  |
| `mailserver` |  |
| `majestic` |  |
| `message` |  |
| `municipality` |  |
| `mx_host` |  |
| `mx_ip` |  |
| `name` |  |
| `national` |  |
| `number` |  |
| `ocid` |  |
| `pagerank` |  |
| `platform` |  |
| `population` |  |
| `postcode` |  |
| `province` |  |
| `province_code` |  |
| `redirect_count` |  |
| `region` |  |
| `register_id` |  |
| `renewal_date` |  |
| `score` |  |
| `sepa` |  |
| `spf` |  |
| `status` |  |
| `street` |  |
| `success` |  |
| `swift` |  |
| `tranco` |  |
| `type` |  |
| `umbrella` |  |
| `url` |  |
| `user` |  |
| `user_agent` |  |
| `valid` |  |
| `verified` |  |
| `verified_checksum` |  |
| `webrank` |  |
| `wrong_email` |  |
| `wrong_format` |  |
| `wrong_password` |  |
| `wrong_phone` |  |

Operations: Create, List, Load.

API path: `/password`

#### NetherlandsApI

| Field | Description |
| --- | --- |
| `active` |  |
| `addition` |  |
| `city` |  |
| `coc` |  |
| `construction_year` |  |
| `floor_area` |  |
| `freeformaddress` |  |
| `id` |  |
| `lat` |  |
| `letter` |  |
| `lon` |  |
| `municipality` |  |
| `name` |  |
| `number` |  |
| `postcode` |  |
| `province` |  |
| `province_code` |  |
| `purpose` |  |
| `street` |  |
| `type` |  |
| `vestiging` |  |

Operations: List.

API path: `/bag`



## Entities


### EuApI

Create an instance: `eu_ap_i = client.EuApI`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `active` | `Integer` |  |
| `address` | `String` |  |
| `buyer` | `String` |  |
| `buyer_country` | `String` |  |
| `city` | `String` |  |
| `contract_nature` | `String` |  |
| `country` | `String` |  |
| `html` | `String` |  |
| `id` | `String` |  |
| `link` | `String` |  |
| `name` | `String` |  |
| `notice_type` | `String` |  |
| `official_language` | `String` |  |
| `pdf` | `String` |  |
| `place_of_performance` | `String` |  |
| `postcode` | `String` |  |
| `procedure_type` | `String` |  |
| `publication_date` | `String` |  |
| `response_date` | `String` |  |
| `title` | `String` |  |
| `vat` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the EuApI record (raises on error).
eu_ap_i = client.EuApI.load({ "id" => "eu_ap_i_id" })
```

#### Example: List

```ruby
# list returns an Array of EuApI records (raises on error).
eu_ap_is = client.EuApI.list
```


### GlobalApI

Create an instance: `global_ap_i = client.GlobalApI`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addition` | `String` |  |
| `address` | `String` |  |
| `admin1` | `String` |  |
| `admin2` | `String` |  |
| `admin3` | `String` |  |
| `bank` | `String` |  |
| `bic` | `String` |  |
| `browser` | `String` |  |
| `builtwith` | `Integer` |  |
| `carrier` | `String` |  |
| `city` | `String` |  |
| `cloudflare` | `Integer` |  |
| `commoncrawl` | `Integer` |  |
| `content_length` | `Integer` |  |
| `content_type` | `String` |  |
| `country` | `String` |  |
| `country_code` | `String` |  |
| `crux` | `Integer` |  |
| `device_family` | `String` |  |
| `device_name` | `String` |  |
| `device_type` | `String` |  |
| `disposable` | `Integer` |  |
| `dns_a` | `Array` |  |
| `dns_mx` | `Array` |  |
| `dns_ns` | `Array` |  |
| `dns_soa` | `Array` |  |
| `dns_txt` | `Array` |  |
| `dns_www_a` | `Array` |  |
| `dnsserver` | `String` |  |
| `domain` | `String` |  |
| `domcop` | `Integer` |  |
| `email` | `String` |  |
| `found` | `Integer` |  |
| `free` | `Integer` |  |
| `freeformaddress` | `String` |  |
| `host` | `String` |  |
| `host_type` | `String` |  |
| `hostio` | `Integer` |  |
| `http_code` | `Integer` |  |
| `iban` | `String` |  |
| `iban_human` | `String` |  |
| `int` | `String` |  |
| `international` | `String` |  |
| `ip` | `String` |  |
| `ipint` | `Integer` |  |
| `ismobile` | `Integer` |  |
| `lat` | `Float` |  |
| `lei` | `String` |  |
| `letter` | `String` |  |
| `local_id` | `String` |  |
| `lon` | `Float` |  |
| `mailserver` | `String` |  |
| `majestic` | `Integer` |  |
| `message` | `String` |  |
| `municipality` | `String` |  |
| `mx_host` | `String` |  |
| `mx_ip` | `String` |  |
| `name` | `String` |  |
| `national` | `String` |  |
| `number` | `Integer` |  |
| `ocid` | `String` |  |
| `pagerank` | `Integer` |  |
| `platform` | `String` |  |
| `population` | `Integer` |  |
| `postcode` | `String` |  |
| `province` | `String` |  |
| `province_code` | `String` |  |
| `redirect_count` | `Integer` |  |
| `region` | `String` |  |
| `register_id` | `String` |  |
| `renewal_date` | `String` |  |
| `score` | `Float` |  |
| `sepa` | `Integer` |  |
| `spf` | `String` |  |
| `status` | `String` |  |
| `street` | `String` |  |
| `success` | `Integer` |  |
| `swift` | `Integer` |  |
| `tranco` | `Integer` |  |
| `type` | `String` |  |
| `umbrella` | `Integer` |  |
| `url` | `String` |  |
| `user` | `String` |  |
| `user_agent` | `String` |  |
| `valid` | `Integer` |  |
| `verified` | `Boolean` |  |
| `verified_checksum` | `Boolean` |  |
| `webrank` | `Integer` |  |
| `wrong_email` | `Integer` |  |
| `wrong_format` | `Integer` |  |
| `wrong_password` | `Integer` |  |
| `wrong_phone` | `Integer` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the GlobalApI record (raises on error).
global_ap_i = client.GlobalApI.load()
```

#### Example: List

```ruby
# list returns an Array of GlobalApI records (raises on error).
global_ap_is = client.GlobalApI.list
```

#### Example: Create

```ruby
global_ap_i = client.GlobalApI.create({
})
```


### NetherlandsApI

Create an instance: `netherlands_ap_i = client.NetherlandsApI`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `active` | `Integer` |  |
| `addition` | `String` |  |
| `city` | `String` |  |
| `coc` | `String` |  |
| `construction_year` | `Integer` |  |
| `floor_area` | `Integer` |  |
| `freeformaddress` | `String` |  |
| `id` | `String` |  |
| `lat` | `Float` |  |
| `letter` | `String` |  |
| `lon` | `Float` |  |
| `municipality` | `String` |  |
| `name` | `String` |  |
| `number` | `String` |  |
| `postcode` | `String` |  |
| `province` | `String` |  |
| `province_code` | `String` |  |
| `purpose` | `String` |  |
| `street` | `String` |  |
| `type` | `String` |  |
| `vestiging` | `String` |  |

#### Example: List

```ruby
# list returns an Array of NetherlandsApI records (raises on error).
netherlands_ap_is = client.NetherlandsApI.list
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

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

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── DutchCustomerData_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`DutchCustomerData_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
euapi = client.EuApI
euapi.list()

# euapi.data_get now returns the euapi data from the last list
# euapi.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
