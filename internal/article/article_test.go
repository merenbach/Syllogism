package article

import (
	"fmt"
	"testing"
)

func TestTypeFromString(t *testing.T) {
	tables := map[string]Type{
		"a":      TypeA,
		"an":     TypeAn,
		"sm":     TypeSm,
		"":       TypeNone,
		"foobar": TypeNone,
	}
	for k, v := range tables {
		out := TypeFromString(k)
		if v != out {
			t.Errorf("Expected %s => type %q, but got type %q instead\n", k, v, out)
		}
	}
}

func ExampleTypeNone() {
	fmt.Println(TypeNone)
	// Output:
}
func ExampleTypeA() {
	fmt.Println(TypeA)
	// Output: a
}

func ExampleTypeAn() {
	fmt.Println(TypeAn)
	// Output: an
}

func ExampleTypeSm() {
	fmt.Println(TypeSm)
	// Output: sm
}
