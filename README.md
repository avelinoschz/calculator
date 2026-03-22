# Calculator

Simple full-stack calculator system built with:

- Backend: Go (net/http)
- Frontend: React + TypeScript

## Overview

This project is a small full-stack calculator designed with a focus on:

- maintainability
- clarity
- correctness
- engineering judgment under time constraints

Rather than maximizing features, the goal was to deliver a clean, well-structured, and production-minded solution.

## AI-Assisted Development

This project was developed using AI-assisted workflows for:

- specification and planning
- implementation guidance
- code review and refinement

All AI-generated outputs were manually reviewed and validated.

Representative prompts used during development can be found in:

- `docs/ai-prompts.md`

## Features

- Addition
- Subtraction
- Multiplication
- Division

## API

- `POST /api/v1/calculations`

See:

- `specs/calculator/api-contract.md`
- `api/openapi.yaml`

## How to Run

Instructions will be provided via:

- Makefile targets
- Docker Compose

## Design

Key design decisions are documented in:

- `docs/adr/0001-architecture-and-api.md`
- `docs/adr/0002-tooling-and-delivery.md`

## Notes

This project intentionally prioritizes simplicity and completeness over feature expansion, as a reflection of real-world pragmatic engineering decisions and simplicity.
