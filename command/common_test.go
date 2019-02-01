package command

import "testing"
import "github.com/franela/goblin"

func Test(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Test image and tag parsing", func() {
		g.It("Should correctly parse <docker/image:anyTag>", func() {
			image, tag := extractImageAndTag([]string{"docker/image:anyTag"})
			g.Assert(image).Equal("docker/image")
			g.Assert(tag).Equal("anyTag")
		})
		g.It("Should correctly parse <docker/image anyTag>", func() {
			image, tag := extractImageAndTag([]string{"docker/image", "anyTag"})
			g.Assert(image).Equal("docker/image")
			g.Assert(tag).Equal("anyTag")
		})
		g.It("Should correctly parse <docker/image>", func() {
			image, tag := extractImageAndTag([]string{"docker/image"})
			g.Assert(image).Equal("docker/image")
			g.Assert(tag).Equal("latest")
		})
		g.It("Should correctly parse <docker/image:latest>", func() {
			image, tag := extractImageAndTag([]string{"docker/image:latest"})
			g.Assert(image).Equal("docker/image")
			g.Assert(tag).Equal("latest")
		})
	})
}
