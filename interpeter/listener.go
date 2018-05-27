package interpeter

import (
	"fmt"
	//"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/parser"
)

type Listener struct {
	*parser.BaseAnzerListener
	Types map[string]BaseType
	Funcs map[string]BaseFunc
}

func NewListener() *Listener {
	l := new(Listener)
	types := map[string]BaseType{}
	funcs := map[string]BaseFunc{}
	l.Types = types
	l.Funcs = funcs
	return l
}

/* func (l *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
} */

func (l *Listener) EnterDataSig(ctx *parser.DataSigContext) {
	name := ctx.DATA_NAME_ID().GetText()
	for _, child := range ctx.DataDefinition().GetChildren() {
		l.processDataDef(name, child)
	}
	l.Types[name] = NewBaseType(name, val)
}

func (l *Listener) processDataDef(fname string, child antlr.Tree) *FuncBody {
	p := child.GetPayload()
	switch t := p.(type) {
	case *antlr.CommonToken:
		return l.appendFunc(fb, t.GetText())
	case *antlr.BaseParserRuleContext:
		return l.appendProduct(fb, t.GetText(), t)
	default:
		_ = t
	}
	return nil
}

func (l *Listener) EnterFuncSig(ctx *parser.FuncSigContext) {
	name := ctx.FUNC_NAME_ID().GetText()
	arg := ctx.DataName(0).GetText()
	ret := ctx.DataName(1).GetText()
	l.Funcs[name] = NewBaseFunc(name, arg, ret)
}

func (l *Listener) EnterFuncDef(ctx *parser.FuncDefContext) {
	name := ctx.FUNC_NAME_ID().GetText()
	fmt.Printf("name %s\n", name)

	funcDef, exists := l.Funcs[name]
	if !exists {
		panic(fmt.Errorf("Can not find func %s", name))
	}

	fb := funcDef.Def

	for _, cf := range ctx.AllComposeFunc() {
		children := cf.GetChildren()
		for _, child := range children {
			fb = l.processFuncDef(fb, name, child)
		}

	}
	fmt.Printf("func %s: %v\n", name, funcDef)
	fmt.Printf("func %s: definition %v\n", name, funcDef.Def)
	fmt.Printf("func %s: definition param %v\n", name, funcDef.Def.Param)
	fmt.Printf("func %s: definition param param %v\n", name, funcDef.Def.Param.Param)
}

func (l *Listener) processFuncDef(fb *FuncBody, name string, child antlr.Tree) *FuncBody {
	p := child.GetPayload()
	switch t := p.(type) {
	case *antlr.CommonToken:
		return l.appendFunc(fb, t.GetText())
	case *antlr.BaseParserRuleContext:
		return l.appendProduct(fb, t.GetText(), t)
	default:
		_ = t
	}
	return nil
}

func (l *Listener) appendFunc(fb *FuncBody, name string) *FuncBody {
	cfb := FuncBody{
		Name: name,
	}
	fb.Param = &cfb
	return fb.Param
}

func (l *Listener) appendProduct(fb *FuncBody, name string, tree antlr.Tree) *FuncBody {
	children := tree.GetChildren()
	product := []FuncBody{}
	for _, child := range children {
		pfb := l.processFuncDef(fb, name, child)
		product = append(product, *pfb)
	}
	fb.Product = product
	return fb
}
