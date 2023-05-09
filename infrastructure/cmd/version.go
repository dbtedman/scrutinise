package cmd

import (
	"github.com/spf13/cobra"
)

// version, commit, and date are populated during build
var (
	version = "latest"
	commit  = "n/a"
	date    = "n/a"
)

func VersionCommand(resultCh *chan int) *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			CobraLog(cmd, "scrutinise version: %s, commit: %s, built at: %s\n", version, commit, date)
			*resultCh <- SuccessResult
		},
	}
}
