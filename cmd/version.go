package cmd

import (
	"fmt"

	"github.com/longkey1/gosla/internal/version"
	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("gosla version %s\n", version.Version)
			fmt.Printf("  commit: %s\n", version.CommitSHA)
			fmt.Printf("  built:  %s\n", version.BuildTime)
		},
	}
}
