package checker

import "github.com/tariel-x/anzer/types"

func validateString(parent, child types.JsonSchema) TypesIdentity {
	return TypesEqual
}