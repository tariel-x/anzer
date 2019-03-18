package wsk

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/pkg/errors"
	"github.com/tariel-x/anzer/platform/models"
)

const (
	Sequence            = "sequence"
	Namespace           = "guest"
	CfgFile             = ".wskprops"
	CfgDefaultNamespace = "_"
	CfgNamespace        = "NAMESPACE"
	CfgAuth             = "AUTH"
	CfgApiGWAccessToken = "APIGW_ACCESS_TOKEN"
	CfgApihost          = "APIHOST"
)

type Wsk struct {
	client    *whisk.Client
	namespace string
}

func New() (*Wsk, error) {
	return &Wsk{
		namespace: Namespace,
	}, nil
}

func (w *Wsk) Connect() error {
	var err error
	w.client, err = whisk.NewClient(http.DefaultClient, nil)
	return err
}

func (w *Wsk) List() ([]models.PublishedFunction, error) {
	actions, _, err := w.client.Actions.List("", nil)
	if err != nil {
		return nil, err
	}
	published := make([]models.PublishedFunction, 0, len(actions))
	for _, action := range actions {
		published = append(published, models.PublishedFunction{
			Name: fmt.Sprintf("/%s/%s", action.Namespace, action.Name),
		})
	}
	return published, nil
}

func (w *Wsk) Update(action io.Reader, name, runtime string) (models.PublishedFunction, error) {
	return w.Create(action, name, runtime)
}

func (w *Wsk) Create(action io.Reader, name, runtime string) (models.PublishedFunction, error) {
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
		Name: fmt.Sprintf("/%s/%s", readyAction.Namespace, readyAction.Name),
	}, err
}

func (w *Wsk) Upsert(action io.Reader, name, runtime string) (models.PublishedFunction, error) {
	return w.Create(action, name, runtime)
}

func (w *Wsk) makeExec(action io.Reader, runtime string) (*whisk.Exec, error) {
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

func (w *Wsk) Link(invoke string, names []string) (models.PublishedFunction, error) {
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

	publishedFunction := models.PublishedFunction{
		Name: fmt.Sprintf("/%s/%s", readyAction.Namespace, readyAction.Name),
	}

	apiCreateReq := &whisk.ApiCreateRequest{
		ApiDoc: &whisk.Api{
			Namespace:       w.namespace,
			ApiName:         invoke + "_api",
			GatewayBasePath: "/anzer",
			GatewayRelPath:  "/" + invoke,
			GatewayMethod:   http.MethodPost,
		},
	}
	apiCreateOpts := &whisk.ApiCreateRequestOptions{
		ActionName:  publishedFunction.Name,
		ApiBasePath: apiCreateReq.ApiDoc.GatewayBasePath,
		ApiRelPath:  apiCreateReq.ApiDoc.GatewayRelPath,
		ApiVerb:     apiCreateReq.ApiDoc.GatewayMethod,
		ApiName:     invoke + "_api",
	}
	apiCreateResp, _, err := w.client.Apis.Insert(apiCreateReq, apiCreateOpts, true)
	if err != nil {
		return models.PublishedFunction{}, err
	}
	publishedFunction.URL = fmt.Sprintf("%s/%s", apiCreateResp.BaseUrl, "")

	return publishedFunction, nil
}

func (w *Wsk) extractNamespace(name string) (string, string, error) {
	parts := strings.Split(name, "/")
	if len(parts) < 2 {
		return "", "", errors.New("Wrong function uri")
	}
	return parts[0], parts[1], nil
}

func (w *Wsk) Init(args map[string]string) error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	f, err := os.Create(fmt.Sprintf("%s/%s", homedir, CfgFile))
	if err != nil {
		return err
	}
	defer f.Close()
	auth, ok := args["auth"]
	if !ok {
		return errors.New("no auth")
	}
	apihost, ok := args["apihost"]
	if !ok {
		return errors.New("no apihost")
	}
	namespace, ok := args["namespace"]
	if !ok {
		namespace = CfgDefaultNamespace
	}

	cfg := fmt.Sprintf(
		"%s=%s\n%s=%s\n%s=%s\n%s=%s\n",
		CfgNamespace, namespace,
		CfgAuth, auth,
		CfgApiGWAccessToken, auth,
		CfgApihost, apihost,
	)

	if _, err := f.WriteString(cfg); err != nil {
		return err
	}

	if err := w.Connect(); err != nil {
		return err
	}
	if _, err := w.List(); err != nil {
		return err
	}

	return nil
}
