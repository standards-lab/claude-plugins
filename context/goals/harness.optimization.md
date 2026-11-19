# goal · harness.optimization

- **State:** building
- **Task:** harness-standards
- **Branch:** harness-standards

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

## Task brief · harness.optimization · harness-standards

**Problem**

claude-plugins has no STANDARDS.md, so the standards-reviewer applies nothing of harness
programming beyond its own profile. No pointer reaches the architecture pages that govern the
repository. The practices marathon held survive only where marathon's own files still express
them: the harness pages the architecture diet deleted, and the Pocock practices 0.16 adopted.
The repository's scripts and configuration carry comments the comments principle doesn't allow.
marathon's configuration template teaches every repository to write commented TOML, and the
coordinator's roadmap holds 229 comment lines. On retro-routes, a session retyped the configured
merge command with `;` for `&&`, and #49 merged with CI red. No rule on main prevented it.

**Behaviors**

1. claude-plugins has a STANDARDS.md of one-line judgement calls:
   - instructions state the present; removals live only in the CHANGELOG
   - a skill programs its own mechanics, never user-scope behavior
   - a skill's configuration and state live in declared files in the repository
   - a command runs through the one pipeline; an extension acts only at hook points
   - SKILL.md indexes its files and never restates them
   - a procedural section in a skill is a candidate tool
   - an instruction that changes no behavior is cut
   - each change gets an Unreleased entry; an extension-contract change bumps the floor
   - an eval case comes from an observed failure

   It points to the architecture layer's comments page for scripts, tests, and every
   configuration file. It also points to rolling-currency, independent-releases,
   context-architecture, and minimal-footprint, and to marathon's "Checks come first". It
   restates nothing the check enforces, and CLAUDE.md points to it.
2. claude-plugins' check fails, naming the fix, when:
   - a skill whose description says "marathon extension" doesn't declare
     `user-invocable: false`
   - a CHANGELOG release heading carries anything after its version
   - a CHANGELOG holds a link definition
3. claude-plugins' scripts and tests carry only the comments the principle allows: each script's
   or module's header. Function doc comments and docstrings are gone, and names carry their
   meaning. The check gives the same failures on the same inputs and passes on main.
4. claude-plugins' configuration files carry only the comments the principle allows. Each key's
   meaning lives in marathon's configuration reference or a script's header. The eval case's
   comments go; its description already gives their reasons.
5. marathon's and marathon-architecture's TOML examples carry no comments, except the roadmap
   index's pointer to its format. init therefore writes an uncommented configuration, and the
   eval fixture's `marathon.toml` and `.gitignore` match it.
6. `marathon <key> [args]` runs the named command from the current repository's marathon
   configuration. The key is one of check, currency, gate, publish, merge, or ci.
   - It runs the command through bash exactly as declared, appends any arguments
     shell-quoted, and runs a multi-line value as one script.
   - It exits with the command's status.
   - It runs only declared commands. An unset key, an unknown key, or a missing configuration
     fails and names the fix.
   - It ships in marathon's `bin/` and uses only Python's standard library.
7. marathon's pipeline Invariants state that every configured command runs through `marathon`,
   never retyped. The steps that run a configured command call it, and the configuration
   reference documents it. No word budget rises.
8. lint-context fails on a comment in a roadmap goal file, naming the fix.
9. claude-plugins' main requires the `check` status in a ruleset no actor can bypass, so a pull
   request whose check is red can't merge.
10. marathon's Unreleased section records:
    - the `marathon` command, as Added
    - the uncommented examples
    - the goal-file rule
    - this repository's new check rules

    None carries an extension-contract marker. marathon-architecture's Unreleased section
    records its uncommented example.
11. The goal record gains a pending edit for release's sync:
    - the coordinator roadmap's header narrative goes to notes
    - its dividers go
    - each goal-attached comment moves to the note it cites, or to its goal's note when it
      cites none
12. The check passes, every word budget holds with none changed, and claude-plugins' currency
    command reports nothing.

**Test seams**

- claude-plugins' check and its unit tests:
  - the two new check rules, each on a fixture tree that breaks it
  - lint-context's goal-file rule
  - `marathon` against scratch repositories whose configuration holds a chained merge, a
    multi-line ci, and a publish that takes arguments
- The check on a scratch copy seeded with each failure it reports: identical output and exit
  code before and after the comment trim.
