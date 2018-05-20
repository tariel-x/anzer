package main

import "fmt"

//go:generate ragel -Z -G2 -o lex.go lex.rl
//go:generate goyacc thermostat.y

func main() {
	lex := newLexer([]byte(`
		target temperature 5
		heat on
		target temperature 10
		heat off
	`))
	e := yyParse(lex)
	fmt.Println(e)
}
