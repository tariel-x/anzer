package generator

import (
	"bytes"
	"strings"
	"text/template"
	"time"

	in "github.com/tariel-x/anzer/internal"
)

func GenerateFunc(inT, outT in.T, packagePath string) (string, error) {
	packageElements := strings.Split(packagePath, "/")
	if len(packageElements) == 0 {
		return "", errInvalidPackage
	}
	packageName := packageElements[len(packageElements)-1]

	var result bytes.Buffer
	err := funcTemplate.Execute(&result, struct {
		Timestamp time.Time
		TypeIn    string
		TypeOut   string
		Package   string
	}{
		Timestamp: time.Now(),
		TypeIn:    genType(inT, "TypeIn"),
		TypeOut:   genType(outT, "TypeOut"),
		Package:   packageName,
	})
	return result.String(), err
}

var funcTemplate = template.Must(template.New("").Parse(`
// This file was generated by robots for you at {{ .Timestamp }}
package {{ .Package }}

{{ .TypeIn}}

{{ .TypeOut}}

func handle(input TypeIn) TypeOut {
	return TypeOut{}
}
`))
