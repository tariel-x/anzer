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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 213,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 3, 2, 6, 2, 70, 10, 2, 13, 2, 14, 2, 71, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 80, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 5, 6, 91, 10, 6, 3, 7, 7, 7, 94, 10, 7, 12, 7, 14, 7, 97, 11,
	7, 3, 7, 3, 7, 7, 7, 101, 10, 7, 12, 7, 14, 7, 104, 11, 7, 3, 7, 3, 7,
	3, 8, 6, 8, 109, 10, 8, 13, 8, 14, 8, 110, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 122, 10, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 5, 12, 130, 10, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 19, 3, 19, 5, 19, 150, 10, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 5, 25, 163, 10, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 29, 6, 29, 181, 10, 29, 13, 29, 14, 29, 182, 3,
	30, 6, 30, 186, 10, 30, 13, 30, 14, 30, 187, 3, 31, 3, 31, 3, 31, 3, 31,
	6, 31, 194, 10, 31, 13, 31, 14, 31, 195, 3, 32, 3, 32, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 6, 33, 205, 10, 33, 13, 33, 14, 33, 206, 3, 33, 3, 33,
	3, 34, 3, 34, 3, 34, 2, 2, 35, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 2, 5, 3, 2, 28, 29, 3, 2, 12, 13, 3, 2, 14, 15, 2, 203,
	2, 69, 3, 2, 2, 2, 4, 79, 3, 2, 2, 2, 6, 81, 3, 2, 2, 2, 8, 86, 3, 2, 2,
	2, 10, 90, 3, 2, 2, 2, 12, 95, 3, 2, 2, 2, 14, 108, 3, 2, 2, 2, 16, 112,
	3, 2, 2, 2, 18, 116, 3, 2, 2, 2, 20, 121, 3, 2, 2, 2, 22, 129, 3, 2, 2,
	2, 24, 131, 3, 2, 2, 2, 26, 134, 3, 2, 2, 2, 28, 137, 3, 2, 2, 2, 30, 139,
	3, 2, 2, 2, 32, 141, 3, 2, 2, 2, 34, 143, 3, 2, 2, 2, 36, 149, 3, 2, 2,
	2, 38, 151, 3, 2, 2, 2, 40, 153, 3, 2, 2, 2, 42, 155, 3, 2, 2, 2, 44, 157,
	3, 2, 2, 2, 46, 159, 3, 2, 2, 2, 48, 162, 3, 2, 2, 2, 50, 173, 3, 2, 2,
	2, 52, 175, 3, 2, 2, 2, 54, 177, 3, 2, 2, 2, 56, 180, 3, 2, 2, 2, 58, 185,
	3, 2, 2, 2, 60, 189, 3, 2, 2, 2, 62, 197, 3, 2, 2, 2, 64, 199, 3, 2, 2,
	2, 66, 210, 3, 2, 2, 2, 68, 70, 5, 4, 3, 2, 69, 68, 3, 2, 2, 2, 70, 71,
	3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2,
	73, 74, 7, 2, 2, 3, 74, 3, 3, 2, 2, 2, 75, 80, 5, 6, 4, 2, 76, 80, 5, 48,
	25, 2, 77, 80, 5, 60, 31, 2, 78, 80, 5, 64, 33, 2, 79, 75, 3, 2, 2, 2,
	79, 76, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 78, 3, 2, 2, 2, 80, 5, 3, 2,
	2, 2, 81, 82, 7, 3, 2, 2, 82, 83, 5, 8, 5, 2, 83, 84, 7, 4, 2, 2, 84, 85,
	5, 10, 6, 2, 85, 7, 3, 2, 2, 2, 86, 87, 7, 29, 2, 2, 87, 9, 3, 2, 2, 2,
	88, 91, 5, 12, 7, 2, 89, 91, 5, 14, 8, 2, 90, 88, 3, 2, 2, 2, 90, 89, 3,
	2, 2, 2, 91, 11, 3, 2, 2, 2, 92, 94, 5, 22, 12, 2, 93, 92, 3, 2, 2, 2,
	94, 97, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 98, 3,
	2, 2, 2, 97, 95, 3, 2, 2, 2, 98, 102, 7, 5, 2, 2, 99, 101, 5, 16, 9, 2,
	100, 99, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103,
	3, 2, 2, 2, 103, 105, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 106, 7, 6,
	2, 2, 106, 13, 3, 2, 2, 2, 107, 109, 5, 20, 11, 2, 108, 107, 3, 2, 2, 2,
	109, 110, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111,
	15, 3, 2, 2, 2, 112, 113, 5, 18, 10, 2, 113, 114, 7, 7, 2, 2, 114, 115,
	5, 10, 6, 2, 115, 17, 3, 2, 2, 2, 116, 117, 9, 2, 2, 2, 117, 19, 3, 2,
	2, 2, 118, 122, 5, 22, 12, 2, 119, 122, 5, 36, 19, 2, 120, 122, 5, 46,
	24, 2, 121, 118, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2,
	122, 21, 3, 2, 2, 2, 123, 130, 5, 24, 13, 2, 124, 130, 5, 26, 14, 2, 125,
	130, 5, 28, 15, 2, 126, 130, 5, 30, 16, 2, 127, 130, 5, 32, 17, 2, 128,
	130, 5, 34, 18, 2, 129, 123, 3, 2, 2, 2, 129, 124, 3, 2, 2, 2, 129, 125,
	3, 2, 2, 2, 129, 126, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 128, 3, 2,
	2, 2, 130, 23, 3, 2, 2, 2, 131, 132, 7, 8, 2, 2, 132, 133, 7, 30, 2, 2,
	133, 25, 3, 2, 2, 2, 134, 135, 7, 9, 2, 2, 135, 136, 7, 30, 2, 2, 136,
	27, 3, 2, 2, 2, 137, 138, 7, 10, 2, 2, 138, 29, 3, 2, 2, 2, 139, 140, 7,
	11, 2, 2, 140, 31, 3, 2, 2, 2, 141, 142, 9, 3, 2, 2, 142, 33, 3, 2, 2,
	2, 143, 144, 9, 4, 2, 2, 144, 35, 3, 2, 2, 2, 145, 150, 5, 38, 20, 2, 146,
	150, 5, 40, 21, 2, 147, 150, 5, 42, 22, 2, 148, 150, 5, 44, 23, 2, 149,
	145, 3, 2, 2, 2, 149, 146, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 148,
	3, 2, 2, 2, 150, 37, 3, 2, 2, 2, 151, 152, 7, 16, 2, 2, 152, 39, 3, 2,
	2, 2, 153, 154, 7, 17, 2, 2, 154, 41, 3, 2, 2, 2, 155, 156, 7, 18, 2, 2,
	156, 43, 3, 2, 2, 2, 157, 158, 7, 19, 2, 2, 158, 45, 3, 2, 2, 2, 159, 160,
	7, 29, 2, 2, 160, 47, 3, 2, 2, 2, 161, 163, 5, 50, 26, 2, 162, 161, 3,
	2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 5, 54, 28,
	2, 165, 166, 7, 20, 2, 2, 166, 167, 5, 52, 27, 2, 167, 168, 7, 21, 2, 2,
	168, 169, 7, 7, 2, 2, 169, 170, 5, 56, 29, 2, 170, 171, 7, 22, 2, 2, 171,
	172, 5, 58, 30, 2, 172, 49, 3, 2, 2, 2, 173, 174, 7, 28, 2, 2, 174, 51,
	3, 2, 2, 2, 175, 176, 7, 28, 2, 2, 176, 53, 3, 2, 2, 2, 177, 178, 7, 31,
	2, 2, 178, 55, 3, 2, 2, 2, 179, 181, 5, 20, 11, 2, 180, 179, 3, 2, 2, 2,
	181, 182, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183,
	57, 3, 2, 2, 2, 184, 186, 5, 20, 11, 2, 185, 184, 3, 2, 2, 2, 186, 187,
	3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 59, 3, 2,
	2, 2, 189, 190, 5, 50, 26, 2, 190, 193, 7, 4, 2, 2, 191, 194, 5, 62, 32,
	2, 192, 194, 7, 23, 2, 2, 193, 191, 3, 2, 2, 2, 193, 192, 3, 2, 2, 2, 194,
	195, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 61, 3,
	2, 2, 2, 197, 198, 7, 28, 2, 2, 198, 63, 3, 2, 2, 2, 199, 200, 7, 24, 2,
	2, 200, 204, 7, 25, 2, 2, 201, 202, 5, 66, 34, 2, 202, 203, 7, 26, 2, 2,
	203, 205, 3, 2, 2, 2, 204, 201, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206,
	204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 209,
	7, 27, 2, 2, 209, 65, 3, 2, 2, 2, 210, 211, 7, 28, 2, 2, 211, 67, 3, 2,
	2, 2, 17, 71, 79, 90, 95, 102, 110, 121, 129, 149, 162, 182, 187, 193,
	195, 206,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'type'", "'='", "'{'", "'}'", "'::'", "'MinLength'", "'MaxLength'",
	"'Right'", "'Left'", "'List'", "'[]'", "'Optional'", "'*'", "'String'",
	"'Integer'", "'Float'", "'Bool'", "'['", "']'", "'->'", "'.'", "'invoke'",
	"'('", "','", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "LowIdent", "UpperIdent", "ConstructorArg",
	"URL", "WS",
}

