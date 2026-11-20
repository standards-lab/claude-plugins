package wrapformats_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/standards-lab/claude-plugins/marathon/internal/pytext"
	"github.com/standards-lab/claude-plugins/marathon/internal/wrapformats"
)

const storedRounds = "testdata/plan-rounds"

var (
	longWord    = "https://example.com/" + strings.Repeat("a", 100)
	prose       = strings.TrimSuffix(strings.Repeat("word ", 15), " ")
	blockStarts = []string{
		"1.", "2)", "**3.", "-", "+", "*", "#", "###", ">", "|", "```", "~~~", "---", "===", "***",
	}
	opensBlock = regexp.MustCompile(`^(?:[*_]*\d+[.)](?:\s|$)|[-*+](?:\s|$)|#{1,6}(?:\s|$)|>|\||` +
		"```|~~~" + `|[-=]+\s*$|[*_]{3,}\s*$)`)
	numberedItem = regexp.MustCompile(`^[ \t>*_]*\d+[.)][ \t]`)
	numberedText = regexp.MustCompile(`^[ \t>*_]*\d+[.)][ \t]+\S`)
	recLine      = regexp.MustCompile(`^[ \t>*_-]*rec:`)
	changesLine  = regexp.MustCompile(`^[ \t>*_-]*changes:`)
)

func runHook(t *testing.T, payload string) string {
	t.Helper()
	var stdout bytes.Buffer
	if status := wrapformats.Run(strings.NewReader(payload), &stdout); status != 0 {
		t.Fatalf("exit status %d, want 0", status)
	}
	return stdout.String()
}

func hookOutput(t *testing.T, payload string) map[string]string {
	t.Helper()
	stdout := runHook(t, payload)
	if stdout == "" {
		return nil
	}
	var output map[string]string
	if err := json.Unmarshal([]byte(stdout), &output); err != nil {
		t.Fatalf("output %q is not a JSON object of strings: %v", stdout, err)
	}
	return output
}

func stopPayload(message string, active bool) string {
	payload, _ := json.Marshal(map[string]any{
		"stop_hook_active": active, "last_assistant_message": message,
	})
	return string(payload)
}

func stop(t *testing.T, message string) map[string]string {
	t.Helper()
	return hookOutput(t, stopPayload(message, false))
}

func reprint(t *testing.T, output map[string]string) string {
	t.Helper()
	_, printed, found := strings.Cut(output["reason"], "\n\n")
	if output["decision"] != "block" || !found {
		t.Fatalf("got %q, want a block with the reprint", output)
	}
	return printed
}

func roundWith(lines ...string) string {
	return strings.Join(slices.Concat(
		[]string{"## PLAN ROUND 1 · goal · task t", ""}, lines, []string{"", "**reply:** ok"}), "\n")
}

func storedRoundReplies(t *testing.T) map[string]string {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join(storedRounds, "*.txt"))
	if err != nil || len(paths) == 0 {
		t.Fatalf("no stored rounds in %s: %v", storedRounds, err)
	}
	replies := map[string]string{}
	for _, path := range paths {
		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		replies[filepath.Base(path)] = string(content)
	}
	return replies
}

func printedReply(t *testing.T, reply string) string {
	t.Helper()
	output := stop(t, reply)
	if output == nil {
		return reply
	}
	return reprint(t, output)
}

func lineWidthGraderMatches(reply string) bool {
	lines := strings.Split(reply, "\n")
	for i, line := range lines {
		if !strings.Contains(line, "PLAN ROUND 1") {
			continue
		}
		for _, later := range lines[i:] {
			if utf8.RuneCountInString(later) > 100 {
				return true
			}
			if strings.Contains(later, "reply:") {
				break
			}
		}
	}
	return false
}

