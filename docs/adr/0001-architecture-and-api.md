# ADR 0001: Architecture and API Design

## Status

Accepted

## Context

This project is a small calculator system. The goal is to prioritize
simplicity, maintainability, and correctness over feature breadth or
architectural complexity.

## Decisions

### 1. Keep the architecture intentionally small

The system will use a minimal architecture optimized for clarity and
maintainability rather than extensibility.

Implications:

- Avoid unnecessary abstractions
- Avoid heavy frameworks
- Keep file and package structure small and explicit

### 2. Separate business logic from transport logic

Calculator operations will be implemented independently from HTTP handlers.

Structure:

- domain/service layer for calculator logic
- HTTP layer for request parsing, validation, and response handling
- a small handler-owned interface for the calculation dependency, with
  the concrete implementation provided by the calculator package

Rationale:

- Improves testability
- Reduces coupling
- Keeps logic easy to reason about
- Makes handler-layer mocks straightforward without forcing the domain
  package to depend on transport concerns

### 2a. Let the consumer own the interface

When an interface is needed between layers, the consuming layer defines
it and the provider supplies the implementation.

Applied here:

- `internal/handler` owns the `CalculatorService` interface because the
  handler consumes that dependency
- `internal/calculator` owns the concrete `Service` implementation

Rationale:

- keeps interfaces close to the seam that needs substitution
- avoids exporting abstraction for callers that do not need it
- supports focused handler tests with mocks while preserving concrete
  domain code by default

### 3. Use a minimal REST API

The API exposes two endpoints:

- `POST /api/v1/calculations` — core calculator operation
- `GET /health` — liveness check for local and container environments

Rationale:

- Keeps API surface small
- Avoids duplication across multiple endpoints
- Simplifies frontend integration

### 4. Use a stable JSON response model

All responses follow consistent success and error shapes defined by the API contract.

Rationale:

- Simplifies client handling
- Improves consistency
- Aligns with OpenAPI contract

### 5. Validate on both frontend and backend

Validation responsibilities:

Frontend:

- UX improvements
- Early feedback

Backend:

- Source of truth
- Guarantees correctness

### 7. Use typed domain errors with generic HTTP mapping

The domain layer (`internal/calculator`) defines errors as exported
pointer-valued variables of type `*calculator.Error`, which carries both
a machine-readable code and a canonical human-readable message:

- `ErrInvalidOperation`
- `ErrDivisionByZero`
- `ErrNegativeSquareRoot`
- `ErrNonFiniteResult`
- `ErrOperandOutOfRange`

The HTTP handler maps them generically via `errors.As` and a static
status-code table, without a per-error branch. See ADR 0006 for the
full rationale.

Rationale:

- Errors are explicit, discoverable, and self-describing
- `errors.Is` continues to work by pointer identity
- Adding a new domain error does not require touching the handler
- Error messages are owned by the domain, not duplicated in the transport

### 6. Use Go standard library for HTTP

The backend will use `net/http` instead of introducing a framework.

Rationale:

- Sufficient for scope
- Reduces dependencies
- Keeps implementation explicit

## Consequences

### Positive

- Simple and easy to understand system
- Strong separation of concerns
- Easy to test core logic

### Trade-offs

- Less extensible for large-scale evolution
- No advanced routing features
- Adds a small amount of indirection in the handler path to improve seam
  testing between transport and domain layers

## Notes

This design intentionally favors simplicity over extensibility, given
the project's current scope.
