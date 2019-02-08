package internal

import "testing"

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
	t1 := Complex{
		Fields: map[string]T{
			"a": TypeInteger,
		},
	}
	t2 := Complex{
		Fields: map[string]T{
			"a": TypeInteger,
		},
	}
	if !t1.Equal(t2) {
		t.Error("t1 == t2")
	}
}

func TestComplexNeq1(t *testing.T) {
	t1 := Complex{
		Fields: map[string]T{
			"a": TypeInteger,
		},
	}
	t2 := Complex{
		Fields: map[string]T{
			"b": TypeInteger,
		},
	}
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestComplexNeq2(t *testing.T) {
	t1 := Complex{
		Fields: map[string]T{
			"a": TypeInteger,
		},
	}
	t2 := Complex{
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
	t1 := Complex{
		Fields: map[string]T{
			"a": TypeInteger,
		},
	}
	t2 := Complex{
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
	t1 := Complex{
		Fields: map[string]T{
			"a": TypeString,
		},
	}
	t2 := Complex{
		Fields: map[string]T{
			"a": Construct(TypeString, TypeMaxLength, []interface{}{10}),
			"b": TypeInteger,
		},
	}
	if t2.Subtype(t1) {
		t.Error("t1 <: t2")
	}
}

func TestConstructorEq1(t *testing.T) {
	t1 := Construct(TypeString, TypeMaxLength, []interface{}{10})
	t2 := Construct(TypeString, TypeMaxLength, []interface{}{10})
	if !t1.Equal(t2) {
		t.Error("t1 == t2")
	}
}

func TestConstructorEq2(t *testing.T) {
	t1 := Construct(Construct(TypeString, TypeMaxLength, []interface{}{10}), TypeMinLength, []interface{}{2})
	t2 := Construct(Construct(TypeString, TypeMaxLength, []interface{}{10}), TypeMinLength, []interface{}{2})
	if !t1.Equal(t2) {
		t.Error("t1 == t2")
	}
}

func TestConstructorNeq1(t *testing.T) {
	t1 := Construct(TypeString, TypeMaxLength, []interface{}{10})
	t2 := Construct(TypeString, TypeMaxLength, []interface{}{10})
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestConstructorNeq2(t *testing.T) {
	t1 := Construct(Construct(TypeString, TypeMaxLength, []interface{}{10}), TypeMinLength, []interface{}{2})
	t2 := Construct(Construct(TypeString, TypeMaxLength, []interface{}{10}), TypeMinLength, []interface{}{2})
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestConstructorNeq3(t *testing.T) {
	t1 := Construct(Construct(TypeString, TypeMaxLength, []interface{}{10}), TypeMinLength, []interface{}{2})
	t2 := Construct(Construct(TypeInteger, TypeMaxLength, []interface{}{10}), TypeMinLength, []interface{}{2})
	if t1.Equal(t2) {
		t.Error("t1 != t2")
	}
}

func TestConstructorParent(t *testing.T) {
	t1 := TypeString
	t2 := Construct(TypeString, TypeMaxLength, []interface{}{10})
	if !t1.Parent(t2) {
		t.Error("t1 <: t2")
	}
}

func TestConstructorNotparent(t *testing.T) {
	t1 := Construct(TypeString, TypeMaxLength, []interface{}{10})
	t2 := TypeString
	if t1.Parent(t2) {
		t.Error("!(t1 <: t2)")
	}
}
