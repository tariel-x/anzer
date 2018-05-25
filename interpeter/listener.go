package interpeter

import (
	"fmt"
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

func (l *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

/*func (l *Listener) EnterDataSig(ctx *parser.DataSigContext) {
	name := ctx.DataNameId().GetText()
	val := ctx.Json().GetText()
	l.Types[name] = NewBaseType(name, val)
}

func (l *Listener) EnterFuncSig(ctx *parser.FuncSigContext) {
	name := ctx.FuncNameId().GetText()
	arg := ctx.DataName(0).GetText()
	ret := ctx.DataName(1).GetText()
	l.Funcs[name] = NewBaseFunc(name, arg, ret)
}*/
