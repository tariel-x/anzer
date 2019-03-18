package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/tariel-x/anzer/platform"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Anzer CLI tool"
	app.Version = "2.0"
	app.Usage = "generate new functions and build system"

	inputFlag := cli.StringFlag{
		Name:  "input, i",
		Usage: "Anzer source file",
	}
	platformFlag := cli.StringFlag{
		Name:  "platform, p",
		Usage: "Target FaaS platform",
	}
	debugFlag := cli.StringFlag{
		Name:  "debug, d",
		Usage: "Debug mode, specific for platform and language",
	}
	outputFlag := cli.StringFlag{
		Name:  "output, o",
		Usage: "Output for generated files",
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init target platform settings",
			Action:  Build,
			Flags: []cli.Flag{
				platformFlag,
				cli.StringFlag{
					Name:  "arg, a",
					Usage: "Platform-specific settings string",
				},
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build anzer project",
			Action:  Build,
			Flags: []cli.Flag{
				inputFlag,
				platformFlag,
			},
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate base for anzer func",
			Action:  Generate,
			Flags: []cli.Flag{
				inputFlag,
				outputFlag,
				cli.StringFlag{
					Name:  "function, f",
					Usage: "Function to generate",
				},
			},
		},
		{
			Name:    "export",
			Aliases: []string{"e"},
			Usage:   "build functions and export to zip files",
			Action:  Export,
			Flags: []cli.Flag{
				inputFlag,
				outputFlag,
				debugFlag,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getPlatform(c *cli.Context) (platform.Platform, error) {
	platName := c.String("platform")
	if platName == "" {
		return nil, fmt.Errorf("no platform")
	}
	plat, err := platform.GetPlatform(platName)
	if err != nil {
		return nil, err
	}
	return plat, nil
}

func getInput(c *cli.Context) (io.Reader, error) {
	input := c.String("input")
	if input == "" {
		return nil, fmt.Errorf("no input")
	}
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	return f, nil
}
