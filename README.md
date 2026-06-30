# claude-plugins

Claude Code plugins for the standards lab — agentic infrastructure for the
reference-architecture effort.

This repository is a plugin **host**: each plugin is independently versioned and released, and
projects install only the ones they need.

## Available Plugins

| Plugin | Description |
|--------|-------------|
| [marathon](./plugins/marathon/) | Sustainable long-haul development workflow with context engineering |

## Installation

```bash
claude plugin marketplace add standards-lab/claude-plugins

# Install the plugins you need
claude plugin install marathon@claude-plugins
```

## Update

```bash
claude plugin marketplace update
claude plugin update marathon@claude-plugins
```

## Remove

```bash
claude plugin remove marathon@claude-plugins
claude plugin marketplace remove claude-plugins
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

`marathon init` writes this file when it scaffolds a project. Opt-in platform extensions (added
later) are enabled per project by adding their plugin to `enabledPlugins` and their skill to
`permissions.allow` — the core never requires them.

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
├── .claude-plugin/
│   └── marketplace.json         # Host manifest
├── .github/
│   └── workflows/
│       └── release.yml          # Tag-triggered release automation
├── plugins/
│   └── marathon/                # Sustainable long-haul development workflow
└── README.md
```
