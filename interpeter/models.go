package interpeter

import "encoding/json"

const OpernadProduction = 0
const OpernadSum = 1

type BaseType struct {
	Type    *json.RawMessage
	Name    *string
	Operand *int
	Arg     *BaseType
}

func NewBaseType(Type string) BaseType {
	b := []byte(Type)
	j := json.RawMessage(b)
	return BaseType{
		Type: &j,
	}
}

func NewBaseTypeComplex(Name string, Operand int) BaseType {
	return BaseType{
		Name: &Name,
		Operand: &Operand,
	}
}

type BaseFunc struct {
	arg string
	ret string
	Def *FuncBody
}

type FuncBody struct {
	Name    string
	Product []FuncBody
	Param   *FuncBody
}

func NewBaseFunc(name string, arg string, ret string) BaseFunc {
	return BaseFunc{
		Def: &FuncBody{
			Name: name,
		},
		arg: arg,
		ret: ret,
	}
}

func (f *BaseFunc) Name() string {
	return f.Def.Name
}

func (f *BaseFunc) Arg() string {
	return f.arg
}

func (f *BaseFunc) Ret() string {
	return f.ret
}
