"""Tests for scripts/check.sh, run against a seeded copy of the repository.

The copy leaves out this file, so the check it runs doesn't recurse into
these tests, and puts a stub claude first on PATH, so plugin validation is
the outer check's job rather than this one's.
"""

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
        (self.root / "scripts/tests" / Path(__file__).name).unlink()
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

    def test_unseeded_copy_passes(self):
        self.assertEqual(self.check(), (0, []))

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
        self.write(SEEDED, "# Seeded\n\nSee `agents/nobody.md`.\n")
        code, lines = self.check()
        self.assertEqual(code, 1)
        self.assertEqual(
            lines,
            [
                f"FAIL: {SEEDED}:3: citation agents/nobody.md does not resolve"
            ],
        )

    def test_consumer_paths_stay_skipped(self):
        self.write(
            SEEDED,
            "# Seeded\n\nSee `context/roadmap.toml`, `scripts/check.sh`, and `src/app/main.py`.\n",
        )
        self.assertEqual(self.check(), (0, []))


if __name__ == "__main__":
    unittest.main()
