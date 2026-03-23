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

## AI-Assisted Development

This repository is being prepared using AI-assisted workflows for:

- specification and planning
- implementation guidance
- code review and refinement

All AI-generated outputs were manually reviewed and validated.

Representative prompts used during development can be found in:

- `docs/ai-prompts.md`

## Target Features

- Addition
- Subtraction
- Multiplication
- Division

## API

- `POST /api/v1/calculations`

This endpoint is defined as part of the current project specification.

See:

- `specs/calculator/api.md`
- `api/openapi.yaml`

## How to Run

### Quick start — local dev

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

### Quick start — Docker Compose

Prerequisites: Docker with Compose v2.

```sh
make up   # builds images and starts the full stack
make down # stops all containers
```

The frontend is served by nginx on `http://localhost:80`. API requests to
`/api/` are proxied to the backend container.

### Verify the backend

```sh
curl http://localhost:8080/health
# {"status":"ok"}

curl -s -X POST http://localhost:8080/api/v1/calculations \
  -H "Content-Type: application/json" \
  -d '{"op":"add","a":10,"b":5}'
# {"result":15}
```

### All Makefile targets

```sh
make help
```

| Target | Description |
| --- | --- |
| `make run` | Run backend and frontend locally in parallel |
| `make backend.run` | Run the Go backend (port 8080) |
| `make frontend.run` | Run the Vite dev server (port 5173) |
| `make test` | Run all tests |
| `make backend.test` | Run Go tests |
| `make frontend.test` | Run Vitest (single-run) |
| `make lint` | Run all linters |
| `make backend.lint` | Run golangci-lint |
| `make frontend.lint` | Run ESLint |
| `make build` | Build backend binary and frontend assets |
| `make backend.build` | Build Go binary → `backend/bin/server` |
| `make frontend.build` | Build frontend static assets → `frontend/dist/` |
| `make docker.build` | Build all Docker images |
| `make backend.docker.build` | Build the backend Docker image |
| `make frontend.docker.build` | Build the frontend Docker image |
| `make up` | Start the full stack with Docker Compose |
| `make down` | Stop the full stack |

## Project Structure

```text
backend/
  cmd/server/         ← entry point, graceful shutdown
  internal/
    calculator/       ← domain logic, sentinel errors, unit tests
    handler/          ← HTTP handlers, request/response models, handler tests
  Dockerfile          ← golang:alpine build stage → distroless runtime stage
  .golangci.yml       ← golangci-lint configuration
```

```text
frontend/
  src/
    api/              ← typed fetch wrapper; no React imports
    components/       ← form components with client-side validation
    App.tsx           ← root component; holds result, error, and loading state
    main.tsx          ← ReactDOM.createRoot entry point
  index.html
  vite.config.ts      ← build config, dev proxy (/api → :8080), Vitest config
  Dockerfile          ← node:20-alpine build stage → nginx:alpine serve stage
```

```text
Makefile              ← common developer targets (run, test, lint, build, docker, compose)
docker-compose.yml    ← orchestrates backend + frontend containers
nginx.conf            ← nginx proxy config for Docker Compose (proxies /api/ to backend)
.github/workflows/
  ci.yml              ← GitHub Actions: lint, test, build (backend and frontend jobs)
```

The key separation of concerns: components never import fetch directly.
All network calls go through `src/api/`, keeping UI logic and data
fetching independently testable.

## Design

Key design decisions are documented in:

- `docs/adr/0001-architecture-and-api.md`
- `docs/adr/0002-tooling-and-delivery.md`

## Notes

This project intentionally prioritizes clear scope, maintainable design,
and correctness over feature volume. See `AGENTS.md` for implementation
guidance and priorities.
