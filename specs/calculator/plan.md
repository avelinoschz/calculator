# Implementation Plan

## Overview

This document records the historical implementation plan that guided
development of the calculator project.

It was used as part of a spec-driven, AI-assisted workflow: define the
requirements and contract first, then execute the work in explicit,
reviewable phases.

This file records the historical implementation process. Current
runtime behavior is still governed by:

- `specs/calculator/requirements.md`
- `api/openapi.yaml`

## Delivered State

All four phases were completed. The first cut of the calculator is
fully functional and includes:

- Go backend with domain/handler separation, structured logging,
  graceful shutdown, and version embedding
- React + TypeScript frontend with isolated API layer, three-layer
  tests, and responsive plain-CSS styling
- `POST /api/v1/calculations` and `GET /health` endpoints
- consistent JSON error responses with typed error codes
- multi-stage Dockerfiles and Docker Compose
- Makefile workflows for setup, run, test, coverage, lint, format,
  build, and Docker tasks
- GitHub Actions CI for lint, test, and build
- supporting ADRs and AI-usage documentation

The purpose of this plan was to keep the work:

- phased
- reviewable
- constrained by scope
- aligned with the requirements and API contract

## Delivery Strategy

The system was intentionally implemented in small sequential phases.

The core rule during development was:

> complete the required end-to-end path before investing in optional
> polish or extras.

That meant the project needed to become, in order:

- functionally correct
- easy to test
- easy to run
- easy to review

Only after that were tooling, delivery polish, and documentation
expanded.

## Implementation Phases

## Phase 0 — Specification Alignment [COMPLETED]

### Phase 0 Goal

Make the scope, contract, and priorities explicit before implementation
begins.

### Phase 0 Deliverables

- `specs/calculator/requirements.md`
- `specs/calculator/api.md`
- `api/openapi.yaml`
- `docs/adr/0001-architecture-and-api.md`
- `docs/adr/0002-tooling-and-delivery.md`
- `specs/calculator/plan.md`
- `AGENTS.md`

### Phase 0 Exit Criteria

- core scope is explicit
- API contract is stable enough to implement
- de-scope rules are defined
- AI collaboration constraints are written down before coding starts

## Phase 1 — Backend Core [COMPLETED]

### Phase 1 Goal

Build the smallest correct backend that satisfies the required
calculator operations and error handling model.

### Phase 1 Tasks

- create the Go module and backend structure
- keep calculator domain logic independent from HTTP handlers
- implement `GET /health`
- implement `POST /api/v1/calculations`
- validate request payloads explicitly
- return consistent JSON error responses
- add structured logging with `log/slog`
- add graceful shutdown handling
- add unit tests for domain logic
- add targeted handler tests for request/response behavior

### Phase 1 Exit Criteria

- all four core operations work
- invalid requests are handled correctly
- division by zero is handled correctly
- tests pass
- backend can run locally

### Phase 1 Deliverables

- calculator domain package
- HTTP handler package
- health and calculation endpoints
- structured logging
- graceful shutdown
- backend unit and handler tests

## Phase 2 — Frontend Core [COMPLETED]

### Phase 2 Goal

Build a minimal UI that consumes the backend cleanly and mirrors the
same separation-of-concerns principles.

### Phase 2 Tasks

- scaffold React + TypeScript frontend with Vite
- build calculator form with operand inputs and operation selector
- add submit flow
- isolate API calls in `src/api/`
- show loading, result, and error states
- add client-side validation before submission
- add basic responsive styling
- add tests at API, component, and integration layers

### Phase 2 Exit Criteria

- a user can perform all core operations from the UI
- invalid input is surfaced clearly
- backend errors are surfaced clearly
- frontend can run locally

### Phase 2 Deliverables

- Vite-based React frontend
- isolated API client
- validated calculator form
- result and error rendering
- responsive plain CSS
- three-layer frontend tests

## Phase 3 — Developer Experience and Quality Gates [COMPLETED]

### Phase 3 Goal

Make the project easy to run, test, lint, build, and evaluate.

### Phase 3 Tasks

- add a Makefile for common workflows
- pin backend tooling locally
- add backend and frontend Dockerfiles
- add Docker Compose for the full stack
- add CI for lint, test, and build
- keep local and CI commands aligned

### Phase 3 Exit Criteria

- common workflows are available through the Makefile
- the full stack can be started through Docker Compose
- CI covers the main quality gates

### Phase 3 Make Targets

- `make setup`
- `make run`
- `make test`
- `make coverage`
- `make lint`
- `make format`
- `make build`
- `make docker.build`
- `make up`
- `make down`

## Phase 4 — Documentation and Delivery Polish [COMPLETED]

### Phase 4 Goal

Make the repository understandable to a reviewer without requiring a
deep code read first.

### Phase 4 Tasks

- write or refine root and service READMEs
- document API behavior and examples
- capture architecture and tooling decisions in ADRs
- document the AI-assisted workflow and representative prompts
- trim obvious redundancy
- keep documentation aligned with implementation

### Phase 4 Exit Criteria

- a reviewer can run the project quickly
- the main design decisions are visible without reading the whole codebase
- AI usage is documented as part of the project history

### Phase 4 Extra Deliverables

- `docs/adr/0003-frontend-architecture.md`
- `docs/adr/0004-environment-variables-for-configuration.md`
- `docs/ai-prompts.md`

## De-scope Rules

If time or complexity forced scope reduction, cut work in this order:

1. skip optional calculator operations
2. skip advanced observability
3. keep Docker Compose but avoid extra polish
4. keep CI minimal
5. keep frontend styling minimal
6. do not cut validation, tests, or core documentation before cutting
   optional extras

### Hard Rule

Do not sacrifice:

- correctness of the four required operations
- clear validation and error handling
- testability
- readability

## Risk Management

### Main Risks

- spending too much time on tooling before the core works
- overengineering the project structure
- adding optional features too early
- letting frontend polish outrun backend completeness
- producing documentation that explains too much but guides too little

### Mitigations

- start with the backend core
- keep a single calculation endpoint
- keep frontend state and styling simple
- use the plan to sequence the work deliberately
- treat requirements and OpenAPI as the contract anchors

## Definition of Done

The project was considered done when:

- the four required operations worked end to end
- frontend and backend validation were implemented
- error responses were consistent
- tests covered the main behavior
- lint, test, and build workflows existed
- Docker Compose could run the full stack
- documentation made the design and workflow reviewable

## Relationship to AI Workflow

This plan and `docs/ai-prompts.md` serve different purposes:

- `specs/calculator/plan.md` records the phased execution strategy used
  during development
- `docs/ai-prompts.md` records representative prompt patterns used to
  execute that strategy with AI assistance

Together they document the spec-driven development process used for the
project.
