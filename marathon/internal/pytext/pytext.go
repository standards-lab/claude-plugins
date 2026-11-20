package pytext

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// SpaceClass is the body of a regular-expression character class, without
// its brackets, that matches exactly the runes [IsSpace] accepts.
const SpaceClass = `\t\n\v\f\r \x1c-\x1f\x{85}\x{a0}\x{1680}\x{2000}-\x{200a}\x{2028}\x{2029}\x{202f}\x{205f}\x{3000}`

// IsSpace reports whether r is whitespace to Python's str.isspace.
func IsSpace(r rune) bool {
	return unicode.IsSpace(r) || '\x1c' <= r && r <= '\x1f'
}

// Fields splits s around each run of Python whitespace, as str.split with no
// arguments does.
func Fields(s string) []string {
	return strings.FieldsFunc(s, IsSpace)
}

// TrimSpace returns s without its leading and trailing Python whitespace, as
// str.strip with no arguments does.
func TrimSpace(s string) string {
	return strings.TrimFunc(s, IsSpace)
}

// Lines splits s at Python's line boundaries, as str.splitlines does: the
// lines carry no boundary, and a boundary ending s starts no further line.
func Lines(s string) []string {
	var lines []string
	start := 0
	for i, r := range s {
		if i < start {
			continue
		}
		width := utf8.RuneLen(r)
		switch r {
		case '\r':
			if strings.HasPrefix(s[i+1:], "\n") {
				width++
			}
		case '\n', '\v', '\f', '\x1c', '\x1d', '\x1e', '\u0085', '\u2028', '\u2029':
		default:
			continue
		}
		lines = append(lines, s[start:i])
		start = i + width
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

// UniversalNewlines returns s with each \r\n and each lone \r turned into \n,
// as Python translates the text it reads in universal newlines mode.
func UniversalNewlines(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\r\n", "\n"), "\r", "\n")
}
