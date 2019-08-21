package renderer

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/hairyhenderson/gomplate"
)

type (
	Renderer interface {
		Render() (*bytes.Buffer, error)
	}

	Options struct {
		TemplateReaders map[string]io.Reader
		ValueReaders    map[string]io.Reader
		RootNamespace   string
	}

	Values map[string]interface{}

	renderer struct {
		templateReaders map[string]io.Reader
		valueReaders    map[string]io.Reader
	}
)

func New(opt *Options) Renderer {
	return &renderer{
		templateReaders: opt.TemplateReaders,
		valueReaders:    opt.ValueReaders,
	}
}

func (r *renderer) Render() (*bytes.Buffer, error) {
	var content []string
	out := new(bytes.Buffer)
	root := Values{}

	for _, reader := range r.templateReaders {
		res, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		content = append(content, string(res))
	}
	for name, reader := range r.valueReaders {
		vals := Values{}
		res, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(res))
		err = yaml.Unmarshal([]byte(string(res)), &vals)
		if err != nil {
			return nil, err
		}
		root[name] = vals

	}

	tmpl := template.New("")
	tmpl.Funcs(gomplate.Funcs(nil))
	tmpl.Parse(strings.Join(content, "\n"))
	err := tmpl.Execute(out, root)
	if err != nil {
		return nil, err
	}
	return out, nil
}
