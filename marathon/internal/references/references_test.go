package references_test

import (
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
