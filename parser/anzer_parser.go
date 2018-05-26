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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 144,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 6,
	2, 36, 10, 2, 13, 2, 14, 2, 37, 3, 3, 3, 3, 3, 3, 5, 3, 43, 10, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 52, 10, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 7, 7, 59, 10, 7, 12, 7, 14, 7, 62, 11, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 5, 8, 70, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 7, 10, 79, 10, 10, 12, 10, 14, 10, 82, 11, 10, 3, 11, 3, 11,
	5, 11, 86, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 92, 10, 12, 12, 12,
	14, 12, 95, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 5, 14, 108, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15,
	7, 15, 114, 10, 15, 12, 15, 14, 15, 117, 11, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 5, 15, 123, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 7, 17, 133, 10, 17, 12, 17, 14, 17, 136, 11, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 5, 17, 142, 10, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 3, 4, 2, 11, 11, 21, 21, 2, 146,
	2, 35, 3, 2, 2, 2, 4, 42, 3, 2, 2, 2, 6, 44, 3, 2, 2, 2, 8, 51, 3, 2, 2,
	2, 10, 53, 3, 2, 2, 2, 12, 55, 3, 2, 2, 2, 14, 63, 3, 2, 2, 2, 16, 71,
	3, 2, 2, 2, 18, 73, 3, 2, 2, 2, 20, 85, 3, 2, 2, 2, 22, 87, 3, 2, 2, 2,
	24, 98, 3, 2, 2, 2, 26, 107, 3, 2, 2, 2, 28, 122, 3, 2, 2, 2, 30, 124,
	3, 2, 2, 2, 32, 141, 3, 2, 2, 2, 34, 36, 5, 4, 3, 2, 35, 34, 3, 2, 2, 2,
	36, 37, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 3, 3, 2,
	2, 2, 39, 43, 5, 6, 4, 2, 40, 43, 5, 14, 8, 2, 41, 43, 5, 18, 10, 2, 42,
	39, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 5, 3, 2, 2,
	2, 44, 45, 7, 3, 2, 2, 45, 46, 7, 21, 2, 2, 46, 47, 7, 4, 2, 2, 47, 48,
	5, 8, 5, 2, 48, 7, 3, 2, 2, 2, 49, 52, 5, 10, 6, 2, 50, 52, 5, 12, 7, 2,
	51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 9, 3, 2, 2, 2, 53, 54, 5, 26,
	14, 2, 54, 11, 3, 2, 2, 2, 55, 60, 7, 21, 2, 2, 56, 57, 7, 20, 2, 2, 57,
	59, 7, 21, 2, 2, 58, 56, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2,
	2, 2, 60, 61, 3, 2, 2, 2, 61, 13, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 64,
	7, 22, 2, 2, 64, 65, 7, 5, 2, 2, 65, 66, 5, 24, 13, 2, 66, 67, 7, 6, 2,
	2, 67, 69, 5, 24, 13, 2, 68, 70, 5, 16, 9, 2, 69, 68, 3, 2, 2, 2, 69, 70,
	3, 2, 2, 2, 70, 15, 3, 2, 2, 2, 71, 72, 5, 26, 14, 2, 72, 17, 3, 2, 2,
	2, 73, 74, 7, 22, 2, 2, 74, 75, 7, 4, 2, 2, 75, 80, 5, 20, 11, 2, 76, 77,
	7, 7, 2, 2, 77, 79, 5, 20, 11, 2, 78, 76, 3, 2, 2, 2, 79, 82, 3, 2, 2,
	2, 80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 19, 3, 2, 2, 2, 82, 80,
	3, 2, 2, 2, 83, 86, 7, 22, 2, 2, 84, 86, 5, 22, 12, 2, 85, 83, 3, 2, 2,
	2, 85, 84, 3, 2, 2, 2, 86, 21, 3, 2, 2, 2, 87, 88, 7, 8, 2, 2, 88, 93,
	5, 20, 11, 2, 89, 90, 7, 9, 2, 2, 90, 92, 5, 20, 11, 2, 91, 89, 3, 2, 2,
	2, 92, 95, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 96,
	3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 96, 97, 7, 10, 2, 2, 97, 23, 3, 2, 2, 2,
	98, 99, 9, 2, 2, 2, 99, 25, 3, 2, 2, 2, 100, 108, 7, 24, 2, 2, 101, 108,
	7, 25, 2, 2, 102, 108, 5, 28, 15, 2, 103, 108, 5, 32, 17, 2, 104, 108,
	7, 12, 2, 2, 105, 108, 7, 13, 2, 2, 106, 108, 7, 14, 2, 2, 107, 100, 3,
	2, 2, 2, 107, 101, 3, 2, 2, 2, 107, 102, 3, 2, 2, 2, 107, 103, 3, 2, 2,
	2, 107, 104, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108,
	27, 3, 2, 2, 2, 109, 110, 7, 15, 2, 2, 110, 115, 5, 30, 16, 2, 111, 112,
	7, 9, 2, 2, 112, 114, 5, 30, 16, 2, 113, 111, 3, 2, 2, 2, 114, 117, 3,
	2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 118, 3, 2, 2,
	2, 117, 115, 3, 2, 2, 2, 118, 119, 7, 16, 2, 2, 119, 123, 3, 2, 2, 2, 120,
	121, 7, 15, 2, 2, 121, 123, 7, 16, 2, 2, 122, 109, 3, 2, 2, 2, 122, 120,
	3, 2, 2, 2, 123, 29, 3, 2, 2, 2, 124, 125, 7, 24, 2, 2, 125, 126, 7, 17,
	2, 2, 126, 127, 5, 26, 14, 2, 127, 31, 3, 2, 2, 2, 128, 129, 7, 18, 2,
	2, 129, 134, 5, 26, 14, 2, 130, 131, 7, 9, 2, 2, 131, 133, 5, 26, 14, 2,
	132, 130, 3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134,
	135, 3, 2, 2, 2, 135, 137, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 137, 138,
	7, 19, 2, 2, 138, 142, 3, 2, 2, 2, 139, 140, 7, 18, 2, 2, 140, 142, 7,
	19, 2, 2, 141, 128, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 142, 33, 3, 2, 2,
	2, 15, 37, 42, 51, 60, 69, 80, 85, 93, 107, 115, 122, 134, 141,
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
	"system", "statement", "dataSig", "dataDefinition", "jsonDataDefinition",
	"recursiveDataDefinition", "funcSig", "funcParams", "funcDef", "composeFunc",
	"productFunc", "dataName", "json", "obj", "pair", "array",
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
	AnzerParserRULE_system                  = 0
	AnzerParserRULE_statement               = 1
	AnzerParserRULE_dataSig                 = 2
	AnzerParserRULE_dataDefinition          = 3
	AnzerParserRULE_jsonDataDefinition      = 4
	AnzerParserRULE_recursiveDataDefinition = 5
	AnzerParserRULE_funcSig                 = 6
	AnzerParserRULE_funcParams              = 7
	AnzerParserRULE_funcDef                 = 8
	AnzerParserRULE_composeFunc             = 9
	AnzerParserRULE_productFunc             = 10
	AnzerParserRULE_dataName                = 11
	AnzerParserRULE_json                    = 12
	AnzerParserRULE_obj                     = 13
	AnzerParserRULE_pair                    = 14
	AnzerParserRULE_array                   = 15
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AnzerParserT__0 || _la == AnzerParserFUNC_NAME_ID {
		{
			p.SetState(32)
			p.Statement()
		}

		p.SetState(35)
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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.DataSig()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.FuncSig()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
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
		p.SetState(42)
		p.Match(AnzerParserT__0)
	}
	{
		p.SetState(43)
		p.Match(AnzerParserDATA_NAME_ID)
	}
	{
		p.SetState(44)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(45)
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

func (s *DataDefinitionContext) JsonDataDefinition() IJsonDataDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonDataDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonDataDefinitionContext)
}

