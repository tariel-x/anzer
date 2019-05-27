//package parser contains both ANTLR-generated files and parser logic.
package parser

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"
	"github.com/tariel-x/anzer/lang"
)

type Logger interface {
	Debug(args ...interface{})
}

type stubLogger struct{}

func (l stubLogger) Debug(args ...interface{}) {}

var (
	ErrTypeRefUnreachable = errors.New("Type reference can not be reached")
	ErrFuncRefUnreachable = errors.New("Function reference can not be reached")
	ErrFuncTypeIncorrect  = errors.New("Func type is incorrect")
)

// Parser contains source code and temporary internal representation which is used to build final internal representation.
type Parser struct {
	logger  Logger
	source  string
	types   map[string]lang.T
	funcs   map[string]lang.Composable
	invokes []lang.InternalReference
	tc      *tContainer
	fc      *fContainer
}

func New(source string) Parser {
	return Parser{
		logger:  stubLogger{},
		source:  source,
		types:   map[string]lang.T{},
		funcs:   map[string]lang.Composable{},
		invokes: []lang.InternalReference{},
	}
}

// ParseLazy returns internal representation only for types and functions that are mentioned in the `invoke` instruction.
func (parser *Parser) ParseLazy() ([]lang.Composable, error) {
	parser.buildTree()

	result := []lang.Composable{}
	for _, invk := range parser.invokes {
		composable, err := parser.resolveFunc(invk)
		if err != nil {
			return nil, err
		}
		result = append(result, composable)
	}
	return result, nil
}

// ParseAll returns internal representation for all types and functions from the source code.
func (parser *Parser) ParseAll() ([]lang.Composable, error) {
	parser.buildTree()

	composesMap, err := parser.resolveFuncs(parser.funcs)
	if err != nil {
		return nil, err
	}
	composes := make([]lang.Composable, 0, len(composesMap))
	for _, c := range composesMap {
		composes = append(composes, c)
	}
	return composes, nil
}

// ParseTypes returns only parsed types
func (parser *Parser) ParseTypes() (map[string]lang.T, error) {
	parser.buildTree()
	return parser.resolveTypes(parser.types)
}

func (parser *Parser) buildTree() {
	input := antlr.NewInputStream(parser.source)
	lexer := NewAnzerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewAnzerParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Forms()
	antlr.ParseTreeWalkerDefault.Walk(Newlistener(parser), tree)
}

func (parser *Parser) SetLogger(logger Logger) {
	parser.logger = logger
}

func (parser *Parser) resolveTypes(types map[string]lang.T) (map[string]lang.T, error) {
	for name, t := range types {
		tr, err := parser.resolveType(t)
		if err != nil {
			return nil, err
		}
		types[name] = tr
	}
	return types, nil
}

func (parser *Parser) resolveType(t lang.T) (lang.T, error) {
	switch tt := t.(type) {
	case lang.Record:
		for fname, ft := range tt.Fields {
			ftr, err := parser.resolveType(ft)
			if err != nil {
				return nil, err
			}
			tt.Fields[fname] = ftr
		}
		return tt, nil
	case lang.Container:
		if !lang.IsEither(tt) {
			if len(tt.Operands) == 0 {
				return nil, errors.New("container without operands")
			}
			opr, err := parser.resolveType(tt.Operands[0])
			if err != nil {
				return nil, err
			}
			tt.Operands[0] = opr
			return tt, nil
		}
	case lang.Ref:
		if target, ok := parser.types[string(tt)]; ok {
			return parser.resolveType(target)
		}
		return nil, errors.Wrap(ErrTypeRefUnreachable, string(tt))
	case lang.Sum:
		for i, ft := range tt {
			ftr, err := parser.resolveType(ft)
			if err != nil {
				return nil, err
			}
			tt[i] = ftr
		}
		return tt, nil
	}
	return t, nil
}

func (parser *Parser) resolveFuncs(funcs map[string]lang.Composable) (map[string]lang.Composable, error) {
	for name, f := range funcs {
		tr, err := parser.resolveFunc(f)
		if err != nil {
			return nil, err
		}
		funcs[name] = tr
	}
	return funcs, nil
}

func (parser *Parser) resolveFunc(f lang.Composable) (lang.Composable, error) {
	switch ff := f.(type) {
	case lang.Alias:
		for idx, f := range ff.Compose {
			fr, err := parser.resolveFunc(f)
			if err != nil {
				return nil, err
			}
			ff.Compose[idx] = fr
		}
		return ff, nil
	case lang.InternalReference:
		if target, ok := parser.funcs[string(ff)]; ok {
			return parser.resolveFunc(target)
		}
		return nil, errors.Wrap(ErrFuncRefUnreachable, string(ff))
	case lang.F:
		var err error
		ff.TypeIn, err = parser.resolveType(ff.TypeIn)
		if err != nil {
			return nil, errors.Wrap(ErrFuncTypeIncorrect, ff.Name)
		}
		ff.TypeOut, err = parser.resolveType(ff.TypeOut)
		if err != nil {
			return nil, errors.Wrap(ErrFuncTypeIncorrect, ff.Name)
		}
		return ff, nil
	case lang.EitherBind:
		argF, err := parser.resolveFunc(ff.Argument)
		if err != nil {
			return nil, errors.Wrap(ErrFuncTypeIncorrect, ">>= "+ff.Argument.GetName())
		}
		ff.Argument = argF
		return ff, nil
	}
	return f, nil
}

