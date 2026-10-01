# claude-plugins

Claude Code plugins for Standards Lab — agentic infrastructure for the
reference-architecture effort.

This repository is a plugin **host**: each plugin is independently versioned and released, and
projects install only the ones they need.

## Available Plugins

| Plugin | Description |
|--------|-------------|
| [marathon](./plugins/marathon/) | Long-haul development run as a software factory: goals and tasks in a roadmap, two architect touchpoints per task, and subagents that implement, review, and keep the context current |
| [marathon-architecture](./plugins/marathon-architecture/) | Marathon extension: the architecture layer that each repository's `STANDARDS.md` points to |

## Installation

```bash
claude plugin marketplace add standards-lab/claude-plugins

# Install the plugins you need
claude plugin install marathon@standards-lab
claude plugin install marathon-architecture@standards-lab
```

## Update

```bash
claude plugin marketplace update
claude plugin update marathon@standards-lab
claude plugin update marathon-architecture@standards-lab
```

## Remove

```bash
claude plugin remove marathon-architecture@standards-lab
claude plugin remove marathon@standards-lab
claude plugin marketplace remove standards-lab
```

## Configuration

Per project, grant the skills and tools a project actually uses in its `.claude/settings.json`.
The `marathon` core is standalone — it needs only its own skill permission:

```json
{
  "plansDirectory": "./.claude/plans",
  "permissions": {
    "allow": [
      "Skill(marathon:marathon)"
    ]
  }
}
```

`marathon init` writes this file when it scaffolds a project. Marathon extensions such as
`marathon-architecture` are opt-in; installing and enabling one is
[`references/extensions.md`](plugins/marathon/skills/marathon/references/extensions.md) in the
marathon skill. The core never requires an extension.

## How It Works

Skills load automatically based on conversational context. When Claude detects relevant triggers
(starting a task, planning goals, asking for status, running a retro, handing off, initializing a
project), it loads the appropriate skill to provide specialized guidance and commands.

User-invocable skills can also be triggered directly with slash commands (e.g., `/marathon:marathon`).

## Releases

Releases are tag-driven. Pushing a tag of the form `<plugin>/v<version>` (e.g. `marathon/v0.1.0`)
triggers [`.github/workflows/release.yml`](./.github/workflows/release.yml), which cuts a GitHub
release from that plugin's `CHANGELOG.md`.

## Repository Structure

```
claude-plugins/
├── .claude/                     # marathon project configuration for this repository
├── .claude-plugin/
│   └── marketplace.json         # Host manifest
├── .github/
│   └── workflows/
│       └── release.yml          # Tag-triggered release automation
├── CLAUDE.md                    # Agent orientation: the repo is developed under marathon
├── context/                     # marathon working context for the repo itself
├── plugins/
│   ├── marathon/                # Long-haul development as a software factory
│   └── marathon-architecture/   # Marathon extension: the architecture layer
├── LICENSE
└── README.md
```

The repository is developed with the workflow it ships: `CLAUDE.md`, `context/`, and `.claude/` are
marathon's own working infrastructure, so the plugin is exercised against its own source.
