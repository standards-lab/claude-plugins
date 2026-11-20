package main

import (
	"os"

	"github.com/standards-lab/claude-plugins/marathon/internal/contextlint"
	"github.com/standards-lab/claude-plugins/marathon/internal/version"
)

func main() {
	if version.Report("lint-context", os.Args[1:], os.Stdout) {
		return
	}
	os.Exit(contextlint.Run(".", os.Stdout, os.Stderr))
}
