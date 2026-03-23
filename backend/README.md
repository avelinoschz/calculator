# Backend

Go HTTP API server for the calculator project.

Built with the Go standard library (`net/http`), structured logging
(`log/slog`), and a single external dependency (`testify`) for testing.

## Prerequisites

- Go 1.21+

## Quick start

```sh
make backend.run
# or
cd backend && go run ./cmd/server
```

The server starts on `http://localhost:8080`.

## Project structure

```text
backend/
  cmd/server/
    main.go             ← entry point, graceful shutdown (SIGINT/SIGTERM)
  internal/
    calculator/
      calculator.go     ← domain logic, sentinel errors
      calculator_test.go
    handler/
      handler.go        ← HTTP handlers
      handler_test.go
      models.go         ← request/response types
  Dockerfile            ← multi-stage: golang:alpine → distroless/static
  .golangci.yml         ← lint configuration
  go.mod / go.sum
```

Business logic lives entirely in `internal/calculator` and has no
knowledge of HTTP. The handler layer maps domain errors to status codes
using `errors.Is`, keeping transport and domain concerns separate.

## API

The server exposes two endpoints. See [`specs/calculator/api.md`](../specs/calculator/api.md)
and [`api/openapi.yaml`](../api/openapi.yaml) for the canonical contract.

### GET /health

```sh
curl http://localhost:8080/health
```

```json
{"status": "ok"}
```

### POST /api/v1/calculations

**Request shape**

```json
{"op": "<operation>", "a": <number>, "b": <number>}
```

Supported operations: `add`, `subtract`, `multiply`, `divide`.

**Examples**

Addition:

```sh
curl -s -X POST http://localhost:8080/api/v1/calculations \
  -H "Content-Type: application/json" \
  -d '{"op":"add","a":10,"b":5}'
# {"result":15}
```

Subtraction:

```sh
curl -s -X POST http://localhost:8080/api/v1/calculations \
  -H "Content-Type: application/json" \
  -d '{"op":"subtract","a":10,"b":3}'
# {"result":7}
```

Multiplication:

```sh
curl -s -X POST http://localhost:8080/api/v1/calculations \
  -H "Content-Type: application/json" \
  -d '{"op":"multiply","a":6,"b":7}'
# {"result":42}
```

Division:

```sh
curl -s -X POST http://localhost:8080/api/v1/calculations \
  -H "Content-Type: application/json" \
  -d '{"op":"divide","a":20,"b":4}'
# {"result":5}
```

**Error examples**

Division by zero (422):

```sh
curl -s -X POST http://localhost:8080/api/v1/calculations \
  -H "Content-Type: application/json" \
  -d '{"op":"divide","a":10,"b":0}'
# {"error":{"code":"DIVISION_BY_ZERO","message":"division by zero is not allowed"}}
```

Invalid operation (400):

```sh
curl -s -X POST http://localhost:8080/api/v1/calculations \
  -H "Content-Type: application/json" \
  -d '{"op":"power","a":2,"b":3}'
# {"error":{"code":"INVALID_OPERATION","message":"operation must be one of add, subtract, multiply, divide"}}
```

**Error codes**

| Code | Status | Description |
| --- | --- | --- |
| `INVALID_REQUEST` | 400 | Malformed JSON or failed basic validation |
| `INVALID_OPERATION` | 400 | Operation is not supported |
| `MISSING_FIELD` | 400 | A required field is absent |
| `INVALID_NUMBER` | 400 | One or more operands are invalid |
| `DIVISION_BY_ZERO` | 422 | Division by zero is not allowed |
| `INTERNAL_ERROR` | 500 | Unexpected server error |

## Makefile targets

| Target | Description |
| --- | --- |
| `make backend.run` | Run the Go server (port 8080) |
| `make backend.test` | Run Go tests |
| `make backend.lint` | Run golangci-lint |
| `make backend.build` | Build binary → `backend/bin/server` |
| `make backend.docker.build` | Build the backend Docker image |

## Testing

```sh
make backend.test
# or
cd backend && go test ./...
```

Tests use `testify` and follow a table-driven pattern. The calculator
domain logic and HTTP handlers are tested independently.

## Linting

```sh
make backend.lint
# or
cd backend && golangci-lint run ./...
```

Enabled linters (configured in `.golangci.yml`):

- `errcheck` — unchecked errors
- `govet` — suspicious constructs
- `staticcheck` — static analysis
- `unused` — unused code
- `gofmt` — formatting
- `goimports` — import ordering

## Build

```sh
make backend.build
# or
cd backend && go build -o bin/server ./cmd/server
```

Output binary: `backend/bin/server`.

## Docker

```sh
make backend.docker.build
# or
docker build -t calculator-backend ./backend
```

The Dockerfile uses a two-stage build:

- **Stage 1 (`build`)** — `golang:alpine`; compiles a static binary with
  `CGO_ENABLED=0`
- **Stage 2 (`runtime`)** — `gcr.io/distroless/static-debian12`; no shell,
  no package manager, minimal attack surface

Run the image standalone:

```sh
docker run -p 8080:8080 calculator-backend
```
