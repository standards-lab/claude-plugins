package contextlint_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/standards-lab/claude-plugins/marathon/internal/contextlint"
	"github.com/standards-lab/claude-plugins/marathon/internal/repotest"
)

const (
	index = "# Roadmap index: the three arrays. The format is marathon's references/manifest.md.\n" +
		"\nactive = [\"goal\"]\nplanned = []\nbacklog = []\n"
	task          = "order = [\"a\"]\n\n[tasks.a]\n"
	contextReadme = "# Repo\n\n## Capability map\n\n- **Note** (built): see [note](note.md).\n"
	entryForm     = "open it as '- **Name** (built): text' or '- **Name** (planned): text'"
	entryFix      = "move detail to the goal's context note"
	linkFix       = "point the link at an existing file or remove it"
	proseFix      = "trim; move detail to docs/"
)

func phrase(count int) string {
	return strings.TrimSuffix(strings.Repeat("word ", count), " ")
}

func text(count int) string {
	return phrase(count) + "\n"
}

type repo struct {
	*repotest.Repo
	t *testing.T
}

func newRepo(t *testing.T) repo {
	t.Helper()
	r := repo{repotest.New(t), t}
	r.Write("README.md", text(10))
	r.Write("context/README.md", contextReadme)
	r.Write("context/note.md", text(10))
	r.Write("roadmap/index.toml", index)
	r.Write("roadmap/goal.toml", "repos = []\n")
	r.Write("context/goals/goal.md", "# Goal\n\n## Decisions\n\n"+text(10))
	return r
}

func (r repo) assertClean() {
	r.t.Helper()
	r.AssertResult(contextlint.Run, 0)
}

func (r repo) assertFails(lines ...string) {
	r.t.Helper()
	r.AssertResult(contextlint.Run, 1, lines...)
}

func (r repo) assertFailsInAnyOrder(lines ...string) {
	r.t.Helper()
	code, got, stderr := r.Run(contextlint.Run)
	slices.Sort(got)
	slices.Sort(lines)
	if code != 1 || !slices.Equal(got, lines) || stderr != "" {
		r.t.Errorf("got (%d, %q, %q), want (1, %q, \"\") in any order", code, got, stderr, lines)
	}
}

type section struct {
	heading string
	words   int
}

func (r repo) record(sections ...section) {
	r.t.Helper()
	var body strings.Builder
	for _, s := range sections {
		body.WriteString(s.heading + "\n\n" + text(s.words) + "\n")
	}
	r.Write("context/goals/goal.md", "# Goal\n\n"+body.String())
}

func (r repo) goalFileWithoutRecord(body string) {
	r.t.Helper()
	r.Remove("context/goals/goal.md")
	r.Write("roadmap/goal.toml", "repos = []\n"+body)
}

func (r repo) capabilityMap(entries, after string) {
	r.t.Helper()
	r.Write("context/README.md", "# Repo\n\nIntro.\n\n## Capability map\n\n"+entries+after)
}

func TestCleanTreeExitsSilently(t *testing.T) {
	newRepo(t).assertClean()
}

func TestMissingContextIsClean(t *testing.T) {
	repo{repotest.New(t), t}.assertClean()
}

func TestNoteBudget(t *testing.T) {
	r := newRepo(t)
	r.Write("context/note.md", text(1500))
	r.assertClean()
	r.Write("context/note.md", text(1501))
	r.assertFails("context/note.md: note is 1,501 words, budget 1,500; " + proseFix)
}

func TestContextReadmeBudget(t *testing.T) {
	r := newRepo(t)
	padding := 1000 - len(strings.Fields(contextReadme))
	r.Write("context/README.md", text(padding)+contextReadme)
	r.assertClean()
	r.Write("context/README.md", text(padding+1)+contextReadme)
	r.assertFails("context/README.md: README is 1,001 words, budget 1,000; " + proseFix)
}

func TestRootReadmeBudget(t *testing.T) {
	r := newRepo(t)
	r.Write("README.md", text(1500))
	r.assertClean()
	r.Write("README.md", text(1501))
	r.assertFails("README.md: README is 1,501 words, budget 1,500; " + proseFix)
}

