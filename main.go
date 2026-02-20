package main

import (
	"os"

	"github.com/longkey1/gosla/cmd"
	"github.com/longkey1/gosla/internal/version"
)

// Version information (set by ldflags)
var (
	ver    = "dev"
	commit = "unknown"
	date   = "unknown"
)

func main() {
	version.Version = ver
	version.CommitSHA = commit
	version.BuildTime = date

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
