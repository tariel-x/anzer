package main

import (
	"os"

	"github.com/tariel-x/anzer/pkg/platform"
	"github.com/urfave/cli"
)

type ValidateCmd struct {
	*BuildCmd
}

func Validate(c *cli.Context) error {
	input := c.String("input")
	if input == "" {
		return errNoInput
	}
	cmd := &ValidateCmd{
		BuildCmd: &BuildCmd{
			input: input,
		},
	}
	return cmd.validate()
}

func (v *ValidateCmd) validate() error {
	f, err := os.Open(v.input)
	if err != nil {
		return err
	}
	defer f.Close()

	_, composes, err := platform.ParseLazy(f)
	if err != nil {
		return err
	}

	for _, compose := range composes {
		if err := compose.Invalid(); err != nil {
			return err
		}

		_, err := v.toChain(compose)
		if err != nil {
			return err
		}
	}
	return nil
}
