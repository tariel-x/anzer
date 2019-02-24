package generator

import (
	"bytes"
	"errors"
	"strings"
	"text/template"
	"time"

	j "github.com/dave/jennifer/jen"
	in "github.com/tariel-x/anzer/internal"
)

var (
	errInvalidPackage = errors.New("invalid package")
)

func Generate(inT, outT in.T, packagePath string) (string, error) {
	packageElements := strings.Split(packagePath, "/")
	if len(packageElements) == 0 {
		return "", errInvalidPackage
	}

	var result bytes.Buffer
	err := execTemplate.Execute(&result, struct {
		Timestamp   time.Time
		AnzerIn     string
		AnzerOut    string
		PackagePath string
		Package     string
	}{
		Timestamp:   time.Now(),
		AnzerIn:     genType(inT, "AnzerIn"),
		AnzerOut:    genType(outT, "AnzerOut"),
		PackagePath: packagePath,
		Package:     packageElements[len(packageElements)-1],
	})
	return result.String(), err
}

func genType(t in.T, name string) string {
	return j.Type().Id(name).Add(genTypeDef(t)).GoString()
}

func genTypeDef(t in.T) *j.Statement {
	switch tt := t.(type) {
	case in.Constructor:
		switch tt.Type {
		case in.TypeList:
			return list(genTypeDef(tt.Operand))
		case in.TypeOptional:
			return pointer(genTypeDef(tt.Operand))
		default:
			return genTypeDef(tt.Operand)
		}
	case in.Basic:
		return basic(tt)
	case in.Complex:
		return complex(tt)
	case in.AnyType:
		return j.Interface()
	default:
		return j.Interface()
	}
}

func basic(tt in.Basic) *j.Statement {
	switch tt {
	case in.TypeString:
		return j.String()
	case in.TypeInteger:
		return j.Int()
	case in.TypeBool:
		return j.Bool()
	default:
		return j.Interface()
	}
}

func complex(tt in.Complex) *j.Statement {
	gofields := make([]j.Code, 0, len(tt.Fields))
	for fn, f := range tt.Fields {
		goftype := genTypeDef(f).Tag(map[string]string{"json": fn})
		gof := j.Id(strings.Title(fn)).Add(goftype)
		gofields = append(gofields, gof)
	}
	return j.Struct(gofields...)
}

func pointer(st *j.Statement) *j.Statement {
	return j.Op("*").Add(st)
}

func list(st *j.Statement) *j.Statement {
	return j.Index().Add(st)
}

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
	// debugging
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
	anzHandlerOutput := {{ .Package }}.Handle(typeIn(anzHandlerInput))
	handlerOutput := AnzerOut(anzHandlerOutput)
	output, err := json.Marshal(handlerOutput)
	return whiskOutput{
		Value: output,
	}, err
}

{{ .AnzerIn}}

{{ .AnzerOut}}

`))
