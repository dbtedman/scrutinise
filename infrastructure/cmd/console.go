package cmd

import (
	"io"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func CobraLog(cmd *cobra.Command, format string, a ...any) {
	logger := Logger(cmd.OutOrStdout())
	logger.Infof(format, a...)
}

func Logger(writer io.Writer) *log.Logger {
	logger := log.New(writer)
	logger.SetReportTimestamp(true)
	return logger
}
