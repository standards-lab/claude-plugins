# Waves

A closeout's Next-focus may name a **wave**: a set of **lanes** the architect judges safe to run
at the same time, each in sessions of its own. A lane is a single step, or a sequence of steps run
in order. Waves run in a standalone project and in a workspace.

## Lanes and the shared repository

The **shared repository** holds the reset file: the project itself when it stands alone, or the
coordinator in a workspace. Every lane keeps its record there, so it is the one repository the
lanes share.

Each lane owns exactly one of these:

- **The standalone project**, for its one main lane. A standalone project's other lanes are
  experiments.
- **A workspace member**, or the coordinator itself.
- **An experiment** (`commands/experiment.md`). Its setup and intake steps run at the shared
  repository, and its spikes run in its own repository.

No two lanes own the same thing. A lane does its step's work in what it owns, where no other
lane works, and changes the shared repository only under the rules below. An experiment lane owns
the experiment's repository, and its setup and intake at the shared repository follow the same
rules. A step that spans several repositories that lanes could own, such as two workspace
members, never runs as a lane.

## Lane records

While a wave runs, each lane keeps its own record at `context/reset/<lane>.md` in the shared
repository, in the reset file's schema (`mechanics/reset-file.md`), named after the slug of the
lane's first step. The record's Next-focus names the lane's next step, and after the lane's last
step it reads `Lane finished.` No lane writes `context/reset.md`.

At the shared repository, a lane changes only its record and the notes its steps own. It records a
change to any other file there, such as another lane's note or the experiment catalog, in its
Disposition instead of applying it. An extension's artifact, such as the roadmap or the
architecture layer, is shared wherever it lives, so a lane records a change to one in its
Disposition as well.

## Checkouts

- **The lane that owns the shared repository** works in its main checkout. That is the main lane
  of a standalone project, or a workspace lane that owns the coordinator.
- **Every other lane** changes the shared repository only in a git worktree of its own, at
  `<shared repository>/.claude/worktrees/<branch>`. The branch has the step's name, the same name
  as in the repository the lane owns. The shared repository gitignores `.claude/worktrees/`. The
  worktree holds only the lane's `context/` changes, so it needs no gitignored files or services.
- **When no lane owns the shared repository,** its main checkout stays on the default branch
  while the wave runs.

A lane creates its worktree at SETTLE with
`git worktree add .claude/worktrees/<branch> -b <branch> <default-branch>`, so the branch starts
from the default branch and not from whatever the main checkout holds. The session then enters
the checkout where its step's work happens through the harness, for a worktree with Claude Code's
EnterWorktree and a `path`, and reaches any other checkout by path. The session creates the
worktree itself before entering it, because Claude Code names a branch it creates
`worktree-<name>`.

A gitignored local file at the shared repository, such as a map of local checkouts, is changed in
the main checkout. It isn't versioned, so the change moves no branch, and a worktree's copy
would be lost when the worktree is removed.

`reset` keeps the worktree, and Next-focus records its path. `close` removes it after publishing.

## Reading a lane's record

LOCATE reads a lane's record from the checkout that holds the lane's open branch, either its
worktree or, for the owning lane, the main checkout. A lane with no open branch reads its record
from the default branch.

## Folding

A wave is folded in one commit once every lane's record reads `Lane finished.` on the default
branch. If the other lanes are already merged, the session that finishes the last lane folds the
wave. Otherwise, the next session whose LOCATE finds every lane finished folds it first. Folding
applies the changes the lanes recorded in their Dispositions, rewrites `context/reset.md`, and
deletes the lane records. No other lane is running by then, so the folding session applies these
changes in whatever checkout it holds, including files outside `context/` and in other
repositories.
