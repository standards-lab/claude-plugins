# marathon init

Set up marathon on a repository, once, from a planning concept. `init` runs the session pipeline
(`mechanics/pipeline.md`) with two differences: its LOCATE checks that marathon isn't set up (no
`context/` here and no sibling declaring itself coordinator), and its CONCLUDE is the setup
commit. Settle every decision in plan mode before creating any file.

## 1. Find the concept

The concept is the document the architect points you to, often a `concept.md` or a charter, or
one you write together. It gives the vision and the scope. The structure comes out of planning.

## 2. Settle the founding decisions

- **Vision**: one paragraph on what the project is for.
- **Project kind**: `code` if a build-and-test loop defines the stage boundaries and validation,
  otherwise `context` (skills, prose, configuration).
- **Capability map**: the project's major capabilities, broad, shallow, and unordered.
- **First step**: the one concrete thing the first session will do.
- **Remote**: the platform and its publish command (`gh pr create`, `glab mr create`), or none for
  a local-only project.

## 3. Create the structure

Keep every file brief, a few sentences each.

```
<repo>/
├── CLAUDE.md          # names the workflow; points to context/README.md
├── .gitignore         # .claude/plans/ and .claude/report.md
├── .claude/
│   ├── settings.json  # plansDirectory = ./.claude/plans; allow Skill(marathon:marathon)
│   └── marathon.toml  # mechanics/configuration.md
└── context/
    ├── README.md      # vision and capability map
    ├── <note>.md      # optional: what the first steps need to know
    └── reset.md       # Status: closeout, Session: init, Next-focus: the first step
```

`marathon.toml` records the project kind and the remote. An experiment adds an `[experiment]`
table (`commands/experiment.md`), and a coordinator adds a `[workspace]` table
(`references/workspace-coordination.md`). Extensions are enabled later, when the project adopts
them (`references/extensions.md`). Don't create `docs/`. A documentation step creates it.

## 4. Commit the setup

Commit everything. The architect starts the real work with `plan` or `start`.