func (s *DataDefinitionContext) RecursiveDataDefinition() IRecursiveDataDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecursiveDataDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecursiveDataDefinitionContext)
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

	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserT__9, AnzerParserT__10, AnzerParserT__11, AnzerParserT__12, AnzerParserT__15, AnzerParserSTRING, AnzerParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.JsonDataDefinition()
		}

	case AnzerParserDATA_NAME_ID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.RecursiveDataDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJsonDataDefinitionContext is an interface to support dynamic dispatch.
type IJsonDataDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonDataDefinitionContext differentiates from other interfaces.
	IsJsonDataDefinitionContext()
}

type JsonDataDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonDataDefinitionContext() *JsonDataDefinitionContext {
	var p = new(JsonDataDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_jsonDataDefinition
	return p
}

func (*JsonDataDefinitionContext) IsJsonDataDefinitionContext() {}

func NewJsonDataDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonDataDefinitionContext {
	var p = new(JsonDataDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_jsonDataDefinition

	return p
}

func (s *JsonDataDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonDataDefinitionContext) Json() IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *JsonDataDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonDataDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonDataDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterJsonDataDefinition(s)
	}
}

func (s *JsonDataDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitJsonDataDefinition(s)
	}
}

func (p *AnzerParser) JsonDataDefinition() (localctx IJsonDataDefinitionContext) {
	localctx = NewJsonDataDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AnzerParserRULE_jsonDataDefinition)

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
		p.SetState(51)
		p.Json()
	}

	return localctx
}

