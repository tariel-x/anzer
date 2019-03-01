package platform

import (
	"errors"
	"github.com/tariel-x/anzer/wsk"
	"io"

	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/wsk/golang"
)

var (
	errUndefinedLanguage = errors.New("undefined language")
	errUndefinedPlatform = errors.New("undefined platform")
)

type CodeGenerator interface {
	GenerateFunc(inT, outT l.T, packagePath string) (string, error)
}

type Builder interface {
	Build(link string, inT l.T, outT l.T) (io.Reader, error)
}

func GetGenerator(runtime string) (CodeGenerator, error) {
	switch runtime {
	case "golang":
		return golang.NewGenerator(), nil
	default:
		return nil, errUndefinedLanguage
	}
}

func GetBuilder(runtime string) (Builder, error) {
	switch runtime {
	case "golang":
		return golang.NewBuilder(), nil
	default:
		return nil, errUndefinedLanguage
	}
}

type Platform interface {
	Update(action io.Reader, name, runtime string) error
	Create(action io.Reader, name, runtime string) error
	Upsert(action io.Reader, name, runtime string) error
}

func GetPlatform(platform string) (Platform, error) {
	switch platform {
	case "golang":
		return wsk.New(), nil
	default:
		return nil, errUndefinedPlatform
	}
}
