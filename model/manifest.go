package model

import (
	"dr/errcodes"
	. "dr/printer"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"gopkg.in/yaml.v2"
)

// ManifestInformation contains information about an image and tag (manifest info for an image)
type ManifestInformation struct {
	Repository string `yaml:"repository,omitempty" json:"repository,omitempty"`
	// Image full name of the image. E.g: busybox
	Image string `yaml:"image" json:"image"`
	// Tag shows the tag of the image. E.g: latest
	Tag string `yaml:"name,omitempty" json:"tag,omitempty"`
	// Digest shows the sha256 digest of the image (last layer)
	Digest string `yaml:"digest,omitempty" json:"digest,omitempty"`
	// Layers number of layers for this image
	Layers int `yaml:"layers,omitempty" json:"layers,omitempty"`
	// CompressedSize is the size of all the layers that make this image. Compressed, so different from
	// the size on disk
	CompressedSize string `yaml:"compressedSize,omitempty" json:"compressedSize,omitempty"`
}

func (printer ManifestInformation) PrintYaml(printable interface{}, options PrintOptions) error {
	manifestInfo, ok := printable.(ManifestInformation)
	if !ok {
		return errcodes.TemplateDisplayError
	}
	s, err := yaml.Marshal(manifestInfo)
	if err != nil {
		return errcodes.TemplateDisplayError
	}

	_, err = fmt.Fprintf(options.Writer, "%v\n", string(s))
	if err != nil {
		return errcodes.TemplateDisplayError
	}
	return nil

}

func (printer ManifestInformation) PrintJson(printable interface{}, options PrintOptions) error {
	manifestInfo, ok := printable.(ManifestInformation)
	if !ok {
		return errcodes.TemplateDisplayError
	}
	s, err := prettyjson.Marshal(manifestInfo)
	if err != nil {
		return errcodes.TemplateDisplayError
	}

	_, err = fmt.Fprintf(options.Writer, "%v\n", string(s))
	if err != nil {
		return errcodes.TemplateDisplayError
	}
	return nil

}
func (printer ManifestInformation) PrintPlain(printable interface{}, options PrintOptions) error {
	return printer.PrintYaml(printable, options)
}
