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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 169,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 6, 2, 42, 10, 2, 13, 2, 14, 2, 43,
	3, 3, 3, 3, 3, 3, 5, 3, 49, 10, 3, 3, 4, 3, 4, 5, 4, 53, 10, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 65, 10, 6, 3,
	7, 3, 7, 5, 7, 69, 10, 7, 3, 8, 3, 8, 3, 8, 7, 8, 74, 10, 8, 12, 8, 14,
	8, 77, 11, 8, 3, 9, 3, 9, 3, 9, 7, 9, 82, 10, 9, 12, 9, 14, 9, 85, 11,
	9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 95, 10,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 104, 10, 13,
	12, 13, 14, 13, 107, 11, 13, 3, 14, 3, 14, 5, 14, 111, 10, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 7, 15, 117, 10, 15, 12, 15, 14, 15, 120, 11, 15, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5,
	17, 133, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 139, 10, 18, 12, 18,
	14, 18, 142, 11, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 148, 10, 18, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 158, 10, 20,
	12, 20, 14, 20, 161, 11, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 167, 10,
	20, 3, 20, 2, 2, 21, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 2, 3, 4, 2, 13, 13, 22, 22, 2, 171, 2, 41, 3, 2, 2,
	2, 4, 48, 3, 2, 2, 2, 6, 52, 3, 2, 2, 2, 8, 54, 3, 2, 2, 2, 10, 59, 3,
	2, 2, 2, 12, 68, 3, 2, 2, 2, 14, 70, 3, 2, 2, 2, 16, 78, 3, 2, 2, 2, 18,
	86, 3, 2, 2, 2, 20, 88, 3, 2, 2, 2, 22, 96, 3, 2, 2, 2, 24, 98, 3, 2, 2,
	2, 26, 110, 3, 2, 2, 2, 28, 112, 3, 2, 2, 2, 30, 123, 3, 2, 2, 2, 32, 132,
	3, 2, 2, 2, 34, 147, 3, 2, 2, 2, 36, 149, 3, 2, 2, 2, 38, 166, 3, 2, 2,
	2, 40, 42, 5, 4, 3, 2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 41,
	3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 3, 3, 2, 2, 2, 45, 49, 5, 6, 4, 2,
	46, 49, 5, 20, 11, 2, 47, 49, 5, 24, 13, 2, 48, 45, 3, 2, 2, 2, 48, 46,
	3, 2, 2, 2, 48, 47, 3, 2, 2, 2, 49, 5, 3, 2, 2, 2, 50, 53, 5, 8, 5, 2,
	51, 53, 5, 10, 6, 2, 52, 50, 3, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 7, 3,
	2, 2, 2, 54, 55, 7, 3, 2, 2, 55, 56, 7, 22, 2, 2, 56, 57, 7, 4, 2, 2, 57,
	58, 5, 32, 17, 2, 58, 9, 3, 2, 2, 2, 59, 60, 7, 3, 2, 2, 60, 61, 7, 22,
	2, 2, 61, 64, 7, 4, 2, 2, 62, 65, 5, 16, 9, 2, 63, 65, 5, 14, 8, 2, 64,
	62, 3, 2, 2, 2, 64, 63, 3, 2, 2, 2, 65, 11, 3, 2, 2, 2, 66, 69, 5, 18,
	10, 2, 67, 69, 5, 32, 17, 2, 68, 66, 3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69,
	13, 3, 2, 2, 2, 70, 75, 5, 18, 10, 2, 71, 72, 7, 5, 2, 2, 72, 74, 5, 18,
	10, 2, 73, 71, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75,
	76, 3, 2, 2, 2, 76, 15, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 83, 5, 18,
	10, 2, 79, 80, 7, 6, 2, 2, 80, 82, 5, 18, 10, 2, 81, 79, 3, 2, 2, 2, 82,
	85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 17, 3, 2, 2,
	2, 85, 83, 3, 2, 2, 2, 86, 87, 7, 22, 2, 2, 87, 19, 3, 2, 2, 2, 88, 89,
	7, 23, 2, 2, 89, 90, 7, 7, 2, 2, 90, 91, 5, 30, 16, 2, 91, 92, 7, 8, 2,
	2, 92, 94, 5, 30, 16, 2, 93, 95, 5, 22, 12, 2, 94, 93, 3, 2, 2, 2, 94,
	95, 3, 2, 2, 2, 95, 21, 3, 2, 2, 2, 96, 97, 5, 32, 17, 2, 97, 23, 3, 2,
	2, 2, 98, 99, 7, 23, 2, 2, 99, 100, 7, 4, 2, 2, 100, 105, 5, 26, 14, 2,
	101, 102, 7, 9, 2, 2, 102, 104, 5, 26, 14, 2, 103, 101, 3, 2, 2, 2, 104,
	107, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 25, 3,
	2, 2, 2, 107, 105, 3, 2, 2, 2, 108, 111, 7, 23, 2, 2, 109, 111, 5, 28,
	15, 2, 110, 108, 3, 2, 2, 2, 110, 109, 3, 2, 2, 2, 111, 27, 3, 2, 2, 2,
	112, 113, 7, 10, 2, 2, 113, 118, 5, 26, 14, 2, 114, 115, 7, 11, 2, 2, 115,
	117, 5, 26, 14, 2, 116, 114, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116,
	3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 121, 3, 2, 2, 2, 120, 118, 3, 2,
	2, 2, 121, 122, 7, 12, 2, 2, 122, 29, 3, 2, 2, 2, 123, 124, 9, 2, 2, 2,
	124, 31, 3, 2, 2, 2, 125, 133, 7, 25, 2, 2, 126, 133, 7, 26, 2, 2, 127,
	133, 5, 34, 18, 2, 128, 133, 5, 38, 20, 2, 129, 133, 7, 14, 2, 2, 130,
	133, 7, 15, 2, 2, 131, 133, 7, 16, 2, 2, 132, 125, 3, 2, 2, 2, 132, 126,
	3, 2, 2, 2, 132, 127, 3, 2, 2, 2, 132, 128, 3, 2, 2, 2, 132, 129, 3, 2,
	2, 2, 132, 130, 3, 2, 2, 2, 132, 131, 3, 2, 2, 2, 133, 33, 3, 2, 2, 2,
	134, 135, 7, 17, 2, 2, 135, 140, 5, 36, 19, 2, 136, 137, 7, 11, 2, 2, 137,
	139, 5, 36, 19, 2, 138, 136, 3, 2, 2, 2, 139, 142, 3, 2, 2, 2, 140, 138,
	3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 143, 3, 2, 2, 2, 142, 140, 3, 2,
	2, 2, 143, 144, 7, 18, 2, 2, 144, 148, 3, 2, 2, 2, 145, 146, 7, 17, 2,
	2, 146, 148, 7, 18, 2, 2, 147, 134, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148,
	35, 3, 2, 2, 2, 149, 150, 7, 25, 2, 2, 150, 151, 7, 19, 2, 2, 151, 152,
	5, 32, 17, 2, 152, 37, 3, 2, 2, 2, 153, 154, 7, 20, 2, 2, 154, 159, 5,
	32, 17, 2, 155, 156, 7, 11, 2, 2, 156, 158, 5, 32, 17, 2, 157, 155, 3,
	2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2,
	2, 160, 162, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 162, 163, 7, 21, 2, 2, 163,
	167, 3, 2, 2, 2, 164, 165, 7, 20, 2, 2, 165, 167, 7, 21, 2, 2, 166, 153,
	3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 167, 39, 3, 2, 2, 2, 18, 43, 48, 52,
	64, 68, 75, 83, 94, 105, 110, 118, 132, 140, 147, 159, 166,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'data'", "'='", "'|'", "'&'", "'::'", "'->'", "'.'", "'<'", "','",
	"'>'", "'_'", "'true'", "'false'", "'null'", "'{'", "'}'", "':'", "'['",
	"']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "DATA_NAME_ID", "FUNC_NAME_ID", "FUNC_PARAM_ID", "STRING", "NUMBER",
	"WS", "LineComment",
}

