package interpeter

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/parser"
)

type Listener struct{
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

// func (l *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Println(ctx.GetText())
// }

func (l *Listener) EnterDataSig(ctx *parser.DataSigContext) {
	name := ""
	val := ""
	for _, child :=  range ctx.GetChildren() {
		payload := child.GetPayload()
		switch t := payload.(type) {
		case *antlr.CommonToken:
		case *antlr.BaseParserRuleContext:
			if t.RuleIndex == parser.AnzerParserRULE_dataNameId {
				name = t.GetText()
			}
			if t.RuleIndex == parser.AnzerParserRULE_json {
				val = t.GetText()
			}
		default:
			_ = t
		}
	}
	l.Types[name] = NewBaseType(name, val)
}

func (l *Listener) EnterFuncSig(ctx *parser.FuncSigContext) {
	name := ""
	val := ""
	for _, child :=  range ctx.GetChildren() {
		payload := child.GetPayload()
		switch t := payload.(type) {
		case *antlr.CommonToken:
		case *antlr.BaseParserRuleContext:
			if t.RuleIndex == parser.AnzerParserRULE_funcNameId {
				name = t.GetText()
			}
			if t.RuleIndex == parser.AnzerParserRULE_dataName {
				val = t.GetText()
			}
		default:
			_ = t
		}
	}
	l.Funcs[name] = NewBaseFunc(name, val, val)
}