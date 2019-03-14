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
	tpath  []lang.T
	parent *container
}

func (c *container) sub(t lang.T) *container {
	return &container{
		t:      t,
		tpath:  []lang.T{},
		parent: c,
	}
}

func (c *container) appendT(t lang.T) {
	if c == nil {
		return
	}
	if c.tpath == nil {
		c.tpath = []lang.T{}
	}
	c.tpath = append(c.tpath, t)
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
	fmt.Println("->  ", ctx.GetText())
}

func (l *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("  <-", ctx.GetText())
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
	t := lang.Construct(nil, lang.TypeMinLength, []interface{}{arg})
	l.parser.container.appendT(t)
}

func (l *TreeShapeListener) ExitTypeMaxLength(ctx *TypeMaxLengthContext) {
	arg, _ := strconv.Atoi(ctx.ConstructorArg().GetText())
	t := lang.Construct(nil, lang.TypeMaxLength, []interface{}{arg})
	l.parser.container.appendT(t)
}

func (l *TreeShapeListener) ExitTypeOptional(ctx *TypeOptionalContext) {
	t := lang.Construct(nil, lang.TypeOptional, nil)
	l.parser.container.appendT(t)
}

func (l *TreeShapeListener) ExitTypeString(ctx *TypeStringContext) {
	l.parser.container.appendT(lang.TypeString)
}

func (l *TreeShapeListener) ExitTypeBool(ctx *TypeBoolContext) {
	l.parser.container.appendT(lang.TypeBool)
}

func (l *TreeShapeListener) ExitTypeInteger(ctx *TypeIntegerContext) {
	l.parser.container.appendT(lang.TypeInteger)
}

func (l *TreeShapeListener) ExitTypeSimpleDefinition(ctx *TypeSimpleDefinitionContext) {
	var finalT lang.T
	for i := len(l.parser.container.tpath) - 1; i >= 0; i-- {
		if finalT == nil {
			finalT = l.parser.container.tpath[i]
			continue
		}
		if constructor, ok := l.parser.container.tpath[i].(lang.Constructor); ok {
			constructor.Operand = finalT
			finalT = constructor
		}
	}
	l.parser.container.t = finalT
}

func (l *TreeShapeListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {
	l.parser.types[ctx.TypeName().GetText()] = l.parser.container.t
	l.parser.container = nil
}
