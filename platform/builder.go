package platform

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
)

func Build(source io.Reader, name string) error {
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
		Tags:           []string{name},
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
