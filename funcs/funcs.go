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
		return fmt.Errorf("No func %q in raw funcs list", name)
	}
	_, err := fr.resoveDefinition(def, nil)
	return err
}

func (fr *FuncResolver) resoveDefinition(def listener.FuncDef, toS *Service) (*Service, error) {
	if def.Body == nil {
		fromS, err := fr.createLambda(def.Name)
		if err != nil {
			return nil, err
		}
		return fromS, err
	} else {
		return fr.resolveBody(*def.Body, toS)
	}
}

func (fr *FuncResolver) resolveBody(body listener.FuncBody, toS *Service) (*Service, error) {
	var err error
	if body.ProductEls != nil {
		// if production
		toS, err = fr.resolveProduction(body.ProductEls, toS)
		if err != nil {
			return nil, err
		}
	} else if body.Ref != nil {
		// if reference to another
		if def, exists := fr.RawFuncs[*body.Ref]; exists {
			toS, err = fr.resoveDefinition(def, toS)
			if err != nil {
				return nil, err
			}
		}
	}

	// process next composition
	if body.ComposeTo != nil {
		fromS, err := fr.resolveBody(*body.ComposeTo, toS)
		if err != nil {
			return nil, err
		}
		if toS != nil {
			fr.addDependency(*fromS, *toS)
		}
	}
	return toS, nil
}

func (fr *FuncResolver) resolveProduction(body listener.Production, toS *Service) (*Service, error) {
	services := []Service{}
	prodName := ""
	for _, productBody := range body {
		s, err := fr.resolveBody(productBody, toS)
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
		Name:         prodName,
		Type:         TypeProduction,
		Index:        num,
		UniqueName:   fmt.Sprintf("%s.%d", prodName, num),
		ProductionOf: services,
	}

	fr.Services[prodName][num] = s

	for _, prodService := range services {
		fr.Dependencies = append(fr.Dependencies, Dependency{
			From: prodService.UniqueName,
			To:   s.UniqueName,
		})
	}

	return &s, nil
}

func (fr *FuncResolver) addDependency(fromS, toS Service) error {

	if toS.Type == TypeLambda {
		fr.Dependencies = append(fr.Dependencies, Dependency{
			From: fromS.UniqueName,
			To:   toS.UniqueName,
		})
	} else if toS.Type == TypeProduction {
		for _, prodService := range toS.ProductionOf {
			fr.Dependencies = append(fr.Dependencies, Dependency{
				From: fromS.UniqueName,
				To:   prodService.UniqueName,
			})
		}
	}

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
			Config:     *def.Params,
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
