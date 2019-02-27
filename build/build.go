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
	builder, err := platform.GetBuilder(f.Runtime)
	if err != nil {
		return err
	}
	dockersource, err := builder.Build(f.Link, f.In(), f.Out())
	if err != nil {
		return err
	}
	return platform.Build(dockersource)
}

func getScheme() (l.Composable, error) {
	bFunc := l.F{
		Link:    "github.com/tariel-x/anzer-examples/b",
		Runtime: "golang",
		TypeIn:  l.MaxLength(l.TypeString, 10),
		TypeOut: l.Complex{
			Fields: map[string]l.T{
				"f1": l.Optional(l.TypeInteger),
				"f2": l.List(l.TypeString),
			},
		},
	}

	compose := l.Alias{
		Name: "c",
		Compose: []l.Composable{
			l.F{
				Link:    "github.com/tariel-x/anzer-examples/a",
				Runtime: "golang",
				TypeIn:  l.TypeString,
				TypeOut: l.MaxLength(l.TypeString, 10),
			},
			bFunc,
			l.F{
				Link:    "github.com/tariel-x/anzer-examples/c",
				Runtime: "golang",
				TypeIn: l.Complex{
					Fields: map[string]l.T{
						"f1": l.Optional(l.TypeInteger),
						"f2": l.List(l.TypeString),
					},
				},
				TypeOut: l.TypeBool,
			},
		},
	}
	return compose, nil
}
