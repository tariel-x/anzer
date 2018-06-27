package listener

import "encoding/json"

const OpernadProduction = 0
const OpernadSum = 1

type Types map[string]BaseType
type Funcs map[string]FuncDef

type BaseType struct {
	Type    *json.RawMessage
	Operand *int
	Args    []string
}

func NewBaseType(Type string) *BaseType {
	b := []byte(Type)
	j := json.RawMessage(b)
	return &BaseType{
		Type: &j,
	}
}

type FuncDef struct {
	Name   string
	Arg    string
	Ret    string
	Body   *FuncBody
	Params *FuncParams
}

type FuncParams struct {
	EnvIn  string
	EnvOut string
	Envs   map[string]string
}

type FuncBody struct {
	Ref        *string
	ProductEls Production
	SumEls     Sum
	ComposeTo  *FuncBody
}

type Production []FuncBody
type Sum []FuncBody

func NewFunc(name, arg, ret string) FuncDef {
	return FuncDef{
		Name: name,
		Arg:  arg,
		Ret:  ret,
	}
}

func ProductFunc(product Production) *FuncBody {
	return &FuncBody{
		ProductEls: product,
	}
}

func SumFunc(sum Sum) *FuncBody {
	return &FuncBody{
		SumEls: sum,
	}
}

func SimpleFunc(name string) *FuncBody {
	return &FuncBody{
		Ref: &name,
	}
}

func (fb *FuncBody) AppendTo(parent *FuncBody) {
	parent.ComposeTo = fb
}

func (fb *FuncBody) Append(child *FuncBody) {
	fb.ComposeTo = child
}
