# DutchCustomerData TypeScript SDK



The TypeScript SDK for the DutchCustomerData API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.EuApI()` — each with a small set of operations (`list`, `load`, `create`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/dutch-customer-data-sdk/releases](https://github.com/voxgig-sdk/dutch-customer-data-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { DutchCustomerDataSDK } from '@voxgig-sdk/dutch-customer-data'

const client = new DutchCustomerDataSDK()
```

### 2. List euapi records

`list()` resolves to an array of EuApI objects — iterate it directly:

```ts
const euapis = await client.EuApI().list()

for (const euapi of euapis) {
  console.log(euapi)
}
```

### 3. Load an euapi

`load()` returns the entity directly and throws on failure:

```ts
try {
  const euapi = await client.EuApI().load({ id: 'example_id' })
  console.log(euapi)
} catch (err) {
  console.error('load failed:', err)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const euapis = await client.EuApI().list()
  console.log(euapis)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = DutchCustomerDataSDK.test()

const euapi = await client.EuApI().list()
// euapi is a bare entity populated with mock response data
console.log(euapi)
```

You can also use the instance method:

```ts
const client = new DutchCustomerDataSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.EuApI()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data.id)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new DutchCustomerDataSDK({
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
DUTCH_CUSTOMER_DATA_TEST_LIVE=TRUE
```

Then run:

```bash
cd ts && npm test
```


## Reference

### DutchCustomerDataSDK

#### Constructor

```ts
new DutchCustomerDataSDK(options?: {
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `EuApI(data?)` | `EuApIEntity` | Create an EuApI entity instance. |
| `GlobalApI(data?)` | `GlobalApIEntity` | Create a GlobalApI entity instance. |
| `NetherlandsApI(data?)` | `NetherlandsApIEntity` | Create a NetherlandsApI entity instance. |
| `tester(testopts?, sdkopts?)` | `DutchCustomerDataSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `DutchCustomerDataSDK.test(testopts?, sdkopts?)` | `DutchCustomerDataSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): DutchCustomerDataSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load` and `create` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

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

Operations: list, load.

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

Operations: create, list, load.

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

Operations: list.

API path: `/bag`



## Entities


### EuApI

Create an instance: `const eu_ap_i = client.EuApI()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `buyer` | `string` |  |
| `buyer_country` | `string` |  |
| `contract_nature` | `string` |  |
| `html` | `string` |  |
| `id` | `string` |  |
| `link` | `string` |  |
| `notice_type` | `string` |  |
| `official_language` | `string` |  |
| `pdf` | `string` |  |
| `place_of_performance` | `string` |  |
| `procedure_type` | `string` |  |
| `publication_date` | `string` |  |
| `status` | `string` |  |
| `title` | `string` |  |
| `vat` | `Record<string, any>` |  |

#### Example: Load

```ts
const eu_ap_i = await client.EuApI().load({ id: 'eu_ap_i_id' })
```

#### Example: List

```ts
const eu_ap_is = await client.EuApI().list()
```


### GlobalApI

Create an instance: `const global_ap_i = client.GlobalApI()`

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
| `admin1` | `string` |  |
| `admin2` | `string` |  |
| `admin3` | `string` |  |
| `bic` | `Record<string, any>` |  |
| `city` | `string` |  |
| `currency` | `Record<string, any>` |  |
| `date` | `string` |  |
| `dns` | `Record<string, any>` |  |
| `email` | `Record<string, any>` |  |
| `found` | `number` |  |
| `freeformaddress` | `string` |  |
| `from_currency` | `string` |  |
| `iban` | `Record<string, any>` |  |
| `ip` | `Record<string, any>` |  |
| `lat` | `number` |  |
| `lei` | `Record<string, any>` |  |
| `letter` | `string` |  |
| `lon` | `number` |  |
| `municipality` | `string` |  |
| `number` | `number` |  |
| `password` | `Record<string, any>` |  |
| `phone` | `Record<string, any>` |  |
| `population` | `number` |  |
| `postcode` | `string` |  |
| `province` | `string` |  |
| `province_code` | `string` |  |
| `score` | `number` |  |
| `status` | `string` |  |
| `street` | `string` |  |
| `type` | `string` |  |
| `url` | `Record<string, any>` |  |
| `webrank` | `Record<string, any>` |  |

#### Example: Load

```ts
const global_ap_i = await client.GlobalApI().load()
```

#### Example: List

```ts
const global_ap_is = await client.GlobalApI().list()
```

#### Example: Create

```ts
const global_ap_i = await client.GlobalApI().create({
})
```


### NetherlandsApI

Create an instance: `const netherlands_ap_i = client.NetherlandsApI()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `active` | `number` |  |
| `addition` | `string` |  |
| `city` | `string` |  |
| `coc` | `string` |  |
| `construction_year` | `number` |  |
| `floor_area` | `number` |  |
| `freeformaddress` | `string` |  |
| `id` | `string` |  |
| `lat` | `number` |  |
| `letter` | `string` |  |
| `lon` | `number` |  |
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

```ts
const netherlands_ap_is = await client.NetherlandsApI().list()
```


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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
dutch-customer-data/
├── src/
│   ├── DutchCustomerDataSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { DutchCustomerDataSDK } from '@voxgig-sdk/dutch-customer-data'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const euapi = client.EuApI()
await euapi.list()

// euapi.data() now returns the euapi data from the last `list`
// euapi.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
