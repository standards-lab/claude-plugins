package pytext_test

import (
	"regexp"
	"slices"
	"testing"
	"unicode"

	"github.com/standards-lab/claude-plugins/marathon/internal/pytext"
)

var pythonSpace = []rune{
	'\t', '\n', '\v', '\f', '\r', '\x1c', '\x1d', '\x1e', '\x1f', ' ', '\u0085', '\u00a0',
	'\u1680', '\u2000', '\u2001', '\u2002', '\u2003', '\u2004', '\u2005', '\u2006', '\u2007',
	'\u2008', '\u2009', '\u200a', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000',
}

func TestIsSpaceAcceptsExactlyPythonWhitespace(t *testing.T) {
	for r := rune(0); r <= unicode.MaxRune; r++ {
		if got, want := pytext.IsSpace(r), slices.Contains(pythonSpace, r); got != want {
			t.Errorf("IsSpace(%U) = %v, want %v", r, got, want)
		}
	}
}

func TestSpaceClassMatchesExactlyPythonWhitespace(t *testing.T) {
	class := regexp.MustCompile(`^[` + pytext.SpaceClass + `]$`)
	for r := rune(0); r <= unicode.MaxRune; r++ {
		if got, want := class.MatchString(string(r)), slices.Contains(pythonSpace, r); got != want {
			t.Errorf("SpaceClass on %U = %v, want %v", r, got, want)
		}
	}
}

func TestFieldsSplitsOnRunsOfPythonWhitespace(t *testing.T) {
	got := pytext.Fields("\x1c one\u3000\u3000two\x1fthree\u200b ")
	want := []string{"one", "two", "three\u200b"}
	if !slices.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestTrimSpaceTrimsPythonWhitespace(t *testing.T) {
	if got := pytext.TrimSpace("\x1f\u00a0 a b\u2029\t"); got != "a b" {
		t.Errorf("got %q, want %q", got, "a b")
	}
}

func TestLinesBreaksAtEachPythonBoundary(t *testing.T) {
	text := "a\nb\rc\r\nd\ve\ff\x1cg\x1dh\x1ei\u0085j\u2028k\u2029l\x1fm\u00a0n\n"
	want := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l\x1fm\u00a0n"}
	if got := pytext.Lines(text); !slices.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestLinesKeepsEmptyLinesButNoTrailingOne(t *testing.T) {
	want := []string{"", "a", ""}
	if got := pytext.Lines("\na\n\n"); !slices.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
	if got := pytext.Lines(""); got != nil {
		t.Errorf("got %q, want none", got)
	}
}
