package marathon

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"

	"github.com/BurntSushi/toml"

	"github.com/standards-lab/claude-plugins/marathon/internal/pytext"
)

const (
	configuration       = ".claude/marathon.toml"
	usageFailure        = 2
	signaledStatusBase  = 128
	commandFailure      = 1
	positionalArgument0 = "marathon"
)

var keys = []string{"check", "currency", "gate", "publish", "merge", "ci"}

var tableOfKey = map[string]string{
	"check":    "project",
	"currency": "project",
	"gate":     "project",
	"publish":  "remote",
	"merge":    "remote",
	"ci":       "remote",
}

var shellUnsafe = regexp.MustCompile(`[^\w@%+=:,./-]`)

type configurationError string

func (e configurationError) Error() string {
	return string(e)
}

// Run runs the command that args[0] names in the configuration of the git
// repository holding dir, passing it args[1:], and returns marathon's exit
// status.
func Run(dir string, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		_, _ = fmt.Fprintf(stderr, "usage: marathon <key> [args]; keys: %s\n", strings.Join(keys, ", "))
		return usageFailure
	}
	root, command, err := declared(dir, args[0])
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "marathon: %v\n", err)
		if _, ok := errors.AsType[configurationError](err); ok {
			return usageFailure
		}
		return commandFailure
	}
	return run(root, bashInvocation(command, args[1:]), stdout, stderr)
}

func declared(dir, key string) (root, command string, err error) {
	if dir, err = absolute(dir); err != nil {
		return "", "", err
	}
	if root, err = repositoryRoot(dir); err != nil {
		return "", "", err
	}
	command, err = declaredCommand(root, key)
	return root, command, err
}

func absolute(dir string) (string, error) {
	if filepath.IsAbs(dir) {
		return dir, nil
	}
	wd, err := syscall.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, dir), nil
}

func repositoryRoot(dir string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = dir
	cmd.Env = os.Environ()
	out, err := cmd.Output()
	if _, exited := errors.AsType[*exec.ExitError](err); exited {
		return "", configurationError(fmt.Sprintf(
			"%s is not in a git repository; run marathon inside the repository whose %s "+
				"declares the command", dir, configuration))
	}
	if err != nil {
		return "", err
	}
	return filepath.Clean(pytext.TrimSpace(string(out))), nil
}

func declaredCommand(root, key string) (string, error) {
	table, known := tableOfKey[key]
	if !known {
		return "", configurationError(fmt.Sprintf("unknown key '%s'; keys: %s", key,
			strings.Join(keys, ", ")))
	}
	path := filepath.Join(root, configuration)
	if info, err := os.Stat(path); err != nil || !info.Mode().IsRegular() {
		return "", configurationError(fmt.Sprintf(
			"no %s in %s; write one (marathon's init does), then set [%s] %s in it",
			configuration, root, table, key))
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	text := pytext.UniversalNewlines(string(content))
	var parsed map[string]any
	if _, err := toml.Decode(text, &parsed); err != nil {
		return "", configurationError(fmt.Sprintf("%s does not parse: %v; fix it", path, err))
	}
	unset := configurationError(fmt.Sprintf("set [%s] %s in %s", table, key, configuration))
	entries, _ := parsed[table].(map[string]any)
	value, isString := entries[key].(string)
	if !isString || pytext.TrimSpace(value) == "" {
		return "", unset
	}
	return value, nil
}

func bashInvocation(command string, args []string) []string {
	declared := strings.TrimRight(command, "\n")
	if strings.Contains(declared, "\n") {
		return append([]string{"bash", "-c", command, positionalArgument0}, args...)
	}
	words := []string{declared}
	for _, arg := range args {
		words = append(words, shellQuote(arg))
	}
	return []string{"bash", "-c", strings.Join(words, " ")}
}

func shellQuote(s string) string {
	if s == "" {
		return "''"
	}
	if !shellUnsafe.MatchString(s) {
		return s
	}
	return "'" + strings.ReplaceAll(s, "'", `'"'"'`) + "'"
}

func run(root string, argv []string, stdout, stderr io.Writer) int {
	cmd := exec.Command(argv[0], argv[1:]...)
	cmd.Dir = root
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	if exit, ok := errors.AsType[*exec.ExitError](err); ok {
		if status, ok := exit.Sys().(syscall.WaitStatus); ok && status.Signaled() {
			return signaledStatusBase + int(status.Signal())
		}
		return exit.ExitCode()
	}
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "marathon: %v\n", err)
		return commandFailure
	}
	return 0
}
