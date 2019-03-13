package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	l "github.com/tariel-x/anzer/lang"
)

type Parser struct {
	source string
	types []l.T
	funcs []l.F
	lastObj *container
}

const (
	kindT = iota
	kindF
)

type container struct {
	kind int
	t l.T
	f l.F
}

func New(source string) Parser {
	return Parser{
		source: source,
	}
}

func (parser *Parser) Parse()  {
	input := antlr.NewInputStream(parser.source)
	lexer := NewAnzerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := NewAnzerParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Forms()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(parser), tree)
}

type TreeShapeListener struct {
	*BaseAnzerListener
	parser *Parser
}

func NewTreeShapeListener(parser *Parser) *TreeShapeListener {
	l := new(TreeShapeListener)
	l.parser = parser
	return l
}

func (l *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func (l *TreeShapeListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {
	c := &container{
		kind: kindT,
	}
	l.parser.lastObj = c
}

func (l *TreeShapeListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {
	l.parser.types = append(l.parser.types, l.parser.lastObj.t)
	l.parser.lastObj = nil
}
