---
name: "calculator-final-check"
description: "Use when wrapping up work in this calculator repo and you need a concise final validation pass across implementation impact, contract drift, docs, and the required lint/test/build checks."
---

# Calculator Final Check

Use this skill near the end of a task.

## Workflow

1. Review the changed files and summarize the user-facing impact.
2. Check for drift against:
   - `specs/calculator/requirements.md`
   - `specs/calculator/api.md`
   - `api/openapi.yaml`
3. Confirm the smallest necessary docs were updated.
4. Run the validation helper script:
   - `sh .agents/skills/calculator-final-check/scripts/run_validation.sh`
5. If the helper script fails, inspect the failing step and fix the root cause.
6. Report:
   - what changed,
   - what was verified,
   - any remaining risk or follow-up.

## Notes

- Keep the final summary concise.
- Prefer real validation results over assumptions.
- Flag any contract mismatch before closing the task.
- This skill includes a script so the repo has one example of a more
  robust, tool-assisted skill definition.
