package main

import (
	"github.com/tariel-x/anzer/platform"
	"github.com/urfave/cli"
	"os"
)

func Validate(c *cli.Context) error {
	input := c.String("input")
	if input == "" {
		return errNoInput
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

		_, err := toChain(compose)
		if err != nil {
			return err
		}
	}
	return nil
}
