package interpeter

type BaseType struct {
	name string
	val  string
}

func (t *BaseType) Name() string {
	return t.name
}

func (t *BaseType) Val() string {
	return t.val
}

func NewBaseType(name string, val string) BaseType {
	return BaseType{
		name: name,
		val:  val,
	}
}

type BaseFunc struct {
	name string
	arg  BaseType
	ret  BaseType
}

func NewBaseFunc(name string, arg BaseType, ret BaseType) BaseFunc {
	return BaseFunc{
		name: name,
		arg:  arg,
		ret:  ret,
	}
}

func (f *BaseFunc) Name() string {
	return f.name
}

func (f *BaseFunc) Arg() BaseType {
	return f.arg
}

func (f *BaseFunc) Ret() BaseType {
	return f.ret
}
