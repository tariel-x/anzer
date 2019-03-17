package platform

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/lang/parser"
)

func Parse(sourceStream io.Reader) ([]lang.F, error) {
	source, err := ioutil.ReadAll(sourceStream)
	if err != nil {
		return nil, err
	}

	p := parser.New(string(source))
	result, err := p.Parse()
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n\n", result)

	return nil, nil
}
