package contextlint

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/BurntSushi/toml"

	"github.com/standards-lab/claude-plugins/marathon/internal/pytext"
)

const (
	contextReadmeBudget = 1000
	rootReadmeBudget    = 1500
	noteBudget          = 1500
	proseFix            = "trim; move detail to docs/"
	indexCommentLines   = 2
	criterionBudget     = 60
	entryFix            = "move detail to the goal's context note"
	capabilityMap       = "## Capability map"
	entryForm           = "open it as '- **Name** (built): text' or '- **Name** (planned): text'"
	linkFix             = "point the link at an existing file or remove it"
	space               = "[" + pytext.SpaceClass + "]"
	nonSpace            = "[^" + pytext.SpaceClass + "]"
)

var recordSectionBudgets = []struct {
	prefix string
	budget int
	fix    string
}{
	{"## Task brief", 1200, "shorten"},
	{"## Decisions", 2000, "prune what no longer binds later tasks"},
	{"## Pending edits", 1000, "apply or shorten"},
}

type fieldBudget struct {
	field  string
	budget int
}

var (
	indexArrays      = []string{"active", "planned", "backlog"}
	goalEntryBudgets = []fieldBudget{{"summary", 100}}
	taskEntryBudgets = []fieldBudget{{"summary", 100}, {"proof", 40}}
	tomlDelimiters   = []string{`"""`, `'''`, `"`, `'`}

	recordTask      = regexp.MustCompile(`^\p{Nd}+\. \[[ x]\] (` + nonSpace + `+)$`)
	capabilityEntry = regexp.MustCompile(`^- \*\*(.+?)\*\* \((built|planned)\): ` + nonSpace)
	capabilityName  = regexp.MustCompile(`^- \*\*(.+?)\*\*`)
	fence           = regexp.MustCompile("^" + space + "*(`{3,}|~{3,})")
	linkTitle       = `(?:` + space + `+(?:"[^"]*"|'[^']*'|\([^)]*\)))?`
	inlineLink      = regexp.MustCompile(`\]\(` + space + `*(<[^>]*>|[^` + pytext.SpaceClass + `)]*)` +
		linkTitle + space + `*\)`)
	linkDefinition = regexp.MustCompile(`^ {0,3}\[[^\]^][^\]]*\]:` + space + `*(<[^>]*>|` + nonSpace + `+)`)
	urlScheme      = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9+.-]*:`)
)

// Run checks the context of the repository at dir, printing one line per
// failure to stdout, and returns the exit status: 1 if anything fails or a
// file can't be read as UTF-8, with the reason on stderr and nothing on
// stdout, and 0 otherwise.
func Run(dir string, stdout, stderr io.Writer) int {
	c := &checker{root: dir, parsed: map[string]parsed{}}
	c.files = append(c.nonDotFiles("context", ""), c.nonDotFiles("roadmap", "")...)
	checks := []func(){
		c.checkKinds,
		c.checkBudgets,
		c.checkTOMLParses,
		c.checkIndex,
		c.checkGoalFileComments,
		c.checkEntryBudgets,
		c.checkRepos,
		c.checkOrder,
		c.checkRecordTasks,
		c.checkCapabilityMap,
		c.checkLinks,
	}
	for _, check := range checks {
		check()
	}
	if c.err != nil {
		_, _ = fmt.Fprintf(stderr, "lint-context: %v\n", c.err)
		return 1
	}
	var report strings.Builder
	for _, line := range c.failures {
		report.WriteString(line + "\n")
	}
	if _, err := io.WriteString(stdout, report.String()); err != nil {
		_, _ = fmt.Fprintf(stderr, "lint-context: %v\n", err)
		return 1
	}
	if len(c.failures) > 0 {
		return 1
	}
	return 0
}

type checker struct {
	root     string
	files    []string
	parsed   map[string]parsed
	failures []string
	err      error
}

type parsed struct {
	data map[string]any
	keys []toml.Key
	err  error
}

func (c *checker) fail(format string, args ...any) {
	c.failures = append(c.failures, fmt.Sprintf(format, args...))
}

func (c *checker) path(rel string) string {
	return filepath.Join(c.root, rel)
}

func (c *checker) read(rel string) string {
	content, err := os.ReadFile(c.path(rel))
	if err == nil && !utf8.Valid(content) {
		err = fmt.Errorf("%s: not UTF-8", rel)
	}
	if err != nil {
		if c.err == nil {
			c.err = err
		}
		return ""
	}
	return strings.ReplaceAll(strings.ReplaceAll(string(content), "\r\n", "\n"), "\r", "\n")
}

func (c *checker) isFile(rel string) bool {
	info, err := os.Stat(c.path(rel))
	return err == nil && info.Mode().IsRegular()
}

func (c *checker) nonDotFiles(top, suffix string) []string {
	if info, err := os.Stat(c.path(top)); err != nil || !info.IsDir() {
		return nil
	}
	var files []string
	var walk func(dir string)
	walk = func(dir string) {
		entries, _ := os.ReadDir(c.path(dir))
		for _, entry := range entries {
			rel := dir + "/" + entry.Name()
			switch {
			case strings.HasPrefix(entry.Name(), "."):
			case entry.IsDir():
				walk(rel)
			case strings.HasSuffix(entry.Name(), suffix) && c.isFile(rel):
				files = append(files, rel)
			}
		}
	}
	walk(top)
	slices.SortFunc(files, func(a, b string) int {
		return slices.Compare(strings.Split(a, "/"), strings.Split(b, "/"))
	})
	return files
}

type kind int

const (
	unknown kind = iota
	roadmapIndex
	goalFile
	contextReadme
	movedRoadmapIndex
	note
	record
	movedGoalFile
)

func fileKind(rel string) kind {
	parts := strings.Split(rel, "/")
	name := parts[len(parts)-1]
	if parts[0] == "roadmap" {
		switch {
		case rel == "roadmap/index.toml":
			return roadmapIndex
		case len(parts) == 2 && strings.HasSuffix(name, ".toml"):
			return goalFile
		}
		return unknown
	}
	switch {
	case rel == "context/README.md":
		return contextReadme
	case rel == "context/roadmap.toml":
		return movedRoadmapIndex
	case len(parts) == 2 && strings.HasSuffix(name, ".md"):
		return note
	case len(parts) == 3 && parts[1] == "goals" && strings.HasSuffix(name, ".md"):
		return record
	case len(parts) == 3 && parts[1] == "roadmap" && strings.HasSuffix(name, ".toml"):
		return movedGoalFile
	}
	return unknown
}

func words(text string) int {
	return len(pytext.Fields(text))
}

func thousands(n int) string {
	digits := strconv.Itoa(n)
	var b strings.Builder
	for i, d := range digits {
		if i > 0 && (len(digits)-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteRune(d)
	}
	return b.String()
}

func (c *checker) over(rel, what string, count, budget int, fix string) {
	c.fail("%s: %s is %s words, budget %s; %s", rel, what, thousands(count), thousands(budget), fix)
}

func (c *checker) checkKinds() {
	for _, rel := range c.files {
		switch fileKind(rel) {
		case movedRoadmapIndex:
			c.fail("%s: the roadmap index moved; move it to roadmap/index.toml", rel)
		case movedGoalFile:
			c.fail("%s: the roadmap moved; move it to roadmap/%s", rel, filepath.Base(rel))
		case unknown:
			if strings.HasPrefix(rel, "roadmap/") {
				c.fail("%s: not a roadmap file; name it index.toml or <goal>.toml, "+
					"or move it out of roadmap/", rel)
			} else {
				c.fail("%s: not a context file, which is README.md, a note (*.md), "+
					"or goals/*.md; move it to docs/ or rename it to one of those kinds", rel)
			}
		}
	}
}

type section struct {
	heading string
	body    string
}

func sections(text string) []section {
	var all []section
	var body []string
	for _, line := range pytext.Lines(text) {
		if strings.HasPrefix(line, "## ") {
			if len(all) > 0 {
				all[len(all)-1].body = strings.Join(body, "\n")
			}
			all = append(all, section{heading: line})
			body = nil
		} else if len(all) > 0 {
			body = append(body, line)
		}
	}
	if len(all) > 0 {
		all[len(all)-1].body = strings.Join(body, "\n")
	}
	return all
}

func (c *checker) checkBudgets() {
	if c.isFile("README.md") {
		if count := words(c.read("README.md")); count > rootReadmeBudget {
			c.over("README.md", "README", count, rootReadmeBudget, proseFix)
		}
	}
	for _, rel := range c.files {
		switch fileKind(rel) {
		case contextReadme:
			if count := words(c.read(rel)); count > contextReadmeBudget {
				c.over(rel, "README", count, contextReadmeBudget, proseFix)
			}
		case note:
			if count := words(c.read(rel)); count > noteBudget {
				c.over(rel, "note", count, noteBudget, proseFix)
			}
		case record:
			for _, s := range sections(c.read(rel)) {
				for _, b := range recordSectionBudgets {
					if s.heading != b.prefix && !strings.HasPrefix(s.heading, b.prefix+" ") {
						continue
					}
					if count := words(s.body); count > b.budget {
						name := pytext.TrimSpace(s.heading[3:])
						c.over(rel, "section '"+name+"'", count, b.budget, b.fix)
					}
				}
			}
		}
	}
}

func (c *checker) parse(rel string) parsed {
	if p, ok := c.parsed[rel]; ok {
		return p
	}
	var p parsed
	text := c.read(rel)
	if strings.HasPrefix(text, "\ufeff") {
		p.err = errors.New("toml: line 1: a byte order mark, which TOML doesn't allow")
	} else {
		var meta toml.MetaData
		meta, p.err = toml.Decode(text, &p.data)
		p.keys = meta.Keys()
	}
	c.parsed[rel] = p
	return p
}

func (c *checker) goalFiles() []string {
	var rels []string
	for _, rel := range c.files {
		if fileKind(rel) == goalFile && c.parse(rel).err == nil {
			rels = append(rels, rel)
		}
	}
	return rels
}

func (p parsed) keysUnder(path ...string) []string {
	var keys []string
	for _, key := range p.keys {
		if len(key) > len(path) && slices.Equal(key[:len(path)], path) &&
			!slices.Contains(keys, key[len(path)]) {
			keys = append(keys, key[len(path)])
		}
	}
	return keys
}

func (p parsed) taskSlugs() []string {
	tasks, ok := p.data["tasks"].(map[string]any)
	if !ok {
		return nil
	}
	var slugs []string
	for _, slug := range p.keysUnder("tasks") {
		if _, ok := tasks[slug].(map[string]any); ok {
			slugs = append(slugs, slug)
		}
	}
	return slugs
}

func list(value any) ([]any, bool) {
	switch v := value.(type) {
	case []any:
		return v, true
	case []map[string]any:
		items := make([]any, len(v))
		for i, item := range v {
			items[i] = item
		}
		return items, true
	}
	return nil, false
}

func (c *checker) checkTOMLParses() {
	for _, rel := range c.files {
		if k := fileKind(rel); k != roadmapIndex && k != goalFile {
			continue
		}
		if err := c.parse(rel).err; err != nil {
			c.fail("%s: does not parse: %s; fix the TOML at that position", rel, err)
		}
	}
}

func (c *checker) checkIndex() {
	for _, rel := range c.files {
		if fileKind(rel) != roadmapIndex {
			continue
		}
		comments := 0
		for _, line := range pytext.Lines(c.read(rel)) {
			if strings.HasPrefix(strings.TrimLeftFunc(line, pytext.IsSpace), "#") {
				comments++
			}
		}
		if comments > indexCommentLines {
			c.fail("%s: %d comment lines, at most %d; keep one comment that points to the format",
				rel, comments, indexCommentLines)
		}
		p := c.parse(rel)
		if p.err != nil {
			continue
		}
		for _, key := range p.keysUnder() {
			if !slices.Contains(indexArrays, key) {
				c.fail("%s: '%s' is not one of active, planned, backlog; "+
					"move it into a goal file under roadmap/ or remove it", rel, key)
			} else if _, ok := list(p.data[key]); !ok {
				c.fail("%s: '%s' is not an array; write it as an array of goal slugs, [] for none",
					rel, key)
			}
		}
	}
}

func tomlCommentLineNumbers(text string) []int {
	var numbers []int
	number, delimiter := 1, ""
	for i := 0; i < len(text); {
		char := text[i]
		if char == '\n' {
			number++
		}
		switch {
		case delimiter == "":
			if char == '#' {
				numbers = append(numbers, number)
				if end := strings.IndexByte(text[i:], '\n'); end == -1 {
					i = len(text)
				} else {
					i += end
				}
				continue
			}
			if d := opening(text[i:]); d != "" {
				delimiter = d
				i += len(d)
				continue
			}
		case char == '\\' && delimiter[0] == '"':
			if strings.HasPrefix(text[i+1:], "\n") {
				number++
			}
			i += 2
			continue
		case strings.HasPrefix(text[i:], delimiter):
			i += len(delimiter)
			delimiter = ""
			continue
		case char == '\n' && len(delimiter) == 1:
			delimiter = ""
		}
		i++
	}
	return numbers
}

func opening(rest string) string {
	for _, d := range tomlDelimiters {
		if strings.HasPrefix(rest, d) {
			return d
		}
	}
	return ""
}

func (c *checker) checkGoalFileComments() {
	for _, rel := range c.files {
		if fileKind(rel) != goalFile {
			continue
		}
		for _, number := range tomlCommentLineNumbers(c.read(rel)) {
			c.fail("%s:%d: comment; move the note into the goal's context note", rel, number)
		}
	}
}

func (c *checker) entryBudgets(rel, entry string, table map[string]any, budgets []fieldBudget) {
	for _, b := range budgets {
		if value, ok := table[b.field].(string); ok {
			if count := words(value); count > b.budget {
				c.over(rel, entry+" "+b.field, count, b.budget, entryFix)
			}
		}
	}
}

func (c *checker) checkEntryBudgets() {
	for _, rel := range c.goalFiles() {
		p := c.parse(rel)
		c.entryBudgets(rel, "goal", p.data, goalEntryBudgets)
		var labeled [][2]string
		if criteria, ok := p.data["criteria"].(string); ok {
			labeled = append(labeled, [2]string{"criteria", criteria})
		} else if criteria, ok := list(p.data["criteria"]); ok {
			for i, statement := range criteria {
				if s, ok := statement.(string); ok {
					labeled = append(labeled, [2]string{fmt.Sprintf("criteria[%d]", i), s})
				}
			}
		}
		for _, l := range labeled {
			if count := words(l[1]); count > criterionBudget {
				c.over(rel, "goal "+l[0], count, criterionBudget, entryFix)
			}
		}
		tasks, _ := p.data["tasks"].(map[string]any)
		for _, slug := range p.taskSlugs() {
			c.entryBudgets(rel, "tasks."+slug, tasks[slug].(map[string]any), taskEntryBudgets)
		}
	}
}

func (c *checker) checkRepos() {
	for _, rel := range c.goalFiles() {
		if _, ok := c.parse(rel).data["repos"]; !ok {
			c.fail("%s: repos is missing; every goal file writes it, [] for none", rel)
		}
	}
}

func slugList(value any) ([]string, bool) {
	items, ok := list(value)
	if !ok {
		return nil, false
	}
	slugs := make([]string, len(items))
	for i, item := range items {
		if slugs[i], ok = item.(string); !ok {
			return nil, false
		}
	}
	return slugs, true
}

func (c *checker) checkOrder() {
	for _, rel := range c.goalFiles() {
		p := c.parse(rel)
		slugs := p.taskSlugs()
		value, present := p.data["order"]
		if !present {
			if len(slugs) > 0 {
				c.fail("%s: order is missing; list each task slug once, in run order", rel)
			}
			continue
		}
		order, ok := slugList(value)
		if !ok {
			c.fail("%s: order is not an array of task slugs; "+
				"write it as an array of the task slugs, in run order", rel)
			continue
		}
		var seen []string
		for _, slug := range order {
			if slices.Contains(seen, slug) {
				c.fail("%s: order lists '%s' more than once; list it once", rel, slug)
			} else if !slices.Contains(slugs, slug) {
				c.fail("%s: order lists '%s', which has no [tasks.%s] table; "+
					"add a [tasks.%s] table or drop it from order", rel, slug, slug, slug)
			}
			seen = append(seen, slug)
		}
		for _, slug := range slugs {
			if !slices.Contains(seen, slug) {
				c.fail("%s: order lacks task '%s'; add it to order or drop its [tasks.%s] table",
					rel, slug, slug)
			}
		}
	}
}

func recordTaskSlugs(text string) []string {
	slugs := []string{}
	for _, s := range sections(text) {
		if s.heading != "## Tasks" {
			continue
		}
		for _, line := range pytext.Lines(s.body) {
			if match := recordTask.FindStringSubmatch(pytext.TrimSpace(line)); match != nil {
				slugs = append(slugs, match[1])
			}
		}
		break
	}
	return slugs
}

func (c *checker) checkRecordTasks() {
	for _, rel := range c.goalFiles() {
		name := filepath.Base(rel)
		recordRel := "context/goals/" + strings.TrimSuffix(name, ".toml") + ".md"
		if !slices.Contains(c.files, recordRel) {
			continue
		}
		value, present := c.parse(rel).data["order"]
		order, ok := []string{}, true
		if present {
			order, ok = slugList(value)
		}
		if !ok {
			continue
		}
		listed := recordTaskSlugs(c.read(recordRel))
		if !slices.Equal(listed, order) {
			c.fail("%s: ## Tasks lists [%s], but %s order is [%s]; make the list mirror order",
				recordRel, strings.Join(listed, ", "), rel, strings.Join(order, ", "))
		}
	}
}

type numberedLine struct {
	number int
	text   string
}

func linesOutsideCodeWithoutSpans(text string) []numberedLine {
	var lines []numberedLine
	open := ""
	for i, line := range pytext.Lines(text) {
		match := fence.FindStringSubmatch(line)
		switch {
		case open == "" && match != nil:
			open = match[1]
		case open != "":
			if match != nil && strings.HasPrefix(match[1], open) {
				open = ""
			}
		default:
			lines = append(lines, numberedLine{i + 1, withoutCodeSpans(line)})
		}
	}
	return lines
}

func withoutCodeSpans(line string) string {
	var b strings.Builder
	for i := 0; i < len(line); {
		if line[i] != '`' {
			b.WriteByte(line[i])
			i++
			continue
		}
		if end := codeSpanEnd(line, i); end > 0 {
			i = end
			continue
		}
		b.WriteByte(line[i])
		i++
	}
	return b.String()
}