// IRecursiveDataDefinitionContext is an interface to support dynamic dispatch.
type IRecursiveDataDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecursiveDataDefinitionContext differentiates from other interfaces.
	IsRecursiveDataDefinitionContext()
}

type RecursiveDataDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecursiveDataDefinitionContext() *RecursiveDataDefinitionContext {
	var p = new(RecursiveDataDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_recursiveDataDefinition
	return p
}

func (*RecursiveDataDefinitionContext) IsRecursiveDataDefinitionContext() {}

func NewRecursiveDataDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecursiveDataDefinitionContext {
	var p = new(RecursiveDataDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_recursiveDataDefinition

	return p
}

func (s *RecursiveDataDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *RecursiveDataDefinitionContext) AllDATA_NAME_ID() []antlr.TerminalNode {
	return s.GetTokens(AnzerParserDATA_NAME_ID)
}

func (s *RecursiveDataDefinitionContext) DATA_NAME_ID(i int) antlr.TerminalNode {
	return s.GetToken(AnzerParserDATA_NAME_ID, i)
}

func (s *RecursiveDataDefinitionContext) AllDATA_AND_OR() []antlr.TerminalNode {
	return s.GetTokens(AnzerParserDATA_AND_OR)
}

func (s *RecursiveDataDefinitionContext) DATA_AND_OR(i int) antlr.TerminalNode {
	return s.GetToken(AnzerParserDATA_AND_OR, i)
}

func (s *RecursiveDataDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecursiveDataDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecursiveDataDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterRecursiveDataDefinition(s)
	}
}

func (s *RecursiveDataDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitRecursiveDataDefinition(s)
	}
}

