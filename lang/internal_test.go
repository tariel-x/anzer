package lang

import (
	"testing"
)

func TestAliasValid(t *testing.T) {
	c := Alias{
		Name: "c",
		Compose: []Composable{
			F{
				Name:    "a",
				TypeIn:  TypeString,
				TypeOut: MaxLength(TypeString, 10),
			},
			F{
				Name:   "b",
				TypeIn: MaxLength(TypeString, 10),
				TypeOut: Complex{
					Fields: map[string]T{
						"f1": TypeInteger,
						"f2": TypeString,
					},
				},
			},
			F{
				Name: "c",
				TypeIn: Complex{
					Fields: map[string]T{
						"f1": TypeInteger,
						"f2": TypeString,
					},
				},
				TypeOut: TypeBool,
			},
		},
	}
	err := c.Invalid()
	if err != nil {
		t.Errorf("c must be valid, but: %s", err)
	}
}

func TestAliasInvalid(t *testing.T) {
	c := Alias{
		Name: "c",
		Compose: []Composable{
			F{
				Name:    "a",
				TypeIn:  TypeString,
				TypeOut: TypeString,
			},
			F{
				Name:   "b",
				TypeIn: MaxLength(TypeString, 10),
				TypeOut: Complex{
					Fields: map[string]T{
						"f2": TypeString,
					},
				},
			},
			F{
				Name: "c",
				TypeIn: Complex{
					Fields: map[string]T{
						"f1": TypeInteger,
						"f2": TypeString,
					},
				},
				TypeOut: TypeBool,
			},
		},
	}
	err := c.Invalid()
	if err == nil {
		t.Error("c must be invalid, but err is nil")
	}
}
