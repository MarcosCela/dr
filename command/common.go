package command

import (
	"dr/config"
	"dr/errcodes"
	"github.com/urfave/cli"
	"strings"
)

// getRepositoryFromParameters extracts the repository from the first parameter issued by the user
func getRepositoryFromParameters(context *cli.Context) (string, error) {
	if context.Args().Present() {
		return context.Args().First(), nil
	}
	return "", cli.NewExitError("You must provide the repository and image to get the available tags", errcodes.InsufficientArgs)

}

func getRepositoryAndTagFromParameters(context *cli.Context) (string, string, error) {
	if !context.Args().Present() {
		return "", "", cli.NewExitError("You must provide at least a valid image name", errcodes.InsufficientArgs)
	}
	if context.NArg() > 2 {
		return "", "", cli.NewExitError("Too many arguments, call takes only '<image>:<tag>' or '<image>'  or '<image> <tag>' (tag assumed as latest)", errcodes.TooManyArgs)

	}
	image, tag := extractImageAndTag(context.Args())
	return image, tag, nil

}

func extractImageAndTag(args []string) (string, string) {
	image := ""
	tag := "latest"
	fusedImageAndTag := args[0]
	if len(args) == 1 {
		// first arg is of the form "<myimage>:<mytag>" or just "<myimage>"
		if strings.Contains(fusedImageAndTag, ":") {
			// Image is of the form "<myimage>:<mytag>"
			split := strings.Split(fusedImageAndTag, ":")
			image = split[0]
			tag = split[1]
		} else {
			// Image is of the form "<myimage>"
			image = fusedImageAndTag
		}

	} else if len(args) == 2 {
		image = args[0]
		tag = args[1]
	}

	return image, tag
}

func choose(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func contextExists(context string, allContexts []config.DrContext) bool {
	for _, isContext := range allContexts {
		if isContext.Name == context {
			return true
		}
	}
	return false
}
