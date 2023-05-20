package cmd

import (
	"net/url"

	"github.com/dbtedman/scrutinise/application"
	"github.com/spf13/cobra"
)

func URLCommand(resultCh *chan int) *cobra.Command {
	return &cobra.Command{
		Use: "url",
		Run: func(cmd *cobra.Command, args []string) {
			requestUrl, _ := url.Parse("https://example.com")

			CobraLog(cmd, "preparing to scrutinise url; url=[%s]", requestUrl)

			uc, ucError := application.NewScrutiniseURLUC()

			if ucError != nil {
				CobraError(cmd, "cannot setup scrutinise url use case; error=[%s]", ucError.Error())
				*resultCh <- ErrorResult
				return
			}

			result, resultError := uc.Execute(application.ScrutiniseURLUCRequest{
				Url: *requestUrl,
			})

			if resultError != nil {
				CobraError(cmd, "error occurred while executing url use case; error=[%s]", resultError.Error())
				*resultCh <- ErrorResult
				return
			}

			if result == nil {
				CobraError(cmd, "error occurred while executing url use case; error=[%s]", resultError.Error())
				*resultCh <- ErrorResult
				return
			}

			*resultCh <- SuccessResult
		},
	}
}
