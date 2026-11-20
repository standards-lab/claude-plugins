package version

import (
	"fmt"
	"io"
)

var version = "dev"

// Report writes "<name> <version>" and a newline to stdout and returns true
// when args is --version alone, and otherwise writes nothing and returns
// false.
func Report(name string, args []string, stdout io.Writer) bool {
	if len(args) != 1 || args[0] != "--version" {
		return false
	}
	_, _ = fmt.Fprintf(stdout, "%s %s\n", name, version)
	return true
}
