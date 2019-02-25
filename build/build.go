package build

import (
	"io/ioutil"
	"strings"

	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform"
	"github.com/urfave/cli"
)

func Build(c *cli.Context) error {
	bFunc := l.F{
		Link:    "github.com/tariel-x/anzer-examples/b",
		Runtime: "go",
		TypeIn:  l.MaxLength(l.TypeString, 10),
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

	cg, err := platform.GetGenerator(bFunc.Runtime)
	if err != nil {
		return err
	}

	directory := strings.Replace(bFunc.Link, "/", "_", -1)
	output := "/tmp/" + directory
	execPath := output + "/exec"
	dockerfilePath := output + "/Dockerfile"

	generated, err := cg.Generate(bFunc.In(), bFunc.Out(), bFunc.Link)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(execPath, []byte(generated), 0666); err != nil {
		return err
	}

	// TODO: write Dockercompose here and build function archive with docker
	if err := ioutil.WriteFile(dockerfilePath, []byte(cg.GenerateDocker()), 0666); err != nil {
		return err
	}

	return nil
}
