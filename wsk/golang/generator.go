package golang

import (
	"bytes"
	"errors"
	"github.com/tariel-x/anzer/platform/models"
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

func (g Generator) Generate(f l.Runnable) (string, error) {
	packagePath := string(f.GetLink())
	packageElements := strings.Split(packagePath, "/")
	if len(packageElements) == 0 {
		return "", errInvalidPackage
	}

	var result bytes.Buffer

	templateArgs := struct {
		Timestamp   time.Time
		AnzerIn     string
		AnzerOut    string
		PackagePath string
		Package     string
		Either      bool
	}{
		Timestamp:   time.Now(),
		AnzerIn:     genType(f.In(), "AnzerIn"),
		AnzerOut:    genType(f.Out(), "AnzerOut"),
		PackagePath: packagePath,
		Package:     packageElements[len(packageElements)-1],
		Either:      f.IsEither(),
	}

	err := execTemplate.Execute(&result, templateArgs)
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

	templateArgs := struct {
		Timestamp time.Time
		TypeIn    string
		TypeOut   string
		Package   string
	}{
		Timestamp: time.Now(),
		TypeIn:    genType(inT, "TypeIn"),
		TypeOut:   genType(outT, "TypeOut"),
		Package:   packageName,
	}
	err := funcTemplate.Execute(&result, templateArgs)
	return result.String(), err
}

func (g Generator) GenerateDocker(opts *models.BuildOpts) (string, error) {
	var result bytes.Buffer

	templateArgs := struct {
		Debug       bool
		CommitID    string
		PackagePath string
	}{
		Debug:       opts.Debug,
		CommitID:    opts.CommitID,
		PackagePath: opts.F.GetName(),
	}
	err := dockerfileTemplate.Execute(&result, templateArgs)
	return result.String(), err
}

func genType(t l.T, name string) string {
	return j.Type().Id(name).Add(genTypeDef(t)).GoString()
}

func genTypeDef(t l.T) *j.Statement {
	switch tt := t.(type) {
	case l.Container:
		if len(tt.Operands) == 0 {
			return nil
		}
		if l.IsOptional(tt) {
			return pointer(genTypeDef(tt.Operands[0]))
		} else if tt.Type() == l.Type(l.TypeList) {
			return list(genTypeDef(tt.Operands[0]))
		}
		return genTypeDef(tt.Operands[0])

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
		switch {
		case l.IsOptional(tt):
			return pointer(genTypeDef(tt[1]))
		case l.IsEither(tt):
			leftType := pointer(genTypeDef(tt[0]).Tag(map[string]string{"json": "left"}))
			rightType := pointer(genTypeDef(tt[1]).Tag(map[string]string{"json": "right"}))
			left := j.Id(strings.Title("Left")).Add(leftType)
			right := j.Id(strings.Title("Right")).Add(rightType)
			return j.Struct(left, right)
		}
	}
	return j.Interface()
}
