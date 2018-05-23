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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 104,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 6, 2, 24, 10,
	2, 13, 2, 14, 2, 25, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 46, 10,
	4, 12, 4, 14, 4, 49, 11, 4, 5, 4, 51, 10, 4, 3, 5, 3, 5, 5, 5, 55, 10,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5,
	8, 68, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 74, 10, 9, 12, 9, 14, 9, 77,
	11, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 83, 10, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 93, 10, 11, 12, 11, 14, 11, 96,
	11, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 102, 10, 11, 3, 11, 2, 2, 12,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 2, 2, 108, 2, 23, 3, 2, 2, 2, 4,
	27, 3, 2, 2, 2, 6, 50, 3, 2, 2, 2, 8, 54, 3, 2, 2, 2, 10, 56, 3, 2, 2,
	2, 12, 58, 3, 2, 2, 2, 14, 67, 3, 2, 2, 2, 16, 82, 3, 2, 2, 2, 18, 84,
	3, 2, 2, 2, 20, 101, 3, 2, 2, 2, 22, 24, 5, 4, 3, 2, 23, 22, 3, 2, 2, 2,
	24, 25, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 3, 3, 2,
	2, 2, 27, 28, 5, 6, 4, 2, 28, 5, 3, 2, 2, 2, 29, 30, 7, 3, 2, 2, 30, 31,
	5, 10, 6, 2, 31, 32, 7, 4, 2, 2, 32, 33, 5, 14, 8, 2, 33, 51, 3, 2, 2,
	2, 34, 35, 5, 12, 7, 2, 35, 36, 7, 5, 2, 2, 36, 37, 5, 8, 5, 2, 37, 38,
	7, 6, 2, 2, 38, 39, 5, 8, 5, 2, 39, 51, 3, 2, 2, 2, 40, 41, 5, 12, 7, 2,
	41, 42, 7, 4, 2, 2, 42, 47, 5, 12, 7, 2, 43, 44, 7, 7, 2, 2, 44, 46, 5,
	12, 7, 2, 45, 43, 3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47,
	48, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 29, 3, 2, 2,
	2, 50, 34, 3, 2, 2, 2, 50, 40, 3, 2, 2, 2, 51, 7, 3, 2, 2, 2, 52, 55, 5,
	10, 6, 2, 53, 55, 7, 8, 2, 2, 54, 52, 3, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55,
	9, 3, 2, 2, 2, 56, 57, 7, 18, 2, 2, 57, 11, 3, 2, 2, 2, 58, 59, 7, 19,
	2, 2, 59, 13, 3, 2, 2, 2, 60, 68, 7, 20, 2, 2, 61, 68, 7, 21, 2, 2, 62,
	68, 5, 16, 9, 2, 63, 68, 5, 20, 11, 2, 64, 68, 7, 9, 2, 2, 65, 68, 7, 10,
	2, 2, 66, 68, 7, 11, 2, 2, 67, 60, 3, 2, 2, 2, 67, 61, 3, 2, 2, 2, 67,
	62, 3, 2, 2, 2, 67, 63, 3, 2, 2, 2, 67, 64, 3, 2, 2, 2, 67, 65, 3, 2, 2,
	2, 67, 66, 3, 2, 2, 2, 68, 15, 3, 2, 2, 2, 69, 70, 7, 12, 2, 2, 70, 75,
	5, 18, 10, 2, 71, 72, 7, 13, 2, 2, 72, 74, 5, 18, 10, 2, 73, 71, 3, 2,
	2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 78,
	3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 79, 7, 14, 2, 2, 79, 83, 3, 2, 2, 2,
	80, 81, 7, 12, 2, 2, 81, 83, 7, 14, 2, 2, 82, 69, 3, 2, 2, 2, 82, 80, 3,
	2, 2, 2, 83, 17, 3, 2, 2, 2, 84, 85, 7, 20, 2, 2, 85, 86, 7, 15, 2, 2,
	86, 87, 5, 14, 8, 2, 87, 19, 3, 2, 2, 2, 88, 89, 7, 16, 2, 2, 89, 94, 5,
	14, 8, 2, 90, 91, 7, 13, 2, 2, 91, 93, 5, 14, 8, 2, 92, 90, 3, 2, 2, 2,
	93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 97, 3,
	2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 98, 7, 17, 2, 2, 98, 102, 3, 2, 2, 2,
	99, 100, 7, 16, 2, 2, 100, 102, 7, 17, 2, 2, 101, 88, 3, 2, 2, 2, 101,
	99, 3, 2, 2, 2, 102, 21, 3, 2, 2, 2, 11, 25, 47, 50, 54, 67, 75, 82, 94,
	101,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'data'", "'='", "'::'", "'->'", "'.'", "'_'", "'true'", "'false'",
	"'null'", "'{'", "','", "'}'", "':'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "DATA_NAME_ID",
	"FUNC_NAME_ID", "STRING", "NUMBER", "WS",
}

