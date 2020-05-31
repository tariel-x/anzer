package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"

	"github.com/tariel-x/anzer/pkg/platform"
)

var (
	errOutputUndefined   = errors.New("output is undefined")
	errFunctionUndefined = errors.New("function is undefined")
	errNoInput           = errors.New("no input")
	errNoOutput          = errors.New("no output")
)

func main() {
	app := cli.NewApp()
	app.Name = "Anzer"
	app.Version = "1.0"
	app.Usage = "generate new functions and build typesafe serverless system"

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
	cacheFlag := cli.StringFlag{
		Name:  "cacheLocation, cl",
		Usage: "Location of the build cache",
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init target platform settings",
			Action:  Init,
			Flags: []cli.Flag{
				platformFlag,
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
				cacheFlag,
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
				platformFlag,
			},
		},
		{
			Name:    "validate",
			Aliases: []string{"v"},
			Usage:   "validate anzer file",
			Action:  Validate,
			Flags: []cli.Flag{
				inputFlag,
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
		return nil, errors.New("no platform")
	}
	plat, err := platform.GetPlatform(platName)
	if err != nil {
		return nil, err
	}
	return plat, nil
}

func getEnv() (map[string]*string, error) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return nil, nil
	}

	readedEnvs, err := godotenv.Read()
	if err != nil {
		return nil, err
	}
	env := map[string]*string{}
	for k, v := range readedEnvs {
		env[k] = &v
	}
	return env, nil
}
