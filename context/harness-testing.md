# Harness testing

This note covers what is worth testing about a workflow skill, and where those tests belong in
CI. The deeper evaluation is tracked as `v1.harness.testing`. It waits for real failure modes to
appear, and for `claude plugin eval` to be available in this environment.

## What CI checks

CI validates what can be checked mechanically: that version numbers agree, that file references
resolve, and that each marketplace source points to a plugin. `scripts/check.sh` runs these
checks. They stay cheap, deterministic, and always on.

## Behavioral testing

Behavioral tests of the skill have no place in CI until an evaluation justifies them. The
evaluation answers four questions:

- Which failure modes have occurred in real sessions.
- Whether a test fixture can catch each one cheaply.
- What a test run costs.
- Whether `claude plugin eval` is the right tool to run the fixtures.

Fixtures written before any failure has occurred test guesses, not regressions. The workflow
changes through real use, so the first fixtures come from failures that actually happen.

## Assumptions

- Real sessions surface concrete failure modes, and each one is worth fixing in the skill before
  any is worth encoding as a fixture.
