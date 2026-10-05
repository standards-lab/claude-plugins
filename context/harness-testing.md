# Harness testing

This note covers how the plugins in this repository are tested. Two layers run: a cheap check on
every push, and a release gate before each tag that adds a plugin's eval suite.

## Per push: the check

`scripts/check.sh` runs in CI on every push and locally. It checks that:

- each plugin's manifest version, top CHANGELOG heading, and skill `Version:` lines agree
- the `@` pointers and `./` links in the plugins' markdown resolve
- each marketplace source points to a plugin
- `claude plugin validate --strict` passes for each plugin and the marketplace

CI installs Claude Code at an exact version on the stable channel, pinned in
`.github/workflows/ci.yml`. `scripts/currency.sh` reports the pin when it trails npm's `stable`
dist-tag. The check is deterministic and runs no model.

## Per release: the gate

`scripts/gate.sh` is this repository's `[project] gate`. SHIP runs it on the merge commit before
each `<plugin>/v<version>` tag, with the tag as its argument. It runs `scripts/check.sh`, then:

```
claude plugin eval plugins/<plugin> --threshold 1.0 --runs 3 --ablation none --scaffold \
  --trust-plugin --no-publish --allow-tools EnterPlanMode
```

Every case must pass all three runs, or the tag is held. A plugin with no case releases on the
check alone. Evals never run per push or in CI, because each run is a paid session.

## Eval suites

A plugin has a suite only once it has a case. marathon has one, in `plugins/marathon/evals/`;
marathon-architecture has none, and the gate skips it. `results/` is gitignored.

- **Cases come from observed failures.** `retro` turns a ticked plugin eval finding into a case on
  its `retro-<topic>` branch, and the case's `description` cites the failing session or PR
  (marathon's `commands/retro.md`). The fix is its own task; the gate holds the next release
  until it merges.
- **Graders are deterministic where possible.** `tool_used` and `regex` read the same way every
  run. An `llm` grader judges only short output, with its rubric written as concrete PASS and FAIL
  conditions.
- **Fixtures are offline scaffold scripts.** A case that needs a repository builds it with a
  scaffold script. Replaying a recorded session through `history_file` isn't used, because it
  carries the skill text of the session it recorded.

## The seed case

`plan-round-in-reply` runs `start` on a task with no brief and grades the reply: a `PLAN ROUND 1`
header, numbered questions each with `rec:` and `changes:` lines, and questions that stand on
their own. It comes from marathon 0.16 sessions in which rounds asked through AskUserQuestion were
rejected.

It checks the round's shape only. `claude plugin eval` 2.1.289 runs each case as
`claude -p --permission-mode dontAsk` with no permission-prompt tool, which drops
AskUserQuestion, EnterPlanMode, and ExitPlanMode. The case can't reproduce the original failure,
which happened in plan mode. Its AskUserQuestion grader stays for when eval offers that tool.

## Held

- **The standards-reviewer writing into `~/go/pkg/mod`.** A real failure, held until an offline
  fixture can provoke it.
