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
	arg  string
	ret  string
}

func NewBaseFunc(name string, arg string, ret string) BaseFunc {
	return BaseFunc{
		name: name,
		arg:  arg,
		ret:  ret,
	}
}

func (f *BaseFunc) Name() string {
	return f.name
}

func (f *BaseFunc) Arg() string {
	return f.arg
}

func (f *BaseFunc) Ret() string {
	return f.ret
}
