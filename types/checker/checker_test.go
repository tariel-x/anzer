package checker

import (
	"testing"

	"github.com/tariel-x/anzer/types"
)

func int2point(x int) *int {
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

func TestSubtypeObjectsPropertiesNotEqual3(t *testing.T) {
	objType := types.Object
	strType := types.String
	numType := types.Number
	type1 := types.JsonSchema{
		Type: &objType,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &objType,
					JSTypeObj: types.JSTypeObj{
						Properties: map[string]types.JsonSchema{
							"a": types.JsonSchema{
								Type: &numType,
							},
						},
					},
				},
			},
		},
	}
	type2 := types.JsonSchema{
		Type: &objType,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &objType,
					JSTypeObj: types.JSTypeObj{
						Properties: map[string]types.JsonSchema{
							"a": types.JsonSchema{
								Type: &strType,
							},
						},
					},
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsPropertiesSubtype2(t *testing.T) {
	objType := types.Object
	strType := types.String
	type1 := types.JsonSchema{
		Type: &objType,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &objType,
					JSTypeObj: types.JSTypeObj{
						Properties: map[string]types.JsonSchema{
							"a": types.JsonSchema{
								Type: &strType,
							},
						},
					},
				},
			},
		},
	}
	type2 := types.JsonSchema{
		Type: &objType,
		JSTypeObj: types.JSTypeObj{
			Properties: map[string]types.JsonSchema{
				"a": types.JsonSchema{
					Type: &objType,
					JSTypeObj: types.JSTypeObj{
						Properties: map[string]types.JsonSchema{
							"a": types.JsonSchema{
								Type: &strType,
							},
							"b": types.JsonSchema{
								Type: &strType,
							},
						},
					},
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesSubtype {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesSubtype)
	}
}

func TestSubtypeStringCheck(t *testing.T) {
	typeName := types.String
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeString: types.JSTypeString{
			MaxLength: int2point(1),
		},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeString: types.JSTypeString{
			MaxLength: int2point(1),
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesEqual)
	}
}

func TestSubtypeStringCheck2(t *testing.T) {
	typeName := types.String
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeString: types.JSTypeString{},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeString: types.JSTypeString{
			MaxLength: int2point(1),
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesSubtype {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesSubtype)
	}
}

func TestSubtypeStringCheckNotEqual(t *testing.T) {
	typeName := types.String
	type1 := types.JsonSchema{
		Type: &typeName,
		JSTypeString: types.JSTypeString{
			MaxLength: int2point(1),
		},
	}
	type2 := types.JsonSchema{
		Type: &typeName,
		JSTypeString: types.JSTypeString{
			MaxLength: int2point(2),
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

