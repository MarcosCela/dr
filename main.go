package main

import (
	"dr/config"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	disableLogging()
	err := setUpApplication().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
func disableLogging() {
	log.SetOutput(ioutil.Discard)
}

// setUpApplication initializes the app context
func setUpApplication() *cli.App {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Marcos Cela Lopez (Github https://github.com/MarcosCela)"
	app.Email = "marcos.cela.lopez@gmail.com"
	app.Usage = "Docker remote CLI utility"
	app.Copyright = License
	app.EnableBashCompletion = true
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound
	app.Metadata = map[string]interface{}{
		config.CONFIG: getConfiguration(),
	}
	return app
}

// getConfigurationOrFail gets the currently default configuration, and if it is not found or is not parseable, fails
func getConfiguration() config.DrConfig {
	cfg, e := config.New()
	if e != nil {
		log.Println("Could not load configuration file! You need a valid configuration file!\nError:\t", e)
		return config.DrConfig{}
	}
	return *cfg
}