var ruleNames = []string{
	"system", "statement", "dataDef", "dataName", "dataNameId", "funcNameId",
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
	AnzerParserEOF          = antlr.TokenEOF
	AnzerParserT__0         = 1
	AnzerParserT__1         = 2
	AnzerParserT__2         = 3
	AnzerParserT__3         = 4
	AnzerParserT__4         = 5
	AnzerParserT__5         = 6
	AnzerParserT__6         = 7
	AnzerParserT__7         = 8
	AnzerParserT__8         = 9
	AnzerParserT__9         = 10
	AnzerParserT__10        = 11
	AnzerParserT__11        = 12
	AnzerParserT__12        = 13
	AnzerParserT__13        = 14
	AnzerParserT__14        = 15
	AnzerParserDATA_NAME_ID = 16
	AnzerParserFUNC_NAME_ID = 17
	AnzerParserSTRING       = 18
	AnzerParserNUMBER       = 19
	AnzerParserWS           = 20
)

// AnzerParser rules.
const (
	AnzerParserRULE_system     = 0
	AnzerParserRULE_statement  = 1
	AnzerParserRULE_dataDef    = 2
	AnzerParserRULE_dataName   = 3
	AnzerParserRULE_dataNameId = 4
	AnzerParserRULE_funcNameId = 5
	AnzerParserRULE_json       = 6
	AnzerParserRULE_obj        = 7
	AnzerParserRULE_pair       = 8
	AnzerParserRULE_array      = 9
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

func (s *SystemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AnzerVisitor:
		return t.VisitSystem(s)

	default:
		return t.VisitChildren(s)
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AnzerParserT__0 || _la == AnzerParserFUNC_NAME_ID {
		{
			p.SetState(20)
			p.Statement()
		}

		p.SetState(23)
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

func (s *StatementContext) DataDef() IDataDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataDefContext)
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

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AnzerVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.DataDef()
	}

	return localctx
}

// IDataDefContext is an interface to support dynamic dispatch.
type IDataDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataDefContext differentiates from other interfaces.
	IsDataDefContext()
}

type DataDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataDefContext() *DataDefContext {
	var p = new(DataDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_dataDef
	return p
}

func (*DataDefContext) IsDataDefContext() {}

func NewDataDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataDefContext {
	var p = new(DataDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_dataDef

	return p
}

func (s *DataDefContext) GetParser() antlr.Parser { return s.parser }

func (s *DataDefContext) DataNameId() IDataNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNameIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataNameIdContext)
}

func (s *DataDefContext) Json() IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *DataDefContext) AllFuncNameId() []IFuncNameIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncNameIdContext)(nil)).Elem())
	var tst = make([]IFuncNameIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncNameIdContext)
		}
	}

	return tst
}

func (s *DataDefContext) FuncNameId(i int) IFuncNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncNameIdContext)
}

func (s *DataDefContext) AllDataName() []IDataNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataNameContext)(nil)).Elem())
	var tst = make([]IDataNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataNameContext)
		}
	}

	return tst
}

func (s *DataDefContext) DataName(i int) IDataNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataNameContext)
}

func (s *DataDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterDataDef(s)
	}
}

func (s *DataDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitDataDef(s)
	}
}

