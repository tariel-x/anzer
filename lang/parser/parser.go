package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/lang"
	"strconv"
)

type Parser struct {
	source  string
	types   map[string]lang.T
	funcs   map[string]lang.Composable
	invokes []lang.F
	tc      *tContainer
	fc      *fContainer
}

type ParseResult struct {
	Types   map[string]lang.T
	Funcs   map[string]lang.Composable
	Invokes []lang.F
}

type fContainer struct {
	f lang.F
}

type tContainer struct {
	t      lang.T
	tpath  []lang.T
	parent *tContainer
}

func (c *tContainer) sub(t lang.T) *tContainer {
	return &tContainer{
		t:      t,
		tpath:  []lang.T{},
		parent: c,
	}
}

func (c *tContainer) appendT(t lang.T) {
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
	antlr.ParseTreeWalkerDefault.Walk(Newlistener(parser), tree)

	return ParseResult{
		Types: parser.types,
		Funcs: parser.funcs,
	}
}

type listener struct {
	*BaseAnzerListener
	parser *Parser
}

func Newlistener(parser *Parser) *listener {
	l := new(listener)
	l.parser = parser
	return l
}

func (l *listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("->  ", ctx.GetText())
}

func (l *listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("  <-", ctx.GetText())
}

func (l *listener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {
	l.parser.tc = &tContainer{}
}

func (l *listener) EnterTypeComplexDefinition(ctx *TypeComplexDefinitionContext) {
	l.parser.tc.t = lang.Complex{
		Fields: map[string]lang.T{},
	}
}

func (l *listener) EnterFieldName(ctx *FieldNameContext) {
	l.parser.tc = l.parser.tc.sub(nil)
}

func (l *listener) ExitTypeField(ctx *TypeFieldContext) {
	parentc := l.parser.tc.parent
	if complext, ok := parentc.t.(lang.Complex); ok {
		fieldName := ctx.FieldName().GetText()
		complext.Fields[fieldName] = l.parser.tc.t
	}
	l.parser.tc = parentc
}

func (l *listener) ExitTypeMinLength(ctx *TypeMinLengthContext) {
	arg, _ := strconv.Atoi(ctx.ConstructorArg().GetText())
	t := lang.Construct(nil, lang.TypeMinLength, []interface{}{arg})
	l.parser.tc.appendT(t)
}

func (l *listener) ExitTypeMaxLength(ctx *TypeMaxLengthContext) {
	arg, _ := strconv.Atoi(ctx.ConstructorArg().GetText())
	t := lang.Construct(nil, lang.TypeMaxLength, []interface{}{arg})
	l.parser.tc.appendT(t)
}

func (l *listener) ExitTypeOptional(ctx *TypeOptionalContext) {
	t := lang.Construct(nil, lang.TypeOptional, nil)
	l.parser.tc.appendT(t)
}

func (l *listener) ExitTypeString(ctx *TypeStringContext) {
	l.parser.tc.appendT(lang.TypeString)
}

func (l *listener) ExitTypeBool(ctx *TypeBoolContext) {
	l.parser.tc.appendT(lang.TypeBool)
}

func (l *listener) ExitTypeInteger(ctx *TypeIntegerContext) {
	l.parser.tc.appendT(lang.TypeInteger)
}

func (l *listener) ExitTypeSimpleDefinition(ctx *TypeSimpleDefinitionContext) {
	var finalT lang.T
	for i := len(l.parser.tc.tpath) - 1; i >= 0; i-- {
		if finalT == nil {
			finalT = l.parser.tc.tpath[i]
			continue
		}
		if constructor, ok := l.parser.tc.tpath[i].(lang.Constructor); ok {
			constructor.Operand = finalT
			finalT = constructor
		}
	}
	l.parser.tc.t = finalT
}

func (l *listener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {
	l.parser.types[ctx.TypeName().GetText()] = l.parser.tc.t
	l.parser.tc = nil
}

// Link functions

func (l *listener) EnterFuncDeclaration(ctx *FuncDeclarationContext) {
	l.parser.fc = &fContainer{
		f: lang.F{},
	}
}

func (l *listener) EnterFuncName(ctx *FuncNameContext) {
	l.parser.fc.f.Name = ctx.GetText()
}

func (l *listener) EnterUrl(ctx *UrlContext) {
	l.parser.fc.f.Link = lang.FunctionLink(ctx.GetText())
}
