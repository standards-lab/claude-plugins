package version_test

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/standards-lab/claude-plugins/marathon/internal/version"
)

const stamp = "-X github.com/standards-lab/claude-plugins/marathon/internal/version.version="

var commands = []string{"lint-context", "lint-references", "marathon", "wrap-formats"}

func TestReportAnswersVersionAlone(t *testing.T) {
	var stdout bytes.Buffer
	if !version.Report("tool", []string{"--version"}, &stdout) {
		t.Fatal("Report(--version) = false, want true")
	}
	if got := stdout.String(); got != "tool dev\n" {
		t.Errorf("Report(--version) wrote %q, want %q", got, "tool dev\n")
	}
}

func TestReportLeavesOtherArgumentsAlone(t *testing.T) {
	for _, args := range [][]string{nil, {"check"}, {"-version"}, {"--version", "check"}, {"check", "--version"}} {
		var stdout bytes.Buffer
		if version.Report("tool", args, &stdout) || stdout.Len() != 0 {
			t.Errorf("Report(%q) answered with %q, want false and no output", args, stdout.String())
		}
	}
}

func TestEachCommandReportsItsBuildVersion(t *testing.T) {
	for _, build := range []struct{ ldflags, want string }{
		{"", "dev"},
		{stamp + "0.21.0", "0.21.0"},
	} {
		bin := t.TempDir()
		cmd := exec.Command("go", "build", "-C", filepath.Join("..", "..", ".."),
			"-ldflags", build.ldflags, "-o", bin+string(filepath.Separator), "./cmd/...")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("go build -ldflags %q: %v\n%s", build.ldflags, err, out)
		}
		for _, name := range commands {
			out, err := exec.Command(filepath.Join(bin, name), "--version").Output()
			want := name + " " + build.want + "\n"
			if err != nil || string(out) != want {
				t.Errorf("%s --version built with %q: got (%q, %v), want (%q, nil)",
					name, build.ldflags, out, err, strings.TrimSpace(want))
			}
		}
	}
}