func TestWordsAreSeparatedByPythonWhitespace(t *testing.T) {
	r := newRepo(t)
	r.Write("README.md", strings.Repeat("word\x1c", 1500)+"word\u00a0word\u200bword\n")
	r.assertFails("README.md: README is 1,502 words, budget 1,500; " + proseFix)
}

func TestDocsPagesAreUnbudgeted(t *testing.T) {
	r := newRepo(t)
	r.Write("docs/guide.md", text(5000))
	r.assertClean()
}

func TestRecordSectionBudgets(t *testing.T) {
	cases := []struct {
		heading, name string
		budget        int
		fix           string
	}{
		{"## Task brief · some-task", "Task brief · some-task", 1200, "shorten"},
		{"## Decisions", "Decisions", 2000, "prune what no longer binds later tasks"},
		{"## Pending edits", "Pending edits", 1000, "apply or shorten"},
	}
	for _, c := range cases {
		t.Run(c.heading, func(t *testing.T) {
			r := newRepo(t)
			r.record(section{c.heading, c.budget})
			r.assertClean()
			r.record(section{c.heading, c.budget + 1})
			r.assertFails(fmt.Sprintf("context/goals/goal.md: section '%s' is %d,%03d words, "+
				"budget %d,%03d; %s", c.name, (c.budget+1)/1000, (c.budget+1)%1000,
				c.budget/1000, c.budget%1000, c.fix))
		})
	}
}

func TestRecordSectionsCountedSeparately(t *testing.T) {
	r := newRepo(t)
	r.record(section{"## Task brief · a", 1200}, section{"## Decisions", 2000},
		section{"## Pending edits", 1000}, section{"## Progress", 5000})
	r.assertClean()
}

func TestEachTaskBriefCountedOnItsOwn(t *testing.T) {
	r := newRepo(t)
	r.record(section{"## Task brief · a", 1200}, section{"## Task brief · b", 1201})
	r.assertFails("context/goals/goal.md: section 'Task brief · b' is 1,201 words, " +
		"budget 1,200; shorten")
}

func TestStrayFilesFail(t *testing.T) {
	r := newRepo(t)
	stray := []string{
		"context/notes.txt",
		"context/sub/note.md",
		"context/goals/goal.toml",
		"context/roadmap/goal.md",
		"context/roadmap/sub/goal.toml",
	}
	var want []string
	for _, rel := range stray {
		r.Write(rel, "")
		want = append(want, rel+": not a context file, which is README.md, a note (*.md), "+
			"or goals/*.md; move it to docs/ or rename it to one of those kinds")
	}
	r.assertFailsInAnyOrder(want...)
}

func TestFilesAreReportedInPathPartOrder(t *testing.T) {
	r := newRepo(t)
	r.Write("context/a-b/x", "")
	r.Write("context/a/x", "")
	kind := ": not a context file, which is README.md, a note (*.md), " +
		"or goals/*.md; move it to docs/ or rename it to one of those kinds"
	r.assertFails("context/a/x"+kind, "context/a-b/x"+kind)
}

func TestDotfilesAreSkipped(t *testing.T) {
	r := newRepo(t)
	r.Write("context/.DS_Store", "")
	r.Write("context/.cache/x.bin", "")
	r.assertClean()
}

func TestGoalEntryBudgets(t *testing.T) {
	cases := []struct {
		entry  string
		budget int
		body   func(string) string
	}{
		{"goal summary", 100, func(w string) string { return `summary = "` + w + "\"\n" }},
		{"tasks.a summary", 100, func(w string) string { return task + `summary = "` + w + "\"\n" }},
		{"tasks.a proof", 40, func(w string) string { return task + `proof = "` + w + "\"\n" }},
	}
	for _, c := range cases {
		t.Run(c.entry, func(t *testing.T) {
			r := newRepo(t)
			r.goalFileWithoutRecord(c.body(phrase(c.budget)))
			r.assertClean()
			r.Write("roadmap/goal.toml", "repos = []\n"+c.body(phrase(c.budget+1)))
			r.assertFails(fmt.Sprintf("roadmap/goal.toml: %s is %d words, budget %d; %s",
				c.entry, c.budget+1, c.budget, entryFix))
		})
	}
}

