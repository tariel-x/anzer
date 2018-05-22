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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 49, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 6, 2, 18, 10, 2, 13, 2, 14, 2, 19, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 33, 10, 3, 3, 4, 3,
	4, 3, 5, 3, 5, 5, 5, 39, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 2, 2, 44, 2, 17, 3,
	2, 2, 2, 4, 32, 3, 2, 2, 2, 6, 34, 3, 2, 2, 2, 8, 38, 3, 2, 2, 2, 10, 40,
	3, 2, 2, 2, 12, 42, 3, 2, 2, 2, 14, 46, 3, 2, 2, 2, 16, 18, 5, 4, 3, 2,
	17, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3,
	2, 2, 2, 20, 3, 3, 2, 2, 2, 21, 22, 7, 3, 2, 2, 22, 23, 5, 10, 6, 2, 23,
	24, 7, 4, 2, 2, 24, 25, 5, 12, 7, 2, 25, 33, 3, 2, 2, 2, 26, 27, 5, 6,
	4, 2, 27, 28, 7, 5, 2, 2, 28, 29, 5, 8, 5, 2, 29, 30, 7, 6, 2, 2, 30, 31,
	5, 8, 5, 2, 31, 33, 3, 2, 2, 2, 32, 21, 3, 2, 2, 2, 32, 26, 3, 2, 2, 2,
	33, 5, 3, 2, 2, 2, 34, 35, 7, 9, 2, 2, 35, 7, 3, 2, 2, 2, 36, 39, 5, 10,
	6, 2, 37, 39, 7, 7, 2, 2, 38, 36, 3, 2, 2, 2, 38, 37, 3, 2, 2, 2, 39, 9,
	3, 2, 2, 2, 40, 41, 7, 10, 2, 2, 41, 11, 3, 2, 2, 2, 42, 43, 7, 8, 2, 2,
	43, 44, 5, 14, 8, 2, 44, 45, 7, 8, 2, 2, 45, 13, 3, 2, 2, 2, 46, 47, 7,
	11, 2, 2, 47, 15, 3, 2, 2, 2, 5, 19, 32, 38,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'data'", "'='", "'::'", "'->'", "'_'", "'\"'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "FUNC_NAME", "DATA_NAME_ID", "JSON_CONTENT",
	"WS",
}

var ruleNames = []string{
	"system", "statement", "funcName", "dataName", "dataNameId", "dataContent",
	"jsonContent",
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
	AnzerParserFUNC_NAME    = 7
	AnzerParserDATA_NAME_ID = 8
	AnzerParserJSON_CONTENT = 9
	AnzerParserWS           = 10
)

// AnzerParser rules.
const (
	AnzerParserRULE_system      = 0
	AnzerParserRULE_statement   = 1
	AnzerParserRULE_funcName    = 2
	AnzerParserRULE_dataName    = 3
	AnzerParserRULE_dataNameId  = 4
	AnzerParserRULE_dataContent = 5
	AnzerParserRULE_jsonContent = 6
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AnzerParserT__0 || _la == AnzerParserFUNC_NAME {
		{
			p.SetState(14)
			p.Statement()
		}

		p.SetState(17)
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

func (s *StatementContext) DataNameId() IDataNameIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNameIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataNameIdContext)
}

func (s *StatementContext) DataContent() IDataContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataContentContext)
}

func (s *StatementContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *StatementContext) AllDataName() []IDataNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataNameContext)(nil)).Elem())
	var tst = make([]IDataNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataNameContext)
		}
	}

	return tst
}

func (s *StatementContext) DataName(i int) IDataNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataNameContext)
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

	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
			p.Match(AnzerParserT__0)
		}
		{
			p.SetState(20)
			p.DataNameId()
		}
		{
			p.SetState(21)
			p.Match(AnzerParserT__1)
		}
		{
			p.SetState(22)
			p.DataContent()
		}

	case AnzerParserFUNC_NAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.FuncName()
		}
		{
			p.SetState(25)
			p.Match(AnzerParserT__2)
		}
		{
			p.SetState(26)
			p.DataName()
		}
		{
			p.SetState(27)
			p.Match(AnzerParserT__3)
		}
		{
			p.SetState(28)
			p.DataName()
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

func (s *FuncNameContext) FUNC_NAME() antlr.TerminalNode {
	return s.GetToken(AnzerParserFUNC_NAME, 0)
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
	p.EnterRule(localctx, 4, AnzerParserRULE_funcName)

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
		p.SetState(32)
		p.Match(AnzerParserFUNC_NAME)
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

	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserDATA_NAME_ID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.DataNameId()
		}

	case AnzerParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Match(AnzerParserT__4)
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
		p.SetState(38)
		p.Match(AnzerParserDATA_NAME_ID)
	}

	return localctx
}

// IDataContentContext is an interface to support dynamic dispatch.
type IDataContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataContentContext differentiates from other interfaces.
	IsDataContentContext()
}

type DataContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataContentContext() *DataContentContext {
	var p = new(DataContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_dataContent
	return p
}

func (*DataContentContext) IsDataContentContext() {}

func NewDataContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContentContext {
	var p = new(DataContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_dataContent

	return p
}

func (s *DataContentContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContentContext) JsonContent() IJsonContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJsonContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContentContext)
}

func (s *DataContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterDataContent(s)
	}
}

func (s *DataContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitDataContent(s)
	}
}

func (p *AnzerParser) DataContent() (localctx IDataContentContext) {
	localctx = NewDataContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AnzerParserRULE_dataContent)

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
		p.SetState(40)
		p.Match(AnzerParserT__5)
	}
	{
		p.SetState(41)
		p.JsonContent()
	}
	{
		p.SetState(42)
		p.Match(AnzerParserT__5)
	}

	return localctx
}

// IJsonContentContext is an interface to support dynamic dispatch.
type IJsonContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonContentContext differentiates from other interfaces.
	IsJsonContentContext()
}

type JsonContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContentContext() *JsonContentContext {
	var p = new(JsonContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_jsonContent
	return p
}

func (*JsonContentContext) IsJsonContentContext() {}

func NewJsonContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContentContext {
	var p = new(JsonContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_jsonContent

	return p
}

func (s *JsonContentContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonContentContext) JSON_CONTENT() antlr.TerminalNode {
	return s.GetToken(AnzerParserJSON_CONTENT, 0)
}

func (s *JsonContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterJsonContent(s)
	}
}

func (s *JsonContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitJsonContent(s)
	}
}

func (p *AnzerParser) JsonContent() (localctx IJsonContentContext) {
	localctx = NewJsonContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AnzerParserRULE_jsonContent)

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
		p.SetState(44)
		p.Match(AnzerParserJSON_CONTENT)
	}

	return localctx
}
