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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 61, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 6, 2, 14,
	10, 2, 13, 2, 14, 2, 15, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 22, 10, 3, 12, 3,
	14, 3, 25, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 31, 10, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 41, 10, 5, 12, 5, 14, 5, 44,
	11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 50, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 59, 10, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2,
	2, 66, 2, 13, 3, 2, 2, 2, 4, 30, 3, 2, 2, 2, 6, 32, 3, 2, 2, 2, 8, 49,
	3, 2, 2, 2, 10, 58, 3, 2, 2, 2, 12, 14, 5, 10, 6, 2, 13, 12, 3, 2, 2, 2,
	14, 15, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 3, 3, 2,
	2, 2, 17, 18, 7, 3, 2, 2, 18, 23, 5, 6, 4, 2, 19, 20, 7, 4, 2, 2, 20, 22,
	5, 6, 4, 2, 21, 19, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2,
	23, 24, 3, 2, 2, 2, 24, 26, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 26, 27, 7,
	5, 2, 2, 27, 31, 3, 2, 2, 2, 28, 29, 7, 3, 2, 2, 29, 31, 7, 5, 2, 2, 30,
	17, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32, 33, 7, 12, 2,
	2, 33, 34, 7, 6, 2, 2, 34, 35, 5, 10, 6, 2, 35, 7, 3, 2, 2, 2, 36, 37,
	7, 7, 2, 2, 37, 42, 5, 10, 6, 2, 38, 39, 7, 4, 2, 2, 39, 41, 5, 10, 6,
	2, 40, 38, 3, 2, 2, 2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43,
	3, 2, 2, 2, 43, 45, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 46, 7, 8, 2, 2,
	46, 50, 3, 2, 2, 2, 47, 48, 7, 7, 2, 2, 48, 50, 7, 8, 2, 2, 49, 36, 3,
	2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 9, 3, 2, 2, 2, 51, 59, 7, 12, 2, 2, 52,
	59, 7, 13, 2, 2, 53, 59, 5, 4, 3, 2, 54, 59, 5, 8, 5, 2, 55, 59, 7, 9,
	2, 2, 56, 59, 7, 10, 2, 2, 57, 59, 7, 11, 2, 2, 58, 51, 3, 2, 2, 2, 58,
	52, 3, 2, 2, 2, 58, 53, 3, 2, 2, 2, 58, 54, 3, 2, 2, 2, 58, 55, 3, 2, 2,
	2, 58, 56, 3, 2, 2, 2, 58, 57, 3, 2, 2, 2, 59, 11, 3, 2, 2, 2, 8, 15, 23,
	30, 42, 49, 58,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "','", "'}'", "':'", "'['", "']'", "'true'", "'false'", "'null'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "WS",
}

var ruleNames = []string{
	"system", "obj", "pair", "array", "value",
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
	AnzerParserEOF    = antlr.TokenEOF
	AnzerParserT__0   = 1
	AnzerParserT__1   = 2
	AnzerParserT__2   = 3
	AnzerParserT__3   = 4
	AnzerParserT__4   = 5
	AnzerParserT__5   = 6
	AnzerParserT__6   = 7
	AnzerParserT__7   = 8
	AnzerParserT__8   = 9
	AnzerParserSTRING = 10
	AnzerParserNUMBER = 11
	AnzerParserWS     = 12
)

// AnzerParser rules.
const (
	AnzerParserRULE_system = 0
	AnzerParserRULE_obj    = 1
	AnzerParserRULE_pair   = 2
	AnzerParserRULE_array  = 3
	AnzerParserRULE_value  = 4
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

func (s *SystemContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *SystemContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
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
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AnzerParserT__0)|(1<<AnzerParserT__4)|(1<<AnzerParserT__6)|(1<<AnzerParserT__7)|(1<<AnzerParserT__8)|(1<<AnzerParserSTRING)|(1<<AnzerParserNUMBER))) != 0) {
		{
			p.SetState(10)
			p.Value()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 2, AnzerParserRULE_obj)
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

	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(15)
			p.Match(AnzerParserT__0)
		}
		{
			p.SetState(16)
			p.Pair()
		}
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__1 {
			{
				p.SetState(17)
				p.Match(AnzerParserT__1)
			}
			{
				p.SetState(18)
				p.Pair()
			}

			p.SetState(23)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(24)
			p.Match(AnzerParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
			p.Match(AnzerParserT__0)
		}
		{
			p.SetState(27)
			p.Match(AnzerParserT__2)
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

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
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
	p.EnterRule(localctx, 4, AnzerParserRULE_pair)

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
		p.SetState(30)
		p.Match(AnzerParserSTRING)
	}
	{
		p.SetState(31)
		p.Match(AnzerParserT__3)
	}
	{
		p.SetState(32)
		p.Value()
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

func (s *ArrayContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
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
	p.EnterRule(localctx, 6, AnzerParserRULE_array)
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

	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(AnzerParserT__4)
		}
		{
			p.SetState(35)
			p.Value()
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AnzerParserT__1 {
			{
				p.SetState(36)
				p.Match(AnzerParserT__1)
			}
			{
				p.SetState(37)
				p.Value()
			}

			p.SetState(42)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(43)
			p.Match(AnzerParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(AnzerParserT__4)
		}
		{
			p.SetState(46)
			p.Match(AnzerParserT__5)
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(AnzerParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AnzerParserNUMBER, 0)
}

func (s *ValueContext) Obj() IObjContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ValueContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *AnzerParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AnzerParserRULE_value)

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

	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Match(AnzerParserSTRING)
		}

	case AnzerParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Match(AnzerParserNUMBER)
		}

	case AnzerParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(51)
			p.Obj()
		}

	case AnzerParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(52)
			p.Array()
		}

	case AnzerParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(53)
			p.Match(AnzerParserT__6)
		}

	case AnzerParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(54)
			p.Match(AnzerParserT__7)
		}

	case AnzerParserT__8:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(55)
			p.Match(AnzerParserT__8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
