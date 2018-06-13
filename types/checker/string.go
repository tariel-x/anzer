package checker

import "github.com/tariel-x/anzer/types"

func checkString(parent, child types.JsonSchema) TypesIdentity {

	if parent.MaxLength != child.MaxLength {
		return TypesNotEqual
	}

	if parent.MinLength != child.MinLength {
		return TypesNotEqual
	}

	if parent.Pattern != child.Pattern {
		return TypesNotEqual
	}

	return TypesEqual
}