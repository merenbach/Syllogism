package symbol

import (
	"fmt"

	"github.com/merenbach/syllogism/internal/term"
)

const (
	articleBlankString = ""
	articleAString     = "a "
	articleAnString    = "an "
	articleSmString    = "sm "
)

// A Symbol is a logical symbol.
type Symbol struct {
	Term              string
	ArticleType       int
	TermType          term.Type
	Occurrences       int
	DistributionCount int
}

// Empty determines whether a symbol is empty.
func (s *Symbol) Empty() bool {
	return s.Occurrences == 0
}

// ArticleTypeString returns the article type (blank, A, An, Sm) for the symbol.
func (s *Symbol) ArticleTypeString() string {
	a := []string{
		articleBlankString,
		articleAString,
		articleAnString,
		articleSmString,
	}
	return a[s.ArticleType]
}

// Dump values of variables in a Symbol.
func (s *Symbol) Dump() string {
	return fmt.Sprintf("%s\t%s\t%d\t%d\t%d",
		s.ArticleTypeString(),
		s.Term,
		s.TermType,
		s.Occurrences,
		s.DistributionCount)
}
