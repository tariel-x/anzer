package golang

import (
	"bytes"
	"errors"
	"sort"
	"strings"
	"time"

	j "github.com/dave/jennifer/jen"
	l "github.com/tariel-x/anzer/lang"
)

var (
	errInvalidPackage = errors.New("invalid package")
)

type Generator struct{}

func NewGenerator() Generator {
	return Generator{}
}

func (g Generator) Generate(inT, outT l.T, link l.FunctionLink) (string, error) {
	packagePath := string(link)
	packageElements := strings.Split(packagePath, "/")
	if len(packageElements) == 0 {
		return "", errInvalidPackage
	}

	var result bytes.Buffer
	err := execTemplate.Execute(&result, struct {
		Timestamp   time.Time
		AnzerIn     string
		AnzerOut    string
		PackagePath string
		Package     string
	}{
		Timestamp:   time.Now(),
		AnzerIn:     genType(inT, "AnzerIn"),
		AnzerOut:    genType(outT, "AnzerOut"),
		PackagePath: packagePath,
		Package:     packageElements[len(packageElements)-1],
	})
	return result.String(), err
}

func (g Generator) GenerateFunc(inT, outT l.T, link l.FunctionLink) (string, error) {
	packagePath := string(link)
	packageElements := strings.Split(packagePath, "/")
	if len(packageElements) == 0 {
		return "", errInvalidPackage
	}
	packageName := packageElements[len(packageElements)-1]

	var result bytes.Buffer
	err := funcTemplate.Execute(&result, struct {
		Timestamp time.Time
		TypeIn    string
		TypeOut   string
		Package   string
	}{
		Timestamp: time.Now(),
		TypeIn:    genType(inT, "TypeIn"),
		TypeOut:   genType(outT, "TypeOut"),
		Package:   packageName,
	})
	return result.String(), err
}

func (g Generator) GenerateDocker(debug bool) string {
	if debug {
		return dockerfileDebug
	}
	return dockerfile
}

func (g Generator) GenerateMakefile() string {
	return makefile
}

func genType(t l.T, name string) string {
	return j.Type().Id(name).Add(genTypeDef(t)).GoString()
}

func genTypeDef(t l.T) *j.Statement {
	switch tt := t.(type) {
	case l.Constructor:
		if len(tt.Operands) == 0 {
			return nil
		}
		switch tt.Type {
		case l.TypeList:
			return list(genTypeDef(tt.Operands[0]))
		case l.TypeOptional:
			return pointer(genTypeDef(tt.Operands[0]))
		default:
			return genTypeDef(tt.Operands[0])
		}
	case l.Basic:
		return basic(tt)
	case l.Record:
		return complex(tt)
	case l.AnyType:
		return j.Interface()
	case l.Sum:
		return sum(tt)
	default:
		return j.Interface()
	}
}

func basic(tt l.Basic) *j.Statement {
	switch tt {
	case l.TypeString:
		return j.String()
	case l.TypeInteger:
		return j.Int()
	case l.TypeFloat:
		return j.Float64()
	case l.TypeBool:
		return j.Bool()
	default:
		return j.Interface()
	}
}

func complex(tt l.Record) *j.Statement {
	keys := make([]string, 0, len(tt.Fields))
	for key := range tt.Fields {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	gofields := make([]j.Code, 0, len(tt.Fields))
	for _, key := range keys {
		f := tt.Fields[key]
		goftype := genTypeDef(f).Tag(map[string]string{"json": key})
		gof := j.Id(strings.Title(key)).Add(goftype)
		gofields = append(gofields, gof)
	}
	return j.Struct(gofields...)
}

func pointer(st *j.Statement) *j.Statement {
	return j.Op("*").Add(st)
}

func list(st *j.Statement) *j.Statement {
	return j.Index().Add(st)
}

// sum is hack to properly convert Optional Integer into *int
func sum(tt l.Sum) *j.Statement {
	switch len(tt) {
	case 0:
		return j.Interface()
	case 1:
		return genTypeDef(tt[0])
	default:
		if tt[0] == l.Nothing {
			return pointer(genTypeDef(tt[1]))
		}
	}
	return j.Interface()
}
