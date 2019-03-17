package main

import (
	"log"
	"os"

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