func codeSpanEnd(line string, start int) int {
	run := start
	for run < len(line) && line[run] == '`' {
		run++
	}
	n := run - start
	for k := run + 1; k+n <= len(line); k++ {
		if line[k-1] != '`' && strings.Count(line[k:k+n], "`") == n &&
			(k+n == len(line) || line[k+n] != '`') {
			return k + n
		}
	}
	return 0
}

func numberedLinkTargets(text string) []numberedLine {
	var targets []numberedLine
	for _, line := range linesOutsideCodeWithoutSpans(text) {
		if definition := linkDefinition.FindStringSubmatch(line.text); definition != nil {
			targets = append(targets, numberedLine{line.number, definition[1]})
		}
		for _, match := range inlineLink.FindAllStringSubmatch(line.text, -1) {
			targets = append(targets, numberedLine{line.number, match[1]})
		}
	}
	return targets
}

func repoPath(target string) (string, bool) {
	if strings.HasPrefix(target, "<") && strings.HasSuffix(target, ">") {
		target = target[1 : len(target)-1]
	}
	if urlScheme.MatchString(target) || strings.HasPrefix(target, "#") {
		return "", false
	}
	if i := strings.IndexAny(target, "#?"); i >= 0 {
		target = target[:i]
	}
	return unquote(target), true
}

