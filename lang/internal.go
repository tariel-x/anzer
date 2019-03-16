package lang

import (
	"errors"
	"fmt"
	"path"
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
	previous := a.Compose[0]
	for _, c := range a.Compose[1:] {
		if !c.In().Equal(previous.Out()) && !c.In().Parent(previous.Out()) {
			return errTypeInconsistent
		}
		previous = c
	}
	return nil
}

type FunctionLink string

type F struct {
	Name    string
	Link    FunctionLink
	Runtime string
	TypeIn  T
	TypeOut T
}

func (f F) Definition() string {
	return f.GetName()
}

func (f F) GetName() string {
	if f.Name != "" {
		return f.Name
	}
	if f.Link == "" {
		return ""
	}
	return path.Base(string(f.Link))
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

type FRef string

func (f FRef) Definition() string {
	return string(f)
}

func (f FRef) In() T {
	return nil
}

func (f FRef) Out() T {
	return nil
}

func (f FRef) Invalid() error {
	return nil
}
