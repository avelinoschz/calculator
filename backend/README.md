# Backend

Go API server for the calculator project.

## Run

```sh
make backend.setup
make backend.run
```

The server listens on `http://localhost:8080`.

## Structure

```text
backend/
  cmd/server/              entry point, env config, graceful shutdown
  internal/calculator/     calculation domain logic
  internal/handler/        HTTP transport, request validation, JSON responses
  Dockerfile               multi-stage build
  .golangci.yml            Go lint configuration
```

## Configuration

Optional environment variables:

| Variable | Description |
| --- | --- |
| `CALC_MIN` | Minimum allowed operand value |
| `CALC_MAX` | Maximum allowed operand value |

`make backend.run` sources the repository `.env` file automatically when
present.

If a configured value cannot be parsed as a float, startup logs a warning
and falls back to no limit on that side.

## API Behavior

Endpoints:

- `GET /health`
- `POST /api/v1/calculations`

Request shape:

```json
{
  "op": "add",
  "a": 10,
  "b": 5
}
```

Unary request shape:

```json
{
  "op": "sqrt",
  "a": 9
}
```

Success:

```json
{
  "result": 15
}
```

Error:

```json
{
  "error": {
    "code": "MISSING_FIELD",
    "message": "a is required"
  }
}
```

Validation rules:

- malformed JSON returns `400 INVALID_REQUEST`
- unknown request fields return `400 INVALID_REQUEST`
- trailing JSON after the first object returns `400 INVALID_REQUEST`
- missing `op` or `a` returns `400 MISSING_FIELD`
- missing `b` for binary operations returns `400 MISSING_FIELD`
- `sqrt` rejects `b` with `400 INVALID_REQUEST`
- unsupported operations return `400 INVALID_OPERATION`
- division by zero returns `422 DIVISION_BY_ZERO`
- negative square root returns `422 NEGATIVE_SQUARE_ROOT`
- non-finite results return `422 NON_FINITE_RESULT`
- configured operand-limit violations return `422 OPERAND_OUT_OF_RANGE`

Supported operations:

- `add`
- `subtract`
- `multiply`
- `divide`
- `power`
- `sqrt`
- `percentage`

Percentage is defined as `a% of b`, computed as `(a / 100) * b`.

See [`../api/openapi.yaml`](../api/openapi.yaml) and
[`../specs/calculator/api.md`](../specs/calculator/api.md) for the
contract.

## Make Targets

| Target | Description |
| --- | --- |
| `make backend.setup` | Install backend tooling and download Go module dependencies |
| `make backend.run` | Run the backend locally |
| `make backend.test` | Run Go tests |
| `make backend.coverage` | Run Go tests with `coverage.out` output |
| `make backend.coverage.html` | Generate `backend/coverage.html` from `coverage.out` |
| `make backend.lint` | Run `golangci-lint` (use `FIX=1` to auto-fix) |
| `make backend.format` | Format Go source files (`go fmt`) |
| `make backend.build` | Build `backend/bin/server` |
| `make backend.clean` | Remove backend build artifacts |
| `make backend.docker.build` | Build the backend Docker image |

## Testing and Linting

```sh
make backend.test
make backend.coverage
make backend.lint
```

Tests are table-driven where that improves clarity and cover both domain
logic and HTTP handler behavior.

## Build and Docker

```sh
make backend.build
make backend.docker.build
```

The build embeds the current git SHA in `main.version`. Docker builds
receive the same version value through `VERSION`.
