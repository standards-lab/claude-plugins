# The reset file

`context/reset.md` is the session record and the pointer to the next step: written at the end of
one session, read at the start of the next. It is a purely ephemeral artifact — its sole purpose
is contextual bootstrapping between sessions, and durable detail belongs in the context layers,
never here. It is a single file, rewritten at each `reset` or `close`; git is the archive, and
keeps every older version.

## Where it lives

- A **standalone project** keeps its own `context/reset.md`.
- A **workspace** maintains exactly one reset file, at the coordinator. Member projects carry
  none: a session anywhere in the workspace reads the coordinator's reset at LOCATE and writes it
  at CONCLUDE — a commit in the coordinator's repository. The record names the member projects it
  concerns in its **Project** line.

## Schema

```markdown
# reset · wire-config-loader

- **Status:** closeout            # handoff | closeout
- **Session:** start              # init | plan | start | experiment | review | docs
- **Project:** core-lib           # workspace reset only: the member repo(s) the step touched
- **Branch:** wire-config-loader

## Disposition
- **Integrated:** removed the "three-phase load" note from design/config.md — the code now expresses it (config/loader).
- **Promoted:** concepts/config-validation.md → design/ (validation rules settled this session).
- **Culled:** dropped the env-override idea — the loader implementation went another way.
- **Retained:** design/config.md "secret sourcing" — still unbuilt.

## Next-focus
Add secret sourcing on top of the validated loader. Start here next session.
```

A standalone project's reset omits the Project line. In a workspace reset, Next-focus also names
the member project the next session continues in, so a session entered anywhere in the workspace
finds its way. A step that spanned several member repos lists them all on the Project line; its
branches share the step's slug, so the Branch line stays one value.

## Status semantics

- **closeout** — the recorded session finished and published. The next session starts a fresh
  step from the Next-focus, on a new branch.
- **handoff** — the recorded session stopped mid-work with its branch open. The Session line
  names the command that resumes it; the Branch line (and, in a workspace, the Project line)
  names where; Next-focus records the in-progress state and the exact next move, so a fresh
  context can resume without working it out again. A working session records its stage list and
  position there, for example `Stages: 3/7 · stage 3 implemented, unreviewed · list: …`. A step
  spanning member repos records each touched repo's branch state there too; whatever the shape of
  the interrupted work, Next-focus carries enough bootstrap state to resume it cold.

The Disposition speaks the tending ledger of `references/context-engineering.md`: Integrated
(a decayed note), Promoted, Culled, Retained — and **Cross-repo**, for an edit the session made
in another repository, such as a member session's edit to an artifact held at the coordinator.

Older reset files may carry a `Session type:` line naming `development`/`context`/`experiment`;
read it as the equivalent working session.
