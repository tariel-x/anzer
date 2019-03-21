package parser

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/tariel-x/anzer/lang"
)

var (
	Source = `
type GreetingText = {
    text       :: String
    formatting :: *String
}
//Comment for testing
type Gift = {
    address :: MinLength 10 MaxLength 20 String
    gift    :: {
        name :: String
        size :: Integer
        age  :: *Integer
    }
    greeting :: []{
        author :: String
        text   :: GreetingText
    }
}
type DeliverResult = Integer
`

	Expected = map[string]lang.T{
		"GreetingText": lang.Complex{
			Fields: map[string]lang.T{
				"text": lang.TypeString,
				"formatting": lang.Constructor{
					Operands:  []lang.T{lang.TypeString},
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
							Operands:  []lang.T{lang.TypeInteger},
							Type:      lang.TypeOptional,
							Arguments: []interface{}(nil),
						},
					},
				},
				"greeting": lang.Constructor{
					Operands: []lang.T{
						lang.Complex{
							Fields: map[string]lang.T{
								"author": lang.TypeString,
								"text": lang.Complex{
									Fields: map[string]lang.T{
										"text": lang.TypeString,
										"formatting": lang.Constructor{
											Operands:  []lang.T{lang.TypeString},
											Type:      lang.TypeOptional,
											Arguments: []interface{}(nil),
										},
									},
								},
							},
						},
					},
					Type:      lang.TypeList,
					Arguments: []interface{}(nil),
				},
				"address": lang.MinLength(lang.MaxLength(lang.TypeString, 20), 10),
			},
		},
		"DeliverResult": lang.TypeInteger,
	}
)

func TestParseType(t *testing.T) {
	parser := New(Source)
	result, err := parser.ParseTypes()
	if err != nil {
		t.Error(err)
	}

	if diff := deep.Equal(result, Expected); diff != nil {
		t.Error(diff)
	}
}