func TestCriteriaListCountedPerStatement(t *testing.T) {
	r := newRepo(t)
	r.goalFileWithoutRecord(`criteria = ["` + phrase(60) + `", "` + phrase(60) + "\"]\n")
	r.assertClean()
	r.Write("roadmap/goal.toml", `repos = []
criteria = ["`+phrase(60)+`", "`+phrase(61)+"\"]\n")
	r.assertFails("roadmap/goal.toml: goal criteria[1] is 61 words, budget 60; " + entryFix)
}

func TestCriteriaStringIsOneStatement(t *testing.T) {
	r := newRepo(t)
	r.goalFileWithoutRecord(`criteria = "` + phrase(60) + "\"\n")
	r.assertClean()
	r.Write("roadmap/goal.toml", "repos = []\ncriteria = \""+phrase(61)+"\"\n")
	r.assertFails("roadmap/goal.toml: goal criteria is 61 words, budget 60; " + entryFix)
}

func TestIndexHoldsOnlyTheArrays(t *testing.T) {
	r := newRepo(t)
	r.Write("roadmap/index.toml", index+"extra = \"x\"\n")
	r.assertFails("roadmap/index.toml: 'extra' is not one of active, planned, backlog; " +
		"move it into a goal file under roadmap/ or remove it")
	r.Write("roadmap/index.toml", index+"\n[goal]\nname = \"x\"\n")
	r.assertFails("roadmap/index.toml: 'goal' is not one of active, planned, backlog; " +
		"move it into a goal file under roadmap/ or remove it")
	r.Write("roadmap/index.toml", "active = \"goal\"\n")
	r.assertFails("roadmap/index.toml: 'active' is not an array; " +
		"write it as an array of goal slugs, [] for none")
}

func TestIndexKeysAreReportedInDocumentOrder(t *testing.T) {
	r := newRepo(t)
	r.Write("roadmap/index.toml", "zeta = 1\nactive = []\nalpha.x = 1\n[[mid]]\n")
	fix := " is not one of active, planned, backlog; " +
		"move it into a goal file under roadmap/ or remove it"
	r.assertFails("roadmap/index.toml: 'zeta'"+fix, "roadmap/index.toml: 'alpha'"+fix,
		"roadmap/index.toml: 'mid'"+fix)
}

func TestIndexCommentLines(t *testing.T) {
	r := newRepo(t)
	r.Write("roadmap/index.toml", "# one\n# two\n"+strings.SplitN(index, "\n", 2)[1])
	r.assertClean()
	r.Write("roadmap/index.toml", "# one\n\n# two\n"+index)
	r.assertFails("roadmap/index.toml: 3 comment lines, at most 2; " +
		"keep one comment that points to the format")
}

func TestGoalFileCarriesNoComments(t *testing.T) {
	r := newRepo(t)
	strs := "repos = []\n" +
		"summary = \"\"\"\n# inside a basic string\\\n\"\"\"\n" +
		"proof = '''\n# inside a literal string\n'''\n" +
		"name = \"# not a comment\"\n"
	r.Write("roadmap/goal.toml", strs)
	r.assertClean()
	r.Write("roadmap/goal.toml", "# a note\n"+strs+"  # indented\nroot = \".\" # trailing\n")
	fix := "comment; move the note into the goal's context note"
	r.assertFails("roadmap/goal.toml:1: "+fix, "roadmap/goal.toml:10: "+fix,
		"roadmap/goal.toml:11: "+fix)
}

func TestCarriageReturnsReadAsNewlines(t *testing.T) {
	r := newRepo(t)
	r.Write("roadmap/goal.toml", "repos = []\r\nname = \"x\"\r# a note\r\n")
	r.assertFails("roadmap/goal.toml:3: comment; move the note into the goal's context note")
}

func TestGoalFileWritesRepos(t *testing.T) {
	r := newRepo(t)
	r.Write("roadmap/goal.toml", "name = \"Goal\"\n")
	r.assertFails("roadmap/goal.toml: repos is missing; every goal file writes it, [] for none")
}

