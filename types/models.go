package types

const (
	TypeObject = "object"
	TypeNumber = "number"
	TypeArray  = "array"
)

type JsonSchema struct {
	JSTypeString
	JSTypeInt
	JSTypeArr
	JSTypeObj

	Title       *string               `json:"title,omitempty"`
	Type        *string               `json:"type,omitempty"`
	Enum        []string              `json:"enum,omitempty"`
	Definitions map[string]JsonSchema `json:"definitions,omitempty"`
}

type JSTypeObj struct {
	Required             []string              `json:"required,omitempty"`
	Properties           map[string]JsonSchema `json:"properties,omitempty"`
	AdditionalProperties *bool                 `json:"additionalProperties,omitempty"`
}

type JSTypeString struct {
	MaxLength *int    `json:"maxLength,omitempty"`
	MinLength *int    `json:"minLength,omitempty"`
	Pattern   *string `json:"pattern,omitempty"`
}

type JSTypeInt struct {
	Minimum *int `json:"minimum,omitempty"`
	Maximum *int `json:"maximum,omitempty"`
}

type JSTypeArr struct {
	MaxItems *int        `json:"maxItems,omitempty"`
	MinItems *int        `json:"minItems,omitempty"`
	Items    *JsonSchema `json:"items,omitempty"`
}
