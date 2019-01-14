package main

import (
	"fmt"

	"github.com/tariel-x/anzer/internal"
)

func main() {
	fmt.Println("Hello")
	m := internal.Alias{
		Alias: []internal.Composable{
			internal.Func{
				Name: "f",
			},
		},
	}
	fmt.Printf("%#v\n", m)
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
