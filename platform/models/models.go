package models

import "io"

type BuildWithImageOpts struct {
	Source     io.Reader
	ActionPath string
	Debug      bool
}

type PublishedFunction struct {
	Name string
	URL  string
}
