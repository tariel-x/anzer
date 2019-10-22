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
	"github.com/tariel-x/anzer/platform/models"
	"github.com/urfave/cli"
)

type ExportCmd struct {
	*BuildCmd
	output string
	debug  bool
}

func Export(c *cli.Context) error {
	input := c.String("input")
	if input == "" {
		return errNoInput
	}

	output := strings.TrimRight(c.String("output"), "/")
	if output == "" {
		output = "."
	}

	cmd := &ExportCmd{
		BuildCmd: &BuildCmd{
			input:    input,
			platform: nil,
			cache:    nil,
		},
		output: output,
		debug:  c.Bool("debug"),
	}
	return cmd.export()
}

func (e *ExportCmd) export() error {
	f, err := os.Open(e.input)
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

		chain, err := e.toChain(compose)
		if err != nil {
			return err
		}

		for _, el := range chain {
			log.Printf("build function %s", el.Definition())
			if err := e.exportFunc(el); err != nil {
				return errors.Wrap(err, fmt.Sprintf("build function %s", el.Definition()))
			}
		}
	}
	return nil
}

func (e *ExportCmd) exportFunc(f l.Runnable) error {
	dockerGenerator, err := platform.GetDockerGenerator(f.GetRuntime())
	if err != nil {
		return err
	}
	opts, err := dockerGenerator.GetBuildOptions(&models.BuildOpts{
		Debug:    e.debug,
		F:        f,
		CommitID: "",
	})
	if err != nil {
		return err
	}
	env, err := getEnv()
	if err != nil {
		return err
	}
	opts.Env = env

	builder, err := platform.NewBuilder()
	if err != nil {
		return err
	}

	action, err := builder.BuildWithImage(opts, f.GetLink())
	if err != nil {
		return err
	}
	zipFile, err := ioutil.ReadAll(action)
	if err != nil {
		return err
	}

	name := strings.Replace(string(f.GetLink()), "/", "_", -1)

	return ioutil.WriteFile(fmt.Sprintf("%s/%s.zip", e.output, name), zipFile, 0666)
}