var ruleNames = []string{
	"system", "statement", "dataSig", "jsonDataDef", "logicDataDef", "dataArg",
	"dataOr", "dataAnd", "dataNameId", "funcSig", "funcParams", "funcDef",
	"composeFunc", "productFunc", "dataName", "json", "obj", "pair", "array",
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
	AnzerParserT__17         = 18
	AnzerParserT__18         = 19
	AnzerParserDATA_NAME_ID  = 20
	AnzerParserFUNC_NAME_ID  = 21
	AnzerParserFUNC_PARAM_ID = 22
	AnzerParserSTRING        = 23
	AnzerParserNUMBER        = 24
	AnzerParserWS            = 25
	AnzerParserLineComment   = 26
)

// AnzerParser rules.
const (
	AnzerParserRULE_system       = 0
	AnzerParserRULE_statement    = 1
	AnzerParserRULE_dataSig      = 2
	AnzerParserRULE_jsonDataDef  = 3
	AnzerParserRULE_logicDataDef = 4
	AnzerParserRULE_dataArg      = 5
	AnzerParserRULE_dataOr       = 6
	AnzerParserRULE_dataAnd      = 7
	AnzerParserRULE_dataNameId   = 8
	AnzerParserRULE_funcSig      = 9
	AnzerParserRULE_funcParams   = 10
	AnzerParserRULE_funcDef      = 11
	AnzerParserRULE_composeFunc  = 12
	AnzerParserRULE_productFunc  = 13
	AnzerParserRULE_dataName     = 14
	AnzerParserRULE_json         = 15
	AnzerParserRULE_obj          = 16
	AnzerParserRULE_pair         = 17
	AnzerParserRULE_array        = 18
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
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AnzerParserT__0 || _la == AnzerParserFUNC_NAME_ID {
		{
			p.SetState(38)
			p.Statement()
		}

		p.SetState(41)
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

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.DataSig()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.FuncSig()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)
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

func (s *DataSigContext) JsonDataDef() IJsonDataDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonDataDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonDataDefContext)
}

