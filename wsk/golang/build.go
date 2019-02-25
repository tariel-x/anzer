package golang

import (
	"io/ioutil"
	"strings"

	l "github.com/tariel-x/anzer/lang"
)

type Builder struct {
	generator Generator
}

func NewBuilder() Builder {
	return Builder{
		generator: NewGenerator(),
	}
}

func (b Builder) Build(link string, inT l.T, outT l.T) error {
	directory := strings.Replace(link, "/", "_", -1)
	output := "/tmp/" + directory
	execPath := output + "/exec"
	dockerfilePath := output + "/Dockerfile"

	generated, err := b.generator.Generate(inT, outT, link)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(execPath, []byte(generated), 0666); err != nil {
		return err
	}

	if err := ioutil.WriteFile(dockerfilePath, []byte(b.generator.GenerateDocker()), 0666); err != nil {
		return err
	}

	return nil
}
