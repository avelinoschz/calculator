# AI Usage Summary

This project was developed using an AI-assisted workflow to improve productivity, structure, and code quality.

AI was used as a collaborative assistant for:

- translating requirements into structured specifications
- defining architecture and design decisions (ADRs)
- proposing project structure
- guiding implementation sequencing
- generating and refining code
- reviewing the implementation from a senior-level code review perspective
- drafting and improving documentation

All AI-generated or AI-assisted outputs were manually reviewed, validated, and adjusted before inclusion.

The goal was not to rely blindly on AI, but to use it as a tool to accelerate iteration while maintaining full ownership of design decisions and code quality.

## AI Prompts

This document captures a small set of representative prompts used during development.

The goal is not to document every single prompt, but to preserve the most useful prompts that shaped the project structure, implementation flow, and review process.

---

## 1. Initial Structure Prompt

Use this prompt before generating implementation code.

```text
Read the following files:
- AGENTS.md
- specs/calculator/requirements.md
- specs/calculator/plan.md
- specs/calculator/api.md
- api/openapi.yaml
- docs/adr/0001-architecture-and-api.md
- docs/adr/0002-tooling-and-delivery.md

Propose a minimal and clean project structure for:
- Go backend
- React + TypeScript frontend

The structure should:
- follow the documented requirements and ADRs
- prioritize simplicity and maintainability
- avoid overengineering
- separate domain logic from transport logic in the backend
- support the documented developer experience requirements

Do not generate implementation code yet.
Only propose the project structure and explain the reasoning behind it.
```

## 2. Reviewer Prompt

Use this prompt after a meaningful implementation step, such as backend core, frontend core, or a final review pass.

```text
Act as a strict and pragmatic code reviewer.

Read the following files first:
- README.md
- AGENTS.md
- specs/calculator/requirements.md
- specs/calculator/plan.md
- specs/calculator/api.md
- api/openapi.yaml
- docs/adr/0001-architecture-and-api.md
- docs/adr/0002-tooling-and-delivery.md

Then review the current implementation for:
- correctness against requirements
- maintainability
- unnecessary complexity
- API contract mismatches
- missing validation
- missing tests
- frontend/backend inconsistencies
- naming or structure issues

Do not implement changes yet.

Return the review in this format:
1. Critical issues
2. Medium issues
3. Low-priority improvements
4. Items that should be de-scoped
5. Final review of code quality and maintainability
```

## Notes

These prompts are intentionally concise and reusable.

They are meant to demonstrate structured use of AI for:

- planning
- implementation guidance
- review and quality control

Additional prompts used during implementation can be added later if needed.
