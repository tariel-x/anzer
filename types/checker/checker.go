package checker

import (
	"github.com/tariel-x/anzer/types"
)

type TypesIdentity int8

const (
	TypesEqual    TypesIdentity = 0
	TypesNotEqual TypesIdentity = -1
	TypesSubtype  TypesIdentity = 1
)

func Subtype(parent, child types.JsonSchema) TypesIdentity {
	if eq := checkTypeIdentity(parent, child); eq != TypesNotEqual {
		return TypesNotEqual
	}

	switch *parent.Type {
	case types.Object:
		return validateObject(parent, child)
	}

	return TypesNotEqual
}

func checkTypeIdentity(schema1, schema2 types.JsonSchema) TypesIdentity {
	if schema1.Type == nil || schema2.Type == nil {
		return TypesNotEqual
	}

	if schema1.Type != schema2.Type {
		return TypesNotEqual
	}

	return TypesEqual
}
