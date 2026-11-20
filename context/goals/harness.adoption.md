# goal · harness.adoption

- **State:** building
- **Task:** references
- **Branch:** references

## Tasks

1. [ ] references
2. [ ] ai-hosting
3. [ ] core
4. [ ] docproc
5. [ ] go-ai
6. [ ] go-libraries
7. [ ] templates
8. [ ] web
9. [ ] go-web-service
10. [ ] architecture
11. [ ] spikes

## Task brief · references

**Problem**

The comment principle lets source carry a reference comment pointing a file at the document that
governs it, and it requires every repository's check to resolve those references. Only
architecture's check does so today, and only for `#` comments in its scripts. No code repository
can adopt the rule until marathon ships one check that any repository's check and CI can call.
The goal also moves the roadmap out of `context/` to a root `roadmap/` directory with
`index.toml`. The coordinator's roadmap moves in a plan session after this release installs, so
the release lands first. Every later adoption task depends on both.

**Behaviors**
1. A new marathon check, `lint-references`, runs on the repository it starts in and reads every
   git-tracked file that isn't markdown. A line is a reference line when its first non-blank
   characters are `#`, `//`, `--`, `/*`, or `*`. Every `*.md` path on such a line is a
   reference, except a path inside a URL.
2. A reference resolves from the repository root first. If nothing in the repository matches
   its first segment, it resolves in the sibling directory of that name beside the repository.
   A standalone project's top-level `architecture/` resolves as an ordinary path; in a
   workspace, `architecture/...` resolves against the architecture repository beside it.
3. A reference that resolves in neither place fails with
   `<file>:<line>: reference <path> does not resolve; fix the path or remove the reference`.
   This covers a missing path in the repository or in an existing named sibling, and a
   one-segment path missing at the root.
4. When the first segment names nothing in the repository and no sibling of that name exists,
   as in CI, the check skips the reference and prints
   `<file>:<line>: skipped <path>: no <segment>/ in this repository or beside it`. Skips alone
   pass. The check exits nonzero only when a reference fails. A file that isn't UTF-8 is skipped
   silently.
5. `lint-references` ships beside `lint-context` with the same requirements: stdlib Python 3.11,
   and on PATH in a session. marathon documents the two together: every repository's check
   calls both, and a CI without the plugin downloads both from one pinned `marathon/vX.Y.Z`
   tag, shown with a v0.20.0 example.
6. claude-plugins' own check runs `lint-references` and its tests. Any broken reference it finds
   is fixed, and the check passes.
7. marathon reads the roadmap from a root `roadmap/` directory: `roadmap/index.toml` for the
   three arrays and `roadmap/<goal>.toml` for each goal. This holds at the coordinator and in a
   standalone project, everywhere the skill names, reads, writes, or creates the roadmap,
   including LOCATE, `status`, `plan`, `intake`, `init`, and sync. There is no fallback.
8. `lint-context` checks the root `roadmap/` exactly as it checked the roadmap in `context/`:
   - the index's arrays and its two-line comment limit;
   - goal files carry no comments;
   - `repos`, `order`, and the summary, proof, and criteria budgets.

   A file in `roadmap/` that is neither `index.toml` nor a `<goal>.toml` fails and names its fix.
   `context/` no longer counts roadmap files among its kinds.
9. `lint-context` fails on a `context/roadmap.toml` with
   `context/roadmap.toml: the roadmap index moved; move it to roadmap/index.toml`. It fails on
   each `context/roadmap/<goal>.toml` with `move it to roadmap/<goal>.toml`.
10. The `plan-round-in-reply` eval builds its fixture on the new layout, and the release gate
    passes.
11. marathon's CHANGELOG `v0.20.0` lists the new check and the roadmap move. Its
    `### Migrating from 0.19.0` has numbered steps:
    1. Install 0.20.0.
    2. At the coordinator or in a standalone project, move the index and every goal file into
       the root `roadmap/`, and update the notes that name the old paths.
    3. In each repository, add `lint-references` to the check, and have CI download it beside
       `lint-context` at `marathon/v0.20.0` or later.
    4. Run it, then fix or remove each reference that fails.
12. The plugin's version, the skill's Version line, and the CHANGELOG all read 0.20.0, and the
    read-first and description budgets still hold.

**Test seams**
- `lint-references`, run on a temporary repository and its siblings; the tests compare stdout,
  stderr, and the exit status.
- `lint-context`, run on a temporary repository with a root `roadmap/` or the old layout.
- claude-plugins' check and release gate, for the skill and the eval fixture.

**Slices**
1. Upgrade: CI's Claude Code pin 2.1.328 → 2.1.329; the currency command exits 0 and the check
   passes.
2. The reference check (behaviors 1–4, 6): the command, its tests, and claude-plugins' check
   calling it, with broken references fixed.
3. The reference check documented (behavior 5): configuration, the pinned CI download, `init`'s
   check, and `Unreleased` CHANGELOG entries with migration steps 3 and 4.
4. `lint-context` follows the roadmap (behaviors 8, 9), with its tests.
5. The skill follows the roadmap (behaviors 7, 10): every skill file, the README, the eval
   scaffold, and the CHANGELOG entry with migration step 2.
6. The release (behaviors 11, 12): `Unreleased` becomes `v0.20.0` with migration step 1, the
   versions agree, the budgets hold, and the gate passes.

**Out of scope**

- Moving the coordinator's roadmap (a `plan harness.adoption` session, after install).
- Switching architecture's check to `lint-references` (the `architecture` task).
- Every other repository's check and CI (the adoption tasks).
- A currency command that reports a trailing `marathon/vX.Y.Z` download pin.
- Any marathon-architecture release.

**Door**

The `marathon/v0.20.0` tag is one-way, with the layout change it ships: once pushed, a roadmap
still in `context/` fails `lint-context`. The work before the tag is two-way.

**Release**

- `marathon/v0.20.0`

## Progress

slices 6/6 committed · standards ✓ · spec ✓ · editor ✓

## Decisions

- plan: staged over the locks of v1.ai.processing, v1.client.stack, v1.ai.inference, and
  v1.ai.hosting, a one-time exception; no session runs on them until harness.adoption and then
  v1.ai.hosting's `secrets` task ship.
- plan: the coordinator's roadmap moves to `roadmap/index.toml` in a `plan harness.adoption`
  session right after `references` ships and is installed, not at sync.
- plan: architecture's `scripts/check.sh` switches to marathon's `lint-references` (the
  `architecture` task).
- plan: the 33 archived spikes stay as they are; only the 4 unarchived spikes adopt.
- plan: webappgen sits on its local, unpublished `template` branch: its adoption branches from
  main, and once merged, main merges into `template`. go-storage's modified `go.work.sum` is
  committed with its adoption.
- references: rejected a pinned architecture checkout in CI, so `architecture/...` references
  skip there: architecture has no tags, so it would pin a bare SHA no currency command tracks.

## Pending edits

- none
