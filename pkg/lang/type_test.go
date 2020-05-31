package lang

import (
	"testing"
)

func TestBasicNeq(t *testing.T) {
	t1 := TypeInteger
	t2 := TypeString
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestBasicEq(t *testing.T) {
	t1 := TypeInteger
	t2 := TypeInteger
	if !t1.Equal(t2) {
		t.Error("t1 == t2")
	}
}

func TestBasicSubtype(t *testing.T) {
	t1 := TypeInteger
	t2 := TypeString
	if t1.Subtype(t2) {
		t.Error("!(t2 <: t1)")
	}
}

func TestComplexEq(t *testing.T) {
	t1 := Record{
		Fields: map[string]T{
			"a": TypeInteger,
		},
	}
	t2 := Record{
		Fields: map[string]T{
			"a": TypeInteger,
		},
	}
	if !t1.Equal(t2) {
		t.Error("t1 == t2")
	}
}

func TestComplexNeq1(t *testing.T) {
	t1 := Record{
		Fields: map[string]T{
			"a": TypeInteger,
		},
	}
	t2 := Record{
		Fields: map[string]T{
			"b": TypeInteger,
		},
	}
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestComplexNeq2(t *testing.T) {
	t1 := Record{
		Fields: map[string]T{
			"a": TypeInteger,
		},
	}
	t2 := Record{
		Fields: map[string]T{
			"a": TypeInteger,
			"b": TypeInteger,
		},
	}
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestComplexParent(t *testing.T) {
	t1 := Record{
		Fields: map[string]T{
			"a": TypeInteger,
		},
	}
	t2 := Record{
		Fields: map[string]T{
			"a": TypeInteger,
			"b": TypeInteger,
		},
	}
	if t1.Parent(t2) {
		t.Error("t1 <: t2")
	}
}

func TestComplexSubtype(t *testing.T) {
	t1 := Record{
		Fields: map[string]T{
			"a": TypeString,
		},
	}
	t2 := Record{
		Fields: map[string]T{
			"a": MaxLength(TypeString, 10),
			"b": TypeInteger,
		},
	}
	if t2.Subtype(t1) {
		t.Error("t1 <: t2")
	}
}

func TestConstructorEq1(t *testing.T) {
	t1 := MaxLength(TypeString, 10)
	t2 := MaxLength(TypeString, 10)
	if !t1.Equal(t2) {
		t.Error("t1 == t2")
	}
}

func TestConstructorEq2(t *testing.T) {
	t1 := MaxLength(MinLength(TypeString, 2), 10)
	t2 := MaxLength(MinLength(TypeString, 2), 10)
	if !t1.Equal(t2) {
		t.Error("t1 == t2")
	}
}

func TestConstructorNeq1(t *testing.T) {
	t1 := MaxLength(TypeString, 10)
	t2 := MaxLength(TypeString, 2)
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestConstructorNeq2(t *testing.T) {
	t1 := MaxLength(MinLength(TypeString, 2), 10)
	t2 := MaxLength(MinLength(TypeString, 2), 2)
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestConstructorNeq3(t *testing.T) {
	t1 := MaxLength(MinLength(TypeString, 2), 10)
	t2 := MaxLength(Optional(TypeInteger), 10)
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestConstructorParent(t *testing.T) {
	t1 := TypeString
	t2 := MaxLength(TypeString, 10)
	if !t1.Parent(t2) {
		t.Error("t1 <: t2")
	}
}

func TestConstructorNotparent(t *testing.T) {
	t1 := MaxLength(TypeString, 10)
	t2 := TypeString
	if t1.Parent(t2) {
		t.Error("!(t1 <: t2)")
	}
}

func TestMinLength(t *testing.T) {
	t1 := Contain([]T{TypeString}, TypeMinLength, []interface{}{10})
	t2 := MinLength(TypeString, 10)
	if !t1.Equal(t2) {
		t.Error("t1 == t2")
	}
}

func TestMinLengthIncorrect(t *testing.T) {
	t1 := TypeString
	t2 := MinLength(TypeString, 10)
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestMaxStringOfConstructor(t *testing.T) {
	embedded := MinLength(TypeString, 2)
	t1 := MaxLength(embedded, 10)
	t2 := Contain([]T{MinLength(TypeString, 2)}, TypeMaxLength, []interface{}{10})
	if !t1.Equal(t2) {
		t.Error("t1 == t2")
	}
}

func TestMaxStringOfConstructorIncorrect(t *testing.T) {
	t1 := MaxLength(Right(TypeInteger), 10)
	if t1 != nil {
		t.Error("t1 == nil")
	}
}

func TestEitherSubtype(t *testing.T) {
	left1 := MaxLength(TypeString, 10)
	right1 := List(TypeInteger)

	left2 := TypeString
	right2 := List(TypeInteger)

	t1 := Either(left1, right1)
	t2 := Either(left2, right2)

	if !t1.Subtype(t2) {
		t.Error("t2 <: t2")
	}
}

func TestTypeSumEqual1(t *testing.T) {
	t1 := NewSum(TypeString, TypeInteger)
	t2 := NewSum(TypeInteger, TypeString)
	if !t1.Equal(t2) {
		t.Error("t1 == t2")
	}
}

func TestTypeSumEqual2(t *testing.T) {
	t1 := NewSum(TypeString, TypeInteger)
	t2 := NewSum(TypeInteger)
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestTypeSumSubtype1(t *testing.T) {
	t1 := NewSum(TypeString, TypeInteger)
	t2 := NewSum(TypeInteger)
	if !t1.Parent(t2) {
		t.Error("t1 <: t2")
	}
}

func TestTypeSumSubtype2(t *testing.T) {
	t1 := NewSum(TypeInteger)
	t2 := NewSum(TypeString, TypeInteger)
	if t1.Parent(t2) {
		t.Error("!(t1 <: t2)")
	}
}

func TestTypeSumSubtype3(t *testing.T) {
	t1 := NewSum(TypeInteger)
	t2 := NewSum(TypeString, TypeInteger)
	if t2.Subtype(t1) {
		t.Error("!(t1 <: t2)")
	}
}

func TestTypeSumSubtype4(t *testing.T) {
	t1 := NewSum(TypeString, TypeInteger)
	t2 := NewSum(MaxLength(TypeString, 10))
	if !t2.Subtype(t1) {
		t.Error("t1 <: t2")
	}
}

func TestTypeSumSubtype5(t *testing.T) {
	t1 := NewSum(TypeString, TypeInteger)
	t2 := TypeString
	if !t2.Subtype(t1) {
		t.Error("t1 <: t2")
	}
}
