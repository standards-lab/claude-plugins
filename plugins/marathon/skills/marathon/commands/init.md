# marathon init

Set up marathon on a repository, once, from a project-planning concept. `init` runs the session
pipeline (`mechanics/pipeline.md`): its LOCATE checks that neither mark exists (no `context/` here,
no sibling declaring a coordinator), and its CONCLUDE is the setup commit. Settle everything in
plan mode before creating any file.

## 1. Find the concept

The document the architect points you to, often a `concept.md` or a charter, or one you write
together. It gives the vision and the territory; the structure comes out of planning.

## 2. Settle the founding decisions

- **Vision** — one paragraph on what the project is for.
- **Project kind** — `code` if a build-and-test loop defines stage boundaries and validation;
  otherwise `context` (skills, prose, configuration).
- **Capability map** — the major capabilities, broad, unordered, and shallow.
- **First step** — the one concrete thing the first session will do.
- **Remote** — the platform and its publish command (`gh pr create`, `glab mr create`), or none
  for local-only.

## 3. Create the structure

Seed everything shallow, a few sentences each.

```
<repo>/
├── CLAUDE.md          # names the workflow; points to context/README.md
├── .gitignore         # .claude/plans/ and .claude/report.md
├── .claude/
│   ├── settings.json  # plansDirectory = ./.claude/plans; allow Skill(marathon:marathon)
│   └── marathon.toml  # mechanics/configuration.md
└── context/
    ├── README.md      # vision + capability map
    ├── <note>.md      # optional: intent the first steps need, each stating its settledness
    └── reset.md       # Status: closeout, Session: init, Next-focus = the first step
```

`marathon.toml` records the kind and remote. An experiment adds `[experiment]`
(`commands/experiment.md`); a coordinator adds `[workspace]`
(`references/workspace-coordination.md`). Extensions are enabled later, whenever adopted
(`references/extensions.md`). Don't create `docs/`; a documentation step establishes it.

## 4. Commit the setup

Commit everything. The architect starts real work with `plan` or `start`.
