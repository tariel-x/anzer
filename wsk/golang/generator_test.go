package golang

import (
	"testing"

	l "github.com/tariel-x/anzer/lang"
)

func TestGenTypeBasic1(t *testing.T) {
	inT := l.Optional(l.TypeInteger)
	goT := genType(inT, "inT")
	expected := `type inT *int`
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

func TestGenTypeBasic2(t *testing.T) {
	inT := l.TypeString
	goT := genType(inT, "inT")
	expected := `type inT string`
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

func TestGenTypeList(t *testing.T) {
	inT := l.List(l.TypeString)
	goT := genType(inT, "inT")
	expected := `type inT []string`
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

func TestGenTypeOptionalList(t *testing.T) {
	inT := l.List(l.Optional(l.TypeString))
	goT := genType(inT, "inT")
	expected := `type inT []*string`
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

func TestGenTypeListOptional(t *testing.T) {
	inT := l.Optional(l.List(l.TypeString))
	goT := genType(inT, "inT")
	expected := `type inT *[]string`
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

func TestGenTypeStruct(t *testing.T) {
	inT := l.Complex{
		Fields: map[string]l.T{
			"f1": l.Optional(l.TypeInteger),
			"f2": l.List(l.TypeString),
		},
	}
	goT := genType(inT, "inT")
	expected := "type inT struct {\n\tF1 *int     `json:\"f1\"`\n\tF2 []string `json:\"f2\"`\n}"
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

/*func TestGenTypeStructListStruct(t *testing.T) {
	inT := l.Complex{
		Fields: map[string]l.T{
			"f1": l.Optional(l.TypeInteger),
			"f2": l.List(l.Complex{
				Fields: map[string]l.T{
					"subfield": l.Optional(l.TypeInteger),
				},
			}),
		},
	}
	goT := genType(inT, "inT")
	expected := "type inT struct {\n\tF1 *int `json:\"f1\"`\n\tF2 []struct {\n\t\tSubfield *int `json:\"subfield\"`\n\t} `json:\"f2\"`\n}"
	if goT != expected {
		t.Errorf("Definition %s != expected %s", goT, expected)
	}
}*/
