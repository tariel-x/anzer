package golang

import (
	"archive/tar"
	"bytes"

	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform/models"
)

const (
	permissionMode int64 = 0666
)

type DockerGenerator struct {
	generator Generator
}

func NewDockerGenerator() DockerGenerator {
	return DockerGenerator{
		generator: NewGenerator(),
	}
}

func (dg DockerGenerator) GetBuildOptions(f l.Runnable, debug bool) (*models.BuildWithImageOpts, error) {
	generated, err := dg.generator.Generate(f)
	if err != nil {
		return nil, err
	}
	dockerfile := dg.generator.GenerateDocker(debug)

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	defer tw.Close()
	// TODO: handle error

	if err := writeTar("main.go", []byte(generated), tw); err != nil {
		return nil, err
	}
	if err := writeTar("Dockerfile", []byte(dockerfile), tw); err != nil {
		return nil, err
	}
	return &models.BuildWithImageOpts{
		Source:     bytes.NewBuffer(buf.Bytes()),
		ActionPath: "/exec/action.zip",
		Debug:      debug,
	}, nil
}

func writeTar(name string, file []byte, tw *tar.Writer) error {
	hdr := &tar.Header{
		Name: name,
		Mode: permissionMode,
		Size: int64(len(file)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}
	if _, err := tw.Write(file); err != nil {
		return err
	}
	return nil
}
