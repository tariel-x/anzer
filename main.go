package main

import (
	"fmt"

	"github.com/tariel-x/anzer/internal"
)

func main() {
	fmt.Println("Hello")

	fout := internal.TypeString
	ain := fout
	aout := internal.Construct(ain, internal.ConstructorMaxLength, 10)

	s := internal.Alias{
		Compose: []internal.Composable{
			internal.F{
				Name:    "f",
				TypeOut: fout,
			},
			internal.Applied{
				internal.F{
					Name:    "a",
					TypeIn:  ain,
					TypeOut: aout,
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
