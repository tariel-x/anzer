package checker

import "github.com/tariel-x/anzer/types"

func checkNumber(parent, child types.JsonSchema) TypesIdentity {

	if parent.Minimum != child.Minimum {
		return TypesNotEqual
	}
	if parent.ExclusiveMinimum != child.ExclusiveMinimum {
		return TypesNotEqual
	}
	if parent.Maximum != child.Maximum {
		return TypesNotEqual
	}
	if parent.ExclusiveMaximum != child.ExclusiveMaximum {
		return TypesNotEqual
	}
	if parent.MultipleOf != child.MultipleOf {
		return TypesNotEqual
	}

	return TypesEqual
}