func TestOrderListsEachTaskOnce(t *testing.T) {
	tasks := "\n[tasks.a]\nname = \"A\"\n\n[tasks.b]\nname = \"B\"\n"
	rel := "roadmap/goal.toml"
	cases := []struct {
		order string
		lines []string
	}{
		{"order = [\"a\", \"b\"]\n", nil},
		{"", []string{rel + ": order is missing; list each task slug once, in run order"}},
		{"order = [\"a\", \"b\", \"a\"]\n",
			[]string{rel + ": order lists 'a' more than once; list it once"}},
		{"order = [\"a\", \"b\", \"c\"]\n", []string{rel + ": order lists 'c', which has no " +
			"[tasks.c] table; add a [tasks.c] table or drop it from order"}},
		{"order = [\"b\"]\n",
			[]string{rel + ": order lacks task 'a'; add it to order or drop its [tasks.a] table"}},
		{"order = \"a b\"\n", []string{rel + ": order is not an array of task slugs; " +
			"write it as an array of the task slugs, in run order"}},
	}
	for _, c := range cases {
		t.Run(c.order, func(t *testing.T) {
			r := newRepo(t)
			r.goalFileWithoutRecord(c.order + tasks)
			if c.lines == nil {
				r.assertClean()
			} else {
				r.assertFails(c.lines...)
			}
		})
	}
}

func TestOrderLacksTasksInDocumentOrder(t *testing.T) {
	r := newRepo(t)
	r.goalFileWithoutRecord("order = []\n\n[tasks.z]\n[tasks.b]\n[tasks.m]\n")
	r.assertFails(
		"roadmap/goal.toml: order lacks task 'z'; add it to order or drop its [tasks.z] table",
		"roadmap/goal.toml: order lacks task 'b'; add it to order or drop its [tasks.b] table",
		"roadmap/goal.toml: order lacks task 'm'; add it to order or drop its [tasks.m] table")
}

func TestOrderWithoutTasksListsNothing(t *testing.T) {
	r := newRepo(t)
	r.goalFileWithoutRecord("order = [\"a\"]\n")
	r.assertFails("roadmap/goal.toml: order lists 'a', which has no [tasks.a] table; " +
		"add a [tasks.a] table or drop it from order")
}

func TestRecordTasksMirrorOrder(t *testing.T) {
	cases := []struct{ listed, slugs string }{
		{"1. [ ] b\n2. [x] a\n", "b, a"},
		{"1. [x] a\n", "a"},
		{"", ""},
	}
	for _, c := range cases {
		t.Run(c.listed, func(t *testing.T) {
			r := newRepo(t)
			r.goalFileWithoutRecord(
				"order = [\"a\", \"b\"]\n\n[tasks.a]\nname = \"A\"\n\n[tasks.b]\nname = \"B\"\n")
			r.Write("context/goals/goal.md",
				"# Goal\n\n## Tasks\n\n1. [x] a\n2. [ ] b\n\n## Decisions\n\nNone.\n")
			r.assertClean()
			r.Write("context/goals/goal.md", "# Goal\n\n## Tasks\n\n"+c.listed)
			r.assertFails("context/goals/goal.md: ## Tasks lists [" + c.slugs + "], but " +
				"roadmap/goal.toml order is [a, b]; make the list mirror order")
		})
	}
}

func TestRecordWithoutGoalFileIsUnchecked(t *testing.T) {
	r := newRepo(t)
	r.Write("context/goals/other.md", "# Other\n\n## Tasks\n\n1. [ ] x\n")
	r.assertClean()
}

func TestUnparseableTOMLFailsWithItsError(t *testing.T) {
	for _, c := range []struct{ rel, content string }{
		{"roadmap/index.toml", "active = [\n"},
		{"roadmap/goal.toml", "active = [\n"},
		{"roadmap/goal.toml", "\ufeffrepos = []\n"},
	} {
		t.Run(c.rel+" "+c.content, func(t *testing.T) {
			r := newRepo(t)
			r.Write(c.rel, c.content)
			code, lines, stderr := r.Run(contextlint.Run)
			if code != 1 || len(lines) != 1 || stderr != "" ||
				!strings.HasPrefix(lines[0], c.rel+": does not parse: toml: line 1") ||
				!strings.HasSuffix(lines[0], "; fix the TOML at that position") {
				t.Errorf("got (%d, %q, %q)", code, lines, stderr)
			}
		})
	}
}