type fContainer struct {
	f       lang.F
	a       lang.Alias
	refpath []lang.Composable
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

type gadtContainer int

const (
	modeOptional gadtContainer = iota
	modeEither
)

func (gc gadtContainer) Equal(to lang.T) bool   { return false }
func (gc gadtContainer) Subtype(of lang.T) bool { return false }
func (gc gadtContainer) Parent(of lang.T) bool  { return false }
func (gc gadtContainer) Type() lang.Type        { return -1 }

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
	l.parser.logger.Debug("->   " + ctx.GetText())
}

func (l *listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	l.parser.logger.Debug("  <- " + ctx.GetText())
}

func (l *listener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {
	l.parser.tc = &tContainer{}
}

func (l *listener) EnterTypeComplexDefinition(ctx *TypeComplexDefinitionContext) {
	l.parser.tc.t = lang.Record{
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
	if complext, ok := parentc.t.(lang.Record); ok {
		fieldName := ctx.FieldName().GetText()
		complext.Fields[fieldName] = l.parser.tc.t
	}
	l.parser.tc = parentc
}

func (l *listener) EnterTypeMinLength(ctx *TypeMinLengthContext) {
	arg, _ := strconv.Atoi(ctx.ConstructorArg().GetText())
	t := lang.Contain(nil, lang.TypeMinLength, []interface{}{arg})
	l.parser.tc.appendT(t)
}

func (l *listener) EnterTypeMaxLength(ctx *TypeMaxLengthContext) {
	arg, _ := strconv.Atoi(ctx.ConstructorArg().GetText())
	t := lang.Contain(nil, lang.TypeMaxLength, []interface{}{arg})
	l.parser.tc.appendT(t)
}

func (l *listener) EnterTypeOptional(ctx *TypeOptionalContext) {
	l.parser.tc.appendT(modeOptional)
}

func (l *listener) EnterTypeList(ctx *TypeListContext) {
	t := lang.Contain(nil, lang.TypeList, nil)
	l.parser.tc.appendT(t)
}

func (l *listener) EnterTypeEither(ctx *TypeEitherContext) {
	l.parser.tc.appendT(modeEither)
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

func (l *listener) EnterTypeFloat(ctx *TypeFloatContext) {
	l.parser.tc.appendT(lang.TypeFloat)
}

func (l *listener) EnterTypeOther(ctx *TypeOtherContext) {
	l.parser.tc.appendT(lang.Ref(ctx.GetText()))
}

func (l *listener) ExitTypeSimpleDefinition(ctx *TypeSimpleDefinitionContext) {
	l.parser.tc.t = l.reduceTpath(l.parser.tc.tpath)
	l.parser.tc.tpath = []lang.T{}
}

func (l *listener) reduceTpath(tpath []lang.T) lang.T {
	if len(tpath) == 0 {
		return nil
	}

	if gc, ok := tpath[0].(gadtContainer); ok && gc == modeEither {
		return l.reduceEitherTpath(tpath)
	}

	var finalT lang.T
	for i := len(tpath) - 1; i >= 0; i-- {
		if finalT == nil {
			finalT = tpath[i]
			continue
		}
		if container, ok := tpath[i].(lang.Container); ok {
			container.Operands = []lang.T{finalT}
			finalT = container
		}
	}

	if gc, ok := tpath[0].(gadtContainer); ok && gc == modeOptional {
		return lang.Optional(finalT)
	}

	return finalT
}

func (l *listener) reduceEitherTpath(tpath []lang.T) lang.T {
	if len(tpath) < 3 {
		return nil
	}
	operands := make([]lang.T, 2, 2)
	ptr := 1

	for i := len(tpath) - 1; i > 0; i-- {
		if operands[ptr] == nil {
			operands[ptr] = tpath[i]
			continue
		}
		if container, ok := tpath[i].(lang.Container); ok {
			container.Operands = []lang.T{operands[ptr]}
			operands[ptr] = container
		} else {
			if ptr == 1 {
				ptr = 0
				operands[ptr] = tpath[i]
			}
		}
	}

	return lang.Either(operands[0], operands[1])
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
		refpath: []lang.Composable{},
		a:       lang.Alias{},
	}
	if ctx.FuncName() != nil {
		l.parser.fc.a.Name = ctx.FuncName().GetText()
	}
}

func (l *listener) EnterFuncRef(ctx *FuncRefContext) {
	l.parser.fc.refpath = append(l.parser.fc.refpath, lang.InternalReference(ctx.GetText()))
}

func (l *listener) EnterFuncBind(ctx *FuncBindContext) {
	bind := lang.Bind(lang.InternalReference(ctx.FuncApplied().GetText()))
	l.parser.fc.refpath = append(l.parser.fc.refpath, bind)
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
	l.parser.invokes = append(l.parser.invokes, lang.InternalReference(ctx.GetText()))
}
