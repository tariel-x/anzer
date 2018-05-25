// Code generated from Anzer.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 169,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 6, 2,
	46, 10, 2, 13, 2, 14, 2, 47, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 54, 10, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 65, 10, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 7, 9, 80, 10, 9, 12, 9, 14, 9, 83, 11, 9, 3, 10, 3, 10, 5, 10, 87, 10,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 107, 10,
	13, 12, 13, 14, 13, 110, 11, 13, 3, 13, 3, 13, 5, 13, 114, 10, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 133, 10, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 7, 20, 139, 10, 20, 12, 20, 14, 20, 142, 11, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 5, 20, 148, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 22, 3, 22, 7, 22, 158, 10, 22, 12, 22, 14, 22, 161, 11, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 5, 22, 167, 10, 22, 3, 22, 2, 2, 23, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	2, 2, 2, 166, 2, 45, 3, 2, 2, 2, 4, 53, 3, 2, 2, 2, 6, 55, 3, 2, 2, 2,
	8, 60, 3, 2, 2, 2, 10, 64, 3, 2, 2, 2, 12, 66, 3, 2, 2, 2, 14, 68, 3, 2,
	2, 2, 16, 74, 3, 2, 2, 2, 18, 86, 3, 2, 2, 2, 20, 88, 3, 2, 2, 2, 22, 95,
	3, 2, 2, 2, 24, 113, 3, 2, 2, 2, 26, 115, 3, 2, 2, 2, 28, 117, 3, 2, 2,
	2, 30, 119, 3, 2, 2, 2, 32, 121, 3, 2, 2, 2, 34, 123, 3, 2, 2, 2, 36, 132,
	3, 2, 2, 2, 38, 147, 3, 2, 2, 2, 40, 149, 3, 2, 2, 2, 42, 166, 3, 2, 2,
	2, 44, 46, 5, 4, 3, 2, 45, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 45,
	3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 3, 3, 2, 2, 2, 49, 54, 5, 6, 4, 2,
	50, 54, 5, 14, 8, 2, 51, 54, 5, 16, 9, 2, 52, 54, 5, 18, 10, 2, 53, 49,
	3, 2, 2, 2, 53, 50, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 52, 3, 2, 2, 2,
	54, 5, 3, 2, 2, 2, 55, 56, 7, 3, 2, 2, 56, 57, 5, 12, 7, 2, 57, 58, 7,
	4, 2, 2, 58, 59, 5, 8, 5, 2, 59, 7, 3, 2, 2, 2, 60, 61, 5, 36, 19, 2, 61,
	9, 3, 2, 2, 2, 62, 65, 5, 12, 7, 2, 63, 65, 7, 5, 2, 2, 64, 62, 3, 2, 2,
	2, 64, 63, 3, 2, 2, 2, 65, 11, 3, 2, 2, 2, 66, 67, 7, 21, 2, 2, 67, 13,
	3, 2, 2, 2, 68, 69, 5, 28, 15, 2, 69, 70, 7, 6, 2, 2, 70, 71, 5, 10, 6,
	2, 71, 72, 7, 7, 2, 2, 72, 73, 5, 10, 6, 2, 73, 15, 3, 2, 2, 2, 74, 75,
	5, 26, 14, 2, 75, 76, 7, 4, 2, 2, 76, 81, 5, 24, 13, 2, 77, 78, 7, 8, 2,
	2, 78, 80, 5, 24, 13, 2, 79, 77, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 17, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2,
	84, 87, 5, 22, 12, 2, 85, 87, 5, 20, 11, 2, 86, 84, 3, 2, 2, 2, 86, 85,
	3, 2, 2, 2, 87, 19, 3, 2, 2, 2, 88, 89, 5, 26, 14, 2, 89, 90, 7, 9, 2,
	2, 90, 91, 5, 32, 17, 2, 91, 92, 7, 10, 2, 2, 92, 93, 7, 4, 2, 2, 93, 94,
	5, 34, 18, 2, 94, 21, 3, 2, 2, 2, 95, 96, 5, 26, 14, 2, 96, 97, 7, 8, 2,
	2, 97, 98, 5, 30, 16, 2, 98, 99, 7, 4, 2, 2, 99, 100, 5, 34, 18, 2, 100,
	23, 3, 2, 2, 2, 101, 114, 5, 26, 14, 2, 102, 103, 7, 11, 2, 2, 103, 108,
	5, 26, 14, 2, 104, 105, 7, 12, 2, 2, 105, 107, 5, 26, 14, 2, 106, 104,
	3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2,
	2, 2, 109, 111, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 112, 7, 13, 2, 2,
	112, 114, 3, 2, 2, 2, 113, 101, 3, 2, 2, 2, 113, 102, 3, 2, 2, 2, 114,
	25, 3, 2, 2, 2, 115, 116, 5, 28, 15, 2, 116, 27, 3, 2, 2, 2, 117, 118,
	7, 22, 2, 2, 118, 29, 3, 2, 2, 2, 119, 120, 7, 23, 2, 2, 120, 31, 3, 2,
	2, 2, 121, 122, 7, 24, 2, 2, 122, 33, 3, 2, 2, 2, 123, 124, 7, 24, 2, 2,
	124, 35, 3, 2, 2, 2, 125, 133, 7, 24, 2, 2, 126, 133, 7, 25, 2, 2, 127,
	133, 5, 38, 20, 2, 128, 133, 5, 42, 22, 2, 129, 133, 7, 14, 2, 2, 130,
	133, 7, 15, 2, 2, 131, 133, 7, 16, 2, 2, 132, 125, 3, 2, 2, 2, 132, 126,
	3, 2, 2, 2, 132, 127, 3, 2, 2, 2, 132, 128, 3, 2, 2, 2, 132, 129, 3, 2,
	2, 2, 132, 130, 3, 2, 2, 2, 132, 131, 3, 2, 2, 2, 133, 37, 3, 2, 2, 2,
	134, 135, 7, 17, 2, 2, 135, 140, 5, 40, 21, 2, 136, 137, 7, 12, 2, 2, 137,
	139, 5, 40, 21, 2, 138, 136, 3, 2, 2, 2, 139, 142, 3, 2, 2, 2, 140, 138,
	3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 143, 3, 2, 2, 2, 142, 140, 3, 2,
	2, 2, 143, 144, 7, 18, 2, 2, 144, 148, 3, 2, 2, 2, 145, 146, 7, 17, 2,
	2, 146, 148, 7, 18, 2, 2, 147, 134, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148,
	39, 3, 2, 2, 2, 149, 150, 7, 24, 2, 2, 150, 151, 7, 19, 2, 2, 151, 152,
	5, 36, 19, 2, 152, 41, 3, 2, 2, 2, 153, 154, 7, 20, 2, 2, 154, 159, 5,
	36, 19, 2, 155, 156, 7, 12, 2, 2, 156, 158, 5, 36, 19, 2, 157, 155, 3,
	2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2,
	2, 160, 162, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 162, 163, 7, 10, 2, 2, 163,
	167, 3, 2, 2, 2, 164, 165, 7, 20, 2, 2, 165, 167, 7, 10, 2, 2, 166, 153,
	3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 167, 43, 3, 2, 2, 2, 14, 47, 53, 64,
	81, 86, 108, 113, 132, 140, 147, 159, 166,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'data'", "'='", "'_'", "'::'", "'->'", "'.'", "'.env['", "']'", "'<'",
	"','", "'>'", "'true'", "'false'", "'null'", "'{'", "'}'", "':'", "'['",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "DATA_NAME_ID", "FUNC_NAME_ID", "FUNC_PARAM_ID", "STRING", "NUMBER",
	"WS", "LineComment",
}

