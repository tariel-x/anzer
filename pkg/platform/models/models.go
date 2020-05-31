// package models stores types that are using in both platform package and FaaS-specific implementation package
package models

import (
	"io"

	l "github.com/tariel-x/anzer/pkg/lang"
)

type DockerBuildOpts struct {
	Source     io.Reader
	ActionPath string
	Debug      bool
	Env        map[string]*string
}

type BuildOpts struct {
	Debug    bool
	F        l.Runnable
	CommitID string
}

type PublishedFunction struct {
	Name string
	URL  string
}
