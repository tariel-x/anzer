package platform

import (
	"io"
	"io/ioutil"

	"github.com/tariel-x/anzer/pkg/lang"
	"github.com/tariel-x/anzer/pkg/lang/parser"
)

func ParseLazy(sourceStream io.Reader) (string, []lang.Composable, error) {
	source, err := ioutil.ReadAll(sourceStream)
	if err != nil {
		return "", nil, err
	}
	p := parser.New(string(source))

	pkg, err := p.GetPackage()
	if err != nil {
		return "", nil, err
	}

	composes, err := p.ParseLazy()
	if err != nil {
		return "", nil, err
	}
	return pkg, composes, nil
}

func ParseAll(sourceStream io.Reader) (string, []lang.Composable, error) {
	source, err := ioutil.ReadAll(sourceStream)
	if err != nil {
		return "", nil, err
	}

	p := parser.New(string(source))
	pkg, err := p.GetPackage()
	if err != nil {
		return "", nil, err
	}

	composes, err := p.ParseLazy()
	if err != nil {
		return "", nil, err
	}
	return pkg, composes, nil
}
