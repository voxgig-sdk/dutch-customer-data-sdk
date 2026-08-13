# DutchCustomerData Lua SDK Reference

Complete API reference for the DutchCustomerData Lua SDK.


## DutchCustomerDataSDK

### Constructor

```lua
local sdk = require("dutch-customer-data_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `EuApI(data)`

Create a new `EuApI` entity instance. Pass `nil` for no initial data.

#### `GlobalApI(data)`

Create a new `GlobalApI` entity instance. Pass `nil` for no initial data.

#### `NetherlandsApI(data)`

Create a new `NetherlandsApI` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## EuApIEntity

```lua
local eu_ap_i = client:EuApI(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active` | `number` | No |  |
| `address` | `string` | No |  |
| `buyer` | `string` | No |  |
| `buyer_country` | `string` | No |  |
| `city` | `string` | No |  |
| `contract_nature` | `string` | No |  |
| `country` | `string` | No |  |
| `html` | `string` | No |  |
| `id` | `string` | No |  |
| `link` | `string` | No |  |
| `name` | `string` | No |  |
| `notice_type` | `string` | No |  |
| `official_language` | `string` | No |  |
| `pdf` | `string` | No |  |
| `place_of_performance` | `string` | No |  |
| `postcode` | `string` | No |  |
| `procedure_type` | `string` | No |  |
| `publication_date` | `string` | No |  |
| `response_date` | `string` | No |  |
| `title` | `string` | No |  |
| `vat` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EuApI():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EuApI():load({ id = "eu_ap_i_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EuApIEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## GlobalApIEntity

```lua
local global_ap_i = client:GlobalApI(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addition` | `string` | No |  |
| `address` | `string` | No |  |
| `admin1` | `string` | No |  |
| `admin2` | `string` | No |  |
| `admin3` | `string` | No |  |
| `bank` | `string` | No |  |
| `bic` | `string` | No |  |
| `browser` | `string` | No |  |
| `builtwith` | `number` | No |  |
| `carrier` | `string` | No |  |
| `city` | `string` | No |  |
| `cloudflare` | `number` | No |  |
| `commoncrawl` | `number` | No |  |
| `content_length` | `number` | No |  |
| `content_type` | `string` | No |  |
| `country` | `string` | No |  |
| `country_code` | `string` | No |  |
| `crux` | `number` | No |  |
| `device_family` | `string` | No |  |
| `device_name` | `string` | No |  |
| `device_type` | `string` | No |  |
| `disposable` | `number` | No |  |
| `dns_a` | `table` | No |  |
| `dns_mx` | `table` | No |  |
| `dns_ns` | `table` | No |  |
| `dns_soa` | `table` | No |  |
| `dns_txt` | `table` | No |  |
| `dns_www_a` | `table` | No |  |
| `dnsserver` | `string` | No |  |
| `domain` | `string` | No |  |
| `domcop` | `number` | No |  |
| `email` | `string` | No |  |
| `found` | `number` | No |  |
| `free` | `number` | No |  |
| `freeformaddress` | `string` | No |  |
| `host` | `string` | No |  |
| `host_type` | `string` | No |  |
| `hostio` | `number` | No |  |
| `http_code` | `number` | No |  |
| `iban` | `string` | No |  |
| `iban_human` | `string` | No |  |
| `int` | `string` | No |  |
| `international` | `string` | No |  |
| `ip` | `string` | No |  |
| `ipint` | `number` | No |  |
| `ismobile` | `number` | No |  |
| `lat` | `number` | No |  |
| `lei` | `string` | No |  |
| `letter` | `string` | No |  |
| `local_id` | `string` | No |  |
| `lon` | `number` | No |  |
| `mailserver` | `string` | No |  |
| `majestic` | `number` | No |  |
| `message` | `string` | No |  |
| `municipality` | `string` | No |  |
| `mx_host` | `string` | No |  |
| `mx_ip` | `string` | No |  |
| `name` | `string` | No |  |
| `national` | `string` | No |  |
| `number` | `number` | No |  |
| `ocid` | `string` | No |  |
| `pagerank` | `number` | No |  |
| `platform` | `string` | No |  |
| `population` | `number` | No |  |
| `postcode` | `string` | No |  |
| `province` | `string` | No |  |
| `province_code` | `string` | No |  |
| `redirect_count` | `number` | No |  |
| `region` | `string` | No |  |
| `register_id` | `string` | No |  |
| `renewal_date` | `string` | No |  |
| `score` | `number` | No |  |
| `sepa` | `number` | No |  |
| `spf` | `string` | No |  |
| `status` | `string` | No |  |
| `street` | `string` | No |  |
| `success` | `number` | No |  |
| `swift` | `number` | No |  |
| `tranco` | `number` | No |  |
| `type` | `string` | No |  |
| `umbrella` | `number` | No |  |
| `url` | `string` | No |  |
| `user` | `string` | No |  |
| `user_agent` | `string` | No |  |
| `valid` | `number` | No |  |
| `verified` | `boolean` | No |  |
| `verified_checksum` | `boolean` | No |  |
| `webrank` | `number` | No |  |
| `wrong_email` | `number` | No |  |
| `wrong_format` | `number` | No |  |
| `wrong_password` | `number` | No |  |
| `wrong_phone` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:GlobalApI():create({
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:GlobalApI():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:GlobalApI():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GlobalApIEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NetherlandsApIEntity

```lua
local netherlands_ap_i = client:NetherlandsApI(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active` | `number` | No |  |
| `addition` | `string` | No |  |
| `city` | `string` | No |  |
| `coc` | `string` | No |  |
| `construction_year` | `number` | No |  |
| `floor_area` | `number` | No |  |
| `freeformaddress` | `string` | No |  |
| `id` | `string` | No |  |
| `lat` | `number` | No |  |
| `letter` | `string` | No |  |
| `lon` | `number` | No |  |
| `municipality` | `string` | No |  |
| `name` | `string` | No |  |
| `number` | `string` | No |  |
| `postcode` | `string` | No |  |
| `province` | `string` | No |  |
| `province_code` | `string` | No |  |
| `purpose` | `string` | No |  |
| `street` | `string` | No |  |
| `type` | `string` | No |  |
| `vestiging` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NetherlandsApI():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NetherlandsApIEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

