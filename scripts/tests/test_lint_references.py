"""Tests for marathon's lint-references, run on a git repository built per test."""

import subprocess
import tempfile
import unittest
from pathlib import Path

LINT = Path(__file__).resolve().parents[2] / "plugins/marathon/bin/lint-references"
FIX = "fix the path or remove the reference"


class LintReferencesTest(unittest.TestCase):
    def setUp(self):
        tmp = tempfile.TemporaryDirectory()
        self.addCleanup(tmp.cleanup)
        self.parent = Path(tmp.name)
        self.root = self.parent / "repo"
        self.root.mkdir()
        subprocess.run(["git", "init", "-q"], cwd=self.root, check=True)
        self.write("docs/guide.md", "# Guide\n")

    def write(self, rel, content, root=None, track=True):
        root = root or self.root
        path = root / rel
        path.parent.mkdir(parents=True, exist_ok=True)
        if isinstance(content, bytes):
            path.write_bytes(content)
        else:
            path.write_text(content)
        if track and root == self.root:
            subprocess.run(["git", "add", rel], cwd=self.root, check=True)

    def sibling(self, rel, content="# Page\n"):
        self.write(rel, content, root=self.parent)

    def lint(self):
        result = subprocess.run(
            [str(LINT)], cwd=self.root, capture_output=True, text=True
        )
        return result.returncode, result.stdout.splitlines(), result.stderr

    def assert_result(self, code, *lines):
        self.assertEqual(self.lint(), (code, list(lines), ""))

    def test_each_marker_makes_a_reference_line(self):
        for marker in ("#", "//", "--", "/*", "*"):
            with self.subTest(marker=marker):
                self.write("src/code", f"  {marker} See docs/missing.md\n")
                self.assert_result(
                    1, f"src/code:1: reference docs/missing.md does not resolve; {FIX}"
                )

    def test_other_lines_are_not_reference_lines(self):
        self.write("src/code", 'path = "docs/missing.md"\nx = 1  # docs/missing.md\n')
        self.assert_result(0)

    def test_markdown_files_are_ignored(self):
        self.write("docs/page.md", "# See docs/missing.md\n")
        self.assert_result(0)

    def test_untracked_files_are_ignored(self):
        self.write("src/code", "# See docs/missing.md\n", track=False)
        self.assert_result(0)

    def test_resolves_at_the_root(self):
        self.write("src/code", "# See docs/guide.md, and ./docs/guide.md.\n")
        self.assert_result(0)

    def test_resolves_in_a_sibling(self):
        self.sibling("architecture/principles/comments.md")
        self.write(
            "src/code",
            "# See architecture/principles/comments.md\n"
            "# See ../architecture/principles/comments.md\n",
        )
        self.assert_result(0)

    def test_top_level_directory_shadows_the_sibling(self):
        self.sibling("architecture/principles/comments.md")
        self.write("architecture/principles/other.md", "# Other\n")
        self.write("src/code", "# See architecture/principles/comments.md\n")
        self.assert_result(
            1,
            "src/code:1: reference architecture/principles/comments.md "
            f"does not resolve; {FIX}",
        )

    def test_missing_in_the_repository_fails(self):
        self.write("src/code", "\n# See docs/missing.md\n")
        self.assert_result(
            1, f"src/code:2: reference docs/missing.md does not resolve; {FIX}"
        )

    def test_missing_in_an_existing_sibling_fails(self):
        self.sibling("architecture/principles/comments.md")
        self.write("src/code", "# See architecture/principles/missing.md\n")
        self.assert_result(
            1,
            "src/code:1: reference architecture/principles/missing.md "
            f"does not resolve; {FIX}",
        )

    def test_one_segment_missing_at_the_root_fails(self):
        self.write("src/code", "# See NOTES.md\n")
        self.assert_result(1, f"src/code:1: reference NOTES.md does not resolve; {FIX}")

    def test_no_sibling_skips_and_passes(self):
        self.write("src/code", "# See architecture/principles/comments.md\n")
        self.assert_result(
            0,
            "src/code:1: skipped architecture/principles/comments.md: "
            "no architecture/ in this repository or beside it",
        )

    def test_url_paths_are_not_references(self):
        self.write("src/code", "# See https://example.com/docs/missing.md\n")
        self.assert_result(0)

    def test_patterns_and_templates_are_not_references(self):
        self.write("src/code", "# Every *.md file, and each <run>.last-message.md\n")
        self.assert_result(0)

    def test_non_utf8_files_are_ignored(self):
        self.write("src/blob", b"# See docs/missing.md \xff\xfe\n")
        self.assert_result(0)

    def test_a_skip_with_a_failure_fails(self):
        self.write(
            "src/code",
            "# See architecture/principles/comments.md\n# See docs/missing.md\n",
        )
        self.assert_result(
            1,
            "src/code:1: skipped architecture/principles/comments.md: "
            "no architecture/ in this repository or beside it",
            f"src/code:2: reference docs/missing.md does not resolve; {FIX}",
        )


if __name__ == "__main__":
    unittest.main()
