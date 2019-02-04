package build

import (
	"fmt"

	"github.com/tariel-x/anzer/internal"
	"github.com/urfave/cli"
)

func Build(c *cli.Context) error {
	s := internal.Alias{
		Name: "composition1",
		Compose: []internal.Composable{
			internal.F{
				Name:    "f",
				Link:    "github.com/tariel-x/f",
				TypeIn:  internal.TypeString,
				TypeOut: internal.Construct(internal.TypeString, internal.MaxLength, 10),
			},
			internal.F{
				Name: "a",
				Link: "github.com/tariel-x/a",
			},
			internal.F{
				Name: "b",
				Link: "github.com/tariel-x/b",
			},
		},
	}

	m := internal.Alias{
		Name: "compisition2",
		Compose: []internal.Composable{
			internal.F{
				Link: "github.com/tariel-x/abc",
				Name: "abc",
			},
			s,
		},
	}

	if err := m.Invalid(); err != nil {
		return err
	}
	fmt.Println(m.Definition())
	return nil
}
