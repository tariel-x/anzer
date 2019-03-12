package build

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

func Tst(c *cli.Context) error {
	input := c.String("input")
	if input == "" {
		return fmt.Errorf("no input")
	}
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	_, err = platform.Parse(f)
	return err
}

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

	plat, err := platform.GetPlatform("wsk")
	if err != nil {
		return err
	}

	for _, el := range chain {
		log.Printf("build function %s", el.Definition())
		if err := buildFunc(el, plat); err != nil {
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

func buildFunc(f l.F, plat platform.Platform) error {
	dockerGenerator, err := platform.GetDockerGenerator(f.Runtime)
	if err != nil {
		return err
	}
	opts, err := dockerGenerator.GetBuildOptions(f.Link, f.In(), f.Out(), false)
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

	name := strings.Replace(string(f.Link), "/", "_", -1)
	return plat.Create(action, name, f.Runtime)
}

func getScheme() (l.Composable, error) {
	loadImages := l.F{
		Link:    "github.com/tariel-x/anzer-example/load_images",
		Runtime: "go",
		TypeIn: l.Complex{
			Fields: map[string]l.T{
				"brand":      l.TypeString,
				"model":      l.TypeString,
				"generation": l.TypeString,
				"price":      l.TypeInteger,
				"year":       l.TypeInteger,
				"body":       l.Optional(l.TypeString),
				"rawImages":  l.List(l.TypeString),
				"phone":      l.TypeString,
			},
		},
		TypeOut: l.Complex{
			Fields: map[string]l.T{
				"brand":      l.TypeString,
				"model":      l.TypeString,
				"generation": l.TypeString,
				"price":      l.TypeInteger,
				"year":       l.TypeInteger,
				"body":       l.Optional(l.TypeString),
				"photos":     l.List(l.TypeString),
				"phone":      l.TypeString,
			},
		},
	}

	compose := l.Alias{
		Name: "c",
		Compose: []l.Composable{
			l.F{
				Link:    "github.com/tariel-x/anzer-example/transform",
				Runtime: "go",
				TypeIn: l.Complex{
					Fields: map[string]l.T{
						"brand":  l.TypeString,
						"model":  l.TypeString,
						"price":  l.TypeFloat,
						"year":   l.TypeInteger,
						"body":   l.Optional(l.TypeString),
						"photos": l.List(l.TypeString),
						"phone":  l.TypeString,
					},
				},
				TypeOut: l.Complex{
					Fields: map[string]l.T{
						"brand":      l.TypeString,
						"model":      l.TypeString,
						"generation": l.TypeString,
						"price":      l.TypeInteger,
						"year":       l.TypeInteger,
						"body":       l.Optional(l.TypeString),
						"rawImages":  l.List(l.TypeString),
						"phone":      l.TypeString,
					},
				},
			},
			loadImages,
			l.F{
				Link:    "github.com/tariel-x/anzer-example/create",
				Runtime: "go",
				TypeIn: l.Complex{
					Fields: map[string]l.T{
						"brand":      l.TypeString,
						"model":      l.TypeString,
						"generation": l.TypeString,
						"price":      l.TypeInteger,
						"year":       l.TypeInteger,
						"body":       l.Optional(l.TypeString),
						"photos":     l.List(l.TypeString),
						"phone":      l.TypeString,
					},
				},
				TypeOut: l.Complex{
					Fields: map[string]l.T{
						"id":         l.TypeInteger,
						"brand":      l.TypeString,
						"model":      l.TypeString,
						"generation": l.TypeString,
						"price":      l.TypeInteger,
						"year":       l.TypeInteger,
						"body":       l.Optional(l.TypeString),
						"photos":     l.List(l.TypeString),
						"phone":      l.TypeString,
					},
				},
			},
		},
	}
	return compose, nil
}
