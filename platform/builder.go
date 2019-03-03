package platform

import (
	"archive/tar"
	"context"
	"github.com/docker/docker/api/types/container"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	"github.com/pkg/errors"
	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/platform/models"
	"github.com/tariel-x/anzer/wsk/golang"
)

var (
	errEmptyOptions = errors.New("empty build options")
)

type DockerGenerator interface {
	GetBuildOptions(link l.FunctionLink, inT l.T, outT l.T) (*models.BuildWithImageOpts, error)
}

func GetDockerGenerator(runtime string) (DockerGenerator, error) {
	switch runtime {
	case "golang":
		p := golang.NewDockerGenerator()
		return p, nil
	default:
		return nil, errUndefinedLanguage
	}
}

type Builder struct {
	cli *client.Client
}

func NewBuilder() (Builder, error) {
	dockerCli, err := client.NewEnvClient()
	if err != nil {
		return Builder{}, err
	}
	return Builder{
		cli: dockerCli,
	}, nil
}

func (b Builder) BuildWithImage(opts *models.BuildWithImageOpts, link l.FunctionLink) error {
	if opts == nil {
		return errEmptyOptions
	}
	name := strings.Replace(string(link), "/", "_", -1)
	tag := name + ":latest"
	tags := []string{tag}
	if err := b.buildImage(opts.Source, tags); err != nil {
		return err
	}
	containerID, err := b.runImage(tag, nil)
	if err != nil {
		return err
	}
	data, err := b.copyFromContainer(containerID, opts.ActionPath)
	if err != nil {
		return err
	}
	reader := tar.NewReader(data)
	for {
		h, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "reading next header from tar")
		}
		f, err := os.Create(name + ".zip")
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.CopyN(f, reader, h.Size)
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "coping contents of tar")
		}
	}
	return nil
}

func (b Builder) buildImage(source io.Reader, tags []string) error {
	args := map[string]*string{}
	options := types.ImageBuildOptions{
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		Tags:           tags,
		Dockerfile:     "Dockerfile",
		BuildArgs:      args,
	}

	buildResponse, err := b.cli.ImageBuild(context.Background(), source, options)
	if err != nil {
		return err
	}
	defer buildResponse.Body.Close()
	termFd, isTerm := term.GetFdInfo(os.Stderr)

	return jsonmessage.DisplayJSONMessagesStream(buildResponse.Body, os.Stderr, termFd, isTerm, nil)
}

func (b Builder) runImage(tag string, cmd []string) (string, error) {
	if cmd == nil {
		cmd = []string{"echo", "hello world"}
	}
	ctx := context.Background()
	resp, err := b.cli.ContainerCreate(ctx, &container.Config{
		Image: tag,
		Cmd:   cmd,
		Tty:   true,
	}, nil, nil, "")

	if err := b.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}

	_, err = b.cli.ContainerWait(ctx, resp.ID)
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

func (b Builder) copyFromContainer(containerID, path string) (io.ReadCloser, error) {
	ctx := context.Background()
	data, _, err := b.cli.CopyFromContainer(ctx, containerID, path)
	return data, err
}
