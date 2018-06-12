package checker

import "github.com/tariel-x/anzer/types"

func validateObject(parent, child types.JsonSchema) TypesIdentity {
	requiredIdent := validateRequired(parent, child)
	if requiredIdent == TypesNotEqual {
		return TypesNotEqual
	}

	propsIdent := validateProperties(parent, child)
	if propsIdent == TypesNotEqual {
		return TypesNotEqual
	}

	if requiredIdent == TypesEqual && propsIdent == TypesEqual {
		return TypesEqual
	}

	return TypesSubtype
}

// validateRequired checks identity of parent and child required lists.
// Returns 0 if types are equal, 1 if child is subtype of parent and -1 if child is not equal and not subtype of parent.
func validateRequired(parent, child types.JsonSchema) TypesIdentity {
	var identity TypesIdentity
	identity = TypesSubtype

	childMap := map[string]bool{}
	for _, name := range child.Required {
		childMap[name] = true
	}

	for _, name := range parent.Required {
		_, exists := childMap[name]
		if !exists {
			return TypesNotEqual
		}
	}

	if len(parent.Required) == len(child.Required) {
		identity = TypesEqual
	}

	return identity
}

func validateProperties(parent, child types.JsonSchema) TypesIdentity {
	var identity TypesIdentity
	identity = TypesSubtype

	for name, parentProp := range parent.Properties {
		childProp, exists := child.Properties[name]
		if !exists {
			return TypesNotEqual
		}
		if typesIdent := Subtype(parentProp, childProp); typesIdent == TypesNotEqual {
			return TypesNotEqual
		}
	}

	if len(parent.Properties) == len(child.Properties) {
		identity = TypesEqual
	}

	return identity
}