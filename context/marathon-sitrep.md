# marathon-sitrep: situation reports for a chosen audience

marathon-sitrep is a planned marathon extension, tracked as `v1.harness.sitrep`. The architect
invokes it with a date range, an audience, and an output format. It produces a situation report:
an overview of the project or workspace as it stands, followed by an account of what the range
accomplished, at the depth and in the register the audience needs.

## The problem it solves

marathon already records everything a situation report needs:

- The git history of the coordinator's reset file records each session's dispositions and focus.
- The git history of the roadmap manifest shows each task deleted as it finished and each goal
  closed.
- Each member repository's log and merged pull requests hold the work itself.

Nothing reads that record back out. An architect who owes a stakeholder an account of the last
month walks the history by hand and adjusts the account to the audience by feel.

There is a second problem. `references/extensions.md` defines an extension by the hooks it
declares. A capability that the architect invokes, and that fires at no hook, has no place in
that definition.

## Two kinds of extension

This proposal names two kinds of extension:

- An **integration** fires at marathon's hook points and is enabled in `marathon.toml`. Sessions
  apply it. marathon-roadmap is an integration.
- An **enhancement** adds capability on top of marathon and its enabled integrations by reading
  the structures marathon defines. It fires at no hook, and the architect invokes it.

A hybrid extension has both parts. marathon-sitrep is a pure enhancement. The session that
builds it adds these two kinds to `references/extensions.md`.

An enhancement finds everything it reads through marathon's own configuration:

- The coordinator's `[workspace] order` lists the member repositories.
- `[remote]` names the platform whose merged pull requests it reads.
- The `extensions` lists name the enabled integrations. When marathon-roadmap is enabled, the
  report includes the tasks and goals closed in the range and the path that remains.

## Data sources

- For each repository in `order`: `git log` over the range, with merges resolved to pull
  requests through the `[remote]` platform.
- At the coordinator: `git log -p context/reset.md`, the record of each session's dispositions.
- The roadmap manifest, when present: `git log -p context/roadmap.toml`, for the tasks deleted
  and goals closed in the range.
- The capability maps in each `context/README.md`, for the overview the report opens with.

## The audience file

The architect's stakeholders are the audiences, so the audiences are repository data, not the
agent's guess. The extension owns `context/sitrep.toml`, resolved the way marathon-roadmap
resolves its manifest: at the coordinator in a workspace, or in the project itself when it is
standalone. Each `[audiences.<slug>]` table gives the audience's name, who they are, and how to
write for them: the depth, the register, what they care about, and what to leave out. The
extension creates the file on its first run and fills it in with the architect.

## Publication

Reports are published to a GitHub Pages dev blog, as `v1.harness.sitrep` states. A separate
dev-blog skill owns the blog's scaffold and publishing, and the extension calls it. The extension
owns reading the workspace and drafting the report. Leadership briefs are the blog's second kind
of entry.

## Limits

A report run only reads. It opens no branch, runs no session pipeline, writes no reset file, and
never changes `context/`. The report lands only where the invocation sends it.

## Open questions

- How an enhancement records that it is enabled: an entry in the `extensions` list, like an
  integration, or the presence of the file it owns. An entry keeps one way to enable extensions,
  but sessions would then resolve names they never fire.
- How an enhancement declares the marathon version it targets. No session checks it, so the
  invocation has to report a mismatch itself.
- The default scope of a run: the whole workspace, narrowed to one member by an argument, or the
  reverse.
