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
| `buyer` | `string` | No |  |
| `buyer_country` | `string` | No |  |
| `contract_nature` | `string` | No |  |
| `html` | `string` | No |  |
| `id` | `string` | No |  |
| `link` | `string` | No |  |
| `notice_type` | `string` | No |  |
| `official_language` | `string` | No |  |
| `pdf` | `string` | No |  |
| `place_of_performance` | `string` | No |  |
| `procedure_type` | `string` | No |  |
| `publication_date` | `string` | No |  |
| `status` | `string` | No |  |
| `title` | `string` | No |  |
| `vat` | `table` | No |  |

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
| `admin1` | `string` | No |  |
| `admin2` | `string` | No |  |
| `admin3` | `string` | No |  |
| `bic` | `table` | No |  |
| `city` | `string` | No |  |
| `currency` | `table` | No |  |
| `date` | `string` | No |  |
| `dns` | `table` | No |  |
| `email` | `table` | No |  |
| `found` | `number` | No |  |
| `freeformaddress` | `string` | No |  |
| `from_currency` | `string` | No |  |
| `iban` | `table` | No |  |
| `ip` | `table` | No |  |
| `lat` | `number` | No |  |
| `lei` | `table` | No |  |
| `letter` | `string` | No |  |
| `lon` | `number` | No |  |
| `municipality` | `string` | No |  |
| `number` | `number` | No |  |
| `password` | `table` | No |  |
| `phone` | `table` | No |  |
| `population` | `number` | No |  |
| `postcode` | `string` | No |  |
| `province` | `string` | No |  |
| `province_code` | `string` | No |  |
| `score` | `number` | No |  |
| `status` | `string` | No |  |
| `street` | `string` | No |  |
| `type` | `string` | No |  |
| `url` | `table` | No |  |
| `webrank` | `table` | No |  |

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

