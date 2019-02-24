package generate

import (
	"errors"
	"fmt"

	"github.com/tariel-x/anzer/go/generator"
	l "github.com/tariel-x/anzer/lang"
	"github.com/urfave/cli"
)

type codeGenerator interface {
	GenerateFunc(inT, outT l.T, packagePath string) (string, error)
}

func Generate(c *cli.Context) error {
	cg, err := getGenerator(c)
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
