package model

import (
	"dr/errcodes"
	"dr/printer"
	"dr/templates"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"gopkg.in/yaml.v2"
	"text/template"
)

type TagList struct {
	Tags []string `yaml:"tags" json:"tags"`
}

// Print prints a message with the JSON formatter
func (printer TagList) PrintPlain(printable interface{}, options printer.PrintOptions) error {
	tagList, ok := printable.(TagList)
	if !ok {
		return errcodes.NotImplementedError
	}
	return template.Must(template.New("list-repository-tags").
		Parse(templates.TagListPlain)).
		Execute(options.Writer, tagList)
}

func (printer TagList) PrintYaml(printable interface{}, options printer.PrintOptions) error {
	tagList, ok := printable.(TagList)
	if !ok {
		return errcodes.TemplateDisplayError
	}
	s, err := yaml.Marshal(tagList)
	if err != nil {
		return errcodes.TemplateDisplayError
	}

	_, err = fmt.Fprintf(options.Writer, "%v\n", string(s))
	if err != nil {
		return errcodes.TemplateDisplayError
	}
	return nil
}

func (printer TagList) PrintJson(printable interface{}, options printer.PrintOptions) error {
	tagList, ok := printable.(TagList)
	if !ok {
		return errcodes.TemplateDisplayError
	}
	s, err := prettyjson.Marshal(tagList)
	if err != nil {
		return errcodes.TemplateDisplayError
	}

	_, err = fmt.Fprintf(options.Writer, "%v\n", string(s))
	if err != nil {
		return errcodes.TemplateDisplayError
	}
	return nil
}
