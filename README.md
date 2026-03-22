# Calculator

Simple full-stack calculator system built with:

- Backend: Go (net/http)
- Frontend: React + TypeScript

## Overview

This repository contains the specification, API contract, and design decisions for a small full-stack calculator take-home assessment.

The intended implementation is designed with a focus on:

- maintainability
- clarity
- correctness
- engineering judgment under time constraints

Rather than maximizing features, the goal is to deliver a clean, well-structured, and production-minded solution.

## Document Guide

- `specs/calculator/requirements.md` is the source of truth for scope and acceptance criteria.
- `specs/calculator/plan.md` describes the intended implementation sequence.
- `specs/calculator/api.md` is the human-readable API guide.
- `api/openapi.yaml` is the canonical API contract.
- `docs/adr/0001-architecture-and-api.md` and `docs/adr/0002-tooling-and-delivery.md` capture architectural decisions and trade-offs.
- `AGENTS.md` provides implementation guidance for AI-assisted workflows.

## AI-Assisted Development

This repository is being prepared using AI-assisted workflows for:

- specification and planning
- implementation guidance
- code review and refinement

All AI-generated outputs were manually reviewed and validated.

Representative prompts used during development can be found in:

- `docs/ai-prompts.md`

## Target Features

- Addition
- Subtraction
- Multiplication
- Division

## API

- `POST /api/v1/calculations`

This endpoint is defined as part of the current project specification.

See:

- `specs/calculator/api.md`
- `api/openapi.yaml`

## How to Run

Setup and run instructions will be completed alongside the implementation.

They are expected to be provided via:

- Makefile targets
- Docker Compose

## Design

Key design decisions are documented in:

- `docs/adr/0001-architecture-and-api.md`
- `docs/adr/0002-tooling-and-delivery.md`

## Notes

This repository intentionally prioritizes clear scope, maintainable design, and pragmatic decision-making before implementation begins.
