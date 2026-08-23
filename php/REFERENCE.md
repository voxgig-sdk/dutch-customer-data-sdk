# DutchCustomerData PHP SDK Reference

Complete API reference for the DutchCustomerData PHP SDK.


## DutchCustomerDataSDK

### Constructor

```php
require_once __DIR__ . '/dutchcustomerdata_sdk.php';

$client = new DutchCustomerDataSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DutchCustomerDataSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = DutchCustomerDataSDK::test();
```


### Instance Methods

#### `EuApI($data = null)`

Create a new `EuApIEntity` instance. Pass `null` for no initial data.

#### `GlobalApI($data = null)`

Create a new `GlobalApIEntity` instance. Pass `null` for no initial data.

#### `NetherlandsApI($data = null)`

Create a new `NetherlandsApIEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): DutchCustomerDataUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## EuApIEntity

```php
$eu_ap_i = $client->EuApI();
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

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EuApI()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EuApI()->load(["id" => "eu_ap_i_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EuApIEntity`

Create a new `EuApIEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GlobalApIEntity

```php
$global_ap_i = $client->GlobalApI();
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
| `carrier` | `string` | No | Carrier name |
| `city` | `string` | No |  |
| `cloudflare` | `int` | No |  |
| `commoncrawl` | `int` | No |  |
| `content_length` | `int` | No |  |
| `content_type` | `string` | No |  |
| `country` | `string` | No | ISO country code |
| `country_code` | `string` | No |  |
| `crux` | `int` | No |  |
| `device_family` | `string` | No |  |
| `device_name` | `string` | No |  |
| `device_type` | `string` | No |  |
| `disposable` | `int` | No |  |
| `dns_a` | `array` | No |  |
| `dns_mx` | `array` | No |  |
| `dns_ns` | `array` | No |  |
| `dns_soa` | `array` | No |  |
| `dns_txt` | `array` | No |  |
| `dns_www_a` | `array` | No |  |
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
| `int` | `string` | No | International format without plus sign |
| `international` | `string` | No | International formatted phone number |
| `ip` | `string` | No |  |
| `ipint` | `int` | No |  |
| `ismobile` | `int` | No | 1 if mobile, 0 if not |
| `lat` | `float` | No |  |
| `lei` | `string` | No |  |
| `letter` | `string` | No |  |
| `local_id` | `string` | No |  |
| `lon` | `float` | No |  |
| `mailserver` | `string` | No |  |
| `majestic` | `int` | No |  |
| `message` | `string` | No |  |
| `municipality` | `string` | No |  |
| `mx_host` | `string` | No |  |
| `mx_ip` | `string` | No |  |
| `name` | `string` | No |  |
| `national` | `string` | No | National formatted phone number |
| `number` | `int` | No |  |
| `ocid` | `string` | No |  |
| `pagerank` | `int` | No |  |
| `platform` | `string` | No |  |
| `population` | `int` | No |  |
| `postcode` | `string` | No |  |
| `province` | `string` | No |  |
| `province_code` | `string` | No |  |
| `redirect_count` | `int` | No |  |
| `region` | `string` | No | Geographic region |
| `register_id` | `string` | No |  |
| `renewal_date` | `string` | No |  |
| `score` | `float` | No |  |
| `sepa` | `int` | No |  |
| `spf` | `string` | No |  |
| `status` | `string` | No |  |
| `street` | `string` | No |  |
| `success` | `int` | No | 1 if successful, 0 if not |
| `swift` | `int` | No |  |
| `tranco` | `int` | No |  |
| `type` | `string` | No |  |
| `umbrella` | `int` | No |  |
| `url` | `string` | No |  |
| `user` | `string` | No |  |
| `user_agent` | `string` | No |  |
| `valid` | `int` | No | 1 if valid, 0 if not |
| `verified` | `bool` | No |  |
| `verified_checksum` | `bool` | No |  |
| `webrank` | `int` | No |  |
| `wrong_email` | `int` | No |  |
| `wrong_format` | `int` | No |  |
| `wrong_password` | `int` | No |  |
| `wrong_phone` | `int` | No | 1 if wrong, 0 if correct |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->GlobalApI()->create([
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->GlobalApI()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->GlobalApI()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GlobalApIEntity`

Create a new `GlobalApIEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NetherlandsApIEntity

```php
$netherlands_ap_i = $client->NetherlandsApI();
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
| `lat` | `float` | No |  |
| `letter` | `string` | No |  |
| `lon` | `float` | No |  |
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

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NetherlandsApI()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NetherlandsApIEntity`

Create a new `NetherlandsApIEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new DutchCustomerDataSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

