package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/lang/parser/parser"
)

type Parser struct {
}

func New() Parser {
	return Parser{}
}

func (pp Parser) Parse(source string)  {
	input := antlr.NewInputStream(source)
	lexer := parser.NewAnzerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := parser.NewAnzerParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Forms()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

type TreeShapeListener struct {
	*parser.BaseAnzerListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}