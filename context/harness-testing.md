# Harness testing

This note covers how the plugins in this repository are tested. A cheap check runs on every push,
and a release gate runs the check and the plugin's eval suite before each tag.

## Per push: the check

`scripts/check.sh` runs in CI on every push and locally. It checks that:

- each plugin's manifest version, top CHANGELOG heading, and skill `Version:` lines agree
- no plugin's markdown contains an `@`-include line, which Claude Code expands only in CLAUDE.md
- the `./` and `../` links in the plugins' markdown resolve
- the backtick file citations in the plugins' markdown resolve, including `${CLAUDE_PLUGIN_ROOT}`
  paths and paths another plugin owns, as in marathon's `references/briefs.md`. A bare path counts
  as a citation only when its first segment is a directory in a plugin or skill root, and a
  CHANGELOG's released sections are skipped.
- each command's read-first set and each agent profile stay within the word budget recorded in
  `scripts/budgets.txt`. A command's set is its skill's SKILL.md, the playbook, and every file
  reached through "Read first" lists from there. A file with no budget fails.
- each skill description is at most 1,024 characters
- each marketplace source points to a plugin
- `claude plugin validate --strict` passes for each plugin and the marketplace
- marathon's `lint-context`, run from the working tree, passes on this repository's `context/`,
  READMEs, and links, and its tests in `scripts/tests/` pass

CI installs Claude Code at an exact version, its latest release, pinned in
`.github/workflows/ci.yml`. `scripts/currency.sh` reports the pin when it trails npm's `latest`
dist-tag. The check is deterministic and runs no model.

## Per release: the gate

`scripts/gate.sh` is this repository's `[project] gate`. SHIP runs it on the merge commit before
each `<plugin>/v<version>` tag, with the tag as its argument. It runs `scripts/check.sh`, then:

```
claude plugin eval plugins/<plugin> --threshold 1.0 --runs 3 --ablation none --scaffold \
  --trust-plugin --no-publish
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

`plan-round-in-reply` runs `start` on a task whose note leaves two independent decisions open, and
grades the final reply with six graders:

- AskUserQuestion is never called.
- The reply contains `PLAN ROUND 1`.
- The round doesn't sit inside a code fence.
- No line of the round exceeds 100 characters.
- At least two numbered questions each carry a `rec:` line and a `changes:` line.
- An `llm` grader judges that every question names its subject and can be decided on its own. It
  finds the round by the line that contains `PLAN ROUND 1`.

The case comes from marathon 0.16 sessions in which the architect rejected rounds shown as
AskUserQuestion headlines, and a question asked without its context.

The case checks the round's shape in the reply, not the original failure. `claude plugin eval`
2.1.289 runs each case as `claude -p --permission-mode dontAsk` with no permission-prompt tool,
which drops AskUserQuestion, EnterPlanMode, and ExitPlanMode, so the AskUserQuestion grader
passes whatever the skill does. The grader stays for when eval offers that tool.

## Failures without a case

- **The standards-reviewer writing into `~/go/pkg/mod`.** This failure happened in a real session.
  Its case waits until an offline fixture can provoke it.
