package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
	"github.com/tariel-x/anzer/cache"
	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform"
	"github.com/tariel-x/anzer/platform/models"
	"github.com/urfave/cli"
)

const (
	defaultCacheLocation = "~/.anzer_cache"
)

type BuildCmd struct {
	input    string
	platform platform.Platform
	cache    *cache.Manager
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

	f, err := os.Open("anzer.sum")
	if err != nil && os.IsNotExist(err) {
		f = &os.File{}
	} else if err != nil {
		return err
	}
	defer f.Close()
	cm, err := cache.NewManager(f, cacheLocation)
	if err != nil {
		return err
	}

	cmd := &BuildCmd{
		input:    input,
		platform: plat,
		cache:    cm,
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

	chain, err := b.toChain(compose)
	if err != nil {
		return err
	}

	components := make([]models.PublishedFunction, 0, len(chain))
	for _, el := range chain {
		action, err := b.loadCached(el)
		if err != nil {
			log.Printf("build function %s", el.Definition())
			action, err = b.buildFunc(el)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("build function %s", el.Definition()))
			}
		}
		component, err := b.publishFunc(el, action)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("publish function %s", el.Definition()))
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

func (b *BuildCmd) toChain(f l.Composable) ([]l.Runnable, error) {
	chain := []l.Runnable{}
	switch ft := f.(type) {
	case l.Alias:
		for _, sf := range ft.Compose {
			subchain, err := b.toChain(sf)
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

func (b *BuildCmd) loadCached(f l.Runnable) (io.Reader, error) {
	location, _, err := b.cache.FunctionWithCommit(f.GetName(), nil)
	if err != nil {
		return nil, err
	}
	return os.Open(location)
}

func (b *BuildCmd) buildFunc(f l.Runnable) (io.Reader, error) {
	dockerGenerator, err := platform.GetDockerGenerator(f.GetRuntime())
	if err != nil {
		return nil, err
	}
	opts, err := dockerGenerator.GetBuildOptions(f, false)
	if err != nil {
		return nil, err
	}
	env, err := getEnv()
	if err != nil {
		return nil, err
	}
	opts.Env = env

	builder, err := platform.NewBuilder()
	if err != nil {
		return nil, err
	}

	return builder.BuildWithImage(opts, f.GetLink())
}

func (b *BuildCmd) publishFunc(f l.Runnable, action io.Reader) (models.PublishedFunction, error) {
	name := strings.Replace(string(f.GetLink()), "/", "_", -1)
	function, err := b.platform.Create(action, name, f.GetRuntime())
	if err != nil {
		return models.PublishedFunction{}, err
	}
	return function, nil
}
