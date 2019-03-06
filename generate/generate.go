package generate

import (
	"errors"
	"io/ioutil"

	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform"
	"github.com/urfave/cli"
)

var (
	errOutputUndefined   = errors.New("Output is undefined")
	errFunctionUndefined = errors.New("Function is undefined")
)

func Generate(c *cli.Context) error {
	output := c.String("output")
	if output == "" {
		return errOutputUndefined
	}

	name := c.String("function")
	if name == "" {
		return errFunctionUndefined
	}

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

	if err := compose.Invalid(); err != nil {
		return err
	}

	f := findFunc(name, compose)
	if f == nil {
		return errors.New("function not found")
	}

	cg, err := platform.GetGenerator(f.Runtime)
	if err != nil {
		return err
	}

	generated, err := cg.GenerateFunc(f.In(), f.Out(), f.Link)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(output, []byte(generated), 0666)
}

func findFunc(name string, f l.Composable) *l.F {
	switch ft := f.(type) {
	case l.Alias:
		for _, sf := range ft.Compose {
			if found := findFunc(name, sf); found != nil {
				return found
			}
		}
	case l.F:
		if string(ft.Link) == name {
			return &ft
		}
	}
	return nil
}
