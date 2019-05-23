package parser

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/tariel-x/anzer/lang"
)

var (
	SourceType = `
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

	ExpectedType = map[string]lang.T{
		"GreetingText": lang.Record{
			Fields: map[string]lang.T{
				"text": lang.TypeString,
				"formatting": lang.Sum{
					lang.NothingType{},
					lang.Just(lang.TypeString),
				},
				"err": lang.Sum{
					lang.Left(lang.MaxLength(lang.TypeString, 20)),
					lang.Right(lang.TypeInteger),
				},
			},
		},
		"Gift": lang.Record{
			Fields: map[string]lang.T{
				"gift": lang.Record{
					Fields: map[string]lang.T{
						"name": lang.TypeString,
						"size": lang.TypeInteger,
						"age":  lang.Optional(lang.TypeInteger),
					},
				},
				"greeting": lang.List(lang.Record{
					Fields: map[string]lang.T{
						"author": lang.TypeString,
						"text": lang.Record{
							Fields: map[string]lang.T{
								"text":       lang.TypeString,
								"formatting": lang.Optional(lang.TypeString),
								"err":        lang.Either(lang.MaxLength(lang.TypeString, 20), lang.TypeInteger),
							},
						},
					},
				}),
				"address": lang.MinLength(lang.MaxLength(lang.TypeString, 20), 10),
			},
		},
		"DeliverResult": lang.TypeInteger,
	}
)

func TestParseType(t *testing.T) {
	parser := New(SourceType)
	result, err := parser.ParseTypes()
	if err != nil {
		t.Error(err)
	}

	if diff := deep.Equal(result, ExpectedType); diff != nil {
		t.Error(diff)
	}
}
