package cmd

import (
	"github.com/spf13/cobra"
)

func RootCommand(resultCh *chan int) *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "scrutinise",
		Short: "Tool to scrutinise website development security.",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()

			if err != nil {
				*resultCh <- ErrorResult
			} else {
				*resultCh <- SuccessResult
			}
		},
	}

	rootCommand.AddCommand(VersionCommand(resultCh))

	return rootCommand
}
