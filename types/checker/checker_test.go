package checker

import (
	"testing"

	"github.com/tariel-x/anzer/types"
)

func bool2point(x bool) *bool {
	return &x
}

func TestSubtypeOnlyType(t *testing.T) {
	typeName1 := types.String
	type1 := types.JsonSchema{
		Type: &typeName1,
	}
	type2 := types.JsonSchema{
		Type: &typeName1,
	}

	ident := Subtype(type1, type2)
	if ident != TypesEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesEqual)
	}
}

func TestSubtypeOnlyTypeFails(t *testing.T) {
	typeName1 := types.String
	typeName2 := types.Number
	type1 := types.JsonSchema{
		Type: &typeName1,
	}
	type2 := types.JsonSchema{
		Type: &typeName2,
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsEqual(t *testing.T) {
	typeName := types.Object
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Required: []string{"a"},
		},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Required: []string{"a"},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesEqual)
	}
}

func TestSubtypeObjectsNotEqual(t *testing.T) {
	typeName := types.Object
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Required: []string{"a"},
		},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Required: []string{"b"},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsNotEqual2(t *testing.T) {
	typeName := types.Object
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Required: []string{"a", "b"},
		},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Required: []string{"b"},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsSubtype(t *testing.T) {
	typeName := types.Object
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Required: []string{"a", "b"},
		},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Required: []string{"a", "b", "c"},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesSubtype {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesSubtype)
	}
}

func TestSubtypeObjectsPropertiesEqual(t *testing.T) {
	typeName := types.Object
	subtypeName := types.String
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &subtypeName,
				},
			},
		},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &subtypeName,
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesEqual)
	}
}

func TestSubtypeObjectsPropertiesNotEqual(t *testing.T) {
	typeName := types.Object
	subtypeName1 := types.String
	subtypeName2 := types.Number
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &subtypeName1,
				},
			},
		},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &subtypeName2,
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsPropertiesNotEqual2(t *testing.T) {
	typeName := types.Object
	subtypeName1 := types.String
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &subtypeName1,
				},
			},
		},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"b": types.JsonSchema{
					Type: &subtypeName1,
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsPropertiesSubtype(t *testing.T) {
	typeName := types.Object
	subtypeName1 := types.String
	subtypeName2 := types.Number
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &subtypeName1,
				},
			},
		},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &subtypeName1,
				},
				"b": types.JsonSchema{
					Type: &subtypeName2,
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesSubtype {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesSubtype)
	}
}

func TestSubtypeObjectsPropertiesNotSubtype(t *testing.T) {
	typeName := types.Object
	subtypeName1 := types.String
	subtypeName2 := types.Number
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &subtypeName1,
				},
			},
			AdditionalProperties: bool2point(false),
		},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &subtypeName1,
				},
				"b": types.JsonSchema{
					Type: &subtypeName2,
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}