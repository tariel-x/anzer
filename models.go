package main

type Type interface {
	Name() string
	Val() string
}

type Func interface {
	Name() string
	Arg() Type
	Ret() Type
}
