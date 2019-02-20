package command

import (
	"dr/config"
	"dr/errcodes"
	"fmt"
	"github.com/urfave/cli"
	"github.com/zalando/go-keyring"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

// Login ensures that a password is saved for a given context in the OS keyring
func Login(c *cli.Context) error {
	if !c.Args().Present() {
		return cli.NewExitError("You must provide a context name to login", errcodes.InsufficientArgs)
	}
	var contextName = c.Args().First()
	if !contextExists(contextName, config.GetCurrentConfiguration(c).Contexts) {
		return cli.NewExitError("The context you are trying to login-to does not exist", errcodes.InvalidContext)
	}
	user, err := getUserForContext(contextName, config.GetCurrentConfiguration(c))
	if err != nil {
		return err
	}
	fmt.Println("Enter a password:")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return cli.NewExitError("There was an error getting the password from the terminal: "+err.Error(), errcodes.TerminalReadError)
	}

	err = keyring.Set(contextName, user, string(bytePassword))
	if err != nil {
		return cli.NewExitError("There was an error saving the password to the keyring: "+err.Error(), errcodes.KeyringError)
	}
	fmt.Println("Correctly set the password for context: <" + contextName + "> and user: <" + user + ">")
	return nil
}

func getUserForContext(contextName string, configuration config.DrConfig) (string, error) {
	for _, context := range configuration.Contexts {
		if context.Name == contextName {
			if context.User == "" {
				return "", cli.NewExitError("The username for the context: "+contextName+" is empty!", errcodes.InvalidContext)
			}
			return context.User, nil
		}
	}
	return "", cli.NewExitError("Could not recover a valid username for the context: <"+contextName+">", errcodes.InvalidContext)
}

func LoginComplete(c *cli.Context) {
	currentConfiguration := config.GetCurrentConfiguration(c)

	if len(currentConfiguration.Contexts) == 0 {
		return
	}

	contextNames := getContextNames(currentConfiguration)
	for _, contextName := range contextNames {
		fmt.Println(contextName)
	}
}

func getContextNames(cfg config.DrConfig) []string {
	contextNames := make([]string, len(cfg.Contexts))
	for i, context := range cfg.Contexts {
		contextNames[i] = context.Name
	}
	return contextNames
}
