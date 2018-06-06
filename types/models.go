package types

type JsonSchema struct {
	JSTypeString
	JSTypeInt
	JSTypeObj
	JSTypeArr
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

type JSTypeObj struct {
	Required             []string              `json:"required"`
	Properties           map[string]JsonSchema `json:"properties"`
	AdditionalProperties *bool                 `json:"additionalProperties"`
}

type JSTypeArr struct {
	MaxItems *int        `json:"maxItems"`
	MinItems *int        `json:"minItems"`
	Items    *JsonSchema `json:"items"`
}
