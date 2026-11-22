# Configuration

`.claude/marathon.toml` declares a repository's marathon configuration. `init` writes it once, and
every session reads it.

```toml
[project]
kind     = "code"
check    = "mise run check"
currency = "mise run currency"
gate     = "mise run gate"

[remote]
platform = "github"
publish  = "gh pr create"
checks   = "gh pr checks --watch"
merge    = "gh pr merge --merge --delete-branch"
ci       = '''
sha=$(git rev-parse HEAD)
for _ in $(seq 30); do
  runs=$(gh run list --commit "$sha" --event push --json databaseId --jq '.[].databaseId')
  [ -n "$runs" ] && break || sleep 10
done
[ -n "$runs" ] || exit 1
for run in $runs; do gh run watch --exit-status "$run" || exit 1; done
'''

[workspace]
role  = "coordinator"
order = [
  "core-lib",
  ["service-a", "service-b"],
  "gateway",
]

[workspace.paths]
core-lib = "~/code/core-lib"

[workspace.experiments]
path   = "~/experiments/spike-<slug>"
remote = "https://github.com/<owner>/spike-<slug>.git"
```

## Keys

- **`[project] kind`**: `code` when the repository holds production source with a build-and-test
  loop. `context` when the repository is itself context, such as prose, configuration, or skills,
  which the agent writes directly and which has no tests. A context project can still version and
  release what it ships.
- **`[project] check`**: the repository's one deterministic check, which every slice and review
  runs (`references/build.md`). A context project names its consistency script, if it has one.
  Every repository's check calls `lint-context` and `lint-references`. `lint-context` holds
  `context/` and `roadmap/` to their file kinds, word budgets, and shapes, and checks the relative
  links in the root README, `context/`, and `docs/`. `lint-references` resolves each markdown
  reference in the comments of the tracked files that aren't markdown: from the repository root,
  or in the sibling directory its first segment names. It skips a reference whose directory
  exists in neither place, as in CI. Both are Go binaries. The plugin's `bin/` shims put them on a
  session's PATH, each running the plugin's local build when present, or the plugin's release
  through mise, so a session needs mise on PATH. A repository pins the release in its
  `mise.toml`, so `mise install` puts both on PATH and `mise outdated` reports a newer tag:

  ```toml
  [tools]
  "github:standards-lab/claude-plugins" = { version = "0.22.0", version_prefix = "marathon/v" }
  ```

  CI calls both after `mise install`.
- **`[project] currency`**: optional. The repository's read-only command that reports what trails
  its latest release: the direct dependencies, the toolchain, and the pins, such as CI actions and
  images. Indirect dependencies are left to the ecosystem's resolver. PLAN runs it at development
  time; CI never does, since a release upstream would fail a build nothing in the repository
  changed. It exits 0 with empty stdout when everything is current. It exits nonzero when something
  trails, with one stdout line per trailing item, suggested as `<where>: <item> <pin> -> <latest>`,
  where the latest is what the ecosystem's tooling reports. Nonzero with no stdout lines means the
  command itself failed, which PLAN reports as a fact and doesn't block on. Diagnostics go to
  stderr. marathon prescribes no tool.
- **`[project] gate`**: optional. The repository's release gate, the release-time part of its
  automated checks, such as a plugin's eval suite (`references/build.md`). SHIP runs it at
  development time on the merge commit, before each Release tag, with that tag as its one
  argument, as in `mise run gate v1.4.0` (`mechanics/releasing.md`). It exits 0 when the tag may
  be pushed, and nonzero to hold it. CI never runs it. marathon fixes only this contract and
  prescribes no tool.
- **`[remote]`**: the platform, and the commands SHIP runs to publish a branch and, optionally,
  to merge it.
- **`[remote] checks`**: optional, for a repository that runs CI on pull requests. A command that
  waits for the pull request's checks and exits nonzero if one fails. Every merge runs it first.
- **`[remote] ci`**: optional. A command that waits for the default branch's CI runs on the merge
  commit, with that commit checked out, and exits nonzero if one fails or none starts. CI often
  runs more after a merge than on the pull request, such as integration tests only on the default
  branch. Runs take a few seconds to appear, so the command polls for them first, as the
  example's does. SHIP runs it before tagging a release; without it, tagging follows the merge
  directly (`mechanics/releasing.md`).
- **`[workspace]`**: for the coordinator only. `order` lists layers, lowest first, and an array
  entry is a layer of peers. `exclusive` lists groups of repositories that at most one active goal's
  `repos` may touch at a time (`mechanics/goals.md`, "The repository lock"). `[workspace.paths]`
  gives the location of a key that isn't a sibling directory
  (`references/workspace-coordination.md`). `[workspace.experiments]` gives the local path and
  remote a new spike takes, with `<slug>` replaced (`commands/experiment.md`).
- **`extensions`**: the enabled extensions, by skill name. Under `[project]`, they apply to this
  repository. Under `[workspace]`, they apply to every member (`references/extensions.md`).

## Running a command

`marathon <key> [args]` runs the `check`, `currency`, `gate`, `publish`, `checks`, `merge`, or
`ci` command above, and is on PATH in a session beside `lint-context` and `lint-references`. It
reads `.claude/marathon.toml` at the root of the git repository it runs in, and runs the key's
value with bash from that root, exactly as declared. A single-line command gets the arguments
appended, each shell-quoted, as in `marathon gate v1.4.0`. A multi-line command, such as `ci`,
runs as one script and gets them as its positional parameters. `marathon` exits with the command's
status. An unset or unknown key fails with the fix, such as "set [remote] merge in
.claude/marathon.toml".
