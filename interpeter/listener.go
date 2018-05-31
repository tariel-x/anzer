package interpeter

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/parser"
)

type Listener struct {
	*parser.BaseAnzerListener
	Types map[string]BaseType
	Funcs map[string]FuncDef
}

func NewListener() *Listener {
	l := new(Listener)
	types := map[string]BaseType{}
	funcs := map[string]FuncDef{}
	l.Types = types
	l.Funcs = funcs
	return l
}

/* func (l *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
} */

func (l *Listener) EnterJsonDataDef(ctx *parser.JsonDataDefContext) {
	name := ctx.DATA_NAME_ID().GetText()
	value := ctx.Json().GetText()
	l.Types[name] = *NewBaseType(value)
}

func (l *Listener) EnterLogicDataDef(ctx *parser.LogicDataDefContext) {
	name := ctx.DATA_NAME_ID().GetText()
	bt := &BaseType{}
	if ctx.DataOr() != nil {
		l.makeLogicDataDef(ctx.DataOr().GetChildren(), OpernadSum, bt)
	} else if ctx.DataAnd() != nil {
		l.makeLogicDataDef(ctx.DataAnd().GetChildren(), OpernadProduction, bt)
	}
	l.Types[name] = *bt
}

func (l *Listener) makeLogicDataDef(children []antlr.Tree, op int, def *BaseType) {
	def.Operand = &op
	for _, child := range children {
		p := child.GetPayload().(*antlr.BaseParserRuleContext)
		if p.GetRuleIndex() == parser.AnzerParserRULE_dataNameId {
			def.Args = append(def.Args, p.GetText())
		}
	}
}

func (l *Listener) EnterFuncSig(ctx *parser.FuncSigContext) {
	name := ctx.FUNC_NAME_ID().GetText()
	arg := ctx.DataName(0).GetText()
	ret := ctx.DataName(1).GetText()
	l.Funcs[name] = NewFuncDef(name, arg, ret)
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
			fb = l.processFuncDef(fb, child)
		}

	}
	fmt.Printf("func %s: %v\n", name, funcDef)
	fmt.Printf("func %s: definition %v\n", name, funcDef.Def)
	fmt.Printf("func %s: definition param %v\n", name, funcDef.Def.ComposeTo)
	fmt.Printf("func %s: definition param param %v\n", name, funcDef.Def.ComposeTo.ComposeTo)
}

func (l *Listener) processFuncDef(fb *FuncBody, child antlr.Tree) *FuncBody {
	p := child.GetPayload()
	switch t := p.(type) {
	case *antlr.CommonToken:
		return fb.AppendComposition(t.GetText())
	case *antlr.BaseParserRuleContext:
		return l.processFuncProduct(fb, t)
	default:
		_ = t
	}
	return nil
}

func (l *Listener) processFuncProduct(fb *FuncBody, tree antlr.Tree) *FuncBody {
	p := tree.GetPayload()
	switch t := p.(type) {
	case *antlr.BaseParserRuleContext:
		for _, child := range t.GetChildren() {
			return l.processFuncDef(fb, child)
		}
	default:
		_ = t
	}
	return nil
}
