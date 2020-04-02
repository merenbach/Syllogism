package symbol

import (
	"fmt"

	"github.com/merenbach/syllogism/internal/article"
	"github.com/merenbach/syllogism/internal/term"
)

// A Symbol is a logical symbol.
type Symbol struct {
	Term              string
	ArticleType       article.Type
	TermType          term.Type
	Occurrences       int
	DistributionCount int
}

func (s *Symbol) String() string {
	return fmt.Sprintf("%s %q", s.TermType, s.Term)
}

// ReduceDistributionCount reduces the occurrences and (optionally) distribution count of the symbol.
// TODO: Golang may present a more idiomatic way to do this
func (s *Symbol) ReduceDistributionCount(decrement bool) {
	s.Occurrences--
	if s.Occurrences == 0 {
		s.Term = ""
		s.ArticleType = article.TypeNone
		s.TermType = term.TypeUndetermined
	}

	if decrement {
		s.DistributionCount--
	}
}

// Dump values of variables in a Symbol.
// TODO: can we improve alignment?
func (s *Symbol) Dump() string {
	return fmt.Sprintf("%s\t%s\t%d\t%d\t%d",
		s.ArticleType,
		s.Term,
		s.TermType,
		s.Occurrences,
		s.DistributionCount)
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
