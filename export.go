package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"

	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform"
)

type ExportCmd struct {
	*BuildCmd
	output string
	debug  bool
}

func Export(c *cli.Context) error {
	buildCmd, err := newBuildCmd(c)
	if err != nil {
		return err
	}

	output := strings.TrimRight(c.String("output"), "/")
	if output == "" {
		output = "."
	}

	cmd := &ExportCmd{
		BuildCmd: buildCmd,
		output:   output,
		debug:    c.Bool("debug"),
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

		for _, f := range chain {
			action, err := e.resolveFunc(f)
			if err != nil {
				return err
			}

			log.Printf("export function %s", f.Definition())
			if err := e.exportFunc(f, action); err != nil {
				return fmt.Errorf("build function %s: %w", f.Definition(), err)
			}
		}
	}
	sumFile, err := os.Create(defaultSumFileName)
	if err != nil {
		return fmt.Errorf("can not write to %s: %w", defaultSumFileName, err)
	}
	defer sumFile.Close()
	return e.cache.Flush(sumFile)
}

func (e *ExportCmd) exportFunc(f l.Runnable, action io.Reader) error {
	zipFile, err := ioutil.ReadAll(action)
	if err != nil {
		return err
	}

	name := strings.Replace(string(f.GetLink()), "/", "_", -1)

	return ioutil.WriteFile(fmt.Sprintf("%s/%s.zip", e.output, name), zipFile, 0666)
}
