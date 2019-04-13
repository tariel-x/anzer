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

func TestBindValid(t *testing.T) {
	errT := Complex{
		Fields: map[string]T{
			"error": TypeString,
			"code":  TypeInteger,
		},
	}
	c := Alias{
		Name: "c",
		Compose: []Composable{
			F{
				Name:    "a",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, errT),
			},
			Bind(F{
				Name:    "b",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, errT),
			}),
			F{
				Name:    "c",
				TypeIn:  Either(TypeString, errT),
				TypeOut: Optional(errT),
			},
		},
	}
	err := c.Invalid()
	if err != nil {
		t.Errorf("c must be valid, but err is %s", err)
	}
}

func TestBindInvalid1(t *testing.T) {
	errT := Complex{
		Fields: map[string]T{
			"error": TypeString,
			"code":  TypeInteger,
		},
	}
	c := Alias{
		Name: "c",
		Compose: []Composable{
			F{
				Name:    "a",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, errT),
			},
			Bind(F{
				Name:    "b",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, errT),
			}),
			F{
				Name:    "c",
				TypeIn:  Either(TypeInteger, errT),
				TypeOut: Optional(errT),
			},
		},
	}
	err := c.Invalid()
	if err == nil {
		t.Error("c must be invalid, but err is nil")
	}
}

func TestBindInvalid2(t *testing.T) {
	errT := Complex{
		Fields: map[string]T{
			"error": TypeString,
			"code":  TypeInteger,
		},
	}
	c := Alias{
		Name: "c",
		Compose: []Composable{
			F{
				Name:    "a",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, errT),
			},
			Bind(F{
				Name:    "b",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, errT),
			}),
			F{
				Name:    "c",
				TypeIn:  TypeString,
				TypeOut: Optional(errT),
			},
		},
	}
	err := c.Invalid()
	if err == nil {
		t.Error("c must be invalid, but err is nil")
	}
}

func TestBindInvalid3(t *testing.T) {
	errT := Complex{
		Fields: map[string]T{
			"error": TypeString,
			"code":  TypeInteger,
		},
	}
	c := Alias{
		Name: "c",
		Compose: []Composable{
			F{
				Name:    "a",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, errT),
			},
			Bind(F{
				Name:    "b",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, TypeInteger),
			}),
			F{
				Name:    "c",
				TypeIn:  Either(TypeString, errT),
				TypeOut: Optional(errT),
			},
		},
	}
	err := c.Invalid()
	if err == nil {
		t.Error("c must be invalid, but err is nil")
	}
}

func TestBindInvalid4(t *testing.T) {
	errT := Complex{
		Fields: map[string]T{
			"error": TypeString,
			"code":  TypeInteger,
		},
	}
	c := Alias{
		Name: "c",
		Compose: []Composable{
			F{
				Name:    "a",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, TypeInteger),
			},
			Bind(F{
				Name:    "b",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, errT),
			}),
			F{
				Name:    "c",
				TypeIn:  Either(TypeString, errT),
				TypeOut: Optional(errT),
			},
		},
	}
	err := c.Invalid()
	if err == nil {
		t.Error("c must be invalid, but err is nil")
	}
}

func TestBindInvalid5(t *testing.T) {
	errT := Complex{
		Fields: map[string]T{
			"error": TypeString,
			"code":  TypeInteger,
		},
	}
	c := Alias{
		Name: "c",
		Compose: []Composable{
			F{
				Name:    "a",
				TypeIn:  TypeString,
				TypeOut: TypeString,
			},
			Bind(F{
				Name:    "b",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, errT),
			}),
			F{
				Name:    "c",
				TypeIn:  Either(TypeString, errT),
				TypeOut: Optional(errT),
			},
		},
	}
	err := c.Invalid()
	if err == nil {
		t.Error("c must be invalid, but err is nil")
	}
}

func TestReturnValid(t *testing.T) {
	errT := Complex{
		Fields: map[string]T{
			"error": TypeString,
			"code[": TypeInteger,
		},
	}
	c := Alias{
		Name: "c",
		Compose: []Composable{
			F{
				Name:    "a",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, errT),
			},
			Return(F{
				Name:    "b",
				TypeIn:  TypeString,
				TypeOut: TypeString,
			}),
			F{
				Name:    "c",
				TypeIn:  Either(TypeString, errT),
				TypeOut: Optional(errT),
			},
		},
	}
	err := c.Invalid()
	if err != nil {
		t.Errorf("c must be valid, but err is %s", err)
	}
}

func TestReturnInvalid1(t *testing.T) {
	errT := Complex{
		Fields: map[string]T{
			"error": TypeString,
			"code[": TypeInteger,
		},
	}
	c := Alias{
		Name: "c",
		Compose: []Composable{
			F{
				Name:    "a",
				TypeIn:  TypeString,
				TypeOut: Either(TypeString, errT),
			},
			Return(F{
				Name:    "b",
				TypeIn:  TypeString,
				TypeOut: TypeString,
			}),
			F{
				Name:    "c",
				TypeIn:  TypeString,
				TypeOut: Optional(errT),
			},
		},
	}
	err := c.Invalid()
	if err == nil {
		t.Error("c must be invalid, but err is nil")
	}
}
