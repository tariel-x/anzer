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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 114,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 6, 2, 32, 10, 2, 13, 2, 14, 2,
	33, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 42, 10, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6, 53, 10, 6, 12, 6, 14, 6, 56,
	11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 63, 10, 7, 12, 7, 14, 7, 66,
	11, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 5, 10, 73, 10, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 81, 10, 10, 13, 10, 14, 10, 82, 3, 10,
	3, 10, 6, 10, 87, 10, 10, 13, 10, 14, 10, 88, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 6, 14, 101, 10, 14, 13, 14, 14,
	14, 102, 3, 15, 3, 15, 3, 15, 6, 15, 108, 10, 15, 13, 15, 14, 15, 109,
	3, 15, 3, 15, 3, 15, 2, 2, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 2, 3, 3, 2, 15, 16, 2, 111, 2, 31, 3, 2, 2, 2, 4, 41, 3, 2,
	2, 2, 6, 43, 3, 2, 2, 2, 8, 48, 3, 2, 2, 2, 10, 50, 3, 2, 2, 2, 12, 59,
	3, 2, 2, 2, 14, 67, 3, 2, 2, 2, 16, 69, 3, 2, 2, 2, 18, 72, 3, 2, 2, 2,
	20, 90, 3, 2, 2, 2, 22, 92, 3, 2, 2, 2, 24, 94, 3, 2, 2, 2, 26, 96, 3,
	2, 2, 2, 28, 104, 3, 2, 2, 2, 30, 32, 5, 4, 3, 2, 31, 30, 3, 2, 2, 2, 32,
	33, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 35, 3, 2, 2,
	2, 35, 36, 7, 2, 2, 3, 36, 3, 3, 2, 2, 2, 37, 42, 5, 6, 4, 2, 38, 42, 5,
	18, 10, 2, 39, 42, 5, 26, 14, 2, 40, 42, 5, 28, 15, 2, 41, 37, 3, 2, 2,
	2, 41, 38, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 40, 3, 2, 2, 2, 42, 5, 3,
	2, 2, 2, 43, 44, 7, 3, 2, 2, 44, 45, 5, 8, 5, 2, 45, 46, 7, 4, 2, 2, 46,
	47, 5, 10, 6, 2, 47, 7, 3, 2, 2, 2, 48, 49, 7, 16, 2, 2, 49, 9, 3, 2, 2,
	2, 50, 54, 7, 5, 2, 2, 51, 53, 5, 12, 7, 2, 52, 51, 3, 2, 2, 2, 53, 56,
	3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 57, 3, 2, 2, 2,
	56, 54, 3, 2, 2, 2, 57, 58, 7, 6, 2, 2, 58, 11, 3, 2, 2, 2, 59, 60, 5,
	14, 8, 2, 60, 64, 7, 7, 2, 2, 61, 63, 5, 16, 9, 2, 62, 61, 3, 2, 2, 2,
	63, 66, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 13, 3,
	2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 68, 9, 2, 2, 2, 68, 15, 3, 2, 2, 2, 69,
	70, 9, 2, 2, 2, 70, 17, 3, 2, 2, 2, 71, 73, 5, 20, 11, 2, 72, 71, 3, 2,
	2, 2, 72, 73, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75, 5, 24, 13, 2, 75,
	76, 7, 8, 2, 2, 76, 77, 5, 22, 12, 2, 77, 78, 7, 9, 2, 2, 78, 80, 7, 7,
	2, 2, 79, 81, 5, 8, 5, 2, 80, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 80,
	3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 86, 7, 10, 2, 2,
	85, 87, 5, 8, 5, 2, 86, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 86, 3,
	2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 19, 3, 2, 2, 2, 90, 91, 7, 15, 2, 2, 91,
	21, 3, 2, 2, 2, 92, 93, 7, 15, 2, 2, 93, 23, 3, 2, 2, 2, 94, 95, 7, 17,
	2, 2, 95, 25, 3, 2, 2, 2, 96, 97, 5, 20, 11, 2, 97, 100, 7, 4, 2, 2, 98,
	101, 5, 20, 11, 2, 99, 101, 7, 11, 2, 2, 100, 98, 3, 2, 2, 2, 100, 99,
	3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2,
	2, 2, 103, 27, 3, 2, 2, 2, 104, 105, 7, 12, 2, 2, 105, 107, 7, 13, 2, 2,
	106, 108, 5, 20, 11, 2, 107, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109,
	107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 112,
	7, 14, 2, 2, 112, 29, 3, 2, 2, 2, 12, 33, 41, 54, 64, 72, 82, 88, 100,
	102, 109,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'type'", "'='", "'{'", "'}'", "'::'", "'['", "']'", "'->'", "'.'",
	"'invoke'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "LowIdent", "UpperIdent",
	"URL", "WS",
}

