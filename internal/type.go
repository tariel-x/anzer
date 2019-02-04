package internal

type T interface {
	Equal(to T) bool
	Subtype(of T) bool
	Parent(of T) bool
}

type NothingType struct{}

var (
	Nothing NothingType = NothingType{}
)

func (n NothingType) Equal(to T) bool {
	if _, ok := to.(NothingType); ok {
		return true
	}
	return false
}

func (n NothingType) Subtype(of T) bool {
	return false
}

func (n NothingType) Parent(of T) bool {
	return false
}

type AnyType struct{}

var (
	Any AnyType = AnyType{}
)

func (a AnyType) Equal(to T) bool {
	if _, ok := to.(AnyType); ok {
		return true
	}
	return false
}

func (a AnyType) Subtype(of T) bool {
	return false
}

func (a AnyType) Parent(of T) bool {
	return true
}

type Basic int

const (
	TypeString Basic = iota
	TypeInteger
	TypeBool
)

func (b Basic) Equal(to T) bool {
	switch t := to.(type) {
	case Basic:
		return t == b
	}
	return false
}

func (b Basic) Parent(of T) bool {
	if of == nil {
		return false
	}
	return of.Subtype(b)
}

func (b Basic) Subtype(of T) bool {
	return false
}

type Complex struct {
	Fields map[string]T
}

func (c Complex) Equal(to T) bool {
	switch t := to.(type) {
	case Complex:
		if len(c.Fields) != len(t.Fields) {
			return false
		}
		for n1, f1 := range c.Fields {
			if f2, ok := t.Fields[n1]; ok {
				if !f1.Equal(f2) {
					return false
				}
			} else {
				return false
			}
		}
	default:
		return false
	}
	return true
}

func (c Complex) Parent(of T) bool {
	if of == nil {
		return false
	}
	return of.Subtype(c)
}

func (c Complex) Subtype(of T) bool {
	switch t := of.(type) {
	case Complex:
		for n1, f1 := range c.Fields {
			if f2, ok := t.Fields[n1]; ok {
				if !f1.Parent(f2) && !f1.Equal(f2) {
					return false
				}
			} else {
				return false
			}
		}
	default:
		return false
	}
	return true
}

type ConstructorType int

const (
	MaxLength ConstructorType = iota
	MinLength
	Right
)

type Constructor struct {
	Operand  T
	Type     ConstructorType
	Argument interface{}
}

func Construct(parent T, constructor ConstructorType, argument interface{}) T {
	return Constructor{
		Operand:  parent,
		Type:     constructor,
		Argument: argument,
	}
}

func (c Constructor) Equal(to T) bool {
	switch t := to.(type) {
	case Constructor:
		return c.Operand.Equal(t.Operand) &&
			c.Argument == t.Argument &&
			c.Type == t.Type
	}
	return false
}

func (c Constructor) Parent(of T) bool {
	if of == nil {
		return false
	}
	return of.Subtype(c)
}

func (c Constructor) Subtype(of T) bool {
	switch t := of.(type) {
	case Constructor:
		if c.Equal(t) {
			return false
		}
		if c.Equal(t.Operand) || c.Subtype(t.Operand) {
			return true
		}
	default:
		if c.Operand.Equal(of) {
			return true
		}
	}

	return false
}
