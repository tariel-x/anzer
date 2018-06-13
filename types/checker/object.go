package checker

import (
	"github.com/tariel-x/anzer/types"
)

// checkObject checks identity of parent and child of type `object`
// Returns 0 if types are equal, 1 if child is subtype of parent and -1 if child is not equal and not subtype of parent.
func checkObject(parent, child types.JsonSchema) TypesIdentity {
	requiredIdent := checkObjRequired(parent, child)
	if requiredIdent == TypesNotEqual {
		return TypesNotEqual
	}

	propsIdent := checkObjProperties(parent, child)
	if propsIdent == TypesNotEqual {
		return TypesNotEqual
	}

	if requiredIdent == TypesEqual && propsIdent == TypesEqual {
		return TypesEqual
	}

	return TypesSubtype
}

// checkObjRequired checks identity of parent and child required lists.
// Returns 0 if types are equal, 1 if child is subtype of parent and -1 if child is not equal and not subtype of parent.
func checkObjRequired(parent, child types.JsonSchema) TypesIdentity {
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

func checkObjProperties(parent, child types.JsonSchema) TypesIdentity {
	var subIdentity TypesIdentity
	subIdentity = TypesEqual

	for name, parentProp := range parent.Properties {
		childProp, exists := child.Properties[name]
		if !exists {
			return TypesNotEqual
		}

		typesIdent := Subtype(parentProp, childProp)

		if typesIdent == TypesNotEqual {
			return TypesNotEqual
		}

		if typesIdent == TypesSubtype {
			subIdentity = TypesSubtype
		}
	}

	if len(parent.Properties) == len(child.Properties) && subIdentity == TypesEqual {
		return TypesEqual
	}

	return TypesSubtype
}