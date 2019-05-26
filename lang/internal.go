package lang

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

var (
	errTypeInconsistent = errors.New("types are not consistent")
	ErrBindNoArg        = errors.New("bind has no argument")
	ErrReturnNoArg      = errors.New("return has no argument")
	ErrBindArgNotEither = errors.New("argument of bind has not either output type")
	ErrBindArgNotF      = errors.New("argument of bind is not a function")
)

type Composable interface {
	Definition() string
	GetName() string
	In() T
	Out() T
	Invalid() error
}

type Runnable interface {
	Definition() string
	GetName() string
	GetLink() FunctionLink
	GetRuntime() string
	In() T
	Out() T
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

func (a Alias) GetName() string {
	return a.Name
}

func (a Alias) In() T {
	if len(a.Compose) == 0 {
		panic(fmt.Errorf("type %s is undefined", a.Name))
	}
	return a.Compose[0].In()
}

func (a Alias) Out() T {
	if len(a.Compose) == 0 {
		panic(fmt.Errorf("type %s is undefined", a.Name))
	}
	return a.Compose[len(a.Compose)-1].Out()
}

func (a Alias) Invalid() error {
	if len(a.Compose) <= 1 {
		return nil
	}
	previous := a.Compose[0]
	for _, c := range a.Compose[1:] {
		cIn := c.In()
		if cIn == nil {
			return errTypeInconsistent
		}
		previousOut := previous.Out()
		if !cIn.Equal(previousOut) && !cIn.Parent(previousOut) {
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

func (f F) GetLink() FunctionLink {
	return f.Link
}

func (f F) GetRuntime() string {
	return f.Runtime
}

// InternalReference stores reference to another function in a source code.
// Lets consider `f = a . b`. `a` and `b` would be stored in InternalReference type during parsing source code.
type InternalReference string

func (f InternalReference) Definition() string {
	return string(f)
}

func (f InternalReference) GetName() string {
	return string(f)
}

func (f InternalReference) In() T {
	return nil
}

func (f InternalReference) Out() T {
	return nil
}

func (f InternalReference) Invalid() error {
	return nil
}

// EitherBind is HOF that stores reference to another function.
type EitherBind struct {
	Argument Composable
}

func (b EitherBind) Definition() string {
	return fmt.Sprint(b.GetName(), " ", b.Argument.Definition())
}

func (b EitherBind) GetName() string {
	return ">>="
}

func (b EitherBind) In() T {
	if b.Argument == nil {
		return nil
	}
	if b.Argument.In() == nil {
		return nil
	}
	argOut := b.Argument.Out()
	if IsEither(argOut) {
		return Either(b.Argument.In(), argOut.(Sum)[1].(Constructor).FirstOperand())
	}
	return nil
}

func (b EitherBind) Out() T {
	if b.Argument == nil {
		return nil
	}
	if b.Argument.Out() == nil {
		return nil
	}
	argOut := b.Argument.Out()
	if IsEither(argOut) {
		return argOut
	}
	return nil
}

func (b EitherBind) Invalid() error {
	if b.Argument == nil {
		return ErrBindNoArg
	}
	if _, ok := b.Argument.(F); !ok {
		return ErrBindArgNotF
	}
	if b.Argument.In() == nil {
		return ErrBindArgNotEither
	}

	argOut := b.Argument.Out()

	if !IsEither(argOut) {
		return ErrBindArgNotEither
	}

	return nil
}

func (b EitherBind) GetLink() FunctionLink {
	if f, ok := b.Argument.(F); ok {
		return f.Link
	}
	return ""
}

func (b EitherBind) GetRuntime() string {
	if f, ok := b.Argument.(F); ok {
		return f.Runtime
	}
	return ""
}

func Bind(arg Composable) EitherBind {
	return EitherBind{
		Argument: arg,
	}
}

// EitherReturn is HOF that stores reference to another function.
type EitherReturn struct {
	Argument Composable
}

func (r EitherReturn) Definition() string {
	return fmt.Sprint(r.GetName(), r.Argument.Definition())
}

func (r EitherReturn) GetName() string {
	return "return"
}

func (r EitherReturn) In() T {
	if r.Argument == nil {
		return nil
	}
	if r.Argument.In() == nil {
		return nil
	}
	return r.Argument.In()
}

func (r EitherReturn) Out() T {
	if r.Argument == nil {
		return nil
	}
	if r.Argument.Out() == nil {
		return nil
	}
	return Right(r.Argument.Out())
}

func (r EitherReturn) Invalid() error {
	if r.Argument == nil {
		return ErrReturnNoArg
	}
	if r.Argument.In() == nil {
		return ErrReturnNoArg
	}
	//TODO: check that invalid function is correct
	return nil
}

func Return(arg Composable) EitherReturn {
	return EitherReturn{
		Argument: arg,
	}
}
