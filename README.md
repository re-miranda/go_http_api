# A Minimal API in Go

> **Status:** Phase 1 — learning project focused on clarity, maintainability, and performance.

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

- `net/http` + a lightweight router (e.g., `chi` or `httprouter`)
- Simple request/response flow and basic middleware (logging and recovery)
- Health check endpoint and one example resource
- Clear project layout

---

## Quick start

**Prerequisite:** Go 1.22+

```bash
# clone
git clone https://github.com/re-miranda/go_http_api.git
cd go_http_api

# (optional) ensure modules are ready
go mod download

# run
make run

# test
make test

# clean
make clean
```
---

## Principles

- **Small surface area:** Fewer moving parts, fewer surprises.
- **Explicit over magic:** Prefer standard patterns and clear wiring.
- **Testable units:** Domain logic separated from HTTP concerns.

---

##  API Endpoint

- **GET /healthz** -> {"status": "ok"}
- **GET /v1/ping** -> pong
- **POST /v1/reverse** -> {"input": "your_input", "output": "your_input_reversed"}

### Error format
All non-2xx responses use the same JSON shape:
```
{
  "error": "<human-readable HTTP status code applicable>",
  "details": { ... } // optional
}
```

Examples:

- 400 invalid JSON:
```
{
    "error" : "Bad Request",
    "details" : [ "json: unknown field \"wrong_field\"" ]
}
```

- 405 wrong method:
```
{ "error" : "Method not allowed" }
```

---

## License

BSD-2-Clause © 2025 Renato Miranda Goncalves. See LICENSE for details.
