---
name: "fix-lint"
description: "Use when lint is failing in this repository and you want an agent to run the appropriate Makefile lint targets, apply safe fixes, and iterate until the target is clean."
---

# Fix Lint

Use this skill to resolve lint failures in this repository.

## Scope

The repo exposes these lint-related commands:

- `make lint`
- `make backend.lint`
- `make frontend.lint`
- `make docs.lint`
- `make format`
- `make backend.format`
- `make frontend.format`

## Workflow

1. Choose the narrowest target that matches the task.
2. Run the linter first and capture the failures.
3. Apply auto-fixes where the target supports them.
4. Fix remaining issues manually.
5. Re-run the linter until it is clean.
6. Summarize what was auto-fixed versus manually fixed.

## Target Selection

- Use `make backend.lint` for backend-only failures.
- Use `make frontend.lint` for frontend-only failures.
- Use `make docs.lint` for Markdown-only failures.
- Use `make lint` when the scope is mixed or unknown.

## Notes

- Do not suppress rules to make failures disappear.
- Prefer fixing the root cause.
- If the lint run fails because tooling is missing, call out the relevant
  setup target.
