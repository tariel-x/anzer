package models

import "io"

type BuildWithImageOpts struct {
	Source     io.Reader
	ActionPath string
}
