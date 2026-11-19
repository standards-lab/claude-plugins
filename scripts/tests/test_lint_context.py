"""Tests for marathon's lint-context, run against fixture trees built per test."""

import subprocess
import tempfile
import unittest
from pathlib import Path

LINT = Path(__file__).resolve().parents[2] / "plugins/marathon/bin/lint-context"


INDEX = """# Roadmap index: the three arrays. The format is marathon's references/manifest.md.

active = ["goal"]
planned = []
backlog = []
"""
TASK = 'order = ["a"]\n\n[tasks.a]\n'
CONTEXT_README = "# Repo\n\n## Capability map\n\n- **Note** (built): see [note](note.md).\n"
ENTRY_FORM = "open it as '- **Name** (built): text' or '- **Name** (planned): text'"
ENTRY_FIX = "move detail to the goal's context note"
LINK_FIX = "point the link at an existing file or remove it"


def phrase(count):
    return " ".join(["word"] * count)


def text(count):
    return phrase(count) + "\n"


class LintContextTest(unittest.TestCase):
    def setUp(self):
        tmp = tempfile.TemporaryDirectory()
        self.addCleanup(tmp.cleanup)
        self.root = Path(tmp.name)
        self.write("README.md", text(10))
        self.write("context/README.md", CONTEXT_README)
        self.write("context/note.md", text(10))
        self.write("context/roadmap.toml", INDEX)
        self.write("context/roadmap/goal.toml", 'repos = []\n')
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
        padding = 1000 - len(CONTEXT_README.split())
        self.write("context/README.md", text(padding) + CONTEXT_README)
        self.assert_clean()
        self.write("context/README.md", text(padding + 1) + CONTEXT_README)
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
            f"{rel}: not a context file, which is README.md, a note (*.md), "
            "goals/*.md, roadmap.toml, or roadmap/*.toml; "
            "move it to docs/ or rename it to one of those kinds"
            for rel in stray
        )
        code, lines = self.lint()
        self.assertEqual((code, sorted(lines)), (1, expected))

    def test_dotfiles_are_skipped(self):
        self.write("context/.DS_Store", "")
        self.write("context/.cache/x.bin", "")
        self.assert_clean()

    def write_goal_file_without_record(self, body):
        (self.root / "context/goals/goal.md").unlink(missing_ok=True)
        self.write("context/roadmap/goal.toml", "repos = []\n" + body)

    def test_goal_entry_budgets(self):
        cases = (
            ("goal summary", 100, lambda words: f'summary = "{words}"\n'),
            ("tasks.a summary", 100, lambda words: f'{TASK}summary = "{words}"\n'),
            ("tasks.a proof", 40, lambda words: f'{TASK}proof = "{words}"\n'),
        )
        for entry, budget, body in cases:
            with self.subTest(entry=entry):
                self.write_goal_file_without_record(body(phrase(budget)))
                self.assert_clean()
                self.write_goal_file_without_record(body(phrase(budget + 1)))
                self.assert_fails(
                    f"context/roadmap/goal.toml: {entry} is {budget + 1:,} words, "
                    f"budget {budget:,}; {ENTRY_FIX}"
                )

    def test_criteria_list_counted_per_statement(self):
        self.write_goal_file_without_record(f'criteria = ["{phrase(60)}", "{phrase(60)}"]\n')
        self.assert_clean()
        self.write_goal_file_without_record(f'criteria = ["{phrase(60)}", "{phrase(61)}"]\n')
        self.assert_fails(
            "context/roadmap/goal.toml: goal criteria[1] is 61 words, budget 60; "
            + ENTRY_FIX
        )

    def test_criteria_string_is_one_statement(self):
        self.write_goal_file_without_record(f'criteria = "{phrase(60)}"\n')
        self.assert_clean()
        self.write_goal_file_without_record(f'criteria = "{phrase(61)}"\n')
        self.assert_fails(
            "context/roadmap/goal.toml: goal criteria is 61 words, budget 60; "
            + ENTRY_FIX
        )

    def test_index_holds_only_the_arrays(self):
        self.write("context/roadmap.toml", INDEX + 'extra = "x"\n')
        self.assert_fails(
            "context/roadmap.toml: 'extra' is not one of active, planned, backlog; "
            "move it into a goal file under context/roadmap/ or remove it"
        )
        self.write("context/roadmap.toml", INDEX + "\n[goal]\nname = \"x\"\n")
        self.assert_fails(
            "context/roadmap.toml: 'goal' is not one of active, planned, backlog; "
            "move it into a goal file under context/roadmap/ or remove it"
        )
        self.write("context/roadmap.toml", 'active = "goal"\n')
        self.assert_fails(
            "context/roadmap.toml: 'active' is not an array; "
            "write it as an array of goal slugs, [] for none"
        )

    def test_index_comment_lines(self):
        self.write("context/roadmap.toml", "# one\n# two\n" + INDEX.split("\n", 1)[1])
        self.assert_clean()
        self.write("context/roadmap.toml", "# one\n\n# two\n" + INDEX)
        self.assert_fails(
            "context/roadmap.toml: 3 comment lines, at most 2; "
            "keep one comment that points to the format"
        )

    def test_goal_file_writes_repos(self):
        self.write("context/roadmap/goal.toml", 'name = "Goal"\n')
        self.assert_fails(
            "context/roadmap/goal.toml: repos is missing; every goal file writes it, "
            "[] for none"
        )

    def test_order_lists_each_task_once(self):
        tasks = "\n[tasks.a]\nname = \"A\"\n\n[tasks.b]\nname = \"B\"\n"
        rel = "context/roadmap/goal.toml"
        cases = (
            ('order = ["a", "b"]\n', []),
            ("", [f"{rel}: order is missing; list each task slug once, in run order"]),
            (
                'order = ["a", "b", "a"]\n',
                [f"{rel}: order lists 'a' more than once; list it once"],
            ),
            (
                'order = ["a", "b", "c"]\n',
                [
                    f"{rel}: order lists 'c', which has no [tasks.c] table; "
                    "add a [tasks.c] table or drop it from order"
                ],
            ),
            (
                'order = ["b"]\n',
                [f"{rel}: order lacks task 'a'; add it to order or drop its [tasks.a] table"],
            ),
            (
                'order = "a b"\n',
                [
                    f"{rel}: order is not an array of task slugs; "
                    "write it as an array of the task slugs, in run order"
                ],
            ),
        )
        for order, lines in cases:
            with self.subTest(order=order):
                self.write_goal_file_without_record(order + tasks)
                self.assertEqual(self.lint(), (1 if lines else 0, lines))

    def test_order_without_tasks_lists_nothing(self):
        self.write_goal_file_without_record('order = ["a"]\n')
        self.assert_fails(
            "context/roadmap/goal.toml: order lists 'a', which has no [tasks.a] table; "
            "add a [tasks.a] table or drop it from order"
        )

    def test_record_tasks_mirror_order(self):
        self.write_goal_file_without_record(
            'order = ["a", "b"]\n\n[tasks.a]\nname = "A"\n\n[tasks.b]\nname = "B"\n'
        )
        self.write(
            "context/goals/goal.md",
            "# Goal\n\n## Tasks\n\n1. [x] a\n2. [ ] b\n\n## Decisions\n\nNone.\n",
        )
        self.assert_clean()
        for listed, slugs in (
            ("1. [ ] b\n2. [x] a\n", "b, a"),
            ("1. [x] a\n", "a"),
            ("", ""),
        ):
            with self.subTest(listed=listed):
                self.write("context/goals/goal.md", f"# Goal\n\n## Tasks\n\n{listed}")
                self.assert_fails(
                    f"context/goals/goal.md: ## Tasks lists [{slugs}], but "
                    "context/roadmap/goal.toml order is [a, b]; make the list mirror order"
                )

    def test_record_without_goal_file_is_unchecked(self):
        self.write("context/goals/other.md", "# Other\n\n## Tasks\n\n1. [ ] x\n")
        self.assert_clean()

    def test_unparseable_toml_fails_with_its_error(self):
        for rel in ("context/roadmap.toml", "context/roadmap/goal.toml"):
            with self.subTest(rel=rel):
                self.setUp()
                self.write(rel, "active = [\n")
                code, lines = self.lint()
                self.assertEqual(code, 1)
                self.assertEqual(len(lines), 1)
                self.assertTrue(lines[0].startswith(f"{rel}: does not parse: "), lines)
                self.assertTrue(lines[0].endswith("; fix the TOML at that position"), lines)

    def test_missing_roadmap_is_clean(self):
        (self.root / "context/roadmap.toml").unlink()
        (self.root / "context/roadmap/goal.toml").unlink()
        self.assert_clean()

    def capability_map(self, entries, after=""):
        self.write(
            "context/README.md",
            "# Repo\n\nIntro.\n\n## Capability map\n\n" + entries + after,
        )

    def test_capability_map_entries(self):
        self.capability_map(
            "- **Note** (built): see [note](note.md).\n"
            "- **Two** (built): spans lines\n"
            "  and links on a continuation, [notes](./).\n"
            "\n"
            "  - a nested item\n"
            "- **Later** (planned): links nothing yet.\n"
            "- **Soon** (planned): see [note](note.md).\n",
            "\n## Other\n\nProse after the map is not an entry.\n",
        )
        self.assert_clean()

    def test_capability_map_heading_required(self):
        self.write("context/README.md", "# Repo\n\n## Capabilities\n\n- **A** (planned): a\n")
        self.assert_fails(
            "context/README.md: no '## Capability map' heading; add one listing each "
            "feature as built or planned"
        )

    def test_capability_map_without_readme_is_unchecked(self):
        (self.root / "context/README.md").unlink()
        self.assert_clean()

    def test_capability_map_entry_form(self):
        cases = (
            ("- **Name**: text\n", "entry 'Name'"),
            ("- **Name** (active): text\n", "entry 'Name'"),
            ("- **Name** (built) text [n](note.md)\n", "entry 'Name'"),
            ("- Name (built): text [n](note.md)\n", "entry without a bold name"),
        )
        for entry, label in cases:
            with self.subTest(entry=entry):
                self.capability_map(entry)
                self.assert_fails(f"context/README.md:7: capability map {label}: {ENTRY_FORM}")

    def test_capability_map_prose_is_not_an_entry(self):
        self.capability_map("Some prose.\n\n- **A** (planned): a\n")
        self.assert_fails(
            f"context/README.md:7: capability map line is not an entry; {ENTRY_FORM}"
        )

    def test_built_entry_needs_a_relative_link(self):
        for entry in (
            "- **A** (built): links nothing.\n",
            "- **A** (built): only [a site](https://example.com).\n",
            "- **A** (built): only `[code](note.md)`.\n",
        ):
            with self.subTest(entry=entry):
                self.capability_map(entry + "- **B** (built): [n](note.md)\n")
                self.assert_fails(
                    "context/README.md:7: capability map entry 'A' is built but links "
                    "nothing; link its home with a relative link"
                )

    def test_capability_map_links_must_resolve(self):
        for status in ("built", "planned"):
            with self.subTest(status=status):
                self.capability_map(f"- **A** ({status}): see [gone](gone.md).\n")
                self.assert_fails(
                    f"context/README.md:7: link gone.md does not resolve; {LINK_FIX}"
                )

    def test_links_resolve(self):
        self.write("docs/guide.md", "# Guide\n")
        self.write("docs/img/a b.png", "")
        self.write(
            "README.md",
            "[guide](docs/guide.md#top) [dir](docs/) [q](docs/guide.md?x=1)\n"
            "[angle](<docs/img/a b.png>) [encoded](docs/img/a%20b.png)\n"
            '[titled](docs/guide.md "Title") [rooted](/docs/guide.md)\n'
            "[web](https://example.com/x) [mail](mailto:a@example.com) [anchor](#top)\n"
            "[ref]: docs/guide.md\n"
            "`[span](gone.md)` and ``[double `tick`](gone.md)``\n"
            "```\n[fenced](gone.md)\n```\n"
            "~~~md\n[tilde](gone.md)\n```\n~~~\n",
        )
        self.write("docs/sub/page.md", "[up](../guide.md) [root](/README.md)\n")
        self.write("context/goals/goal.md", "# Goal\n\n[note](../note.md)\n")
        self.assert_clean()

    def test_unresolved_links_fail(self):
        self.write("README.md", "Intro.\n\n[a](missing.md) and [b](<gone dir/>)\n")
        self.write("context/note.md", "[c](gone.md#x)\n\n[d]: ../nowhere.md\n")
        self.write("docs/page.md", "[e](/docs/missing.md \"title\")\n")
        self.write("context/goals/goal.md", "[f](note.md)\n")
        self.assert_fails(
            f"README.md:3: link missing.md does not resolve; {LINK_FIX}",
            f"README.md:3: link <gone dir/> does not resolve; {LINK_FIX}",
            f"context/goals/goal.md:1: link note.md does not resolve; {LINK_FIX}",
            f"context/note.md:1: link gone.md#x does not resolve; {LINK_FIX}",
            f"context/note.md:3: link ../nowhere.md does not resolve; {LINK_FIX}",
            f"docs/page.md:1: link /docs/missing.md does not resolve; {LINK_FIX}",
        )

    def test_links_outside_scanned_files_are_unchecked(self):
        self.write("other/page.md", "[a](missing.md)\n")
        self.write("docs/page.txt", "[a](missing.md)\n")
        self.assert_clean()


if __name__ == "__main__":
    unittest.main()
