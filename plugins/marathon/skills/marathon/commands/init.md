# marathon init

Set up marathon on a repository, starting from a project-planning concept. Run this once per project.
When it's done, the repo has a top-level `context/` directory that the rest of the workflow reads and
maintains, plus the Claude configuration the project needs.

`init` commits to the project's initial shape, so do the thinking in plan mode and agree on it with the
developer before creating any files.

## 1. Find the concept

Find the project-planning concept: a document the developer points you to (often a `concept.md` or a
charter), or, if there isn't one, work with the developer to write it. The concept gives you the vision
and the rough territory. It's the starting point, not the finished structure.

## 2. Plan it out first

Enter plan mode, read the concept closely, and settle these with the developer rather than assuming them:

- **Vision** — one paragraph on what the project is for.
- **Capability map** — the major pieces the project will need. Broad, not ordered, and shallow; you'll
  add detail later, when a piece is about to be built.
- **Settled vs. candidate** — which intent is solid enough to start in `design/`, and which is still a
  candidate for `concepts/`. When in doubt, it's a candidate.
- **First step** — the one concrete thing the first session will do.
- **Remote platform** — which remote the project publishes to, and the command that proposes a change
  there: `gh pr create` for GitHub, `glab mr create` for GitLab, the equivalent for another platform, or
  none for local-only. This is required; closeout uses it to publish.
- **Project-management extension** — optionally, whether to turn on a project-management extension now
  (issues, boards) or stay standalone. The core works without one, and you can add it later.

Don't create anything until the developer approves the plan.

## 3. Create the structure

Set up the directories and files. Seed them shallow — a few sentences each is right; resist writing
detail up front.

```
<repo>/
├── CLAUDE.md                  # short: says the repo uses marathon; points into context/; states the role boundary
├── .claude/
│   ├── settings.json          # plansDirectory; permissions; any chosen extension
│   └── marathon.toml          # remote platform + publish command
└── context/
    ├── README.md              # vision + capability map
    ├── design/                # settled-intent notes (shallow)
    ├── concepts/              # candidate notes (shallow)
    ├── guide.md               # placeholder; the first development session fills it in
    └── reset.md               # Status: closeout; Next-focus = the first step
```

- `context/reset.md` — write it with `Status: closeout` and the first step as Next-focus, so the first
  `start` begins from a clean handoff and creates a new branch.
- `.claude/marathon.toml` — record the remote platform and its publish command from the planning step.
- `.claude/settings.json` — set `plansDirectory` to `./.claude/plans` and allow `Skill(marathon:marathon)`.
  If you chose a project-management extension, enable its plugin and allow its skill and tools here too;
  otherwise leave it out.
- `CLAUDE.md` — keep it short: name the workflow, point to `context/README.md` for orientation, and
  state the role boundary in plain, language-neutral terms (the developer owns the production code; the
  agent writes tests, documentation, and context). The detail goes in `context/`, not here.
- Don't create `experiments/` now; an experiment session makes it when it's needed.

If you set up a platform extension, run its `on-init` hook so it can establish its outward side (creating
the project's board, for example). With no extension, this does nothing.

## 4. Commit the setup

Stage everything and make the first commit. The repo is now marathon-managed, and the developer starts
real work with `start`.
