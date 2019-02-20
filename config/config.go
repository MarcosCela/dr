// Package config contains all configuration related code, and helper methods to extract clients from the given urfave/cli context
package config

import (
	"dr/errcodes"
	"dr/printer"
	"errors"
	"fmt"
	"github.com/appscode/docker-registry-client/registry"
	"github.com/urfave/cli"
	"github.com/zalando/go-keyring"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// CONFIG contains the key where to store the configuration in the urfave/cli app metadata
const CONFIG = "config"
const OutputFormatFlagName = "output"

const (
	JSONOutputFormat  = "json"
	PLAINOutputFormat = "plain"
	YAMLOutputFormat  = "yaml"
)

// DrContext contains all the necessary data to contact a given docker registry
type DrContext struct {
	Name    string `yaml:"name" json:"name"`
	URL     string `yaml:"url" json:"url"`
	User    string `yaml:"user" json:"user"`
	Trusted bool   `yaml:"trusted" json:"trusted"`
}

// DrConfig contains all available configuration data for the application
type DrConfig struct {
	// Name of the currently configured context, as a string
	CurrentContext string `yaml:"currentContext"`
	// List of available contexts. The user should be able to change the active context manually or via config
	Contexts []DrContext `yaml:"contexts"`
	// OutputFormat is the preferred output format when no other option is available
	OutputFormat string `yaml:"outputFormat" json:"outputFormat"`
}

type OutputFormat struct {
	Name        string   `yaml:"name"`
	Aliases     []string `yaml:"aliases"`
	Description string   `yaml:"description"`
}

var OutputFormatJSON = OutputFormat{
	Name:        JSONOutputFormat,
	Aliases:     []string{},
	Description: "JSON based formatter, usually the default for complex outputs that are not possible to represent in 'plaintext'.",
}
var OutputFormatPLAIN = OutputFormat{
	Name:        PLAINOutputFormat,
	Aliases:     []string{"text", "plaintext", "bash"},
	Description: "Plain bash-like syntax (e.g: lists have each element in a new line) useful to chain executions or to use in bash scripts",
}
var OutputFormatYAML = OutputFormat{
	Name:        YAMLOutputFormat,
	Description: "YAML based formatter, alternative to JSON when showing complex outputs that require structured text",
}
var validOutputFormats = []OutputFormat{OutputFormatJSON, OutputFormatPLAIN, OutputFormatYAML}

// GetClient generates a docker registry v2 client, with data from the current context configured by the user
// passed as metadata
func GetClient(c *cli.Context) *registry.Registry {
	context := c.App.Metadata[CONFIG].(DrConfig)
	if context.CurrentContext == "" {
		log.Fatal("Need a valid configuration file!")
	}

	currentContext := getCurrentContext(context)
	password := getPasswordForContext(currentContext)
	hub, err := registry.New(currentContext.URL, currentContext.User, password)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
	return hub
}

func getPasswordForContext(context DrContext) string {
	// get password
	secret, err := keyring.Get(context.Name, context.User)
	if err != nil {
		log.Fatal(err)
	}
	return secret
}

// getCurrentContext extracts the current context for the user configuration from the CLI context (params etc...)
func getCurrentContext(context DrConfig) DrContext {
	var currentContext DrContext
	for _, c := range context.Contexts {
		if c.Name == context.CurrentContext {
			currentContext = c
			break
		}
	}
	return currentContext
}

// GetJSONPrinter Returns the JSON printer, without checking if the user has the JSON printer configured or not
func GetJSONPrinter(c *cli.Context) printer.Printer {
	return printer.JSONPrinter(c.App.Writer)
}

// GetPrinter returns a printer that dumps content to the App writer depending on what kind of writer is configured
// as an example, it could return a writer that dumps json-formatted content to stdout, or a writer that
// dumps html-like content to a file
func GetPrinter(c *cli.Context) printer.Printer {
	outputFormat := getConfiguredOutputFormat(c)
	return printerByFormat(outputFormat, c.App.Writer)

}

func getConfiguredOutputFormat(c *cli.Context) string {
	// Default value
	outputFormat := JSONOutputFormat
	// Acquire value from configuration
	outputFormat = GetCurrentConfiguration(c).OutputFormat
	// If flag is set, acquire value from flag and overwrite
	if c.String(OutputFormatFlagName) != "" {
		outputFormat = c.String(OutputFormatFlagName)
	}
	return outputFormat
}

// printerByFormat attempts to return a valid Printer implementation depending on the user preferences,
// defaulting to JSON output
func printerByFormat(outputFormat string, writer io.Writer) printer.Printer {
	switch outputFormat {
	case PLAINOutputFormat:
		return printer.PlainPrinter(writer)
	case JSONOutputFormat:
		return printer.JSONPrinter(writer)
	default:
		return printer.JSONPrinter(writer)
	}
}

// GetCurrentConfiguration returns the whole configuration for the application
func GetCurrentConfiguration(c *cli.Context) DrConfig {
	return c.App.Metadata[CONFIG].(DrConfig)

}

// New gets a new configuration, trying to parse the configuration in ORDER from the given list of system paths
// The function will attempt all paths, returning when the first configuration is correctly loaded
func New(paths ...string) (*DrConfig, error) {
	if len(paths) == 0 {
		return loadConfigFromPath(defaultConfigurationFilePath())
	}
	for _, path := range paths {
		config, err := loadConfigFromPath(path)
		if err == nil {
			return config, nil
		}
	}
	return nil, errors.New("Could not load any valid configuration from the given paths: " + strings.Join(paths, " "))
}

// loadConfigFromPath Loads configuration from a YAML string
func loadConfigFromPath(s string) (*DrConfig, error) {
	yamlFile, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}
	configuration := DrConfig{}
	err = yaml.Unmarshal(yamlFile, &configuration)
	if err != nil {
		return nil, err
	}
	return &configuration, nil
}

// SaveConfigToPath saves the provided configuration to the configured file path for the app configuration
func SaveConfigToPath(config DrConfig) error {
	data, e := yaml.Marshal(config)
	if e != nil {
		return cli.NewExitError("Could not save the configuration file to the path: "+defaultConfigurationFilePath()+" error: "+e.Error(), errcodes.WriteConfigError)
	}
	return ioutil.WriteFile(defaultConfigurationFilePath(), data, 0644)
}

// defaultConfigurationFilePath returns the default configuration file path (~/.config/dr/config.yaml),
// or alternatively, returns the configured file in the DR_CONFIG environment variable
func defaultConfigurationFilePath() string {
	value := os.Getenv("DR_CONFIG")
	if len(value) == 0 {
		return filepath.FromSlash(getUserHomeDir() + "/.config/dr/config.yaml")
	}
	return value
}

// getUserHomeDir returns the home directory for the user (in UNIX and Windows)
func getUserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func GetValidOutputFormatsShort() (retVal []string) {
	for _, outputFormat := range validOutputFormats {
		retVal = append(retVal, outputFormat.Name)
	}
	return retVal
}
func GetValidOutputFormatsLong() []OutputFormat {
	return validOutputFormats
}
