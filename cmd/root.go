package cmd

import (
	"github.com/codefresh-io/pikolo/pkg/logger"
	"github.com/spf13/cobra"
)

var rootCmdOptions struct {
	Verbose bool
}

var rootCmd = &cobra.Command{
	Use: "pikolo",
}

// Execute - execute the root command
func Execute() {
	err := rootCmd.Execute()
	dieOnError(err, logger.New(nil))
}
