package command

import (
	"dr/config"
	"dr/errcodes"
	"github.com/urfave/cli"
	"strings"
)

// ShowCurrentConfiguration shows the current configuration
func ShowCurrentConfiguration(c *cli.Context) error {
	currentConfiguration := config.GetCurrentConfiguration(c)

	if len(currentConfiguration.Contexts) == 0 {
		return cli.NewExitError("Invalid configuration file. You need at least one context!", errcodes.InvalidContext)
	}
	currentConfiguration.CensorPasswords()
	config.GetJSONPrinter(c).Print(currentConfiguration)

	return nil
}
func SetOutputFormat(c *cli.Context) error {
	if c.Args().Present() {
		// See if the first argument is a valid output format
		outputFormat := c.Args().First()
		if !outputFormatIsValid(outputFormat) {
			return cli.NewExitError("The selected output format is not valid. Valid values: "+strings.Join(config.GetValidOutputFormatsShort(), " "), errcodes.InvalidOutputFormat)
		}
		// Set the output format to the provided format
		currentConfiguration := config.GetCurrentConfiguration(c)
		currentConfiguration.OutputFormat = outputFormat
		return config.SaveConfigToPath(currentConfiguration)
	} else {
		config.GetPrinter(c).Print(config.GetValidOutputFormatsLong)
		return cli.NewExitError("You need to provide at least an argument to set the output format, see the "+
			"previous output for mor information!", errcodes.InsufficientArgs)
	}

}

func outputFormatIsValid(format string) bool {
	for _, validFormat := range config.GetValidOutputFormatsShort() {
		if format == validFormat {
			return true
		}
	}
	return false
}
