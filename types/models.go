package types

type TypeName string

const (
	Object       TypeName = "object"
	Number       TypeName = "number"
	Array        TypeName = "array"
	String       TypeName = "string"
	Unrecognized TypeName = "unrecognized"
)

type Types map[string]JsonSchema

type JsonSchema struct {
	JSTypeString
	JSTypeInt
	JSTypeArr
	JSTypeObj

	Title       *string               `json:"title,omitempty"`
	Type        *TypeName             `json:"type,omitempty"`
	Enum        []string              `json:"enum,omitempty"`
	Definitions map[string]JsonSchema `json:"definitions,omitempty"`
}

type JSTypeObj struct {
	Required   []string              `json:"required,omitempty"`
	Properties map[string]JsonSchema `json:"properties,omitempty"`
}

type JSTypeString struct {
	MaxLength *int    `json:"maxLength,omitempty"`
	MinLength *int    `json:"minLength,omitempty"`
	Pattern   *string `json:"pattern,omitempty"`
}

type JSTypeInt struct {
	Minimum          *int `json:"minimum,omitempty"`
	ExclusiveMinimum *int `json:"exclusiveMinimum,omitempty"`
	Maximum          *int `json:"maximum,omitempty"`
	ExclusiveMaximum *int `json:"exclusiveMaximum,omitempty"`
	MultipleOf       *int `json:"multipleOf,omitempty"`
}

type JSTypeArr struct {
	MaxItems        *int        `json:"maxItems,omitempty"`
	MinItems        *int        `json:"minItems,omitempty"`
	Items           *JsonSchema `json:"items,omitempty"`
	Contains        *JsonSchema `json:"contains,omitempty"`
	AdditionalItems *JsonSchema `json:"additionalItems,omitempty"`
	UniqueItems     *bool       `json:"uniqueItems,omitempty"`
}
