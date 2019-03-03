package build

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform"
	"github.com/urfave/cli"
)

func Build(c *cli.Context) error {
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
		if err := buildFunc(el); err != nil {
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

func buildFunc(f l.F) error {
	dockerGenerator, err := platform.GetDockerGenerator(f.Runtime)
	if err != nil {
		return err
	}
	opts, err := dockerGenerator.GetBuildOptions(f.Link, f.In(), f.Out())
	if err != nil {
		return err
	}

	builder, err := platform.NewBuilder()
	if err != nil {
		return err
	}

	return builder.BuildWithImage(opts, f.Link)
}

func getScheme() (l.Composable, error) {
	loadImages := l.F{
		Link:    "github.com/tariel-x/anzer-example/load_images",
		Runtime: "golang",
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
				Runtime: "golang",
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
				Runtime: "golang",
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
