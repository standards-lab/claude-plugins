package marathon_test

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/standards-lab/claude-plugins/marathon/internal/marathon"
	"github.com/standards-lab/claude-plugins/marathon/internal/repotest"
)

const (
	keys = "keys: check, currency, gate, publish, merge, ci, checks"

	configuration = `[project]
kind = "context"
check = "printf 'check in %s' \"$PWD\""
gate = "printf '[%s]'"

[remote]
publish = "printf '<%s>'"
merge = "printf a && printf b"
ci = '''
printf 'ci %s,' "$#"
for arg in "$@"; do printf '<%s>' "$arg"; done
exit 3
'''
`
)

type result struct {
	status         int
	stdout, stderr string
}

type repo struct {
	*repotest.Repo
	t    *testing.T
	root string
}

func newRepo(t *testing.T) repo {
	t.Helper()
	r := repotest.New(t)
	parent, err := filepath.EvalSymlinks(r.Parent)
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv("GIT_CEILING_DIRECTORIES", parent)
	r.Git("init", "-q")
	rr := repo{r, t, filepath.Join(parent, "repo")}
	rr.Write(".claude/marathon.toml", configuration)
	return rr
}

func (r repo) marathonIn(dir string, args ...string) result {
	var stdout, stderr bytes.Buffer
	status := marathon.Run(dir, args, &stdout, &stderr)
	return result{status, stdout.String(), stderr.String()}
}

func (r repo) marathon(args ...string) result {
	return r.marathonIn(r.root, args...)
}

func (r repo) assertResult(got result, want result) {
	r.t.Helper()
	if got != want {
		r.t.Errorf("got %+v, want %+v", got, want)
	}
}

func (r repo) assertUsageFailure(got result, message string) {
	r.t.Helper()
	r.assertResult(got, result{2, "", "marathon: " + message + "\n"})
}

func TestChainedMergeRunsAsDeclared(t *testing.T) {
	r := newRepo(t)
	r.assertResult(r.marathon("merge"), result{0, "ab", ""})
}

func TestChainedMergeStopsAtAFailingFirstCommand(t *testing.T) {
	r := newRepo(t)
	r.Write(".claude/marathon.toml", "[remote]\nmerge = \"printf a && (exit 4) && printf b\"\n")
	r.assertResult(r.marathon("merge"), result{4, "a", ""})
	r.Write(".claude/marathon.toml", "[remote]\nmerge = \"(exit 5) && printf b\"\n")
	r.assertResult(r.marathon("merge"), result{5, "", ""})
}

func TestPublishArgumentsArriveQuoted(t *testing.T) {
	r := newRepo(t)
	tricky := "a b; echo $HOME && `id` 'q' \"d\""
	r.assertResult(r.marathon("publish", "--body-file", tricky),
		result{0, "<--body-file><" + tricky + ">", ""})
}

func TestArgumentsQuoteAsShlexQuotes(t *testing.T) {
	r := newRepo(t)
	declared := `printf '%s' "$BASH_EXECUTION_STRING" #`
	r.Write(".claude/marathon.toml", "[project]\ngate = '''"+declared+"'''\n")
	quoted := map[string]string{
		"":                 "''",
		"plain@%+=:,./-_9": "plain@%+=:,./-_9",
		"é":                "'é'",
		"it's":             `'it'"'"'s'`,
		"tab\there":        "'tab\there'",
		"$x":               "'$x'",
		"~home":            "'~home'",
	}
	for arg, want := range quoted {
		r.assertResult(r.marathon("gate", arg), result{0, declared + " " + want, ""})
	}
}

func TestGateTakesItsTag(t *testing.T) {
	r := newRepo(t)
	r.assertResult(r.marathon("gate", "v1.4.0"), result{0, "[v1.4.0]", ""})
}

func TestMultiLineCIRunsAsOneScriptWithPositionalArguments(t *testing.T) {
	r := newRepo(t)
	r.assertResult(r.marathon("ci", "x y", "$z;"), result{3, "ci 2,<x y><$z;>", ""})
}

