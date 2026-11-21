# goal · harness.adoption

- **State:** building
- **Task:** architecture
- **Branch:** architecture

## Tasks

1. [x] references
2. [x] go-tools
3. [x] ai-hosting
4. [x] shim
5. [ ] architecture

## Task brief · architecture

**Problem**

architecture has no pin putting marathon's lints on PATH; its check duplicates lint-references and
never runs lint-context, which fails on the missing capability map; its configuration carries
comments outside the five kinds; comments.md says no code repository follows it and depends on
other repositories' current state. 28 repositories' context changed in one sweep with no review
across them before the adoption note, which holds the settled rules, is deleted.

**Behaviors**
1. Every comment in architecture's tracked scripts and configuration is one of the five kinds;
   STANDARDS.md's comment rule covers every configuration file.
2. `context/README.md` has a capability map: built entries linked (the architecture and its
   principles, the standards, the page schema and conventions, the check); planned: site hosting
   (`sites`, `architecture-site`), the Web Elemental standard (v1.client.stack `standard`), the
   development-time secrets principle (v1.ai.hosting `secrets`). lint-context passes.
3. `mise install` puts marathon's tools at the latest release on PATH through the pin, with
   `task.run_auto_install = false`.
4. The currency command prints one line per pin behind its latest release and exits nonzero, or
   prints nothing and exits 0.
5. The check keeps its link and front-matter checks, drops its own reference loop, runs
   lint-context and lint-references (both, even when the first fails), and fails when any part
   fails; a missing lint fails with `FAIL: <lint> is not on PATH; run mise install`. Its header,
   task description, STANDARDS.md and CLAUDE.md name what it runs.
6. The README's Conventions state that a page's claims hold regardless of any other repository's
   current state; it links a repository or tool only as an example. comments.md follows it: each
   repository's check resolves every reference, naming no tool; "Which code follows the
   principle" is gone.
7. comments.md states eight settled rules: every configuration file is covered; a configuration
   header is a short description plus `See` lines; docs on unexported identifiers and tests are
   not package documentation; each ecosystem's meaning of exported (Rust `pub`, C#
   public/protected, CSS custom properties, classes, layers); a data file's header (SQL) is its
   package documentation; lint-required comments (`// SAFETY:`) stay; licence notices a ported
   file requires stay; provenance is one line per source with a short "what changed" clause.
8. A Fable review covers all 28 workspace repositories once 1–7 land; each finding the architect
   accepts at triage is fixed in its repository, each fixed repository's check passes, and
   nothing beyond accepted findings changes. standards-lab findings land through the sync.

**Test seams**
- `marathon check` in architecture and each repository that receives fixes.
- lint-context and lint-references output; `marathon currency` in architecture.

**Slices**
1. architecture comments (1). 2. Capability map (2). 3. Pin, currency, check (3–5).
4. Conventions and comments.md (6, 7). 5. Fable review, 12 read-only agents, one triage
   escalation (8). 6. Upgrades for fixed repositories that trail. 7. Fix slices per repository,
   coordinator order, architecture last.

**Out of scope**

CI for architecture, shellcheck, replacing the link and front-matter checks, anchor checking in
lint-references (backlog note), spike evidence content, the marathon adjustments themselves
(recorded only), releases.

**Door**

Two-way: prose, configuration, one script; no tags.

## Progress

slices 4/4 (architecture) committed · review done: 361 findings, triage escalated (.claude/plans/triage-harness.adoption-architecture.md) · standards — · spec — · editor —

## Decisions

- plan: staged over the locks of v1.ai.processing, v1.client.stack, v1.ai.inference, and
  v1.ai.hosting, a one-time exception that the sibling adoption goals share; no session runs on
  those goals until every adoption goal syncs (standards-lab/context/harness-adoption.md).
- plan: the rest of the adoption runs as 14 sibling `harness.adoption-<slug>` goals, in one
  orchestrating session once `shim` releases marathon/v0.22.0; the rules they share live in
  standards-lab/context/harness-adoption.md.
- plan: architecture's `scripts/check.sh` switches to marathon's `lint-references` (the
  `architecture` task).
- references: rejected a pinned architecture checkout in CI, so `architecture/...` references
  skip there: architecture has no tags, so it would pin a bare SHA no currency command tracks.
- architecture: no CI for architecture; the merge waits on no checks.
- architecture: a page's claims hold regardless of another repository's current state.
- architecture: the Fable review's fixes land inside this task, while the lock exception holds;
  the goal's `repos` widen to each fixed repository through the sync.

## Pending edits

- coordinator · standards-lab: marathon 0.22.0, the release the adoptions pin, in place of 0.21.0:
  `context/harness-adoption.md`'s opening sentence, and the `harness.adoption-*` goal files'
  names and summaries.
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
- coordinator · standards-lab: `roadmap/v1.client.stack.toml` `tasks.template` summary restates
  the pattern in place of citing the adoption note: "It adopts marathon's tools at their latest
  release: mise pins them, the check runs lint-context and lint-references, `context/README.md`
  has a capability map, and every comment is one of `architecture/principles/comments.md`'s five
  kinds."
- coordinator · standards-lab: goal `marathon.repo-access` (tasks `ci-awareness`: a declared
  per-repository CI flag so no session waits on CI that doesn't exist; `shared-repos`: a goal owns
  its root, and non-root repositories lock only while off their default branch, so goals queue on
  them) at the top of `planned`, not staged.
- coordinator · standards-lab: `roadmap/sites.toml` `tasks.architecture-site` summary: drop
  "It repoints `architecture/context/README.md`, which names backlog.docs-site …"; the
  architecture task named the site in its capability map.
- claude-plugins: `context/repo-access.md`, the note for `marathon.repo-access`, with its open
  points; a backlog note for anchor checking in lint-references;
  `context/harness-optimization.md` "Code repositories adopt it in `harness.adoption`" becomes
  "every repository's check runs it".
