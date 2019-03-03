package platform

import (
	"context"
	"fmt"
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

type Builder interface {
	GetBuildOptions(link l.FunctionLink, inT l.T, outT l.T) (*models.BuildWithImageOpts, error)
}

func GetBuilder(runtime string) (Builder, error) {
	switch runtime {
	case "golang":
		p := golang.NewBuilder()
		return p, nil
	default:
		return nil, errUndefinedLanguage
	}
}

func Build(opts *models.BuildWithImageOpts, link l.FunctionLink) error {
	if opts == nil {
		return errEmptyOptions
	}
	name := strings.Replace(string(link), "/", "_", -1)
	tags := []string{name + ":latest"}
	if err := buildImage(opts.Source, tags); err != nil {
		return err
	}
	return nil
}

func buildImage(source io.Reader, tags []string) error {
	dockerCli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

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

	buildResponse, err := dockerCli.ImageBuild(context.Background(), source, options)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	defer buildResponse.Body.Close()
	termFd, isTerm := term.GetFdInfo(os.Stderr)

	return jsonmessage.DisplayJSONMessagesStream(buildResponse.Body, os.Stderr, termFd, isTerm, nil)
}
