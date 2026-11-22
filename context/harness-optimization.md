# harness.optimization: marathon checked, not described

marathon's rules are checks, not prose. Goal `harness.optimization` turns them into checks, and this
note holds its sources, what it settled, and the findings it leaves for later.

## Sources

- **Lauren Tan, "All Code Is Slop (Except Mine)"** (AI Engineer Code 2026). Each correction
  becomes a deterministic check rather than another line of rules, because prose rules, docs, and
  comments are where agents write justifications for not doing the right thing, and they go
  stale. Context comes through structure, review reads verification artifacts, a feature map
  gives one vocabulary, and CLIs inside skills replace procedure.
- **Brandon Thomas, "Let's Open Source the World"** (AI Engineer Code 2026). CLI and MCP are
  first-class, and humans review what is hard to reverse. The second point is the backlog task
  `v1.harness.door-scaled-review`.
- **Matt Pocock, "Fixing the PR Bottleneck"** is marathon's baseline, since 0.16.

## Settled

- **Budgets are checks** (built). marathon's `lint-context` holds `context/` to its file
  kinds, the READMEs, notes, and goal-record sections to word budgets, and the roadmap and
  capability map to their shapes. It fails on a comment in a goal file and resolves relative
  links. Each repository's `[project] check` calls it (marathon's `mechanics/configuration.md`).
- **Comments** (built). The architecture layer's `principles/comments.md` allows five kinds of
  comment in source: package documentation, a lint suppression's reason, a pin hold's reason,
  provenance, and a reference, one line pointing a file at the document that governs it instead
  of restating it. Each ecosystem's stock lint rule requires a suppression's reason, each
  repository's check resolves its references with marathon's `lint-references`, and the
  standards-reviewer judges every other comment.
  marathon's standards-reviewer applies the comment rule a repository's `STANDARDS.md` declares.
  claude-plugins' scripts, tests, and configuration files follow it, as its `STANDARDS.md`
  declares, and marathon's configuration examples carry no comments. Every repository's check
  runs `lint-references`.
- **Roadmap layout** (built). marathon reads the root `roadmap/`: `roadmap/index.toml`, of the
  three arrays, plus one `roadmap/<goal>.toml` per goal, holding its own `root` and `repos`, an
  `order` array, and its tasks (marathon's `references/manifest.md`).
- **Harness standards** (built). claude-plugins' `STANDARDS.md` lists the judgement calls of
  harness programming, one line each, with pointers to the architecture pages that govern the
  repository, and the standards-reviewer applies it. `scripts/check.sh` enforces the mechanical
  rules: a CHANGELOG release heading carries only its version, a CHANGELOG holds no link
  definitions, and a marathon extension declares `user-invocable: false`.
- **Configured commands run through `marathon`** (built). marathon's `marathon` command runs a
  `.claude/marathon.toml` command by its key, exactly as declared, and every pipeline step that
  runs one calls it (marathon's `mechanics/configuration.md`, "Running a command"), so a session
  runs `[remote] merge` with its `&&` instead of retyping it. A hook that refuses a chained
  `gh pr merge` or `git push`, and a required `check` status on main, wait on
  `v1.harness.tooling` (`tool-based-skills.md`). The rule to run each git and `gh` command on its
  own call stays in the architect's user-scope configuration.
- **Verification artifacts** (built). Following Lauren Tan, review reads verification
  artifacts. The session brief's Verify section gives the steps the architect can run to see the
  work behave, and what they should see, or "none" for a prose-only change. The spec-reviewer
  has already followed them (marathon's `references/briefs/session-brief.md`).
- **Formats wrapped by a hook** (built). marathon's MessageDisplay hook runs `wrap-formats`,
  which rewraps a marathon format at 100 columns on screen as the reply streams, so the model never
  reprints it.

## Findings for later

A retro can take these up.

- marathon's manifest reference says `repos = []` means "only `plan` works on it", but `intake`
  runs on experiment goals, which write `repos = []`.
- A spike's `docs/` pages that link to its goal record break when sync deletes the record.
- Reviewers read the whole diff and the whole goal record. They could read `git diff --stat`
  first and receive excerpts of the brief.
- marathon's `references/context-engineering.md` says the session lists the editor's note operations
  in the session brief's Decided without you, but `references/briefs/session-brief.md` lists there
  only the decisions the brief didn't cover and the dropped Decisions. One of them should give way.
- The standards-reviewer rewraps markdown to 100 columns by hand; a check could enforce the
  width instead.
- `marathon` treats any value that holds a newline as multi-line, so a one-line TOML string with
  a `\n` escape gets its arguments as positional parameters, which the command ignores unless it
  reads `$1`.
- `claude plugin eval` stores no judge reasoning and makes the judge answer in one word, so a
  flipped verdict can't be traced. Both are upstream requests.
- `scripts/check.sh`'s header lists every check it runs, so each new check must be added there by
  hand, and nothing checks that it is.
- `scripts/check.sh`'s `unreleased_path_citations_with_owner` reads every markdown file whole and
  stops only a CHANGELOG at its first release heading, so "unreleased" describes only part of
  what it returns.
- `lint-references` and `lint-context` resolve a relative link's path but not its `#anchor`
  fragment, so reviews keep finding broken anchors by hand. Checking anchors is backlog for
  `lint-references`. A pinned architecture checkout in CI, to resolve `architecture/...`
  references there, was rejected: architecture has no tags, so the pin would be a bare SHA that no
  currency command tracks. Those references skip in CI instead.
