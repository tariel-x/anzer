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
		stringDef, err := json.Marshal(def)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Type %s %s\n\n", name, string(stringDef))
	}
}

func (c *Checker) simplifyAll() error {
	for name, _ := range c.RawTypes {
		schema, err := c.getType(name)
		if err != nil {
			return err
		}
		if schema == nil {
			continue
		}
		c.Types[name] = *schema
	}
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
	if def.Type == nil && def.Args != nil && def.Operand != nil {
		return c.simplifyComplexType(def)
	}
	return nil, fmt.Errorf("Can not simplify type %s", name)
}

func (c *Checker) simplifySimple(raw json.RawMessage) (*JsonSchema, error) {
	schema := JsonSchema{}
	err := json.Unmarshal(raw, &schema)
	if err != nil {
		return nil, err
	}
	return &schema, nil
}

func (c *Checker) simplifyComplexType(def interpeter.BaseType) (*JsonSchema, error) {
	schema := c.makeProdBootstrap()
	required := []string{}
	for _, childName := range def.Args {
		childSchema, err := c.getType(childName)
		if err != nil {
			return nil, err
		}
		schema.Properties[childName] = *childSchema
		required = append(required, childName)
	}
	schema.Required = required
	return &schema, nil
}

func (c *Checker) makeProdBootstrap() JsonSchema {
	return JsonSchema{
		Type: string2point(TypeObject),
		JSTypeObj: JSTypeObj{
			AdditionalProperties: bool2point(true),
			Properties:           map[string]JsonSchema{},
		},
	}
}

func string2point(x string) *string {
	return &x
}

func bool2point(x bool) *bool {
	return &x
}
