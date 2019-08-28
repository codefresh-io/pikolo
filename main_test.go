package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

type (
	Tests []E2E

	E2E struct {
		Title     string   `yaml:"title"`
		Arguments []string `yaml:"arguments"`
		Output    string   `yaml:"output"`
	}
)

func Test_main(t *testing.T) {

	table := Tests{}
	tests, err := ioutil.ReadFile("./testdata/tests.yaml")
	dieOnError(err)
	err = yaml.Unmarshal(tests, &table)
	dieOnError(err)

	for _, test := range table {
		fmt.Println(test.Title)
		t.Run(test.Title, func(tt *testing.T) {
			fmt.Println("Runnig test")
			pikolo := exec.Command("/tmp/pikolo-test", test.Arguments...)
			out, err := pikolo.Output()
			if err != nil {
				assert.Fail(tt, "Should not fail")
			}
			assert.Equal(tt, test.Output, string(out))
		})
	}

}

func dieOnError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}
