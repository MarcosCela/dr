package command

import (
	"dr/config"
	"dr/errcodes"
	"github.com/urfave/cli"
)

// SetCurrentContext sets the current context to the provided context on the first parameter
func SetCurrentContext(c *cli.Context) error {
	currentConfiguration := config.GetCurrentConfiguration(c)
	context := ""

	if len(currentConfiguration.Contexts) == 0 {
		return cli.NewExitError("Invalid configuration file. You need at least one context!", errcodes.InvalidContext)
	}
	if c.Args().Present() {
		context = c.Args().First()
	} else {
		return cli.NewExitError("You need to provide the name of the context", errcodes.InvalidContext)

	}

	if !contextExists(context, currentConfiguration.Contexts) {
		return cli.NewExitError("The provided context: '"+context+"' is not a valid context defined on the configuration!", errcodes.InvalidContext)
	}

	// Context is valid. Set the current context to the provided context
	currentConfiguration.CurrentContext = context
	return config.SaveConfigToPath(currentConfiguration)
}

// GetCurrentContexts returns the list of currently available contexts
func GetCurrentContexts(c *cli.Context) error {
	currentConfiguration := config.GetCurrentConfiguration(c)

	if len(currentConfiguration.Contexts) == 0 {
		return cli.NewExitError("Invalid configuration file. You need at least one context!", errcodes.InvalidContext)
	}

	contextNames := make([]string, len(currentConfiguration.Contexts))
	for i, context := range currentConfiguration.Contexts {
		contextNames[i] = context.Name
	}
	config.GetPrinter(c).Print(contextNames)
	return nil
}
