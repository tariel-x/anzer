package main

import (
	"log"
	"os"

	"github.com/tariel-x/anzer/build"
	"github.com/tariel-x/anzer/generate"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Anzer CLI tool"
	app.Version = "2.0"
	app.Usage = "generate new functions and build system"

	app.Commands = []cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build anzer project",
			Action:  build.Build,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input, i",
					Usage: "Anzer source file",
				},
			},
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate base for anzer func",
			Action:  generate.Generate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input, i",
					Usage: "Anzer source file",
				},
				cli.StringFlag{
					Name:  "lang, l",
					Usage: "Implementation language",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
