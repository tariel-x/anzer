package build

import (
	"fmt"

	"github.com/tariel-x/anzer/generator"
	in "github.com/tariel-x/anzer/internal"
	"github.com/urfave/cli"
)

func Build(c *cli.Context) error {
	bFunc := in.F{
		Link:   "github.com/tariel-x/anzer-examples/b",
		TypeIn: in.MaxLength(in.TypeString, 10),
		TypeOut: in.Complex{
			Fields: map[string]in.T{
				"f1": in.TypeInteger,
				"f2": in.TypeString,
			},
		},
	}

	compose := in.Alias{
		Name: "c",
		Compose: []in.Composable{
			in.F{
				Link:    "github.com/tariel-x/anzer-examples/a",
				TypeIn:  in.TypeString,
				TypeOut: in.MaxLength(in.TypeString, 10),
			},
			bFunc,
			in.F{
				Link: "github.com/tariel-x/anzer-examples/c",
				TypeIn: in.Complex{
					Fields: map[string]in.T{
						"f1": in.TypeInteger,
						"f2": in.TypeString,
					},
				},
				TypeOut: in.TypeBool,
			},
		},
	}

	if err := compose.Invalid(); err != nil {
		return err
	}
	fmt.Println(compose.Definition())
	output, err := generator.Generate(bFunc.In(), bFunc.Out(), bFunc.Link)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
