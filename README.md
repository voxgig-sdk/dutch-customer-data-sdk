# DutchCustomerData SDK

Verify and standardize Dutch customer and business data — addresses, phone, email, VAT, IBAN, KvK and postal codes

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About Dutch Customer Data API

This SDK wraps the free validation endpoint at `https://free.bedrijfsdata.nl/v1.1`, operated by [Bedrijfsdata.nl](https://bedrijfsdata.nl) — a Dutch B2B data company based in Amsterdam that maintains profiles on more than 3.7 million Dutch companies and organisations.

The free service is aimed at verifying and standardising Dutch customer and business records. Typical inputs and outputs cover:

- Dutch and international phone numbers (with country code)
- Postal addresses and Dutch postal codes
- Email addresses and website URLs
- VAT (BTW) numbers
- IBAN bank account numbers
- KvK (Chamber of Commerce) registration numbers

The endpoint requires no API key and is intended for lightweight validation use. Rate limits and a formal licence are not published; for higher-volume or commercial use, contact Bedrijfsdata.nl about their paid datasets and enrichment products.

## Try it

**TypeScript**
```bash
npm install dutch-customer-data
```

**Python**
```bash
pip install dutch-customer-data-sdk
```

**PHP**
```bash
composer require voxgig/dutch-customer-data-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/dutch-customer-data-sdk/go
```

**Ruby**
```bash
gem install dutch-customer-data-sdk
```

**Lua**
```bash
luarocks install dutch-customer-data-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { DutchCustomerDataSDK } from 'dutch-customer-data'

const client = new DutchCustomerDataSDK({})

// List all euapis
const euapis = await client.EuApI().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o dutch-customer-data-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "dutch-customer-data": {
      "command": "/abs/path/to/dutch-customer-data-mcp"
    }
  }
}
```

## Entities

The API exposes 3 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **EuApI** | EU-scoped validation helpers — checks for data formats that are common across the European Union, such as VAT (BTW) numbers and IBAN bank account numbers. | `/tender` |
| **GlobalApI** | Globally-scoped validation helpers — generic checks that are not country-specific, such as email address and URL validation. | `/password` |
| **NetherlandsApI** | Netherlands-specific validation helpers — Dutch phone numbers, postal codes, addresses and KvK (Chamber of Commerce) numbers, e.g. `GET /v1.1/phone`. | `/bag` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from dutchcustomerdata_sdk import DutchCustomerDataSDK

client = DutchCustomerDataSDK({})

# List all euapis
euapis, err = client.EuApI(None).list(None, None)

# Load a specific euapi
euapi, err = client.EuApI(None).load(
    {"id": "example_id"}, None
)
```

### PHP

```php
<?php
require_once 'dutchcustomerdata_sdk.php';

$client = new DutchCustomerDataSDK([]);

// List all euapis
[$euapis, $err] = $client->EuApI(null)->list(null, null);

// Load a specific euapi
[$euapi, $err] = $client->EuApI(null)->load(
    ["id" => "example_id"], null
);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/dutch-customer-data-sdk/go"

client := sdk.NewDutchCustomerDataSDK(map[string]any{})

// List all euapis
euapis, err := client.EuApI(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "DutchCustomerData_sdk"

client = DutchCustomerDataSDK.new({})

# List all euapis
euapis, err = client.EuApI(nil).list(nil, nil)

# Load a specific euapi
euapi, err = client.EuApI(nil).load(
  { "id" => "example_id" }, nil
)
```

### Lua

```lua
local sdk = require("dutch-customer-data_sdk")

local client = sdk.new({})

-- List all euapis
local euapis, err = client:EuApI(nil):list(nil, nil)

-- Load a specific euapi
local euapi, err = client:EuApI(nil):load(
  { id = "example_id" }, nil
)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = DutchCustomerDataSDK.test()
const result = await client.EuApI().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = DutchCustomerDataSDK.test(None, None)
result, err = client.EuApI(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = DutchCustomerDataSDK::test(null, null);
[$result, $err] = $client->EuApI(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.EuApI(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = DutchCustomerDataSDK.test(nil, nil)
result, err = client.EuApI(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:EuApI(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the Dutch Customer Data API

- Upstream: [https://bedrijfsdata.nl](https://bedrijfsdata.nl)
- API docs: [https://free.bedrijfsdata.nl/v1.1](https://free.bedrijfsdata.nl/v1.1)

- Hosted as a free public endpoint at `free.bedrijfsdata.nl/v1.1` by [Bedrijfsdata.nl](https://bedrijfsdata.nl), an Amsterdam-based Dutch business-data provider.
- No authentication is required for the free tier.
- No explicit licence or terms of use are published alongside the free endpoint — treat the data as for development and validation use only, and check directly with the operator before commercial redistribution.
- CORS is not enabled, so calls must be made from a server-side context.

---

Generated from the Dutch Customer Data API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