func (s *DataSigContext) LogicDataDef() ILogicDataDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicDataDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicDataDefContext)
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

	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.JsonDataDef()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.LogicDataDef()
		}

	}

	return localctx
}

// IJsonDataDefContext is an interface to support dynamic dispatch.
type IJsonDataDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonDataDefContext differentiates from other interfaces.
	IsJsonDataDefContext()
}

type JsonDataDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonDataDefContext() *JsonDataDefContext {
	var p = new(JsonDataDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_jsonDataDef
	return p
}

func (*JsonDataDefContext) IsJsonDataDefContext() {}

func NewJsonDataDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonDataDefContext {
	var p = new(JsonDataDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_jsonDataDef

	return p
}

func (s *JsonDataDefContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonDataDefContext) DATA_NAME_ID() antlr.TerminalNode {
	return s.GetToken(AnzerParserDATA_NAME_ID, 0)
}

func (s *JsonDataDefContext) Json() IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *JsonDataDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonDataDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonDataDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterJsonDataDef(s)
	}
}

func (s *JsonDataDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitJsonDataDef(s)
	}
}

func (p *AnzerParser) JsonDataDef() (localctx IJsonDataDefContext) {
	localctx = NewJsonDataDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AnzerParserRULE_jsonDataDef)

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
		p.SetState(52)
		p.Match(AnzerParserT__0)
	}
	{
		p.SetState(53)
		p.Match(AnzerParserDATA_NAME_ID)
	}
	{
		p.SetState(54)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(55)
		p.Json()
	}

	return localctx
}

// ILogicDataDefContext is an interface to support dynamic dispatch.
type ILogicDataDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicDataDefContext differentiates from other interfaces.
	IsLogicDataDefContext()
}

type LogicDataDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicDataDefContext() *LogicDataDefContext {
	var p = new(LogicDataDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_logicDataDef
	return p
}

func (*LogicDataDefContext) IsLogicDataDefContext() {}

func NewLogicDataDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicDataDefContext {
	var p = new(LogicDataDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_logicDataDef

	return p
}

func (s *LogicDataDefContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicDataDefContext) DATA_NAME_ID() antlr.TerminalNode {
	return s.GetToken(AnzerParserDATA_NAME_ID, 0)
}

func (s *LogicDataDefContext) DataAnd() IDataAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataAndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataAndContext)
}

func (s *LogicDataDefContext) DataOr() IDataOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataOrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataOrContext)
}

func (s *LogicDataDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicDataDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicDataDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterLogicDataDef(s)
	}
}

func (s *LogicDataDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitLogicDataDef(s)
	}
}

func (p *AnzerParser) LogicDataDef() (localctx ILogicDataDefContext) {
	localctx = NewLogicDataDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AnzerParserRULE_logicDataDef)

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
		p.SetState(57)
		p.Match(AnzerParserT__0)
	}
	{
		p.SetState(58)
		p.Match(AnzerParserDATA_NAME_ID)
	}
	{
		p.SetState(59)
		p.Match(AnzerParserT__1)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(60)
			p.DataAnd()
		}

	case 2:
		{
			p.SetState(61)
			p.DataOr()
		}

	}

	return localctx
}

