package wrapformats

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
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
	unsafeName = regexp.MustCompile(`[^A-Za-z0-9_-]`)
)

func length(s string) int {
	return utf8.RuneCountInString(s)
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

type state struct {
	format, fence bool
}

func statePath(sessionID, messageID string) string {
	return filepath.Join(os.TempDir(), "marathon-wrap-formats."+
		unsafeName.ReplaceAllString(sessionID, "_")+"."+unsafeName.ReplaceAllString(messageID, "_"))
}

func loadState(path string) state {
	content, err := os.ReadFile(path)
	if err != nil {
		return state{}
	}
	words := strings.Fields(string(content))
	return state{slices.Contains(words, "format"), slices.Contains(words, "fence")}
}

func saveState(path string, s state) {
	if s == (state{}) {
		_ = os.Remove(path)
		return
	}
	var words []string
	if s.format {
		words = append(words, "format")
	}
	if s.fence {
		words = append(words, "fence")
	}
	_ = os.WriteFile(path, []byte(strings.Join(words, " ")+"\n"), 0o600)
}

func displayed(delta string, s state) (string, state) {
	lines := strings.Split(delta, "\n")
	for i, line := range lines {
		if fence.MatchString(line) {
			s.fence = !s.fence
			continue
		}
		prose := !s.fence && !strings.HasPrefix(strings.TrimLeft(line, " \t>"), "|")
		if prose && formatHeading.MatchString(line) {
			s.format = true
		}
		if prose && s.format && length(line) > width {
			lines[i] = rewrappedLine(line)
		}
	}
	return strings.Join(lines, "\n"), s
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

// Run reads one MessageDisplay batch from stdin, writes the batch rewrapped
// to stdout when the message holds a marathon format and the rewrap changes
// the batch, and returns the exit status, which is always 0.
func Run(stdin io.Reader, stdout io.Writer) int {
	input, err := io.ReadAll(stdin)
	if err != nil || !utf8.Valid(input) || !json.Valid(input) {
		return 0
	}
	var batch struct {
		SessionID string  `json:"session_id"`
		MessageID *string `json:"message_id"`
		Final     bool    `json:"final"`
		Delta     *string `json:"delta"`
	}
	if json.Unmarshal(input, &batch) != nil || batch.MessageID == nil || batch.Delta == nil {
		return 0
	}
	path := statePath(batch.SessionID, *batch.MessageID)
	before := loadState(path)
	shown, after := displayed(*batch.Delta, before)
	if batch.Final {
		_ = os.Remove(path)
	} else if after != before {
		saveState(path, after)
	}
	if shown != *batch.Delta {
		_, _ = fmt.Fprintln(stdout, `{"hookSpecificOutput": {"hookEventName": "MessageDisplay", `+
			`"displayContent": `+pythonJSONString(shown)+`}}`)
	}
	return 0
}
