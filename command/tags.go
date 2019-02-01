package command

import (
	"dr/config"
	"dr/errcodes"
	. "dr/model"
	"fmt"
	"github.com/urfave/cli"
	"sort"
	"strings"
)

// Tags provides the available tags for a given image. The image repository is extracted from the first parameter,
// passed thought the context.
func Tags(c *cli.Context) error {

	// get a valid client
	client := GetCustomClient(c)

	// get the repository from params
	repository, e := getRepositoryFromParameters(c)
	if e != nil {
		return e
	}
	tags, e := client.Tags(repository)
	if e != nil {
		return cli.NewExitError("Could not list repositories. Error:  "+e.Error(), errcodes.CannotListRepo)
	}
	return Print(c, tags)
}

// TagsAuto Auto-complete for tags
func TagsAuto(c *cli.Context) {
	hub := config.GetClient(c)
	repos, e := hub.Repositories()
	if e != nil {
		// On autocomplete function, silently fail
		return
	}
	// Try to get the first parameter, that will be a prefix
	// of the repository that the user is trying to list tags of
	// dr tags <userInput>
	userInput := ""
	if c.Args().Present() {
		userInput = c.Args().First()
	}

	sort.Strings(repos)
	for _, repo := range repos {
		if strings.HasPrefix(repo, userInput) {
			// Print only those repositories that start with the
			// given prefix.
			// If the user does not input any prefix, then just
			// print all repositories
			fmt.Println(repo)
		}
	}

}
