# marathon-extraction: carrying proven patterns out of a consumer

marathon-extraction is a planned marathon extension, tracked as `v1.harness.extraction`. It
records the patterns a consumer of the standard proves, and it carries that record back to the
blueprint's context and documentation. The consumer can be a workspace member or a repository
the workspace can never see.

## The problem it solves

A pattern proven inside the workspace reaches the blueprint through the normal session loop:
the session's notes, the documentation step, and the roadmap. That loop breaks when the pattern
is proven in a consumer outside the workspace:

- The record has to cross by hand. A production consumer may sit on a network the blueprint
  can't reach, where nothing about it may be cataloged publicly. What it proved has to become a
  self-contained, sanitized prose record before anything else happens to it.
- Awareness runs downward. No member repository, and nothing in the coordinator's public context,
  may name a consumer. Knowledge flowing from a consumer to the blueprint therefore has to pass
  through sessions and artifacts that only the coordinator holds.

## What the extension carries

The extension carries a description of each pattern and never a decision about it. Each pattern
is stated once, in general terms, as the consumer implemented it. No code crosses. Whether a
pattern moves into a library or the template is decided the way the workspace already decides
it: a note or a documentation page names the pattern, and a later session builds it.

## Design

The extension has three parts:

- **Capture, on the consumer side.** The consumer enables it in its `marathon.toml`. It owns
  `context/extraction.toml`, a ledger of the patterns the consumer nominates. At `on-close`, the
  closing session compares the step's changes with the want-list (below) and with any nomination
  markers, and drafts ledger entries. The architect accepts or rejects each one during the close.
- **Package, on the consumer side.** The architect invokes it. It renders the accepted entries
  into a bundle: one self-contained markdown file with the entries and a manifest of the sources
  they came from. The bundle is what crosses the boundary by hand. Package refuses an entry that
  has no `evidence`, so the judgment of what is general is made where the code is.
- **Intake, on the blueprint side.** The architect invokes it at the coordinator against a
  bundle. It turns each entry into work, not pages: a task under the roadmap entry that the
  entry's `satisfies` field cites, or a backlog task and a note when the entry cites nothing. The
  bundle is stored under `private/` beside the private references catalog, and each task cites it
  in its `context`. Intake never edits a note, a documentation page, a library, or the template.
  The sessions that run the new tasks do that work.

This design makes marathon-extraction a hybrid extension: capture fires at a hook, and package
and intake are invoked by the architect (see the taxonomy in `marathon-sitrep.md`).

### Ledger entries

Each ledger entry and each bundle entry has these fields:

- `kind`: `design`, for a pattern that belongs in a module's documentation or an architecture
  page, or `convention`, for an organization or standard principle that no single module owns.
- `claim`: the pattern, stated as a documentation page would state it, with no domain nouns.
- `evidence`: why the pattern belongs to the standard and not to the consumer's domain. For
  example, status filters, search, and hierarchy queries belong to the domain, never to the SQL
  library.
- `satisfies`: the dotted path of the coordinator's roadmap entry the pattern serves, when one
  exists.
- `origin`: the consumer's key in the private references catalog and the commit that proved the
  pattern, so packaging the same ledger twice never duplicates an entry.
- `sources`: the references keys of anything else the pattern came from.

### Adding sources to the references catalog

A bundle names its sources: the consumer, and any prior repository or external example the
pattern came from. Intake checks each source against the references catalog and adds any that
are missing, following the catalog's file contract (the header of
`standards-lab/references.toml`):

- A private source goes into `private/references.toml` (its key and remote) and
  `private/references.md` (what it is and what to draw from it). A public source goes into the
  public catalog.
- Keys share one namespace. Intake never redefines a key that already resolves. When it can't
  find a remote for a source, it adds the entry without one and flags it in the session record
  for the architect.
- Intake never writes `references.local.toml`, because local paths belong to one machine.

A consumer's key is created by the first intake that carries it, so a consumer exists on the
blueprint side only once intake runs. Once `v1.harness.references` is built, intake calls it to
add sources. Until then, intake edits the catalog files directly.

### The want-list

Capture nominates well only when the consumer knows what the blueprint is waiting for. The
consumer may read the blueprint's remaining goals, because it sits downstream, but copying them
into its context would duplicate the roadmap. Instead, the extension generates a want-list file
from `context/roadmap.toml`. The file is stamped with the roadmap's source commit, regenerated
at each crossing, never edited by hand, and gitignored. When both workspaces are on the same
machine, the extension reads the roadmap directly through its references key, and no file is
generated.

### Nomination markers

The cheapest signal is a marker comment in the consumer's code, such as `// extract:`, placed
on a helper an SDK should own, a reworked middleware, or a forked template element. `on-close`
collects the markers and drafts entries from them. Without markers, the session nominates from
the step's diff against the want-list, which is noisier and gives the architect more to review.

### Tooling

Bundle schema validation, the `evidence` check, origin de-duplication, adding references, and
generating the want-list are deterministic. They belong in a tool the skill calls
(`v1.harness.tooling`), per `architecture/harness/tool-based-skills.md`. The skill's text keeps
only the judgment of what is general and where it belongs.

## The first consumer

go-web-service is already a consumer. `v1.data.evaluation` needs a record of the data layer
before it can decide what moves into the libraries. Running capture on a ledger at the
coordinator, with a workspace member as the origin and nothing crossing a boundary, tests the
ledger and the `on-close` nomination before any bundle exists. Package, private origins, adding
references, and intake are built when a bundle first needs to cross.

## Open questions

- Where the ledger lives: at the coordinator's `context/extraction.toml` when the consumer is a
  workspace, matching the roadmap, or in the project itself when it is standalone.
- Whether intake may read a ledger on the same machine directly, or only ever a bundle. The
  design assumes a bundle and treats the same-machine case as a special case of it.
- Whether the standard states the marker convention on a per-module page, or each consumer
  chooses its own.
- How the extension records that it is enabled and which marathon version it targets.
  `marathon-sitrep.md` has the same questions, and one answer serves both.
- Whether adding references belongs to intake or to the references extension that intake calls.
