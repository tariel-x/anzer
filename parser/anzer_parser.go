// Generated from /home/nikita/go/src/github.com/tariel-x/anzer/parser/Anzer.g4 by ANTLR 4.7.

package parser // Anzer
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 137,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 6, 2, 32, 10, 2, 13, 2, 14, 2,
	33, 3, 3, 3, 3, 3, 3, 5, 3, 39, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 7, 5, 50, 10, 5, 12, 5, 14, 5, 53, 11, 5, 5, 5, 55,
	10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 63, 10, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 72, 10, 8, 12, 8, 14, 8, 75, 11, 8,
	3, 9, 3, 9, 5, 9, 79, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 85, 10,
	10, 12, 10, 14, 10, 88, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 101, 10, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 7, 13, 107, 10, 13, 12, 13, 14, 13, 110, 11, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 5, 13, 116, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 7, 15, 126, 10, 15, 12, 15, 14, 15, 129, 11, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 5, 15, 135, 10, 15, 3, 15, 2, 2, 16, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 2, 3, 4, 2, 11, 11, 21, 21, 2, 141,
	2, 31, 3, 2, 2, 2, 4, 38, 3, 2, 2, 2, 6, 40, 3, 2, 2, 2, 8, 54, 3, 2, 2,
	2, 10, 56, 3, 2, 2, 2, 12, 64, 3, 2, 2, 2, 14, 66, 3, 2, 2, 2, 16, 78,
	3, 2, 2, 2, 18, 80, 3, 2, 2, 2, 20, 91, 3, 2, 2, 2, 22, 100, 3, 2, 2, 2,
	24, 115, 3, 2, 2, 2, 26, 117, 3, 2, 2, 2, 28, 134, 3, 2, 2, 2, 30, 32,
	5, 4, 3, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2,
	33, 34, 3, 2, 2, 2, 34, 3, 3, 2, 2, 2, 35, 39, 5, 6, 4, 2, 36, 39, 5, 10,
	6, 2, 37, 39, 5, 14, 8, 2, 38, 35, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38,
	37, 3, 2, 2, 2, 39, 5, 3, 2, 2, 2, 40, 41, 7, 3, 2, 2, 41, 42, 7, 21, 2,
	2, 42, 43, 7, 4, 2, 2, 43, 44, 5, 8, 5, 2, 44, 7, 3, 2, 2, 2, 45, 55, 5,
	22, 12, 2, 46, 51, 7, 21, 2, 2, 47, 48, 7, 20, 2, 2, 48, 50, 7, 21, 2,
	2, 49, 47, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 52,
	3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 45, 3, 2, 2, 2,
	54, 46, 3, 2, 2, 2, 55, 9, 3, 2, 2, 2, 56, 57, 7, 22, 2, 2, 57, 58, 7,
	5, 2, 2, 58, 59, 5, 20, 11, 2, 59, 60, 7, 6, 2, 2, 60, 62, 5, 20, 11, 2,
	61, 63, 5, 12, 7, 2, 62, 61, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 11, 3,
	2, 2, 2, 64, 65, 5, 22, 12, 2, 65, 13, 3, 2, 2, 2, 66, 67, 7, 22, 2, 2,
	67, 68, 7, 4, 2, 2, 68, 73, 5, 16, 9, 2, 69, 70, 7, 7, 2, 2, 70, 72, 5,
	16, 9, 2, 71, 69, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73,
	74, 3, 2, 2, 2, 74, 15, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 79, 7, 22,
	2, 2, 77, 79, 5, 18, 10, 2, 78, 76, 3, 2, 2, 2, 78, 77, 3, 2, 2, 2, 79,
	17, 3, 2, 2, 2, 80, 81, 7, 8, 2, 2, 81, 86, 5, 16, 9, 2, 82, 83, 7, 9,
	2, 2, 83, 85, 5, 16, 9, 2, 84, 82, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86,
	84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 2, 2, 2, 88, 86, 3, 2, 2,
	2, 89, 90, 7, 10, 2, 2, 90, 19, 3, 2, 2, 2, 91, 92, 9, 2, 2, 2, 92, 21,
	3, 2, 2, 2, 93, 101, 7, 24, 2, 2, 94, 101, 7, 25, 2, 2, 95, 101, 5, 24,
	13, 2, 96, 101, 5, 28, 15, 2, 97, 101, 7, 12, 2, 2, 98, 101, 7, 13, 2,
	2, 99, 101, 7, 14, 2, 2, 100, 93, 3, 2, 2, 2, 100, 94, 3, 2, 2, 2, 100,
	95, 3, 2, 2, 2, 100, 96, 3, 2, 2, 2, 100, 97, 3, 2, 2, 2, 100, 98, 3, 2,
	2, 2, 100, 99, 3, 2, 2, 2, 101, 23, 3, 2, 2, 2, 102, 103, 7, 15, 2, 2,
	103, 108, 5, 26, 14, 2, 104, 105, 7, 9, 2, 2, 105, 107, 5, 26, 14, 2, 106,
	104, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109,
	3, 2, 2, 2, 109, 111, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 112, 7, 16,
	2, 2, 112, 116, 3, 2, 2, 2, 113, 114, 7, 15, 2, 2, 114, 116, 7, 16, 2,
	2, 115, 102, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 25, 3, 2, 2, 2, 117,
	118, 7, 24, 2, 2, 118, 119, 7, 17, 2, 2, 119, 120, 5, 22, 12, 2, 120, 27,
	3, 2, 2, 2, 121, 122, 7, 18, 2, 2, 122, 127, 5, 22, 12, 2, 123, 124, 7,
	9, 2, 2, 124, 126, 5, 22, 12, 2, 125, 123, 3, 2, 2, 2, 126, 129, 3, 2,
	2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 130, 3, 2, 2, 2,
	129, 127, 3, 2, 2, 2, 130, 131, 7, 19, 2, 2, 131, 135, 3, 2, 2, 2, 132,
	133, 7, 18, 2, 2, 133, 135, 7, 19, 2, 2, 134, 121, 3, 2, 2, 2, 134, 132,
	3, 2, 2, 2, 135, 29, 3, 2, 2, 2, 15, 33, 38, 51, 54, 62, 73, 78, 86, 100,
	108, 115, 127, 134,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'data'", "'='", "'::'", "'->'", "'.'", "'<'", "','", "'>'", "'_'",
	"'true'", "'false'", "'null'", "'{'", "'}'", "':'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"DATA_AND_OR", "DATA_NAME_ID", "FUNC_NAME_ID", "FUNC_PARAM_ID", "STRING",
	"NUMBER", "WS", "LineComment",
}

