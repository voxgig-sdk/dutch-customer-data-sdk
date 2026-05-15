# DutchCustomerData PHP SDK Reference

Complete API reference for the DutchCustomerData PHP SDK.


## DutchCustomerDataSDK

### Constructor

```php
require_once __DIR__ . '/dutch-customer-data_sdk.php';

$client = new DutchCustomerDataSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
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

#### `optionsMap(): array`

Return a deep copy of the current SDK options.

#### `getUtility(): ProjectNameUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. Returns `[$result, $err]`.

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

**Returns:** `array [$result, $err]`

#### `prepare(array $fetchargs = []): array`

Prepare a fetch definition without sending the request. Returns `[$fetchdef, $err]`.


---

## EuApIEntity

```php
$eu_ap_i = $client->EuApI();
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

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->EuApI()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->EuApI()->load(["id" => "eu_ap_i_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): EuApIEntity`

Create a new `EuApIEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## GlobalApIEntity

```php
$global_ap_i = $client->GlobalApI();
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

#### `create(array $reqdata, ?array $ctrl = null): array`

Create a new entity with the given data.

```php
[$result, $err] = $client->GlobalApI()->create([
]);
```

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->GlobalApI()->list([]);
```

#### `load(array $reqmatch, ?array $ctrl = null): array`

Load a single entity matching the given criteria.

```php
[$result, $err] = $client->GlobalApI()->load(["id" => "global_ap_i_id"]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): GlobalApIEntity`

Create a new `GlobalApIEntity` instance with the same client and
options.

#### `getName(): string`

Return the entity name.


---

## NetherlandsApIEntity

```php
$netherlands_ap_i = $client->NetherlandsApI();
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

#### `list(array $reqmatch, ?array $ctrl = null): array`

List entities matching the given criteria. Returns an array.

```php
[$results, $err] = $client->NetherlandsApI()->list([]);
```

### Common Methods

#### `dataGet(): array`

Get the entity data. Returns a copy of the current data.

#### `dataSet($data): void`

Set the entity data.

#### `matchGet(): array`

Get the entity match criteria.

#### `matchSet($match): void`

Set the entity match criteria.

#### `make(): NetherlandsApIEntity`

Create a new `NetherlandsApIEntity` instance with the same client and
options.

#### `getName(): string`

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

