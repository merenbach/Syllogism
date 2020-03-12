package symbol

import "fmt"

const (
	articleBlankString    = ""
	articleAString        = "a "
	articleAStringnString = "an "
	articleSmString       = "sm "

	// UndeterminedTypeString is a placeholder for an undetermined type.
	UndeterminedTypeString = "undetermined type"

	// GeneralTermString is a placeholder for a general term.
	GeneralTermString = "general term"

	// DesignatorString is a placeholder for a designator.
	DesignatorString = "designator"
)

// A Symbol is a logical symbol.
type Symbol struct {
	Term              string
	ArticleType       int
	TermType          int
	Occurrences       int
	DistributionCount int
}

// Empty determines whether a symbol is empty.
func (s *Symbol) Empty() bool {
	return s.Occurrences == 0
}

func (s *Symbol) ArticleTypeString() string {
	a := []string{
		articleBlankString,
		articleAString,
		articleAStringnString,
		articleSmString,
	}
	return a[s.ArticleType]
}

func (s *Symbol) TermTypeString() string {
	g := []string{
		UndeterminedTypeString,
		GeneralTermString,
		DesignatorString,
	}
	return g[s.TermType]
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
