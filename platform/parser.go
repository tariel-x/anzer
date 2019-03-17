package platform

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/lang/parser"
)

func Parse(sourceStream io.Reader) ([]lang.F, error) {
	source, err := ioutil.ReadAll(sourceStream)
	if err != nil {
		return nil, err
	}

	p := parser.New(string(source))
	result, err := p.Parse()
	if err != nil {
		return nil, err
	}

	fmt.Printf("%#v\n\n", result.Types)
	fmt.Printf("%+v\n\n", result.Funcs)

	return nil, nil
}

func tst() {
	tt := map[string]lang.T{
		"GreetingText": lang.Complex{
			Fields: map[string]lang.T{
				"text": lang.TypeString,
				"formatting": lang.Constructor{
					Operand:   lang.TypeString,
					Type:      lang.TypeOptional,
					Arguments: []interface{}(nil),
				},
			},
		},
		"Gift": lang.Complex{
			Fields: map[string]lang.T{
				"gift": lang.Complex{
					Fields: map[string]lang.T{
						"name": lang.TypeString,
						"size": lang.TypeInteger,
						"age": lang.Constructor{
							Operand:   lang.TypeInteger,
							Type:      lang.TypeOptional,
							Arguments: []interface{}(nil),
						},
					},
				},
				"greeting": lang.Constructor{
					Operand: lang.Complex{
						Fields: map[string]lang.T{
							"author": lang.TypeString,
							"text": lang.Complex{
								Fields: map[string]lang.T{
									"text": lang.TypeString,
									"formatting": lang.Constructor{
										Operand:   lang.TypeString,
										Type:      lang.TypeOptional,
										Arguments: []interface{}(nil),
									},
								},
							},
						},
					},
					Type:      lang.TypeList,
					Arguments: []interface{}(nil),
				},
				"address": lang.Constructor{
					Operand: lang.Constructor{
						Operand:   lang.TypeString,
						Type:      lang.TypeMaxLength,
						Arguments: []interface{}{20},
					},
					Type:      lang.TypeMinLength,
					Arguments: []interface{}{10},
				},
			},
		},
		"DeliverResult": lang.TypeBool,
	}
	fmt.Println(tt)
}