var ruleNames = []string{
	"forms", "form", "typeDeclaration", "typeName", "typeDefinition", "typeField",
	"fieldName", "typeId", "funcDeclaration", "funcName", "runtime", "url",
	"localFuncDeclaration", "invokeCmd",
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
	AnzerParserEOF        = antlr.TokenEOF
	AnzerParserT__0       = 1
	AnzerParserT__1       = 2
	AnzerParserT__2       = 3
	AnzerParserT__3       = 4
	AnzerParserT__4       = 5
	AnzerParserT__5       = 6
	AnzerParserT__6       = 7
	AnzerParserT__7       = 8
	AnzerParserT__8       = 9
	AnzerParserT__9       = 10
	AnzerParserT__10      = 11
	AnzerParserT__11      = 12
	AnzerParserLowIdent   = 13
	AnzerParserUpperIdent = 14
	AnzerParserURL        = 15
	AnzerParserWS         = 16
)

// AnzerParser rules.
const (
	AnzerParserRULE_forms                = 0
	AnzerParserRULE_form                 = 1
	AnzerParserRULE_typeDeclaration      = 2
	AnzerParserRULE_typeName             = 3
	AnzerParserRULE_typeDefinition       = 4
	AnzerParserRULE_typeField            = 5
	AnzerParserRULE_fieldName            = 6
	AnzerParserRULE_typeId               = 7
	AnzerParserRULE_funcDeclaration      = 8
	AnzerParserRULE_funcName             = 9
	AnzerParserRULE_runtime              = 10
	AnzerParserRULE_url                  = 11
	AnzerParserRULE_localFuncDeclaration = 12
	AnzerParserRULE_invokeCmd            = 13
)

// IFormsContext is an interface to support dynamic dispatch.
type IFormsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormsContext differentiates from other interfaces.
	IsFormsContext()
}

type FormsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormsContext() *FormsContext {
	var p = new(FormsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_forms
	return p
}

func (*FormsContext) IsFormsContext() {}

func NewFormsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormsContext {
	var p = new(FormsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_forms

	return p
}

func (s *FormsContext) GetParser() antlr.Parser { return s.parser }

func (s *FormsContext) EOF() antlr.TerminalNode {
	return s.GetToken(AnzerParserEOF, 0)
}

func (s *FormsContext) AllForm() []IFormContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormContext)(nil)).Elem())
	var tst = make([]IFormContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormContext)
		}
	}

	return tst
}

func (s *FormsContext) Form(i int) IFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormContext)
}

func (s *FormsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterForms(s)
	}
}

func (s *FormsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitForms(s)
	}
}

func (p *AnzerParser) Forms() (localctx IFormsContext) {
	localctx = NewFormsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AnzerParserRULE_forms)
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

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AnzerParserT__0)|(1<<AnzerParserT__9)|(1<<AnzerParserLowIdent)|(1<<AnzerParserURL))) != 0) {
		{
			p.SetState(28)
			p.Form()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
		p.Match(AnzerParserEOF)
	}

	return localctx
}

// IFormContext is an interface to support dynamic dispatch.
type IFormContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormContext differentiates from other interfaces.
	IsFormContext()
}

type FormContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormContext() *FormContext {
	var p = new(FormContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_form
	return p
}

func (*FormContext) IsFormContext() {}

func NewFormContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormContext {
	var p = new(FormContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_form

	return p
}

func (s *FormContext) GetParser() antlr.Parser { return s.parser }

func (s *FormContext) TypeDeclaration() ITypeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *FormContext) FuncDeclaration() IFuncDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncDeclarationContext)
}

