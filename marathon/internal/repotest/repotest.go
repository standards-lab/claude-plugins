package repotest

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// Command is a tool's Run function: it runs on the repository at dir, writes
// to stdout and stderr, and returns the exit status.
type Command func(dir string, stdout, stderr io.Writer) int

// Repo is a temporary directory, Root, that a test builds a repository in.
// Root sits alone in Parent, so a test can put siblings beside it.
type Repo struct {
	t      *testing.T
	Parent string
	Root   string
}

// New creates an empty Repo, removed when the test ends.
func New(t *testing.T) *Repo {
	t.Helper()
	parent := t.TempDir()
	r := &Repo{t: t, Parent: parent, Root: filepath.Join(parent, "repo")}
	if err := os.Mkdir(r.Root, 0o755); err != nil {
		t.Fatal(err)
	}
	return r
}

// Write writes content to the file at rel under Root, creating its
// directories.
func (r *Repo) Write(rel, content string) {
	r.t.Helper()
	r.put(r.Root, rel, content)
}

// WriteBeside writes content to the file at rel under Parent, beside Root.
func (r *Repo) WriteBeside(rel, content string) {
	r.t.Helper()
	r.put(r.Parent, rel, content)
}

func (r *Repo) put(base, rel, content string) {
	r.t.Helper()
	name := filepath.Join(base, rel)
	if err := os.MkdirAll(filepath.Dir(name), 0o755); err != nil {
		r.t.Fatal(err)
	}
	if err := os.WriteFile(name, []byte(content), 0o644); err != nil {
		r.t.Fatal(err)
	}
}

// Remove removes the file or empty directory at rel under Root.
func (r *Repo) Remove(rel string) {
	r.t.Helper()
	if err := os.Remove(filepath.Join(r.Root, rel)); err != nil {
		r.t.Fatal(err)
	}
}

// Git runs git with args in Root, failing the test if it fails.
func (r *Repo) Git(args ...string) {
	r.t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = r.Root
	if out, err := cmd.CombinedOutput(); err != nil {
		r.t.Fatalf("git %v: %v\n%s", args, err, out)
	}
}

// Run runs cmd on Root and returns its exit status, its stdout split into
// lines, and its stderr.
func (r *Repo) Run(cmd Command) (int, []string, string) {
	var stdout, stderr bytes.Buffer
	code := cmd(r.Root, &stdout, &stderr)
	var lines []string
	if stdout.Len() != 0 {
		lines = strings.Split(strings.TrimSuffix(stdout.String(), "\n"), "\n")
	}
	return code, lines, stderr.String()
}

// AssertResult runs cmd on Root and fails the test unless it exits with
// wantCode, prints exactly wantLines, and writes nothing to stderr.
func (r *Repo) AssertResult(cmd Command, wantCode int, wantLines ...string) {
	r.t.Helper()
	code, lines, stderr := r.Run(cmd)
	if code != wantCode || !slices.Equal(lines, wantLines) || stderr != "" {
		r.t.Errorf("got (%d, %q, %q), want (%d, %q, \"\")",
			code, lines, stderr, wantCode, wantLines)
	}
}
