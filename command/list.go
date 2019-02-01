package command

import (
	. "dr/model"
	"github.com/urfave/cli"
)

// ListRepositoryImages lists all available docker images in the configured repository
func ListRepositoryImages(c *cli.Context) error {

	// get a valid client
	client := GetCustomClient(c)
	// perform operation on server
	imageList, e := client.Repositories()
	if e != nil {
		return e
	}

	// Automatically print
	return Print(c, imageList)
}
