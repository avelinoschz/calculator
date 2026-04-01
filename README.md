# Calculator

Small full-stack calculator built with Go (`net/http`) and React + TypeScript.

![Calculator UI](docs/assets/screenshot.png)

## Overview

The project is intentionally narrow in scope:

- seven operations: add, subtract, multiply, divide, power, sqrt,
  percentage
- one calculation endpoint: `POST /api/v1/calculations`
- one health endpoint: `GET /health`
- backend-enforced operand limits via `CALC_MIN` / `CALC_MAX`
- matching frontend validation defaults via `VITE_CALC_MIN` / `VITE_CALC_MAX`

Primary references:

- `specs/calculator/requirements.md` — scope and acceptance criteria
- `api/openapi.yaml` — canonical API contract
- `specs/calculator/api.md` — human-readable API guide
- `specs/calculator/plan.md` — historical phased implementation plan used
  during the AI-assisted development process
- `backend/README.md` — backend usage and behavior
- `frontend/README.md` — frontend usage and behavior

## AI-Assisted Development

This repository was developed with AI assistants.

The workflow was intentionally spec-driven:

- define requirements and API contract first
- break implementation into explicit phases
- use AI to help plan, implement, review, and refine each phase
- keep all outputs manually reviewed and aligned to the repo contract

The two main process artifacts are:

- `specs/calculator/plan.md` — the phased execution guide used during
  development
- `docs/ai-prompts.md` — representative prompts preserved from that
  workflow

The repo now also includes local AI workflow artifacts under `.agents/`:

- reusable skills in `.agents/skills/` for spec review, backend changes,
  final validation, lint fixing, and commit suggestion
- agent roles in `.agents/agents/` for planning, backend work, and review

These are intentionally small and repo-specific. They capture practical
skills, sub-agent responsibilities, and context boundaries for work in
this repository.

See `.agents/README.md` for the purpose of the skills, the sub-agent
roles, and suggested MCP categories that fit this repository.

## Run

Prerequisites:

- Go 1.26+
- Node.js 24+
- npm
- Docker with Compose v2 for containerized usage

If you use `asdf`, the repository pins the recommended local toolchain in
`.tool-versions`.

Local development:

```sh
make setup
make run
```

- backend: `http://localhost:8080`
- frontend: `http://localhost:5173`

Docker Compose:

```sh
make up
make down
```

- frontend: `http://localhost:80`
- API requests to `/api/` are proxied to the backend by the frontend image

## API

Quick example:

```sh
curl -s -X POST http://localhost:8080/api/v1/calculations \
  -H 'Content-Type: application/json' \
  -d '{"op":"power","a":2,"b":3}'
```

```json
{
  "result": 8
}
```

Unary example:

```sh
curl -s -X POST http://localhost:8080/api/v1/calculations \
  -H 'Content-Type: application/json' \
  -d '{"op":"sqrt","a":9}'
```

Error shape:

```json
{
  "error": {
    "code": "DIVISION_BY_ZERO",
    "message": "division by zero is not allowed"
  }
}
```

Validation notes:

- malformed JSON, extra fields, and trailing payloads return `400 INVALID_REQUEST`
- missing `op` or `a`, and missing `b` for binary operations, return `400 MISSING_FIELD`
- unsupported operations return `400 INVALID_OPERATION`
- `sqrt` rejects a second operand with `400 INVALID_REQUEST`
- division by zero returns `422 DIVISION_BY_ZERO`
- negative square root returns `422 NEGATIVE_SQUARE_ROOT`
- non-finite results return `422 NON_FINITE_RESULT`
- configured operand-limit violations return `422 OPERAND_OUT_OF_RANGE`

## Make Targets

```sh
make help
```

| Target | Description |
| --- | --- |
| `make setup` | Bootstrap local environment (tools + dependencies) |
| `make backend.setup` | Install backend tooling and download Go module dependencies |
| `make frontend.setup` | Install Node dependencies (`npm ci`) |
| `make run` | Run backend and frontend locally in parallel |
| `make test` | Run all tests |
| `make coverage` | Run all tests with coverage reports |
| `make lint` | Run all linters (use `FIX=1` to auto-fix) |
| `make format` | Format source files (Go: `go fmt`, frontend: Prettier) |
| `make build` | Build backend binary and frontend assets |
| `make ci` | Run the full validation gate: lint, test, and build |
| `make clean` | Remove all build artifacts and installed tools |
| `make clean.bin` | Remove only installed tools (`bin/`) |
| `make docker.build` | Build both Docker images |
| `make up` | Start the full stack with Docker Compose |
| `make down` | Stop the full stack |

Per-service targets are documented in `backend/README.md` and `frontend/README.md`.

## Project Structure

```text
backend/    Go API server
frontend/   React + TypeScript UI
api/        OpenAPI contract
.agents/    Repo-local skills and agent roles
docs/adr/   Architecture and tooling decisions
specs/      Requirements, API guide, and implementation plan
```

## Design Summary

- business logic lives in `backend/internal/calculator`, separate from HTTP handling
- the HTTP layer depends on a small calculator service interface so
  handler tests can mock the layer boundary explicitly
- frontend API calls are isolated in `frontend/src/api/`
- validation happens on both sides, with the backend as the source of truth
- unary and binary operations share one endpoint with explicit request schemas
- Docker images are multi-stage and self-contained
- CI runs lint, test, and build through the documented Make targets

## Documentation Notes

The repo also includes historical process and design artifacts:

- `docs/adr/0001-architecture-and-api.md`
- `docs/adr/0002-tooling-and-delivery.md`
- `docs/adr/0003-frontend-architecture.md`
- `docs/adr/0004-environment-variables-for-configuration.md`
- `docs/adr/0005-operation-arity-and-finite-result-handling.md`
- `specs/calculator/plan.md`
- `docs/ai-prompts.md`
