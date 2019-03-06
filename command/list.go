package command

import (
	. "dr/model"
	"github.com/urfave/cli"
	"strings"
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

	// If the user passes a parameter to "ls", consider it a prefix and return only images that start with that prefix
	prefix := getPrefix(c)
	imageList = filterImageListBasedOnPrefix(imageList, prefix)

	// Automatically print
	return Print(c, imageList)
}

/* filterImageListBasedOnPrefix returns a subset of the original list based on the name of each Docker image,
where only the images that start with the prefix are included.
*/
func filterImageListBasedOnPrefix(list DockerImageList, prefix string) (filtered DockerImageList) {
	for _, v := range list.Images {
		if strings.HasPrefix(v.Name, prefix) {
			filtered.Images = append(filtered.Images, v)
		}
	}
	return filtered

}

// getPrefix returns the first argument passed by the user, or an empty string if the user did not add a first parameter
func getPrefix(context *cli.Context) string {
	if len(context.Args()) >= 0 {
		return context.Args().First()
	} else {
		return ""
	}
}
