# DutchCustomerData TypeScript SDK Reference

Complete API reference for the DutchCustomerData TypeScript SDK.


## DutchCustomerDataSDK

### Constructor

```ts
new DutchCustomerDataSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `DutchCustomerDataSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = DutchCustomerDataSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `DutchCustomerDataSDK` instance in test mode.


### Instance Methods

#### `EuApI(data?: object)`

Create a new `EuApI` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EuApIEntity` instance.

#### `GlobalApI(data?: object)`

Create a new `GlobalApI` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GlobalApIEntity` instance.

#### `NetherlandsApI(data?: object)`

Create a new `NetherlandsApI` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NetherlandsApIEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `DutchCustomerDataSDK.test()`.

**Returns:** `DutchCustomerDataSDK` instance in test mode.


---

## EuApIEntity

```ts
const eu_ap_i = client.EuApI()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EuApI().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EuApI().load({ id: 'eu_ap_i_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EuApIEntity` instance with the same client and
options.

#### `client()`

Return the parent `DutchCustomerDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GlobalApIEntity

```ts
const global_ap_i = client.GlobalApI()
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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.GlobalApI().create({
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.GlobalApI().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.GlobalApI().load({ id: 'global_ap_i_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GlobalApIEntity` instance with the same client and
options.

#### `client()`

Return the parent `DutchCustomerDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NetherlandsApIEntity

```ts
const netherlands_ap_i = client.NetherlandsApI()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NetherlandsApI().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NetherlandsApIEntity` instance with the same client and
options.

#### `client()`

Return the parent `DutchCustomerDataSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new DutchCustomerDataSDK({
  feature: {
    test: { active: true },
  }
})
```

