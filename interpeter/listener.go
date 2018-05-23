package interpeter

import (
	"github.com/tariel-x/anzer/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Listener struct{
	*parser.BaseAnzerListener
}

func NewListener() *Listener {
	return new(Listener)
}

//func (l *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	fmt.Println(ctx.GetText())
//}

func (l *Listener) EnterDataSig(ctx *parser.DataSigContext) {
	fmt.Printf("---enter data sig %s\n", ctx.GetText())

	name := ""
	val := ""
	for _, child :=  range ctx.GetChildren() {
		payload := child.GetPayload()
		switch t := payload.(type) {
		case *antlr.CommonToken:
			fmt.Printf("%T\n", payload)
			fmt.Printf("comm tkn\n")
			fmt.Printf("string: %s\n", t.String())
		case *antlr.BaseParserRuleContext:
			fmt.Printf("%T\n", payload)
			fmt.Printf("ctx\n")

			fmt.Printf("rule index %d ", t.RuleIndex)
			fmt.Printf("text: %s\n", t.GetText())
			if t.RuleIndex == 6 {
				name = t.GetText()
			}
			if t.RuleIndex == 8 {
				val = t.GetText()
			}
		default:
			_ = t
			fmt.Printf("unkn\n")
			fmt.Printf("%T\n", payload)
		}
	}

	fmt.Printf("Name: %s val: %s\n", name, val)
}