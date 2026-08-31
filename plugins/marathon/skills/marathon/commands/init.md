# marathon init

Set up marathon on a repository, starting from a project-planning concept. Run this once per project.
When it's done, the repo has a top-level `context/` directory that the rest of the workflow reads and
maintains, plus the Claude configuration the project needs.

`init` runs the session pipeline (`mechanics/pipeline.md`) on a repository that has none of it yet:
its LOCATE is the check that both marks are absent — no `context/` here, no sibling project declaring
a coordinator — and its CONCLUDE is its own setup commit. `init` commits to the project's founding
decisions, so do the thinking in plan mode and agree on them with the architect before creating any
files.

## 1. Find the concept

Find the project-planning concept: a document the architect points you to (often a `concept.md` or a
charter), or, if there isn't one, work with the architect to write it. The concept gives you the vision
and the rough territory to start from; the structure comes out of planning.

## 2. Settle the founding decisions

Read the concept closely, and settle these with the architect rather than assuming them:

- **Vision** — one paragraph on what the project is for.
- **Project kind** — whether the repository holds production source with a build-and-test loop
  (`code`) or is entirely context the agent authors directly (`context`: skills, prose,
  configuration). This sets how sessions execute. When in doubt: is there a build-and-test loop
  that defines stage boundaries and a validation phase? If there is, it's `code`.
- **Capability map** — the major capabilities the project will need. Broad, not ordered, and shallow;
  you'll add detail later, when a capability is about to be built.
- **Settled vs. concept** — which intent is solid enough to start in `design/`, and which is still a
  concept for `concepts/`. When in doubt, it's a concept.
- **First step** — the one concrete thing the first session will do.
- **Remote platform** — which remote the project publishes to, and the command that proposes a change
  there: `gh pr create` for GitHub, `glab mr create` for GitLab, the equivalent for another platform, or
  none for local-only. This is required; closeout uses it to publish.

Don't create anything until the architect approves the plan.

## 3. Create the structure

Set up the directories and files. Seed them shallow — a few sentences each is right; resist writing
detail up front.

```
<repo>/
├── CLAUDE.md                  # short: says the repo uses marathon; points into context/
├── .claude/
│   ├── settings.json          # plansDirectory; permissions
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
  remote platform with its publish command, in the canonical layout of
  `mechanics/configuration.md`. If this project coordinates a workspace, add the
  `[workspace]` block too (`role = "coordinator"` and a layered `order`); most projects don't. See
  `references/workspace-coordination.md`. Extensions are not an `init` decision: a repository enables
  one whenever the convention is adopted, by adding the `extensions` key (see
  `references/extensions.md`).
- `.claude/settings.json` — set `plansDirectory` to `./.claude/plans` and allow `Skill(marathon:marathon)`.
- `CLAUDE.md` — keep it short: name the workflow and point to `context/README.md` for
  orientation. The kind lives in `marathon.toml` and the skill formalizes what it means; the
  detail goes in `context/`, not here.
- Don't create `experiments/` now; an experiment session makes it when it's needed.

## 4. Commit the setup

Stage everything and make the first commit. The repo is now marathon-managed, and the architect starts
real work with `plan` or `start`.
