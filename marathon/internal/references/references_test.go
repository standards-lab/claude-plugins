package references_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/standards-lab/claude-plugins/marathon/internal/references"
	"github.com/standards-lab/claude-plugins/marathon/internal/repotest"
)

const fix = "fix the path or remove the reference"

type repo struct {
	*repotest.Repo
	t *testing.T
}

func newRepo(t *testing.T) repo {
	t.Helper()
	r := repo{repotest.New(t), t}
	r.Git("init", "-q")
	r.write("docs/guide.md", "# Guide\n")
	return r
}

func (r repo) write(rel, content string) {
	r.t.Helper()
	r.Write(rel, content)
	r.Git("add", rel)
}

func (r repo) writeUntracked(rel, content string) {
	r.t.Helper()
	r.Write(rel, content)
}

func (r repo) sibling(rel string) {
	r.t.Helper()
	r.WriteBeside(rel, "# Page\n")
}

func (r repo) assertResult(wantCode int, wantLines ...string) {
	r.t.Helper()
	r.AssertResult(references.Run, wantCode, wantLines...)
}

func TestEachMarkerMakesAReferenceLine(t *testing.T) {
	for _, marker := range []string{"#", "//", "--", "/*"} {
		t.Run(marker, func(t *testing.T) {
			r := newRepo(t)
			r.write("src/code", "  "+marker+" See docs/missing.md\n")
			r.assertResult(1, "src/code:1: reference docs/missing.md does not resolve; "+fix)
		})
	}
}

func TestAStarLineInsideABlockCommentIsAReferenceLine(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "/*\n * See docs/missing.md\n */\n")
	r.assertResult(1, "src/code:2: reference docs/missing.md does not resolve; "+fix)
}

func TestAStarLineOutsideABlockCommentIsNotAReferenceLine(t *testing.T) {
	r := newRepo(t)
	r.write(".gitignore", "* docs/missing.md\n")
	r.write("src/code", "/* Closed. */\n* See docs/missing.md\n/*\n */\n* See docs/missing.md\n")
	r.assertResult(0)
}

func TestAStarLineInsideABlockOpenedMidLineIsAReferenceLine(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "int x; /* see\n * docs/missing.md\n */\n")
	r.assertResult(1, "src/code:2: reference docs/missing.md does not resolve; "+fix)
}

func TestAStarLineInsideABlockReopenedAfterOneClosedIsAReferenceLine(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "/* a */ int x; /* b\n * docs/missing.md\n */\n")
	r.assertResult(1, "src/code:2: reference docs/missing.md does not resolve; "+fix)
}

func TestAStarLineInsideABlockLeftOpenAfterAnOpenerIsAReferenceLine(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "/*/\n * docs/missing.md\n */\n")
	r.assertResult(1, "src/code:2: reference docs/missing.md does not resolve; "+fix)
}

func TestAStarLineAfterABlockClosedMidLineIsNotAReferenceLine(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "int x; /* see */ int y;\n* docs/missing.md\n")
	r.write("src/other", "/*\n */ int x; /* again */\n* docs/missing.md\n")
	r.assertResult(0)
}

func TestAnOpenerInsideAPatternOrLineCommentDoesNotOpenABlock(t *testing.T) {
	r := newRepo(t)
	r.write(".gitignore", "build/*\n* docs/missing.md\n")
	r.write("src/code", "x = 1; // not /* a block\n* docs/missing.md\n")
	r.write("src/script", "url = \"http://host/*\"\n* docs/missing.md\n")
	r.assertResult(0)
}

func TestAnOpenerInsideAnyLineCommentDoesNotOpenABlock(t *testing.T) {
	r := newRepo(t)
	r.write(".gitignore", "# drop everything under /* at root\n* docs/missing.md\n")
	r.write("s.sh", "echo hi # glob /* here\n* docs/missing.md\n")
	r.write("q.sql", "select 1; -- not /* a block\n* docs/missing.md\n")
	r.assertResult(0)
}

