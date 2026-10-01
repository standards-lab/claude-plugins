# Harness testing

This note covers how the plugins in this repository are tested. Plugin evals are their automated
checks, the first of the three quality layers in marathon's `references/build.md`. The work is
tracked as `factory.evals`.

## Per push: consistency

`scripts/check.sh` runs in CI on every push. It checks that:

- the version numbers agree
- the `@` pointers and `./` links resolve
- each marketplace source points to a plugin

It stays cheap, deterministic and always on.

## Per release: plugin evals

`claude plugin eval` runs a plugin against a suite of cases and scores each one. Each case is a
realistic prompt plus one or more graders. It is available in this environment.

- **Layout.** Each plugin keeps its suite in `plugins/<name>/evals/`, one directory per case, with
  `prompt.md`, `graders/`, and a `case.yaml` and scaffold script when the case needs a fixture
  repository. `results/` is gitignored.
- **Cases come from failures.** A case is added when a real session fails in a way the plugin
  should prevent. `retro` sends that failure here (marathon's `commands/retro.md`). A case
  written before any failure would test a guess, not a regression.
- **Graders are deterministic where possible.** `tool_used`, `tool_order`, `regex` and
  `file_exists` cost nothing and read the same way every run. An `llm` grader judges only short
  output, with its rubric written as concrete PASS and FAIL conditions. Each case grades both
  the result and the steps that produced it.
- **The release gate.** The release script runs `scripts/check.sh` and then
  `claude plugin eval --threshold 1.0`. A plugin doesn't release while a case fails. Evals
  don't run on every push, because each run costs model calls.

## The seed suite

The seed suite covers the contract marathon 0.16 introduces:

- **implementer-no-standards:** the implementer never reads `STANDARDS.md` (`tool_used Read`,
  `input_match STANDARDS`, `min 0`, `max 0`).
- **reviewer-commits:** the standards-reviewer commits its fixes (`tool_used Bash`,
  `input_match git commit`).
- **brief-shape:** the session brief has Summary, Core changes, Evidence and Merge danger
  (`regex`).
- **plan-round-format:** plan-round questions are numbered, each with a recommendation
  (`regex`).

## Assumptions

- Fixture repositories made by a scaffold script are enough to exercise a whole task without
  network access.
- The no-plugin baseline means little for a workflow plugin. Cases run with `--ablation none`
  unless they test whether a skill fires.
