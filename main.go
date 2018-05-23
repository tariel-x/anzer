package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/parser"
	"github.com/tariel-x/anzer/interpeter"
	"os"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewAnzerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAnzerParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.System()
	antlr.ParseTreeWalkerDefault.Walk(interpeter.NewListener(), tree)
}
