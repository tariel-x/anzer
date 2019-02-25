package platform

import (
	"errors"

	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/wsk/golang"
)

type CodeGenerator interface {
	Generate(inT, outT l.T, packagePath string) (string, error)
	GenerateFunc(inT, outT l.T, packagePath string) (string, error)
	GenerateDocker() string
}

var (
	errUndefinedLanguage = errors.New("undefined language")
)

func GetGenerator(platform string) (CodeGenerator, error) {
	switch platform {
	case "golang":
		return golang.NewGenerator(), nil
	default:
		return nil, errUndefinedLanguage
	}
}
