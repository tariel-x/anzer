package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func Init(c *cli.Context) error {
	plat, err := getPlatform(c)
	if err != nil {
		return err
	}

	arg := c.String("arg")
	if arg == "" {
		return fmt.Errorf("no argument")
	}

	return plat.Init(arg)
}
