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

const (
	defaultCacheLocation = "~/.anzer_cache"
)

type BuildCmd struct {
	input         string
	cacheLocation string
	platform      platform.Platform
}

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

	cacheLocation := defaultCacheLocation
	if c.String("cacheLocation") != "" {
		cacheLocation = c.String("cacheLocation")
	}

	cmd := &BuildCmd{
		input:         input,
		cacheLocation: cacheLocation,
		platform:      plat,
	}
	return cmd.build(c)
}

func (b *BuildCmd) build(c *cli.Context) error {
	f, err := os.Open(b.input)
	if err != nil {
		return err
	}
	defer f.Close()

	composes, err := platform.ParseLazy(f)
	if err != nil {
		return err
	}

	for _, compose := range composes {
		if err := b.buildCompose(compose); err != nil {
			return err
		}
	}
	return nil
}

func (b *BuildCmd) buildCompose(compose l.Composable) error {
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
		component, err := buildFunc(el, b.platform)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("build function %s", el.Definition()))
		}
		components = append(components, component)
	}

	log.Printf("make link for %v", components)
	lnk, err := b.platform.Link(compose.GetName(), components)
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
	//TODO: somewhere here is needed to calculate scheme hash and commit id
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