func (s *FormContext) LocalFuncDeclaration() ILocalFuncDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalFuncDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalFuncDeclarationContext)
}

func (s *FormContext) InvokeCmd() IInvokeCmdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvokeCmdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvokeCmdContext)
}

func (s *FormContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterForm(s)
	}
}

func (s *FormContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitForm(s)
	}
}

func (p *AnzerParser) Form() (localctx IFormContext) {
	localctx = NewFormContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AnzerParserRULE_form)

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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.TypeDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
			p.FuncDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(37)
			p.LocalFuncDeclaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(38)
			p.InvokeCmd()
		}

	}

	return localctx
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TypeDeclarationContext) TypeDefinition() ITypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefinitionContext)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (p *AnzerParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AnzerParserRULE_typeDeclaration)

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
		p.SetState(41)
		p.Match(AnzerParserT__0)
	}
	{
		p.SetState(42)
		p.TypeName()
	}
	{
		p.SetState(43)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(44)
		p.TypeDefinition()
	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) UpperIdent() antlr.TerminalNode {
	return s.GetToken(AnzerParserUpperIdent, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (p *AnzerParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AnzerParserRULE_typeName)

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
		p.SetState(46)
		p.Match(AnzerParserUpperIdent)
	}

	return localctx
}

// ITypeDefinitionContext is an interface to support dynamic dispatch.
type ITypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefinitionContext differentiates from other interfaces.
	IsTypeDefinitionContext()
}

type TypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefinitionContext() *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeDefinition
	return p
}

func (*TypeDefinitionContext) IsTypeDefinitionContext() {}

func NewTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeDefinition

	return p
}

func (s *TypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefinitionContext) AllTypeField() []ITypeFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeFieldContext)(nil)).Elem())
	var tst = make([]ITypeFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeFieldContext)
		}
	}

	return tst
}

func (s *TypeDefinitionContext) TypeField(i int) ITypeFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeFieldContext)
}

func (s *TypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeDefinition(s)
	}
}

func (s *TypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeDefinition(s)
	}
}

func (p *AnzerParser) TypeDefinition() (localctx ITypeDefinitionContext) {
	localctx = NewTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AnzerParserRULE_typeDefinition)
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
		p.SetState(48)
		p.Match(AnzerParserT__2)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserLowIdent || _la == AnzerParserUpperIdent {
		{
			p.SetState(49)
			p.TypeField()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
		p.Match(AnzerParserT__3)
	}

	return localctx
}

// ITypeFieldContext is an interface to support dynamic dispatch.
type ITypeFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeFieldContext differentiates from other interfaces.
	IsTypeFieldContext()
}

type TypeFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFieldContext() *TypeFieldContext {
	var p = new(TypeFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeField
	return p
}

func (*TypeFieldContext) IsTypeFieldContext() {}

func NewTypeFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFieldContext {
	var p = new(TypeFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeField

	return p
}

func (s *TypeFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *TypeFieldContext) AllTypeId() []ITypeIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeIdContext)(nil)).Elem())
	var tst = make([]ITypeIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeIdContext)
		}
	}

	return tst
}

func (s *TypeFieldContext) TypeId(i int) ITypeIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeIdContext)
}

func (s *TypeFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeField(s)
	}
}

func (s *TypeFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeField(s)
	}
}

func (p *AnzerParser) TypeField() (localctx ITypeFieldContext) {
	localctx = NewTypeFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AnzerParserRULE_typeField)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.FieldName()
	}
	{
		p.SetState(58)
		p.Match(AnzerParserT__4)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(59)
				p.TypeId()
			}

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) LowIdent() antlr.TerminalNode {
	return s.GetToken(AnzerParserLowIdent, 0)
}

func (s *FieldNameContext) UpperIdent() antlr.TerminalNode {
	return s.GetToken(AnzerParserUpperIdent, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *AnzerParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AnzerParserRULE_fieldName)
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
		p.SetState(65)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AnzerParserLowIdent || _la == AnzerParserUpperIdent) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeIdContext is an interface to support dynamic dispatch.
type ITypeIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeIdContext differentiates from other interfaces.
	IsTypeIdContext()
}

type TypeIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeIdContext() *TypeIdContext {
	var p = new(TypeIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeId
	return p
}

func (*TypeIdContext) IsTypeIdContext() {}

func NewTypeIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeIdContext {
	var p = new(TypeIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeId

	return p
}

func (s *TypeIdContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeIdContext) LowIdent() antlr.TerminalNode {
	return s.GetToken(AnzerParserLowIdent, 0)
}

func (s *TypeIdContext) UpperIdent() antlr.TerminalNode {
	return s.GetToken(AnzerParserUpperIdent, 0)
}

func (s *TypeIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeId(s)
	}
}

func (s *TypeIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeId(s)
	}
}

func (p *AnzerParser) TypeId() (localctx ITypeIdContext) {
	localctx = NewTypeIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AnzerParserRULE_typeId)
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
		p.SetState(67)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AnzerParserLowIdent || _la == AnzerParserUpperIdent) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFuncDeclarationContext is an interface to support dynamic dispatch.
type IFuncDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDeclarationContext differentiates from other interfaces.
	IsFuncDeclarationContext()
}

type FuncDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclarationContext() *FuncDeclarationContext {
	var p = new(FuncDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcDeclaration
	return p
}

func (*FuncDeclarationContext) IsFuncDeclarationContext() {}

func NewFuncDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclarationContext {
	var p = new(FuncDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcDeclaration

	return p
}

func (s *FuncDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclarationContext) Url() IUrlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUrlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUrlContext)
}

func (s *FuncDeclarationContext) Runtime() IRuntimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuntimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuntimeContext)
}

func (s *FuncDeclarationContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *FuncDeclarationContext) AllTypeName() []ITypeNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeNameContext)(nil)).Elem())
	var tst = make([]ITypeNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeNameContext)
		}
	}

	return tst
}

func (s *FuncDeclarationContext) TypeName(i int) ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *FuncDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncDeclaration(s)
	}
}

func (s *FuncDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncDeclaration(s)
	}
}

func (p *AnzerParser) FuncDeclaration() (localctx IFuncDeclarationContext) {
	localctx = NewFuncDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AnzerParserRULE_funcDeclaration)
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
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AnzerParserLowIdent {
		{
			p.SetState(69)
			p.FuncName()
		}

	}
	{
		p.SetState(72)
		p.Url()
	}
	{
		p.SetState(73)
		p.Match(AnzerParserT__5)
	}
	{
		p.SetState(74)
		p.Runtime()
	}
	{
		p.SetState(75)
		p.Match(AnzerParserT__6)
	}
	{
		p.SetState(76)
		p.Match(AnzerParserT__4)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AnzerParserUpperIdent {
		{
			p.SetState(77)
			p.TypeName()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
		p.Match(AnzerParserT__7)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AnzerParserUpperIdent {
		{
			p.SetState(83)
			p.TypeName()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *FuncNameContext) LowIdent() antlr.TerminalNode {
	return s.GetToken(AnzerParserLowIdent, 0)
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
	p.EnterRule(localctx, 18, AnzerParserRULE_funcName)

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
		p.SetState(88)
		p.Match(AnzerParserLowIdent)
	}

	return localctx
}

// IRuntimeContext is an interface to support dynamic dispatch.
type IRuntimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuntimeContext differentiates from other interfaces.
	IsRuntimeContext()
}

type RuntimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuntimeContext() *RuntimeContext {
	var p = new(RuntimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_runtime
	return p
}

func (*RuntimeContext) IsRuntimeContext() {}

func NewRuntimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuntimeContext {
	var p = new(RuntimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_runtime

	return p
}

func (s *RuntimeContext) GetParser() antlr.Parser { return s.parser }

func (s *RuntimeContext) LowIdent() antlr.TerminalNode {
	return s.GetToken(AnzerParserLowIdent, 0)
}

func (s *RuntimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuntimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuntimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterRuntime(s)
	}
}

func (s *RuntimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitRuntime(s)
	}
}

func (p *AnzerParser) Runtime() (localctx IRuntimeContext) {
	localctx = NewRuntimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AnzerParserRULE_runtime)

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
		p.SetState(90)
		p.Match(AnzerParserLowIdent)
	}

	return localctx
}

