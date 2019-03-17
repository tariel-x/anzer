package platform

import (
	"errors"
	"io"

	l "github.com/tariel-x/anzer/lang"
	"github.com/tariel-x/anzer/wsk"
	"github.com/tariel-x/anzer/wsk/golang"
)

var (
	errUndefinedLanguage = errors.New("undefined language")
	errUndefinedPlatform = errors.New("undefined platform")
)

type CodeGenerator interface {
	GenerateFunc(inT, outT l.T, link l.FunctionLink) (string, error)
}

func GetGenerator(runtime string) (CodeGenerator, error) {
	switch runtime {
	case "go":
		return golang.NewGenerator(), nil
	default:
		return nil, errUndefinedLanguage
	}
}

var (
	platforms = map[string]Platform{}
)

type Platform interface {
	Update(action io.Reader, name, runtime string) error
	Create(action io.Reader, name, runtime string) error
	Upsert(action io.Reader, name, runtime string) error
	Link(invoke string, names []string) error
}

func GetPlatform(name string) (Platform, error) {
	if p, ok := platforms[name]; ok {
		return p, nil
	}
	var err error

	switch name {
	case "wsk":
		platforms[name], err = wsk.New()
		return platforms[name], err
	default:
		return nil, errUndefinedPlatform
	}
}
