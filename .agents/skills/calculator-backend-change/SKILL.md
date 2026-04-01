---
name: "calculator-backend-change"
description: "Use when changing the Go backend in this calculator repo, especially for domain logic, handler behavior, backend tests, or contract-aligned refactors that should stay small and maintainable."
---

# Calculator Backend Change

Use this skill for backend work in `backend/`.

## Workflow

1. Start with `$calculator-spec-review` or manually read the contract files first.
2. Inspect the current backend path before editing:
   - `backend/cmd/server/main.go`
   - `backend/internal/calculator/`
   - `backend/internal/handler/`
3. Keep changes small and explicit:
   - business logic stays out of HTTP handlers,
   - handlers validate and map errors consistently,
   - interfaces are only added when they create a real testing seam.
4. Update backend tests in the same change.
5. If behavior or tooling changes, update the smallest relevant docs.

## Backend Checklist

- contract still matches `api/openapi.yaml`,
- error codes and statuses remain stable,
- table-driven tests still cover core behavior,
- `make lint test build` stays green.

## Notes

- Prefer concrete types by default.
- If an interface is introduced, define it in the consumer layer and keep it small.
- Do not add new endpoints unless the spec truly requires them.
