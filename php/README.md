# DutchCustomerData PHP SDK



The PHP SDK for the DutchCustomerData API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->EuApI()` — with named operations (`list`/`load`/`create`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/dutch-customer-data-sdk/releases](https://github.com/voxgig-sdk/dutch-customer-data-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'dutchcustomerdata_sdk.php';

$client = new DutchCustomerDataSDK();
```

### 2. List euapi records

```php
try {
    // list() returns an array of EuApI records — iterate directly.
    $euapis = $client->EuApI()->list();
    foreach ($euapis as $item) {
        echo $item["id"] . " " . $item["active"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load an euapi

```php
try {
    // load() returns the ENTITY — call data_get() for the EuApI record (throws on error).
    $euapi = $client->EuApI()->load(["vat" => "example_vat"]);
    print_r($euapi);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $euapis = $client->EuApI()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required:

```php
$client = DutchCustomerDataSDK::test();

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
$euapi = $client->EuApI()->list();
print_r($euapi);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new DutchCustomerDataSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
DUTCH_CUSTOMER_DATA_TEST_LIVE=TRUE
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### DutchCustomerDataSDK

```php
require_once 'dutchcustomerdata_sdk.php';
$client = new DutchCustomerDataSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = DutchCustomerDataSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### DutchCustomerDataSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `EuApI` | `($data): EuApIEntity` | Create an EuApI entity instance. |
| `GlobalApI` | `($data): GlobalApIEntity` | Create a GlobalApI entity instance. |
| `NetherlandsApI` | `($data): NetherlandsApIEntity` | Create a NetherlandsApI entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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
| `carrier` | Carrier name |
| `city` |  |
| `cloudflare` |  |
| `commoncrawl` |  |
| `content_length` |  |
| `content_type` |  |
| `country` | ISO country code |
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
| `int` | International format without plus sign |
| `international` | International formatted phone number |
| `ip` |  |
| `ipint` |  |
| `ismobile` | 1 if mobile, 0 if not |
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
| `national` | National formatted phone number |
| `number` |  |
| `ocid` |  |
| `pagerank` |  |
| `platform` |  |
| `population` |  |
| `postcode` |  |
| `province` |  |
| `province_code` |  |
| `redirect_count` |  |
| `region` | Geographic region |
| `register_id` |  |
| `renewal_date` |  |
| `score` |  |
| `sepa` |  |
| `spf` |  |
| `status` |  |
| `street` |  |
| `success` | 1 if successful, 0 if not |
| `swift` |  |
| `tranco` |  |
| `type` |  |
| `umbrella` |  |
| `url` |  |
| `user` |  |
| `user_agent` |  |
| `valid` | 1 if valid, 0 if not |
| `verified` |  |
| `verified_checksum` |  |
| `webrank` |  |
| `wrong_email` |  |
| `wrong_format` |  |
| `wrong_password` |  |
| `wrong_phone` | 1 if wrong, 0 if correct |

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

Create an instance: `$eu_ap_i = $client->EuApI();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `active` | `int` |  |
| `address` | `string` |  |
| `buyer` | `string` |  |
| `buyer_country` | `string` |  |
| `city` | `string` |  |
| `contract_nature` | `string` |  |
| `country` | `string` |  |
| `html` | `string` |  |
| `id` | `string` |  |
| `link` | `string` |  |
| `name` | `string` |  |
| `notice_type` | `string` |  |
| `official_language` | `string` |  |
| `pdf` | `string` |  |
| `place_of_performance` | `string` |  |
| `postcode` | `string` |  |
| `procedure_type` | `string` |  |
| `publication_date` | `string` |  |
| `response_date` | `string` |  |
| `title` | `string` |  |
| `vat` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the EuApI record (throws on error).
$eu_ap_i = $client->EuApI()->load(["vat" => "vat"]);
```

#### Example: List

```php
// list() returns an array of EuApI records (throws on error).
$eu_ap_is = $client->EuApI()->list();
```


### GlobalApI

Create an instance: `$global_ap_i = $client->GlobalApI();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addition` | `string` |  |
| `address` | `string` |  |
| `admin1` | `string` |  |
| `admin2` | `string` |  |
| `admin3` | `string` |  |
| `bank` | `string` |  |
| `bic` | `string` |  |
| `browser` | `string` |  |
| `builtwith` | `int` |  |
| `carrier` | `string` | Carrier name |
| `city` | `string` |  |
| `cloudflare` | `int` |  |
| `commoncrawl` | `int` |  |
| `content_length` | `int` |  |
| `content_type` | `string` |  |
| `country` | `string` | ISO country code |
| `country_code` | `string` |  |
| `crux` | `int` |  |
| `device_family` | `string` |  |
| `device_name` | `string` |  |
| `device_type` | `string` |  |
| `disposable` | `int` |  |
| `dns_a` | `array` |  |
| `dns_mx` | `array` |  |
| `dns_ns` | `array` |  |
| `dns_soa` | `array` |  |
| `dns_txt` | `array` |  |
| `dns_www_a` | `array` |  |
| `dnsserver` | `string` |  |
| `domain` | `string` |  |
| `domcop` | `int` |  |
| `email` | `string` |  |
| `found` | `int` |  |
| `free` | `int` |  |
| `freeformaddress` | `string` |  |
| `host` | `string` |  |
| `host_type` | `string` |  |
| `hostio` | `int` |  |
| `http_code` | `int` |  |
| `iban` | `string` |  |
| `iban_human` | `string` |  |
| `int` | `string` | International format without plus sign |
| `international` | `string` | International formatted phone number |
| `ip` | `string` |  |
| `ipint` | `int` |  |
| `ismobile` | `int` | 1 if mobile, 0 if not |
| `lat` | `float` |  |
| `lei` | `string` |  |
| `letter` | `string` |  |
| `local_id` | `string` |  |
| `lon` | `float` |  |
| `mailserver` | `string` |  |
| `majestic` | `int` |  |
| `message` | `string` |  |
| `municipality` | `string` |  |
| `mx_host` | `string` |  |
| `mx_ip` | `string` |  |
| `name` | `string` |  |
| `national` | `string` | National formatted phone number |
| `number` | `int` |  |
| `ocid` | `string` |  |
| `pagerank` | `int` |  |
| `platform` | `string` |  |
| `population` | `int` |  |
| `postcode` | `string` |  |
| `province` | `string` |  |
| `province_code` | `string` |  |
| `redirect_count` | `int` |  |
| `region` | `string` | Geographic region |
| `register_id` | `string` |  |
| `renewal_date` | `string` |  |
| `score` | `float` |  |
| `sepa` | `int` |  |
| `spf` | `string` |  |
| `status` | `string` |  |
| `street` | `string` |  |
| `success` | `int` | 1 if successful, 0 if not |
| `swift` | `int` |  |
| `tranco` | `int` |  |
| `type` | `string` |  |
| `umbrella` | `int` |  |
| `url` | `string` |  |
| `user` | `string` |  |
| `user_agent` | `string` |  |
| `valid` | `int` | 1 if valid, 0 if not |
| `verified` | `bool` |  |
| `verified_checksum` | `bool` |  |
| `webrank` | `int` |  |
| `wrong_email` | `int` |  |
| `wrong_format` | `int` |  |
| `wrong_password` | `int` |  |
| `wrong_phone` | `int` | 1 if wrong, 0 if correct |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the GlobalApI record (throws on error).
$global_ap_i = $client->GlobalApI()->load();
```

#### Example: List

```php
// list() returns an array of GlobalApI records (throws on error).
$global_ap_is = $client->GlobalApI()->list();
```

#### Example: Create

```php
$global_ap_i = $client->GlobalApI()->create([
]);
```


### NetherlandsApI

Create an instance: `$netherlands_ap_i = $client->NetherlandsApI();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `active` | `int` |  |
| `addition` | `string` |  |
| `city` | `string` |  |
| `coc` | `string` |  |
| `construction_year` | `int` |  |
| `floor_area` | `int` |  |
| `freeformaddress` | `string` |  |
| `id` | `string` |  |
| `lat` | `float` |  |
| `letter` | `string` |  |
| `lon` | `float` |  |
| `municipality` | `string` |  |
| `name` | `string` |  |
| `number` | `string` |  |
| `postcode` | `string` |  |
| `province` | `string` |  |
| `province_code` | `string` |  |
| `purpose` | `string` |  |
| `street` | `string` |  |
| `type` | `string` |  |
| `vestiging` | `string` |  |

#### Example: List

```php
// list() returns an array of NetherlandsApI records (throws on error).
$netherlands_ap_is = $client->NetherlandsApI()->list();
```

## Features

This SDK ships 1 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`test`](#test) | In-memory mock transport for testing without a live server |

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.


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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── dutchcustomerdata_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`dutchcustomerdata_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$euapi = $client->EuApI();
$euapi->list();

// $euapi->data_get() now returns the euapi data from the last list
// $euapi->match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
