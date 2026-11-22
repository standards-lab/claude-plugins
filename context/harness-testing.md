# Harness testing

This note covers how the plugins in this repository are tested. A cheap check runs on every push,
and a release gate runs the check and the plugin's eval suite before each tag.

## Per push: the check

`scripts/check.sh` runs in CI on every push and locally. It checks that:

- each plugin's manifest version, top CHANGELOG heading, and skill `Version:` lines agree
- each CHANGELOG release heading carries only `## vX.Y.Z` or `## Unreleased`, and no CHANGELOG
  holds a link reference definition
- no plugin's markdown contains an `@`-include line, which Claude Code expands only in CLAUDE.md
- the `./` and `../` links in the plugins' markdown resolve
- the backtick file citations in the plugins' markdown resolve, including `${CLAUDE_PLUGIN_ROOT}`
  paths and paths another plugin owns, as in marathon's `references/briefs/plan-round.md`. A bare
  path counts as a citation only when its first segment is a directory in a plugin or skill root; a
  first segment one edit from such a directory fails as a typo. A citation inside a skill also
  resolves against its plugin's root, and a CHANGELOG's released sections are skipped.
- each command playbook has exactly one "Read first" line
- each command's read-first set and each agent profile stay within the word budget recorded in
  `scripts/budgets.txt`. A command's set is its skill's SKILL.md, the playbook, and every file
  reached through "Read first" lists from there. A file with no budget fails.
- no word budget in `scripts/budgets.txt` sits above its read-first set's or profile's size
  rounded up to the next hundred
- each command's read-first set contains every brief file (`references/briefs.md` or a file
  under `references/briefs/`) that its playbook, or a brief format in the set, cites
- each skill description is at most 1,024 characters
- each skill whose description calls it a marathon extension declares `user-invocable: false`
- each marketplace source points to a plugin
- `claude plugin validate --strict` passes for each plugin and the marketplace
- no tracked file is Python
- the `marathon/` Go module passes the Go check: build, vet, gofmt, `go fix`, `go mod tidy`,
  `go test -race`, and golangci-lint
- marathon's `lint-context` and `lint-references`, built from the working tree, pass on this
  repository
- the tests in `scripts/tests/` pass: `node --test` files for the eval graders' patterns, and bash
  scripts for the check, the `bin/` shims, the gate, and the release build

CI installs the exact versions that `mise.toml` and `mise.ci.toml` pin. mise reads `mise.ci.toml`
only where `MISE_ENV=ci` is set, which CI's workflow (`.github/workflows/ci.yml`) does and a
developer's machine doesn't. Claude Code is pinned in `mise.ci.toml` alone, so a session run from a
developer's machine keeps its own claude. `scripts/currency.sh` reports any tool or GitHub Action
pin that trails its latest release. The check is deterministic and runs no model. Each failure
message names what to change for the check to pass.

## Per release: the gate

`scripts/gate.sh` is this repository's `[project] gate`. SHIP runs it on the merge commit before
each `<plugin>/v<version>` tag, with the tag as its argument. It runs `scripts/check.sh`, then:

```
claude plugin eval plugins/<plugin> --threshold 1.0 --runs 3 --ablation none --scaffold \
  --trust-plugin --no-publish
```

For a marathon tag, the gate also cross-builds the release binaries after the check, and builds
the host's binaries into `plugins/marathon/.build/` so the eval's sessions run the working
tree's tools. It removes both when it exits.

The gate rests on deterministic graders alone. eval keeps no judge reasoning and can't re-grade a
stored reply, so an `llm` verdict can't be reproduced: eval's sonnet judge failed a round in all
three runs that the same model, given the same criteria outside eval, passed. eval has no
non-scoring grader under `--ablation none`, where `arm: with-only` still scores and `weight` must be
positive, so a plugin keeps its judges in a second eval directory the gate doesn't run.

Every case must pass all three runs, or the tag is held. A plugin with no case releases on the
check alone. Evals never run per push or in CI, because each run is a paid session.

## Eval suites

A plugin has a suite only once it has a case. marathon has one, in `plugins/marathon/evals/`, and
its judged suite in `plugins/marathon/evals-judged/`; marathon-architecture has none, and the gate
skips it. Each suite's `results/` is gitignored.

- **Cases come from observed failures.** `retro` turns a ticked plugin eval finding into a case on
  its `retro-<topic>` branch, and the case's `description` cites the failing session or PR
  (marathon's `commands/retro.md`). The fix is its own task; the gate holds the next release
  until it merges.
- **Gated graders are deterministic.** `tool_used` and `regex` read the same way every run. An
  `llm` grader goes in the judged suite, judges only short output, and states concrete PASS and
  FAIL conditions. It runs on sonnet, since eval's default judge, haiku, flipped verdicts.
- **Fixtures are offline scaffold scripts.** A case that needs a repository builds it with a
  scaffold script. Replaying a recorded session through `history_file` isn't used, because it
  carries the skill text of the session it recorded.

## The seed case

`plan-round-in-reply` runs `start` on a task whose note leaves two independent decisions open, and
grades the final reply with five graders:

- AskUserQuestion is never called.
- The reply contains `PLAN ROUND 1`.
- The round doesn't sit inside a code fence.
- At least one numbered question carries a `rec:` line and then a `changes:` line.
- No numbered question points outside itself with "above", "below", "option" and a capital
  letter, "Facts found", "see the plan file", or "details to follow". Its `node --test` test runs
  the pattern on the stored rounds in `scripts/tests/fixtures/plan-rounds/`.

`plan-round-judged`, in the judged suite, runs the same fixture with one `llm` grader, which
judges that every question names its subject and options and can be decided on its own. Run it
by hand with `claude plugin eval plugins/marathon --eval-dir evals-judged --ablation none
--scaffold --trust-plugin --no-publish --judge-model sonnet`.

The case comes from marathon 0.16 sessions in which the architect rejected rounds shown as
AskUserQuestion headlines, and a question asked without its context.

The case checks the round's shape in the reply, not the original failure. `claude plugin eval`
runs each case as `claude -p --permission-mode dontAsk` with no permission-prompt tool, which
drops AskUserQuestion, EnterPlanMode, and ExitPlanMode, so the AskUserQuestion grader passes
whatever the skill does. The grader stays for when eval offers that tool.

## Failures without a case

- **The standards-reviewer writing into `~/go/pkg/mod`.** This failure happened in a real session.
  Its case waits until an offline fixture can provoke it.
