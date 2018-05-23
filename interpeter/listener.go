package interpeter

import (
	"github.com/tariel-x/anzer/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
)

type Listener struct{
	*parser.BaseAnzerListener
}

func NewListener() *Listener {
	return new(Listener)
}

func (l *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}