func (c *checker) resolves(rel, target string) bool {
	base := filepath.Dir(c.path(rel))
	if strings.HasPrefix(target, "/") {
		base = c.root
	}
	var parts []string
	for part := range strings.SplitSeq(target, "/") {
		if part != "" && part != "." {
			parts = append(parts, part)
		}
	}
	_, err := os.Stat(strings.Join(append([]string{base}, parts...), "/"))
	return err == nil
}

func (c *checker) markdown() []string {
	var rels []string
	if c.isFile("README.md") {
		rels = append(rels, "README.md")
	}
	rels = append(rels, c.nonDotFiles("context", ".md")...)
	return append(rels, c.nonDotFiles("docs", ".md")...)
}

func (c *checker) checkLinks() {
	for _, rel := range c.markdown() {
		for _, target := range numberedLinkTargets(c.read(rel)) {
			if p, ok := repoPath(target.text); ok && !c.resolves(rel, p) {
				c.fail("%s:%d: link %s does not resolve; "+linkFix, rel, target.number, target.text)
			}
		}
	}
}

type entry struct {
	number int
	line   string
	lines  []int
}

func capabilitySectionEntries(lines []string, start int) []entry {
	var entries []entry
	var current *entry
	for number := start; number <= len(lines); number++ {
		line := lines[number-1]
		if strings.HasPrefix(line, "## ") {
			break
		}
		blank := pytext.TrimSpace(line) == ""
		first, _ := utf8.DecodeRuneInString(line)
		unindentedProse := !blank && !pytext.IsSpace(first)
		switch {
		case strings.HasPrefix(line, "- ") || unindentedProse:
			if current != nil {
				entries = append(entries, *current)
			}
			current = &entry{number, line, []int{number}}
		case current != nil:
			current.lines = append(current.lines, number)
		case !blank:
			entries = append(entries, entry{number, line, []int{number}})
		}
	}
	if current != nil {
		entries = append(entries, *current)
	}
	return entries
}

