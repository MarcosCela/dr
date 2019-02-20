// Package errcodes :Contains all information related to errors.
// Usually error codes and other information
package errcodes

import "github.com/urfave/cli"

// CannotListRepo signals a problem when listing the images present in a repository
const CannotListRepo int = 1

// InvalidContext is used when the context is present, but has problems
// (bad format, requested a context that does not exist...)
const InvalidContext int = 2

// InsufficientArgs marks that the command needs parameters that were not provided
const InsufficientArgs int = 3

// CommandNotFound signals that the command supplied is not a valid command
const CommandNotFound int = 4

// TooManyArgs marks that the command was issued with more parameters than the needed parameters
const TooManyArgs int = 5

// CannotRetrieveManifest happens when there is an error retrieving the manifest for an image (E.g: the image does not exist, cannot contact the repo...)
const CannotRetrieveManifest int = 6

// ImageNotFound happens when an operation that depends of the existence of an image finds that the image does not exist
const ImageNotFound int = 7

// TemplateDisplayError happens when there is an unexpected error when showing output to terminal (usually printing results) and there is an error on the last phase (printing)
const templateDisplayError int = 8

// CannotListTags an error is blocking the content display when listing tags
const CannotListTags int = 9

// CannotListTags general error when writing configuration to the configuration file. Usually signals bad permissions, I/O error or other
// file-io related error (path not existing...)
const WriteConfigError int = 10

// InvalidOutputFormat is used when the user selects an output format that is not valid
const InvalidOutputFormat int = 11

// TerminalReadError is used when the user needs to input some data from the terminal and any kind of I/O related error happens
const TerminalReadError int = 12

// KeyringError signals any kind of error that happened when saving a password to the terminal
const KeyringError int = 13

// pre-built errors

var TemplateDisplayError = cli.NewExitError("There was an unexpected error when showing output to terminal", templateDisplayError)
var NotImplementedError = cli.NewExitError("This function is not implemented yet!", -1)
