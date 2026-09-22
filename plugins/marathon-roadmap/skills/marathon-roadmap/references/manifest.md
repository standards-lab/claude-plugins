# The roadmap manifest

`context/roadmap.toml` holds only what remains on the path to the target end state.

## Structure

- Root goals sit under `[goals.<slug>]`; child goals nest directly (`[goals.v1.data]`).
- A goal's tasks live under its reserved `tasks` table (`[goals.v1.data.tasks.reads]`).
- Backlog tasks — assigned to no goal — live under the root `backlog` table
  (`[backlog.<slug>]`).
- `next`, a root-level list, holds the file's only sequences, with each lane; the tree itself is
  unordered. Each entry is a dotted path to a task or goal, or an array forming a wave whose
  members run at the same time (marathon's `mechanics/reset-file.md`). A wave member is a dotted
  path, or a nested array of paths: a lane, run in order. The first entry is what comes next.

The composed key is the identity; renaming a slug updates every reference with it.

## Fields

Goal:

- `name` — the outcome, as a title.
- `summary` — the claim: what holding this outcome means.
- `criteria` (optional) — the claims that, all holding, close the goal.
- `context` (optional) — linked context files carrying the detail.

Task:

- `name` — the requirement, as a title.
- `summary` — one session's worth of work: what it is.
- `repos` (optional) — the repositories it touches.
- `proof` (optional) — the observable result that shows it done.
- `context` (optional) — linked context files carrying the detail.

`context` entries are file paths: workspace-relative in a workspace
(`go-web-service/context/data-layer.md`), repository-relative otherwise.

## Dotted citations

Everything outside the manifest cites a task or goal by its dotted slug path, omitting the
structural `goals` and `tasks` segments: `v1.data.reads`, `backlog.docs-site`. Slugs never contain
dots.

## Lifecycle

- **Ephemeral** — a finished task is deleted, and a goal whose criteria hold goes with it; the
  session record keeps the disposition. A stale claim is a defect the finding session fixes.
- **Unordered** — `next` and each lane are the only sequences; a wave's lanes have no order among
  themselves.
- **Proximate** — the task in front carries detail; everything else stays a claim, its depth in
  linked context files.

## Bootstrap

A manifest created at `on-start`, populated with the architect when the session settles scope:

```toml
# Roadmap — the path to the target end state, holding only what remains. Maintained by
# the marathon-roadmap extension; the format is the plugin's references/manifest.md.

next = []
```

## Worked example

```toml
# Roadmap — the path to the target end state, holding only what remains. Maintained by
# the marathon-roadmap extension; the format is the plugin's references/manifest.md.

next = [
  [["v1.data.reads", "v1.data.writes"], "v1.messaging"],
  "backlog.docs-site",
]

[goals.v1]
name = "The service at v1.0"
summary = "The service reaches v1.0.0 with every capability layer complete."
criteria = [
  "Every capability goal beneath this one is closed.",
]

[goals.v1.data]
name = "Data layer"
summary = "A composed data model over plain SQL with a CQRS-oriented interface."
context = ["service/context/data-layer.md"]

[goals.v1.data.tasks.reads]
name = "Reads"
repos = ["service"]
summary = "The reads slice: query vocabulary, pagination, and the first domain package."
proof = "Paginated, filtered, sorted reads over HTTP."

[goals.v1.data.tasks.writes]
name = "Writes"
repos = ["service"]
summary = "Commands over the reads slice's first domain package."

[goals.v1.messaging]
name = "Messaging"
summary = "A messaging library and the reactor services it enables."

[backlog.docs-site]
name = "The docs site"
summary = "The documentation site serving the landing zone's content."
context = ["context/docs-site.md"]
```
