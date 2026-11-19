# marathon retro

Turn what the architect had to say about past work into changes to the system, so no comment is
written twice. `retro` reviews a task, a goal, or a date range, and includes the drift pass: notes
the code now expresses or contradicts. It runs LOCATE and START (`mechanics/pipeline.md`).

Read first, before acting:

- [`../mechanics/pipeline.md`](../mechanics/pipeline.md)
- [`../references/briefs.md`](../references/briefs.md)
- [`../references/context-engineering.md`](../references/context-engineering.md)

## Gather

For the range, read:

- the merged pull requests' bodies, which are the session briefs, including their Redirects and
  escalations and Decided without you
- the commits the standards-reviewer made, which show what the implementer kept missing
- the notes, `context/README.md`, and `docs/` pages the range touched, for drift
- when present, the local Claude Code transcripts of the range's sessions, under
  `~/.claude/projects/`, as optional evidence. They show the tool calls and their cost, which
  tool-economy findings cite.

A repository whose transcripts are missing is named in the retro's header, and the retro goes on
without them.

## Route

Route each finding to the cheapest layer that would have caught it, in this order:

1. **design**: a change to the work that makes the mistake impossible, which beats a check that
   only catches it
2. **check/eval**: a rule a check can enforce, or for a plugin, an eval case. A proposed check
   carries its failure message (`references/build.md`, "Checks come first").
3. **standard**: a judgement call for the repository's `STANDARDS.md`, or, when the same
   principle is expressed in more than one repository's merged code, a page in the architecture
   layer, which becomes a task in the repository that holds the layer
4. **skill/profile**: a change to a skill or a subagent profile
5. **context**: a note to cull, sharpen, or write, or a pointer to add

Look also for two kinds of finding, routed the same way: **tool economy**, a costly tool call that
a cheaper call or a tool could replace, and **instruction trim**, an instruction that changes no
behavior, whose proposed change is to cut it.

When the architecture layer is enabled, the drift pass also checks that every pointer in each
touched repository's `STANDARDS.md` resolves to a page in the layer; a broken pointer routes to
**standard**.

Drift findings route the same way: a note the built work expresses is culled, a note it
contradicts is fixed, and a `docs/` page it made wrong becomes a task.

Show the findings in the retro format (`references/briefs.md`, "Retro"), each one line with its
proposed change, and let the architect tick the ones to apply.

## Apply

Each repository the ticked findings change gets a `retro-<topic>` branch, published and merged with
the ticked findings as its body (`mechanics/pipeline.md`, "Branches and pull requests"):

- A ticked finding in the coordinator is applied on the coordinator's branch.
- A ticked finding in a repository no active goal locks is applied on that repository's branch.
- A ticked finding in a repository another active goal locks becomes a task in that goal's
  file, on the coordinator's branch.

A ticked design finding becomes a task in the goal that owns the work, on the coordinator's
branch; retro doesn't make the change. When no active goal owns the repository, the finding's line
proposes a goal, and the tick confirms it.

A ticked plugin eval finding becomes a new case under that plugin's `evals/<case>/`, in the
`claude plugin eval` case format, on its repository's `retro-<topic>` branch. The case's
`description` cites the failure it came from: the session ID or the PR. Retro writes the case and
does not fix the plugin; the fix is its own task. Until that task merges, the case may fail, and
the release gate (`mechanics/configuration.md`, `[project] gate`) holds the plugin's next release.