var ruleNames = []string{
	"forms", "form", "typeDeclaration", "typeName", "typeDefinition", "typeComplexDefinition",
	"typeSimpleDefinition", "typeField", "fieldName", "typeId", "typeConstructor",
	"typeMinLength", "typeMaxLength", "typeRight", "typeLeft", "typeList",
	"typeOptional", "typeScalar", "typeString", "typeInteger", "typeFloat",
	"typeBool", "typeOther", "funcDeclaration", "funcName", "runtime", "url",
	"funcArgument", "funcResult", "localFuncDeclaration", "funcRef", "invokeCmd",
	"invokeFuncName",
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
	AnzerParserEOF            = antlr.TokenEOF
	AnzerParserT__0           = 1
	AnzerParserT__1           = 2
	AnzerParserT__2           = 3
	AnzerParserT__3           = 4
	AnzerParserT__4           = 5
	AnzerParserT__5           = 6
	AnzerParserT__6           = 7
	AnzerParserT__7           = 8
	AnzerParserT__8           = 9
	AnzerParserT__9           = 10
	AnzerParserT__10          = 11
	AnzerParserT__11          = 12
	AnzerParserT__12          = 13
	AnzerParserT__13          = 14
	AnzerParserT__14          = 15
	AnzerParserT__15          = 16
	AnzerParserT__16          = 17
	AnzerParserT__17          = 18
	AnzerParserT__18          = 19
	AnzerParserT__19          = 20
	AnzerParserT__20          = 21
	AnzerParserT__21          = 22
	AnzerParserT__22          = 23
	AnzerParserT__23          = 24
	AnzerParserT__24          = 25
	AnzerParserLowIdent       = 26
	AnzerParserUpperIdent     = 27
	AnzerParserConstructorArg = 28
	AnzerParserURL            = 29
	AnzerParserWS             = 30
)