var ruleNames = []string{
	"system", "statement", "dataSig", "dataDefinition", "dataName", "dataNameId",
	"funcSig", "funcDef", "funcParam", "funcParamEnv", "funcParamConfig", "composeFunc",
	"funcName", "funcNameId", "funcParamId", "funcEnvName", "funcParamValue",
	"json", "obj", "pair", "array",
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
	AnzerParserRULE_system          = 0
	AnzerParserRULE_statement       = 1
	AnzerParserRULE_dataSig         = 2
	AnzerParserRULE_dataDefinition  = 3
	AnzerParserRULE_dataName        = 4
	AnzerParserRULE_dataNameId      = 5
	AnzerParserRULE_funcSig         = 6
	AnzerParserRULE_funcDef         = 7
	AnzerParserRULE_funcParam       = 8
	AnzerParserRULE_funcParamEnv    = 9
	AnzerParserRULE_funcParamConfig = 10
	AnzerParserRULE_composeFunc     = 11
	AnzerParserRULE_funcName        = 12
	AnzerParserRULE_funcNameId      = 13
	AnzerParserRULE_funcParamId     = 14
	AnzerParserRULE_funcEnvName     = 15
	AnzerParserRULE_funcParamValue  = 16
	AnzerParserRULE_json            = 17
	AnzerParserRULE_obj             = 18
	AnzerParserRULE_pair            = 19
	AnzerParserRULE_array           = 20
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
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AnzerParserT__0 || _la == AnzerParserFUNC_NAME_ID {
		{
			p.SetState(42)
			p.Statement()
		}

		p.SetState(45)
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

func (s *StatementContext) FuncParam() IFuncParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncParamContext)
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

	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.DataSig()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.FuncSig()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)
			p.FuncDef()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(50)
			p.FuncParam()
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

