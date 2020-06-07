package platform

import (
	"errors"

	"github.com/tariel-x/anzer/pkg/platform/models"

	"github.com/tariel-x/anzer/pkg/drivers/wsk"
	"github.com/tariel-x/anzer/pkg/drivers/wsk/golang"
	l "github.com/tariel-x/anzer/pkg/lang"
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
	Update(details models.PublicationDetails) (models.PublishedFunction, error)
	Create(details models.PublicationDetails) (models.PublishedFunction, error)
	Upsert(details models.PublicationDetails) (models.PublishedFunction, error)
	Link(pkg, invoke string, funcs []models.PublishedFunction) (models.PublishedFunction, error)
	Init(args map[string]string) error
	Connect() error
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
