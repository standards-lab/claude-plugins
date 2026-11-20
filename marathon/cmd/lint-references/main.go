package main

import (
	"os"

	"github.com/standards-lab/claude-plugins/marathon/internal/references"
)

func main() {
	os.Exit(references.Run(".", os.Stdout, os.Stderr))
}
