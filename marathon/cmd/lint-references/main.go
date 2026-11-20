package main

import (
	"os"

	"github.com/standards-lab/claude-plugins/marathon/internal/references"
	"github.com/standards-lab/claude-plugins/marathon/internal/version"
)

func main() {
	if version.Report("lint-references", os.Args[1:], os.Stdout) {
		return
	}
	os.Exit(references.Run(".", os.Stdout, os.Stderr))
}