// AnzerParser rules.
const (
	AnzerParserRULE_forms                 = 0
	AnzerParserRULE_form                  = 1
	AnzerParserRULE_typeDeclaration       = 2
	AnzerParserRULE_typeName              = 3
	AnzerParserRULE_typeDefinition        = 4
	AnzerParserRULE_typeComplexDefinition = 5
	AnzerParserRULE_typeSimpleDefinition  = 6
	AnzerParserRULE_typeField             = 7
	AnzerParserRULE_fieldName             = 8
	AnzerParserRULE_typeId                = 9
	AnzerParserRULE_typeConstructor       = 10
	AnzerParserRULE_typeMinLength         = 11
	AnzerParserRULE_typeMaxLength         = 12
	AnzerParserRULE_typeRight             = 13
	AnzerParserRULE_typeLeft              = 14
	AnzerParserRULE_typeList              = 15
	AnzerParserRULE_typeOptional          = 16
	AnzerParserRULE_typeScalar            = 17
	AnzerParserRULE_typeString            = 18
	AnzerParserRULE_typeInteger           = 19
	AnzerParserRULE_typeFloat             = 20
	AnzerParserRULE_typeBool              = 21
	AnzerParserRULE_typeOther             = 22
	AnzerParserRULE_funcDeclaration       = 23
	AnzerParserRULE_funcName              = 24
	AnzerParserRULE_runtime               = 25
	AnzerParserRULE_url                   = 26
	AnzerParserRULE_funcArgument          = 27
	AnzerParserRULE_funcResult            = 28
	AnzerParserRULE_localFuncDeclaration  = 29
	AnzerParserRULE_funcRef               = 30
	AnzerParserRULE_invokeCmd             = 31
	AnzerParserRULE_invokeFuncName        = 32
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AnzerParserT__0)|(1<<AnzerParserT__21)|(1<<AnzerParserLowIdent)|(1<<AnzerParserURL))) != 0) {
		{
			p.SetState(66)
			p.Form()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(71)
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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.TypeDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.FuncDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.LocalFuncDeclaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
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
		p.SetState(79)
		p.Match(AnzerParserT__0)
	}
	{
		p.SetState(80)
		p.TypeName()
	}
	{
		p.SetState(81)
		p.Match(AnzerParserT__1)
	}
	{
		p.SetState(82)
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
		p.SetState(84)
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

func (s *TypeDefinitionContext) TypeComplexDefinition() ITypeComplexDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeComplexDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeComplexDefinitionContext)
}

func (s *TypeDefinitionContext) TypeSimpleDefinition() ITypeSimpleDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSimpleDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSimpleDefinitionContext)
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

	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.TypeComplexDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.TypeSimpleDefinition()
		}

	}

	return localctx
}

// ITypeComplexDefinitionContext is an interface to support dynamic dispatch.
type ITypeComplexDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeComplexDefinitionContext differentiates from other interfaces.
	IsTypeComplexDefinitionContext()
}

type TypeComplexDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeComplexDefinitionContext() *TypeComplexDefinitionContext {
	var p = new(TypeComplexDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeComplexDefinition
	return p
}

func (*TypeComplexDefinitionContext) IsTypeComplexDefinitionContext() {}

func NewTypeComplexDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeComplexDefinitionContext {
	var p = new(TypeComplexDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeComplexDefinition

	return p
}

func (s *TypeComplexDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeComplexDefinitionContext) AllTypeConstructor() []ITypeConstructorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeConstructorContext)(nil)).Elem())
	var tst = make([]ITypeConstructorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeConstructorContext)
		}
	}

	return tst
}

func (s *TypeComplexDefinitionContext) TypeConstructor(i int) ITypeConstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConstructorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeConstructorContext)
}

func (s *TypeComplexDefinitionContext) AllTypeField() []ITypeFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeFieldContext)(nil)).Elem())
	var tst = make([]ITypeFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeFieldContext)
		}
	}

	return tst
}

func (s *TypeComplexDefinitionContext) TypeField(i int) ITypeFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeFieldContext)
}

func (s *TypeComplexDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeComplexDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeComplexDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeComplexDefinition(s)
	}
}

