# goal · harness.optimization

- **State:** idle
- **Task:** none
- **Branch:** none

## Tasks

1. [x] version-floor
2. [x] skill-footprint
3. [x] roadmap-layout
4. [x] content-budgets
5. [x] retro-routes
6. [x] evidence-artifacts
7. [x] comment-standard
8. [ ] harness-standards
9. [ ] release

## Task brief · harness.optimization · comment-standard

**Problem**

The architecture layer's only comment rule is Go's. tests-and-docs' "doc.go and godoc" section
asks for a doc comment on every production symbol and limits nothing else, and no check enforces
it. The web repos follow their own, looser rule, "comments only for a non-obvious why", which
sits outside the layer. Comments are where agents write justifications, and those go stale
(Lauren Tan). The goal settled a near-ban: source keeps package documentation, provenance lines,
and the reasons that lint suppressions and pin holds require. Each ecosystem's stock lint rules
enforce what they can, and review covers the rest. marathon's standards-reviewer sweeps only for
comments that excuse a workaround. It needs to apply whatever comment rule the repository
declares.

**Behaviors**

1. The architecture layer has an architecture-level comments principle. It allows exactly four
   kinds of comment in source:
   - package documentation: a package or module comment (a script's or module's header counts
     as one) and a doc comment on each exported identifier
   - the reason a lint suppression requires, on the suppression
   - the reason a pin hold requires, beside the pin
   - a provenance line naming the source repository and commit

   Any other comment is removed, or the code changes so it needs none.
2. For each ecosystem, the principle names the stock rule that requires a suppression's reason:
   golangci-lint's nolintlint (require-explanation, require-specific), typescript-eslint's
   ban-ts-comment, eslint-comments' require-description, and clippy's
   allow_attributes_without_reason. It states that the standards-reviewer judges every other
   comment. It names no custom tool.
3. Both of architecture's indexes list the principle, and architecture's STANDARDS.md points to
   it.
4. go-elemental's Go pages agree with the principle without restating it. tests-and-docs keeps
   the doc.go inventory and the one-contract-per-symbol rule, and links the principle for which
   comments are allowed. Doc comments are required on exported identifiers, not on every
   production symbol. No other Go page states a comment rule of its own.
5. architecture's check carries only the comments the principle allows. It behaves as before:
   the same failures on the same inputs, and a pass on main.
6. marathon's standards-reviewer flags each comment that the repository's declared comment rule
   doesn't permit. Where the repository declares no rule, it flags a comment that excuses a
   workaround, as it does today. The implementer's rule and delegation's editor rule stay as
   they are.
7. marathon's CHANGELOG records the reviewer change under Unreleased, with no extension-contract
   marker. The contract doesn't cover that profile, and marathon-architecture only gives the
   reviewer the layer's location.
8. Both repositories' checks pass. Every word budget holds with no budget changed, and
   claude-plugins' currency command reports nothing.

**Test seams**

- Each repository's check: links, front matter, citations, and budgets.
- architecture's check, run on a scratch copy with a broken link and a page missing front
  matter. It reports the same failures before and after the comment trim.
- The branch's standards-reviewer, engaged through `claude -p --plugin-dir` on a fixture diff
  with one disallowed comment. Run it twice: in a repository whose STANDARDS.md declares the
  principle, and in one that declares no rule. It flags the comment in the first and leaves it
  alone in the second, unless the comment excuses a workaround.

**Slices**

The slices follow the coordinator's order: claude-plugins first, architecture last. There is no
upgrade slice.

1. claude-plugins: the standards-reviewer applies the repository's declared comment rule, and
   the CHANGELOG gets its Unreleased entry. Demo: the two-repository reviewer fixture, with the
   check passing within budget.
2. architecture: the comments principle, listed in both indexes and pointed to from
   STANDARDS.md. Demo: the page reads on its own, and the check passes.
3. architecture: go-elemental's Go comment wording aligned with the principle by linking it,
   with doc comments narrowed to exported identifiers. Demo: no Go page restates the near-ban,
   and the check passes.
4. architecture: the check script brought under the principle. Demo: the scratch copy's failures
   match before and after, and main passes.

**Out of scope**

- Everything harness.adoption owns:
  - each Go and web repository's STANDARDS.md and comments
  - enabling the named rules in `.golangci.yml` and the ESLint configs
  - docproc's hold list and hold comments
  - web-tooling's `eslint-disable` that gives no reason
- claude-plugins' own scripts and lint-context, which go to task harness-standards.
- Any custom comment checker or Go analyzer (round 1, question 6).
- Releasing or tagging, which is task `release`.

**Door**

Two-way. Nothing is released, and each change is a page or profile edit that can be reverted.
The reviewer change reaches no session until `release`. Once the architecture pages merge, they
apply to new Go diffs at once, through tests-and-docs (round 2).

## Decisions

- Nothing releases until the last task, `release`; sessions run the installed 0.18.0 until then
  (plan).
- harness-standards: harness programming principles live in claude-plugins' STANDARDS.md, not the
  architecture layer, whose pages need more than one repository to express them; the task
  evaluates this goal's practices, those before it, and the Pocock practices 0.16 adopted (plan).
- version-floor: unreleased changes sit under `## Unreleased` in each CHANGELOG; `release` renames
  them to marathon v0.19.0 and marathon-architecture's next version and releases both together,
  which satisfies the 0.19.0 floor.
- roadmap-layout: GitHub Issues were rejected for the roadmap, for losing PR-reviewed planning,
  offline sessions, and checks; Markdown and HTML were rejected because the roadmap is data that
  LOCATE, staging, lint-context, and the planned CLI parse (plan).
- roadmap-layout: nested goal directories, a full `[goals.<path>]` table in a goal file, and
  ancestor inheritance of `root` and `repos` were rejected (plan).
- roadmap-layout: a GitHub Project mirror of the cross-goal view is deferred to
  `v1.harness.tooling` (plan).
- roadmap-layout: the coordinator's roadmap is split during this goal's final sync, which also
  moves the index's narrative header into notes; marathon 0.19.0 reads only the split layout (plan).
- content-budgets: CI without the plugin fetches `lint-context` from the pinned `marathon/vX.Y.Z`
  tag; skipping when absent was rejected. The fetch can't be verified until `release` tags (plan).
- content-budgets: budget options and per-repository overrides were rejected (plan).
- retro-routes: retro reads redirects and escalations from the session brief in the pull request
  body; a goal-record section, which sync deletes, and transcripts alone were rejected (plan).
- retro-routes: a citation's first directory fails only when one edit from a real one; an
  allowlist and owner-marking were rejected. No retro eval case (plan).
- evidence-artifacts: an artifact is text in the session brief; stored files were rejected
  because the brief is gitignored and deleted at SHIP (plan).
- comment-standard: moving provenance lines into the context README was rejected (plan).
- comment-standard: a hold list read by the currency command, docproc's form, was rejected for
  pin-hold reasons (plan).
- comment-standard: a portable comment checker beside lint-context and a go/ast analyzer for
  free comments were rejected; stock lint rules and standards review enforce comments (plan).
- comment-standard: making the near-ban marathon's default in every repository was rejected; the
  standards-reviewer applies the repository's declared rule (plan).
- comment-standard: claude-plugins' scripts adopt the near-ban in harness-standards (plan).

## Pending edits

- standards-lab · `.claude/marathon.toml`: set `[project] check` to `lint-context`, applied at
  `release`'s sync after the roadmap split (content-budgets).
