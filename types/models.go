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

	Required             []string              `json:"required"`
	Properties           map[string]JsonSchema `json:"properties"`
	AdditionalProperties *bool                 `json:"additionalProperties"`

	Title       *string               `json:"title"`
	Type        *string               `json:"type"`
	Enum        []string              `json:"enum"`
	Definitions map[string]JsonSchema `json:"definitions"`
}

type JSTypeString struct {
	MaxLength *int    `json:"maxLength"`
	MinLength *int    `json:"minLength"`
	Pattern   *string `json:"pattern"`
}

type JSTypeInt struct {
	Minimum *int `json:"minimum"`
	Maximum *int `json:"maximum"`
}

type JSTypeArr struct {
	MaxItems *int        `json:"maxItems"`
	MinItems *int        `json:"minItems"`
	Items    *JsonSchema `json:"items"`
}
