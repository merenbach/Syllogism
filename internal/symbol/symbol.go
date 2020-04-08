package symbol

import (
	"fmt"

	"github.com/merenbach/syllogism/internal/article"
	"github.com/merenbach/syllogism/internal/term"
)

// A Symbol is a logical symbol.
type Symbol struct {
	Term        string
	ArticleType article.Type
	TermType    term.Type
}

func (s *Symbol) String() string {
	return fmt.Sprintf("<%s %q>", s.TermType, s.Term)
}

// MatchesWordAndTermType determines if this symbol has a given term type (or is undetermined).
func (s *Symbol) MatchesWordAndTermType(w string, tt term.Type) bool {
	if w == s.Term {
		switch s.TermType {
		case term.TypeUndetermined:
			fmt.Printf("Note: %q used in premises taken to be %s\n", s.Term, tt)
			return true
		case tt:
			return true
		}
	}
	return false
}

// Dump values of variables in a Symbol.
// TODO: can we improve alignment?
func (s *Symbol) Dump() string {
	return fmt.Sprintf("%s\t%s\t%d",
		s.ArticleType,
		s.Term,
		s.TermType)
}

// ConclusionForAllIs returns a conclusion for "all X is Y"
// TODO: update these comments
func (s *Symbol) ConclusionForAllIs(o *Symbol) string {
	if s.TermType == term.TypeDesignator {
		return fmt.Sprintf("%s is %s%s", s.Term, o.ArticleType, o.Term)
	}

	return fmt.Sprintf("All %s is %s", s.Term, o.Term)
}

// ConclusionForSomeIs returns a conclusion for "some X is Y"
// TODO: update these comments
func (s *Symbol) ConclusionForSomeIs(o *Symbol) string {
	return fmt.Sprintf("Some %s is %s%s", s.Term, o.ArticleType, o.Term)
}
