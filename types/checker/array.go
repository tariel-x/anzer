package checker

import "github.com/tariel-x/anzer/types"

func checkArray(parent, child types.JsonSchema) TypesIdentity {

	if parent.MaxItems != child.MaxItems {
		return TypesNotEqual
	}
	if parent.MinItems != child.MinItems {
		return TypesNotEqual
	}
	if parent.Items != child.Items {
		return TypesNotEqual
	}
	if parent.Contains != child.Contains {
		return TypesNotEqual
	}
	if parent.AdditionalItems != child.AdditionalItems {
		return TypesNotEqual
	}
	if parent.UniqueItems != child.UniqueItems {
		return TypesNotEqual
	}

	return TypesEqual
}
