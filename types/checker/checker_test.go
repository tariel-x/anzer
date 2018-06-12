package checker

import (
	"testing"

	"github.com/tariel-x/anzer/types"
)

func TestSubtypeOnlyType(t *testing.T) {
	typeName1 := types.TypeName("string")
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