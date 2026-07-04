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
local result, err = client:GlobalApI():load({ id = "global_ap_i_id" })
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

