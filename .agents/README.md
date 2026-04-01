# Agents

This folder keeps the repository's reusable AI collaboration artifacts in
an agent-agnostic layout.

## Purpose

- `skills/` contains reusable workflows that an AI coding agent can invoke
  while working in this repo.
- `agents/` contains role definitions for a multi-agent strategy used in
  implementation and review.

The goal is to preserve useful working patterns without coupling the repo
to a single client or vendor-specific directory convention.

## Skills

Skills in this repo serve two purposes:

- repo-specific workflows, such as contract review or backend changes
- generally useful workflows that still reflect this repo's Makefile and
  delivery rules, such as fixing lint issues or suggesting commit messages

Each skill should stay small and practical. When a skill benefits from
more deterministic behavior, it may include scripts or reference files.

Example:

- `skills/calculator-final-check/` includes a script that runs the full
  validation gate used by this project

## Sub-Agents

Sub-agent definitions in `agents/` describe how work can be split when
using a multi-agent strategy:

- `planner` for scoping and execution planning
- `backend-worker` for backend-focused implementation or review
- `reviewer` for findings, drift checks, and final review

## Suggested MCPs

Potential MCPs that fit these skills and agent roles:

- docs/spec lookup for requirements and API contract review
- Git or GitHub workflow support for review and commit-oriented tasks
- filesystem or code-search tooling for fast repository exploration
- OpenAPI-oriented tooling for contract verification
