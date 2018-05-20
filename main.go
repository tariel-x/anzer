package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/parser"
	"os"
)

type TreeShapeListener struct {
	*parser.BaseAnzerListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewAnzerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAnzerParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	p.GetRuleNames()
	dd := p.DataDef()
	txt := dd.GetText()
	fmt.Printf("%s\n", txt)
	/*tree := p
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)*/
}