func (c *checker) checkCapabilityMap() {
	rel := "context/README.md"
	if !c.isFile(rel) {
		return
	}
	text := c.read(rel)
	lines := pytext.Lines(text)
	heading := slices.Index(lines, capabilityMap)
	if heading < 0 {
		c.fail("%s: no '%s' heading; add one listing each feature as built or planned",
			rel, capabilityMap)
		return
	}
	linked := map[int]bool{}
	for _, target := range numberedLinkTargets(text) {
		if _, ok := repoPath(target.text); ok {
			linked[target.number] = true
		}
	}
	for _, e := range capabilitySectionEntries(lines, heading+2) {
		if !strings.HasPrefix(e.line, "- ") {
			c.fail("%s:%d: capability map line is not an entry; %s", rel, e.number, entryForm)
			continue
		}
		match := capabilityEntry.FindStringSubmatch(e.line)
		if match == nil {
			name := "without a bold name"
			if named := capabilityName.FindStringSubmatch(e.line); named != nil {
				name = "'" + named[1] + "'"
			}
			c.fail("%s:%d: capability map entry %s: %s", rel, e.number, name, entryForm)
			continue
		}
		if match[2] == "built" && !slices.ContainsFunc(e.lines, func(n int) bool { return linked[n] }) {
			c.fail("%s:%d: capability map entry '%s' is built but links nothing; "+
				"link its home with a relative link", rel, e.number, match[1])
		}
	}
}

