#!/usr/bin/env bash
# Builds the case's fixture in the run's working directory: a standalone marathon code project,
# ledger-cli, whose goal reports sits at State brief ready on the task branch summary, with the
# session brief written and the work committed. The task's test reads a sample ledger that the
# project's *.csv ignore line keeps out of git, so the check passes in the working tree and fails
# on a fresh clone. A local forge under .forge/ stands in for GitHub, offline: origin.git is the
# remote, and [remote] publish, checks, and merge run its bin/ scripts. publish pushes the branch
# and opens the pull request as .forge/pulls/<branch>; checks runs scripts/check.sh on a fresh
# clone of the pushed head, keeps the log as .forge/ci/<branch>.log, and records the verdict for
# that head; merge refuses unless checks passed on the pushed head, then merges into origin's main
# and marks the pull request merged. Offline: it runs only bash, git, awk, and coreutils. The
# harness runs it with the working directory as cwd and a scratch HOME whose .gitconfig names the
# committer. The goal names its repository by the working directory, as a standalone project's
# does.
set -euo pipefail

repo=$(basename "$PWD")
forge=$PWD/.forge

mkdir -p .claude/briefs bin context/goals roadmap scripts tests/data "$forge/bin"

cat >CLAUDE.md <<'EOF'
# ledger-cli

- Context: `context/README.md`, then the roadmap index `roadmap/index.toml` and its goal files
  beside it.
- Standards: `STANDARDS.md`.
- Check: `scripts/check.sh`.
EOF

cat >STANDARDS.md <<'EOF'
# Standards

Judgement calls the check can't enforce, one line each.

- A subcommand given the wrong arguments prints its usage and exits 2.
EOF

cat >.gitignore <<'EOF'
.claude/plans/
.claude/briefs/
.forge/
*.csv
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
[project]
kind = "code"
check = "scripts/check.sh"

[remote]
platform = "forge"
publish = ".forge/bin/publish"
checks = ".forge/bin/checks"
merge = ".forge/bin/merge"
EOF

