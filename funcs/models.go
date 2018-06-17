package funcs

import "github.com/tariel-x/anzer/types"

type ServiceType string

const (
	TypeLambda     ServiceType = "lambda"
	TypeProduction ServiceType = "production"
)

type SystemGraph struct {
	Services     Services
	Dependencies Dependencies
}

type Services map[string]Service
type Dependencies []Dependency

type Service struct {
	InType  types.JsonSchema
	OutType types.JsonSchema
	Index   int
	Type    ServiceType
	Config struct {
		EnvIn  string
		EnvOut string
		Envs   map[string]string
	}
}

type Dependency struct {
	From Service
	To   Service
}
