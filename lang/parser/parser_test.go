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
	err        :: Either MaxLength 20 String Integer
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
		"GreetingText": lang.Record{
			Fields: map[string]lang.T{
				"text": lang.TypeString,
				"formatting": lang.Constructor{
					Operands:  []lang.T{lang.TypeString},
					Type:      lang.TypeOptional,
					Arguments: []interface{}(nil),
				},
				"err": lang.Constructor{
					Operands: []lang.T{
						lang.Constructor{
							Operands:  []lang.T{lang.TypeString},
							Type:      lang.TypeMaxLength,
							Arguments: []interface{}{20},
						},
						lang.TypeInteger,
					},
					Type:      lang.TypeEither,
					Arguments: []interface{}(nil),
				},
			},
		},
		"Gift": lang.Record{
			Fields: map[string]lang.T{
				"gift": lang.Record{
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
						lang.Record{
							Fields: map[string]lang.T{
								"author": lang.TypeString,
								"text": lang.Record{
									Fields: map[string]lang.T{
										"text": lang.TypeString,
										"formatting": lang.Constructor{
											Operands:  []lang.T{lang.TypeString},
											Type:      lang.TypeOptional,
											Arguments: []interface{}(nil),
										},
										"err": lang.Constructor{
											Operands: []lang.T{
												lang.Constructor{
													Operands:  []lang.T{lang.TypeString},
													Type:      lang.TypeMaxLength,
													Arguments: []interface{}{20},
												},
												lang.TypeInteger,
											},
											Type:      lang.TypeEither,
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
