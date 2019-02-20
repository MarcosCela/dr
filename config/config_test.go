package config

import "testing"
import "github.com/franela/goblin"

func TestConfigurationLoadingFromFile(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Check YAML parsing of example test.yaml configuration [valid]", func() {
		config, e := New("resources/test.yaml")

		g.It("No error when parsing", func() {
			g.Assert(e == nil).IsTrue()

		})
		g.It("Config is not nil", func() {
			g.Assert(config != nil).IsTrue()

		})
		g.It("The default context exists and is the first context", func() {
			g.Assert(config.Contexts[0].Name).Equal("default")

		})
		g.It("Current context is set to the default context", func() {
			g.Assert(config.CurrentContext).Equal("default")

		})
		g.It("Default context is trusted", func() {
			g.Assert(config.Contexts[0].Trusted).Equal(true)

		})
		g.It("Default context has valid URL", func() {
			g.Assert(config.Contexts[0].URL).Equal("http://localhost:5000")

		})
		g.It("Default context has empty user", func() {
			g.Assert(config.Contexts[0].User).Equal("")

		})

	})
	g.Describe("Check YAML parsing of path that does not exist", func() {
		config, e := New("invalid/path/does/not/exist/test.yaml")

		g.It("When loading from invalid file, error is not nil", func() {
			g.Assert(e == nil).IsFalse()

		})
		g.It("When loading from invalid file, config is nil", func() {
			g.Assert(config == nil).IsTrue()

		})

	})
}
