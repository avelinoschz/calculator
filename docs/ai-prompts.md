# AI Usage Summary

This project was developed using an AI-assisted workflow to improve
productivity, structure, and code quality.

AI was used as a collaborative assistant for:

- translating requirements into structured specifications
- defining architecture and design decisions (ADRs)
- proposing project structure
- guiding implementation sequencing
- generating and refining code
- reviewing the implementation from a senior-level code review perspective
- drafting and improving documentation

All AI-generated or AI-assisted outputs were manually reviewed,
validated, and adjusted before inclusion.

The goal was not to rely blindly on AI, but to use it as a tool to
accelerate iteration while maintaining full ownership of design
decisions and code quality.

## AI Prompts

This document captures a small set of representative prompts used during development.

The goal is not to document every single prompt, but to preserve the
most useful prompts that shaped the project structure, implementation
flow, and review process.

## 1. Initial Structure Prompt

Used this prompt at the start of the project to propose an initial
full-stack structure before implementation began.

Its purpose was to translate the requirements, ADRs, and delivery plan
into a simple project layout that favored maintainability over
abstraction.

In practice, it mostly informed the Phase 1 backend structure because
the implementation was intentionally sequenced backend-first.

```text
Read the following files:
- AGENTS.md
- specs/calculator/requirements.md
- specs/calculator/plan.md
- specs/calculator/api.md
- api/openapi.yaml
- docs/adr/0001-architecture-and-api.md
- docs/adr/0002-tooling-and-delivery.md

Propose a minimal and clean project structure for:
- Go backend
- React + TypeScript frontend

The structure should:
- follow the documented requirements and ADRs
- prioritize simplicity and maintainability
- avoid overengineering
- separate domain logic from transport logic in the backend
- support the documented developer experience requirements

Do not generate implementation code yet.
Only propose the project structure and explain the reasoning behind it.
```

## 2. Frontend Implementation Prompt

Used this prompt to implement the frontend phase after the backend API
contract and core handlers were already in place.

Its purpose was to keep the frontend work aligned with the documented
project priorities: simple structure, clear separation between UI and
API code, client-side validation, and compatibility with the existing
backend contract.

This prompt was most useful once the backend endpoints, request shape,
and implementation boundaries were already stable.

```text
This is a full-stack calculator project. The repository is at /github.com/avelinoschz/calculator.

Phase 1 (Go backend) is complete. Read the following files before doing anything:

AGENTS.md — implementation rules and priorities
specs/calculator/requirements.md — scope and acceptance criteria
specs/calculator/plan.md — phased delivery plan
specs/calculator/api.md — human-readable API contract
api/openapi.yaml — canonical API contract
What is already built


backend/
  cmd/server/main.go                          ← entry point, graceful shutdown
  internal/calculator/calculator.go           ← domain logic, sentinel errors
  internal/calculator/calculator_test.go      ← 19 table-driven tests
  internal/handler/handler.go                 ← HTTP handler
  internal/handler/handler_test.go            ← 9 handler tests
  internal/handler/models.go                  ← request/response structs
  go.mod / go.sum
The backend runs on :8080. Two endpoints: GET /health and POST /api/v1/calculations. The request field for the operation is op (not operation).

Your task

Implement Phase 2 — Frontend Core — as described in specs/calculator/plan.md.

The agreed project structure for the frontend is:


frontend/
  src/
    api/
      calculator.ts       ← fetch client, isolated from UI
    components/
      CalculatorForm.tsx  ← inputs, operation selector, submit
    App.tsx               ← root, holds result/error state
    main.tsx
  index.html
  package.json
  tsconfig.json
  vite.config.ts
  Dockerfile
Key constraints from AGENTS.md:

React + TypeScript, Vite build tool
No heavy UI frameworks, no complex state management
Separate API calls from UI logic (src/api/ must not bleed into components)
Client-side validation before submitting
Display result and error states clearly
Basic responsive layout for mobile
Start by reading the spec files listed above, then plan before implementing.
```

## 3. Tooling and Delivery Prompt

Used this prompt to implement the developer-experience and quality-gate
phase after both backend and frontend core functionality were already in
place.

Its purpose was to keep the final project setup aligned with the
documented delivery requirements: simple developer workflows, Docker
support, CI checks, and minimal tooling that improved reliability
without expanding scope.

This prompt was most useful once the application code and API behavior
were already stable, since it explicitly focused on Make targets,
containerization, and automation rather than product features.

```text
This is a full-stack calculator project. The repository is at /Users/avelino/repos/github.com/avelinoschz/calculator.

Phases 1 (Go backend) and 2 (React frontend) are complete. Read the following
files before doing anything:

  AGENTS.md                          — implementation rules and priorities
  specs/calculator/requirements.md   — scope and acceptance criteria
  specs/calculator/plan.md           — phased delivery plan

What is already built:

  backend/
    cmd/server/main.go                  ← entry point, graceful shutdown
    internal/calculator/calculator.go   ← domain logic, sentinel errors
    internal/calculator/calculator_test.go
    internal/handler/handler.go         ← HTTP handlers (GET /health, POST /api/v1/calculations)
    internal/handler/handler_test.go
    internal/handler/models.go
    go.mod / go.sum

  frontend/
    src/api/calculator.ts               ← fetch client, isolated from UI
    src/components/CalculatorForm.tsx   ← form with client-side validation
    src/App.tsx                         ← root component
    src/main.tsx
    src/App.css
    index.html
    package.json / tsconfig.json
    vite.config.ts                      ← dev proxy /api → :8080, Vitest config
    Dockerfile                          ← node:20-alpine build → nginx:alpine serve

The backend runs on :8080. The frontend dev server runs on :5173 and proxies
/api to the backend. The frontend Dockerfile already exists.

Your task

Implement Phase 3 — Developer Experience and Quality Gates — as described in
specs/calculator/plan.md.

The expected Makefile targets are:

  make help
  make backend.run
  make frontend.run
  make test
  make lint
  make build
  make docker.build
  make up
  make down

Key constraints:

  - The backend Dockerfile must use a multi-stage build (build stage + minimal
    runtime image such as distroless or alpine)
  - Docker Compose must orchestrate both frontend and backend together
  - The GitHub Actions workflow must run lint, test, and build
  - The frontend Dockerfile already exists — do not recreate it
  - Do not change any backend or frontend source code

Start by reading the spec files listed above, then plan before implementing.
```

## 4. Reviewer Prompt

Use this prompt after a meaningful implementation step, such as backend
core, frontend core, or a final review pass.

```text
Act as a strict and pragmatic code reviewer.

Read the following files first:
- README.md
- AGENTS.md
- specs/calculator/requirements.md
- specs/calculator/plan.md
- specs/calculator/api.md
- api/openapi.yaml
- docs/adr/0001-architecture-and-api.md
- docs/adr/0002-tooling-and-delivery.md

Then review the current implementation for:
- correctness against requirements
- maintainability
- unnecessary complexity
- API contract mismatches
- missing validation
- missing tests
- frontend/backend inconsistencies
- naming or structure issues

Do not implement changes yet.

Return the review in this format:
1. Critical issues
2. Medium issues
3. Low-priority improvements
4. Items that should be de-scoped
5. Final review of code quality and maintainability
```

## Notes

These prompts are intentionally concise and reusable.

They are meant to demonstrate structured use of AI for:

- planning
- implementation guidance
- review and quality control

Additional prompts used during implementation can be added later if needed.
