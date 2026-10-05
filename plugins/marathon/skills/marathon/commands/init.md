# marathon init

Set up marathon on a repository, once, from a planning concept. `init` first checks that marathon
isn't set up: no `context/` here and no sibling declaring itself coordinator. Settle every decision
in plan mode, in plan rounds, before creating any file.

## 1. Find the concept

The concept is the document the architect points you to, often a `concept.md` or a charter, or
one you write together. It gives the vision and the scope. The structure comes out of planning.

## 2. Settle the founding decisions

- **Vision**: one paragraph on what the project is for.
- **Project kind**: `code` if a build-and-test loop defines the slices and validation, otherwise
  `context` (skills, prose, configuration).
- **The check**: the one deterministic command that runs the build, tests, and lint, or the
  consistency script of a context project (`references/build.md`).
- **The currency command**: optional. The read-only command that reports which direct
  dependencies, toolchain versions, and pins trail their latest release, run at PLAN and never in
  CI (`mechanics/configuration.md`).
- **Capability map**: the project's major capabilities, broad, shallow, and unordered.
- **The first goal**: the one outcome the first sessions work toward, and its first task.
- **Remote**: the platform, its publish command (`gh pr create`, `glab mr create`), and its merge
  command, or none for a local-only project.

## 3. Create the structure

Keep every file brief, a few sentences each.

```
<repo>/
├── CLAUDE.md          # navigation pointers only: context/README.md, STANDARDS.md, the check
├── STANDARDS.md       # judgement calls for the standards-reviewer; starts as a stub
├── .gitignore         # .claude/plans/ and .claude/briefs/
├── .claude/
│   ├── settings.json  # plansDirectory = ./.claude/plans; allow Skill(marathon:marathon)
│   └── marathon.toml  # mechanics/configuration.md
└── context/
    ├── README.md      # vision and capability map
    ├── roadmap.toml   # active = [the first goal]; references/manifest.md
    └── <note>.md      # optional: what the first goal needs to know
```

A member of a workspace has no `roadmap.toml`; its goals live in the coordinator's, and so do a
spike's. A coordinator adds a `[workspace]` table (`references/workspace-coordination.md`).
Extensions are enabled later, when the project adopts them (`references/extensions.md`). Don't
create `docs/`; a task that documents creates it.

The `STANDARDS.md` stub says what the file is for: judgement calls the check can't enforce, each
one line, with pointers to the architecture pages that apply. It never restates what the check
enforces.

## 4. Commit the setup

Fire `on-start`, then commit everything. The architect runs the first goal with `start <goal>`.