func unquote(s string) string {
	var decoded []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '%' && i+2 < len(s) && isHex(s[i+1]) && isHex(s[i+2]) {
			high, low := unhex(s[i+1]), unhex(s[i+2])
			decoded = append(decoded, high<<4|low)
			i += 2
			continue
		}
		decoded = append(decoded, s[i])
	}
	return replacingInvalidUTF8(decoded)
}

func isHex(c byte) bool {
	return '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F'
}

func unhex(c byte) byte {
	switch {
	case c <= '9':
		return c - '0'
	case c <= 'F':
		return c - 'A' + 10
	}
	return c - 'a' + 10
}

func replacingInvalidUTF8(b []byte) string {
	var s strings.Builder
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		if r == utf8.RuneError && size == 1 {
			size = maximalSubpart(b)
			s.WriteRune(utf8.RuneError)
		} else {
			s.Write(b[:size])
		}
		b = b[size:]
	}
	return s.String()
}

func maximalSubpart(b []byte) int {
	need, low, high := 0, byte(0x80), byte(0xbf)
	switch c := b[0]; {
	case 0xc2 <= c && c <= 0xdf:
		need = 1
	case c == 0xe0:
		need, low = 2, 0xa0
	case c == 0xed:
		need, high = 2, 0x9f
	case 0xe1 <= c && c <= 0xef:
		need = 2
	case c == 0xf0:
		need, low = 3, 0x90
	case c == 0xf4:
		need, high = 3, 0x8f
	case 0xf1 <= c && c <= 0xf3:
		need = 3
	}
	n := 1
	for n <= need && n < len(b) && low <= b[n] && b[n] <= high {
		n, low, high = n+1, 0x80, 0xbf
	}
	return n
}
