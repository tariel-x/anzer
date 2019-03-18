package wsk

import (
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/tariel-x/anzer/platform/models"
)

const (
	Sequence  = "sequence"
	Namespace = "guest"
)

type Wsk struct {
	client    *whisk.Client
	namespace string
}

func New() (Wsk, error) {
	client, err := whisk.NewClient(http.DefaultClient, nil)
	return Wsk{
		client:    client,
		namespace: Namespace,
	}, err
}

func (w Wsk) List() ([]models.PublishedFunction, error) {
	actions, _, err := w.client.Actions.List("", nil)
	if err != nil {
		return nil, err
	}
	published := make([]models.PublishedFunction, 0, len(actions))
	for _, action := range actions {
		published = append(published, models.PublishedFunction{
			Name: fmt.Sprintf("%s/%s", action.Namespace, action.Name),
		})
	}
	return published, nil
}

func (w Wsk) Update(action io.Reader, name, runtime string) (models.PublishedFunction, error) {
	return w.Create(action, name, runtime)
}

func (w Wsk) Create(action io.Reader, name, runtime string) (models.PublishedFunction, error) {
	if len(strings.Split(runtime, ":")) == 1 {
		runtime = runtime + ":default"
	}

	exec, err := w.makeExec(action, runtime)
	if err != nil {
		return models.PublishedFunction{}, err
	}
	publish := true
	wskaction := whisk.Action{
		Exec:      exec,
		Name:      name,
		Namespace: w.namespace,
		Publish:   &publish,
		Annotations: whisk.KeyValueArr{
			whisk.KeyValue{
				Key:   "web-export",
				Value: true,
			},
		},
	}
	readyAction, _, err := w.client.Actions.Insert(&wskaction, true)
	return models.PublishedFunction{
		Name: fmt.Sprintf("%s/%s", readyAction.Namespace, readyAction.Name),
	}, err
}

func (w Wsk) Upsert(action io.Reader, name, runtime string) (models.PublishedFunction, error) {
	return w.Create(action, name, runtime)
}

func (w Wsk) makeExec(action io.Reader, runtime string) (*whisk.Exec, error) {
	exec := new(whisk.Exec)
	var err error
	exec.Kind = runtime

	dat, err := ioutil.ReadAll(action)
	if err != nil {
		return nil, err
	}
	code := base64.StdEncoding.EncodeToString([]byte(dat))
	exec.Code = &code

	return exec, err
}

func (w Wsk) Link(invoke string, names []string) (models.PublishedFunction, error) {
	publish := true
	wskaction := whisk.Action{
		Exec: &whisk.Exec{
			Kind:       Sequence,
			Components: names,
		},
		Annotations: whisk.KeyValueArr{
			whisk.KeyValue{
				Key:   "web-export",
				Value: true,
			},
			whisk.KeyValue{
				Key:   "web",
				Value: true,
			},
		},
		Name:      invoke + "_sequence",
		Namespace: w.namespace,
		Publish:   &publish,
	}
	readyAction, _, err := w.client.Actions.Insert(&wskaction, true)
	if err != nil {
		return models.PublishedFunction{}, err
	}
	return models.PublishedFunction{
		Name: fmt.Sprintf("%s/%s", readyAction.Namespace, readyAction.Name),
	}, nil
}

func (w Wsk) extractNamespace(name string) (string, string, error) {
	parts := strings.Split(name, "/")
	if len(parts) < 2 {
		return "", "", errors.New("Wrong function uri")
	}
	return parts[0], parts[1], nil
}

func (w Wsk) Init(arg string) error {
	return nil
}