var ruleNames = []string{
	"system", "statement", "dataSig", "dataDefinition", "funcSig", "funcParams",
	"funcDef", "composeFunc", "productFunc", "dataName", "json", "obj", "pair",
	"array",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AnzerParser struct {
	*antlr.BaseParser
}

func NewAnzerParser(input antlr.TokenStream) *AnzerParser {
	this := new(AnzerParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Anzer.g4"

	return this
}

// AnzerParser tokens.
const (
	AnzerParserEOF           = antlr.TokenEOF
	AnzerParserT__0          = 1
	AnzerParserT__1          = 2
	AnzerParserT__2          = 3
	AnzerParserT__3          = 4
	AnzerParserT__4          = 5
	AnzerParserT__5          = 6
	AnzerParserT__6          = 7
	AnzerParserT__7          = 8
	AnzerParserT__8          = 9
	AnzerParserT__9          = 10
	AnzerParserT__10         = 11
	AnzerParserT__11         = 12
	AnzerParserT__12         = 13
	AnzerParserT__13         = 14
	AnzerParserT__14         = 15
	AnzerParserT__15         = 16
	AnzerParserT__16         = 17
	AnzerParserDATA_AND_OR   = 18
	AnzerParserDATA_NAME_ID  = 19
	AnzerParserFUNC_NAME_ID  = 20
	AnzerParserFUNC_PARAM_ID = 21
	AnzerParserSTRING        = 22
	AnzerParserNUMBER        = 23
	AnzerParserWS            = 24
	AnzerParserLineComment   = 25
)

// AnzerParser rules.
const (
	AnzerParserRULE_system         = 0
	AnzerParserRULE_statement      = 1
	AnzerParserRULE_dataSig        = 2
	AnzerParserRULE_dataDefinition = 3
	AnzerParserRULE_funcSig        = 4
	AnzerParserRULE_funcParams     = 5
	AnzerParserRULE_funcDef        = 6
	AnzerParserRULE_composeFunc    = 7
	AnzerParserRULE_productFunc    = 8
	AnzerParserRULE_dataName       = 9
	AnzerParserRULE_json           = 10
	AnzerParserRULE_obj            = 11
	AnzerParserRULE_pair           = 12
	AnzerParserRULE_array          = 13
)

// ISystemContext is an interface to support dynamic dispatch.
type ISystemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSystemContext differentiates from other interfaces.
	IsSystemContext()
}

type SystemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySystemContext() *SystemContext {
	var p = new(SystemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_system
	return p
}

func (*SystemContext) IsSystemContext() {}

func NewSystemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SystemContext {
	var p = new(SystemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_system

	return p
}

func (s *SystemContext) GetParser() antlr.Parser { return s.parser }

func (s *SystemContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *SystemContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SystemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SystemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SystemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterSystem(s)
	}
}

func (s *SystemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitSystem(s)
	}
}

func (p *AnzerParser) System() (localctx ISystemContext) {
	localctx = NewSystemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AnzerParserRULE_system)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AnzerParserT__0 || _la == AnzerParserFUNC_NAME_ID {
		{
			p.SetState(28)
			p.Statement()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) DataSig() IDataSigContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataSigContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataSigContext)
}

