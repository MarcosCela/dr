package model

import (
	"dr/config"
	"dr/errcodes"
	. "dr/printer"
	"fmt"
	"github.com/appscode/docker-registry-client/registry"
	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/urfave/cli"
	"strings"
)

type CustomClient interface {
	Repositories() (imageList DockerImageList, error error)
	Manifest(repository string, tag string) (manifest ManifestInformation, error error)
	RegistryUrl() (registry string)
	Tags(imageRepository string) (tagList TagList, error error)
}

func NewCustomClient(registry *registry.Registry) CustomClient {
	return &defaultCustomClient{registry}
}

type defaultCustomClient struct {
	registry *registry.Registry
}

func (client defaultCustomClient) RegistryUrl() (registry string) {
	return strings.TrimPrefix(strings.TrimPrefix(client.registry.URL, "http://"), "https://")
}

// Get the list of repositories, and generate models
func (client defaultCustomClient) Repositories() (imageList DockerImageList, error error) {
	repos, e := client.registry.Repositories()
	if e != nil {
		return DockerImageList{}, cli.NewExitError("Could not list repositories. Error: "+e.Error(), errcodes.CannotListRepo)
	}
	var images []DockerImage
	for _, repo := range repos {
		tags, e := client.registry.Tags(repo)
		if e != nil {
			continue
		}
		for _, tag := range tags {
			images = append(images, DockerImage{
				Repository: client.RegistryUrl(),
				Name:       repo,
				Tag:        tag,
			})
		}
	}
	return DockerImageList{Images: images}, nil
}

func (client defaultCustomClient) Manifest(image string, tag string) (manifest ManifestInformation, error error) {
	imageManifest, e := client.registry.ManifestV2(image, tag)
	if e != nil {
		return
	}
	if e != nil {
		return ManifestInformation{}, cli.NewExitError("Error when getting the manifest: "+e.Error(), errcodes.CannotRetrieveManifest)
	}
	return WrapManifestInfo(client.RegistryUrl(), image, tag, imageManifest), nil
}

func WrapManifestInfo(repository string, image string, tag string, imageManifest *schema2.DeserializedManifest) ManifestInformation {
	manifestInfo := ManifestInformation{
		Repository:     repository,
		Image:          image,
		Tag:            tag,
		Digest:         imageManifest.Manifest.Config.Digest.String(),
		Layers:         len(imageManifest.Layers),
		CompressedSize: byteSizeToHumanReadable(calculateFullSize(append(imageManifest.Layers, imageManifest.Config.Descriptor()))),
	}
	return manifestInfo
}
func byteSizeToHumanReadable(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
func calculateFullSize(layers []distribution.Descriptor) int64 {
	totalSize := int64(0)
	for _, layer := range layers {
		totalSize += layer.Size
	}
	return totalSize
}
func printByFormat(outputFormat string,
	printable Printable, options PrintOptions) error {
	switch outputFormat {
	case config.PLAINOutputFormat:
		return printable.PrintPlain(printable, options)
	case config.JSONOutputFormat:
		return printable.PrintJson(printable, options)
	case config.YAMLOutputFormat:
		return printable.PrintYaml(printable, options)
	default:
		return printable.PrintJson(printable, options)
	}
}

func GetCustomClient(c *cli.Context) CustomClient {
	return NewCustomClient(config.GetClient(c))
}

func Print(c *cli.Context, printable Printable) error {
	return printByFormat(getConfiguredOutputFormat(c), printable, GetPrintOptions(c))
}

func getConfiguredOutputFormat(c *cli.Context) string {
	// Default value
	outputFormat := config.JSONOutputFormat
	// Acquire value from configuration
	outputFormat = config.GetCurrentConfiguration(c).OutputFormat
	// If flag is set, acquire value from flag and overwrite
	if c.String(config.OutputFormatFlagName) != "" {
		outputFormat = c.String(config.OutputFormatFlagName)
	}
	return outputFormat
}
func (client defaultCustomClient) Tags(imageRepository string) (tagList TagList, error error) {
	tags, e := client.registry.Tags(imageRepository)
	if e != nil {
		return TagList{}, cli.NewExitError("Could not list tags. Error: "+e.Error(), errcodes.CannotListTags)
	}

	for _, tag := range tags {
		tagList.Tags = append(tagList.Tags, tag)
	}
	return tagList, nil
}
