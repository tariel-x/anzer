package build

import (
	"fmt"

	in "github.com/tariel-x/anzer/internal"
	"github.com/urfave/cli"
)

func Build(c *cli.Context) error {
	compose := in.Alias{
		Name: "c",
		Compose: []in.Composable{
			in.F{
				Link:    "github.com/tariel-x/anzer-examples/a",
				TypeIn:  in.TypeString,
				TypeOut: in.MaxLength(in.TypeString, 10),
			},
			in.F{
				Link:   "github.com/tariel-x/anzer-examples/b",
				TypeIn: in.MaxLength(in.TypeString, 10),
				TypeOut: in.Complex{
					Fields: map[string]in.T{
						"f1": in.TypeInteger,
						"f2": in.TypeString,
					},
				},
			},
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
	return nil
}
