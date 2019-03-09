package export

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/pkg/errors"
	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform"
	"github.com/urfave/cli"
)

func Export(c *cli.Context) error {
	debug := c.Bool("debug")

	compose, err := getScheme()
	if err != nil {
		return err
	}

	if err := compose.Invalid(); err != nil {
		return err
	}

	chain, err := toChain(compose)
	if err != nil {
		return err
	}

	for _, el := range chain {
		log.Printf("build function %s", el.Definition())
		if err := buildFunc(el, debug); err != nil {
			return errors.Wrap(err, fmt.Sprintf("build function %s", el.Definition()))
		}
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

func buildFunc(f l.F, debug bool) error {
	dockerGenerator, err := platform.GetDockerGenerator(f.Runtime)
	if err != nil {
		return err
	}
	opts, err := dockerGenerator.GetBuildOptions(f.Link, f.In(), f.Out(), debug)
	if err != nil {
		return err
	}

	builder, err := platform.NewBuilder()
	if err != nil {
		return err
	}

	action, err := builder.BuildWithImage(opts, f.Link)
	if err != nil {
		return err
	}
	zipFile, err := ioutil.ReadAll(action)
	if err != nil {
		return err
	}

	name := strings.Replace(string(f.Link), "/", "_", -1)

	return ioutil.WriteFile(name+".zip", zipFile, 0666)
}

func getScheme() (l.Composable, error) {
	loadImages := l.F{
		Link:    "github.com/tariel-x/anzer-example/load_images",
		Runtime: "go",
		TypeIn: l.Complex{
			Fields: map[string]l.T{
				"name":        l.TypeString,
				"description": l.TypeString,
				"price":       l.TypeInteger,
				"rawImages":   l.List(l.TypeString),
				"phone":       l.TypeString,
				"year":        l.Optional(l.TypeInteger),
			},
		},
		TypeOut: l.Complex{
			Fields: map[string]l.T{
				"name":        l.TypeString,
				"description": l.TypeString,
				"price":       l.TypeInteger,
				"photos":      l.List(l.TypeString),
				"phone":       l.TypeString,
				"year":        l.Optional(l.TypeInteger),
			},
		},
	}

	compose := l.Alias{
		Name: "c",
		Compose: []l.Composable{
			l.F{
				Link:    "github.com/tariel-x/anzer-example/transform",
				Runtime: "go",
				TypeIn:  l.TypeString,
				TypeOut: l.Complex{
					Fields: map[string]l.T{
						"name":        l.TypeString,
						"description": l.TypeString,
						"price":       l.TypeInteger,
						"rawImages":   l.List(l.TypeString),
						"phone":       l.TypeString,
						"year":        l.Optional(l.TypeInteger),
					},
				},
			},
			loadImages,
			l.F{
				Link:    "github.com/tariel-x/anzer-example/create",
				Runtime: "go",
				TypeIn: l.Complex{
					Fields: map[string]l.T{
						"name":        l.TypeString,
						"description": l.TypeString,
						"price":       l.TypeInteger,
						"photos":      l.List(l.TypeString),
						"phone":       l.TypeString,
						"year":        l.Optional(l.TypeInteger),
					},
				},
				TypeOut: l.TypeInteger,
			},
		},
	}
	return compose, nil
}