func (s *StatementContext) FuncSig() IFuncSigContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncSigContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncSigContext)
}

func (s *StatementContext) FuncDef() IFuncDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncDefContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *AnzerParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AnzerParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.DataSig()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.FuncSig()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.FuncDef()
		}

	}

	return localctx
}

// IDataSigContext is an interface to support dynamic dispatch.
type IDataSigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataSigContext differentiates from other interfaces.
	IsDataSigContext()
}

type DataSigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataSigContext() *DataSigContext {
	var p = new(DataSigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_dataSig
	return p
}

func (*DataSigContext) IsDataSigContext() {}

func NewDataSigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataSigContext {
	var p = new(DataSigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_dataSig

	return p
}

func (s *DataSigContext) GetParser() antlr.Parser { return s.parser }

func (s *DataSigContext) DATA_NAME_ID() antlr.TerminalNode {
	return s.GetToken(AnzerParserDATA_NAME_ID, 0)
}

func (s *DataSigContext) DataDefinition() IDataDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataDefinitionContext)
}

func (s *DataSigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataSigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataSigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterDataSig(s)
	}
}

func (s *DataSigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitDataSig(s)
	}
}

func (p *AnzerParser) DataSig() (localctx IDataSigContext) {
	localctx = NewDataSigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AnzerParserRULE_dataSig)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(AnzerParserT__0)
	}
	{
		p.SetState(39)
		p.Match(AnzerParserDATA_NAME_ID)
	}
	{
		p.SetState(40)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(41)
		p.DataDefinition()
	}

	return localctx
}

// IDataDefinitionContext is an interface to support dynamic dispatch.
type IDataDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataDefinitionContext differentiates from other interfaces.
	IsDataDefinitionContext()
}

type DataDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataDefinitionContext() *DataDefinitionContext {
	var p = new(DataDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_dataDefinition
	return p
}

func (*DataDefinitionContext) IsDataDefinitionContext() {}

func NewDataDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataDefinitionContext {
	var p = new(DataDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_dataDefinition

	return p
}

func (s *DataDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DataDefinitionContext) Json() IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *DataDefinitionContext) AllDATA_NAME_ID() []antlr.TerminalNode {
	return s.GetTokens(AnzerParserDATA_NAME_ID)
}

func (s *DataDefinitionContext) DATA_NAME_ID(i int) antlr.TerminalNode {
	return s.GetToken(AnzerParserDATA_NAME_ID, i)
}

func (s *DataDefinitionContext) AllDATA_AND_OR() []antlr.TerminalNode {
	return s.GetTokens(AnzerParserDATA_AND_OR)
}

func (s *DataDefinitionContext) DATA_AND_OR(i int) antlr.TerminalNode {
	return s.GetToken(AnzerParserDATA_AND_OR, i)
}

func (s *DataDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterDataDefinition(s)
	}
}

func (s *DataDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitDataDefinition(s)
	}
}

func (p *AnzerParser) DataDefinition() (localctx IDataDefinitionContext) {
	localctx = NewDataDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AnzerParserRULE_dataDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(52)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserT__9, AnzerParserT__10, AnzerParserT__11, AnzerParserT__12, AnzerParserT__15, AnzerParserSTRING, AnzerParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.Json()
		}

	case AnzerParserDATA_NAME_ID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(AnzerParserDATA_NAME_ID)
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserDATA_AND_OR {
			{
				p.SetState(45)
				p.Match(AnzerParserDATA_AND_OR)
			}
			{
				p.SetState(46)
				p.Match(AnzerParserDATA_NAME_ID)
			}

			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFuncSigContext is an interface to support dynamic dispatch.
type IFuncSigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncSigContext differentiates from other interfaces.
	IsFuncSigContext()
}

type FuncSigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSigContext() *FuncSigContext {
	var p = new(FuncSigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcSig
	return p
}

func (*FuncSigContext) IsFuncSigContext() {}

func NewFuncSigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSigContext {
	var p = new(FuncSigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcSig

	return p
}

func (s *FuncSigContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSigContext) FUNC_NAME_ID() antlr.TerminalNode {
	return s.GetToken(AnzerParserFUNC_NAME_ID, 0)
}

func (s *FuncSigContext) AllDataName() []IDataNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataNameContext)(nil)).Elem())
	var tst = make([]IDataNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataNameContext)
		}
	}

	return tst
}

func (s *FuncSigContext) DataName(i int) IDataNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataNameContext)
}

func (s *FuncSigContext) FuncParams() IFuncParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncParamsContext)
}

func (s *FuncSigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncSig(s)
	}
}

func (s *FuncSigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncSig(s)
	}
}

func (p *AnzerParser) FuncSig() (localctx IFuncSigContext) {
	localctx = NewFuncSigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AnzerParserRULE_funcSig)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(AnzerParserFUNC_NAME_ID)
	}
	{
		p.SetState(55)
		p.Match(AnzerParserT__2)
	}
	{
		p.SetState(56)
		p.DataName()
	}
	{
		p.SetState(57)
		p.Match(AnzerParserT__3)
	}
	{
		p.SetState(58)
		p.DataName()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AnzerParserT__9)|(1<<AnzerParserT__10)|(1<<AnzerParserT__11)|(1<<AnzerParserT__12)|(1<<AnzerParserT__15)|(1<<AnzerParserSTRING)|(1<<AnzerParserNUMBER))) != 0 {
		{
			p.SetState(59)
			p.FuncParams()
		}

	}

	return localctx
}

// IFuncParamsContext is an interface to support dynamic dispatch.
type IFuncParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncParamsContext differentiates from other interfaces.
	IsFuncParamsContext()
}

type FuncParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamsContext() *FuncParamsContext {
	var p = new(FuncParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcParams
	return p
}

func (*FuncParamsContext) IsFuncParamsContext() {}

func NewFuncParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamsContext {
	var p = new(FuncParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcParams

	return p
}

func (s *FuncParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamsContext) Json() IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *FuncParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncParams(s)
	}
}

func (s *FuncParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncParams(s)
	}
}

