package cmd

import (
	"fmt"
	"io"
	"os"
	"path"

	"strings"

	"github.com/codefresh-io/pikolo/pkg/logger"
	"github.com/codefresh-io/pikolo/pkg/renderer"
	"github.com/spf13/cobra"
)

var renderCmdOptions struct {
	templates   []string
	values      []string
	rootContext string
	rightDelim  string
	leftDelim   string
}

var renderCmd = &cobra.Command{
	Use: "render",
	Run: func(cmd *cobra.Command, args []string) {
		templateReaders := make(map[string]io.Reader)
		valueReaders := make(map[string][]io.Reader)
		log := logger.New(nil)
		if len(renderCmdOptions.templates) == 0 {
			dieOnError(fmt.Errorf("No tempalte given"), log)
		}
		for _, templatePath := range renderCmdOptions.templates {
			file, err := os.Open(templatePath)
			dieOnError(err, log)
			templateReaders[templatePath] = file
			defer file.Close()
		}

		for _, valuePath := range renderCmdOptions.values {
			var file *os.File
			var err error
			var valueReader io.Reader
			values := strings.Split(valuePath, "=")
			if len(values) == 1 {
				file, err = os.Open(values[0])
				dieOnError(err, log)
				if strings.HasSuffix(values[0], ".json") {
					valueReader, err = jsonToYaml(file)
					dieOnError(err, log)
				} else {
					valueReader = file
				}
				if _, ok := valueReaders[renderCmdOptions.rootContext]; !ok {
					valueReaders[renderCmdOptions.rootContext] = []io.Reader{valueReader}
				} else {
					valueReaders[renderCmdOptions.rootContext] = append(valueReaders[renderCmdOptions.rootContext], valueReader)
				}
			} else {
				file, err = os.Open(values[1])
				dieOnError(err, log)
				if strings.HasSuffix(values[1], ".json") {
					valueReader, err = jsonToYaml(file)
					dieOnError(err, log)
				} else {
					valueReader = file
				}
				if _, ok := valueReaders[values[0]]; !ok {
					valueReaders[values[0]] = []io.Reader{valueReader}
				} else {
					valueReaders[values[0]] = append(valueReaders[values[0]], valueReader)
				}
			}
			defer file.Close()
		}

		engine := renderer.New(&renderer.Options{
			TemplateReaders: templateReaders,
			ValueReaders:    valueReaders,
			LeftDelim:       renderCmdOptions.leftDelim,
			RightDelim:      renderCmdOptions.rightDelim,
			Name:            path.Base(renderCmdOptions.templates[0]),
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
	renderCmd.Flags().StringVar(&renderCmdOptions.leftDelim, "left-delim", "{{", "Left delimiter (default: <<)")
	renderCmd.Flags().StringVar(&renderCmdOptions.rightDelim, "right-delim", "}}", "Right delimiter (default: >>)")
}
