package checker

import "github.com/tariel-x/anzer/types"

func checkNumber(parent, child types.JsonSchema) TypesIdentity {

	var identity TypesIdentity
	identity = TypesEqual

	// minimum
	if parent.Minimum != nil && child.Minimum != nil && *parent.Minimum != *child.Minimum {
		return TypesNotEqual
	}

	if parent.Minimum == nil && child.Minimum != nil {
		identity = TypesSubtype
	}

	// exclusiveMinimum
	if parent.ExclusiveMinimum != nil && child.ExclusiveMinimum != nil && *parent.ExclusiveMinimum != *child.ExclusiveMinimum {
		return TypesNotEqual
	}

	if parent.ExclusiveMinimum == nil && child.ExclusiveMinimum != nil {
		identity = TypesSubtype
	}

	// maximum
	if parent.Maximum != nil && child.Maximum != nil && *parent.Maximum != *child.Maximum {
		return TypesNotEqual
	}

	if parent.Maximum == nil && child.Maximum != nil {
		identity = TypesSubtype
	}

	// exclusiveMaximum
	if parent.ExclusiveMaximum != nil && child.ExclusiveMaximum != nil && *parent.ExclusiveMaximum != *child.ExclusiveMaximum {
		return TypesNotEqual
	}

	if parent.ExclusiveMaximum == nil && child.ExclusiveMaximum != nil {
		identity = TypesSubtype
	}

	// multipleOf
	if parent.MultipleOf != nil && child.MultipleOf != nil && *parent.MultipleOf != *child.MultipleOf {
		return TypesNotEqual
	}

	if parent.MultipleOf == nil && child.MultipleOf != nil {
		identity = TypesSubtype
	}

	return identity
}
