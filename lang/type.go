package lang

/*
+------------------------------+
|  ADT<---------------Basic    |
|   ^ <-                ^      |
|   |   \---            |      |
|   |       \---        |      |
|   |           \---    |      |
|   |               \-  |      |
|Complex<----------Constructor |
+------------------------------+
*/

type T interface {
	Equal(to T) bool
	Subtype(of T) bool
	Parent(of T) bool
}

const (
	TypeString Basic = iota
	TypeInteger
	TypeFloat
	TypeBool
	TypeMaxLength ConstructorType = iota
	TypeMinLength
	TypeRight
	TypeLeft
	TypeList
	TypeOptional
	TypeEither
)

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
	return true
}

func (a AnyType) Subtype(of T) bool {
	return false
}

func (a AnyType) Parent(of T) bool {
	return true
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
	if of == nil {
		return false
	}
	return of.Subtype(b)
}

func (b Basic) Subtype(of T) bool {
	switch of.(type) {
	case Sum:
		return of.Parent(b)
	default:
		return false
	}
}

type Ref string

func (r Ref) Equal(to T) bool {
	return r == to
}

func (r Ref) Parent(of T) bool {
	return false
}

func (r Ref) Subtype(of T) bool {
	return false
}

type Record struct {
	Fields map[string]T
}

func (r Record) Equal(to T) bool {
	switch t := to.(type) {
	case Record:
		if len(r.Fields) != len(t.Fields) {
			return false
		}
		for n1, f1 := range r.Fields {
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

func (r Record) Parent(of T) bool {
	if of == nil {
		return false
	}
	return of.Subtype(r)
}

func (r Record) Subtype(of T) bool {
	switch t := of.(type) {
	case Record:
		for n1, f1 := range r.Fields {
			if f2, ok := t.Fields[n1]; ok {
				if !f1.Parent(f2) && !f1.Equal(f2) {
					return false
				}
			} else {
				return false
			}
		}
	case Sum:
		return of.Parent(r)
	default:
		return false
	}
	return true
}

type Sum []T

func (s Sum) Equal(to T) bool {
	switch t := to.(type) {
	case Sum:
		if len(s) != len(t) {
			return false
		}
		for _, f1 := range s {
			existsEqual := false
			for _, f2 := range t {
				if f1.Equal(f2) {
					existsEqual = true
					break
				}
			}
			if !existsEqual {
				return false
			}
		}
	default:
		return false
	}
	return true
}

func (s Sum) Parent(of T) bool {
	if of == nil {
		return false
	}
	switch oft := of.(type) {
	case Sum:
		for _, off := range oft {
			existsEqual := false
			for _, f := range s {
				if off.Subtype(f) || off.Equal(f) {
					existsEqual = true
					break
				}
			}
			if !existsEqual {
				return false
			}
		}
		return true
	default:
		for _, f := range s {
			if of.Subtype(f) || of.Equal(f) {
				return true
			}
		}
		return false
	}
}

func (s Sum) Subtype(of T) bool {
	switch of.(type) {
	case Sum:
		return of.Parent(s)
	}
	return false
}

func NewSum(types ...T) T {
	return Sum(types)
}

type ConstructorType int

type Constructor struct {
	Operands  []T
	Type      ConstructorType
	Arguments []interface{}
}

func Construct(parents []T, constructor ConstructorType, arguments []interface{}) T {
	return Constructor{
		Operands:  parents,
		Type:      constructor,
		Arguments: arguments,
	}
}

func (c Constructor) Equal(to T) bool {
	switch t := to.(type) {
	case Constructor:
		if len(c.Operands) != len(t.Operands) {
			return false
		}
		for i, op1 := range c.Operands {
			if !op1.Equal(t.Operands[i]) ||
				c.Type != t.Type {
				return false
			}
		}
		if len(c.Arguments) != len(t.Arguments) {
			return false
		}
		for i, arg1 := range c.Arguments {
			if arg1 != t.Arguments[i] {
				return false
			}
		}
		return true

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
		for i, op1 := range c.Operands {
			if !op1.Equal(t.Operands[i]) && !op1.Subtype(t.Operands[i]) {
				return false
			}
		}
	case Basic, Record:
		for _, op := range c.Operands {
			if !op.Equal(of) {
				return false
			}
		}
	default:
		return false
	}

	return true
}

func MaxLength(parent T, length int) T {
	if parent.Subtype(TypeString) || parent.Equal(TypeString) {
		return Construct([]T{parent}, TypeMaxLength, []interface{}{length})
	}
	return nil
}

func MinLength(parent T, length int) T {
	if parent.Subtype(TypeString) || parent.Equal(TypeString) {
		return Construct([]T{parent}, TypeMinLength, []interface{}{length})
	}
	return nil
}

func Right(parent T) T {
	return Construct([]T{parent}, TypeRight, nil)
}

func Left(parent T) T {
	return Construct([]T{parent}, TypeLeft, nil)
}

func List(parent T) T {
	return Construct([]T{parent}, TypeList, nil)
}

func Optional(parent T) T {
	return NewSum(Nothing, parent)
}

func Either(left, right T) T {
	return Construct([]T{left, right}, TypeEither, nil)
}
