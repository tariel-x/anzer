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

type Builder struct {
	generator Generator
}

func NewBuilder() Builder {
	return Builder{
		generator: NewGenerator(),
	}
}

func (b Builder) GetBuildOptions(link l.FunctionLink, inT l.T, outT l.T) (*models.BuildWithImageOpts, error) {
	generated, err := b.generator.Generate(inT, outT, link)
	if err != nil {
		return nil, err
	}
	makefile := b.generator.GenerateMakefile()

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	defer tw.Close()
	// TODO: handler error

	if err := writeTar("main.go", []byte(generated), tw); err != nil {
		return nil, err
	}
	if err := writeTar("Makefile", []byte(makefile), tw); err != nil {
		return nil, err
	}
	return &models.BuildWithImageOpts{
		Source:     bytes.NewBuffer(buf.Bytes()),
		ActionPath: "",
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
