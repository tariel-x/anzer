package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/lang"
	"strconv"
)

type Parser struct {
	source    string
	types     map[string]lang.T
	funcs     map[string]lang.F
	container *container
}

type ParseResult struct {
	Types map[string]lang.T
	Funcs map[string]lang.F
}

type container struct {
	t      lang.T
	parent *container
}

func (c *container) sub(t lang.T) *container {
	return &container{
		t:      t,
		parent: c,
	}
}

func New(source string) Parser {
	return Parser{
		source: source,
		types:  map[string]lang.T{},
	}
}

func (parser *Parser) Parse() ParseResult {
	input := antlr.NewInputStream(parser.source)
	lexer := NewAnzerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewAnzerParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Forms()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(parser), tree)

	return ParseResult{
		Types: parser.types,
		Funcs: parser.funcs,
	}
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
	fmt.Println("->", ctx.GetText())
}

func (l *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("<-", ctx.GetText())
}

func (l *TreeShapeListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {
	l.parser.container = &container{}
}

func (l *TreeShapeListener) EnterTypeComplexDefinition(ctx *TypeComplexDefinitionContext) {
	l.parser.container.t = lang.Complex{
		Fields: map[string]lang.T{},
	}
}

func (l *TreeShapeListener) EnterFieldName(ctx *FieldNameContext) {
	l.parser.container = l.parser.container.sub(nil)
}

func (l *TreeShapeListener) ExitTypeField(ctx *TypeFieldContext) {
	parentc := l.parser.container.parent
	if complext, ok := parentc.t.(lang.Complex); ok {
		fieldName := ctx.FieldName().GetText()
		complext.Fields[fieldName] = l.parser.container.t
	}
	l.parser.container = parentc
}

func (l *TreeShapeListener) ExitTypeMinLength(ctx *TypeMinLengthContext) {
	arg, _ := strconv.Atoi(ctx.ConstructorArg().GetText())
	// TODO: HOW TO DETECT NEXT TYPE HERE?
	t := lang.MinLength(l.parser.container.t, arg)
	l.parser.container = l.parser.container.sub(t)
}

func (l *TreeShapeListener) ExitTypeMaxLength(ctx *TypeMaxLengthContext) {
	arg, _ := strconv.Atoi(ctx.ConstructorArg().GetText())
	l.parser.container.t = lang.MaxLength(l.parser.container.t, arg)
}

func (l *TreeShapeListener) ExitTypeOptional(ctx *TypeOptionalContext) {
	l.parser.container.t = lang.Optional(l.parser.container.t)
}

func (l *TreeShapeListener) ExitTypeString(ctx *TypeStringContext) {
	l.parser.container.t = lang.TypeString
}

func (l *TreeShapeListener) ExitTypeBool(ctx *TypeBoolContext) {
	l.parser.container.t = lang.TypeBool
}

func (l *TreeShapeListener) ExitTypeInteger(ctx *TypeIntegerContext) {
	l.parser.container.t = lang.TypeInteger
}

func (l *TreeShapeListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {
	l.parser.types[ctx.TypeName().GetText()] = l.parser.container.t
	l.parser.container = nil
}
