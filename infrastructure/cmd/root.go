package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func RootCommand(resultCh *chan int) *cobra.Command {
	return &cobra.Command{
		Use:   "scrutinise",
		Short: "Tool to scrutinise website development security.",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()

			if err != nil {
				log.Println(err)
				*resultCh <- ErrorResult
			} else {
				*resultCh <- SuccessResult
			}
		},
	}
}
