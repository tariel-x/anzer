package internal

import (
	"strings"
)

type Composable interface {
	Definition() string
}

type T interface{}

type Alias struct {
	Name    string
	Compose []Composable
}

func (a Alias) Definition() string {
	definitions := []string{}
	l := len(a.Compose) - 1
	for i := range a.Compose {
		definitions = append(definitions, a.Compose[l-i].Definition())
	}
	return strings.Join(definitions, " . ")
}

type F struct {
	Name    string
	Link    string
	TypeIn  T
	TypeOut T
}

func (f F) Definition() string {
	return f.Name
}

type Applied []Composable

func (a Applied) Definition() string {
	definitions := []string{}
	for _, f := range a {
		definitions = append(definitions, f.Definition())
	}
	return strings.Join(definitions, " ")
}

type EitherBind bool

func (e EitherBind) Definition() string {
	return ">>="
}

type EitherReturn bool

func (e EitherReturn) Definition() string {
	return "return"
}
