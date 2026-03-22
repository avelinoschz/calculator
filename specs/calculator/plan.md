# Implementation Plan

## Overview

This document describes the execution plan for the calculator project.

Its purpose is to translate requirements and design decisions into a
realistic delivery strategy.

This plan derives from `specs/calculator/requirements.md`. The API
contract described here should be validated against `api/openapi.yaml`.

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

> finish the required path end-to-end before investing in optional
> enhancements.

This means the project should first become:

- functionally correct
- easy to run
- easy to review
- reasonably tested

Only after that should optional polish or production-minded extras be
added.

---

## Implementation Phases

## Phase 0 — Specification Alignment

### Phase 0 Goal

Ensure requirements, design decisions, and API contract are clear before
implementation starts.

### Phase 0 Deliverables

- `specs/calculator/requirements.md`
- `docs/adr/0001-architecture-and-api.md`
- `specs/calculator/api.md`
- `api/openapi.yaml`
- `specs/calculator/plan.md`

### Phase 0 Exit Criteria

- Core scope is clearly defined
- API contract is stable enough to implement
- Priorities and de-scope rules are explicit

---

## Phase 1 — Backend Core

### Phase 1 Goal

Build the smallest correct backend that supports the required
calculator operations.

### Phase 1 Tasks

- Create Go module and backend project structure
- Implement calculator domain logic separate from HTTP handlers
- Define request and response models
- Implement `POST /api/v1/calculations`
- Add request validation
- Add consistent JSON error responses
- Add structured logging with `log/slog`
- Add backend unit tests for core logic
- Add selected handler tests where they add value

### Phase 1 Exit Criteria

- All 4 core operations work
- Invalid requests are handled correctly
- Division by zero is handled correctly
- Tests pass
- Backend can be run locally

---

## Phase 2 — Frontend Core

### Phase 2 Goal

Build a minimal but clear UI that consumes the backend API.

### Phase 2 Tasks

- Scaffold React + TypeScript frontend
- Build calculator form with operand inputs and operation selector
- Add submit action
- Call backend API
- Show result state
- Show error state
- Add client-side validation
- Add basic component or interaction tests
- Add basic responsive layout support

### Phase 2 Exit Criteria

- A user can perform all core operations from the UI
- Invalid input is surfaced clearly
- Backend errors are surfaced clearly
- Frontend can be run locally

---

## Phase 3 — Developer Experience and Quality Gates

### Phase 3 Goal

Make the project easy to run, test, lint, and review.

### Phase 3 Tasks

- Add `Makefile`
- Add linting configuration
- Add backend Dockerfile using multi-stage build
- Add frontend Dockerfile
- Add Docker Compose for full-stack execution
- Add basic GitHub Actions workflow for lint, test, and build

### Phase 3 Expected Make Targets

- `make help`
- `make backend.run`
- `make frontend.run`
- `make test`
- `make lint`
- `make build`
- `make docker.build`
- `make up`
- `make down`

### Phase 3 Exit Criteria

- Common workflows are available via Makefile
- Full stack can be started with Docker Compose
- Basic lint/test/build workflow is codified

---

## Phase 4 — Documentation and Delivery Polish

### Phase 4 Goal

Make the project easy to understand and review.

### Phase 4 Tasks

- Write or refine README
- Add local run instructions
- Add Docker run instructions
- Add API usage examples
- Add design rationale summary
- Add trade-offs and future improvements
- Verify file structure and naming consistency
- Remove dead code or unnecessary complexity

### Phase 4 Exit Criteria

- Reviewers can run the project quickly
- Reviewers can understand the design without reading the entire
  codebase first
- Trade-offs are explicit

---

## De-scope Rules

If constraints require de-scoping, reduce scope in this order:

1. Skip optional calculator operations
2. Skip advanced observability
3. Keep Docker Compose, but avoid extra container polish
4. Keep CI minimal
5. Keep frontend styling minimal
6. Do not cut core validation, tests, or README before cutting optional
   extras

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
- Expanding Docker/CI/observability beyond the project's needs

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
- the README is sufficient for setup and usage
- the project can be run locally with clear commands