func TestUnreadableTextFailsWithNoReport(t *testing.T) {
	for _, rel := range []string{"context/note.md", "docs/page.md", "roadmap/goal.toml"} {
		t.Run(rel, func(t *testing.T) {
			r := newRepo(t)
			r.Write("context/stray.txt", "")
			r.Write(rel, "word \xff\n")
			code, lines, stderr := r.Run(contextlint.Run)
			want := "lint-context: " + rel + ": not UTF-8\n"
			if code != 1 || lines != nil || stderr != want {
				t.Errorf("got (%d, %q, %q), want (1, [], %q)", code, lines, stderr, want)
			}
		})
	}
}

func TestMissingRoadmapIsClean(t *testing.T) {
	r := newRepo(t)
	r.Remove("roadmap/index.toml")
	r.Remove("roadmap/goal.toml")
	r.Remove("roadmap")
	r.assertClean()
}

func TestStrayRoadmapFilesFail(t *testing.T) {
	r := newRepo(t)
	stray := []string{"roadmap/notes.md", "roadmap/goal.json", "roadmap/sub/goal.toml"}
	var want []string
	for _, rel := range stray {
		r.Write(rel, "")
		want = append(want, rel+": not a roadmap file; name it index.toml or <goal>.toml, "+
			"or move it out of roadmap/")
	}
	r.Write("roadmap/.DS_Store", "")
	r.assertFailsInAnyOrder(want...)
}

func TestOldRoadmapLayoutFailsUnchecked(t *testing.T) {
	r := newRepo(t)
	r.Write("context/roadmap.toml", "extra = \"x\"\n")
	r.Write("context/roadmap/goal.toml", "# a note\nactive = [\n")
	r.assertFails(
		"context/roadmap/goal.toml: the roadmap moved; move it to roadmap/goal.toml",
		"context/roadmap.toml: the roadmap index moved; move it to roadmap/index.toml")
}

func TestCapabilityMapEntries(t *testing.T) {
	r := newRepo(t)
	r.capabilityMap(
		"- **Note** (built): see [note](note.md).\n"+
			"- **Two** (built): spans lines\n"+
			"  and links on a continuation, [notes](./).\n"+
			"\n"+
			"  - a nested item\n"+
			"- **Later** (planned): links nothing yet.\n"+
			"- **Soon** (planned): see [note](note.md).\n",
		"\n## Other\n\nProse after the map is not an entry.\n")
	r.assertClean()
}

func TestCapabilityMapHeadingRequired(t *testing.T) {
	r := newRepo(t)
	r.Write("context/README.md", "# Repo\n\n## Capabilities\n\n- **A** (planned): a\n")
	r.assertFails("context/README.md: no '## Capability map' heading; add one listing each " +
		"feature as built or planned")
}

func TestCapabilityMapWithoutReadmeIsUnchecked(t *testing.T) {
	r := newRepo(t)
	r.Remove("context/README.md")
	r.assertClean()
}

func TestCapabilityMapEntryForm(t *testing.T) {
	cases := []struct{ entry, label string }{
		{"- **Name**: text\n", "entry 'Name'"},
		{"- **Name** (active): text\n", "entry 'Name'"},
		{"- **Name** (built) text [n](note.md)\n", "entry 'Name'"},
		{"- Name (built): text [n](note.md)\n", "entry without a bold name"},
		{"- **A** b** (built):\u00a0text\n", "entry 'A'"},
	}
	for _, c := range cases {
		t.Run(c.entry, func(t *testing.T) {
			r := newRepo(t)
			r.capabilityMap(c.entry, "")
			r.assertFails("context/README.md:7: capability map " + c.label + ": " + entryForm)
		})
	}
}

func TestCapabilityMapNameIsTheShortestBeforeTheStatus(t *testing.T) {
	r := newRepo(t)
	r.capabilityMap("- **A** b** (built): links nothing.\n", "")
	r.assertFails("context/README.md:7: capability map entry 'A** b' is built but links " +
		"nothing; link its home with a relative link")
}

func TestCapabilityMapProseIsNotAnEntry(t *testing.T) {
	r := newRepo(t)
	r.capabilityMap("Some prose.\n\n- **A** (planned): a\n", "")
	r.assertFails("context/README.md:7: capability map line is not an entry; " + entryForm)
}

