package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/codefresh-io/pikolo/pkg/logger"
	"github.com/ghodss/yaml"
)

type (
	JSON map[string]interface{}
)

func dieOnError(err error, logger logger.Logger) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func jsonToYaml(j io.Reader) (io.Reader, error) {
	buf, err := io.ReadAll(j)
	if err != nil {
		return nil, err
	}
	js := &JSON{}
	err = json.Unmarshal(buf, js)
	if err != nil {
		return nil, err
	}
	b, err := yaml.Marshal(js)
	return bytes.NewReader(b), nil
}
