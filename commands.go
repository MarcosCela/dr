package main

import (
	"dr/command"
	"dr/config"
	"dr/errcodes"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

var outputFlag = cli.StringFlag{
	Name:  "output, o",
	Value: config.JSONOutputFormat,
	Usage: "output format for all commands [json, plain]",
}

// Commands available to the application. Array of all commands that can be used in dr
var Commands = []cli.Command{
	{
		Name:      "list",
		ShortName: "ls",
		Usage:     "List remote images for the default repository",
		Action:    command.ListRepositoryImages,
		Flags: append([]cli.Flag{
			cli.BoolFlag{
				Name:   "t, tags, show-tags",
				Hidden: false,
				Usage:  "When showing images, also show the tag (as a suffix)",
			},
			cli.BoolFlag{
				Name:   "repo, show-repo",
				Hidden: false,
				Usage:  "Prefix each image name with the repository where it comes from",
			},
		}, outputFlag),
		UsageText:    "dr list",
		BashComplete: command.TagsAuto,
	},
	{
		Name:         "tags",
		ShortName:    "t",
		Aliases:      []string{"tag"},
		Usage:        "List available tags for a remote image",
		Action:       command.Tags,
		Flags:        []cli.Flag{outputFlag},
		UsageText:    "dr tags busybox",
		ArgsUsage:    "dr tags <repository>",
		BashComplete: command.TagsAuto,
	},
	{
		Name:      "context",
		ShortName: "cont",
		Usage:     "Manage configuration for contexts",
		UsageText: "dr context <subcommand>",
		Subcommands: []cli.Command{
			{
				Name:      "set",
				Aliases:   []string{"use"},
				Usage:     "Set the current context. The name of the context must exist",
				Action:    command.SetCurrentContext,
				ArgsUsage: "Name of the context that you want to use as the 'currentContext'.",
				UsageText: "dr context set <name-of-the-context>",
			},
			{
				Name:      "get",
				Aliases:   []string{"list", "ls", "show"},
				Usage:     "Get the currently configured contexts, including the name",
				Action:    command.GetCurrentContexts,
				UsageText: "dr context get",
				Flags:     []cli.Flag{outputFlag},
			},
		},
	},
	{
		Name:      "config",
		ShortName: "conf",
		Usage:     "Manage configuration",
		Flags:     []cli.Flag{},
		UsageText: "dr config <subcommand>",
		Subcommands: []cli.Command{
			{
				Name:      "get",
				Aliases:   []string{"show"},
				Usage:     "Get the current configuration",
				Action:    command.ShowCurrentConfiguration,
				UsageText: "dr config get",
			},
			{
				Name:      "set",
				Usage:     "Set a configuration parameter",
				UsageText: "dr config set <subcommand>",
				Subcommands: []cli.Command{
					{
						Name:      "output",
						Aliases:   []string{"o", "format", "fmt"},
						Usage:     "Set the default output to the specified",
						UsageText: "dr config set output <output format> -- Valid formats are: [" + strings.Join(config.GetValidOutputFormatsShort(), ", ") + "]",
						Action:    command.SetOutputFormat,
					},
				},
			},
		},
	},
	{
		Name:         "info",
		ShortName:    "i",
		Aliases:      []string{"manifest", "mf", "m"},
		Usage:        "Show additional information about a layer",
		Flags:        []cli.Flag{outputFlag},
		UsageText:    "dr info image:tag",
		Action:       command.ManifestInfo,
		BashComplete: command.ManifestInfoAuto,
	},
}

// CommandNotFound default handler for when a command is not found, with custom exit error
func CommandNotFound(c *cli.Context, command string) {
	_, _ = fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.\n", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(errcodes.CommandNotFound)
}
