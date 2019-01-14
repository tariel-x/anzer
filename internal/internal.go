package internal

type Alias struct {
	Name  string
	Alias []Composable
}

func (a Alias) Definiiton() string {
	return a.Name
}

type Func struct {
	Name    string
	Link    string
	TypeIn  interface{}
	TypeOut interface{}
}

func (f Func) Definiiton() string {
	return f.Name
}

type Application struct {
	Name string
}

func (a Application) Definiiton() string {
	return a.Name
}

type Composable interface {
	Definiiton() string
}
