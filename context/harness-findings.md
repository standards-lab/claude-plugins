# Harness findings: inputs for later harness work

This note holds what later harness work draws on: the source behind a backlog task, and findings
a retro can take up.

## Sources

- **Brandon Thomas, "Let's Open Source the World"** (AI Engineer Code 2026). CLI and MCP are
  first-class, and humans review what is hard to reverse. The second point is the backlog task
  `v1.harness.door-scaled-review`.

## Findings for later

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
