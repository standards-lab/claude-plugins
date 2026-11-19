"""Tests for scripts/check.sh, run against a seeded copy of the repository.

The outer check runs the real test suite and validates the plugins, so the
copy's scripts/tests holds one passing test in their place, which keeps the
check from recursing into these tests, and a stub claude goes first on PATH.
"""

import json
import os
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path

REPO = Path(__file__).resolve().parents[2]
SKILL = "plugins/marathon/skills/marathon"
# A new file in the skill, outside every read-first set, so seeding it moves no
# word budget.
SEEDED = f"{SKILL}/references/seeded.md"
READ_FIRST_FIX = 'give it exactly one "Read first, before acting:" list'
RERUN = "fix the lines above, then rerun"
EXTENSION = "plugins/marathon-architecture"


class CheckTest(unittest.TestCase):
    def setUp(self):
        tmp = tempfile.TemporaryDirectory()
        self.addCleanup(tmp.cleanup)
        self.root = Path(tmp.name) / "repo"
        shutil.copytree(
            REPO,
            self.root,
            ignore=shutil.ignore_patterns(".git", "__pycache__"),
        )
        tests = self.root / "scripts/tests"
        shutil.rmtree(tests)
        tests.mkdir()
        (tests / "test_passes.py").write_text(
            "import unittest\n\n\nclass T(unittest.TestCase):\n"
            "    def test_passes(self):\n        pass\n"
        )
        stub = Path(tmp.name) / "bin/claude"
        stub.parent.mkdir()
        stub.write_text("#!/bin/sh\nexit 0\n")
        stub.chmod(0o755)
        self.env = {**os.environ, "PATH": f"{stub.parent}{os.pathsep}{os.environ['PATH']}"}

    def write(self, rel, content):
        (self.root / rel).write_text(content)

    def append(self, rel, content):
        with (self.root / rel).open("a") as f:
            f.write(content)

    def replace(self, rel, old, new):
        path = self.root / rel
        text = path.read_text()
        self.assertEqual(text.count(old), 1)
        path.write_text(text.replace(old, new))

    def check(self):
        result = subprocess.run(
            ["scripts/check.sh"], cwd=self.root, env=self.env, capture_output=True, text=True
        )
        return result.returncode, result.stderr.splitlines()

    def test_playbook_needs_exactly_one_read_first_line(self):
        self.append(f"{SKILL}/commands/status.md", "\nRead first, before acting:\n")
        self.replace(f"{SKILL}/commands/reset.md", "\nRead first", "\nRead these first")
        code, lines = self.check()
        self.assertEqual(code, 1)
        self.assertEqual(
            lines,
            [
                f'FAIL: {SKILL}/commands/reset.md: has 0 "Read first" lines; {READ_FIRST_FIX}',
                f'FAIL: {SKILL}/commands/status.md: has 2 "Read first" lines; {READ_FIRST_FIX}',
            ],
        )

    def test_typo_in_first_directory_fails(self):
        self.write(SEEDED, "# Seeded\n\nSee `refrences/build.md`.\n")
        code, lines = self.check()
        self.assertEqual(code, 1)
        self.assertEqual(
            lines,
            [
                f"FAIL: {SEEDED}:3: citation refrences/build.md: "
                f"directory 'refrences' is not in {SKILL}; did you mean 'references'?"
            ],
        )

    def test_skill_cites_plugin_agents(self):
        self.write(SEEDED, "# Seeded\n\nSee `agents/planner.md` and `agents/nobody.md`.\n")
        code, lines = self.check()
        self.assertEqual(code, 1)
        self.assertEqual(
            lines,
            [
                f"FAIL: {SEEDED}:3: citation agents/nobody.md does not resolve; "
                "cite an existing file or remove the citation"
            ],
        )

    def test_consumer_paths_pass(self):
        self.write(
            SEEDED,
            "# Seeded\n\nSee `context/roadmap.toml`, `scripts/check.sh`, and `src/app/main.py`.\n",
        )
        self.assertEqual(self.check(), (0, []))

    def test_version_mismatch_names_the_files_to_align(self):
        manifest = self.root / EXTENSION / ".claude-plugin/plugin.json"
        data = json.loads(manifest.read_text())
        version, data["version"] = data["version"], "9.9.9"
        manifest.write_text(json.dumps(data))
        code, lines = self.check()
        self.assertEqual(code, 1)
        skill = f"{EXTENSION}/skills/marathon-architecture/SKILL.md"
        self.assertEqual(
            lines,
            [
                "FAIL: marathon-architecture: plugin.json version 9.9.9 != CHANGELOG top "
                f"heading {version}; give {EXTENSION}/CHANGELOG.md a top '## v9.9.9' heading "
                "or correct plugin.json",
                f"FAIL: marathon-architecture: plugin.json version 9.9.9 != {skill} "
                f"Version: {version}; set its Version: line to 9.9.9 or correct plugin.json",
            ],
        )

    def test_marketplace_source_must_be_a_plugin(self):
        self.replace(".claude-plugin/marketplace.json", f"./{EXTENSION}", "./plugins/gone")
        code, lines = self.check()
        self.assertEqual(code, 1)
        self.assertEqual(
            lines,
            [
                "FAIL: marketplace source ./plugins/gone does not resolve to a plugin; "
                "point it at a directory holding .claude-plugin/plugin.json"
            ],
        )

    def test_unresolved_link_fails(self):
        self.write(SEEDED, "# Seeded\n\nSee [gone](./gone.md).\n")
        code, lines = self.check()
        self.assertEqual(code, 1)
        self.assertEqual(
            lines,
            [
                f"FAIL: {SEEDED}: link ./gone.md does not resolve; "
                "point the link at an existing file or remove it"
            ],
        )

    def test_budget_messages_give_the_budget_to_record(self):
        implementer = "plugins/marathon/agents/implementer.md"
        planner = "plugins/marathon/agents/planner.md"
        self.replace("scripts/budgets.txt", f"{implementer} 300\n", f"{implementer} 100\n")
        self.replace("scripts/budgets.txt", f"{planner} 300\n", "")
        sizes = {
            rel: len((self.root / rel).read_text().split()) for rel in (implementer, planner)
        }
        rounded = {rel: -(-size // 100) * 100 for rel, size in sizes.items()}
        code, lines = self.check()
        self.assertEqual(code, 1)
        self.assertEqual(
            lines,
            [
                f"FAIL: {implementer}: {sizes[implementer]} words, over its budget of 100 "
                f"in scripts/budgets.txt; trim it or raise the budget to {rounded[implementer]}",
                f"FAIL: {planner}: {sizes[planner]} words, no budget in scripts/budgets.txt; "
                f"add the line '{planner} {rounded[planner]}' to it",
            ],
        )

    def test_long_description_fails(self):
        skill = f"{EXTENSION}/skills/marathon-architecture/SKILL.md"
        self.replace(skill, "description: >\n", "description: >\n  " + "x" * 1024 + "\n")
        code, lines = self.check()
        self.assertEqual(code, 1)
        self.assertEqual(len(lines), 1, lines)
        self.assertRegex(
            lines[0],
            rf"^FAIL: {skill}: description is \d+ characters, over the 1024 limit; "
            "shorten it to 1024 characters or fewer$",
        )

    def test_wrapped_tools_point_at_their_output(self):
        stub = Path(self.env["PATH"].split(os.pathsep)[0]) / "claude"
        stub.write_text("#!/bin/sh\necho invalid\nexit 1\n")
        self.write(".claude-plugin/marketplace.json", "{")
        self.write("context/seeded.md", "[gone](gone.md)\n")
        self.write(
            "scripts/tests/test_seeded.py",
            "import unittest\n\n\nclass T(unittest.TestCase):\n"
            "    def test_errs(self):\n        raise RuntimeError\n",
        )
        code, lines = self.check()
        self.assertEqual(code, 1)
        self.assertEqual(
            [line for line in lines if line.startswith("FAIL: ")],
            [
                f"FAIL: .claude-plugin/marketplace.json does not parse; {RERUN}",
                f"FAIL: claude plugin validate failed for plugins/marathon/; {RERUN}",
                f"FAIL: claude plugin validate failed for {EXTENSION}/; {RERUN}",
                f"FAIL: claude plugin validate failed for .; {RERUN}",
                f"FAIL: lint-context failed; {RERUN}",
                f"FAIL: scripts/tests failed; {RERUN}",
            ],
        )


if __name__ == "__main__":
    unittest.main()
