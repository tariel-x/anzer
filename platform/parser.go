package platform

import (
	"fmt"
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

	p := parser.New(string(source))
	result := p.Parse()

	fmt.Printf("%#v\n\n", result.Types)

	return nil, nil
}
