package interpeter

import "encoding/json"

const OpernadProduction = 0
const OpernadSum = 1

type BaseType struct {
	Type    *json.RawMessage
	Operand *int
	Args     []string
}

func NewBaseType(Type string) *BaseType {
	b := []byte(Type)
	j := json.RawMessage(b)
	return &BaseType{
		Type: &j,
	}
}

type BaseFunc struct {
	Name string
	Arg  string
	Ret  string
	Def  *FuncBody
}

type FuncBody struct {
	Name    *string
	Product []FuncBody
	Param   *FuncBody
}

func NewBaseFunc(name string, arg string, ret string) BaseFunc {
	return BaseFunc{
		Name: name,
		Arg:  arg,
		Ret:  ret,
	}
}

func (f *BaseFunc) AppendComposition(name string) *FuncBody {
	fb := &FuncBody{
		Name: &name,
	}
	f.Def = fb
	return fb
}

func (b *FuncBody) AppendComposition(name string) *FuncBody {
	fb := &FuncBody{
		Name: &name,
	}
	b.Param = fb
	return fb
}

func (b *FuncBody) AppendProdComposition(names []string) []FuncBody {
	fbs := []FuncBody{}
	for _, name := range names {
		fb := &FuncBody{
			Name: &name,
		}
		fbs = append(fbs, *fb)
	}

	b.Product = fbs
	return fbs
}
