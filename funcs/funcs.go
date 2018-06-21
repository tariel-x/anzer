package funcs

import (
	"fmt"
	"strings"

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
	_, err := fr.resolveRaw("main")
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (fr *FuncResolver) resolveRaw(name string) ([]Service, error) {
	rawDef, exists := fr.RawFuncs[name]
	if !exists {
		return nil, fmt.Errorf("No such func %q in raw funcs list", name)
	}

	services := []Service{}

	if rawDef.Def == nil {
		fmt.Printf("%s: New empty func found, create lambda\n", name)
		l, err := fr.createLambda(rawDef)
		if err != nil {
			return nil, err
		}
		services = append(services, *l)
	} else {
		// create function here

		// iterate over child services or product child services and recursive add to graph
		if rawDef.Def.ComposeTo != nil {
			fmt.Printf("%s: New composing func found\n", name)
			appended, err := fr.resolveCompose(rawDef.Def.ComposeTo)
			if err != nil {
				return nil, err
			}
			services = append(services, appended...)
		}
	}

	return services, nil
}

func (fr *FuncResolver) resolveCompose(fb *listener.FuncBody) ([]Service, error) {
	if fb.Name != nil {
		return fr.resolveRaw(*fb.Name)
	}
	return nil, nil
}

func (fr *FuncResolver) resolveProduction(fb *listener.FuncBody) ([]Service, error) {
	fmt.Printf("%v: Product func\n", fb.Name)
	product := []Service{}
	prodNames := []string{}
	if fb.ProductEls != nil {
		for _, pFunc := range fb.ProductEls {
			productItem, err := fr.resolveRaw(*pFunc.Name)
			prodNames = append(prodNames, *pFunc.Name)
			if err != nil {
				return nil, err
			}
			product = append(product, productItem...)
		}
	}
	prodService, err := fr.createProduction(prodNames)
	if err != nil {
		return nil, err
	}
	product = append(product, *prodService)
	return product, nil
}

func (fr *FuncResolver) createProduction(names []string) (*Service, error) {
	name := strings.Join(names, "_")
	s := Service{
		Name:  name,
		Index: 0,
		Type:  TypeProduction,
	}
	return &s, nil
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
