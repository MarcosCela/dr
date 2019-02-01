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
	return template.Must(template.New("list-repository-images").
		Parse(templates.ImageListPlain)).
		Execute(options.Writer, ImageListTemplate{
			Images:            dockerImageList.Images,
			IncludeRepository: options.IncludeRegistry,
			IncludeTag:        options.IncludeTags,
		})
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
