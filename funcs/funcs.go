package funcs

import (
	"fmt"
	_ "strings"

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
	err := fr.resolveFunc("main")
	return nil, err
}

func (fr *FuncResolver) resolveFunc(name string) error {
	def, exists := fr.RawFuncs[name]
	if !exists {
		return fmt.Errorf("No main func in raw funcs list")
	}
	return fr.resoveDefinition(def)
}

func (fr *FuncResolver) resoveDefinition(def listener.FuncDef) error {
	if def.Body == nil {
		fr.createLambda(def.Name)
	}
	return nil
}

func (fr *FuncResolver) resolveBody(body listener.FuncBody) error {
	return nil
}

func (fr *FuncResolver) createLambda(name string) (*Service, error) {
	if def, exists := fr.RawFuncs[name]; exists {
		inType, err := fr.getType(def.Arg)
		if err != nil {
			return nil, err
		}
		outType, err := fr.getType(def.Ret)
		if err != nil {
			return nil, err
		}

		serviceSet, exists := fr.Services[def.Name]
		if !exists {
			serviceSet = ServiceSet{}
			fr.Services[def.Name] = serviceSet
		}

		num := len(serviceSet)

		s := Service{
			InType:  *inType,
			OutType: *outType,
			Name:    def.Name,
			Index:   num,
			Type:    TypeLambda,
		}

		fr.Services[def.Name][num] = s

		return &s, nil
	}

	return nil, fmt.Errorf("No func with name %q", name)
}

func (fr *FuncResolver) getType(name string) (*types.JsonSchema, error) {
	typeDef, exists := fr.Types[name]
	if !exists {
		return nil, fmt.Errorf("No such type %q in types list")
	}
	return &typeDef, nil
}