func (s *TypeComplexDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeComplexDefinition(s)
	}
}

func (p *AnzerParser) TypeComplexDefinition() (localctx ITypeComplexDefinitionContext) {
	localctx = NewTypeComplexDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AnzerParserRULE_typeComplexDefinition)
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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AnzerParserT__5)|(1<<AnzerParserT__6)|(1<<AnzerParserT__7)|(1<<AnzerParserT__8)|(1<<AnzerParserT__9)|(1<<AnzerParserT__10)|(1<<AnzerParserT__11)|(1<<AnzerParserT__12))) != 0 {
		{
			p.SetState(90)
			p.TypeConstructor()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
		p.Match(AnzerParserT__2)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnzerParserLowIdent || _la == AnzerParserUpperIdent {
		{
			p.SetState(97)
			p.TypeField()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(AnzerParserT__3)
	}

	return localctx
}

// ITypeSimpleDefinitionContext is an interface to support dynamic dispatch.
type ITypeSimpleDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSimpleDefinitionContext differentiates from other interfaces.
	IsTypeSimpleDefinitionContext()
}

type TypeSimpleDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSimpleDefinitionContext() *TypeSimpleDefinitionContext {
	var p = new(TypeSimpleDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeSimpleDefinition
	return p
}

func (*TypeSimpleDefinitionContext) IsTypeSimpleDefinitionContext() {}

func NewTypeSimpleDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSimpleDefinitionContext {
	var p = new(TypeSimpleDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeSimpleDefinition

	return p
}

func (s *TypeSimpleDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSimpleDefinitionContext) AllTypeId() []ITypeIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeIdContext)(nil)).Elem())
	var tst = make([]ITypeIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeIdContext)
		}
	}

	return tst
}

func (s *TypeSimpleDefinitionContext) TypeId(i int) ITypeIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeIdContext)
}

func (s *TypeSimpleDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSimpleDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSimpleDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeSimpleDefinition(s)
	}
}

func (s *TypeSimpleDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeSimpleDefinition(s)
	}
}

func (p *AnzerParser) TypeSimpleDefinition() (localctx ITypeSimpleDefinitionContext) {
	localctx = NewTypeSimpleDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AnzerParserRULE_typeSimpleDefinition)

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
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(105)
				p.TypeId()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
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

func (s *TypeFieldContext) TypeDefinition() ITypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefinitionContext)
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
	p.EnterRule(localctx, 14, AnzerParserRULE_typeField)

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
		p.FieldName()
	}
	{
		p.SetState(111)
		p.Match(AnzerParserT__4)
	}
	{
		p.SetState(112)
		p.TypeDefinition()
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
	p.EnterRule(localctx, 16, AnzerParserRULE_fieldName)
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
		p.SetState(114)
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

func (s *TypeIdContext) TypeConstructor() ITypeConstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeConstructorContext)
}

func (s *TypeIdContext) TypeScalar() ITypeScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeScalarContext)
}

func (s *TypeIdContext) TypeOther() ITypeOtherContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeOtherContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeOtherContext)
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
	p.EnterRule(localctx, 18, AnzerParserRULE_typeId)

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

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserT__5, AnzerParserT__6, AnzerParserT__7, AnzerParserT__8, AnzerParserT__9, AnzerParserT__10, AnzerParserT__11, AnzerParserT__12:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.TypeConstructor()
		}

	case AnzerParserT__13, AnzerParserT__14, AnzerParserT__15, AnzerParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.TypeScalar()
		}

	case AnzerParserUpperIdent:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.TypeOther()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeConstructorContext is an interface to support dynamic dispatch.
type ITypeConstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeConstructorContext differentiates from other interfaces.
	IsTypeConstructorContext()
}

type TypeConstructorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeConstructorContext() *TypeConstructorContext {
	var p = new(TypeConstructorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeConstructor
	return p
}

func (*TypeConstructorContext) IsTypeConstructorContext() {}

func NewTypeConstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeConstructorContext {
	var p = new(TypeConstructorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeConstructor

	return p
}

func (s *TypeConstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeConstructorContext) TypeMinLength() ITypeMinLengthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeMinLengthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeMinLengthContext)
}

func (s *TypeConstructorContext) TypeMaxLength() ITypeMaxLengthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeMaxLengthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeMaxLengthContext)
}

func (s *TypeConstructorContext) TypeRight() ITypeRightContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRightContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRightContext)
}

func (s *TypeConstructorContext) TypeLeft() ITypeLeftContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLeftContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeLeftContext)
}

func (s *TypeConstructorContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *TypeConstructorContext) TypeOptional() ITypeOptionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeOptionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeOptionalContext)
}

func (s *TypeConstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeConstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeConstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeConstructor(s)
	}
}

func (s *TypeConstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeConstructor(s)
	}
}

func (p *AnzerParser) TypeConstructor() (localctx ITypeConstructorContext) {
	localctx = NewTypeConstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AnzerParserRULE_typeConstructor)

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

	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.TypeMinLength()
		}

	case AnzerParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.TypeMaxLength()
		}

	case AnzerParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.TypeRight()
		}

	case AnzerParserT__8:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(124)
			p.TypeLeft()
		}

	case AnzerParserT__9, AnzerParserT__10:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(125)
			p.TypeList()
		}

	case AnzerParserT__11, AnzerParserT__12:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(126)
			p.TypeOptional()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeMinLengthContext is an interface to support dynamic dispatch.
type ITypeMinLengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeMinLengthContext differentiates from other interfaces.
	IsTypeMinLengthContext()
}

type TypeMinLengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeMinLengthContext() *TypeMinLengthContext {
	var p = new(TypeMinLengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeMinLength
	return p
}

func (*TypeMinLengthContext) IsTypeMinLengthContext() {}

func NewTypeMinLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeMinLengthContext {
	var p = new(TypeMinLengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeMinLength

	return p
}

func (s *TypeMinLengthContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeMinLengthContext) ConstructorArg() antlr.TerminalNode {
	return s.GetToken(AnzerParserConstructorArg, 0)
}

func (s *TypeMinLengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeMinLengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeMinLengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeMinLength(s)
	}
}

func (s *TypeMinLengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeMinLength(s)
	}
}

func (p *AnzerParser) TypeMinLength() (localctx ITypeMinLengthContext) {
	localctx = NewTypeMinLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AnzerParserRULE_typeMinLength)

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
		p.SetState(129)
		p.Match(AnzerParserT__5)
	}
	{
		p.SetState(130)
		p.Match(AnzerParserConstructorArg)
	}

	return localctx
}

// ITypeMaxLengthContext is an interface to support dynamic dispatch.
type ITypeMaxLengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeMaxLengthContext differentiates from other interfaces.
	IsTypeMaxLengthContext()
}

type TypeMaxLengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeMaxLengthContext() *TypeMaxLengthContext {
	var p = new(TypeMaxLengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeMaxLength
	return p
}

func (*TypeMaxLengthContext) IsTypeMaxLengthContext() {}

func NewTypeMaxLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeMaxLengthContext {
	var p = new(TypeMaxLengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeMaxLength

	return p
}

func (s *TypeMaxLengthContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeMaxLengthContext) ConstructorArg() antlr.TerminalNode {
	return s.GetToken(AnzerParserConstructorArg, 0)
}

func (s *TypeMaxLengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeMaxLengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeMaxLengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeMaxLength(s)
	}
}

func (s *TypeMaxLengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeMaxLength(s)
	}
}

func (p *AnzerParser) TypeMaxLength() (localctx ITypeMaxLengthContext) {
	localctx = NewTypeMaxLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AnzerParserRULE_typeMaxLength)

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
		p.SetState(132)
		p.Match(AnzerParserT__6)
	}
	{
		p.SetState(133)
		p.Match(AnzerParserConstructorArg)
	}

	return localctx
}

// ITypeRightContext is an interface to support dynamic dispatch.
type ITypeRightContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeRightContext differentiates from other interfaces.
	IsTypeRightContext()
}

type TypeRightContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRightContext() *TypeRightContext {
	var p = new(TypeRightContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeRight
	return p
}

func (*TypeRightContext) IsTypeRightContext() {}

func NewTypeRightContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRightContext {
	var p = new(TypeRightContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeRight

	return p
}

func (s *TypeRightContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeRightContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRightContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRightContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeRight(s)
	}
}

func (s *TypeRightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeRight(s)
	}
}

func (p *AnzerParser) TypeRight() (localctx ITypeRightContext) {
	localctx = NewTypeRightContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AnzerParserRULE_typeRight)

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
		p.SetState(135)
		p.Match(AnzerParserT__7)
	}

	return localctx
}

// ITypeLeftContext is an interface to support dynamic dispatch.
type ITypeLeftContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeLeftContext differentiates from other interfaces.
	IsTypeLeftContext()
}

type TypeLeftContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeLeftContext() *TypeLeftContext {
	var p = new(TypeLeftContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeLeft
	return p
}

func (*TypeLeftContext) IsTypeLeftContext() {}

func NewTypeLeftContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLeftContext {
	var p = new(TypeLeftContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeLeft

	return p
}

func (s *TypeLeftContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeLeftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeLeftContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeLeftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeLeft(s)
	}
}

func (s *TypeLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeLeft(s)
	}
}

func (p *AnzerParser) TypeLeft() (localctx ITypeLeftContext) {
	localctx = NewTypeLeftContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AnzerParserRULE_typeLeft)

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
		p.SetState(137)
		p.Match(AnzerParserT__8)
	}

	return localctx
}

// ITypeListContext is an interface to support dynamic dispatch.
type ITypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeListContext differentiates from other interfaces.
	IsTypeListContext()
}

type TypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeListContext() *TypeListContext {
	var p = new(TypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeList
	return p
}

func (*TypeListContext) IsTypeListContext() {}

func NewTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeListContext {
	var p = new(TypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeList

	return p
}

func (s *TypeListContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeList(s)
	}
}

func (s *TypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeList(s)
	}
}

