# Workspaces

Related repositories, such as libraries, a template, and the service that uses them, are often
developed together, and a step sometimes changes several of them. A workspace runs that step as
one session, while each repository stays a complete marathon project of its own.

## The coordinator and the dependency order

A workspace is the directory that holds the projects as siblings. It has no `context/` of its
own. Instead, the project that already describes the whole set of repositories, usually an
organization-context project, declares itself coordinator with a `[workspace]` table
(`mechanics/configuration.md`).

The coordinator's `order` lists **layers**, lowest dependency first. An array entry is a layer of
peers with no dependencies between them. Each key is the name of a sibling directory, or a
location given in `[workspace.paths]`. Ask the architect about a key that matches neither. When
no project declares itself coordinator, list the sibling projects and ask.

## Steps that span repositories

A step that changes several member repositories runs as one session, lowest layer first, so each
higher layer builds against the real change below it. The step has one stage list, grouped by
repository in `order`. It creates a branch in each touched repository, all with the same name,
and `close` publishes each branch as that repository's own change proposal. The coordinator's
reset file records the step's state between sessions, and a resumed step reads `order` again.

## Concurrent lanes

The lanes of a wave (`mechanics/reset-file.md`) run at the same time, and every lane commits its
record to the coordinator. If lanes share one checkout, one lane's new branch can start from
another lane's branch, and one lane's commit can land on the other's branch. So each lane works in
a git worktree of its own:

- **Location.** The worktree for a lane's branch is `<repo>/.claude/worktrees/<branch>`, in every
  repository the lane touches, including the coordinator. Each of those repositories gitignores
  `.claude/worktrees/`.
- **Harness first.** The session enters the worktree through the harness's own worktree support,
  such as Claude Code's EnterWorktree with a `path`, or a session launched from the worktree's
  directory. Claude Code names a branch it creates `worktree-<name>`, not the step's branch, so
  the session creates the worktree itself with
  `git worktree add .claude/worktrees/<branch> -b <branch> <default-branch>` and then enters it.
- **The main checkout stays on the default branch** while the wave runs. No lane checks a branch
  out in it.

### Finalizing a worktree

A new worktree holds only tracked files, and its services can collide with the main checkout's.
So the lane finalizes each worktree right after creating it, before any stage runs:

1. **Copy the gitignored files the repository needs** from the main checkout, such as local
   secrets, `mise.local.toml`, or a `go.work`. The repository lists them in `.worktreeinclude` at
   its root, in `.gitignore` syntax, which is Claude Code's own convention. Only files that
   match and are gitignored are copied. Without the file, the lane judges what the repository
   needs from its README and build configuration. Copy the files rather than link them, so the
   lane can change its own without touching the main checkout's.
2. **Run the repository's setup**, the `[worktree] setup` command in its `.claude/marathon.toml`
   (`mechanics/configuration.md`), from the worktree's root. The setup handles what copying
   can't, such as remapping ports or seeding a database. Without a setup command, the lane
   isolates services itself. Docker Compose names its project after the directory, so a
   worktree's stack already gets its own containers and volumes, but its published host ports
   collide with the main checkout's stack. The lane sets `COMPOSE_PROJECT_NAME` and free host
   ports in a gitignored local file, such as `.env`, and never changes tracked files.
3. **Verify** with the repository's quick check, such as its build, so a missing file surfaces
   now and not partway through a stage.

A worktree kept at `reset` is already finalized, and a resumed lane only enters it. At `close`,
the lane stops the worktree's services and removes their volumes, then removes the worktree.

A session that runs alone keeps working in the main checkout. A worktree gains it nothing, and the
setup above still costs it.

## Experiments

An experiment is a standalone project outside the workspace, with its own reset file, so it runs
in parallel with the workspace's sessions. The coordinator lists it in a catalog and takes in its
results (`commands/experiment.md`).

## Projects know only what they depend on

A project knows what it builds on, and never what builds on it. Only the coordinator sees the
whole set of repositories. A lower project stays unaware of its consumers, even during a step that
spans repositories.

The conventions the coordinator keeps for its members, such as naming and writing rules, apply
through sessions, not citations: a member's `review` checks the member against them, and a step
that spans repositories applies them.
