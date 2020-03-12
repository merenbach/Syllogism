package stringutil

import "testing"

func TestHasAnyPrefix(t *testing.T) {
	dataTrue := [][]string{
		{"corgies", "cor", "foo"},
		{"corgies", "cor", ""},
		{"kitties", "kit"},
		{"kitties", ""},
		{"", ""},
	}

	dataFalse := [][]string{
		{"corgies", "foo", "bar", "baz"},
		{"kitties", "foo"},
		{"", "baz"},
	}

	for _, d := range dataTrue {
		if !hasAnyPrefix(d[0], d[1:]...) {
			t.Errorf("Expected %q to have prefix in: %s", d[0], d[1:])
		}
	}

	for _, d := range dataFalse {
		if hasAnyPrefix(d[0], d[1:]...) {
			t.Errorf("Expected %q to have no prefix in: %s", d[0], d[1:])
		}
	}
}

func TestHasAnySuffix(t *testing.T) {
	dataTrue := [][]string{
		{"corgies", "ies", "foo"},
		{"corgies", "foo", ""},
		{"kitties", "ies"},
		{"kitties", ""},
		{"", ""},
	}

	dataFalse := [][]string{
		{"corgies", "foo", "bar", "baz"},
		{"kitties", "foo"},
		{"", "baz"},
	}

	for _, d := range dataTrue {
		if !hasAnySuffix(d[0], d[1:]...) {
			t.Errorf("Expected %q to have suffix in: %s", d[0], d[1:])
		}
	}

	for _, d := range dataFalse {
		if hasAnySuffix(d[0], d[1:]...) {
			t.Errorf("Expected %q to have no suffix in: %s", d[0], d[1:])
		}
	}
}

func TestSingularize(t *testing.T) {
	data := map[string]string{
		"arches":        "arch",
		"class":         "class",
		"colossus":      "colossus",
		"dishes":        "dish",
		"foxes":         "fox",
		"friend's":      "friend's",
		"friends":       "friend",
		"iris":          "iris",
		"oxen":          "ox",
		"passes":        "pass",
		"pastries":      "pastry",
		"people":        "person",
		"pies":          "pie",
		"postmen":       "postman",
		"seraphim":      "seraph",
		"syllogism":     "syllogism",
		"the employees": "the employees",
	}

	for k, v := range data {
		if out := Singularize(k); out != v {
			t.Errorf("Expected %q to singularize to %q; instead got %q", k, v, out)
		}
	}
}
