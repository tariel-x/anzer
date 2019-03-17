package parser

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/lang"
)

type Parser struct {
	source  string
	types   map[string]lang.T
	funcs   map[string]lang.Composable
	invokes []lang.FRef
	tc      *tContainer
	fc      *fContainer
}

type ParseResult struct {
	Types   map[string]lang.T
	Funcs   map[string]lang.Composable
	Invokes []lang.F
}

func New(source string) Parser {
	return Parser{
		source:  source,
		types:   map[string]lang.T{},
		funcs:   map[string]lang.Composable{},
		invokes: []lang.FRef{},
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

type fContainer struct {
	f       lang.F
	a       lang.Alias
	refpath []lang.FRef
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

func (l *listener) ExitTypeComplexDefinition(ctx *TypeComplexDefinitionContext) {
	tpath := append(l.parser.tc.tpath, l.parser.tc.t)
	l.parser.tc.tpath = []lang.T{}
	l.parser.tc.t = l.reduceTpath(tpath)
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

func (l *listener) EnterTypeMinLength(ctx *TypeMinLengthContext) {
	arg, _ := strconv.Atoi(ctx.ConstructorArg().GetText())
	t := lang.Construct(nil, lang.TypeMinLength, []interface{}{arg})
	l.parser.tc.appendT(t)
}

func (l *listener) EnterTypeMaxLength(ctx *TypeMaxLengthContext) {
	arg, _ := strconv.Atoi(ctx.ConstructorArg().GetText())
	t := lang.Construct(nil, lang.TypeMaxLength, []interface{}{arg})
	l.parser.tc.appendT(t)
}

func (l *listener) EnterTypeOptional(ctx *TypeOptionalContext) {
	t := lang.Construct(nil, lang.TypeOptional, nil)
	l.parser.tc.appendT(t)
}

func (l *listener) EnterTypeList(ctx *TypeListContext) {
	t := lang.Construct(nil, lang.TypeList, nil)
	l.parser.tc.appendT(t)
}

func (l *listener) EnterTypeString(ctx *TypeStringContext) {
	l.parser.tc.appendT(lang.TypeString)
}

func (l *listener) EnterTypeBool(ctx *TypeBoolContext) {
	l.parser.tc.appendT(lang.TypeBool)
}

func (l *listener) EnterTypeInteger(ctx *TypeIntegerContext) {
	l.parser.tc.appendT(lang.TypeInteger)
}

func (l *listener) EnterTypeOther(ctx *TypeOtherContext) {
	l.parser.tc.appendT(lang.Ref(ctx.GetText()))
}

func (l *listener) ExitTypeSimpleDefinition(ctx *TypeSimpleDefinitionContext) {
	l.parser.tc.t = l.reduceTpath(l.parser.tc.tpath)
	l.parser.tc.tpath = []lang.T{}
}

func (l *listener) reduceTpath(tpath []lang.T) lang.T {
	var finalT lang.T
	for i := len(tpath) - 1; i >= 0; i-- {
		if finalT == nil {
			finalT = tpath[i]
			continue
		}
		if constructor, ok := tpath[i].(lang.Constructor); ok {
			constructor.Operand = finalT
			finalT = constructor
		}
	}
	return finalT
}

func (l *listener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {
	l.parser.types[ctx.TypeName().GetText()] = l.parser.tc.t
	l.parser.tc = nil
}

// Link functions

func (l *listener) EnterFuncDeclaration(ctx *FuncDeclarationContext) {
	l.parser.fc = &fContainer{
		f: lang.F{
			Link:    lang.FunctionLink(ctx.Url().GetText()),
			Runtime: ctx.Runtime().GetText(),
		},
	}
	if ctx.FuncName() != nil {
		l.parser.fc.f.Name = ctx.FuncName().GetText()
	}
}

func (l *listener) EnterFuncArgument(ctx *FuncArgumentContext) {
	l.parser.tc = &tContainer{}
}

func (l *listener) ExitFuncArgument(ctx *FuncArgumentContext) {
	if len(l.parser.tc.tpath) > 0 {
		l.parser.tc.t = l.reduceTpath(l.parser.tc.tpath)
	}
	l.parser.fc.f.TypeIn = l.parser.tc.t
	l.parser.tc = nil
}

func (l *listener) EnterFuncResult(ctx *FuncResultContext) {
	l.parser.tc = &tContainer{}
}

func (l *listener) ExitFuncResult(ctx *FuncResultContext) {
	if len(l.parser.tc.tpath) > 0 {
		l.parser.tc.t = l.reduceTpath(l.parser.tc.tpath)
	}
	l.parser.fc.f.TypeOut = l.parser.tc.t
	l.parser.tc = nil
}

func (l *listener) ExitFuncDeclaration(ctx *FuncDeclarationContext) {
	l.parser.funcs[l.parser.fc.f.GetName()] = l.parser.fc.f
	l.parser.fc = nil
}

func (l *listener) EnterLocalFuncDeclaration(ctx *LocalFuncDeclarationContext) {
	l.parser.fc = &fContainer{
		refpath: []lang.FRef{},
		a:       lang.Alias{},
	}
	if ctx.FuncName() != nil {
		l.parser.fc.a.Name = ctx.FuncName().GetText()
	}
}

func (l *listener) EnterFuncRef(ctx *FuncRefContext) {
	l.parser.fc.refpath = append(l.parser.fc.refpath, lang.FRef(ctx.GetText()))
}

func (l *listener) ExitLocalFuncDeclaration(ctx *LocalFuncDeclarationContext) {
	composables := make([]lang.Composable, 0, len(l.parser.fc.refpath))
	for _, ref := range l.parser.fc.refpath {
		composables = append(composables, ref)
	}
	l.parser.fc.a.Compose = composables
	l.parser.funcs[l.parser.fc.a.Name] = l.parser.fc.a
	l.parser.fc = nil
}

func (l *listener) EnterInvokeFuncName(ctx *InvokeFuncNameContext) {
	l.parser.invokes = append(l.parser.invokes, lang.FRef(ctx.GetText()))
}
