package main

import (
	"fmt"

	"github.com/tariel-x/anzer/internal"
)

func main() {
	s := internal.Alias{
		Name: "composition1",
		Compose: []internal.Composable{
			internal.F{
				Name:    "f",
				Link:    "github.com/tariel-x/f",
				TypeIn:  internal.TypeString,
				TypeOut: internal.Construct(internal.TypeString, internal.MaxLength, 10),
			},
			internal.F{
				Name: "a",
				Link: "github.com/tariel-x/a",
			},
			internal.F{
				Name: "b",
				Link: "github.com/tariel-x/b",
			},
		},
	}

	m := internal.Alias{
		Name: "compisition2",
		Compose: []internal.Composable{
			internal.F{
				Link: "github.com/tariel-x/abc",
				Name: "abc",
			},
			s,
		},
	}

	//fmt.Printf("%#v\n", s)
	fmt.Println(m.Definition())
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