func (s *DataSigContext) DataNameId() IDataNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNameIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataNameIdContext)
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
		p.SetState(53)
		p.Match(AnzerParserT__0)
	}
	{
		p.SetState(54)
		p.DataNameId()
	}
	{
		p.SetState(55)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(56)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Json()
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

func (s *DataNameContext) DataNameId() IDataNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNameIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataNameIdContext)
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
	p.EnterRule(localctx, 8, AnzerParserRULE_dataName)

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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserDATA_NAME_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.DataNameId()
		}

	case AnzerParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Match(AnzerParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 10, AnzerParserRULE_dataNameId)

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

func (s *FuncSigContext) FuncNameId() IFuncNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameIdContext)
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
		p.SetState(66)
		p.FuncNameId()
	}
	{
		p.SetState(67)
		p.Match(AnzerParserT__3)
	}
	{
		p.SetState(68)
		p.DataName()
	}
	{
		p.SetState(69)
		p.Match(AnzerParserT__4)
	}
	{
		p.SetState(70)
		p.DataName()
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

func (s *FuncDefContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
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
	p.EnterRule(localctx, 14, AnzerParserRULE_funcDef)
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
		p.SetState(72)
		p.FuncName()
	}
	{
		p.SetState(73)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(74)
		p.ComposeFunc()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserT__5 {
		{
			p.SetState(75)
			p.Match(AnzerParserT__5)
		}
		{
			p.SetState(76)
			p.ComposeFunc()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFuncParamContext is an interface to support dynamic dispatch.
type IFuncParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncParamContext differentiates from other interfaces.
	IsFuncParamContext()
}

type FuncParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamContext() *FuncParamContext {
	var p = new(FuncParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcParam
	return p
}

func (*FuncParamContext) IsFuncParamContext() {}

func NewFuncParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamContext {
	var p = new(FuncParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcParam

	return p
}

func (s *FuncParamContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamContext) FuncParamConfig() IFuncParamConfigContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamConfigContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncParamConfigContext)
}

func (s *FuncParamContext) FuncParamEnv() IFuncParamEnvContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamEnvContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncParamEnvContext)
}

func (s *FuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncParam(s)
	}
}

func (s *FuncParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncParam(s)
	}
}

func (p *AnzerParser) FuncParam() (localctx IFuncParamContext) {
	localctx = NewFuncParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AnzerParserRULE_funcParam)

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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.FuncParamConfig()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.FuncParamEnv()
		}

	}

	return localctx
}

// IFuncParamEnvContext is an interface to support dynamic dispatch.
type IFuncParamEnvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncParamEnvContext differentiates from other interfaces.
	IsFuncParamEnvContext()
}

type FuncParamEnvContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamEnvContext() *FuncParamEnvContext {
	var p = new(FuncParamEnvContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcParamEnv
	return p
}

func (*FuncParamEnvContext) IsFuncParamEnvContext() {}

func NewFuncParamEnvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamEnvContext {
	var p = new(FuncParamEnvContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcParamEnv

	return p
}

func (s *FuncParamEnvContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamEnvContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *FuncParamEnvContext) FuncEnvName() IFuncEnvNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncEnvNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncEnvNameContext)
}

func (s *FuncParamEnvContext) FuncParamValue() IFuncParamValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncParamValueContext)
}

func (s *FuncParamEnvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamEnvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamEnvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncParamEnv(s)
	}
}

func (s *FuncParamEnvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncParamEnv(s)
	}
}

