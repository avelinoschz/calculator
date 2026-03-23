# Calculator

Simple full-stack calculator system built with:

- Backend: Go (net/http)
- Frontend: React + TypeScript

## Overview

This repository contains the specification, API contract, and design
decisions for a small full-stack calculator project.

The intended implementation is designed with a focus on:

- maintainability
- clarity
- correctness
- thoughtful engineering judgment

Rather than maximizing features, the goal is to deliver a clean,
well-structured, and production-minded solution.

## Document Guide

- `specs/calculator/requirements.md` is the source of truth for scope
  and acceptance criteria.
- `specs/calculator/plan.md` describes the intended implementation sequence.
- `specs/calculator/api.md` is the human-readable API guide.
- `api/openapi.yaml` is the canonical API contract.
- `docs/adr/0001-architecture-and-api.md` captures backend architecture and API design decisions.
- `docs/adr/0002-tooling-and-delivery.md` captures tooling and delivery decisions.
- `docs/adr/0003-frontend-architecture.md` captures frontend architecture decisions.
- `AGENTS.md` provides implementation guidance for AI-assisted workflows.
- `backend/README.md` covers the Go backend in full detail.
- `frontend/README.md` covers the React frontend in full detail.

## AI-Assisted Development

This repository is being prepared using AI-assisted workflows for:

- specification and planning
- implementation guidance
- code review and refinement

All AI-generated outputs were manually reviewed and validated.

Representative prompts used during development can be found in:

- `docs/ai-prompts.md`

## Features

- Addition
- Subtraction
- Multiplication
- Division

## API

`POST /api/v1/calculations`

### Quick example

```sh
curl -s -X POST http://localhost:8080/api/v1/calculations \
  -H 'Content-Type: application/json' \
  -d '{"op": "divide", "a": 10, "b": 3}' | jq .
```

```json
{
  "result": 3.3333333333333335
}
```

Division by zero returns HTTP 422 with `"code": "DIVISION_BY_ZERO"`.

See [`specs/calculator/api.md`](specs/calculator/api.md) and
[`api/openapi.yaml`](api/openapi.yaml) for the full contract, or
[`backend/README.md`](backend/README.md) for more examples.

## How to Run

### Local dev

Prerequisites: Go 1.21+, Node.js 20+, npm.

```sh
make run          # starts backend (:8080) and frontend (:5173) in parallel
```

Or start each service individually:

```sh
make backend.run  # Go server on http://localhost:8080
make frontend.run # Vite dev server on http://localhost:5173
```

All frontend API requests are proxied to `:8080` by the Vite dev server.

### Docker Compose

Prerequisites: Docker with Compose v2.

```sh
make up   # builds images and starts the full stack
make down # stops all containers
```

The frontend is served by nginx on `http://localhost:80`. API requests to
`/api/` are proxied to the backend container.

### Makefile targets

```sh
make help
```

| Target | Description |
| --- | --- |
| `make run` | Run backend and frontend locally in parallel |
| `make test` | Run all tests |
| `make lint` | Run all linters |
| `make build` | Build backend binary and frontend assets |
| `make docker.build` | Build all Docker images |
| `make up` | Start the full stack with Docker Compose |
| `make down` | Stop the full stack |

For per-service targets (`backend.test`, `frontend.lint`, etc.) see
[`backend/README.md`](backend/README.md) and
[`frontend/README.md`](frontend/README.md).

## Project Structure

```text
backend/    ← Go API server
frontend/   ← React + TypeScript UI
Makefile
docker-compose.yml
nginx.conf
.github/workflows/ci.yml
```

See [`backend/README.md`](backend/README.md) and
[`frontend/README.md`](frontend/README.md) for detailed structure and
per-service commands.

## Design Summary

- **Single endpoint** (`POST /api/v1/calculations`) — one stable API surface; operation type is a request field, not a route.
- **Separated layers** — calculator domain logic (`internal/calculator`) is independent of HTTP handlers; testable without network.
- **Dual validation** — frontend validates for UX; backend is the authoritative source and always validates.
- **Standard library only** — Go `net/http` and `log/slog`; no frameworks needed at this scale.
- **Isolated API layer** — frontend `src/api/` is decoupled from UI components and mocked independently in tests.

Full rationale in [`docs/adr/0001-architecture-and-api.md`](docs/adr/0001-architecture-and-api.md),
[`docs/adr/0002-tooling-and-delivery.md`](docs/adr/0002-tooling-and-delivery.md), and
[`docs/adr/0003-frontend-architecture.md`](docs/adr/0003-frontend-architecture.md).

## Trade-offs

| Decision | Rationale |
| --- | --- |
| Single calculation endpoint | Simpler API surface; avoids per-operation route proliferation |
| No persistence | Out of scope; stateless API is easier to test and deploy |
| Plain CSS, no UI framework | Minimal frontend dependency footprint |
| No authentication | Out of scope for a local demo |
| Distroless runtime image | Minimal attack surface; no shell in the production container |
| `net/http` over a framework | Sufficient for one endpoint; avoids unnecessary abstractions |

### Future improvements

If the scope were to grow:

- Add calculation history backed by a database
- Add OpenTelemetry tracing (noted as optional in requirements)
- Add more operations (exponentiation, square root, percentage)
- Add rate limiting and authentication for a public deployment

## Notes

This project intentionally prioritizes clear scope, maintainable design,
and correctness over feature volume. See `AGENTS.md` for implementation
guidance and priorities.
