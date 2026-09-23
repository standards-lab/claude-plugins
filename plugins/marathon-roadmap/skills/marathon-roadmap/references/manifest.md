# The roadmap manifest

`context/roadmap.toml` holds only what remains on the path to the target end state.

## Structure

- Root goals sit under `[goals.<slug>]`, and child goals nest directly under them
  (`[goals.v1.data]`).
- A goal's tasks sit under its reserved `tasks` table (`[goals.v1.data.tasks.reads]`).
- Backlog tasks, which belong to no goal, sit under the root `backlog` table (`[backlog.<slug>]`).
- `next` is a root-level list and holds the file's only sequences. The rest of the tree has no
  order. Each entry of `next` is either a dotted path to a task or goal, or an array that forms a
  wave, whose members run at the same time (marathon's `mechanics/reset-file.md`). A wave member
  is a dotted path, or a nested array of paths: a lane, whose steps run in order. The first entry
  is what comes next.

An entry's full key is its identity. Renaming a slug means updating every reference to it.

## Fields

A goal has these fields:

- `name`: the outcome, as a title.
- `summary`: what it means for this outcome to hold.
- `criteria` (optional): the statements that close the goal once they all hold.
- `context` (optional): the context files that carry the detail.

A task has these fields:

- `name`: the requirement, as a title.
- `summary`: what the task is, sized to one session's work.
- `repos` (optional): the repositories it touches.
- `proof` (optional): the observable result that shows it is done.
- `context` (optional): the context files that carry the detail.

`context` entries are file paths: relative to the workspace in a workspace
(`go-web-service/context/data-layer.md`), and relative to the repository otherwise.

## Citing entries

Everything outside the manifest cites a task or goal by its dotted path, leaving out the `goals`
and `tasks` segments: `v1.data.reads`, `backlog.docs-site`. Slugs never contain dots.

## Lifecycle

- **Only what remains**: a finished task is deleted, and a goal whose criteria hold is deleted
  with it. The session record keeps the disposition. A stale entry is a defect, and the session
  that finds it fixes it.
- **No order outside `next`**: `next` and each lane are the only sequences. The lanes of a wave
  have no order among themselves.
- **Detail only at the front**: the task at the front of `next` carries detail. Every other
  entry stays a short statement, with its detail in the linked context files.

## Starting file

`on-start` creates this manifest, and the architect fills it in when the session settles its scope:

```toml
# Roadmap: what remains on the path to the target end state. Maintained by the marathon-roadmap
# extension; the format is the plugin's references/manifest.md.

next = []
```

## Example

```toml
# Roadmap: what remains on the path to the target end state. Maintained by the marathon-roadmap
# extension; the format is the plugin's references/manifest.md.

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
summary = "The documentation site that serves the project's published pages."
context = ["context/docs-site.md"]
```
