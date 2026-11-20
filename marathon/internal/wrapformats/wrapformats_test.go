package wrapformats_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/standards-lab/claude-plugins/marathon/internal/pytext"
	"github.com/standards-lab/claude-plugins/marathon/internal/wrapformats"
)

var (
	longWord    = "https://example.com/" + strings.Repeat("a", 100)
	prose       = strings.TrimSuffix(strings.Repeat("word ", 15), " ")
	long        = prose + " " + prose
	heading     = "## PLAN ROUND 1 · goal · task t"
	blockStarts = []string{
		"1.", "2)", "**3.", "-", "+", "*", "#", "###", ">", "|", "```", "~~~", "---", "===", "***",
	}
	opensBlock = regexp.MustCompile(`^(?:[*_]*\d+[.)](?:\s|$)|[-*+](?:\s|$)|#{1,6}(?:\s|$)|>|\||` +
		"```|~~~" + `|[-=]+\s*$|[*_]{3,}\s*$)`)
)

type message struct {
	session, id string
	index       int
}

func isolatedTempDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("TMPDIR", dir)
	return dir
}

func assertNoState(t *testing.T, dir string) {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Errorf("state left in the temp dir: %v", entries)
	}
}

func runHook(t *testing.T, payload string) string {
	t.Helper()
	var stdout bytes.Buffer
	if status := wrapformats.Run(strings.NewReader(payload), &stdout); status != 0 {
		t.Fatalf("exit status %d, want 0", status)
	}
	return stdout.String()
}

func (m *message) batch(t *testing.T, final bool, delta string) (string, bool) {
	t.Helper()
	payload, _ := json.Marshal(map[string]any{
		"session_id": m.session, "hook_event_name": "MessageDisplay", "turn_id": "turn",
		"message_id": m.id, "index": m.index, "final": final, "delta": delta,
	})
	m.index++
	stdout := runHook(t, string(payload))
	if stdout == "" {
		return "", false
	}
	var output struct {
		HookSpecificOutput map[string]string `json:"hookSpecificOutput"`
	}
	if err := json.Unmarshal([]byte(stdout), &output); err != nil {
		t.Fatalf("output %q is not hook JSON: %v", stdout, err)
	}
	content, found := output.HookSpecificOutput["displayContent"]
	if output.HookSpecificOutput["hookEventName"] != "MessageDisplay" || !found {
		t.Fatalf("got %q, want MessageDisplay output with displayContent", stdout)
	}
	return content, true
}

func display(t *testing.T, reply string) (string, bool) {
	t.Helper()
	isolatedTempDir(t)
	return (&message{session: "session", id: "message"}).batch(t, true, reply)
}

func shown(t *testing.T, reply string) string {
	t.Helper()
	content, changed := display(t, reply)
	if !changed {
		t.Fatalf("reply %q displays untouched, want it rewrapped", reply)
	}
	return content
}

func passes(t *testing.T, reply string) bool {
	t.Helper()
	_, changed := display(t, reply)
	return !changed
}

func roundWith(lines ...string) string {
	return strings.Join(slices.Concat(
		[]string{heading, ""}, lines, []string{"", "**reply:** ok"}), "\n")
}

func TestNoContinuationLineOpensABlock(t *testing.T) {
	var items []string
	for _, word := range blockStarts {
		for shift := range 30 {
			for _, lead := range []string{"- **rec:** ", "> - "} {
				for _, tail := range []string{"", " after"} {
					items = append(items,
						fmt.Sprintf("%s%s %s %s%s", lead, prose, strings.Repeat("x", shift), word, tail))
				}
			}
		}
	}
	reply := roundWith(items...)
	lines := strings.Split(shown(t, reply), "\n")
	var continuations, unquoted []string
	for _, line := range lines {
		if strings.HasPrefix(line, "  ") || strings.HasPrefix(line, ">   ") {
			continuations = append(continuations, line)
		}
		unquoted = append(unquoted, strings.TrimPrefix(line, ">   "))
	}
	if !slices.Equal(pytext.Fields(strings.Join(unquoted, " ")), pytext.Fields(reply)) {
		t.Error("the display changes the words")
	}
	if len(continuations) == 0 {
		t.Fatal("the display has no continuation lines")
	}
	for _, line := range continuations {
		if utf8.RuneCountInString(line) > 100 {
			t.Errorf("continuation line over 100 columns: %q", line)
		}
		if opensBlock.MatchString(strings.TrimLeft(strings.TrimPrefix(line, ">   "), " \t")) {
			t.Errorf("continuation line opens a block: %q", line)
		}
	}
}

