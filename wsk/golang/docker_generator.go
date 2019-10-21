package golang

import (
	"archive/tar"
	"bytes"

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

func (dg DockerGenerator) GetBuildOptions(opts *models.BuildOpts) (*models.DockerBuildOpts, error) {
	generated, err := dg.generator.Generate(opts.F)
	if err != nil {
		return nil, err
	}
	dockerfile, err := dg.generator.GenerateDocker(opts)
	if err != nil {
		return nil, err
	}

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
	return &models.DockerBuildOpts{
		Source:     bytes.NewBuffer(buf.Bytes()),
		ActionPath: "/exec/action.zip",
		Debug:      opts.Debug,
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
