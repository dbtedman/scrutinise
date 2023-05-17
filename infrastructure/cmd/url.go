package cmd

import (
	"github.com/spf13/cobra"
)

func URLCommand(resultCh *chan int) *cobra.Command {
	return &cobra.Command{
		Use: "url",
		Run: func(cmd *cobra.Command, args []string) {
			*resultCh <- SuccessResult
		},
	}
}