func TestContinuationLinesIndentUnderTheText(t *testing.T) {
	lines := strings.Split(shown(t, roundWith(
		"1. **Question?** "+long,
		"   - **rec:** "+long,
		"> - "+long,
	)), "\n")
	var starts []string
	for _, line := range lines[2:8] {
		starts = append(starts, line[:6])
	}
	want := []string{"1. **Q", "   wor", "   - *", "     w", "> - wo", ">   wo"}
	if !slices.Equal(starts, want) {
		t.Errorf("line starts %q, want %q", starts, want)
	}
	for _, line := range lines {
		if utf8.RuneCountInString(line) > 100 {
			t.Errorf("line over 100 columns: %q", line)
		}
	}
}

func TestReplyWithoutAMarathonFormatDisplaysUntouched(t *testing.T) {
	if !passes(t, "## Summary\n\n"+long+"\n") {
		t.Error("got output, want none")
	}
}

func TestEveryArchitectFormatHeadingIsRecognized(t *testing.T) {
	headings := []string{
		"PLAN ROUND 2 · goal · task t",
		"INTAKE ROUND 1 · experiment.ai → v1.ai",
		"ESCALATION · goal · task t",
		"Task brief · goal · task",
		"Status · 2026-10-01",
		"── STATUS · 2026-10-01 ──",
		"Retro · goal v1.storage · 2026-09-14..09-24",
		"RETRO · goal v1.storage",
	}
	for _, heading := range headings {
		for _, form := range []string{heading, "**" + heading + "**", "# " + heading, "## " + heading} {
			if passes(t, form+"\n\n"+long+"\n") {
				t.Errorf("%q: displays untouched, want it rewrapped", form)
			}
		}
	}
}

func TestSessionBriefHeadingIsRecognized(t *testing.T) {
	for _, form := range []string{"## goal · task 3/5 · task", "> ### **goal · task 3/5 · task**"} {
		if passes(t, form+"\n\n"+long+"\n") {
			t.Errorf("%q: displays untouched, want it rewrapped", form)
		}
	}
}

func TestHeadingsMatchUnicodeAsPythonDoes(t *testing.T) {
	for form, recognized := range map[string]bool{
		"PLAN ROUND ٣ · goal · task t": true,
		"## goal · task ٣/٥ · task":    true,
		"ESCALATION· goal":             true,
		"ESCALATIONé goal":             false,
		"ESCALATION_ goal":             false,
		"## \u3000goal · task 3/5 · t": false,
	} {
		if got := !passes(t, form+"\n\n"+long+"\n"); got != recognized {
			t.Errorf("%q: got recognized %v, want %v", form, got, recognized)
		}
	}
}

func TestFormatNamedMidLineIsNotAHeading(t *testing.T) {
	for _, line := range []string{
		"The next plan round asks about the audience.",
		"I'll print PLAN ROUND 1 once the facts are in.",
		"Next: ESCALATION · goal · task t",
	} {
		if !passes(t, line+"\n\n"+long+"\n") {
			t.Errorf("%q: got output, want none", line)
		}
	}
}

func TestWrappedReplyDisplaysUntouched(t *testing.T) {
	if !passes(t, roundWith(prose)) {
		t.Error("got output, want none")
	}
}

func TestLinesBeforeTheHeadingDisplayUntouched(t *testing.T) {
	lines := strings.Split(shown(t, long+"\n"+roundWith(long)), "\n")
	if lines[0] != long || lines[1] != heading || len(lines) != 7 {
		t.Errorf("got %q, want the first line kept and the round's long line rewrapped", lines)
	}
}