func (p *AnzerParser) FuncParamEnv() (localctx IFuncParamEnvContext) {
	localctx = NewFuncParamEnvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AnzerParserRULE_funcParamEnv)

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
		p.FuncName()
	}
	{
		p.SetState(87)
		p.Match(AnzerParserT__6)
	}
	{
		p.SetState(88)
		p.FuncEnvName()
	}
	{
		p.SetState(89)
		p.Match(AnzerParserT__7)
	}
	{
		p.SetState(90)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(91)
		p.FuncParamValue()
	}

	return localctx
}

// IFuncParamConfigContext is an interface to support dynamic dispatch.
type IFuncParamConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncParamConfigContext differentiates from other interfaces.
	IsFuncParamConfigContext()
}

type FuncParamConfigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamConfigContext() *FuncParamConfigContext {
	var p = new(FuncParamConfigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcParamConfig
	return p
}

func (*FuncParamConfigContext) IsFuncParamConfigContext() {}

func NewFuncParamConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamConfigContext {
	var p = new(FuncParamConfigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcParamConfig

	return p
}

func (s *FuncParamConfigContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamConfigContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *FuncParamConfigContext) FuncParamId() IFuncParamIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncParamIdContext)
}

func (s *FuncParamConfigContext) FuncParamValue() IFuncParamValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncParamValueContext)
}

func (s *FuncParamConfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamConfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamConfigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncParamConfig(s)
	}
}

func (s *FuncParamConfigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncParamConfig(s)
	}
}

func (p *AnzerParser) FuncParamConfig() (localctx IFuncParamConfigContext) {
	localctx = NewFuncParamConfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AnzerParserRULE_funcParamConfig)

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
		p.SetState(93)
		p.FuncName()
	}
	{
		p.SetState(94)
		p.Match(AnzerParserT__5)
	}
	{
		p.SetState(95)
		p.FuncParamId()
	}
	{
		p.SetState(96)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(97)
		p.FuncParamValue()
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

func (s *ComposeFuncContext) AllFuncName() []IFuncNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncNameContext)(nil)).Elem())
	var tst = make([]IFuncNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncNameContext)
		}
	}

	return tst
}

func (s *ComposeFuncContext) FuncName(i int) IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
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
	p.EnterRule(localctx, 22, AnzerParserRULE_composeFunc)
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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserFUNC_NAME_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.FuncName()
		}

	case AnzerParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.Match(AnzerParserT__8)
		}
		{
			p.SetState(101)
			p.FuncName()
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__9 {
			{
				p.SetState(102)
				p.Match(AnzerParserT__9)
			}
			{
				p.SetState(103)
				p.FuncName()
			}

			p.SetState(108)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(109)
			p.Match(AnzerParserT__10)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFuncNameContext is an interface to support dynamic dispatch.
type IFuncNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncNameContext differentiates from other interfaces.
	IsFuncNameContext()
}

type FuncNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncNameContext() *FuncNameContext {
	var p = new(FuncNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcName
	return p
}

func (*FuncNameContext) IsFuncNameContext() {}

func NewFuncNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncNameContext {
	var p = new(FuncNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcName

	return p
}

func (s *FuncNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncNameContext) FuncNameId() IFuncNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameIdContext)
}

func (s *FuncNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncName(s)
	}
}

func (s *FuncNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncName(s)
	}
}

func (p *AnzerParser) FuncName() (localctx IFuncNameContext) {
	localctx = NewFuncNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AnzerParserRULE_funcName)

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
		p.SetState(113)
		p.FuncNameId()
	}

	return localctx
}

// IFuncNameIdContext is an interface to support dynamic dispatch.
type IFuncNameIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncNameIdContext differentiates from other interfaces.
	IsFuncNameIdContext()
}

type FuncNameIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncNameIdContext() *FuncNameIdContext {
	var p = new(FuncNameIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcNameId
	return p
}

func (*FuncNameIdContext) IsFuncNameIdContext() {}

func NewFuncNameIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncNameIdContext {
	var p = new(FuncNameIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcNameId

	return p
}

func (s *FuncNameIdContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncNameIdContext) FUNC_NAME_ID() antlr.TerminalNode {
	return s.GetToken(AnzerParserFUNC_NAME_ID, 0)
}

func (s *FuncNameIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncNameIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncNameIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncNameId(s)
	}
}

func (s *FuncNameIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncNameId(s)
	}
}

