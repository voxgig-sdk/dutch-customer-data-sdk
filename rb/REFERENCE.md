# DutchCustomerData Ruby SDK Reference

Complete API reference for the DutchCustomerData Ruby SDK.


## DutchCustomerDataSDK

### Constructor

```ruby
require_relative 'DutchCustomerData_sdk'

client = DutchCustomerDataSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DutchCustomerDataSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = DutchCustomerDataSDK.test
```


### Instance Methods

#### `EuApI(data = nil)`

Create a new `EuApI` entity instance. Pass `nil` for no initial data.

#### `GlobalApI(data = nil)`

Create a new `GlobalApI` entity instance. Pass `nil` for no initial data.

#### `NetherlandsApI(data = nil)`

Create a new `NetherlandsApI` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## EuApIEntity

```ruby
eu_ap_i = client.EuApI
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `buyer` | `String` | No |  |
| `buyer_country` | `String` | No |  |
| `contract_nature` | `String` | No |  |
| `html` | `String` | No |  |
| `id` | `String` | No |  |
| `link` | `String` | No |  |
| `notice_type` | `String` | No |  |
| `official_language` | `String` | No |  |
| `pdf` | `String` | No |  |
| `place_of_performance` | `String` | No |  |
| `procedure_type` | `String` | No |  |
| `publication_date` | `String` | No |  |
| `status` | `String` | No |  |
| `title` | `String` | No |  |
| `vat` | `Hash` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.EuApI.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.EuApI.load({ "id" => "eu_ap_i_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EuApIEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## GlobalApIEntity

```ruby
global_ap_i = client.GlobalApI
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addition` | `String` | No |  |
| `admin1` | `String` | No |  |
| `admin2` | `String` | No |  |
| `admin3` | `String` | No |  |
| `bic` | `Hash` | No |  |
| `city` | `String` | No |  |
| `currency` | `Hash` | No |  |
| `date` | `String` | No |  |
| `dns` | `Hash` | No |  |
| `email` | `Hash` | No |  |
| `found` | `Integer` | No |  |
| `freeformaddress` | `String` | No |  |
| `from_currency` | `String` | No |  |
| `iban` | `Hash` | No |  |
| `ip` | `Hash` | No |  |
| `lat` | `Float` | No |  |
| `lei` | `Hash` | No |  |
| `letter` | `String` | No |  |
| `lon` | `Float` | No |  |
| `municipality` | `String` | No |  |
| `number` | `Integer` | No |  |
| `password` | `Hash` | No |  |
| `phone` | `Hash` | No |  |
| `population` | `Integer` | No |  |
| `postcode` | `String` | No |  |
| `province` | `String` | No |  |
| `province_code` | `String` | No |  |
| `score` | `Float` | No |  |
| `status` | `String` | No |  |
| `street` | `String` | No |  |
| `type` | `String` | No |  |
| `url` | `Hash` | No |  |
| `webrank` | `Hash` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.GlobalApI.create({
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.GlobalApI.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.GlobalApI.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `GlobalApIEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NetherlandsApIEntity

```ruby
netherlands_ap_i = client.NetherlandsApI
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active` | `Integer` | No |  |
| `addition` | `String` | No |  |
| `city` | `String` | No |  |
| `coc` | `String` | No |  |
| `construction_year` | `Integer` | No |  |
| `floor_area` | `Integer` | No |  |
| `freeformaddress` | `String` | No |  |
| `id` | `String` | No |  |
| `lat` | `Float` | No |  |
| `letter` | `String` | No |  |
| `lon` | `Float` | No |  |
| `municipality` | `String` | No |  |
| `name` | `String` | No |  |
| `number` | `String` | No |  |
| `postcode` | `String` | No |  |
| `province` | `String` | No |  |
| `province_code` | `String` | No |  |
| `purpose` | `String` | No |  |
| `street` | `String` | No |  |
| `type` | `String` | No |  |
| `vestiging` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NetherlandsApI.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NetherlandsApIEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = DutchCustomerDataSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

