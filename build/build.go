package build

import (
	"fmt"

	"github.com/tariel-x/anzer/internal"
	"github.com/urfave/cli"
)

func Build(c *cli.Context) error {
	compose := internal.Alias{
		Name: "c",
		Compose: []internal.Composable{
			internal.F{
				Link:    "github.com/tariel-x/anzer-examples/a",
				TypeIn:  internal.TypeString,
				TypeOut: internal.Construct(internal.TypeString, internal.TypeMaxLength, []interface{}{10}),
			},
			internal.F{
				Link:   "github.com/tariel-x/anzer-examples/b",
				TypeIn: internal.Construct(internal.TypeString, internal.TypeMaxLength, []interface{}{10}),
				TypeOut: internal.Complex{
					Fields: map[string]internal.T{
						"f1": internal.TypeInteger,
						"f2": internal.TypeString,
					},
				},
			},
			internal.F{
				Link: "github.com/tariel-x/anzer-examples/c",
				TypeIn: internal.Complex{
					Fields: map[string]internal.T{
						"f1": internal.TypeInteger,
						"f2": internal.TypeString,
					},
				},
				TypeOut: internal.TypeBool,
			},
		},
	}

	if err := compose.Invalid(); err != nil {
		return err
	}
	fmt.Println(compose.Definition())
	return nil
}
