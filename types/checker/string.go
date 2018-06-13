package checker

import "github.com/tariel-x/anzer/types"

func checkString(parent, child types.JsonSchema) TypesIdentity {
	var identity TypesIdentity
	identity = TypesEqual

	if parent.MaxLength != nil && child.MaxLength != nil && *parent.MaxLength != *child.MaxLength {
		return TypesNotEqual
	}

	if parent.MaxLength == nil && child.MaxLength != nil {
		identity =  TypesSubtype
	}

	if parent.MinLength != nil && child.MinLength != nil && *parent.MinLength != *child.MinLength {
		return TypesNotEqual
	}

	if parent.MinLength == nil && child.MinLength != nil {
		identity =  TypesSubtype
	}

	if parent.Pattern != nil && child.Pattern != nil && *parent.Pattern != *child.Pattern {
		return TypesNotEqual
	}

	if parent.Pattern == nil && child.Pattern != nil {
		identity =  TypesSubtype
	}

	return identity
}