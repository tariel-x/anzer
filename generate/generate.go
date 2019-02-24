package generate

import (
	"errors"
	"fmt"

	"github.com/tariel-x/anzer/go/generator"
	in "github.com/tariel-x/anzer/internal"
	"github.com/urfave/cli"
)

type codeGenerator interface {
	GenerateFunc(inT, outT in.T, packagePath string) (string, error)
}

func Generate(c *cli.Context) error {
	cg, err := getGenerator(c)
	if err != nil {
		return err
	}

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

	funcOutput, err := cg.GenerateFunc(bFunc.In(), bFunc.Out(), bFunc.Link)
	if err != nil {
		return err
	}
	fmt.Println(funcOutput)
	return nil
}

func getGenerator(c *cli.Context) (codeGenerator, error) {
	switch c.String("lang") {
	case "go":
		return generator.CodeGenerator{}, nil
	default:
		return nil, errors.New("undefined language")
	}
}
