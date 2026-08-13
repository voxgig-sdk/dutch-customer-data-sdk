# DutchCustomerData Golang SDK Reference

Complete API reference for the DutchCustomerData Golang SDK.


## DutchCustomerDataSDK

### Constructor

```go
func NewDutchCustomerDataSDK(options map[string]any) *DutchCustomerDataSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *DutchCustomerDataSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *DutchCustomerDataSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `EuApI(data map[string]any) DutchCustomerDataEntity`

Create a new `EuApI` entity instance. Pass `nil` for no initial data.

#### `GlobalApI(data map[string]any) DutchCustomerDataEntity`

Create a new `GlobalApI` entity instance. Pass `nil` for no initial data.

#### `NetherlandsApI(data map[string]any) DutchCustomerDataEntity`

Create a new `NetherlandsApI` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## EuApIEntity

```go
euApI := client.EuApI(nil)
fmt.Println(euApI.GetName()) // "eu_ap_i"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active` | `int` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EuApI(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EuApI(nil).Load(map[string]any{"id": "eu_ap_i_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EuApIEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GlobalApIEntity

```go
globalApI := client.GlobalApI(nil)
fmt.Println(globalApI.GetName()) // "global_ap_i"
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
| `builtwith` | `int` | No |  |
| `carrier` | `string` | No |  |
| `city` | `string` | No |  |
| `cloudflare` | `int` | No |  |
| `commoncrawl` | `int` | No |  |
| `content_length` | `int` | No |  |
| `content_type` | `string` | No |  |
| `country` | `string` | No |  |
| `country_code` | `string` | No |  |
| `crux` | `int` | No |  |
| `device_family` | `string` | No |  |
| `device_name` | `string` | No |  |
| `device_type` | `string` | No |  |
| `disposable` | `int` | No |  |
| `dns_a` | `[]any` | No |  |
| `dns_mx` | `[]any` | No |  |
| `dns_ns` | `[]any` | No |  |
| `dns_soa` | `[]any` | No |  |
| `dns_txt` | `[]any` | No |  |
| `dns_www_a` | `[]any` | No |  |
| `dnsserver` | `string` | No |  |
| `domain` | `string` | No |  |
| `domcop` | `int` | No |  |
| `email` | `string` | No |  |
| `found` | `int` | No |  |
| `free` | `int` | No |  |
| `freeformaddress` | `string` | No |  |
| `host` | `string` | No |  |
| `host_type` | `string` | No |  |
| `hostio` | `int` | No |  |
| `http_code` | `int` | No |  |
| `iban` | `string` | No |  |
| `iban_human` | `string` | No |  |
| `int` | `string` | No |  |
| `international` | `string` | No |  |
| `ip` | `string` | No |  |
| `ipint` | `int` | No |  |
| `ismobile` | `int` | No |  |
| `lat` | `float64` | No |  |
| `lei` | `string` | No |  |
| `letter` | `string` | No |  |
| `local_id` | `string` | No |  |
| `lon` | `float64` | No |  |
| `mailserver` | `string` | No |  |
| `majestic` | `int` | No |  |
| `message` | `string` | No |  |
| `municipality` | `string` | No |  |
| `mx_host` | `string` | No |  |
| `mx_ip` | `string` | No |  |
| `name` | `string` | No |  |
| `national` | `string` | No |  |
| `number` | `int` | No |  |
| `ocid` | `string` | No |  |
| `pagerank` | `int` | No |  |
| `platform` | `string` | No |  |
| `population` | `int` | No |  |
| `postcode` | `string` | No |  |
| `province` | `string` | No |  |
| `province_code` | `string` | No |  |
| `redirect_count` | `int` | No |  |
| `region` | `string` | No |  |
| `register_id` | `string` | No |  |
| `renewal_date` | `string` | No |  |
| `score` | `float64` | No |  |
| `sepa` | `int` | No |  |
| `spf` | `string` | No |  |
| `status` | `string` | No |  |
| `street` | `string` | No |  |
| `success` | `int` | No |  |
| `swift` | `int` | No |  |
| `tranco` | `int` | No |  |
| `type` | `string` | No |  |
| `umbrella` | `int` | No |  |
| `url` | `string` | No |  |
| `user` | `string` | No |  |
| `user_agent` | `string` | No |  |
| `valid` | `int` | No |  |
| `verified` | `bool` | No |  |
| `verified_checksum` | `bool` | No |  |
| `webrank` | `int` | No |  |
| `wrong_email` | `int` | No |  |
| `wrong_format` | `int` | No |  |
| `wrong_password` | `int` | No |  |
| `wrong_phone` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.GlobalApI(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.GlobalApI(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.GlobalApI(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GlobalApIEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NetherlandsApIEntity

```go
netherlandsApI := client.NetherlandsApI(nil)
fmt.Println(netherlandsApI.GetName()) // "netherlands_ap_i"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active` | `int` | No |  |
| `addition` | `string` | No |  |
| `city` | `string` | No |  |
| `coc` | `string` | No |  |
| `construction_year` | `int` | No |  |
| `floor_area` | `int` | No |  |
| `freeformaddress` | `string` | No |  |
| `id` | `string` | No |  |
| `lat` | `float64` | No |  |
| `letter` | `string` | No |  |
| `lon` | `float64` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NetherlandsApI(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NetherlandsApIEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewDutchCustomerDataSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

