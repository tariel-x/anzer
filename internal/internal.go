package internal

import (
	"fmt"
	"strings"
)

/*
return a = b a
*/

type Composable interface {
	Definition() string
	In() T
	Out() T
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

func (a Alias) In() T {
	if len(a.Compose) == 0 {
		panic(fmt.Errorf("Type %s is undefined", a.Name))
	}
	return a.Compose[0].In()
}

func (a Alias) Out() T {
	if len(a.Compose) == 0 {
		panic(fmt.Errorf("Type %s is undefined", a.Name))
	}
	return a.Compose[len(a.Compose)-1].Out()
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

func (f F) In() T {
	return f.TypeIn
}

func (f F) Out() T {
	return f.TypeOut
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

func (e EitherBind) In() T {
	return nil
}

func (e EitherBind) Out() T {
	return nil
}

type EitherReturn bool

func (e EitherReturn) Definition() string {
	return "return"
}

func (e EitherReturn) In() T {
	return nil
}

func (e EitherReturn) Out() T {
	return nil
}
