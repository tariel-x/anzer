package build

import (
	"fmt"

	"github.com/tariel-x/anzer/go/generator"
	in "github.com/tariel-x/anzer/internal"
	"github.com/urfave/cli"
)

type codeGenerator interface {
	Generate(inT, outT in.T, packagePath string) (string, error)
	GenerateFunc(inT, outT in.T, packagePath string) (string, error)
}

var (
	cg = generator.GoGenerator{}
)

func Build(c *cli.Context) error {
	bFunc := in.F{
		Link:   "github.com/tariel-x/anzer-examples/b",
		TypeIn: in.MaxLength(in.TypeString, 10),
		TypeOut: in.Complex{
			Fields: map[string]in.T{
				"f1": in.Optional(in.TypeInteger),
				"f2": in.List(in.TypeString),
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
						"f1": in.Optional(in.TypeInteger),
						"f2": in.List(in.TypeString),
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
	output, err := cg.Generate(bFunc.In(), bFunc.Out(), bFunc.Link)
	if err != nil {
		return err
	}
	fmt.Println(output)

	funcOutput, err := cg.GenerateFunc(bFunc.In(), bFunc.Out(), bFunc.Link)
	if err != nil {
		return err
	}
	fmt.Println(funcOutput)
	return nil
}
