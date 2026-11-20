package main

import (
	"os"

	"github.com/standards-lab/claude-plugins/marathon/internal/version"
	"github.com/standards-lab/claude-plugins/marathon/internal/wrapformats"
)

func main() {
	if version.Report("wrap-formats", os.Args[1:], os.Stdout) {
		return
	}
	os.Exit(wrapformats.Run(os.Stdin, os.Stdout))
}
