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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 67, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 6, 7, 40, 10, 7, 13, 7, 14, 7, 41, 3, 8, 6, 8, 45, 10, 8, 13, 8,
	14, 8, 46, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 6, 10, 55, 10, 10, 13,
	10, 14, 10, 56, 3, 10, 3, 10, 3, 11, 6, 11, 62, 10, 11, 13, 11, 14, 11,
	63, 3, 11, 3, 11, 2, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 12, 3, 2, 6, 5, 2, 50, 59, 67, 92, 99, 124, 4, 2, 50,
	59, 67, 92, 3, 2, 48, 48, 5, 2, 11, 12, 15, 15, 34, 34, 2, 70, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 23, 3, 2, 2, 2, 5, 28, 3, 2, 2, 2, 7,
	30, 3, 2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 36, 3, 2, 2, 2, 13, 39, 3, 2, 2,
	2, 15, 44, 3, 2, 2, 2, 17, 48, 3, 2, 2, 2, 19, 52, 3, 2, 2, 2, 21, 61,
	3, 2, 2, 2, 23, 24, 7, 102, 2, 2, 24, 25, 7, 99, 2, 2, 25, 26, 7, 118,
	2, 2, 26, 27, 7, 99, 2, 2, 27, 4, 3, 2, 2, 2, 28, 29, 7, 63, 2, 2, 29,
	6, 3, 2, 2, 2, 30, 31, 7, 60, 2, 2, 31, 32, 7, 60, 2, 2, 32, 8, 3, 2, 2,
	2, 33, 34, 7, 47, 2, 2, 34, 35, 7, 64, 2, 2, 35, 10, 3, 2, 2, 2, 36, 37,
	7, 97, 2, 2, 37, 12, 3, 2, 2, 2, 38, 40, 9, 2, 2, 2, 39, 38, 3, 2, 2, 2,
	40, 41, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 14, 3,
	2, 2, 2, 43, 45, 9, 3, 2, 2, 44, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46,
	44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 16, 3, 2, 2, 2, 48, 49, 7, 36,
	2, 2, 49, 50, 5, 19, 10, 2, 50, 51, 7, 36, 2, 2, 51, 18, 3, 2, 2, 2, 52,
	54, 7, 125, 2, 2, 53, 55, 9, 4, 2, 2, 54, 53, 3, 2, 2, 2, 55, 56, 3, 2,
	2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59,
	7, 127, 2, 2, 59, 20, 3, 2, 2, 2, 60, 62, 9, 5, 2, 2, 61, 60, 3, 2, 2,
	2, 62, 63, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 65,
	3, 2, 2, 2, 65, 66, 8, 11, 2, 2, 66, 22, 3, 2, 2, 2, 7, 2, 41, 46, 56,
	63, 3, 8, 2, 2,
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
	"", "'data'", "'='", "'::'", "'->'", "'_'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "FUNC_NAME", "DATA_NAME", "DATA_CONTENT", "JSON_STRING",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "FUNC_NAME", "DATA_NAME", "DATA_CONTENT",
	"JSON_STRING", "WS",
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
	AnzerLexerFUNC_NAME    = 6
	AnzerLexerDATA_NAME    = 7
	AnzerLexerDATA_CONTENT = 8
	AnzerLexerJSON_STRING  = 9
	AnzerLexerWS           = 10
)
