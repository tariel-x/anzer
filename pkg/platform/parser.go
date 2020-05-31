package platform

import (
	"io"
	"io/ioutil"

	"github.com/tariel-x/anzer/pkg/lang"
	"github.com/tariel-x/anzer/pkg/lang/parser"
)

func ParseLazy(sourceStream io.Reader) ([]lang.Composable, error) {
	source, err := ioutil.ReadAll(sourceStream)
	if err != nil {
		return nil, err
	}

	p := parser.New(string(source))
	return p.ParseLazy()
}

func ParseAll(sourceStream io.Reader) ([]lang.Composable, error) {
	source, err := ioutil.ReadAll(sourceStream)
	if err != nil {
		return nil, err
	}

	p := parser.New(string(source))
	return p.ParseAll()
}
