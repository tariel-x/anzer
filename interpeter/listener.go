package interpeter

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/parser"
)

type Listener struct{
	*parser.BaseAnzerListener
	Types []BaseType
	Funcs []BaseFunc
}

func NewListener() *Listener {
	l := new(Listener)
	types := []BaseType{}
	funcs := []BaseFunc{}
	l.Types = types
	l.Funcs = funcs
	return l
}

func (l *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func (l *Listener) EnterDataSig(ctx *parser.DataSigContext) {
	name := ""
	val := ""
	for _, child :=  range ctx.GetChildren() {
		payload := child.GetPayload()
		switch t := payload.(type) {
		case *antlr.CommonToken:
		case *antlr.BaseParserRuleContext:
			if t.RuleIndex == 6 {
				name = t.GetText()
			}
			if t.RuleIndex == 8 {
				val = t.GetText()
			}
		default:
			_ = t
		}
	}

	l.Types = append(l.Types, NewBaseType(name, val))
}