package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/interpeter"
	"github.com/tariel-x/anzer/parser"
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
	listener := interpeter.NewListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	fmt.Println("-----------------")
	for _, t := range listener.Types {
		fmt.Printf("%s = %s\n", t.Name, t.Val)
	}
	fmt.Println("-----------------")
	for _, t := range listener.Funcs {
		fmt.Printf("%s :: %s -> %s\n", t.Name(), t.Arg(), t.Ret())
	}
}
