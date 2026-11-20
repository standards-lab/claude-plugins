package main

import (
	"os"

	"github.com/standards-lab/claude-plugins/marathon/internal/contextlint"
)

func main() {
	os.Exit(contextlint.Run(".", os.Stdout, os.Stderr))
}
