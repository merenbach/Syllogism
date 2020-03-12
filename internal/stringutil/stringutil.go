package stringutil

import "strings"

// Plurals for singularization.
var plurals = map[string]string{
	"socrates":    "socrates",
	"parmenides":  "parmenides",
	"epimenides":  "epimenides",
	"mice":        "mouse",
	"lice":        "louse",
	"geese":       "goose",
	"children":    "child",
	"oxen":        "ox",
	"people":      "person",
	"teeth":       "tooth",
	"wolves":      "wolf",
	"wives":       "wife",
	"selves":      "self",
	"lives":       "life",
	"leaves":      "leaf",
	"shelves":     "shelf",
	"elves":       "elf",
	"dwarves":     "dwarf",
	"knives":      "knife",
	"thieves":     "thief",
	"neckties":    "necktie",
	"hippies":     "hippie",
	"yippies":     "yippie",
	"yuppies":     "yuppie",
	"moonies":     "moonie",
	"druggies":    "druggie",
	"cookies":     "cookie",
	"commies":     "commie",
	"groupies":    "groupie",
	"tomatoes":    "tomato",
	"alcibiades":  "alcibiades",
	"thales":      "thales",
	"aries":       "aries",
	"athens":      "athens",
	"species":     "species",
	"feces":       "feces",
	"geniuses":    "genius",
	"sorites":     "sorites",
	"crises":      "crisis",
	"emphases":    "emphasis",
	"memoranda":   "memorandum",
	"theses":      "thesis",
	"automata":    "automaton",
	"formulae":    "formula",
	"stigmata":    "stigma",
	"lemmata":     "lemma",
	"vertices":    "vertex",
	"vortices":    "vortex",
	"indices":     "index",
	"codices":     "codex",
	"matrices":    "matrix",
	"gasses":      "gas",
	"gases":       "gas",
	"buses":       "bus",
	"aches":       "ache",
	"headaches":   "headache",
	"grits":       "grits",
	"molasses":    "molasses",
	"gas":         "gas",
	"christmas":   "christmas",
	"mathematics": "mathematics",
	"semantics":   "semantics",
	"physics":     "physics",
	"metaphysics": "metaphysics",
	"ethics":      "ethics",
	"linguistics": "linguistics",
	"kiwis":       "kiwi",
	"israelis":    "israeli",
	"goyim":       "goy",
	"seraphim":    "seraph",
	"cherubim":    "cherub",
	"semen":       "semen",
	"amen":        "amen",
}

// HasAnyPrefix saves less-maintainable repeated calls to strings.HasPrefix()
func hasAnyPrefix(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// HasAnySuffix saves less-maintainable repeated calls to strings.HasSuffix()
func hasAnySuffix(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// HasPrefixVowel returns `true` if a string starts with a vowel, `false` otherwise.
func HasPrefixVowel(s string) bool {
	return hasAnyPrefix(s, "a", "e", "i", "o", "u")
}

// Singularize tries to convert a plural term to singular.
// Porting notes: All variable use is encapsulated, even in the original, so if porting needs to be re-done in future, re-porting this function can be avoided by invoking the equivalent of `w$ = singularize(w$)`.
func Singularize(term string) string {
	// 4040
	//---Convert W$ to singular---
	if strings.HasPrefix(term, "the ") {
		return term
	}

	words := make([]string, 0)
	for _, word := range strings.Fields(term) {
		if singular, found := plurals[word]; found {
			word = singular

		} else if strings.HasSuffix(word, "men") {
			word = strings.TrimSuffix(word, "en") + "an"

		} else if strings.HasSuffix(word, "s") && len(word) > 1 && !hasAnySuffix(word, "ss", "us", "is", "'s") {
			if hasAnySuffix(word, "xes", "sses", "shes", "ches") {
				// foxes => fox
				// passes => pass
				// dishes => dish
				// arches => arch
				word = strings.TrimSuffix(word, "es")

			} else if strings.HasSuffix(word, "ies") && len(word) > 4 {
				// pastries => pastry
				word = strings.TrimSuffix(word, "ies") + "y"

			} else {
				// cats => cat
				word = strings.TrimSuffix(word, "s")

			}
		}

		words = append(words, word)
	}

	return strings.Join(words, " ")
}