func TestChecksRunsAndPropagatesItsStatus(t *testing.T) {
	r := newRepo(t)
	r.Write(".claude/marathon.toml", "[remote]\nchecks = \"printf checks && exit 6\"\n")
	r.assertResult(r.marathon("checks"), result{6, "checks", ""})
}

func TestUnsetChecks(t *testing.T) {
	r := newRepo(t)
	r.assertUsageFailure(r.marathon("checks"), "set [remote] checks in .claude/marathon.toml")
}

func TestSignaledCommandExitsAsAShellReportsIt(t *testing.T) {
	r := newRepo(t)
	r.Write(".claude/marathon.toml", "[project]\ncheck = \"kill -TERM $$\"\n")
	r.assertResult(r.marathon("check"), result{143, "", ""})
}

func TestCommandRunsFromTheRepositoryRoot(t *testing.T) {
	r := newRepo(t)
	subdir := filepath.Join(r.root, "sub/dir")
	if err := os.MkdirAll(subdir, 0o755); err != nil {
		t.Fatal(err)
	}
	r.assertResult(r.marathonIn(subdir, "check"), result{0, "check in " + r.root, ""})
}

func TestNoArguments(t *testing.T) {
	r := newRepo(t)
	r.assertResult(r.marathon(), result{2, "", "usage: marathon <key> [args]; " + keys + "\n"})
}

func TestUnknownKey(t *testing.T) {
	r := newRepo(t)
	r.assertUsageFailure(r.marathon("deploy"), "unknown key 'deploy'; "+keys)
}

func TestUnsetKey(t *testing.T) {
	r := newRepo(t)
	r.assertUsageFailure(r.marathon("currency"), "set [project] currency in .claude/marathon.toml")
}

func TestBlankOrNonStringKeyIsUnset(t *testing.T) {
	r := newRepo(t)
	for _, value := range []string{`" \t\u3000"`, "3", "[\"printf\"]"} {
		r.Write(".claude/marathon.toml", "[project]\ncheck = "+value+"\n")
		r.assertUsageFailure(r.marathon("check"), "set [project] check in .claude/marathon.toml")
	}
}

func TestKeyInTheWrongTableIsUnset(t *testing.T) {
	r := newRepo(t)
	r.Write(".claude/marathon.toml", "[project]\nmerge = \"printf no\"\n")
	r.assertUsageFailure(r.marathon("merge"), "set [remote] merge in .claude/marathon.toml")
}

func TestMissingConfiguration(t *testing.T) {
	r := newRepo(t)
	r.Remove(".claude/marathon.toml")
	r.assertUsageFailure(r.marathon("merge"), "no .claude/marathon.toml in "+r.root+
		"; write one (marathon's init does), then set [remote] merge in it")
}

func TestConfigurationThatDoesNotParse(t *testing.T) {
	r := newRepo(t)
	message := regexp.MustCompile(`^marathon: .*/\.claude/marathon\.toml does not parse: .*; fix it\n$`)
	for _, text := range []string{"[remote\n", "[remote]\nmerge = \"\xff\"\n"} {
		r.Write(".claude/marathon.toml", text)
		got := r.marathon("merge")
		if got.status != 2 || got.stdout != "" || !message.MatchString(got.stderr) {
			t.Errorf("%q: got %+v, want status 2 and a parse failure", text, got)
		}
	}
}

func TestConfigurationReadsWithUniversalNewlines(t *testing.T) {
	r := newRepo(t)
	r.Write(".claude/marathon.toml", "[remote]\r\nci = '''\r\nprintf a\rprintf b\r\n'''\r\n")
	r.assertResult(r.marathon("ci"), result{0, "ab", ""})
}

func TestOutsideAGitRepository(t *testing.T) {
	r := newRepo(t)
	outside := filepath.Join(r.root, "outside")
	if err := os.Mkdir(outside, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(filepath.Join(r.root, ".git"), filepath.Join(r.root, "not-git")); err != nil {
		t.Fatal(err)
	}
	r.assertUsageFailure(r.marathonIn(outside, "merge"), outside+
		" is not in a git repository; run marathon inside the repository whose "+
		".claude/marathon.toml declares the command")
}
