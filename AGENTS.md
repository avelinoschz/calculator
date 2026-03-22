
# AI Context

## Project Overview

This repository contains a small full-stack calculator application.

The goal is to demonstrate strong engineering judgment, with mantainability in mind, not feature volume.

The system consists of:

- Frontend: React + TypeScript
- Backend: Go (net/http)
- API: REST (JSON)

## Primary Objective

Deliver a clean, maintainable, and well-structured solution within a time constraint.

Priorities (in order):

1. Correctness
2. Maintainability
3. Readability
4. Clear error handling
5. Testability
6. Developer experience
7. Optional enhancements

## Core Requirements

- Support 4 operations: add, subtract, multiply, divide
- Single endpoint: `POST /api/v1/calculations`
- Frontend UI for input and result display
- Validation on frontend and backend
- Consistent JSON error model
- Unit tests for critical logic
- README with setup and rationale

## Non-Goals

Avoid implementing:

- authentication
- persistence
- caching
- complex state management
- advanced UI frameworks
- distributed architecture
- unnecessary abstractions

## Architecture Rules

### Backend

- Keep business logic separate from HTTP handlers
- Use small, focused packages
- Use Go standard library (`net/http`) for API
- Validate inputs explicitly
- Return consistent JSON responses

### Frontend

- Keep UI minimal and intuitive
- Prefer simple state management (no heavy libraries)
- Clearly separate API calls from UI logic

## API Rules

- Follow the contract defined in:
  - `docs/04-api-contract.md`
  - `api/openapi.yaml`

- Do not introduce additional endpoints unless strictly necessary
- Keep request/response shapes stable

## Logging

- Use structured logging (`log/slog`)
- Log meaningful events (request handling, errors)
- Avoid noisy or redundant logs

## Testing Rules

### Backend testing

- Test calculator logic thoroughly
- Use `testify`
- Prefer table-driven tests

### Frontend testing

- Test critical user flows:
  - valid submission
  - validation errors
  - API error handling

## Tooling Rules

- Use `golangci-lint` for Go
- Provide a `Makefile` for common tasks
- Support both:
  - local execution
  - Docker Compose execution

## Docker Rules

### Backend container

- Use multi-stage Docker build
- Final image should be minimal

### System

- Docker Compose runs frontend + backend

## CI Rules

- Provide a minimal GitHub Actions workflow
- Should run:
  - lint
  - tests
  - build

Do not implement full CD pipelines.

## Coding Guidelines

- Prefer clarity over cleverness
- Use descriptive naming
- Keep functions small
- Avoid premature abstraction
- Avoid unnecessary dependencies

## When Generating Code

Always:

- follow existing file structure
- respect separation of concerns
- match existing naming conventions
- keep implementations simple
- include validation where required

Never:

- introduce new frameworks without strong justification
- add features outside defined scope
- overengineer the solution

## Definition of Done

The implementation is complete when:

- All 4 operations work end-to-end
- API contract is respected
- Validation and error handling are correct
- Tests pass
- Code is readable and maintainable
- README explains setup and usage
- Project can be run locally or via Docker Compose

## Iteration Strategy

Work in small steps:

1. Backend core
2. Frontend core
3. Tests
4. Tooling (Makefile, Docker, CI)
5. Documentation

At each step:

- ensure code compiles
- ensure tests pass
- avoid partial or broken states

## Trade-off Philosophy

Prefer:

- simple and complete

Over:

- complex and incomplete

If time is limited:

- finish core features
- skip optional enhancements

## Final Reminder

This project is an evaluation of engineering judgment.

The best solution is not the most complex one, but the one that is:

- clear
- correct
- maintainable
- well-prioritized
