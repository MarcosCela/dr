package printer

import (
	"github.com/urfave/cli"
	"io"
)

const includeTagsFlag string = "tags"
const includeRepoFlag string = "repo"

// Printer represents a generic printer
type PrintOptions struct {
	IncludeTags     bool      `yaml:"includeTags"`
	IncludeRegistry bool      `yaml:"includeRegistry"`
	Writer          io.Writer `yaml:"writer"`
}

func GetPrintOptions(c *cli.Context) PrintOptions {
	return PrintOptions{
		IncludeRegistry: c.Bool(includeRepoFlag),
		IncludeTags:     c.Bool(includeTagsFlag),
		Writer:          c.App.Writer,
	}
}
