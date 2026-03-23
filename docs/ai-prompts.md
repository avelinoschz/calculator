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

## 3. Phase 3 — Developer Experience Prompt

The same implementation-prompt pattern was reused for Phase 3, with the
repository state, completed phases, and constraints updated to match the
next step in the plan.

The prompt shifted from product functionality to developer experience and
delivery requirements: Make targets, Docker, Docker Compose, and CI automation.

It followed the same core structure as the frontend implementation prompt
above, so it is summarized here rather than repeated in full.

## 4. Phase 4 — Documentation and Delivery Polish Prompt

Used this prompt to implement the documentation and delivery polish phase
after the full application stack and tooling were already in place.

Its purpose was to focus exclusively on documentation quality, reviewer
usability, and concise delivery polish without reopening stable application
code.

```text
This is a full-stack calculator project. The repository is at /github.com/avelinoschz/calculator.

Phases 1 (Go backend), 2 (React frontend), and 3 (Developer Experience and
Quality Gates) are complete. Read the following files before doing anything:

AGENTS.md                          — implementation rules and priorities
specs/calculator/requirements.md   — scope and acceptance criteria
specs/calculator/plan.md           — phased delivery plan

What is already built:

All source code, tests, Dockerfiles, docker-compose, nginx, CI, and Makefile
are in place.

Your task:

Implement Phase 4 — Documentation and Delivery Polish — as described in
specs/calculator/plan.md.

Key constraints from AGENTS.md:

- Do not reopen stable application code
- Prioritize reviewer usability and first-time reader clarity
- All documentation should reflect the actual implementation
- Keep documentation concise and scannable
```

## 5. Reviewer Prompt

Use this prompt for a final review pass after all phases are complete.
It covers the full application stack: backend, frontend, tooling,
developer experience, and documentation.

```text
Act as a strict and pragmatic senior engineer doing a final review
of a completed project before handoff.

Read the following files first:
- README.md
- AGENTS.md
- specs/calculator/requirements.md
- specs/calculator/plan.md
- specs/calculator/api.md
- api/openapi.yaml
- docs/adr/0001-architecture-and-api.md
- docs/adr/0002-tooling-and-delivery.md
- docs/adr/0003-frontend-architecture.md
- docs/adr/0004-environment-variables-for-configuration.md
- docs/ai-prompts.md
- Makefile
- docker-compose.yml
- .github/ (CI workflows)

Then review the current implementation across all four phases:

Backend (Phase 1):
- correctness of domain logic against requirements and API contract
- request validation and error response shape
- test coverage and quality (table-driven, edge cases)
- structured logging and graceful shutdown
- naming, package structure, and separation of concerns

Frontend (Phase 2):
- correctness of UI behavior against acceptance criteria
- isolation of API layer from components
- client-side validation before submission
- error and result state handling
- test coverage (unit, integration, end-to-end if applicable)
- frontend/backend contract alignment (field names, status codes, error codes)

Developer experience (Phase 3):
- Makefile targets: completeness, naming clarity, and correctness
- Docker builds: multi-stage correctness, image hygiene
- Docker Compose: service wiring, environment variable propagation
- CI pipeline: coverage of lint, test, and build steps

Documentation (Phase 4):
- README accuracy against actual implementation
- OpenAPI spec alignment with handler behavior
- ADR completeness and relevance
- First-time reader clarity and scannability

Do not implement changes yet.

Return the review in this format:
1. Critical issues (correctness, contract mismatches, broken tooling)
2. Medium issues (missing validation, weak tests, misleading docs)
3. Low-priority improvements (naming, structure, style)
4. Items that are out of scope and should not be added
5. Final verdict on readiness: is this project ready for handoff as-is?
```

## Notes

These prompts are intentionally concise and reusable.

They are meant to demonstrate structured use of AI for:

- planning
- implementation guidance
- review and quality control

Additional prompts used during implementation can be added later if needed.
