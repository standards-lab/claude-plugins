# claude-plugins

The standards-lab organization's Claude Code plugin marketplace: independently versioned plugins that
codify the organization's development processes. This is the harness level of the reference architecture —
the worked example for the agent harness the other levels are built with. Managed with the marathon
workflow; start from `context/README.md`.

## Plugin design lives in the skill

Each plugin's design and behavior is expressed in its own files under `plugins/<name>/` — that is the
source of truth. `context/` holds the repository's vision, candidates, and session record, not a
restatement of how the plugins work. See `context/design/skill-is-the-source-of-truth.md`.

## Role boundary

The developer owns the production source — the plugin and skill files — and answers for it. The agent
writes everything else: tests, prose documentation, the files in `context/`, the implementation guide, and
the reset file.

## Repository specifics

- **Marketplace host.** Each plugin under `plugins/` is independently versioned and released; consumers
  install only what they need. The marketplace manifest is `.claude-plugin/marketplace.json`.
- **The active skill is the installed copy.** This repository is the *source* of `marathon`; the workflow
  the agent runs is the installed marketplace version. Changes here take effect once reinstalled.
