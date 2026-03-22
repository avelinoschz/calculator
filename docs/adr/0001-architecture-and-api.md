# ADR 0001: Architecture and API Design

## Status
Accepted

## Context

This project is a small calculator system. The goal is to prioritize simplicity, maintainability, and correctness over feature breadth or architectural complexity.

## Decisions

### 1. Keep the architecture intentionally small

The system will use a minimal architecture optimized for clarity and maintainability rather than extensibility.

Implications:
- Avoid unnecessary abstractions
- Avoid heavy frameworks
- Keep file and package structure small and explicit

---

### 2. Separate business logic from transport logic

Calculator operations will be implemented independently from HTTP handlers.

Structure:
- domain/service layer for calculator logic
- HTTP layer for request parsing, validation, and response handling

Rationale:
- Improves testability
- Reduces coupling
- Keeps logic easy to reason about

---

### 3. Use a minimal REST API with a single endpoint

The API will expose a single endpoint, `POST /api/v1/calculations`.

Rationale:
- Keeps API surface small
- Avoids duplication across multiple endpoints
- Simplifies frontend integration

---

### 4. Use a stable JSON response model

All responses follow consistent success and error shapes defined by the API contract.

Rationale:
- Simplifies client handling
- Improves consistency
- Aligns with OpenAPI contract

---

### 5. Validate on both frontend and backend

Validation responsibilities:

Frontend:
- UX improvements
- Early feedback

Backend:
- Source of truth
- Guarantees correctness

---

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

---

## Notes

This design intentionally favors simplicity over extensibility, given the project's current scope.
