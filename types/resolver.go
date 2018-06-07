package types

import (
	"encoding/json"
	"fmt"

	"github.com/tariel-x/anzer/listener"
)

type Resolver struct {
	RawTypes listener.Types
	Types    Types
}

func NewResolver(types listener.Types) Resolver {
	return Resolver{
		RawTypes: types,
		Types:    Types{},
	}
}

func (c *Resolver) GetTypes() Types {
	return c.Types
}

func (c *Resolver) Resolve() error {
	err := c.resolveAll()
	if err != nil {
		return err
	}
	return nil
}

func (c *Resolver) resolveAll() error {
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

func (c *Resolver) getType(name string) (*JsonSchema, error) {
	schema, exists := c.Types[name]
	if exists {
		return &schema, nil
	}
	return c.simplifyType(name)
}

func (c *Resolver) simplifyType(name string) (*JsonSchema, error) {
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

func (c *Resolver) simplifySimple(raw json.RawMessage) (*JsonSchema, error) {
	schema := JsonSchema{}
	err := json.Unmarshal(raw, &schema)
	if err != nil {
		return nil, err
	}
	return &schema, nil
}

func (c *Resolver) simplifyComplexType(def listener.BaseType) (*JsonSchema, error) {
	if len(def.Args) == 0 {
		return c.getType(def.Args[0])
	}

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

func (c *Resolver) makeProdBootstrap() JsonSchema {
	return JsonSchema{
		Type: string2point(TypeObject),
		JSTypeObj: JSTypeObj{
			AdditionalProperties: bool2point(false),
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
