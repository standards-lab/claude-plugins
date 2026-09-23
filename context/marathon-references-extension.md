# marathon-references: the references catalog as an extension

marathon-references is a proposed marathon extension, tracked as `v1.harness.references`. It
would own the references catalog that this workspace keeps at its coordinator as a convention of
its own.

## The catalog today

The workspace coordinator keeps three files:

- `references.toml` maps one key per repository to that repository's canonical remote.
- `references.local.toml` is gitignored and maps the same keys to local checkouts on one machine.
- `references.md` describes what each repository is: the repositories the effort builds, external
  references, and prior work.

The contract between the files is stated in the header of `standards-lab/references.toml`.
marathon itself never mentions the catalog.

## Why it generalizes

Any marathon workspace draws on repositories beyond its members: prior work to reuse, external
examples, and related efforts. Each one needs a portable identity plus a map of where it is
checked out on this machine, so the workspace can be rebuilt anywhere. The roadmap manifest has
the same shape: a file with a schema, owned by an extension.

## Proposal

The extension owns the three files and resolves them the way marathon-roadmap resolves its
manifest: at the coordinator in a workspace, or in the project itself when it is standalone. It
enforces the file contract: keys join the files, a location is never duplicated between the two
TOML files, and the local file stays gitignored. It prescribes nothing about what a workspace
catalogs.

## Open questions

- Which kind of extension it is: one that fires at hooks, one the architect invokes, or both (see
  the taxonomy in `marathon-sitrep.md`). Sessions read the catalog but never advance it, so it
  may need no hooks at all. Whether the extension contract supports a pure file convention is
  the main question.
- Whether `on-start` adds a convention, such as citing catalog keys in prose the way roadmap
  tasks are cited by dotted path, or whether nothing fires and sessions read the catalog when
  they need it.
- Where the files live. The workspace keeps them at the coordinator's repository root. An
  extension either adopts that location or standardizes on `context/`, and moving the existing
  workspace follows from that choice.
- Whether `references.md` stays hand-written or is generated from the TOML plus the descriptions.
  One authored home for each fact argues for generating it.
