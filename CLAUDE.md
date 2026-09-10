# claude-plugins

The Standards Lab organization's Claude Code plugin marketplace: independently versioned plugins that
codify the organization's development processes. This is the harness level of the reference architecture —
the worked example for the agent harness the other levels are built with. Managed with the marathon
workflow; start from `context/README.md`.

## Plugin design lives in the skill

Each plugin's design and behavior is expressed in its own files under `plugins/<name>/` — that is the
source of truth, and `scripts/check.sh` verifies that every pointer inside them resolves. `context/`
holds the repository's vision, concepts, and session record, not a restatement of how the plugins
work.

## Repository specifics

- **Marketplace host.** Each plugin under `plugins/` is independently versioned and released; consumers
  install only what they need. The marketplace manifest is `.claude-plugin/marketplace.json`.
- **The active skill is the installed copy.** This repository authors `marathon`; the workflow
  the agent runs is the installed marketplace copy. Changes here take effect once reinstalled.
