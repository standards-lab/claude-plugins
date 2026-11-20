package wrapformats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/standards-lab/claude-plugins/marathon/internal/pytext"
)

const width = 100

var (
	formatHeading = regexp.MustCompile(`^[ \t]*(?:>[ \t]?)*(?:` +
		`(?:#{1,6}[ \t]+)?[*_]{0,3}(?:PLAN ROUND \p{Nd}|INTAKE ROUND \p{Nd}` +
		`|ESCALATION(?:[^\p{L}\p{N}_]|$)|Task brief ·` +
		`|(?:── )?(?:Status|STATUS) ·|(?:Retro|RETRO) ·)` +
		`|#{2,3}[ \t]+[*_]{0,3}[^` + pytext.SpaceClass + `].* · task \p{Nd}+/\p{Nd}+ ·)`)
	fence      = regexp.MustCompile("^[ \t>]*(?:```|~~~)")
	marker     = regexp.MustCompile(`^([ \t]*(?:>[ \t]?)*[ \t]*)((?:[-*+]|\p{Nd}+[.)]|#{1,6})[ \t]+)?`)
	blockStart = regexp.MustCompile(`^(?:[*_]*\p{Nd}+[.)]|[-*+]|#{1,6}|[>|].*|(?:` + "```|~~~" +
		`).*|[-=]+|[*_]{3,})$`)
)

type proseLine struct {
	text  string
	prose bool
}

func proseLines(text string) []proseLine {
	inFence := false
	var lines []proseLine
	for line := range strings.SplitSeq(text, "\n") {
		if fence.MatchString(line) {
			inFence = !inFence
			lines = append(lines, proseLine{line, false})
			continue
		}
		table := strings.HasPrefix(strings.TrimLeft(line, " \t>"), "|")
		lines = append(lines, proseLine{line, !inFence && !table})
	}
	return lines
}

func length(s string) int {
	return utf8.RuneCountInString(s)
}

func holdsFormat(text string) bool {
	return slices.ContainsFunc(proseLines(text), func(l proseLine) bool {
		return l.prose && formatHeading.MatchString(l.text)
	})
}

func overlong(text string) int {
	count := 0
	for _, l := range proseLines(text) {
		if l.prose && length(l.text) > width {
			count++
		}
	}
	return count
}

func filled(words []string, first, rest string) string {
	var lines, line []string
	prefix := func() string {
		if len(lines) == 0 {
			return first
		}
		return rest
	}
	for _, word := range words {
		if len(line) == 0 || length(prefix()+strings.Join(line, " ")+" "+word) <= width {
			line = append(line, word)
			continue
		}
		carried := []string{word}
		for blockStart.MatchString(carried[0]) && len(line) > 1 {
			carried = slices.Insert(carried, 0, line[len(line)-1])
			line = line[:len(line)-1]
		}
		if blockStart.MatchString(carried[0]) {
			line = append(line, carried...)
		} else {
			lines = append(lines, prefix()+strings.Join(line, " "))
			line = carried
		}
	}
	lines = append(lines, prefix()+strings.Join(line, " "))
	return strings.Join(lines, "\n")
}

func wrapped(line string) string {
	groups := marker.FindStringSubmatch(line)
	lead, item := groups[1], groups[2]
	continuation := strings.Repeat(" ", length(lead+item))
	if strings.Contains(lead, ">") {
		continuation = lead + strings.Repeat(" ", length(item))
	}
	return filled(pytext.Fields(line[len(lead+item):]), lead+item, continuation)
}

func rewrappedLine(line string) string {
	lines := wrapped(line)
	for part := range strings.SplitSeq(lines, "\n") {
		if length(part) > width {
			return line
		}
	}
	return lines
}

func rewrap(text string) string {
	var lines []string
	for _, l := range proseLines(text) {
		if l.prose && length(l.text) > width {
			lines = append(lines, rewrappedLine(l.text))
		} else {
			lines = append(lines, l.text)
		}
	}
	return strings.Join(lines, "\n")
}

type field struct {
	key, value string
}

func verdict(hookInput map[string]any) []field {
	reply, isString := hookInput["last_assistant_message"].(string)
	if truthy(hookInput["stop_hook_active"]) || !isString {
		return nil
	}
	if !holdsFormat(reply) || overlong(reply) == 0 {
		return nil
	}
	rewrapped := rewrap(reply)
	if rewrapped == reply {
		return []field{{"systemMessage", fmt.Sprintf("marathon: %d line(s) of this reply run past "+
			"%d characters on a single word, so they stay unwrapped.", overlong(reply), width)}}
	}
	return []field{
		{"decision", "block"},
		{"reason", fmt.Sprintf("marathon prints its formats wrapped at %d columns, and this reply "+
			"has longer lines. Print the reply again exactly as below, rewrapped at %d columns, "+
			"with nothing before or after it:\n\n", width, width) + rewrapped},
	}
}

func truthy(value any) bool {
	switch v := value.(type) {
	case nil:
		return false
	case bool:
		return v
	case json.Number:
		f, _ := strconv.ParseFloat(string(v), 64)
		return f != 0
	case string:
		return v != ""
	case []any:
		return len(v) != 0
	case map[string]any:
		return len(v) != 0
	}
	return true
}

func pythonJSON(fields []field) string {
	var b strings.Builder
	b.WriteByte('{')
	for i, f := range fields {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(pythonJSONString(f.key) + ": " + pythonJSONString(f.value))
	}
	b.WriteByte('}')
	return b.String()
}

func pythonJSONString(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for _, r := range s {
		switch {
		case r == '"' || r == '\\':
			b.WriteString(`\` + string(r))
		case r == '\n':
			b.WriteString(`\n`)
		case r == '\r':
			b.WriteString(`\r`)
		case r == '\t':
			b.WriteString(`\t`)
		case r == '\b':
			b.WriteString(`\b`)
		case r == '\f':
			b.WriteString(`\f`)
		case ' ' <= r && r <= '~':
			b.WriteRune(r)
		case r < 0x10000:
			fmt.Fprintf(&b, `\u%04x`, r)
		default:
			r -= 0x10000
			fmt.Fprintf(&b, `\u%04x\u%04x`, 0xd800+(r>>10), 0xdc00+(r&0x3ff))
		}
	}
	b.WriteByte('"')
	return b.String()
}

// Run reads the Stop hook's input from stdin, writes the hook's output to
// stdout when the reply needs one, and returns the exit status, which is
// always 0.
func Run(stdin io.Reader, stdout io.Writer) int {
	input, err := io.ReadAll(stdin)
	if err != nil || !utf8.Valid(input) || !json.Valid(input) {
		return 0
	}
	decoder := json.NewDecoder(bytes.NewReader(input))
	decoder.UseNumber()
	var hookInput any
	if decoder.Decode(&hookInput) != nil {
		return 0
	}
	if object, ok := hookInput.(map[string]any); ok {
		if output := verdict(object); output != nil {
			_, _ = fmt.Fprintln(stdout, pythonJSON(output))
		}
	}
	return 0
}