func (p *AnzerParser) TypeList() (localctx ITypeListContext) {
	localctx = NewTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AnzerParserRULE_typeList)
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
		p.SetState(139)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AnzerParserT__9 || _la == AnzerParserT__10) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeOptionalContext is an interface to support dynamic dispatch.
type ITypeOptionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeOptionalContext differentiates from other interfaces.
	IsTypeOptionalContext()
}

type TypeOptionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeOptionalContext() *TypeOptionalContext {
	var p = new(TypeOptionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeOptional
	return p
}

func (*TypeOptionalContext) IsTypeOptionalContext() {}

func NewTypeOptionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeOptionalContext {
	var p = new(TypeOptionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeOptional

	return p
}

func (s *TypeOptionalContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeOptionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeOptionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeOptionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeOptional(s)
	}
}

func (s *TypeOptionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeOptional(s)
	}
}

func (p *AnzerParser) TypeOptional() (localctx ITypeOptionalContext) {
	localctx = NewTypeOptionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AnzerParserRULE_typeOptional)
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
		p.SetState(141)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AnzerParserT__11 || _la == AnzerParserT__12) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeScalarContext is an interface to support dynamic dispatch.
type ITypeScalarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeScalarContext differentiates from other interfaces.
	IsTypeScalarContext()
}

type TypeScalarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeScalarContext() *TypeScalarContext {
	var p = new(TypeScalarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeScalar
	return p
}

func (*TypeScalarContext) IsTypeScalarContext() {}

func NewTypeScalarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeScalarContext {
	var p = new(TypeScalarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeScalar

	return p
}

func (s *TypeScalarContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeScalarContext) TypeString() ITypeStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeStringContext)
}

func (s *TypeScalarContext) TypeInteger() ITypeIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeIntegerContext)
}

func (s *TypeScalarContext) TypeFloat() ITypeFloatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFloatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFloatContext)
}

func (s *TypeScalarContext) TypeBool() ITypeBoolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeBoolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeBoolContext)
}

func (s *TypeScalarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeScalarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeScalarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeScalar(s)
	}
}

func (s *TypeScalarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeScalar(s)
	}
}

func (p *AnzerParser) TypeScalar() (localctx ITypeScalarContext) {
	localctx = NewTypeScalarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AnzerParserRULE_typeScalar)

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

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AnzerParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.TypeString()
		}

	case AnzerParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.TypeInteger()
		}

	case AnzerParserT__15:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(145)
			p.TypeFloat()
		}

	case AnzerParserT__16:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(146)
			p.TypeBool()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeStringContext is an interface to support dynamic dispatch.
type ITypeStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeStringContext differentiates from other interfaces.
	IsTypeStringContext()
}

type TypeStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeStringContext() *TypeStringContext {
	var p = new(TypeStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeString
	return p
}

func (*TypeStringContext) IsTypeStringContext() {}

func NewTypeStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeStringContext {
	var p = new(TypeStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeString

	return p
}

func (s *TypeStringContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeString(s)
	}
}

func (s *TypeStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeString(s)
	}
}

func (p *AnzerParser) TypeString() (localctx ITypeStringContext) {
	localctx = NewTypeStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AnzerParserRULE_typeString)

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
		p.SetState(149)
		p.Match(AnzerParserT__13)
	}

	return localctx
}

// ITypeIntegerContext is an interface to support dynamic dispatch.
type ITypeIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeIntegerContext differentiates from other interfaces.
	IsTypeIntegerContext()
}

type TypeIntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeIntegerContext() *TypeIntegerContext {
	var p = new(TypeIntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeInteger
	return p
}

func (*TypeIntegerContext) IsTypeIntegerContext() {}

func NewTypeIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeIntegerContext {
	var p = new(TypeIntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeInteger

	return p
}

func (s *TypeIntegerContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeIntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeIntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeInteger(s)
	}
}

func (s *TypeIntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeInteger(s)
	}
}

func (p *AnzerParser) TypeInteger() (localctx ITypeIntegerContext) {
	localctx = NewTypeIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AnzerParserRULE_typeInteger)

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
		p.SetState(151)
		p.Match(AnzerParserT__14)
	}

	return localctx
}

// ITypeFloatContext is an interface to support dynamic dispatch.
type ITypeFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeFloatContext differentiates from other interfaces.
	IsTypeFloatContext()
}

type TypeFloatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFloatContext() *TypeFloatContext {
	var p = new(TypeFloatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeFloat
	return p
}

func (*TypeFloatContext) IsTypeFloatContext() {}

func NewTypeFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFloatContext {
	var p = new(TypeFloatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeFloat

	return p
}

func (s *TypeFloatContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeFloat(s)
	}
}

func (s *TypeFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeFloat(s)
	}
}

func (p *AnzerParser) TypeFloat() (localctx ITypeFloatContext) {
	localctx = NewTypeFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AnzerParserRULE_typeFloat)

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
		p.SetState(153)
		p.Match(AnzerParserT__15)
	}

	return localctx
}

// ITypeBoolContext is an interface to support dynamic dispatch.
type ITypeBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeBoolContext differentiates from other interfaces.
	IsTypeBoolContext()
}

type TypeBoolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeBoolContext() *TypeBoolContext {
	var p = new(TypeBoolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeBool
	return p
}

func (*TypeBoolContext) IsTypeBoolContext() {}

func NewTypeBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBoolContext {
	var p = new(TypeBoolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeBool

	return p
}

func (s *TypeBoolContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeBool(s)
	}
}

func (s *TypeBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeBool(s)
	}
}

func (p *AnzerParser) TypeBool() (localctx ITypeBoolContext) {
	localctx = NewTypeBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AnzerParserRULE_typeBool)

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
		p.SetState(155)
		p.Match(AnzerParserT__16)
	}

	return localctx
}

// ITypeOtherContext is an interface to support dynamic dispatch.
type ITypeOtherContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeOtherContext differentiates from other interfaces.
	IsTypeOtherContext()
}

type TypeOtherContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeOtherContext() *TypeOtherContext {
	var p = new(TypeOtherContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_typeOther
	return p
}

func (*TypeOtherContext) IsTypeOtherContext() {}

func NewTypeOtherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeOtherContext {
	var p = new(TypeOtherContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_typeOther

	return p
}

func (s *TypeOtherContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeOtherContext) UpperIdent() antlr.TerminalNode {
	return s.GetToken(AnzerParserUpperIdent, 0)
}

func (s *TypeOtherContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeOtherContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeOtherContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterTypeOther(s)
	}
}

func (s *TypeOtherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitTypeOther(s)
	}
}

func (p *AnzerParser) TypeOther() (localctx ITypeOtherContext) {
	localctx = NewTypeOtherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AnzerParserRULE_typeOther)

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
		p.SetState(157)
		p.Match(AnzerParserUpperIdent)
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

func (s *FuncDeclarationContext) FuncArgument() IFuncArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncArgumentContext)
}

func (s *FuncDeclarationContext) FuncResult() IFuncResultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncResultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncResultContext)
}

func (s *FuncDeclarationContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
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
	p.EnterRule(localctx, 46, AnzerParserRULE_funcDeclaration)
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
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AnzerParserLowIdent {
		{
			p.SetState(159)
			p.FuncName()
		}

	}
	{
		p.SetState(162)
		p.Url()
	}
	{
		p.SetState(163)
		p.Match(AnzerParserT__17)
	}
	{
		p.SetState(164)
		p.Runtime()
	}
	{
		p.SetState(165)
		p.Match(AnzerParserT__18)
	}
	{
		p.SetState(166)
		p.Match(AnzerParserT__4)
	}
	{
		p.SetState(167)
		p.FuncArgument()
	}
	{
		p.SetState(168)
		p.Match(AnzerParserT__19)
	}
	{
		p.SetState(169)
		p.FuncResult()
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
	p.EnterRule(localctx, 48, AnzerParserRULE_funcName)

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
		p.SetState(171)
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
	p.EnterRule(localctx, 50, AnzerParserRULE_runtime)

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
		p.SetState(173)
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
	p.EnterRule(localctx, 52, AnzerParserRULE_url)

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
		p.SetState(175)
		p.Match(AnzerParserURL)
	}

	return localctx
}

// IFuncArgumentContext is an interface to support dynamic dispatch.
type IFuncArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncArgumentContext differentiates from other interfaces.
	IsFuncArgumentContext()
}

type FuncArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgumentContext() *FuncArgumentContext {
	var p = new(FuncArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcArgument
	return p
}

func (*FuncArgumentContext) IsFuncArgumentContext() {}

func NewFuncArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgumentContext {
	var p = new(FuncArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcArgument

	return p
}

func (s *FuncArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgumentContext) AllTypeId() []ITypeIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeIdContext)(nil)).Elem())
	var tst = make([]ITypeIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeIdContext)
		}
	}

	return tst
}

func (s *FuncArgumentContext) TypeId(i int) ITypeIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeIdContext)
}

func (s *FuncArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncArgument(s)
	}
}

func (s *FuncArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncArgument(s)
	}
}

func (p *AnzerParser) FuncArgument() (localctx IFuncArgumentContext) {
	localctx = NewFuncArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, AnzerParserRULE_funcArgument)
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
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AnzerParserT__5)|(1<<AnzerParserT__6)|(1<<AnzerParserT__7)|(1<<AnzerParserT__8)|(1<<AnzerParserT__9)|(1<<AnzerParserT__10)|(1<<AnzerParserT__11)|(1<<AnzerParserT__12)|(1<<AnzerParserT__13)|(1<<AnzerParserT__14)|(1<<AnzerParserT__15)|(1<<AnzerParserT__16)|(1<<AnzerParserUpperIdent))) != 0) {
		{
			p.SetState(177)
			p.TypeId()
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFuncResultContext is an interface to support dynamic dispatch.
type IFuncResultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncResultContext differentiates from other interfaces.
	IsFuncResultContext()
}

type FuncResultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncResultContext() *FuncResultContext {
	var p = new(FuncResultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcResult
	return p
}

func (*FuncResultContext) IsFuncResultContext() {}

func NewFuncResultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncResultContext {
	var p = new(FuncResultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcResult

	return p
}

func (s *FuncResultContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncResultContext) AllTypeId() []ITypeIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeIdContext)(nil)).Elem())
	var tst = make([]ITypeIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeIdContext)
		}
	}

	return tst
}

