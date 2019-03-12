package platform

import (
	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/lang/parser"
	"io"
	"io/ioutil"
)

func Parse(sourceStream io.Reader) ([]l.F, error) {
	source, err := ioutil.ReadAll(sourceStream)
	if err != nil {
		return nil, err
	}

	p := parser.New()
	p.Parse(string(source))

	return nil, nil
}