---
name: suggest-commit
description: Inspects currently staged files and suggests a conventional commit message. Offers to run git commit with the suggested message.
allowed-tools: Bash(git *)
---

You are helping the user craft a high-quality git commit message for their staged changes.

## Steps

1. Run `git diff --cached` to see all staged changes in detail.
2. Run `git diff --cached --name-status` to get the list of staged files and their change type (Added, Modified, Deleted).
3. Analyze the changes and determine:
   - The **type** of change: `feat`, `fix`, `chore`, `refactor`, `test`, `docs`, `style`, `ci`, or `build`
   - The **scope** (optional): the affected area or module (e.g. `auth`, `api`, `frontend`)
   - A concise **subject** line (max 72 chars) in imperative mood that describes *what* the change does
   - An optional **body** if the change is non-trivial and benefits from a short explanation of *why*

4. Present the suggested commit message in a code block using the Conventional Commits format:
   ```
   <type>(<scope>): <subject>

   <body — optional>
   ```

5. Briefly explain the reasoning behind your choice of type and subject (1–2 sentences max).

6. Ask the user: **"Would you like me to run `git commit` with this message?"**
   - If yes: run `git commit -m "<message>"` using the exact message suggested (use a heredoc to preserve formatting if there is a body).
   - If no: invite the user to tweak the message or ask for an alternative.

## Guidelines

- Use imperative mood: "add", "fix", "remove" — not "added", "fixes", "removed"
- Do not include issue numbers unless the user mentions them
- Keep the subject line under 72 characters
- If there are no staged changes, tell the user and stop
