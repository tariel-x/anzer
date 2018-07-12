package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tariel-x/anzer/funcs"
	"github.com/tariel-x/anzer/listener"
	"github.com/tariel-x/anzer/parser"
	"github.com/tariel-x/anzer/types"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/fatih/color"
)

var (
	Debug = false
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Specify input and output or -d for debug\n")
		return
	}

	if os.Args[2] == "-d" {
		Debug = true
	}

	rawTypes, rawFuncs, err := readInput(os.Args[1])
	die(err)

	if Debug {
		displayFuncs(rawFuncs)
	}
	types, err := types.Resolve(rawTypes)
	die(err)

	if Debug {
		displayTypes(types)
	}

	sysgraph, err := funcs.Resolve(rawFuncs, types)
	if err != nil {
		fmt.Printf("funcs resolving error: %s\n", err)
	}

	if Debug {
		printOutput(*sysgraph)
	} else {
		writeOutput(*sysgraph, os.Args[2])
	}
}

func readInput(fileName string) (listener.Types, listener.Funcs, error) {
	input, _ := antlr.NewFileStream(fileName)
	lexer := parser.NewAnzerLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAnzerParser(stream)
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.System()
	listener := listener.NewListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Types, listener.Funcs, nil
}

func displayFuncs(fbs listener.Funcs) {
	color.Green("RAW FUNCS:\n\n")

	for name, t := range fbs {
		fmt.Printf("%s :: %s -> %s\n", name, t.Arg, t.Ret)
		fmt.Printf("%s params: %v\n", name, t.Params)
		if t.Body != nil {
			displayFunc(*t.Body)
		}
		fmt.Printf("\n")
	}

	color.Green("END RAW FUNCS\n\n\n")
}

func displayFunc(fb listener.FuncBody) {
	if fb.Ref != nil {
		fmt.Printf(" %s", *fb.Ref)
		if fb.ComposeTo != nil {
			displayFunc(*fb.ComposeTo)
		}
	} else if fb.ProductEls != nil {
		fmt.Print(" <")
		for _, childFd := range fb.ProductEls {
			fmt.Print(",")
			displayFunc(childFd)
		}
		fmt.Print(">")
	}
}

func displayTypes(types types.Types) {
	color.Green("TYPES:\n\n")
	for name, td := range types {
		displayType(name, td)
	}
	color.Green("END TYPES\n\n\n")
}

func displayType(name string, td types.JsonSchema) {
	typeStr, err := json.Marshal(td)
	die(err)
	fmt.Printf("%s: %s\n\n", name, typeStr)
}

func printOutput(system funcs.SystemGraph) {
	jsonSystem, _ := json.Marshal(system)
	fmt.Printf("System:\n\n%s\n", jsonSystem)
}

func writeOutput(system funcs.SystemGraph, output string) {
	jsonSystem, _ := json.MarshalIndent(system, "", "    ")
	err := ioutil.WriteFile(output, jsonSystem, 0644)
	die(err)
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
