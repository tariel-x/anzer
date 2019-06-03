// package models stores types that are using in both platform package and FaaS-specific implementation package
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
