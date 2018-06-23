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
	rawDef, exists := fr.RawFuncs["main"]
	if !exists {
		return nil, fmt.Errorf("No such func %q in raw funcs list", "main")
	}

	_, err := fr.resolveRaw(*rawDef.Def)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (fr *FuncResolver) resolveRaw(body listener.FuncBody) ([]Service, error) {
	services := []Service{}

	if fr.isLambda(body) {
		fmt.Printf("%s: New empty func found, create lambda\n", body.Name)
		l, err := fr.createLambda(body)
		if err != nil {
			return nil, err
		}
		services = append(services, *l)
	} else {
		// create function here

		// iterate over child services or product child services and recursive add to graph
	}

	if body.ComposeTo != nil {
		if body.ComposeTo.Name != nil {
			fmt.Printf("%s: New composing func found\n", *body.ComposeTo.Name)
		} else {
			fmt.Printf("New composing func with out name found\n")
		}

		appended, err := fr.resolveCompose(*body.ComposeTo)
		if err != nil {
			return nil, err
		}
		services = append(services, appended...)
	}

	return services, nil
}

func (fr *FuncResolver) isLambda(body listener.FuncBody) bool {
	def, exists := fr.RawFuncs[*body.Name]
	if !exists {
		return false
	}
	if def.Def == nil {
		return true
	}
	return false
}

func (fr *FuncResolver) resolveCompose(body listener.FuncBody) ([]Service, error) {
	if body.Name != nil {
		return fr.resolveRaw(body)
	}
	return nil, nil
}

func (fr *FuncResolver) resolveProduction(body *listener.FuncBody) ([]Service, error) {
	fmt.Printf("%v: Product func\n", body.Name)
	services := []Service{}
	names := []string{}
	if body.ProductEls != nil {
		for _, product := range body.ProductEls {
			productItem, err := fr.resolveRaw(product)
			names = append(names, *product.Name)
			if err != nil {
				return nil, err
			}
			services = append(services, productItem...)
		}
	}
	service, err := fr.createProduction(names)
	if err != nil {
		return nil, err
	}
	services = append(services, *service)
	return services, nil
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

func (fr *FuncResolver) createLambda(body listener.FuncBody) (*Service, error) {
	def, err := fr.getDef(*body.Name)
	if def == nil {
		return nil, err
	}

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

func (fr *FuncResolver) getDef(name string) (*listener.FuncDef, error) {
	def, exists := fr.RawFuncs[name]
	if !exists {
		return nil, fmt.Errorf("No definition with name %q", name)
	}
	return &def, nil
}

func (fr *FuncResolver) getType(name string) (*types.JsonSchema, error) {
	typeDef, exists := fr.Types[name]
	if !exists {
		return nil, fmt.Errorf("No such type %q in types list")
	}
	return &typeDef, nil
}
