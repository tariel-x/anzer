package main

import (
	"strings"

	"github.com/urfave/cli/v2"
)

func Init(c *cli.Context) error {
	plat, err := getPlatform(c)
	if err != nil {
		return err
	}

	args := map[string]string{}

	num := len(c.Args().Tail())
	for i := 0; i <= num; i++ {
		arg := c.Args().Get(i)
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) < 2 {
			continue
		}
		args[parts[0]] = parts[1]
	}

	return plat.Init(args)
}
