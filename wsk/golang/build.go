package golang

import (
	"archive/tar"
	"bytes"
	"io"

	l "github.com/tariel-x/anzer/lang"
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

func (b Builder) Build(link string, inT l.T, outT l.T) (io.Reader, error) {
	generated, err := b.generator.Generate(inT, outT, link)
	if err != nil {
		return nil, err
	}
	dockercontent := b.generator.GenerateDocker()

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	defer tw.Close()
	// TODO: handler error

	if err := writeTar("exec", []byte(generated), tw); err != nil {
		return nil, err
	}
	if err := writeTar("Dockerfile", []byte(dockercontent), tw); err != nil {
		return nil, err
	}
	return bytes.NewBuffer(buf.Bytes()), nil
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
