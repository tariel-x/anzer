package funcs

import (
	"fmt"
	_ "strings"

	"github.com/tariel-x/anzer/listener"
	"github.com/tariel-x/anzer/types"
)

type FuncResolver struct {
	RawFuncs     listener.Funcs
	Types        types.Types
	Services     Services
	Dependencies Dependencies
	LastService  *Service
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
	if err := fr.resolveFunc("main"); err != nil {
		return nil, err
	}

	g := SystemGraph{
		Services:     fr.Services,
		Dependencies: fr.Dependencies,
	}
	return &g, nil
}

func (fr *FuncResolver) resolveFunc(name string) error {
	def, exists := fr.RawFuncs[name]
	if !exists {
		return fmt.Errorf("No main func in raw funcs list")
	}
	_, err := fr.resoveDefinition(def, nil)
	return err
}

func (fr *FuncResolver) resoveDefinition(def listener.FuncDef, objS *Service) (*Service, error) {
	if def.Body == nil {
		subjS, err := fr.createLambda(def.Name)
		if err != nil {
			return nil, err
		}
		if objS != nil {
			fr.addDependency(*objS, *subjS)
		}

		return subjS, err
	} else {
		return fr.resolveBody(*def.Body, objS)
	}
}

func (fr *FuncResolver) resolveBody(body listener.FuncBody, objS *Service) (*Service, error) {
	var err error
	if body.ProductEls != nil {
		// if production
		objS, err = fr.resolveProduction(body.ProductEls, objS)
		if err != nil {
			return nil, err
		}
	} else if body.Ref != nil {
		// if reference to another
		if def, exists := fr.RawFuncs[*body.Ref]; exists {
			objS, err = fr.resoveDefinition(def, objS)
			if err != nil {
				return nil, err
			}
		}
	}

	// process next composition
	if body.ComposeTo != nil {
		return fr.resolveBody(*body.ComposeTo, objS)
	}
	return objS, nil
}

func (fr *FuncResolver) resolveProduction(body listener.Production, objS *Service) (*Service, error) {
	services := []Service{}
	prodName := ""
	for _, productBody := range body {
		s, err := fr.resolveBody(productBody, objS)
		if err != nil {
			return nil, err
		}
		services = append(services, *s)
		prodName = prodName + "_" + s.Name
	}

	serviceSet, exists := fr.Services[prodName]
	if !exists {
		serviceSet = ServiceSet{}
		fr.Services[prodName] = serviceSet
	}
	num := len(serviceSet)

	s := Service{
		Name:       prodName,
		Type:       TypeProduction,
		Index:      num,
		UniqueName: fmt.Sprintf("%s.%d", prodName, num),
	}

	fr.Services[prodName][num] = s

	for _, prodService := range services {
		fr.addDependency(prodService, s)
	}

	return &s, nil
}

func (fr *FuncResolver) addDependency(to, from Service) error {
	fr.Dependencies = append(fr.Dependencies, Dependency{
		From: from.UniqueName,
		To:   to.UniqueName,
	})
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
			InType:     inType,
			OutType:    outType,
			Name:       def.Name,
			Index:      num,
			UniqueName: fmt.Sprintf("%s.%d", def.Name, num),
			Type:       TypeLambda,
		}

		fr.Services[def.Name][num] = s

		return &s, nil
	}

	return nil, fmt.Errorf("No func with name %q", name)
}

func (fr *FuncResolver) getType(name string) (*types.JsonSchema, error) {
	if name == "_" {
		return nil, nil
	}
	typeDef, exists := fr.Types[name]
	if !exists {
		return nil, fmt.Errorf("No such type %q in types list", name)
	}
	return &typeDef, nil
}
