package main

import (
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tariel-x/anzer/funcs"
	"github.com/tariel-x/anzer/listener"
	"github.com/tariel-x/anzer/parser"
	"github.com/tariel-x/anzer/types"

	"os"
)

func main() {
	rawTypes, rawFuncs, err := readInput(os.Args[1])
	die(err)

	fmt.Println("-----------------")

	for name, t := range rawFuncs {
		fmt.Printf("%s :: %s -> %s\n", name, t.Arg, t.Ret)
		if t.Body != nil {
			displayFunc(*t.Body)
		}
		fmt.Printf("\n")
	}

	fmt.Println("-----------------")

	types, err := types.Resolve(rawTypes)
	die(err)

	for name, td := range types {
		displayType(name, td)
	}

	sysgraph, err := funcs.Resolve(rawFuncs, types)
	if err != nil {
		fmt.Printf("funcs resolving error: %s\n", err)
	}
	fmt.Printf("services: %v\n", sysgraph.Services)
	fmt.Printf("dependencies: %v\n", sysgraph.Dependencies)
}

func readInput(fileName string) (listener.Types, listener.Funcs, error) {
	input, _ := antlr.NewFileStream(fileName)
	lexer := parser.NewAnzerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAnzerParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.System()
	listener := listener.NewListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Types, listener.Funcs, nil
}

func displayFunc(fd listener.FuncBody) {
	if fd.Ref != nil {
		fmt.Printf(" %s", *fd.Ref)
		if fd.ComposeTo != nil {
			displayFunc(*fd.ComposeTo)
		}
	} else if fd.ProductEls != nil {
		fmt.Print(" <")
		for _, childFd := range fd.ProductEls {
			fmt.Print(",")
			displayFunc(childFd)
		}
		fmt.Print(">")
	}
}

func displayType(name string, td types.JsonSchema) {
	typeStr, err := json.Marshal(td)
	die(err)
	fmt.Printf("%s: %s\n\n", name, typeStr)
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
