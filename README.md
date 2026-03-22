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

### Backend

```sh
cd backend
go run ./cmd/server/
```

The server starts on `http://localhost:8080`.

Check the server is up:

```sh
curl http://localhost:8080/health
# {"status":"ok"}
```

Run a calculation:

```sh
curl -s -X POST http://localhost:8080/api/v1/calculations \
  -H "Content-Type: application/json" \
  -d '{"op":"add","a":10,"b":5}'
# {"result":15}
```

Makefile targets and Docker Compose support will be added in a later phase.

### Frontend

Prerequisites: Node.js 20+ and npm.

```sh
cd frontend
npm install
npm run dev
```

The dev server starts on `http://localhost:5173`.

All API requests are proxied to `http://localhost:8080`. The backend must
be running for calculations to work.

Run the tests:

```sh
npm test
```

To build and run the frontend with Docker:

```sh
docker build -t calculator-frontend ./frontend
docker run -p 3000:80 calculator-frontend
```

The container serves the static build via nginx on port 80, mapped to
`http://localhost:3000`.

Makefile targets and Docker Compose support will be added in a later phase.

## Project Structure

```text
backend/
  cmd/server/         ← entry point, graceful shutdown
  internal/
    calculator/       ← domain logic, sentinel errors, unit tests
    handler/          ← HTTP handlers, request/response models, handler tests
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

The key separation of concerns: components never import fetch directly.
All network calls go through `src/api/`, keeping UI logic and data
fetching independently testable.

## Design

Key design decisions are documented in:

- `docs/adr/0001-architecture-and-api.md`
- `docs/adr/0002-tooling-and-delivery.md`

## Notes

This repository intentionally prioritizes clear scope, maintainable
design, and pragmatic decision-making before implementation begins.
