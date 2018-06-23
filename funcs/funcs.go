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

	_, err := fr.resolveByDefName(rawDef.Name)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (fr *FuncResolver) resolveByDefName(name string) ([]Service, error) {
	def, exists := fr.RawFuncs[name]
	if !exists {
		return nil, fmt.Errorf("No func with name %q", name)
	}

	services := []Service{}

	if fr.isLambda(name) {
		fmt.Printf("%s: New empty func found, create lambda\n", name)
		l, err := fr.createLambda(name)
		if err != nil {
			return nil, err
		}
		services = append(services, *l)
	} else {
		if def.Body.Ref == nil && len(def.Body.ProductEls) == 0 {
			return nil, fmt.Errorf("Ref func body has no name")
		}

		if len(def.Body.ProductEls) > 0 {
			return fr.resolveProduction(*def.Body)
		}

		refDef, exists := fr.RawFuncs[*def.Body.Ref]
		if !exists {
			return nil, fmt.Errorf("No such referenced func %q in raw funcs list", *def.Body.Ref)
		}

		appended, err := fr.resolveByDefName(refDef.Name)
		if err != nil {
			return nil, err
		}

		services = append(services, appended...)
	}

	if def.Body != nil && def.Body.ComposeTo != nil {
		appended, err := fr.resolveNextComposition(*def.Body.ComposeTo)
		if err != nil {
			return nil, err
		}

		services = append(services, appended...)
	}

	return services, nil
}

func (fr *FuncResolver) resolveNextComposition(body listener.FuncBody) ([]Service, error) {
	var err error
	appended := []Service{}
	if body.Ref != nil {
		fmt.Printf("%s: New composing func found\n", *body.Ref)
		appended, err = fr.resolveComposed(*body.Ref)
		if err != nil {
			return nil, err
		}
	} else if len(body.ProductEls) > 0 {
		fmt.Printf("New composing product func found\n")
		appended, err = fr.resolveProduction(body)
		if err != nil {
			return nil, err
		}
	}

	return appended, nil
}

func (fr *FuncResolver) isLambda(name string) bool {
	def, exists := fr.RawFuncs[name]
	if !exists {
		return false
	}
	if def.Body != nil {
		return false
	}
	return true
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

func (fr *FuncResolver) resolveComposed(name string) ([]Service, error) {

	def, exists := fr.RawFuncs[name]
	if !exists {
		return nil, fmt.Errorf("No func with name %q", name)
	}

	if def.Body == nil {
		return nil, fmt.Errorf("Func %q might to be lambda, not composed", name)
	}

	if def.Body.Ref == nil && len(def.Body.ProductEls) == 0 {
		return nil, fmt.Errorf("Ref func body has no name")
	}

	if len(def.Body.ProductEls) > 0 {
		return fr.resolveProduction(*def.Body)
	}

	refDef, exists := fr.RawFuncs[*def.Body.Ref]
	if !exists {
		return nil, fmt.Errorf("No such referenced func %q in raw funcs list", *def.Body.Ref)
	}

	resolved, err := fr.resolveByDefName(refDef.Name)
	if err != nil {
		return nil, err
	}
	return resolved, nil
}

func (fr *FuncResolver) resolveProduction(body listener.FuncBody) ([]Service, error) {
	fmt.Printf("%v: Product func\n", body.Ref)
	services := []Service{}
	names := []string{}
	if body.ProductEls != nil {
		for _, product := range body.ProductEls {
			if product.Ref == nil {
				return nil, fmt.Errorf("One func of product has no reference name")
			}
			productItem, err := fr.resolveByDefName(*product.Ref)
			names = append(names, *product.Ref)
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