func TestAnOpenerAfterPunctuationOpensABlock(t *testing.T) {
	for _, before := range []string{";", ")", "}", ",", "="} {
		t.Run(before, func(t *testing.T) {
			r := newRepo(t)
			r.write("b.c", "x"+before+"/* tight\n * docs/missing.md\n */\n")
			r.assertResult(1, "b.c:2: reference docs/missing.md does not resolve; "+fix)
		})
	}
}

func TestAnOpenerInsideAPathOrGlobDoesNotOpenABlock(t *testing.T) {
	for _, line := range []string{"build/*", "http://host/*", "src/**/*.go", "~/*", "[ab]/*", "\"/*\""} {
		t.Run(line, func(t *testing.T) {
			r := newRepo(t)
			r.write("src/script", "x = "+line+"\n* docs/missing.md\n")
			r.assertResult(0)
		})
	}
}

func TestOtherLinesAreNotReferenceLines(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "path = \"docs/missing.md\"\nx = 1  # docs/missing.md\n")
	r.assertResult(0)
}

func TestMarkdownFilesAreIgnored(t *testing.T) {
	r := newRepo(t)
	r.write("docs/page.md", "# See docs/missing.md\n")
	r.assertResult(0)
}

func TestUntrackedFilesAreIgnored(t *testing.T) {
	r := newRepo(t)
	r.writeUntracked("src/code", "# See docs/missing.md\n")
	r.assertResult(0)
}

func TestResolvesAtTheRoot(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "# See docs/guide.md, and ./docs/guide.md.\n")
	r.assertResult(0)
}

func TestResolvesInASibling(t *testing.T) {
	r := newRepo(t)
	r.sibling("architecture/principles/comments.md")
	r.write("src/code",
		"# See architecture/principles/comments.md\n"+
			"# See ../architecture/principles/comments.md\n")
	r.assertResult(0)
}

func TestTopLevelDirectoryShadowsTheSibling(t *testing.T) {
	r := newRepo(t)
	r.sibling("architecture/principles/comments.md")
	r.write("architecture/principles/other.md", "# Other\n")
	r.write("src/code", "# See architecture/principles/comments.md\n")
	r.assertResult(1,
		"src/code:1: reference architecture/principles/comments.md does not resolve; "+fix)
}

func TestMissingInTheRepositoryFails(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "\n# See docs/missing.md\n")
	r.assertResult(1, "src/code:2: reference docs/missing.md does not resolve; "+fix)
}

func TestMissingInAnExistingSiblingFails(t *testing.T) {
	r := newRepo(t)
	r.sibling("architecture/principles/comments.md")
	r.write("src/code", "# See architecture/principles/missing.md\n")
	r.assertResult(1,
		"src/code:1: reference architecture/principles/missing.md does not resolve; "+fix)
}

func TestOneSegmentMissingAtTheRootFails(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "# See NOTES.md\n")
	r.assertResult(1, "src/code:1: reference NOTES.md does not resolve; "+fix)
}

func TestNoSiblingSkipsAndPasses(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "# See architecture/principles/comments.md\n")
	r.assertResult(0,
		"src/code:1: skipped architecture/principles/comments.md: "+
			"no architecture/ in this repository or beside it")
}

func TestURLPathsAreNotReferences(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "# See https://example.com/docs/missing.md\n")
	r.assertResult(0)
}

func TestPatternsAndTemplatesAreNotReferences(t *testing.T) {
	r := newRepo(t)
	r.write("src/code", "# Every *.md file, and each <run>.last-message.md\n")
	r.assertResult(0)
}

func TestNonUTF8FilesAreIgnored(t *testing.T) {
	r := newRepo(t)
	r.write("src/blob", "# See docs/missing.md \xff\xfe\n")
	r.assertResult(0)
}

func TestASkipWithAFailureFails(t *testing.T) {
	r := newRepo(t)
	r.write("src/code",
		"# See architecture/principles/comments.md\n# See docs/missing.md\n")
	r.assertResult(1,
		"src/code:1: skipped architecture/principles/comments.md: "+
			"no architecture/ in this repository or beside it",
		"src/code:2: reference docs/missing.md does not resolve; "+fix)
}

