---
name: "calculator-spec-review"
description: "Use when starting work on this calculator repo and you need to align the request with AGENTS.md, requirements, the human-readable API guide, and the OpenAPI contract before planning, reviewing, or implementing changes."
---

# Calculator Spec Review

Use this skill at the start of any non-trivial task in this repository.

## Workflow

1. Read `AGENTS.md`.
2. Read `specs/calculator/requirements.md`.
3. Read `specs/calculator/api.md`.
4. Read `api/openapi.yaml`.
5. Summarize:
   - the requested change,
   - the contract constraints that matter,
   - which files are most likely to change.
6. Call out any mismatch between the request and the documented scope before implementation.

## Output

Keep the result short:

- task summary,
- contract constraints,
- likely code/test/doc touchpoints,
- risks or drift to watch for.

## Notes

- Treat the backend contract as authoritative.
- Prefer concrete file references over broad architectural discussion.
- Load additional files only after the contract is clear.
