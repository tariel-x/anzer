package main

import (
	"fmt"

	"github.com/tariel-x/anzer/internal"
)

func main() {
	fmt.Println("Hello")
	s := internal.Alias{
		Compose: []internal.Composable{
			internal.F{
				Name: "f",
			},
			internal.Applied{
				internal.F{
					Name: "a",
				},
				internal.EitherBind(true),
				internal.F{
					Name: "b",
				},
			},
		},
	}

	m := internal.Alias{
		Compose: []internal.Composable{
			s,
			internal.F{
				Name: "abc",
			},
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