const reference = "# See architecture/principles/comments.md\n"

func newCoordinator(t *testing.T, config string) *repotest.Repo {
	t.Helper()
	c := repotest.New(t)
	c.Write(".claude/marathon.toml", config)
	return c
}

func (r repo) member(coordinator string) {
	r.t.Helper()
	r.write(".claude/marathon.toml", "[project]\ncoordinator = \""+coordinator+"\"\n")
}

func TestAHomeRelativeKeyResolvesBesideTheCoordinator(t *testing.T) {
	c := newCoordinator(t, "[workspace]\n")
	c.WriteBeside("architecture/principles/comments.md", "# Page\n")
	t.Setenv("HOME", c.Parent)
	r := newRepo(t)
	r.member("~/repo")
	r.write("src/code", reference)
	r.assertResult(0)
}

func TestARelativeKeyIsReadFromTheRepositoryRoot(t *testing.T) {
	c := newCoordinator(t, "[workspace]\n")
	c.WriteBeside("architecture/principles/comments.md", "# Page\n")
	r := newRepo(t)
	key, err := filepath.Rel(r.Root, c.Root)
	if err != nil {
		t.Fatal(err)
	}
	r.member(key)
	r.write("src/code", reference)
	r.assertResult(0)
}

func TestTheKeyIsReadAtTheGitTopLevel(t *testing.T) {
	r := newRepo(t)
	missing := filepath.Join(r.Parent, "absent")
	r.member(missing)
	r.write("sub/code", reference)
	var stdout, stderr bytes.Buffer
	code := references.Run(filepath.Join(r.Root, "sub"), &stdout, &stderr)
	want := "code:1: skipped architecture/principles/comments.md: no architecture/ in this " +
		"repository or beside it, and coordinator " + missing + " does not exist\n"
	if code != 0 || stdout.String() != want || stderr.Len() != 0 {
		t.Errorf("got (%d, %q, %q), want (0, %q, \"\")", code, stdout.String(), stderr.String(), want)
	}
}

func TestAWorkspaceIsItsOwnCoordinator(t *testing.T) {
	elsewhere := repotest.New(t)
	elsewhere.Write("principles/comments.md", "# Page\n")
	r := newRepo(t)
	r.write(".claude/marathon.toml",
		"[workspace.paths]\narchitecture = \""+elsewhere.Root+"\"\n")
	r.write("src/code", reference)
	r.assertResult(0)
}