func TestCodeFencesAndTablesAreLeftLong(t *testing.T) {
	fenced := []string{"```", long, "```", "| " + prose + " | " + prose + " |"}
	if !passes(t, roundWith(fenced...)) {
		t.Error("got output, want none")
	}
	lines := strings.Split(shown(t, roundWith(append(fenced, long)...)), "\n")
	if !slices.Equal(lines[2:6], fenced) {
		t.Errorf("got %q, want %q", lines[2:6], fenced)
	}
}

func TestHeadingInsideACodeFenceIsNotAFormat(t *testing.T) {
	if !passes(t, "```\n"+heading+"\n```\n"+long+"\n") {
		t.Error("got output, want none")
	}
}

func TestUnwrappableWordStaysAsItIs(t *testing.T) {
	if !passes(t, roundWith("- see "+longWord)) {
		t.Error("got output, want none")
	}
}

func TestUnwrappableWordStaysWholeBesideRewrappedLines(t *testing.T) {
	if content := shown(t, roundWith("- see "+longWord, long)); !strings.Contains(
		content, "\n- see "+longWord+"\n") {
		t.Errorf("display %q splits the long word", content)
	}
}

func TestFormatSplitAcrossBatches(t *testing.T) {
	dir := isolatedTempDir(t)
	m := &message{session: "session", id: "message"}
	if content, changed := m.batch(t, false, long+"\n"+heading+"\n\n"); changed {
		t.Errorf("heading batch: got %q, want no output", content)
	}
	content, changed := m.batch(t, false, "- "+long+"\n"+long+"\n")
	want := "- " + strings.Repeat("word ", 18) + "word\n  " + strings.Repeat("word ", 10) + "word\n" +
		strings.Repeat("word ", 19) + "word\nword word word word word word word word word word\n"
	if !changed || content != want {
		t.Errorf("got %q (changed %v), want %q", content, changed, want)
	}
	if _, changed := m.batch(t, true, ""); changed {
		t.Error("empty final batch: got output, want none")
	}
	assertNoState(t, dir)
}

func TestCodeFenceAcrossBatches(t *testing.T) {
	dir := isolatedTempDir(t)
	m := &message{session: "session", id: "message"}
	m.batch(t, false, heading+"\n```\n")
	if content, changed := m.batch(t, false, long+"\n"); changed {
		t.Errorf("inside the fence: got %q, want no output", content)
	}
	if content, changed := m.batch(t, false, long+"\n```\n"); changed {
		t.Errorf("closing the fence: got %q, want no output", content)
	}
	if _, changed := m.batch(t, true, long); !changed {
		t.Error("after the fence: displays untouched, want it rewrapped")
	}
	assertNoState(t, dir)
}

func TestHeadingInAFenceOpenedInAnEarlierBatchIsNotAFormat(t *testing.T) {
	dir := isolatedTempDir(t)
	m := &message{session: "session", id: "message"}
	for _, delta := range []string{"```\n", heading + "\n" + long + "\n", "```\n" + long + "\n"} {
		if content, changed := m.batch(t, false, delta); changed {
			t.Errorf("%q: got %q, want no output", delta, content)
		}
	}
	m.batch(t, true, "")
	assertNoState(t, dir)
}

func TestReplyWithoutAFormatAcrossBatchesLeavesNoState(t *testing.T) {
	dir := isolatedTempDir(t)
	m := &message{session: "session", id: "message"}
	for _, delta := range []string{"## Summary\n", long + "\n", "| " + long + " |\n"} {
		if content, changed := m.batch(t, false, delta); changed {
			t.Errorf("%q: got %q, want no output", delta, content)
		}
		assertNoState(t, dir)
	}
	m.batch(t, true, long)
}