func (p *AnzerParser) FuncParams() (localctx IFuncParamsContext) {
	localctx = NewFuncParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AnzerParserRULE_funcParams)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Json()
	}

	return localctx
}

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDefContext differentiates from other interfaces.
	IsFuncDefContext()
}

type FuncDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefContext() *FuncDefContext {
	var p = new(FuncDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcDef
	return p
}

func (*FuncDefContext) IsFuncDefContext() {}

func NewFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefContext {
	var p = new(FuncDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcDef

	return p
}

func (s *FuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefContext) FUNC_NAME_ID() antlr.TerminalNode {
	return s.GetToken(AnzerParserFUNC_NAME_ID, 0)
}

func (s *FuncDefContext) AllComposeFunc() []IComposeFuncContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComposeFuncContext)(nil)).Elem())
	var tst = make([]IComposeFuncContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComposeFuncContext)
		}
	}

	return tst
}

func (s *FuncDefContext) ComposeFunc(i int) IComposeFuncContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComposeFuncContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComposeFuncContext)
}

func (s *FuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncDef(s)
	}
}

func (s *FuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncDef(s)
	}
}

func (p *AnzerParser) FuncDef() (localctx IFuncDefContext) {
	localctx = NewFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AnzerParserRULE_funcDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(AnzerParserFUNC_NAME_ID)
	}
	{
		p.SetState(65)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(66)
		p.ComposeFunc()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserT__4 {
		{
			p.SetState(67)
			p.Match(AnzerParserT__4)
		}
		{
			p.SetState(68)
			p.ComposeFunc()
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IComposeFuncContext is an interface to support dynamic dispatch.
type IComposeFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComposeFuncContext differentiates from other interfaces.
	IsComposeFuncContext()
}

type ComposeFuncContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComposeFuncContext() *ComposeFuncContext {
	var p = new(ComposeFuncContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_composeFunc
	return p
}

func (*ComposeFuncContext) IsComposeFuncContext() {}

func NewComposeFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComposeFuncContext {
	var p = new(ComposeFuncContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_composeFunc

	return p
}

func (s *ComposeFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *ComposeFuncContext) FUNC_NAME_ID() antlr.TerminalNode {
	return s.GetToken(AnzerParserFUNC_NAME_ID, 0)
}

func (s *ComposeFuncContext) ProductFunc() IProductFuncContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProductFuncContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProductFuncContext)
}

func (s *ComposeFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComposeFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComposeFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterComposeFunc(s)
	}
}

func (s *ComposeFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitComposeFunc(s)
	}
}

func (p *AnzerParser) ComposeFunc() (localctx IComposeFuncContext) {
	localctx = NewComposeFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AnzerParserRULE_composeFunc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserFUNC_NAME_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Match(AnzerParserFUNC_NAME_ID)
		}

	case AnzerParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.ProductFunc()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IProductFuncContext is an interface to support dynamic dispatch.
type IProductFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProductFuncContext differentiates from other interfaces.
	IsProductFuncContext()
}

type ProductFuncContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProductFuncContext() *ProductFuncContext {
	var p = new(ProductFuncContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_productFunc
	return p
}

func (*ProductFuncContext) IsProductFuncContext() {}

func NewProductFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProductFuncContext {
	var p = new(ProductFuncContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_productFunc

	return p
}

func (s *ProductFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *ProductFuncContext) AllComposeFunc() []IComposeFuncContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComposeFuncContext)(nil)).Elem())
	var tst = make([]IComposeFuncContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComposeFuncContext)
		}
	}

	return tst
}

func (s *ProductFuncContext) ComposeFunc(i int) IComposeFuncContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComposeFuncContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComposeFuncContext)
}

func (s *ProductFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProductFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProductFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterProductFunc(s)
	}
}

func (s *ProductFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitProductFunc(s)
	}
}

