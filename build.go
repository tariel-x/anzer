package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform"
	"github.com/urfave/cli"
)

func Build(c *cli.Context) error {
	plat, err := getPlatform(c)
	if err != nil {
		return err
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

	composes, err := platform.ParseLazy(f)
	if err != nil {
		return err
	}

	for _, compose := range composes {
		if err := buildCompose(compose, plat); err != nil {
			return err
		}
	}
	return nil
}

func buildCompose(compose l.Composable, plat platform.Platform) error {
	log.Printf("build composition %s = %s", compose.GetName(), compose.Definition())

	if err := compose.Invalid(); err != nil {
		return err
	}

	chain, err := toChain(compose)
	if err != nil {
		return err
	}

	names := make([]string, 0, len(chain))
	for _, el := range chain {
		log.Printf("build function %s", el.Definition())
		name, err := buildFunc(el, plat)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("build function %s", el.Definition()))
		}
		names = append(names, name)
	}

	log.Printf("make link for %v", names)
	_, err = plat.Link(compose.GetName(), names)
	if err != nil {
		return err
	}

	return nil
}

func toChain(f l.Composable) ([]l.F, error) {
	chain := []l.F{}
	switch ft := f.(type) {
	case l.Alias:
		for _, sf := range ft.Compose {
			subchain, err := toChain(sf)
			if err != nil {
				return nil, err
			}
			chain = append(chain, subchain...)
		}
	case l.F:
		chain = append(chain, ft)
	}
	return chain, nil
}

func buildFunc(f l.F, plat platform.Platform) (string, error) {
	dockerGenerator, err := platform.GetDockerGenerator(f.Runtime)
	if err != nil {
		return "", err
	}
	opts, err := dockerGenerator.GetBuildOptions(f.Link, f.In(), f.Out(), false)
	if err != nil {
		return "", err
	}

	builder, err := platform.NewBuilder()
	if err != nil {
		return "", err
	}

	action, err := builder.BuildWithImage(opts, f.Link)
	if err != nil {
		return "", err
	}

	name := strings.Replace(string(f.Link), "/", "_", -1)
	function, err := plat.Create(action, name, f.Runtime)
	if err != nil {
		return "", err
	}
	return function.Name, nil
}
