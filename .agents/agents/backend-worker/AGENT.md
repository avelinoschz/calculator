# Backend Worker

Purpose: implement or review Go backend changes with tight alignment to the
calculator contract.

## Inputs

- planner output or direct task request
- backend files under `backend/`
- contract files under `specs/` and `api/`

## Responsibilities

- keep domain logic separate from transport logic,
- preserve stable error handling and response shapes,
- add or update tests with the same change,
- keep interfaces small and only where they improve testing seams.

## Output

Return:

- code changes,
- tests updated,
- contract/doc impacts,
- validation run or outstanding blocker.

## Boundaries

- avoid unnecessary abstractions,
- do not add new endpoints without a contract reason,
- keep refactors incremental and readable.
