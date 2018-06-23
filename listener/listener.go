package listener

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/parser"
)

type Listener struct {
	*parser.BaseAnzerListener
	Types Types
	Funcs Funcs
}

func NewListener() *Listener {
	l := new(Listener)
	types := Types{}
	funcs := Funcs{}
	l.Types = types
	l.Funcs = funcs
	return l
}

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
	l.Funcs[name] = NewFunc(name, arg, ret)
}

func (l *Listener) EnterFuncDef(ctx *parser.FuncDefContext) {
	name := ctx.FUNC_NAME_ID().GetText()
	def := l.Funcs[name]
	composes := ctx.AllComposeFunc()
	fb := &FuncBody{}
	childFb := fb
	for _, compose := range composes {
		childFb = l.processCompose(compose, childFb)
	}
	def.Body = fb.ComposeTo
	l.Funcs[name] = def
}

func (l *Listener) processCompose(ctx antlr.ParserRuleContext, fb *FuncBody) *FuncBody {
	if ctx.GetRuleIndex() != parser.AnzerParserRULE_composeFunc {
		panic("Not a compose func")
	}
	for _, child := range ctx.GetChildren() {
		p := child.GetPayload()
		switch t := p.(type) {
		case *antlr.CommonToken:
			name := t.GetText()
			return l.processCommonToken(name, fb)
		case *antlr.BaseParserRuleContext:
			return l.processProductFunc(t, fb)
		}
	}
	return nil
}

func (l *Listener) processCommonToken(name string, fb *FuncBody) *FuncBody {
	childFb := SimpleFunc(name)
	fb.Append(childFb)
	return childFb
}

func (l *Listener) processProductFunc(ctx antlr.ParserRuleContext, fb *FuncBody) *FuncBody {
	if ctx.GetRuleIndex() != parser.AnzerParserRULE_productFunc {
		panic("Not a product func")
	}
	prods := Production{}
	for _, child := range ctx.GetChildren() {
		p := child.GetPayload()
		switch t := p.(type) {
		case *antlr.BaseParserRuleContext:
			childFb := l.processCompose(t, fb)
			prods = append(prods, *childFb)
		}
	}
	childFb := ProductFunc(prods)
	childFb.AppendTo(fb)
	return childFb
}
