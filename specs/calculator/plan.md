# Implementation Plan

## Overview

This document describes the execution plan for the take-home calculator project.

Its purpose is to translate requirements and design decisions into a realistic delivery strategy under a 2–4 hour time constraint.

The plan prioritizes:

- core completeness
- maintainability
- testability
- reviewer experience
- disciplined scope control

---

## Delivery Strategy

The project should be implemented in small, sequential phases.

Each phase should produce something concrete and reviewable.

The core rule is:

> finish the required path end-to-end before investing in optional enhancements.

This means the project should first become:

- functionally correct
- easy to run
- easy to review
- reasonably tested

Only after that should optional polish or production-minded extras be added.

---

## Priorities

### P0 — Must Have

- Core calculator operations:
  - addition
  - subtraction
  - multiplication
  - division
- Single REST endpoint
- Frontend form for calculation input and result display
- Backend validation and error handling
- Frontend validation and user-friendly errors
- Unit tests for critical backend logic
- Basic frontend tests for key behavior
- README with setup, usage, and design rationale
- OpenAPI contract

### P1 — Should Have

- Structured logging
- Makefile
- Dockerfiles for frontend and backend
- Docker Compose for full-stack local run
- Basic linting setup
- Basic GitHub Actions CI workflow
- Basic responsive UI polish

### P2 — Nice to Have

- Optional calculator operations
- Observability extensions such as OpenTelemetry + Jaeger
- Coverage target helpers
- Extra Make targets
- Additional DX polish

---

## Implementation Phases

## Phase 0 — Specification Alignment

### Goal

Ensure requirements, design decisions, and API contract are clear before implementation starts.

### Deliverables

- `specs/calculator/requirements.md`
- `docs/adr/0001-architecture-and-api.md`
- `specs/calculator/api.md`
- `api/openapi.yaml`
- `specs/calculator/plan.md`

### Exit Criteria

- Core scope is clearly defined
- API contract is stable enough to implement
- Priorities and de-scope rules are explicit

---

## Phase 1 — Backend Core

### Goal

Build the smallest correct backend that supports the required calculator operations.

### Tasks

- Create Go module and backend project structure
- Implement calculator domain logic separate from HTTP handlers
- Define request and response models
- Implement `POST /api/v1/calculations`
- Add request validation
- Add consistent JSON error responses
- Add structured logging with `log/slog`
- Add backend unit tests for core logic
- Add selected handler tests where they add value

### Exit Criteria

- All 4 core operations work
- Invalid requests are handled correctly
- Division by zero is handled correctly
- Tests pass
- Backend can be run locally

---

## Phase 2 — Frontend Core

### Goal

Build a minimal but clear UI that consumes the backend API.

### Tasks

- Scaffold React + TypeScript frontend
- Build calculator form with operand inputs and operation selector
- Add submit action
- Call backend API
- Show result state
- Show error state
- Add client-side validation
- Add basic component or interaction tests
- Add basic responsive layout support

### Exit Criteria

- A user can perform all core operations from the UI
- Invalid input is surfaced clearly
- Backend errors are surfaced clearly
- Frontend can be run locally

---

## Phase 3 — Developer Experience and Quality Gates

### Goal

Make the project easy to run, test, lint, and review.

### Tasks

- Add `Makefile`
- Add linting configuration
- Add backend Dockerfile using multi-stage build
- Add frontend Dockerfile
- Add Docker Compose for full-stack execution
- Add basic GitHub Actions workflow for lint, test, and build

### Expected Make Targets

- `make help`
- `make backend.run`
- `make frontend.run`
- `make test`
- `make lint`
- `make build`
- `make docker.build`
- `make up`
- `make down`

### Exit Criteria
- Common workflows are available via Makefile
- Full stack can be started with Docker Compose
- Basic lint/test/build workflow is codified

---

## Phase 4 — Documentation and Submission Polish

### Goal
Make the submission easy for a reviewer to understand and evaluate.

### Tasks
- Write or refine README
- Add local run instructions
- Add Docker run instructions
- Add API usage examples
- Add design rationale summary
- Add trade-offs and future improvements
- Verify file structure and naming consistency
- Remove dead code or unnecessary complexity

### Exit Criteria
- Reviewer can run the project quickly
- Reviewer can understand the design without reading the entire codebase first
- Trade-offs are explicit

---

## Suggested Timeboxes

These are approximate and should be adjusted dynamically.

| Time Block | Focus |
|---|---|
| 20–30 min | Specification and structure |
| 60–90 min | Backend core |
| 45–75 min | Frontend core |
| 30–45 min | Tests, README, cleanup |
| Remaining | Docker, Makefile, CI, polish |

---

## De-scope Rules

If time becomes tight, reduce scope in this order:

1. Skip optional calculator operations
2. Skip advanced observability
3. Keep Docker Compose, but avoid extra container polish
4. Keep CI minimal
5. Keep frontend styling minimal
6. Do not cut core validation, tests, or README before cutting optional extras

### Hard Rule
Do not sacrifice:
- correctness of the 4 required operations
- clear validation and error handling
- testability
- readability

---

## Risk Management

### Main Risks
- Spending too much time on tooling before the core works
- Overengineering project structure
- Adding optional features too early
- Losing time in frontend polish beyond what is needed
- Expanding Docker/CI/observability beyond the assignment value

### Mitigations
- Build backend core first
- Keep a single API endpoint
- Keep frontend intentionally simple
- Add tooling only after the end-to-end path works
- Use the plan as the source of priority decisions

---

## Definition of Done

The project is considered done when:
- the 4 required operations work end-to-end
- frontend and backend validation are implemented
- errors are handled consistently
- the code is readable and reasonably structured
- unit tests cover key functionality
- the API contract is documented
- the README is sufficient for evaluation
- the project can be run locally with clear commands

---

## Submission Mindset

The target is not to build the most feature-rich calculator.

The target is to submit a small system that communicates:
- sound prioritization
- maintainable design
- pragmatic engineering judgment
- awareness of production concerns
- disciplined execution under constraint
