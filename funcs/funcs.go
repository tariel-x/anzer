package funcs

import (
	"github.com/tariel-x/anzer/listener"
	"github.com/tariel-x/anzer/types"
)

type FuncResolver struct {
	RawFuncs listener.Funcs
	Types    types.Types
}

func Resolve(funcs listener.Funcs, types types.Types) (*SystemGraph, error) {
	fr := FuncResolver{
		RawFuncs: funcs,
		Types:    types,
	}

	return fr.ResolveAll()
}

func (fr *FuncResolver) ResolveAll() (*SystemGraph, error) {
	for name, _ := range fr.RawFuncs {
		fr.resolveRaw(name)
	}

	return nil, nil
}

func (fr *FuncResolver) resolveRaw(name string) {
	// get raw func from the list and try to simplify it
}
