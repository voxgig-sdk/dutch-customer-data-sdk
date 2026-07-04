# DutchCustomerData Python SDK Reference

Complete API reference for the DutchCustomerData Python SDK.


## DutchCustomerDataSDK

### Constructor

```python
from dutch-customer-data_sdk import DutchCustomerDataSDK

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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.eu_ap_i.list({})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.eu_ap_i.load({"id": "eu_ap_i_id"})
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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.global_ap_i.create({
})
```

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.global_ap_i.list({})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.global_ap_i.load({"id": "global_ap_i_id"})
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

#### `list(reqmatch, ctrl=None) -> list`

List entities matching the given criteria. Returns a list and raises on error.

```python
results = client.netherlands_ap_i.list({})
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

