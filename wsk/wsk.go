package wsk

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
)

type Wsk struct {
	client *whisk.Client
}

func New() (Wsk, error) {
	client, err := whisk.NewClient(http.DefaultClient, nil)
	return Wsk{
		client: client,
	}, err
}

func (w Wsk) List() error {
	actions, _, err := w.client.Actions.List("", nil)
	if err != nil {
		return err
	}
	for _, action := range actions {
		fmt.Println(action.Name)
	}
	return nil
}

func (w Wsk) Update(action io.Reader, name, runtime string) error {
	return nil
}

func (w Wsk) Create(action io.Reader, name, runtime string) error {
	if len(strings.Split(runtime, ":")) == 1 {
		runtime = runtime + ":default"
	}

	exec, err := w.makeExec(action, runtime)
	if err != nil {
		return err
	}
	publish := true
	wskaction := whisk.Action{
		Exec:      exec,
		Name:      name,
		Namespace: "",
		Publish:   &publish,
	}
	wskaction.Annotations.AddOrReplace(&whisk.KeyValue{
		Key:   "web-export",
		Value: true,
	})
	readyAction, _, err := w.client.Actions.Insert(&wskaction, true)
	fmt.Println(readyAction)
	return err
}

func (w Wsk) Upsert(action io.Reader, name, runtime string) error {
	return nil
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