func TestFinalPartialLineIsRewrapped(t *testing.T) {
	for _, tail := range []string{"", "\n"} {
		dir := isolatedTempDir(t)
		m := &message{session: "session", id: "message"}
		m.batch(t, false, heading+"\n")
		content, changed := m.batch(t, true, long+tail)
		want := strings.Repeat("word ", 19) + "word\n" +
			strings.TrimSuffix(strings.Repeat("word ", 10), " ") + tail
		if !changed || content != want {
			t.Errorf("tail %q: got %q (changed %v), want %q", tail, content, changed, want)
		}
		assertNoState(t, dir)
	}
}

func TestInterleavedMessagesKeepTheirOwnState(t *testing.T) {
	dir := isolatedTempDir(t)
	round := &message{session: "session", id: "round"}
	other := &message{session: "session", id: "other"}
	elsewhere := &message{session: "elsewhere", id: "round"}
	round.batch(t, false, heading+"\n")
	if _, changed := other.batch(t, false, long+"\n"); changed {
		t.Error("other message: got output, want none")
	}
	if _, changed := elsewhere.batch(t, false, long+"\n"); changed {
		t.Error("same message_id in another session: got output, want none")
	}
	if _, changed := round.batch(t, false, long+"\n"); !changed {
		t.Error("round: displays untouched, want it rewrapped")
	}
	if _, changed := other.batch(t, true, long); changed {
		t.Error("other message's final batch: got output, want none")
	}
	if _, changed := round.batch(t, true, long); !changed {
		t.Error("round's final batch: displays untouched, want it rewrapped")
	}
	elsewhere.batch(t, true, "")
	assertNoState(t, dir)
}

func TestStateIsRemovedAfterTheFinalBatch(t *testing.T) {
	dir := isolatedTempDir(t)
	m := &message{session: "session/../x", id: "message\x00?"}
	m.batch(t, false, heading+"\n")
	entries, _ := os.ReadDir(dir)
	if len(entries) != 1 {
		t.Fatalf("got state %v, want one file in the temp dir", entries)
	}
	m.batch(t, true, "")
	assertNoState(t, dir)
	m.index = 0
	if _, changed := m.batch(t, true, long); changed {
		t.Error("after the final batch: got output, want none")
	}
}

func TestOutputIsJSONAsPythonWritesIt(t *testing.T) {
	isolatedTempDir(t)
	payload, _ := json.Marshal(map[string]any{
		"session_id": "s", "message_id": "m", "index": 0, "final": true,
		"delta": roundWith(`"q" \ é 😀 `+"\x01\x7f "+long, "| a\tb |\r"),
	})
	want := `{"hookSpecificOutput": {"hookEventName": "MessageDisplay", "displayContent": ` +
		`"## PLAN ROUND 1 \u00b7 goal \u00b7 task t\n\n\"q\" \\ \u00e9 \ud83d\ude00 \u0001\u007f ` +
		`word word word word word word word word word word word word word word word word word\n` +
		`word word word word word word word word word word word word word\n| a\tb |\r\n\n` +
		`**reply:** ok"}}` + "\n"
	if got := runHook(t, string(payload)); got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestUnreadableInputPrintsNothing(t *testing.T) {
	dir := isolatedTempDir(t)
	delta, _ := json.Marshal(heading + "\n" + long)
	for _, payload := range []string{
		"not json", "[]", "null", "{}", "\xff{}", "\ufeff{}", "{} {}",
		`{"message_id": "m", "final": true}`,
		fmt.Sprintf(`{"delta": %s, "final": true}`, delta),
		fmt.Sprintf(`{"message_id": 3, "delta": %s, "final": true}`, delta),
		`{"message_id": "m", "delta": 3, "final": true}`,
		fmt.Sprintf(`{"message_id": "m", "delta": %s, "final": "yes"}`, delta),
		fmt.Sprintf(`{"session_id": 3, "message_id": "m", "delta": %s, "final": true}`, delta),
	} {
		if stdout := runHook(t, payload); stdout != "" {
			t.Errorf("%q: got %q, want no output", payload, stdout)
		}
	}
	assertNoState(t, dir)
}