func (p *AnzerParser) FuncNameId() (localctx IFuncNameIdContext) {
	localctx = NewFuncNameIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AnzerParserRULE_funcNameId)

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
		p.Match(AnzerParserFUNC_NAME_ID)
	}

	return localctx
}

// IFuncParamIdContext is an interface to support dynamic dispatch.
type IFuncParamIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncParamIdContext differentiates from other interfaces.
	IsFuncParamIdContext()
}

type FuncParamIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamIdContext() *FuncParamIdContext {
	var p = new(FuncParamIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcParamId
	return p
}

func (*FuncParamIdContext) IsFuncParamIdContext() {}

func NewFuncParamIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamIdContext {
	var p = new(FuncParamIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcParamId

	return p
}

func (s *FuncParamIdContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamIdContext) FUNC_PARAM_ID() antlr.TerminalNode {
	return s.GetToken(AnzerParserFUNC_PARAM_ID, 0)
}

func (s *FuncParamIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncParamId(s)
	}
}

func (s *FuncParamIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncParamId(s)
	}
}

func (p *AnzerParser) FuncParamId() (localctx IFuncParamIdContext) {
	localctx = NewFuncParamIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AnzerParserRULE_funcParamId)

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
		p.SetState(117)
		p.Match(AnzerParserFUNC_PARAM_ID)
	}

	return localctx
}

// IFuncEnvNameContext is an interface to support dynamic dispatch.
type IFuncEnvNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncEnvNameContext differentiates from other interfaces.
	IsFuncEnvNameContext()
}

type FuncEnvNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncEnvNameContext() *FuncEnvNameContext {
	var p = new(FuncEnvNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcEnvName
	return p
}

func (*FuncEnvNameContext) IsFuncEnvNameContext() {}

func NewFuncEnvNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncEnvNameContext {
	var p = new(FuncEnvNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcEnvName

	return p
}

func (s *FuncEnvNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncEnvNameContext) STRING() antlr.TerminalNode {
	return s.GetToken(AnzerParserSTRING, 0)
}

func (s *FuncEnvNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncEnvNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncEnvNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncEnvName(s)
	}
}

func (s *FuncEnvNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncEnvName(s)
	}
}

func (p *AnzerParser) FuncEnvName() (localctx IFuncEnvNameContext) {
	localctx = NewFuncEnvNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AnzerParserRULE_funcEnvName)

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
		p.SetState(119)
		p.Match(AnzerParserSTRING)
	}

	return localctx
}

// IFuncParamValueContext is an interface to support dynamic dispatch.
type IFuncParamValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncParamValueContext differentiates from other interfaces.
	IsFuncParamValueContext()
}

type FuncParamValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamValueContext() *FuncParamValueContext {
	var p = new(FuncParamValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcParamValue
	return p
}

func (*FuncParamValueContext) IsFuncParamValueContext() {}

func NewFuncParamValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamValueContext {
	var p = new(FuncParamValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcParamValue

	return p
}

func (s *FuncParamValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(AnzerParserSTRING, 0)
}

func (s *FuncParamValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncParamValue(s)
	}
}

func (s *FuncParamValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncParamValue(s)
	}
}

func (p *AnzerParser) FuncParamValue() (localctx IFuncParamValueContext) {
	localctx = NewFuncParamValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AnzerParserRULE_funcParamValue)

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
		p.SetState(121)
		p.Match(AnzerParserSTRING)
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
	p.EnterRule(localctx, 34, AnzerParserRULE_json)

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
	p.EnterRule(localctx, 36, AnzerParserRULE_obj)
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
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

		for _la == AnzerParserT__9 {
			{
				p.SetState(134)
				p.Match(AnzerParserT__9)
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
	p.EnterRule(localctx, 38, AnzerParserRULE_pair)

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
	p.EnterRule(localctx, 40, AnzerParserRULE_array)
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
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

		for _la == AnzerParserT__9 {
			{
				p.SetState(153)
				p.Match(AnzerParserT__9)
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
			p.Match(AnzerParserT__7)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(AnzerParserT__17)
		}
		{
			p.SetState(163)
			p.Match(AnzerParserT__7)
		}

	}

	return localctx
}