// IDataArgContext is an interface to support dynamic dispatch.
type IDataArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataArgContext differentiates from other interfaces.
	IsDataArgContext()
}

type DataArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataArgContext() *DataArgContext {
	var p = new(DataArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_dataArg
	return p
}

func (*DataArgContext) IsDataArgContext() {}

func NewDataArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataArgContext {
	var p = new(DataArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_dataArg

	return p
}

func (s *DataArgContext) GetParser() antlr.Parser { return s.parser }

func (s *DataArgContext) DataNameId() IDataNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNameIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataNameIdContext)
}

func (s *DataArgContext) Json() IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *DataArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterDataArg(s)
	}
}

func (s *DataArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitDataArg(s)
	}
}

func (p *AnzerParser) DataArg() (localctx IDataArgContext) {
	localctx = NewDataArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AnzerParserRULE_dataArg)

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

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserDATA_NAME_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.DataNameId()
		}

	case AnzerParserT__11, AnzerParserT__12, AnzerParserT__13, AnzerParserT__14, AnzerParserT__17, AnzerParserSTRING, AnzerParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Json()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDataOrContext is an interface to support dynamic dispatch.
type IDataOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataOrContext differentiates from other interfaces.
	IsDataOrContext()
}

type DataOrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataOrContext() *DataOrContext {
	var p = new(DataOrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_dataOr
	return p
}

func (*DataOrContext) IsDataOrContext() {}

func NewDataOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataOrContext {
	var p = new(DataOrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_dataOr

	return p
}

func (s *DataOrContext) GetParser() antlr.Parser { return s.parser }

func (s *DataOrContext) AllDataNameId() []IDataNameIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataNameIdContext)(nil)).Elem())
	var tst = make([]IDataNameIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataNameIdContext)
		}
	}

	return tst
}

func (s *DataOrContext) DataNameId(i int) IDataNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNameIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataNameIdContext)
}

func (s *DataOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterDataOr(s)
	}
}

func (s *DataOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitDataOr(s)
	}
}

func (p *AnzerParser) DataOr() (localctx IDataOrContext) {
	localctx = NewDataOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AnzerParserRULE_dataOr)
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
		p.SetState(68)
		p.DataNameId()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserT__2 {
		{
			p.SetState(69)
			p.Match(AnzerParserT__2)
		}
		{
			p.SetState(70)
			p.DataNameId()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDataAndContext is an interface to support dynamic dispatch.
type IDataAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataAndContext differentiates from other interfaces.
	IsDataAndContext()
}

type DataAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataAndContext() *DataAndContext {
	var p = new(DataAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_dataAnd
	return p
}

func (*DataAndContext) IsDataAndContext() {}

func NewDataAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataAndContext {
	var p = new(DataAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_dataAnd

	return p
}

func (s *DataAndContext) GetParser() antlr.Parser { return s.parser }

func (s *DataAndContext) AllDataNameId() []IDataNameIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataNameIdContext)(nil)).Elem())
	var tst = make([]IDataNameIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataNameIdContext)
		}
	}

	return tst
}

func (s *DataAndContext) DataNameId(i int) IDataNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNameIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataNameIdContext)
}

func (s *DataAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterDataAnd(s)
	}
}

func (s *DataAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitDataAnd(s)
	}
}

