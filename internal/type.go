package internal

type T interface {
	Equal(to T) bool
	Subtype(of T) bool
	Parent(of T) bool
}

type Basic int

func (b Basic) Equal(to T) bool {
	switch t := to.(type) {
	case Basic:
		return t == b
	}
	return false
}

func (b Basic) Parent(of T) bool {
	return false
}

func (b Basic) Subtype(of T) bool {
	return of.Parent(b)
}

type Constructor struct {
	Operand  T
	Type     string
	Argument interface{}
}

func (c Constructor) Equal(to T) bool {
	switch t := to.(type) {
	case Constructor:
		return c.Operand.Equal(c) &&
			c.Argument == t.Argument &&
			c.Type == t.Type
	}
	return false
}

func (c Constructor) Parent(of T) bool {
	switch t := of.(type) {
	case Constructor:
		if c.Equal(t) {
			return false
		}
	}
	if c.Operand.Equal(of) {
		return true
	}
	return false
}

func (c Constructor) Subtype(of T) bool {
	return of.Parent(c)
}

type Complex struct {
	Fields map[string]T
}
