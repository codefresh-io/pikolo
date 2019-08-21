package cmd

import (
	"os"

	"github.com/codefresh-io/pikolo/pkg/logger"
)

func dieOnError(err error, logger logger.Logger) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
