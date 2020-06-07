package main

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"

	l "github.com/tariel-x/anzer/pkg/lang"
	"github.com/tariel-x/anzer/pkg/platform"
)

func Generate(c *cli.Context) error {
	output := c.String("output")
	if output == "" {
		return errOutputUndefined
	}

	name := c.String("function")
	if name == "" {
		return errFunctionUndefined
	}

	input := c.String("input")
	if input == "" {
		return errNoInput
	}
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer f.Close()

	_, composes, err := platform.ParseAll(f)
	if err != nil {
		return err
	}

	for _, compose := range composes {
		if err := compose.Invalid(); err != nil {
			return err
		}
		f := findFunc(name, compose)
		if f == nil {
			continue
		}

		cg, err := platform.GetGenerator(f.Runtime)
		if err != nil {
			return err
		}

		generated, err := cg.GenerateFunc(f.In(), f.Out(), f.Link)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(output, []byte(generated), 0666)
	}

	return errors.New("function not found")
}

func findFunc(name string, f l.Composable) *l.F {
	switch ft := f.(type) {
	case l.Alias:
		for _, sf := range ft.Compose {
			if found := findFunc(name, sf); found != nil {
				return found
			}
		}
	case l.F:
		if string(ft.Link) == name {
			return &ft
		}
	case l.EitherBind:
		if found := findFunc(name, ft.Argument); found != nil {
			return found
		}
	}
	return nil
}
