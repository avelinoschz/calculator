# Backend

Go API server for the calculator project.

## Run

```sh
make backend.setup
make backend.run
```

The server listens on `http://localhost:8080`.

If you use `asdf`, the repository includes a root `.tool-versions` file
with the recommended Go and Node.js versions for local development.

## Structure

```text
backend/
  cmd/server/              entry point, env config, graceful shutdown
  internal/calculator/     calculation domain logic and concrete service
  internal/handler/        HTTP transport, request validation, JSON responses,
                           and a small service interface for handler-layer mocks
  Dockerfile               multi-stage build
  .golangci.yml            Go lint configuration
```

## Layering Notes

- `internal/calculator` is the authority for all business rules: operand
  range validation, operation arity, and domain error definitions
- `internal/handler` owns the `CalculatorService` interface (consumer
  owns the interface) and is responsible only for parsing HTTP input,
  delegating to the service, and mapping domain errors to HTTP responses
- `internal/calculator.NewService(min, max)` validates configuration
  invariants at construction time; invalid limits cause a startup failure
- domain errors carry their own machine-readable code and message;
  the handler maps them generically without per-error branching (see ADR 0006)

## Configuration

Optional environment variables:

| Variable | Description |
| --- | --- |
| `CALC_MIN` | Minimum allowed operand value |
| `CALC_MAX` | Maximum allowed operand value |

`make backend.run` sources the repository `.env` file automatically when
present.

If a configured value cannot be parsed as a float, startup logs a warning
and falls back to no limit on that side. If the parsed values are
structurally invalid (e.g. `CALC_MIN` greater than `CALC_MAX`), the
server exits with a fatal error at startup.

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
| `make ci` | Run the root validation gate (`lint`, `test`, and `build`) |
| `make backend.clean` | Remove all backend build artifacts |
| `make backend.clean.bin` | Remove only the backend binary (`backend/bin/`) |
| `make backend.clean.coverage` | Remove only coverage files (`coverage.out`, `coverage.html`) |
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