func (p *AnzerParser) ProductFunc() (localctx IProductFuncContext) {
	localctx = NewProductFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AnzerParserRULE_productFunc)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(AnzerParserT__5)
	}
	{
		p.SetState(79)
		p.ComposeFunc()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserT__6 {
		{
			p.SetState(80)
			p.Match(AnzerParserT__6)
		}
		{
			p.SetState(81)
			p.ComposeFunc()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(AnzerParserT__7)
	}

	return localctx
}

// IDataNameContext is an interface to support dynamic dispatch.
type IDataNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataNameContext differentiates from other interfaces.
	IsDataNameContext()
}

type DataNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataNameContext() *DataNameContext {
	var p = new(DataNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_dataName
	return p
}

func (*DataNameContext) IsDataNameContext() {}

func NewDataNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataNameContext {
	var p = new(DataNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_dataName

	return p
}

func (s *DataNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DataNameContext) DATA_NAME_ID() antlr.TerminalNode {
	return s.GetToken(AnzerParserDATA_NAME_ID, 0)
}

func (s *DataNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterDataName(s)
	}
}

func (s *DataNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitDataName(s)
	}
}

func (p *AnzerParser) DataName() (localctx IDataNameContext) {
	localctx = NewDataNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AnzerParserRULE_dataName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	_la = p.GetTokenStream().LA(1)

	if !(_la == AnzerParserT__8 || _la == AnzerParserDATA_NAME_ID) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IJsonContext is an interface to support dynamic dispatch.
type IJsonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonContext differentiates from other interfaces.
	IsJsonContext()
}

type JsonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContext() *JsonContext {
	var p = new(JsonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_json
	return p
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_json

	return p
}

func (s *JsonContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonContext) STRING() antlr.TerminalNode {
	return s.GetToken(AnzerParserSTRING, 0)
}

func (s *JsonContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AnzerParserNUMBER, 0)
}

func (s *JsonContext) Obj() IObjContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *JsonContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterJson(s)
	}
}

func (s *JsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitJson(s)
	}
}

func (p *AnzerParser) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AnzerParserRULE_json)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.Match(AnzerParserSTRING)
		}

	case AnzerParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Match(AnzerParserNUMBER)
		}

	case AnzerParserT__12:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Obj()
		}

	case AnzerParserT__15:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(94)
			p.Array()
		}

	case AnzerParserT__9:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(95)
			p.Match(AnzerParserT__9)
		}

	case AnzerParserT__10:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(96)
			p.Match(AnzerParserT__10)
		}

	case AnzerParserT__11:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(97)
			p.Match(AnzerParserT__11)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *ObjContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitObj(s)
	}
}

func (p *AnzerParser) Obj() (localctx IObjContext) {
	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AnzerParserRULE_obj)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Match(AnzerParserT__12)
		}
		{
			p.SetState(101)
			p.Pair()
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__6 {
			{
				p.SetState(102)
				p.Match(AnzerParserT__6)
			}
			{
				p.SetState(103)
				p.Pair()
			}

			p.SetState(108)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(109)
			p.Match(AnzerParserT__13)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Match(AnzerParserT__12)
		}
		{
			p.SetState(112)
			p.Match(AnzerParserT__13)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(AnzerParserSTRING, 0)
}

func (s *PairContext) Json() IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *AnzerParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AnzerParserRULE_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(AnzerParserSTRING)
	}
	{
		p.SetState(116)
		p.Match(AnzerParserT__14)
	}
	{
		p.SetState(117)
		p.Json()
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllJson() []IJsonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJsonContext)(nil)).Elem())
	var tst = make([]IJsonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonContext)
		}
	}

	return tst
}

func (s *ArrayContext) Json(i int) IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *AnzerParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AnzerParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(AnzerParserT__15)
		}
		{
			p.SetState(120)
			p.Json()
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__6 {
			{
				p.SetState(121)
				p.Match(AnzerParserT__6)
			}
			{
				p.SetState(122)
				p.Json()
			}

			p.SetState(127)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(128)
			p.Match(AnzerParserT__16)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Match(AnzerParserT__15)
		}
		{
			p.SetState(131)
			p.Match(AnzerParserT__16)
		}

	}

	return localctx
}
