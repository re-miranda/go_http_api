# A Minimal API in Go

> **Status:** Phase 1 — learning project focused on clarity, maintainability, and performance with multi-server configuration support.

![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go)
![Status](https://img.shields.io/badge/status-experimental-blueviolet)
[![License](https://img.shields.io/badge/license-BSD--2--Clause-lightgrey)](LICENSE)

## Why minimal?

For this API, I’m deliberately using Go’s standard library with a **lightweight router** (no full framework). I evaluated options like **Gin** and **Echo**—they’re powerful and convenient, but introduce extra abstractions and dependencies. For this project’s goals, minimal wins:

- **Lean & transparent:** Easier to read, reason about, and maintain
- **Low onboarding cost:** New contributors don’t need to learn a framework first; just Go
- **Closer to the metal:** Deepens understanding of Go’s native HTTP primitives

This phase optimizes for **clarity, maintainability, and performance** over convenience.

---

## Current scope

- `net/http` + `httprouter` for efficient routing with low memory allocation
- **Multi-server configuration:** Run multiple API servers with different endpoints on different ports from a single binary
- JSON-based configuration for flexible server setup
- Simple request/response flow with proper error handling
- Comprehensive test coverage for handlers

---

### Router

`net/http` + `httprouter` — escolha orientada a **muitas rotas**, **baixa alocação** e **patterns de rota expressivos** (ex.: `/v1/users/:id`, `/*filepath`)

---

## Quick start

**Prerequisite:** Go 1.22+

```bash
# clone
git clone https://github.com/re-miranda/go_http_api.git
cd go_http_api

# ensure modules are ready
go mod download

# run with default config
make run

# run with custom config
./bin/api -config api_config.json

# test
make test

# clean
make clean
```
---

## Configuration

The API supports flexible configuration through JSON files. By default, it runs two servers on different ports:

### Example Configuration [api_config.json](api_config.json)

```json
{
  "Global": {
    "Name": "Go API",
    "ReadTimeout": 5,
    "WriteTimeout": 10,
    "IdleTimeout": 60
  },
  "Servers": [
    {
      "Name": "Reverse API",
      "Host": "0.0.0.0",
      "Port": "8080",
      "Routes": [...]
    },
    {
      "Name": "Ping API",
      "Host": "0.0.0.0",
      "Port": "8081",
      "ReadTimeout": 10,
      "WriteTimeout": 20,
      "IdleTimeout": 120,
      "Routes": [...]
    }
  ]
}
```

Servers inherit global timeout settings unless overridden.

---

##  API Endpoints (From default config)

### Server 1 (Port 8080) - Reverse API
- **GET /healthz** → `{"status": "ok"}`
- **POST /v1/reverse** → `{"input": "your_input", "output": "your_input_reversed"}`

### Server 2 (Port 8081) - Ping API
- **GET /healthz** → `{"status": "ok"}`
- **GET /v1/ping** → `pong`

### Error format
All non-2xx responses use consistent JSON structure:
```json
{
  "error": "<HTTP status description>",
  "details": ["optional", "error", "details"]
}
```

Examples:

- 400 Bad Request (invalid JSON):
```json
{
  "error": "Bad Request",
  "details": ["json: unknown field \"wrong_field\""]
}
```

- 405 Method Not Allowed:
```json
{
  "error": "Method not allowed"
}
```

---

## Features

### Implemented
- ✅ Multi-server support from single binary
- ✅ JSON-based configuration
- ✅ Request size limiting (1MB max)
- ✅ Content-Type validation
- ✅ Custom error handlers (404, 405)
- ✅ Consistent error response format
- ✅ UTF-8 support (including emoji and international characters)

### Security & Validation
- Enforces `Content-Type: application/json` for POST requests
- Limits request body size to prevent abuse
- Disallows unknown JSON fields in resquest body
- Proper timeout configuration per server

---

## Router Choice

**`httprouter`** was chosen for:
- **High performance:** Optimized for many routes with minimal allocations
- **Expressive patterns:** Support for parameters (`/users/:id`) and wildcards (`/*filepath`)
- **Zero garbage:** Careful memory management for production workloads
- **Simplicity:** Clean API that complements Go's `net/http`

---

## License

BSD-2-Clause © 2025 Renato Miranda Goncalves. See [LICENSE](LICENSE) for details.
