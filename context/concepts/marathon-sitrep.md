# marathon-sitrep: audience-calibrated situation reports

Captured 2026-08-27 from a planning session in the reference-architecture workspace. marathon's
discipline already leaves a complete record of what a period accomplished — the question a
situation report answers — but reading that record today is manual archaeology. This concept
proposes the extension that narrates it, and the taxonomy that gives such an extension a place in
the specification. Everything here is candidate direction for a future session on this repository.

## The gap

The repository-as-source-of-truth principle produces an accomplishment ledger as a byproduct of
the workflow. The coordinator reset file's git history is a session-by-session record of
dispositions and focus; the roadmap manifest's history shows each task deleted as it finished and
each goal closed; every member repository's log and merged pull requests carry the work itself.
Nothing reads that ledger back out. An architect who owes a stakeholder an account of the last
month re-walks the history by hand and calibrates the telling by feel, and the result is bound to
one audience at one depth.

There is a second gap underneath the first. `references/extensions.md` defines an extension by
the hooks it declares: "it declares what happens at the hook points every session fires; marathon
applies it at those points and otherwise never names it." A capability that layers on top of
marathon — invoked by the architect, firing at no hook — has no shape in the specification at
all.

## Proposal

`marathon-sitrep`, a user-invocable marathon extension. An invocation names a date range, an
audience, and an output format; the product is a situation report: an overview of the project or
workspace as it stands, then a narrative of what the range accomplished, at the depth and in the
register the audience calls for.

### The extension taxonomy

The concept treats "extension" as the higher-level idea, with two implementation facets:

- An **integration** fires at marathon's hook points and requires registration in
  `marathon.toml`; the sessions apply it. marathon-roadmap is the existing member.
- An **enhancement** layers functionality on top of the marathon skill — and potentially its
  enabled integrations — by reading the specification's own structures. It fires at no hook and
  the sessions never apply it; the architect invokes it.

A hybrid carries both facets. marathon-sitrep is the first pure enhancement. Building it amends
`references/extensions.md` to name the taxonomy; that amendment is part of the build session, not
a prerequisite.

An enhancement navigates by the specification rather than by private knowledge: coordinator
resolution and the `[workspace] order` map give it the member repositories; `[remote]` names the
platform whose merged pull requests it reads; the `extensions` lists in `marathon.toml` tell it
which integrations are enabled. With marathon-roadmap enabled, the report gains the roadmap
dimension — tasks and goals closed within the range, and the remaining path the manifest asserts.

### Data sources

- Each repository in the order map: `git log` over the range, with merges resolved to pull
  requests through the `[remote]` platform.
- The coordinator: `git log -p context/reset.md`, the session-by-session disposition ledger.
- The manifest, when present: `git log -p context/roadmap.toml`, the tasks deleted and goals
  closed within the range.
- The capability maps (`context/README.md` at the coordinator and members) for the standing
  overview the narrative opens with.

### The owned artifact: `context/sitrep.toml`

Audiences are the architect's stakeholders, not the agent's guess, so they are repository data.
The extension owns `context/sitrep.toml`, resolved the way the roadmap manifest is: at the
coordinator when serving a workspace, at the project itself when standalone. Each
`[audiences.<slug>]` table carries a name, a description of who the audience is, and calibration
guidance — depth, register, what this audience cares about and what it should be spared.
Bootstrapped at first invocation and populated with the architect.

### Output formats

The invocation selects a format: rendered in the conversation, written to a markdown file, an
HTML page, the clipboard. Candidate direction: the format set is itself extensible, with
`[formats.<slug>]` tables in `sitrep.toml` letting a workspace declare its own — a house template,
a distribution channel — so formats stay as workspace-defined as audiences are.

### Boundaries

A sitrep run is read-only against every repository. It opens no branch, runs no session
pipeline, writes no reset file, and never mutates `context/`. The report is a projection of
repository state in the sense of the source-of-truth rule: it lands only where the invocation
directs it and feeds nothing back.

## Open questions

- How an enhancement records its enablement: a `marathon.toml` `extensions` entry like an
  integration's, or the presence of its owned artifact. The entry keeps one enablement mechanism;
  the artifact would make the sessions' resolution list carry names they never fire.
- The format-extensibility mechanism: what a `[formats.<slug>]` declaration contains, and where
  the line sits between a format and a template.
- How an enhancement declares the marathon version it targets, since no session checks it — the
  invocation itself has to surface a mismatch.
- Scoping a run: whole workspace by default, with a single member project as a narrowing
  argument, or the reverse.