func (p *AnzerParser) RecursiveDataDefinition() (localctx IRecursiveDataDefinitionContext) {
	localctx = NewRecursiveDataDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AnzerParserRULE_recursiveDataDefinition)
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
		p.SetState(53)
		p.Match(AnzerParserDATA_NAME_ID)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserDATA_AND_OR {
		{
			p.SetState(54)
			p.Match(AnzerParserDATA_AND_OR)
		}
		{
			p.SetState(55)
			p.Match(AnzerParserDATA_NAME_ID)
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 12, AnzerParserRULE_funcSig)
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
		p.SetState(61)
		p.Match(AnzerParserFUNC_NAME_ID)
	}
	{
		p.SetState(62)
		p.Match(AnzerParserT__2)
	}
	{
		p.SetState(63)
		p.DataName()
	}
	{
		p.SetState(64)
		p.Match(AnzerParserT__3)
	}
	{
		p.SetState(65)
		p.DataName()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AnzerParserT__9)|(1<<AnzerParserT__10)|(1<<AnzerParserT__11)|(1<<AnzerParserT__12)|(1<<AnzerParserT__15)|(1<<AnzerParserSTRING)|(1<<AnzerParserNUMBER))) != 0 {
		{
			p.SetState(66)
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
	p.EnterRule(localctx, 14, AnzerParserRULE_funcParams)

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
		p.SetState(69)
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
	p.EnterRule(localctx, 16, AnzerParserRULE_funcDef)
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
		p.SetState(71)
		p.Match(AnzerParserFUNC_NAME_ID)
	}
	{
		p.SetState(72)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(73)
		p.ComposeFunc()
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserT__4 {
		{
			p.SetState(74)
			p.Match(AnzerParserT__4)
		}
		{
			p.SetState(75)
			p.ComposeFunc()
		}

		p.SetState(80)
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
	p.EnterRule(localctx, 18, AnzerParserRULE_composeFunc)

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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserFUNC_NAME_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.Match(AnzerParserFUNC_NAME_ID)
		}

	case AnzerParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
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
	p.EnterRule(localctx, 20, AnzerParserRULE_productFunc)
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
		p.SetState(85)
		p.Match(AnzerParserT__5)
	}
	{
		p.SetState(86)
		p.ComposeFunc()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserT__6 {
		{
			p.SetState(87)
			p.Match(AnzerParserT__6)
		}
		{
			p.SetState(88)
			p.ComposeFunc()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
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
	p.EnterRule(localctx, 22, AnzerParserRULE_dataName)
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
	p.SetState(96)
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
	p.EnterRule(localctx, 24, AnzerParserRULE_json)

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

	p.SetState(105)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Match(AnzerParserSTRING)
		}

	case AnzerParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Match(AnzerParserNUMBER)
		}

	case AnzerParserT__12:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			p.Obj()
		}

	case AnzerParserT__15:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(101)
			p.Array()
		}

	case AnzerParserT__9:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(102)
			p.Match(AnzerParserT__9)
		}

	case AnzerParserT__10:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(103)
			p.Match(AnzerParserT__10)
		}

	case AnzerParserT__11:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(104)
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
	p.EnterRule(localctx, 26, AnzerParserRULE_obj)
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

	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Match(AnzerParserT__12)
		}
		{
			p.SetState(108)
			p.Pair()
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__6 {
			{
				p.SetState(109)
				p.Match(AnzerParserT__6)
			}
			{
				p.SetState(110)
				p.Pair()
			}

			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(116)
			p.Match(AnzerParserT__13)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(AnzerParserT__12)
		}
		{
			p.SetState(119)
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
	p.EnterRule(localctx, 28, AnzerParserRULE_pair)

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
		p.SetState(122)
		p.Match(AnzerParserSTRING)
	}
	{
		p.SetState(123)
		p.Match(AnzerParserT__14)
	}
	{
		p.SetState(124)
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
	p.EnterRule(localctx, 30, AnzerParserRULE_array)
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

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Match(AnzerParserT__15)
		}
		{
			p.SetState(127)
			p.Json()
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__6 {
			{
				p.SetState(128)
				p.Match(AnzerParserT__6)
			}
			{
				p.SetState(129)
				p.Json()
			}

			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(135)
			p.Match(AnzerParserT__16)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Match(AnzerParserT__15)
		}
		{
			p.SetState(138)
			p.Match(AnzerParserT__16)
		}

	}

	return localctx
}
