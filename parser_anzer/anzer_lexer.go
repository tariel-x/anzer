// Code generated from Anzer.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 65, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 6, 8, 42, 10, 8, 13, 8, 14, 8, 43, 3, 9, 6, 9, 47, 10,
	9, 13, 9, 14, 9, 48, 3, 10, 3, 10, 6, 10, 53, 10, 10, 13, 10, 14, 10, 54,
	3, 10, 3, 10, 3, 11, 6, 11, 60, 10, 11, 13, 11, 14, 11, 61, 3, 11, 3, 11,
	2, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 3, 2, 6, 5, 2, 50, 59, 67, 92, 99, 124, 4, 2, 50, 59, 67, 92, 3,
	2, 48, 48, 5, 2, 11, 12, 15, 15, 34, 34, 2, 68, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 3, 23, 3, 2, 2, 2, 5, 28, 3, 2, 2, 2, 7, 30, 3, 2, 2, 2,
	9, 33, 3, 2, 2, 2, 11, 36, 3, 2, 2, 2, 13, 38, 3, 2, 2, 2, 15, 41, 3, 2,
	2, 2, 17, 46, 3, 2, 2, 2, 19, 50, 3, 2, 2, 2, 21, 59, 3, 2, 2, 2, 23, 24,
	7, 102, 2, 2, 24, 25, 7, 99, 2, 2, 25, 26, 7, 118, 2, 2, 26, 27, 7, 99,
	2, 2, 27, 4, 3, 2, 2, 2, 28, 29, 7, 63, 2, 2, 29, 6, 3, 2, 2, 2, 30, 31,
	7, 60, 2, 2, 31, 32, 7, 60, 2, 2, 32, 8, 3, 2, 2, 2, 33, 34, 7, 47, 2,
	2, 34, 35, 7, 64, 2, 2, 35, 10, 3, 2, 2, 2, 36, 37, 7, 97, 2, 2, 37, 12,
	3, 2, 2, 2, 38, 39, 7, 36, 2, 2, 39, 14, 3, 2, 2, 2, 40, 42, 9, 2, 2, 2,
	41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3,
	2, 2, 2, 44, 16, 3, 2, 2, 2, 45, 47, 9, 3, 2, 2, 46, 45, 3, 2, 2, 2, 47,
	48, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 18, 3, 2, 2,
	2, 50, 52, 7, 125, 2, 2, 51, 53, 9, 4, 2, 2, 52, 51, 3, 2, 2, 2, 53, 54,
	3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2,
	56, 57, 7, 127, 2, 2, 57, 20, 3, 2, 2, 2, 58, 60, 9, 5, 2, 2, 59, 58, 3,
	2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62,
	63, 3, 2, 2, 2, 63, 64, 8, 11, 2, 2, 64, 22, 3, 2, 2, 2, 7, 2, 43, 48,
	54, 61, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'data'", "'='", "'::'", "'->'", "'_'", "'\"'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "FUNC_NAME", "DATA_NAME_ID", "JSON_CONTENT",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "FUNC_NAME", "DATA_NAME_ID",
	"JSON_CONTENT", "WS",
}

type AnzerLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewAnzerLexer(input antlr.CharStream) *AnzerLexer {

	l := new(AnzerLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Anzer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AnzerLexer tokens.
const (
	AnzerLexerT__0         = 1
	AnzerLexerT__1         = 2
	AnzerLexerT__2         = 3
	AnzerLexerT__3         = 4
	AnzerLexerT__4         = 5
	AnzerLexerT__5         = 6
	AnzerLexerFUNC_NAME    = 7
	AnzerLexerDATA_NAME_ID = 8
	AnzerLexerJSON_CONTENT = 9
	AnzerLexerWS           = 10
)
