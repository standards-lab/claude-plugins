package references

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"unicode"
	"unicode/utf8"

	"github.com/standards-lab/claude-plugins/marathon/internal/pytext"
)

const fix = "fix the path or remove the reference"

// Run checks the repository at dir, printing one line per failing or skipped
// reference to stdout, and returns the exit status: 1 if any reference fails
// or the repository can't be read, with the reason on stderr, and 0 otherwise.
func Run(dir string, stdout, stderr io.Writer) int {
	lines, failed, err := check(dir)
	var report strings.Builder
	for _, line := range lines {
		report.WriteString(line + "\n")
	}
	if _, writeErr := io.WriteString(stdout, report.String()); writeErr != nil && err == nil {
		err = writeErr
	}
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "lint-references: %v\n", err)
		return 1
	}
	if failed {
		return 1
	}
	return 0
}

func check(dir string) ([]string, bool, error) {
	root, err := filepath.Abs(dir)
	if err != nil {
		return nil, false, err
	}
	files, err := trackedNonMarkdownFiles(root)
	if err != nil {
		return nil, false, err
	}
	var lines []string
	failed := false
	for _, rel := range files {
		text, ok, err := readUTF8(filepath.Join(root, rel))
		if err != nil {
			return lines, failed, err
		}
		if !ok {
			continue
		}
		for _, ref := range numberedReferences(text) {
			line, fails := outcome(root, rel, ref)
			if line != "" {
				lines = append(lines, line)
				failed = failed || fails
			}
		}
	}
	return lines, failed, nil
}

func trackedNonMarkdownFiles(root string) ([]string, error) {
	cmd := exec.Command("git", "ls-files", "-z")
	cmd.Dir = root
	var gitStderr bytes.Buffer
	cmd.Stderr = &gitStderr
	listing, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git ls-files: %w: %s", err, strings.TrimSpace(gitStderr.String()))
	}
	var files []string
	for rel := range strings.SplitSeq(string(listing), "\x00") {
		if rel != "" && !strings.HasSuffix(rel, ".md") {
			files = append(files, rel)
		}
	}
	return files, nil
}

func readUTF8(name string) (string, bool, error) {
	content, err := os.ReadFile(name)
	if errors.Is(err, fs.ErrNotExist) || errors.Is(err, syscall.EISDIR) {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	if !utf8.Valid(content) {
		return "", false, nil
	}
	return string(content), true, nil
}

type reference struct {
	number int
	path   string
}

func numberedReferences(text string) []reference {
	var refs []reference
	inBlockComment := false
	for i, line := range pytext.Lines(text) {
		trimmed := strings.TrimLeftFunc(line, unicode.IsSpace)
		opensBlock := strings.HasPrefix(trimmed, "/*")
		isReferenceLine := opensBlock ||
			strings.HasPrefix(trimmed, "#") ||
			strings.HasPrefix(trimmed, "//") ||
			strings.HasPrefix(trimmed, "--") ||
			(inBlockComment && strings.HasPrefix(trimmed, "*"))
		switch {
		case opensBlock:
			inBlockComment = !strings.Contains(trimmed[2:], "*/")
		case inBlockComment:
			inBlockComment = !strings.Contains(trimmed, "*/")
		}
		if !isReferenceLine {
			continue
		}
		for _, p := range referencePaths(line) {
			refs = append(refs, reference{number: i + 1, path: p})
		}
	}
	return refs
}

func referencePaths(line string) []string {
	var paths []string
	for i := 0; i < len(line); {
		end := referenceEnd(line, i)
		if end == 0 {
			i++
			continue
		}
		paths = append(paths, line[i:end])
		i = end
	}
	return paths
}

func referenceEnd(line string, start int) int {
	if !isStartChar(line[start]) || (start > 0 && precedesNonReference(line[start-1])) {
		return 0
	}
	for end := start + 1; end <= len(line); end++ {
		if end >= start+4 && line[end-3:end] == ".md" && endsReference(line[end:]) {
			return end
		}
		if end < len(line) && !isPathChar(line[end]) {
			return 0
		}
	}
	return 0
}

func endsReference(rest string) bool {
	if rest == "" {
		return true
	}
	next := rest[0]
	if isAlnum(next) || next == '_' || next == '/' || next == '-' {
		return false
	}
	return next != '.' || len(rest) == 1 || !isAlnum(rest[1]) && rest[1] != '_'
}

func isAlnum(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9'
}

func isStartChar(c byte) bool {
	return isAlnum(c) || c == '_' || c == '.'
}

func isPathChar(c byte) bool {
	return isStartChar(c) || c == '/' || c == '-'
}

func precedesNonReference(c byte) bool {
	return isPathChar(c) || strings.IndexByte("*${}<>", c) >= 0
}

func outcome(root, rel string, ref reference) (string, bool) {
	parts := strings.Split(path.Clean(ref.path), "/")
	ups := 0
	for _, part := range parts {
		if part == ".." {
			ups++
		}
	}
	base := root + strings.Repeat("/..", ups)
	segment := parts[ups]
	if ups == 0 && len(parts) > 1 && !exists(filepath.Join(root, segment)) {
		base = filepath.Dir(root)
	}
	target := base + "/" + strings.Join(parts[ups:], "/")
	if exists(target) {
		return "", false
	}
	hasMoreSegments := len(parts)-ups > 1
	if hasMoreSegments && !exists(base+"/"+segment) {
		return fmt.Sprintf("%s:%d: skipped %s: no %s/ in this repository or beside it",
			rel, ref.number, ref.path, segment), false
	}
	return fmt.Sprintf("%s:%d: reference %s does not resolve; %s", rel, ref.number, ref.path, fix), true
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}
