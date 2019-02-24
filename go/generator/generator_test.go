package generator

import (
	"testing"

	in "github.com/tariel-x/anzer/internal"
)

func TestGenTypeBasic1(t *testing.T) {
	inT := in.Optional(in.TypeInteger)
	goT := genType(inT, "inT")
	expected := `type inT *int`
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

func TestGenTypeBasic2(t *testing.T) {
	inT := in.TypeString
	goT := genType(inT, "inT")
	expected := `type inT string`
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

func TestGenTypeList(t *testing.T) {
	inT := in.List(in.TypeString)
	goT := genType(inT, "inT")
	expected := `type inT []string`
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

func TestGenTypeOptionalList(t *testing.T) {
	inT := in.List(in.Optional(in.TypeString))
	goT := genType(inT, "inT")
	expected := `type inT []*string`
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

func TestGenTypeListOptional(t *testing.T) {
	inT := in.Optional(in.List(in.TypeString))
	goT := genType(inT, "inT")
	expected := `type inT *[]string`
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

func TestGenTypeStruct(t *testing.T) {
	inT := in.Complex{
		Fields: map[string]in.T{
			"f1": in.Optional(in.TypeInteger),
			"f2": in.List(in.TypeString),
		},
	}
	goT := genType(inT, "inT")
	expected := "type inT struct {\n\tF1 *int     `json:\"f1\"`\n\tF2 []string `json:\"f2\"`\n}"
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}

func TestGenTypeStructListStruct(t *testing.T) {
	inT := in.Complex{
		Fields: map[string]in.T{
			"f1": in.Optional(in.TypeInteger),
			"f2": in.List(in.Complex{
				Fields: map[string]in.T{
					"subfield": in.Optional(in.TypeInteger),
				},
			}),
		},
	}
	goT := genType(inT, "inT")
	expected := "type inT struct {\n\tF1 *int `json:\"f1\"`\n\tF2 []struct {\n\t\tSubfield *int `json:\"subfield\"`\n\t} `json:\"f2\"`\n}"
	if goT != expected {
		t.Errorf("Definition %q != expected %q", goT, expected)
	}
}