- This task's own standards review, which applies the branch's STANDARDS.md; its commits cite
  the lines they apply.
- The ruleset, read back from GitHub: `check` required, no bypass actors.

**Slices**

claude-plugins only, with no upgrade slice.

1. STANDARDS.md and CLAUDE.md's pointer to it.
   Demo: each line is one judgement call or pointer, each pointed-to page exists, and the
   check passes.
2. Scripts and tests brought under the comments principle.
   Demo: the scratch copy's failures match before and after, and main passes.
3. Configuration files, the two plugins' TOML examples, and the eval fixture without comments,
   with their CHANGELOG entries.
   Demo: no comment outside the allowed kinds, and init's read-first set is about 190 words
   smaller.
4. The two check rules and lint-context's goal-file rule, with tests and a CHANGELOG entry.
   Demo: each fixture fails with its fix named, and main passes.
5. `marathon`, the Invariants line, the step edits, and its configuration entry, with tests and
   an Added entry.
   Demo: the merge fixture runs exactly as declared, a chained value keeps its `&&`, and every
   read-first set stays within budget.
6. main's ruleset requires `check` with no bypass.
   Demo: the ruleset read back, and this task's own pull request merges only once green.

**Out of scope**

- A hook that refuses a chained `gh pr merge` or `git push`: deferred to the backlog.
- A rule to run each git and gh command in its own call: left to the architect's user-scope
  configuration.
- Folding lint-context or other tools under `marathon` as subcommands: `v1.harness.tooling`.
- Any change to architecture's comments page, and the comments in other repositories'
  configuration: both go to harness.adoption.
- The coordinator roadmap itself, which changes only at release's sync.
- A 100-column wrap check and the read-first budget ratchet: both left for a retro.
- Releasing or tagging: task `release`.

**Door**

Two-way. Nothing is released, every change is a revertible commit, and `marathon` reaches no
session until `release`. The ruleset sits outside the diff and binds every claude-plugins pull
request once set, other sessions' included. Removing it reverses it.

## Progress

slices 0/6 committed · standards — · spec — · editor —

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
- harness-standards: claude-plugins only; STANDARDS.md points to architecture's pages without
  editing them (plan).
- harness-standards: STANDARDS.md is one line per judgement call or pointer; rationale sections
  were rejected as the prose that goes stale (plan).
- harness-standards: staged execution and `@`-loading were dropped from the old harness pages as
  superseded (plan).
- harness-standards: the extension `user-invocable: false` rule and the CHANGELOG heading rule
  are checks, not STANDARDS lines (plan).
- harness-standards: a 100-column wrap check was left for a retro (plan).
- harness-standards: the comment near-ban covers TOML and YAML configuration, which drifts like
  source. The convention lives in the STANDARDS pointer; amending architecture's comments page
  was deferred to harness.adoption (plan).
- harness-standards: script function doc comments and docstrings were removed; names carry
  their meaning (plan).
- harness-standards: main requires a non-bypassable `check` status. `[remote] merge` keeps
  `gh pr checks --watch &&`; `--auto` was rejected because auto-merge is off and it returns
  before the merge (plan).
- harness-standards: configured commands run through `marathon <key>`, built here rather than
  deferred to v1.harness.tooling. `marathon-run`, `marathon-exec`, `marathon-command`, and
  per-table tools were rejected; the root command takes the key as its subcommand (plan).
- harness-standards: the command runs all six keys through bash with appended arguments; refusing
  multi-line values and arguments was rejected (plan).
- harness-standards: a hook refusing chained merges was deferred; a marathon hook fires in every
  session (plan).
- harness-standards: the one-command-per-call rule stays out of marathon as user-scope
  behavior (plan).
- harness-standards: the comment-standard Decision stands; lint-context's goal-file rule shapes
  the manifest and is not a comment checker for source (plan).
- harness-standards: release's sync moves goal-attached roadmap comments into the note each cites;
  a block-by-block judgement pass was rejected (plan).

## Pending edits

- standards-lab · `.claude/marathon.toml`: set `[project] check` to `lint-context`, applied at
  `release`'s sync after the roadmap split (content-budgets).
- standards-lab · `context/roadmap.toml`: at release's sync, move the header narrative into notes, drop the section
  dividers, and move each goal-attached comment into the note it cites, or its goal's note when it
  cites none (harness-standards).
