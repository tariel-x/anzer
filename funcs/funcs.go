package funcs

import (
	"fmt"
	"github.com/tariel-x/anzer/listener"
	"github.com/tariel-x/anzer/types"
)

type FuncResolver struct {
	RawFuncs listener.Funcs
	Types    types.Types
	Services Services
}

func Resolve(funcs listener.Funcs, types types.Types) (*SystemGraph, error) {
	fr := FuncResolver{
		RawFuncs: funcs,
		Types:    types,
		Services: Services{},
	}

	return fr.ResolveAll()
}

func (fr *FuncResolver) ResolveAll() (*SystemGraph, error) {
	for name, _ := range fr.RawFuncs {
		err := fr.resolveRaw(name)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (fr *FuncResolver) resolveRaw(name string) error {
	// get raw func from the list and try to simplify it
	rawDef, exists := fr.RawFuncs[name]
	if !exists {
		return fmt.Errorf("No such func %q in raw funcs list", name)
	}

	if rawDef.Def == nil {
		fmt.Printf("New empty func found")
		fr.createLambda(rawDef)
	}

	return nil
}

func (fr *FuncResolver) createLambda(rawDef listener.FuncDef) (*Service, error) {
	inType, err := fr.getType(rawDef.Arg)
	if err != nil {
		return nil, err
	}
	outType, err := fr.getType(rawDef.Ret)
	if err != nil {
		return nil, err
	}

	serviceSet, exists := fr.Services[rawDef.Name]
	if !exists {
		serviceSet = ServiceSet{}
		fr.Services[rawDef.Name] = serviceSet
	}

	num := len(serviceSet)

	s := Service{
		InType:  *inType,
		OutType: *outType,
		Name:    rawDef.Name,
		Index:   num,
		Type:    TypeLambda,
	}

	fr.Services[rawDef.Name][num] = s

	return &s, nil
}

func (fr *FuncResolver) getType(name string) (*types.JsonSchema, error) {
	typeDef, exists := fr.Types[name]
	if !exists {
		return nil, fmt.Errorf("No such type %q in types list")
	}
	return &typeDef, nil
}
