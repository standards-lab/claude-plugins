# goal · marathon.repo-access

- **State:** brief ready
- **Task:** review-findings
- **Branch:** review-findings

## Tasks

1. [x] ci-awareness
2. [x] read-first-split
3. [ ] review-findings
4. [ ] shared-repos

## Task brief · review-findings

**Problem**

The `harness.adoption` review left fourteen claude-plugins findings: notes describe a test suite,
CI pins, and history that no longer match the repository; rules are stated twice or undefined;
marathon-architecture's layer reference misses the spike pointer form and the self-containment
rule. The read-first budget rule lives only in the goal record, and the notes' figures are wrong.

**Behaviors**
1. The harness-testing note says the seed case has five graders and no line-length grader, and
   requires at least one numbered question with a `rec:` line then a `changes:` line.
2. The harness-testing note says CI installs what the development and CI tool files pin, the CI
   file read only under `MISE_ENV=ci`, with the reason Claude Code is pinned there, and that the
   currency script reports any tool or GitHub Action pin trailing its latest release.
3. The CI tool file's header reads "The tools CI adds to mise.toml's, read only where
   MISE_ENV=ci is set.", keeps the rolling-currency reference, adds one to the harness-testing
   note, and `lint-references` resolves both.
4. The harness-testing note names no eval version: "`claude plugin eval` runs each case as".
5. The root README's tree lists every tracked top-level entry (CI workflow, CI tool file,
   standards file, `scripts/`, `.gitignore`) and doesn't say CI installs only the development
   tool file.
6. The completed harness.optimization goal's note is renamed `harness-findings.md` and holds only
   what has no other home: the Brandon Thomas source behind `door-scaled-review` and the findings
   for later. The capability map points harness.optimization at marathon's own files.
7. The tool-based-skills note lists a CLI for LOCATE and `status` as a first consumer.
8. The marathon-extraction note's first paragraph defines "the blueprint" and "a consumer of the
   standard"; its open questions drop the extension-enablement and adding-references questions.
9. The standards file's comments line ends at "(TOML, YAML, `.gitignore`)"; "Plugin design lives
   in the skill" appears only in the agent orientation file.
10. The standards file states: "A read-first budget rises only with the architect's approval,
    with the reason in the Unreleased entry."
11. The check fails when a word budget sits above its set's or profile's size rounded up to the
    next hundred, naming the budget to set; a test covers failing and passing; the check's header
    and the harness-testing check list name the rule; marathon's Unreleased has an entry.
12. marathon-architecture's layer reference says a spike outside the workspace directory writes
    its pointer relative to the architecture repository, as `principles/<page>.md`.
13. marathon-architecture's README section on what the layer holds becomes a link to the layer
    reference, which states the self-containment rule under what belongs in a layer.
14. marathon-architecture's CHANGELOG has an Unreleased section covering 12 and 13; no version
    changes.
15. The repository's check passes.

**Test seams**
- The repository's check (links, citations, budgets, `lint-context`, `lint-references`, its test
  suite), plus a coherence read of each prose change.

**Slices**
1. Budget ceiling in the check (11).
2. CI pins and eval text: harness-testing note and CI tool file (1–4).
3. Budget rule to the standards file; harness-optimization budget findings dropped (6 part, 10).
4. harness-optimization tidy-up and tool-based-skills (6 rest, 7).
5. Orientation: standards comments line, context README, root README (5, 9).
6. marathon-extraction note (8).
7. marathon-architecture: README, layer reference, CHANGELOG (12–14).
8. The harness note renamed and reduced; the capability map repointed (6, redirect).

**Out of scope**

Tags and version bumps; any budget value; marathon skill files; architecture's README; the eval
case's provenance text; the layer's starting index template.

**Door**

Two-way: prose, one check rule, and comments; nothing tagged.

## Progress

slices 8/8 committed · standards ✓ · spec ✓ (after the redirect) · editor ✓ (before it)
- redirect, brief: harness-optimization.md is a completed goal's note; reduce it to what has no
  other home, and rename it to fit its purpose? yes

## Decisions

- goal: one release, at shared-repos: `marathon/v0.23.0` and `marathon-architecture/v0.4.1`;
  ci-awareness and review-findings tag nothing.
- shared-repos: moves claude-plugins' own configuration to `[remote] checks` with the release,
  since the installed 0.22.0 copy and binary don't know the key.

## Pending edits

- coordinator · roadmap/marathon.repo-access.toml: end the `review-findings` summary's first
  sentence with "and the stale read-first sizes in claude-plugins' harness note".
- coordinator · roadmap/v1.harness.toml: `door-scaled-review`'s `context` becomes
  `claude-plugins/context/harness-findings.md`.
