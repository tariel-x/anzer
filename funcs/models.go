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

type Services map[string]ServiceSet
type ServiceSet map[int]Service
type Dependencies []Dependency

type Service struct {
	InType          *types.JsonSchema `json:"-"`
	OutType         *types.JsonSchema `json:"-"`
	Index           int
	Type            ServiceType
	ProductionTypes []string `json:",omitempty"`
	Name            string
	UniqueName      string
	//Here is temporary element because type checker and code generator are not ready yet
	ProductionOf []Service `json:"-"`
	TypeNameIn   string    `json:"-"`
	TypeNameOut  string    `json:"-"`
	Config       Config
}

type Config struct {
	EnvIn  string
	EnvOut string
	Envs   map[string]string
}

type Dependency struct {
	From string
	To   string
}
