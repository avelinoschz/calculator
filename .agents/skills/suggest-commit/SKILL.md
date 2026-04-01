---
name: "suggest-commit"
description: "Use when staged changes are ready and you want an agent to inspect them, suggest a concise conventional commit message, and explain the reasoning."
---

# Suggest Commit

Use this skill when changes are already staged.

## Workflow

1. Inspect the staged diff.
2. Inspect the staged file list.
3. Infer:
   - conventional commit type,
   - optional scope,
   - concise imperative subject,
   - optional body if the change is non-trivial.
4. Propose the commit message in Conventional Commits format.
5. Briefly explain the choice.

## Guidelines

- Keep the subject under 72 characters.
- Use imperative mood.
- Do not invent issue numbers.
- If nothing is staged, stop and say so.

## Notes

- This skill suggests the message; it does not require making the commit.
- Prefer scopes that match actual repo areas such as `backend`,
  `frontend`, `docs`, or `tooling`.
