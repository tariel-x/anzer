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

type FuncProduction []FuncBody

type FuncBody struct {
	Name      *string
	Product   FuncProduction
	ComposeTo *FuncBody
}

func NewFuncDef(name string, arg string, ret string) FuncDef {
	return FuncDef{
		Name: name,
		Arg:  arg,
		Ret:  ret,
	}
}

func (f *FuncDef) AppendComposition(name string) *FuncBody {
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
	b.ComposeTo = fb
	return fb
}

func (b *FuncBody) AppendProdComposition(names []string) *FuncBody {
	cfb := FuncBody{}
	fbs := FuncProduction{}
	for _, name := range names {
		fb := &FuncBody{
			Name: &name,
		}
		fbs = append(fbs, *fb)
	}

	cfb.Product = fbs
	b.ComposeTo = &cfb
	return &cfb
}
