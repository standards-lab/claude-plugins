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
- **Project kind** — whether the repository holds production source the developer authors (`code`) or
  is entirely context the agent authors directly (`context`: skills, prose, configuration). This sets
  the role boundary and whether sessions hand off an implementation guide. When in doubt, ask whether
  `git blame` on the real deliverable should point at the developer (`code`) or at the agent under the
  developer's review (`context`). See `references/role-boundary.md`.
- **Capability map** — the major pieces the project will need. Broad, not ordered, and shallow; you'll
  add detail later, when a piece is about to be built.
- **Settled vs. concept** — which intent is solid enough to start in `design/`, and which is still a
  concept for `concepts/`. When in doubt, it's a concept.
- **First step** — the one concrete thing the first session will do.
- **Remote platform** — which remote the project publishes to, and the command that proposes a change
  there: `gh pr create` for GitHub, `glab mr create` for GitLab, the equivalent for another platform, or
  none for local-only. This is required; closeout uses it to publish.
- **Project-management extension** — optionally, whether to turn on a project-management extension now
  (issues, boards) or stay standalone. The core works without one, and you can add it later.

Don't create anything until the developer approves the plan.

## 3. Create the structure

Set up the directories and files. Seed them shallow — a few sentences each is right; resist writing
detail up front. Write the prose per the voice standard in `references/writing-voice.md`.

```
<repo>/
├── CLAUDE.md                  # short: says the repo uses marathon; points into context/; states the role boundary
├── .claude/
│   ├── settings.json          # plansDirectory; permissions; any chosen extension
│   └── marathon.toml          # project kind; remote platform + publish command
└── context/
    ├── README.md              # vision + capability map
    ├── design/                # settled-intent notes (shallow)
    ├── concepts/              # concepts: unsettled ideas (shallow)
    └── reset.md               # Status: closeout; Next-focus = the first step
```

- `context/reset.md` — write it with `Status: closeout`, `Session: init`, a Disposition listing what
  was scaffolded, and the first step as Next-focus, so the first `start` begins from a closeout and
  creates a new branch.
- `.claude/marathon.toml` — record the project kind (`[project] kind = "code"` or `"context"`) and the
  remote platform with its publish command. If this project coordinates a workspace, add the
  `[workspace]` block too (`role = "coordinator"` and a layered `order`); most projects don't. See
  `references/workspace-coordination.md`.
- `.claude/settings.json` — set `plansDirectory` to `./.claude/plans` and allow `Skill(marathon:marathon)`.
  If you chose a project-management extension, enable its plugin and allow its skill and tools here too;
  otherwise leave it out.
- `CLAUDE.md` — keep it short: name the workflow, point to `context/README.md` for orientation, and
  state the role boundary for the project's kind in plain, language-neutral terms — for a `code`
  project, the developer owns the production code and the agent writes tests, documentation, and
  context; for a `context` project, the agent authors the repository directly under the developer's
  review. The detail goes in `context/`, not here.
- Don't create `context/guide.md`; it isn't a scaffolded file. On a `code` project a `start` creates it
  when the session needs it and closeout deletes it. A `context` project never has one.
- Don't create `experiments/` now; an experiment session makes it when it's needed.

If you set up a platform extension, run its `on-init` hook so it can establish its outward side (creating
the project's board, for example). With no extension, this does nothing.

## 4. Commit the setup

Stage everything and make the first commit; this fires the `on-commit` hook. The repo is now
marathon-managed, and the developer starts real work with `plan` or `start`.
