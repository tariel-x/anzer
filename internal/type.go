package internal

type T interface {
	Equal(to T) bool
	Subtype(of T) bool
	Parent(of T) bool
}

/*
Alternatives to type Basic string:
- type String string
- type String Basic
*/

type Basic string

func (b Basic) Equal(to T) bool {
	return false
}

type Constructor struct {
	Operand  T
	Type     string
	Argument interface{}
}
