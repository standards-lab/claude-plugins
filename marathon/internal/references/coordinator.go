package references

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"

	"github.com/standards-lab/claude-plugins/marathon/internal/pytext"
)

const (
	configuration  = ".claude/marathon.toml"
	coordinatorKey = "[project] coordinator"
)

type parseError string

func (e parseError) Error() string {
	return string(e)
}

type coordinator struct {
	dir     string
	found   bool
	members map[string]string
}

func findCoordinator(root string) (coordinator, error) {
	top, err := gitOutput(root, "rev-parse", "--show-toplevel")
	if err != nil {
		return coordinator{}, err
	}
	top = filepath.Clean(pytext.TrimSpace(top))
	own := filepath.Join(top, configuration)
	parsed, err := readConfiguration(own)
	if errors.Is(err, fs.ErrNotExist) {
		return coordinator{}, nil
	}
	if err != nil {
		return coordinator{}, err
	}
	if workspace, ok := parsed["workspace"].(map[string]any); ok {
		return coordinatorAt(top, own, workspace)
	}
	project, _ := parsed["project"].(map[string]any)
	value, set := project["coordinator"]
	if !set {
		return coordinator{}, nil
	}
	name, isString := value.(string)
	if !isString || pytext.TrimSpace(name) == "" {
		return coordinator{}, fmt.Errorf(
			"%s in %s is not a path; set it to the coordinator's directory, such as \"~/architecture\"",
			coordinatorKey, own)
	}
	dir, err := expandPath(top, name)
	if err != nil {
		return coordinator{}, err
	}
	if _, err := os.Stat(dir); errors.Is(err, fs.ErrNotExist) {
		return coordinator{dir: dir}, nil
	} else if err != nil {
		return coordinator{}, err
	}
	theirs := filepath.Join(dir, configuration)
	parsed, err = readConfiguration(theirs)
	if _, unparsed := errors.AsType[parseError](err); unparsed {
		return coordinator{}, err
	}
	if err != nil {
		return coordinator{}, fmt.Errorf(
			"%s in %s names %s, which has no readable %s: %v; "+
				"point the key at the coordinator's directory, or give the coordinator a %s with a [workspace] table",
			coordinatorKey, own, dir, configuration, err, configuration)
	}
	workspace, ok := parsed["workspace"].(map[string]any)
	if !ok {
		return coordinator{}, fmt.Errorf(
			"%s in %s names %s, whose %s has no [workspace] table; "+
				"point the key at the coordinator's directory, or add a [workspace] table to %s",
			coordinatorKey, own, dir, configuration, theirs)
	}
	return coordinatorAt(dir, theirs, workspace)
}

func coordinatorAt(dir, file string, workspace map[string]any) (coordinator, error) {
	c := coordinator{dir: dir, found: true, members: map[string]string{}}
	paths, _ := workspace["paths"].(map[string]any)
	for name, value := range paths {
		location, isString := value.(string)
		if !isString {
			return coordinator{}, fmt.Errorf(
				"[workspace.paths] %s in %s is not a path; set it to the directory's location", name, file)
		}
		expanded, err := expandPath(dir, location)
		if err != nil {
			return coordinator{}, err
		}
		c.members[name] = expanded
	}
	return c, nil
}

func readConfiguration(name string) (map[string]any, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	var parsed map[string]any
	if _, err := toml.Decode(pytext.UniversalNewlines(string(content)), &parsed); err != nil {
		return nil, parseError(fmt.Sprintf("%s does not parse: %v; fix it", name, err))
	}
	return parsed, nil
}

func expandPath(base, name string) (string, error) {
	if rest, ok := strings.CutPrefix(name, "~/"); ok {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, rest), nil
	}
	if filepath.IsAbs(name) {
		return filepath.Clean(name), nil
	}
	return filepath.Join(base, name), nil
}

func (c coordinator) locate(root, segment string) string {
	candidates := []string{filepath.Join(root, segment), filepath.Join(filepath.Dir(root), segment)}
	if c.found {
		if member, ok := c.members[segment]; ok {
			candidates = append(candidates, member)
		}
		candidates = append(candidates, filepath.Join(filepath.Dir(c.dir), segment))
	}
	for _, candidate := range candidates {
		if exists(candidate) {
			return candidate
		}
	}
	return ""
}

func (c coordinator) unresolved(segment string) string {
	switch {
	case c.found:
		return fmt.Sprintf("no %s/ in this repository, beside it, in the [workspace.paths] of coordinator %s, or beside it",
			segment, c.dir)
	case c.dir != "":
		return fmt.Sprintf("no %s/ in this repository or beside it, and coordinator %s does not exist",
			segment, c.dir)
	default:
		return fmt.Sprintf("no %s/ in this repository or beside it", segment)
	}
}
