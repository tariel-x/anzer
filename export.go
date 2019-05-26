package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform"
	"github.com/urfave/cli"
)

func Export(c *cli.Context) error {
	debug := c.Bool("debug")

	input := c.String("input")
	if input == "" {
		return errNoInput
	}

	output := strings.TrimRight(c.String("output"), "/")
	if output == "" {
		output = "."
	}

	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer f.Close()

	composes, err := platform.ParseLazy(f)
	if err != nil {
		return err
	}

	for _, compose := range composes {
		if err := compose.Invalid(); err != nil {
			return err
		}

		chain, err := toChain(compose)
		if err != nil {
			return err
		}

		for _, el := range chain {
			log.Printf("build function %s", el.Definition())
			if err := exportFunc(el, debug, output); err != nil {
				return errors.Wrap(err, fmt.Sprintf("build function %s", el.Definition()))
			}
		}
	}
	return nil
}

func exportFunc(f l.F, debug bool, output string) error {
	dockerGenerator, err := platform.GetDockerGenerator(f.Runtime)
	if err != nil {
		return err
	}
	opts, err := dockerGenerator.GetBuildOptions(f.Link, f.In(), f.Out(), debug)
	if err != nil {
		return err
	}

	builder, err := platform.NewBuilder()
	if err != nil {
		return err
	}

	action, err := builder.BuildWithImage(opts, f.Link)
	if err != nil {
		return err
	}
	zipFile, err := ioutil.ReadAll(action)
	if err != nil {
		return err
	}

	name := strings.Replace(string(f.Link), "/", "_", -1)

	return ioutil.WriteFile(fmt.Sprintf("%s/%s.zip", output, name), zipFile, 0666)
}
