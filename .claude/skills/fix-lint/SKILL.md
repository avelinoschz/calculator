---
name: fix-lint
description: Run linters via Makefile targets and fix reported warnings. Accepts an optional argument to target a specific linter (backend, frontend, markdown/docs). Defaults to all linters.
argument-hint: [backend|frontend|markdown]
allowed-tools: Bash(make *), Read, Edit, Grep
---

You are helping the user run linters and fix all reported warnings in this project.

The project uses a Makefile with the following lint targets:
- `make lint` — run all linters (backend + frontend + docs)
- `make backend.lint` — golangci-lint on Go code
- `make frontend.lint` — ESLint on frontend code
- `make docs.lint` — markdownlint on all Markdown files

And the following auto-fix targets (where available):
- `make format` — auto-fix backend + frontend lint issues
- `make backend.format` — auto-fix Go lint issues via golangci-lint --fix
- `make frontend.format` — auto-fix frontend lint issues via eslint --fix
- (no auto-fix target for docs/markdown — fixes must be applied manually)

## Steps

1. **Determine the target** based on `$ARGUMENTS`:
   - `backend` → use `make backend.lint` / `make backend.format`
   - `frontend` → use `make frontend.lint` / `make frontend.format`
   - `markdown` or `docs` → use `make docs.lint` (manual fixes only)
   - empty or anything else → use `make lint` / `make format` (all linters)

2. **Run the linter** to collect the initial output. Capture all warnings and errors.

3. **Apply auto-fixes** where available:
   - For backend: run `make backend.format`
   - For frontend: run `make frontend.format`
   - For all: run `make format`
   - For markdown/docs: skip auto-fix (no target available — proceed to manual fixes)

4. **Re-run the linter** to see what issues remain after auto-fix.

5. **Fix remaining issues manually** by reading the flagged files and editing them directly. Address each remaining warning one by one.

6. **Re-run the linter again** to confirm it passes clean. If issues still remain, repeat step 5 until the linter reports no errors.

7. **Report a summary** of what was fixed (auto-fixed vs. manually fixed), and confirm the linter is now clean.

## Guidelines

- Never skip or suppress lint rules to make warnings disappear — fix the root cause.
- If a warning is ambiguous, prefer the more idiomatic fix for the language/tool.
- If a warning cannot be safely auto-fixed and you are unsure, explain it to the user and ask before editing.
- If the linter fails to run (e.g. missing tool), tell the user and suggest running the appropriate setup target (e.g. `make backend.setup`).