func TestBuiltEntryNeedsARelativeLink(t *testing.T) {
	for _, entry := range []string{
		"- **A** (built): links nothing.\n",
		"- **A** (built): only [a site](https://example.com).\n",
		"- **A** (built): only `[code](note.md)`.\n",
	} {
		t.Run(entry, func(t *testing.T) {
			r := newRepo(t)
			r.capabilityMap(entry+"- **B** (built): [n](note.md)\n", "")
			r.assertFails("context/README.md:7: capability map entry 'A' is built but links " +
				"nothing; link its home with a relative link")
		})
	}
}

func TestCapabilityMapLinksMustResolve(t *testing.T) {
	for _, status := range []string{"built", "planned"} {
		t.Run(status, func(t *testing.T) {
			r := newRepo(t)
			r.capabilityMap("- **A** ("+status+"): see [gone](gone.md).\n", "")
			r.assertFails("context/README.md:7: link gone.md does not resolve; " + linkFix)
		})
	}
}

func TestLinksResolve(t *testing.T) {
	r := newRepo(t)
	r.Write("docs/guide.md", "# Guide\n")
	r.Write("docs/img/a b.png", "")
	r.Write("docs/img/é\ufffd.png", "")
	r.Write("docs/img/%zz%4.png", "")
	r.Write("README.md",
		"[guide](docs/guide.md#top) [dir](docs/) [q](docs/guide.md?x=1)\n"+
			"[angle](<docs/img/a b.png>) [encoded](docs/img/a%20b.png)\n"+
			"[titled](docs/guide.md \"Title\") [rooted](/docs/guide.md)\n"+
			"[web](https://example.com/x) [mail](mailto:a@example.com) [anchor](#top)\n"+
			"[ref]: docs/guide.md\n"+
			"`[span](gone.md)` and ``[double `tick`](gone.md)``\n"+
			"```\n[fenced](gone.md)\n```\n"+
			"~~~md\n[tilde](gone.md)\n```\n~~~\n"+
			"[file slash](docs/guide.md/) [dot](docs/./guide.md/.) [empty]()\n"+
			"[invalid](docs/img/%C3%A9%E2%82.png) [bad escape](docs/img/%zz%4.png)\n")
	r.Write("docs/sub/page.md", "[up](../guide.md) [root](/README.md)\n")
	r.Write("context/goals/goal.md", "# Goal\n\n[note](../note.md)\n")
	r.assertClean()
}

func TestUnresolvedLinksFail(t *testing.T) {
	r := newRepo(t)
	r.Write("README.md", "Intro.\n\n[a](missing.md) and [b](<gone dir/>)\n")
	r.Write("context/note.md", "[c](gone.md#x)\n\n[d]: ../nowhere.md\n")
	r.Write("docs/page.md", "[e](/docs/missing.md \"title\")\n")
	r.Write("context/goals/goal.md", "[f](note.md)\n")
	r.assertFails(
		"README.md:3: link missing.md does not resolve; "+linkFix,
		"README.md:3: link <gone dir/> does not resolve; "+linkFix,
		"context/goals/goal.md:1: link note.md does not resolve; "+linkFix,
		"context/note.md:1: link gone.md#x does not resolve; "+linkFix,
		"context/note.md:3: link ../nowhere.md does not resolve; "+linkFix,
		"docs/page.md:1: link /docs/missing.md does not resolve; "+linkFix)
}

func TestAnUnmatchedBacktickRunLeavesItsLeadingBackticks(t *testing.T) {
	r := newRepo(t)
	r.Write("README.md", "``[a](gone.md)` and [b](lost.md)\n")
	r.assertFails("README.md:1: link lost.md does not resolve; " + linkFix)
	r.Write("README.md", "x ```[a](gone.md)`` y [b](lost.md)\n")
	r.assertFails("README.md:1: link lost.md does not resolve; " + linkFix)
}

func TestLinksOutsideScannedFilesAreUnchecked(t *testing.T) {
	r := newRepo(t)
	r.Write("other/page.md", "[a](missing.md)\n")
	r.Write("docs/page.txt", "[a](missing.md)\n")
	r.assertClean()
}
