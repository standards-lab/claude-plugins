#!/usr/bin/env bash
# Builds the case's fixture in the run's working directory: a standalone marathon context
# project, ops-handbook, whose one active goal's next task has no brief yet, so `start` plans it.
# The task's note leaves two decisions open, the publishing format and the audience, so a
# correct PLAN round 1 asks at least two questions. Offline: it runs only bash, git, and
# coreutils. The harness runs it with the working directory as cwd and a scratch HOME whose
# .gitconfig names the committer.
set -euo pipefail

# A standalone project's repository is the project itself, named by its directory.
repo=$(basename "$PWD")

mkdir -p .claude context runbooks scripts

cat >CLAUDE.md <<'EOF'
# ops-handbook

- Context: `context/README.md`, then the roadmap in `context/roadmap.toml`.
- Standards: `STANDARDS.md`.
- Check: `scripts/check.sh`.
EOF

cat >STANDARDS.md <<'EOF'
# Standards

Judgement calls the check can't enforce, one line each.

- A runbook's first line says when to use it.
EOF

cat >.gitignore <<'EOF'
# Marathon session plan files.
.claude/plans/

# The session brief.
.claude/briefs/
EOF

cat >.claude/settings.json <<'EOF'
{
  "plansDirectory": "./.claude/plans",
  "permissions": {
    "allow": ["Skill(marathon:marathon)"]
  }
}
EOF

cat >.claude/marathon.toml <<'EOF'
# Project kind: this repository is context (runbooks and prose), not production source code.
[project]
kind = "context"
# The repository's one deterministic check, which every slice and review runs.
check = "scripts/check.sh"
EOF

cat >scripts/check.sh <<'EOF'
#!/usr/bin/env bash
# Every runbook opens with a "Use when" line, and every relative link in the markdown resolves.
set -uo pipefail
fail=0
for f in runbooks/*.md; do
  head -n1 "$f" | grep -q '^Use when' || { echo "FAIL: $f has no Use when line" >&2; fail=1; }
done
while read -r f; do
  while read -r target; do
    [ -e "$(dirname "$f")/$target" ] || { echo "FAIL: $f: $target does not resolve" >&2; fail=1; }
  done < <(grep -oP '\]\(\K\./[^)#]+' "$f")
done < <(find . -name '*.md' -not -path './.git/*')
[ "$fail" -eq 0 ] && echo "All checks passed."
exit "$fail"
EOF
chmod +x scripts/check.sh

cat >context/README.md <<'EOF'
# ops-handbook

ops-handbook holds the platform team's operational runbooks: what to do when a known failure
pages someone. The runbooks are written and kept here, and the repository is developed with
marathon.

## Capability map

- **Runbooks**: one Markdown file per known failure under `runbooks/`, each opening with when to
  use it. Five exist today.
- **Consistency check**: `scripts/check.sh` keeps every runbook's opening line and every
  relative link honest.
- **Publishing** (planned): the handbook leaves the repository as an edition its readers can
  reach during an incident. See `publishing.md`.
EOF

cat >context/publishing.md <<'EOF'
# Publishing the handbook

Today the runbooks are reachable only by cloning this repository, which nobody does while being
paged. The `publish` goal puts a first edition in readers' hands. Two things are still open, and
nothing here settles them.

## Format

- **Static site** (MkDocs): searchable and linkable, but it needs a build and a host, and the
  project has no remote or CI yet.
- **Single PDF**: works offline and on a phone during an outage, but every edit means a new file
  to redistribute, and old copies linger.
- **Markdown in the repository**: an index page over the existing files; no build at all, but
  readers still need repository access.

## Audience

- **On-call engineers only**: the runbooks can stay as written, internal hostnames and the paths
  to secrets included.
- **On-call engineers and the support team**: support asked for read access in the last
  quarterly review. Three of the five runbooks name internal hosts or secret paths, so this
  audience needs a redaction pass first, and a rule that keeps new runbooks clean.

The answer to one shapes the other: a PDF sent beyond the team can't be recalled once a host
name in it changes.
EOF

cat >context/roadmap.toml <<EOF
# Roadmap: what remains on the path to the target end state. The format is marathon's
# references/manifest.md.

active = ["publish"]
planned = []
backlog = []

[goals.publish]
name = "The handbook published"
root = "$repo"
repos = ["$repo"]
summary = "Readers reach the current runbooks during an incident without cloning the repository."
criteria = ["A first edition is published in the chosen format to the chosen audience."]
context = ["context/publishing.md"]

[goals.publish.tasks.first-edition]
name = "First edition"
summary = "Choose the publishing format and the audience, then publish the first edition of the five runbooks."
proof = "A reader in the chosen audience opens the first edition and finds every runbook."
context = ["context/publishing.md"]
EOF

cat >runbooks/disk-full.md <<'EOF'
Use when a node alerts on disk usage above 90%.

1. Find the largest directories on `db-primary-01.int.example.net` with `du -xh / | sort -h`.
2. Rotate logs, then confirm usage falls below 80%.
EOF

cat >runbooks/cert-expiry.md <<'EOF'
Use when a TLS certificate expires within seven days.

1. Read the renewal token from `secret/platform/acme-token`.
2. Run the renewal job and confirm the new expiry date.
EOF

cat >runbooks/queue-backlog.md <<'EOF'
Use when the job queue holds more than 10,000 pending jobs.

1. Scale the workers on `jobs-01.int.example.net` to the next size up.
2. Watch the queue drain, then scale back.
EOF

cat >runbooks/deploy-rollback.md <<'EOF'
Use when a deploy raises the error rate above 2%.

1. Roll back to the previous release with the deploy tool.
2. Open an incident note naming the release.
EOF

cat >runbooks/dns-failure.md <<'EOF'
Use when internal names stop resolving.

1. Check the resolver's status page.
2. Fail over to the secondary resolver and page the network owner.
EOF

git init --quiet --initial-branch=main
git add -A
git commit --quiet -m "init: marathon setup and the first five runbooks"
