// Generated from /home/nikita/go/src/github.com/tariel-x/anzer/parser/Anzer.g4 by ANTLR 4.7.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 28, 216,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	20, 3, 20, 3, 21, 3, 21, 6, 21, 127, 10, 21, 13, 21, 14, 21, 128, 3, 22,
	3, 22, 7, 22, 133, 10, 22, 12, 22, 14, 22, 136, 11, 22, 3, 23, 6, 23, 139,
	10, 23, 13, 23, 14, 23, 140, 3, 24, 3, 24, 3, 24, 7, 24, 146, 10, 24, 12,
	24, 14, 24, 149, 11, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 5, 25, 156,
	10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 29, 5, 29, 169, 10, 29, 3, 29, 3, 29, 3, 29, 6, 29, 174, 10,
	29, 13, 29, 14, 29, 175, 5, 29, 178, 10, 29, 3, 29, 5, 29, 181, 10, 29,
	3, 30, 3, 30, 3, 30, 7, 30, 186, 10, 30, 12, 30, 14, 30, 189, 11, 30, 5,
	30, 191, 10, 30, 3, 31, 3, 31, 5, 31, 195, 10, 31, 3, 31, 3, 31, 3, 32,
	6, 32, 200, 10, 32, 13, 32, 14, 32, 201, 3, 32, 3, 32, 3, 33, 3, 33, 3,
	33, 3, 33, 7, 33, 210, 10, 33, 12, 33, 14, 33, 213, 11, 33, 3, 33, 3, 33,
	2, 2, 34, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 2, 51, 2, 53, 2, 55, 2, 57,
	26, 59, 2, 61, 2, 63, 27, 65, 28, 3, 2, 17, 4, 2, 44, 44, 126, 126, 3,
	2, 67, 92, 5, 2, 50, 59, 67, 92, 97, 97, 3, 2, 99, 124, 5, 2, 50, 59, 67,
	92, 99, 124, 5, 2, 50, 59, 97, 97, 99, 124, 10, 2, 36, 36, 49, 49, 94,
	94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2, 50, 59, 67,
	72, 99, 104, 5, 2, 2, 33, 36, 36, 94, 94, 3, 2, 50, 59, 3, 2, 51, 59, 4,
	2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 5, 2, 11, 12, 15, 15, 34, 34,
	4, 2, 12, 12, 15, 15, 2, 224, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2,
	2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3,
	2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65,
	3, 2, 2, 2, 3, 67, 3, 2, 2, 2, 5, 72, 3, 2, 2, 2, 7, 74, 3, 2, 2, 2, 9,
	77, 3, 2, 2, 2, 11, 80, 3, 2, 2, 2, 13, 82, 3, 2, 2, 2, 15, 88, 3, 2, 2,
	2, 17, 90, 3, 2, 2, 2, 19, 92, 3, 2, 2, 2, 21, 94, 3, 2, 2, 2, 23, 96,
	3, 2, 2, 2, 25, 98, 3, 2, 2, 2, 27, 103, 3, 2, 2, 2, 29, 109, 3, 2, 2,
	2, 31, 114, 3, 2, 2, 2, 33, 116, 3, 2, 2, 2, 35, 118, 3, 2, 2, 2, 37, 120,
	3, 2, 2, 2, 39, 122, 3, 2, 2, 2, 41, 124, 3, 2, 2, 2, 43, 130, 3, 2, 2,
	2, 45, 138, 3, 2, 2, 2, 47, 142, 3, 2, 2, 2, 49, 152, 3, 2, 2, 2, 51, 157,
	3, 2, 2, 2, 53, 163, 3, 2, 2, 2, 55, 165, 3, 2, 2, 2, 57, 168, 3, 2, 2,
	2, 59, 190, 3, 2, 2, 2, 61, 192, 3, 2, 2, 2, 63, 199, 3, 2, 2, 2, 65, 205,
	3, 2, 2, 2, 67, 68, 7, 102, 2, 2, 68, 69, 7, 99, 2, 2, 69, 70, 7, 118,
	2, 2, 70, 71, 7, 99, 2, 2, 71, 4, 3, 2, 2, 2, 72, 73, 7, 63, 2, 2, 73,
	6, 3, 2, 2, 2, 74, 75, 7, 60, 2, 2, 75, 76, 7, 60, 2, 2, 76, 8, 3, 2, 2,
	2, 77, 78, 7, 47, 2, 2, 78, 79, 7, 64, 2, 2, 79, 10, 3, 2, 2, 2, 80, 81,
	7, 48, 2, 2, 81, 12, 3, 2, 2, 2, 82, 83, 7, 48, 2, 2, 83, 84, 7, 103, 2,
	2, 84, 85, 7, 112, 2, 2, 85, 86, 7, 120, 2, 2, 86, 87, 7, 93, 2, 2, 87,
	14, 3, 2, 2, 2, 88, 89, 7, 95, 2, 2, 89, 16, 3, 2, 2, 2, 90, 91, 7, 62,
	2, 2, 91, 18, 3, 2, 2, 2, 92, 93, 7, 46, 2, 2, 93, 20, 3, 2, 2, 2, 94,
	95, 7, 64, 2, 2, 95, 22, 3, 2, 2, 2, 96, 97, 7, 97, 2, 2, 97, 24, 3, 2,
	2, 2, 98, 99, 7, 118, 2, 2, 99, 100, 7, 116, 2, 2, 100, 101, 7, 119, 2,
	2, 101, 102, 7, 103, 2, 2, 102, 26, 3, 2, 2, 2, 103, 104, 7, 104, 2, 2,
	104, 105, 7, 99, 2, 2, 105, 106, 7, 110, 2, 2, 106, 107, 7, 117, 2, 2,
	107, 108, 7, 103, 2, 2, 108, 28, 3, 2, 2, 2, 109, 110, 7, 112, 2, 2, 110,
	111, 7, 119, 2, 2, 111, 112, 7, 110, 2, 2, 112, 113, 7, 110, 2, 2, 113,
	30, 3, 2, 2, 2, 114, 115, 7, 125, 2, 2, 115, 32, 3, 2, 2, 2, 116, 117,
	7, 127, 2, 2, 117, 34, 3, 2, 2, 2, 118, 119, 7, 60, 2, 2, 119, 36, 3, 2,
	2, 2, 120, 121, 7, 93, 2, 2, 121, 38, 3, 2, 2, 2, 122, 123, 9, 2, 2, 2,
	123, 40, 3, 2, 2, 2, 124, 126, 9, 3, 2, 2, 125, 127, 9, 4, 2, 2, 126, 125,
	3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2,
	2, 2, 129, 42, 3, 2, 2, 2, 130, 134, 9, 5, 2, 2, 131, 133, 9, 6, 2, 2,
	132, 131, 3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134,
	135, 3, 2, 2, 2, 135, 44, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 137, 139, 9,
	7, 2, 2, 138, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 138, 3, 2, 2,
	2, 140, 141, 3, 2, 2, 2, 141, 46, 3, 2, 2, 2, 142, 147, 7, 36, 2, 2, 143,
	146, 5, 49, 25, 2, 144, 146, 5, 55, 28, 2, 145, 143, 3, 2, 2, 2, 145, 144,
	3, 2, 2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2,
	2, 2, 148, 150, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 151, 7, 36, 2, 2,
	151, 48, 3, 2, 2, 2, 152, 155, 7, 94, 2, 2, 153, 156, 9, 8, 2, 2, 154,
	156, 5, 51, 26, 2, 155, 153, 3, 2, 2, 2, 155, 154, 3, 2, 2, 2, 156, 50,
	3, 2, 2, 2, 157, 158, 7, 119, 2, 2, 158, 159, 5, 53, 27, 2, 159, 160, 5,
	53, 27, 2, 160, 161, 5, 53, 27, 2, 161, 162, 5, 53, 27, 2, 162, 52, 3,
	2, 2, 2, 163, 164, 9, 9, 2, 2, 164, 54, 3, 2, 2, 2, 165, 166, 10, 10, 2,
	2, 166, 56, 3, 2, 2, 2, 167, 169, 7, 47, 2, 2, 168, 167, 3, 2, 2, 2, 168,
	169, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 177, 5, 59, 30, 2, 171, 173,
	7, 48, 2, 2, 172, 174, 9, 11, 2, 2, 173, 172, 3, 2, 2, 2, 174, 175, 3,
	2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 178, 3, 2, 2,
	2, 177, 171, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 180, 3, 2, 2, 2, 179,
	181, 5, 61, 31, 2, 180, 179, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 58,
	3, 2, 2, 2, 182, 191, 7, 50, 2, 2, 183, 187, 9, 12, 2, 2, 184, 186, 9,
	11, 2, 2, 185, 184, 3, 2, 2, 2, 186, 189, 3, 2, 2, 2, 187, 185, 3, 2, 2,
	2, 187, 188, 3, 2, 2, 2, 188, 191, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190,
	182, 3, 2, 2, 2, 190, 183, 3, 2, 2, 2, 191, 60, 3, 2, 2, 2, 192, 194, 9,
	13, 2, 2, 193, 195, 9, 14, 2, 2, 194, 193, 3, 2, 2, 2, 194, 195, 3, 2,
	2, 2, 195, 196, 3, 2, 2, 2, 196, 197, 5, 59, 30, 2, 197, 62, 3, 2, 2, 2,
	198, 200, 9, 15, 2, 2, 199, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201,
	199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 204,
	8, 32, 2, 2, 204, 64, 3, 2, 2, 2, 205, 206, 7, 49, 2, 2, 206, 207, 7, 49,
	2, 2, 207, 211, 3, 2, 2, 2, 208, 210, 10, 16, 2, 2, 209, 208, 3, 2, 2,
	2, 210, 213, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212,
	214, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 214, 215, 8, 33, 2, 2, 215, 66,
	3, 2, 2, 2, 18, 2, 128, 134, 140, 145, 147, 155, 168, 175, 177, 180, 187,
	190, 194, 201, 211, 3, 8, 2, 2,
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
	"", "'data'", "'='", "'::'", "'->'", "'.'", "'.env['", "']'", "'<'", "','",
	"'>'", "'_'", "'true'", "'false'", "'null'", "'{'", "'}'", "':'", "'['",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "DATA_AND_OR", "DATA_NAME_ID", "FUNC_NAME_ID", "FUNC_PARAM_ID", "STRING",
	"NUMBER", "WS", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "DATA_AND_OR", "DATA_NAME_ID", "FUNC_NAME_ID", "FUNC_PARAM_ID",
	"STRING", "ESC", "UNICODE", "HEX", "SAFECODEPOINT", "NUMBER", "INT", "EXP",
	"WS", "LineComment",
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
	AnzerLexerT__0          = 1
	AnzerLexerT__1          = 2
	AnzerLexerT__2          = 3
	AnzerLexerT__3          = 4
	AnzerLexerT__4          = 5
	AnzerLexerT__5          = 6
	AnzerLexerT__6          = 7
	AnzerLexerT__7          = 8
	AnzerLexerT__8          = 9
	AnzerLexerT__9          = 10
	AnzerLexerT__10         = 11
	AnzerLexerT__11         = 12
	AnzerLexerT__12         = 13
	AnzerLexerT__13         = 14
	AnzerLexerT__14         = 15
	AnzerLexerT__15         = 16
	AnzerLexerT__16         = 17
	AnzerLexerT__17         = 18
	AnzerLexerDATA_AND_OR   = 19
	AnzerLexerDATA_NAME_ID  = 20
	AnzerLexerFUNC_NAME_ID  = 21
	AnzerLexerFUNC_PARAM_ID = 22
	AnzerLexerSTRING        = 23
	AnzerLexerNUMBER        = 24
	AnzerLexerWS            = 25
	AnzerLexerLineComment   = 26
)
