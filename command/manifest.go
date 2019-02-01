package command

import (
	"dr/config"
	. "dr/model"
	"fmt"
	"github.com/urfave/cli"
	"sort"
	"strings"
)

// ManifestInfo shows additional information for a given manifest
func ManifestInfo(c *cli.Context) error {
	// get a valid client
	client := GetCustomClient(c)

	repository, tag, e := getRepositoryAndTagFromParameters(c)
	if e != nil {
		return e
	}

	imageManifest, e := client.Manifest(repository, tag)
	if e != nil {
		return e
	}
	return Print(c, imageManifest)

}

func ManifestInfoAuto(c *cli.Context) {
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
			tags, e := hub.Tags(repo)
			if e != nil {
				continue
			}
			for _, tag := range tags {
				fmt.Println(repo + "\\:" + tag)
			}
		}
	}

}
