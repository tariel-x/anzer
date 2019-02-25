package golang

import (
	"text/template"
)

var dockerfile = `
FROM golang:latest
`

var funcTemplate = template.Must(template.New("").Parse(`
// This file was generated by robots for you at {{ .Timestamp }}
package {{ .Package }}

{{ .TypeIn}}

{{ .TypeOut}}

func handle(input TypeIn) TypeOut {
	return TypeOut{}
}
`))

var execTemplate = template.Must(template.New("").Parse(`
// This file was generated by robots for you at {{ .Timestamp }}
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"{{ .PackagePath }}"
)

type whiskInput struct {
	Value json.RawMessage ` + "`" + `json:"value"` + "`" + `
}

type whiskOutput struct {
	Value json.RawMessage ` + "`" + `json:"value"` + "`" + `
}

type rawInput map[string]interface{}

func main() {
	// debucging
	var debug = os.Getenv("OW_DEBUG") != ""

	if debug {
		filename := os.Getenv("OW_DEBUG")
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(f)
			defer f.Close()
		}
		log.Printf("ACTION ENV: %v", os.Environ())
	}

	// input
	out := os.NewFile(3, "pipe")
	defer out.Close()
	reader := bufio.NewReader(os.Stdin)

	// read-eval-print loop
	if debug {
		log.Println("started")
	}
	for {
		// read one line
		inbuf, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		if debug {
			log.Printf(">>>'%s'>>>", inbuf)
		}

		// parse one line
		raw := rawInput{}
		if err := json.Unmarshal(inbuf, &raw); err != nil {
			printError(out, err)
			continue
		}
		input := whiskInput{}
		if err := json.Unmarshal(inbuf, &input); err != nil {
			printError(out, err)
			continue
		}
		if debug {
			log.Printf("%v\n", raw)
		}

		setEnvironment(raw)

		// process the request
		result, err := callHandler(input)
		if err != nil {
			printError(out, err)
			continue
		}

		// encode the answer
		output, err := json.Marshal(&result)
		if err != nil {
			printError(out, err)
			continue
		}
		output = bytes.Replace(output, []byte("\n"), []byte(""), -1)
		if debug {
			log.Printf("'<<<%s'<<<", output)
		}
		fmt.Fprintf(out, "%s\n", output)
	}
}

func printError(out io.Writer, err error) {
	log.Println(err.Error())
	fmt.Fprintf(out, "{ error: %q}\n", err.Error())
}

func setEnvironment(raw map[string]interface{}) {
	for k, v := range raw {
		if k == "value" {
			continue
		}
		if s, ok := v.(string); ok {
			os.Setenv("__OW_"+strings.ToUpper(k), s)
		}
	}
}

func callHandler(input whiskInput) (whiskOutput, error) {
	anzHandlerInput := AnzerIn{}
	if err := json.Unmarshal(input.Value, &anzHandlerInput); err != nil {
		return whiskOutput{}, err
	}
	anzHandlerOutput := {{ .Package }}.Handle(TypeIn(anzHandlerInput))
	handlerOutput := AnzerOut(anzHandlerOutput)
	output, err := json.Marshal(handlerOutput)
	return whiskOutput{
		Value: output,
	}, err
}

{{ .AnzerIn}}

{{ .AnzerOut}}

`))
