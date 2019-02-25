package generate

import (
	"errors"
	"io/ioutil"

	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform"
	"github.com/urfave/cli"
)

var (
	errOutputUndefined = errors.New("Output is undefined")
)

func Generate(c *cli.Context) error {
	cg, err := platform.GetGenerator(c.String("lang"))
	if err != nil {
		return err
	}

	output := c.String("output")
	if output == "" {
		return errOutputUndefined
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

	generated, err := cg.GenerateFunc(bFunc.In(), bFunc.Out(), bFunc.Link)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(output, []byte(generated), 0666)
}
