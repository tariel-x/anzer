package interpeter

import "encoding/json"

const OpernadProduction = 0
const OpernadSum = 1

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
	Name string
	Arg  string
	Ret  string
	Def  *FuncBody
}

type FuncBody struct {
	Name       *string
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
		Name: &name,
	}
}

func (fb *FuncBody) AppendTo(parent *FuncBody) {
	parent.ComposeTo = fb
}

func (fb *FuncBody) Append(child *FuncBody) {
	fb.ComposeTo = child
}
