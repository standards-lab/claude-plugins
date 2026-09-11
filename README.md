# claude-plugins

Claude Code plugins for Standards Lab — agentic infrastructure for the
reference-architecture effort.

This repository is a plugin **host**: each plugin is independently versioned and released, and
projects install only the ones they need.

## Available Plugins

| Plugin | Description |
|--------|-------------|
| [marathon](./plugins/marathon/) | Sustainable long-haul development workflow with context engineering |
| [marathon-roadmap](./plugins/marathon-roadmap/) | Marathon extension: a roadmap manifest of goals, tasks, and backlog the sessions keep current |
| [marathon-architecture](./plugins/marathon-architecture/) | Marathon extension: an architecture layer for the principles and conventions that outgrow a single repository, kept current as settled design proves out across the project |

## Installation

```bash
claude plugin marketplace add standards-lab/claude-plugins

# Install the plugins you need
claude plugin install marathon@standards-lab
claude plugin install marathon-roadmap@standards-lab
claude plugin install marathon-architecture@standards-lab
```

## Update

```bash
claude plugin marketplace update
claude plugin update marathon@standards-lab
claude plugin update marathon-roadmap@standards-lab
claude plugin update marathon-architecture@standards-lab
```

## Remove

```bash
claude plugin remove marathon-architecture@standards-lab
claude plugin remove marathon-roadmap@standards-lab
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
`marathon-roadmap` are opt-in; installing and enabling one is
[`references/extensions.md`](plugins/marathon/skills/marathon/references/extensions.md) in the
marathon skill. The core never requires an extension.

## How It Works

Skills load automatically based on conversational context. When Claude detects relevant triggers
(starting a development session, handing off or closing one out, reviewing drift, initializing a
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
│   ├── marathon/                # Sustainable long-haul development workflow
│   ├── marathon-roadmap/        # Marathon extension: the roadmap manifest convention
│   └── marathon-architecture/   # Marathon extension: the architecture layer
├── LICENSE
└── README.md
```

The repository is developed with the workflow it ships: `CLAUDE.md`, `context/`, and `.claude/` are
marathon's own working infrastructure, so the plugin is exercised against its own source.