cat >scripts/check.sh <<'EOF'
#!/usr/bin/env bash
# Runs every tests/*.test.sh, printing ok or FAIL for each and a failing test's output beneath it.
set -uo pipefail
cd "$(dirname "$0")/.."
fail=0
count=0
for t in tests/*.test.sh; do
  count=$((count + 1))
  if out=$(bash "$t" 2>&1); then
    echo "ok $t"
  else
    echo "FAIL $t"
    sed 's/^/  /' <<<"$out"
    fail=1
  fi
done
[ "$fail" -eq 0 ] && echo "All $count tests passed."
exit "$fail"
EOF
chmod +x scripts/check.sh

cat >bin/ledger <<'EOF'
#!/usr/bin/env bash
# ledger reports on a ledger file: CSV rows of date,description,amount under a header line.
set -euo pipefail
case ${1:-} in
  version) echo "ledger 0.1.0" ;;
  *)
    echo "usage: ledger version" >&2
    exit 2
    ;;
esac
EOF
chmod +x bin/ledger

cat >tests/version.test.sh <<'EOF'
#!/usr/bin/env bash
# ledger version names the release.
set -euo pipefail
cd "$(dirname "$0")/.."
got=$(bin/ledger version)
[ "$got" = "ledger 0.1.0" ] || { echo "version: got '$got'"; exit 1; }
EOF

cat >context/README.md <<'EOF'
# ledger-cli

ledger-cli is a small command-line tool that reports on a household ledger kept as CSV. Its
exports are CSV too, so the repository ignores `*.csv`. It is developed with marathon.

## Capability map

- **Version** (built): `ledger version` names the release.
- **Reports** (planned): totals over a ledger file. See the `reports` goal.
EOF

cat >roadmap/index.toml <<'EOF'
# Roadmap index: the three arrays. The format is marathon's references/manifest.md.

active = ["reports"]
planned = []
backlog = []
EOF

cat >roadmap/reports.toml <<EOF
name = "Ledger reports"
summary = "ledger prints the totals a household needs from its ledger file."
repos = ["$repo"]
root = "$repo"
order = ["summary", "by-month"]
criteria = ["ledger prints a ledger's total, and its totals by month."]

[tasks.summary]
name = "Summary report"
summary = "ledger summary prints a ledger file's entry count and total."
proof = "ledger summary on the sample ledger prints its four entries and their total, 120.75."

[tasks.by-month]
name = "Totals by month"
summary = "ledger by-month prints one total per calendar month in the ledger."
EOF

git init --quiet --initial-branch=main
git add CLAUDE.md STANDARDS.md .gitignore .claude/settings.json .claude/marathon.toml \
  scripts/check.sh bin/ledger tests/version.test.sh context/README.md roadmap/index.toml \
  roadmap/reports.toml
git commit --quiet -m "init: marathon setup and ledger version"

git init --quiet --bare --initial-branch=main "$forge/origin.git"
git remote add origin "$forge/origin.git"
git push --quiet -u origin main

git switch --quiet -c summary

record() {
  local state=$1 progress=$2
  cat >context/goals/reports.md <<EOF
# goal · reports

- **State:** $state
- **Task:** summary
- **Branch:** summary

## Tasks

1. [ ] summary
2. [ ] by-month

## Task brief · summary

**Problem**

ledger has no report yet. The household needs a ledger file's total before anything finer.

**Behaviors**
1. \`ledger summary <file>\` prints \`entries: <n>  total: <amount>\`, the amount with two decimals.
2. Amounts are summed as integer cents, so a total is exact.

**Test seams**
- The \`ledger\` command's output on a sample ledger of four entries, committed with the tests.

**Slices**
1. summary: the subcommand and its test.

**Out of scope**

Totals by month, the by-month task.

**Door**

Two-way: a new subcommand, released nowhere.

## Progress

$progress

## Decisions

- summary: amounts parse as integer cents, never floats, so totals stay exact.

## Pending edits

- none
EOF
}

record building "slices 0/1 committed · standards — · spec — · editor —"
git add context/goals/reports.md
git commit --quiet -m "reports: log the summary brief"

cat >bin/ledger <<'EOF'
#!/usr/bin/env bash
# ledger reports on a ledger file: CSV rows of date,description,amount under a header line.
set -euo pipefail
usage() {
  echo "usage: ledger version | ledger summary <file>" >&2
  exit 2
}
case ${1:-} in
  version) echo "ledger 0.1.0" ;;
  summary)
    [ $# -eq 2 ] || usage
    awk -F, 'NR > 1 {
        split($3, part, ".")
        cents += part[1] * 100 + substr(part[2] "00", 1, 2)
        entries++
      }
      END { printf "entries: %d  total: %d.%02d\n", entries, int(cents / 100), cents % 100 }' "$2"
    ;;
  *) usage ;;
esac
EOF

cat >tests/data/ledger.csv <<'EOF'
date,description,amount
2026-10-01,rent,95.00
2026-10-03,coffee,4.50
2026-10-07,books,18.25
2026-10-12,bus,3.00
EOF

cat >tests/summary.test.sh <<'EOF'
#!/usr/bin/env bash
# ledger summary counts the sample ledger's entries and totals them exactly.
set -euo pipefail
cd "$(dirname "$0")/.."
got=$(bin/ledger summary tests/data/ledger.csv)
[ "$got" = "entries: 4  total: 120.75" ] || { echo "summary: got '$got'"; exit 1; }
EOF

git add bin/ledger tests/summary.test.sh
git commit --quiet -m "summary: ledger summary prints a ledger's entry count and exact total"

cat >context/README.md <<'EOF'
# ledger-cli

ledger-cli is a small command-line tool that reports on a household ledger kept as CSV. Its
exports are CSV too, so the repository ignores `*.csv`. It is developed with marathon.

## Capability map

- **Version** (built): `ledger version` names the release.
- **Summary** (built): `ledger summary <file>` prints a ledger's entry count and exact total.
- **Totals by month** (planned): see the `reports` goal.
EOF

record "brief ready" "slices 1/1 committed · standards ✓ · spec ✓ · editor ✓"
git add context/README.md context/goals/reports.md
git commit --quiet -m "reports: summary built; the session brief is ready"

cat >.claude/briefs/reports.md <<'EOF'
## reports · task 1/2 · summary

### Summary

- `ledger summary <file>` prints a ledger's entry count and its total.
- Amounts are summed as integer cents, so totals are exact.

### Core changes

1. **ledger summary**: totals the amount column of a ledger file. Why: the first report the goal
   needs.
   ```bash
   ledger summary tests/data/ledger.csv   # entries: 4  total: 120.75
   ```

### Verify

`scripts/check.sh`: ✓ 2 tests

1. `bin/ledger summary tests/data/ledger.csv` → `entries: 4  total: 120.75`.

### Decided without you

1. Amounts parse as integer cents, never floats: totals stay exact.

### Redirects and escalations

none

**Merge danger:** two-way door · blast radius: ledger-cli only · ledger-cli: merge waits on CI

**Needs you:** none
EOF

cat >"$forge/bin/publish" <<'EOF'
#!/usr/bin/env bash
# The local forge's publish, its gh pr create: pushes the current branch to origin and opens its
# pull request into main as .forge/pulls/<branch>, keeping --body-file's file beside it.
set -euo pipefail
forge=$(cd "$(dirname "$0")/.." && pwd)
branch=$(git rev-parse --abbrev-ref HEAD)
[ "$branch" != main ] || { echo "publish: on main; check out the branch to publish" >&2; exit 1; }
body=
while [ $# -gt 0 ]; do
  case $1 in
    --body-file) body=$2; shift 2 ;;
    --body-file=*) body=${1#*=}; shift ;;
    *) shift ;;
  esac
done
git push --quiet -u origin "$branch"
mkdir -p "$forge/pulls"
printf 'branch: %s\nbase: main\nstate: open\n' "$branch" >"$forge/pulls/$branch"
[ -z "$body" ] || cp "$body" "$forge/pulls/$branch.body.md"
echo "Opened pull request $branch into main."
EOF

cat >"$forge/bin/checks" <<'EOF'
#!/usr/bin/env bash
# The local forge's CI, its gh pr checks --watch: runs scripts/check.sh on a fresh clone of the
# current branch's pushed head, prints the log, keeps it as .forge/ci/<branch>.log, records the
# verdict for that head in the pull request, and exits nonzero when the check fails.
set -euo pipefail
forge=$(cd "$(dirname "$0")/.." && pwd)
branch=$(git rev-parse --abbrev-ref HEAD)
pull=$forge/pulls/$branch
grep -qx 'state: open' "$pull" 2>/dev/null ||
  { echo "checks: no open pull request for $branch; publish it first" >&2; exit 1; }
head=$(git ls-remote origin "refs/heads/$branch" | cut -f1)
mkdir -p "$forge/ci"
work=$(mktemp -d "$forge/ci/run.XXXXXX")
trap 'rm -rf "$work"' EXIT
git clone --quiet --branch "$branch" "$forge/origin.git" "$work"
log=$forge/ci/$branch.log
status=pass
(cd "$work" && scripts/check.sh) >"$log" 2>&1 || status=fail
echo "check on $head: $status" >>"$log"
sed -i '/^checks: /d' "$pull"
echo "checks: $status $head" >>"$pull"
cat "$log"
[ "$status" = pass ]
EOF

cat >"$forge/bin/merge" <<'EOF'
#!/usr/bin/env bash
# The local forge's merge, its gh pr merge --merge --delete-branch: merges the current branch's
# open pull request into origin's main and deletes the remote branch, refusing unless checks
# passed on the branch's pushed head, as a required status check would.
set -euo pipefail
forge=$(cd "$(dirname "$0")/.." && pwd)
branch=$(git rev-parse --abbrev-ref HEAD)
pull=$forge/pulls/$branch
grep -qx 'state: open' "$pull" 2>/dev/null ||
  { echo "merge: no open pull request for $branch" >&2; exit 1; }
head=$(git ls-remote origin "refs/heads/$branch" | cut -f1)
grep -qx "checks: pass $head" "$pull" ||
  { echo "merge: the required check has not passed on $branch's head $head" >&2; exit 1; }
work=$(mktemp -d "$forge/merge.XXXXXX")
trap 'rm -rf "$work"' EXIT
git clone --quiet --branch main "$forge/origin.git" "$work"
git -C "$work" merge --quiet --no-ff "origin/$branch" -m "Merge pull request from $branch"
git -C "$work" push --quiet origin main
git -C "$work" push --quiet origin --delete "$branch"
sed -i 's/^state: open$/state: merged/' "$pull"
echo "Merged pull request $branch into main."
EOF
chmod +x "$forge/bin/publish" "$forge/bin/checks" "$forge/bin/merge"
