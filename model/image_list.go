package model

import (
	"dr/errcodes"
	. "dr/printer"
	"dr/templates"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"gopkg.in/yaml.v2"
	"text/template"
)

type ImageListTemplate struct {
	Images            []DockerImage
	IncludeTag        bool
	IncludeRepository bool
}

type DockerImage struct {
	Name       string `yaml:"name" json:"name"`
	Tag        string `yaml:"tag,omitempty" json:"tag,omitempty"`
	Repository string `yaml:"repository,omitempty" json:"repository,omitempty"`
}

type DockerImageList struct {
	Images []DockerImage `yaml:"images" json:"images"`
}

// Print prints a message with the JSON formatter
func (printer DockerImageList) PrintPlain(printable interface{}, options PrintOptions) error {

	dockerImageList, ok := printable.(DockerImageList)
	if !ok {
		return errcodes.NotImplementedError
	}
	// If we do not include tags, we can end up with a lot of images
	// for different tags. Ensure that in this case, we do not
	// print a lot of lines for images that only differ on tag (since we are not showing it!)
	if !options.IncludeTags {
		dockerImageList.Images = removeDuplicatedEntries(dockerImageList.Images)
	}
	return template.Must(template.New("list-repository-images").
		Parse(templates.ImageListPlain)).
		Execute(options.Writer, ImageListTemplate{
			Images:            dockerImageList.Images,
			IncludeRepository: options.IncludeRegistry,
			IncludeTag:        options.IncludeTags,
		})
}

/* removeDuplicatedEntries filters the given slice, based only the name of the docker image, and ignoring the tag.
This means that if two images are only different on the tag, only one of them will be included.
*/
func removeDuplicatedEntries(images []DockerImage) (result []DockerImage) {
	// Use  a map to record duplicates as we find them
	encountered := map[string]bool{}
	for v := range images {
		if encountered[images[v].Name] == true {
			// Do not add the duplicate value
		} else {
			encountered[images[v].Name] = true
			// Append the result slice
			result = append(result, images[v])
		}
	}
	return result
}

func (printer DockerImageList) PrintYaml(printable interface{}, options PrintOptions) error {
	dockerImageList, e := getImageList(printable, options)
	if e != nil {
		return e
	}

	s, err := yaml.Marshal(dockerImageList)
	if err != nil {
		return errcodes.TemplateDisplayError
	}

	_, err = fmt.Fprintf(options.Writer, "%v\n", string(s))
	if err != nil {
		return errcodes.TemplateDisplayError
	}
	return nil
}

func (printer DockerImageList) PrintJson(printable interface{}, options PrintOptions) error {
	dockerImageList, e := getImageList(printable, options)
	if e != nil {
		return e
	}
	s, err := prettyjson.Marshal(dockerImageList)
	if err != nil {
		return errcodes.TemplateDisplayError
	}

	_, err = fmt.Fprintf(options.Writer, "%v\n", string(s))
	if err != nil {
		return errcodes.TemplateDisplayError
	}
	return nil
}

func getImageList(printable interface{}, options PrintOptions) (list DockerImageList, err error) {
	dockerImageList, ok := printable.(DockerImageList)
	if !ok {
		return DockerImageList{}, errcodes.TemplateDisplayError
	}
	return removeUnusedFields(dockerImageList, options), nil
}

func removeUnusedFields(list DockerImageList, options PrintOptions) DockerImageList {
	for i, image := range list.Images {
		tag := getTagOrEmpty(image, options)
		repo := getRepoOrEmpty(image, options)

		list.Images[i] = DockerImage{
			Name:       image.Name,
			Tag:        tag,
			Repository: repo,
		}
	}
	return list
}

func getRepoOrEmpty(image DockerImage, options PrintOptions) string {
	repo := image.Repository
	if !options.IncludeRegistry {
		repo = ""
	}
	return repo
}

func getTagOrEmpty(image DockerImage, options PrintOptions) string {
	tag := image.Tag
	if !options.IncludeTags {
		tag = ""
	}
	return tag
}