func (s *FuncResultContext) TypeId(i int) ITypeIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeIdContext)
}

func (s *FuncResultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncResultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncResultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncResult(s)
	}
}

func (s *FuncResultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncResult(s)
	}
}

func (p *AnzerParser) FuncResult() (localctx IFuncResultContext) {
	localctx = NewFuncResultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, AnzerParserRULE_funcResult)
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
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AnzerParserT__5)|(1<<AnzerParserT__6)|(1<<AnzerParserT__7)|(1<<AnzerParserT__8)|(1<<AnzerParserT__9)|(1<<AnzerParserT__10)|(1<<AnzerParserT__11)|(1<<AnzerParserT__12)|(1<<AnzerParserT__13)|(1<<AnzerParserT__14)|(1<<AnzerParserT__15)|(1<<AnzerParserT__16)|(1<<AnzerParserUpperIdent))) != 0) {
		{
			p.SetState(182)
			p.TypeId()
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *LocalFuncDeclarationContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *LocalFuncDeclarationContext) AllFuncRef() []IFuncRefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncRefContext)(nil)).Elem())
	var tst = make([]IFuncRefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncRefContext)
		}
	}

	return tst
}

func (s *LocalFuncDeclarationContext) FuncRef(i int) IFuncRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncRefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncRefContext)
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
	p.EnterRule(localctx, 58, AnzerParserRULE_localFuncDeclaration)

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
		p.SetState(187)
		p.FuncName()
	}
	{
		p.SetState(188)
		p.Match(AnzerParserT__1)
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(191)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case AnzerParserLowIdent:
				{
					p.SetState(189)
					p.FuncRef()
				}

			case AnzerParserT__20:
				{
					p.SetState(190)
					p.Match(AnzerParserT__20)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncRefContext is an interface to support dynamic dispatch.
type IFuncRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncRefContext differentiates from other interfaces.
	IsFuncRefContext()
}

type FuncRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncRefContext() *FuncRefContext {
	var p = new(FuncRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_funcRef
	return p
}

func (*FuncRefContext) IsFuncRefContext() {}

func NewFuncRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncRefContext {
	var p = new(FuncRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_funcRef

	return p
}

func (s *FuncRefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncRefContext) LowIdent() antlr.TerminalNode {
	return s.GetToken(AnzerParserLowIdent, 0)
}

func (s *FuncRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterFuncRef(s)
	}
}

func (s *FuncRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitFuncRef(s)
	}
}

func (p *AnzerParser) FuncRef() (localctx IFuncRefContext) {
	localctx = NewFuncRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, AnzerParserRULE_funcRef)

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
		p.SetState(195)
		p.Match(AnzerParserLowIdent)
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

func (s *InvokeCmdContext) AllInvokeFuncName() []IInvokeFuncNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInvokeFuncNameContext)(nil)).Elem())
	var tst = make([]IInvokeFuncNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInvokeFuncNameContext)
		}
	}

	return tst
}

func (s *InvokeCmdContext) InvokeFuncName(i int) IInvokeFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvokeFuncNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInvokeFuncNameContext)
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
	p.EnterRule(localctx, 62, AnzerParserRULE_invokeCmd)
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
		p.SetState(197)
		p.Match(AnzerParserT__21)
	}
	{
		p.SetState(198)
		p.Match(AnzerParserT__22)
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AnzerParserLowIdent {
		{
			p.SetState(199)
			p.InvokeFuncName()
		}
		{
			p.SetState(200)
			p.Match(AnzerParserT__23)
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(206)
		p.Match(AnzerParserT__24)
	}

	return localctx
}

// IInvokeFuncNameContext is an interface to support dynamic dispatch.
type IInvokeFuncNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvokeFuncNameContext differentiates from other interfaces.
	IsInvokeFuncNameContext()
}

type InvokeFuncNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvokeFuncNameContext() *InvokeFuncNameContext {
	var p = new(InvokeFuncNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnzerParserRULE_invokeFuncName
	return p
}

func (*InvokeFuncNameContext) IsInvokeFuncNameContext() {}

func NewInvokeFuncNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvokeFuncNameContext {
	var p = new(InvokeFuncNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnzerParserRULE_invokeFuncName

	return p
}

func (s *InvokeFuncNameContext) GetParser() antlr.Parser { return s.parser }

func (s *InvokeFuncNameContext) LowIdent() antlr.TerminalNode {
	return s.GetToken(AnzerParserLowIdent, 0)
}

func (s *InvokeFuncNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvokeFuncNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvokeFuncNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.EnterInvokeFuncName(s)
	}
}

func (s *InvokeFuncNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnzerListener); ok {
		listenerT.ExitInvokeFuncName(s)
	}
}

func (p *AnzerParser) InvokeFuncName() (localctx IInvokeFuncNameContext) {
	localctx = NewInvokeFuncNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, AnzerParserRULE_invokeFuncName)

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
		p.SetState(208)
		p.Match(AnzerParserLowIdent)
	}

	return localctx
}
