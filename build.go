package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform"
	"github.com/tariel-x/anzer/platform/models"
	"github.com/urfave/cli"
)

func Build(c *cli.Context) error {
	plat, err := getPlatform(c)
	if err != nil {
		return err
	}
	if err := plat.Connect(); err != nil {
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

	components := make([]models.PublishedFunction, 0, len(chain))
	for _, el := range chain {
		log.Printf("build function %s", el.Definition())
		component, err := buildFunc(el, plat)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("build function %s", el.Definition()))
		}
		components = append(components, component)
	}

	log.Printf("make link for %v", components)
	lnk, err := plat.Link(compose.GetName(), components)
	if err != nil {
		return err
	}
	fmt.Println("Your function name is:", aurora.Cyan(lnk.Name))
	return nil
}

func toChain(f l.Composable) ([]l.Runnable, error) {
	chain := []l.Runnable{}
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
	case l.EitherBind:
		chain = append(chain, ft)
	}
	return chain, nil
}

func buildFunc(f l.Runnable, plat platform.Platform) (models.PublishedFunction, error) {
	dockerGenerator, err := platform.GetDockerGenerator(f.GetRuntime())
	if err != nil {
		return models.PublishedFunction{}, err
	}
	opts, err := dockerGenerator.GetBuildOptions(f, false)
	if err != nil {
		return models.PublishedFunction{}, err
	}
	env, err := getEnv()
	if err != nil {
		return models.PublishedFunction{}, err
	}
	opts.Env = env

	builder, err := platform.NewBuilder()
	if err != nil {
		return models.PublishedFunction{}, err
	}

	action, err := builder.BuildWithImage(opts, f.GetLink())
	if err != nil {
		return models.PublishedFunction{}, err
	}

	name := strings.Replace(string(f.GetLink()), "/", "_", -1)
	function, err := plat.Create(action, name, f.GetRuntime())
	if err != nil {
		return models.PublishedFunction{}, err
	}
	return function, nil
}
