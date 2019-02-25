package build

import (
	"fmt"

	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform"
	"github.com/urfave/cli"
)

func Build(c *cli.Context) error {
	cg, err := platform.GetGenerator(c.String("lang"))
	if err != nil {
		return err
	}

	bFunc := l.F{
		Link:   "github.com/tariel-x/anzer-examples/b",
		TypeIn: l.MaxLength(l.TypeString, 10),
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
				TypeIn:  l.TypeString,
				TypeOut: l.MaxLength(l.TypeString, 10),
			},
			bFunc,
			l.F{
				Link: "github.com/tariel-x/anzer-examples/c",
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

	if err := compose.Invalid(); err != nil {
		return err
	}

	output, err := cg.Generate(bFunc.In(), bFunc.Out(), bFunc.Link)
	if err != nil {
		return err
	}
	fmt.Println(output)

	return nil
}
