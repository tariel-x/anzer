package env

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type FunctionParams map[string]string

type FunctionsParams map[string]FunctionParams

func LoadParams(reader io.Reader) (FunctionsParams, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	i := FunctionsParams{}
	return i, yaml.Unmarshal(data, &i)
}
