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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 35, 273,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 33, 5, 33, 214, 10, 33, 3, 34, 3, 34, 3, 34,
	7, 34, 219, 10, 34, 12, 34, 14, 34, 222, 11, 34, 3, 35, 3, 35, 3, 35, 7,
	35, 227, 10, 35, 12, 35, 14, 35, 230, 11, 35, 3, 36, 6, 36, 233, 10, 36,
	13, 36, 14, 36, 234, 3, 37, 3, 37, 6, 37, 239, 10, 37, 13, 37, 14, 37,
	240, 3, 37, 3, 37, 6, 37, 245, 10, 37, 13, 37, 14, 37, 246, 3, 37, 3, 37,
	3, 37, 6, 37, 252, 10, 37, 13, 37, 14, 37, 253, 3, 38, 3, 38, 3, 38, 3,
	38, 7, 38, 260, 10, 38, 12, 38, 14, 38, 263, 11, 38, 3, 38, 3, 38, 3, 39,
	6, 39, 268, 10, 39, 13, 39, 14, 39, 269, 3, 39, 3, 39, 2, 2, 40, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 2, 59, 2, 61,
	2, 63, 2, 65, 2, 67, 30, 69, 31, 71, 32, 73, 33, 75, 34, 77, 35, 3, 2,
	8, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 99, 124, 3, 2, 67, 92, 3, 2, 50,
	59, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 15, 15, 34, 34, 2, 281, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2,
	2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2,
	2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 67, 3,
	2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75,
	3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 3, 79, 3, 2, 2, 2, 5, 84, 3, 2, 2, 2, 7,
	86, 3, 2, 2, 2, 9, 88, 3, 2, 2, 2, 11, 90, 3, 2, 2, 2, 13, 93, 3, 2, 2,
	2, 15, 103, 3, 2, 2, 2, 17, 113, 3, 2, 2, 2, 19, 120, 3, 2, 2, 2, 21, 126,
	3, 2, 2, 2, 23, 131, 3, 2, 2, 2, 25, 136, 3, 2, 2, 2, 27, 139, 3, 2, 2,
	2, 29, 148, 3, 2, 2, 2, 31, 150, 3, 2, 2, 2, 33, 157, 3, 2, 2, 2, 35, 165,
	3, 2, 2, 2, 37, 171, 3, 2, 2, 2, 39, 176, 3, 2, 2, 2, 41, 178, 3, 2, 2,
	2, 43, 180, 3, 2, 2, 2, 45, 183, 3, 2, 2, 2, 47, 185, 3, 2, 2, 2, 49, 189,
	3, 2, 2, 2, 51, 196, 3, 2, 2, 2, 53, 198, 3, 2, 2, 2, 55, 200, 3, 2, 2,
	2, 57, 202, 3, 2, 2, 2, 59, 204, 3, 2, 2, 2, 61, 206, 3, 2, 2, 2, 63, 208,
	3, 2, 2, 2, 65, 213, 3, 2, 2, 2, 67, 215, 3, 2, 2, 2, 69, 223, 3, 2, 2,
	2, 71, 232, 3, 2, 2, 2, 73, 238, 3, 2, 2, 2, 75, 255, 3, 2, 2, 2, 77, 267,
	3, 2, 2, 2, 79, 80, 7, 118, 2, 2, 80, 81, 7, 123, 2, 2, 81, 82, 7, 114,
	2, 2, 82, 83, 7, 103, 2, 2, 83, 4, 3, 2, 2, 2, 84, 85, 7, 63, 2, 2, 85,
	6, 3, 2, 2, 2, 86, 87, 7, 125, 2, 2, 87, 8, 3, 2, 2, 2, 88, 89, 7, 127,
	2, 2, 89, 10, 3, 2, 2, 2, 90, 91, 7, 60, 2, 2, 91, 92, 7, 60, 2, 2, 92,
	12, 3, 2, 2, 2, 93, 94, 7, 79, 2, 2, 94, 95, 7, 107, 2, 2, 95, 96, 7, 112,
	2, 2, 96, 97, 7, 78, 2, 2, 97, 98, 7, 103, 2, 2, 98, 99, 7, 112, 2, 2,
	99, 100, 7, 105, 2, 2, 100, 101, 7, 118, 2, 2, 101, 102, 7, 106, 2, 2,
	102, 14, 3, 2, 2, 2, 103, 104, 7, 79, 2, 2, 104, 105, 7, 99, 2, 2, 105,
	106, 7, 122, 2, 2, 106, 107, 7, 78, 2, 2, 107, 108, 7, 103, 2, 2, 108,
	109, 7, 112, 2, 2, 109, 110, 7, 105, 2, 2, 110, 111, 7, 118, 2, 2, 111,
	112, 7, 106, 2, 2, 112, 16, 3, 2, 2, 2, 113, 114, 7, 71, 2, 2, 114, 115,
	7, 107, 2, 2, 115, 116, 7, 118, 2, 2, 116, 117, 7, 106, 2, 2, 117, 118,
	7, 103, 2, 2, 118, 119, 7, 116, 2, 2, 119, 18, 3, 2, 2, 2, 120, 121, 7,
	84, 2, 2, 121, 122, 7, 107, 2, 2, 122, 123, 7, 105, 2, 2, 123, 124, 7,
	106, 2, 2, 124, 125, 7, 118, 2, 2, 125, 20, 3, 2, 2, 2, 126, 127, 7, 78,
	2, 2, 127, 128, 7, 103, 2, 2, 128, 129, 7, 104, 2, 2, 129, 130, 7, 118,
	2, 2, 130, 22, 3, 2, 2, 2, 131, 132, 7, 78, 2, 2, 132, 133, 7, 107, 2,
	2, 133, 134, 7, 117, 2, 2, 134, 135, 7, 118, 2, 2, 135, 24, 3, 2, 2, 2,
	136, 137, 7, 93, 2, 2, 137, 138, 7, 95, 2, 2, 138, 26, 3, 2, 2, 2, 139,
	140, 7, 81, 2, 2, 140, 141, 7, 114, 2, 2, 141, 142, 7, 118, 2, 2, 142,
	143, 7, 107, 2, 2, 143, 144, 7, 113, 2, 2, 144, 145, 7, 112, 2, 2, 145,
	146, 7, 99, 2, 2, 146, 147, 7, 110, 2, 2, 147, 28, 3, 2, 2, 2, 148, 149,
	7, 44, 2, 2, 149, 30, 3, 2, 2, 2, 150, 151, 7, 85, 2, 2, 151, 152, 7, 118,
	2, 2, 152, 153, 7, 116, 2, 2, 153, 154, 7, 107, 2, 2, 154, 155, 7, 112,
	2, 2, 155, 156, 7, 105, 2, 2, 156, 32, 3, 2, 2, 2, 157, 158, 7, 75, 2,
	2, 158, 159, 7, 112, 2, 2, 159, 160, 7, 118, 2, 2, 160, 161, 7, 103, 2,
	2, 161, 162, 7, 105, 2, 2, 162, 163, 7, 103, 2, 2, 163, 164, 7, 116, 2,
	2, 164, 34, 3, 2, 2, 2, 165, 166, 7, 72, 2, 2, 166, 167, 7, 110, 2, 2,
	167, 168, 7, 113, 2, 2, 168, 169, 7, 99, 2, 2, 169, 170, 7, 118, 2, 2,
	170, 36, 3, 2, 2, 2, 171, 172, 7, 68, 2, 2, 172, 173, 7, 113, 2, 2, 173,
	174, 7, 113, 2, 2, 174, 175, 7, 110, 2, 2, 175, 38, 3, 2, 2, 2, 176, 177,
	7, 93, 2, 2, 177, 40, 3, 2, 2, 2, 178, 179, 7, 95, 2, 2, 179, 42, 3, 2,
	2, 2, 180, 181, 7, 47, 2, 2, 181, 182, 7, 64, 2, 2, 182, 44, 3, 2, 2, 2,
	183, 184, 7, 48, 2, 2, 184, 46, 3, 2, 2, 2, 185, 186, 7, 64, 2, 2, 186,
	187, 7, 64, 2, 2, 187, 188, 7, 63, 2, 2, 188, 48, 3, 2, 2, 2, 189, 190,
	7, 107, 2, 2, 190, 191, 7, 112, 2, 2, 191, 192, 7, 120, 2, 2, 192, 193,
	7, 113, 2, 2, 193, 194, 7, 109, 2, 2, 194, 195, 7, 103, 2, 2, 195, 50,
	3, 2, 2, 2, 196, 197, 7, 42, 2, 2, 197, 52, 3, 2, 2, 2, 198, 199, 7, 46,
	2, 2, 199, 54, 3, 2, 2, 2, 200, 201, 7, 43, 2, 2, 201, 56, 3, 2, 2, 2,
	202, 203, 9, 2, 2, 2, 203, 58, 3, 2, 2, 2, 204, 205, 9, 3, 2, 2, 205, 60,
	3, 2, 2, 2, 206, 207, 9, 4, 2, 2, 207, 62, 3, 2, 2, 2, 208, 209, 9, 5,
	2, 2, 209, 64, 3, 2, 2, 2, 210, 214, 7, 47, 2, 2, 211, 214, 5, 57, 29,
	2, 212, 214, 5, 63, 32, 2, 213, 210, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2,
	213, 212, 3, 2, 2, 2, 214, 66, 3, 2, 2, 2, 215, 220, 5, 59, 30, 2, 216,
	219, 5, 57, 29, 2, 217, 219, 5, 63, 32, 2, 218, 216, 3, 2, 2, 2, 218, 217,
	3, 2, 2, 2, 219, 222, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 220, 221, 3, 2,
	2, 2, 221, 68, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 223, 228, 5, 61, 31, 2,
	224, 227, 5, 57, 29, 2, 225, 227, 5, 63, 32, 2, 226, 224, 3, 2, 2, 2, 226,
	225, 3, 2, 2, 2, 227, 230, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 228, 229,
	3, 2, 2, 2, 229, 70, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 231, 233, 5, 63,
	32, 2, 232, 231, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2,
	234, 235, 3, 2, 2, 2, 235, 72, 3, 2, 2, 2, 236, 239, 5, 65, 33, 2, 237,
	239, 7, 48, 2, 2, 238, 236, 3, 2, 2, 2, 238, 237, 3, 2, 2, 2, 239, 240,
	3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 242, 3, 2,
	2, 2, 242, 244, 7, 48, 2, 2, 243, 245, 5, 57, 29, 2, 244, 243, 3, 2, 2,
	2, 245, 246, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247,
	248, 3, 2, 2, 2, 248, 251, 7, 49, 2, 2, 249, 252, 5, 65, 33, 2, 250, 252,
	7, 49, 2, 2, 251, 249, 3, 2, 2, 2, 251, 250, 3, 2, 2, 2, 252, 253, 3, 2,
	2, 2, 253, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 74, 3, 2, 2, 2,
	255, 256, 7, 49, 2, 2, 256, 257, 7, 49, 2, 2, 257, 261, 3, 2, 2, 2, 258,
	260, 10, 6, 2, 2, 259, 258, 3, 2, 2, 2, 260, 263, 3, 2, 2, 2, 261, 259,
	3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 264, 3, 2, 2, 2, 263, 261, 3, 2,
	2, 2, 264, 265, 8, 38, 2, 2, 265, 76, 3, 2, 2, 2, 266, 268, 9, 7, 2, 2,
	267, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 269,
	270, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 272, 8, 39, 2, 2, 272, 78,
	3, 2, 2, 2, 16, 2, 213, 218, 220, 226, 228, 234, 238, 240, 246, 251, 253,
	261, 269, 3, 8, 2, 2,
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
	"", "'type'", "'='", "'{'", "'}'", "'::'", "'MinLength'", "'MaxLength'",
	"'Either'", "'Right'", "'Left'", "'List'", "'[]'", "'Optional'", "'*'",
	"'String'", "'Integer'", "'Float'", "'Bool'", "'['", "']'", "'->'", "'.'",
	"'>>='", "'invoke'", "'('", "','", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "LowIdent", "UpperIdent", "ConstructorArg",
	"URL", "LINE_COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "Letter", "LowLetter", "UpperLetter", "DecimalDigit",
	"Urlpart", "LowIdent", "UpperIdent", "ConstructorArg", "URL", "LINE_COMMENT",
	"WS",
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
	AnzerLexerT__0           = 1
	AnzerLexerT__1           = 2
	AnzerLexerT__2           = 3
	AnzerLexerT__3           = 4
	AnzerLexerT__4           = 5
	AnzerLexerT__5           = 6
	AnzerLexerT__6           = 7
	AnzerLexerT__7           = 8
	AnzerLexerT__8           = 9
	AnzerLexerT__9           = 10
	AnzerLexerT__10          = 11
	AnzerLexerT__11          = 12
	AnzerLexerT__12          = 13
	AnzerLexerT__13          = 14
	AnzerLexerT__14          = 15
	AnzerLexerT__15          = 16
	AnzerLexerT__16          = 17
	AnzerLexerT__17          = 18
	AnzerLexerT__18          = 19
	AnzerLexerT__19          = 20
	AnzerLexerT__20          = 21
	AnzerLexerT__21          = 22
	AnzerLexerT__22          = 23
	AnzerLexerT__23          = 24
	AnzerLexerT__24          = 25
	AnzerLexerT__25          = 26
	AnzerLexerT__26          = 27
	AnzerLexerLowIdent       = 28
	AnzerLexerUpperIdent     = 29
	AnzerLexerConstructorArg = 30
	AnzerLexerURL            = 31
	AnzerLexerLINE_COMMENT   = 32
	AnzerLexerWS             = 33
)
