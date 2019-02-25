package platform

import (
	"errors"

	"github.com/tariel-x/anzer/go/generator"
	l "github.com/tariel-x/anzer/lang"
)

type CodeGenerator interface {
	Generate(inT, outT l.T, packagePath string) (string, error)
	GenerateFunc(inT, outT l.T, packagePath string) (string, error)
	GenerateDocker() string
}

var (
	errUndefinedLanguage = errors.New("undefined language")
)

func GetGenerator(lang string) (CodeGenerator, error) {
	switch lang {
	case "go":
		return generator.New(), nil
	default:
		return nil, errUndefinedLanguage
	}
}
