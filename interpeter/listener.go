package interpeter

import (
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
	l.Funcs[name] = NewFunc(name, arg, ret)
}

func (l *Listener) EnterFuncDef(ctx *parser.FuncDefContext) {
	name := ctx.FUNC_NAME_ID().GetText()
	def := l.Funcs[name]
	composes := ctx.AllComposeFunc()
	def.Def = &FuncBody{}
	for _, compose := range composes {
		l.processCompose(compose, def.Def)
	}
}

func (l *Listener) processCompose(ctx parser.IComposeFuncContext, fb *FuncBody) {
	children := ctx.GetChildren()
	for
}
