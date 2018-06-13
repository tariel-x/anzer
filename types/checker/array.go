package checker

import "github.com/tariel-x/anzer/types"

func checkArray(parent, child types.JsonSchema) TypesIdentity {

	var identity TypesIdentity
	identity = TypesEqual

	// maxItems
	if parent.MaxItems != nil && child.MaxItems != nil && *parent.MaxItems != *child.MaxItems {
		return TypesNotEqual
	}

	if parent.MaxItems == nil && child.MaxItems != nil {
		identity = TypesSubtype
	}

	// minItems
	if parent.MinItems != nil && child.MinItems != nil && *parent.MinItems != *child.MinItems {
		return TypesNotEqual
	}

	if parent.MinItems == nil && child.MinItems != nil {
		identity = TypesSubtype
	}

	// items
	if parent.Items != nil && child.Items != nil{
		switch Subtype(*parent.Items, *child.Items) {
		case TypesNotEqual:
			return TypesNotEqual
		case TypesSubtype:
			identity = TypesSubtype
		}
	}

	if parent.Items == nil && child.Items != nil {
		identity = TypesSubtype
	}

	// contains
	if parent.Contains != nil && child.Contains != nil{
		switch Subtype(*parent.Contains, *child.Contains) {
		case TypesNotEqual:
			return TypesNotEqual
		case TypesSubtype:
			identity = TypesSubtype
		}
	}

	if parent.Contains == nil && child.Contains != nil {
		identity = TypesSubtype
	}

	// additionalItems
	if parent.AdditionalItems != nil && child.AdditionalItems != nil{
		switch Subtype(*parent.AdditionalItems, *child.AdditionalItems) {
		case TypesNotEqual:
			return TypesNotEqual
		case TypesSubtype:
			identity = TypesSubtype
		}
	}

	if parent.AdditionalItems == nil && child.AdditionalItems != nil {
		identity = TypesSubtype
	}

	// uniqueItems
	if parent.UniqueItems != nil && child.UniqueItems != nil && *parent.UniqueItems != *child.UniqueItems {
		return TypesNotEqual
	}

	if parent.UniqueItems == nil && child.UniqueItems != nil {
		identity = TypesSubtype
	}

	return identity
}
