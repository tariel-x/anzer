package internal

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errTypeInconsistent = errors.New("types are not consistent")
)

type Composable interface {
	Definition() string
	In() T
	Out() T
	Invalid() error
}

type Alias struct {
	Name    string
	Compose []Composable
}

func (a Alias) Definition() string {
	definitions := []string{}
	for _, c := range a.Compose {
		definitions = append(definitions, c.Definition())
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

func (a Alias) Invalid() error {
	if len(a.Compose) <= 1 {
		return nil
	}
	last := a.Compose[0]
	for _, c := range a.Compose[1:] {
		if !c.In().Equal(last.Out()) && !c.In().Parent(last.Out()) {
			return errTypeInconsistent
		}
		last = c
	}
	return nil
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

func (f F) Invalid() error {
	return nil
}

type Applied []Composable

func (a Applied) Definition() string {
	definitions := []string{}
	for _, f := range a {
		definitions = append(definitions, f.Definition())
	}
	return strings.Join(definitions, " ")
}

func (a Applied) In() T {
	return nil
}

func (a Applied) Out() T {
	return nil
}

func (a Applied) Invalid() error {
	return nil
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

func (e EitherBind) Invalid() error {
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

func (e EitherReturn) Invalid() error {
	return nil
}