// IUrlContext is an interface to support dynamic dispatch.
type IUrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUrlContext differentiates from other interfaces.
	IsUrlContext()
}

type UrlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrlContext() *UrlContext {
	var p = new(UrlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_url
	return p
}

func (*UrlContext) IsUrlContext() {}

func NewUrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UrlContext {
	var p = new(UrlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_url

	return p
}

func (s *UrlContext) GetParser() antlr.Parser { return s.parser }

func (s *UrlContext) URL() antlr.TerminalNode {
	return s.GetToken(AnzerParserURL, 0)
}

func (s *UrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterUrl(s)
	}
}

func (s *UrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitUrl(s)
	}
}

func (p *AnzerParser) Url() (localctx IUrlContext) {
	localctx = NewUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AnzerParserRULE_url)

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
		p.SetState(92)
		p.Match(AnzerParserURL)
	}

	return localctx
}

// ILocalFuncDeclarationContext is an interface to support dynamic dispatch.
type ILocalFuncDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalFuncDeclarationContext differentiates from other interfaces.
	IsLocalFuncDeclarationContext()
}

type LocalFuncDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalFuncDeclarationContext() *LocalFuncDeclarationContext {
	var p = new(LocalFuncDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_localFuncDeclaration
	return p
}

func (*LocalFuncDeclarationContext) IsLocalFuncDeclarationContext() {}

func NewLocalFuncDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalFuncDeclarationContext {
	var p = new(LocalFuncDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_localFuncDeclaration

	return p
}

func (s *LocalFuncDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalFuncDeclarationContext) AllFuncName() []IFuncNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncNameContext)(nil)).Elem())
	var tst = make([]IFuncNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncNameContext)
		}
	}

	return tst
}

func (s *LocalFuncDeclarationContext) FuncName(i int) IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *LocalFuncDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalFuncDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalFuncDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterLocalFuncDeclaration(s)
	}
}

func (s *LocalFuncDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitLocalFuncDeclaration(s)
	}
}

func (p *AnzerParser) LocalFuncDeclaration() (localctx ILocalFuncDeclarationContext) {
	localctx = NewLocalFuncDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AnzerParserRULE_localFuncDeclaration)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.FuncName()
	}
	{
		p.SetState(95)
		p.Match(AnzerParserT__1)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(98)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case AnzerParserLowIdent:
				{
					p.SetState(96)
					p.FuncName()
				}

			case AnzerParserT__8:
				{
					p.SetState(97)
					p.Match(AnzerParserT__8)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IInvokeCmdContext is an interface to support dynamic dispatch.
type IInvokeCmdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvokeCmdContext differentiates from other interfaces.
	IsInvokeCmdContext()
}

type InvokeCmdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvokeCmdContext() *InvokeCmdContext {
	var p = new(InvokeCmdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_invokeCmd
	return p
}

func (*InvokeCmdContext) IsInvokeCmdContext() {}

func NewInvokeCmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvokeCmdContext {
	var p = new(InvokeCmdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_invokeCmd

	return p
}

func (s *InvokeCmdContext) GetParser() antlr.Parser { return s.parser }

func (s *InvokeCmdContext) AllFuncName() []IFuncNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncNameContext)(nil)).Elem())
	var tst = make([]IFuncNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncNameContext)
		}
	}

	return tst
}

func (s *InvokeCmdContext) FuncName(i int) IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *InvokeCmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvokeCmdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvokeCmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterInvokeCmd(s)
	}
}

func (s *InvokeCmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitInvokeCmd(s)
	}
}

func (p *AnzerParser) InvokeCmd() (localctx IInvokeCmdContext) {
	localctx = NewInvokeCmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AnzerParserRULE_invokeCmd)
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
		p.SetState(102)
		p.Match(AnzerParserT__9)
	}
	{
		p.SetState(103)
		p.Match(AnzerParserT__10)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AnzerParserLowIdent {
		{
			p.SetState(104)
			p.FuncName()
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(109)
		p.Match(AnzerParserT__11)
	}

	return localctx
}