func numberedQuestionsGraderMatches(reply string) bool {
	lines := strings.Split(reply, "\n")
	last := len(lines) - 1
	for i := range last {
		if !numberedText.MatchString(lines[i]) {
			continue
		}
		for j := i + 1; j < last && !numberedItem.MatchString(lines[j]); j++ {
			if !recLine.MatchString(lines[j]) {
				continue
			}
			for k := j + 1; k <= last && !numberedItem.MatchString(lines[k]); k++ {
				if changesLine.MatchString(lines[k]) {
					return true
				}
			}
		}
	}
	return false
}

func TestStoredRoundsPrintWithinTheLineWidthGrader(t *testing.T) {
	failing := 0
	for name, reply := range storedRoundReplies(t) {
		if lineWidthGraderMatches(reply) {
			failing++
		}
		printed := printedReply(t, reply)
		if lineWidthGraderMatches(printed) {
			t.Errorf("%s: the printed reply fails the line-width grader", name)
		}
		if !slices.Equal(pytext.Fields(printed), pytext.Fields(reply)) {
			t.Errorf("%s: the printed reply changes the words", name)
		}
	}
	if failing == 0 {
		t.Error("no stored round fails the line-width grader unwrapped")
	}
}

func TestStoredRoundsKeepTheirNumberedQuestions(t *testing.T) {
	kept := 0
	for name, reply := range storedRoundReplies(t) {
		if !numberedQuestionsGraderMatches(reply) {
			continue
		}
		kept++
		if !numberedQuestionsGraderMatches(printedReply(t, reply)) {
			t.Errorf("%s: the printed reply loses its numbered questions", name)
		}
	}
	if kept == 0 {
		t.Error("no stored round has numbered questions")
	}
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
	lines := strings.Split(reprint(t, stop(t, reply)), "\n")
	var continuations, unquoted []string
	for _, line := range lines {
		if strings.HasPrefix(line, "  ") || strings.HasPrefix(line, ">   ") {
			continuations = append(continuations, line)
		}
		unquoted = append(unquoted, strings.TrimPrefix(line, ">   "))
	}
	if !slices.Equal(pytext.Fields(strings.Join(unquoted, " ")), pytext.Fields(reply)) {
		t.Error("the reprint changes the words")
	}
	if len(continuations) == 0 {
		t.Fatal("the reprint has no continuation lines")
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
	lines := strings.Split(reprint(t, stop(t, roundWith(
		"1. **Question?** "+prose+" "+prose,
		"   - **rec:** "+prose+" "+prose,
		"> - "+prose+" "+prose,
	))), "\n")
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

func TestReplyWithoutAMarathonFormatPasses(t *testing.T) {
	if output := stop(t, "## Summary\n\n"+prose+" "+prose+"\n"); output != nil {
		t.Errorf("got %q, want no output", output)
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
			if output := stop(t, form+"\n\n"+prose+" "+prose+"\n"); output["decision"] != "block" {
				t.Errorf("%q: got %q, want a block", form, output)
			}
		}
	}
}

func TestSessionBriefHeadingIsRecognized(t *testing.T) {
	for _, form := range []string{"## goal · task 3/5 · task", "> ### **goal · task 3/5 · task**"} {
		if output := stop(t, form+"\n\n"+prose+" "+prose+"\n"); output["decision"] != "block" {
			t.Errorf("%q: got %q, want a block", form, output)
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
		output := stop(t, form+"\n\n"+prose+" "+prose+"\n")
		if got := output["decision"] == "block"; got != recognized {
			t.Errorf("%q: got %q, want recognized %v", form, output, recognized)
		}
	}
}

func TestFormatNamedMidLineIsNotAHeading(t *testing.T) {
	for _, line := range []string{
		"The next plan round asks about the audience.",
		"I'll print PLAN ROUND 1 once the facts are in.",
		"Next: ESCALATION · goal · task t",
	} {
		if output := stop(t, line+"\n\n"+prose+" "+prose+"\n"); output != nil {
			t.Errorf("%q: got %q, want no output", line, output)
		}
	}
}

func TestWrappedReplyPasses(t *testing.T) {
	if output := stop(t, roundWith(prose)); output != nil {
		t.Errorf("got %q, want no output", output)
	}
}

func TestReprintIsNotBlockedAgain(t *testing.T) {
	if output := hookOutput(t, stopPayload(roundWith(prose+" "+prose), true)); output != nil {
		t.Errorf("got %q, want no output", output)
	}
}

func TestStopHookActiveCountsAsPythonTruth(t *testing.T) {
	reply, _ := json.Marshal(roundWith(prose + " " + prose))
	for value, active := range map[string]bool{
		"0": false, "0.0": false, "1e-999": false, `""`: false, "[]": false, "{}": false,
		"null": false, "1": true, "1e999": true, `"no"`: true, "[0]": true, `{"a": 0}`: true,
	} {
		payload := fmt.Sprintf(`{"stop_hook_active": %s, "last_assistant_message": %s}`, value, reply)
		if output := hookOutput(t, payload); (output == nil) != active {
			t.Errorf("stop_hook_active %s: got %q, want active %v", value, output, active)
		}
	}
}

func TestCodeFencesAndTablesAreLeftLong(t *testing.T) {
	fenced := []string{"```", prose + " " + prose, "```", "| " + prose + " | " + prose + " |"}
	if output := stop(t, roundWith(fenced...)); output != nil {
		t.Errorf("got %q, want no output", output)
	}
	lines := strings.Split(reprint(t, stop(t, roundWith(append(fenced, prose+" "+prose)...))), "\n")
	if !slices.Equal(lines[2:6], fenced) {
		t.Errorf("got %q, want %q", lines[2:6], fenced)
	}
}

func TestHeadingInsideACodeFenceIsNotAFormat(t *testing.T) {
	if output := stop(t, "```\n## PLAN ROUND 1 · g · t\n```\n"+prose+" "+prose+"\n"); output != nil {
		t.Errorf("got %q, want no output", output)
	}
}

func TestUnwrappableWordIsReportedNotBlocked(t *testing.T) {
	output := stop(t, roundWith("- see "+longWord))
	if _, blocked := output["decision"]; blocked || !strings.Contains(output["systemMessage"], "1 line(s)") {
		t.Errorf("got %q, want a report of 1 line", output)
	}
}

func TestUnwrappableWordStaysWholeInAReprint(t *testing.T) {
	printed := reprint(t, stop(t, roundWith("- see "+longWord, prose+" "+prose)))
	if !strings.Contains(printed, "\n- see "+longWord+"\n") {
		t.Errorf("reprint %q splits the long word", printed)
	}
}

func TestOutputIsJSONAsPythonWritesIt(t *testing.T) {
	block := stopPayload(roundWith(`"q" \ é 😀 `+"\x01\x7f "+prose+" "+prose, "| a\tb |\r"), false)
	want := `{"decision": "block", "reason": "marathon prints its formats wrapped at 100 columns, ` +
		`and this reply has longer lines. Print the reply again exactly as below, rewrapped at ` +
		`100 columns, with nothing before or after it:\n\n## PLAN ROUND 1 \u00b7 goal \u00b7 ` +
		`task t\n\n\"q\" \\ \u00e9 \ud83d\ude00 \u0001\u007f word word word word word word word ` +
		`word word word word word word word word word word\nword word word word word word word ` +
		`word word word word word word\n| a\tb |\r\n\n**reply:** ok"}` + "\n"
	if got := runHook(t, block); got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
	report := stopPayload(roundWith("- see https://example.com/"+strings.Repeat("é", 100)), false)
	want = `{"systemMessage": "marathon: 1 line(s) of this reply run past 100 characters on a ` +
		`single word, so they stay unwrapped."}` + "\n"
	if got := runHook(t, report); got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestUnreadableInputPasses(t *testing.T) {
	for _, payload := range []string{
		"not json", "[]", `{"stop_hook_active": false}`, "\xff{}", "\ufeff{}", "{} {}",
		`{"last_assistant_message": 3}`,
	} {
		if stdout := runHook(t, payload); stdout != "" {
			t.Errorf("%q: got %q, want no output", payload, stdout)
		}
	}
}
