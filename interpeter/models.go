package interpeter

type BaseType struct {
	Name string
	Val  string
	Op   string
	Arg  *BaseType
}

func NewBaseType(name string, val string) BaseType {
	return BaseType{
		Name: name,
		Val:  val,
	}
}

type BaseFunc struct {
	arg string
	ret string
	Def *FuncBody
}

type FuncBody struct {
	Name    string
	Product []FuncBody
	Param   *FuncBody
}

func NewBaseFunc(name string, arg string, ret string) BaseFunc {
	return BaseFunc{
		Def: &FuncBody{
			Name: name,
		},
		arg: arg,
		ret: ret,
	}
}

func (f *BaseFunc) Name() string {
	return f.Def.Name
}

func (f *BaseFunc) Arg() string {
	return f.arg
}

func (f *BaseFunc) Ret() string {
	return f.ret
}
