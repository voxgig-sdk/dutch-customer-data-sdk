# DutchCustomerData Python SDK Reference

Complete API reference for the DutchCustomerData Python SDK.


## DutchCustomerDataSDK

### Constructor

```python
from dutchcustomerdata_sdk import DutchCustomerDataSDK

client = DutchCustomerDataSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DutchCustomerDataSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = DutchCustomerDataSDK.test()
```


### Instance Methods

#### `EuApI(data=None)`

Create a new `EuApIEntity` instance. Pass `None` for no initial data.

#### `GlobalApI(data=None)`

Create a new `GlobalApIEntity` instance. Pass `None` for no initial data.

#### `NetherlandsApI(data=None)`

Create a new `NetherlandsApIEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## EuApIEntity

```python
eu_ap_i = client.EuApI()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `buyer` | `str` | No |  |
| `buyer_country` | `str` | No |  |
| `contract_nature` | `str` | No |  |
| `html` | `str` | No |  |
| `id` | `str` | No |  |
| `link` | `str` | No |  |
| `notice_type` | `str` | No |  |
| `official_language` | `str` | No |  |
| `pdf` | `str` | No |  |
| `place_of_performance` | `str` | No |  |
| `procedure_type` | `str` | No |  |
| `publication_date` | `str` | No |  |
| `status` | `str` | No |  |
| `title` | `str` | No |  |
| `vat` | `dict` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EuApI().list()
for eu_ap_i in results:
    print(eu_ap_i)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EuApI().load({"id": "eu_ap_i_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EuApIEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GlobalApIEntity

```python
global_ap_i = client.GlobalApI()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addition` | `str` | No |  |
| `admin1` | `str` | No |  |
| `admin2` | `str` | No |  |
| `admin3` | `str` | No |  |
| `bic` | `dict` | No |  |
| `city` | `str` | No |  |
| `currency` | `dict` | No |  |
| `date` | `str` | No |  |
| `dns` | `dict` | No |  |
| `email` | `dict` | No |  |
| `found` | `int` | No |  |
| `freeformaddress` | `str` | No |  |
| `from_currency` | `str` | No |  |
| `iban` | `dict` | No |  |
| `ip` | `dict` | No |  |
| `lat` | `float` | No |  |
| `lei` | `dict` | No |  |
| `letter` | `str` | No |  |
| `lon` | `float` | No |  |
| `municipality` | `str` | No |  |
| `number` | `int` | No |  |
| `password` | `dict` | No |  |
| `phone` | `dict` | No |  |
| `population` | `int` | No |  |
| `postcode` | `str` | No |  |
| `province` | `str` | No |  |
| `province_code` | `str` | No |  |
| `score` | `float` | No |  |
| `status` | `str` | No |  |
| `street` | `str` | No |  |
| `type` | `str` | No |  |
| `url` | `dict` | No |  |
| `webrank` | `dict` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.GlobalApI().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.GlobalApI().list()
for global_ap_i in results:
    print(global_ap_i)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.GlobalApI().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GlobalApIEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NetherlandsApIEntity

```python
netherlands_ap_i = client.NetherlandsApI()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active` | `int` | No |  |
| `addition` | `str` | No |  |
| `city` | `str` | No |  |
| `coc` | `str` | No |  |
| `construction_year` | `int` | No |  |
| `floor_area` | `int` | No |  |
| `freeformaddress` | `str` | No |  |
| `id` | `str` | No |  |
| `lat` | `float` | No |  |
| `letter` | `str` | No |  |
| `lon` | `float` | No |  |
| `municipality` | `str` | No |  |
| `name` | `str` | No |  |
| `number` | `str` | No |  |
| `postcode` | `str` | No |  |
| `province` | `str` | No |  |
| `province_code` | `str` | No |  |
| `purpose` | `str` | No |  |
| `street` | `str` | No |  |
| `type` | `str` | No |  |
| `vestiging` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NetherlandsApI().list()
for netherlands_ap_i in results:
    print(netherlands_ap_i)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NetherlandsApIEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = DutchCustomerDataSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

