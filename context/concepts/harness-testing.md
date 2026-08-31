# Harness testing

What is functionally helpful to test about a living workflow skill, and how it should fit the CI
pipeline. Deferred from the v0.9.0 session, where a first cut of per-command
`claude plugin eval` fixtures was authored and then dropped as speculative: the workflow evolves
through real engagement, and fixtures written ahead of observed failure modes test guesses, not
regressions.

The position on record:

- CI's job is tangible validation: version numbers aligned and correct, file references
  resolving, marketplace sources present — what `scripts/check.sh` does today. It stays cheap,
  deterministic, and always on.
- Behavioral testing of the skill, if it earns a place, needs its own evaluation: which failure
  modes have actually occurred, whether a fixture can catch them cheaply, what the run cost is,
  and whether `claude plugin eval` (early access; not yet enabled for this environment) is the
  right harness.

## Assumptions

- Assumes real-world regressions will surface concrete failure modes worth fixing in place
  before any are worth encoding as fixtures.