func (s *DataDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AnzerVisitor:
		return t.VisitDataDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AnzerParser) DataDef() (localctx IDataDefContext) {
	localctx = NewDataDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AnzerParserRULE_dataDef)
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

	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Match(AnzerParserT__0)
		}
		{
			p.SetState(28)
			p.DataNameId()
		}
		{
			p.SetState(29)
			p.Match(AnzerParserT__1)
		}
		{
			p.SetState(30)
			p.Json()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.FuncNameId()
		}
		{
			p.SetState(33)
			p.Match(AnzerParserT__2)
		}
		{
			p.SetState(34)
			p.DataName()
		}
		{
			p.SetState(35)
			p.Match(AnzerParserT__3)
		}
		{
			p.SetState(36)
			p.DataName()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.FuncNameId()
		}
		{
			p.SetState(39)
			p.Match(AnzerParserT__1)
		}
		{
			p.SetState(40)
			p.FuncNameId()
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__4 {
			{
				p.SetState(41)
				p.Match(AnzerParserT__4)
			}
			{
				p.SetState(42)
				p.FuncNameId()
			}

			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

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

func (s *DataNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AnzerVisitor:
		return t.VisitDataName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AnzerParser) DataName() (localctx IDataNameContext) {
	localctx = NewDataNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AnzerParserRULE_dataName)

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
	case AnzerParserDATA_NAME_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.DataNameId()
		}

	case AnzerParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Match(AnzerParserT__5)
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

func (s *DataNameIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AnzerVisitor:
		return t.VisitDataNameId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AnzerParser) DataNameId() (localctx IDataNameIdContext) {
	localctx = NewDataNameIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AnzerParserRULE_dataNameId)

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
		p.Match(AnzerParserDATA_NAME_ID)
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

func (s *FuncNameIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AnzerVisitor:
		return t.VisitFuncNameId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AnzerParser) FuncNameId() (localctx IFuncNameIdContext) {
	localctx = NewFuncNameIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AnzerParserRULE_funcNameId)

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
		p.SetState(56)
		p.Match(AnzerParserFUNC_NAME_ID)
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

func (s *JsonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AnzerVisitor:
		return t.VisitJson(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AnzerParser) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AnzerParserRULE_json)

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

	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Match(AnzerParserSTRING)
		}

	case AnzerParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Match(AnzerParserNUMBER)
		}

	case AnzerParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.Obj()
		}

	case AnzerParserT__13:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(61)
			p.Array()
		}

	case AnzerParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(62)
			p.Match(AnzerParserT__6)
		}

	case AnzerParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(63)
			p.Match(AnzerParserT__7)
		}

	case AnzerParserT__8:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(64)
			p.Match(AnzerParserT__8)
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

func (s *ObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AnzerVisitor:
		return t.VisitObj(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AnzerParser) Obj() (localctx IObjContext) {
	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AnzerParserRULE_obj)
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

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Match(AnzerParserT__9)
		}
		{
			p.SetState(68)
			p.Pair()
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__10 {
			{
				p.SetState(69)
				p.Match(AnzerParserT__10)
			}
			{
				p.SetState(70)
				p.Pair()
			}

			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(76)
			p.Match(AnzerParserT__11)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.Match(AnzerParserT__9)
		}
		{
			p.SetState(79)
			p.Match(AnzerParserT__11)
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

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AnzerVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AnzerParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AnzerParserRULE_pair)

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
		p.SetState(82)
		p.Match(AnzerParserSTRING)
	}
	{
		p.SetState(83)
		p.Match(AnzerParserT__12)
	}
	{
		p.SetState(84)
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

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AnzerVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AnzerParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AnzerParserRULE_array)
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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(AnzerParserT__13)
		}
		{
			p.SetState(87)
			p.Json()
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__10 {
			{
				p.SetState(88)
				p.Match(AnzerParserT__10)
			}
			{
				p.SetState(89)
				p.Json()
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(95)
			p.Match(AnzerParserT__14)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Match(AnzerParserT__13)
		}
		{
			p.SetState(98)
			p.Match(AnzerParserT__14)
		}

	}

	return localctx
}
