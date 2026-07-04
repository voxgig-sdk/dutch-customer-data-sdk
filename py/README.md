# DutchCustomerData Python SDK



The Python SDK for the DutchCustomerData API — an entity-oriented client following Pythonic conventions.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/dutch-customer-data-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
from dutchcustomerdata_sdk import DutchCustomerDataSDK

client = DutchCustomerDataSDK()
```

### 2. List euapi records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    euapis = client.EuApI().list({})
    for euapi in euapis:
        print(euapi)
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load an euapi

`load()` returns the bare record (a `dict`) and raises on error.

```python
try:
    euapi = client.EuApI().load({"id": "example_id"})
    print(euapi)
except Exception as err:
    print(f"load failed: {err}")
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    print(result["err"])     # error value
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = DutchCustomerDataSDK.test()

# Entity ops return the bare record and raise on error.
euapi = client.EuApI().load({"id": "test01"})
# euapi contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = DutchCustomerDataSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### DutchCustomerDataSDK

```python
from dutchcustomerdata_sdk import DutchCustomerDataSDK

client = DutchCustomerDataSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = DutchCustomerDataSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### DutchCustomerDataSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
| `EuApI` | `(data) -> EuApIEntity` | Create an EuApI entity instance. |
| `GlobalApI` | `(data) -> GlobalApIEntity` | Create a GlobalApI entity instance. |
| `NetherlandsApI` | `(data) -> NetherlandsApIEntity` | Create a NetherlandsApI entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the bare result data (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

### Entities

#### EuApI

| Field | Description |
| --- | --- |
| `buyer` |  |
| `buyer_country` |  |
| `contract_nature` |  |
| `html` |  |
| `id` |  |
| `link` |  |
| `notice_type` |  |
| `official_language` |  |
| `pdf` |  |
| `place_of_performance` |  |
| `procedure_type` |  |
| `publication_date` |  |
| `status` |  |
| `title` |  |
| `vat` |  |

Operations: List, Load.

API path: `/tender`

#### GlobalApI

| Field | Description |
| --- | --- |
| `addition` |  |
| `admin1` |  |
| `admin2` |  |
| `admin3` |  |
| `bic` |  |
| `city` |  |
| `currency` |  |
| `date` |  |
| `dns` |  |
| `email` |  |
| `found` |  |
| `freeformaddress` |  |
| `from_currency` |  |
| `iban` |  |
| `ip` |  |
| `lat` |  |
| `lei` |  |
| `letter` |  |
| `lon` |  |
| `municipality` |  |
| `number` |  |
| `password` |  |
| `phone` |  |
| `population` |  |
| `postcode` |  |
| `province` |  |
| `province_code` |  |
| `score` |  |
| `status` |  |
| `street` |  |
| `type` |  |
| `url` |  |
| `webrank` |  |

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

Create an instance: `eu_ap_i = client.EuApI()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `buyer` | ``$STRING`` |  |
| `buyer_country` | ``$STRING`` |  |
| `contract_nature` | ``$STRING`` |  |
| `html` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `link` | ``$STRING`` |  |
| `notice_type` | ``$STRING`` |  |
| `official_language` | ``$STRING`` |  |
| `pdf` | ``$STRING`` |  |
| `place_of_performance` | ``$STRING`` |  |
| `procedure_type` | ``$STRING`` |  |
| `publication_date` | ``$STRING`` |  |
| `status` | ``$STRING`` |  |
| `title` | ``$STRING`` |  |
| `vat` | ``$OBJECT`` |  |

#### Example: Load

```python
eu_ap_i = client.EuApI().load({"id": "eu_ap_i_id"})
```

#### Example: List

```python
eu_ap_is = client.EuApI().list({})
```


### GlobalApI

Create an instance: `global_ap_i = client.GlobalApI()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addition` | ``$STRING`` |  |
| `admin1` | ``$STRING`` |  |
| `admin2` | ``$STRING`` |  |
| `admin3` | ``$STRING`` |  |
| `bic` | ``$OBJECT`` |  |
| `city` | ``$STRING`` |  |
| `currency` | ``$OBJECT`` |  |
| `date` | ``$STRING`` |  |
| `dns` | ``$OBJECT`` |  |
| `email` | ``$OBJECT`` |  |
| `found` | ``$INTEGER`` |  |
| `freeformaddress` | ``$STRING`` |  |
| `from_currency` | ``$STRING`` |  |
| `iban` | ``$OBJECT`` |  |
| `ip` | ``$OBJECT`` |  |
| `lat` | ``$NUMBER`` |  |
| `lei` | ``$OBJECT`` |  |
| `letter` | ``$STRING`` |  |
| `lon` | ``$NUMBER`` |  |
| `municipality` | ``$STRING`` |  |
| `number` | ``$INTEGER`` |  |
| `password` | ``$OBJECT`` |  |
| `phone` | ``$OBJECT`` |  |
| `population` | ``$INTEGER`` |  |
| `postcode` | ``$STRING`` |  |
| `province` | ``$STRING`` |  |
| `province_code` | ``$STRING`` |  |
| `score` | ``$NUMBER`` |  |
| `status` | ``$STRING`` |  |
| `street` | ``$STRING`` |  |
| `type` | ``$STRING`` |  |
| `url` | ``$OBJECT`` |  |
| `webrank` | ``$OBJECT`` |  |

#### Example: Load

```python
global_ap_i = client.GlobalApI().load({"id": "global_ap_i_id"})
```

#### Example: List

```python
global_ap_is = client.GlobalApI().list({})
```

#### Example: Create

```python
global_ap_i = client.GlobalApI().create({
})
```


### NetherlandsApI

Create an instance: `netherlands_ap_i = client.NetherlandsApI()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `active` | ``$INTEGER`` |  |
| `addition` | ``$STRING`` |  |
| `city` | ``$STRING`` |  |
| `coc` | ``$STRING`` |  |
| `construction_year` | ``$INTEGER`` |  |
| `floor_area` | ``$INTEGER`` |  |
| `freeformaddress` | ``$STRING`` |  |
| `id` | ``$STRING`` |  |
| `lat` | ``$NUMBER`` |  |
| `letter` | ``$STRING`` |  |
| `lon` | ``$NUMBER`` |  |
| `municipality` | ``$STRING`` |  |
| `name` | ``$STRING`` |  |
| `number` | ``$STRING`` |  |
| `postcode` | ``$STRING`` |  |
| `province` | ``$STRING`` |  |
| `province_code` | ``$STRING`` |  |
| `purpose` | ``$STRING`` |  |
| `street` | ``$STRING`` |  |
| `type` | ``$STRING`` |  |
| `vestiging` | ``$STRING`` |  |

#### Example: List

```python
netherlands_ap_is = client.NetherlandsApI().list({})
```


## Explanation

### The operation pipeline

Every entity operation (load, list, create, update, remove) follows a
six-stage pipeline. Each stage fires a feature hook before executing:

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

If any stage returns an error, the pipeline short-circuits and the
error is returned to the caller as the second element in the return tuple.

### Features and hooks

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── dutchcustomerdata_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`dutchcustomerdata_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```python
euapi = client.EuApI()
euapi.load({"id": "example_id"})

# euapi.data_get() now returns the loaded euapi data
# euapi.match_get() returns the last match criteria
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
