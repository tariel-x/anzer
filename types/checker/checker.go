package checker

import (
	"github.com/tariel-x/anzer/types"
)

func Equal(schema1, schema2 types.JsonSchema) bool {
	return false
}

func Subtype(parent, child types.JsonSchema) bool {
	return false
}