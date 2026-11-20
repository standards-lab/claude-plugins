package main

import (
	"os"

	"github.com/standards-lab/claude-plugins/marathon/internal/marathon"
	"github.com/standards-lab/claude-plugins/marathon/internal/version"
)

func main() {
	if version.Report("marathon", os.Args[1:], os.Stdout) {
		return
	}
	os.Exit(marathon.Run(".", os.Args[1:], os.Stdout, os.Stderr))
}
