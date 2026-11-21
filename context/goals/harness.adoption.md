# goal · harness.adoption

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] references
2. [x] go-tools
3. [x] ai-hosting
4. [x] shim
5. [x] architecture

## Decisions

- plan: staged over the locks of v1.ai.processing, v1.client.stack, v1.ai.inference, and
  v1.ai.hosting, a one-time exception that the sibling adoption goals share; no session runs on
  those goals until every adoption goal syncs (standards-lab/context/harness-adoption.md).
- architecture: no CI for architecture; the merge waits on no checks.

## Pending edits

- coordinator · standards-lab: delete `context/harness-adoption.md` and its entry in
  `context/README.md`; delete `roadmap/harness.adoption.toml`, and remove `harness.adoption` from
  `roadmap/index.toml` `active`. The architecture task fixed the review's findings in every
  repository it covered, which would widen the goal's `repos` and `tasks.architecture` summary to
  each of them; the goal file is deleted instead, so the session brief is the record of which
  repositories changed.
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
- coordinator · standards-lab: v1.client.stack's `template` summary adds webappgen's review
  findings (triage OUT-1..3); its `standard` summary adds the web stack's repeated rules and the
  pnpm release-age reasoning as the standard's to hold (WTL-7, WTL-8), and web-components' peer
  range excluding web-foundation v0.4.0 (WCM-7).
- coordinator · standards-lab: `roadmap/sites.toml` `tasks.architecture-site` summary: drop
  "It repoints `architecture/context/README.md`, which names backlog.docs-site …"; the
  architecture task named the site in its capability map.
- coordinator · standards-lab: goal `marathon.repo-access` at the top of `planned`, not staged,
  with tasks `ci-awareness` (a declared per-repository CI flag, so no session waits on CI that
  doesn't exist) and `shared-repos` (a goal owns its root, and a non-root repository locks only
  while off its default branch, so goals queue on it). The goal also carries every claude-plugins
  finding of the review: the triage's CLP items, among them currency.sh's null bug (CLP-1) and the
  marathon-architecture plugin docs (CLP-14, CLP-15), and the mise-action v5.3.0 bump. The triage
  file is gitignored, so the goal copies each item's text rather than citing the file.
- coordinator · standards-lab: apply the review's accepted standards-lab findings, triage STL-1..37
  and STL-40..46, as each item's fix states. STL-38 and STL-39 are the sites.toml and design.md
  edits above.
- claude-plugins: `context/repo-access.md`, the note for `marathon.repo-access`, with its open
  points; a backlog note for anchor checking in lint-references, which also records the rejected
  alternative of a pinned architecture checkout in CI (architecture has no tags, so the pin would
  be a bare SHA that no currency command tracks; `architecture/...` references skip in CI
  instead); `context/harness-optimization.md:52` "Code repositories adopt it in
  `harness.adoption`" becomes "every repository's check runs it".
