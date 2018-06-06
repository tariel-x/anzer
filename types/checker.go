package types

import (
	"encoding/json"
	"fmt"

	"github.com/tariel-x/anzer/interpeter"
)

type Checker struct {
	RawTypes map[string]interpeter.BaseType
	Types    map[string]JsonSchema
}

func NewChecker(types map[string]interpeter.BaseType) Checker {
	return Checker{
		RawTypes: types,
		Types:    map[string]JsonSchema{},
	}
}

func (c *Checker) Simplify() {
	err := c.simplifyAll()
	if err != nil {
		panic(err)
	}
	for name, def := range c.Types {
		fmt.Printf("Type %s %v\n", name, def)
	}
}

func (c *Checker) simplifySimple(raw json.RawMessage) (*JsonSchema, error) {
	schema := JsonSchema{}
	err := json.Unmarshal(raw, &schema)
	if err != nil {
		return nil, err
	}
	return &schema, nil
}

func (c *Checker) simplifyAll() error {
	types := map[string]JsonSchema{}
	for name, _ := range c.RawTypes {
		schema, err := c.getType(name)
		if err != nil {
			return err
		}
		if schema == nil {
			continue
		}
		types[name] = *schema
	}
	c.Types = types
	return nil
}

func (c *Checker) getType(name string) (*JsonSchema, error) {
	schema, exists := c.Types[name]
	if exists {
		return &schema, nil
	}
	return c.simplifyType(name)
}

func (c *Checker) simplifyType(name string) (*JsonSchema, error) {
	def, exists := c.RawTypes[name]
	if !exists {
		return nil, fmt.Errorf("No such type %s", name)
	}
	if def.Type != nil {
		schema, err := c.simplifySimple(*def.Type)
		if err != nil {
			return nil, err
		}
		return schema, nil
	}
	return nil, nil
}