func (p *AnzerParser) DataAnd() (localctx IDataAndContext) {
	localctx = NewDataAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AnzerParserRULE_dataAnd)
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
		p.SetState(76)
		p.DataNameId()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserT__3 {
		{
			p.SetState(77)
			p.Match(AnzerParserT__3)
		}
		{
			p.SetState(78)
			p.DataNameId()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDataNameIdContext is an interface to support dynamic dispatch.
type IDataNameIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataNameIdContext differentiates from other interfaces.
	IsDataNameIdContext()
}

type DataNameIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataNameIdContext() *DataNameIdContext {
	var p = new(DataNameIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_dataNameId
	return p
}

func (*DataNameIdContext) IsDataNameIdContext() {}

func NewDataNameIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataNameIdContext {
	var p = new(DataNameIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_dataNameId

	return p
}

func (s *DataNameIdContext) GetParser() antlr.Parser { return s.parser }

func (s *DataNameIdContext) DATA_NAME_ID() antlr.TerminalNode {
	return s.GetToken(AnzerParserDATA_NAME_ID, 0)
}

func (s *DataNameIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataNameIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataNameIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterDataNameId(s)
	}
}

func (s *DataNameIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitDataNameId(s)
	}
}

func (p *AnzerParser) DataNameId() (localctx IDataNameIdContext) {
	localctx = NewDataNameIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AnzerParserRULE_dataNameId)

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
		p.SetState(84)
		p.Match(AnzerParserDATA_NAME_ID)
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
	p.EnterRule(localctx, 18, AnzerParserRULE_funcSig)
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
		p.SetState(86)
		p.Match(AnzerParserFUNC_NAME_ID)
	}
	{
		p.SetState(87)
		p.Match(AnzerParserT__4)
	}
	{
		p.SetState(88)
		p.DataName()
	}
	{
		p.SetState(89)
		p.Match(AnzerParserT__5)
	}
	{
		p.SetState(90)
		p.DataName()
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AnzerParserT__11)|(1<<AnzerParserT__12)|(1<<AnzerParserT__13)|(1<<AnzerParserT__14)|(1<<AnzerParserT__17)|(1<<AnzerParserSTRING)|(1<<AnzerParserNUMBER))) != 0 {
		{
			p.SetState(91)
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
	p.EnterRule(localctx, 20, AnzerParserRULE_funcParams)

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
		p.SetState(94)
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
	p.EnterRule(localctx, 22, AnzerParserRULE_funcDef)
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
		p.SetState(96)
		p.Match(AnzerParserFUNC_NAME_ID)
	}
	{
		p.SetState(97)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(98)
		p.ComposeFunc()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserT__6 {
		{
			p.SetState(99)
			p.Match(AnzerParserT__6)
		}
		{
			p.SetState(100)
			p.ComposeFunc()
		}

		p.SetState(105)
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
	p.EnterRule(localctx, 24, AnzerParserRULE_composeFunc)

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

	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserFUNC_NAME_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(AnzerParserFUNC_NAME_ID)
		}

	case AnzerParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
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
	p.EnterRule(localctx, 26, AnzerParserRULE_productFunc)
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
		p.SetState(110)
		p.Match(AnzerParserT__7)
	}
	{
		p.SetState(111)
		p.ComposeFunc()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserT__8 {
		{
			p.SetState(112)
			p.Match(AnzerParserT__8)
		}
		{
			p.SetState(113)
			p.ComposeFunc()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(119)
		p.Match(AnzerParserT__9)
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
	p.EnterRule(localctx, 28, AnzerParserRULE_dataName)
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
	p.SetState(121)
	_la = p.GetTokenStream().LA(1)

	if !(_la == AnzerParserT__10 || _la == AnzerParserDATA_NAME_ID) {
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
	p.EnterRule(localctx, 30, AnzerParserRULE_json)

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

	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(AnzerParserSTRING)
		}

	case AnzerParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Match(AnzerParserNUMBER)
		}

	case AnzerParserT__14:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Obj()
		}

	case AnzerParserT__17:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)
			p.Array()
		}

	case AnzerParserT__11:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(127)
			p.Match(AnzerParserT__11)
		}

	case AnzerParserT__12:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(128)
			p.Match(AnzerParserT__12)
		}

	case AnzerParserT__13:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(129)
			p.Match(AnzerParserT__13)
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
	p.EnterRule(localctx, 32, AnzerParserRULE_obj)
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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Match(AnzerParserT__14)
		}
		{
			p.SetState(133)
			p.Pair()
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__8 {
			{
				p.SetState(134)
				p.Match(AnzerParserT__8)
			}
			{
				p.SetState(135)
				p.Pair()
			}

			p.SetState(140)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(141)
			p.Match(AnzerParserT__15)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.Match(AnzerParserT__14)
		}
		{
			p.SetState(144)
			p.Match(AnzerParserT__15)
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
	p.EnterRule(localctx, 34, AnzerParserRULE_pair)

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
		p.SetState(147)
		p.Match(AnzerParserSTRING)
	}
	{
		p.SetState(148)
		p.Match(AnzerParserT__16)
	}
	{
		p.SetState(149)
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
	p.EnterRule(localctx, 36, AnzerParserRULE_array)
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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.Match(AnzerParserT__17)
		}
		{
			p.SetState(152)
			p.Json()
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__8 {
			{
				p.SetState(153)
				p.Match(AnzerParserT__8)
			}
			{
				p.SetState(154)
				p.Json()
			}

			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(160)
			p.Match(AnzerParserT__18)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(AnzerParserT__17)
		}
		{
			p.SetState(163)
			p.Match(AnzerParserT__18)
		}

	}

	return localctx
}