func TestAWorkspacePathResolves(t *testing.T) {
	for name, location := range map[string]func(home string) string{
		"absolute":      func(home string) string { return home + "/arch" },
		"home-relative": func(string) string { return "~/arch" },
	} {
		t.Run(name, func(t *testing.T) {
			home := t.TempDir()
			t.Setenv("HOME", home)
			c := newCoordinator(t,
				"[workspace.paths]\narchitecture = \""+location(home)+"\"\n")
			c.WriteBeside("architecture/principles/other.md", "# Other\n")
			if err := os.MkdirAll(home+"/arch/principles", 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(home+"/arch/principles/comments.md", nil, 0o644); err != nil {
				t.Fatal(err)
			}
			r := newRepo(t)
			r.member(c.Root)
			r.write("src/code", reference)
			r.assertResult(0)
		})
	}
}

func TestTheFirstPlaceWithTheSegmentDecides(t *testing.T) {
	for _, place := range []string{"repository", "beside", "paths"} {
		t.Run(place, func(t *testing.T) {
			elsewhere := repotest.New(t)
			c := newCoordinator(t,
				"[workspace.paths]\narchitecture = \""+elsewhere.Root+"\"\n")
			c.WriteBeside("architecture/principles/comments.md", "# Page\n")
			elsewhere.Write("principles/comments.md", "# Page\n")
			r := newRepo(t)
			r.member(c.Root)
			switch place {
			case "repository":
				r.write("architecture/principles/other.md", "# Other\n")
			case "beside":
				r.WriteBeside("architecture/principles/other.md", "# Other\n")
			case "paths":
				elsewhere.Remove("principles/comments.md")
			}
			r.write("src/code", reference)
			r.assertResult(1,
				"src/code:1: reference architecture/principles/comments.md does not resolve; "+fix)
		})
	}
}

func TestMissingWhereTheCoordinatorResolvedFails(t *testing.T) {
	c := newCoordinator(t, "[workspace]\n")
	c.WriteBeside("architecture/principles/other.md", "# Other\n")
	r := newRepo(t)
	r.member(c.Root)
	r.write("src/code", reference)
	r.assertResult(1,
		"src/code:1: reference architecture/principles/comments.md does not resolve; "+fix)
}

func TestTheCoordinatorDoesNotServeParentOrOneSegmentReferences(t *testing.T) {
	c := newCoordinator(t, "[workspace.paths]\n\"NOTES.md\" = \"~/NOTES.md\"\n")
	c.WriteBeside("architecture/principles/comments.md", "# Page\n")
	c.WriteBeside("NOTES.md", "# Notes\n")
	t.Setenv("HOME", c.Parent)
	r := newRepo(t)
	r.member(c.Root)
	r.write("src/code", "# See ../architecture/principles/comments.md\n# See NOTES.md\n")
	r.assertResult(1,
		"src/code:1: skipped ../architecture/principles/comments.md: "+
			"no architecture/ in this repository or beside it",
		"src/code:2: reference NOTES.md does not resolve; "+fix)
}

func TestASegmentFoundNowhereIsSkipped(t *testing.T) {
	c := newCoordinator(t, "[workspace]\n")
	r := newRepo(t)
	r.member(c.Root)
	r.write("src/code", reference)
	r.assertResult(0,
		"src/code:1: skipped architecture/principles/comments.md: no architecture/ in this "+
			"repository, beside it, in the [workspace.paths] of coordinator "+c.Root+", or beside it")
}

func TestAMissingCoordinatorIsNamedInTheSkip(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	r := newRepo(t)
	r.member("~/architecture")
	r.write("src/code", reference)
	r.assertResult(0,
		"src/code:1: skipped architecture/principles/comments.md: no architecture/ in this "+
			"repository or beside it, and coordinator "+home+"/architecture does not exist")
}

func TestAMisconfiguredCoordinatorFails(t *testing.T) {
	for name, config := range map[string]string{
		"no configuration":    "",
		"no workspace table":  "[project]\nkind = \"context\"\n",
		"unparseable":         "[workspace\n",
		"unreadable":          "directory",
		"non-string location": "[workspace.paths]\narchitecture = 1\n",
	} {
		t.Run(name, func(t *testing.T) {
			c := repotest.New(t)
			switch config {
			case "":
			case "directory":
				c.Write(".claude/marathon.toml/inner", "")
			default:
				c.Write(".claude/marathon.toml", config)
			}
			r := newRepo(t)
			r.member(c.Root)
			r.write("src/code", reference)
			code, lines, stderr := r.Run(references.Run)
			if code != 1 || lines != nil || !strings.Contains(stderr, c.Root) ||
				!strings.Contains(stderr, "; ") {
				t.Errorf("got (%d, %q, %q), want exit 1 with a reason naming %s and its fix",
					code, lines, stderr, c.Root)
			}
		})
	}
}

func TestAKeyThatIsNotAPathFails(t *testing.T) {
	r := newRepo(t)
	r.write(".claude/marathon.toml", "[project]\ncoordinator = 1\n")
	r.write("src/code", reference)
	code, _, stderr := r.Run(references.Run)
	if code != 1 || !strings.Contains(stderr, "set it to the coordinator's directory") {
		t.Errorf("got (%d, %q), want exit 1 naming the fix", code, stderr)
	}
}

func TestAConfigurationWithoutTheKeyChangesNothing(t *testing.T) {
	r := newRepo(t)
	r.write(".claude/marathon.toml", "[project]\nkind = \"code\"\n")
	r.write("src/code", reference)
	r.assertResult(0,
		"src/code:1: skipped architecture/principles/comments.md: "+
			"no architecture/ in this repository or beside it")
}
