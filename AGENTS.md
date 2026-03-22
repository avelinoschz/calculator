
# AI Context

## Project Overview

This repository contains a small full-stack calculator application.

The goal is to demonstrate strong engineering judgment, with
maintainability in mind rather than feature volume.

The system consists of:

- Frontend: React + TypeScript
- Backend: Go (net/http)
- API: REST (JSON)

## Primary Objective

Deliver a clean, maintainable, and well-structured solution with clear priorities.

Priorities (in order):

1. Correctness
2. Maintainability
3. Readability
4. Clear error handling
5. Testability
6. Developer experience
7. Optional enhancements

## Core Requirements

- Follow `specs/calculator/requirements.md` for scope and acceptance criteria.
- Follow `api/openapi.yaml` for the canonical API contract.
- Keep the implementation aligned with the documented constraints and priorities.

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

- Use `specs/calculator/requirements.md` as the source of truth for scope.
- Use `api/openapi.yaml` as the source of truth for the API contract.
- Use `specs/calculator/api.md` as the human-readable companion to the API contract.

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

- Use `vitest` as the test runner
- Use `React Testing Library` for component rendering and interaction
- Use `@testing-library/user-event` for realistic user interaction simulation
- Test across three layers:
  - API layer (`src/api/`) — plain TypeScript, mock `fetch` with `vi.stubGlobal`
  - Component layer (`src/components/`) — render in isolation, no network calls
  - Integration (`App`) — full component tree with mocked fetch, assert on visible output
- Cover critical user flows:
  - valid submission
  - validation errors
  - API error handling

## Tooling Rules

- Use `golangci-lint` for Go
- Provide a `Makefile` for common tasks
- Support both:
  - local execution
  - Docker Compose execution

## Coding Guidelines

- Prefer clarity over cleverness
- Use descriptive naming
- Keep functions small
- Avoid premature abstraction
- Avoid unnecessary dependencies

## When Generating Code

Always:

- Follow existing file structure
- Respect separation of concerns
- Match existing naming conventions
- Keep implementations simple
- Include validation where required

Never:

- Introduce new frameworks without strong justification
- Add features outside the defined scope
- Overengineer the solution

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

If constraints require de-scoping, preserve core features and skip
optional enhancements first.

## Final Reminder

This project values strong engineering judgment.

The best solution is not the most complex one, but the one that is:

- clear
- correct
- maintainable
- well-prioritized
