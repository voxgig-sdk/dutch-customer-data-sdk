# DutchCustomerData Ruby SDK Reference

Complete API reference for the DutchCustomerData Ruby SDK.


## DutchCustomerDataSDK

### Constructor

```ruby
require_relative 'dutch-customer-data_sdk'

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
eu_ap_i = client.eu_ap_i
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `buyer` | ``$STRING`` | No |  |
| `buyer_country` | ``$STRING`` | No |  |
| `contract_nature` | ``$STRING`` | No |  |
| `html` | ``$STRING`` | No |  |
| `id` | ``$STRING`` | No |  |
| `link` | ``$STRING`` | No |  |
| `notice_type` | ``$STRING`` | No |  |
| `official_language` | ``$STRING`` | No |  |
| `pdf` | ``$STRING`` | No |  |
| `place_of_performance` | ``$STRING`` | No |  |
| `procedure_type` | ``$STRING`` | No |  |
| `publication_date` | ``$STRING`` | No |  |
| `status` | ``$STRING`` | No |  |
| `title` | ``$STRING`` | No |  |
| `vat` | ``$OBJECT`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.eu_ap_i.list(nil)
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.eu_ap_i.load({ "id" => "eu_ap_i_id" })
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
global_ap_i = client.global_ap_i
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addition` | ``$STRING`` | No |  |
| `admin1` | ``$STRING`` | No |  |
| `admin2` | ``$STRING`` | No |  |
| `admin3` | ``$STRING`` | No |  |
| `bic` | ``$OBJECT`` | No |  |
| `city` | ``$STRING`` | No |  |
| `currency` | ``$OBJECT`` | No |  |
| `date` | ``$STRING`` | No |  |
| `dns` | ``$OBJECT`` | No |  |
| `email` | ``$OBJECT`` | No |  |
| `found` | ``$INTEGER`` | No |  |
| `freeformaddress` | ``$STRING`` | No |  |
| `from_currency` | ``$STRING`` | No |  |
| `iban` | ``$OBJECT`` | No |  |
| `ip` | ``$OBJECT`` | No |  |
| `lat` | ``$NUMBER`` | No |  |
| `lei` | ``$OBJECT`` | No |  |
| `letter` | ``$STRING`` | No |  |
| `lon` | ``$NUMBER`` | No |  |
| `municipality` | ``$STRING`` | No |  |
| `number` | ``$INTEGER`` | No |  |
| `password` | ``$OBJECT`` | No |  |
| `phone` | ``$OBJECT`` | No |  |
| `population` | ``$INTEGER`` | No |  |
| `postcode` | ``$STRING`` | No |  |
| `province` | ``$STRING`` | No |  |
| `province_code` | ``$STRING`` | No |  |
| `score` | ``$NUMBER`` | No |  |
| `status` | ``$STRING`` | No |  |
| `street` | ``$STRING`` | No |  |
| `type` | ``$STRING`` | No |  |
| `url` | ``$OBJECT`` | No |  |
| `webrank` | ``$OBJECT`` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.global_ap_i.create({
})
```

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.global_ap_i.list(nil)
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.global_ap_i.load({ "id" => "global_ap_i_id" })
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
netherlands_ap_i = client.netherlands_ap_i
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active` | ``$INTEGER`` | No |  |
| `addition` | ``$STRING`` | No |  |
| `city` | ``$STRING`` | No |  |
| `coc` | ``$STRING`` | No |  |
| `construction_year` | ``$INTEGER`` | No |  |
| `floor_area` | ``$INTEGER`` | No |  |
| `freeformaddress` | ``$STRING`` | No |  |
| `id` | ``$STRING`` | No |  |
| `lat` | ``$NUMBER`` | No |  |
| `letter` | ``$STRING`` | No |  |
| `lon` | ``$NUMBER`` | No |  |
| `municipality` | ``$STRING`` | No |  |
| `name` | ``$STRING`` | No |  |
| `number` | ``$STRING`` | No |  |
| `postcode` | ``$STRING`` | No |  |
| `province` | ``$STRING`` | No |  |
| `province_code` | ``$STRING`` | No |  |
| `purpose` | ``$STRING`` | No |  |
| `street` | ``$STRING`` | No |  |
| `type` | ``$STRING`` | No |  |
| `vestiging` | ``$STRING`` | No |  |

### Operations

#### `list(reqmatch, ctrl = nil) -> Array`

List entities matching the given criteria. Returns an array. Raises on error.

```ruby
results = client.netherlands_ap_i.list(nil)
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

