package env

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type FunctionEnvs map[string]string

type input map[string]FunctionEnvs

func LoadEnvs(reader io.Reader) (input, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	i := input{}
	return i, yaml.Unmarshal(data, &i)
}
