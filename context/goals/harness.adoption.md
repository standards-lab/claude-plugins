# goal · harness.adoption

- **State:** brief ready
- **Task:** ai-hosting
- **Branch:** ai-hosting

## Tasks

1. [x] references
2. [x] go-tools
3. [ ] ai-hosting
4. [ ] core
5. [ ] docproc
6. [ ] go-ai
7. [ ] go-libraries
8. [ ] templates
9. [ ] web
10. [ ] go-web-service
11. [ ] architecture
12. [ ] spikes

## Task brief · ai-hosting

**Problem**

ai-hosting is the first repository to adopt marathon 0.21.0, and the later adoption tasks copy
its pattern. Its check doesn't run marathon's lints, its design note is about four times the note
budget, its capability map has no built/planned markers, fifteen comment lines fall outside
architecture's five kinds, and its orientation files point at a roadmap file that no longer
exists.

**Behaviors**
1. `mise install` puts lint-context and lint-references from marathon 0.21.0 on PATH through the
   repository's pin; the currency command reports a newer marathon tag when one exists.
2. The check runs shellcheck, the existing link check over every tracked Markdown file,
   lint-context, and lint-references, and fails when any fails. When either lint isn't on PATH the
   check fails and names `mise install` as the fix. The check's header and task description name
   all four.
3. lint-context passes: the design note becomes one note per planned capability, each under 1,500
   words (engines, the `hosting` suite, host-class profiles, gateway, secrets, Foundry embedding
   overflow, cleanup), each linked from a `(planned)` capability map entry; every map entry is
   marked `(planned)`; each point left open for a task's plan sits in that capability's note.
4. The findings table is gone; each settlement only it stated sits in its task's note, with no
   dangling finding IDs.
5. The Foundry note keeps its prices, figures, and source URLs, with no dates.
6. Every comment in tracked code and configuration is one of the five kinds: the 15 lines are
   gone, each script section becomes a named function, planned work moves to its note, no `docs/`
   is created, and lint-references passes.
7. The standards file's comment rule covers every configuration file (TOML, YAML, `.gitignore`),
   and its source line names the notes in place of the design note.
8. The orientation files point at `standards-lab/roadmap/index.toml`.
9. In the `v1.ai.hosting` goal record, the `design:` decision and the P14 pending edit name the
   notes that now hold their subjects; its `secrets` brief is unchanged byte for byte.
10. CI runs the check after `mise install` and passes on the task's pull request.

**Test seams**
- The check (`mise run check`), with lint-context and lint-references inside it.

**Slices**
1. Comments: five kinds only, section comments become named functions, the standards comment
   rule covers configuration (behaviors 6, 7).
2. Context: design note split, capability map, Foundry dates, findings table removed,
   `v1.ai.hosting` record reworded, roadmap pointers; lint-context passes by hand (3, 4, 5, 7, 8,
   9).
3. Pin and check: the mise pin, the check calling both lints, CI green (1, 2, 10).

**Out of scope**

The coordinator edits to standards-lab's ai-hosting note, docs page, and `v1.ai.hosting` goal
file (pending edits until sync); any `secrets` work; a `docs/` directory; replacing the existing
link check.

**Door**

Two-way: prose, configuration, and scripts in one repository, with no release.

## Progress

slices 3/3 committed · standards ✓ (4f9364b) · spec ✓ · editor ✓
- slice 1, decided: currency.sh `[ -n ] && report` became an `if` so a function returns 0 when nothing trails.
- slice 2, decided: a 10th map entry, "One hosting repository per workstation" (planned), links `cleanup.md`; P14's label dropped with the finding IDs.
- slice 3, decided: currency.sh unchanged; mise's versions host lags a fresh tag (a stderr warning), yet the newer tag still reports.

## Decisions

- plan: staged over the locks of v1.ai.processing, v1.client.stack, v1.ai.inference, and
  v1.ai.hosting, a one-time exception; no session runs on them until harness.adoption and then
  v1.ai.hosting's `secrets` task ship.
- plan: architecture's `scripts/check.sh` switches to marathon's `lint-references` (the
  `architecture` task).
- plan: the 33 archived spikes stay as they are; only the 4 unarchived spikes adopt.
- plan: webappgen sits on its local, unpublished `template` branch: its adoption branches from
  main, and once merged, main merges into `template`. go-storage's modified `go.work.sum` is
  committed with its adoption.
- references: rejected a pinned architecture checkout in CI, so `architecture/...` references
  skip there: architecture has no tags, so it would pin a bare SHA no currency command tracks.
- adoption: comments outside the five kinds go where the code already says them (a key name, a
  mise task description) or where they describe planned work; a script's section comments become
  named functions; a reference comment points where the reasoning has a home, and a new `docs/`
  page holds only reasoning with none. In ai-hosting this removes 15 lines and creates no `docs/`.
- adoption: each repository's `STANDARDS.md` comment rule covers every configuration file (TOML,
  YAML, `.gitignore`), as claude-plugins' does.
- ai-hosting: after ai-hosting ships, the order is (1) a marathon 0.21.1 task fixing `bin/`'s shims,
  whose `MISE_OVERRIDE_CONFIG_FILENAMES`/`MISE_OVERRIDE_TOOL_VERSIONS_FILENAMES` reach the command
  `marathon` runs, so `marathon check` and `marathon currency` ignore the project's mise config;
  (2) a `plan` session settling the goal-wide questions (unexported and test doc comments,
  nolintlint, PostgreSQL 19, over-budget READMEs and built-work notes to `docs/`), splitting tasks
  (go-libraries ×4, web ×3, spikes ×4, docinfer and docproc-comments added, webappgen to
  v1.client.stack's `template`), and splitting the goal into sub-goals so tasks run concurrently;
  (3) the concurrent run. Pre-plans: `.claude/plans/preplan-harness.adoption-*.md` (local).

## Pending edits

- coordinator · standards-lab: repoint each mention of ai-hosting's removed `context/design.md`
  at the ai-hosting note that now holds its subject:
  - `context/ai-hosting.md`, "What `ai-hosting` holds", the suite's design:
    `ai-hosting/context/hosting-suite.md`.
  - `docs/ai-hosting.md`, the opening paragraph's design note (each task's spike parts and open
    points): ai-hosting's capability notes, mapped in `ai-hosting/context/README.md`; "The spikes'
    answers", e5's scope and price on Foundry: `ai-hosting/context/foundry-overflow.md`; "The
    `gateway` task", the six open points: `ai-hosting/context/gateway.md`, "Open at the gateway
    plan".
  - `roadmap/v1.ai.hosting.toml`: `tasks.secrets`' "design's note": `ai-hosting/context/secrets.md`;
    `tasks.gateway`'s "what design's note defers to it": `ai-hosting/context/gateway.md`, "Open
    at the gateway plan"; `tasks.cleanup`'s "the design note's inventory":
    `ai-hosting/context/cleanup.md`. `tasks.design`'s summary describes the note that task wrote,
    and stays.
