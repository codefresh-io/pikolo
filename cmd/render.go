package cmd

import (
	"fmt"
	"io"
	"os"

	"strings"

	"github.com/codefresh-io/pikolo/pkg/logger"
	"github.com/codefresh-io/pikolo/pkg/renderer"
	"github.com/spf13/cobra"
)

var renderCmdOptions struct {
	templates   []string
	values      []string
	rootContext string
}

var renderCmd = &cobra.Command{
	Use: "render",
	Run: func(cmd *cobra.Command, args []string) {
		templateReaders := make(map[string]io.Reader)
		valueReaders := make(map[string]io.Reader)
		log := logger.New(nil)
		for _, templatePath := range renderCmdOptions.templates {
			file, err := os.Open(templatePath)
			dieOnError(err, log)
			templateReaders[templatePath] = file
			defer file.Close()
		}

		for _, valuePath := range renderCmdOptions.values {
			var file *os.File
			var err error
			values := strings.Split(valuePath, "=")
			if len(values) == 1 {
				file, err = os.Open(values[0])
				dieOnError(err, log)
				valueReaders[renderCmdOptions.rootContext] = file
			} else {
				file, err = os.Open(values[1])
				dieOnError(err, log)
				valueReaders[values[0]] = file
			}
			defer file.Close()
		}

		engine := renderer.New(&renderer.Options{
			TemplateReaders: templateReaders,
			ValueReaders:    valueReaders,
		})
		res, err := engine.Render()
		dieOnError(err, log)
		fmt.Println(res.String())
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)
	renderCmd.Flags().StringArrayVar(&renderCmdOptions.templates, "template", []string{}, "Path to template file")
	renderCmd.Flags().StringArrayVar(&renderCmdOptions.values, "value", []string{}, "Path to value file")
	renderCmd.Flags().StringVar(&renderCmdOptions.rootContext, "root-namespace", "Values", "Name of the root namespace (default: Values)")
}
