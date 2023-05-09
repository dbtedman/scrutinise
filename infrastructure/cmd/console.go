package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CobraLog(cmd *cobra.Command, format string, a ...any) {
	_, _ = fmt.Fprintf(cmd.OutOrStdout(), format, a...)
}
