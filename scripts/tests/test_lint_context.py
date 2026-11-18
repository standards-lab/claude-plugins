"""Tests for marathon's lint-context, run against fixture trees built per test."""

import subprocess
import tempfile
import unittest
from pathlib import Path

LINT = Path(__file__).resolve().parents[2] / "plugins/marathon/bin/lint-context"


def text(count):
    return " ".join(["word"] * count) + "\n"


class LintContextTest(unittest.TestCase):
    def setUp(self):
        tmp = tempfile.TemporaryDirectory()
        self.addCleanup(tmp.cleanup)
        self.root = Path(tmp.name)
        self.write("README.md", text(10))
        self.write("context/README.md", text(10))
        self.write("context/note.md", text(10))
        self.write("context/roadmap.toml", "")
        self.write("context/roadmap/goal.toml", "")
        self.write("context/goals/goal.md", "# Goal\n\n## Decisions\n\n" + text(10))

    def write(self, rel, content):
        path = self.root / rel
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(content)

    def lint(self):
        result = subprocess.run(
            [str(LINT)], cwd=self.root, capture_output=True, text=True
        )
        return result.returncode, result.stdout.splitlines()

    def assert_clean(self):
        self.assertEqual(self.lint(), (0, []))

    def assert_fails(self, *lines):
        self.assertEqual(self.lint(), (1, list(lines)))

    def test_clean_tree_exits_silently(self):
        self.assert_clean()

    def test_missing_context_is_clean(self):
        empty = tempfile.TemporaryDirectory()
        self.addCleanup(empty.cleanup)
        self.root = Path(empty.name)
        self.assert_clean()

    def test_note_budget(self):
        self.write("context/note.md", text(1500))
        self.assert_clean()
        self.write("context/note.md", text(1501))
        self.assert_fails(
            "context/note.md: note is 1,501 words, budget 1,500; "
            "trim; move detail to docs/"
        )

    def test_context_readme_budget(self):
        self.write("context/README.md", text(1000))
        self.assert_clean()
        self.write("context/README.md", text(1001))
        self.assert_fails(
            "context/README.md: README is 1,001 words, budget 1,000; "
            "trim; move detail to docs/"
        )

    def test_root_readme_budget(self):
        self.write("README.md", text(1500))
        self.assert_clean()
        self.write("README.md", text(1501))
        self.assert_fails(
            "README.md: README is 1,501 words, budget 1,500; "
            "trim; move detail to docs/"
        )

    def test_docs_pages_are_unbudgeted(self):
        self.write("docs/guide.md", text(5000))
        self.assert_clean()

    def record(self, *sections):
        body = "".join(f"{heading}\n\n{text(count)}\n" for heading, count in sections)
        self.write("context/goals/goal.md", "# Goal\n\n" + body)

    def test_record_section_budgets(self):
        cases = (
            ("## Task brief · some-task", "Task brief · some-task", 1200, "shorten"),
            ("## Decisions", "Decisions", 2000, "prune what no longer binds later tasks"),
            ("## Pending edits", "Pending edits", 1000, "apply or shorten"),
        )
        for heading, name, budget, fix in cases:
            with self.subTest(heading=heading):
                self.record((heading, budget))
                self.assert_clean()
                self.record((heading, budget + 1))
                self.assert_fails(
                    f"context/goals/goal.md: section '{name}' is "
                    f"{budget + 1:,} words, budget {budget:,}; {fix}"
                )

    def test_record_sections_counted_separately(self):
        self.record(
            ("## Task brief · a", 1200),
            ("## Decisions", 2000),
            ("## Pending edits", 1000),
            ("## Progress", 5000),
        )
        self.assert_clean()

    def test_each_task_brief_counted_on_its_own(self):
        self.record(("## Task brief · a", 1200), ("## Task brief · b", 1201))
        self.assert_fails(
            "context/goals/goal.md: section 'Task brief · b' is 1,201 words, "
            "budget 1,200; shorten"
        )

    def test_stray_files_fail(self):
        stray = (
            "context/notes.txt",
            "context/sub/note.md",
            "context/goals/goal.toml",
            "context/roadmap/goal.md",
        )
        for rel in stray:
            self.write(rel, "")
        expected = sorted(
            f"{rel}: not a context file; context/ holds only README.md, "
            "notes (*.md), goals/*.md, roadmap.toml, and roadmap/*.toml"
            for rel in stray
        )
        code, lines = self.lint()
        self.assertEqual((code, sorted(lines)), (1, expected))

    def test_dotfiles_are_skipped(self):
        self.write("context/.DS_Store", "")
        self.write("context/.cache/x.bin", "")
        self.assert_clean()


if __name__ == "__main__":
    unittest.main()
