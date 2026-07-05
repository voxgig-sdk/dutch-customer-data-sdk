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
| `vat` | `Record<string, any>` | No |  |

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
| `addition` | `string` | No |  |
| `admin1` | `string` | No |  |
| `admin2` | `string` | No |  |
| `admin3` | `string` | No |  |
| `bic` | `Record<string, any>` | No |  |
| `city` | `string` | No |  |
| `currency` | `Record<string, any>` | No |  |
| `date` | `string` | No |  |
| `dns` | `Record<string, any>` | No |  |
| `email` | `Record<string, any>` | No |  |
| `found` | `number` | No |  |
| `freeformaddress` | `string` | No |  |
| `from_currency` | `string` | No |  |
| `iban` | `Record<string, any>` | No |  |
| `ip` | `Record<string, any>` | No |  |
| `lat` | `number` | No |  |
| `lei` | `Record<string, any>` | No |  |
| `letter` | `string` | No |  |
| `lon` | `number` | No |  |
| `municipality` | `string` | No |  |
| `number` | `number` | No |  |
| `password` | `Record<string, any>` | No |  |
| `phone` | `Record<string, any>` | No |  |
| `population` | `number` | No |  |
| `postcode` | `string` | No |  |
| `province` | `string` | No |  |
| `province_code` | `string` | No |  |
| `score` | `number` | No |  |
| `status` | `string` | No |  |
| `street` | `string` | No |  |
| `type` | `string` | No |  |
| `url` | `Record<string, any>` | No |  |
| `webrank` | `Record<string, any>` | No |  |

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
const result = await client.GlobalApI().load()
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

